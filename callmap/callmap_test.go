package callmap

import (
	"testing"
)

func TestAddFile(t *testing.T) {
	multJS := "./testfiles/mult.js"
	cm := New()

	if err := cm.AddFile(multJS); err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	if s := len(cm.files[multJS].functions); s != 2 {
		t.Errorf("got %d functions; want '2'", s)
	}
	if _, ok := cm.files[multJS].functions["plus"]; !ok {
		t.Errorf("got %+v; want 'plus' in map", cm.files[multJS].functions)
	}
	if _, ok := cm.files[multJS].functions["mult"]; !ok {
		t.Errorf("got %+v; want 'mult' in map", cm.files[multJS].functions)
	}
}
