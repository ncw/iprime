// Go alternatives for assembler routines

// +build !amd64

package main

// Define addition as x - (p - y).  This seemingly backwards way of doing
// things works very well as it is much easier to make sure the result of a
// subtraction is in the range 0..p-1 than an addition (the conditional
// instructions for comparing > FFFFFFFF00000000 just don't work out
//
// The addition sequence for x + y (mod p) works whenever
//
//   0 <= x <= 2^64 - 1
//   0 <= y <= p
//   x + y <= 2p - 1
func mod_add(x, y uint64) uint64 {
	x = MOD_P - x // do addition by negating y then subracting
	// This can't overflow if y < p
	//
	// what if y is 0, -y = p-0 which is out of range
	// however subtracting y=p from x < p is guaranteed to carry which will fix it
	// x = 0, y = 0, -y = p, x - -y = 0 - p = -p => carry, add p == 0 ok
	// x = p-1, y = 0, -y = p, x - -y = p-1 - p = -1 => carry, add p == p-1 ok

	r := y - x // y - (-x)
	// if borrow generated - hopefully the compiler will optimise this!
	if y < x {
		r += MOD_P // Add back p if borrow
	}
	return r
}

// The subtraction sequence for x - y (mod p) works whenever
//
//    0 <= x <= 2^64 - 1
//    0 <= y <= 2^64 - 1
//    -p <= x - y <= p - 1
func mod_sub(x, y uint64) uint64 {

	r := x - y
	// if borrow generated - hopefully the compiler will optimise this!
	if x < y {
		r += MOD_P // Add back p if borrow
	}
	return r
}

// This reduces a 128 bit product mod p
//
//   (x3,x2,x1,x0) mod p
//   = (x2,x1,x0-x3)	[2^96 mod p = -1]
//   = (x1+x2,x0-x3-x2)	[2^64 mod p = 2^32 -1]
//
// Care is needed with the carries
func mod_reduce(b, a uint64) uint64 {
	d := uint32(b >> 32)
	c := uint32(b)
	if a >= MOD_P { // (x1, x0)
		a -= MOD_P
	}
	a = mod_sub(a, uint64(c))
	a = mod_sub(a, uint64(d))
	a = mod_add(a, uint64(c)<<32)
	return a
}

// Mod Multiply
//
// x,y must be in range 0..p-1
// z will be in range 0..p-1
func mod_mul(x, y uint64) uint64 {
	a := uint32(x)
	b := uint32(x >> 32)
	c := uint32(y)
	d := uint32(y >> 32)

	// first synthesize the product using 32*32 -> 64 bit multiplies
	x = uint64(b) * uint64(c)  // b*c
	y = uint64(a) * uint64(d)  // a*d
	e := uint64(a) * uint64(c) // a*c
	f := uint64(b) * uint64(d) // b*d

	x += y // b*c + a*d
	// carry?
	if x < y {
		f += 1 << 32 // carry into upper 32 bits - can't overflow
	}

	t := x << 32
	e += t // a*c + LSW(b*c + a*d)
	// carry?
	if e < t {
		f += 1 // carry into upper 64 bits - can't overflow
	}
	t = x >> 32
	f += t // b*d + MSW(b*c + a*d)
	// can't overflow

	// now reduce: (b*d + MSW(b*c + a*d), a*c + LSW(b*c + a*d))
	return mod_reduce(f, e)
}

