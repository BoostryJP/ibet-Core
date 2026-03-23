package fetcher

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/ethereum/go-ethereum/cmd/utils/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type NodeKeyAwsSecretsManagerFetcher struct {
	aws *common.AwsClient
}

func NewNodeKeyAwsSecretsManagerFetcher(configBytes []byte) (*NodeKeyAwsSecretsManagerFetcher, error) {
	awsClient, err := common.NewAwsClient(configBytes)
	return &NodeKeyAwsSecretsManagerFetcher{aws: awsClient}, err
}

func (f *NodeKeyAwsSecretsManagerFetcher) FetchNodeKey() (*ecdsa.PrivateKey, error) {
	secretData, err := f.fetch()
	if err != nil {
		return nil, err
	}
	// For non-encrypted mode, secret payload is expected to be raw hex key data.
	privateKey, err := crypto.HexToECDSA(secretData)
	if err != nil {
		return nil, fmt.Errorf("unable to convert node key from secret manager to private key: %w", err)
	}
	return privateKey, nil
}

func (f *NodeKeyAwsSecretsManagerFetcher) FetchEncryptedNodeKey() (string, error) {
	return f.fetch()
}

func (f *NodeKeyAwsSecretsManagerFetcher) fetch() (string, error) {
	secretName := f.aws.Config.SecretName
	secretVersionID := f.aws.Config.SecretVersionId
	secretVersionStage := f.aws.Config.SecretVersionStage

	log.Info("Fetching node key from AWS Secrets Manager", "secret", secretName)

	input, err := buildGetSecretValueInput(secretName, secretVersionID, secretVersionStage)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	resp, err := f.aws.SecretsClient.GetSecretValue(ctx, input)
	if err != nil {
		return "", err
	}
	if resp.SecretString == nil || *resp.SecretString == "" {
		return "", fmt.Errorf("using key [%s], data from secret manager is empty", secretName)
	}
	return *resp.SecretString, nil
}

func buildGetSecretValueInput(secretName, versionID, versionStage string) (*secretsmanager.GetSecretValueInput, error) {
	input := &secretsmanager.GetSecretValueInput{SecretId: aws.String(secretName)}
	if versionID == "" && versionStage == "" {
		return nil, errors.New("either secret version id or secret version stage must be specified")
	}
	if versionID != "" && versionStage != "" {
		return nil, errors.New("secret version id and secret version stage cannot both be specified")
	}
	if versionID != "" {
		input.VersionId = aws.String(versionID)
	}
	if versionStage != "" {
		input.VersionStage = aws.String(versionStage)
	}
	return input, nil
}
