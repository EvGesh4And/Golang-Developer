package practic_test

import (
	"testing"

	"github.com/kulti/titlecase"
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

func TestEmpty(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}
}

func TestWithoutMinor(t *testing.T) {
	const str, minor, want = "CAPITAN THE KATALKIN", "the", "Capitan Katalkin"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v)=%v; want=%v", str, minor, got, want)
	}
}

func TestWithMinorInFirst(t *testing.T) {
	const str, minor, want = "capitan katalkin", "capitan", "Capitan Katalkin"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v)=%v; want=%v", str, minor, got, want)
	}
}
