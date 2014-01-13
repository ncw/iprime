// Test the modular integer ffts

package main

import (
	"testing"
)

var (
	rnd []uint64
	z   []uint64
)

func test_fft_init(_log_n uint8) {
	log_n = _log_n
	n = uint(1) << log_n
	fft_init()
	rnd = make([]uint64, n)
	z = make([]uint64, n)
	/*
		four_step = 0;
		log_fft_rows = log_n / 2;
		log_fft_cols = log_n - log_fft_rows;
		fft_rows = 1<<log_fft_rows;
		fft_cols = 1<<log_fft_cols;
	*/
	/* make some random numbers */
	for i := range rnd {
		rnd[i] = mod_rnd()
	}
}

func normaliseAndCompare(t *testing.T, what string) {
	// normalise
	inv_n := mod_inv(uint64(n))
	for i := uint(0); i < n; i++ {
		x[i] = mod_mul(x[i], inv_n)
	}

	// compare
	for i := uint(0); i < n; i++ {
		if x[i] != rnd[i] {
			t.Errorf("%s %i: orig=%d xform=%d ERROR\n", what, i, x[i], rnd[i])
		}
	}
}

func testSlowVsSlow(t *testing.T) {
	copy(x, rnd)
	fft_slow(log_n, x)
	invfft_slow(log_n, x)
	normaliseAndCompare(t, "Slow vs Slow")
}

func testSlowVsFastish(t *testing.T) {
	copy(x, rnd)
	fft_slow(log_n, x)
	bit_reverse(log_n, x)
	invfft_fastish(log_n, x)
	normaliseAndCompare(t, "Slow vs Fastish")

	copy(x, rnd)
	fft_fastish(log_n, x)
	bit_reverse(log_n, x)
	invfft_slow(log_n, x)
	normaliseAndCompare(t, "Slow vs Fastish Inv")
}

func testFastishVsFastish(t *testing.T) {
	copy(x, rnd)
	fft_fastish(log_n, x)
	invfft_fastish(log_n, x)
	normaliseAndCompare(t, "Fastish vs Fastish")
}

func testFastishVsShift(t *testing.T) {
	copy(x, rnd)
	fft_shift(log_n, x)
	invfft_fastish(log_n, x)
	normaliseAndCompare(t, "Fastish vs Shift")

	// copy(x, rnd)
	// fft_fastish(log_n, x)
	// invfft_shift(log_n, x)
	// normaliseAndCompare(t, "Inv Fastish vs Shift")
}

func TestFFT1(t *testing.T) {
	test_fft_init(1)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
	testFastishVsShift(t)
}

func TestFFT2(t *testing.T) {
	test_fft_init(2)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
	testFastishVsShift(t)
}

func TestFFT3(t *testing.T) {
	test_fft_init(3)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
	testFastishVsShift(t)
}

func TestFFT4(t *testing.T) {
	test_fft_init(4)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
	testFastishVsShift(t)
}

func TestFFT5(t *testing.T) {
	test_fft_init(5)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
	testFastishVsShift(t)
}

func TestFFT6(t *testing.T) {
	test_fft_init(6)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
	testFastishVsShift(t)
}

func TestFFT7(t *testing.T) {
	test_fft_init(7)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT8(t *testing.T) {
	test_fft_init(8)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT9(t *testing.T) {
	test_fft_init(9)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT10(t *testing.T) {
	test_fft_init(10)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT11(t *testing.T) {
	test_fft_init(11)
	testSlowVsSlow(t)
	testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT12(t *testing.T) {
	test_fft_init(12)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT13(t *testing.T) {
	test_fft_init(13)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT14(t *testing.T) {
	test_fft_init(14)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT15(t *testing.T) {
	test_fft_init(15)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT16(t *testing.T) {
	test_fft_init(16)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT17(t *testing.T) {
	test_fft_init(17)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT18(t *testing.T) {
	test_fft_init(18)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT19(t *testing.T) {
	test_fft_init(19)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func TestFFT20(t *testing.T) {
	test_fft_init(20)
	// testSlowVsSlow(t)
	// testSlowVsFastish(t)
	testFastishVsFastish(t)
}

func test_fft(log_n uint8) {

	// /* check fast vs slow */
	// printf("Check fast vs slow\n");
	// memcpy(x, rnd, size);
	// printf("fast\n");
	// if (four_step)
	//     fft_four_step(x, log_fft_cols, log_fft_rows, FALSE);
	// else
	//     fft_fastish(log_n, x);

	// if (four_step)
	// {
	//     for (i = 0; i < fft_rows; i++)
	//         bit_reverse(log_fft_cols, &x[i << log_fft_rows]);
	//     mod_transpose_square(x, fft_cols);
	//     for (i = 0; i < fft_rows; i++)
	//         bit_reverse(log_fft_cols, &x[i << log_fft_rows]);
	//     mod_transpose_square(x, fft_cols);
	// }

	// printf("vs slow\n");
	// memcpy(z, rnd, size);
	// fft_slow(log_n, z);
	// bit_reverse((int)log_n, z);
	// //fft_fastish((int)log_n, z);
	// if (four_step)
	//     bit_reverse((int)log_n, z);

	// printf("done\n");

	// for (i = 0; i < n; i++)
	// {
	//     if (x[i] != z[i])
	//     {
	//         printf("%04i: slow=%s fast=%s ERROR\n", i, mod_string(x[i]), mod_string(z[i]));
	//     }
	// }

	// /* check fast + fast inverse */
	// printf("Check fast + inverse\n");
	// memcpy(x, rnd, size);
	// if (four_step)
	//     fft_four_step(x, log_fft_cols, log_fft_rows, FALSE);
	// else
	//     fft_fastish(log_n, x);
	// if (four_step)
	//     fft_four_step(x, log_fft_cols, log_fft_rows, TRUE);
	// else
	//     invfft_fastish(log_n, x);
	// inv_n = (u64)n;
	// inv_n = mod_inv(inv_n);
	// /* normalise */
	// for (i = 0; i < n; i++)
	//     x[i] = mod_mul(x[i], inv_n);
	// for (i = 0; i < n; i++)
	// {
	//     if (x[i] != rnd[i])
	//     {
	// 	    printf("%04i: orig=%s fast=%s ERROR\n", i, mod_string(x[i]), mod_string(rnd[i]));
	//     }
	// }
	// printf("done\n");

	// printf("freeing blocks\n");
	// //printf("x=%p\n", x);
	// //printf("y=%p\n", y);
	// //printf("z=%p\n", z);
	// xafree(x);
	// xafree(z);
	// xafree(rnd);
}

// int main(int argc, char ** argv)
// {
//     int log_n;

//     argc--;
//     argv++;
//     if (argc != 1)
//         fatal("Usage:\n %s N [n]\nwhere N is log2(FFT) size to test\n", argv[-1]);
//     log_n = atoi(argv[0]);

//     test_fft(log_n);
//     return EXIT_SUCCESS;
// }
