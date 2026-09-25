// Copyright 2026 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package misc

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// ApplyEIP7997 inserts the deterministic deployment factory into the state as an
// irregular state transition, as specified by EIP-7997.
func ApplyEIP7997(statedb *state.StateDB) {
	wantHash := crypto.Keccak256Hash(params.DeterministicFactoryCode)
	if statedb.GetCodeHash(params.DeterministicFactoryAddress) == wantHash {
		return
	}
	if !statedb.Exist(params.DeterministicFactoryAddress) {
		statedb.CreateAccount(params.DeterministicFactoryAddress)
	}
	statedb.SetCode(params.DeterministicFactoryAddress, params.DeterministicFactoryCode)
	if statedb.GetNonce(params.DeterministicFactoryAddress) == 0 {
		statedb.SetNonce(params.DeterministicFactoryAddress, 1)
	}
}
