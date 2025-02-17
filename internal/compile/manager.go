// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package compile

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bluele/gcache"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/observability/metrics"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage"
)

const (
	cacheKind             = "compile"
	negativeCacheEntryTTL = 10 * time.Second
	storeFetchTimeout     = 2 * time.Second
	updateQueueSize       = 32
)

type Manager struct {
	log           *zap.SugaredLogger
	store         storage.SourceStore
	schemaMgr     schema.Manager
	updateQueue   chan storage.Event
	cache         gcache.Cache
	sf            singleflight.Group
	cacheDuration time.Duration
}

func NewManager(ctx context.Context, store storage.SourceStore, schemaMgr schema.Manager) (*Manager, error) {
	conf := &Conf{}
	if err := config.GetSection(conf); err != nil {
		return nil, err
	}

	return NewManagerFromConf(ctx, conf, store, schemaMgr), nil
}

func NewManagerFromDefaultConf(ctx context.Context, store storage.SourceStore, schemaMgr schema.Manager) *Manager {
	return NewManagerFromConf(ctx, DefaultConf(), store, schemaMgr)
}

func NewManagerFromConf(ctx context.Context, conf *Conf, store storage.SourceStore, schemaMgr schema.Manager) *Manager {
	c := &Manager{
		log:           zap.S().Named("compiler"),
		store:         store,
		schemaMgr:     schemaMgr,
		updateQueue:   make(chan storage.Event, updateQueueSize),
		cache:         mkCache(int(conf.CacheSize)),
		cacheDuration: conf.CacheDuration,
	}

	go c.processUpdateQueue(ctx)
	store.Subscribe(c)

	return c
}

func (c *Manager) SubscriberID() string {
	return "compile.Manager"
}

func (c *Manager) OnStorageEvent(events ...storage.Event) {
	for _, evt := range events {
		c.log.Debugw("Received storage event", "event", evt)
		c.updateQueue <- evt
	}
}

func (c *Manager) processUpdateQueue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-c.updateQueue:
			c.log.Debugw("Processing storage event", "event", evt)
			switch evt.Kind {
			case storage.EventReload:
				c.log.Info("Purging compile cache")
				c.cache.Purge()
			case storage.EventAddOrUpdatePolicy, storage.EventDeletePolicy:
				if err := c.recompile(evt); err != nil {
					c.log.Warnw("Error while processing storage event", "event", evt, "error", err)
				}
			default:
				c.log.Debugw("Ignoring storage event", "event", evt)
			}
		}
	}
}

func (c *Manager) recompile(evt storage.Event) error {
	// if this is a delete event, remove the module from the cache
	if evt.Kind == storage.EventDeletePolicy {
		c.evict(evt.PolicyID)
	}

	// find the modules that will be affected by this policy getting updated or deleted.
	var toRecompile []namer.ModuleID
	if evt.Kind == storage.EventAddOrUpdatePolicy {
		toRecompile = append(toRecompile, evt.PolicyID)

		// if the policy ID has changed, remove the old cached entry
		if evt.OldPolicyID != nil {
			c.evict(*evt.OldPolicyID)
		}
	}

	dependents, err := c.getDependents(evt.PolicyID)
	if err != nil {
		return err
	}

	// only recompile the ones that are already cached.
	for _, d := range dependents {
		if c.cache.Has(d) {
			toRecompile = append(toRecompile, d)
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), storeFetchTimeout)
	defer cancelFunc()

	compileUnits, err := c.store.GetCompilationUnits(ctx, toRecompile...)
	if err != nil {
		return fmt.Errorf("failed to get compilation units: %w", err)
	}

	for modID, cu := range compileUnits {
		if cu.MainPolicy() == nil || cu.MainPolicy().Disabled {
			c.evict(cu.ModID)
			c.log.Debugw("Evicted the disabled policy", "id", cu.ModID.String())
			continue
		}
		if _, err := c.compile(cu); err != nil {
			// log and remove the module that failed to compile.
			c.log.Errorw("Failed to recompile", "id", modID, "error", err)
			c.evict(modID)
		}
	}

	return nil
}

func (c *Manager) getDependents(modID namer.ModuleID) ([]namer.ModuleID, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), storeFetchTimeout)
	defer cancelFunc()

	dependents, err := c.store.GetDependents(ctx, modID)
	if err != nil {
		return nil, fmt.Errorf("failed to find dependents: %w", err)
	}

	if len(dependents) > 0 {
		return dependents[modID], nil
	}

	return nil, nil
}

