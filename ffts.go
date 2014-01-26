// Automatically generated - DO NOT EDIT
// Regenerate with: go run mod_math.go gen_ffts.go | gofmt >ffts.go

// Unrolled FFTs

package main

func butterfly_null(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_sub(a, b)
	return
}

func butterfly_shift3(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift3(mod_sub(a, b))
	return
}

func butterfly_shift6(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift6(mod_sub(a, b))
	return
}

func butterfly_shift9(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift9(mod_sub(a, b))
	return
}

func butterfly_shift12(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift12(mod_sub(a, b))
	return
}

func butterfly_shift15(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift15(mod_sub(a, b))
	return
}

func butterfly_shift18(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift18(mod_sub(a, b))
	return
}

func butterfly_shift21(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift21(mod_sub(a, b))
	return
}

func butterfly_shift24(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift24(mod_sub(a, b))
	return
}

func butterfly_shift27(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift27(mod_sub(a, b))
	return
}

func butterfly_shift30(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift30(mod_sub(a, b))
	return
}

func butterfly_shift33(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift33(mod_sub(a, b))
	return
}

func butterfly_shift36(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift36(mod_sub(a, b))
	return
}

func butterfly_shift39(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift39(mod_sub(a, b))
	return
}

func butterfly_shift42(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift42(mod_sub(a, b))
	return
}

func butterfly_shift45(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift45(mod_sub(a, b))
	return
}

func butterfly_shift48(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift48(mod_sub(a, b))
	return
}

func butterfly_shift51(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift51(mod_sub(a, b))
	return
}

func butterfly_shift54(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift54(mod_sub(a, b))
	return
}

func butterfly_shift57(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift57(mod_sub(a, b))
	return
}

func butterfly_shift60(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift60(mod_sub(a, b))
	return
}

func butterfly_shift63(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift63(mod_sub(a, b))
	return
}

func butterfly_shift66(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift66(mod_sub(a, b))
	return
}

func butterfly_shift69(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift69(mod_sub(a, b))
	return
}

func butterfly_shift72(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift72(mod_sub(a, b))
	return
}

func butterfly_shift75(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift75(mod_sub(a, b))
	return
}

func butterfly_shift78(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift78(mod_sub(a, b))
	return
}

func butterfly_shift81(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift81(mod_sub(a, b))
	return
}

func butterfly_shift84(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift84(mod_sub(a, b))
	return
}

func butterfly_shift87(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift87(mod_sub(a, b))
	return
}

func butterfly_shift90(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift90(mod_sub(a, b))
	return
}

