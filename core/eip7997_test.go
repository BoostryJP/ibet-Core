// Copyright 2026 BOOSTRY Co., Ltd.
// This file is part of ibet-Core.
//
// ibet-Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ibet-Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with ibet-Core. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func newEIP7997State(t *testing.T) (*state.StateDB, state.Database) {
	t.Helper()

	db := state.NewDatabase(rawdb.NewMemoryDatabase())
	statedb, err := state.New(common.Hash{}, db, nil)
	if err != nil {
		t.Fatal(err)
	}
	return statedb, db
}

func stateAtRoot(t *testing.T, db state.Database, root common.Hash) *state.StateDB {
	t.Helper()

	statedb, err := state.New(root, db, nil)
	if err != nil {
		t.Fatal(err)
	}
	return statedb
}

func amsterdamTestConfig(block uint64) *params.ChainConfig {
	config := *params.TestChainConfig
	config.AmsterdamBlock = new(big.Int).SetUint64(block)
	return &config
}

// Verify that ApplyEIP7997 installs the factory code and nonce=1 into an empty public state.
func TestApplyEIP7997(t *testing.T) {
	statedb, _ := newEIP7997State(t)

	// Apply the irregular state transition to an address that does not exist yet.
	misc.ApplyEIP7997(statedb)

	if got := statedb.GetCode(params.DeterministicFactoryAddress); !bytes.Equal(got, params.DeterministicFactoryCode) {
		t.Fatalf("factory code mismatch:\n got %x\nwant %x", got, params.DeterministicFactoryCode)
	}
	if got := statedb.GetNonce(params.DeterministicFactoryAddress); got != 1 {
		t.Fatalf("factory nonce = %d, want 1", got)
	}
}

// Verify that ApplyEIP7997 preserves an existing non-zero nonce when canonical code is already installed.
func TestApplyEIP7997Existing(t *testing.T) {
	statedb, _ := newEIP7997State(t)

	// Simulate a state where the factory has already been installed and used.
	statedb.SetCode(params.DeterministicFactoryAddress, params.DeterministicFactoryCode)
	statedb.SetNonce(params.DeterministicFactoryAddress, 5)

	misc.ApplyEIP7997(statedb)

	if got := statedb.GetNonce(params.DeterministicFactoryAddress); got != 5 {
		t.Fatalf("existing factory nonce overwritten: got %d, want 5", got)
	}
}

// Verify that ApplyEIP7997 replaces wrong factory code with canonical code while preserving a non-zero nonce.
func TestApplyEIP7997WrongCode(t *testing.T) {
	statedb, _ := newEIP7997State(t)

	// Simulate a pre-existing account at the factory address with unexpected code.
	statedb.SetCode(params.DeterministicFactoryAddress, []byte{0x60, 0x00})
	statedb.SetNonce(params.DeterministicFactoryAddress, 7)

	misc.ApplyEIP7997(statedb)

	if got := statedb.GetCode(params.DeterministicFactoryAddress); !bytes.Equal(got, params.DeterministicFactoryCode) {
		t.Fatalf("factory code not overwritten:\n got %x\nwant %x", got, params.DeterministicFactoryCode)
	}
	if got := statedb.GetNonce(params.DeterministicFactoryAddress); got != 7 {
		t.Fatalf("factory nonce = %d, want 7", got)
	}
}

