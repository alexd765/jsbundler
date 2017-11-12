package callmap

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	f, err := newFile("./testfiles/mult.js")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	t.Logf(": %+v", f)
}
