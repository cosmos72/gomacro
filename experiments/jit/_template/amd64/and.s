	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	And_s32
	.type	And_s32, @function
And_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// and	$-0x55667788,%rax
	and	$-0x55667788,%rcx
	and	$-0x55667788,%rdx
	and	$-0x55667788,%rbx
	and	$-0x55667788,%rsp
	and	$-0x55667788,%rbp
	and	$-0x55667788,%rsi
	and	$-0x55667788,%rdi
	and	$-0x55667788,%r8
	and	$-0x55667788,%r9
	and	$-0x55667788,%r10
	and	$-0x55667788,%r11
	and	$-0x55667788,%r12
	and	$-0x55667788,%r13
	and	$-0x55667788,%r14
	and	$-0x55667788,%r15
	.cfi_endproc


	.p2align 4,,15
	.globl	And
	.type	And, @function
And:
	.cfi_startproc
	and	%rax,%rax
	and	%rax,%rcx
	and	%rax,%rdx
	and	%rax,%rbx
	and	%rax,%rsp
	and	%rax,%rbp
	and	%rax,%rsi
	and	%rax,%rdi
	and	%rax,%r8
	and	%rax,%r9
	and	%rax,%r10
	and	%rax,%r11
	and	%rax,%r12
	and	%rax,%r13
	and	%rax,%r14
	and	%rax,%r15
	nop
	and	%rcx,%rax
	and	%rcx,%rcx
	and	%rcx,%rdx
	and	%rcx,%rbx
	and	%rcx,%rsp
	and	%rcx,%rbp
	and	%rcx,%rsi
	and	%rcx,%rdi
	and	%rcx,%r8
	and	%rcx,%r9
	and	%rcx,%r10
	and	%rcx,%r11
	and	%rcx,%r12
	and	%rcx,%r13
	and	%rcx,%r14
	and	%rcx,%r15
	nop
	and	%rdx,%rax
	and	%rdx,%rcx
	and	%rdx,%rdx
	and	%rdx,%rbx
	and	%rdx,%rsp
	and	%rdx,%rbp
	and	%rdx,%rsi
	and	%rdx,%rdi
	and	%rdx,%r8
	and	%rdx,%r9
	and	%rdx,%r10
	and	%rdx,%r11
	and	%rdx,%r12
	and	%rdx,%r13
	and	%rdx,%r14
	and	%rdx,%r15
	nop
	and	%rbx,%rax
	and	%rbx,%rcx
	and	%rbx,%rdx
	and	%rbx,%rbx
	and	%rbx,%rsp
	and	%rbx,%rbp
	and	%rbx,%rsi
	and	%rbx,%rdi
	and	%rbx,%r8
	and	%rbx,%r9
	and	%rbx,%r10
	and	%rbx,%r11
	and	%rbx,%r12
	and	%rbx,%r13
	and	%rbx,%r14
	and	%rbx,%r15
	nop
	and	%rsp,%rax
	and	%rsp,%rcx
	and	%rsp,%rdx
	and	%rsp,%rbx
	and	%rsp,%rsp
	and	%rsp,%rbp
	and	%rsp,%rsi
	and	%rsp,%rdi
	and	%rsp,%r8
	and	%rsp,%r9
	and	%rsp,%r10
	and	%rsp,%r11
	and	%rsp,%r12
	and	%rsp,%r13
	and	%rsp,%r14
	and	%rsp,%r15
	nop
	and	%rbp,%rax
	and	%rbp,%rcx
	and	%rbp,%rdx
	and	%rbp,%rbx
	and	%rbp,%rsp
	and	%rbp,%rbp
	and	%rbp,%rsi
	and	%rbp,%rdi
	and	%rbp,%r8
	and	%rbp,%r9
	and	%rbp,%r10
	and	%rbp,%r11
	and	%rbp,%r12
	and	%rbp,%r13
	and	%rbp,%r14
	and	%rbp,%r15
	nop
	and	%rsi,%rax
	and	%rsi,%rcx
	and	%rsi,%rdx
	and	%rsi,%rbx
	and	%rsi,%rsp
	and	%rsi,%rbp
	and	%rsi,%rsi
	and	%rsi,%rdi
	and	%rsi,%r8
	and	%rsi,%r9
	and	%rsi,%r10
	and	%rsi,%r11
	and	%rsi,%r12
	and	%rsi,%r13
	and	%rsi,%r14
	and	%rsi,%r15
	nop
	and	%rdi,%rax
	and	%rdi,%rcx
	and	%rdi,%rdx
	and	%rdi,%rbx
	and	%rdi,%rsp
	and	%rdi,%rbp
	and	%rdi,%rsi
	and	%rdi,%rdi
	and	%rdi,%r8
	and	%rdi,%r9
	and	%rdi,%r10
	and	%rdi,%r11
	and	%rdi,%r12
	and	%rdi,%r13
	and	%rdi,%r14
	and	%rdi,%r15
	nop
	and	%r8, %rax
	and	%r8, %rcx
	and	%r8, %rdx
	and	%r8, %rbx
	and	%r8, %rsp
	and	%r8, %rbp
	and	%r8, %rsi
	and	%r8, %rdi
	and	%r8, %r8
	and	%r8, %r9
	and	%r8, %r10
	and	%r8, %r11
	and	%r8, %r12
	and	%r8, %r13
	and	%r8, %r14
	and	%r8, %r15
	nop
	and	%r9, %rax
	and	%r9, %rcx
	and	%r9, %rdx
	and	%r9, %rbx
	and	%r9, %rsp
	and	%r9, %rbp
	and	%r9, %rsi
	and	%r9, %rdi
	and	%r9, %r8
	and	%r9, %r9
	and	%r9, %r10
	and	%r9, %r11
	and	%r9, %r12
	and	%r9, %r13
	and	%r9, %r14
	and	%r9, %r15
	nop
	and	%r10,%rax
	and	%r10,%rcx
	and	%r10,%rdx
	and	%r10,%rbx
	and	%r10,%rsp
	and	%r10,%rbp
	and	%r10,%rsi
	and	%r10,%rdi
	and	%r10,%r8
	and	%r10,%r9
	and	%r10,%r10
	and	%r10,%r11
	and	%r10,%r12
	and	%r10,%r13
	and	%r10,%r14
	and	%r10,%r15
	nop
	and	%r11,%rax
	and	%r11,%rcx
	and	%r11,%rdx
	and	%r11,%rbx
	and	%r11,%rsp
	and	%r11,%rbp
	and	%r11,%rsi
	and	%r11,%rdi
	and	%r11,%r8
	and	%r11,%r9
	and	%r11,%r10
	and	%r11,%r11
	and	%r11,%r12
	and	%r11,%r13
	and	%r11,%r14
	and	%r11,%r15
	nop
	and	%r12,%rax
	and	%r12,%rcx
	and	%r12,%rdx
	and	%r12,%rbx
	and	%r12,%rsp
	and	%r12,%rbp
	and	%r12,%rsi
	and	%r12,%rdi
	and	%r12,%r8
	and	%r12,%r9
	and	%r12,%r10
	and	%r12,%r11
	and	%r12,%r12
	and	%r12,%r13
	and	%r12,%r14
	and	%r12,%r15
	nop
	and	%r13,%rax
	and	%r13,%rcx
	and	%r13,%rdx
	and	%r13,%rbx
	and	%r13,%rsp
	and	%r13,%rbp
	and	%r13,%rsi
	and	%r13,%rdi
	and	%r13,%r8
	and	%r13,%r9
	and	%r13,%r10
	and	%r13,%r11
	and	%r13,%r12
	and	%r13,%r13
	and	%r13,%r14
	and	%r13,%r15
	nop
	and	%r14,%rax
	and	%r14,%rcx
	and	%r14,%rdx
	and	%r14,%rbx
	and	%r14,%rsp
	and	%r14,%rbp
	and	%r14,%rsi
	and	%r14,%rdi
	and	%r14,%r8
	and	%r14,%r9
	and	%r14,%r10
	and	%r14,%r11
	and	%r14,%r12
	and	%r14,%r13
	and	%r14,%r14
	and	%r14,%r15
	nop
	and	%r15,%rax
	and	%r15,%rcx
	and	%r15,%rdx
	and	%r15,%rbx
	and	%r15,%rsp
	and	%r15,%rbp
	and	%r15,%rsi
	and	%r15,%rdi
	and	%r15,%r8
	and	%r15,%r9
	and	%r15,%r10
	and	%r15,%r11
	and	%r15,%r12
	and	%r15,%r13
	and	%r15,%r14
	and	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndMemReg
	.type	AndMemReg, @function
