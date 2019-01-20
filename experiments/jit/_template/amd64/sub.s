	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Sub_s32
	.type	Sub_s32, @function
Sub_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// sub	$-0x55667788,%rax
	sub	$-0x55667788,%rcx
	sub	$-0x55667788,%rdx
	sub	$-0x55667788,%rbx
	sub	$-0x55667788,%rsp
	sub	$-0x55667788,%rbp
	sub	$-0x55667788,%rsi
	sub	$-0x55667788,%rdi
	sub	$-0x55667788,%r8
	sub	$-0x55667788,%r9
	sub	$-0x55667788,%r10
	sub	$-0x55667788,%r11
	sub	$-0x55667788,%r12
	sub	$-0x55667788,%r13
	sub	$-0x55667788,%r14
	sub	$-0x55667788,%r15
	.cfi_endproc


	.p2align 4,,15
	.globl	Sub
	.type	Sub, @function
Sub:
	.cfi_startproc
	sub	%rax,%rax
	sub	%rax,%rcx
	sub	%rax,%rdx
	sub	%rax,%rbx
	sub	%rax,%rsp
	sub	%rax,%rbp
	sub	%rax,%rsi
	sub	%rax,%rdi
	sub	%rax,%r8
	sub	%rax,%r9
	sub	%rax,%r10
	sub	%rax,%r11
	sub	%rax,%r12
	sub	%rax,%r13
	sub	%rax,%r14
	sub	%rax,%r15
	nop
	sub	%rcx,%rax
	sub	%rcx,%rcx
	sub	%rcx,%rdx
	sub	%rcx,%rbx
	sub	%rcx,%rsp
	sub	%rcx,%rbp
	sub	%rcx,%rsi
	sub	%rcx,%rdi
	sub	%rcx,%r8
	sub	%rcx,%r9
	sub	%rcx,%r10
	sub	%rcx,%r11
	sub	%rcx,%r12
	sub	%rcx,%r13
	sub	%rcx,%r14
	sub	%rcx,%r15
	nop
	sub	%rdx,%rax
	sub	%rdx,%rcx
	sub	%rdx,%rdx
	sub	%rdx,%rbx
	sub	%rdx,%rsp
	sub	%rdx,%rbp
	sub	%rdx,%rsi
	sub	%rdx,%rdi
	sub	%rdx,%r8
	sub	%rdx,%r9
	sub	%rdx,%r10
	sub	%rdx,%r11
	sub	%rdx,%r12
	sub	%rdx,%r13
	sub	%rdx,%r14
	sub	%rdx,%r15
	nop
	sub	%rbx,%rax
	sub	%rbx,%rcx
	sub	%rbx,%rdx
	sub	%rbx,%rbx
	sub	%rbx,%rsp
	sub	%rbx,%rbp
	sub	%rbx,%rsi
	sub	%rbx,%rdi
	sub	%rbx,%r8
	sub	%rbx,%r9
	sub	%rbx,%r10
	sub	%rbx,%r11
	sub	%rbx,%r12
	sub	%rbx,%r13
	sub	%rbx,%r14
	sub	%rbx,%r15
	nop
	sub	%rsp,%rax
	sub	%rsp,%rcx
	sub	%rsp,%rdx
	sub	%rsp,%rbx
	sub	%rsp,%rsp
	sub	%rsp,%rbp
	sub	%rsp,%rsi
	sub	%rsp,%rdi
	sub	%rsp,%r8
	sub	%rsp,%r9
	sub	%rsp,%r10
	sub	%rsp,%r11
	sub	%rsp,%r12
	sub	%rsp,%r13
	sub	%rsp,%r14
	sub	%rsp,%r15
	nop
	sub	%rbp,%rax
	sub	%rbp,%rcx
	sub	%rbp,%rdx
	sub	%rbp,%rbx
	sub	%rbp,%rsp
	sub	%rbp,%rbp
	sub	%rbp,%rsi
	sub	%rbp,%rdi
	sub	%rbp,%r8
	sub	%rbp,%r9
	sub	%rbp,%r10
	sub	%rbp,%r11
	sub	%rbp,%r12
	sub	%rbp,%r13
	sub	%rbp,%r14
	sub	%rbp,%r15
	nop
	sub	%rsi,%rax
	sub	%rsi,%rcx
	sub	%rsi,%rdx
	sub	%rsi,%rbx
	sub	%rsi,%rsp
	sub	%rsi,%rbp
	sub	%rsi,%rsi
	sub	%rsi,%rdi
	sub	%rsi,%r8
	sub	%rsi,%r9
	sub	%rsi,%r10
	sub	%rsi,%r11
	sub	%rsi,%r12
	sub	%rsi,%r13
	sub	%rsi,%r14
	sub	%rsi,%r15
	nop
	sub	%rdi,%rax
	sub	%rdi,%rcx
	sub	%rdi,%rdx
	sub	%rdi,%rbx
	sub	%rdi,%rsp
	sub	%rdi,%rbp
	sub	%rdi,%rsi
	sub	%rdi,%rdi
	sub	%rdi,%r8
	sub	%rdi,%r9
	sub	%rdi,%r10
	sub	%rdi,%r11
	sub	%rdi,%r12
	sub	%rdi,%r13
	sub	%rdi,%r14
	sub	%rdi,%r15
	nop
	sub	%r8, %rax
	sub	%r8, %rcx
	sub	%r8, %rdx
	sub	%r8, %rbx
	sub	%r8, %rsp
	sub	%r8, %rbp
	sub	%r8, %rsi
	sub	%r8, %rdi
	sub	%r8, %r8
	sub	%r8, %r9
	sub	%r8, %r10
	sub	%r8, %r11
	sub	%r8, %r12
	sub	%r8, %r13
	sub	%r8, %r14
	sub	%r8, %r15
	nop
	sub	%r9, %rax
	sub	%r9, %rcx
	sub	%r9, %rdx
	sub	%r9, %rbx
	sub	%r9, %rsp
	sub	%r9, %rbp
	sub	%r9, %rsi
	sub	%r9, %rdi
	sub	%r9, %r8
	sub	%r9, %r9
	sub	%r9, %r10
	sub	%r9, %r11
	sub	%r9, %r12
	sub	%r9, %r13
	sub	%r9, %r14
	sub	%r9, %r15
	nop
	sub	%r10,%rax
	sub	%r10,%rcx
	sub	%r10,%rdx
	sub	%r10,%rbx
	sub	%r10,%rsp
	sub	%r10,%rbp
	sub	%r10,%rsi
	sub	%r10,%rdi
	sub	%r10,%r8
	sub	%r10,%r9
	sub	%r10,%r10
	sub	%r10,%r11
	sub	%r10,%r12
	sub	%r10,%r13
	sub	%r10,%r14
	sub	%r10,%r15
	nop
	sub	%r11,%rax
	sub	%r11,%rcx
	sub	%r11,%rdx
	sub	%r11,%rbx
	sub	%r11,%rsp
	sub	%r11,%rbp
	sub	%r11,%rsi
	sub	%r11,%rdi
	sub	%r11,%r8
	sub	%r11,%r9
	sub	%r11,%r10
	sub	%r11,%r11
	sub	%r11,%r12
	sub	%r11,%r13
	sub	%r11,%r14
	sub	%r11,%r15
	nop
	sub	%r12,%rax
	sub	%r12,%rcx
	sub	%r12,%rdx
	sub	%r12,%rbx
	sub	%r12,%rsp
	sub	%r12,%rbp
	sub	%r12,%rsi
	sub	%r12,%rdi
	sub	%r12,%r8
	sub	%r12,%r9
	sub	%r12,%r10
	sub	%r12,%r11
	sub	%r12,%r12
	sub	%r12,%r13
	sub	%r12,%r14
	sub	%r12,%r15
	nop
	sub	%r13,%rax
	sub	%r13,%rcx
	sub	%r13,%rdx
	sub	%r13,%rbx
	sub	%r13,%rsp
	sub	%r13,%rbp
	sub	%r13,%rsi
	sub	%r13,%rdi
	sub	%r13,%r8
	sub	%r13,%r9
	sub	%r13,%r10
	sub	%r13,%r11
	sub	%r13,%r12
	sub	%r13,%r13
	sub	%r13,%r14
	sub	%r13,%r15
	nop
	sub	%r14,%rax
	sub	%r14,%rcx
	sub	%r14,%rdx
	sub	%r14,%rbx
	sub	%r14,%rsp
	sub	%r14,%rbp
	sub	%r14,%rsi
	sub	%r14,%rdi
	sub	%r14,%r8
	sub	%r14,%r9
	sub	%r14,%r10
	sub	%r14,%r11
	sub	%r14,%r12
	sub	%r14,%r13
	sub	%r14,%r14
	sub	%r14,%r15
	nop
	sub	%r15,%rax
	sub	%r15,%rcx
	sub	%r15,%rdx
	sub	%r15,%rbx
	sub	%r15,%rsp
	sub	%r15,%rbp
	sub	%r15,%rsi
	sub	%r15,%rdi
	sub	%r15,%r8
	sub	%r15,%r9
	sub	%r15,%r10
	sub	%r15,%r11
	sub	%r15,%r12
	sub	%r15,%r13
	sub	%r15,%r14
	sub	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	SubMemReg
	.type	SubMemReg, @function
