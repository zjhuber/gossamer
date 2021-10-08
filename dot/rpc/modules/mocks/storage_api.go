// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	common "github.com/ChainSafe/gossamer/lib/common"
	mock "github.com/stretchr/testify/mock"

	state "github.com/ChainSafe/gossamer/dot/state"

	trie "github.com/ChainSafe/gossamer/lib/trie"
)

// MockStorageAPI is an autogenerated mock type for the StorageAPI type
type MockStorageAPI struct {
	mock.Mock
}

// Entries provides a mock function with given fields: root
func (_m *MockStorageAPI) Entries(root *common.Hash) (map[string][]byte, error) {
	ret := _m.Called(root)

	var r0 map[string][]byte
	if rf, ok := ret.Get(0).(func(*common.Hash) map[string][]byte); ok {
		r0 = rf(root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash) error); ok {
		r1 = rf(root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeysWithPrefix provides a mock function with given fields: root, prefix
func (_m *MockStorageAPI) GetKeysWithPrefix(root *common.Hash, prefix []byte) ([][]byte, error) {
	ret := _m.Called(root, prefix)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(*common.Hash, []byte) [][]byte); ok {
		r0 = rf(root, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash, []byte) error); ok {
		r1 = rf(root, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateRootFromBlock provides a mock function with given fields: bhash
func (_m *MockStorageAPI) GetStateRootFromBlock(bhash *common.Hash) (*common.Hash, error) {
	ret := _m.Called(bhash)

	var r0 *common.Hash
	if rf, ok := ret.Get(0).(func(*common.Hash) *common.Hash); ok {
		r0 = rf(bhash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash) error); ok {
		r1 = rf(bhash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorage provides a mock function with given fields: root, key
func (_m *MockStorageAPI) GetStorage(root *common.Hash, key []byte) ([]byte, error) {
	ret := _m.Called(root, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*common.Hash, []byte) []byte); ok {
		r0 = rf(root, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash, []byte) error); ok {
		r1 = rf(root, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageByBlockHash provides a mock function with given fields: bhash, key
func (_m *MockStorageAPI) GetStorageByBlockHash(bhash common.Hash, key []byte) ([]byte, error) {
	ret := _m.Called(bhash, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(common.Hash, []byte) []byte); ok {
		r0 = rf(bhash, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash, []byte) error); ok {
		r1 = rf(bhash, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageChild provides a mock function with given fields: root, keyToChild
func (_m *MockStorageAPI) GetStorageChild(root *common.Hash, keyToChild []byte) (*trie.Trie, error) {
	ret := _m.Called(root, keyToChild)

	var r0 *trie.Trie
	if rf, ok := ret.Get(0).(func(*common.Hash, []byte) *trie.Trie); ok {
		r0 = rf(root, keyToChild)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*trie.Trie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash, []byte) error); ok {
		r1 = rf(root, keyToChild)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageFromChild provides a mock function with given fields: root, keyToChild, key
func (_m *MockStorageAPI) GetStorageFromChild(root *common.Hash, keyToChild []byte, key []byte) ([]byte, error) {
	ret := _m.Called(root, keyToChild, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*common.Hash, []byte, []byte) []byte); ok {
		r0 = rf(root, keyToChild, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash, []byte, []byte) error); ok {
		r1 = rf(root, keyToChild, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterStorageObserver provides a mock function with given fields: observer
func (_m *MockStorageAPI) RegisterStorageObserver(observer state.Observer) {
	_m.Called(observer)
}

// UnregisterStorageObserver provides a mock function with given fields: observer
func (_m *MockStorageAPI) UnregisterStorageObserver(observer state.Observer) {
	_m.Called(observer)
}