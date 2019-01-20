	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Xor_s32
	.type	Xor_s32, @function
Xor_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// xor	$-0x55667788,%rax
	xor	$-0x55667788,%rcx
	xor	$-0x55667788,%rdx
	xor	$-0x55667788,%rbx
	xor	$-0x55667788,%rsp
	xor	$-0x55667788,%rbp
	xor	$-0x55667788,%rsi
	xor	$-0x55667788,%rdi
	xor	$-0x55667788,%r8
	xor	$-0x55667788,%r9
	xor	$-0x55667788,%r10
	xor	$-0x55667788,%r11
	xor	$-0x55667788,%r12
	xor	$-0x55667788,%r13
	xor	$-0x55667788,%r14
	xor	$-0x55667788,%r15
	.cfi_endproc


	.p2align 4,,15
	.globl	Xor
	.type	Xor, @function
Xor:
	.cfi_startproc
	xor	%rax,%rax
	xor	%rax,%rcx
	xor	%rax,%rdx
	xor	%rax,%rbx
	xor	%rax,%rsp
	xor	%rax,%rbp
	xor	%rax,%rsi
	xor	%rax,%rdi
	xor	%rax,%r8
	xor	%rax,%r9
	xor	%rax,%r10
	xor	%rax,%r11
	xor	%rax,%r12
	xor	%rax,%r13
	xor	%rax,%r14
	xor	%rax,%r15
	nop
	xor	%rcx,%rax
	xor	%rcx,%rcx
	xor	%rcx,%rdx
	xor	%rcx,%rbx
	xor	%rcx,%rsp
	xor	%rcx,%rbp
	xor	%rcx,%rsi
	xor	%rcx,%rdi
	xor	%rcx,%r8
	xor	%rcx,%r9
	xor	%rcx,%r10
	xor	%rcx,%r11
	xor	%rcx,%r12
	xor	%rcx,%r13
	xor	%rcx,%r14
	xor	%rcx,%r15
	nop
	xor	%rdx,%rax
	xor	%rdx,%rcx
	xor	%rdx,%rdx
	xor	%rdx,%rbx
	xor	%rdx,%rsp
	xor	%rdx,%rbp
	xor	%rdx,%rsi
	xor	%rdx,%rdi
	xor	%rdx,%r8
	xor	%rdx,%r9
	xor	%rdx,%r10
	xor	%rdx,%r11
	xor	%rdx,%r12
	xor	%rdx,%r13
	xor	%rdx,%r14
	xor	%rdx,%r15
	nop
	xor	%rbx,%rax
	xor	%rbx,%rcx
	xor	%rbx,%rdx
	xor	%rbx,%rbx
	xor	%rbx,%rsp
	xor	%rbx,%rbp
	xor	%rbx,%rsi
	xor	%rbx,%rdi
	xor	%rbx,%r8
	xor	%rbx,%r9
	xor	%rbx,%r10
	xor	%rbx,%r11
	xor	%rbx,%r12
	xor	%rbx,%r13
	xor	%rbx,%r14
	xor	%rbx,%r15
	nop
	xor	%rsp,%rax
	xor	%rsp,%rcx
	xor	%rsp,%rdx
	xor	%rsp,%rbx
	xor	%rsp,%rsp
	xor	%rsp,%rbp
	xor	%rsp,%rsi
	xor	%rsp,%rdi
	xor	%rsp,%r8
	xor	%rsp,%r9
	xor	%rsp,%r10
	xor	%rsp,%r11
	xor	%rsp,%r12
	xor	%rsp,%r13
	xor	%rsp,%r14
	xor	%rsp,%r15
	nop
	xor	%rbp,%rax
	xor	%rbp,%rcx
	xor	%rbp,%rdx
	xor	%rbp,%rbx
	xor	%rbp,%rsp
	xor	%rbp,%rbp
	xor	%rbp,%rsi
	xor	%rbp,%rdi
	xor	%rbp,%r8
	xor	%rbp,%r9
	xor	%rbp,%r10
	xor	%rbp,%r11
	xor	%rbp,%r12
	xor	%rbp,%r13
	xor	%rbp,%r14
	xor	%rbp,%r15
	nop
	xor	%rsi,%rax
	xor	%rsi,%rcx
	xor	%rsi,%rdx
	xor	%rsi,%rbx
	xor	%rsi,%rsp
	xor	%rsi,%rbp
	xor	%rsi,%rsi
	xor	%rsi,%rdi
	xor	%rsi,%r8
	xor	%rsi,%r9
	xor	%rsi,%r10
	xor	%rsi,%r11
	xor	%rsi,%r12
	xor	%rsi,%r13
	xor	%rsi,%r14
	xor	%rsi,%r15
	nop
	xor	%rdi,%rax
	xor	%rdi,%rcx
	xor	%rdi,%rdx
	xor	%rdi,%rbx
	xor	%rdi,%rsp
	xor	%rdi,%rbp
	xor	%rdi,%rsi
	xor	%rdi,%rdi
	xor	%rdi,%r8
	xor	%rdi,%r9
	xor	%rdi,%r10
	xor	%rdi,%r11
	xor	%rdi,%r12
	xor	%rdi,%r13
	xor	%rdi,%r14
	xor	%rdi,%r15
	nop
	xor	%r8, %rax
	xor	%r8, %rcx
	xor	%r8, %rdx
	xor	%r8, %rbx
	xor	%r8, %rsp
	xor	%r8, %rbp
	xor	%r8, %rsi
	xor	%r8, %rdi
	xor	%r8, %r8
	xor	%r8, %r9
	xor	%r8, %r10
	xor	%r8, %r11
	xor	%r8, %r12
	xor	%r8, %r13
	xor	%r8, %r14
	xor	%r8, %r15
	nop
	xor	%r9, %rax
	xor	%r9, %rcx
	xor	%r9, %rdx
	xor	%r9, %rbx
	xor	%r9, %rsp
	xor	%r9, %rbp
	xor	%r9, %rsi
	xor	%r9, %rdi
	xor	%r9, %r8
	xor	%r9, %r9
	xor	%r9, %r10
	xor	%r9, %r11
	xor	%r9, %r12
	xor	%r9, %r13
	xor	%r9, %r14
	xor	%r9, %r15
	nop
	xor	%r10,%rax
	xor	%r10,%rcx
	xor	%r10,%rdx
	xor	%r10,%rbx
	xor	%r10,%rsp
	xor	%r10,%rbp
	xor	%r10,%rsi
	xor	%r10,%rdi
	xor	%r10,%r8
	xor	%r10,%r9
	xor	%r10,%r10
	xor	%r10,%r11
	xor	%r10,%r12
	xor	%r10,%r13
	xor	%r10,%r14
	xor	%r10,%r15
	nop
	xor	%r11,%rax
	xor	%r11,%rcx
	xor	%r11,%rdx
	xor	%r11,%rbx
	xor	%r11,%rsp
	xor	%r11,%rbp
	xor	%r11,%rsi
	xor	%r11,%rdi
	xor	%r11,%r8
	xor	%r11,%r9
	xor	%r11,%r10
	xor	%r11,%r11
	xor	%r11,%r12
	xor	%r11,%r13
	xor	%r11,%r14
	xor	%r11,%r15
	nop
	xor	%r12,%rax
	xor	%r12,%rcx
	xor	%r12,%rdx
	xor	%r12,%rbx
	xor	%r12,%rsp
	xor	%r12,%rbp
	xor	%r12,%rsi
	xor	%r12,%rdi
	xor	%r12,%r8
	xor	%r12,%r9
	xor	%r12,%r10
	xor	%r12,%r11
	xor	%r12,%r12
	xor	%r12,%r13
	xor	%r12,%r14
	xor	%r12,%r15
	nop
	xor	%r13,%rax
	xor	%r13,%rcx
	xor	%r13,%rdx
	xor	%r13,%rbx
	xor	%r13,%rsp
	xor	%r13,%rbp
	xor	%r13,%rsi
	xor	%r13,%rdi
	xor	%r13,%r8
	xor	%r13,%r9
	xor	%r13,%r10
	xor	%r13,%r11
	xor	%r13,%r12
	xor	%r13,%r13
	xor	%r13,%r14
	xor	%r13,%r15
	nop
	xor	%r14,%rax
	xor	%r14,%rcx
	xor	%r14,%rdx
	xor	%r14,%rbx
	xor	%r14,%rsp
	xor	%r14,%rbp
	xor	%r14,%rsi
	xor	%r14,%rdi
	xor	%r14,%r8
	xor	%r14,%r9
	xor	%r14,%r10
	xor	%r14,%r11
	xor	%r14,%r12
	xor	%r14,%r13
	xor	%r14,%r14
	xor	%r14,%r15
	nop
	xor	%r15,%rax
	xor	%r15,%rcx
	xor	%r15,%rdx
	xor	%r15,%rbx
	xor	%r15,%rsp
	xor	%r15,%rbp
	xor	%r15,%rsi
	xor	%r15,%rdi
	xor	%r15,%r8
	xor	%r15,%r9
	xor	%r15,%r10
	xor	%r15,%r11
	xor	%r15,%r12
	xor	%r15,%r13
	xor	%r15,%r14
	xor	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XorMem
	.type	XorMem, @function
