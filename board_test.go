package gochan

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBoards(t *testing.T) {
	boards, err := GetBoards()

	require.NoError(t, err)
	require.NotNil(t, boards)
}
