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

package validatorcontract

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/validatorcontract/contract"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	operationAdd    uint8 = 0
	operationRemove uint8 = 1
)

var validatorContractChainID = big.NewInt(1337)

type testAccount struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

func newTestAccount(t *testing.T) testAccount {
	t.Helper()

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}
	return testAccount{
		key:  key,
		addr: crypto.PubkeyToAddress(key.PublicKey),
	}
}

func newTestAuth(t *testing.T, account testAccount) *bind.TransactOpts {
	t.Helper()

	auth, err := bind.NewKeyedTransactorWithChainID(account.key, validatorContractChainID)
	if err != nil {
		t.Fatalf("failed to create transactor: %v", err)
	}
	return auth
}

func deployValidatorSet(t *testing.T) (*backends.SimulatedBackend, *contract.ValidatorSet, []testAccount) {
	t.Helper()

	accounts := make([]testAccount, 6)
	alloc := make(core.GenesisAlloc)
	for i := range accounts {
		accounts[i] = newTestAccount(t)
		alloc[accounts[i].addr] = core.GenesisAccount{Balance: big.NewInt(1000000000000000000)}
	}

	backend := backends.NewSimulatedBackend(alloc, 10000000)
	initialValidators := []common.Address{
		accounts[0].addr,
		accounts[1].addr,
		accounts[2].addr,
		accounts[3].addr,
	}
	_, _, validatorSet, err := contract.DeployValidatorSet(newTestAuth(t, accounts[0]), backend, initialValidators)
	if err != nil {
		backend.Close()
		t.Fatalf("failed to deploy validator set: %v", err)
	}
	backend.Commit()

	return backend, validatorSet, accounts
}

func deployValidatorSetWithValidators(t *testing.T, validators []common.Address) (*backends.SimulatedBackend, *types.Transaction, error) {
	t.Helper()

	deployer := newTestAccount(t)
	alloc := core.GenesisAlloc{deployer.addr: core.GenesisAccount{Balance: big.NewInt(1000000000000000000)}}
	for _, validator := range validators {
		alloc[validator] = core.GenesisAccount{Balance: big.NewInt(1000000000000000000)}
	}
	backend := backends.NewSimulatedBackend(alloc, 10000000)

	_, tx, _, err := contract.DeployValidatorSet(newTestAuth(t, deployer), backend, validators)
	return backend, tx, err
}

func mustTransact(t *testing.T, backend *backends.SimulatedBackend, transact func() (*types.Transaction, error)) {
	t.Helper()

	tx, err := transact()
	if err != nil {
		t.Fatalf("transaction failed before mining: %v", err)
	}
	backend.Commit()
	receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("failed to fetch transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Fatalf("transaction reverted: %s", tx.Hash())
	}
}

func mustReject(t *testing.T, backend *backends.SimulatedBackend, transact func() (*types.Transaction, error)) {
	t.Helper()

	tx, err := transact()
	if err != nil {
		return
	}
	backend.Commit()
	receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("failed to fetch transaction receipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusFailed {
		t.Fatalf("transaction unexpectedly succeeded: %s", tx.Hash())
	}
}

func assertAddresses(t *testing.T, got []common.Address, want []common.Address) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("validator mismatch\n got: %v\nwant: %v", got, want)
	}
}

func TestValidatorSetInitialState(t *testing.T) {
	// Viewpoint:
	//  A freshly deployed 4-validator set exposes the same active/full validator list and derives f=1, quorum=3.

	backend, validatorSet, accounts := deployValidatorSet(t)
	defer backend.Close()

	wantValidators := []common.Address{accounts[0].addr, accounts[1].addr, accounts[2].addr, accounts[3].addr}
	activeValidators, err := validatorSet.GetValidators(nil)
	if err != nil {
		t.Fatalf("failed to get validators: %v", err)
	}
	assertAddresses(t, activeValidators, wantValidators)

	allValidators, err := validatorSet.GetAllValidators(nil)
	if err != nil {
		t.Fatalf("failed to get all validators: %v", err)
	}
	assertAddresses(t, allValidators, wantValidators)

	f, err := validatorSet.FaultTolerance(nil)
	if err != nil {
		t.Fatalf("failed to get fault tolerance: %v", err)
	}
	if f.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("unexpected fault tolerance: %s", f)
	}

	quorum, err := validatorSet.QuorumSize(nil)
	if err != nil {
		t.Fatalf("failed to get quorum size: %v", err)
	}
	if quorum.Cmp(big.NewInt(3)) != 0 {
		t.Fatalf("unexpected quorum size: %s", quorum)
	}
}

