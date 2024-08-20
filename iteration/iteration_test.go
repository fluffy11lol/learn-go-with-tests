package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat1("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Got %s, expected %s", repeated, expected)
	}
}
func BenchmarkRepeat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat1("a", 50)
	}
}
func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a", 50)
	}
}
func BenchmarkRepeat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat3("a", 50)
	}
}
func BenchmarkRepeat4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat4("a", 50)
	}
}
