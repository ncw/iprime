//************************************************************
// Assembler routines for arithmetic mod p = 2^64-2^32+1
//
// The special form of p means that we can make a mod p routine which only
// involves a few shifts and arithmetic operations.  We use these facts
//
// 	2^64  == 2^32 -1 mod p
// 	2^96  == -1 mod p
// 	2^128 == -2^32 mod p
// 	2^192 == 1 mod p
//	2^n * 2^(192-n) = 1 mod p
//
// Also we use the fact that 2^64 - p is 2^32-1 ie 00000000FFFFFFFF so instead
// of adding FFFFFFFF00000001 mod 2^64 we subract 00000000FFFFFFFF.  This is
// convenient because the ARM carry flag is inverted for ADD and SUBtract so
// it is best (conditional execution wise) to follow an ADD with an ADDCS and
// a SUB with a SUBCC.  Read on and you will understand!
//
// Note that some of the comments use positional notation where (a,b,c) is
// 2^64 * a + 2^32 * b + c
// Note also that p = (2^32 -1, 1) in this notation
//
// -----
// 
// Factors of p-1 are 2^32 * 3 * 5 * 17 * 257 * 65537
// 
// Note that a 64-th root of unity is 8 mod p.
// 
// This means that practically an FFT can be defined of length up to 2^32 with
// optional factors of 3 and 5.
// 
// For the Discrete Weighted Transform we need an n-th root of 2.  2 has order
// 192 mod p (ie 2^192 mod p = 1) so we can have the
// 
// (p-1)/192 = (2^58 - 2^26) / 3 = 2^26 * 5 * 17 * 257 * 65537 th root of 2.
// 
// This means that we can do the DWT for lengths up to 2^26 with an optional
// factor of 5
// 
// 7 is a primitive root mod p
// 
// An n-th root of unity can be generated by 7^(5*(p-1)/n) mod p.
// 
// An n-th root of two can be generated by 7^(5*(p-1)/192/n) mod p
// 
// So a suitable 5 * 2^26-th root of 1 is 0xED41D05B78D6E286 and the 5 * 2^26-th
// root of 2 is &C47FC73D33F80E14
// 
// -----
//
// Many thanks to Peter-Lawrence Montgomery for working out the maths behind
// how to do the 128 bit to 64 bit reduction mod p and the shifts mod p so
// efficiently.  Peter also suggested the idea of using shifts in the
// transform which really makes a lot of difference in execution speed on ARM
//
// -----
//
// We could shave a few cycles off here and there by using redundant
// representation probably where the numbers are represented in the range
// 0..2^64-1.
//************************************************************

#define MOD_P 0xFFFFFFFF00000001

//The subtraction sequence for x - y (mod p) works whenever
// 
//   0 <= x <= 2^64 - 1
//   0 <= y <= 2^64 - 1
//   -p <= x - y <= p - 1

#define MOD_SUB(t0,t1, x0,x1, y0,y1, neg1) \
		SUB.S	y0, x0, t0		/* x - y */; \
		SBC.S	y1, x1, t1		/* CC if borrow */; \
		SUB.CC.S	neg1, t0, t0	/* Add back p if borrow by subratcing -p, set the flags */; \
		SBC.CC	$0, t1, t1		/* ripple the carry if CC */; \

#define MOD_NEGATE(t0,t1, x0,x1, neg1) \
		RSBS	$0, x0, t0		/* 0 - x */; \
		RSCS	$0, x1, t1		/* CC if borrow */; \
		SUBCCS	neg1, t0, t0		/* Add back p if borrow by subratcing -p, set the flags */; \
		SBCCC	$0, t1, t1		/* ripple the carry if CC */; \

// 4 words at 4..16(R13) for called routine parameters
TEXT ·mod_sub(SB),7,$0-24
	MOVW	$-1, R12
	MOVM.IB	(R13),[R0,R1,R2,R3]
        MOD_SUB(R4,R5, R0,R1, R2,R3, R12)
	MOVW	R4, ret+16(FP)
	MOVW	R5, ret+20(FP)
	RET

