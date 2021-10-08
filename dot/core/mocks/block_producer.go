// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	types "github.com/ChainSafe/gossamer/dot/types"
	mock "github.com/stretchr/testify/mock"
)

// MockBlockProducer is an autogenerated mock type for the BlockProducer type
type MockBlockProducer struct {
	mock.Mock
}

// GetBlockChannel provides a mock function with given fields:
func (_m *MockBlockProducer) GetBlockChannel() <-chan types.Block {
	ret := _m.Called()

	var r0 <-chan types.Block
	if rf, ok := ret.Get(0).(func() <-chan types.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan types.Block)
		}
	}

	return r0
}

// SetOnDisabled provides a mock function with given fields: authorityIndex
func (_m *MockBlockProducer) SetOnDisabled(authorityIndex uint32) {
	_m.Called(authorityIndex)
}