SubMemReg:
	.cfi_startproc
	sub	%rax,(%rax)
	sub	%rax,(%rcx)
	sub	%rax,(%rdx)
	sub	%rax,(%rbx)
	sub	%rax,(%rsp)
	sub	%rax,(%rbp)
	sub	%rax,(%rsi)
	sub	%rax,(%rdi)
	sub	%rax,(%r8)
	sub	%rax,(%r9)
	sub	%rax,(%r10)
	sub	%rax,(%r11)
	sub	%rax,(%r12)
	sub	%rax,(%r13)
	sub	%rax,(%r14)
	sub	%rax,(%r15)
	nop
	sub	%rcx,(%rax)
	sub	%rcx,(%rcx)
	sub	%rcx,(%rdx)
	sub	%rcx,(%rbx)
	sub	%rcx,(%rsp)
	sub	%rcx,(%rbp)
	sub	%rcx,(%rsi)
	sub	%rcx,(%rdi)
	sub	%rcx,(%r8)
	sub	%rcx,(%r9)
	sub	%rcx,(%r10)
	sub	%rcx,(%r11)
	sub	%rcx,(%r12)
	sub	%rcx,(%r13)
	sub	%rcx,(%r14)
	sub	%rcx,(%r15)
	nop
	sub	%rdx,(%rax)
	sub	%rdx,(%rcx)
	sub	%rdx,(%rdx)
	sub	%rdx,(%rbx)
	sub	%rdx,(%rsp)
	sub	%rdx,(%rbp)
	sub	%rdx,(%rsi)
	sub	%rdx,(%rdi)
	sub	%rdx,(%r8)
	sub	%rdx,(%r9)
	sub	%rdx,(%r10)
	sub	%rdx,(%r11)
	sub	%rdx,(%r12)
	sub	%rdx,(%r13)
	sub	%rdx,(%r14)
	sub	%rdx,(%r15)
	nop
	sub	%rbx,(%rax)
	sub	%rbx,(%rcx)
	sub	%rbx,(%rdx)
	sub	%rbx,(%rbx)
	sub	%rbx,(%rsp)
	sub	%rbx,(%rbp)
	sub	%rbx,(%rsi)
	sub	%rbx,(%rdi)
	sub	%rbx,(%r8)
	sub	%rbx,(%r9)
	sub	%rbx,(%r10)
	sub	%rbx,(%r11)
	sub	%rbx,(%r12)
	sub	%rbx,(%r13)
	sub	%rbx,(%r14)
	sub	%rbx,(%r15)
	nop
	sub	%rsp,(%rax)
	sub	%rsp,(%rcx)
	sub	%rsp,(%rdx)
	sub	%rsp,(%rbx)
	sub	%rsp,(%rsp)
	sub	%rsp,(%rbp)
	sub	%rsp,(%rsi)
	sub	%rsp,(%rdi)
	sub	%rsp,(%r8)
	sub	%rsp,(%r9)
	sub	%rsp,(%r10)
	sub	%rsp,(%r11)
	sub	%rsp,(%r12)
	sub	%rsp,(%r13)
	sub	%rsp,(%r14)
	sub	%rsp,(%r15)
	nop
	sub	%rbp,(%rax)
	sub	%rbp,(%rcx)
	sub	%rbp,(%rdx)
	sub	%rbp,(%rbx)
	sub	%rbp,(%rsp)
	sub	%rbp,(%rbp)
	sub	%rbp,(%rsi)
	sub	%rbp,(%rdi)
	sub	%rbp,(%r8)
	sub	%rbp,(%r9)
	sub	%rbp,(%r10)
	sub	%rbp,(%r11)
	sub	%rbp,(%r12)
	sub	%rbp,(%r13)
	sub	%rbp,(%r14)
	sub	%rbp,(%r15)
	nop
	sub	%rsi,(%rax)
	sub	%rsi,(%rcx)
	sub	%rsi,(%rdx)
	sub	%rsi,(%rbx)
	sub	%rsi,(%rsp)
	sub	%rsi,(%rbp)
	sub	%rsi,(%rsi)
	sub	%rsi,(%rdi)
	sub	%rsi,(%r8)
	sub	%rsi,(%r9)
	sub	%rsi,(%r10)
	sub	%rsi,(%r11)
	sub	%rsi,(%r12)
	sub	%rsi,(%r13)
	sub	%rsi,(%r14)
	sub	%rsi,(%r15)
	nop
	sub	%rdi,(%rax)
	sub	%rdi,(%rcx)
	sub	%rdi,(%rdx)
	sub	%rdi,(%rbx)
	sub	%rdi,(%rsp)
	sub	%rdi,(%rbp)
	sub	%rdi,(%rsi)
	sub	%rdi,(%rdi)
	sub	%rdi,(%r8)
	sub	%rdi,(%r9)
	sub	%rdi,(%r10)
	sub	%rdi,(%r11)
	sub	%rdi,(%r12)
	sub	%rdi,(%r13)
	sub	%rdi,(%r14)
	sub	%rdi,(%r15)
	nop
	sub	%r8, (%rax)
	sub	%r8, (%rcx)
	sub	%r8, (%rdx)
	sub	%r8, (%rbx)
	sub	%r8, (%rsp)
	sub	%r8, (%rbp)
	sub	%r8, (%rsi)
	sub	%r8, (%rdi)
	sub	%r8, (%r8)
	sub	%r8, (%r9)
	sub	%r8, (%r10)
	sub	%r8, (%r11)
	sub	%r8, (%r12)
	sub	%r8, (%r13)
	sub	%r8, (%r14)
	sub	%r8, (%r15)
	nop
	sub	%r9, (%rax)
	sub	%r9, (%rcx)
	sub	%r9, (%rdx)
	sub	%r9, (%rbx)
	sub	%r9, (%rsp)
	sub	%r9, (%rbp)
	sub	%r9, (%rsi)
	sub	%r9, (%rdi)
	sub	%r9, (%r8)
	sub	%r9, (%r9)
	sub	%r9, (%r10)
	sub	%r9, (%r11)
	sub	%r9, (%r12)
	sub	%r9, (%r13)
	sub	%r9, (%r14)
	sub	%r9, (%r15)
	nop
	sub	%r10,(%rax)
	sub	%r10,(%rcx)
	sub	%r10,(%rdx)
	sub	%r10,(%rbx)
	sub	%r10,(%rsp)
	sub	%r10,(%rbp)
	sub	%r10,(%rsi)
	sub	%r10,(%rdi)
	sub	%r10,(%r8)
	sub	%r10,(%r9)
	sub	%r10,(%r10)
	sub	%r10,(%r11)
	sub	%r10,(%r12)
	sub	%r10,(%r13)
	sub	%r10,(%r14)
	sub	%r10,(%r15)
	nop
	sub	%r11,(%rax)
	sub	%r11,(%rcx)
	sub	%r11,(%rdx)
	sub	%r11,(%rbx)
	sub	%r11,(%rsp)
	sub	%r11,(%rbp)
	sub	%r11,(%rsi)
	sub	%r11,(%rdi)
	sub	%r11,(%r8)
	sub	%r11,(%r9)
	sub	%r11,(%r10)
	sub	%r11,(%r11)
	sub	%r11,(%r12)
	sub	%r11,(%r13)
	sub	%r11,(%r14)
	sub	%r11,(%r15)
	nop
	sub	%r12,(%rax)
	sub	%r12,(%rcx)
	sub	%r12,(%rdx)
	sub	%r12,(%rbx)
	sub	%r12,(%rsp)
	sub	%r12,(%rbp)
	sub	%r12,(%rsi)
	sub	%r12,(%rdi)
	sub	%r12,(%r8)
	sub	%r12,(%r9)
	sub	%r12,(%r10)
	sub	%r12,(%r11)
	sub	%r12,(%r12)
	sub	%r12,(%r13)
	sub	%r12,(%r14)
	sub	%r12,(%r15)
	nop
	sub	%r13,(%rax)
	sub	%r13,(%rcx)
	sub	%r13,(%rdx)
	sub	%r13,(%rbx)
	sub	%r13,(%rsp)
	sub	%r13,(%rbp)
	sub	%r13,(%rsi)
	sub	%r13,(%rdi)
	sub	%r13,(%r8)
	sub	%r13,(%r9)
	sub	%r13,(%r10)
	sub	%r13,(%r11)
	sub	%r13,(%r12)
	sub	%r13,(%r13)
	sub	%r13,(%r14)
	sub	%r13,(%r15)
	nop
	sub	%r14,(%rax)
	sub	%r14,(%rcx)
	sub	%r14,(%rdx)
	sub	%r14,(%rbx)
	sub	%r14,(%rsp)
	sub	%r14,(%rbp)
	sub	%r14,(%rsi)
	sub	%r14,(%rdi)
	sub	%r14,(%r8)
	sub	%r14,(%r9)
	sub	%r14,(%r10)
	sub	%r14,(%r11)
	sub	%r14,(%r12)
	sub	%r14,(%r13)
	sub	%r14,(%r14)
	sub	%r14,(%r15)
	nop
	sub	%r15,(%rax)
	sub	%r15,(%rcx)
	sub	%r15,(%rdx)
	sub	%r15,(%rbx)
	sub	%r15,(%rsp)
	sub	%r15,(%rbp)
	sub	%r15,(%rsi)
	sub	%r15,(%rdi)
	sub	%r15,(%r8)
	sub	%r15,(%r9)
	sub	%r15,(%r10)
	sub	%r15,(%r11)
	sub	%r15,(%r12)
	sub	%r15,(%r13)
	sub	%r15,(%r14)
	sub	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	SubMem8Reg
	.type	SubMem8Reg, @function
