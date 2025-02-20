package main_test

import (
	main1 "codes"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	const s, sub, want = "chicken", "ken", 3
	got := strings.Index(s, sub)

	if got != want {
		t.Fatalf("Index(%q,%q) = %v; want %v", s, sub, got, want)
	}
}

func TestGetTrue(t *testing.T) {
	const want = true
	got := main1.GetTrue()

	if got {
		t.Errorf("GetTrue() = %v; want %v", got, want)
	}
}
