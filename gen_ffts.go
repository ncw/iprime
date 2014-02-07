// +build ignore

// FIXME factor add/sub/mul/shift out into butterfly
// which will stop excessive inlining and maybe be quicker
// x[1], x[2] = butterfly(x[1], x[2], w)
// x[1], x[2] = butterfly_shift(x[1], x[2], 5)
// x[1], x[2] = butterfly_shift5(x[1], x[2])

// FIXME specialise shifts by generating code for them mod_shift5
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
)

// Generate ffts.go like this
const REGENERATE = "go run mod_math.go mod_math_noasm.go gen_ffts.go"

var Data = struct {
	FftSizes   []int
	Shifts     []int
	Regenerate string
}{
	FftSizes:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, //11, 12}
	Shifts:     []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93},
	Regenerate: REGENERATE,
}

var PowersOfTwo = make([]uint64, 193)

func init() {
	x := uint64(1)
	for i := range PowersOfTwo {
		PowersOfTwo[i] = x
		x = mod_mul(x, uint64(2))
	}
}

// Read a twiddle and work out what sort of multiply we need to do
//
// Returns 0 for null, -1 for full multiply or +ve for shift
func multiplyType(w uint64) int {
	if w == 0 {
		panic("Butterflies can't be 0")
	}
	if w == 1 {
		return 0
	}
	for i, x := range PowersOfTwo {
		if x == w {
			return i
		}
	}
	return -1
}

type ButterflyFn func(a, b uint, w uint64)

func MakeFftGeneric(_log_n int, butterfly ButterflyFn) {
	log_n := uint8(_log_n)
	if log_n == 0 {
		return
	}
	n := uint(1) << log_n
	mod_w := mod_pow(7, (MOD_P-1)/uint64(n))
	mod_w = mod_pow(mod_w, 5)
	var d uint64 = mod_w

	for k := log_n; k >= 1; k-- {
		m := uint(1) << k
		c := m >> 1
		var w uint64 = 1
		for j := uint(0); j < c; j++ {
			for r := uint(0); r < n; r += m {
				a := r + j
				b := a + c
				butterfly(a, b, w)
			}
			w = mod_mul(w, d)
		}
		d = mod_mul(d, d)
	}
}

func MakeInvFftGeneric(_log_n int, butterfly ButterflyFn) string {
	log_n := uint8(_log_n)
	if log_n == 0 {
		return ""
	}
	n := uint(1) << log_n
	mod_w := mod_pow(7, (MOD_P-1)/uint64(n))
	mod_w = mod_pow(mod_w, 5)
	mod_invw := mod_inv(mod_w)

	out := new(bytes.Buffer)
	for k := uint8(1); k <= log_n; k++ {
		m := uint(1) << k
		c := m >> 1
		z := uint64(1 << (log_n - k))
		d := mod_pow(mod_invw, z)
		w := uint64(1)
		for j := uint(0); j < c; j++ {
			for r := uint(0); r < n; r += m {
				a := r + j
				b := a + c
				butterfly(a, b, w)
			}
			w = mod_mul(w, d)
		}
	}
	return out.String()
}

func MakeFft(_log_n int) string {
	out := new(bytes.Buffer)
	MakeFftGeneric(_log_n, func(a, b uint, w uint64) {
		switch i := multiplyType(w); i {
		case 0:
			fmt.Fprintf(out, "x[%d], x[%d] = butterfly_null(x[%d], x[%d])\n", a, b, a, b)
		case -1:
			fmt.Fprintf(out, "x[%d], x[%d] = butterfly_mul(x[%d], x[%d], %d)\n", a, b, a, b, w)
		default:
			fmt.Fprintf(out, "x[%d], x[%d] = butterfly_shift%d(x[%d], x[%d])\n", a, b, i, a, b)
		}
	})
	return out.String()
}

