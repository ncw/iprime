// Compute x = x - y mod p
// Preserves y, p
#define MOD_SUB(x, y, p, LABEL) \
	SUBQ    y, x; \
	JCC     LABEL; \
	ADDQ    p, x; \
LABEL:	\

// Compute x = x + y mod p, using t
// Preserves y, p
#define MOD_ADD(x, y, t, p, LABEL) \
        MOVQ    p, t; \
	SUBQ    y, t; \
        MOD_SUB(x, t, p, LABEL); \

TEXT ·mod_sub(SB),7,$0-24
	MOVQ    $-4294967295,R8
	MOVQ    x+0(FP),AX
	MOVQ    y+8(FP),CX
        MOD_SUB(AX, CX, R8, sub1)
	MOVQ    AX,ret+16(FP)
	RET

TEXT ·mod_add(SB),7,$0-24
	MOVQ    $-4294967295,R8
	MOVQ    x+0(FP),AX
	MOVQ    y+8(FP),BP
        MOD_ADD(AX, BP, CX, R8, add1)
	MOVQ    AX,ret+16(FP)
	RET

// Reduce 128 bits mod p (b, a) -> a
// Using t0, t1
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
	MOVQ    $-4294967295,R8
	MOVQ    b+0(FP),DI
	MOVQ    a+8(FP),AX

        MOD_REDUCE(DI, AX, SI, BX, R8, mod_reduce)

	MOVQ    AX,ret+16(FP)
	RET 

TEXT ·mod_mul(SB),7,$0-24
	MOVQ    x+0(FP),BX
	MOVQ    y+8(FP),AX
	MOVQ    $-4294967295,R8
        MULQ	BX /* BX * AX -> (DX, AX) */
        MOD_REDUCE(DX, AX, SI, BX, R8, mod_mul)
	MOVQ    AX,ret+16(FP)
	RET 

TEXT ·mod_sqr(SB),7,$0-16
	MOVQ    x+0(FP),AX
	MOVQ    $-4294967295,R8
        MULQ	AX /* AX * AX -> (DX, AX) */
        MOD_REDUCE(DX, AX, SI, BX, R8, mod_sqr)
	MOVQ    AX,ret+8(FP)
	RET 