// Define addition as x - (p - y).  This seemingly backwards way of doing
// things works very well as it is much easier to make sure the result of a
// subtraction is in the range 0..p-1 than an addition (the conditional
// instructions for comparing > FFFFFFFF00000000 just don't work out
//
// The addition sequence for x + y (mod p) works whenever
// 
//   0 <= x <= 2^64 - 1
//   0 <= y <= p
//   x + y <= 2p - 1


#define MOD_ADD(t0,t1, x0,x1, y0,y1, neg1) \
		RSB.S	$1, y0, t0		/* do addition by negating y then subracting */; \
		RSC	neg1, y1, t1		/* This can't overflow if y < p */; \
		\
		/* what if y is 0 p-0 which is out of range */; \
		/* however subtracting y=p from x < p is guaranteed to carry which will fix it */; \
		/* x = 0, y = 0, -y = p, x - -y = 0 - p = -p => carry, add p == 0 ok */; \
		/* x = p-1, y = 0, -y = p, x - -y = p-1 - p = -1 => carry, add p == p-1 ok */; \
		\
		SUB.S	t0, x0, t0		/* x - (-y) */; \
		SBC.S	t1, x1, t1		/* CC if borrow */; \
		SUB.CC.S	neg1, t0, t0		/* Add back p if borrow by subratcing -p, set the flags */; \
		SBC.CC	$0, t1, t1		/* ripple the carry if CC */; \

// 4 words at 4..16(R13) for called routine parameters
TEXT ·mod_add(SB),7,$0-24
	MOVW	$-1, R12
	MOVM.IB  (R13),[R0,R1,R2,R3]
        MOD_ADD(R4,R5, R0,R1, R2,R3, R12)
	MOVW    R4,ret+16(FP)
	MOVW    R5,ret+20(FP)
	RET

//------------------------------------------------------------
// This reduces a 128 bit product mod p
// (x3,x2,x1,x0) mod p
// = (x2,x1,x0-x3)	[2^96 mod p = -1]
// = (x1+x2,x0-x3-x2)	[2^64 mod p = 2^32 -1]
// This ignores a whole pile of carries.
//
// temp1 = (x2, 2^32 - 1 - x3) - (0, x2)	/* In [0, p] */
// temp2 = temp1 + (x1, x0)
// if (carry clear) then
//   temp2 = temp2 - (0, 2^32 - 1)
//   if (borrow set) then
//     temp2 = temp2 - (0, 2^32 - 1)
//   end if
// end if
//
// If the temp2 = temp1 + (x1, x0) add has a carry (2^64), the intention
//   is to skip further processing, ending with
//   temp2 = (x1, x0) + (x2, 2^32 - 1 - x3 - x2) - 2^64
//         = (x1, x0) + (x2, - x3 - x2) - 2^64 + 2^32 - 1
//         = (x1, x0) + (x2, - x3 - x2) - p
//         = (x1, x0) + (x2, -x3 - x2) - 2^64 + 2^32 - 1.
//   This will be nonnegative (because it had a carry), and at most
//   (when x0 = x1 = x2 = 2^32-1 and x3 = 0)
//   (2^32-1, 2^32-1) + (2^32-1, -2^32+1) - p
//   = (2^33-2, 0) -p = 2^32*(2^33-2) - p = 2^65 - 2^33 - 2^64 + 2^32 - 1
//   = 2^64 - 2^32 - 1 = p - 2
// If there is no carry in the temp1 + (x1, x0) add
//   we subtract off the bias from the definition of temp1, ie
//   temp2' = temp2 - (2^32 - 1)
//          = (x1, x0) + (x2, -x3 - x2) + 2^32 - 1 - 2^32 + 1
//          = (x1, x0) + (x2, -x3 - x2)
//   Remember that temp2 is <= 2^64-1 since it did not carry therefore temp2' is at most
//   (2^64 - 1) - (2^32 - 1) = p - 1.
//   In the rare event where the difference is negative,
//   we subtract another 2^32 - 1 (effectively adding p mod 2^64)
//
// This works out very well in ARM code using the fact that carry has opposite
// senses for add and subract on the ARM (I am usually irritated by this not
// pleased!)
//
// Entry
//	(z3, z2, z1, z0)
// Exit
//	(z1, z0) = (z3, z2, z1, z0) mod p
//	z2,z3 destroyed
//------------------------------------------------------------


