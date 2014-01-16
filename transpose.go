// Transpose a matrix

// Based on code posted on sci.math.num-analysis by Jason Stratos Papadopoulos
//
// https://groups.google.com/forum/?hl=en#!msg/sci.math.num-analysis/qbqu8eEbaZ0/xDCs-7MRAmgJ

// The functions here are for a very fast matrix transpose, needed
// for Bailey's four-step FFT. Some assumptions: the matrix must
// be square, must be a power of two on a side, and must be stored
// as a 1-D array. It uses block transposition; divide the matrix
// into blocks of a certain optimal size; then transpose the diagonal
// blocks and transpose/switch the off-diagonal blocks. BL below is the
// power of two defining block size.
//
// There's plenty of room for assembly optimization, but right now the
// time is dominated by cache misses, and tweaking code here and there
// won't help too much.

package main

const BL = 5

// Move a BLxBL block from matrix "startpt" (of size
// 1<<startlength on a side) to matrix "destpt" (of
// size 1<<destlength on a side).
func copyblock(start []uint64, startpt, startlength int, dest []uint64, destpt, destlength int) {
	for i := 0; i < (1 << BL); i++ {
		// FIXME do we need to get dest in the cache before starting the copy?
		copy(dest[destpt:destpt+(1<<BL)], start[startpt:startpt+(1<<BL)])
		startpt += startlength /* next row of startpt */
		destpt += destlength   /* next row of destpt */
	}
}

// This function switches the columns in the BLxBL block
// "startpt" with the rows of the BLxBL block in "destpt".
// This also has the effect of transposing the block in
// destpt, which need only then be copied into its appropriate
// place. The code below pairs nicely, and performs the
// copying while all the data involved is in cache.
func copytranspose(start []uint64, startpt int, dest []uint64, destpt int, destlength int) {
	for i := 0; i < (1 << BL); i++ {
		src := startpt
		dst := destpt
		for j := 0; j < (1 << BL); j += 8 {
			dest[dst], start[src] = start[src], dest[dst]
			dest[dst+1], start[src+1<<BL] = start[src+1<<BL], dest[dst+1]
			dest[dst+2], start[src+2<<BL] = start[src+2<<BL], dest[dst+2]
			dest[dst+3], start[src+3<<BL] = start[src+3<<BL], dest[dst+3]
			dest[dst+4], start[src+4<<BL] = start[src+4<<BL], dest[dst+4]
			dest[dst+5], start[src+5<<BL] = start[src+5<<BL], dest[dst+5]
			dest[dst+6], start[src+6<<BL] = start[src+6<<BL], dest[dst+6]
			dest[dst+7], start[src+7<<BL] = start[src+7<<BL], dest[dst+7]
			src += 8 << BL
			dst += 8
		}
		startpt += 1         /* move to next column */
		destpt += destlength /* move to next row */
	}
}

// Transpose an array "x". "scratch" must be the size
// of one full block (BL*BL), and "size" is the power of two x
// is on a side (e.g. 512x512 matrix has side=9). The loop
// below transposes block (1,1) and then moves down the first
// row and first column of blocks, switching their transposes.
// Then the process repeats for the second row and column, etc
func TransposeSquareFast(x, scratch []uint64, log_side uint8) ([]uint64, []uint64) {
	if log_side < BL {
		panic("Too small to transpose")
	}
	if len(scratch) < (1 << (2 * BL)) {
		panic("Scratch space too small")
	}
	corner := 0
	blocks := log_side - BL
	for col := 1; col <= 1<<blocks; col++ {
		copyblock(x, corner, 1<<log_side, scratch, 0, 1<<BL) /* diagonals */
		copytranspose(scratch, 0, x, corner, 1<<log_side)
		for row := 1; row < (1<<blocks)-col+1; row++ {
			copyblock(x, corner+(row<<BL), 1<<log_side, scratch, 0, 1<<BL) /* off-diagonals */
			copytranspose(scratch, 0, x, corner+(row<<(log_side+BL)), 1<<log_side)
			copyblock(scratch, 0, 1<<BL, x, corner+(row<<BL), 1<<log_side)
		}
		/* move down and to the right */
		corner += (1 << (log_side + BL)) + (1 << BL)
	}
	return x, scratch
}

// Transpose a square matrix
func TransposeSquareSlow(x, scratch []uint64, log_side uint8) ([]uint64, []uint64) {
	side := 1 << log_side
	for i := 1; i < side; i++ {
		a := i
		b := i << log_side
		for j := 0; j < i; j++ {
			x[b], x[a] = x[a], x[b]
			a += side
			b++
		}
	}
	return x, scratch
}

// Transpose a non square matrix
func TransposeSlow(x, scratch []uint64, log_cols uint8, log_rows uint8) ([]uint64, []uint64) {
	rows := 1 << log_rows
	cols := 1 << log_cols
	a := 0
	for i := 0; i < rows; i++ {
		b := i
		for j := 0; j < cols; j++ {
			scratch[a] = x[b]
			a++
			b += rows
		}
	}
	return scratch, x
}

// Transpose a matrix
//
// Pass in some scratch space which should be the same size as x
//
// It returns x, scratch which it may swap over
func Transpose(x, scratch []uint64, log_cols uint8, log_rows uint8) ([]uint64, []uint64) {
	if log_cols == log_rows {
		if log_cols < BL {
			return TransposeSquareSlow(x, scratch, log_cols)
		} else {
			return TransposeSquareFast(x, scratch, log_cols)
		}
	} else {
		return TransposeSlow(x, scratch, log_cols, log_rows)
	}
	panic("unreachable")
}
