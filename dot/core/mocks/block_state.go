// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ChainSafe/gossamer/lib/common"

	mock "github.com/stretchr/testify/mock"

	runtime "github.com/ChainSafe/gossamer/lib/runtime"

	storage "github.com/ChainSafe/gossamer/lib/runtime/storage"

	types "github.com/ChainSafe/gossamer/dot/types"
)

// MockBlockState is an autogenerated mock type for the BlockState type
type MockBlockState struct {
	mock.Mock
}

// AddBlock provides a mock function with given fields: _a0
func (_m *MockBlockState) AddBlock(_a0 *types.Block) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Block) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockBlockState) AddBlockVdt(_a0 *types.BlockVdt) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.BlockVdt) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BestBlock provides a mock function with given fields:
func (_m *MockBlockState) BestBlock() (*types.Block, error) {
	ret := _m.Called()

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func() *types.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
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

// BestBlockHash provides a mock function with given fields:
func (_m *MockBlockState) BestBlockHash() common.Hash {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// BestBlockHeader provides a mock function with given fields:
func (_m *MockBlockState) BestBlockHeader() (*types.Header, error) {
	ret := _m.Called()

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func() *types.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
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

func (_m *MockBlockState) BestBlockHeaderVdt() (*types.HeaderVdt, error) {
	ret := _m.Called()

	var r0 *types.HeaderVdt
	if rf, ok := ret.Get(0).(func() *types.HeaderVdt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.HeaderVdt)
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

// BestBlockNumber provides a mock function with given fields:
func (_m *MockBlockState) BestBlockNumber() (*big.Int, error) {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
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

// BestBlockStateRoot provides a mock function with given fields:
func (_m *MockBlockState) BestBlockStateRoot() (common.Hash, error) {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
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

// GenesisHash provides a mock function with given fields:
func (_m *MockBlockState) GenesisHash() common.Hash {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// GetAllBlocksAtDepth provides a mock function with given fields: hash
func (_m *MockBlockState) GetAllBlocksAtDepth(hash common.Hash) []common.Hash {
	ret := _m.Called(hash)

	var r0 []common.Hash
	if rf, ok := ret.Get(0).(func(common.Hash) []common.Hash); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	return r0
}

// GetBlockBody provides a mock function with given fields: hash
func (_m *MockBlockState) GetBlockBody(hash common.Hash) (*types.Body, error) {
	ret := _m.Called(hash)

	var r0 *types.Body
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Body); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Body)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHash provides a mock function with given fields: _a0
func (_m *MockBlockState) GetBlockByHash(_a0 common.Hash) (*types.Block, error) {
	ret := _m.Called(_a0)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Block); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFinalisedHash provides a mock function with given fields: _a0, _a1
func (_m *MockBlockState) GetFinalisedHash(_a0 uint64, _a1 uint64) (common.Hash, error) {
	ret := _m.Called(_a0, _a1)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(uint64, uint64) common.Hash); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFinalisedHeader provides a mock function with given fields: _a0, _a1
func (_m *MockBlockState) GetFinalisedHeader(_a0 uint64, _a1 uint64) (*types.Header, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(uint64, uint64) *types.Header); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRuntime provides a mock function with given fields: _a0
func (_m *MockBlockState) GetRuntime(_a0 *common.Hash) (runtime.Instance, error) {
	ret := _m.Called(_a0)

	var r0 runtime.Instance
	if rf, ok := ret.Get(0).(func(*common.Hash) runtime.Instance); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlotForBlock provides a mock function with given fields: _a0
func (_m *MockBlockState) GetSlotForBlock(_a0 common.Hash) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(common.Hash) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleRuntimeChanges provides a mock function with given fields: newState, in, bHash
func (_m *MockBlockState) HandleRuntimeChanges(newState *storage.TrieState, in runtime.Instance, bHash common.Hash) error {
	ret := _m.Called(newState, in, bHash)

	var r0 error
	if rf, ok := ret.Get(0).(func(*storage.TrieState, runtime.Instance, common.Hash) error); ok {
		r0 = rf(newState, in, bHash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HighestCommonAncestor provides a mock function with given fields: a, b
func (_m *MockBlockState) HighestCommonAncestor(a common.Hash, b common.Hash) (common.Hash, error) {
	ret := _m.Called(a, b)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Hash, common.Hash) common.Hash); ok {
		r0 = rf(a, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash, common.Hash) error); ok {
		r1 = rf(a, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterFinalizedChannel provides a mock function with given fields: ch
func (_m *MockBlockState) RegisterFinalizedChannel(ch chan<- *types.FinalisationInfo) (byte, error) {
	ret := _m.Called(ch)

	var r0 byte
	if rf, ok := ret.Get(0).(func(chan<- *types.FinalisationInfo) byte); ok {
		r0 = rf(ch)
	} else {
		r0 = ret.Get(0).(byte)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(chan<- *types.FinalisationInfo) error); ok {
		r1 = rf(ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterImportedChannel provides a mock function with given fields: ch
func (_m *MockBlockState) RegisterImportedChannel(ch chan<- *types.Block) (byte, error) {
	ret := _m.Called(ch)

	var r0 byte
	if rf, ok := ret.Get(0).(func(chan<- *types.Block) byte); ok {
		r0 = rf(ch)
	} else {
		r0 = ret.Get(0).(byte)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(chan<- *types.Block) error); ok {
		r1 = rf(ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetFinalisedHash provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockBlockState) SetFinalisedHash(_a0 common.Hash, _a1 uint64, _a2 uint64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, uint64, uint64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreRuntime provides a mock function with given fields: _a0, _a1
func (_m *MockBlockState) StoreRuntime(_a0 common.Hash, _a1 runtime.Instance) {
	_m.Called(_a0, _a1)
}

// SubChain provides a mock function with given fields: start, end
func (_m *MockBlockState) SubChain(start common.Hash, end common.Hash) ([]common.Hash, error) {
	ret := _m.Called(start, end)

	var r0 []common.Hash
	if rf, ok := ret.Get(0).(func(common.Hash, common.Hash) []common.Hash); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash, common.Hash) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnregisterFinalisedChannel provides a mock function with given fields: id
func (_m *MockBlockState) UnregisterFinalisedChannel(id byte) {
	_m.Called(id)
}

// UnregisterImportedChannel provides a mock function with given fields: id
func (_m *MockBlockState) UnregisterImportedChannel(id byte) {
	_m.Called(id)
}
