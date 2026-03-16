package common

import (
	"testing"

	"github.com/ethereum/go-ethereum/p2p/nodekey"
	"github.com/stretchr/testify/require"
)

func TestValidateConfigurationValues(t *testing.T) {
	err := validateConfigurationValues(nodekey.AwsConfig{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "secret name")

	err = validateConfigurationValues(nodekey.AwsConfig{SecretName: "boostry/ibet-network/quorum/nodekey"})
	require.NoError(t, err)
}