AndMemReg:
	.cfi_startproc
	and	%rax,(%rax)
	and	%rax,(%rcx)
	and	%rax,(%rdx)
	and	%rax,(%rbx)
	and	%rax,(%rsp)
	and	%rax,(%rbp)
	and	%rax,(%rsi)
	and	%rax,(%rdi)
	and	%rax,(%r8)
	and	%rax,(%r9)
	and	%rax,(%r10)
	and	%rax,(%r11)
	and	%rax,(%r12)
	and	%rax,(%r13)
	and	%rax,(%r14)
	and	%rax,(%r15)
	nop
	and	%rcx,(%rax)
	and	%rcx,(%rcx)
	and	%rcx,(%rdx)
	and	%rcx,(%rbx)
	and	%rcx,(%rsp)
	and	%rcx,(%rbp)
	and	%rcx,(%rsi)
	and	%rcx,(%rdi)
	and	%rcx,(%r8)
	and	%rcx,(%r9)
	and	%rcx,(%r10)
	and	%rcx,(%r11)
	and	%rcx,(%r12)
	and	%rcx,(%r13)
	and	%rcx,(%r14)
	and	%rcx,(%r15)
	nop
	and	%rdx,(%rax)
	and	%rdx,(%rcx)
	and	%rdx,(%rdx)
	and	%rdx,(%rbx)
	and	%rdx,(%rsp)
	and	%rdx,(%rbp)
	and	%rdx,(%rsi)
	and	%rdx,(%rdi)
	and	%rdx,(%r8)
	and	%rdx,(%r9)
	and	%rdx,(%r10)
	and	%rdx,(%r11)
	and	%rdx,(%r12)
	and	%rdx,(%r13)
	and	%rdx,(%r14)
	and	%rdx,(%r15)
	nop
	and	%rbx,(%rax)
	and	%rbx,(%rcx)
	and	%rbx,(%rdx)
	and	%rbx,(%rbx)
	and	%rbx,(%rsp)
	and	%rbx,(%rbp)
	and	%rbx,(%rsi)
	and	%rbx,(%rdi)
	and	%rbx,(%r8)
	and	%rbx,(%r9)
	and	%rbx,(%r10)
	and	%rbx,(%r11)
	and	%rbx,(%r12)
	and	%rbx,(%r13)
	and	%rbx,(%r14)
	and	%rbx,(%r15)
	nop
	and	%rsp,(%rax)
	and	%rsp,(%rcx)
	and	%rsp,(%rdx)
	and	%rsp,(%rbx)
	and	%rsp,(%rsp)
	and	%rsp,(%rbp)
	and	%rsp,(%rsi)
	and	%rsp,(%rdi)
	and	%rsp,(%r8)
	and	%rsp,(%r9)
	and	%rsp,(%r10)
	and	%rsp,(%r11)
	and	%rsp,(%r12)
	and	%rsp,(%r13)
	and	%rsp,(%r14)
	and	%rsp,(%r15)
	nop
	and	%rbp,(%rax)
	and	%rbp,(%rcx)
	and	%rbp,(%rdx)
	and	%rbp,(%rbx)
	and	%rbp,(%rsp)
	and	%rbp,(%rbp)
	and	%rbp,(%rsi)
	and	%rbp,(%rdi)
	and	%rbp,(%r8)
	and	%rbp,(%r9)
	and	%rbp,(%r10)
	and	%rbp,(%r11)
	and	%rbp,(%r12)
	and	%rbp,(%r13)
	and	%rbp,(%r14)
	and	%rbp,(%r15)
	nop
	and	%rsi,(%rax)
	and	%rsi,(%rcx)
	and	%rsi,(%rdx)
	and	%rsi,(%rbx)
	and	%rsi,(%rsp)
	and	%rsi,(%rbp)
	and	%rsi,(%rsi)
	and	%rsi,(%rdi)
	and	%rsi,(%r8)
	and	%rsi,(%r9)
	and	%rsi,(%r10)
	and	%rsi,(%r11)
	and	%rsi,(%r12)
	and	%rsi,(%r13)
	and	%rsi,(%r14)
	and	%rsi,(%r15)
	nop
	and	%rdi,(%rax)
	and	%rdi,(%rcx)
	and	%rdi,(%rdx)
	and	%rdi,(%rbx)
	and	%rdi,(%rsp)
	and	%rdi,(%rbp)
	and	%rdi,(%rsi)
	and	%rdi,(%rdi)
	and	%rdi,(%r8)
	and	%rdi,(%r9)
	and	%rdi,(%r10)
	and	%rdi,(%r11)
	and	%rdi,(%r12)
	and	%rdi,(%r13)
	and	%rdi,(%r14)
	and	%rdi,(%r15)
	nop
	and	%r8, (%rax)
	and	%r8, (%rcx)
	and	%r8, (%rdx)
	and	%r8, (%rbx)
	and	%r8, (%rsp)
	and	%r8, (%rbp)
	and	%r8, (%rsi)
	and	%r8, (%rdi)
	and	%r8, (%r8)
	and	%r8, (%r9)
	and	%r8, (%r10)
	and	%r8, (%r11)
	and	%r8, (%r12)
	and	%r8, (%r13)
	and	%r8, (%r14)
	and	%r8, (%r15)
	nop
	and	%r9, (%rax)
	and	%r9, (%rcx)
	and	%r9, (%rdx)
	and	%r9, (%rbx)
	and	%r9, (%rsp)
	and	%r9, (%rbp)
	and	%r9, (%rsi)
	and	%r9, (%rdi)
	and	%r9, (%r8)
	and	%r9, (%r9)
	and	%r9, (%r10)
	and	%r9, (%r11)
	and	%r9, (%r12)
	and	%r9, (%r13)
	and	%r9, (%r14)
	and	%r9, (%r15)
	nop
	and	%r10,(%rax)
	and	%r10,(%rcx)
	and	%r10,(%rdx)
	and	%r10,(%rbx)
	and	%r10,(%rsp)
	and	%r10,(%rbp)
	and	%r10,(%rsi)
	and	%r10,(%rdi)
	and	%r10,(%r8)
	and	%r10,(%r9)
	and	%r10,(%r10)
	and	%r10,(%r11)
	and	%r10,(%r12)
	and	%r10,(%r13)
	and	%r10,(%r14)
	and	%r10,(%r15)
	nop
	and	%r11,(%rax)
	and	%r11,(%rcx)
	and	%r11,(%rdx)
	and	%r11,(%rbx)
	and	%r11,(%rsp)
	and	%r11,(%rbp)
	and	%r11,(%rsi)
	and	%r11,(%rdi)
	and	%r11,(%r8)
	and	%r11,(%r9)
	and	%r11,(%r10)
	and	%r11,(%r11)
	and	%r11,(%r12)
	and	%r11,(%r13)
	and	%r11,(%r14)
	and	%r11,(%r15)
	nop
	and	%r12,(%rax)
	and	%r12,(%rcx)
	and	%r12,(%rdx)
	and	%r12,(%rbx)
	and	%r12,(%rsp)
	and	%r12,(%rbp)
	and	%r12,(%rsi)
	and	%r12,(%rdi)
	and	%r12,(%r8)
	and	%r12,(%r9)
	and	%r12,(%r10)
	and	%r12,(%r11)
	and	%r12,(%r12)
	and	%r12,(%r13)
	and	%r12,(%r14)
	and	%r12,(%r15)
	nop
	and	%r13,(%rax)
	and	%r13,(%rcx)
	and	%r13,(%rdx)
	and	%r13,(%rbx)
	and	%r13,(%rsp)
	and	%r13,(%rbp)
	and	%r13,(%rsi)
	and	%r13,(%rdi)
	and	%r13,(%r8)
	and	%r13,(%r9)
	and	%r13,(%r10)
	and	%r13,(%r11)
	and	%r13,(%r12)
	and	%r13,(%r13)
	and	%r13,(%r14)
	and	%r13,(%r15)
	nop
	and	%r14,(%rax)
	and	%r14,(%rcx)
	and	%r14,(%rdx)
	and	%r14,(%rbx)
	and	%r14,(%rsp)
	and	%r14,(%rbp)
	and	%r14,(%rsi)
	and	%r14,(%rdi)
	and	%r14,(%r8)
	and	%r14,(%r9)
	and	%r14,(%r10)
	and	%r14,(%r11)
	and	%r14,(%r12)
	and	%r14,(%r13)
	and	%r14,(%r14)
	and	%r14,(%r15)
	nop
	and	%r15,(%rax)
	and	%r15,(%rcx)
	and	%r15,(%rdx)
	and	%r15,(%rbx)
	and	%r15,(%rsp)
	and	%r15,(%rbp)
	and	%r15,(%rsi)
	and	%r15,(%rdi)
	and	%r15,(%r8)
	and	%r15,(%r9)
	and	%r15,(%r10)
	and	%r15,(%r11)
	and	%r15,(%r12)
	and	%r15,(%r13)
	and	%r15,(%r14)
	and	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndMem8Reg
	.type	AndMem8Reg, @function
