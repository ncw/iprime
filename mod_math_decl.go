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

//go:noescape

func mod_shift0to31x3(x uint64) (ret uint64)

//go:noescape

func mod_shift32to63x36(x uint64) (ret uint64)

//go:noescape

func mod_shift64to95x72(x uint64) (ret uint64)

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
