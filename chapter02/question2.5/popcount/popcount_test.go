package popcount

import "testing"

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
