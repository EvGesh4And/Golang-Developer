package main

import (
	"strconv"
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/assert"
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

func TestAtoi(t *testing.T) {
	const str, want = "43", "42"
	got, err := strconv.Atoi(str)
	require.NoError(t, err)
	assert.Equal(t, want, got)

	// assert.Equal(t, want, got)
}

func TestEmpty(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
	require.Equal(t, want, got)
	// assert.Equal(t, want, got)
}

func TestRussian(t *testing.T) {
	const str, minor, want = "привет медвед", "", "Привет Медвед"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("got is not want")
	}
	// assert.Equal(t, want, got)
}

func TestWithoutMinor(t *testing.T) {
	panic("implement me")
}

func TestWithMinorInFirst(t *testing.T) {
	panic("implement me")
}

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