func TestValidatorSetRejectsInvalidInitialValidators(t *testing.T) {
	// Viewpoint:
	//  Constructor validation rejects empty, zero-address, and duplicate initial validator lists.

	for _, tt := range []struct {
		name       string
		validators func() []common.Address
	}{
		{
			name: "empty",
			validators: func() []common.Address {
				return nil
			},
		},
		{
			name: "zero address",
			validators: func() []common.Address {
				return []common.Address{newTestAccount(t).addr, common.Address{}}
			},
		},
		{
			name: "duplicate",
			validators: func() []common.Address {
				account := newTestAccount(t)
				return []common.Address{account.addr, account.addr}
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			backend, tx, err := deployValidatorSetWithValidators(t, tt.validators())
			defer backend.Close()
			if err != nil {
				return
			}
			backend.Commit()
			receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				t.Fatalf("failed to fetch deploy receipt: %v", err)
			}
			if receipt.Status != types.ReceiptStatusFailed {
				t.Fatalf("invalid validator deployment unexpectedly succeeded: %s", tx.Hash())
			}
		})
	}
}

func TestValidatorSetAddAndRemoveByQuorum(t *testing.T) {
	// Viewpoint:
	//  Add/Remove proposals do not execute before quorum and execute immediately when the third vote reaches 2f+1.

	backend, validatorSet, accounts := deployValidatorSet(t)
	defer backend.Close()

	candidate := accounts[4].addr
	// Two votes are recorded for the add proposal, but the candidate is not registered yet.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[0]), candidate)
	})
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[1]), candidate)
	})

	isValidator, err := validatorSet.IsValidator(nil, candidate)
	if err != nil {
		t.Fatalf("failed to check candidate before quorum: %v", err)
	}
	if isValidator {
		t.Fatal("candidate became validator before quorum")
	}
	votes, err := validatorSet.ProposalVotes(nil, operationAdd, candidate)
	if err != nil {
		t.Fatalf("failed to get add proposal votes: %v", err)
	}
	if votes.Cmp(big.NewInt(2)) != 0 {
		t.Fatalf("unexpected add proposal votes: %s", votes)
	}

	// The third vote reaches quorum and appends the candidate to the registered validator list.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[2]), candidate)
	})
	isValidator, err = validatorSet.IsValidator(nil, candidate)
	if err != nil {
		t.Fatalf("failed to check candidate after quorum: %v", err)
	}
	if !isValidator {
		t.Fatal("candidate did not become validator after quorum")
	}

	allValidators, err := validatorSet.GetAllValidators(nil)
	if err != nil {
		t.Fatalf("failed to get all validators after add: %v", err)
	}
	assertAddresses(t, allValidators, []common.Address{accounts[0].addr, accounts[1].addr, accounts[2].addr, accounts[3].addr, candidate})

	// Two votes are recorded for the remove proposal, but the candidate remains registered.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteRemoveValidator(newTestAuth(t, accounts[0]), candidate)
	})
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteRemoveValidator(newTestAuth(t, accounts[1]), candidate)
	})
	isValidator, err = validatorSet.IsValidator(nil, candidate)
	if err != nil {
		t.Fatalf("failed to check removal candidate before quorum: %v", err)
	}
	if !isValidator {
		t.Fatal("candidate was removed before quorum")
	}

	// The third remove vote reaches quorum and removes the candidate while preserving remaining validator order.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteRemoveValidator(newTestAuth(t, accounts[2]), candidate)
	})
	isValidator, err = validatorSet.IsValidator(nil, candidate)
	if err != nil {
		t.Fatalf("failed to check removal candidate after quorum: %v", err)
	}
	if isValidator {
		t.Fatal("candidate was not removed after quorum")
	}

	allValidators, err = validatorSet.GetAllValidators(nil)
	if err != nil {
		t.Fatalf("failed to get all validators after remove: %v", err)
	}
	assertAddresses(t, allValidators, []common.Address{accounts[0].addr, accounts[1].addr, accounts[2].addr, accounts[3].addr})
}