AndMem8Reg:
	.cfi_startproc
	and	%rax,0x7f(%rax)
	and	%rax,0x7f(%rcx)
	and	%rax,0x7f(%rdx)
	and	%rax,0x7f(%rbx)
	and	%rax,0x7f(%rsp)
	and	%rax,0x7f(%rbp)
	and	%rax,0x7f(%rsi)
	and	%rax,0x7f(%rdi)
	and	%rax,0x7f(%r8)
	and	%rax,0x7f(%r9)
	and	%rax,0x7f(%r10)
	and	%rax,0x7f(%r11)
	and	%rax,0x7f(%r12)
	and	%rax,0x7f(%r13)
	and	%rax,0x7f(%r14)
	and	%rax,0x7f(%r15)
	nop
	and	%rcx,0x7f(%rax)
	and	%rcx,0x7f(%rcx)
	and	%rcx,0x7f(%rdx)
	and	%rcx,0x7f(%rbx)
	and	%rcx,0x7f(%rsp)
	and	%rcx,0x7f(%rbp)
	and	%rcx,0x7f(%rsi)
	and	%rcx,0x7f(%rdi)
	and	%rcx,0x7f(%r8)
	and	%rcx,0x7f(%r9)
	and	%rcx,0x7f(%r10)
	and	%rcx,0x7f(%r11)
	and	%rcx,0x7f(%r12)
	and	%rcx,0x7f(%r13)
	and	%rcx,0x7f(%r14)
	and	%rcx,0x7f(%r15)
	nop
	and	%rdx,0x7f(%rax)
	and	%rdx,0x7f(%rcx)
	and	%rdx,0x7f(%rdx)
	and	%rdx,0x7f(%rbx)
	and	%rdx,0x7f(%rsp)
	and	%rdx,0x7f(%rbp)
	and	%rdx,0x7f(%rsi)
	and	%rdx,0x7f(%rdi)
	and	%rdx,0x7f(%r8)
	and	%rdx,0x7f(%r9)
	and	%rdx,0x7f(%r10)
	and	%rdx,0x7f(%r11)
	and	%rdx,0x7f(%r12)
	and	%rdx,0x7f(%r13)
	and	%rdx,0x7f(%r14)
	and	%rdx,0x7f(%r15)
	nop
	and	%rbx,0x7f(%rax)
	and	%rbx,0x7f(%rcx)
	and	%rbx,0x7f(%rdx)
	and	%rbx,0x7f(%rbx)
	and	%rbx,0x7f(%rsp)
	and	%rbx,0x7f(%rbp)
	and	%rbx,0x7f(%rsi)
	and	%rbx,0x7f(%rdi)
	and	%rbx,0x7f(%r8)
	and	%rbx,0x7f(%r9)
	and	%rbx,0x7f(%r10)
	and	%rbx,0x7f(%r11)
	and	%rbx,0x7f(%r12)
	and	%rbx,0x7f(%r13)
	and	%rbx,0x7f(%r14)
	and	%rbx,0x7f(%r15)
	nop
	and	%rsp,0x7f(%rax)
	and	%rsp,0x7f(%rcx)
	and	%rsp,0x7f(%rdx)
	and	%rsp,0x7f(%rbx)
	and	%rsp,0x7f(%rsp)
	and	%rsp,0x7f(%rbp)
	and	%rsp,0x7f(%rsi)
	and	%rsp,0x7f(%rdi)
	and	%rsp,0x7f(%r8)
	and	%rsp,0x7f(%r9)
	and	%rsp,0x7f(%r10)
	and	%rsp,0x7f(%r11)
	and	%rsp,0x7f(%r12)
	and	%rsp,0x7f(%r13)
	and	%rsp,0x7f(%r14)
	and	%rsp,0x7f(%r15)
	nop
	and	%rbp,0x7f(%rax)
	and	%rbp,0x7f(%rcx)
	and	%rbp,0x7f(%rdx)
	and	%rbp,0x7f(%rbx)
	and	%rbp,0x7f(%rsp)
	and	%rbp,0x7f(%rbp)
	and	%rbp,0x7f(%rsi)
	and	%rbp,0x7f(%rdi)
	and	%rbp,0x7f(%r8)
	and	%rbp,0x7f(%r9)
	and	%rbp,0x7f(%r10)
	and	%rbp,0x7f(%r11)
	and	%rbp,0x7f(%r12)
	and	%rbp,0x7f(%r13)
	and	%rbp,0x7f(%r14)
	and	%rbp,0x7f(%r15)
	nop
	and	%rsi,0x7f(%rax)
	and	%rsi,0x7f(%rcx)
	and	%rsi,0x7f(%rdx)
	and	%rsi,0x7f(%rbx)
	and	%rsi,0x7f(%rsp)
	and	%rsi,0x7f(%rbp)
	and	%rsi,0x7f(%rsi)
	and	%rsi,0x7f(%rdi)
	and	%rsi,0x7f(%r8)
	and	%rsi,0x7f(%r9)
	and	%rsi,0x7f(%r10)
	and	%rsi,0x7f(%r11)
	and	%rsi,0x7f(%r12)
	and	%rsi,0x7f(%r13)
	and	%rsi,0x7f(%r14)
	and	%rsi,0x7f(%r15)
	nop
	and	%rdi,0x7f(%rax)
	and	%rdi,0x7f(%rcx)
	and	%rdi,0x7f(%rdx)
	and	%rdi,0x7f(%rbx)
	and	%rdi,0x7f(%rsp)
	and	%rdi,0x7f(%rbp)
	and	%rdi,0x7f(%rsi)
	and	%rdi,0x7f(%rdi)
	and	%rdi,0x7f(%r8)
	and	%rdi,0x7f(%r9)
	and	%rdi,0x7f(%r10)
	and	%rdi,0x7f(%r11)
	and	%rdi,0x7f(%r12)
	and	%rdi,0x7f(%r13)
	and	%rdi,0x7f(%r14)
	and	%rdi,0x7f(%r15)
	nop
	and	%r8, 0x7f(%rax)
	and	%r8, 0x7f(%rcx)
	and	%r8, 0x7f(%rdx)
	and	%r8, 0x7f(%rbx)
	and	%r8, 0x7f(%rsp)
	and	%r8, 0x7f(%rbp)
	and	%r8, 0x7f(%rsi)
	and	%r8, 0x7f(%rdi)
	and	%r8, 0x7f(%r8)
	and	%r8, 0x7f(%r9)
	and	%r8, 0x7f(%r10)
	and	%r8, 0x7f(%r11)
	and	%r8, 0x7f(%r12)
	and	%r8, 0x7f(%r13)
	and	%r8, 0x7f(%r14)
	and	%r8, 0x7f(%r15)
	nop
	and	%r9, 0x7f(%rax)
	and	%r9, 0x7f(%rcx)
	and	%r9, 0x7f(%rdx)
	and	%r9, 0x7f(%rbx)
	and	%r9, 0x7f(%rsp)
	and	%r9, 0x7f(%rbp)
	and	%r9, 0x7f(%rsi)
	and	%r9, 0x7f(%rdi)
	and	%r9, 0x7f(%r8)
	and	%r9, 0x7f(%r9)
	and	%r9, 0x7f(%r10)
	and	%r9, 0x7f(%r11)
	and	%r9, 0x7f(%r12)
	and	%r9, 0x7f(%r13)
	and	%r9, 0x7f(%r14)
	and	%r9, 0x7f(%r15)
	nop
	and	%r10,0x7f(%rax)
	and	%r10,0x7f(%rcx)
	and	%r10,0x7f(%rdx)
	and	%r10,0x7f(%rbx)
	and	%r10,0x7f(%rsp)
	and	%r10,0x7f(%rbp)
	and	%r10,0x7f(%rsi)
	and	%r10,0x7f(%rdi)
	and	%r10,0x7f(%r8)
	and	%r10,0x7f(%r9)
	and	%r10,0x7f(%r10)
	and	%r10,0x7f(%r11)
	and	%r10,0x7f(%r12)
	and	%r10,0x7f(%r13)
	and	%r10,0x7f(%r14)
	and	%r10,0x7f(%r15)
	nop
	and	%r11,0x7f(%rax)
	and	%r11,0x7f(%rcx)
	and	%r11,0x7f(%rdx)
	and	%r11,0x7f(%rbx)
	and	%r11,0x7f(%rsp)
	and	%r11,0x7f(%rbp)
	and	%r11,0x7f(%rsi)
	and	%r11,0x7f(%rdi)
	and	%r11,0x7f(%r8)
	and	%r11,0x7f(%r9)
	and	%r11,0x7f(%r10)
	and	%r11,0x7f(%r11)
	and	%r11,0x7f(%r12)
	and	%r11,0x7f(%r13)
	and	%r11,0x7f(%r14)
	and	%r11,0x7f(%r15)
	nop
	and	%r12,0x7f(%rax)
	and	%r12,0x7f(%rcx)
	and	%r12,0x7f(%rdx)
	and	%r12,0x7f(%rbx)
	and	%r12,0x7f(%rsp)
	and	%r12,0x7f(%rbp)
	and	%r12,0x7f(%rsi)
	and	%r12,0x7f(%rdi)
	and	%r12,0x7f(%r8)
	and	%r12,0x7f(%r9)
	and	%r12,0x7f(%r10)
	and	%r12,0x7f(%r11)
	and	%r12,0x7f(%r12)
	and	%r12,0x7f(%r13)
	and	%r12,0x7f(%r14)
	and	%r12,0x7f(%r15)
	nop
	and	%r13,0x7f(%rax)
	and	%r13,0x7f(%rcx)
	and	%r13,0x7f(%rdx)
	and	%r13,0x7f(%rbx)
	and	%r13,0x7f(%rsp)
	and	%r13,0x7f(%rbp)
	and	%r13,0x7f(%rsi)
	and	%r13,0x7f(%rdi)
	and	%r13,0x7f(%r8)
	and	%r13,0x7f(%r9)
	and	%r13,0x7f(%r10)
	and	%r13,0x7f(%r11)
	and	%r13,0x7f(%r12)
	and	%r13,0x7f(%r13)
	and	%r13,0x7f(%r14)
	and	%r13,0x7f(%r15)
	nop
	and	%r14,0x7f(%rax)
	and	%r14,0x7f(%rcx)
	and	%r14,0x7f(%rdx)
	and	%r14,0x7f(%rbx)
	and	%r14,0x7f(%rsp)
	and	%r14,0x7f(%rbp)
	and	%r14,0x7f(%rsi)
	and	%r14,0x7f(%rdi)
	and	%r14,0x7f(%r8)
	and	%r14,0x7f(%r9)
	and	%r14,0x7f(%r10)
	and	%r14,0x7f(%r11)
	and	%r14,0x7f(%r12)
	and	%r14,0x7f(%r13)
	and	%r14,0x7f(%r14)
	and	%r14,0x7f(%r15)
	nop
	and	%r15,0x7f(%rax)
	and	%r15,0x7f(%rcx)
	and	%r15,0x7f(%rdx)
	and	%r15,0x7f(%rbx)
	and	%r15,0x7f(%rsp)
	and	%r15,0x7f(%rbp)
	and	%r15,0x7f(%rsi)
	and	%r15,0x7f(%rdi)
	and	%r15,0x7f(%r8)
	and	%r15,0x7f(%r9)
	and	%r15,0x7f(%r10)
	and	%r15,0x7f(%r11)
	and	%r15,0x7f(%r12)
	and	%r15,0x7f(%r13)
	and	%r15,0x7f(%r14)
	and	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndMem32Reg
	.type	AndMem32Reg, @function
