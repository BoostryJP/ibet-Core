package nodekey

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/cmd/utils/nodekey/constants"
	"github.com/ethereum/go-ethereum/cmd/utils/nodekey/decrypter"
	"github.com/ethereum/go-ethereum/cmd/utils/nodekey/fetcher"
)

type keyFetcher interface {
	FetchNodeKey() (*ecdsa.PrivateKey, error)
	FetchEncryptedNodeKey() (string, error)
}

type keyDecrypter interface {
	DecryptNodeKey(base64StringData string) (*ecdsa.PrivateKey, error)
}

type Manager struct {
	fetcher    keyFetcher
	decrypter  keyDecrypter
	decryption string
}

func NewManager(source, decryptionScheme string, config []byte) (*Manager, error) {
	var (
		fetch keyFetcher
		dec   keyDecrypter
		err   error
	)

	// Select how encrypted/plaintext node key data is fetched.
	switch source {
	case constants.SourceAwsSm:
		fetch, err = fetcher.NewNodeKeyAwsSecretsManagerFetcher(config)
	default:
		return nil, fmt.Errorf("unsupported source type %q", source)
	}
	if err != nil {
		return nil, err
	}

	// Select how fetched data is turned into an ECDSA private key.
	switch decryptionScheme {
	case constants.DecryptionNone:
		// no-op; plaintext key is expected in the secret
	case constants.DecryptionAwsKms:
		dec, err = decrypter.NewNodeKeyAwsKmsDecrypter(config)
	default:
		return nil, fmt.Errorf("invalid decryption scheme %q", decryptionScheme)
	}
	if err != nil {
		return nil, err
	}

	return &Manager{fetcher: fetch, decrypter: dec, decryption: decryptionScheme}, nil
}

func (m *Manager) GetNodeKey() (*ecdsa.PrivateKey, error) {
	if m.decryption == constants.DecryptionNone {
		// Secret already contains a hex private key.
		return m.fetcher.FetchNodeKey()
	}
	// Secret contains encrypted bytes (base64) that must be KMS-decrypted first.
	encryptedData, err := m.fetcher.FetchEncryptedNodeKey()
	if err != nil {
		return nil, err
	}
	return m.decrypter.DecryptNodeKey(encryptedData)
}
