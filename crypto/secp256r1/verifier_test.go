package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func TestVerify(t *testing.T) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	msg := []byte("quorum-secp256r1")
	hash := sha256.Sum256(msg)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		t.Fatal(err)
	}
	if !Verify(hash[:], r, s, priv.X, priv.Y) {
		t.Fatal("expected valid signature")
	}
	if Verify(hash[:], r, s, priv.X, nil) {
		t.Fatal("expected invalid signature for nil public key component")
	}
}
