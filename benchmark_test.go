package patricia_test

import (
	"testing"

	. "github.com/gbrlsnchs/patricia"
)

func BenchmarkLongString(b *testing.B) {
	b.ReportAllocs()

	t := New("BenchmarkLongString")

	t.Add("This is a very, very long string, so let's benchmark it.", "foo")

	for i := 0; i < b.N; i++ {
		_ = t.Get("This is a very, very long string, so let's benchmark it.")
	}
}

func BenchmarkManyWords(b *testing.B) {
	b.ReportAllocs()

	t := New("BenchmarkManyWords")

	t.Add("romane", 1)
	t.Add("romanus", 2)
	t.Add("romulus", 3)
	t.Add("rubens", 4)
	t.Add("ruber", 5)
	t.Add("rubicon", 6)
	t.Add("rubicundus", 7)

	for i := 0; i < b.N; i++ {
		_ = t.Get("romanus")
	}
}