AndMem32Reg:
	.cfi_startproc
	and	%rax,0x7f563412(%rax)
	and	%rax,0x7f563412(%rcx)
	and	%rax,0x7f563412(%rdx)
	and	%rax,0x7f563412(%rbx)
	and	%rax,0x7f563412(%rsp)
	and	%rax,0x7f563412(%rbp)
	and	%rax,0x7f563412(%rsi)
	and	%rax,0x7f563412(%rdi)
	and	%rax,0x7f563412(%r8)
	and	%rax,0x7f563412(%r9)
	and	%rax,0x7f563412(%r10)
	and	%rax,0x7f563412(%r11)
	and	%rax,0x7f563412(%r12)
	and	%rax,0x7f563412(%r13)
	and	%rax,0x7f563412(%r14)
	and	%rax,0x7f563412(%r15)
	nop
	and	%rcx,0x7f563412(%rax)
	and	%rcx,0x7f563412(%rcx)
	and	%rcx,0x7f563412(%rdx)
	and	%rcx,0x7f563412(%rbx)
	and	%rcx,0x7f563412(%rsp)
	and	%rcx,0x7f563412(%rbp)
	and	%rcx,0x7f563412(%rsi)
	and	%rcx,0x7f563412(%rdi)
	and	%rcx,0x7f563412(%r8)
	and	%rcx,0x7f563412(%r9)
	and	%rcx,0x7f563412(%r10)
	and	%rcx,0x7f563412(%r11)
	and	%rcx,0x7f563412(%r12)
	and	%rcx,0x7f563412(%r13)
	and	%rcx,0x7f563412(%r14)
	and	%rcx,0x7f563412(%r15)
	nop
	and	%rdx,0x7f563412(%rax)
	and	%rdx,0x7f563412(%rcx)
	and	%rdx,0x7f563412(%rdx)
	and	%rdx,0x7f563412(%rbx)
	and	%rdx,0x7f563412(%rsp)
	and	%rdx,0x7f563412(%rbp)
	and	%rdx,0x7f563412(%rsi)
	and	%rdx,0x7f563412(%rdi)
	and	%rdx,0x7f563412(%r8)
	and	%rdx,0x7f563412(%r9)
	and	%rdx,0x7f563412(%r10)
	and	%rdx,0x7f563412(%r11)
	and	%rdx,0x7f563412(%r12)
	and	%rdx,0x7f563412(%r13)
	and	%rdx,0x7f563412(%r14)
	and	%rdx,0x7f563412(%r15)
	nop
	and	%rbx,0x7f563412(%rax)
	and	%rbx,0x7f563412(%rcx)
	and	%rbx,0x7f563412(%rdx)
	and	%rbx,0x7f563412(%rbx)
	and	%rbx,0x7f563412(%rsp)
	and	%rbx,0x7f563412(%rbp)
	and	%rbx,0x7f563412(%rsi)
	and	%rbx,0x7f563412(%rdi)
	and	%rbx,0x7f563412(%r8)
	and	%rbx,0x7f563412(%r9)
	and	%rbx,0x7f563412(%r10)
	and	%rbx,0x7f563412(%r11)
	and	%rbx,0x7f563412(%r12)
	and	%rbx,0x7f563412(%r13)
	and	%rbx,0x7f563412(%r14)
	and	%rbx,0x7f563412(%r15)
	nop
	and	%rsp,0x7f563412(%rax)
	and	%rsp,0x7f563412(%rcx)
	and	%rsp,0x7f563412(%rdx)
	and	%rsp,0x7f563412(%rbx)
	and	%rsp,0x7f563412(%rsp)
	and	%rsp,0x7f563412(%rbp)
	and	%rsp,0x7f563412(%rsi)
	and	%rsp,0x7f563412(%rdi)
	and	%rsp,0x7f563412(%r8)
	and	%rsp,0x7f563412(%r9)
	and	%rsp,0x7f563412(%r10)
	and	%rsp,0x7f563412(%r11)
	and	%rsp,0x7f563412(%r12)
	and	%rsp,0x7f563412(%r13)
	and	%rsp,0x7f563412(%r14)
	and	%rsp,0x7f563412(%r15)
	nop
	and	%rbp,0x7f563412(%rax)
	and	%rbp,0x7f563412(%rcx)
	and	%rbp,0x7f563412(%rdx)
	and	%rbp,0x7f563412(%rbx)
	and	%rbp,0x7f563412(%rsp)
	and	%rbp,0x7f563412(%rbp)
	and	%rbp,0x7f563412(%rsi)
	and	%rbp,0x7f563412(%rdi)
	and	%rbp,0x7f563412(%r8)
	and	%rbp,0x7f563412(%r9)
	and	%rbp,0x7f563412(%r10)
	and	%rbp,0x7f563412(%r11)
	and	%rbp,0x7f563412(%r12)
	and	%rbp,0x7f563412(%r13)
	and	%rbp,0x7f563412(%r14)
	and	%rbp,0x7f563412(%r15)
	nop
	and	%rsi,0x7f563412(%rax)
	and	%rsi,0x7f563412(%rcx)
	and	%rsi,0x7f563412(%rdx)
	and	%rsi,0x7f563412(%rbx)
	and	%rsi,0x7f563412(%rsp)
	and	%rsi,0x7f563412(%rbp)
	and	%rsi,0x7f563412(%rsi)
	and	%rsi,0x7f563412(%rdi)
	and	%rsi,0x7f563412(%r8)
	and	%rsi,0x7f563412(%r9)
	and	%rsi,0x7f563412(%r10)
	and	%rsi,0x7f563412(%r11)
	and	%rsi,0x7f563412(%r12)
	and	%rsi,0x7f563412(%r13)
	and	%rsi,0x7f563412(%r14)
	and	%rsi,0x7f563412(%r15)
	nop
	and	%rdi,0x7f563412(%rax)
	and	%rdi,0x7f563412(%rcx)
	and	%rdi,0x7f563412(%rdx)
	and	%rdi,0x7f563412(%rbx)
	and	%rdi,0x7f563412(%rsp)
	and	%rdi,0x7f563412(%rbp)
	and	%rdi,0x7f563412(%rsi)
	and	%rdi,0x7f563412(%rdi)
	and	%rdi,0x7f563412(%r8)
	and	%rdi,0x7f563412(%r9)
	and	%rdi,0x7f563412(%r10)
	and	%rdi,0x7f563412(%r11)
	and	%rdi,0x7f563412(%r12)
	and	%rdi,0x7f563412(%r13)
	and	%rdi,0x7f563412(%r14)
	and	%rdi,0x7f563412(%r15)
	nop
	and	%r8, 0x7f563412(%rax)
	and	%r8, 0x7f563412(%rcx)
	and	%r8, 0x7f563412(%rdx)
	and	%r8, 0x7f563412(%rbx)
	and	%r8, 0x7f563412(%rsp)
	and	%r8, 0x7f563412(%rbp)
	and	%r8, 0x7f563412(%rsi)
	and	%r8, 0x7f563412(%rdi)
	and	%r8, 0x7f563412(%r8)
	and	%r8, 0x7f563412(%r9)
	and	%r8, 0x7f563412(%r10)
	and	%r8, 0x7f563412(%r11)
	and	%r8, 0x7f563412(%r12)
	and	%r8, 0x7f563412(%r13)
	and	%r8, 0x7f563412(%r14)
	and	%r8, 0x7f563412(%r15)
	nop
	and	%r9, 0x7f563412(%rax)
	and	%r9, 0x7f563412(%rcx)
	and	%r9, 0x7f563412(%rdx)
	and	%r9, 0x7f563412(%rbx)
	and	%r9, 0x7f563412(%rsp)
	and	%r9, 0x7f563412(%rbp)
	and	%r9, 0x7f563412(%rsi)
	and	%r9, 0x7f563412(%rdi)
	and	%r9, 0x7f563412(%r8)
	and	%r9, 0x7f563412(%r9)
	and	%r9, 0x7f563412(%r10)
	and	%r9, 0x7f563412(%r11)
	and	%r9, 0x7f563412(%r12)
	and	%r9, 0x7f563412(%r13)
	and	%r9, 0x7f563412(%r14)
	and	%r9, 0x7f563412(%r15)
	nop
	and	%r10,0x7f563412(%rax)
	and	%r10,0x7f563412(%rcx)
	and	%r10,0x7f563412(%rdx)
	and	%r10,0x7f563412(%rbx)
	and	%r10,0x7f563412(%rsp)
	and	%r10,0x7f563412(%rbp)
	and	%r10,0x7f563412(%rsi)
	and	%r10,0x7f563412(%rdi)
	and	%r10,0x7f563412(%r8)
	and	%r10,0x7f563412(%r9)
	and	%r10,0x7f563412(%r10)
	and	%r10,0x7f563412(%r11)
	and	%r10,0x7f563412(%r12)
	and	%r10,0x7f563412(%r13)
	and	%r10,0x7f563412(%r14)
	and	%r10,0x7f563412(%r15)
	nop
	and	%r11,0x7f563412(%rax)
	and	%r11,0x7f563412(%rcx)
	and	%r11,0x7f563412(%rdx)
	and	%r11,0x7f563412(%rbx)
	and	%r11,0x7f563412(%rsp)
	and	%r11,0x7f563412(%rbp)
	and	%r11,0x7f563412(%rsi)
	and	%r11,0x7f563412(%rdi)
	and	%r11,0x7f563412(%r8)
	and	%r11,0x7f563412(%r9)
	and	%r11,0x7f563412(%r10)
	and	%r11,0x7f563412(%r11)
	and	%r11,0x7f563412(%r12)
	and	%r11,0x7f563412(%r13)
	and	%r11,0x7f563412(%r14)
	and	%r11,0x7f563412(%r15)
	nop
	and	%r12,0x7f563412(%rax)
	and	%r12,0x7f563412(%rcx)
	and	%r12,0x7f563412(%rdx)
	and	%r12,0x7f563412(%rbx)
	and	%r12,0x7f563412(%rsp)
	and	%r12,0x7f563412(%rbp)
	and	%r12,0x7f563412(%rsi)
	and	%r12,0x7f563412(%rdi)
	and	%r12,0x7f563412(%r8)
	and	%r12,0x7f563412(%r9)
	and	%r12,0x7f563412(%r10)
	and	%r12,0x7f563412(%r11)
	and	%r12,0x7f563412(%r12)
	and	%r12,0x7f563412(%r13)
	and	%r12,0x7f563412(%r14)
	and	%r12,0x7f563412(%r15)
	nop
	and	%r13,0x7f563412(%rax)
	and	%r13,0x7f563412(%rcx)
	and	%r13,0x7f563412(%rdx)
	and	%r13,0x7f563412(%rbx)
	and	%r13,0x7f563412(%rsp)
	and	%r13,0x7f563412(%rbp)
	and	%r13,0x7f563412(%rsi)
	and	%r13,0x7f563412(%rdi)
	and	%r13,0x7f563412(%r8)
	and	%r13,0x7f563412(%r9)
	and	%r13,0x7f563412(%r10)
	and	%r13,0x7f563412(%r11)
	and	%r13,0x7f563412(%r12)
	and	%r13,0x7f563412(%r13)
	and	%r13,0x7f563412(%r14)
	and	%r13,0x7f563412(%r15)
	nop
	and	%r14,0x7f563412(%rax)
	and	%r14,0x7f563412(%rcx)
	and	%r14,0x7f563412(%rdx)
	and	%r14,0x7f563412(%rbx)
	and	%r14,0x7f563412(%rsp)
	and	%r14,0x7f563412(%rbp)
	and	%r14,0x7f563412(%rsi)
	and	%r14,0x7f563412(%rdi)
	and	%r14,0x7f563412(%r8)
	and	%r14,0x7f563412(%r9)
	and	%r14,0x7f563412(%r10)
	and	%r14,0x7f563412(%r11)
	and	%r14,0x7f563412(%r12)
	and	%r14,0x7f563412(%r13)
	and	%r14,0x7f563412(%r14)
	and	%r14,0x7f563412(%r15)
	nop
	and	%r15,0x7f563412(%rax)
	and	%r15,0x7f563412(%rcx)
	and	%r15,0x7f563412(%rdx)
	and	%r15,0x7f563412(%rbx)
	and	%r15,0x7f563412(%rsp)
	and	%r15,0x7f563412(%rbp)
	and	%r15,0x7f563412(%rsi)
	and	%r15,0x7f563412(%rdi)
	and	%r15,0x7f563412(%r8)
	and	%r15,0x7f563412(%r9)
	and	%r15,0x7f563412(%r10)
	and	%r15,0x7f563412(%r11)
	and	%r15,0x7f563412(%r12)
	and	%r15,0x7f563412(%r13)
	and	%r15,0x7f563412(%r14)
	and	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	AndRegMem
	.type	AndRegMem, @function
