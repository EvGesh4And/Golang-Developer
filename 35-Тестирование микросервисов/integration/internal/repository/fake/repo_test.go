package fake

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOne(t *testing.T) {
	require.Equal(t, 2, 1+1)
}
