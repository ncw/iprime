// mersenne - mersenne testing
//
// FIXME this code was converted from the ARM prime project and needs a
// lot of tidying up - getting rid of global variables etc.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

var (
	digit_weight   []uint64
	digit_unweight []uint64

	exponent          uint64
	digit_width0      uint8  // bits in a small digit
	digit_width1      uint8  // bits in a large digit
	digit_width_0_max uint32 // max size of a small digit
	digit_widths      []uint8

	// Flags
	cpuprofile = flag.String("cpuprofile", "", "Write cpu profile to file")
	iterations = flag.Uint64("iterations", 0, "Number of iterations to check run - 0 for full test")
)

// Try to do dwt...
//
// n is size of transform
// p is the exponent we want to test
// i is the number of the element
//
//     if (2*(pMersenne/FFTLEN) + LG2_FFTLEN >= 62*NPFFT) {
//         fprintf(stderr, "FFTLEN = %s insufficient for pMersenne = %s\n",
//                 uint64out(FFTLEN), uint64out(pMersenne));
//         exit(1);
//     }
//
// return false for failed true for ok
func mersenne_initialise() bool {
	width := exponent / uint64(n)

	// Make sure the FFT is long enough so that each 'digit' can't
	// overflow a 63 bit number (mod p is slightly less that 64
	// bits) after the convolution
	// Some digits are (w+1) wide so use this for safety
	// (w+1)*2+log_n >= 63
	if 2*width+uint64(log_n) >= 61 {
		return false
	}

	// fft_initialise (to make root2)
	fft_init()

	digit_width0 = uint8(width)
	digit_width_0_max = uint32(1) << width
	digit_width1 = uint8(width) + 1

	// memory allocation
	digit_weight = make([]uint64, n)
	digit_unweight = make([]uint64, n)
	digit_widths = make([]uint8, n)

	// digit weights
	digit_weight[0] = 1
	digit_unweight[0] = mod_inv(uint64(n))
	old_addr := uint64(0)
	for i := uint(0); i <= n; i++ {
		t := uint64(exponent) * uint64(i)
		r := t % uint64(n)
		addr := t / uint64(n)
		if r>>32 != 0 {
			return false
		}
		if uint32(r) != 0 { // do ceil
			addr++
		}
		if addr>>32 != 0 {
			return false
		}

		// bit position for digit[i] is ceil((exponent * i) / n)
		if i > 0 {
			digit_width := addr - old_addr
			digit_widths[i-1] = uint8(digit_width)
			if digit_width != width && digit_width != width+1 {
				return false
			}
			// printf("digit_widths[%i] = %i from %i to %i\n", i-1, digit_widths[i-1], o, a-1);

			// dwt weight is 2^(1 - ((exponent * i mod n)/n))
			if i < n {
				r = uint64(n) - r
				digit_weight[i] = mod_pow(root2, r)
				digit_unweight[i] = mod_inv(mod_mul(digit_weight[i], uint64(n)))
			}
		}

		old_addr = addr
	}

	return true
}

// Calls mersenne_initialise with increasing sizes until we find a bit enough FFT size
func mersenne_auto_initialise() {
	for log_n = 0; log_n <= 26; log_n++ {
		n = 1 << log_n
		if mersenne_initialise() {
			return
		}
	}
	log.Fatal("Exponent too big")
}

// Return the bottom 64 bits
// Assumes a carry propagated array where all digits are within their widths
// And that all digit widths are <= 32
//
// If the residue is 0 then it checks the whole array to double check
// that is zero for a proper primality check
func mersenne_residue() uint64 {
	i := uint(0)
	j := uint(0)
	r := uint64(0)
	for ; i < 64 && j < n; i, j = i+uint(digit_widths[j]), j+1 {
		r |= x[j] << i
	}
	if r != 0 {
		return r
	}
	r = 0
	for j = 0; j < n; j++ {
		r |= x[j]
	}
	return r
}