func MakeInvFft(_log_n int) string {
	out := new(bytes.Buffer)
	MakeInvFftGeneric(_log_n, func(a, b uint, w uint64) {
		switch i := multiplyType(w); i {
		case 0:
			fmt.Fprintf(out, "x[%d], x[%d] = invbutterfly_null(x[%d], x[%d])\n", a, b, a, b)
		case -1:
			fmt.Fprintf(out, "x[%d], x[%d] = invbutterfly_mul(x[%d], x[%d], %d)\n", a, b, a, b, w)
		default:
			// Since 2^96 mod p = -1, we subtract -1 and use a negative butterfly
			fmt.Fprintf(out, "x[%d], x[%d] = invbutterfly_shift%d(x[%d], x[%d])\n", a, b, i-96, a, b)
		}
	})
	return out.String()
}

// Make ffts.go
func generateFfts() {
	t := template.Must(template.New("main").Funcs(template.FuncMap{
		"Fft":    MakeFft,
		"InvFft": MakeInvFft,
	}).Parse(program))

	cmd := exec.Command("gofmt")

	out, err := os.Create("ffts.go")
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = out

	gofmt, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err = t.Execute(gofmt, Data); err != nil {
		log.Fatal(err)
	}
	if err = gofmt.Close(); err != nil {
		log.Fatal(err)
	}
	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	if err = out.Close(); err != nil {
		log.Fatal(err)
	}
}

var program = `
// Automatically generated - DO NOT EDIT
// Regenerate with: {{ .Regenerate }}

// +build !amd64,!arm

// Unrolled FFTs

package main

func butterfly_null(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_sub(a, b)
	return
}

{{ range .Shifts }}
func butterfly_shift{{.}}(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift{{.}}(mod_sub(a, b))
	return
}
{{ end }}

func butterfly_mul(a, b uint64, w uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_mul(mod_sub(a, b), w)
	return
}

func invbutterfly_null(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_sub(a, b)
	return
}

{{ range .Shifts }}
// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift{{.}}(a, b uint64) (u, v uint64) {
	b = mod_shift{{.}}(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}
{{ end }}

func invbutterfly_mul(a, b uint64, w uint64) (u, v uint64) {
	b = mod_mul(b, w)
	u = mod_add(a, b)
	v = mod_sub(a, b)
	return
}


{{ range $size := .FftSizes }}
// Fft for size 2**{{$size}}
//
// This is an in place FFT with a bit reversed output
func fft{{$size}}(x []uint64) {
{{ Fft $size }}
}

// InvFft for size 2**{{$size}}
//
// This is an in place Inverse FFT with a bit reversed input
func invfft{{$size}}(x []uint64) {
{{ InvFft $size }}
}
{{ end }}
`

func MakeAmd64Fft(_log_n int) string {
	out := new(bytes.Buffer)
	MakeFftGeneric(_log_n, func(a, b uint, w uint64) {
		fmt.Fprintf(out, "\tMOVQ (%d*8)(R12), AX\n", a)
		fmt.Fprintf(out, "\tMOVQ (%d*8)(R12), BX\n", b)
		switch i := multiplyType(w); i {
		case 0:
			fmt.Fprintf(out, "\tCALL ·butterfly_null(SB)\n")
		case -1:
			fmt.Fprintf(out, "\tMOVQ $0x%016X, DI\n", w)
			fmt.Fprintf(out, "\tCALL ·butterfly_mul(SB)\n")
		default:
			fmt.Fprintf(out, "\tCALL ·butterfly_shift%d(SB)\n", i)
		}
		fmt.Fprintf(out, "\tMOVQ CX, (%d*8)(R12)\n", a)
		fmt.Fprintf(out, "\tMOVQ DX, (%d*8)(R12)\n", b)
	})
	return out.String()
}

func MakeAmd64InvFft(_log_n int) string {
	out := new(bytes.Buffer)
	MakeInvFftGeneric(_log_n, func(a, b uint, w uint64) {
		fmt.Fprintf(out, "\tMOVQ (%d*8)(R12), AX\n", a)
		fmt.Fprintf(out, "\tMOVQ (%d*8)(R12), BX\n", b)
		switch i := multiplyType(w); i {
		case 0:
			fmt.Fprintf(out, "\tCALL ·invbutterfly_null(SB)\n")
		case -1:
			fmt.Fprintf(out, "\tMOVQ $0x%016X, DI\n", w)
			fmt.Fprintf(out, "\tCALL ·invbutterfly_mul(SB)\n")
		default:
			// Since 2^96 mod p = -1, we subtract -1 and use a negative butterfly
			fmt.Fprintf(out, "\tCALL ·invbutterfly_shift%d(SB)\n", i-96)
		}
		fmt.Fprintf(out, "\tMOVQ CX, (%d*8)(R12)\n", a)
		fmt.Fprintf(out, "\tMOVQ DX, (%d*8)(R12)\n", b)
	})
	return out.String()
}

