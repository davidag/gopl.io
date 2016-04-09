// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"gopl.io/ch2/popcount"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShifting2(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClearing2(0x1234567890ABCDEF)
	}
}

func TestPopCounts(t *testing.T) {
	var v uint64 = 0x1234567890ABCDEF
	pc := popcount.PopCount(v)
	pcbs2 := popcount.PopCountByShifting2(v)
	pcbc2 := popcount.PopCountByClearing2(v)
	if pc != pcbs2 {
		t.Errorf("PopCount = %v, PopCountByShifting2 = %v",
			pc, pcbs2)
	}
	if pc != pcbc2 {
		t.Errorf("PopCount = %v, PopCountByClearing2 = %v",
			pc, pcbc2)
	}
}

// 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// testing: warning: no tests to run
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// 1.3GHz Intel Core2 Duo U7300
// $ go test -cpu=2 -bench=. gopl.io/ch2/popcount
// PASS
// BenchmarkPopCount-2             100000000           14.2 ns/op
// BenchmarkPopCountLoop-2         30000000            39.3 ns/op
// BenchmarkBitCount-2             200000000           10.0 ns/op
// BenchmarkPopCountByClearing-2   10000000           125 ns/op
// BenchmarkPopCountByShifting-2   2000000            660 ns/op
// BenchmarkPopCountByShifting2-2  5000000            380 ns/op
// BenchmarkPopCountByClearing2-2  10000000           125 ns/op
// ok      gopl.io/ch2/popcount    12.729s
