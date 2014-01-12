iprime
======

This is intended to be an optimised prime finder using integer only code.

It is written in Go for portability between platforms and the easy of
adding assembler to the code.

The optimised bits aren't finished yet though!

This was based on
[my ioccc 2012 entry](http://www.craig-wood.com/nick/articles/ioccc-2012-mersenne-prime-checker/)
where you can find some more information until I get round to
rewriting it specifically for this.

Notes
-----

This program proves the primality of [Mersenne
primes](http://en.wikipedia.org/wiki/Mersenne_prime) of the form `2^p
- 1`.  The biggest known primes are of this form because there is an
efficient test for primality - the [Lucas Lehmer
test](http://mathworld.wolfram.com/Lucas-LehmerTest.html).

If you run the program and it produces `0x0000000000000000` (0 in hex)
as an answer then you've found a prime of the form `2^p-1`.  If it
produces anything else then `2^p-1` isn't prime.  The printed result is
the bottom 64 bits of the last iteration of the Lucas Lehmer test.

The test is implemented using arithmetic module a prime just less than
`2^64`, `0xFFFFFFFF00000001` = `2^64-2^32+1` (p from now on), with
some rather special properties.

The most important of which is that arithmetic modulo p can all be
done without using divisions or modulo operations (which are really
slow).

The next important property is that p-1 has the following factors

    2^32 * 3 * 5 * 17 * 257 * 65537

All those factors of 2 make a [Fast Fourier
Transform](http://mathworld.wolfram.com/FastFourierTransform.html)
over the [Galois Field](http://mathworld.wolfram.com/FiniteField.html)
GF(p) which is arithmetic modulo p possible.

To make a truly efficient Mersenne primality prover it implements the
[IBDWT](http://en.wikipedia.org/wiki/Irrational_base_discrete_weighted_transform)
a variant of the FFT using an irrational base.  This halves the size
of the FFT array by combining the modulo and multiply steps from the
Lucas Lehmer test into one.

It can use an irrational base in a discrete field where every element
is an integer, because the p chosen has n-th roots of 2 where n is up
to 26, which means that an IBDWT can be defined for FFT sizes up to
`2^26`. 7 is a primitive root of GF(p) so all follows from that!

This all integer IBDWT isn't susceptible to the round off errors that
plague the floating point implementation, but does suffer in the speed
stakes from the fact that modern CPUs have much faster floating point
units than integer units.

For example to check if `2^18000000-1` is prime (a 5,418,539 digit
number) [Prime95](http://www.mersenne.org/freesoft/) (the current
speed demon) uses a `2^20` FFT size and each iteration takes it about
25ms. On similar hardware it this program takes about 1.2s also using
an FFT size of `2^20`!  Prime95 is written in optimised assembler to use
SSE3 whereas this program was squeezed into 2k of portable C with a
completely unoptimised FFT!

This program can do an FFT of up to `2^26` entries.  Each entry can have
up to 18 bits in (as `2*18+26 <= 63`), which means that the biggest
exponent that it can test for primality is `18*(2^26)-1 = 1,207,959,551`.
This would be number of 363 million decimal digits so could be used to
find a 100 million digit prime and scoop the [EFF $150,000
prize](https://www.eff.org/awards/coop)!  It would take rather a long
time though...

The `-iterations` argument to the program allows the number of
iterations to be put in, and using this it has been validated against
the Prime95 test suite.

The program comes with a unit test - run it like

  go test -v -short

If if you have lots of time to spare

  go test -v

Roadmap
-------

Implement
  * FFT for small sizes which only have shifts
  * Optimised ffts with shifts
  * Four step fft so can use these ffts and improve cache locality
  * Assembler primitives
  * Assembler FFTs

License
-------

This is free software under the terms of MIT license (check COPYING file
included in this package).

Author
------

- Nick Craig-Wood <nick@craig-wood.com>
