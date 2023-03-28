package pack

import "testing"

func TestGet1(t *testing.T) {
	if Get1() != 1 {
		t.Errorf("Not 1")
	}
}
