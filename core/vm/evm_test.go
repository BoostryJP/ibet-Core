package vm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func TestActivePrecompiles(t *testing.T) {
	tests := []struct {
		name string
		evm  *EVM
		want []common.Address
	}{
		{
			name: "berlin",
			evm: &EVM{
				chainRules: params.Rules{
					IsBerlin:            true,
					IsPrivacyPrecompile: false,
				},
			},
			want: []common.Address{
				common.BytesToAddress([]byte{0x01}),
				common.BytesToAddress([]byte{0x02}),
				common.BytesToAddress([]byte{0x03}),
				common.BytesToAddress([]byte{0x04}),
				common.BytesToAddress([]byte{0x05}),
				common.BytesToAddress([]byte{0x06}),
				common.BytesToAddress([]byte{0x07}),
				common.BytesToAddress([]byte{0x08}),
				common.BytesToAddress([]byte{0x09}),
				common.BytesToAddress([]byte{0x0b}),
				common.BytesToAddress([]byte{0x0c}),
				common.BytesToAddress([]byte{0x0d}),
				common.BytesToAddress([]byte{0x0e}),
				common.BytesToAddress([]byte{0x0f}),
				common.BytesToAddress([]byte{0x10}),
				common.BytesToAddress([]byte{0x11}),
				common.BytesToAddress([]byte{0x01, 0x00}),
			},
		},
		{
			name: "istanbul-plus-quorum-privacy",
			evm: &EVM{
				chainRules: params.Rules{
					IsIstanbul:          true,
					IsPrivacyPrecompile: true,
				},
			},
			want: []common.Address{
				common.BytesToAddress([]byte{0x01}),
				common.BytesToAddress([]byte{0x02}),
				common.BytesToAddress([]byte{0x03}),
				common.BytesToAddress([]byte{0x04}),
				common.BytesToAddress([]byte{0x05}),
				common.BytesToAddress([]byte{0x06}),
				common.BytesToAddress([]byte{0x07}),
				common.BytesToAddress([]byte{0x08}),
				common.BytesToAddress([]byte{0x09}),
				common.QuorumPrivacyPrecompileContractAddress(),
			},
		},
		{
			name: "homestead-plus-quorum-privacy",
			evm: &EVM{
				chainRules: params.Rules{
					IsHomestead:         true,
					IsPrivacyPrecompile: true,
				},
			},
			want: []common.Address{
				common.BytesToAddress([]byte{0x01}),
				common.BytesToAddress([]byte{0x02}),
				common.BytesToAddress([]byte{0x03}),
				common.BytesToAddress([]byte{0x04}),
				common.QuorumPrivacyPrecompileContractAddress(),
			},
		},
		{
			name: "istanbul",
			evm: &EVM{
				chainRules: params.Rules{
					IsIstanbul:          true,
					IsPrivacyPrecompile: false,
				},
			},
			want: []common.Address{
				common.BytesToAddress([]byte{0x01}),
				common.BytesToAddress([]byte{0x02}),
				common.BytesToAddress([]byte{0x03}),
				common.BytesToAddress([]byte{0x04}),
				common.BytesToAddress([]byte{0x05}),
				common.BytesToAddress([]byte{0x06}),
				common.BytesToAddress([]byte{0x07}),
				common.BytesToAddress([]byte{0x08}),
				common.BytesToAddress([]byte{0x09}),
			},
		},
		{
			name: "homestead",
			evm: &EVM{
				chainRules: params.Rules{
					IsHomestead:         true,
					IsPrivacyPrecompile: false,
				},
			},
			want: []common.Address{
				common.BytesToAddress([]byte{0x01}),
				common.BytesToAddress([]byte{0x02}),
				common.BytesToAddress([]byte{0x03}),
				common.BytesToAddress([]byte{0x04}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ActivePrecompiles(tt.evm.chainRules)
			require.ElementsMatchf(t, tt.want, got, "want: %v, got: %v", tt.want, got)
		})
	}
}
