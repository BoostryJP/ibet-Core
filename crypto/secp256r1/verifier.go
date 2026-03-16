// Copyright 2024 The go-ethereum Authors
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

// Package secp256r1 implements signature verification for the P256VERIFY precompile.
package secp256r1

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

// Verify checks the given signature (r, s) for the given hash and public key (x, y).
func Verify(hash []byte, r, s, x, y *big.Int) bool {
	if !isValidP256PublicKey(x, y) {
		return false
	}
	pk := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
	return ecdsa.Verify(pk, hash, r, s)
}

func isValidP256PublicKey(x, y *big.Int) bool {
	if x == nil || y == nil || x.Sign() < 0 || y.Sign() < 0 {
		return false
	}
	xb, yb := x.Bytes(), y.Bytes()
	if len(xb) > 32 || len(yb) > 32 {
		return false
	}
	encoded := make([]byte, 65)
	encoded[0] = 0x04
	copy(encoded[1+32-len(xb):33], xb)
	copy(encoded[33+32-len(yb):], yb)
	_, err := ecdh.P256().NewPublicKey(encoded)
	return err == nil
}
