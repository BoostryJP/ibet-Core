package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/ethereum/go-ethereum/p2p/nodekey"
	"github.com/naoina/toml"
)

type AwsClient struct {
	Config        nodekey.AwsConfig
	SecretsClient *secretsmanager.Client
	KMSClient     *kms.Client
}

func NewAwsClient(configBytes []byte) (*AwsClient, error) {
	ctx := context.Background()

	// Parse only the AWS-related nodekey sub-config from TOML bytes.
	var cfg nodekey.AwsConfig
	if err := toml.Unmarshal(configBytes, &cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration passed: %w", err)
	}
	if err := validateConfigurationValues(cfg); err != nil {
		return nil, err
	}

	// Credentials and region are resolved by the AWS default provider chain.
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &AwsClient{
		Config:        cfg,
		SecretsClient: secretsmanager.NewFromConfig(awsConfig),
		KMSClient:     kms.NewFromConfig(awsConfig),
	}, nil
}

func validateConfigurationValues(config nodekey.AwsConfig) error {
	if config.SecretName == "" {
		return errors.New("need to specify secret name to retrieve data from AWS Secrets Manager")
	}
	return nil
}