// Make ffts_amd64.go
func generateAmd64Ffts() {
	t := template.Must(template.New("main").Funcs(template.FuncMap{
		"Fft":    MakeAmd64Fft,
		"InvFft": MakeAmd64InvFft,
	}).Parse(amd64Assembler))

	out, err := os.Create("ffts_amd64.h")
	if err != nil {
		log.Fatal(err)
	}
	if err = t.Execute(out, Data); err != nil {
		log.Fatal(err)
	}
	if err = out.Close(); err != nil {
		log.Fatal(err)
	}
}

var amd64Assembler = `
// Automatically generated - DO NOT EDIT
// Regenerate with: {{ .Regenerate }}

// Non golang calling convention here for the butterflies
// Input in AX, BX
// Output in CX, DX
// Twiddle in DI if present
// R8 is MOD_P

TEXT ·butterfly_null(SB),7,$0-0
	// u = mod_add(a, b)
	// v = mod_sub(a, b)
	MOVQ AX, CX
	MOVQ AX, DX
	MOD_ADD(CX, BX, R9, R8, butterfly_null_a)
	MOD_SUB(DX, BX, R8, butterfly_null_b)
	RET

{{ range .Shifts }}
TEXT ·butterfly_shift{{.}}(SB),7,$0-0
	// u = mod_add(a, b)
	// v = mod_shift{{.}}(mod_sub(a, b))
	MOVQ AX, CX
	MOVQ AX, DX
	MOD_ADD(CX, BX, R9, R8, butterfly_shift{{.}}_a)
	MOD_SUB(DX, BX, R8, butterfly_shift{{.}}_b)
{{ if lt . 32 }}
	MOD_SHIFT_0_TO_31(DX, {{.}}, AX, BX, R8, butterfly_shift{{.}}_c)
{{ else if lt . 64 }}
	MOD_SHIFT_32_TO_63(DX, {{.}}, AX, BX, R9, R8, butterfly_shift{{.}}_d)
{{ else }}
	MOD_SHIFT_64_TO_95(DX, {{.}}, AX, BX, R8, butterfly_shift{{.}}_e)
{{ end }}
	RET
{{ end }}

TEXT ·butterfly_mul(SB),7,$0-0
	// u = mod_add(a, b)
	// v = mod_mul(mod_sub(a, b), w)
	MOVQ AX, CX
	MOD_ADD(CX, BX, R9, R8, butterfly_mul_a)
	MOD_SUB(AX, BX, R8, butterfly_mul_b)
        MULQ	DI /* DI * AX -> (DX, AX) */
        MOD_REDUCE(DX, AX, BX, R9, R8, butterfly_mul_c)
	MOVQ AX, DX
	RET

TEXT ·invbutterfly_null(SB),7,$0-0
	// u = mod_add(a, b)
	// v = mod_sub(a, b)
	MOVQ AX, CX
	MOVQ AX, DX
	MOD_ADD(CX, BX, R9, R8, invbutterfly_null_a)
	MOD_SUB(DX, BX, R8, invbutterfly_null_b)
	RET

{{ range .Shifts }}
// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
TEXT ·invbutterfly_shift{{.}}(SB),7,$0-0
	// b = mod_shift{{.}}(b)
	// u = mod_sub(a, b)
	// v = mod_add(a, b)
{{ if lt . 32 }}
	MOD_SHIFT_0_TO_31(BX, {{.}}, CX, DX, R8, invbutterfly_shift{{.}}_c)
{{ else if lt . 64 }}
	MOD_SHIFT_32_TO_63(BX, {{.}}, CX, DX, R9, R8, invbutterfly_shift{{.}}_d)
{{ else }}
	MOD_SHIFT_64_TO_95(BX, {{.}}, CX, DX, R8, invbutterfly_shift{{.}}_e)
{{ end }}
	MOVQ AX, CX
	MOVQ AX, DX
	MOD_SUB(CX, BX, R8, invbutterfly_shift{{.}}_a)
	MOD_ADD(DX, BX, R9, R8, invbutterfly_shift{{.}}_b)
	RET
{{ end }}

TEXT ·invbutterfly_mul(SB),7,$0-0
	// b = mod_mul(b, w)
	// u = mod_add(a, b)
	// v = mod_sub(a, b)
	MOVQ AX, CX
	MOVQ BX, AX
        MULQ	DI /* DI * AX -> (DX, AX) */
        MOD_REDUCE(DX, AX, SI, BX, R8, invbutterfly_mul_a)
	MOVQ CX, DX
	MOD_ADD(CX, AX, R9, R8, invbutterfly_mul_b)
	MOD_SUB(DX, AX, R8, invbutterfly_mul_c)
	RET

{{ range $size := .FftSizes }}
// Fft for size 2**{{$size}}
//
// This is an in place FFT with a bit reversed output
TEXT ·fft{{$size}}(SB),7,$0-24
	MOVQ $MOD_P, R8
	MOVQ x_len+8(FP), AX
	MOVQ x+0(FP), R12
	CMPQ AX, $(1<<{{$size}})
	JAE fft{{$size}}ok
	CALL runtime·panicslice(SB)
fft{{$size}}ok:
{{ Fft $size }}
	RET

// InvFft for size 2**{{$size}}
//
// This is an in place Inverse FFT with a bit reversed input
TEXT ·invfft{{$size}}(SB),7,$0-24
	MOVQ $MOD_P, R8
	MOVQ x_len+8(FP), AX
	MOVQ x+0(FP), R12
	CMPQ AX, $(1<<{{$size}})
	JAE invfft{{$size}}ok
	CALL runtime·panicslice(SB)
invfft{{$size}}ok:
{{ InvFft $size }}
	RET
{{ end }}
`

