package callmap

import "testing"

func TestAddFile(t *testing.T) {
	cm := New()
	if err := cm.AddFile("./testfiles/mult.js"); err != nil {
		t.Errorf("unexpected err: %s", err)
	}
}
