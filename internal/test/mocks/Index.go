// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	index "github.com/cerbos/cerbos/internal/storage/index"

	mock "github.com/stretchr/testify/mock"

	namer "github.com/cerbos/cerbos/internal/namer"

	policy "github.com/cerbos/cerbos/internal/policy"

	storage "github.com/cerbos/cerbos/internal/storage"
)

// Index is an autogenerated mock type for the Index type
type Index struct {
	mock.Mock
}

type Index_Expecter struct {
	mock *mock.Mock
}

func (_m *Index) EXPECT() *Index_Expecter {
	return &Index_Expecter{mock: &_m.Mock}
}

// AddOrUpdate provides a mock function with given fields: _a0
func (_m *Index) AddOrUpdate(_a0 index.Entry) (storage.Event, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddOrUpdate")
	}

	var r0 storage.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(index.Entry) (storage.Event, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(index.Entry) storage.Event); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(storage.Event)
	}

	if rf, ok := ret.Get(1).(func(index.Entry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_AddOrUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddOrUpdate'
type Index_AddOrUpdate_Call struct {
	*mock.Call
}

// AddOrUpdate is a helper method to define mock.On call
//   - _a0 index.Entry
func (_e *Index_Expecter) AddOrUpdate(_a0 interface{}) *Index_AddOrUpdate_Call {
	return &Index_AddOrUpdate_Call{Call: _e.mock.On("AddOrUpdate", _a0)}
}

func (_c *Index_AddOrUpdate_Call) Run(run func(_a0 index.Entry)) *Index_AddOrUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(index.Entry))
	})
	return _c
}

func (_c *Index_AddOrUpdate_Call) Return(_a0 storage.Event, _a1 error) *Index_AddOrUpdate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_AddOrUpdate_Call) RunAndReturn(run func(index.Entry) (storage.Event, error)) *Index_AddOrUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// Clear provides a mock function with given fields:
func (_m *Index) Clear() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Index_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type Index_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
func (_e *Index_Expecter) Clear() *Index_Clear_Call {
	return &Index_Clear_Call{Call: _e.mock.On("Clear")}
}

func (_c *Index_Clear_Call) Run(run func()) *Index_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Index_Clear_Call) Return(_a0 error) *Index_Clear_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Index_Clear_Call) RunAndReturn(run func() error) *Index_Clear_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *Index) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Index_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Index_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Index_Expecter) Close() *Index_Close_Call {
	return &Index_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Index_Close_Call) Run(run func()) *Index_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Index_Close_Call) Return(_a0 error) *Index_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Index_Close_Call) RunAndReturn(run func() error) *Index_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *Index) Delete(_a0 index.Entry) (storage.Event, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 storage.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(index.Entry) (storage.Event, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(index.Entry) storage.Event); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(storage.Event)
	}

	if rf, ok := ret.Get(1).(func(index.Entry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Index_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 index.Entry
func (_e *Index_Expecter) Delete(_a0 interface{}) *Index_Delete_Call {
	return &Index_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *Index_Delete_Call) Run(run func(_a0 index.Entry)) *Index_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(index.Entry))
	})
	return _c
}

func (_c *Index_Delete_Call) Return(_a0 storage.Event, _a1 error) *Index_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_Delete_Call) RunAndReturn(run func(index.Entry) (storage.Event, error)) *Index_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllCompilationUnits provides a mock function with given fields: _a0
func (_m *Index) GetAllCompilationUnits(_a0 context.Context) <-chan *policy.CompilationUnit {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCompilationUnits")
	}

	var r0 <-chan *policy.CompilationUnit
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *policy.CompilationUnit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *policy.CompilationUnit)
		}
	}

	return r0
}

// Index_GetAllCompilationUnits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllCompilationUnits'
type Index_GetAllCompilationUnits_Call struct {
	*mock.Call
}

