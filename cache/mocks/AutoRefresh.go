// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import cache "github.com/lyft/flytestdlib/cache"
import context "context"
import mock "github.com/stretchr/testify/mock"

// AutoRefresh is an autogenerated mock type for the AutoRefresh type
type AutoRefresh struct {
	mock.Mock
}

type AutoRefresh_Get struct {
	*mock.Call
}

func (_m AutoRefresh_Get) Return(_a0 cache.Item, _a1 error) *AutoRefresh_Get {
	return &AutoRefresh_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AutoRefresh) OnGet(id string) *AutoRefresh_Get {
	c := _m.On("Get", id)
	return &AutoRefresh_Get{Call: c}
}

func (_m *AutoRefresh) OnGetMatch(matchers ...interface{}) *AutoRefresh_Get {
	c := _m.On("Get", matchers...)
	return &AutoRefresh_Get{Call: c}
}

// Get provides a mock function with given fields: id
func (_m *AutoRefresh) Get(id string) (cache.Item, error) {
	ret := _m.Called(id)

	var r0 cache.Item
	if rf, ok := ret.Get(0).(func(string) cache.Item); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AutoRefresh_GetOrCreate struct {
	*mock.Call
}

func (_m AutoRefresh_GetOrCreate) Return(_a0 cache.Item, _a1 error) *AutoRefresh_GetOrCreate {
	return &AutoRefresh_GetOrCreate{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AutoRefresh) OnGetOrCreate(id string, item cache.Item) *AutoRefresh_GetOrCreate {
	c := _m.On("GetOrCreate", id, item)
	return &AutoRefresh_GetOrCreate{Call: c}
}

func (_m *AutoRefresh) OnGetOrCreateMatch(matchers ...interface{}) *AutoRefresh_GetOrCreate {
	c := _m.On("GetOrCreate", matchers...)
	return &AutoRefresh_GetOrCreate{Call: c}
}

// GetOrCreate provides a mock function with given fields: id, item
func (_m *AutoRefresh) GetOrCreate(id string, item cache.Item) (cache.Item, error) {
	ret := _m.Called(id, item)

	var r0 cache.Item
	if rf, ok := ret.Get(0).(func(string, cache.Item) cache.Item); ok {
		r0 = rf(id, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, cache.Item) error); ok {
		r1 = rf(id, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AutoRefresh_Start struct {
	*mock.Call
}

func (_m AutoRefresh_Start) Return(_a0 error) *AutoRefresh_Start {
	return &AutoRefresh_Start{Call: _m.Call.Return(_a0)}
}

func (_m *AutoRefresh) OnStart(ctx context.Context) *AutoRefresh_Start {
	c := _m.On("Start", ctx)
	return &AutoRefresh_Start{Call: c}
}

func (_m *AutoRefresh) OnStartMatch(matchers ...interface{}) *AutoRefresh_Start {
	c := _m.On("Start", matchers...)
	return &AutoRefresh_Start{Call: c}
}

// Start provides a mock function with given fields: ctx
func (_m *AutoRefresh) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
