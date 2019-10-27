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

