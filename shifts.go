// Automatically generated - DO NOT EDIT
// Regenerate with: go run mod_math.go gen_shifts.go | gofmt >shifts.go

// +build !amd64

// Unrolled mod_shifts

package main

// Shift a value by 3 bits (multiply by 2^3).
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
func mod_shift3(x uint64) uint64 {
	xmid_xlow := x << 3 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 3))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 6 bits (multiply by 2^6).
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
func mod_shift6(x uint64) uint64 {
	xmid_xlow := x << 6 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 6))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 9 bits (multiply by 2^9).
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
func mod_shift9(x uint64) uint64 {
	xmid_xlow := x << 9 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 9))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 12 bits (multiply by 2^12).
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
func mod_shift12(x uint64) uint64 {
	xmid_xlow := x << 12 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 12))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 15 bits (multiply by 2^15).
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
func mod_shift15(x uint64) uint64 {
	xmid_xlow := x << 15 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 15))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 18 bits (multiply by 2^18).
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
func mod_shift18(x uint64) uint64 {
	xmid_xlow := x << 18 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 18))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 21 bits (multiply by 2^21).
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
func mod_shift21(x uint64) uint64 {
	xmid_xlow := x << 21 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 21))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 24 bits (multiply by 2^24).
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
func mod_shift24(x uint64) uint64 {
	xmid_xlow := x << 24 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 24))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 27 bits (multiply by 2^27).
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
func mod_shift27(x uint64) uint64 {
	xmid_xlow := x << 27 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 27))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 30 bits (multiply by 2^30).
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
func mod_shift30(x uint64) uint64 {
	xmid_xlow := x << 30 // (xmid, xlow)
	xhigh := uint32(x >> (64 - 30))
	t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1) // (2^32 - 1 - xhigh, xhigh + 1)
	r := xmid_xlow - t                                  // (xmid, xlow) + (xhigh, -xhigh)
	// carry
	if xmid_xlow < t {
		r += MOD_P
	}
	return r
}

// Shift a value by 33 bits (multiply by 2^33).
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
func mod_shift33(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 33))
	xmid := uint32(x >> (64 - 33))
	xlow := uint32(x << (33 - 32))
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

// Shift a value by 36 bits (multiply by 2^36).
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
func mod_shift36(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 36))
	xmid := uint32(x >> (64 - 36))
	xlow := uint32(x << (36 - 32))
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

// Shift a value by 39 bits (multiply by 2^39).
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
func mod_shift39(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 39))
	xmid := uint32(x >> (64 - 39))
	xlow := uint32(x << (39 - 32))
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

// Shift a value by 42 bits (multiply by 2^42).
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
func mod_shift42(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 42))
	xmid := uint32(x >> (64 - 42))
	xlow := uint32(x << (42 - 32))
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

// Shift a value by 45 bits (multiply by 2^45).
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
func mod_shift45(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 45))
	xmid := uint32(x >> (64 - 45))
	xlow := uint32(x << (45 - 32))
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

// Shift a value by 48 bits (multiply by 2^48).
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
func mod_shift48(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 48))
	xmid := uint32(x >> (64 - 48))
	xlow := uint32(x << (48 - 32))
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

// Shift a value by 51 bits (multiply by 2^51).
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
func mod_shift51(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 51))
	xmid := uint32(x >> (64 - 51))
	xlow := uint32(x << (51 - 32))
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

// Shift a value by 54 bits (multiply by 2^54).
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
func mod_shift54(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 54))
	xmid := uint32(x >> (64 - 54))
	xlow := uint32(x << (54 - 32))
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

// Shift a value by 57 bits (multiply by 2^57).
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
func mod_shift57(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 57))
	xmid := uint32(x >> (64 - 57))
	xlow := uint32(x << (57 - 32))
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

// Shift a value by 60 bits (multiply by 2^60).
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
func mod_shift60(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 60))
	xmid := uint32(x >> (64 - 60))
	xlow := uint32(x << (60 - 32))
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

// Shift a value by 63 bits (multiply by 2^63).
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
func mod_shift63(x uint64) uint64 {
	xhigh := uint32(x >> (96 - 63))
	xmid := uint32(x >> (64 - 63))
	xlow := uint32(x << (63 - 32))
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

// Shift a value  by 66 bits (multiply by 2^66).
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
func mod_shift66(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 66))
	xmid := uint32(x >> (96 - 66))
	xlow := uint32(x << (66 - 64))
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

// Shift a value  by 69 bits (multiply by 2^69).
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
func mod_shift69(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 69))
	xmid := uint32(x >> (96 - 69))
	xlow := uint32(x << (69 - 64))
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

// Shift a value  by 72 bits (multiply by 2^72).
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
func mod_shift72(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 72))
	xmid := uint32(x >> (96 - 72))
	xlow := uint32(x << (72 - 64))
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

// Shift a value  by 75 bits (multiply by 2^75).
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
func mod_shift75(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 75))
	xmid := uint32(x >> (96 - 75))
	xlow := uint32(x << (75 - 64))
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

// Shift a value  by 78 bits (multiply by 2^78).
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
func mod_shift78(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 78))
	xmid := uint32(x >> (96 - 78))
	xlow := uint32(x << (78 - 64))
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

// Shift a value  by 81 bits (multiply by 2^81).
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
func mod_shift81(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 81))
	xmid := uint32(x >> (96 - 81))
	xlow := uint32(x << (81 - 64))
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

// Shift a value  by 84 bits (multiply by 2^84).
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
func mod_shift84(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 84))
	xmid := uint32(x >> (96 - 84))
	xlow := uint32(x << (84 - 64))
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

// Shift a value  by 87 bits (multiply by 2^87).
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
func mod_shift87(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 87))
	xmid := uint32(x >> (96 - 87))
	xlow := uint32(x << (87 - 64))
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

// Shift a value  by 90 bits (multiply by 2^90).
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
func mod_shift90(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 90))
	xmid := uint32(x >> (96 - 90))
	xlow := uint32(x << (90 - 64))
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

// Shift a value  by 93 bits (multiply by 2^93).
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
func mod_shift93(x uint64) uint64 {
	xhigh := uint32(x >> (128 - 93))
	xmid := uint32(x >> (96 - 93))
	xlow := uint32(x << (93 - 64))
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
