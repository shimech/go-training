package main

import "testing"

var testArguments = []string{"hoge", "fuga", "piyo"}

func BenchmarkStringifyArgsWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringifyArgsWithLoop(testArguments)
	}
}

func BenchmarkStringifyArgsWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringifyArgsWithJoin(testArguments)
	}
}
