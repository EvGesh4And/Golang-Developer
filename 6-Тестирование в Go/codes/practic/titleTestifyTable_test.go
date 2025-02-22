package practic_test

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/require"
)

// TitleCase(str, minor) returns a str string with all words capitalized except minor words.
// The first word is always capitalized.
//
// E.g.
// TitleCase("the quick fox in the bag", "") = "The Quick Fox In The Bag"
// TitleCase("the quick fox in the bag", "in the") = "The Quick Fox in the Bag"

// Задание
// 1. Дописать существующие тесты.
// 2. Придумать один новый тест.

func TestTestifyTable(t *testing.T) {
	test := []struct {
		str      string
		minor    string
		expected string
	}{
		{"da pa", "", "Da Pa"},
		{"trust ther bust", "trust", "Trust Ther Bust"},
		{"BUBA 454dsad", "", "Buba 454dsad"},
	}

	for _, tc := range test {
		t.Run(tc.str, func(t *testing.T) {
			got := titlecase.TitleCase(tc.str, tc.minor)
			require.Equal(t, tc.expected, got)
		})
	}
}
