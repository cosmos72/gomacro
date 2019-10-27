	.file	"call_gcc.s"
	.text
	.p2align 4,,15
	.globl	call_0b
	.type	call_0b, @function
call_0b:
	.cfi_startproc
	popq	%rdx
	movq	0x0(%rsp), %rax
	movq	%rdx, 0x0(%rsp)
	call	*%rax
	movq	0x0(%rsp), %rdx
	pushq	%rdx
	ret
	.cfi_endproc

	.text
	.p2align 4,,15
	.globl	asm_loop
	.type	asm_loop, @function
asm_loop:
	.cfi_startproc
	jmp		asm_loop
	ret
	.cfi_endproc