#define MOD_REDUCE(z0,z1,z2,z3, neg1, label) \
		SUB	z3, neg1, z3		/* z3 = 2^32 - 1 - z3 (can't carry) */; \
		SUB.S	z2, z3, z3		/* (z2,z3) = (z2, 2^32 - 1 - z3) - (0, z2) */; \
		SBC	$0, z2, z2		/* can't carry */; \
		\
		ADD.S	z3, z0, z0		/* (x1,x0) = (z1,z0)+(z2,z3) */; \
		ADC.S	z2, z1, z1		; \
		\
		/* CC = carry clear is borrow from subtract or no carry from add */; \
		\
label:		SUB.CC.S	neg1, z0, z0		/* if no carry (x1,x0) -= (0,2^32-1) */; \
		SBC.CC.S	$0, z1, z1		; \
		\
		BCC	label			/* Sub again if borrow - infrequent */; \

// 4 words at 4..16(R13) for called routine parameters
TEXT ·mod_reduce(SB),7,$0-24
	MOVW	$-1, R12
	MOVM.IB	(R13),[R0,R1,R2,R3]
        MOD_REDUCE(R0,R1,R2,R3, R12, mod_reduce)
	MOVW	R4, ret+16(FP)
	MOVW	R5, ret+20(FP)
	RET 

//------------------------------------------------------------
// Mod Multiply
//
// Note that this multiplies x0, y0 first so for best pipelining make sure x0
// and y0 aren't in limbo from a LDR when using this.  There are no pipelining
// timings to worry about on the way out of this macro
//
// x,y must be in range 0..p-1
// z will be in range 0..p-1
//
// Entry
//	(x1, x0)
//	(y1, y0)
//	x0 and y0 shouldn't be in the pipeline
// Exit
//	(z1, z0) = (x1, x0) * (y1, y0) mod p
//	y0 destroyed
//	z2, z3, t0 destroyed
//------------------------------------------------------------


// 64 x 64 -> 128 bit multiply - takes 22 cycles
// (x1,x0) * (y1,y0) -> (z3, z2, z1, z0)
#define	MOD_MUL(z0,z1, x0,x1, y0,y1, z2,z3,t0, neg1, label) \
		MULLU	y0, x0, (z0, z1)	/* odd order is to reduce pipeline stalls */; \
		MULLU	y0, x1, (y0, t0)	/* [not z0 or z1] */; \
		MULLU	y1, x1, (z2, z3)	/* [not y0 or t0] */; \
		ADD.S	y0, z1, z1		/* [not z2 or z3] */; \
		ADC.S	t0, z2, z2		; \
		MULLU	y1, x0, (y0, t0)	; \
		ADC	$0, z3, z3 		/* [not y0 or t0] */; \
		ADD.S	y0, z1, z1		; \
		ADC.S	t0, z2, z2		; \
		ADC	$0, z3, z3		; \
		MOD_REDUCE(z0, z1, z2, z3, neg1, label)	; \

// 4 words at 4..20(R13) for called routine parameters
TEXT ·mod_mul(SB),7,$0-24
	MOVW	$-1, R12
	MOVM.IB	(R13),[R0,R1,R2,R3]
        MOD_MUL(R4,R5, R0,R1, R2,R3, R7,R8,R11, R12, mod_mul)
	MOVW	R4, ret+16(FP)
	MOVW	R5, ret+20(FP)
	RET 


//------------------------------------------------------------
// Mod Square
//
// Note that this multiplies x0, x0 first so for best pipelining make sure x0
// isn't in limbo from a LDR when using this.  There are no pipelining
// timings to worry about on the way out of this macro
//
// x must be in range 0..p-1
// z will be in range 0..p-1
//
// Entry
//	(x1, x0)
//	x0 shouldn't be in the pipeline
// Exit
//	(z1, z0) = (x1, x0) ^ 2 mod p
//	x0 destroyed
//	z2, z3, t0 destroyed
//------------------------------------------------------------