XorMem:
	.cfi_startproc
	xor	%rax,(%rax)
	xor	%rax,(%rcx)
	xor	%rax,(%rdx)
	xor	%rax,(%rbx)
	xor	%rax,(%rsp)
	xor	%rax,(%rbp)
	xor	%rax,(%rsi)
	xor	%rax,(%rdi)
	xor	%rax,(%r8)
	xor	%rax,(%r9)
	xor	%rax,(%r10)
	xor	%rax,(%r11)
	xor	%rax,(%r12)
	xor	%rax,(%r13)
	xor	%rax,(%r14)
	xor	%rax,(%r15)
	nop
	xor	%rcx,(%rax)
	xor	%rcx,(%rcx)
	xor	%rcx,(%rdx)
	xor	%rcx,(%rbx)
	xor	%rcx,(%rsp)
	xor	%rcx,(%rbp)
	xor	%rcx,(%rsi)
	xor	%rcx,(%rdi)
	xor	%rcx,(%r8)
	xor	%rcx,(%r9)
	xor	%rcx,(%r10)
	xor	%rcx,(%r11)
	xor	%rcx,(%r12)
	xor	%rcx,(%r13)
	xor	%rcx,(%r14)
	xor	%rcx,(%r15)
	nop
	xor	%rdx,(%rax)
	xor	%rdx,(%rcx)
	xor	%rdx,(%rdx)
	xor	%rdx,(%rbx)
	xor	%rdx,(%rsp)
	xor	%rdx,(%rbp)
	xor	%rdx,(%rsi)
	xor	%rdx,(%rdi)
	xor	%rdx,(%r8)
	xor	%rdx,(%r9)
	xor	%rdx,(%r10)
	xor	%rdx,(%r11)
	xor	%rdx,(%r12)
	xor	%rdx,(%r13)
	xor	%rdx,(%r14)
	xor	%rdx,(%r15)
	nop
	xor	%rbx,(%rax)
	xor	%rbx,(%rcx)
	xor	%rbx,(%rdx)
	xor	%rbx,(%rbx)
	xor	%rbx,(%rsp)
	xor	%rbx,(%rbp)
	xor	%rbx,(%rsi)
	xor	%rbx,(%rdi)
	xor	%rbx,(%r8)
	xor	%rbx,(%r9)
	xor	%rbx,(%r10)
	xor	%rbx,(%r11)
	xor	%rbx,(%r12)
	xor	%rbx,(%r13)
	xor	%rbx,(%r14)
	xor	%rbx,(%r15)
	nop
	xor	%rsp,(%rax)
	xor	%rsp,(%rcx)
	xor	%rsp,(%rdx)
	xor	%rsp,(%rbx)
	xor	%rsp,(%rsp)
	xor	%rsp,(%rbp)
	xor	%rsp,(%rsi)
	xor	%rsp,(%rdi)
	xor	%rsp,(%r8)
	xor	%rsp,(%r9)
	xor	%rsp,(%r10)
	xor	%rsp,(%r11)
	xor	%rsp,(%r12)
	xor	%rsp,(%r13)
	xor	%rsp,(%r14)
	xor	%rsp,(%r15)
	nop
	xor	%rbp,(%rax)
	xor	%rbp,(%rcx)
	xor	%rbp,(%rdx)
	xor	%rbp,(%rbx)
	xor	%rbp,(%rsp)
	xor	%rbp,(%rbp)
	xor	%rbp,(%rsi)
	xor	%rbp,(%rdi)
	xor	%rbp,(%r8)
	xor	%rbp,(%r9)
	xor	%rbp,(%r10)
	xor	%rbp,(%r11)
	xor	%rbp,(%r12)
	xor	%rbp,(%r13)
	xor	%rbp,(%r14)
	xor	%rbp,(%r15)
	nop
	xor	%rsi,(%rax)
	xor	%rsi,(%rcx)
	xor	%rsi,(%rdx)
	xor	%rsi,(%rbx)
	xor	%rsi,(%rsp)
	xor	%rsi,(%rbp)
	xor	%rsi,(%rsi)
	xor	%rsi,(%rdi)
	xor	%rsi,(%r8)
	xor	%rsi,(%r9)
	xor	%rsi,(%r10)
	xor	%rsi,(%r11)
	xor	%rsi,(%r12)
	xor	%rsi,(%r13)
	xor	%rsi,(%r14)
	xor	%rsi,(%r15)
	nop
	xor	%rdi,(%rax)
	xor	%rdi,(%rcx)
	xor	%rdi,(%rdx)
	xor	%rdi,(%rbx)
	xor	%rdi,(%rsp)
	xor	%rdi,(%rbp)
	xor	%rdi,(%rsi)
	xor	%rdi,(%rdi)
	xor	%rdi,(%r8)
	xor	%rdi,(%r9)
	xor	%rdi,(%r10)
	xor	%rdi,(%r11)
	xor	%rdi,(%r12)
	xor	%rdi,(%r13)
	xor	%rdi,(%r14)
	xor	%rdi,(%r15)
	nop
	xor	%r8, (%rax)
	xor	%r8, (%rcx)
	xor	%r8, (%rdx)
	xor	%r8, (%rbx)
	xor	%r8, (%rsp)
	xor	%r8, (%rbp)
	xor	%r8, (%rsi)
	xor	%r8, (%rdi)
	xor	%r8, (%r8)
	xor	%r8, (%r9)
	xor	%r8, (%r10)
	xor	%r8, (%r11)
	xor	%r8, (%r12)
	xor	%r8, (%r13)
	xor	%r8, (%r14)
	xor	%r8, (%r15)
	nop
	xor	%r9, (%rax)
	xor	%r9, (%rcx)
	xor	%r9, (%rdx)
	xor	%r9, (%rbx)
	xor	%r9, (%rsp)
	xor	%r9, (%rbp)
	xor	%r9, (%rsi)
	xor	%r9, (%rdi)
	xor	%r9, (%r8)
	xor	%r9, (%r9)
	xor	%r9, (%r10)
	xor	%r9, (%r11)
	xor	%r9, (%r12)
	xor	%r9, (%r13)
	xor	%r9, (%r14)
	xor	%r9, (%r15)
	nop
	xor	%r10,(%rax)
	xor	%r10,(%rcx)
	xor	%r10,(%rdx)
	xor	%r10,(%rbx)
	xor	%r10,(%rsp)
	xor	%r10,(%rbp)
	xor	%r10,(%rsi)
	xor	%r10,(%rdi)
	xor	%r10,(%r8)
	xor	%r10,(%r9)
	xor	%r10,(%r10)
	xor	%r10,(%r11)
	xor	%r10,(%r12)
	xor	%r10,(%r13)
	xor	%r10,(%r14)
	xor	%r10,(%r15)
	nop
	xor	%r11,(%rax)
	xor	%r11,(%rcx)
	xor	%r11,(%rdx)
	xor	%r11,(%rbx)
	xor	%r11,(%rsp)
	xor	%r11,(%rbp)
	xor	%r11,(%rsi)
	xor	%r11,(%rdi)
	xor	%r11,(%r8)
	xor	%r11,(%r9)
	xor	%r11,(%r10)
	xor	%r11,(%r11)
	xor	%r11,(%r12)
	xor	%r11,(%r13)
	xor	%r11,(%r14)
	xor	%r11,(%r15)
	nop
	xor	%r12,(%rax)
	xor	%r12,(%rcx)
	xor	%r12,(%rdx)
	xor	%r12,(%rbx)
	xor	%r12,(%rsp)
	xor	%r12,(%rbp)
	xor	%r12,(%rsi)
	xor	%r12,(%rdi)
	xor	%r12,(%r8)
	xor	%r12,(%r9)
	xor	%r12,(%r10)
	xor	%r12,(%r11)
	xor	%r12,(%r12)
	xor	%r12,(%r13)
	xor	%r12,(%r14)
	xor	%r12,(%r15)
	nop
	xor	%r13,(%rax)
	xor	%r13,(%rcx)
	xor	%r13,(%rdx)
	xor	%r13,(%rbx)
	xor	%r13,(%rsp)
	xor	%r13,(%rbp)
	xor	%r13,(%rsi)
	xor	%r13,(%rdi)
	xor	%r13,(%r8)
	xor	%r13,(%r9)
	xor	%r13,(%r10)
	xor	%r13,(%r11)
	xor	%r13,(%r12)
	xor	%r13,(%r13)
	xor	%r13,(%r14)
	xor	%r13,(%r15)
	nop
	xor	%r14,(%rax)
	xor	%r14,(%rcx)
	xor	%r14,(%rdx)
	xor	%r14,(%rbx)
	xor	%r14,(%rsp)
	xor	%r14,(%rbp)
	xor	%r14,(%rsi)
	xor	%r14,(%rdi)
	xor	%r14,(%r8)
	xor	%r14,(%r9)
	xor	%r14,(%r10)
	xor	%r14,(%r11)
	xor	%r14,(%r12)
	xor	%r14,(%r13)
	xor	%r14,(%r14)
	xor	%r14,(%r15)
	nop
	xor	%r15,(%rax)
	xor	%r15,(%rcx)
	xor	%r15,(%rdx)
	xor	%r15,(%rbx)
	xor	%r15,(%rsp)
	xor	%r15,(%rbp)
	xor	%r15,(%rsi)
	xor	%r15,(%rdi)
	xor	%r15,(%r8)
	xor	%r15,(%r9)
	xor	%r15,(%r10)
	xor	%r15,(%r11)
	xor	%r15,(%r12)
	xor	%r15,(%r13)
	xor	%r15,(%r14)
	xor	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XorMem8
	.type	XorMem8, @function
