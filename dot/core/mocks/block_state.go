// Code generated by mockery v2.9.4. DO NOT EDIT.

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

// FreeFinalisedNotifierChannel provides a mock function with given fields: ch
func (_m *MockBlockState) FreeFinalisedNotifierChannel(ch chan *types.FinalisationInfo) {
	_m.Called(ch)
}

// FreeImportedBlockNotifierChannel provides a mock function with given fields: ch
func (_m *MockBlockState) FreeImportedBlockNotifierChannel(ch chan *types.Block) {
	_m.Called(ch)
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

// GetBlockStateRoot provides a mock function with given fields: bhash
func (_m *MockBlockState) GetBlockStateRoot(bhash common.Hash) (common.Hash, error) {
	ret := _m.Called(bhash)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Hash) common.Hash); ok {
		r0 = rf(bhash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(bhash)
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

// GetFinalisedNotifierChannel provides a mock function with given fields:
func (_m *MockBlockState) GetFinalisedNotifierChannel() chan *types.FinalisationInfo {
	ret := _m.Called()

	var r0 chan *types.FinalisationInfo
	if rf, ok := ret.Get(0).(func() chan *types.FinalisationInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.FinalisationInfo)
		}
	}

	return r0
}

// GetImportedBlockNotifierChannel provides a mock function with given fields:
func (_m *MockBlockState) GetImportedBlockNotifierChannel() chan *types.Block {
	ret := _m.Called()

	var r0 chan *types.Block
	if rf, ok := ret.Get(0).(func() chan *types.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.Block)
		}
	}

	return r0
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
