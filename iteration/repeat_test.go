package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func TestRepeatBuiltIn(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
