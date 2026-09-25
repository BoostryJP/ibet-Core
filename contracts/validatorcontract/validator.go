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

// Package validatorcontract is an on-chain validator set contract wrapper.
package validatorcontract

//go:generate abigen --sol contract/validator_set.sol --pkg contract --out contract/validator_set.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/validatorcontract/contract"
)

// ValidatorSet is a Go wrapper around an on-chain validator set contract.
type ValidatorSet struct {
	address  common.Address
	contract *contract.ValidatorSet
}

// NewValidatorSet binds validator set contract and returns a wrapper instance.
func NewValidatorSet(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorSet, error) {
	c, err := contract.NewValidatorSet(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSet{address: contractAddr, contract: c}, nil
}

// ContractAddr returns the address of contract.
func (set *ValidatorSet) ContractAddr() common.Address {
	return set.address
}

// Contract returns the underlying contract instance.
func (set *ValidatorSet) Contract() *contract.ValidatorSet {
	return set.contract
}
