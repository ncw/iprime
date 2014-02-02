// Declarations for assembler routines

// +build amd64

package main

//go:noescape

func mod_sub(x, y uint64) (ret uint64)

//go:noescape

func mod_add(x, y uint64) (ret uint64)

//go:noescape

func mod_reduce(b, a uint64) (ret uint64)

//go:noescape

func mod_mul(x, y uint64) (ret uint64)

//go:noescape

func mod_sqr(x uint64) (ret uint64)


func mod_shift3(x uint64) (ret uint64)
func mod_shift6(x uint64) (ret uint64)
func mod_shift9(x uint64) (ret uint64)
func mod_shift12(x uint64) (ret uint64)
func mod_shift15(x uint64) (ret uint64)
func mod_shift18(x uint64) (ret uint64)
func mod_shift21(x uint64) (ret uint64)
func mod_shift24(x uint64) (ret uint64)
func mod_shift27(x uint64) (ret uint64)
func mod_shift30(x uint64) (ret uint64)
func mod_shift33(x uint64) (ret uint64)
func mod_shift36(x uint64) (ret uint64)
func mod_shift39(x uint64) (ret uint64)
func mod_shift42(x uint64) (ret uint64)
func mod_shift45(x uint64) (ret uint64)
func mod_shift48(x uint64) (ret uint64)
func mod_shift51(x uint64) (ret uint64)
func mod_shift54(x uint64) (ret uint64)
func mod_shift57(x uint64) (ret uint64)
func mod_shift60(x uint64) (ret uint64)
func mod_shift63(x uint64) (ret uint64)
func mod_shift66(x uint64) (ret uint64)
func mod_shift69(x uint64) (ret uint64)
func mod_shift72(x uint64) (ret uint64)
func mod_shift75(x uint64) (ret uint64)
func mod_shift78(x uint64) (ret uint64)
func mod_shift81(x uint64) (ret uint64)
func mod_shift84(x uint64) (ret uint64)
func mod_shift87(x uint64) (ret uint64)
func mod_shift90(x uint64) (ret uint64)
func mod_shift93(x uint64) (ret uint64)

func fft0(x []uint64)
func fft1(x []uint64)
func fft2(x []uint64)
func fft3(x []uint64)
func fft4(x []uint64)
func fft5(x []uint64)
func fft6(x []uint64)
func fft7(x []uint64)
func fft8(x []uint64)
func fft9(x []uint64)
func fft10(x []uint64)
func fft11(x []uint64)
func fft12(x []uint64)

func invfft0(x []uint64)
func invfft1(x []uint64)
func invfft2(x []uint64)
func invfft3(x []uint64)
func invfft4(x []uint64)
func invfft5(x []uint64)
func invfft6(x []uint64)
func invfft7(x []uint64)
func invfft8(x []uint64)
func invfft9(x []uint64)
func invfft10(x []uint64)
func invfft11(x []uint64)
func invfft12(x []uint64)
