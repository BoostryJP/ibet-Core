// Copyright 2025 The go-ethereum Authors
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

package nat

import (
	"fmt"
	"net"
	"testing"

	stunV2 "github.com/pion/stun/v2"
	"github.com/pion/stun/v2/stuntest"
	"github.com/stretchr/testify/assert"
)

func TestNatStun(t *testing.T) {
	mappedIP := net.ParseIP("203.0.113.10").To4()
	serverAddr, closeServer, err := stuntest.NewUDPServer(t, "udp4", 2048, func(req []byte) ([]byte, error) {
		reqMsg := new(stunV2.Message)
		if err := stunV2.Decode(req, reqMsg); err != nil {
			return nil, err
		}
		if reqMsg.Type != stunV2.BindingRequest {
			return nil, fmt.Errorf("unexpected STUN message type: %v", reqMsg.Type)
		}
		resp, err := stunV2.Build(reqMsg, stunV2.BindingSuccess, &stunV2.XORMappedAddress{
			IP:   mappedIP,
			Port: 54321,
		})
		if err != nil {
			return nil, err
		}
		return resp.Raw, nil
	})
	assert.NoError(t, err)
	defer closeServer(t)

	nat, err := newSTUN(serverAddr.String())
	assert.NoError(t, err)
	ip, err := nat.ExternalIP()
	assert.NoError(t, err)
	assert.Equal(t, mappedIP, ip)
}

func TestUnreachedNatServer(t *testing.T) {
	stun := &stun{
		serverList: []string{"198.51.100.2:1234", "198.51.100.5"},
	}
	_, err := stun.ExternalIP()
	if err != errSTUNFailed {
		t.Fatal("wrong error:", err)
	}
}