SubMem8Reg:
	.cfi_startproc
	sub	%rax,0x7f(%rax)
	sub	%rax,0x7f(%rcx)
	sub	%rax,0x7f(%rdx)
	sub	%rax,0x7f(%rbx)
	sub	%rax,0x7f(%rsp)
	sub	%rax,0x7f(%rbp)
	sub	%rax,0x7f(%rsi)
	sub	%rax,0x7f(%rdi)
	sub	%rax,0x7f(%r8)
	sub	%rax,0x7f(%r9)
	sub	%rax,0x7f(%r10)
	sub	%rax,0x7f(%r11)
	sub	%rax,0x7f(%r12)
	sub	%rax,0x7f(%r13)
	sub	%rax,0x7f(%r14)
	sub	%rax,0x7f(%r15)
	nop
	sub	%rcx,0x7f(%rax)
	sub	%rcx,0x7f(%rcx)
	sub	%rcx,0x7f(%rdx)
	sub	%rcx,0x7f(%rbx)
	sub	%rcx,0x7f(%rsp)
	sub	%rcx,0x7f(%rbp)
	sub	%rcx,0x7f(%rsi)
	sub	%rcx,0x7f(%rdi)
	sub	%rcx,0x7f(%r8)
	sub	%rcx,0x7f(%r9)
	sub	%rcx,0x7f(%r10)
	sub	%rcx,0x7f(%r11)
	sub	%rcx,0x7f(%r12)
	sub	%rcx,0x7f(%r13)
	sub	%rcx,0x7f(%r14)
	sub	%rcx,0x7f(%r15)
	nop
	sub	%rdx,0x7f(%rax)
	sub	%rdx,0x7f(%rcx)
	sub	%rdx,0x7f(%rdx)
	sub	%rdx,0x7f(%rbx)
	sub	%rdx,0x7f(%rsp)
	sub	%rdx,0x7f(%rbp)
	sub	%rdx,0x7f(%rsi)
	sub	%rdx,0x7f(%rdi)
	sub	%rdx,0x7f(%r8)
	sub	%rdx,0x7f(%r9)
	sub	%rdx,0x7f(%r10)
	sub	%rdx,0x7f(%r11)
	sub	%rdx,0x7f(%r12)
	sub	%rdx,0x7f(%r13)
	sub	%rdx,0x7f(%r14)
	sub	%rdx,0x7f(%r15)
	nop
	sub	%rbx,0x7f(%rax)
	sub	%rbx,0x7f(%rcx)
	sub	%rbx,0x7f(%rdx)
	sub	%rbx,0x7f(%rbx)
	sub	%rbx,0x7f(%rsp)
	sub	%rbx,0x7f(%rbp)
	sub	%rbx,0x7f(%rsi)
	sub	%rbx,0x7f(%rdi)
	sub	%rbx,0x7f(%r8)
	sub	%rbx,0x7f(%r9)
	sub	%rbx,0x7f(%r10)
	sub	%rbx,0x7f(%r11)
	sub	%rbx,0x7f(%r12)
	sub	%rbx,0x7f(%r13)
	sub	%rbx,0x7f(%r14)
	sub	%rbx,0x7f(%r15)
	nop
	sub	%rsp,0x7f(%rax)
	sub	%rsp,0x7f(%rcx)
	sub	%rsp,0x7f(%rdx)
	sub	%rsp,0x7f(%rbx)
	sub	%rsp,0x7f(%rsp)
	sub	%rsp,0x7f(%rbp)
	sub	%rsp,0x7f(%rsi)
	sub	%rsp,0x7f(%rdi)
	sub	%rsp,0x7f(%r8)
	sub	%rsp,0x7f(%r9)
	sub	%rsp,0x7f(%r10)
	sub	%rsp,0x7f(%r11)
	sub	%rsp,0x7f(%r12)
	sub	%rsp,0x7f(%r13)
	sub	%rsp,0x7f(%r14)
	sub	%rsp,0x7f(%r15)
	nop
	sub	%rbp,0x7f(%rax)
	sub	%rbp,0x7f(%rcx)
	sub	%rbp,0x7f(%rdx)
	sub	%rbp,0x7f(%rbx)
	sub	%rbp,0x7f(%rsp)
	sub	%rbp,0x7f(%rbp)
	sub	%rbp,0x7f(%rsi)
	sub	%rbp,0x7f(%rdi)
	sub	%rbp,0x7f(%r8)
	sub	%rbp,0x7f(%r9)
	sub	%rbp,0x7f(%r10)
	sub	%rbp,0x7f(%r11)
	sub	%rbp,0x7f(%r12)
	sub	%rbp,0x7f(%r13)
	sub	%rbp,0x7f(%r14)
	sub	%rbp,0x7f(%r15)
	nop
	sub	%rsi,0x7f(%rax)
	sub	%rsi,0x7f(%rcx)
	sub	%rsi,0x7f(%rdx)
	sub	%rsi,0x7f(%rbx)
	sub	%rsi,0x7f(%rsp)
	sub	%rsi,0x7f(%rbp)
	sub	%rsi,0x7f(%rsi)
	sub	%rsi,0x7f(%rdi)
	sub	%rsi,0x7f(%r8)
	sub	%rsi,0x7f(%r9)
	sub	%rsi,0x7f(%r10)
	sub	%rsi,0x7f(%r11)
	sub	%rsi,0x7f(%r12)
	sub	%rsi,0x7f(%r13)
	sub	%rsi,0x7f(%r14)
	sub	%rsi,0x7f(%r15)
	nop
	sub	%rdi,0x7f(%rax)
	sub	%rdi,0x7f(%rcx)
	sub	%rdi,0x7f(%rdx)
	sub	%rdi,0x7f(%rbx)
	sub	%rdi,0x7f(%rsp)
	sub	%rdi,0x7f(%rbp)
	sub	%rdi,0x7f(%rsi)
	sub	%rdi,0x7f(%rdi)
	sub	%rdi,0x7f(%r8)
	sub	%rdi,0x7f(%r9)
	sub	%rdi,0x7f(%r10)
	sub	%rdi,0x7f(%r11)
	sub	%rdi,0x7f(%r12)
	sub	%rdi,0x7f(%r13)
	sub	%rdi,0x7f(%r14)
	sub	%rdi,0x7f(%r15)
	nop
	sub	%r8, 0x7f(%rax)
	sub	%r8, 0x7f(%rcx)
	sub	%r8, 0x7f(%rdx)
	sub	%r8, 0x7f(%rbx)
	sub	%r8, 0x7f(%rsp)
	sub	%r8, 0x7f(%rbp)
	sub	%r8, 0x7f(%rsi)
	sub	%r8, 0x7f(%rdi)
	sub	%r8, 0x7f(%r8)
	sub	%r8, 0x7f(%r9)
	sub	%r8, 0x7f(%r10)
	sub	%r8, 0x7f(%r11)
	sub	%r8, 0x7f(%r12)
	sub	%r8, 0x7f(%r13)
	sub	%r8, 0x7f(%r14)
	sub	%r8, 0x7f(%r15)
	nop
	sub	%r9, 0x7f(%rax)
	sub	%r9, 0x7f(%rcx)
	sub	%r9, 0x7f(%rdx)
	sub	%r9, 0x7f(%rbx)
	sub	%r9, 0x7f(%rsp)
	sub	%r9, 0x7f(%rbp)
	sub	%r9, 0x7f(%rsi)
	sub	%r9, 0x7f(%rdi)
	sub	%r9, 0x7f(%r8)
	sub	%r9, 0x7f(%r9)
	sub	%r9, 0x7f(%r10)
	sub	%r9, 0x7f(%r11)
	sub	%r9, 0x7f(%r12)
	sub	%r9, 0x7f(%r13)
	sub	%r9, 0x7f(%r14)
	sub	%r9, 0x7f(%r15)
	nop
	sub	%r10,0x7f(%rax)
	sub	%r10,0x7f(%rcx)
	sub	%r10,0x7f(%rdx)
	sub	%r10,0x7f(%rbx)
	sub	%r10,0x7f(%rsp)
	sub	%r10,0x7f(%rbp)
	sub	%r10,0x7f(%rsi)
	sub	%r10,0x7f(%rdi)
	sub	%r10,0x7f(%r8)
	sub	%r10,0x7f(%r9)
	sub	%r10,0x7f(%r10)
	sub	%r10,0x7f(%r11)
	sub	%r10,0x7f(%r12)
	sub	%r10,0x7f(%r13)
	sub	%r10,0x7f(%r14)
	sub	%r10,0x7f(%r15)
	nop
	sub	%r11,0x7f(%rax)
	sub	%r11,0x7f(%rcx)
	sub	%r11,0x7f(%rdx)
	sub	%r11,0x7f(%rbx)
	sub	%r11,0x7f(%rsp)
	sub	%r11,0x7f(%rbp)
	sub	%r11,0x7f(%rsi)
	sub	%r11,0x7f(%rdi)
	sub	%r11,0x7f(%r8)
	sub	%r11,0x7f(%r9)
	sub	%r11,0x7f(%r10)
	sub	%r11,0x7f(%r11)
	sub	%r11,0x7f(%r12)
	sub	%r11,0x7f(%r13)
	sub	%r11,0x7f(%r14)
	sub	%r11,0x7f(%r15)
	nop
	sub	%r12,0x7f(%rax)
	sub	%r12,0x7f(%rcx)
	sub	%r12,0x7f(%rdx)
	sub	%r12,0x7f(%rbx)
	sub	%r12,0x7f(%rsp)
	sub	%r12,0x7f(%rbp)
	sub	%r12,0x7f(%rsi)
	sub	%r12,0x7f(%rdi)
	sub	%r12,0x7f(%r8)
	sub	%r12,0x7f(%r9)
	sub	%r12,0x7f(%r10)
	sub	%r12,0x7f(%r11)
	sub	%r12,0x7f(%r12)
	sub	%r12,0x7f(%r13)
	sub	%r12,0x7f(%r14)
	sub	%r12,0x7f(%r15)
	nop
	sub	%r13,0x7f(%rax)
	sub	%r13,0x7f(%rcx)
	sub	%r13,0x7f(%rdx)
	sub	%r13,0x7f(%rbx)
	sub	%r13,0x7f(%rsp)
	sub	%r13,0x7f(%rbp)
	sub	%r13,0x7f(%rsi)
	sub	%r13,0x7f(%rdi)
	sub	%r13,0x7f(%r8)
	sub	%r13,0x7f(%r9)
	sub	%r13,0x7f(%r10)
	sub	%r13,0x7f(%r11)
	sub	%r13,0x7f(%r12)
	sub	%r13,0x7f(%r13)
	sub	%r13,0x7f(%r14)
	sub	%r13,0x7f(%r15)
	nop
	sub	%r14,0x7f(%rax)
	sub	%r14,0x7f(%rcx)
	sub	%r14,0x7f(%rdx)
	sub	%r14,0x7f(%rbx)
	sub	%r14,0x7f(%rsp)
	sub	%r14,0x7f(%rbp)
	sub	%r14,0x7f(%rsi)
	sub	%r14,0x7f(%rdi)
	sub	%r14,0x7f(%r8)
	sub	%r14,0x7f(%r9)
	sub	%r14,0x7f(%r10)
	sub	%r14,0x7f(%r11)
	sub	%r14,0x7f(%r12)
	sub	%r14,0x7f(%r13)
	sub	%r14,0x7f(%r14)
	sub	%r14,0x7f(%r15)
	nop
	sub	%r15,0x7f(%rax)
	sub	%r15,0x7f(%rcx)
	sub	%r15,0x7f(%rdx)
	sub	%r15,0x7f(%rbx)
	sub	%r15,0x7f(%rsp)
	sub	%r15,0x7f(%rbp)
	sub	%r15,0x7f(%rsi)
	sub	%r15,0x7f(%rdi)
	sub	%r15,0x7f(%r8)
	sub	%r15,0x7f(%r9)
	sub	%r15,0x7f(%r10)
	sub	%r15,0x7f(%r11)
	sub	%r15,0x7f(%r12)
	sub	%r15,0x7f(%r13)
	sub	%r15,0x7f(%r14)
	sub	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	SubMem32Reg
	.type	SubMem32Reg, @function
