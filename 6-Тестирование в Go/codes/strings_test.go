package main_test

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	// Инициализация
	const s, sub, want = "Context", "text", 3

	// Получение ответа
	got := strings.Index(s, sub)

	if got != want {
		t.Errorf("Index(%q, %q) = %v; want %v", s, sub, got, want)
	}
}
