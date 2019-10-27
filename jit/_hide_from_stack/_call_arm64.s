	.file	"call_gcc.s"
	.text
	.p2align 4,,15
	.globl	call_0b
	.type	call_0b, @function
call_0b:
	.cfi_startproc
	ldr	x1, [sp], #8
	ldr	x0, [sp]
	str	x1, [sp]
	blr	x0
	ldr	x1, [sp]
	str	x1, [sp, #-8]!
	ret
	.cfi_endproc
