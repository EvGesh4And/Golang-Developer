package practic_test

import (
	"simpletest/practic"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	const ch, want = 2, 4
	t.Run("chislo: "+strconv.Itoa(ch), func(t *testing.T) {
		got, err := practic.FuncGetSq(ch)
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}

func TestInvalid(t *testing.T) {
	const ch = -3
	t.Run("chislo: "+strconv.Itoa(ch), func(t *testing.T) {
		_, err := practic.FuncGetSq(ch)
		require.Error(t, err)
	})
}