// 64 x 64 -> 128 bit square - takes 18 cycles
// (x1,x0) * (x1,x0) -> (z3, z2, z1, z0)
#define MOD_SQR(z0,z1, x0,x1, z2,z3,t0, neg1, label) \
		MULLU	x0, x0, (z0, z1)	/* odd order is to reduce pipeline stalls */	; \
		MULLU	x0, x1, (x0, t0)	/* [not z0 or z1] */	; \
		MULLU	x1, x1, (z2, z3)	/* [not x0 or t0] */	; \
		ADD.S	x0, z1, z1	/* [not z2 or z3] */	; \
		ADC.S	t0, z2, z2	; \
		ADC	$0, z3, z3	; \
		ADD.S	x0, z1, z1	; \
		ADC.S	t0, z2, z2	; \
		ADC	$0, z3, z3	; \
		MOD_REDUCE(z0, z1, z2, z3, neg1, label); \

// 4 words at 4..20(R13) for called routine parameters
TEXT ·mod_sqr(SB),7,$0-24
	MOVW	$-1, R12
	MOVM.IB	(R13),[R0,R1,R2,R3]
        MOD_SQR(R4,R5, R0,R1, R2,R3,R6, R12, mod_sqr)
	MOVW	R4, ret+16(FP)
	MOVW	R5, ret+20(FP)
	RET 

//------------------------------------------------------------
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
//------------------------------------------------------------

#define MOD_SHIFT_0_TO_31(shift, t0, t1, x0, x1, neg1, rzero) \
		/* xhigh = x1, LSR #32 - shift */; \
		/* xmid = (x1, LSL #shift) OR (x0, LSR #32 - shift) */; \
		/* xlow  = x0, LSL #shift */; \
		/* calculate (t1,t0) = (2^32 - 1 - xhigh, xhigh + 1) */; \
		SUB.S	x1>>(32-shift), neg1, t1		/* 2^32 - 1 - xhigh, can't carry */; \
		 						/* trick alert!  no carry => CS, we use this to add 1 */; \
		ADC	x1>>(32 -shift), rzero, t0		/* xhigh + 1, can't carry (since xhigh < 2^31) */; \
		/* calculate (t1,t0) = (xhigh, xlow) - (2^32 - 1 - xhigh, xhigh + 1) */; \
		RSB.S	x0<<shift, t0, t0			/* xlow - t0 */; \
		MOVW	x1<<shift, x1	 			/* x1 = xmid */; \
		ORR	x0>>(32-shift), x1, x1			; \
		SBC.S	t1, x1, t1				/* xhigh - t1 - carry */; \
		SUB.CC.S	neg1, t0, t0			; \
		SBC.CC	$0, t1, t1				; \

#define MOD_SHIFT_0_TO_31_PROC(shift) \
	MOVW	$-1, R12	; \
	MOVW	$0, R14	; \
	MOVM.IB	(R13),[R0,R1]	; \
        MOD_SHIFT_0_TO_31(shift, R4, R5, R0, R1, R12, R14)	; \
	MOVW	R4, ret+8(FP)	; \
	MOVW	R5, ret+12(FP)	; \
	RET			; \

