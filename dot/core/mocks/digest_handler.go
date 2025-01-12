// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	types "github.com/ChainSafe/gossamer/dot/types"
	mock "github.com/stretchr/testify/mock"
)

// MockDigestHandler is an autogenerated mock type for the DigestHandler type
type MockDigestHandler struct {
	mock.Mock
}

// HandleDigests provides a mock function with given fields: header
func (_m *MockDigestHandler) HandleDigests(header *types.Header) {
	_m.Called(header)
}
