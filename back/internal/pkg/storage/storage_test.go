package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnection(t *testing.T) {
	store, err := New("postgresql://localhost:5433/sport?user=web101&password=123")
	require.NoError(t, err)
	require.NotEqual(t, Gres{}, store)
}

func TestPingConnection(t *testing.T) {
	store, err := New("postgresql://localhost:5433/sport?user=web101&password=123")
	require.NoError(t, err)
	store.Pool.Ping(context.Background())
	require.NoError(t, err)
}
