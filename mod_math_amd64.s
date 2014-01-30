/*
Quicker way of loading MOD_P ?
        MOVL    $-1, R8
        NEGQ    R8
*/

#define MOD_P 0xFFFFFFFF00000001

// Compute x = x - y mod p
// Preserves y, p
#define MOD_SUB(x, y, p, LABEL) \
	SUBQ    y, x; \
	JCC     LABEL; \
	ADDQ    p, x; \
LABEL:	\

// Compute x = x - y mod p
// Preserves y, p
#define MOD_SUB2(x, y, p, LABEL) \
        XORQ R10,R10; \
	SUBQ    y, x; \
        CMOVQCS  p, R9; \
        CMOVQCC  R10, R9; \
        ADDQ    R9, x; \

// Compute x = x + y mod p, using t
// Preserves y, p
#define MOD_ADD(x, y, t, p, LABEL) \
        MOVQ    p, t; \
	SUBQ    y, t; \
        MOD_SUB(x, t, p, LABEL); \

TEXT ·mod_sub(SB),7,$0-24
	MOVQ    $MOD_P,R8
	MOVQ    x+0(FP),AX
	MOVQ    y+8(FP),CX
        MOD_SUB(AX, CX, R8, sub1)
	MOVQ    AX,ret+16(FP)
	RET

TEXT ·mod_add(SB),7,$0-24
	MOVQ    $MOD_P,R8
	MOVQ    x+0(FP),AX
	MOVQ    y+8(FP),BP
        MOD_ADD(AX, BP, CX, R8, add1)
	MOVQ    AX,ret+16(FP)
	RET

// Reduce 128 bits mod p (b, a) -> a
// Using t0, t1
// This is much faster (2 or 3 times) than DIVQ
#define MOD_REDUCE(b, a, t0, t1, p, label) \
        MOVL    b, t0;	/* Also sets upper 32 bits to 0 */ \
	SHRQ    $32,b ; \
 ; \
	CMPQ    a,p ; \
	JCS     label/**/1 ; \
	SUBQ    p,a ; \
label/**/1: ; \
 ; \
	MOVLQZX t0,t1 ; \
        MOD_SUB(a, t1, p, label/**/2) ; \
 ; \
	MOVLQZX b,t1 ; \
        MOD_SUB(a, t1, p, label/**/3) ; \
 ; \
	SHLQ    $32,t0 ; \
        MOD_ADD(a, t0, t1, p, label/**/4) ; \

TEXT ·mod_reduce(SB),7,$0-24
	MOVQ    $MOD_P,R8
	MOVQ    b+0(FP),DI
	MOVQ    a+8(FP),AX

        MOD_REDUCE(DI, AX, SI, BX, R8, mod_reduce)

	MOVQ    AX,ret+16(FP)
	RET 

TEXT ·mod_mul(SB),7,$0-24
	MOVQ    x+0(FP),BX
	MOVQ    y+8(FP),AX
	MOVQ    $MOD_P,R8
        MULQ	BX /* BX * AX -> (DX, AX) */
        MOD_REDUCE(DX, AX, SI, BX, R8, mod_mul)
	MOVQ    AX,ret+16(FP)
	RET 

TEXT ·mod_sqr(SB),7,$0-16
	MOVQ    x+0(FP),AX
	MOVQ    $MOD_P,R8
        MULQ	AX /* AX * AX -> (DX, AX) */
        MOD_REDUCE(DX, AX, SI, BX, R8, mod_sqr)
	MOVQ    AX,ret+8(FP)
	RET 

#define MOD_SHIFT_0_TO_31(x, shift, t0, t1, p, label) \
        MOVQ     x, t0; \
	/* xmid_xlow := x << shift, (xmid, xlow) */ \
        SHLQ    $shift, x; \
	/* xhigh := uint32(x >> (64 - shift)) */ \
        SHRQ    $(64-shift), t0 ; \
	/* t := uint64(0xFFFFFFFF-xhigh)<<32 + uint64(xhigh+1), (2^32 - 1 - xhigh, xhigh + 1) */ \
        MOVL     t0, t1; \
        SHLQ    $32, t0; \
        SUBQ     t1, t0; \
	/* r := xmid_xlow - t, (xmid, xlow) + (xhigh, -xhigh) */ \
        MOD_ADD(x, t0, t1, p, label); \

TEXT ·mod_shift0to31x3(SB),7,$0-16
	MOVQ    x+0(FP),AX
	MOVQ    $MOD_P,R8

        MOD_SHIFT_0_TO_31(AX, 3, BX, CX, R8, mod_shift)

	MOVQ    AX,ret+8(FP)
	RET 

#define MOD_SHIFT_32_TO_63(x, shift, t0, t1, t2, p, label) \
        MOVQ x, t0; \
	/* xmid := uint32(x >> (64 - shift)) */ \
	/* xlow := uint32(x << (shift - 32)) */ \
        SHLQ $(shift-32),x; \
	/* xhigh := uint32(x >> (96 - shift)) */ \
        SHRQ $(96-shift),t0; \
        MOVL x, t1; /* xlow */ \
        SHRQ $32, x; /* xmid */ \
        SHLQ $32, t1; \
        MOVL x, t2; /* t1 */ \
	/* t0 := uint64(xmid) << 32 // (xmid, 0) */ \
        SHLQ $32, x; /* t0 */ \
	/* t1 := uint64(xmid), (0, xmid) */ \
	/* t0 -= t1, (xmid, -xmid) no carry and must be in range 0..p-1 */ \
        SUBQ t2, x; \
	/* t1 = uint64(xhigh), (0, xhigh) */ \
	/* r := t0 - t1, (xmid, - xhigh - xmid) */ \
        MOD_SUB(x, t0, p, label/**/1); \
	/* add (xlow, 0) */ \
        MOD_ADD(x, t1, t0, p, label/**/2); \

        
TEXT ·mod_shift32to63x36(SB),7,$0-16
	MOVQ    x+0(FP),AX
	MOVQ    $MOD_P,R8

        MOD_SHIFT_32_TO_63(AX, 36, BX, CX, DX, R8, mod_shift36)

	MOVQ    AX,ret+8(FP)
	RET 

#define MOD_SHIFT_64_TO_95(x, shift, t0, t1, p, label) \
        MOVQ x, t0; \
	/* xlow := uint32(x << (shift - 64)) */ \
        SHLL $(shift - 64), x; \
	/* xhigh := uint32(x >> (128 - shift)) */ \
	/* xmid := uint32(x >> (96 - shift)) */ \
        SHRQ $(96 - shift), t0; \
	/* t0 := uint64(xlow) << 32, (xlow, 0) */ \
        MOVL x, t1; \
        SHLQ $32, x; \
	/* t1 := uint64(xlow), (0, xlow) */ \
	/* t0 -= t1, (xlow, -xlow) - no carry possible */ \
        SUBQ t1, x; \
	/* t1 = uint64(xhigh)<<32 + uint64(xmid), (xhigh, xmid) */ \
	/* r := t0 - t1, (xlow, -xlow) - (xhigh, xmid) */ \
        MOD_SUB(x, t0, p, label); \

TEXT ·mod_shift64to95x72(SB),7,$0-16
	MOVQ    x+0(FP),AX
	MOVQ    $MOD_P,R8

        MOD_SHIFT_64_TO_95(AX, 72, BX, CX, R8, mod_shift72)

	MOVQ    AX,ret+8(FP)
	RET 

#include "ffts_amd64.h"