SubMem32Reg:
	.cfi_startproc
	sub	%rax,0x7f563412(%rax)
	sub	%rax,0x7f563412(%rcx)
	sub	%rax,0x7f563412(%rdx)
	sub	%rax,0x7f563412(%rbx)
	sub	%rax,0x7f563412(%rsp)
	sub	%rax,0x7f563412(%rbp)
	sub	%rax,0x7f563412(%rsi)
	sub	%rax,0x7f563412(%rdi)
	sub	%rax,0x7f563412(%r8)
	sub	%rax,0x7f563412(%r9)
	sub	%rax,0x7f563412(%r10)
	sub	%rax,0x7f563412(%r11)
	sub	%rax,0x7f563412(%r12)
	sub	%rax,0x7f563412(%r13)
	sub	%rax,0x7f563412(%r14)
	sub	%rax,0x7f563412(%r15)
	nop
	sub	%rcx,0x7f563412(%rax)
	sub	%rcx,0x7f563412(%rcx)
	sub	%rcx,0x7f563412(%rdx)
	sub	%rcx,0x7f563412(%rbx)
	sub	%rcx,0x7f563412(%rsp)
	sub	%rcx,0x7f563412(%rbp)
	sub	%rcx,0x7f563412(%rsi)
	sub	%rcx,0x7f563412(%rdi)
	sub	%rcx,0x7f563412(%r8)
	sub	%rcx,0x7f563412(%r9)
	sub	%rcx,0x7f563412(%r10)
	sub	%rcx,0x7f563412(%r11)
	sub	%rcx,0x7f563412(%r12)
	sub	%rcx,0x7f563412(%r13)
	sub	%rcx,0x7f563412(%r14)
	sub	%rcx,0x7f563412(%r15)
	nop
	sub	%rdx,0x7f563412(%rax)
	sub	%rdx,0x7f563412(%rcx)
	sub	%rdx,0x7f563412(%rdx)
	sub	%rdx,0x7f563412(%rbx)
	sub	%rdx,0x7f563412(%rsp)
	sub	%rdx,0x7f563412(%rbp)
	sub	%rdx,0x7f563412(%rsi)
	sub	%rdx,0x7f563412(%rdi)
	sub	%rdx,0x7f563412(%r8)
	sub	%rdx,0x7f563412(%r9)
	sub	%rdx,0x7f563412(%r10)
	sub	%rdx,0x7f563412(%r11)
	sub	%rdx,0x7f563412(%r12)
	sub	%rdx,0x7f563412(%r13)
	sub	%rdx,0x7f563412(%r14)
	sub	%rdx,0x7f563412(%r15)
	nop
	sub	%rbx,0x7f563412(%rax)
	sub	%rbx,0x7f563412(%rcx)
	sub	%rbx,0x7f563412(%rdx)
	sub	%rbx,0x7f563412(%rbx)
	sub	%rbx,0x7f563412(%rsp)
	sub	%rbx,0x7f563412(%rbp)
	sub	%rbx,0x7f563412(%rsi)
	sub	%rbx,0x7f563412(%rdi)
	sub	%rbx,0x7f563412(%r8)
	sub	%rbx,0x7f563412(%r9)
	sub	%rbx,0x7f563412(%r10)
	sub	%rbx,0x7f563412(%r11)
	sub	%rbx,0x7f563412(%r12)
	sub	%rbx,0x7f563412(%r13)
	sub	%rbx,0x7f563412(%r14)
	sub	%rbx,0x7f563412(%r15)
	nop
	sub	%rsp,0x7f563412(%rax)
	sub	%rsp,0x7f563412(%rcx)
	sub	%rsp,0x7f563412(%rdx)
	sub	%rsp,0x7f563412(%rbx)
	sub	%rsp,0x7f563412(%rsp)
	sub	%rsp,0x7f563412(%rbp)
	sub	%rsp,0x7f563412(%rsi)
	sub	%rsp,0x7f563412(%rdi)
	sub	%rsp,0x7f563412(%r8)
	sub	%rsp,0x7f563412(%r9)
	sub	%rsp,0x7f563412(%r10)
	sub	%rsp,0x7f563412(%r11)
	sub	%rsp,0x7f563412(%r12)
	sub	%rsp,0x7f563412(%r13)
	sub	%rsp,0x7f563412(%r14)
	sub	%rsp,0x7f563412(%r15)
	nop
	sub	%rbp,0x7f563412(%rax)
	sub	%rbp,0x7f563412(%rcx)
	sub	%rbp,0x7f563412(%rdx)
	sub	%rbp,0x7f563412(%rbx)
	sub	%rbp,0x7f563412(%rsp)
	sub	%rbp,0x7f563412(%rbp)
	sub	%rbp,0x7f563412(%rsi)
	sub	%rbp,0x7f563412(%rdi)
	sub	%rbp,0x7f563412(%r8)
	sub	%rbp,0x7f563412(%r9)
	sub	%rbp,0x7f563412(%r10)
	sub	%rbp,0x7f563412(%r11)
	sub	%rbp,0x7f563412(%r12)
	sub	%rbp,0x7f563412(%r13)
	sub	%rbp,0x7f563412(%r14)
	sub	%rbp,0x7f563412(%r15)
	nop
	sub	%rsi,0x7f563412(%rax)
	sub	%rsi,0x7f563412(%rcx)
	sub	%rsi,0x7f563412(%rdx)
	sub	%rsi,0x7f563412(%rbx)
	sub	%rsi,0x7f563412(%rsp)
	sub	%rsi,0x7f563412(%rbp)
	sub	%rsi,0x7f563412(%rsi)
	sub	%rsi,0x7f563412(%rdi)
	sub	%rsi,0x7f563412(%r8)
	sub	%rsi,0x7f563412(%r9)
	sub	%rsi,0x7f563412(%r10)
	sub	%rsi,0x7f563412(%r11)
	sub	%rsi,0x7f563412(%r12)
	sub	%rsi,0x7f563412(%r13)
	sub	%rsi,0x7f563412(%r14)
	sub	%rsi,0x7f563412(%r15)
	nop
	sub	%rdi,0x7f563412(%rax)
	sub	%rdi,0x7f563412(%rcx)
	sub	%rdi,0x7f563412(%rdx)
	sub	%rdi,0x7f563412(%rbx)
	sub	%rdi,0x7f563412(%rsp)
	sub	%rdi,0x7f563412(%rbp)
	sub	%rdi,0x7f563412(%rsi)
	sub	%rdi,0x7f563412(%rdi)
	sub	%rdi,0x7f563412(%r8)
	sub	%rdi,0x7f563412(%r9)
	sub	%rdi,0x7f563412(%r10)
	sub	%rdi,0x7f563412(%r11)
	sub	%rdi,0x7f563412(%r12)
	sub	%rdi,0x7f563412(%r13)
	sub	%rdi,0x7f563412(%r14)
	sub	%rdi,0x7f563412(%r15)
	nop
	sub	%r8, 0x7f563412(%rax)
	sub	%r8, 0x7f563412(%rcx)
	sub	%r8, 0x7f563412(%rdx)
	sub	%r8, 0x7f563412(%rbx)
	sub	%r8, 0x7f563412(%rsp)
	sub	%r8, 0x7f563412(%rbp)
	sub	%r8, 0x7f563412(%rsi)
	sub	%r8, 0x7f563412(%rdi)
	sub	%r8, 0x7f563412(%r8)
	sub	%r8, 0x7f563412(%r9)
	sub	%r8, 0x7f563412(%r10)
	sub	%r8, 0x7f563412(%r11)
	sub	%r8, 0x7f563412(%r12)
	sub	%r8, 0x7f563412(%r13)
	sub	%r8, 0x7f563412(%r14)
	sub	%r8, 0x7f563412(%r15)
	nop
	sub	%r9, 0x7f563412(%rax)
	sub	%r9, 0x7f563412(%rcx)
	sub	%r9, 0x7f563412(%rdx)
	sub	%r9, 0x7f563412(%rbx)
	sub	%r9, 0x7f563412(%rsp)
	sub	%r9, 0x7f563412(%rbp)
	sub	%r9, 0x7f563412(%rsi)
	sub	%r9, 0x7f563412(%rdi)
	sub	%r9, 0x7f563412(%r8)
	sub	%r9, 0x7f563412(%r9)
	sub	%r9, 0x7f563412(%r10)
	sub	%r9, 0x7f563412(%r11)
	sub	%r9, 0x7f563412(%r12)
	sub	%r9, 0x7f563412(%r13)
	sub	%r9, 0x7f563412(%r14)
	sub	%r9, 0x7f563412(%r15)
	nop
	sub	%r10,0x7f563412(%rax)
	sub	%r10,0x7f563412(%rcx)
	sub	%r10,0x7f563412(%rdx)
	sub	%r10,0x7f563412(%rbx)
	sub	%r10,0x7f563412(%rsp)
	sub	%r10,0x7f563412(%rbp)
	sub	%r10,0x7f563412(%rsi)
	sub	%r10,0x7f563412(%rdi)
	sub	%r10,0x7f563412(%r8)
	sub	%r10,0x7f563412(%r9)
	sub	%r10,0x7f563412(%r10)
	sub	%r10,0x7f563412(%r11)
	sub	%r10,0x7f563412(%r12)
	sub	%r10,0x7f563412(%r13)
	sub	%r10,0x7f563412(%r14)
	sub	%r10,0x7f563412(%r15)
	nop
	sub	%r11,0x7f563412(%rax)
	sub	%r11,0x7f563412(%rcx)
	sub	%r11,0x7f563412(%rdx)
	sub	%r11,0x7f563412(%rbx)
	sub	%r11,0x7f563412(%rsp)
	sub	%r11,0x7f563412(%rbp)
	sub	%r11,0x7f563412(%rsi)
	sub	%r11,0x7f563412(%rdi)
	sub	%r11,0x7f563412(%r8)
	sub	%r11,0x7f563412(%r9)
	sub	%r11,0x7f563412(%r10)
	sub	%r11,0x7f563412(%r11)
	sub	%r11,0x7f563412(%r12)
	sub	%r11,0x7f563412(%r13)
	sub	%r11,0x7f563412(%r14)
	sub	%r11,0x7f563412(%r15)
	nop
	sub	%r12,0x7f563412(%rax)
	sub	%r12,0x7f563412(%rcx)
	sub	%r12,0x7f563412(%rdx)
	sub	%r12,0x7f563412(%rbx)
	sub	%r12,0x7f563412(%rsp)
	sub	%r12,0x7f563412(%rbp)
	sub	%r12,0x7f563412(%rsi)
	sub	%r12,0x7f563412(%rdi)
	sub	%r12,0x7f563412(%r8)
	sub	%r12,0x7f563412(%r9)
	sub	%r12,0x7f563412(%r10)
	sub	%r12,0x7f563412(%r11)
	sub	%r12,0x7f563412(%r12)
	sub	%r12,0x7f563412(%r13)
	sub	%r12,0x7f563412(%r14)
	sub	%r12,0x7f563412(%r15)
	nop
	sub	%r13,0x7f563412(%rax)
	sub	%r13,0x7f563412(%rcx)
	sub	%r13,0x7f563412(%rdx)
	sub	%r13,0x7f563412(%rbx)
	sub	%r13,0x7f563412(%rsp)
	sub	%r13,0x7f563412(%rbp)
	sub	%r13,0x7f563412(%rsi)
	sub	%r13,0x7f563412(%rdi)
	sub	%r13,0x7f563412(%r8)
	sub	%r13,0x7f563412(%r9)
	sub	%r13,0x7f563412(%r10)
	sub	%r13,0x7f563412(%r11)
	sub	%r13,0x7f563412(%r12)
	sub	%r13,0x7f563412(%r13)
	sub	%r13,0x7f563412(%r14)
	sub	%r13,0x7f563412(%r15)
	nop
	sub	%r14,0x7f563412(%rax)
	sub	%r14,0x7f563412(%rcx)
	sub	%r14,0x7f563412(%rdx)
	sub	%r14,0x7f563412(%rbx)
	sub	%r14,0x7f563412(%rsp)
	sub	%r14,0x7f563412(%rbp)
	sub	%r14,0x7f563412(%rsi)
	sub	%r14,0x7f563412(%rdi)
	sub	%r14,0x7f563412(%r8)
	sub	%r14,0x7f563412(%r9)
	sub	%r14,0x7f563412(%r10)
	sub	%r14,0x7f563412(%r11)
	sub	%r14,0x7f563412(%r12)
	sub	%r14,0x7f563412(%r13)
	sub	%r14,0x7f563412(%r14)
	sub	%r14,0x7f563412(%r15)
	nop
	sub	%r15,0x7f563412(%rax)
	sub	%r15,0x7f563412(%rcx)
	sub	%r15,0x7f563412(%rdx)
	sub	%r15,0x7f563412(%rbx)
	sub	%r15,0x7f563412(%rsp)
	sub	%r15,0x7f563412(%rbp)
	sub	%r15,0x7f563412(%rsi)
	sub	%r15,0x7f563412(%rdi)
	sub	%r15,0x7f563412(%r8)
	sub	%r15,0x7f563412(%r9)
	sub	%r15,0x7f563412(%r10)
	sub	%r15,0x7f563412(%r11)
	sub	%r15,0x7f563412(%r12)
	sub	%r15,0x7f563412(%r13)
	sub	%r15,0x7f563412(%r14)
	sub	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	SubRegMem
	.type	SubRegMem, @function