// Load a quad word in r0, r1 to x[i]
func LDRQ(out io.Writer, r0, r1 string, i uint) {
	if i*8 < (1 << 12) {
		// ARM offsets are 12 bits for LDR/STR
		fmt.Fprintf(out, "\tMOVW (%d)(R11), %s\n", i*8, r0)
		fmt.Fprintf(out, "\tMOVW (%d)(R11), %s\n", i*8+4, r1)
	} else {
		// ARM immediate constants are 8 bits shifted in multiples of 2
		i <<= 1
		fmt.Fprintf(out, "\tADD $(%d<<2), R11, R8\n", i&0xFF00)
		if i&0xFF != 0 {
			fmt.Fprintf(out, "\tADD $(%d<<2), R8, R8\n", i&0xFF)
		}
		fmt.Fprintf(out, "\tMOVM.IA (R8), [%s,%s]\n", r0, r1)
	}
}

// Save a quad word in r0, r1 to x[i]
func STRQ(out io.Writer, r0, r1 string, i uint) {
	if i*8 < (1 << 12) {
		// ARM offsets are 12 bits for LDR/STR
		fmt.Fprintf(out, "\tMOVW %s, (%d)(R11)\n", r0, i*8)
		fmt.Fprintf(out, "\tMOVW %s, (%d)(R11)\n", r1, i*8+4)
	} else {
		// ARM immediate constants are 8 bits shifted in multiples of 2
		i <<= 1
		fmt.Fprintf(out, "\tADD $(%d<<2), R11, R8\n", i&0xFF00)
		if i&0xFF != 0 {
			fmt.Fprintf(out, "\tADD $(%d<<2), R8, R8\n", i&0xFF)
		}
		fmt.Fprintf(out, "\tMOVM.IA [%s,%s], (R8)\n", r0, r1)
	}
}