// GetAllCompilationUnits is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Index_Expecter) GetAllCompilationUnits(_a0 interface{}) *Index_GetAllCompilationUnits_Call {
	return &Index_GetAllCompilationUnits_Call{Call: _e.mock.On("GetAllCompilationUnits", _a0)}
}

func (_c *Index_GetAllCompilationUnits_Call) Run(run func(_a0 context.Context)) *Index_GetAllCompilationUnits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Index_GetAllCompilationUnits_Call) Return(_a0 <-chan *policy.CompilationUnit) *Index_GetAllCompilationUnits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Index_GetAllCompilationUnits_Call) RunAndReturn(run func(context.Context) <-chan *policy.CompilationUnit) *Index_GetAllCompilationUnits_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompilationUnits provides a mock function with given fields: _a0
func (_m *Index) GetCompilationUnits(_a0 ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCompilationUnits")
	}

	var r0 map[namer.ModuleID]*policy.CompilationUnit
	var r1 error
	if rf, ok := ret.Get(0).(func(...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...namer.ModuleID) map[namer.ModuleID]*policy.CompilationUnit); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[namer.ModuleID]*policy.CompilationUnit)
		}
	}

	if rf, ok := ret.Get(1).(func(...namer.ModuleID) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_GetCompilationUnits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompilationUnits'
type Index_GetCompilationUnits_Call struct {
	*mock.Call
}

// GetCompilationUnits is a helper method to define mock.On call
//   - _a0 ...namer.ModuleID
func (_e *Index_Expecter) GetCompilationUnits(_a0 ...interface{}) *Index_GetCompilationUnits_Call {
	return &Index_GetCompilationUnits_Call{Call: _e.mock.On("GetCompilationUnits",
		append([]interface{}{}, _a0...)...)}
}

func (_c *Index_GetCompilationUnits_Call) Run(run func(_a0 ...namer.ModuleID)) *Index_GetCompilationUnits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]namer.ModuleID, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(namer.ModuleID)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Index_GetCompilationUnits_Call) Return(_a0 map[namer.ModuleID]*policy.CompilationUnit, _a1 error) *Index_GetCompilationUnits_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_GetCompilationUnits_Call) RunAndReturn(run func(...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error)) *Index_GetCompilationUnits_Call {
	_c.Call.Return(run)
	return _c
}

// GetDependents provides a mock function with given fields: _a0
func (_m *Index) GetDependents(_a0 ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDependents")
	}

	var r0 map[namer.ModuleID][]namer.ModuleID
	var r1 error
	if rf, ok := ret.Get(0).(func(...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...namer.ModuleID) map[namer.ModuleID][]namer.ModuleID); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[namer.ModuleID][]namer.ModuleID)
		}
	}

	if rf, ok := ret.Get(1).(func(...namer.ModuleID) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_GetDependents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDependents'
type Index_GetDependents_Call struct {
	*mock.Call
}

// GetDependents is a helper method to define mock.On call
//   - _a0 ...namer.ModuleID
func (_e *Index_Expecter) GetDependents(_a0 ...interface{}) *Index_GetDependents_Call {
	return &Index_GetDependents_Call{Call: _e.mock.On("GetDependents",
		append([]interface{}{}, _a0...)...)}
}

func (_c *Index_GetDependents_Call) Run(run func(_a0 ...namer.ModuleID)) *Index_GetDependents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]namer.ModuleID, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(namer.ModuleID)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Index_GetDependents_Call) Return(_a0 map[namer.ModuleID][]namer.ModuleID, _a1 error) *Index_GetDependents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_GetDependents_Call) RunAndReturn(run func(...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error)) *Index_GetDependents_Call {
	_c.Call.Return(run)
	return _c
}

// GetFiles provides a mock function with given fields:
func (_m *Index) GetFiles() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFiles")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Index_GetFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFiles'
type Index_GetFiles_Call struct {
	*mock.Call
}

// GetFiles is a helper method to define mock.On call
func (_e *Index_Expecter) GetFiles() *Index_GetFiles_Call {
	return &Index_GetFiles_Call{Call: _e.mock.On("GetFiles")}
}