// This adds an uint32 to x
// We assume that x < 2^minimum_digit_width
//
// It assumes that x has had the first round of carry propagation done on it
// already so each digit[i] is < 2^digit_widths[i] < 2^32
func mersenne_add32(c uint32, i uint) {
	for c != 0 {
		for ; i < n; i++ {
			y := uint64(1) << digit_widths[i]
			x[i] += uint64(c)
			if x[i] >= y {
				x[i] -= y
				c = 1
			} else {
				return // done if no carry
			}
		}
		//         printf("Wrapping round the end in mersenne_add32\n");
		i = 0
	}
}

// This subtracts an uint32 from x
// We assume that x < 2^minimum_digit_width
//
// and that x has had the first round of carry propagation done on it
// already so each digit[i] is < 2^digit_widths[i] < 2^32
func mersenne_sub32(c uint32) {
	for c != 0 {
		for i := uint(0); i < n; i++ {
			y := uint64(1) << digit_widths[i]
			x[i] -= uint64(c)
			if x[i] >= y {
				x[i] += y
				c = 1
			} else {
				return // done if no carry
			}
		}
		//         printf("Wrapping round the end in mersenne_sub32\n");
	}
}

// This adds an uint64 to x
//
// It assumes that x has had the first round of carry propagation done on it
// already so each digit[i] is < 2^digit_widths[i] < 2^32
func mersenne_add64(c uint64) {
	for c != 0 {
		for i := uint(0); i < n; i++ {
			x[i] = mod_adc(x[i], digit_widths[i], &c)
			t := uint32(c)
			if (c>>32) != 0 && t < digit_width_0_max {
				if t != 0 {
					mersenne_add32(t, i+1) // carry in 32 bits if possible
				}
				return // finished if carry is 0
			}
		}
		// printf("Wrapping round the end in mersenne_add64\n");
	}
}

// This does one interation
func mersenne_mul() {
	c := uint64(0)

	// weight the input
	mod_vector_mul(n, x, digit_weight)

	// transform
	fft_fastish(log_n, x)

	// point multiply
	mod_vector_sqr(n, x)

	// untransform
	invfft_fastish(log_n, x)

	// unweight and normalise the output
	mod_vector_mul(n, x, digit_unweight)

	// carry propagation
	for i := uint(0); i < n; i++ {
		// printf("x[%i]=0x%016llX, carry=0x%016llX\n", i, x[i], carry);
		x[i] = mod_adc(x[i], digit_widths[i], &c)
		// printf("x[%i]=0x%016llX, carry=0x%016llX\n", i, x[i], carry);
	}
	if c != 0 {
		// printf("Wrapping carry in mersenne_mul carry propagation\n");
		mersenne_add64(c)
	}

	// subtract 2
	mersenne_sub32(2)
}

// Sets the mersenne array up and runs it for the number of iterations asked for
func mersenne_run(iterations uint64) {
	if iterations == 0 {
		iterations = exponent - 2
	}
	x[0] = 4
	for i := uint64(0); i < iterations; i++ {
		mersenne_mul()
	}
}

// syntaxError prints the syntax
func syntaxError() {
	fmt.Fprintf(os.Stderr, `Mersenne prime tester

Usage:

prog [options] q

where q = Mersenne exponent to test

Options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.Usage = syntaxError
	flag.Parse()
	args := flag.Args()

	// Setup profiling if desired
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if len(args) != 1 {
		syntaxError()
		log.Fatal("No exponent supplied")
	}

	var err error
	exponent, err = strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		syntaxError()
		log.Fatalf("Couldn't parse exponent: %v\n", err)
	}

	mersenne_auto_initialise()

	fmt.Printf("Testing 2**%d-1 with fft size 2**%d for %d iterations\n", exponent, log_n, *iterations)
	mersenne_run(*iterations)
	fmt.Printf("Residue 0x%016X\n", mersenne_residue())
}
