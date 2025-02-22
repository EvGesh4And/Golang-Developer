package practic_test

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/assert"
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

func TestEmptyTestify(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
}

func TestWithoutMinorTestify(t *testing.T) {
	const str, minor, want = "CAPITAN THE KATALKIN", "the", "Capitan Katalkin"
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
}

func TestWithMinorInFirstTestify(t *testing.T) {
	const str, minor, want = "capitan katalkin", "capitan", "Capitan Katalkin"
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
}