XorMem8:
	.cfi_startproc
	xor	%rax,0x7f(%rax)
	xor	%rax,0x7f(%rcx)
	xor	%rax,0x7f(%rdx)
	xor	%rax,0x7f(%rbx)
	xor	%rax,0x7f(%rsp)
	xor	%rax,0x7f(%rbp)
	xor	%rax,0x7f(%rsi)
	xor	%rax,0x7f(%rdi)
	xor	%rax,0x7f(%r8)
	xor	%rax,0x7f(%r9)
	xor	%rax,0x7f(%r10)
	xor	%rax,0x7f(%r11)
	xor	%rax,0x7f(%r12)
	xor	%rax,0x7f(%r13)
	xor	%rax,0x7f(%r14)
	xor	%rax,0x7f(%r15)
	nop
	xor	%rcx,0x7f(%rax)
	xor	%rcx,0x7f(%rcx)
	xor	%rcx,0x7f(%rdx)
	xor	%rcx,0x7f(%rbx)
	xor	%rcx,0x7f(%rsp)
	xor	%rcx,0x7f(%rbp)
	xor	%rcx,0x7f(%rsi)
	xor	%rcx,0x7f(%rdi)
	xor	%rcx,0x7f(%r8)
	xor	%rcx,0x7f(%r9)
	xor	%rcx,0x7f(%r10)
	xor	%rcx,0x7f(%r11)
	xor	%rcx,0x7f(%r12)
	xor	%rcx,0x7f(%r13)
	xor	%rcx,0x7f(%r14)
	xor	%rcx,0x7f(%r15)
	nop
	xor	%rdx,0x7f(%rax)
	xor	%rdx,0x7f(%rcx)
	xor	%rdx,0x7f(%rdx)
	xor	%rdx,0x7f(%rbx)
	xor	%rdx,0x7f(%rsp)
	xor	%rdx,0x7f(%rbp)
	xor	%rdx,0x7f(%rsi)
	xor	%rdx,0x7f(%rdi)
	xor	%rdx,0x7f(%r8)
	xor	%rdx,0x7f(%r9)
	xor	%rdx,0x7f(%r10)
	xor	%rdx,0x7f(%r11)
	xor	%rdx,0x7f(%r12)
	xor	%rdx,0x7f(%r13)
	xor	%rdx,0x7f(%r14)
	xor	%rdx,0x7f(%r15)
	nop
	xor	%rbx,0x7f(%rax)
	xor	%rbx,0x7f(%rcx)
	xor	%rbx,0x7f(%rdx)
	xor	%rbx,0x7f(%rbx)
	xor	%rbx,0x7f(%rsp)
	xor	%rbx,0x7f(%rbp)
	xor	%rbx,0x7f(%rsi)
	xor	%rbx,0x7f(%rdi)
	xor	%rbx,0x7f(%r8)
	xor	%rbx,0x7f(%r9)
	xor	%rbx,0x7f(%r10)
	xor	%rbx,0x7f(%r11)
	xor	%rbx,0x7f(%r12)
	xor	%rbx,0x7f(%r13)
	xor	%rbx,0x7f(%r14)
	xor	%rbx,0x7f(%r15)
	nop
	xor	%rsp,0x7f(%rax)
	xor	%rsp,0x7f(%rcx)
	xor	%rsp,0x7f(%rdx)
	xor	%rsp,0x7f(%rbx)
	xor	%rsp,0x7f(%rsp)
	xor	%rsp,0x7f(%rbp)
	xor	%rsp,0x7f(%rsi)
	xor	%rsp,0x7f(%rdi)
	xor	%rsp,0x7f(%r8)
	xor	%rsp,0x7f(%r9)
	xor	%rsp,0x7f(%r10)
	xor	%rsp,0x7f(%r11)
	xor	%rsp,0x7f(%r12)
	xor	%rsp,0x7f(%r13)
	xor	%rsp,0x7f(%r14)
	xor	%rsp,0x7f(%r15)
	nop
	xor	%rbp,0x7f(%rax)
	xor	%rbp,0x7f(%rcx)
	xor	%rbp,0x7f(%rdx)
	xor	%rbp,0x7f(%rbx)
	xor	%rbp,0x7f(%rsp)
	xor	%rbp,0x7f(%rbp)
	xor	%rbp,0x7f(%rsi)
	xor	%rbp,0x7f(%rdi)
	xor	%rbp,0x7f(%r8)
	xor	%rbp,0x7f(%r9)
	xor	%rbp,0x7f(%r10)
	xor	%rbp,0x7f(%r11)
	xor	%rbp,0x7f(%r12)
	xor	%rbp,0x7f(%r13)
	xor	%rbp,0x7f(%r14)
	xor	%rbp,0x7f(%r15)
	nop
	xor	%rsi,0x7f(%rax)
	xor	%rsi,0x7f(%rcx)
	xor	%rsi,0x7f(%rdx)
	xor	%rsi,0x7f(%rbx)
	xor	%rsi,0x7f(%rsp)
	xor	%rsi,0x7f(%rbp)
	xor	%rsi,0x7f(%rsi)
	xor	%rsi,0x7f(%rdi)
	xor	%rsi,0x7f(%r8)
	xor	%rsi,0x7f(%r9)
	xor	%rsi,0x7f(%r10)
	xor	%rsi,0x7f(%r11)
	xor	%rsi,0x7f(%r12)
	xor	%rsi,0x7f(%r13)
	xor	%rsi,0x7f(%r14)
	xor	%rsi,0x7f(%r15)
	nop
	xor	%rdi,0x7f(%rax)
	xor	%rdi,0x7f(%rcx)
	xor	%rdi,0x7f(%rdx)
	xor	%rdi,0x7f(%rbx)
	xor	%rdi,0x7f(%rsp)
	xor	%rdi,0x7f(%rbp)
	xor	%rdi,0x7f(%rsi)
	xor	%rdi,0x7f(%rdi)
	xor	%rdi,0x7f(%r8)
	xor	%rdi,0x7f(%r9)
	xor	%rdi,0x7f(%r10)
	xor	%rdi,0x7f(%r11)
	xor	%rdi,0x7f(%r12)
	xor	%rdi,0x7f(%r13)
	xor	%rdi,0x7f(%r14)
	xor	%rdi,0x7f(%r15)
	nop
	xor	%r8, 0x7f(%rax)
	xor	%r8, 0x7f(%rcx)
	xor	%r8, 0x7f(%rdx)
	xor	%r8, 0x7f(%rbx)
	xor	%r8, 0x7f(%rsp)
	xor	%r8, 0x7f(%rbp)
	xor	%r8, 0x7f(%rsi)
	xor	%r8, 0x7f(%rdi)
	xor	%r8, 0x7f(%r8)
	xor	%r8, 0x7f(%r9)
	xor	%r8, 0x7f(%r10)
	xor	%r8, 0x7f(%r11)
	xor	%r8, 0x7f(%r12)
	xor	%r8, 0x7f(%r13)
	xor	%r8, 0x7f(%r14)
	xor	%r8, 0x7f(%r15)
	nop
	xor	%r9, 0x7f(%rax)
	xor	%r9, 0x7f(%rcx)
	xor	%r9, 0x7f(%rdx)
	xor	%r9, 0x7f(%rbx)
	xor	%r9, 0x7f(%rsp)
	xor	%r9, 0x7f(%rbp)
	xor	%r9, 0x7f(%rsi)
	xor	%r9, 0x7f(%rdi)
	xor	%r9, 0x7f(%r8)
	xor	%r9, 0x7f(%r9)
	xor	%r9, 0x7f(%r10)
	xor	%r9, 0x7f(%r11)
	xor	%r9, 0x7f(%r12)
	xor	%r9, 0x7f(%r13)
	xor	%r9, 0x7f(%r14)
	xor	%r9, 0x7f(%r15)
	nop
	xor	%r10,0x7f(%rax)
	xor	%r10,0x7f(%rcx)
	xor	%r10,0x7f(%rdx)
	xor	%r10,0x7f(%rbx)
	xor	%r10,0x7f(%rsp)
	xor	%r10,0x7f(%rbp)
	xor	%r10,0x7f(%rsi)
	xor	%r10,0x7f(%rdi)
	xor	%r10,0x7f(%r8)
	xor	%r10,0x7f(%r9)
	xor	%r10,0x7f(%r10)
	xor	%r10,0x7f(%r11)
	xor	%r10,0x7f(%r12)
	xor	%r10,0x7f(%r13)
	xor	%r10,0x7f(%r14)
	xor	%r10,0x7f(%r15)
	nop
	xor	%r11,0x7f(%rax)
	xor	%r11,0x7f(%rcx)
	xor	%r11,0x7f(%rdx)
	xor	%r11,0x7f(%rbx)
	xor	%r11,0x7f(%rsp)
	xor	%r11,0x7f(%rbp)
	xor	%r11,0x7f(%rsi)
	xor	%r11,0x7f(%rdi)
	xor	%r11,0x7f(%r8)
	xor	%r11,0x7f(%r9)
	xor	%r11,0x7f(%r10)
	xor	%r11,0x7f(%r11)
	xor	%r11,0x7f(%r12)
	xor	%r11,0x7f(%r13)
	xor	%r11,0x7f(%r14)
	xor	%r11,0x7f(%r15)
	nop
	xor	%r12,0x7f(%rax)
	xor	%r12,0x7f(%rcx)
	xor	%r12,0x7f(%rdx)
	xor	%r12,0x7f(%rbx)
	xor	%r12,0x7f(%rsp)
	xor	%r12,0x7f(%rbp)
	xor	%r12,0x7f(%rsi)
	xor	%r12,0x7f(%rdi)
	xor	%r12,0x7f(%r8)
	xor	%r12,0x7f(%r9)
	xor	%r12,0x7f(%r10)
	xor	%r12,0x7f(%r11)
	xor	%r12,0x7f(%r12)
	xor	%r12,0x7f(%r13)
	xor	%r12,0x7f(%r14)
	xor	%r12,0x7f(%r15)
	nop
	xor	%r13,0x7f(%rax)
	xor	%r13,0x7f(%rcx)
	xor	%r13,0x7f(%rdx)
	xor	%r13,0x7f(%rbx)
	xor	%r13,0x7f(%rsp)
	xor	%r13,0x7f(%rbp)
	xor	%r13,0x7f(%rsi)
	xor	%r13,0x7f(%rdi)
	xor	%r13,0x7f(%r8)
	xor	%r13,0x7f(%r9)
	xor	%r13,0x7f(%r10)
	xor	%r13,0x7f(%r11)
	xor	%r13,0x7f(%r12)
	xor	%r13,0x7f(%r13)
	xor	%r13,0x7f(%r14)
	xor	%r13,0x7f(%r15)
	nop
	xor	%r14,0x7f(%rax)
	xor	%r14,0x7f(%rcx)
	xor	%r14,0x7f(%rdx)
	xor	%r14,0x7f(%rbx)
	xor	%r14,0x7f(%rsp)
	xor	%r14,0x7f(%rbp)
	xor	%r14,0x7f(%rsi)
	xor	%r14,0x7f(%rdi)
	xor	%r14,0x7f(%r8)
	xor	%r14,0x7f(%r9)
	xor	%r14,0x7f(%r10)
	xor	%r14,0x7f(%r11)
	xor	%r14,0x7f(%r12)
	xor	%r14,0x7f(%r13)
	xor	%r14,0x7f(%r14)
	xor	%r14,0x7f(%r15)
	nop
	xor	%r15,0x7f(%rax)
	xor	%r15,0x7f(%rcx)
	xor	%r15,0x7f(%rdx)
	xor	%r15,0x7f(%rbx)
	xor	%r15,0x7f(%rsp)
	xor	%r15,0x7f(%rbp)
	xor	%r15,0x7f(%rsi)
	xor	%r15,0x7f(%rdi)
	xor	%r15,0x7f(%r8)
	xor	%r15,0x7f(%r9)
	xor	%r15,0x7f(%r10)
	xor	%r15,0x7f(%r11)
	xor	%r15,0x7f(%r12)
	xor	%r15,0x7f(%r13)
	xor	%r15,0x7f(%r14)
	xor	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XorMem32
	.type	XorMem32, @function
