	.file	"_store.s"
	.text


	.p2align 4,,15
	.globl	Store8
	.type	Store8, @function
Store8:
	.cfi_startproc
	strb	w0, [x0, #0]
	strb	w0, [x1, #0]
	strb	w0, [x2, #0]
	strb	w0, [x3, #0]
	strb	w0, [x28, #0]
	strb	w0, [x29, #0]
	strb	w0, [x30, #0]
	strb	w0, [sp, #0]
	nop
	strb	w0, [x0, #0]
	strb	w0, [x0, #1]
	strb	w0, [x0, #2]
	strb	w0, [x0, #8]
	strb	w0, [x0, #8]
	strb	w0, [x0, #16]
	strb	w0, [x0, #32]
	strb	w0, [x0, #64]
	strb	w0, [x0, #128]
	strb	w0, [x0, #256]
	strb	w0, [x0, #512]
	strb	w0, [x0, #1024]
	strb	w0, [x0, #2048]
	strb	w0, [x0, #4095]
	nop
	strb	w1, [x0, #0]
	strb	w1, [x0, #1]
	strb	w1, [x0, #2]
	strb	w1, [x0, #4]
	strb	w1, [x0, #4095]
	nop
	strb	w2, [x0, #0]
	strb	w2, [x0, #1]
	strb	w2, [x0, #2]
	strb	w2, [x0, #4]
	strb	w2, [x0, #4095]
	nop
	strb	w3, [x0, #0]
	strb	w3, [x0, #1]
	strb	w3, [x0, #2]
	strb	w3, [x0, #4]
	strb	w3, [x0, #4095]
	nop
	strb	w30, [x0, #0]
	strb	w30, [x0, #1]
	strb	w30, [x0, #2]
	strb	w30, [x0, #4]
	strb	w30, [x0, #4095]
	nop
	nop
	strb	w0, [x0, x0]
	strb	w0, [x0, x1]
	strb	w0, [x0, x2]
	strb	w0, [x0, x3]
	strb	w0, [x0, x30]
	nop
	strb	w0, [x0, x0]
	strb	w0, [x1, x0]
	strb	w0, [x2, x0]
	strb	w0, [x3, x0]
	strb	w0, [x30, x0]
	strb	w0, [sp, x0]
	nop
	strb	w0, [x0, x0]
	strb	w1, [x0, x0]
	strb	w2, [x0, x0]
	strb	w3, [x0, x0]
	strb	w30, [x0, x0]
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Store16
	.type	Store16, @function
Store16:
	.cfi_startproc
	strh	w0, [x0, #0]
	strh	w0, [x1, #0]
	strh	w0, [x2, #0]
	strh	w0, [x3, #0]
	strh	w0, [x28, #0]
	strh	w0, [x29, #0]
	strh	w0, [x30, #0]
	strh	w0, [sp, #0]
	nop
	strh	w0, [x0, #0]
	strh	w0, [x0, #2]
	strh	w0, [x0, #4]
	strh	w0, [x0, #8]
	strh	w0, [x0, #16]
	strh	w0, [x0, #32]
	strh	w0, [x0, #64]
	strh	w0, [x0, #128]
	strh	w0, [x0, #256]
	strh	w0, [x0, #512]
	strh	w0, [x0, #1024]
	strh	w0, [x0, #2048]
	strh	w0, [x0, #4096]
	strh	w0, [x0, #8190]
	nop
	strh	w1, [x0, #0]
	strh	w1, [x0, #2]
	strh	w1, [x0, #4]
	strh	w1, [x0, #8]
	strh	w1, [x0, #16]
	strh	w1, [x0, #8190]
	nop
	strh	w2, [x0, #0]
	strh	w2, [x0, #2]
	strh	w2, [x0, #4]
	strh	w2, [x0, #8]
	strh	w2, [x0, #16]
	strh	w2, [x0, #8190]
	nop
	strh	w3, [x0, #0]
	strh	w3, [x0, #2]
	strh	w3, [x0, #4]
	strh	w3, [x0, #8]
	strh	w3, [x0, #16]
	strh	w3, [x0, #8190]
	nop
	strh	w30, [x0, #0]
	strh	w30, [x0, #2]
	strh	w30, [x0, #4]
	strh	w30, [x0, #8]
	strh	w30, [x0, #16]
	strh	w30, [x0, #8190]
	nop
	nop
	strh	w0, [x0, x0]
	strh	w0, [x0, x1]
	strh	w0, [x0, x2]
	strh	w0, [x0, x3]
	strh	w0, [x0, x30]
	nop
	strh	w0, [x0, x0]
	strh	w0, [x1, x0]
	strh	w0, [x2, x0]
	strh	w0, [x3, x0]
	strh	w0, [x30, x0]
	strh	w0, [sp, x0]
	nop
	strh	w0, [x0, x0]
	strh	w1, [x0, x0]
	strh	w2, [x0, x0]
	strh	w3, [x0, x0]
	strh	w30, [x0, x0]
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Store32
	.type	Store32, @function
Store32:
	.cfi_startproc
	str	w0, [x0, #0]
	str	w0, [x1, #0]
	str	w0, [x2, #0]
	str	w0, [x3, #0]
	str	w0, [x28, #0]
	str	w0, [x29, #0]
	str	w0, [x30, #0]
	str w0, [sp, #0]
	nop
	str	w0, [x0, #0]
	str	w0, [x0, #4]
	str	w0, [x0, #8]
	str	w0, [x0, #16]
	str	w0, [x0, #32]
	str	w0, [x0, #64]
	str	w0, [x0, #128]
	str	w0, [x0, #256]
	str	w0, [x0, #512]
	str	w0, [x0, #1024]
	str	w0, [x0, #2048]
	str	w0, [x0, #4096]
	str	w0, [x0, #8192]
	str	w0, [x0, #16380]
	nop
	str	w1, [x0, #0]
	str	w1, [x0, #4]
	str	w1, [x0, #8]
	str	w1, [x0, #16]
	str	w1, [x0, #32]
	str	w1, [x0, #16380]
	nop
	str	w2, [x0, #0]
	str	w2, [x0, #8]
	str	w2, [x0, #16]
	str	w2, [x0, #32]
	str	w2, [x0, #16380]
	nop
	str	w3, [x0, #0]
	str	w3, [x0, #8]
	str	w3, [x0, #16]
	str	w3, [x0, #32]
	str	w3, [x0, #16380]
	nop
	str	w30, [x0, #0]
	str	w30, [x0, #8]
	str	w30, [x0, #16]
	str	w30, [x0, #32]
	str	w30, [x0, #16380]
	nop
	nop
	str	w0, [x0, x0]
	str	w0, [x0, x1]
	str	w0, [x0, x2]
	str	w0, [x0, x3]
	str	w0, [x0, x30]
	nop
	str	w0, [x0, x0]
	str	w0, [x1, x0]
	str	w0, [x2, x0]
	str	w0, [x3, x0]
	str	w0, [x30, x0]
	str	w0, [sp, x0]
	nop
	str	w0, [x0, x0]
	str	w1, [x0, x0]
	str	w2, [x0, x0]
	str	w3, [x0, x0]
	str	w30, [x0, x0]
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Store64
	.type	Store64, @function
Store64:
	.cfi_startproc
	str	x0, [x0, #0]
	str	x0, [x1, #0]
	str	x0, [x2, #0]
	str	x0, [x3, #0]
	str	x0, [x28, #0]
	str	x0, [x29, #0]
	str	x0, [x30, #0]
	str x0, [sp, #0]
	nop
	str	x0, [x0, #0]
	str	x0, [x0, #8]
	str	x0, [x0, #16]
	str	x0, [x0, #32]
	str	x0, [x0, #64]
	str	x0, [x0, #120]
	str	x0, [x0, #128]
	str	x0, [x0, #248]
	str	x0, [x0, #256]
	str	x0, [x0, #504]
	str	x0, [x0, #512]
	str	x0, [x0, #1016]
	str	x0, [x0, #1024]
	str	x0, [x0, #2040]
	str	x0, [x0, #2048]
	str	x0, [x0, #4088]
	str	x0, [x0, #4096]
	str	x0, [x0, #32752]
	str	x0, [x0, #32760]
	nop
	str	x1, [x0, #0]
	str	x1, [x0, #8]
	str	x1, [x0, #16]
	str	x1, [x0, #32]
	str	x1, [x0, #32760]
	nop
	str	x2, [x0, #0]
	str	x2, [x0, #8]
	str	x2, [x0, #16]
	str	x2, [x0, #32]
	str	x2, [x0, #32760]
	nop
	str	x3, [x0, #0]
	str	x3, [x0, #8]
	str	x3, [x0, #16]
	str	x3, [x0, #32]
	str	x3, [x0, #32760]
	nop
	str	x30, [x0, #0]
	str	x30, [x0, #8]
	str	x30, [x0, #16]
	str	x30, [x0, #32]
	str	x30, [x0, #32760]
	nop
	nop
	str	x0, [x0, x0]
	str	x0, [x0, x1]
	str	x0, [x0, x2]
	str	x0, [x0, x3]
	str	x0, [x0, x30]
	nop
	str	x0, [x0, x0]
	str	x0, [x1, x0]
	str	x0, [x2, x0]
	str	x0, [x3, x0]
	str	x0, [x30, x0]
	str	x0, [sp, x0]
	nop
	str	x0, [x0, x0]
	str	x1, [x0, x0]
	str	x2, [x0, x0]
	str	x3, [x0, x0]
	str	x30, [x0, x0]
	ret
	.cfi_endproc