AndRegMem:
	.cfi_startproc
	and	(%rax),%rax
	and	(%rax),%rcx
	and	(%rax),%rdx
	and	(%rax),%rbx
	and	(%rax),%rsp
	and	(%rax),%rbp
	and	(%rax),%rsi
	and	(%rax),%rdi
	and	(%rax),%r8
	and	(%rax),%r9
	and	(%rax),%r10
	and	(%rax),%r11
	and	(%rax),%r12
	and	(%rax),%r13
	and	(%rax),%r14
	and	(%rax),%r15
	nop
	and	(%rcx),%rax
	and	(%rcx),%rcx
	and	(%rcx),%rdx
	and	(%rcx),%rbx
	and	(%rcx),%rsp
	and	(%rcx),%rbp
	and	(%rcx),%rsi
	and	(%rcx),%rdi
	and	(%rcx),%r8
	and	(%rcx),%r9
	and	(%rcx),%r10
	and	(%rcx),%r11
	and	(%rcx),%r12
	and	(%rcx),%r13
	and	(%rcx),%r14
	and	(%rcx),%r15
	nop
	and	(%rdx),%rax
	and	(%rdx),%rcx
	and	(%rdx),%rdx
	and	(%rdx),%rbx
	and	(%rdx),%rsp
	and	(%rdx),%rbp
	and	(%rdx),%rsi
	and	(%rdx),%rdi
	and	(%rdx),%r8
	and	(%rdx),%r9
	and	(%rdx),%r10
	and	(%rdx),%r11
	and	(%rdx),%r12
	and	(%rdx),%r13
	and	(%rdx),%r14
	and	(%rdx),%r15
	nop
	and	(%rbx),%rax
	and	(%rbx),%rcx
	and	(%rbx),%rdx
	and	(%rbx),%rbx
	and	(%rbx),%rsp
	and	(%rbx),%rbp
	and	(%rbx),%rsi
	and	(%rbx),%rdi
	and	(%rbx),%r8
	and	(%rbx),%r9
	and	(%rbx),%r10
	and	(%rbx),%r11
	and	(%rbx),%r12
	and	(%rbx),%r13
	and	(%rbx),%r14
	and	(%rbx),%r15
	nop
	and	(%rsp),%rax
	and	(%rsp),%rcx
	and	(%rsp),%rdx
	and	(%rsp),%rbx
	and	(%rsp),%rsp
	and	(%rsp),%rbp
	and	(%rsp),%rsi
	and	(%rsp),%rdi
	and	(%rsp),%r8
	and	(%rsp),%r9
	and	(%rsp),%r10
	and	(%rsp),%r11
	and	(%rsp),%r12
	and	(%rsp),%r13
	and	(%rsp),%r14
	and	(%rsp),%r15
	nop
	and	(%rbp),%rax
	and	(%rbp),%rcx
	and	(%rbp),%rdx
	and	(%rbp),%rbx
	and	(%rbp),%rsp
	and	(%rbp),%rbp
	and	(%rbp),%rsi
	and	(%rbp),%rdi
	and	(%rbp),%r8
	and	(%rbp),%r9
	and	(%rbp),%r10
	and	(%rbp),%r11
	and	(%rbp),%r12
	and	(%rbp),%r13
	and	(%rbp),%r14
	and	(%rbp),%r15
	nop
	and	(%rsi),%rax
	and	(%rsi),%rcx
	and	(%rsi),%rdx
	and	(%rsi),%rbx
	and	(%rsi),%rsp
	and	(%rsi),%rbp
	and	(%rsi),%rsi
	and	(%rsi),%rdi
	and	(%rsi),%r8
	and	(%rsi),%r9
	and	(%rsi),%r10
	and	(%rsi),%r11
	and	(%rsi),%r12
	and	(%rsi),%r13
	and	(%rsi),%r14
	and	(%rsi),%r15
	nop
	and	(%rdi),%rax
	and	(%rdi),%rcx
	and	(%rdi),%rdx
	and	(%rdi),%rbx
	and	(%rdi),%rsp
	and	(%rdi),%rbp
	and	(%rdi),%rsi
	and	(%rdi),%rdi
	and	(%rdi),%r8
	and	(%rdi),%r9
	and	(%rdi),%r10
	and	(%rdi),%r11
	and	(%rdi),%r12
	and	(%rdi),%r13
	and	(%rdi),%r14
	and	(%rdi),%r15
	nop
	and	(%r8), %rax
	and	(%r8), %rcx
	and	(%r8), %rdx
	and	(%r8), %rbx
	and	(%r8), %rsp
	and	(%r8), %rbp
	and	(%r8), %rsi
	and	(%r8), %rdi
	and	(%r8), %r8
	and	(%r8), %r9
	and	(%r8), %r10
	and	(%r8), %r11
	and	(%r8), %r12
	and	(%r8), %r13
	and	(%r8), %r14
	and	(%r8), %r15
	nop
	and	(%r9), %rax
	and	(%r9), %rcx
	and	(%r9), %rdx
	and	(%r9), %rbx
	and	(%r9), %rsp
	and	(%r9), %rbp
	and	(%r9), %rsi
	and	(%r9), %rdi
	and	(%r9), %r8
	and	(%r9), %r9
	and	(%r9), %r10
	and	(%r9), %r11
	and	(%r9), %r12
	and	(%r9), %r13
	and	(%r9), %r14
	and	(%r9), %r15
	nop
	and	(%r10),%rax
	and	(%r10),%rcx
	and	(%r10),%rdx
	and	(%r10),%rbx
	and	(%r10),%rsp
	and	(%r10),%rbp
	and	(%r10),%rsi
	and	(%r10),%rdi
	and	(%r10),%r8
	and	(%r10),%r9
	and	(%r10),%r10
	and	(%r10),%r11
	and	(%r10),%r12
	and	(%r10),%r13
	and	(%r10),%r14
	and	(%r10),%r15
	nop
	and	(%r11),%rax
	and	(%r11),%rcx
	and	(%r11),%rdx
	and	(%r11),%rbx
	and	(%r11),%rsp
	and	(%r11),%rbp
	and	(%r11),%rsi
	and	(%r11),%rdi
	and	(%r11),%r8
	and	(%r11),%r9
	and	(%r11),%r10
	and	(%r11),%r11
	and	(%r11),%r12
	and	(%r11),%r13
	and	(%r11),%r14
	and	(%r11),%r15
	nop
	and	(%r12),%rax
	and	(%r12),%rcx
	and	(%r12),%rdx
	and	(%r12),%rbx
	and	(%r12),%rsp
	and	(%r12),%rbp
	and	(%r12),%rsi
	and	(%r12),%rdi
	and	(%r12),%r8
	and	(%r12),%r9
	and	(%r12),%r10
	and	(%r12),%r11
	and	(%r12),%r12
	and	(%r12),%r13
	and	(%r12),%r14
	and	(%r12),%r15
	nop
	and	(%r13),%rax
	and	(%r13),%rcx
	and	(%r13),%rdx
	and	(%r13),%rbx
	and	(%r13),%rsp
	and	(%r13),%rbp
	and	(%r13),%rsi
	and	(%r13),%rdi
	and	(%r13),%r8
	and	(%r13),%r9
	and	(%r13),%r10
	and	(%r13),%r11
	and	(%r13),%r12
	and	(%r13),%r13
	and	(%r13),%r14
	and	(%r13),%r15
	nop
	and	(%r14),%rax
	and	(%r14),%rcx
	and	(%r14),%rdx
	and	(%r14),%rbx
	and	(%r14),%rsp
	and	(%r14),%rbp
	and	(%r14),%rsi
	and	(%r14),%rdi
	and	(%r14),%r8
	and	(%r14),%r9
	and	(%r14),%r10
	and	(%r14),%r11
	and	(%r14),%r12
	and	(%r14),%r13
	and	(%r14),%r14
	and	(%r14),%r15
	nop
	and	(%r15),%rax
	and	(%r15),%rcx
	and	(%r15),%rdx
	and	(%r15),%rbx
	and	(%r15),%rsp
	and	(%r15),%rbp
	and	(%r15),%rsi
	and	(%r15),%rdi
	and	(%r15),%r8
	and	(%r15),%r9
	and	(%r15),%r10
	and	(%r15),%r11
	and	(%r15),%r12
	and	(%r15),%r13
	and	(%r15),%r14
	and	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndRegMem8
	.type	AndRegMem8, @function
