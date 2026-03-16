package fetcher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildGetSecretValueInput_WithVersionID(t *testing.T) {
	input, err := buildGetSecretValueInput("boostry/ibet-network/quorum/nodekey", "f368ae7f-41e6-4d25-8e8e-a3aad0130846", "")
	require.NoError(t, err)
	require.NotNil(t, input)
	require.Equal(t, "boostry/ibet-network/quorum/nodekey", *input.SecretId)
	require.NotNil(t, input.VersionId)
	require.Equal(t, "f368ae7f-41e6-4d25-8e8e-a3aad0130846", *input.VersionId)
	require.Nil(t, input.VersionStage)
}

func TestBuildGetSecretValueInput_WithVersionStage(t *testing.T) {
	input, err := buildGetSecretValueInput("boostry/ibet-network/quorum/nodekey", "", "AWSCURRENT")
	require.NoError(t, err)
	require.NotNil(t, input)
	require.Equal(t, "boostry/ibet-network/quorum/nodekey", *input.SecretId)
	require.NotNil(t, input.VersionStage)
	require.Equal(t, "AWSCURRENT", *input.VersionStage)
	require.Nil(t, input.VersionId)
}

func TestBuildGetSecretValueInput_WithoutVersionSelector(t *testing.T) {
	_, err := buildGetSecretValueInput("boostry/ibet-network/quorum/nodekey", "", "")
	require.Error(t, err)
	require.Contains(t, err.Error(), "must be specified")
}

func TestBuildGetSecretValueInput_WithBothSelectors(t *testing.T) {
	_, err := buildGetSecretValueInput("boostry/ibet-network/quorum/nodekey", "f368ae7f-41e6-4d25-8e8e-a3aad0130846", "AWSCURRENT")
	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot both be specified")

}