func (_c *Index_GetFiles_Call) Run(run func()) *Index_GetFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Index_GetFiles_Call) Return(_a0 []string) *Index_GetFiles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Index_GetFiles_Call) RunAndReturn(run func() []string) *Index_GetFiles_Call {
	_c.Call.Return(run)
	return _c
}

// GetFirstMatch provides a mock function with given fields: _a0
func (_m *Index) GetFirstMatch(_a0 []namer.ModuleID) (*policy.CompilationUnit, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetFirstMatch")
	}

	var r0 *policy.CompilationUnit
	var r1 error
	if rf, ok := ret.Get(0).(func([]namer.ModuleID) (*policy.CompilationUnit, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]namer.ModuleID) *policy.CompilationUnit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*policy.CompilationUnit)
		}
	}

	if rf, ok := ret.Get(1).(func([]namer.ModuleID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_GetFirstMatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirstMatch'
type Index_GetFirstMatch_Call struct {
	*mock.Call
}

// GetFirstMatch is a helper method to define mock.On call
//   - _a0 []namer.ModuleID
func (_e *Index_Expecter) GetFirstMatch(_a0 interface{}) *Index_GetFirstMatch_Call {
	return &Index_GetFirstMatch_Call{Call: _e.mock.On("GetFirstMatch", _a0)}
}

func (_c *Index_GetFirstMatch_Call) Run(run func(_a0 []namer.ModuleID)) *Index_GetFirstMatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]namer.ModuleID))
	})
	return _c
}

func (_c *Index_GetFirstMatch_Call) Return(_a0 *policy.CompilationUnit, _a1 error) *Index_GetFirstMatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_GetFirstMatch_Call) RunAndReturn(run func([]namer.ModuleID) (*policy.CompilationUnit, error)) *Index_GetFirstMatch_Call {
	_c.Call.Return(run)
	return _c
}

// ListPolicyIDs provides a mock function with given fields: _a0
func (_m *Index) ListPolicyIDs(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ListPolicyIDs")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_ListPolicyIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPolicyIDs'
type Index_ListPolicyIDs_Call struct {
	*mock.Call
}

// ListPolicyIDs is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Index_Expecter) ListPolicyIDs(_a0 interface{}) *Index_ListPolicyIDs_Call {
	return &Index_ListPolicyIDs_Call{Call: _e.mock.On("ListPolicyIDs", _a0)}
}

func (_c *Index_ListPolicyIDs_Call) Run(run func(_a0 context.Context)) *Index_ListPolicyIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Index_ListPolicyIDs_Call) Return(_a0 []string, _a1 error) *Index_ListPolicyIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_ListPolicyIDs_Call) RunAndReturn(run func(context.Context) ([]string, error)) *Index_ListPolicyIDs_Call {
	_c.Call.Return(run)
	return _c
}

// ListSchemaIDs provides a mock function with given fields: _a0
func (_m *Index) ListSchemaIDs(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ListSchemaIDs")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_ListSchemaIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSchemaIDs'
type Index_ListSchemaIDs_Call struct {
	*mock.Call
}

// ListSchemaIDs is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Index_Expecter) ListSchemaIDs(_a0 interface{}) *Index_ListSchemaIDs_Call {
	return &Index_ListSchemaIDs_Call{Call: _e.mock.On("ListSchemaIDs", _a0)}
}

func (_c *Index_ListSchemaIDs_Call) Run(run func(_a0 context.Context)) *Index_ListSchemaIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Index_ListSchemaIDs_Call) Return(_a0 []string, _a1 error) *Index_ListSchemaIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_ListSchemaIDs_Call) RunAndReturn(run func(context.Context) ([]string, error)) *Index_ListSchemaIDs_Call {
	_c.Call.Return(run)
	return _c
}