SubRegMem:
	.cfi_startproc
	sub	(%rax),%rax
	sub	(%rax),%rcx
	sub	(%rax),%rdx
	sub	(%rax),%rbx
	sub	(%rax),%rsp
	sub	(%rax),%rbp
	sub	(%rax),%rsi
	sub	(%rax),%rdi
	sub	(%rax),%r8
	sub	(%rax),%r9
	sub	(%rax),%r10
	sub	(%rax),%r11
	sub	(%rax),%r12
	sub	(%rax),%r13
	sub	(%rax),%r14
	sub	(%rax),%r15
	nop
	sub	(%rcx),%rax
	sub	(%rcx),%rcx
	sub	(%rcx),%rdx
	sub	(%rcx),%rbx
	sub	(%rcx),%rsp
	sub	(%rcx),%rbp
	sub	(%rcx),%rsi
	sub	(%rcx),%rdi
	sub	(%rcx),%r8
	sub	(%rcx),%r9
	sub	(%rcx),%r10
	sub	(%rcx),%r11
	sub	(%rcx),%r12
	sub	(%rcx),%r13
	sub	(%rcx),%r14
	sub	(%rcx),%r15
	nop
	sub	(%rdx),%rax
	sub	(%rdx),%rcx
	sub	(%rdx),%rdx
	sub	(%rdx),%rbx
	sub	(%rdx),%rsp
	sub	(%rdx),%rbp
	sub	(%rdx),%rsi
	sub	(%rdx),%rdi
	sub	(%rdx),%r8
	sub	(%rdx),%r9
	sub	(%rdx),%r10
	sub	(%rdx),%r11
	sub	(%rdx),%r12
	sub	(%rdx),%r13
	sub	(%rdx),%r14
	sub	(%rdx),%r15
	nop
	sub	(%rbx),%rax
	sub	(%rbx),%rcx
	sub	(%rbx),%rdx
	sub	(%rbx),%rbx
	sub	(%rbx),%rsp
	sub	(%rbx),%rbp
	sub	(%rbx),%rsi
	sub	(%rbx),%rdi
	sub	(%rbx),%r8
	sub	(%rbx),%r9
	sub	(%rbx),%r10
	sub	(%rbx),%r11
	sub	(%rbx),%r12
	sub	(%rbx),%r13
	sub	(%rbx),%r14
	sub	(%rbx),%r15
	nop
	sub	(%rsp),%rax
	sub	(%rsp),%rcx
	sub	(%rsp),%rdx
	sub	(%rsp),%rbx
	sub	(%rsp),%rsp
	sub	(%rsp),%rbp
	sub	(%rsp),%rsi
	sub	(%rsp),%rdi
	sub	(%rsp),%r8
	sub	(%rsp),%r9
	sub	(%rsp),%r10
	sub	(%rsp),%r11
	sub	(%rsp),%r12
	sub	(%rsp),%r13
	sub	(%rsp),%r14
	sub	(%rsp),%r15
	nop
	sub	(%rbp),%rax
	sub	(%rbp),%rcx
	sub	(%rbp),%rdx
	sub	(%rbp),%rbx
	sub	(%rbp),%rsp
	sub	(%rbp),%rbp
	sub	(%rbp),%rsi
	sub	(%rbp),%rdi
	sub	(%rbp),%r8
	sub	(%rbp),%r9
	sub	(%rbp),%r10
	sub	(%rbp),%r11
	sub	(%rbp),%r12
	sub	(%rbp),%r13
	sub	(%rbp),%r14
	sub	(%rbp),%r15
	nop
	sub	(%rsi),%rax
	sub	(%rsi),%rcx
	sub	(%rsi),%rdx
	sub	(%rsi),%rbx
	sub	(%rsi),%rsp
	sub	(%rsi),%rbp
	sub	(%rsi),%rsi
	sub	(%rsi),%rdi
	sub	(%rsi),%r8
	sub	(%rsi),%r9
	sub	(%rsi),%r10
	sub	(%rsi),%r11
	sub	(%rsi),%r12
	sub	(%rsi),%r13
	sub	(%rsi),%r14
	sub	(%rsi),%r15
	nop
	sub	(%rdi),%rax
	sub	(%rdi),%rcx
	sub	(%rdi),%rdx
	sub	(%rdi),%rbx
	sub	(%rdi),%rsp
	sub	(%rdi),%rbp
	sub	(%rdi),%rsi
	sub	(%rdi),%rdi
	sub	(%rdi),%r8
	sub	(%rdi),%r9
	sub	(%rdi),%r10
	sub	(%rdi),%r11
	sub	(%rdi),%r12
	sub	(%rdi),%r13
	sub	(%rdi),%r14
	sub	(%rdi),%r15
	nop
	sub	(%r8), %rax
	sub	(%r8), %rcx
	sub	(%r8), %rdx
	sub	(%r8), %rbx
	sub	(%r8), %rsp
	sub	(%r8), %rbp
	sub	(%r8), %rsi
	sub	(%r8), %rdi
	sub	(%r8), %r8
	sub	(%r8), %r9
	sub	(%r8), %r10
	sub	(%r8), %r11
	sub	(%r8), %r12
	sub	(%r8), %r13
	sub	(%r8), %r14
	sub	(%r8), %r15
	nop
	sub	(%r9), %rax
	sub	(%r9), %rcx
	sub	(%r9), %rdx
	sub	(%r9), %rbx
	sub	(%r9), %rsp
	sub	(%r9), %rbp
	sub	(%r9), %rsi
	sub	(%r9), %rdi
	sub	(%r9), %r8
	sub	(%r9), %r9
	sub	(%r9), %r10
	sub	(%r9), %r11
	sub	(%r9), %r12
	sub	(%r9), %r13
	sub	(%r9), %r14
	sub	(%r9), %r15
	nop
	sub	(%r10),%rax
	sub	(%r10),%rcx
	sub	(%r10),%rdx
	sub	(%r10),%rbx
	sub	(%r10),%rsp
	sub	(%r10),%rbp
	sub	(%r10),%rsi
	sub	(%r10),%rdi
	sub	(%r10),%r8
	sub	(%r10),%r9
	sub	(%r10),%r10
	sub	(%r10),%r11
	sub	(%r10),%r12
	sub	(%r10),%r13
	sub	(%r10),%r14
	sub	(%r10),%r15
	nop
	sub	(%r11),%rax
	sub	(%r11),%rcx
	sub	(%r11),%rdx
	sub	(%r11),%rbx
	sub	(%r11),%rsp
	sub	(%r11),%rbp
	sub	(%r11),%rsi
	sub	(%r11),%rdi
	sub	(%r11),%r8
	sub	(%r11),%r9
	sub	(%r11),%r10
	sub	(%r11),%r11
	sub	(%r11),%r12
	sub	(%r11),%r13
	sub	(%r11),%r14
	sub	(%r11),%r15
	nop
	sub	(%r12),%rax
	sub	(%r12),%rcx
	sub	(%r12),%rdx
	sub	(%r12),%rbx
	sub	(%r12),%rsp
	sub	(%r12),%rbp
	sub	(%r12),%rsi
	sub	(%r12),%rdi
	sub	(%r12),%r8
	sub	(%r12),%r9
	sub	(%r12),%r10
	sub	(%r12),%r11
	sub	(%r12),%r12
	sub	(%r12),%r13
	sub	(%r12),%r14
	sub	(%r12),%r15
	nop
	sub	(%r13),%rax
	sub	(%r13),%rcx
	sub	(%r13),%rdx
	sub	(%r13),%rbx
	sub	(%r13),%rsp
	sub	(%r13),%rbp
	sub	(%r13),%rsi
	sub	(%r13),%rdi
	sub	(%r13),%r8
	sub	(%r13),%r9
	sub	(%r13),%r10
	sub	(%r13),%r11
	sub	(%r13),%r12
	sub	(%r13),%r13
	sub	(%r13),%r14
	sub	(%r13),%r15
	nop
	sub	(%r14),%rax
	sub	(%r14),%rcx
	sub	(%r14),%rdx
	sub	(%r14),%rbx
	sub	(%r14),%rsp
	sub	(%r14),%rbp
	sub	(%r14),%rsi
	sub	(%r14),%rdi
	sub	(%r14),%r8
	sub	(%r14),%r9
	sub	(%r14),%r10
	sub	(%r14),%r11
	sub	(%r14),%r12
	sub	(%r14),%r13
	sub	(%r14),%r14
	sub	(%r14),%r15
	nop
	sub	(%r15),%rax
	sub	(%r15),%rcx
	sub	(%r15),%rdx
	sub	(%r15),%rbx
	sub	(%r15),%rsp
	sub	(%r15),%rbp
	sub	(%r15),%rsi
	sub	(%r15),%rdi
	sub	(%r15),%r8
	sub	(%r15),%r9
	sub	(%r15),%r10
	sub	(%r15),%r11
	sub	(%r15),%r12
	sub	(%r15),%r13
	sub	(%r15),%r14
	sub	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	SubRegMem8
	.type	SubRegMem8, @function
