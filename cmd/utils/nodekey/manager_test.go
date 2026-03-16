package nodekey

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewManager_InvalidSource(t *testing.T) {
	_, err := NewManager("invalid", "none", []byte(""))
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported source type")
}
