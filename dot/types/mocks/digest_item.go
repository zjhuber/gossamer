// Code generated by mockery v2.8.0. DO NOT EDIT.

package types

import (
	"bytes"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockDigestItem is an autogenerated mock type for the DigestItem type
type MockDigestItem struct {
	mock.Mock
}

// Decode provides a mock function with given fields: _a0
func (_m *MockDigestItem) Decode(_a0 *bytes.Buffer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Reader) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Encode provides a mock function with given fields:
func (_m *MockDigestItem) Encode() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// String provides a mock function with given fields:
func (_m *MockDigestItem) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *MockDigestItem) Type() byte {
	ret := _m.Called()

	var r0 byte
	if rf, ok := ret.Get(0).(func() byte); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(byte)
	}

	return r0
}