XorMem32:
	.cfi_startproc
	xor	%rax,0x7f563412(%rax)
	xor	%rax,0x7f563412(%rcx)
	xor	%rax,0x7f563412(%rdx)
	xor	%rax,0x7f563412(%rbx)
	xor	%rax,0x7f563412(%rsp)
	xor	%rax,0x7f563412(%rbp)
	xor	%rax,0x7f563412(%rsi)
	xor	%rax,0x7f563412(%rdi)
	xor	%rax,0x7f563412(%r8)
	xor	%rax,0x7f563412(%r9)
	xor	%rax,0x7f563412(%r10)
	xor	%rax,0x7f563412(%r11)
	xor	%rax,0x7f563412(%r12)
	xor	%rax,0x7f563412(%r13)
	xor	%rax,0x7f563412(%r14)
	xor	%rax,0x7f563412(%r15)
	nop
	xor	%rcx,0x7f563412(%rax)
	xor	%rcx,0x7f563412(%rcx)
	xor	%rcx,0x7f563412(%rdx)
	xor	%rcx,0x7f563412(%rbx)
	xor	%rcx,0x7f563412(%rsp)
	xor	%rcx,0x7f563412(%rbp)
	xor	%rcx,0x7f563412(%rsi)
	xor	%rcx,0x7f563412(%rdi)
	xor	%rcx,0x7f563412(%r8)
	xor	%rcx,0x7f563412(%r9)
	xor	%rcx,0x7f563412(%r10)
	xor	%rcx,0x7f563412(%r11)
	xor	%rcx,0x7f563412(%r12)
	xor	%rcx,0x7f563412(%r13)
	xor	%rcx,0x7f563412(%r14)
	xor	%rcx,0x7f563412(%r15)
	nop
	xor	%rdx,0x7f563412(%rax)
	xor	%rdx,0x7f563412(%rcx)
	xor	%rdx,0x7f563412(%rdx)
	xor	%rdx,0x7f563412(%rbx)
	xor	%rdx,0x7f563412(%rsp)
	xor	%rdx,0x7f563412(%rbp)
	xor	%rdx,0x7f563412(%rsi)
	xor	%rdx,0x7f563412(%rdi)
	xor	%rdx,0x7f563412(%r8)
	xor	%rdx,0x7f563412(%r9)
	xor	%rdx,0x7f563412(%r10)
	xor	%rdx,0x7f563412(%r11)
	xor	%rdx,0x7f563412(%r12)
	xor	%rdx,0x7f563412(%r13)
	xor	%rdx,0x7f563412(%r14)
	xor	%rdx,0x7f563412(%r15)
	nop
	xor	%rbx,0x7f563412(%rax)
	xor	%rbx,0x7f563412(%rcx)
	xor	%rbx,0x7f563412(%rdx)
	xor	%rbx,0x7f563412(%rbx)
	xor	%rbx,0x7f563412(%rsp)
	xor	%rbx,0x7f563412(%rbp)
	xor	%rbx,0x7f563412(%rsi)
	xor	%rbx,0x7f563412(%rdi)
	xor	%rbx,0x7f563412(%r8)
	xor	%rbx,0x7f563412(%r9)
	xor	%rbx,0x7f563412(%r10)
	xor	%rbx,0x7f563412(%r11)
	xor	%rbx,0x7f563412(%r12)
	xor	%rbx,0x7f563412(%r13)
	xor	%rbx,0x7f563412(%r14)
	xor	%rbx,0x7f563412(%r15)
	nop
	xor	%rsp,0x7f563412(%rax)
	xor	%rsp,0x7f563412(%rcx)
	xor	%rsp,0x7f563412(%rdx)
	xor	%rsp,0x7f563412(%rbx)
	xor	%rsp,0x7f563412(%rsp)
	xor	%rsp,0x7f563412(%rbp)
	xor	%rsp,0x7f563412(%rsi)
	xor	%rsp,0x7f563412(%rdi)
	xor	%rsp,0x7f563412(%r8)
	xor	%rsp,0x7f563412(%r9)
	xor	%rsp,0x7f563412(%r10)
	xor	%rsp,0x7f563412(%r11)
	xor	%rsp,0x7f563412(%r12)
	xor	%rsp,0x7f563412(%r13)
	xor	%rsp,0x7f563412(%r14)
	xor	%rsp,0x7f563412(%r15)
	nop
	xor	%rbp,0x7f563412(%rax)
	xor	%rbp,0x7f563412(%rcx)
	xor	%rbp,0x7f563412(%rdx)
	xor	%rbp,0x7f563412(%rbx)
	xor	%rbp,0x7f563412(%rsp)
	xor	%rbp,0x7f563412(%rbp)
	xor	%rbp,0x7f563412(%rsi)
	xor	%rbp,0x7f563412(%rdi)
	xor	%rbp,0x7f563412(%r8)
	xor	%rbp,0x7f563412(%r9)
	xor	%rbp,0x7f563412(%r10)
	xor	%rbp,0x7f563412(%r11)
	xor	%rbp,0x7f563412(%r12)
	xor	%rbp,0x7f563412(%r13)
	xor	%rbp,0x7f563412(%r14)
	xor	%rbp,0x7f563412(%r15)
	nop
	xor	%rsi,0x7f563412(%rax)
	xor	%rsi,0x7f563412(%rcx)
	xor	%rsi,0x7f563412(%rdx)
	xor	%rsi,0x7f563412(%rbx)
	xor	%rsi,0x7f563412(%rsp)
	xor	%rsi,0x7f563412(%rbp)
	xor	%rsi,0x7f563412(%rsi)
	xor	%rsi,0x7f563412(%rdi)
	xor	%rsi,0x7f563412(%r8)
	xor	%rsi,0x7f563412(%r9)
	xor	%rsi,0x7f563412(%r10)
	xor	%rsi,0x7f563412(%r11)
	xor	%rsi,0x7f563412(%r12)
	xor	%rsi,0x7f563412(%r13)
	xor	%rsi,0x7f563412(%r14)
	xor	%rsi,0x7f563412(%r15)
	nop
	xor	%rdi,0x7f563412(%rax)
	xor	%rdi,0x7f563412(%rcx)
	xor	%rdi,0x7f563412(%rdx)
	xor	%rdi,0x7f563412(%rbx)
	xor	%rdi,0x7f563412(%rsp)
	xor	%rdi,0x7f563412(%rbp)
	xor	%rdi,0x7f563412(%rsi)
	xor	%rdi,0x7f563412(%rdi)
	xor	%rdi,0x7f563412(%r8)
	xor	%rdi,0x7f563412(%r9)
	xor	%rdi,0x7f563412(%r10)
	xor	%rdi,0x7f563412(%r11)
	xor	%rdi,0x7f563412(%r12)
	xor	%rdi,0x7f563412(%r13)
	xor	%rdi,0x7f563412(%r14)
	xor	%rdi,0x7f563412(%r15)
	nop
	xor	%r8, 0x7f563412(%rax)
	xor	%r8, 0x7f563412(%rcx)
	xor	%r8, 0x7f563412(%rdx)
	xor	%r8, 0x7f563412(%rbx)
	xor	%r8, 0x7f563412(%rsp)
	xor	%r8, 0x7f563412(%rbp)
	xor	%r8, 0x7f563412(%rsi)
	xor	%r8, 0x7f563412(%rdi)
	xor	%r8, 0x7f563412(%r8)
	xor	%r8, 0x7f563412(%r9)
	xor	%r8, 0x7f563412(%r10)
	xor	%r8, 0x7f563412(%r11)
	xor	%r8, 0x7f563412(%r12)
	xor	%r8, 0x7f563412(%r13)
	xor	%r8, 0x7f563412(%r14)
	xor	%r8, 0x7f563412(%r15)
	nop
	xor	%r9, 0x7f563412(%rax)
	xor	%r9, 0x7f563412(%rcx)
	xor	%r9, 0x7f563412(%rdx)
	xor	%r9, 0x7f563412(%rbx)
	xor	%r9, 0x7f563412(%rsp)
	xor	%r9, 0x7f563412(%rbp)
	xor	%r9, 0x7f563412(%rsi)
	xor	%r9, 0x7f563412(%rdi)
	xor	%r9, 0x7f563412(%r8)
	xor	%r9, 0x7f563412(%r9)
	xor	%r9, 0x7f563412(%r10)
	xor	%r9, 0x7f563412(%r11)
	xor	%r9, 0x7f563412(%r12)
	xor	%r9, 0x7f563412(%r13)
	xor	%r9, 0x7f563412(%r14)
	xor	%r9, 0x7f563412(%r15)
	nop
	xor	%r10,0x7f563412(%rax)
	xor	%r10,0x7f563412(%rcx)
	xor	%r10,0x7f563412(%rdx)
	xor	%r10,0x7f563412(%rbx)
	xor	%r10,0x7f563412(%rsp)
	xor	%r10,0x7f563412(%rbp)
	xor	%r10,0x7f563412(%rsi)
	xor	%r10,0x7f563412(%rdi)
	xor	%r10,0x7f563412(%r8)
	xor	%r10,0x7f563412(%r9)
	xor	%r10,0x7f563412(%r10)
	xor	%r10,0x7f563412(%r11)
	xor	%r10,0x7f563412(%r12)
	xor	%r10,0x7f563412(%r13)
	xor	%r10,0x7f563412(%r14)
	xor	%r10,0x7f563412(%r15)
	nop
	xor	%r11,0x7f563412(%rax)
	xor	%r11,0x7f563412(%rcx)
	xor	%r11,0x7f563412(%rdx)
	xor	%r11,0x7f563412(%rbx)
	xor	%r11,0x7f563412(%rsp)
	xor	%r11,0x7f563412(%rbp)
	xor	%r11,0x7f563412(%rsi)
	xor	%r11,0x7f563412(%rdi)
	xor	%r11,0x7f563412(%r8)
	xor	%r11,0x7f563412(%r9)
	xor	%r11,0x7f563412(%r10)
	xor	%r11,0x7f563412(%r11)
	xor	%r11,0x7f563412(%r12)
	xor	%r11,0x7f563412(%r13)
	xor	%r11,0x7f563412(%r14)
	xor	%r11,0x7f563412(%r15)
	nop
	xor	%r12,0x7f563412(%rax)
	xor	%r12,0x7f563412(%rcx)
	xor	%r12,0x7f563412(%rdx)
	xor	%r12,0x7f563412(%rbx)
	xor	%r12,0x7f563412(%rsp)
	xor	%r12,0x7f563412(%rbp)
	xor	%r12,0x7f563412(%rsi)
	xor	%r12,0x7f563412(%rdi)
	xor	%r12,0x7f563412(%r8)
	xor	%r12,0x7f563412(%r9)
	xor	%r12,0x7f563412(%r10)
	xor	%r12,0x7f563412(%r11)
	xor	%r12,0x7f563412(%r12)
	xor	%r12,0x7f563412(%r13)
	xor	%r12,0x7f563412(%r14)
	xor	%r12,0x7f563412(%r15)
	nop
	xor	%r13,0x7f563412(%rax)
	xor	%r13,0x7f563412(%rcx)
	xor	%r13,0x7f563412(%rdx)
	xor	%r13,0x7f563412(%rbx)
	xor	%r13,0x7f563412(%rsp)
	xor	%r13,0x7f563412(%rbp)
	xor	%r13,0x7f563412(%rsi)
	xor	%r13,0x7f563412(%rdi)
	xor	%r13,0x7f563412(%r8)
	xor	%r13,0x7f563412(%r9)
	xor	%r13,0x7f563412(%r10)
	xor	%r13,0x7f563412(%r11)
	xor	%r13,0x7f563412(%r12)
	xor	%r13,0x7f563412(%r13)
	xor	%r13,0x7f563412(%r14)
	xor	%r13,0x7f563412(%r15)
	nop
	xor	%r14,0x7f563412(%rax)
	xor	%r14,0x7f563412(%rcx)
	xor	%r14,0x7f563412(%rdx)
	xor	%r14,0x7f563412(%rbx)
	xor	%r14,0x7f563412(%rsp)
	xor	%r14,0x7f563412(%rbp)
	xor	%r14,0x7f563412(%rsi)
	xor	%r14,0x7f563412(%rdi)
	xor	%r14,0x7f563412(%r8)
	xor	%r14,0x7f563412(%r9)
	xor	%r14,0x7f563412(%r10)
	xor	%r14,0x7f563412(%r11)
	xor	%r14,0x7f563412(%r12)
	xor	%r14,0x7f563412(%r13)
	xor	%r14,0x7f563412(%r14)
	xor	%r14,0x7f563412(%r15)
	nop
	xor	%r15,0x7f563412(%rax)
	xor	%r15,0x7f563412(%rcx)
	xor	%r15,0x7f563412(%rdx)
	xor	%r15,0x7f563412(%rbx)
	xor	%r15,0x7f563412(%rsp)
	xor	%r15,0x7f563412(%rbp)
	xor	%r15,0x7f563412(%rsi)
	xor	%r15,0x7f563412(%rdi)
	xor	%r15,0x7f563412(%r8)
	xor	%r15,0x7f563412(%r9)
	xor	%r15,0x7f563412(%r10)
	xor	%r15,0x7f563412(%r11)
	xor	%r15,0x7f563412(%r12)
	xor	%r15,0x7f563412(%r13)
	xor	%r15,0x7f563412(%r14)
	xor	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
