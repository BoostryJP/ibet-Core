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
	require.Error(t, err)
	require.Contains(t, err.Error(), "either secret version id or secret version stage")

	err = validateConfigurationValues(nodekey.AwsConfig{
		SecretName:         "boostry/ibet-network/quorum/nodekey",
		SecretVersionStage: "AWSCURRENT",
	})
	require.NoError(t, err)

	err = validateConfigurationValues(nodekey.AwsConfig{
		SecretName:      "boostry/ibet-network/quorum/nodekey",
		SecretVersionId: "f368ae7f-41e6-4d25-8e8e-a3aad0130846",
	})
	require.NoError(t, err)

	err = validateConfigurationValues(nodekey.AwsConfig{
		SecretName:         "boostry/ibet-network/quorum/nodekey",
		SecretVersionId:    "f368ae7f-41e6-4d25-8e8e-a3aad0130846",
		SecretVersionStage: "AWSCURRENT",
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "mutually exclusive")
}
