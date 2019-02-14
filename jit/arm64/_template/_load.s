	.file	"_load.s"
	.text


	.p2align 4,,15
	.globl	Load8
	.type	Load8, @function
Load8:
	.cfi_startproc
	ldrb	w0, [x0, #0]
	ldrb	w0, [x1, #0]
	ldrb	w0, [x2, #0]
	ldrb	w0, [x3, #0]
	ldrb	w0, [x28, #0]
	ldrb	w0, [x29, #0]
	ldrb	w0, [x30, #0]
	ldrb	w0, [sp, #0]
	nop
	ldrb	w0, [x0, #0]
	ldrb	w0, [x0, #1]
	ldrb	w0, [x0, #2]
	ldrb	w0, [x0, #8]
	ldrb	w0, [x0, #8]
	ldrb	w0, [x0, #16]
	ldrb	w0, [x0, #32]
	ldrb	w0, [x0, #64]
	ldrb	w0, [x0, #128]
	ldrb	w0, [x0, #256]
	ldrb	w0, [x0, #512]
	ldrb	w0, [x0, #1024]
	ldrb	w0, [x0, #2048]
	ldrb	w0, [x0, #4095]
	nop
	ldrb	w1, [x0, #0]
	ldrb	w1, [x0, #1]
	ldrb	w1, [x0, #2]
	ldrb	w1, [x0, #4]
	ldrb	w1, [x0, #4095]
	nop
	ldrb	w2, [x0, #0]
	ldrb	w2, [x0, #1]
	ldrb	w2, [x0, #2]
	ldrb	w2, [x0, #4]
	ldrb	w2, [x0, #4095]
	nop
	ldrb	w3, [x0, #0]
	ldrb	w3, [x0, #1]
	ldrb	w3, [x0, #2]
	ldrb	w3, [x0, #4]
	ldrb	w3, [x0, #4095]
	nop
	ldrb	w30, [x0, #0]
	ldrb	w30, [x0, #1]
	ldrb	w30, [x0, #2]
	ldrb	w30, [x0, #4]
	ldrb	w30, [x0, #4095]
	nop
	nop
	ldrb	w0, [x0, x0]
	ldrb	w0, [x0, x1]
	ldrb	w0, [x0, x2]
	ldrb	w0, [x0, x3]
	ldrb	w0, [x0, x30]
	nop
	ldrb	w0, [x0, x0]
	ldrb	w0, [x1, x0]
	ldrb	w0, [x2, x0]
	ldrb	w0, [x3, x0]
	ldrb	w0, [x30, x0]
	ldrb	w0, [sp, x0]
	nop
	ldrb	w0, [x0, x0]
	ldrb	w1, [x0, x0]
	ldrb	w2, [x0, x0]
	ldrb	w3, [x0, x0]
	ldrb	w30, [x0, x0]
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Load16
	.type	Load16, @function
Load16:
	.cfi_startproc
	ldrh	w0, [x0, #0]
	ldrh	w0, [x1, #0]
	ldrh	w0, [x2, #0]
	ldrh	w0, [x3, #0]
	ldrh	w0, [x28, #0]
	ldrh	w0, [x29, #0]
	ldrh	w0, [x30, #0]
	ldrh	w0, [sp, #0]
	nop
	ldrh	w0, [x0, #0]
	ldrh	w0, [x0, #2]
	ldrh	w0, [x0, #4]
	ldrh	w0, [x0, #8]
	ldrh	w0, [x0, #16]
	ldrh	w0, [x0, #32]
	ldrh	w0, [x0, #64]
	ldrh	w0, [x0, #128]
	ldrh	w0, [x0, #256]
	ldrh	w0, [x0, #512]
	ldrh	w0, [x0, #1024]
	ldrh	w0, [x0, #2048]
	ldrh	w0, [x0, #4096]
	ldrh	w0, [x0, #8190]
	nop
	ldrh	w1, [x0, #0]
	ldrh	w1, [x0, #2]
	ldrh	w1, [x0, #4]
	ldrh	w1, [x0, #8]
	ldrh	w1, [x0, #16]
	ldrh	w1, [x0, #8190]
	nop
	ldrh	w2, [x0, #0]
	ldrh	w2, [x0, #2]
	ldrh	w2, [x0, #4]
	ldrh	w2, [x0, #8]
	ldrh	w2, [x0, #16]
	ldrh	w2, [x0, #8190]
	nop
	ldrh	w3, [x0, #0]
	ldrh	w3, [x0, #2]
	ldrh	w3, [x0, #4]
	ldrh	w3, [x0, #8]
	ldrh	w3, [x0, #16]
	ldrh	w3, [x0, #8190]
	nop
	ldrh	w30, [x0, #0]
	ldrh	w30, [x0, #2]
	ldrh	w30, [x0, #4]
	ldrh	w30, [x0, #8]
	ldrh	w30, [x0, #16]
	ldrh	w30, [x0, #8190]
	nop
	nop
	ldrh	w0, [x0, x0]
	ldrh	w0, [x0, x1]
	ldrh	w0, [x0, x2]
	ldrh	w0, [x0, x3]
	ldrh	w0, [x0, x30]
	nop
	ldrh	w0, [x0, x0]
	ldrh	w0, [x1, x0]
	ldrh	w0, [x2, x0]
	ldrh	w0, [x3, x0]
	ldrh	w0, [x30, x0]
	ldrh	w0, [sp, x0]
	nop
	ldrh	w0, [x0, x0]
	ldrh	w1, [x0, x0]
	ldrh	w2, [x0, x0]
	ldrh	w3, [x0, x0]
	ldrh	w30, [x0, x0]
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Load32
	.type	Load32, @function
Load32:
	.cfi_startproc
	ldr	w0, [x0, #0]
	ldr	w0, [x1, #0]
	ldr	w0, [x2, #0]
	ldr	w0, [x3, #0]
	ldr	w0, [x28, #0]
	ldr	w0, [x29, #0]
	ldr	w0, [x30, #0]
	ldr w0, [sp, #0]
	nop
	ldr	w0, [x0, #0]
	ldr	w0, [x0, #4]
	ldr	w0, [x0, #8]
	ldr	w0, [x0, #16]
	ldr	w0, [x0, #32]
	ldr	w0, [x0, #64]
	ldr	w0, [x0, #128]
	ldr	w0, [x0, #256]
	ldr	w0, [x0, #512]
	ldr	w0, [x0, #1024]
	ldr	w0, [x0, #2048]
	ldr	w0, [x0, #4096]
	ldr	w0, [x0, #8192]
	ldr	w0, [x0, #16380]
	nop
	ldr	w1, [x0, #0]
	ldr	w1, [x0, #4]
	ldr	w1, [x0, #8]
	ldr	w1, [x0, #16]
	ldr	w1, [x0, #32]
	ldr	w1, [x0, #16380]
	nop
	ldr	w2, [x0, #0]
	ldr	w2, [x0, #8]
	ldr	w2, [x0, #16]
	ldr	w2, [x0, #32]
	ldr	w2, [x0, #16380]
	nop
	ldr	w3, [x0, #0]
	ldr	w3, [x0, #8]
	ldr	w3, [x0, #16]
	ldr	w3, [x0, #32]
	ldr	w3, [x0, #16380]
	nop
	ldr	w30, [x0, #0]
	ldr	w30, [x0, #8]
	ldr	w30, [x0, #16]
	ldr	w30, [x0, #32]
	ldr	w30, [x0, #16380]
	nop
	nop
	ldr	w0, [x0, x0]
	ldr	w0, [x0, x1]
	ldr	w0, [x0, x2]
	ldr	w0, [x0, x3]
	ldr	w0, [x0, x30]
	nop
	ldr	w0, [x0, x0]
	ldr	w0, [x1, x0]
	ldr	w0, [x2, x0]
	ldr	w0, [x3, x0]
	ldr	w0, [x30, x0]
	ldr	w0, [sp, x0]
	nop
	ldr	w0, [x0, x0]
	ldr	w1, [x0, x0]
	ldr	w2, [x0, x0]
	ldr	w3, [x0, x0]
	ldr	w30, [x0, x0]
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Load64
	.type	Load64, @function
Load64:
	.cfi_startproc
	ldr	x0, [x0, #0]
	ldr	x0, [x1, #0]
	ldr	x0, [x2, #0]
	ldr	x0, [x3, #0]
	ldr	x0, [x28, #0]
	ldr	x0, [x29, #0]
	ldr	x0, [x30, #0]
	ldr x0, [sp, #0]
	nop
	ldr	x0, [x0, #0]
	ldr	x0, [x0, #8]
	ldr	x0, [x0, #16]
	ldr	x0, [x0, #32]
	ldr	x0, [x0, #64]
	ldr	x0, [x0, #120]
	ldr	x0, [x0, #128]
	ldr	x0, [x0, #248]
	ldr	x0, [x0, #256]
	ldr	x0, [x0, #504]
	ldr	x0, [x0, #512]
	ldr	x0, [x0, #1016]
	ldr	x0, [x0, #1024]
	ldr	x0, [x0, #2040]
	ldr	x0, [x0, #2048]
	ldr	x0, [x0, #4088]
	ldr	x0, [x0, #4096]
	ldr	x0, [x0, #32752]
	ldr	x0, [x0, #32760]
	nop
	ldr	x1, [x0, #0]
	ldr	x1, [x0, #8]
	ldr	x1, [x0, #16]
	ldr	x1, [x0, #32]
	ldr	x1, [x0, #32760]
	nop
	ldr	x2, [x0, #0]
	ldr	x2, [x0, #8]
	ldr	x2, [x0, #16]
	ldr	x2, [x0, #32]
	ldr	x2, [x0, #32760]
	nop
	ldr	x3, [x0, #0]
	ldr	x3, [x0, #8]
	ldr	x3, [x0, #16]
	ldr	x3, [x0, #32]
	ldr	x3, [x0, #32760]
	nop
	ldr	x30, [x0, #0]
	ldr	x30, [x0, #8]
	ldr	x30, [x0, #16]
	ldr	x30, [x0, #32]
	ldr	x30, [x0, #32760]
	nop
	nop
	ldr	x0, [x0, x0]
	ldr	x0, [x0, x1]
	ldr	x0, [x0, x2]
	ldr	x0, [x0, x3]
	ldr	x0, [x0, x30]
	nop
	ldr	x0, [x0, x0]
	ldr	x0, [x1, x0]
	ldr	x0, [x2, x0]
	ldr	x0, [x3, x0]
	ldr	x0, [x30, x0]
	ldr	x0, [sp, x0]
	nop
	ldr	x0, [x0, x0]
	ldr	x1, [x0, x0]
	ldr	x2, [x0, x0]
	ldr	x3, [x0, x0]
	ldr	x30, [x0, x0]
	ret
	.cfi_endproc


