package main

import (
	"testing"
)

func TestParseRecord(t *testing.T) {
	record := []string{"452600", "8.3252", "41", "880", "129", "322", "126"}
	expected := House{
		Value:    452600,
		Income:   8.3252,
		Age:      41,
		Rooms:    880,
		Bedrooms: 129,
		Pop:      322,
		HH:       126,
	}

	result, err := parseRecord(record)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %+v, got %+v", expected, result)
	}
}
