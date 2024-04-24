package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