// Verify that the installed factory can deploy runtime code with CREATE2 from salt and initcode.
func TestEIP7997FactoryDeploys(t *testing.T) {
	statedb, _ := newEIP7997State(t)
	misc.ApplyEIP7997(statedb)

	var (
		caller   = common.Address{0xca}
		salt     [32]byte
		initcode = common.FromHex("60fe60005360016000f3")
	)
	salt[31] = 0x42

	// The factory calldata is 32 bytes of CREATE2 salt followed by the initcode.
	// This initcode deploys a one-byte runtime code: 0xfe.
	input := append(append([]byte{}, salt[:]...), initcode...)
	blockContext := vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		BlockNumber: big.NewInt(1),
		GasLimit:    10_000_000,
	}
	txContext := vm.TxContext{Origin: caller, GasPrice: new(big.Int)}
	config := amsterdamTestConfig(0)
	evm := vm.NewEVM(blockContext, txContext, statedb, statedb, config, vm.Config{})

	// The test executes with Amsterdam rules, so the factory account must be in the
	// access list before calling it directly through the EVM.
	statedb.PrepareAccessList(caller, &params.DeterministicFactoryAddress, vm.ActivePrecompiles(config.Rules(blockContext.BlockNumber)), nil)

	ret, _, err := evm.Call(vm.AccountRef(caller), params.DeterministicFactoryAddress, input, 10_000_000, new(big.Int))
	if err != nil {
		t.Fatalf("factory call failed: %v", err)
	}

	// The factory returns the deployed address, which should match the CREATE2
	// address derived from the factory address, salt, and initcode hash.
	want := crypto.CreateAddress2(params.DeterministicFactoryAddress, salt, crypto.Keccak256(initcode))
	if len(ret) != 20 {
		t.Fatalf("factory returned %d bytes, want 20", len(ret))
	}
	if got := common.BytesToAddress(ret); got != want {
		t.Fatalf("factory returned address %x, want %x", got, want)
	}
	if code := statedb.GetCode(want); !bytes.Equal(code, []byte{0xfe}) {
		t.Fatalf("deployed runtime code = %x, want fe", code)
	}
}

// Verify that chain generation includes the factory in the state root only at the Amsterdam transition block.
func TestEIP7997AmsterdamTransition(t *testing.T) {
	config := amsterdamTestConfig(1)
	db := rawdb.NewMemoryDatabase()
	genesis := (&Genesis{Config: config, Alloc: GenesisAlloc{}, Difficulty: big.NewInt(1)}).MustCommit(db)

	// The factory must not be present before the Amsterdam transition block.
	statedb := stateAtRoot(t, state.NewDatabase(db), genesis.Root())
	if code := statedb.GetCode(params.DeterministicFactoryAddress); len(code) != 0 {
		t.Fatalf("factory code present before amsterdam transition: %x", code)
	}

	// Generate block 1, where Amsterdam activates, and check the generated state root.
	blocks, _ := GenerateChain(config, genesis, ethash.NewFaker(), db, 1, nil)
	statedb = stateAtRoot(t, state.NewDatabase(db), blocks[0].Root())
	if got := statedb.GetCode(params.DeterministicFactoryAddress); !bytes.Equal(got, params.DeterministicFactoryCode) {
		t.Fatalf("factory code missing from generated amsterdam block: got %x", got)
	}
}

// Verify that block processing inserts the factory into public state at the Amsterdam transition block.
func TestEIP7997AmsterdamTransitionBlockProcessing(t *testing.T) {
	config := amsterdamTestConfig(1)
	db := rawdb.NewMemoryDatabase()
	genesis := (&Genesis{Config: config, Alloc: GenesisAlloc{}, Difficulty: big.NewInt(1)}).MustCommit(db)

	// InsertChain reprocesses the generated block, exercising StateProcessor
	// instead of only trusting the state root produced by GenerateChain.
	blockchain, err := NewBlockChain(db, nil, config, ethash.NewFaker(), vm.Config{}, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer blockchain.Stop()

	blocks, _ := GenerateChain(config, genesis, ethash.NewFaker(), db, 1, nil)
	if _, err := blockchain.InsertChain(blocks); err != nil {
		t.Fatal(err)
	}

	// After processing block 1, the canonical head state should contain the factory.
	statedb := stateAtRoot(t, state.NewDatabase(db), blockchain.CurrentBlock().Root())
	if got := statedb.GetCode(params.DeterministicFactoryAddress); !bytes.Equal(got, params.DeterministicFactoryCode) {
		t.Fatalf("factory code missing after block processing: got %x", got)
	}
}