func MakeArmFft(_log_n int) string {
	out := new(bytes.Buffer)
	MakeFftGeneric(_log_n, func(a, b uint, w uint64) {
		LDRQ(out, "R0", "R1", a)
		LDRQ(out, "R2", "R3", b)
		var a0, a1, b0, b1 string
		switch i := multiplyType(w); i {
		case 0:
			fmt.Fprintf(out, "\tBL ·butterfly_null(SB)\n")
			a0, a1, b0, b1 = "R4", "R5", "R6", "R7"
		case -1:
			fmt.Fprintf(out, "\tMOVW $0x%08X, R4\n", w&0xFFFFFFFF)
			fmt.Fprintf(out, "\tMOVW $0x%08X, R5\n", (w>>32)&0xFFFFFFFF)
			fmt.Fprintf(out, "\tBL ·butterfly_mul(SB)\n")
			a0, a1, b0, b1 = "R6", "R7", "R0", "R1"
		default:
			fmt.Fprintf(out, "\tBL ·butterfly_shift%d(SB)\n", i)
			a0, a1, b0, b1 = "R4", "R5", "R0", "R1"
		}
		STRQ(out, a0, a1, a)
		STRQ(out, b0, b1, b)
	})
	return out.String()
}

func MakeArmInvFft(_log_n int) string {
	out := new(bytes.Buffer)
	MakeInvFftGeneric(_log_n, func(a, b uint, w uint64) {
		LDRQ(out, "R0", "R1", a)
		LDRQ(out, "R2", "R3", b)
		var a0, a1, b0, b1 string
		switch i := multiplyType(w); i {
		case 0:
			fmt.Fprintf(out, "\tBL ·invbutterfly_null(SB)\n")
			a0, a1, b0, b1 = "R4", "R5", "R6", "R7"
		case -1:
			fmt.Fprintf(out, "\tMOVW $0x%08X, R4\n", w&0xFFFFFFFF)
			fmt.Fprintf(out, "\tMOVW $0x%08X, R5\n", (w>>32)&0xFFFFFFFF)
			fmt.Fprintf(out, "\tBL ·invbutterfly_mul(SB)\n")
			a0, a1, b0, b1 = "R2", "R3", "R0", "R1"
		default:
			// Since 2^96 mod p = -1, we subtract -1 and use a negative butterfly
			fmt.Fprintf(out, "\tBL ·invbutterfly_shift%d(SB)\n", i-96)
			a0, a1, b0, b1 = "R6", "R7", "R2", "R3"
		}
		STRQ(out, a0, a1, a)
		STRQ(out, b0, b1, b)
	})
	return out.String()
}

// Make ffts_arm.go
func generateArmFfts() {
	t := template.Must(template.New("main").Funcs(template.FuncMap{
		"Fft":    MakeArmFft,
		"InvFft": MakeArmInvFft,
	}).Parse(armAssembler))

	out, err := os.Create("ffts_arm.h")
	if err != nil {
		log.Fatal(err)
	}
	if err = t.Execute(out, Data); err != nil {
		log.Fatal(err)
	}
	if err = out.Close(); err != nil {
		log.Fatal(err)
	}
}

