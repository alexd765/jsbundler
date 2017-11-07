package callmap

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	_, err := newFile("./testfiles/mult.js")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