// LoadPolicy provides a mock function with given fields: _a0, _a1
func (_m *Index) LoadPolicy(_a0 context.Context, _a1 ...string) ([]*policy.Wrapper, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for LoadPolicy")
	}

	var r0 []*policy.Wrapper
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) ([]*policy.Wrapper, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []*policy.Wrapper); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*policy.Wrapper)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_LoadPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadPolicy'
type Index_LoadPolicy_Call struct {
	*mock.Call
}

// LoadPolicy is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...string
func (_e *Index_Expecter) LoadPolicy(_a0 interface{}, _a1 ...interface{}) *Index_LoadPolicy_Call {
	return &Index_LoadPolicy_Call{Call: _e.mock.On("LoadPolicy",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *Index_LoadPolicy_Call) Run(run func(_a0 context.Context, _a1 ...string)) *Index_LoadPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *Index_LoadPolicy_Call) Return(_a0 []*policy.Wrapper, _a1 error) *Index_LoadPolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_LoadPolicy_Call) RunAndReturn(run func(context.Context, ...string) ([]*policy.Wrapper, error)) *Index_LoadPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// LoadSchema provides a mock function with given fields: _a0, _a1
func (_m *Index) LoadSchema(_a0 context.Context, _a1 string) (io.ReadCloser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LoadSchema")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (io.ReadCloser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_LoadSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadSchema'
type Index_LoadSchema_Call struct {
	*mock.Call
}

// LoadSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *Index_Expecter) LoadSchema(_a0 interface{}, _a1 interface{}) *Index_LoadSchema_Call {
	return &Index_LoadSchema_Call{Call: _e.mock.On("LoadSchema", _a0, _a1)}
}

func (_c *Index_LoadSchema_Call) Run(run func(_a0 context.Context, _a1 string)) *Index_LoadSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Index_LoadSchema_Call) Return(_a0 io.ReadCloser, _a1 error) *Index_LoadSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_LoadSchema_Call) RunAndReturn(run func(context.Context, string) (io.ReadCloser, error)) *Index_LoadSchema_Call {
	_c.Call.Return(run)
	return _c
}

// Reload provides a mock function with given fields: ctx
func (_m *Index) Reload(ctx context.Context) ([]storage.Event, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Reload")
	}

	var r0 []storage.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]storage.Event, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []storage.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index_Reload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reload'
type Index_Reload_Call struct {
	*mock.Call
}

// Reload is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Index_Expecter) Reload(ctx interface{}) *Index_Reload_Call {
	return &Index_Reload_Call{Call: _e.mock.On("Reload", ctx)}
}

func (_c *Index_Reload_Call) Run(run func(ctx context.Context)) *Index_Reload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Index_Reload_Call) Return(_a0 []storage.Event, _a1 error) *Index_Reload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Index_Reload_Call) RunAndReturn(run func(context.Context) ([]storage.Event, error)) *Index_Reload_Call {
	_c.Call.Return(run)
	return _c
}

// RepoStats provides a mock function with given fields: _a0
func (_m *Index) RepoStats(_a0 context.Context) storage.RepoStats {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RepoStats")
	}

	var r0 storage.RepoStats
	if rf, ok := ret.Get(0).(func(context.Context) storage.RepoStats); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(storage.RepoStats)
	}

	return r0
}

// Index_RepoStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RepoStats'
type Index_RepoStats_Call struct {
	*mock.Call
}

// RepoStats is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Index_Expecter) RepoStats(_a0 interface{}) *Index_RepoStats_Call {
	return &Index_RepoStats_Call{Call: _e.mock.On("RepoStats", _a0)}
}

func (_c *Index_RepoStats_Call) Run(run func(_a0 context.Context)) *Index_RepoStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Index_RepoStats_Call) Return(_a0 storage.RepoStats) *Index_RepoStats_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Index_RepoStats_Call) RunAndReturn(run func(context.Context) storage.RepoStats) *Index_RepoStats_Call {
	_c.Call.Return(run)
	return _c
}

// NewIndex creates a new instance of Index. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndex(t interface {
	mock.TestingT
	Cleanup(func())
}) *Index {
	mock := &Index{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
