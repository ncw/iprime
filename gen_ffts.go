// +build ignore

// Generate ffts.go like this
//
// go run mod_math.go gen_ffts.go | gofmt >ffts.go
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var FftSizes = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //,11,12}

var PowersOfTwo = make([]uint64, 193)

func init() {
	x := uint64(1)
	for i := range PowersOfTwo {
		PowersOfTwo[i] = x
		x = mod_mul(x, uint64(2))
	}
}

func make_mod_mul(a string, w uint64) string {
	if w == 0 {
		return "0"
	}
	if w == 1 {
		return a
	}
	for i, x := range PowersOfTwo {
		if x == w {
			return fmt.Sprintf("mod_shift(%s, %d)", a, i)
		}
	}
	return fmt.Sprintf("mod_mul(%s, %d)", a, w)
}

func MakeFft(_log_n int) string {
	log_n := uint8(_log_n)
	if log_n == 0 {
		return ""
	}
	out := "var u, v uint64\n"
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
				out += fmt.Sprintf(`
				u, v = x[%d], x[%d]
				x[%d] = mod_add(u, v)
                                u = mod_sub(u, v)
				x[%d] = %s`, a, b, a, b, make_mod_mul("u", w))
			}
			w = mod_mul(w, d)
		}
		d = mod_mul(d, d)
	}
	return out
}

func MakeInvFft(n int) string {
	return fmt.Sprintf("// InvFft %d", n)
}

func main() {
	t := template.Must(template.New("main").Funcs(template.FuncMap{
		"Fft":    MakeFft,
		"InvFft": MakeInvFft,
	}).Parse(program))
	if err := t.Execute(os.Stdout, FftSizes); err != nil {
		log.Fatal(err)
	}
}

var program = `
// Automatically generated - DO NOT EDIT
// Regenerate with: go run mod_math.go gen_ffts.go | gofmt >ffts.go

// Unrolled FFTs

package main

{{ range $size := . }}
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
