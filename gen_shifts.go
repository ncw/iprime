// +build ignore

// Generate shifts.go like this
//
// go run mod_math.go mod_math_noasm.go gen_shifts.go | gofmt >shifts.go
package main

import (
	"log"
	"os"
	"text/template"
)

var shifts = [][]int{
	{3, 6, 9, 12, 15, 18, 21, 24, 27, 30},
	{33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63},
	{66, 69, 72, 75, 78, 81, 84, 87, 90, 93},
}

func main() {
	t := template.Must(template.New("main").Parse(program))
	if err := t.Execute(os.Stdout, shifts); err != nil {
		log.Fatal(err)
	}
}

var program = `
// Automatically generated - DO NOT EDIT
// Regenerate with: go run mod_math.go gen_shifts.go | gofmt >shifts.go

// +build !amd64,!arm

// Unrolled mod_shifts

package main

{{ range $shift := index . 0 }}
// Shift a value by {{ $shift }} bits (multiply by 2^{{ $shift }}).
//
// For shifts 0..31 bits
// (xhigh, xmid, xlow) = (x1,x0) << shift where shift is 0..31
// (xhigh, xmid, xlow) mod p
// = (xmid + xhigh, xlow- xhigh)
// = (xmid, xlow) + (xhigh, -xhigh)
// negate (xhigh, -xhigh) so we can use a subtract rather than an add
// p - (xhigh, -xhigh) = (2^32 - 1 - xhigh, xhigh + 1)
// is in range 2^32..p.  The add xhigh + 1 cannot overflow since xhigh < 2^31
// = (xmid, xlow) - (2^32 - 1 - xhigh, xhigh + 1)
// (xmid, xlow) is in range 0..2^64-1, therefore result of subtract is in
// range -p..2^64-1-2^32 = -p..p-2 which is ok for the subtract sequence
func mod_shift{{ $shift }}(x uint64) uint64 {
	xmid_xlow := x << {{ $shift }} // (xmid, xlow)
	xhigh := uint32(x >> (64 - {{ $shift }}))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

{{ end }}
{{ range $shift := index . 1 }}
// Shift a value by {{ $shift }} bits (multiply by 2^{{ $shift }}).
//
// For shifts between 32..63 bits
// (xhigh, xmid, xlow, 0) = (x1,x0) << shift where shift is 32..63
// (xhigh, xmid, xlow, 0) mod p
// = (xmid, xlow, - xhigh)
// = (xlow + xmid, - xhigh - xmid)
// This can be negative and can exceed p.
// = (xlow, 0) - (-xmid, xhigh + xmid) mod p
// note that (xlow, 0) < p
//
// xmidneg = 0 - xmid (mod 2^32)
// xmidcomp = xmidneg - (borrow from last subtract)
// temp = (xmidneg, xhigh) - (0, xmidcomp)
//
// If xmid = 0, then xmidneg = xmidcomp = 0 and temp2 = (0, xhigh) which is 0..p-1
// Otherwise xmidneg = 2^32 - xmid and xmidcomp = 2^32 - 1 - xmid
// so temp = (2^32 - xmid, xhigh) - (0, 2^32 - 1 - xmid)
//         = (2^32, -2^32 + 1) + (-xmid, xhigh + xmid)
//         = (2^32 - 1, 1) + (-xmid, xhigh + xmid)
//         = p + (-xmid, xhigh + xmid)
//
// If xmid = 1 then temp = (2^32 - 1, xhigh) - (0, 2^32 - 2) since xhigh < 2^31 this is < p
// if xmid = 2^32 - 1 then temp = (1, xhigh) - (0, 0) which is < p
func mod_shift{{ $shift }}(x uint64) uint64 {
	xhigh := uint32(x >> (96 - {{ $shift }}))
	xmid := uint32(x >> (64 - {{ $shift }}))
	xlow := uint32(x << ({{ $shift }} - 32))
	t0 := uint64(xmid) << 32 // (xmid, 0)
	t1 := uint64(xmid)       // (0, xmid)

	t0 -= t1           // (xmid, -xmid) no carry and must be in range 0..p-1
	t1 = uint64(xhigh) // (0, xhigh)
	r := t0 - t1       // (xmid, - xhigh - xmid)
	// carry?
	if t0 < t1 {
		r += MOD_P
	}
	t0 = r

	// add (xlow, 0) by subtacting p - (xlow, 0) = (2^32 - 1 - xlow, 1)
	t1 = uint64(0xFFFFFFFF-xlow)<<32 + 1 // -(xlow, 0)
	r = t0 - t1                          // (xlow + xmid, - xhigh - xmid)
	// carry?
	if t0 < t1 {
		r += MOD_P
	}

	return r
}

{{ end }}
{{ range $shift := index . 2 }}
// Shift a value  by {{ $shift }} bits (multiply by 2^{{ $shift }}).
//
// For shifts between 64..95 bits
// (xhigh, xmid, xlow, 0, 0) = (x1,x0) << shift where shift is 64..95
// (xhigh, xmid, xlow, 0, 0) mod p
// = (xmid, xlow, -xhigh, 0)
// = (xlow, -xhigh, -xmid)
// = (xlow - xhigh, -xmid - xlow)
// = (xlow, -xlow) - (xhigh, xmid)
// (xhigh, xmid) is < p since xhigh < 2^31
// (xlow, -xlow) can be evaluated as (xlow, 0) - (0, xlow) this is < p
func mod_shift{{ $shift }}(x uint64) uint64 {
	xhigh := uint32(x >> (128 - {{ $shift }}))
	xmid := uint32(x >> (96 - {{ $shift }}))
	xlow := uint32(x << ({{ $shift }} - 64))
	t0 := uint64(xlow) << 32 // (xlow, 0)
	t1 := uint64(xlow)       // (0, xlow)

	t0 -= t1 // (xlow, -xlow) - no carry possible

	t1 = uint64(xhigh)<<32 + uint64(xmid) // (xhigh, xmid)
	r := t0 - t1                          // (xlow, -xlow) - (xhigh, xmid)
	// carry?
	if t0 < t1 {
		r += MOD_P
	}
	return r
}

{{ end }}
`
