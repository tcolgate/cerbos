// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package auxdata

import (
	"fmt"
	"time"

	"go.uber.org/multierr"
)

const (
	confKey = "auxData"
)

// Conf is optional configuration for Auxdata.
type Conf struct {
	// JWT holds the configuration for JWTs used as an auxiliary data source for the engine.
	JWT *JWTConf `yaml:"jwt"`
}

type JWTConf struct {
	// KeySets is the list of keysets to be used to verify tokens.
	KeySets []JWTKeySet `yaml:"keySets"`
	// DisableVerification disables JWT verification.
	DisableVerification bool `yaml:"disableVerification" conf:",example=false"`
	// CacheSize sets the number of verified tokens cached in memory. Set to negative value to disable caching.
	CacheSize int `yaml:"cacheSize" conf:",example=256"`
}

type JWTKeySet struct {
	// Remote defines a remote keyset. Mutually exclusive with Local.
	Remote *RemoteSource `yaml:"remote"`
	// Local defines a local keyset. Mutually exclusive with Remote.
	Local *LocalSource `yaml:"local"`
	// ID is the unique reference to this keyset.
	ID string `yaml:"id" conf:"required,example=ks1"`
}

type RemoteSource struct {
	// URL is the JWKS URL to fetch the keyset from.
	URL string `yaml:"url" conf:"required,example=https://domain.tld/.well-known/keys.jwks"`
	// RefreshInterval is the refresh interval for the keyset.
	RefreshInterval time.Duration `yaml:"refreshInterval" conf:",example=1h"`
}

type LocalSource struct {
	// Data is the encoded JWK data for this keyset. Mutually exclusive with File.
	Data string `yaml:"data" conf:",example=base64encodedJWK"`
	// File is the path to file containing JWK data. Mutually exclusive with Data.
	File string `yaml:"file" conf:",example=/path/to/keys.jwk"`
	// PEM indicates that the data is PEM encoded.
	PEM bool `yaml:"pem" conf:",example=true"`
}

func (c *Conf) Key() string {
	return confKey
}

func (c *Conf) Validate() (errs error) {
	if c.JWT == nil {
		return nil
	}

	if c.JWT.CacheSize == 0 {
		c.JWT.CacheSize = defaultCacheSize
	}

	idSet := make(map[string]struct{}, len(c.JWT.KeySets))
	for _, ks := range c.JWT.KeySets {
		if _, ok := idSet[ks.ID]; ok {
			errs = multierr.Append(errs, fmt.Errorf("duplicate keyset id '%s'", ks.ID))
			continue
		}

		idSet[ks.ID] = struct{}{}

		if ks.Remote == nil && ks.Local == nil {
			errs = multierr.Append(errs, fmt.Errorf("keyset '%s': should have one of `local` or `remote` defined", ks.ID))
			continue
		}

		if ks.Remote != nil && ks.Local != nil {
			errs = multierr.Append(errs, fmt.Errorf("keyset '%s': only one of `local` or `remote` should be defined", ks.ID))
			continue
		}

		if ks.Remote != nil && ks.Remote.URL == "" {
			errs = multierr.Append(errs, fmt.Errorf("keyset '%s': remote URL is empty", ks.ID))
			continue
		}

		if l := ks.Local; l != nil {
			if l.Data == "" && l.File == "" {
				errs = multierr.Append(errs, fmt.Errorf("keyset '%s': at least one of 'local.data' or 'local.file' must be defined", ks.ID))
				continue
			}

			if l.Data != "" && l.File != "" {
				errs = multierr.Append(errs, fmt.Errorf("keyset '%s': only one of 'loca.data' or 'local.file' must be defined", ks.ID))
			}
		}
	}

	return errs
}