TEXT ·mod_shift3(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(3)
TEXT ·mod_shift6(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(6)
TEXT ·mod_shift9(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(9)
TEXT ·mod_shift12(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(12)
TEXT ·mod_shift15(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(15)
TEXT ·mod_shift18(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(18)
TEXT ·mod_shift21(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(21)
TEXT ·mod_shift24(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(24)
TEXT ·mod_shift27(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(27)
TEXT ·mod_shift30(SB),7,$0-16
	MOD_SHIFT_0_TO_31_PROC(30)

//------------------------------------------------------------
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
//------------------------------------------------------------

#define MOD_SHIFT_32_TO_63(shift, t0, t1, x0, x1, neg1) \
		/* xhigh = x1, LSR #64 - shift */		; \
		/* xmid = (x1, LSL #shift - 32) OR (x0, LSR #64 - shift) */	; \
		/* xlow  = x0, LSL #shift - 32 */		; \
		MOVW	x1<<(shift-32), t1			/* t1 = xmid */	; \
		ORR	x0>>(64-shift), t1, t1 			; \
		/* calculate (t1,t0) = (-xmid, xhigh + xmid) */	; \
		RSB.S	$0, t1, t1				/* xmidneg = 0 - xmid */	; \
		SBC	$0, t1, t0				/* xmidcomp = xmidneg - (borrow from last subtract) */	; \
		RSB.S	x1>>(64-shift), t0, t0			/* xhigh - xmidcomp */	; \
		SBC	$0, t1, t1				/* xmidneg - carry */	; \
		/* calculate (t1,t0) = (xlow, 0) - (t1,t0) mod p */	; \
		RSB.S	$0, t0, t0				/* 0 - t0 */	; \
		RSC.S	x0<<(shift-32), t1, t1			/* xlow - t1 - carry */	; \
		SUB.CC.S	neg1, t0, t0			; \
		SBC.CC	$0, t1, t1				; \

#define MOD_SHIFT_32_TO_63_PROC(shift) \
	MOVW	$-1, R12	; \
	MOVM.IB	(R13),[R0,R1]	; \
        MOD_SHIFT_32_TO_63(shift, R4, R5, R0, R1, R12)	; \
	MOVW	R4, ret+8(FP)	; \
	MOVW	R5, ret+12(FP)	; \
	RET			; \

TEXT ·mod_shift33(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(33)
TEXT ·mod_shift36(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(36)
TEXT ·mod_shift39(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(39)
TEXT ·mod_shift42(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(42)
TEXT ·mod_shift45(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(45)
TEXT ·mod_shift48(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(48)
TEXT ·mod_shift51(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(51)
TEXT ·mod_shift54(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(54)
TEXT ·mod_shift57(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(57)
TEXT ·mod_shift60(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(60)
TEXT ·mod_shift63(SB),7,$0-16
	MOD_SHIFT_32_TO_63_PROC(63)

//------------------------------------------------------------
// For shifts between 64..95 bits
// (xhigh, xmid, xlow, 0, 0) = (x1,x0) << shift where shift is 64..95
// (xhigh, xmid, xlow, 0, 0) mod p
// = (xmid, xlow, -xhigh, 0)
// = (xlow, -xhigh, -xmid)
// = (xlow - xhigh, -xmid - xlow)
// = (xlow, -xlow) - (xhigh, xmid)
// (xhigh, xmid) is < p since xhigh < 2^31
// (xlow, -xlow) can be evaluated as (xlow, 0) - (0, xlow) this is < p
//------------------------------------------------------------

#define MOD_SHIFT_64_TO_95(shift, t0, t1, x0, x1, neg1, rzero) \
		/* xhigh = x1, LSR #96 - shift */	; \
		/* xmid = (x1, LSL #shift - 64) OR (x0, LSR #96 - shift) */	; \
		/* xlow  = x0, LSL #shift - 64 */	; \
		/* calculate (xlow, 0) - (0, xlow) */	; \
		SUB.S	x0<<(shift-64), rzero, t0 	/* 0 - xlow */	; \
		RSC	x0<<(shift-64), rzero, t1	/* xlow - 0 - carry, no carry possible */	; \
		/* calculate (xlow, -xlow) - (xhigh, xmid) */	; \
		MOVW	x0>>(96-shift), x0		/* x0 = xmid */	; \
		ORR	x1<<(shift-64), x0, x0		; \
		SUB.S	x0, t0, t0			/* t0 - xmid */	; \
		SBC.S	x1>>(96-shift), t1, t1		/* t1 - xhigh */	; \
		SUB.CC.S	neg1, t0, t0		; \
		SBC.CC	$0, t1, t1			; \

#define MOD_SHIFT_64_TO_95_PROC(shift) \
	MOVW	$-1, R12	; \
	MOVW	$0, R14	; \
	MOVM.IB	(R13),[R0,R1]	; \
        MOD_SHIFT_64_TO_95(shift, R4, R5, R0, R1, R12, R14)	; \
	MOVW	R4, ret+8(FP)	; \
	MOVW	R5, ret+12(FP)	; \
	RET			; \

TEXT ·mod_shift66(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(66)
TEXT ·mod_shift69(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(69)
TEXT ·mod_shift72(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(72)
TEXT ·mod_shift75(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(75)
TEXT ·mod_shift78(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(78)
TEXT ·mod_shift81(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(81)
TEXT ·mod_shift84(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(84)
TEXT ·mod_shift87(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(87)
TEXT ·mod_shift90(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(90)
TEXT ·mod_shift93(SB),7,$0-16
	MOD_SHIFT_64_TO_95_PROC(93)

#include "ffts_arm.h"
