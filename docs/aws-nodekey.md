# AWS KMS/Secrets Manager based node key management

Our project supports loading the P2P node key from AWS Secrets Manager, with optional decryption by AWS KMS. This allows for secure management of node keys without storing them directly on disk.

## CLI options

- --nodekeysource: Node key source (file or aws-sm)
- --nodekeydecryption: Decryption scheme (none or aws-kms)

Default values are:

- --nodekeysource=file
- --nodekeydecryption=none

When --nodekey or --nodekeyhex is explicitly specified, those legacy options take precedence.

## TOML configuration

Add the following section to your config.toml:

```toml
[Node.P2P.NodeKey.ConfigAws]
SecretName = "boostry/ibet-network/quorum/nodekey"
SecretVersionStage = "AWSCURRENT"
KmsKeyId = "alias/mykey"
KmsEncryptionAlgorithm = "RSAES_OAEP_SHA_256"
```

Field details:

- SecretName: Name or ARN of the secret that stores node key data
- SecretVersionId: Optional AWS secret version ID
- SecretVersionStage: Optional AWS secret version stage (for example AWSCURRENT)
- Either SecretVersionId or SecretVersionStage must be specified
- SecretVersionId and SecretVersionStage cannot both be specified
- KmsKeyId: Required when --nodekeydecryption=aws-kms
- KmsEncryptionAlgorithm: Optional KMS algorithm override for asymmetric keys

## Startup examples

Use plain hex private key stored in Secrets Manager:

```bash
./build/bin/geth --config ./config.toml --nodekeysource aws-sm --nodekeydecryption none
```

Use encrypted key stored in Secrets Manager and decrypt with KMS:

```bash
./build/bin/geth --config ./config.toml --nodekeysource aws-sm --nodekeydecryption aws-kms
```

## AWS authentication

AWS SDK v2 default credential chain is used. Typical runtime environment variables:

- AWS_REGION
- AWS_ROLE_ARN and AWS_WEB_IDENTITY_TOKEN_FILE (for IRSA/OIDC)

Ensure the IAM role has permissions for:

- secretsmanager:GetSecretValue
- kms:Decrypt (when using aws-kms)
