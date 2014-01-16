// Test the transpose routines

package main

import (
	"testing"
)

func TestTransposeSquareSlow2(t *testing.T) {
	x := []uint64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	y := []uint64{
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	}
	TransposeSquareSlow(x, nil, 2)
	for i := range x {
		if x[i] != y[i] {
			t.Fatal("Failed %d: expecting %d got %d\n", i, x[i], y[i])
		}
	}

}

func TestTransposeSlow2_2(t *testing.T) {
	x := []uint64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	scratch := make([]uint64, 16)
	y := []uint64{
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	}
	x, scratch = TransposeSlow(x, scratch, 2, 2)
	for i := range x {
		if x[i] != y[i] {
			t.Fatal("Failed %d: expecting %d got %d\n", i, x[i], y[i])
		}
	}

}

func TestTransposeSlow2_3(t *testing.T) {
	x := []uint64{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32,
	}
	scratch := make([]uint64, 32)
	y := []uint64{
		1, 9, 17, 25,
		2, 10, 18, 26,
		3, 11, 19, 27,
		4, 12, 20, 28,
		5, 13, 21, 29,
		6, 14, 22, 30,
		7, 15, 23, 31,
		8, 16, 24, 32,
	}
	x, scratch = TransposeSlow(x, scratch, 2, 3)
	for i := range x {
		if x[i] != y[i] {
			t.Fatal("Failed %d: expecting %d got %d\n", i, x[i], y[i])
		}
	}
	x, scratch = TransposeSlow(x, scratch, 3, 2)
	for i := range x {
		if x[i] != uint64(i)+1 {
			t.Fatal("Failed %d: expecting %d got %d\n", i, x[i], y[i])
		}
	}
}

func testTransposeFastN(t *testing.T, log_side uint8) {
	n := uint(1) << (2 * log_side)
	x := make([]uint64, n)
	scratch := make([]uint64, n)
	for i := range x {
		x[i] = uint64(i)
	}
	TransposeSquareFast(x, scratch, log_side)
	if x[0] != 0 || x[1] != 1<<log_side || x[1<<log_side] != 1 {
		t.Fatalf("Transpose failed x[0]=%d, x[1]=%d, x[%d]=%d", x[0], x[1], 1<<log_side, x[1<<log_side])
	}
	TransposeSquareSlow(x, scratch, log_side)
	for i := range x {
		if x[i] != uint64(i) {
			t.Fatalf("Failed x[%d]: expecting %d got %d\n", i, i, x[i])
		}
	}
}

func TestTransposeFast5(t *testing.T) {
	testTransposeFastN(t, 5)
}

func TestTransposeFast6(t *testing.T) {
	testTransposeFastN(t, 6)
}

func TestTransposeFast7(t *testing.T) {
	testTransposeFastN(t, 7)
}

func TestTransposeFast8(t *testing.T) {
	testTransposeFastN(t, 8)
}

func TestTransposeFast9(t *testing.T) {
	testTransposeFastN(t, 9)
}

func BenchmarkTransposeSlow10(b *testing.B) {
	b.StopTimer()
	n := uint(1) << (2 * 10)
	x := make([]uint64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		TransposeSquareSlow(x, nil, 10)
	}
}

func BenchmarkTransposeFast10(b *testing.B) {
	b.StopTimer()
	n := uint(1) << (2 * 10)
	x := make([]uint64, n)
	scratch := make([]uint64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		TransposeSquareFast(x, scratch, 10)
	}
}
