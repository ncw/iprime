iprime
======

This is a Mersenne prime checker using integer only code.

It is written in Go for portability between platforms and the easy of
adding assembler to the code.

The code is working but the FFTs need more work.

There is optimised assembler for ARM and AMD64 architectures.

This code is a port of
[ARMPrime](http://www.craig-wood.com/nick/armprime/) software which ran
on RISC OS only and was never very portable.

I made a first cut at making ARMPrime more portable in
[my ioccc 2012 entry](http://www.craig-wood.com/nick/articles/ioccc-2012-mersenne-prime-checker/) which was written in C.

I decided to re-write again in Go because of the cross platform
support and the fact that Go integrates really nicely with assembler
code.  The testing suite has come in really useful as have the
benchmarking suite.

Install
-------

Iprime is a Go program and comes as a single binary file.

Download the relevant binary from

- http://www.craig-wood.com/nick/pub/iprime/

Or alternatively if you have Go installed use

    go get github.com/ncw/iprime

and this will build the binary in `$GOPATH/bin`.  You can then modify
the source and submit patches.

Usage
-----

Use `iprime -h` to see all the options.

    Mersenne prime tester
    
    Usage:
    
    ./iprime [options] q
    
    where q = Mersenne exponent to test
    
    Options:
      -cpuprofile="": Write cpu profile to file
      -fft-size=0: minimum size for FFT (2**n)
      -iterations=0: Number of iterations to check run - 0 for full test

The `-iterations` argument to the program allows the number of
iterations to be put in, and using this it has been validated against
the Prime95 test suite.

The program comes with a unit test - run it like

    go test -v -short

If if you have lots of time to spare

    go test -v

Example output

    $ ./iprime 521
    Testing 2**521-1 with fft size 2**5 for 0 iterations
    Residue 0x0000000000000000
    That took 4.424103ms for 519 iterations which is 8.524us per iteration

A residue of 0 after the full number of iterations means you've found
a mersenne prime!

How it works
------------

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
units than integer units.  This is particularly noticeable in the x86
works where 64 bit integer multiplies take 64 cycles.  However ARM
chips, which this was originally designed for, 32x32 -> 64 multiples
only take 1 cycle.

For example to check if `2^18000000-1` is prime (a 5,418,539 digit
number) [Prime95](http://www.mersenne.org/freesoft/) (the current
speed demon) uses a `2^20` FFT size and each iteration takes it about
25ms. On similar hardware it this program takes about 1.2s also using
an FFT size of `2^20`!

This program can do an FFT of up to `2^26` entries.  Each entry can have
up to 18 bits in (as `2*18+26 <= 63`), which means that the biggest
exponent that it can test for primality is `18*(2^26)-1 = 1,207,959,551`.
This would be number of 363 million decimal digits so could be used to
find a 100 million digit prime and scoop the [EFF $150,000
prize](https://www.eff.org/awards/coop)!  It would take rather a long
time though...

Roadmap
-------

Implement
  * FFT for small sizes which only have shifts - DONE
  * Optimised ffts with shifts - DONE
  * Four step fft so can use these ffts and improve cache locality - DONE but needs more work
  * Assembler primitives - DONE
  * Assembler FFTs - DONE
  * New FFT schemes (like the one used in FFTW)
  * Save progress
  * PrimeNet support

License
-------

This is free software under the terms of MIT license (check COPYING file
included in this package).

Author
------

- Nick Craig-Wood <nick@craig-wood.com>
