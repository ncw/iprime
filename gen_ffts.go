// +build ignore

// FIXME factor add/sub/mul/shift out into butterfly
// which will stop excessive inlining and maybe be quicker
// x[1], x[2] = butterfly(x[1], x[2], w)
// x[1], x[2] = butterfly_shift(x[1], x[2], 5)
// x[1], x[2] = butterfly_shift5(x[1], x[2])

// FIXME specialise shifts by generating code for them mod_shift5

// Generate ffts.go like this
//
// go run mod_math.go gen_ffts.go | gofmt >ffts.go
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

var Data = struct {
	FftSizes []int
	Shifts   []int
}{
	FftSizes: []int{0, 1, 2, 3, 4, 5, 6, 7}, //, 8, 9, 10,11,12}
	Shifts:   []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93},
}

var PowersOfTwo = make([]uint64, 193)

func init() {
	x := uint64(1)
	for i := range PowersOfTwo {
		PowersOfTwo[i] = x
		x = mod_mul(x, uint64(2))
	}
}

func make_butterfly(a, b uint, w uint64) string {
	if w == 0 {
		panic("Butterflies can't be 0")
	}
	if w == 1 {
		return fmt.Sprintf("butterfly_null(x[%d], x[%d])", a, b)
	}
	for i, x := range PowersOfTwo {
		if x == w {
			return fmt.Sprintf("butterfly_shift%d(x[%d], x[%d])", i, a, b)
		}
	}
	return fmt.Sprintf("butterfly_mul(x[%d], x[%d], %d)", a, b, w)
}

func make_invbutterfly(a, b uint, w uint64) string {
	if w == 0 {
		panic("Butterflies can't be 0")
	}
	if w == 1 {
		return fmt.Sprintf("invbutterfly_null(x[%d], x[%d])", a, b)
	}
	for i, x := range PowersOfTwo {
		if x == w {
			// Since 2^96 mod p = -1, we subtract -1 and use a negative butterfly
			return fmt.Sprintf("invbutterfly_shift%d(x[%d], x[%d])", i-96, a, b)
		}
	}
	return fmt.Sprintf("invbutterfly_mul(x[%d], x[%d], %d)", a, b, w)
}

func MakeFft(_log_n int) string {
	log_n := uint8(_log_n)
	if log_n == 0 {
		return ""
	}
	n := uint(1) << log_n
	mod_w := mod_pow(7, (MOD_P-1)/uint64(n))
	mod_w = mod_pow(mod_w, 5)
	var d uint64 = mod_w

	out := new(bytes.Buffer)
	for k := log_n; k >= 1; k-- {
		m := uint(1) << k
		c := m >> 1
		var w uint64 = 1
		for j := uint(0); j < c; j++ {
			for r := uint(0); r < n; r += m {
				a := r + j
				b := a + c
				fmt.Fprintf(out, "x[%d], x[%d] = %s\n", a, b, make_butterfly(a, b, w))
			}
			w = mod_mul(w, d)
		}
		d = mod_mul(d, d)
	}
	return out.String()
}

func MakeInvFft(_log_n int) string {
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
				fmt.Fprintf(out, "x[%d], x[%d] = %s\n", a, b, make_invbutterfly(a, b, w))
			}
			w = mod_mul(w, d)
		}
	}
	return out.String()
}

func main() {
	t := template.Must(template.New("main").Funcs(template.FuncMap{
		"Fft":    MakeFft,
		"InvFft": MakeInvFft,
	}).Parse(program))
	if err := t.Execute(os.Stdout, Data); err != nil {
		log.Fatal(err)
	}
}

var program = `
// Automatically generated - DO NOT EDIT
// Regenerate with: go run mod_math.go gen_ffts.go | gofmt >ffts.go

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