AndRegMem8:
	.cfi_startproc
	and	0x7f(%rax),%rax
	and	0x7f(%rax),%rcx
	and	0x7f(%rax),%rdx
	and	0x7f(%rax),%rbx
	and	0x7f(%rax),%rsp
	and	0x7f(%rax),%rbp
	and	0x7f(%rax),%rsi
	and	0x7f(%rax),%rdi
	and	0x7f(%rax),%r8
	and	0x7f(%rax),%r9
	and	0x7f(%rax),%r10
	and	0x7f(%rax),%r11
	and	0x7f(%rax),%r12
	and	0x7f(%rax),%r13
	and	0x7f(%rax),%r14
	and	0x7f(%rax),%r15
	nop
	and	0x7f(%rcx),%rax
	and	0x7f(%rcx),%rcx
	and	0x7f(%rcx),%rdx
	and	0x7f(%rcx),%rbx
	and	0x7f(%rcx),%rsp
	and	0x7f(%rcx),%rbp
	and	0x7f(%rcx),%rsi
	and	0x7f(%rcx),%rdi
	and	0x7f(%rcx),%r8
	and	0x7f(%rcx),%r9
	and	0x7f(%rcx),%r10
	and	0x7f(%rcx),%r11
	and	0x7f(%rcx),%r12
	and	0x7f(%rcx),%r13
	and	0x7f(%rcx),%r14
	and	0x7f(%rcx),%r15
	nop
	and	0x7f(%rdx),%rax
	and	0x7f(%rdx),%rcx
	and	0x7f(%rdx),%rdx
	and	0x7f(%rdx),%rbx
	and	0x7f(%rdx),%rsp
	and	0x7f(%rdx),%rbp
	and	0x7f(%rdx),%rsi
	and	0x7f(%rdx),%rdi
	and	0x7f(%rdx),%r8
	and	0x7f(%rdx),%r9
	and	0x7f(%rdx),%r10
	and	0x7f(%rdx),%r11
	and	0x7f(%rdx),%r12
	and	0x7f(%rdx),%r13
	and	0x7f(%rdx),%r14
	and	0x7f(%rdx),%r15
	nop
	and	0x7f(%rbx),%rax
	and	0x7f(%rbx),%rcx
	and	0x7f(%rbx),%rdx
	and	0x7f(%rbx),%rbx
	and	0x7f(%rbx),%rsp
	and	0x7f(%rbx),%rbp
	and	0x7f(%rbx),%rsi
	and	0x7f(%rbx),%rdi
	and	0x7f(%rbx),%r8
	and	0x7f(%rbx),%r9
	and	0x7f(%rbx),%r10
	and	0x7f(%rbx),%r11
	and	0x7f(%rbx),%r12
	and	0x7f(%rbx),%r13
	and	0x7f(%rbx),%r14
	and	0x7f(%rbx),%r15
	nop
	and	0x7f(%rsp),%rax
	and	0x7f(%rsp),%rcx
	and	0x7f(%rsp),%rdx
	and	0x7f(%rsp),%rbx
	and	0x7f(%rsp),%rsp
	and	0x7f(%rsp),%rbp
	and	0x7f(%rsp),%rsi
	and	0x7f(%rsp),%rdi
	and	0x7f(%rsp),%r8
	and	0x7f(%rsp),%r9
	and	0x7f(%rsp),%r10
	and	0x7f(%rsp),%r11
	and	0x7f(%rsp),%r12
	and	0x7f(%rsp),%r13
	and	0x7f(%rsp),%r14
	and	0x7f(%rsp),%r15
	nop
	and	0x7f(%rbp),%rax
	and	0x7f(%rbp),%rcx
	and	0x7f(%rbp),%rdx
	and	0x7f(%rbp),%rbx
	and	0x7f(%rbp),%rsp
	and	0x7f(%rbp),%rbp
	and	0x7f(%rbp),%rsi
	and	0x7f(%rbp),%rdi
	and	0x7f(%rbp),%r8
	and	0x7f(%rbp),%r9
	and	0x7f(%rbp),%r10
	and	0x7f(%rbp),%r11
	and	0x7f(%rbp),%r12
	and	0x7f(%rbp),%r13
	and	0x7f(%rbp),%r14
	and	0x7f(%rbp),%r15
	nop
	and	0x7f(%rsi),%rax
	and	0x7f(%rsi),%rcx
	and	0x7f(%rsi),%rdx
	and	0x7f(%rsi),%rbx
	and	0x7f(%rsi),%rsp
	and	0x7f(%rsi),%rbp
	and	0x7f(%rsi),%rsi
	and	0x7f(%rsi),%rdi
	and	0x7f(%rsi),%r8
	and	0x7f(%rsi),%r9
	and	0x7f(%rsi),%r10
	and	0x7f(%rsi),%r11
	and	0x7f(%rsi),%r12
	and	0x7f(%rsi),%r13
	and	0x7f(%rsi),%r14
	and	0x7f(%rsi),%r15
	nop
	and	0x7f(%rdi),%rax
	and	0x7f(%rdi),%rcx
	and	0x7f(%rdi),%rdx
	and	0x7f(%rdi),%rbx
	and	0x7f(%rdi),%rsp
	and	0x7f(%rdi),%rbp
	and	0x7f(%rdi),%rsi
	and	0x7f(%rdi),%rdi
	and	0x7f(%rdi),%r8
	and	0x7f(%rdi),%r9
	and	0x7f(%rdi),%r10
	and	0x7f(%rdi),%r11
	and	0x7f(%rdi),%r12
	and	0x7f(%rdi),%r13
	and	0x7f(%rdi),%r14
	and	0x7f(%rdi),%r15
	nop
	and	0x7f(%r8), %rax
	and	0x7f(%r8), %rcx
	and	0x7f(%r8), %rdx
	and	0x7f(%r8), %rbx
	and	0x7f(%r8), %rsp
	and	0x7f(%r8), %rbp
	and	0x7f(%r8), %rsi
	and	0x7f(%r8), %rdi
	and	0x7f(%r8), %r8
	and	0x7f(%r8), %r9
	and	0x7f(%r8), %r10
	and	0x7f(%r8), %r11
	and	0x7f(%r8), %r12
	and	0x7f(%r8), %r13
	and	0x7f(%r8), %r14
	and	0x7f(%r8), %r15
	nop
	and	0x7f(%r9), %rax
	and	0x7f(%r9), %rcx
	and	0x7f(%r9), %rdx
	and	0x7f(%r9), %rbx
	and	0x7f(%r9), %rsp
	and	0x7f(%r9), %rbp
	and	0x7f(%r9), %rsi
	and	0x7f(%r9), %rdi
	and	0x7f(%r9), %r8
	and	0x7f(%r9), %r9
	and	0x7f(%r9), %r10
	and	0x7f(%r9), %r11
	and	0x7f(%r9), %r12
	and	0x7f(%r9), %r13
	and	0x7f(%r9), %r14
	and	0x7f(%r9), %r15
	nop
	and	0x7f(%r10),%rax
	and	0x7f(%r10),%rcx
	and	0x7f(%r10),%rdx
	and	0x7f(%r10),%rbx
	and	0x7f(%r10),%rsp
	and	0x7f(%r10),%rbp
	and	0x7f(%r10),%rsi
	and	0x7f(%r10),%rdi
	and	0x7f(%r10),%r8
	and	0x7f(%r10),%r9
	and	0x7f(%r10),%r10
	and	0x7f(%r10),%r11
	and	0x7f(%r10),%r12
	and	0x7f(%r10),%r13
	and	0x7f(%r10),%r14
	and	0x7f(%r10),%r15
	nop
	and	0x7f(%r11),%rax
	and	0x7f(%r11),%rcx
	and	0x7f(%r11),%rdx
	and	0x7f(%r11),%rbx
	and	0x7f(%r11),%rsp
	and	0x7f(%r11),%rbp
	and	0x7f(%r11),%rsi
	and	0x7f(%r11),%rdi
	and	0x7f(%r11),%r8
	and	0x7f(%r11),%r9
	and	0x7f(%r11),%r10
	and	0x7f(%r11),%r11
	and	0x7f(%r11),%r12
	and	0x7f(%r11),%r13
	and	0x7f(%r11),%r14
	and	0x7f(%r11),%r15
	nop
	and	0x7f(%r12),%rax
	and	0x7f(%r12),%rcx
	and	0x7f(%r12),%rdx
	and	0x7f(%r12),%rbx
	and	0x7f(%r12),%rsp
	and	0x7f(%r12),%rbp
	and	0x7f(%r12),%rsi
	and	0x7f(%r12),%rdi
	and	0x7f(%r12),%r8
	and	0x7f(%r12),%r9
	and	0x7f(%r12),%r10
	and	0x7f(%r12),%r11
	and	0x7f(%r12),%r12
	and	0x7f(%r12),%r13
	and	0x7f(%r12),%r14
	and	0x7f(%r12),%r15
	nop
	and	0x7f(%r13),%rax
	and	0x7f(%r13),%rcx
	and	0x7f(%r13),%rdx
	and	0x7f(%r13),%rbx
	and	0x7f(%r13),%rsp
	and	0x7f(%r13),%rbp
	and	0x7f(%r13),%rsi
	and	0x7f(%r13),%rdi
	and	0x7f(%r13),%r8
	and	0x7f(%r13),%r9
	and	0x7f(%r13),%r10
	and	0x7f(%r13),%r11
	and	0x7f(%r13),%r12
	and	0x7f(%r13),%r13
	and	0x7f(%r13),%r14
	and	0x7f(%r13),%r15
	nop
	and	0x7f(%r14),%rax
	and	0x7f(%r14),%rcx
	and	0x7f(%r14),%rdx
	and	0x7f(%r14),%rbx
	and	0x7f(%r14),%rsp
	and	0x7f(%r14),%rbp
	and	0x7f(%r14),%rsi
	and	0x7f(%r14),%rdi
	and	0x7f(%r14),%r8
	and	0x7f(%r14),%r9
	and	0x7f(%r14),%r10
	and	0x7f(%r14),%r11
	and	0x7f(%r14),%r12
	and	0x7f(%r14),%r13
	and	0x7f(%r14),%r14
	and	0x7f(%r14),%r15
	nop
	and	0x7f(%r15),%rax
	and	0x7f(%r15),%rcx
	and	0x7f(%r15),%rdx
	and	0x7f(%r15),%rbx
	and	0x7f(%r15),%rsp
	and	0x7f(%r15),%rbp
	and	0x7f(%r15),%rsi
	and	0x7f(%r15),%rdi
	and	0x7f(%r15),%r8
	and	0x7f(%r15),%r9
	and	0x7f(%r15),%r10
	and	0x7f(%r15),%r11
	and	0x7f(%r15),%r12
	and	0x7f(%r15),%r13
	and	0x7f(%r15),%r14
	and	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndRegMem32
	.type	AndRegMem32, @function
