// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	ed25519 "github.com/ChainSafe/gossamer/lib/crypto/ed25519"
	"github.com/ChainSafe/gossamer/lib/grandpa"

	mock "github.com/stretchr/testify/mock"
)

// MockBlockFinalityAPI is an autogenerated mock type for the BlockFinalityAPI type
type MockBlockFinalityAPI struct {
	mock.Mock
}

// GetRound provides a mock function with given fields:
func (_m *MockBlockFinalityAPI) GetRound() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetSetID provides a mock function with given fields:
func (_m *MockBlockFinalityAPI) GetSetID() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetVoters provides a mock function with given fields:
func (_m *MockBlockFinalityAPI) GetVoters() grandpa.Voters {
	ret := _m.Called()

	var r0 grandpa.Voters
	if rf, ok := ret.Get(0).(func() grandpa.Voters); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grandpa.Voters)
		}
	}

	return r0
}

// PreCommits provides a mock function with given fields:
func (_m *MockBlockFinalityAPI) PreCommits() []ed25519.PublicKeyBytes {
	ret := _m.Called()

	var r0 []ed25519.PublicKeyBytes
	if rf, ok := ret.Get(0).(func() []ed25519.PublicKeyBytes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ed25519.PublicKeyBytes)
		}
	}

	return r0
}

// PreVotes provides a mock function with given fields:
func (_m *MockBlockFinalityAPI) PreVotes() []ed25519.PublicKeyBytes {
	ret := _m.Called()

	var r0 []ed25519.PublicKeyBytes
	if rf, ok := ret.Get(0).(func() []ed25519.PublicKeyBytes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ed25519.PublicKeyBytes)
		}
	}

	return r0
}