func butterfly_shift93(a, b uint64) (u, v uint64) {
	u = mod_add(a, b)
	v = mod_shift93(mod_sub(a, b))
	return
}

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

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift3(a, b uint64) (u, v uint64) {
	b = mod_shift3(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift6(a, b uint64) (u, v uint64) {
	b = mod_shift6(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift9(a, b uint64) (u, v uint64) {
	b = mod_shift9(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift12(a, b uint64) (u, v uint64) {
	b = mod_shift12(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift15(a, b uint64) (u, v uint64) {
	b = mod_shift15(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift18(a, b uint64) (u, v uint64) {
	b = mod_shift18(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift21(a, b uint64) (u, v uint64) {
	b = mod_shift21(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift24(a, b uint64) (u, v uint64) {
	b = mod_shift24(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift27(a, b uint64) (u, v uint64) {
	b = mod_shift27(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift30(a, b uint64) (u, v uint64) {
	b = mod_shift30(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift33(a, b uint64) (u, v uint64) {
	b = mod_shift33(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift36(a, b uint64) (u, v uint64) {
	b = mod_shift36(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift39(a, b uint64) (u, v uint64) {
	b = mod_shift39(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift42(a, b uint64) (u, v uint64) {
	b = mod_shift42(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift45(a, b uint64) (u, v uint64) {
	b = mod_shift45(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift48(a, b uint64) (u, v uint64) {
	b = mod_shift48(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift51(a, b uint64) (u, v uint64) {
	b = mod_shift51(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift54(a, b uint64) (u, v uint64) {
	b = mod_shift54(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift57(a, b uint64) (u, v uint64) {
	b = mod_shift57(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift60(a, b uint64) (u, v uint64) {
	b = mod_shift60(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift63(a, b uint64) (u, v uint64) {
	b = mod_shift63(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift66(a, b uint64) (u, v uint64) {
	b = mod_shift66(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift69(a, b uint64) (u, v uint64) {
	b = mod_shift69(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift72(a, b uint64) (u, v uint64) {
	b = mod_shift72(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift75(a, b uint64) (u, v uint64) {
	b = mod_shift75(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift78(a, b uint64) (u, v uint64) {
	b = mod_shift78(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift81(a, b uint64) (u, v uint64) {
	b = mod_shift81(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift84(a, b uint64) (u, v uint64) {
	b = mod_shift84(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift87(a, b uint64) (u, v uint64) {
	b = mod_shift87(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift90(a, b uint64) (u, v uint64) {
	b = mod_shift90(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

// This effectively shifts shift+96
// Note signs reversed in the butterfly since 2^96 mod p = -1
func invbutterfly_shift93(a, b uint64) (u, v uint64) {
	b = mod_shift93(b)
	u = mod_sub(a, b)
	v = mod_add(a, b)
	return
}

func invbutterfly_mul(a, b uint64, w uint64) (u, v uint64) {
	b = mod_mul(b, w)
	u = mod_add(a, b)
	v = mod_sub(a, b)
	return
}

// Fft for size 2**0
//
// This is an in place FFT with a bit reversed output
func fft0(x []uint64) {

}

// InvFft for size 2**0
//
// This is an in place Inverse FFT with a bit reversed input
func invfft0(x []uint64) {

}

// Fft for size 2**1
//
// This is an in place FFT with a bit reversed output
func fft1(x []uint64) {
	x[0], x[1] = butterfly_null(x[0], x[1])

}

// InvFft for size 2**1
//
// This is an in place Inverse FFT with a bit reversed input
func invfft1(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])

}

// Fft for size 2**2
//
// This is an in place FFT with a bit reversed output
func fft2(x []uint64) {
	x[0], x[2] = butterfly_null(x[0], x[2])
	x[1], x[3] = butterfly_shift48(x[1], x[3])
	x[0], x[1] = butterfly_null(x[0], x[1])
	x[2], x[3] = butterfly_null(x[2], x[3])

}

// InvFft for size 2**2
//
// This is an in place Inverse FFT with a bit reversed input
func invfft2(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])
	x[2], x[3] = invbutterfly_null(x[2], x[3])
	x[0], x[2] = invbutterfly_null(x[0], x[2])
	x[1], x[3] = invbutterfly_shift48(x[1], x[3])

}

// Fft for size 2**3
//
// This is an in place FFT with a bit reversed output
func fft3(x []uint64) {
	x[0], x[4] = butterfly_null(x[0], x[4])
	x[1], x[5] = butterfly_shift24(x[1], x[5])
	x[2], x[6] = butterfly_shift48(x[2], x[6])
	x[3], x[7] = butterfly_shift72(x[3], x[7])
	x[0], x[2] = butterfly_null(x[0], x[2])
	x[4], x[6] = butterfly_null(x[4], x[6])
	x[1], x[3] = butterfly_shift48(x[1], x[3])
	x[5], x[7] = butterfly_shift48(x[5], x[7])
	x[0], x[1] = butterfly_null(x[0], x[1])
	x[2], x[3] = butterfly_null(x[2], x[3])
	x[4], x[5] = butterfly_null(x[4], x[5])
	x[6], x[7] = butterfly_null(x[6], x[7])

}

// InvFft for size 2**3
//
// This is an in place Inverse FFT with a bit reversed input
func invfft3(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])
	x[2], x[3] = invbutterfly_null(x[2], x[3])
	x[4], x[5] = invbutterfly_null(x[4], x[5])
	x[6], x[7] = invbutterfly_null(x[6], x[7])
	x[0], x[2] = invbutterfly_null(x[0], x[2])
	x[4], x[6] = invbutterfly_null(x[4], x[6])
	x[1], x[3] = invbutterfly_shift48(x[1], x[3])
	x[5], x[7] = invbutterfly_shift48(x[5], x[7])
	x[0], x[4] = invbutterfly_null(x[0], x[4])
	x[1], x[5] = invbutterfly_shift72(x[1], x[5])
	x[2], x[6] = invbutterfly_shift48(x[2], x[6])
	x[3], x[7] = invbutterfly_shift24(x[3], x[7])

}

// Fft for size 2**4
//
// This is an in place FFT with a bit reversed output
func fft4(x []uint64) {
	x[0], x[8] = butterfly_null(x[0], x[8])
	x[1], x[9] = butterfly_shift12(x[1], x[9])
	x[2], x[10] = butterfly_shift24(x[2], x[10])
	x[3], x[11] = butterfly_shift36(x[3], x[11])
	x[4], x[12] = butterfly_shift48(x[4], x[12])
	x[5], x[13] = butterfly_shift60(x[5], x[13])
	x[6], x[14] = butterfly_shift72(x[6], x[14])
	x[7], x[15] = butterfly_shift84(x[7], x[15])
	x[0], x[4] = butterfly_null(x[0], x[4])
	x[8], x[12] = butterfly_null(x[8], x[12])
	x[1], x[5] = butterfly_shift24(x[1], x[5])
	x[9], x[13] = butterfly_shift24(x[9], x[13])
	x[2], x[6] = butterfly_shift48(x[2], x[6])
	x[10], x[14] = butterfly_shift48(x[10], x[14])
	x[3], x[7] = butterfly_shift72(x[3], x[7])
	x[11], x[15] = butterfly_shift72(x[11], x[15])
	x[0], x[2] = butterfly_null(x[0], x[2])
	x[4], x[6] = butterfly_null(x[4], x[6])
	x[8], x[10] = butterfly_null(x[8], x[10])
	x[12], x[14] = butterfly_null(x[12], x[14])
	x[1], x[3] = butterfly_shift48(x[1], x[3])
	x[5], x[7] = butterfly_shift48(x[5], x[7])
	x[9], x[11] = butterfly_shift48(x[9], x[11])
	x[13], x[15] = butterfly_shift48(x[13], x[15])
	x[0], x[1] = butterfly_null(x[0], x[1])
	x[2], x[3] = butterfly_null(x[2], x[3])
	x[4], x[5] = butterfly_null(x[4], x[5])
	x[6], x[7] = butterfly_null(x[6], x[7])
	x[8], x[9] = butterfly_null(x[8], x[9])
	x[10], x[11] = butterfly_null(x[10], x[11])
	x[12], x[13] = butterfly_null(x[12], x[13])
	x[14], x[15] = butterfly_null(x[14], x[15])

}

// InvFft for size 2**4
//
// This is an in place Inverse FFT with a bit reversed input
func invfft4(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])
	x[2], x[3] = invbutterfly_null(x[2], x[3])
	x[4], x[5] = invbutterfly_null(x[4], x[5])
	x[6], x[7] = invbutterfly_null(x[6], x[7])
	x[8], x[9] = invbutterfly_null(x[8], x[9])
	x[10], x[11] = invbutterfly_null(x[10], x[11])
	x[12], x[13] = invbutterfly_null(x[12], x[13])
	x[14], x[15] = invbutterfly_null(x[14], x[15])
	x[0], x[2] = invbutterfly_null(x[0], x[2])
	x[4], x[6] = invbutterfly_null(x[4], x[6])
	x[8], x[10] = invbutterfly_null(x[8], x[10])
	x[12], x[14] = invbutterfly_null(x[12], x[14])
	x[1], x[3] = invbutterfly_shift48(x[1], x[3])
	x[5], x[7] = invbutterfly_shift48(x[5], x[7])
	x[9], x[11] = invbutterfly_shift48(x[9], x[11])
	x[13], x[15] = invbutterfly_shift48(x[13], x[15])
	x[0], x[4] = invbutterfly_null(x[0], x[4])
	x[8], x[12] = invbutterfly_null(x[8], x[12])
	x[1], x[5] = invbutterfly_shift72(x[1], x[5])
	x[9], x[13] = invbutterfly_shift72(x[9], x[13])
	x[2], x[6] = invbutterfly_shift48(x[2], x[6])
	x[10], x[14] = invbutterfly_shift48(x[10], x[14])
	x[3], x[7] = invbutterfly_shift24(x[3], x[7])
	x[11], x[15] = invbutterfly_shift24(x[11], x[15])
	x[0], x[8] = invbutterfly_null(x[0], x[8])
	x[1], x[9] = invbutterfly_shift84(x[1], x[9])
	x[2], x[10] = invbutterfly_shift72(x[2], x[10])
	x[3], x[11] = invbutterfly_shift60(x[3], x[11])
	x[4], x[12] = invbutterfly_shift48(x[4], x[12])
	x[5], x[13] = invbutterfly_shift36(x[5], x[13])
	x[6], x[14] = invbutterfly_shift24(x[6], x[14])
	x[7], x[15] = invbutterfly_shift12(x[7], x[15])

}

// Fft for size 2**5
//
// This is an in place FFT with a bit reversed output
func fft5(x []uint64) {
	x[0], x[16] = butterfly_null(x[0], x[16])
	x[1], x[17] = butterfly_shift6(x[1], x[17])
	x[2], x[18] = butterfly_shift12(x[2], x[18])
	x[3], x[19] = butterfly_shift18(x[3], x[19])
	x[4], x[20] = butterfly_shift24(x[4], x[20])
	x[5], x[21] = butterfly_shift30(x[5], x[21])
	x[6], x[22] = butterfly_shift36(x[6], x[22])
	x[7], x[23] = butterfly_shift42(x[7], x[23])
	x[8], x[24] = butterfly_shift48(x[8], x[24])
	x[9], x[25] = butterfly_shift54(x[9], x[25])
	x[10], x[26] = butterfly_shift60(x[10], x[26])
	x[11], x[27] = butterfly_shift66(x[11], x[27])
	x[12], x[28] = butterfly_shift72(x[12], x[28])
	x[13], x[29] = butterfly_shift78(x[13], x[29])
	x[14], x[30] = butterfly_shift84(x[14], x[30])
	x[15], x[31] = butterfly_shift90(x[15], x[31])
	x[0], x[8] = butterfly_null(x[0], x[8])
	x[16], x[24] = butterfly_null(x[16], x[24])
	x[1], x[9] = butterfly_shift12(x[1], x[9])
	x[17], x[25] = butterfly_shift12(x[17], x[25])
	x[2], x[10] = butterfly_shift24(x[2], x[10])
	x[18], x[26] = butterfly_shift24(x[18], x[26])
	x[3], x[11] = butterfly_shift36(x[3], x[11])
	x[19], x[27] = butterfly_shift36(x[19], x[27])
	x[4], x[12] = butterfly_shift48(x[4], x[12])
	x[20], x[28] = butterfly_shift48(x[20], x[28])
	x[5], x[13] = butterfly_shift60(x[5], x[13])
	x[21], x[29] = butterfly_shift60(x[21], x[29])
	x[6], x[14] = butterfly_shift72(x[6], x[14])
	x[22], x[30] = butterfly_shift72(x[22], x[30])
	x[7], x[15] = butterfly_shift84(x[7], x[15])
	x[23], x[31] = butterfly_shift84(x[23], x[31])
	x[0], x[4] = butterfly_null(x[0], x[4])
	x[8], x[12] = butterfly_null(x[8], x[12])
	x[16], x[20] = butterfly_null(x[16], x[20])
	x[24], x[28] = butterfly_null(x[24], x[28])
	x[1], x[5] = butterfly_shift24(x[1], x[5])
	x[9], x[13] = butterfly_shift24(x[9], x[13])
	x[17], x[21] = butterfly_shift24(x[17], x[21])
	x[25], x[29] = butterfly_shift24(x[25], x[29])
	x[2], x[6] = butterfly_shift48(x[2], x[6])
	x[10], x[14] = butterfly_shift48(x[10], x[14])
	x[18], x[22] = butterfly_shift48(x[18], x[22])
	x[26], x[30] = butterfly_shift48(x[26], x[30])
	x[3], x[7] = butterfly_shift72(x[3], x[7])
	x[11], x[15] = butterfly_shift72(x[11], x[15])
	x[19], x[23] = butterfly_shift72(x[19], x[23])
	x[27], x[31] = butterfly_shift72(x[27], x[31])
	x[0], x[2] = butterfly_null(x[0], x[2])
	x[4], x[6] = butterfly_null(x[4], x[6])
	x[8], x[10] = butterfly_null(x[8], x[10])
	x[12], x[14] = butterfly_null(x[12], x[14])
	x[16], x[18] = butterfly_null(x[16], x[18])
	x[20], x[22] = butterfly_null(x[20], x[22])
	x[24], x[26] = butterfly_null(x[24], x[26])
	x[28], x[30] = butterfly_null(x[28], x[30])
	x[1], x[3] = butterfly_shift48(x[1], x[3])
	x[5], x[7] = butterfly_shift48(x[5], x[7])
	x[9], x[11] = butterfly_shift48(x[9], x[11])
	x[13], x[15] = butterfly_shift48(x[13], x[15])
	x[17], x[19] = butterfly_shift48(x[17], x[19])
	x[21], x[23] = butterfly_shift48(x[21], x[23])
	x[25], x[27] = butterfly_shift48(x[25], x[27])
	x[29], x[31] = butterfly_shift48(x[29], x[31])
	x[0], x[1] = butterfly_null(x[0], x[1])
	x[2], x[3] = butterfly_null(x[2], x[3])
	x[4], x[5] = butterfly_null(x[4], x[5])
	x[6], x[7] = butterfly_null(x[6], x[7])
	x[8], x[9] = butterfly_null(x[8], x[9])
	x[10], x[11] = butterfly_null(x[10], x[11])
	x[12], x[13] = butterfly_null(x[12], x[13])
	x[14], x[15] = butterfly_null(x[14], x[15])
	x[16], x[17] = butterfly_null(x[16], x[17])
	x[18], x[19] = butterfly_null(x[18], x[19])
	x[20], x[21] = butterfly_null(x[20], x[21])
	x[22], x[23] = butterfly_null(x[22], x[23])
	x[24], x[25] = butterfly_null(x[24], x[25])
	x[26], x[27] = butterfly_null(x[26], x[27])
	x[28], x[29] = butterfly_null(x[28], x[29])
	x[30], x[31] = butterfly_null(x[30], x[31])

}

// InvFft for size 2**5
//
// This is an in place Inverse FFT with a bit reversed input
func invfft5(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])
	x[2], x[3] = invbutterfly_null(x[2], x[3])
	x[4], x[5] = invbutterfly_null(x[4], x[5])
	x[6], x[7] = invbutterfly_null(x[6], x[7])
	x[8], x[9] = invbutterfly_null(x[8], x[9])
	x[10], x[11] = invbutterfly_null(x[10], x[11])
	x[12], x[13] = invbutterfly_null(x[12], x[13])
	x[14], x[15] = invbutterfly_null(x[14], x[15])
	x[16], x[17] = invbutterfly_null(x[16], x[17])
	x[18], x[19] = invbutterfly_null(x[18], x[19])
	x[20], x[21] = invbutterfly_null(x[20], x[21])
	x[22], x[23] = invbutterfly_null(x[22], x[23])
	x[24], x[25] = invbutterfly_null(x[24], x[25])
	x[26], x[27] = invbutterfly_null(x[26], x[27])
	x[28], x[29] = invbutterfly_null(x[28], x[29])
	x[30], x[31] = invbutterfly_null(x[30], x[31])
	x[0], x[2] = invbutterfly_null(x[0], x[2])
	x[4], x[6] = invbutterfly_null(x[4], x[6])
	x[8], x[10] = invbutterfly_null(x[8], x[10])
	x[12], x[14] = invbutterfly_null(x[12], x[14])
	x[16], x[18] = invbutterfly_null(x[16], x[18])
	x[20], x[22] = invbutterfly_null(x[20], x[22])
	x[24], x[26] = invbutterfly_null(x[24], x[26])
	x[28], x[30] = invbutterfly_null(x[28], x[30])
	x[1], x[3] = invbutterfly_shift48(x[1], x[3])
	x[5], x[7] = invbutterfly_shift48(x[5], x[7])
	x[9], x[11] = invbutterfly_shift48(x[9], x[11])
	x[13], x[15] = invbutterfly_shift48(x[13], x[15])
	x[17], x[19] = invbutterfly_shift48(x[17], x[19])
	x[21], x[23] = invbutterfly_shift48(x[21], x[23])
	x[25], x[27] = invbutterfly_shift48(x[25], x[27])
	x[29], x[31] = invbutterfly_shift48(x[29], x[31])
	x[0], x[4] = invbutterfly_null(x[0], x[4])
	x[8], x[12] = invbutterfly_null(x[8], x[12])
	x[16], x[20] = invbutterfly_null(x[16], x[20])
	x[24], x[28] = invbutterfly_null(x[24], x[28])
	x[1], x[5] = invbutterfly_shift72(x[1], x[5])
	x[9], x[13] = invbutterfly_shift72(x[9], x[13])
	x[17], x[21] = invbutterfly_shift72(x[17], x[21])
	x[25], x[29] = invbutterfly_shift72(x[25], x[29])
	x[2], x[6] = invbutterfly_shift48(x[2], x[6])
	x[10], x[14] = invbutterfly_shift48(x[10], x[14])
	x[18], x[22] = invbutterfly_shift48(x[18], x[22])
	x[26], x[30] = invbutterfly_shift48(x[26], x[30])
	x[3], x[7] = invbutterfly_shift24(x[3], x[7])
	x[11], x[15] = invbutterfly_shift24(x[11], x[15])
	x[19], x[23] = invbutterfly_shift24(x[19], x[23])
	x[27], x[31] = invbutterfly_shift24(x[27], x[31])
	x[0], x[8] = invbutterfly_null(x[0], x[8])
	x[16], x[24] = invbutterfly_null(x[16], x[24])
	x[1], x[9] = invbutterfly_shift84(x[1], x[9])
	x[17], x[25] = invbutterfly_shift84(x[17], x[25])
	x[2], x[10] = invbutterfly_shift72(x[2], x[10])
	x[18], x[26] = invbutterfly_shift72(x[18], x[26])
	x[3], x[11] = invbutterfly_shift60(x[3], x[11])
	x[19], x[27] = invbutterfly_shift60(x[19], x[27])
	x[4], x[12] = invbutterfly_shift48(x[4], x[12])
	x[20], x[28] = invbutterfly_shift48(x[20], x[28])
	x[5], x[13] = invbutterfly_shift36(x[5], x[13])
	x[21], x[29] = invbutterfly_shift36(x[21], x[29])
	x[6], x[14] = invbutterfly_shift24(x[6], x[14])
	x[22], x[30] = invbutterfly_shift24(x[22], x[30])
	x[7], x[15] = invbutterfly_shift12(x[7], x[15])
	x[23], x[31] = invbutterfly_shift12(x[23], x[31])
	x[0], x[16] = invbutterfly_null(x[0], x[16])
	x[1], x[17] = invbutterfly_shift90(x[1], x[17])
	x[2], x[18] = invbutterfly_shift84(x[2], x[18])
	x[3], x[19] = invbutterfly_shift78(x[3], x[19])
	x[4], x[20] = invbutterfly_shift72(x[4], x[20])
	x[5], x[21] = invbutterfly_shift66(x[5], x[21])
	x[6], x[22] = invbutterfly_shift60(x[6], x[22])
	x[7], x[23] = invbutterfly_shift54(x[7], x[23])
	x[8], x[24] = invbutterfly_shift48(x[8], x[24])
	x[9], x[25] = invbutterfly_shift42(x[9], x[25])
	x[10], x[26] = invbutterfly_shift36(x[10], x[26])
	x[11], x[27] = invbutterfly_shift30(x[11], x[27])
	x[12], x[28] = invbutterfly_shift24(x[12], x[28])
	x[13], x[29] = invbutterfly_shift18(x[13], x[29])
	x[14], x[30] = invbutterfly_shift12(x[14], x[30])
	x[15], x[31] = invbutterfly_shift6(x[15], x[31])

}

// Fft for size 2**6
//
// This is an in place FFT with a bit reversed output
func fft6(x []uint64) {
	x[0], x[32] = butterfly_null(x[0], x[32])
	x[1], x[33] = butterfly_shift3(x[1], x[33])
	x[2], x[34] = butterfly_shift6(x[2], x[34])
	x[3], x[35] = butterfly_shift9(x[3], x[35])
	x[4], x[36] = butterfly_shift12(x[4], x[36])
	x[5], x[37] = butterfly_shift15(x[5], x[37])
	x[6], x[38] = butterfly_shift18(x[6], x[38])
	x[7], x[39] = butterfly_shift21(x[7], x[39])
	x[8], x[40] = butterfly_shift24(x[8], x[40])
	x[9], x[41] = butterfly_shift27(x[9], x[41])
	x[10], x[42] = butterfly_shift30(x[10], x[42])
	x[11], x[43] = butterfly_shift33(x[11], x[43])
	x[12], x[44] = butterfly_shift36(x[12], x[44])
	x[13], x[45] = butterfly_shift39(x[13], x[45])
	x[14], x[46] = butterfly_shift42(x[14], x[46])
	x[15], x[47] = butterfly_shift45(x[15], x[47])
	x[16], x[48] = butterfly_shift48(x[16], x[48])
	x[17], x[49] = butterfly_shift51(x[17], x[49])
	x[18], x[50] = butterfly_shift54(x[18], x[50])
	x[19], x[51] = butterfly_shift57(x[19], x[51])
	x[20], x[52] = butterfly_shift60(x[20], x[52])
	x[21], x[53] = butterfly_shift63(x[21], x[53])
	x[22], x[54] = butterfly_shift66(x[22], x[54])
	x[23], x[55] = butterfly_shift69(x[23], x[55])
	x[24], x[56] = butterfly_shift72(x[24], x[56])
	x[25], x[57] = butterfly_shift75(x[25], x[57])
	x[26], x[58] = butterfly_shift78(x[26], x[58])
	x[27], x[59] = butterfly_shift81(x[27], x[59])
	x[28], x[60] = butterfly_shift84(x[28], x[60])
	x[29], x[61] = butterfly_shift87(x[29], x[61])
	x[30], x[62] = butterfly_shift90(x[30], x[62])
	x[31], x[63] = butterfly_shift93(x[31], x[63])
	x[0], x[16] = butterfly_null(x[0], x[16])
	x[32], x[48] = butterfly_null(x[32], x[48])
	x[1], x[17] = butterfly_shift6(x[1], x[17])
	x[33], x[49] = butterfly_shift6(x[33], x[49])
	x[2], x[18] = butterfly_shift12(x[2], x[18])
	x[34], x[50] = butterfly_shift12(x[34], x[50])
	x[3], x[19] = butterfly_shift18(x[3], x[19])
	x[35], x[51] = butterfly_shift18(x[35], x[51])
	x[4], x[20] = butterfly_shift24(x[4], x[20])
	x[36], x[52] = butterfly_shift24(x[36], x[52])
	x[5], x[21] = butterfly_shift30(x[5], x[21])
	x[37], x[53] = butterfly_shift30(x[37], x[53])
	x[6], x[22] = butterfly_shift36(x[6], x[22])
	x[38], x[54] = butterfly_shift36(x[38], x[54])
	x[7], x[23] = butterfly_shift42(x[7], x[23])
	x[39], x[55] = butterfly_shift42(x[39], x[55])
	x[8], x[24] = butterfly_shift48(x[8], x[24])
	x[40], x[56] = butterfly_shift48(x[40], x[56])
	x[9], x[25] = butterfly_shift54(x[9], x[25])
	x[41], x[57] = butterfly_shift54(x[41], x[57])
	x[10], x[26] = butterfly_shift60(x[10], x[26])
	x[42], x[58] = butterfly_shift60(x[42], x[58])
	x[11], x[27] = butterfly_shift66(x[11], x[27])
	x[43], x[59] = butterfly_shift66(x[43], x[59])
	x[12], x[28] = butterfly_shift72(x[12], x[28])
	x[44], x[60] = butterfly_shift72(x[44], x[60])
	x[13], x[29] = butterfly_shift78(x[13], x[29])
	x[45], x[61] = butterfly_shift78(x[45], x[61])
	x[14], x[30] = butterfly_shift84(x[14], x[30])
	x[46], x[62] = butterfly_shift84(x[46], x[62])
	x[15], x[31] = butterfly_shift90(x[15], x[31])
	x[47], x[63] = butterfly_shift90(x[47], x[63])
	x[0], x[8] = butterfly_null(x[0], x[8])
	x[16], x[24] = butterfly_null(x[16], x[24])
	x[32], x[40] = butterfly_null(x[32], x[40])
	x[48], x[56] = butterfly_null(x[48], x[56])
	x[1], x[9] = butterfly_shift12(x[1], x[9])
	x[17], x[25] = butterfly_shift12(x[17], x[25])
	x[33], x[41] = butterfly_shift12(x[33], x[41])
	x[49], x[57] = butterfly_shift12(x[49], x[57])
	x[2], x[10] = butterfly_shift24(x[2], x[10])
	x[18], x[26] = butterfly_shift24(x[18], x[26])
	x[34], x[42] = butterfly_shift24(x[34], x[42])
	x[50], x[58] = butterfly_shift24(x[50], x[58])
	x[3], x[11] = butterfly_shift36(x[3], x[11])
	x[19], x[27] = butterfly_shift36(x[19], x[27])
	x[35], x[43] = butterfly_shift36(x[35], x[43])
	x[51], x[59] = butterfly_shift36(x[51], x[59])
	x[4], x[12] = butterfly_shift48(x[4], x[12])
	x[20], x[28] = butterfly_shift48(x[20], x[28])
	x[36], x[44] = butterfly_shift48(x[36], x[44])
	x[52], x[60] = butterfly_shift48(x[52], x[60])
	x[5], x[13] = butterfly_shift60(x[5], x[13])
	x[21], x[29] = butterfly_shift60(x[21], x[29])
	x[37], x[45] = butterfly_shift60(x[37], x[45])
	x[53], x[61] = butterfly_shift60(x[53], x[61])
	x[6], x[14] = butterfly_shift72(x[6], x[14])
	x[22], x[30] = butterfly_shift72(x[22], x[30])
	x[38], x[46] = butterfly_shift72(x[38], x[46])
	x[54], x[62] = butterfly_shift72(x[54], x[62])
	x[7], x[15] = butterfly_shift84(x[7], x[15])
	x[23], x[31] = butterfly_shift84(x[23], x[31])
	x[39], x[47] = butterfly_shift84(x[39], x[47])
	x[55], x[63] = butterfly_shift84(x[55], x[63])
	x[0], x[4] = butterfly_null(x[0], x[4])
	x[8], x[12] = butterfly_null(x[8], x[12])
	x[16], x[20] = butterfly_null(x[16], x[20])
	x[24], x[28] = butterfly_null(x[24], x[28])
	x[32], x[36] = butterfly_null(x[32], x[36])
	x[40], x[44] = butterfly_null(x[40], x[44])
	x[48], x[52] = butterfly_null(x[48], x[52])
	x[56], x[60] = butterfly_null(x[56], x[60])
	x[1], x[5] = butterfly_shift24(x[1], x[5])
	x[9], x[13] = butterfly_shift24(x[9], x[13])
	x[17], x[21] = butterfly_shift24(x[17], x[21])
	x[25], x[29] = butterfly_shift24(x[25], x[29])
	x[33], x[37] = butterfly_shift24(x[33], x[37])
	x[41], x[45] = butterfly_shift24(x[41], x[45])
	x[49], x[53] = butterfly_shift24(x[49], x[53])
	x[57], x[61] = butterfly_shift24(x[57], x[61])
	x[2], x[6] = butterfly_shift48(x[2], x[6])
	x[10], x[14] = butterfly_shift48(x[10], x[14])
	x[18], x[22] = butterfly_shift48(x[18], x[22])
	x[26], x[30] = butterfly_shift48(x[26], x[30])
	x[34], x[38] = butterfly_shift48(x[34], x[38])
	x[42], x[46] = butterfly_shift48(x[42], x[46])
	x[50], x[54] = butterfly_shift48(x[50], x[54])
	x[58], x[62] = butterfly_shift48(x[58], x[62])
	x[3], x[7] = butterfly_shift72(x[3], x[7])
	x[11], x[15] = butterfly_shift72(x[11], x[15])
	x[19], x[23] = butterfly_shift72(x[19], x[23])
	x[27], x[31] = butterfly_shift72(x[27], x[31])
	x[35], x[39] = butterfly_shift72(x[35], x[39])
	x[43], x[47] = butterfly_shift72(x[43], x[47])
	x[51], x[55] = butterfly_shift72(x[51], x[55])
	x[59], x[63] = butterfly_shift72(x[59], x[63])
	x[0], x[2] = butterfly_null(x[0], x[2])
	x[4], x[6] = butterfly_null(x[4], x[6])
	x[8], x[10] = butterfly_null(x[8], x[10])
	x[12], x[14] = butterfly_null(x[12], x[14])
	x[16], x[18] = butterfly_null(x[16], x[18])
	x[20], x[22] = butterfly_null(x[20], x[22])
	x[24], x[26] = butterfly_null(x[24], x[26])
	x[28], x[30] = butterfly_null(x[28], x[30])
	x[32], x[34] = butterfly_null(x[32], x[34])
	x[36], x[38] = butterfly_null(x[36], x[38])
	x[40], x[42] = butterfly_null(x[40], x[42])
	x[44], x[46] = butterfly_null(x[44], x[46])
	x[48], x[50] = butterfly_null(x[48], x[50])
	x[52], x[54] = butterfly_null(x[52], x[54])
	x[56], x[58] = butterfly_null(x[56], x[58])
	x[60], x[62] = butterfly_null(x[60], x[62])
	x[1], x[3] = butterfly_shift48(x[1], x[3])
	x[5], x[7] = butterfly_shift48(x[5], x[7])
	x[9], x[11] = butterfly_shift48(x[9], x[11])
	x[13], x[15] = butterfly_shift48(x[13], x[15])
	x[17], x[19] = butterfly_shift48(x[17], x[19])
	x[21], x[23] = butterfly_shift48(x[21], x[23])
	x[25], x[27] = butterfly_shift48(x[25], x[27])
	x[29], x[31] = butterfly_shift48(x[29], x[31])
	x[33], x[35] = butterfly_shift48(x[33], x[35])
	x[37], x[39] = butterfly_shift48(x[37], x[39])
	x[41], x[43] = butterfly_shift48(x[41], x[43])
	x[45], x[47] = butterfly_shift48(x[45], x[47])
	x[49], x[51] = butterfly_shift48(x[49], x[51])
	x[53], x[55] = butterfly_shift48(x[53], x[55])
	x[57], x[59] = butterfly_shift48(x[57], x[59])
	x[61], x[63] = butterfly_shift48(x[61], x[63])
	x[0], x[1] = butterfly_null(x[0], x[1])
	x[2], x[3] = butterfly_null(x[2], x[3])
	x[4], x[5] = butterfly_null(x[4], x[5])
	x[6], x[7] = butterfly_null(x[6], x[7])
	x[8], x[9] = butterfly_null(x[8], x[9])
	x[10], x[11] = butterfly_null(x[10], x[11])
	x[12], x[13] = butterfly_null(x[12], x[13])
	x[14], x[15] = butterfly_null(x[14], x[15])
	x[16], x[17] = butterfly_null(x[16], x[17])
	x[18], x[19] = butterfly_null(x[18], x[19])
	x[20], x[21] = butterfly_null(x[20], x[21])
	x[22], x[23] = butterfly_null(x[22], x[23])
	x[24], x[25] = butterfly_null(x[24], x[25])
	x[26], x[27] = butterfly_null(x[26], x[27])
	x[28], x[29] = butterfly_null(x[28], x[29])
	x[30], x[31] = butterfly_null(x[30], x[31])
	x[32], x[33] = butterfly_null(x[32], x[33])
	x[34], x[35] = butterfly_null(x[34], x[35])
	x[36], x[37] = butterfly_null(x[36], x[37])
	x[38], x[39] = butterfly_null(x[38], x[39])
	x[40], x[41] = butterfly_null(x[40], x[41])
	x[42], x[43] = butterfly_null(x[42], x[43])
	x[44], x[45] = butterfly_null(x[44], x[45])
	x[46], x[47] = butterfly_null(x[46], x[47])
	x[48], x[49] = butterfly_null(x[48], x[49])
	x[50], x[51] = butterfly_null(x[50], x[51])
	x[52], x[53] = butterfly_null(x[52], x[53])
	x[54], x[55] = butterfly_null(x[54], x[55])
	x[56], x[57] = butterfly_null(x[56], x[57])
	x[58], x[59] = butterfly_null(x[58], x[59])
	x[60], x[61] = butterfly_null(x[60], x[61])
	x[62], x[63] = butterfly_null(x[62], x[63])

}

// InvFft for size 2**6
//
// This is an in place Inverse FFT with a bit reversed input
func invfft6(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])
	x[2], x[3] = invbutterfly_null(x[2], x[3])
	x[4], x[5] = invbutterfly_null(x[4], x[5])
	x[6], x[7] = invbutterfly_null(x[6], x[7])
	x[8], x[9] = invbutterfly_null(x[8], x[9])
	x[10], x[11] = invbutterfly_null(x[10], x[11])
	x[12], x[13] = invbutterfly_null(x[12], x[13])
	x[14], x[15] = invbutterfly_null(x[14], x[15])
	x[16], x[17] = invbutterfly_null(x[16], x[17])
	x[18], x[19] = invbutterfly_null(x[18], x[19])
	x[20], x[21] = invbutterfly_null(x[20], x[21])
	x[22], x[23] = invbutterfly_null(x[22], x[23])
	x[24], x[25] = invbutterfly_null(x[24], x[25])
	x[26], x[27] = invbutterfly_null(x[26], x[27])
	x[28], x[29] = invbutterfly_null(x[28], x[29])
	x[30], x[31] = invbutterfly_null(x[30], x[31])
	x[32], x[33] = invbutterfly_null(x[32], x[33])
	x[34], x[35] = invbutterfly_null(x[34], x[35])
	x[36], x[37] = invbutterfly_null(x[36], x[37])
	x[38], x[39] = invbutterfly_null(x[38], x[39])
	x[40], x[41] = invbutterfly_null(x[40], x[41])
	x[42], x[43] = invbutterfly_null(x[42], x[43])
	x[44], x[45] = invbutterfly_null(x[44], x[45])
	x[46], x[47] = invbutterfly_null(x[46], x[47])
	x[48], x[49] = invbutterfly_null(x[48], x[49])
	x[50], x[51] = invbutterfly_null(x[50], x[51])
	x[52], x[53] = invbutterfly_null(x[52], x[53])
	x[54], x[55] = invbutterfly_null(x[54], x[55])
	x[56], x[57] = invbutterfly_null(x[56], x[57])
	x[58], x[59] = invbutterfly_null(x[58], x[59])
	x[60], x[61] = invbutterfly_null(x[60], x[61])
	x[62], x[63] = invbutterfly_null(x[62], x[63])
	x[0], x[2] = invbutterfly_null(x[0], x[2])
	x[4], x[6] = invbutterfly_null(x[4], x[6])
	x[8], x[10] = invbutterfly_null(x[8], x[10])
	x[12], x[14] = invbutterfly_null(x[12], x[14])
	x[16], x[18] = invbutterfly_null(x[16], x[18])
	x[20], x[22] = invbutterfly_null(x[20], x[22])
	x[24], x[26] = invbutterfly_null(x[24], x[26])
	x[28], x[30] = invbutterfly_null(x[28], x[30])
	x[32], x[34] = invbutterfly_null(x[32], x[34])
	x[36], x[38] = invbutterfly_null(x[36], x[38])
	x[40], x[42] = invbutterfly_null(x[40], x[42])
	x[44], x[46] = invbutterfly_null(x[44], x[46])
	x[48], x[50] = invbutterfly_null(x[48], x[50])
	x[52], x[54] = invbutterfly_null(x[52], x[54])
	x[56], x[58] = invbutterfly_null(x[56], x[58])
	x[60], x[62] = invbutterfly_null(x[60], x[62])
	x[1], x[3] = invbutterfly_shift48(x[1], x[3])
	x[5], x[7] = invbutterfly_shift48(x[5], x[7])
	x[9], x[11] = invbutterfly_shift48(x[9], x[11])
	x[13], x[15] = invbutterfly_shift48(x[13], x[15])
	x[17], x[19] = invbutterfly_shift48(x[17], x[19])
	x[21], x[23] = invbutterfly_shift48(x[21], x[23])
	x[25], x[27] = invbutterfly_shift48(x[25], x[27])
	x[29], x[31] = invbutterfly_shift48(x[29], x[31])
	x[33], x[35] = invbutterfly_shift48(x[33], x[35])
	x[37], x[39] = invbutterfly_shift48(x[37], x[39])
	x[41], x[43] = invbutterfly_shift48(x[41], x[43])
	x[45], x[47] = invbutterfly_shift48(x[45], x[47])
	x[49], x[51] = invbutterfly_shift48(x[49], x[51])
	x[53], x[55] = invbutterfly_shift48(x[53], x[55])
	x[57], x[59] = invbutterfly_shift48(x[57], x[59])
	x[61], x[63] = invbutterfly_shift48(x[61], x[63])
	x[0], x[4] = invbutterfly_null(x[0], x[4])
	x[8], x[12] = invbutterfly_null(x[8], x[12])
	x[16], x[20] = invbutterfly_null(x[16], x[20])
	x[24], x[28] = invbutterfly_null(x[24], x[28])
	x[32], x[36] = invbutterfly_null(x[32], x[36])
	x[40], x[44] = invbutterfly_null(x[40], x[44])
	x[48], x[52] = invbutterfly_null(x[48], x[52])
	x[56], x[60] = invbutterfly_null(x[56], x[60])
	x[1], x[5] = invbutterfly_shift72(x[1], x[5])
	x[9], x[13] = invbutterfly_shift72(x[9], x[13])
	x[17], x[21] = invbutterfly_shift72(x[17], x[21])
	x[25], x[29] = invbutterfly_shift72(x[25], x[29])
	x[33], x[37] = invbutterfly_shift72(x[33], x[37])
	x[41], x[45] = invbutterfly_shift72(x[41], x[45])
	x[49], x[53] = invbutterfly_shift72(x[49], x[53])
	x[57], x[61] = invbutterfly_shift72(x[57], x[61])
	x[2], x[6] = invbutterfly_shift48(x[2], x[6])
	x[10], x[14] = invbutterfly_shift48(x[10], x[14])
	x[18], x[22] = invbutterfly_shift48(x[18], x[22])
	x[26], x[30] = invbutterfly_shift48(x[26], x[30])
	x[34], x[38] = invbutterfly_shift48(x[34], x[38])
	x[42], x[46] = invbutterfly_shift48(x[42], x[46])
	x[50], x[54] = invbutterfly_shift48(x[50], x[54])
	x[58], x[62] = invbutterfly_shift48(x[58], x[62])
	x[3], x[7] = invbutterfly_shift24(x[3], x[7])
	x[11], x[15] = invbutterfly_shift24(x[11], x[15])
	x[19], x[23] = invbutterfly_shift24(x[19], x[23])
	x[27], x[31] = invbutterfly_shift24(x[27], x[31])
	x[35], x[39] = invbutterfly_shift24(x[35], x[39])
	x[43], x[47] = invbutterfly_shift24(x[43], x[47])
	x[51], x[55] = invbutterfly_shift24(x[51], x[55])
	x[59], x[63] = invbutterfly_shift24(x[59], x[63])
	x[0], x[8] = invbutterfly_null(x[0], x[8])
	x[16], x[24] = invbutterfly_null(x[16], x[24])
	x[32], x[40] = invbutterfly_null(x[32], x[40])
	x[48], x[56] = invbutterfly_null(x[48], x[56])
	x[1], x[9] = invbutterfly_shift84(x[1], x[9])
	x[17], x[25] = invbutterfly_shift84(x[17], x[25])
	x[33], x[41] = invbutterfly_shift84(x[33], x[41])
	x[49], x[57] = invbutterfly_shift84(x[49], x[57])
	x[2], x[10] = invbutterfly_shift72(x[2], x[10])
	x[18], x[26] = invbutterfly_shift72(x[18], x[26])
	x[34], x[42] = invbutterfly_shift72(x[34], x[42])
	x[50], x[58] = invbutterfly_shift72(x[50], x[58])
	x[3], x[11] = invbutterfly_shift60(x[3], x[11])
	x[19], x[27] = invbutterfly_shift60(x[19], x[27])
	x[35], x[43] = invbutterfly_shift60(x[35], x[43])
	x[51], x[59] = invbutterfly_shift60(x[51], x[59])
	x[4], x[12] = invbutterfly_shift48(x[4], x[12])
	x[20], x[28] = invbutterfly_shift48(x[20], x[28])
	x[36], x[44] = invbutterfly_shift48(x[36], x[44])
	x[52], x[60] = invbutterfly_shift48(x[52], x[60])
	x[5], x[13] = invbutterfly_shift36(x[5], x[13])
	x[21], x[29] = invbutterfly_shift36(x[21], x[29])
	x[37], x[45] = invbutterfly_shift36(x[37], x[45])
	x[53], x[61] = invbutterfly_shift36(x[53], x[61])
	x[6], x[14] = invbutterfly_shift24(x[6], x[14])
	x[22], x[30] = invbutterfly_shift24(x[22], x[30])
	x[38], x[46] = invbutterfly_shift24(x[38], x[46])
	x[54], x[62] = invbutterfly_shift24(x[54], x[62])
	x[7], x[15] = invbutterfly_shift12(x[7], x[15])
	x[23], x[31] = invbutterfly_shift12(x[23], x[31])
	x[39], x[47] = invbutterfly_shift12(x[39], x[47])
	x[55], x[63] = invbutterfly_shift12(x[55], x[63])
	x[0], x[16] = invbutterfly_null(x[0], x[16])
	x[32], x[48] = invbutterfly_null(x[32], x[48])
	x[1], x[17] = invbutterfly_shift90(x[1], x[17])
	x[33], x[49] = invbutterfly_shift90(x[33], x[49])
	x[2], x[18] = invbutterfly_shift84(x[2], x[18])
	x[34], x[50] = invbutterfly_shift84(x[34], x[50])
	x[3], x[19] = invbutterfly_shift78(x[3], x[19])
	x[35], x[51] = invbutterfly_shift78(x[35], x[51])
	x[4], x[20] = invbutterfly_shift72(x[4], x[20])
	x[36], x[52] = invbutterfly_shift72(x[36], x[52])
	x[5], x[21] = invbutterfly_shift66(x[5], x[21])
	x[37], x[53] = invbutterfly_shift66(x[37], x[53])
	x[6], x[22] = invbutterfly_shift60(x[6], x[22])
	x[38], x[54] = invbutterfly_shift60(x[38], x[54])
	x[7], x[23] = invbutterfly_shift54(x[7], x[23])
	x[39], x[55] = invbutterfly_shift54(x[39], x[55])
	x[8], x[24] = invbutterfly_shift48(x[8], x[24])
	x[40], x[56] = invbutterfly_shift48(x[40], x[56])
	x[9], x[25] = invbutterfly_shift42(x[9], x[25])
	x[41], x[57] = invbutterfly_shift42(x[41], x[57])
	x[10], x[26] = invbutterfly_shift36(x[10], x[26])
	x[42], x[58] = invbutterfly_shift36(x[42], x[58])
	x[11], x[27] = invbutterfly_shift30(x[11], x[27])
	x[43], x[59] = invbutterfly_shift30(x[43], x[59])
	x[12], x[28] = invbutterfly_shift24(x[12], x[28])
	x[44], x[60] = invbutterfly_shift24(x[44], x[60])
	x[13], x[29] = invbutterfly_shift18(x[13], x[29])
	x[45], x[61] = invbutterfly_shift18(x[45], x[61])
	x[14], x[30] = invbutterfly_shift12(x[14], x[30])
	x[46], x[62] = invbutterfly_shift12(x[46], x[62])
	x[15], x[31] = invbutterfly_shift6(x[15], x[31])
	x[47], x[63] = invbutterfly_shift6(x[47], x[63])
	x[0], x[32] = invbutterfly_null(x[0], x[32])
	x[1], x[33] = invbutterfly_shift93(x[1], x[33])
	x[2], x[34] = invbutterfly_shift90(x[2], x[34])
	x[3], x[35] = invbutterfly_shift87(x[3], x[35])
	x[4], x[36] = invbutterfly_shift84(x[4], x[36])
	x[5], x[37] = invbutterfly_shift81(x[5], x[37])
	x[6], x[38] = invbutterfly_shift78(x[6], x[38])
	x[7], x[39] = invbutterfly_shift75(x[7], x[39])
	x[8], x[40] = invbutterfly_shift72(x[8], x[40])
	x[9], x[41] = invbutterfly_shift69(x[9], x[41])
	x[10], x[42] = invbutterfly_shift66(x[10], x[42])
	x[11], x[43] = invbutterfly_shift63(x[11], x[43])
	x[12], x[44] = invbutterfly_shift60(x[12], x[44])
	x[13], x[45] = invbutterfly_shift57(x[13], x[45])
	x[14], x[46] = invbutterfly_shift54(x[14], x[46])
	x[15], x[47] = invbutterfly_shift51(x[15], x[47])
	x[16], x[48] = invbutterfly_shift48(x[16], x[48])
	x[17], x[49] = invbutterfly_shift45(x[17], x[49])
	x[18], x[50] = invbutterfly_shift42(x[18], x[50])
	x[19], x[51] = invbutterfly_shift39(x[19], x[51])
	x[20], x[52] = invbutterfly_shift36(x[20], x[52])
	x[21], x[53] = invbutterfly_shift33(x[21], x[53])
	x[22], x[54] = invbutterfly_shift30(x[22], x[54])
	x[23], x[55] = invbutterfly_shift27(x[23], x[55])
	x[24], x[56] = invbutterfly_shift24(x[24], x[56])
	x[25], x[57] = invbutterfly_shift21(x[25], x[57])
	x[26], x[58] = invbutterfly_shift18(x[26], x[58])
	x[27], x[59] = invbutterfly_shift15(x[27], x[59])
	x[28], x[60] = invbutterfly_shift12(x[28], x[60])
	x[29], x[61] = invbutterfly_shift9(x[29], x[61])
	x[30], x[62] = invbutterfly_shift6(x[30], x[62])
	x[31], x[63] = invbutterfly_shift3(x[31], x[63])

}

// Fft for size 2**7
//
// This is an in place FFT with a bit reversed output
func fft7(x []uint64) {
	x[0], x[64] = butterfly_null(x[0], x[64])
	x[1], x[65] = butterfly_mul(x[1], x[65], 2198989700608)
	x[2], x[66] = butterfly_shift3(x[2], x[66])
	x[3], x[67] = butterfly_mul(x[3], x[67], 17591917604864)
	x[4], x[68] = butterfly_shift6(x[4], x[68])
	x[5], x[69] = butterfly_mul(x[5], x[69], 140735340838912)
	x[6], x[70] = butterfly_shift9(x[6], x[70])
	x[7], x[71] = butterfly_mul(x[7], x[71], 1125882726711296)
	x[8], x[72] = butterfly_shift12(x[8], x[72])
	x[9], x[73] = butterfly_mul(x[9], x[73], 9007061813690368)
	x[10], x[74] = butterfly_shift15(x[10], x[74])
	x[11], x[75] = butterfly_mul(x[11], x[75], 72056494509522944)
	x[12], x[76] = butterfly_shift18(x[12], x[76])
	x[13], x[77] = butterfly_mul(x[13], x[77], 576451956076183552)
	x[14], x[78] = butterfly_shift21(x[14], x[78])
	x[15], x[79] = butterfly_mul(x[15], x[79], 4611615648609468416)
	x[16], x[80] = butterfly_shift24(x[16], x[80])
	x[17], x[81] = butterfly_mul(x[17], x[81], 18446181119461163007)
	x[18], x[82] = butterfly_shift27(x[18], x[82])
	x[19], x[83] = butterfly_mul(x[19], x[83], 18442240469787213809)
	x[20], x[84] = butterfly_shift30(x[20], x[84])
	x[21], x[85] = butterfly_mul(x[21], x[85], 18410715272395620225)
	x[22], x[86] = butterfly_shift33(x[22], x[86])
	x[23], x[87] = butterfly_mul(x[23], x[87], 18158513693262871553)
	x[24], x[88] = butterfly_shift36(x[24], x[88])
	x[25], x[89] = butterfly_mul(x[25], x[89], 16140901060200882177)
	x[26], x[90] = butterfly_shift39(x[26], x[90])
	x[27], x[91] = butterfly_mul(x[27], x[91], 18446744065119551490)
	x[28], x[92] = butterfly_shift42(x[28], x[92])
	x[29], x[93] = butterfly_mul(x[29], x[93], 18446744035054321673)
	x[30], x[94] = butterfly_shift45(x[30], x[94])
	x[31], x[95] = butterfly_mul(x[31], x[95], 18446743794532483137)
	x[32], x[96] = butterfly_shift48(x[32], x[96])
	x[33], x[97] = butterfly_mul(x[33], x[97], 18446741870357774849)
	x[34], x[98] = butterfly_shift51(x[34], x[98])
	x[35], x[99] = butterfly_mul(x[35], x[99], 18446726476960108545)
	x[36], x[100] = butterfly_shift54(x[36], x[100])
	x[37], x[101] = butterfly_mul(x[37], x[101], 18446603329778778113)
	x[38], x[102] = butterfly_shift57(x[38], x[102])
	x[39], x[103] = butterfly_mul(x[39], x[103], 18445618152328134657)
	x[40], x[104] = butterfly_shift60(x[40], x[104])
	x[41], x[105] = butterfly_mul(x[41], x[105], 18437736732722987009)
	x[42], x[106] = butterfly_shift63(x[42], x[106])
	x[43], x[107] = butterfly_mul(x[43], x[107], 18374685375881805825)
	x[44], x[108] = butterfly_shift66(x[44], x[108])
	x[45], x[109] = butterfly_mul(x[45], x[109], 17870274521152356353)
	x[46], x[110] = butterfly_shift69(x[46], x[110])
	x[47], x[111] = butterfly_mul(x[47], x[111], 13834987683316760577)
	x[48], x[112] = butterfly_shift72(x[48], x[112])
	x[49], x[113] = butterfly_mul(x[49], x[113], 18446181119461163011)
	x[50], x[114] = butterfly_shift75(x[50], x[114])
	x[51], x[115] = butterfly_mul(x[51], x[115], 18442240469787213841)
	x[52], x[116] = butterfly_shift78(x[52], x[116])
	x[53], x[117] = butterfly_mul(x[53], x[117], 18410715272395620481)
	x[54], x[118] = butterfly_shift81(x[54], x[118])
	x[55], x[119] = butterfly_mul(x[55], x[119], 18158513693262873601)
	x[56], x[120] = butterfly_shift84(x[56], x[120])
	x[57], x[121] = butterfly_mul(x[57], x[121], 16140901060200898561)
	x[58], x[122] = butterfly_shift87(x[58], x[122])
	x[59], x[123] = butterfly_mul(x[59], x[123], 18446744065119682562)
	x[60], x[124] = butterfly_shift90(x[60], x[124])
	x[61], x[125] = butterfly_mul(x[61], x[125], 18446744035055370249)
	x[62], x[126] = butterfly_shift93(x[62], x[126])
	x[63], x[127] = butterfly_mul(x[63], x[127], 18446743794540871745)
	x[0], x[32] = butterfly_null(x[0], x[32])
	x[64], x[96] = butterfly_null(x[64], x[96])
	x[1], x[33] = butterfly_shift3(x[1], x[33])
	x[65], x[97] = butterfly_shift3(x[65], x[97])
	x[2], x[34] = butterfly_shift6(x[2], x[34])
	x[66], x[98] = butterfly_shift6(x[66], x[98])
	x[3], x[35] = butterfly_shift9(x[3], x[35])
	x[67], x[99] = butterfly_shift9(x[67], x[99])
	x[4], x[36] = butterfly_shift12(x[4], x[36])
	x[68], x[100] = butterfly_shift12(x[68], x[100])
	x[5], x[37] = butterfly_shift15(x[5], x[37])
	x[69], x[101] = butterfly_shift15(x[69], x[101])
	x[6], x[38] = butterfly_shift18(x[6], x[38])
	x[70], x[102] = butterfly_shift18(x[70], x[102])
	x[7], x[39] = butterfly_shift21(x[7], x[39])
	x[71], x[103] = butterfly_shift21(x[71], x[103])
	x[8], x[40] = butterfly_shift24(x[8], x[40])
	x[72], x[104] = butterfly_shift24(x[72], x[104])
	x[9], x[41] = butterfly_shift27(x[9], x[41])
	x[73], x[105] = butterfly_shift27(x[73], x[105])
	x[10], x[42] = butterfly_shift30(x[10], x[42])
	x[74], x[106] = butterfly_shift30(x[74], x[106])
	x[11], x[43] = butterfly_shift33(x[11], x[43])
	x[75], x[107] = butterfly_shift33(x[75], x[107])
	x[12], x[44] = butterfly_shift36(x[12], x[44])
	x[76], x[108] = butterfly_shift36(x[76], x[108])
	x[13], x[45] = butterfly_shift39(x[13], x[45])
	x[77], x[109] = butterfly_shift39(x[77], x[109])
	x[14], x[46] = butterfly_shift42(x[14], x[46])
	x[78], x[110] = butterfly_shift42(x[78], x[110])
	x[15], x[47] = butterfly_shift45(x[15], x[47])
	x[79], x[111] = butterfly_shift45(x[79], x[111])
	x[16], x[48] = butterfly_shift48(x[16], x[48])
	x[80], x[112] = butterfly_shift48(x[80], x[112])
	x[17], x[49] = butterfly_shift51(x[17], x[49])
	x[81], x[113] = butterfly_shift51(x[81], x[113])
	x[18], x[50] = butterfly_shift54(x[18], x[50])
	x[82], x[114] = butterfly_shift54(x[82], x[114])
	x[19], x[51] = butterfly_shift57(x[19], x[51])
	x[83], x[115] = butterfly_shift57(x[83], x[115])
	x[20], x[52] = butterfly_shift60(x[20], x[52])
	x[84], x[116] = butterfly_shift60(x[84], x[116])
	x[21], x[53] = butterfly_shift63(x[21], x[53])
	x[85], x[117] = butterfly_shift63(x[85], x[117])
	x[22], x[54] = butterfly_shift66(x[22], x[54])
	x[86], x[118] = butterfly_shift66(x[86], x[118])
	x[23], x[55] = butterfly_shift69(x[23], x[55])
	x[87], x[119] = butterfly_shift69(x[87], x[119])
	x[24], x[56] = butterfly_shift72(x[24], x[56])
	x[88], x[120] = butterfly_shift72(x[88], x[120])
	x[25], x[57] = butterfly_shift75(x[25], x[57])
	x[89], x[121] = butterfly_shift75(x[89], x[121])
	x[26], x[58] = butterfly_shift78(x[26], x[58])
	x[90], x[122] = butterfly_shift78(x[90], x[122])
	x[27], x[59] = butterfly_shift81(x[27], x[59])
	x[91], x[123] = butterfly_shift81(x[91], x[123])
	x[28], x[60] = butterfly_shift84(x[28], x[60])
	x[92], x[124] = butterfly_shift84(x[92], x[124])
	x[29], x[61] = butterfly_shift87(x[29], x[61])
	x[93], x[125] = butterfly_shift87(x[93], x[125])
	x[30], x[62] = butterfly_shift90(x[30], x[62])
	x[94], x[126] = butterfly_shift90(x[94], x[126])
	x[31], x[63] = butterfly_shift93(x[31], x[63])
	x[95], x[127] = butterfly_shift93(x[95], x[127])
	x[0], x[16] = butterfly_null(x[0], x[16])
	x[32], x[48] = butterfly_null(x[32], x[48])
	x[64], x[80] = butterfly_null(x[64], x[80])
	x[96], x[112] = butterfly_null(x[96], x[112])
	x[1], x[17] = butterfly_shift6(x[1], x[17])
	x[33], x[49] = butterfly_shift6(x[33], x[49])
	x[65], x[81] = butterfly_shift6(x[65], x[81])
	x[97], x[113] = butterfly_shift6(x[97], x[113])
	x[2], x[18] = butterfly_shift12(x[2], x[18])
	x[34], x[50] = butterfly_shift12(x[34], x[50])
	x[66], x[82] = butterfly_shift12(x[66], x[82])
	x[98], x[114] = butterfly_shift12(x[98], x[114])
	x[3], x[19] = butterfly_shift18(x[3], x[19])
	x[35], x[51] = butterfly_shift18(x[35], x[51])
	x[67], x[83] = butterfly_shift18(x[67], x[83])
	x[99], x[115] = butterfly_shift18(x[99], x[115])
	x[4], x[20] = butterfly_shift24(x[4], x[20])
	x[36], x[52] = butterfly_shift24(x[36], x[52])
	x[68], x[84] = butterfly_shift24(x[68], x[84])
	x[100], x[116] = butterfly_shift24(x[100], x[116])
	x[5], x[21] = butterfly_shift30(x[5], x[21])
	x[37], x[53] = butterfly_shift30(x[37], x[53])
	x[69], x[85] = butterfly_shift30(x[69], x[85])
	x[101], x[117] = butterfly_shift30(x[101], x[117])
	x[6], x[22] = butterfly_shift36(x[6], x[22])
	x[38], x[54] = butterfly_shift36(x[38], x[54])
	x[70], x[86] = butterfly_shift36(x[70], x[86])
	x[102], x[118] = butterfly_shift36(x[102], x[118])
	x[7], x[23] = butterfly_shift42(x[7], x[23])
	x[39], x[55] = butterfly_shift42(x[39], x[55])
	x[71], x[87] = butterfly_shift42(x[71], x[87])
	x[103], x[119] = butterfly_shift42(x[103], x[119])
	x[8], x[24] = butterfly_shift48(x[8], x[24])
	x[40], x[56] = butterfly_shift48(x[40], x[56])
	x[72], x[88] = butterfly_shift48(x[72], x[88])
	x[104], x[120] = butterfly_shift48(x[104], x[120])
	x[9], x[25] = butterfly_shift54(x[9], x[25])
	x[41], x[57] = butterfly_shift54(x[41], x[57])
	x[73], x[89] = butterfly_shift54(x[73], x[89])
	x[105], x[121] = butterfly_shift54(x[105], x[121])
	x[10], x[26] = butterfly_shift60(x[10], x[26])
	x[42], x[58] = butterfly_shift60(x[42], x[58])
	x[74], x[90] = butterfly_shift60(x[74], x[90])
	x[106], x[122] = butterfly_shift60(x[106], x[122])
	x[11], x[27] = butterfly_shift66(x[11], x[27])
	x[43], x[59] = butterfly_shift66(x[43], x[59])
	x[75], x[91] = butterfly_shift66(x[75], x[91])
	x[107], x[123] = butterfly_shift66(x[107], x[123])
	x[12], x[28] = butterfly_shift72(x[12], x[28])
	x[44], x[60] = butterfly_shift72(x[44], x[60])
	x[76], x[92] = butterfly_shift72(x[76], x[92])
	x[108], x[124] = butterfly_shift72(x[108], x[124])
	x[13], x[29] = butterfly_shift78(x[13], x[29])
	x[45], x[61] = butterfly_shift78(x[45], x[61])
	x[77], x[93] = butterfly_shift78(x[77], x[93])
	x[109], x[125] = butterfly_shift78(x[109], x[125])
	x[14], x[30] = butterfly_shift84(x[14], x[30])
	x[46], x[62] = butterfly_shift84(x[46], x[62])
	x[78], x[94] = butterfly_shift84(x[78], x[94])
	x[110], x[126] = butterfly_shift84(x[110], x[126])
	x[15], x[31] = butterfly_shift90(x[15], x[31])
	x[47], x[63] = butterfly_shift90(x[47], x[63])
	x[79], x[95] = butterfly_shift90(x[79], x[95])
	x[111], x[127] = butterfly_shift90(x[111], x[127])
	x[0], x[8] = butterfly_null(x[0], x[8])
	x[16], x[24] = butterfly_null(x[16], x[24])
	x[32], x[40] = butterfly_null(x[32], x[40])
	x[48], x[56] = butterfly_null(x[48], x[56])
	x[64], x[72] = butterfly_null(x[64], x[72])
	x[80], x[88] = butterfly_null(x[80], x[88])
	x[96], x[104] = butterfly_null(x[96], x[104])
	x[112], x[120] = butterfly_null(x[112], x[120])
	x[1], x[9] = butterfly_shift12(x[1], x[9])
	x[17], x[25] = butterfly_shift12(x[17], x[25])
	x[33], x[41] = butterfly_shift12(x[33], x[41])
	x[49], x[57] = butterfly_shift12(x[49], x[57])
	x[65], x[73] = butterfly_shift12(x[65], x[73])
	x[81], x[89] = butterfly_shift12(x[81], x[89])
	x[97], x[105] = butterfly_shift12(x[97], x[105])
	x[113], x[121] = butterfly_shift12(x[113], x[121])
	x[2], x[10] = butterfly_shift24(x[2], x[10])
	x[18], x[26] = butterfly_shift24(x[18], x[26])
	x[34], x[42] = butterfly_shift24(x[34], x[42])
	x[50], x[58] = butterfly_shift24(x[50], x[58])
	x[66], x[74] = butterfly_shift24(x[66], x[74])
	x[82], x[90] = butterfly_shift24(x[82], x[90])
	x[98], x[106] = butterfly_shift24(x[98], x[106])
	x[114], x[122] = butterfly_shift24(x[114], x[122])
	x[3], x[11] = butterfly_shift36(x[3], x[11])
	x[19], x[27] = butterfly_shift36(x[19], x[27])
	x[35], x[43] = butterfly_shift36(x[35], x[43])
	x[51], x[59] = butterfly_shift36(x[51], x[59])
	x[67], x[75] = butterfly_shift36(x[67], x[75])
	x[83], x[91] = butterfly_shift36(x[83], x[91])
	x[99], x[107] = butterfly_shift36(x[99], x[107])
	x[115], x[123] = butterfly_shift36(x[115], x[123])
	x[4], x[12] = butterfly_shift48(x[4], x[12])
	x[20], x[28] = butterfly_shift48(x[20], x[28])
	x[36], x[44] = butterfly_shift48(x[36], x[44])
	x[52], x[60] = butterfly_shift48(x[52], x[60])
	x[68], x[76] = butterfly_shift48(x[68], x[76])
	x[84], x[92] = butterfly_shift48(x[84], x[92])
	x[100], x[108] = butterfly_shift48(x[100], x[108])
	x[116], x[124] = butterfly_shift48(x[116], x[124])
	x[5], x[13] = butterfly_shift60(x[5], x[13])
	x[21], x[29] = butterfly_shift60(x[21], x[29])
	x[37], x[45] = butterfly_shift60(x[37], x[45])
	x[53], x[61] = butterfly_shift60(x[53], x[61])
	x[69], x[77] = butterfly_shift60(x[69], x[77])
	x[85], x[93] = butterfly_shift60(x[85], x[93])
	x[101], x[109] = butterfly_shift60(x[101], x[109])
	x[117], x[125] = butterfly_shift60(x[117], x[125])
	x[6], x[14] = butterfly_shift72(x[6], x[14])
	x[22], x[30] = butterfly_shift72(x[22], x[30])
	x[38], x[46] = butterfly_shift72(x[38], x[46])
	x[54], x[62] = butterfly_shift72(x[54], x[62])
	x[70], x[78] = butterfly_shift72(x[70], x[78])
	x[86], x[94] = butterfly_shift72(x[86], x[94])
	x[102], x[110] = butterfly_shift72(x[102], x[110])
	x[118], x[126] = butterfly_shift72(x[118], x[126])
	x[7], x[15] = butterfly_shift84(x[7], x[15])
	x[23], x[31] = butterfly_shift84(x[23], x[31])
	x[39], x[47] = butterfly_shift84(x[39], x[47])
	x[55], x[63] = butterfly_shift84(x[55], x[63])
	x[71], x[79] = butterfly_shift84(x[71], x[79])
	x[87], x[95] = butterfly_shift84(x[87], x[95])
	x[103], x[111] = butterfly_shift84(x[103], x[111])
	x[119], x[127] = butterfly_shift84(x[119], x[127])
	x[0], x[4] = butterfly_null(x[0], x[4])
	x[8], x[12] = butterfly_null(x[8], x[12])
	x[16], x[20] = butterfly_null(x[16], x[20])
	x[24], x[28] = butterfly_null(x[24], x[28])
	x[32], x[36] = butterfly_null(x[32], x[36])
	x[40], x[44] = butterfly_null(x[40], x[44])
	x[48], x[52] = butterfly_null(x[48], x[52])
	x[56], x[60] = butterfly_null(x[56], x[60])
	x[64], x[68] = butterfly_null(x[64], x[68])
	x[72], x[76] = butterfly_null(x[72], x[76])
	x[80], x[84] = butterfly_null(x[80], x[84])
	x[88], x[92] = butterfly_null(x[88], x[92])
	x[96], x[100] = butterfly_null(x[96], x[100])
	x[104], x[108] = butterfly_null(x[104], x[108])
	x[112], x[116] = butterfly_null(x[112], x[116])
	x[120], x[124] = butterfly_null(x[120], x[124])
	x[1], x[5] = butterfly_shift24(x[1], x[5])
	x[9], x[13] = butterfly_shift24(x[9], x[13])
	x[17], x[21] = butterfly_shift24(x[17], x[21])
	x[25], x[29] = butterfly_shift24(x[25], x[29])
	x[33], x[37] = butterfly_shift24(x[33], x[37])
	x[41], x[45] = butterfly_shift24(x[41], x[45])
	x[49], x[53] = butterfly_shift24(x[49], x[53])
	x[57], x[61] = butterfly_shift24(x[57], x[61])
	x[65], x[69] = butterfly_shift24(x[65], x[69])
	x[73], x[77] = butterfly_shift24(x[73], x[77])
	x[81], x[85] = butterfly_shift24(x[81], x[85])
	x[89], x[93] = butterfly_shift24(x[89], x[93])
	x[97], x[101] = butterfly_shift24(x[97], x[101])
	x[105], x[109] = butterfly_shift24(x[105], x[109])
	x[113], x[117] = butterfly_shift24(x[113], x[117])
	x[121], x[125] = butterfly_shift24(x[121], x[125])
	x[2], x[6] = butterfly_shift48(x[2], x[6])
	x[10], x[14] = butterfly_shift48(x[10], x[14])
	x[18], x[22] = butterfly_shift48(x[18], x[22])
	x[26], x[30] = butterfly_shift48(x[26], x[30])
	x[34], x[38] = butterfly_shift48(x[34], x[38])
	x[42], x[46] = butterfly_shift48(x[42], x[46])
	x[50], x[54] = butterfly_shift48(x[50], x[54])
	x[58], x[62] = butterfly_shift48(x[58], x[62])
	x[66], x[70] = butterfly_shift48(x[66], x[70])
	x[74], x[78] = butterfly_shift48(x[74], x[78])
	x[82], x[86] = butterfly_shift48(x[82], x[86])
	x[90], x[94] = butterfly_shift48(x[90], x[94])
	x[98], x[102] = butterfly_shift48(x[98], x[102])
	x[106], x[110] = butterfly_shift48(x[106], x[110])
	x[114], x[118] = butterfly_shift48(x[114], x[118])
	x[122], x[126] = butterfly_shift48(x[122], x[126])
	x[3], x[7] = butterfly_shift72(x[3], x[7])
	x[11], x[15] = butterfly_shift72(x[11], x[15])
	x[19], x[23] = butterfly_shift72(x[19], x[23])
	x[27], x[31] = butterfly_shift72(x[27], x[31])
	x[35], x[39] = butterfly_shift72(x[35], x[39])
	x[43], x[47] = butterfly_shift72(x[43], x[47])
	x[51], x[55] = butterfly_shift72(x[51], x[55])
	x[59], x[63] = butterfly_shift72(x[59], x[63])
	x[67], x[71] = butterfly_shift72(x[67], x[71])
	x[75], x[79] = butterfly_shift72(x[75], x[79])
	x[83], x[87] = butterfly_shift72(x[83], x[87])
	x[91], x[95] = butterfly_shift72(x[91], x[95])
	x[99], x[103] = butterfly_shift72(x[99], x[103])
	x[107], x[111] = butterfly_shift72(x[107], x[111])
	x[115], x[119] = butterfly_shift72(x[115], x[119])
	x[123], x[127] = butterfly_shift72(x[123], x[127])
	x[0], x[2] = butterfly_null(x[0], x[2])
	x[4], x[6] = butterfly_null(x[4], x[6])
	x[8], x[10] = butterfly_null(x[8], x[10])
	x[12], x[14] = butterfly_null(x[12], x[14])
	x[16], x[18] = butterfly_null(x[16], x[18])
	x[20], x[22] = butterfly_null(x[20], x[22])
	x[24], x[26] = butterfly_null(x[24], x[26])
	x[28], x[30] = butterfly_null(x[28], x[30])
	x[32], x[34] = butterfly_null(x[32], x[34])
	x[36], x[38] = butterfly_null(x[36], x[38])
	x[40], x[42] = butterfly_null(x[40], x[42])
	x[44], x[46] = butterfly_null(x[44], x[46])
	x[48], x[50] = butterfly_null(x[48], x[50])
	x[52], x[54] = butterfly_null(x[52], x[54])
	x[56], x[58] = butterfly_null(x[56], x[58])
	x[60], x[62] = butterfly_null(x[60], x[62])
	x[64], x[66] = butterfly_null(x[64], x[66])
	x[68], x[70] = butterfly_null(x[68], x[70])
	x[72], x[74] = butterfly_null(x[72], x[74])
	x[76], x[78] = butterfly_null(x[76], x[78])
	x[80], x[82] = butterfly_null(x[80], x[82])
	x[84], x[86] = butterfly_null(x[84], x[86])
	x[88], x[90] = butterfly_null(x[88], x[90])
	x[92], x[94] = butterfly_null(x[92], x[94])
	x[96], x[98] = butterfly_null(x[96], x[98])
	x[100], x[102] = butterfly_null(x[100], x[102])
	x[104], x[106] = butterfly_null(x[104], x[106])
	x[108], x[110] = butterfly_null(x[108], x[110])
	x[112], x[114] = butterfly_null(x[112], x[114])
	x[116], x[118] = butterfly_null(x[116], x[118])
	x[120], x[122] = butterfly_null(x[120], x[122])
	x[124], x[126] = butterfly_null(x[124], x[126])
	x[1], x[3] = butterfly_shift48(x[1], x[3])
	x[5], x[7] = butterfly_shift48(x[5], x[7])
	x[9], x[11] = butterfly_shift48(x[9], x[11])
	x[13], x[15] = butterfly_shift48(x[13], x[15])
	x[17], x[19] = butterfly_shift48(x[17], x[19])
	x[21], x[23] = butterfly_shift48(x[21], x[23])
	x[25], x[27] = butterfly_shift48(x[25], x[27])
	x[29], x[31] = butterfly_shift48(x[29], x[31])
	x[33], x[35] = butterfly_shift48(x[33], x[35])
	x[37], x[39] = butterfly_shift48(x[37], x[39])
	x[41], x[43] = butterfly_shift48(x[41], x[43])
	x[45], x[47] = butterfly_shift48(x[45], x[47])
	x[49], x[51] = butterfly_shift48(x[49], x[51])
	x[53], x[55] = butterfly_shift48(x[53], x[55])
	x[57], x[59] = butterfly_shift48(x[57], x[59])
	x[61], x[63] = butterfly_shift48(x[61], x[63])
	x[65], x[67] = butterfly_shift48(x[65], x[67])
	x[69], x[71] = butterfly_shift48(x[69], x[71])
	x[73], x[75] = butterfly_shift48(x[73], x[75])
	x[77], x[79] = butterfly_shift48(x[77], x[79])
	x[81], x[83] = butterfly_shift48(x[81], x[83])
	x[85], x[87] = butterfly_shift48(x[85], x[87])
	x[89], x[91] = butterfly_shift48(x[89], x[91])
	x[93], x[95] = butterfly_shift48(x[93], x[95])
	x[97], x[99] = butterfly_shift48(x[97], x[99])
	x[101], x[103] = butterfly_shift48(x[101], x[103])
	x[105], x[107] = butterfly_shift48(x[105], x[107])
	x[109], x[111] = butterfly_shift48(x[109], x[111])
	x[113], x[115] = butterfly_shift48(x[113], x[115])
	x[117], x[119] = butterfly_shift48(x[117], x[119])
	x[121], x[123] = butterfly_shift48(x[121], x[123])
	x[125], x[127] = butterfly_shift48(x[125], x[127])
	x[0], x[1] = butterfly_null(x[0], x[1])
	x[2], x[3] = butterfly_null(x[2], x[3])
	x[4], x[5] = butterfly_null(x[4], x[5])
	x[6], x[7] = butterfly_null(x[6], x[7])
	x[8], x[9] = butterfly_null(x[8], x[9])
	x[10], x[11] = butterfly_null(x[10], x[11])
	x[12], x[13] = butterfly_null(x[12], x[13])
	x[14], x[15] = butterfly_null(x[14], x[15])
	x[16], x[17] = butterfly_null(x[16], x[17])
	x[18], x[19] = butterfly_null(x[18], x[19])
	x[20], x[21] = butterfly_null(x[20], x[21])
	x[22], x[23] = butterfly_null(x[22], x[23])
	x[24], x[25] = butterfly_null(x[24], x[25])
	x[26], x[27] = butterfly_null(x[26], x[27])
	x[28], x[29] = butterfly_null(x[28], x[29])
	x[30], x[31] = butterfly_null(x[30], x[31])
	x[32], x[33] = butterfly_null(x[32], x[33])
	x[34], x[35] = butterfly_null(x[34], x[35])
	x[36], x[37] = butterfly_null(x[36], x[37])
	x[38], x[39] = butterfly_null(x[38], x[39])
	x[40], x[41] = butterfly_null(x[40], x[41])
	x[42], x[43] = butterfly_null(x[42], x[43])
	x[44], x[45] = butterfly_null(x[44], x[45])
	x[46], x[47] = butterfly_null(x[46], x[47])
	x[48], x[49] = butterfly_null(x[48], x[49])
	x[50], x[51] = butterfly_null(x[50], x[51])
	x[52], x[53] = butterfly_null(x[52], x[53])
	x[54], x[55] = butterfly_null(x[54], x[55])
	x[56], x[57] = butterfly_null(x[56], x[57])
	x[58], x[59] = butterfly_null(x[58], x[59])
	x[60], x[61] = butterfly_null(x[60], x[61])
	x[62], x[63] = butterfly_null(x[62], x[63])
	x[64], x[65] = butterfly_null(x[64], x[65])
	x[66], x[67] = butterfly_null(x[66], x[67])
	x[68], x[69] = butterfly_null(x[68], x[69])
	x[70], x[71] = butterfly_null(x[70], x[71])
	x[72], x[73] = butterfly_null(x[72], x[73])
	x[74], x[75] = butterfly_null(x[74], x[75])
	x[76], x[77] = butterfly_null(x[76], x[77])
	x[78], x[79] = butterfly_null(x[78], x[79])
	x[80], x[81] = butterfly_null(x[80], x[81])
	x[82], x[83] = butterfly_null(x[82], x[83])
	x[84], x[85] = butterfly_null(x[84], x[85])
	x[86], x[87] = butterfly_null(x[86], x[87])
	x[88], x[89] = butterfly_null(x[88], x[89])
	x[90], x[91] = butterfly_null(x[90], x[91])
	x[92], x[93] = butterfly_null(x[92], x[93])
	x[94], x[95] = butterfly_null(x[94], x[95])
	x[96], x[97] = butterfly_null(x[96], x[97])
	x[98], x[99] = butterfly_null(x[98], x[99])
	x[100], x[101] = butterfly_null(x[100], x[101])
	x[102], x[103] = butterfly_null(x[102], x[103])
	x[104], x[105] = butterfly_null(x[104], x[105])
	x[106], x[107] = butterfly_null(x[106], x[107])
	x[108], x[109] = butterfly_null(x[108], x[109])
	x[110], x[111] = butterfly_null(x[110], x[111])
	x[112], x[113] = butterfly_null(x[112], x[113])
	x[114], x[115] = butterfly_null(x[114], x[115])
	x[116], x[117] = butterfly_null(x[116], x[117])
	x[118], x[119] = butterfly_null(x[118], x[119])
	x[120], x[121] = butterfly_null(x[120], x[121])
	x[122], x[123] = butterfly_null(x[122], x[123])
	x[124], x[125] = butterfly_null(x[124], x[125])
	x[126], x[127] = butterfly_null(x[126], x[127])

}

// InvFft for size 2**7
//
// This is an in place Inverse FFT with a bit reversed input
func invfft7(x []uint64) {
	x[0], x[1] = invbutterfly_null(x[0], x[1])
	x[2], x[3] = invbutterfly_null(x[2], x[3])
	x[4], x[5] = invbutterfly_null(x[4], x[5])
	x[6], x[7] = invbutterfly_null(x[6], x[7])
	x[8], x[9] = invbutterfly_null(x[8], x[9])
	x[10], x[11] = invbutterfly_null(x[10], x[11])
	x[12], x[13] = invbutterfly_null(x[12], x[13])
	x[14], x[15] = invbutterfly_null(x[14], x[15])
	x[16], x[17] = invbutterfly_null(x[16], x[17])
	x[18], x[19] = invbutterfly_null(x[18], x[19])
	x[20], x[21] = invbutterfly_null(x[20], x[21])
	x[22], x[23] = invbutterfly_null(x[22], x[23])
	x[24], x[25] = invbutterfly_null(x[24], x[25])
	x[26], x[27] = invbutterfly_null(x[26], x[27])
	x[28], x[29] = invbutterfly_null(x[28], x[29])
	x[30], x[31] = invbutterfly_null(x[30], x[31])
	x[32], x[33] = invbutterfly_null(x[32], x[33])
	x[34], x[35] = invbutterfly_null(x[34], x[35])
	x[36], x[37] = invbutterfly_null(x[36], x[37])
	x[38], x[39] = invbutterfly_null(x[38], x[39])
	x[40], x[41] = invbutterfly_null(x[40], x[41])
	x[42], x[43] = invbutterfly_null(x[42], x[43])
	x[44], x[45] = invbutterfly_null(x[44], x[45])
	x[46], x[47] = invbutterfly_null(x[46], x[47])
	x[48], x[49] = invbutterfly_null(x[48], x[49])
	x[50], x[51] = invbutterfly_null(x[50], x[51])
	x[52], x[53] = invbutterfly_null(x[52], x[53])
	x[54], x[55] = invbutterfly_null(x[54], x[55])
	x[56], x[57] = invbutterfly_null(x[56], x[57])
	x[58], x[59] = invbutterfly_null(x[58], x[59])
	x[60], x[61] = invbutterfly_null(x[60], x[61])
	x[62], x[63] = invbutterfly_null(x[62], x[63])
	x[64], x[65] = invbutterfly_null(x[64], x[65])
	x[66], x[67] = invbutterfly_null(x[66], x[67])
	x[68], x[69] = invbutterfly_null(x[68], x[69])
	x[70], x[71] = invbutterfly_null(x[70], x[71])
	x[72], x[73] = invbutterfly_null(x[72], x[73])
	x[74], x[75] = invbutterfly_null(x[74], x[75])
	x[76], x[77] = invbutterfly_null(x[76], x[77])
	x[78], x[79] = invbutterfly_null(x[78], x[79])
	x[80], x[81] = invbutterfly_null(x[80], x[81])
	x[82], x[83] = invbutterfly_null(x[82], x[83])
	x[84], x[85] = invbutterfly_null(x[84], x[85])
	x[86], x[87] = invbutterfly_null(x[86], x[87])
	x[88], x[89] = invbutterfly_null(x[88], x[89])
	x[90], x[91] = invbutterfly_null(x[90], x[91])
	x[92], x[93] = invbutterfly_null(x[92], x[93])
	x[94], x[95] = invbutterfly_null(x[94], x[95])
	x[96], x[97] = invbutterfly_null(x[96], x[97])
	x[98], x[99] = invbutterfly_null(x[98], x[99])
	x[100], x[101] = invbutterfly_null(x[100], x[101])
	x[102], x[103] = invbutterfly_null(x[102], x[103])
	x[104], x[105] = invbutterfly_null(x[104], x[105])
	x[106], x[107] = invbutterfly_null(x[106], x[107])
	x[108], x[109] = invbutterfly_null(x[108], x[109])
	x[110], x[111] = invbutterfly_null(x[110], x[111])
	x[112], x[113] = invbutterfly_null(x[112], x[113])
	x[114], x[115] = invbutterfly_null(x[114], x[115])
	x[116], x[117] = invbutterfly_null(x[116], x[117])
	x[118], x[119] = invbutterfly_null(x[118], x[119])
	x[120], x[121] = invbutterfly_null(x[120], x[121])
	x[122], x[123] = invbutterfly_null(x[122], x[123])
	x[124], x[125] = invbutterfly_null(x[124], x[125])
	x[126], x[127] = invbutterfly_null(x[126], x[127])
	x[0], x[2] = invbutterfly_null(x[0], x[2])
	x[4], x[6] = invbutterfly_null(x[4], x[6])
	x[8], x[10] = invbutterfly_null(x[8], x[10])
	x[12], x[14] = invbutterfly_null(x[12], x[14])
	x[16], x[18] = invbutterfly_null(x[16], x[18])
	x[20], x[22] = invbutterfly_null(x[20], x[22])
	x[24], x[26] = invbutterfly_null(x[24], x[26])
	x[28], x[30] = invbutterfly_null(x[28], x[30])
	x[32], x[34] = invbutterfly_null(x[32], x[34])
	x[36], x[38] = invbutterfly_null(x[36], x[38])
	x[40], x[42] = invbutterfly_null(x[40], x[42])
	x[44], x[46] = invbutterfly_null(x[44], x[46])
	x[48], x[50] = invbutterfly_null(x[48], x[50])
	x[52], x[54] = invbutterfly_null(x[52], x[54])
	x[56], x[58] = invbutterfly_null(x[56], x[58])
	x[60], x[62] = invbutterfly_null(x[60], x[62])
	x[64], x[66] = invbutterfly_null(x[64], x[66])
	x[68], x[70] = invbutterfly_null(x[68], x[70])
	x[72], x[74] = invbutterfly_null(x[72], x[74])
	x[76], x[78] = invbutterfly_null(x[76], x[78])
	x[80], x[82] = invbutterfly_null(x[80], x[82])
	x[84], x[86] = invbutterfly_null(x[84], x[86])
	x[88], x[90] = invbutterfly_null(x[88], x[90])
	x[92], x[94] = invbutterfly_null(x[92], x[94])
	x[96], x[98] = invbutterfly_null(x[96], x[98])
	x[100], x[102] = invbutterfly_null(x[100], x[102])
	x[104], x[106] = invbutterfly_null(x[104], x[106])
	x[108], x[110] = invbutterfly_null(x[108], x[110])
	x[112], x[114] = invbutterfly_null(x[112], x[114])
	x[116], x[118] = invbutterfly_null(x[116], x[118])
	x[120], x[122] = invbutterfly_null(x[120], x[122])
	x[124], x[126] = invbutterfly_null(x[124], x[126])
	x[1], x[3] = invbutterfly_shift48(x[1], x[3])
	x[5], x[7] = invbutterfly_shift48(x[5], x[7])
	x[9], x[11] = invbutterfly_shift48(x[9], x[11])
	x[13], x[15] = invbutterfly_shift48(x[13], x[15])
	x[17], x[19] = invbutterfly_shift48(x[17], x[19])
	x[21], x[23] = invbutterfly_shift48(x[21], x[23])
	x[25], x[27] = invbutterfly_shift48(x[25], x[27])
	x[29], x[31] = invbutterfly_shift48(x[29], x[31])
	x[33], x[35] = invbutterfly_shift48(x[33], x[35])
	x[37], x[39] = invbutterfly_shift48(x[37], x[39])
	x[41], x[43] = invbutterfly_shift48(x[41], x[43])
	x[45], x[47] = invbutterfly_shift48(x[45], x[47])
	x[49], x[51] = invbutterfly_shift48(x[49], x[51])
	x[53], x[55] = invbutterfly_shift48(x[53], x[55])
	x[57], x[59] = invbutterfly_shift48(x[57], x[59])
	x[61], x[63] = invbutterfly_shift48(x[61], x[63])
	x[65], x[67] = invbutterfly_shift48(x[65], x[67])
	x[69], x[71] = invbutterfly_shift48(x[69], x[71])
	x[73], x[75] = invbutterfly_shift48(x[73], x[75])
	x[77], x[79] = invbutterfly_shift48(x[77], x[79])
	x[81], x[83] = invbutterfly_shift48(x[81], x[83])
	x[85], x[87] = invbutterfly_shift48(x[85], x[87])
	x[89], x[91] = invbutterfly_shift48(x[89], x[91])
	x[93], x[95] = invbutterfly_shift48(x[93], x[95])
	x[97], x[99] = invbutterfly_shift48(x[97], x[99])
	x[101], x[103] = invbutterfly_shift48(x[101], x[103])
	x[105], x[107] = invbutterfly_shift48(x[105], x[107])
	x[109], x[111] = invbutterfly_shift48(x[109], x[111])
	x[113], x[115] = invbutterfly_shift48(x[113], x[115])
	x[117], x[119] = invbutterfly_shift48(x[117], x[119])
	x[121], x[123] = invbutterfly_shift48(x[121], x[123])
	x[125], x[127] = invbutterfly_shift48(x[125], x[127])
	x[0], x[4] = invbutterfly_null(x[0], x[4])
	x[8], x[12] = invbutterfly_null(x[8], x[12])
	x[16], x[20] = invbutterfly_null(x[16], x[20])
	x[24], x[28] = invbutterfly_null(x[24], x[28])
	x[32], x[36] = invbutterfly_null(x[32], x[36])
	x[40], x[44] = invbutterfly_null(x[40], x[44])
	x[48], x[52] = invbutterfly_null(x[48], x[52])
	x[56], x[60] = invbutterfly_null(x[56], x[60])
	x[64], x[68] = invbutterfly_null(x[64], x[68])
	x[72], x[76] = invbutterfly_null(x[72], x[76])
	x[80], x[84] = invbutterfly_null(x[80], x[84])
	x[88], x[92] = invbutterfly_null(x[88], x[92])
	x[96], x[100] = invbutterfly_null(x[96], x[100])
	x[104], x[108] = invbutterfly_null(x[104], x[108])
	x[112], x[116] = invbutterfly_null(x[112], x[116])
	x[120], x[124] = invbutterfly_null(x[120], x[124])
	x[1], x[5] = invbutterfly_shift72(x[1], x[5])
	x[9], x[13] = invbutterfly_shift72(x[9], x[13])
	x[17], x[21] = invbutterfly_shift72(x[17], x[21])
	x[25], x[29] = invbutterfly_shift72(x[25], x[29])
	x[33], x[37] = invbutterfly_shift72(x[33], x[37])
	x[41], x[45] = invbutterfly_shift72(x[41], x[45])
	x[49], x[53] = invbutterfly_shift72(x[49], x[53])
	x[57], x[61] = invbutterfly_shift72(x[57], x[61])
	x[65], x[69] = invbutterfly_shift72(x[65], x[69])
	x[73], x[77] = invbutterfly_shift72(x[73], x[77])
	x[81], x[85] = invbutterfly_shift72(x[81], x[85])
	x[89], x[93] = invbutterfly_shift72(x[89], x[93])
	x[97], x[101] = invbutterfly_shift72(x[97], x[101])
	x[105], x[109] = invbutterfly_shift72(x[105], x[109])
	x[113], x[117] = invbutterfly_shift72(x[113], x[117])
	x[121], x[125] = invbutterfly_shift72(x[121], x[125])
	x[2], x[6] = invbutterfly_shift48(x[2], x[6])
	x[10], x[14] = invbutterfly_shift48(x[10], x[14])
	x[18], x[22] = invbutterfly_shift48(x[18], x[22])
	x[26], x[30] = invbutterfly_shift48(x[26], x[30])
	x[34], x[38] = invbutterfly_shift48(x[34], x[38])
	x[42], x[46] = invbutterfly_shift48(x[42], x[46])
	x[50], x[54] = invbutterfly_shift48(x[50], x[54])
	x[58], x[62] = invbutterfly_shift48(x[58], x[62])
	x[66], x[70] = invbutterfly_shift48(x[66], x[70])
	x[74], x[78] = invbutterfly_shift48(x[74], x[78])
	x[82], x[86] = invbutterfly_shift48(x[82], x[86])
	x[90], x[94] = invbutterfly_shift48(x[90], x[94])
	x[98], x[102] = invbutterfly_shift48(x[98], x[102])
	x[106], x[110] = invbutterfly_shift48(x[106], x[110])
	x[114], x[118] = invbutterfly_shift48(x[114], x[118])
	x[122], x[126] = invbutterfly_shift48(x[122], x[126])
	x[3], x[7] = invbutterfly_shift24(x[3], x[7])
	x[11], x[15] = invbutterfly_shift24(x[11], x[15])
	x[19], x[23] = invbutterfly_shift24(x[19], x[23])
	x[27], x[31] = invbutterfly_shift24(x[27], x[31])
	x[35], x[39] = invbutterfly_shift24(x[35], x[39])
	x[43], x[47] = invbutterfly_shift24(x[43], x[47])
	x[51], x[55] = invbutterfly_shift24(x[51], x[55])
	x[59], x[63] = invbutterfly_shift24(x[59], x[63])
	x[67], x[71] = invbutterfly_shift24(x[67], x[71])
	x[75], x[79] = invbutterfly_shift24(x[75], x[79])
	x[83], x[87] = invbutterfly_shift24(x[83], x[87])
	x[91], x[95] = invbutterfly_shift24(x[91], x[95])
	x[99], x[103] = invbutterfly_shift24(x[99], x[103])
	x[107], x[111] = invbutterfly_shift24(x[107], x[111])
	x[115], x[119] = invbutterfly_shift24(x[115], x[119])
	x[123], x[127] = invbutterfly_shift24(x[123], x[127])
	x[0], x[8] = invbutterfly_null(x[0], x[8])
	x[16], x[24] = invbutterfly_null(x[16], x[24])
	x[32], x[40] = invbutterfly_null(x[32], x[40])
	x[48], x[56] = invbutterfly_null(x[48], x[56])
	x[64], x[72] = invbutterfly_null(x[64], x[72])
	x[80], x[88] = invbutterfly_null(x[80], x[88])
	x[96], x[104] = invbutterfly_null(x[96], x[104])
	x[112], x[120] = invbutterfly_null(x[112], x[120])
	x[1], x[9] = invbutterfly_shift84(x[1], x[9])
	x[17], x[25] = invbutterfly_shift84(x[17], x[25])
	x[33], x[41] = invbutterfly_shift84(x[33], x[41])
	x[49], x[57] = invbutterfly_shift84(x[49], x[57])
	x[65], x[73] = invbutterfly_shift84(x[65], x[73])
	x[81], x[89] = invbutterfly_shift84(x[81], x[89])
	x[97], x[105] = invbutterfly_shift84(x[97], x[105])
	x[113], x[121] = invbutterfly_shift84(x[113], x[121])
	x[2], x[10] = invbutterfly_shift72(x[2], x[10])
	x[18], x[26] = invbutterfly_shift72(x[18], x[26])
	x[34], x[42] = invbutterfly_shift72(x[34], x[42])
	x[50], x[58] = invbutterfly_shift72(x[50], x[58])
	x[66], x[74] = invbutterfly_shift72(x[66], x[74])
	x[82], x[90] = invbutterfly_shift72(x[82], x[90])
	x[98], x[106] = invbutterfly_shift72(x[98], x[106])
	x[114], x[122] = invbutterfly_shift72(x[114], x[122])
	x[3], x[11] = invbutterfly_shift60(x[3], x[11])
	x[19], x[27] = invbutterfly_shift60(x[19], x[27])
	x[35], x[43] = invbutterfly_shift60(x[35], x[43])
	x[51], x[59] = invbutterfly_shift60(x[51], x[59])
	x[67], x[75] = invbutterfly_shift60(x[67], x[75])
	x[83], x[91] = invbutterfly_shift60(x[83], x[91])
	x[99], x[107] = invbutterfly_shift60(x[99], x[107])
	x[115], x[123] = invbutterfly_shift60(x[115], x[123])
	x[4], x[12] = invbutterfly_shift48(x[4], x[12])
	x[20], x[28] = invbutterfly_shift48(x[20], x[28])
	x[36], x[44] = invbutterfly_shift48(x[36], x[44])
	x[52], x[60] = invbutterfly_shift48(x[52], x[60])
	x[68], x[76] = invbutterfly_shift48(x[68], x[76])
	x[84], x[92] = invbutterfly_shift48(x[84], x[92])
	x[100], x[108] = invbutterfly_shift48(x[100], x[108])
	x[116], x[124] = invbutterfly_shift48(x[116], x[124])
	x[5], x[13] = invbutterfly_shift36(x[5], x[13])
	x[21], x[29] = invbutterfly_shift36(x[21], x[29])
	x[37], x[45] = invbutterfly_shift36(x[37], x[45])
	x[53], x[61] = invbutterfly_shift36(x[53], x[61])
	x[69], x[77] = invbutterfly_shift36(x[69], x[77])
	x[85], x[93] = invbutterfly_shift36(x[85], x[93])
	x[101], x[109] = invbutterfly_shift36(x[101], x[109])
	x[117], x[125] = invbutterfly_shift36(x[117], x[125])
	x[6], x[14] = invbutterfly_shift24(x[6], x[14])
	x[22], x[30] = invbutterfly_shift24(x[22], x[30])
	x[38], x[46] = invbutterfly_shift24(x[38], x[46])
	x[54], x[62] = invbutterfly_shift24(x[54], x[62])
	x[70], x[78] = invbutterfly_shift24(x[70], x[78])
	x[86], x[94] = invbutterfly_shift24(x[86], x[94])
	x[102], x[110] = invbutterfly_shift24(x[102], x[110])
	x[118], x[126] = invbutterfly_shift24(x[118], x[126])
	x[7], x[15] = invbutterfly_shift12(x[7], x[15])
	x[23], x[31] = invbutterfly_shift12(x[23], x[31])
	x[39], x[47] = invbutterfly_shift12(x[39], x[47])
	x[55], x[63] = invbutterfly_shift12(x[55], x[63])
	x[71], x[79] = invbutterfly_shift12(x[71], x[79])
	x[87], x[95] = invbutterfly_shift12(x[87], x[95])
	x[103], x[111] = invbutterfly_shift12(x[103], x[111])
	x[119], x[127] = invbutterfly_shift12(x[119], x[127])
	x[0], x[16] = invbutterfly_null(x[0], x[16])
	x[32], x[48] = invbutterfly_null(x[32], x[48])
	x[64], x[80] = invbutterfly_null(x[64], x[80])
	x[96], x[112] = invbutterfly_null(x[96], x[112])
	x[1], x[17] = invbutterfly_shift90(x[1], x[17])
	x[33], x[49] = invbutterfly_shift90(x[33], x[49])
	x[65], x[81] = invbutterfly_shift90(x[65], x[81])
	x[97], x[113] = invbutterfly_shift90(x[97], x[113])
	x[2], x[18] = invbutterfly_shift84(x[2], x[18])
	x[34], x[50] = invbutterfly_shift84(x[34], x[50])
	x[66], x[82] = invbutterfly_shift84(x[66], x[82])
	x[98], x[114] = invbutterfly_shift84(x[98], x[114])
	x[3], x[19] = invbutterfly_shift78(x[3], x[19])
	x[35], x[51] = invbutterfly_shift78(x[35], x[51])
	x[67], x[83] = invbutterfly_shift78(x[67], x[83])
	x[99], x[115] = invbutterfly_shift78(x[99], x[115])
	x[4], x[20] = invbutterfly_shift72(x[4], x[20])
	x[36], x[52] = invbutterfly_shift72(x[36], x[52])
	x[68], x[84] = invbutterfly_shift72(x[68], x[84])
	x[100], x[116] = invbutterfly_shift72(x[100], x[116])
	x[5], x[21] = invbutterfly_shift66(x[5], x[21])
	x[37], x[53] = invbutterfly_shift66(x[37], x[53])
	x[69], x[85] = invbutterfly_shift66(x[69], x[85])
	x[101], x[117] = invbutterfly_shift66(x[101], x[117])
	x[6], x[22] = invbutterfly_shift60(x[6], x[22])
	x[38], x[54] = invbutterfly_shift60(x[38], x[54])
	x[70], x[86] = invbutterfly_shift60(x[70], x[86])
	x[102], x[118] = invbutterfly_shift60(x[102], x[118])
	x[7], x[23] = invbutterfly_shift54(x[7], x[23])
	x[39], x[55] = invbutterfly_shift54(x[39], x[55])
	x[71], x[87] = invbutterfly_shift54(x[71], x[87])
	x[103], x[119] = invbutterfly_shift54(x[103], x[119])
	x[8], x[24] = invbutterfly_shift48(x[8], x[24])
	x[40], x[56] = invbutterfly_shift48(x[40], x[56])
	x[72], x[88] = invbutterfly_shift48(x[72], x[88])
	x[104], x[120] = invbutterfly_shift48(x[104], x[120])
	x[9], x[25] = invbutterfly_shift42(x[9], x[25])
	x[41], x[57] = invbutterfly_shift42(x[41], x[57])
	x[73], x[89] = invbutterfly_shift42(x[73], x[89])
	x[105], x[121] = invbutterfly_shift42(x[105], x[121])
	x[10], x[26] = invbutterfly_shift36(x[10], x[26])
	x[42], x[58] = invbutterfly_shift36(x[42], x[58])
	x[74], x[90] = invbutterfly_shift36(x[74], x[90])
	x[106], x[122] = invbutterfly_shift36(x[106], x[122])
	x[11], x[27] = invbutterfly_shift30(x[11], x[27])
	x[43], x[59] = invbutterfly_shift30(x[43], x[59])
	x[75], x[91] = invbutterfly_shift30(x[75], x[91])
	x[107], x[123] = invbutterfly_shift30(x[107], x[123])
	x[12], x[28] = invbutterfly_shift24(x[12], x[28])
	x[44], x[60] = invbutterfly_shift24(x[44], x[60])
	x[76], x[92] = invbutterfly_shift24(x[76], x[92])
	x[108], x[124] = invbutterfly_shift24(x[108], x[124])
	x[13], x[29] = invbutterfly_shift18(x[13], x[29])
	x[45], x[61] = invbutterfly_shift18(x[45], x[61])
	x[77], x[93] = invbutterfly_shift18(x[77], x[93])
	x[109], x[125] = invbutterfly_shift18(x[109], x[125])
	x[14], x[30] = invbutterfly_shift12(x[14], x[30])
	x[46], x[62] = invbutterfly_shift12(x[46], x[62])
	x[78], x[94] = invbutterfly_shift12(x[78], x[94])
	x[110], x[126] = invbutterfly_shift12(x[110], x[126])
	x[15], x[31] = invbutterfly_shift6(x[15], x[31])
	x[47], x[63] = invbutterfly_shift6(x[47], x[63])
	x[79], x[95] = invbutterfly_shift6(x[79], x[95])
	x[111], x[127] = invbutterfly_shift6(x[111], x[127])
	x[0], x[32] = invbutterfly_null(x[0], x[32])
	x[64], x[96] = invbutterfly_null(x[64], x[96])
	x[1], x[33] = invbutterfly_shift93(x[1], x[33])
	x[65], x[97] = invbutterfly_shift93(x[65], x[97])
	x[2], x[34] = invbutterfly_shift90(x[2], x[34])
	x[66], x[98] = invbutterfly_shift90(x[66], x[98])
	x[3], x[35] = invbutterfly_shift87(x[3], x[35])
	x[67], x[99] = invbutterfly_shift87(x[67], x[99])
	x[4], x[36] = invbutterfly_shift84(x[4], x[36])
	x[68], x[100] = invbutterfly_shift84(x[68], x[100])
	x[5], x[37] = invbutterfly_shift81(x[5], x[37])
	x[69], x[101] = invbutterfly_shift81(x[69], x[101])
	x[6], x[38] = invbutterfly_shift78(x[6], x[38])
	x[70], x[102] = invbutterfly_shift78(x[70], x[102])
	x[7], x[39] = invbutterfly_shift75(x[7], x[39])
	x[71], x[103] = invbutterfly_shift75(x[71], x[103])
	x[8], x[40] = invbutterfly_shift72(x[8], x[40])
	x[72], x[104] = invbutterfly_shift72(x[72], x[104])
	x[9], x[41] = invbutterfly_shift69(x[9], x[41])
	x[73], x[105] = invbutterfly_shift69(x[73], x[105])
	x[10], x[42] = invbutterfly_shift66(x[10], x[42])
	x[74], x[106] = invbutterfly_shift66(x[74], x[106])
	x[11], x[43] = invbutterfly_shift63(x[11], x[43])
	x[75], x[107] = invbutterfly_shift63(x[75], x[107])
	x[12], x[44] = invbutterfly_shift60(x[12], x[44])
	x[76], x[108] = invbutterfly_shift60(x[76], x[108])
	x[13], x[45] = invbutterfly_shift57(x[13], x[45])
	x[77], x[109] = invbutterfly_shift57(x[77], x[109])
	x[14], x[46] = invbutterfly_shift54(x[14], x[46])
	x[78], x[110] = invbutterfly_shift54(x[78], x[110])
	x[15], x[47] = invbutterfly_shift51(x[15], x[47])
	x[79], x[111] = invbutterfly_shift51(x[79], x[111])
	x[16], x[48] = invbutterfly_shift48(x[16], x[48])
	x[80], x[112] = invbutterfly_shift48(x[80], x[112])
	x[17], x[49] = invbutterfly_shift45(x[17], x[49])
	x[81], x[113] = invbutterfly_shift45(x[81], x[113])
	x[18], x[50] = invbutterfly_shift42(x[18], x[50])
	x[82], x[114] = invbutterfly_shift42(x[82], x[114])
	x[19], x[51] = invbutterfly_shift39(x[19], x[51])
	x[83], x[115] = invbutterfly_shift39(x[83], x[115])
	x[20], x[52] = invbutterfly_shift36(x[20], x[52])
	x[84], x[116] = invbutterfly_shift36(x[84], x[116])
	x[21], x[53] = invbutterfly_shift33(x[21], x[53])
	x[85], x[117] = invbutterfly_shift33(x[85], x[117])
	x[22], x[54] = invbutterfly_shift30(x[22], x[54])
	x[86], x[118] = invbutterfly_shift30(x[86], x[118])
	x[23], x[55] = invbutterfly_shift27(x[23], x[55])
	x[87], x[119] = invbutterfly_shift27(x[87], x[119])
	x[24], x[56] = invbutterfly_shift24(x[24], x[56])
	x[88], x[120] = invbutterfly_shift24(x[88], x[120])
	x[25], x[57] = invbutterfly_shift21(x[25], x[57])
	x[89], x[121] = invbutterfly_shift21(x[89], x[121])
	x[26], x[58] = invbutterfly_shift18(x[26], x[58])
	x[90], x[122] = invbutterfly_shift18(x[90], x[122])
	x[27], x[59] = invbutterfly_shift15(x[27], x[59])
	x[91], x[123] = invbutterfly_shift15(x[91], x[123])
	x[28], x[60] = invbutterfly_shift12(x[28], x[60])
	x[92], x[124] = invbutterfly_shift12(x[92], x[124])
	x[29], x[61] = invbutterfly_shift9(x[29], x[61])
	x[93], x[125] = invbutterfly_shift9(x[93], x[125])
	x[30], x[62] = invbutterfly_shift6(x[30], x[62])
	x[94], x[126] = invbutterfly_shift6(x[94], x[126])
	x[31], x[63] = invbutterfly_shift3(x[31], x[63])
	x[95], x[127] = invbutterfly_shift3(x[95], x[127])
	x[0], x[64] = invbutterfly_null(x[0], x[64])
	x[1], x[65] = invbutterfly_mul(x[1], x[65], 274873712576)
	x[2], x[66] = invbutterfly_shift93(x[2], x[66])
	x[3], x[67] = invbutterfly_mul(x[3], x[67], 34359214072)
	x[4], x[68] = invbutterfly_shift90(x[4], x[68])
	x[5], x[69] = invbutterfly_mul(x[5], x[69], 4294901759)
	x[6], x[70] = invbutterfly_shift87(x[6], x[70])
	x[7], x[71] = invbutterfly_mul(x[7], x[71], 2305843009213685760)
	x[8], x[72] = invbutterfly_shift84(x[8], x[72])
	x[9], x[73] = invbutterfly_mul(x[9], x[73], 288230376151710720)
	x[10], x[74] = invbutterfly_shift81(x[10], x[74])
	x[11], x[75] = invbutterfly_mul(x[11], x[75], 36028797018963840)
	x[12], x[76] = invbutterfly_shift78(x[12], x[76])
	x[13], x[77] = invbutterfly_mul(x[13], x[77], 4503599627370480)
	x[14], x[78] = invbutterfly_shift75(x[14], x[78])
	x[15], x[79] = invbutterfly_mul(x[15], x[79], 562949953421310)
	x[16], x[80] = invbutterfly_shift72(x[16], x[80])
	x[17], x[81] = invbutterfly_mul(x[17], x[81], 4611756386097823744)
	x[18], x[82] = invbutterfly_shift69(x[18], x[82])
	x[19], x[83] = invbutterfly_mul(x[19], x[83], 576469548262227968)
	x[20], x[84] = invbutterfly_shift66(x[20], x[84])
	x[21], x[85] = invbutterfly_mul(x[21], x[85], 72058693532778496)
	x[22], x[86] = invbutterfly_shift63(x[22], x[86])
	x[23], x[87] = invbutterfly_mul(x[23], x[87], 9007336691597312)
	x[24], x[88] = invbutterfly_shift60(x[24], x[88])
	x[25], x[89] = invbutterfly_mul(x[25], x[89], 1125917086449664)
	x[26], x[90] = invbutterfly_shift57(x[26], x[90])
	x[27], x[91] = invbutterfly_mul(x[27], x[91], 140739635806208)
	x[28], x[92] = invbutterfly_shift54(x[28], x[92])
	x[29], x[93] = invbutterfly_mul(x[29], x[93], 17592454475776)
	x[30], x[94] = invbutterfly_shift51(x[30], x[94])
	x[31], x[95] = invbutterfly_mul(x[31], x[95], 2199056809472)
	x[32], x[96] = invbutterfly_shift48(x[32], x[96])
	x[33], x[97] = invbutterfly_mul(x[33], x[97], 274882101184)
	x[34], x[98] = invbutterfly_shift45(x[34], x[98])
	x[35], x[99] = invbutterfly_mul(x[35], x[99], 34360262648)
	x[36], x[100] = invbutterfly_shift42(x[36], x[100])
	x[37], x[101] = invbutterfly_mul(x[37], x[101], 4295032831)
	x[38], x[102] = invbutterfly_shift39(x[38], x[102])
	x[39], x[103] = invbutterfly_mul(x[39], x[103], 2305843009213702144)
	x[40], x[104] = invbutterfly_shift36(x[40], x[104])
	x[41], x[105] = invbutterfly_mul(x[41], x[105], 288230376151712768)
	x[42], x[106] = invbutterfly_shift33(x[42], x[106])
	x[43], x[107] = invbutterfly_mul(x[43], x[107], 36028797018964096)
	x[44], x[108] = invbutterfly_shift30(x[44], x[108])
	x[45], x[109] = invbutterfly_mul(x[45], x[109], 4503599627370512)
	x[46], x[110] = invbutterfly_shift27(x[46], x[110])
	x[47], x[111] = invbutterfly_mul(x[47], x[111], 562949953421314)
	x[48], x[112] = invbutterfly_shift24(x[48], x[112])
	x[49], x[113] = invbutterfly_mul(x[49], x[113], 13835128420805115905)
	x[50], x[114] = invbutterfly_shift21(x[50], x[114])
	x[51], x[115] = invbutterfly_mul(x[51], x[115], 17870292113338400769)
	x[52], x[116] = invbutterfly_shift18(x[52], x[116])
	x[53], x[117] = invbutterfly_mul(x[53], x[117], 18374687574905061377)
	x[54], x[118] = invbutterfly_shift15(x[54], x[118])
	x[55], x[119] = invbutterfly_mul(x[55], x[119], 18437737007600893953)
	x[56], x[120] = invbutterfly_shift12(x[56], x[120])
	x[57], x[121] = invbutterfly_mul(x[57], x[121], 18445618186687873025)
	x[58], x[122] = invbutterfly_shift9(x[58], x[122])
	x[59], x[123] = invbutterfly_mul(x[59], x[123], 18446603334073745409)
	x[60], x[124] = invbutterfly_shift6(x[60], x[124])
	x[61], x[125] = invbutterfly_mul(x[61], x[125], 18446726477496979457)
	x[62], x[126] = invbutterfly_shift3(x[62], x[126])
	x[63], x[127] = invbutterfly_mul(x[63], x[127], 18446741870424883713)

}