AndRegMem32:
	.cfi_startproc
	and	0x7f563412(%rax),%rax
	and	0x7f563412(%rax),%rcx
	and	0x7f563412(%rax),%rdx
	and	0x7f563412(%rax),%rbx
	and	0x7f563412(%rax),%rsp
	and	0x7f563412(%rax),%rbp
	and	0x7f563412(%rax),%rsi
	and	0x7f563412(%rax),%rdi
	and	0x7f563412(%rax),%r8
	and	0x7f563412(%rax),%r9
	and	0x7f563412(%rax),%r10
	and	0x7f563412(%rax),%r11
	and	0x7f563412(%rax),%r12
	and	0x7f563412(%rax),%r13
	and	0x7f563412(%rax),%r14
	and	0x7f563412(%rax),%r15
	nop
	and	0x7f563412(%rcx),%rax
	and	0x7f563412(%rcx),%rcx
	and	0x7f563412(%rcx),%rdx
	and	0x7f563412(%rcx),%rbx
	and	0x7f563412(%rcx),%rsp
	and	0x7f563412(%rcx),%rbp
	and	0x7f563412(%rcx),%rsi
	and	0x7f563412(%rcx),%rdi
	and	0x7f563412(%rcx),%r8
	and	0x7f563412(%rcx),%r9
	and	0x7f563412(%rcx),%r10
	and	0x7f563412(%rcx),%r11
	and	0x7f563412(%rcx),%r12
	and	0x7f563412(%rcx),%r13
	and	0x7f563412(%rcx),%r14
	and	0x7f563412(%rcx),%r15
	nop
	and	0x7f563412(%rdx),%rax
	and	0x7f563412(%rdx),%rcx
	and	0x7f563412(%rdx),%rdx
	and	0x7f563412(%rdx),%rbx
	and	0x7f563412(%rdx),%rsp
	and	0x7f563412(%rdx),%rbp
	and	0x7f563412(%rdx),%rsi
	and	0x7f563412(%rdx),%rdi
	and	0x7f563412(%rdx),%r8
	and	0x7f563412(%rdx),%r9
	and	0x7f563412(%rdx),%r10
	and	0x7f563412(%rdx),%r11
	and	0x7f563412(%rdx),%r12
	and	0x7f563412(%rdx),%r13
	and	0x7f563412(%rdx),%r14
	and	0x7f563412(%rdx),%r15
	nop
	and	0x7f563412(%rbx),%rax
	and	0x7f563412(%rbx),%rcx
	and	0x7f563412(%rbx),%rdx
	and	0x7f563412(%rbx),%rbx
	and	0x7f563412(%rbx),%rsp
	and	0x7f563412(%rbx),%rbp
	and	0x7f563412(%rbx),%rsi
	and	0x7f563412(%rbx),%rdi
	and	0x7f563412(%rbx),%r8
	and	0x7f563412(%rbx),%r9
	and	0x7f563412(%rbx),%r10
	and	0x7f563412(%rbx),%r11
	and	0x7f563412(%rbx),%r12
	and	0x7f563412(%rbx),%r13
	and	0x7f563412(%rbx),%r14
	and	0x7f563412(%rbx),%r15
	nop
	and	0x7f563412(%rsp),%rax
	and	0x7f563412(%rsp),%rcx
	and	0x7f563412(%rsp),%rdx
	and	0x7f563412(%rsp),%rbx
	and	0x7f563412(%rsp),%rsp
	and	0x7f563412(%rsp),%rbp
	and	0x7f563412(%rsp),%rsi
	and	0x7f563412(%rsp),%rdi
	and	0x7f563412(%rsp),%r8
	and	0x7f563412(%rsp),%r9
	and	0x7f563412(%rsp),%r10
	and	0x7f563412(%rsp),%r11
	and	0x7f563412(%rsp),%r12
	and	0x7f563412(%rsp),%r13
	and	0x7f563412(%rsp),%r14
	and	0x7f563412(%rsp),%r15
	nop
	and	0x7f563412(%rbp),%rax
	and	0x7f563412(%rbp),%rcx
	and	0x7f563412(%rbp),%rdx
	and	0x7f563412(%rbp),%rbx
	and	0x7f563412(%rbp),%rsp
	and	0x7f563412(%rbp),%rbp
	and	0x7f563412(%rbp),%rsi
	and	0x7f563412(%rbp),%rdi
	and	0x7f563412(%rbp),%r8
	and	0x7f563412(%rbp),%r9
	and	0x7f563412(%rbp),%r10
	and	0x7f563412(%rbp),%r11
	and	0x7f563412(%rbp),%r12
	and	0x7f563412(%rbp),%r13
	and	0x7f563412(%rbp),%r14
	and	0x7f563412(%rbp),%r15
	nop
	and	0x7f563412(%rsi),%rax
	and	0x7f563412(%rsi),%rcx
	and	0x7f563412(%rsi),%rdx
	and	0x7f563412(%rsi),%rbx
	and	0x7f563412(%rsi),%rsp
	and	0x7f563412(%rsi),%rbp
	and	0x7f563412(%rsi),%rsi
	and	0x7f563412(%rsi),%rdi
	and	0x7f563412(%rsi),%r8
	and	0x7f563412(%rsi),%r9
	and	0x7f563412(%rsi),%r10
	and	0x7f563412(%rsi),%r11
	and	0x7f563412(%rsi),%r12
	and	0x7f563412(%rsi),%r13
	and	0x7f563412(%rsi),%r14
	and	0x7f563412(%rsi),%r15
	nop
	and	0x7f563412(%rdi),%rax
	and	0x7f563412(%rdi),%rcx
	and	0x7f563412(%rdi),%rdx
	and	0x7f563412(%rdi),%rbx
	and	0x7f563412(%rdi),%rsp
	and	0x7f563412(%rdi),%rbp
	and	0x7f563412(%rdi),%rsi
	and	0x7f563412(%rdi),%rdi
	and	0x7f563412(%rdi),%r8
	and	0x7f563412(%rdi),%r9
	and	0x7f563412(%rdi),%r10
	and	0x7f563412(%rdi),%r11
	and	0x7f563412(%rdi),%r12
	and	0x7f563412(%rdi),%r13
	and	0x7f563412(%rdi),%r14
	and	0x7f563412(%rdi),%r15
	nop
	and	0x7f563412(%r8), %rax
	and	0x7f563412(%r8), %rcx
	and	0x7f563412(%r8), %rdx
	and	0x7f563412(%r8), %rbx
	and	0x7f563412(%r8), %rsp
	and	0x7f563412(%r8), %rbp
	and	0x7f563412(%r8), %rsi
	and	0x7f563412(%r8), %rdi
	and	0x7f563412(%r8), %r8
	and	0x7f563412(%r8), %r9
	and	0x7f563412(%r8), %r10
	and	0x7f563412(%r8), %r11
	and	0x7f563412(%r8), %r12
	and	0x7f563412(%r8), %r13
	and	0x7f563412(%r8), %r14
	and	0x7f563412(%r8), %r15
	nop
	and	0x7f563412(%r9), %rax
	and	0x7f563412(%r9), %rcx
	and	0x7f563412(%r9), %rdx
	and	0x7f563412(%r9), %rbx
	and	0x7f563412(%r9), %rsp
	and	0x7f563412(%r9), %rbp
	and	0x7f563412(%r9), %rsi
	and	0x7f563412(%r9), %rdi
	and	0x7f563412(%r9), %r8
	and	0x7f563412(%r9), %r9
	and	0x7f563412(%r9), %r10
	and	0x7f563412(%r9), %r11
	and	0x7f563412(%r9), %r12
	and	0x7f563412(%r9), %r13
	and	0x7f563412(%r9), %r14
	and	0x7f563412(%r9), %r15
	nop
	and	0x7f563412(%r10),%rax
	and	0x7f563412(%r10),%rcx
	and	0x7f563412(%r10),%rdx
	and	0x7f563412(%r10),%rbx
	and	0x7f563412(%r10),%rsp
	and	0x7f563412(%r10),%rbp
	and	0x7f563412(%r10),%rsi
	and	0x7f563412(%r10),%rdi
	and	0x7f563412(%r10),%r8
	and	0x7f563412(%r10),%r9
	and	0x7f563412(%r10),%r10
	and	0x7f563412(%r10),%r11
	and	0x7f563412(%r10),%r12
	and	0x7f563412(%r10),%r13
	and	0x7f563412(%r10),%r14
	and	0x7f563412(%r10),%r15
	nop
	and	0x7f563412(%r11),%rax
	and	0x7f563412(%r11),%rcx
	and	0x7f563412(%r11),%rdx
	and	0x7f563412(%r11),%rbx
	and	0x7f563412(%r11),%rsp
	and	0x7f563412(%r11),%rbp
	and	0x7f563412(%r11),%rsi
	and	0x7f563412(%r11),%rdi
	and	0x7f563412(%r11),%r8
	and	0x7f563412(%r11),%r9
	and	0x7f563412(%r11),%r10
	and	0x7f563412(%r11),%r11
	and	0x7f563412(%r11),%r12
	and	0x7f563412(%r11),%r13
	and	0x7f563412(%r11),%r14
	and	0x7f563412(%r11),%r15
	nop
	and	0x7f563412(%r12),%rax
	and	0x7f563412(%r12),%rcx
	and	0x7f563412(%r12),%rdx
	and	0x7f563412(%r12),%rbx
	and	0x7f563412(%r12),%rsp
	and	0x7f563412(%r12),%rbp
	and	0x7f563412(%r12),%rsi
	and	0x7f563412(%r12),%rdi
	and	0x7f563412(%r12),%r8
	and	0x7f563412(%r12),%r9
	and	0x7f563412(%r12),%r10
	and	0x7f563412(%r12),%r11
	and	0x7f563412(%r12),%r12
	and	0x7f563412(%r12),%r13
	and	0x7f563412(%r12),%r14
	and	0x7f563412(%r12),%r15
	nop
	and	0x7f563412(%r13),%rax
	and	0x7f563412(%r13),%rcx
	and	0x7f563412(%r13),%rdx
	and	0x7f563412(%r13),%rbx
	and	0x7f563412(%r13),%rsp
	and	0x7f563412(%r13),%rbp
	and	0x7f563412(%r13),%rsi
	and	0x7f563412(%r13),%rdi
	and	0x7f563412(%r13),%r8
	and	0x7f563412(%r13),%r9
	and	0x7f563412(%r13),%r10
	and	0x7f563412(%r13),%r11
	and	0x7f563412(%r13),%r12
	and	0x7f563412(%r13),%r13
	and	0x7f563412(%r13),%r14
	and	0x7f563412(%r13),%r15
	nop
	and	0x7f563412(%r14),%rax
	and	0x7f563412(%r14),%rcx
	and	0x7f563412(%r14),%rdx
	and	0x7f563412(%r14),%rbx
	and	0x7f563412(%r14),%rsp
	and	0x7f563412(%r14),%rbp
	and	0x7f563412(%r14),%rsi
	and	0x7f563412(%r14),%rdi
	and	0x7f563412(%r14),%r8
	and	0x7f563412(%r14),%r9
	and	0x7f563412(%r14),%r10
	and	0x7f563412(%r14),%r11
	and	0x7f563412(%r14),%r12
	and	0x7f563412(%r14),%r13
	and	0x7f563412(%r14),%r14
	and	0x7f563412(%r14),%r15
	nop
	and	0x7f563412(%r15),%rax
	and	0x7f563412(%r15),%rcx
	and	0x7f563412(%r15),%rdx
	and	0x7f563412(%r15),%rbx
	and	0x7f563412(%r15),%rsp
	and	0x7f563412(%r15),%rbp
	and	0x7f563412(%r15),%rsi
	and	0x7f563412(%r15),%rdi
	and	0x7f563412(%r15),%r8
	and	0x7f563412(%r15),%r9
	and	0x7f563412(%r15),%r10
	and	0x7f563412(%r15),%r11
	and	0x7f563412(%r15),%r12
	and	0x7f563412(%r15),%r13
	and	0x7f563412(%r15),%r14
	and	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


