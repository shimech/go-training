package popcount

import "testing"

var tests = []struct {
	input uint64
	want  int
}{
	{0, 0},
	{1, 1},
	{3, 2},
}

func TestPopCountLogicalAndDecrement(t *testing.T) {
	for _, test := range tests {
		if got := PopCountLogicalAndDecrement(test.input); got != test.want {
			t.Errorf("PopCountLogicalAndDecrement(%d) = %d", test.input, got)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLogicalAndDecrement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLogicalAndDecrement(uint64(i))
	}
}
