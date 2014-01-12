// Modular FFT routines

package main

import ()

// Bit reverse - slow
func bit_reverse(log_n uint8, x []uint64) {
	n := 1 << log_n
	for i := 0; i < n; i++ {
		m := 0

		for j := uint8(0); j < log_n; j++ {
			m |= ((i >> j) & 1) << (log_n - j - 1)
		}

		if m > i {
			x[i], x[m] = x[m], x[i]
		}
	}
}

// A slow O(n^2) fft for testing purposes
func _fft_slow(n uint, x []uint64, inverse bool) {
	var w uint64 = MOD_W
	if inverse {
		w = MOD_INVW
	}
	wk := uint64(1)
	t := make([]uint64, n)

	for k := uint(0); k < n; k++ {
		sum := uint64(0)
		wj := uint64(1)
		// printf("wk(%02i) = %s\n", k, mod_string(wk));
		for j := uint(0); j < n; j++ {
			// printf("  wj(%02i) = %s\n", k, mod_string(wj));
			sum = mod_add(sum, mod_mul(wj, x[j]))
			wj = mod_mul(wj, wk)
		}
		t[k] = sum
		wk = mod_mul(wk, w)
	}

	copy(x, t)
}

// A slow O(n^2) fft
func fft_slow(log_n uint8, x []uint64) {
	_fft_slow(1<<log_n, x, false)
}

// A slow O(n^2) inverse fft
func invfft_slow(log_n uint8, x []uint64) {
	_fft_slow(1<<log_n, x, true)
}

// A fastish O(n log n) FFT
//
// Output is bit-reversed
func fft_fastish(log_n uint8, x []uint64) {
	var d uint64 = MOD_W

	for k := log_n; k >= 1; k-- {
		m := uint(1) << k
		c := m >> 1
		var w uint64 = 1
		for j := uint(0); j < c; j++ {
			for r := uint(0); r < n; r += m {
				a := r + j
				b := a + c
				u := x[a]
				v := x[b]
				x[a] = mod_add(u, v)
				x[b] = mod_mul(mod_sub(u, v), w)
			}
			w = mod_mul(w, d)
		}
		d = mod_mul(d, d)
	}
}

// A fastish O(n log n) Inverse FFT
//
// Input should be bit-reversed
func invfft_fastish(log_n uint8, x []uint64) {
	for k := uint8(1); k <= log_n; k++ {
		m := uint(1) << k
		c := m >> 1
		z := uint64((1 << (uint32(log_n) - uint32(k))))
		d := mod_pow(MOD_INVW, z)
		w := uint64(1)
		for j := uint(0); j < c; j++ {
			for r := uint(0); r < n; r += m {
				a := r + j
				b := a + c
				u := x[a]
				v := mod_mul(w, x[b])
				x[a] = mod_add(u, v)
				x[b] = mod_sub(u, v)
			}
			w = mod_mul(w, d)
		}
	}
}
