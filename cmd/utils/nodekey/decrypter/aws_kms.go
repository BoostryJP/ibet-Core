package decrypter

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/ethereum/go-ethereum/cmd/utils/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type NodeKeyAwsKmsDecrypter struct {
	aws *common.AwsClient
}

func NewNodeKeyAwsKmsDecrypter(configBytes []byte) (*NodeKeyAwsKmsDecrypter, error) {
	awsClient, err := common.NewAwsClient(configBytes)
	return &NodeKeyAwsKmsDecrypter{aws: awsClient}, err
}

func (d *NodeKeyAwsKmsDecrypter) DecryptNodeKey(base64StringData string) (*ecdsa.PrivateKey, error) {
	ctx := context.Background()
	if d.aws.Config.KmsKeyId == "" {
		return nil, errors.New("configuration [KmsKeyId] should not be empty")
	}

	log.Info("Decrypting node key with AWS KMS", "kmsKeyId", d.aws.Config.KmsKeyId)

	// Secrets Manager stores ciphertext as a base64 string; KMS expects raw bytes.
	bytesData, err := base64.StdEncoding.DecodeString(base64StringData)
	if err != nil {
		return nil, err
	}

	input := &kms.DecryptInput{
		CiphertextBlob: bytesData,
		KeyId:          aws.String(d.aws.Config.KmsKeyId),
	}
	if d.aws.Config.KmsEncryptionAlgorithm != "" {
		// Optional override for asymmetric CMKs; omitted for symmetric default flow.
		algo, err := d.getEncryptionAlgorithmSpec(d.aws.Config.KmsEncryptionAlgorithm)
		if err != nil {
			return nil, err
		}
		input.EncryptionAlgorithm = algo
	}

	result, err := d.aws.KMSClient.Decrypt(ctx, input)
	if err != nil {
		return nil, err
	}
	if result != nil {
		// Decrypted plaintext is expected to be hex private key material.
		privateKey, err := crypto.HexToECDSA(strings.TrimSpace(string(result.Plaintext)))
		if err != nil {
			return nil, fmt.Errorf("unable to convert node key data to private key: %w", err)
		}
		return privateKey, nil
	}
	return nil, fmt.Errorf("unable to decrypt node key using AWS KMS using KeyId [%s]", d.aws.Config.KmsKeyId)
}

func (d *NodeKeyAwsKmsDecrypter) getEncryptionAlgorithmSpec(algo string) (types.EncryptionAlgorithmSpec, error) {
	for _, val := range types.EncryptionAlgorithmSpec("").Values() {
		if string(val) == algo {
			return val, nil
		}
	}
	return "", fmt.Errorf("unsupported encryption algorithm: %s", algo)
}
