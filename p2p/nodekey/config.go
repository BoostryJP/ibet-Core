package nodekey

type NodeKeyConfig struct {
	ConfigFile FileConfig `toml:",omitempty"`
	ConfigAws  AwsConfig  `toml:",omitempty"`
}

type FileConfig struct {
	Hex  string `toml:",omitempty"`
	File string `toml:",omitempty"`
}

type AwsConfig struct {
	SecretName             string `toml:",omitempty"`
	SecretVersion          string `toml:",omitempty"`
	KmsKeyId               string `toml:",omitempty"`
	KmsEncryptionAlgorithm string `toml:",omitempty"`
}
