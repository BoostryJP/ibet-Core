package decrypter

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/stretchr/testify/require"
)

func TestGetEncryptionAlgorithmSpec(t *testing.T) {
	d := &NodeKeyAwsKmsDecrypter{}

	algo, err := d.getEncryptionAlgorithmSpec("RSAES_OAEP_SHA_256")
	require.NoError(t, err)
	require.Equal(t, types.EncryptionAlgorithmSpecRsaesOaepSha256, algo)
}

func TestGetEncryptionAlgorithmSpec_Invalid(t *testing.T) {
	d := &NodeKeyAwsKmsDecrypter{}

	_, err := d.getEncryptionAlgorithmSpec("INVALID")
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported encryption algorithm")
}
