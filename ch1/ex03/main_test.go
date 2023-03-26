package main

import "testing"

func BenchmarkEchoBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoBad()
	}
}

func BenchmarkEchoGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoGood()
	}
}