SubRegMem8:
	.cfi_startproc
	sub	0x7f(%rax),%rax
	sub	0x7f(%rax),%rcx
	sub	0x7f(%rax),%rdx
	sub	0x7f(%rax),%rbx
	sub	0x7f(%rax),%rsp
	sub	0x7f(%rax),%rbp
	sub	0x7f(%rax),%rsi
	sub	0x7f(%rax),%rdi
	sub	0x7f(%rax),%r8
	sub	0x7f(%rax),%r9
	sub	0x7f(%rax),%r10
	sub	0x7f(%rax),%r11
	sub	0x7f(%rax),%r12
	sub	0x7f(%rax),%r13
	sub	0x7f(%rax),%r14
	sub	0x7f(%rax),%r15
	nop
	sub	0x7f(%rcx),%rax
	sub	0x7f(%rcx),%rcx
	sub	0x7f(%rcx),%rdx
	sub	0x7f(%rcx),%rbx
	sub	0x7f(%rcx),%rsp
	sub	0x7f(%rcx),%rbp
	sub	0x7f(%rcx),%rsi
	sub	0x7f(%rcx),%rdi
	sub	0x7f(%rcx),%r8
	sub	0x7f(%rcx),%r9
	sub	0x7f(%rcx),%r10
	sub	0x7f(%rcx),%r11
	sub	0x7f(%rcx),%r12
	sub	0x7f(%rcx),%r13
	sub	0x7f(%rcx),%r14
	sub	0x7f(%rcx),%r15
	nop
	sub	0x7f(%rdx),%rax
	sub	0x7f(%rdx),%rcx
	sub	0x7f(%rdx),%rdx
	sub	0x7f(%rdx),%rbx
	sub	0x7f(%rdx),%rsp
	sub	0x7f(%rdx),%rbp
	sub	0x7f(%rdx),%rsi
	sub	0x7f(%rdx),%rdi
	sub	0x7f(%rdx),%r8
	sub	0x7f(%rdx),%r9
	sub	0x7f(%rdx),%r10
	sub	0x7f(%rdx),%r11
	sub	0x7f(%rdx),%r12
	sub	0x7f(%rdx),%r13
	sub	0x7f(%rdx),%r14
	sub	0x7f(%rdx),%r15
	nop
	sub	0x7f(%rbx),%rax
	sub	0x7f(%rbx),%rcx
	sub	0x7f(%rbx),%rdx
	sub	0x7f(%rbx),%rbx
	sub	0x7f(%rbx),%rsp
	sub	0x7f(%rbx),%rbp
	sub	0x7f(%rbx),%rsi
	sub	0x7f(%rbx),%rdi
	sub	0x7f(%rbx),%r8
	sub	0x7f(%rbx),%r9
	sub	0x7f(%rbx),%r10
	sub	0x7f(%rbx),%r11
	sub	0x7f(%rbx),%r12
	sub	0x7f(%rbx),%r13
	sub	0x7f(%rbx),%r14
	sub	0x7f(%rbx),%r15
	nop
	sub	0x7f(%rsp),%rax
	sub	0x7f(%rsp),%rcx
	sub	0x7f(%rsp),%rdx
	sub	0x7f(%rsp),%rbx
	sub	0x7f(%rsp),%rsp
	sub	0x7f(%rsp),%rbp
	sub	0x7f(%rsp),%rsi
	sub	0x7f(%rsp),%rdi
	sub	0x7f(%rsp),%r8
	sub	0x7f(%rsp),%r9
	sub	0x7f(%rsp),%r10
	sub	0x7f(%rsp),%r11
	sub	0x7f(%rsp),%r12
	sub	0x7f(%rsp),%r13
	sub	0x7f(%rsp),%r14
	sub	0x7f(%rsp),%r15
	nop
	sub	0x7f(%rbp),%rax
	sub	0x7f(%rbp),%rcx
	sub	0x7f(%rbp),%rdx
	sub	0x7f(%rbp),%rbx
	sub	0x7f(%rbp),%rsp
	sub	0x7f(%rbp),%rbp
	sub	0x7f(%rbp),%rsi
	sub	0x7f(%rbp),%rdi
	sub	0x7f(%rbp),%r8
	sub	0x7f(%rbp),%r9
	sub	0x7f(%rbp),%r10
	sub	0x7f(%rbp),%r11
	sub	0x7f(%rbp),%r12
	sub	0x7f(%rbp),%r13
	sub	0x7f(%rbp),%r14
	sub	0x7f(%rbp),%r15
	nop
	sub	0x7f(%rsi),%rax
	sub	0x7f(%rsi),%rcx
	sub	0x7f(%rsi),%rdx
	sub	0x7f(%rsi),%rbx
	sub	0x7f(%rsi),%rsp
	sub	0x7f(%rsi),%rbp
	sub	0x7f(%rsi),%rsi
	sub	0x7f(%rsi),%rdi
	sub	0x7f(%rsi),%r8
	sub	0x7f(%rsi),%r9
	sub	0x7f(%rsi),%r10
	sub	0x7f(%rsi),%r11
	sub	0x7f(%rsi),%r12
	sub	0x7f(%rsi),%r13
	sub	0x7f(%rsi),%r14
	sub	0x7f(%rsi),%r15
	nop
	sub	0x7f(%rdi),%rax
	sub	0x7f(%rdi),%rcx
	sub	0x7f(%rdi),%rdx
	sub	0x7f(%rdi),%rbx
	sub	0x7f(%rdi),%rsp
	sub	0x7f(%rdi),%rbp
	sub	0x7f(%rdi),%rsi
	sub	0x7f(%rdi),%rdi
	sub	0x7f(%rdi),%r8
	sub	0x7f(%rdi),%r9
	sub	0x7f(%rdi),%r10
	sub	0x7f(%rdi),%r11
	sub	0x7f(%rdi),%r12
	sub	0x7f(%rdi),%r13
	sub	0x7f(%rdi),%r14
	sub	0x7f(%rdi),%r15
	nop
	sub	0x7f(%r8), %rax
	sub	0x7f(%r8), %rcx
	sub	0x7f(%r8), %rdx
	sub	0x7f(%r8), %rbx
	sub	0x7f(%r8), %rsp
	sub	0x7f(%r8), %rbp
	sub	0x7f(%r8), %rsi
	sub	0x7f(%r8), %rdi
	sub	0x7f(%r8), %r8
	sub	0x7f(%r8), %r9
	sub	0x7f(%r8), %r10
	sub	0x7f(%r8), %r11
	sub	0x7f(%r8), %r12
	sub	0x7f(%r8), %r13
	sub	0x7f(%r8), %r14
	sub	0x7f(%r8), %r15
	nop
	sub	0x7f(%r9), %rax
	sub	0x7f(%r9), %rcx
	sub	0x7f(%r9), %rdx
	sub	0x7f(%r9), %rbx
	sub	0x7f(%r9), %rsp
	sub	0x7f(%r9), %rbp
	sub	0x7f(%r9), %rsi
	sub	0x7f(%r9), %rdi
	sub	0x7f(%r9), %r8
	sub	0x7f(%r9), %r9
	sub	0x7f(%r9), %r10
	sub	0x7f(%r9), %r11
	sub	0x7f(%r9), %r12
	sub	0x7f(%r9), %r13
	sub	0x7f(%r9), %r14
	sub	0x7f(%r9), %r15
	nop
	sub	0x7f(%r10),%rax
	sub	0x7f(%r10),%rcx
	sub	0x7f(%r10),%rdx
	sub	0x7f(%r10),%rbx
	sub	0x7f(%r10),%rsp
	sub	0x7f(%r10),%rbp
	sub	0x7f(%r10),%rsi
	sub	0x7f(%r10),%rdi
	sub	0x7f(%r10),%r8
	sub	0x7f(%r10),%r9
	sub	0x7f(%r10),%r10
	sub	0x7f(%r10),%r11
	sub	0x7f(%r10),%r12
	sub	0x7f(%r10),%r13
	sub	0x7f(%r10),%r14
	sub	0x7f(%r10),%r15
	nop
	sub	0x7f(%r11),%rax
	sub	0x7f(%r11),%rcx
	sub	0x7f(%r11),%rdx
	sub	0x7f(%r11),%rbx
	sub	0x7f(%r11),%rsp
	sub	0x7f(%r11),%rbp
	sub	0x7f(%r11),%rsi
	sub	0x7f(%r11),%rdi
	sub	0x7f(%r11),%r8
	sub	0x7f(%r11),%r9
	sub	0x7f(%r11),%r10
	sub	0x7f(%r11),%r11
	sub	0x7f(%r11),%r12
	sub	0x7f(%r11),%r13
	sub	0x7f(%r11),%r14
	sub	0x7f(%r11),%r15
	nop
	sub	0x7f(%r12),%rax
	sub	0x7f(%r12),%rcx
	sub	0x7f(%r12),%rdx
	sub	0x7f(%r12),%rbx
	sub	0x7f(%r12),%rsp
	sub	0x7f(%r12),%rbp
	sub	0x7f(%r12),%rsi
	sub	0x7f(%r12),%rdi
	sub	0x7f(%r12),%r8
	sub	0x7f(%r12),%r9
	sub	0x7f(%r12),%r10
	sub	0x7f(%r12),%r11
	sub	0x7f(%r12),%r12
	sub	0x7f(%r12),%r13
	sub	0x7f(%r12),%r14
	sub	0x7f(%r12),%r15
	nop
	sub	0x7f(%r13),%rax
	sub	0x7f(%r13),%rcx
	sub	0x7f(%r13),%rdx
	sub	0x7f(%r13),%rbx
	sub	0x7f(%r13),%rsp
	sub	0x7f(%r13),%rbp
	sub	0x7f(%r13),%rsi
	sub	0x7f(%r13),%rdi
	sub	0x7f(%r13),%r8
	sub	0x7f(%r13),%r9
	sub	0x7f(%r13),%r10
	sub	0x7f(%r13),%r11
	sub	0x7f(%r13),%r12
	sub	0x7f(%r13),%r13
	sub	0x7f(%r13),%r14
	sub	0x7f(%r13),%r15
	nop
	sub	0x7f(%r14),%rax
	sub	0x7f(%r14),%rcx
	sub	0x7f(%r14),%rdx
	sub	0x7f(%r14),%rbx
	sub	0x7f(%r14),%rsp
	sub	0x7f(%r14),%rbp
	sub	0x7f(%r14),%rsi
	sub	0x7f(%r14),%rdi
	sub	0x7f(%r14),%r8
	sub	0x7f(%r14),%r9
	sub	0x7f(%r14),%r10
	sub	0x7f(%r14),%r11
	sub	0x7f(%r14),%r12
	sub	0x7f(%r14),%r13
	sub	0x7f(%r14),%r14
	sub	0x7f(%r14),%r15
	nop
	sub	0x7f(%r15),%rax
	sub	0x7f(%r15),%rcx
	sub	0x7f(%r15),%rdx
	sub	0x7f(%r15),%rbx
	sub	0x7f(%r15),%rsp
	sub	0x7f(%r15),%rbp
	sub	0x7f(%r15),%rsi
	sub	0x7f(%r15),%rdi
	sub	0x7f(%r15),%r8
	sub	0x7f(%r15),%r9
	sub	0x7f(%r15),%r10
	sub	0x7f(%r15),%r11
	sub	0x7f(%r15),%r12
	sub	0x7f(%r15),%r13
	sub	0x7f(%r15),%r14
	sub	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	SubRegMem32
	.type	SubRegMem32, @function
