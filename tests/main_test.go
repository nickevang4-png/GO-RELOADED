package tests

import (
	"GO-RELOADED/internal/textops"
	"testing"
)

func TestHexConversion(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"
	output := textops.Process(input)
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}