func TestValidatorSetRejectsInvalidVotes(t *testing.T) {
	// Viewpoint:
	//  Governance votes must come from validators, target valid candidates,
	//  and be counted at most once per validator.

	backend, validatorSet, accounts := deployValidatorSet(t)
	defer backend.Close()

	candidate := accounts[4].addr
	// A zero-address add candidate and a non-validator voter are both invalid.
	mustReject(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[0]), common.Address{})
	})
	mustReject(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[5]), candidate)
	})

	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[0]), candidate)
	})
	// The same validator cannot vote twice for the same operation/candidate proposal.
	mustReject(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteAddValidator(newTestAuth(t, accounts[0]), candidate)
	})

	hasVoted, err := validatorSet.HasVoted(nil, operationAdd, candidate, accounts[0].addr)
	if err != nil {
		t.Fatalf("failed to check vote status: %v", err)
	}
	if !hasVoted {
		t.Fatal("expected validator vote to be recorded")
	}
}

func TestValidatorSetMaintenance(t *testing.T) {
	// Viewpoint:
	//  Maintenance excludes only the caller from active validators,
	//  enforces the f-validator cap, and restores the active set on exit.

	backend, validatorSet, accounts := deployValidatorSet(t)
	defer backend.Close()

	// With 4 registered validators, f=1, so the first validator may enter maintenance.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.EnterMaintenance(newTestAuth(t, accounts[0]))
	})

	activeValidators, err := validatorSet.GetValidators(nil)
	if err != nil {
		t.Fatalf("failed to get active validators in maintenance: %v", err)
	}
	assertAddresses(t, activeValidators, []common.Address{accounts[1].addr, accounts[2].addr, accounts[3].addr})

	maintenanceSize, err := validatorSet.MaintenanceSize(nil)
	if err != nil {
		t.Fatalf("failed to get maintenance size: %v", err)
	}
	if maintenanceSize.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("unexpected maintenance size: %s", maintenanceSize)
	}

	inMaintenance, err := validatorSet.IsInMaintenance(nil, accounts[0].addr)
	if err != nil {
		t.Fatalf("failed to get maintenance flag: %v", err)
	}
	if !inMaintenance {
		t.Fatal("validator is not marked in maintenance")
	}

	// A second simultaneous maintenance entry exceeds f, and non-validators cannot enter maintenance.
	mustReject(t, backend, func() (*types.Transaction, error) {
		return validatorSet.EnterMaintenance(newTestAuth(t, accounts[1]))
	})
	mustReject(t, backend, func() (*types.Transaction, error) {
		return validatorSet.EnterMaintenance(newTestAuth(t, accounts[5]))
	})

	// Exiting maintenance returns the validator to the active list in its original order.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.ExitMaintenance(newTestAuth(t, accounts[0]))
	})
	activeValidators, err = validatorSet.GetValidators(nil)
	if err != nil {
		t.Fatalf("failed to get active validators after maintenance exit: %v", err)
	}
	assertAddresses(t, activeValidators, []common.Address{accounts[0].addr, accounts[1].addr, accounts[2].addr, accounts[3].addr})
}

func TestValidatorSetRemoveValidatorInMaintenance(t *testing.T) {
	// Viewpoint:
	//  Removing a validator currently in maintenance clears its maintenance flag
	//  and decrements maintenance count.

	backend, validatorSet, accounts := deployValidatorSet(t)
	defer backend.Close()

	// The validator remains eligible to vote while in maintenance, and three registered validators can remove it.
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.EnterMaintenance(newTestAuth(t, accounts[0]))
	})
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteRemoveValidator(newTestAuth(t, accounts[0]), accounts[0].addr)
	})
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteRemoveValidator(newTestAuth(t, accounts[1]), accounts[0].addr)
	})
	mustTransact(t, backend, func() (*types.Transaction, error) {
		return validatorSet.VoteRemoveValidator(newTestAuth(t, accounts[2]), accounts[0].addr)
	})

	isValidator, err := validatorSet.IsValidator(nil, accounts[0].addr)
	if err != nil {
		t.Fatalf("failed to check removed validator: %v", err)
	}
	if isValidator {
		t.Fatal("maintenance validator was not removed")
	}
	maintenanceSize, err := validatorSet.MaintenanceSize(nil)
	if err != nil {
		t.Fatalf("failed to get maintenance size after removal: %v", err)
	}
	if maintenanceSize.Sign() != 0 {
		t.Fatalf("maintenance count did not decrement after removal: %s", maintenanceSize)
	}

	activeValidators, err := validatorSet.GetValidators(nil)
	if err != nil {
		t.Fatalf("failed to get validators after maintenance removal: %v", err)
	}
	assertAddresses(t, activeValidators, []common.Address{accounts[1].addr, accounts[2].addr, accounts[3].addr})
}
