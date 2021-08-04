// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package wasmtime

import (
	"fmt"
	"testing"

	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/runtime"
	"github.com/ChainSafe/gossamer/lib/scale"
	"github.com/stretchr/testify/require"
)

func Test_NewTestInstance(t *testing.T) {
	inst := NewTestInstance(t, runtime.HOST_API_TEST_RUNTIME)
	require.NotNil(t, inst)
}

func Test_ext_hashing_blake2_128_version_1(t *testing.T) {
	inst := NewTestInstance(t, runtime.HOST_API_TEST_RUNTIME)

	data := []byte("helloworld")
	enc, err := scale.Encode(data)
	require.NoError(t, err)
fmt.Printf("enc data %v\n", enc)
	ret, err := inst.Exec("rtm_ext_hashing_blake2_128_version_1", enc)
	require.NoError(t, err)
fmt.Printf("ret %v\n", ret)
	hash, err := scale.Decode(ret, []byte{})
	require.NoError(t, err)

	expected, err := common.Blake2b128(data)
	require.NoError(t, err)
	require.Equal(t, expected[:], hash)
}

func Test_ext_storage_set_version_1(t *testing.T) {
	inst := NewTestInstance(t, runtime.HOST_API_TEST_RUNTIME)

	testkey := []byte("noot")
	testvalue := []byte("washere")

	encKey, err := scale.Encode(testkey)
	require.NoError(t, err)
	encValue, err := scale.Encode(testvalue)
	require.NoError(t, err)

	_, err = inst.Exec("rtm_ext_storage_set_version_1", append(encKey, encValue...))
	require.NoError(t, err)

	//val := inst.ctx.Storage.Get(testkey)
	//require.Equal(t, testvalue, val)
}