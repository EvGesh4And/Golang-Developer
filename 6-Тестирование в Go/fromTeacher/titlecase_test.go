package main

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

// func TestEmpty(t *testing.T) {
// 	const str, minor, want = "", "", ""
// 	got := titlecase.TitleCase(str, minor)
// 	assert.Equal(t, want, got)
// }

// func TestRussian(t *testing.T) {
// 	const str, minor, want = "привет медвед", "", "Привет Медвед"
// 	got := titlecase.TitleCase(str, minor)
// 	assert.Equal(t, want, got)
// }

// func TestCaps(t *testing.T) {
// 	const str, minor, want = "THE QUICK FOX IN THE BAG", "", "The Quick Fox In The Bag"
// 	got := titlecase.TitleCase(str, minor)
// 	assert.Equal(t, want, got)
// }

// func TestWithMinor(t *testing.T) {
// 	const str, minor, want = "the quick fox in the the bag", "the", "The Quick Fox In the the Bag"
// 	got := titlecase.TitleCase(str, minor)
// 	assert.Equal(t, want, got)
// }

func TestTitleCase(t *testing.T) {

	CommonUtil()

	testCases := []struct {
		desc  string
		str   string
		minor string
		want  string
	}{
		{
			desc:  "empty",
			str:   "",
			minor: "",
			want:  "",
		},
		{
			desc:  "rus",
			str:   "привет медвед",
			minor: "",
			want:  "Привет Медвед",
		},
		{
			desc:  "caps",
			str:   "THE QUICK FOX IN THE BAG",
			minor: "",
			want:  "The Quick Fox In The Bag",
		},
		{
			desc:  "minor",
			str:   "the quick fox in the the bag",
			minor: "the",
			want:  "The Quick Fox In the the Bag",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := titlecase.TitleCase(tC.str, tC.minor)
			assert.Equal(t, tC.want, got)
		})
	}
}