SubRegMem32:
	.cfi_startproc
	sub	0x7f563412(%rax),%rax
	sub	0x7f563412(%rax),%rcx
	sub	0x7f563412(%rax),%rdx
	sub	0x7f563412(%rax),%rbx
	sub	0x7f563412(%rax),%rsp
	sub	0x7f563412(%rax),%rbp
	sub	0x7f563412(%rax),%rsi
	sub	0x7f563412(%rax),%rdi
	sub	0x7f563412(%rax),%r8
	sub	0x7f563412(%rax),%r9
	sub	0x7f563412(%rax),%r10
	sub	0x7f563412(%rax),%r11
	sub	0x7f563412(%rax),%r12
	sub	0x7f563412(%rax),%r13
	sub	0x7f563412(%rax),%r14
	sub	0x7f563412(%rax),%r15
	nop
	sub	0x7f563412(%rcx),%rax
	sub	0x7f563412(%rcx),%rcx
	sub	0x7f563412(%rcx),%rdx
	sub	0x7f563412(%rcx),%rbx
	sub	0x7f563412(%rcx),%rsp
	sub	0x7f563412(%rcx),%rbp
	sub	0x7f563412(%rcx),%rsi
	sub	0x7f563412(%rcx),%rdi
	sub	0x7f563412(%rcx),%r8
	sub	0x7f563412(%rcx),%r9
	sub	0x7f563412(%rcx),%r10
	sub	0x7f563412(%rcx),%r11
	sub	0x7f563412(%rcx),%r12
	sub	0x7f563412(%rcx),%r13
	sub	0x7f563412(%rcx),%r14
	sub	0x7f563412(%rcx),%r15
	nop
	sub	0x7f563412(%rdx),%rax
	sub	0x7f563412(%rdx),%rcx
	sub	0x7f563412(%rdx),%rdx
	sub	0x7f563412(%rdx),%rbx
	sub	0x7f563412(%rdx),%rsp
	sub	0x7f563412(%rdx),%rbp
	sub	0x7f563412(%rdx),%rsi
	sub	0x7f563412(%rdx),%rdi
	sub	0x7f563412(%rdx),%r8
	sub	0x7f563412(%rdx),%r9
	sub	0x7f563412(%rdx),%r10
	sub	0x7f563412(%rdx),%r11
	sub	0x7f563412(%rdx),%r12
	sub	0x7f563412(%rdx),%r13
	sub	0x7f563412(%rdx),%r14
	sub	0x7f563412(%rdx),%r15
	nop
	sub	0x7f563412(%rbx),%rax
	sub	0x7f563412(%rbx),%rcx
	sub	0x7f563412(%rbx),%rdx
	sub	0x7f563412(%rbx),%rbx
	sub	0x7f563412(%rbx),%rsp
	sub	0x7f563412(%rbx),%rbp
	sub	0x7f563412(%rbx),%rsi
	sub	0x7f563412(%rbx),%rdi
	sub	0x7f563412(%rbx),%r8
	sub	0x7f563412(%rbx),%r9
	sub	0x7f563412(%rbx),%r10
	sub	0x7f563412(%rbx),%r11
	sub	0x7f563412(%rbx),%r12
	sub	0x7f563412(%rbx),%r13
	sub	0x7f563412(%rbx),%r14
	sub	0x7f563412(%rbx),%r15
	nop
	sub	0x7f563412(%rsp),%rax
	sub	0x7f563412(%rsp),%rcx
	sub	0x7f563412(%rsp),%rdx
	sub	0x7f563412(%rsp),%rbx
	sub	0x7f563412(%rsp),%rsp
	sub	0x7f563412(%rsp),%rbp
	sub	0x7f563412(%rsp),%rsi
	sub	0x7f563412(%rsp),%rdi
	sub	0x7f563412(%rsp),%r8
	sub	0x7f563412(%rsp),%r9
	sub	0x7f563412(%rsp),%r10
	sub	0x7f563412(%rsp),%r11
	sub	0x7f563412(%rsp),%r12
	sub	0x7f563412(%rsp),%r13
	sub	0x7f563412(%rsp),%r14
	sub	0x7f563412(%rsp),%r15
	nop
	sub	0x7f563412(%rbp),%rax
	sub	0x7f563412(%rbp),%rcx
	sub	0x7f563412(%rbp),%rdx
	sub	0x7f563412(%rbp),%rbx
	sub	0x7f563412(%rbp),%rsp
	sub	0x7f563412(%rbp),%rbp
	sub	0x7f563412(%rbp),%rsi
	sub	0x7f563412(%rbp),%rdi
	sub	0x7f563412(%rbp),%r8
	sub	0x7f563412(%rbp),%r9
	sub	0x7f563412(%rbp),%r10
	sub	0x7f563412(%rbp),%r11
	sub	0x7f563412(%rbp),%r12
	sub	0x7f563412(%rbp),%r13
	sub	0x7f563412(%rbp),%r14
	sub	0x7f563412(%rbp),%r15
	nop
	sub	0x7f563412(%rsi),%rax
	sub	0x7f563412(%rsi),%rcx
	sub	0x7f563412(%rsi),%rdx
	sub	0x7f563412(%rsi),%rbx
	sub	0x7f563412(%rsi),%rsp
	sub	0x7f563412(%rsi),%rbp
	sub	0x7f563412(%rsi),%rsi
	sub	0x7f563412(%rsi),%rdi
	sub	0x7f563412(%rsi),%r8
	sub	0x7f563412(%rsi),%r9
	sub	0x7f563412(%rsi),%r10
	sub	0x7f563412(%rsi),%r11
	sub	0x7f563412(%rsi),%r12
	sub	0x7f563412(%rsi),%r13
	sub	0x7f563412(%rsi),%r14
	sub	0x7f563412(%rsi),%r15
	nop
	sub	0x7f563412(%rdi),%rax
	sub	0x7f563412(%rdi),%rcx
	sub	0x7f563412(%rdi),%rdx
	sub	0x7f563412(%rdi),%rbx
	sub	0x7f563412(%rdi),%rsp
	sub	0x7f563412(%rdi),%rbp
	sub	0x7f563412(%rdi),%rsi
	sub	0x7f563412(%rdi),%rdi
	sub	0x7f563412(%rdi),%r8
	sub	0x7f563412(%rdi),%r9
	sub	0x7f563412(%rdi),%r10
	sub	0x7f563412(%rdi),%r11
	sub	0x7f563412(%rdi),%r12
	sub	0x7f563412(%rdi),%r13
	sub	0x7f563412(%rdi),%r14
	sub	0x7f563412(%rdi),%r15
	nop
	sub	0x7f563412(%r8), %rax
	sub	0x7f563412(%r8), %rcx
	sub	0x7f563412(%r8), %rdx
	sub	0x7f563412(%r8), %rbx
	sub	0x7f563412(%r8), %rsp
	sub	0x7f563412(%r8), %rbp
	sub	0x7f563412(%r8), %rsi
	sub	0x7f563412(%r8), %rdi
	sub	0x7f563412(%r8), %r8
	sub	0x7f563412(%r8), %r9
	sub	0x7f563412(%r8), %r10
	sub	0x7f563412(%r8), %r11
	sub	0x7f563412(%r8), %r12
	sub	0x7f563412(%r8), %r13
	sub	0x7f563412(%r8), %r14
	sub	0x7f563412(%r8), %r15
	nop
	sub	0x7f563412(%r9), %rax
	sub	0x7f563412(%r9), %rcx
	sub	0x7f563412(%r9), %rdx
	sub	0x7f563412(%r9), %rbx
	sub	0x7f563412(%r9), %rsp
	sub	0x7f563412(%r9), %rbp
	sub	0x7f563412(%r9), %rsi
	sub	0x7f563412(%r9), %rdi
	sub	0x7f563412(%r9), %r8
	sub	0x7f563412(%r9), %r9
	sub	0x7f563412(%r9), %r10
	sub	0x7f563412(%r9), %r11
	sub	0x7f563412(%r9), %r12
	sub	0x7f563412(%r9), %r13
	sub	0x7f563412(%r9), %r14
	sub	0x7f563412(%r9), %r15
	nop
	sub	0x7f563412(%r10),%rax
	sub	0x7f563412(%r10),%rcx
	sub	0x7f563412(%r10),%rdx
	sub	0x7f563412(%r10),%rbx
	sub	0x7f563412(%r10),%rsp
	sub	0x7f563412(%r10),%rbp
	sub	0x7f563412(%r10),%rsi
	sub	0x7f563412(%r10),%rdi
	sub	0x7f563412(%r10),%r8
	sub	0x7f563412(%r10),%r9
	sub	0x7f563412(%r10),%r10
	sub	0x7f563412(%r10),%r11
	sub	0x7f563412(%r10),%r12
	sub	0x7f563412(%r10),%r13
	sub	0x7f563412(%r10),%r14
	sub	0x7f563412(%r10),%r15
	nop
	sub	0x7f563412(%r11),%rax
	sub	0x7f563412(%r11),%rcx
	sub	0x7f563412(%r11),%rdx
	sub	0x7f563412(%r11),%rbx
	sub	0x7f563412(%r11),%rsp
	sub	0x7f563412(%r11),%rbp
	sub	0x7f563412(%r11),%rsi
	sub	0x7f563412(%r11),%rdi
	sub	0x7f563412(%r11),%r8
	sub	0x7f563412(%r11),%r9
	sub	0x7f563412(%r11),%r10
	sub	0x7f563412(%r11),%r11
	sub	0x7f563412(%r11),%r12
	sub	0x7f563412(%r11),%r13
	sub	0x7f563412(%r11),%r14
	sub	0x7f563412(%r11),%r15
	nop
	sub	0x7f563412(%r12),%rax
	sub	0x7f563412(%r12),%rcx
	sub	0x7f563412(%r12),%rdx
	sub	0x7f563412(%r12),%rbx
	sub	0x7f563412(%r12),%rsp
	sub	0x7f563412(%r12),%rbp
	sub	0x7f563412(%r12),%rsi
	sub	0x7f563412(%r12),%rdi
	sub	0x7f563412(%r12),%r8
	sub	0x7f563412(%r12),%r9
	sub	0x7f563412(%r12),%r10
	sub	0x7f563412(%r12),%r11
	sub	0x7f563412(%r12),%r12
	sub	0x7f563412(%r12),%r13
	sub	0x7f563412(%r12),%r14
	sub	0x7f563412(%r12),%r15
	nop
	sub	0x7f563412(%r13),%rax
	sub	0x7f563412(%r13),%rcx
	sub	0x7f563412(%r13),%rdx
	sub	0x7f563412(%r13),%rbx
	sub	0x7f563412(%r13),%rsp
	sub	0x7f563412(%r13),%rbp
	sub	0x7f563412(%r13),%rsi
	sub	0x7f563412(%r13),%rdi
	sub	0x7f563412(%r13),%r8
	sub	0x7f563412(%r13),%r9
	sub	0x7f563412(%r13),%r10
	sub	0x7f563412(%r13),%r11
	sub	0x7f563412(%r13),%r12
	sub	0x7f563412(%r13),%r13
	sub	0x7f563412(%r13),%r14
	sub	0x7f563412(%r13),%r15
	nop
	sub	0x7f563412(%r14),%rax
	sub	0x7f563412(%r14),%rcx
	sub	0x7f563412(%r14),%rdx
	sub	0x7f563412(%r14),%rbx
	sub	0x7f563412(%r14),%rsp
	sub	0x7f563412(%r14),%rbp
	sub	0x7f563412(%r14),%rsi
	sub	0x7f563412(%r14),%rdi
	sub	0x7f563412(%r14),%r8
	sub	0x7f563412(%r14),%r9
	sub	0x7f563412(%r14),%r10
	sub	0x7f563412(%r14),%r11
	sub	0x7f563412(%r14),%r12
	sub	0x7f563412(%r14),%r13
	sub	0x7f563412(%r14),%r14
	sub	0x7f563412(%r14),%r15
	nop
	sub	0x7f563412(%r15),%rax
	sub	0x7f563412(%r15),%rcx
	sub	0x7f563412(%r15),%rdx
	sub	0x7f563412(%r15),%rbx
	sub	0x7f563412(%r15),%rsp
	sub	0x7f563412(%r15),%rbp
	sub	0x7f563412(%r15),%rsi
	sub	0x7f563412(%r15),%rdi
	sub	0x7f563412(%r15),%r8
	sub	0x7f563412(%r15),%r9
	sub	0x7f563412(%r15),%r10
	sub	0x7f563412(%r15),%r11
	sub	0x7f563412(%r15),%r12
	sub	0x7f563412(%r15),%r13
	sub	0x7f563412(%r15),%r14
	sub	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