var armAssembler = `
// Automatically generated - DO NOT EDIT
// Regenerate with: {{ .Regenerate }}

// NB Non golang calling convention here for the butterflies

// Input in (R0, R1) (R2, R3)
// Output in (R4, R5) (R6, R7)
// Preserves R11
// R12 = -1
TEXT ·butterfly_null(SB),7,$-4-0
	// u = mod_add(a, b)
	// v = mod_sub(a, b)
	MOD_ADD(R4, R5, R0, R1, R2, R3, R12)
	MOD_SUB(R6, R7, R0, R1, R2, R3, R12)
	RET

// Input in (R0, R1) (R2, R3)
// Output in (R4, R5) (R0, R1)
// Preserves R11
// R12 = -1

{{ range .Shifts }}
TEXT ·butterfly_shift{{.}}(SB),7,$-4-0
	// u = mod_add(a, b)
	// v = mod_shift{{.}}(mod_sub(a, b))
	MOD_ADD(R4, R5, R0, R1, R2, R3, R12)
	MOD_SUB(R6, R7, R0, R1, R2, R3, R12)
{{ if lt . 32 }}
	MOVW $0, R8
	MOD_SHIFT_0_TO_31({{.}}, R0, R1, R6, R7, R12, R8)
{{ else if lt . 64 }}
	MOD_SHIFT_32_TO_63({{.}}, R0, R1, R6, R7, R12)
{{ else }}
	MOVW $0, R8
	MOD_SHIFT_64_TO_95({{.}}, R0, R1, R6, R7, R12, R8)
{{ end }}
	RET
{{ end }}

// Input in (R0, R1) (R2, R3)
// Twiddle in (R4, R5)
// Output in (R6, R7) (R0, R1)
// Preserves R11
// R12 = -1

TEXT ·butterfly_mul(SB),7,$4-0
	// u = mod_add(a, b)
	// v = mod_mul(mod_sub(a, b), w)
	MOVW R11, -4(SP)
	MOD_ADD(R6, R7, R0, R1, R2, R3, R12)
	MOD_SUB(R8, R11, R0, R1, R2, R3, R12)
	MOD_MUL(R0,R1, R4,R5, R8,R11, R2,R3,R14, R12, butterfly_mul_a)
	MOVW -4(SP), R11
	RET

// Input in (R0, R1) (R2, R3)
// Output in (R4, R5) (R6, R7)
// Preserves R11
// R12 = -1
TEXT ·invbutterfly_null(SB),7,$-4-0
	// u = mod_add(a, b)
	// v = mod_sub(a, b)
	MOD_ADD(R4, R5, R0, R1, R2, R3, R12)
	MOD_SUB(R6, R7, R0, R1, R2, R3, R12)
	RET

// Input in (R0, R1) (R2, R3)
// Output in (R6, R7) (R0, R1)
// Preserves R11
// R12 = -1
// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1

{{ range .Shifts }}
TEXT ·invbutterfly_shift{{.}}(SB),7,$0-0
	// b = mod_shift{{.}}(b)
	// u = mod_sub(a, b)
	// v = mod_add(a, b)
{{ if lt . 32 }}
	MOVW $0, R8
	MOD_SHIFT_0_TO_31({{.}}, R4, R5, R2, R3, R12, R8)
{{ else if lt . 64 }}
	MOD_SHIFT_32_TO_63({{.}}, R4, R5, R2, R3, R12)
{{ else }}
	MOVW $0, R8
	MOD_SHIFT_64_TO_95({{.}}, R4, R5, R2, R3, R12, R8)
{{ end }}
	MOD_SUB(R6, R7, R0, R1, R4, R5, R12)
	MOD_ADD(R2, R3, R0, R1, R4, R5, R12)
	RET
{{ end }}

// Input in (R0, R1) (R2, R3)
// Twiddle in (R4, R5)
// Output in (R2, R3) (R0, R1)
// Preserves R11
// R12 = -1

TEXT ·invbutterfly_mul(SB),7,$4-0
	// b = mod_mul(b, w)
	// u = mod_add(a, b)
	// v = mod_sub(a, b)
	MOVW R11, -4(SP)
	MOD_MUL(R6,R7, R4,R5, R2,R3, R8,R11,R14, R12, invbutterfly_mul_a)
	MOVW -4(SP), R11
	MOD_ADD(R2, R3, R0, R1, R6, R7, R12)
	MOD_SUB(R0, R1, R0, R1, R6, R7, R12)
	RET

{{ range $size := .FftSizes }}
// Fft for size 2**{{$size}}
//
// This is an in place FFT with a bit reversed output
TEXT ·fft{{$size}}(SB),7,$0-12
	MOVW $-1, R12
	MOVW x_len+4(FP), R0
	MOVW x+0(FP), R11
	CMP $(1<<{{$size}}), R0
	BL.LO runtime·panicslice(SB)
{{ Fft $size }}
	RET

// InvFft for size 2**{{$size}}
//
// This is an in place Inverse FFT with a bit reversed input
TEXT ·invfft{{$size}}(SB),7,$0-12
	MOVW $-1, R12
	MOVW x_len+4(FP), R0
	MOVW x+0(FP), R11
	CMP $(1<<{{$size}}), R0
	BL.LO runtime·panicslice(SB)
{{ InvFft $size }}
	RET
{{ end }}
`

func main() {
	generateFfts()
	generateAmd64Ffts()
	generateArmFfts()
}