func (c *Manager) compile(unit *policy.CompilationUnit) (*runtimev1.RunnablePolicySet, error) {
	startTime := time.Now()
	rps, err := Compile(unit, c.schemaMgr)
	durationMs := float64(time.Since(startTime)) / float64(time.Millisecond)

	if err == nil && rps != nil {
		if c.cacheDuration > 0 {
			_ = c.cache.SetWithExpire(unit.ModID, rps, c.cacheDuration)
		} else {
			_ = c.cache.Set(unit.ModID, rps)
		}
	}

	status := "success"
	if err != nil {
		status = "failure"
	}

	_ = stats.RecordWithTags(
		context.Background(),
		[]tag.Mutator{tag.Upsert(metrics.KeyCompileStatus, status)},
		metrics.CompileDuration.M(durationMs),
	)

	return rps, err
}

func (c *Manager) evict(modID namer.ModuleID) {
	c.cache.Remove(modID)
}

func (c *Manager) GetPolicySet(ctx context.Context, modID namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	key := modID.String()
	defer c.sf.Forget(key)

	rpsVal, err, _ := c.sf.Do(key, func() (any, error) {
		rps, err := c.cache.GetIFPresent(modID)
		if err == nil {
			cacheHit()

			// If the value is nil, it indicates a negative cache entry (see below)
			// Essentially, we tried to find this evaluator before and it wasn't found.
			// We don't want to hit the store over and over again because we know it doesn't exist.
			if rps == nil {
				return nil, nil
			}

			return rps.(*runtimev1.RunnablePolicySet), nil //nolint:forcetypeassert
		}

		cacheMiss()

		compileUnits, err := c.store.GetCompilationUnits(ctx, modID)
		if err != nil {
			return nil, fmt.Errorf("failed to get compilation units: %w", err)
		}

		if len(compileUnits) == 0 {
			// store a nil value in the cache as a negative entry to prevent hitting the database again and again
			_ = c.cache.SetWithExpire(modID, nil, negativeCacheEntryTTL)
			return nil, nil
		}

		var retVal *runtimev1.RunnablePolicySet
		for mID, cu := range compileUnits {
			rps, err := c.compile(cu)
			if err != nil {
				return nil, PolicyCompilationErr{underlying: err}
			}

			if mID == modID {
				retVal = rps
			}
		}

		return retVal, nil
	})
	if err != nil {
		return nil, err
	}

	if rpsVal == nil {
		return nil, nil
	}

	//nolint:forcetypeassert
	return rpsVal.(*runtimev1.RunnablePolicySet), nil
}

func mkCache(size int) gcache.Cache {
	_ = stats.RecordWithTags(context.Background(),
		[]tag.Mutator{tag.Upsert(metrics.KeyCacheKind, cacheKind)},
		metrics.CacheMaxSize.M(int64(size)),
	)

	gauge := metrics.MakeCacheGauge(cacheKind)
	return gcache.New(size).
		ARC().
		AddedFunc(func(_, _ any) {
			gauge.Add(1)
		}).
		EvictedFunc(func(_, _ any) {
			gauge.Add(-1)
		}).Build()
}

func cacheHit() {
	_ = stats.RecordWithTags(context.Background(),
		[]tag.Mutator{tag.Upsert(metrics.KeyCacheKind, cacheKind), tag.Upsert(metrics.KeyCacheResult, "hit")},
		metrics.CacheAccessCount.M(1),
	)
}

func cacheMiss() {
	_ = stats.RecordWithTags(context.Background(),
		[]tag.Mutator{tag.Upsert(metrics.KeyCacheKind, cacheKind), tag.Upsert(metrics.KeyCacheResult, "miss")},
		metrics.CacheAccessCount.M(1),
	)
}

type PolicyCompilationErr struct {
	underlying error
}

func (pce PolicyCompilationErr) Error() string {
	return fmt.Sprintf("policy compilation error: %v", pce.underlying)
}

func (pce PolicyCompilationErr) Unwrap() error {
	return pce.underlying
}

func (pce PolicyCompilationErr) Is(target error) bool {
	return errors.As(target, &PolicyCompilationErr{})
}
