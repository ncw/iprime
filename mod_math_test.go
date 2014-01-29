// Test the modular arithmetic

package main

import (
	"math/big"
	"testing"
)

// FIXME try coverage on this to see whether we are testing everything!

var (

	// Interesting halves to provoke carries and overflows in MOD_P arithmetic
	interesting_halves = []uint32{
		1, 2, 0xFFFFFFFF, 0xFFFFFFFE, 0xFFFFFFFD,
		0x80000000, 0x80000001, 0x80000002, 0x7FFFFFFF, 0x7FFFFFFE,
		0x12345678, 0x89ABCDEF, 0xFEDCBA98, 0x87654321, 0x55555555, 0xAAAAAAAA,
		0,
	}

	// List of numbers to try operations
	numbers []uint64

	// MOD_P in big.Int format
	P = new(big.Int).SetUint64(MOD_P)
)

func init() {
	MakeNumbers()
}

// Make a list of numbers to try out of the interesting halves
func MakeNumbers() {
	if numbers != nil {
		return
	}
	numbers = make([]uint64, 0, len(numbers)*len(numbers))
	for _, a := range interesting_halves {
		for _, b := range interesting_halves {
			n := (uint64(a) << 32) + uint64(b)
			if n < MOD_P {
				numbers = append(numbers, n)
			}
		}
	}
}

func TestModAdd(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for _, b := range numbers {
			B.SetUint64(b)
			c := mod_add(a, b)
			C.Add(A, B)
			C.Mod(C, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("%d + %d: Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func BenchmarkModAdd(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	y := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_add(x, y)
	}
}

func TestModAdc(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	Carry := new(big.Int)
	Mask := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for _, b := range numbers {
			B.SetUint64(b)
			for width := uint8(1); width < 64; width++ {
				carry := b
				c := mod_adc(a, width, &carry)
				C.Add(A, B)
				Carry.Rsh(C, uint(width))
				expectedCarry := Carry.Uint64()
				Mask.SetUint64(uint64(1)<<width - 1)
				C.And(C, Mask)
				expected := C.Uint64()
				if c != expected || expectedCarry != carry {
					t.Fatalf("adc(%d,%d,%d): Expecting %d carry %d but got %d carry %d", a, b, width, expected, expectedCarry, c, carry)
				}
			}
		}
	}
}

func TestModSub(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for _, b := range numbers {
			B.SetUint64(b)
			c := mod_sub(a, b)
			C.Sub(A, B)
			C.Mod(C, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("%d - %d: Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func BenchmarkModSub(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	y := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_sub(x, y)
	}
}

func TestModReduce(t *testing.T) {
	MakeNumbers()
	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for _, b := range numbers {
			B.SetUint64(b)
			c := mod_reduce(a, b)
			C.Lsh(A, 64)
			C.Add(C, B)
			C.Mod(C, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("mod_reduce(%d,%d): Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func TestModMul(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for _, b := range numbers {
			B.SetUint64(b)
			c := mod_mul(a, b)
			C.Mul(A, B)
			C.Mod(C, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("%d * %d: Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func BenchmarkModMul(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	y := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_mul(x, y)
	}
}

func TestModSqr(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		b := mod_sqr(a)
		B.Mul(A, A)
		B.Mod(B, P)
		expected := B.Uint64()
		if b != expected {
			t.Fatalf("%d**2: Expecting %d but got %d", a, expected, b)
		}
	}
}

func BenchmarkModSqr(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_sqr(x)
	}
}

func TestModShift(t *testing.T) {
	A := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for b := uint8(0); b < 96; b++ {
			c := mod_shift(a, b)
			C.Lsh(A, uint(b))
			C.Mod(C, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("%d >> %d: Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func TestModShifts(t *testing.T) {
	A := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for b := uint8(3); b < 96; b += 3 {
			var c uint64
			switch b {
			case 3:
				//c = mod_shift3(a)
				c = mod_shift0to31x3(a)
			case 6:
				c = mod_shift6(a)
			case 9:
				c = mod_shift9(a)
			case 12:
				c = mod_shift12(a)
			case 15:
				c = mod_shift15(a)
			case 18:
				c = mod_shift18(a)
			case 21:
				c = mod_shift21(a)
			case 24:
				c = mod_shift24(a)
			case 27:
				c = mod_shift27(a)
			case 30:
				c = mod_shift30(a)
			case 33:
				c = mod_shift33(a)
			case 36:
				//c = mod_shift36(a)
				c = mod_shift32to63x36(a)
			case 39:
				c = mod_shift39(a)
			case 42:
				c = mod_shift42(a)
			case 45:
				c = mod_shift45(a)
			case 48:
				c = mod_shift48(a)
			case 51:
				c = mod_shift51(a)
			case 54:
				c = mod_shift54(a)
			case 57:
				c = mod_shift57(a)
			case 60:
				c = mod_shift60(a)
			case 63:
				c = mod_shift63(a)
			case 66:
				c = mod_shift66(a)
			case 69:
				c = mod_shift69(a)
			case 72:
				c = mod_shift72(a)
				c = mod_shift64to95x72(a)
			case 75:
				c = mod_shift75(a)
			case 78:
				c = mod_shift78(a)
			case 81:
				c = mod_shift81(a)
			case 84:
				c = mod_shift84(a)
			case 87:
				c = mod_shift87(a)
			case 90:
				c = mod_shift90(a)
			case 93:
				c = mod_shift93(a)
			}
			C.Lsh(A, uint(b))
			C.Mod(C, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("%d >> %d: Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func BenchmarkModShift3(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_shift0to31x3(x)
	}
}

func BenchmarkModShift36(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_shift32to63x36(x)
	}
}

func BenchmarkModShift72(b *testing.B) {
	b.StopTimer()
	x := mod_rnd()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		x = mod_shift64to95x72(x)
	}
}

func TestModPow(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	C := new(big.Int)
	for _, a := range numbers {
		A.SetUint64(a)
		for _, b := range numbers {
			B.SetUint64(b)
			c := mod_pow(a, b)
			C.Exp(A, B, P)
			expected := C.Uint64()
			if c != expected {
				t.Fatalf("%d ** %d: Expecting %d but got %d", a, b, expected, c)
			}
		}
	}
}

func TestModInv(t *testing.T) {
	A := new(big.Int)
	B := new(big.Int)
	for _, a := range numbers {
		if a == 0 {
			continue
		}
		A.SetUint64(a)
		b := mod_inv(a)
		B.ModInverse(A, P)
		expected := B.Uint64()
		if b != expected {
			t.Fatalf("inv(%d): Expecting %d but got %d", a, expected, b)
		}
	}
}

func TestModVectorMul(t *testing.T) {
	// FIXME
}

func TestModVectorSqr(t *testing.T) {
	// FIXME
}

func TestModRnd(t *testing.T) {
	ones := uint64(0)
	zeros := uint64((1 << 64) - 1)
	for i := 0; i < 1000; i++ {
		x := mod_rnd()
		if x > MOD_P {
			t.Fatal("too big")
		}
		ones |= x
		zeros &= x
	}
	if ones != uint64((1<<64)-1) {
		t.Fatal("Didn't set all bits")
	}
	if zeros != 0 {
		t.Fatal("Didn't clear all bits")
	}
}
