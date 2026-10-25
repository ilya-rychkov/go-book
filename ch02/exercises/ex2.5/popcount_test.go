package main

import "testing"

var result int

func BenchmarkPopCountKernighan(b *testing.B) {
	var r int
	x := uint64(10)
	for i := 0; i < b.N; i++ {
		r = PopCountKernighan(x)
	}
	result = r
}

func BenchmarkPopCount(b *testing.B) {
	var r int
	x := uint64(10)
	for i := 0; i < b.N; i++ {
		r = PopCount(x)
	}
	result = r
}

func BenchmarkCountOnesIn64Bits(b *testing.B) {
	var r int
	x := uint64(10)
	for i := 0; i < b.N; i++ {
		r = CountOnesIn64Bits(x)
	}
	result = r
}

func BenchmarkPopCountBytes(b *testing.B) {
	var r int
	x := uint64(10)
	for i := 0; i < b.N; i++ {
		r = PopCountBytes(x)
	}
	result = r
}
