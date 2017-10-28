package callmap

import "testing"

func TestAddFile(t *testing.T) {
	multJS := "./testfiles/mult.js"
	cm := New()

	if err := cm.AddFile(multJS); err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	if s := len(cm.files[multJS].functions); s != 2 {
		t.Errorf("got %d functions; want 2", s)
	}
}
