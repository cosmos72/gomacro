	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Mov_s32
	.type	Mov_s32, @function
Mov_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// mov	$-0x55667788,%rax
	mov	$-0x55667788,%rcx
	mov	$-0x55667788,%rdx
	mov	$-0x55667788,%rbx
	mov	$-0x55667788,%rsp
	mov	$-0x55667788,%rbp
	mov	$-0x55667788,%rsi
	mov	$-0x55667788,%rdi
	mov	$-0x55667788,%r8
	mov	$-0x55667788,%r9
	mov	$-0x55667788,%r10
	mov	$-0x55667788,%r11
	mov	$-0x55667788,%r12
	mov	$-0x55667788,%r13
	mov	$-0x55667788,%r14
	mov	$-0x55667788,%r15
	.cfi_endproc


	.p2align 4,,15
	.globl	Mov
	.type	Mov, @function
Mov:
	.cfi_startproc
	mov	%rax,%rax
	mov	%rax,%rcx
	mov	%rax,%rdx
	mov	%rax,%rbx
	mov	%rax,%rsp
	mov	%rax,%rbp
	mov	%rax,%rsi
	mov	%rax,%rdi
	mov	%rax,%r8
	mov	%rax,%r9
	mov	%rax,%r10
	mov	%rax,%r11
	mov	%rax,%r12
	mov	%rax,%r13
	mov	%rax,%r14
	mov	%rax,%r15
	nop
	mov	%rcx,%rax
	mov	%rcx,%rcx
	mov	%rcx,%rdx
	mov	%rcx,%rbx
	mov	%rcx,%rsp
	mov	%rcx,%rbp
	mov	%rcx,%rsi
	mov	%rcx,%rdi
	mov	%rcx,%r8
	mov	%rcx,%r9
	mov	%rcx,%r10
	mov	%rcx,%r11
	mov	%rcx,%r12
	mov	%rcx,%r13
	mov	%rcx,%r14
	mov	%rcx,%r15
	nop
	mov	%rdx,%rax
	mov	%rdx,%rcx
	mov	%rdx,%rdx
	mov	%rdx,%rbx
	mov	%rdx,%rsp
	mov	%rdx,%rbp
	mov	%rdx,%rsi
	mov	%rdx,%rdi
	mov	%rdx,%r8
	mov	%rdx,%r9
	mov	%rdx,%r10
	mov	%rdx,%r11
	mov	%rdx,%r12
	mov	%rdx,%r13
	mov	%rdx,%r14
	mov	%rdx,%r15
	nop
	mov	%rbx,%rax
	mov	%rbx,%rcx
	mov	%rbx,%rdx
	mov	%rbx,%rbx
	mov	%rbx,%rsp
	mov	%rbx,%rbp
	mov	%rbx,%rsi
	mov	%rbx,%rdi
	mov	%rbx,%r8
	mov	%rbx,%r9
	mov	%rbx,%r10
	mov	%rbx,%r11
	mov	%rbx,%r12
	mov	%rbx,%r13
	mov	%rbx,%r14
	mov	%rbx,%r15
	nop
	mov	%rsp,%rax
	mov	%rsp,%rcx
	mov	%rsp,%rdx
	mov	%rsp,%rbx
	mov	%rsp,%rsp
	mov	%rsp,%rbp
	mov	%rsp,%rsi
	mov	%rsp,%rdi
	mov	%rsp,%r8
	mov	%rsp,%r9
	mov	%rsp,%r10
	mov	%rsp,%r11
	mov	%rsp,%r12
	mov	%rsp,%r13
	mov	%rsp,%r14
	mov	%rsp,%r15
	nop
	mov	%rbp,%rax
	mov	%rbp,%rcx
	mov	%rbp,%rdx
	mov	%rbp,%rbx
	mov	%rbp,%rsp
	mov	%rbp,%rbp
	mov	%rbp,%rsi
	mov	%rbp,%rdi
	mov	%rbp,%r8
	mov	%rbp,%r9
	mov	%rbp,%r10
	mov	%rbp,%r11
	mov	%rbp,%r12
	mov	%rbp,%r13
	mov	%rbp,%r14
	mov	%rbp,%r15
	nop
	mov	%rsi,%rax
	mov	%rsi,%rcx
	mov	%rsi,%rdx
	mov	%rsi,%rbx
	mov	%rsi,%rsp
	mov	%rsi,%rbp
	mov	%rsi,%rsi
	mov	%rsi,%rdi
	mov	%rsi,%r8
	mov	%rsi,%r9
	mov	%rsi,%r10
	mov	%rsi,%r11
	mov	%rsi,%r12
	mov	%rsi,%r13
	mov	%rsi,%r14
	mov	%rsi,%r15
	nop
	mov	%rdi,%rax
	mov	%rdi,%rcx
	mov	%rdi,%rdx
	mov	%rdi,%rbx
	mov	%rdi,%rsp
	mov	%rdi,%rbp
	mov	%rdi,%rsi
	mov	%rdi,%rdi
	mov	%rdi,%r8
	mov	%rdi,%r9
	mov	%rdi,%r10
	mov	%rdi,%r11
	mov	%rdi,%r12
	mov	%rdi,%r13
	mov	%rdi,%r14
	mov	%rdi,%r15
	nop
	mov	%r8, %rax
	mov	%r8, %rcx
	mov	%r8, %rdx
	mov	%r8, %rbx
	mov	%r8, %rsp
	mov	%r8, %rbp
	mov	%r8, %rsi
	mov	%r8, %rdi
	mov	%r8, %r8
	mov	%r8, %r9
	mov	%r8, %r10
	mov	%r8, %r11
	mov	%r8, %r12
	mov	%r8, %r13
	mov	%r8, %r14
	mov	%r8, %r15
	nop
	mov	%r9, %rax
	mov	%r9, %rcx
	mov	%r9, %rdx
	mov	%r9, %rbx
	mov	%r9, %rsp
	mov	%r9, %rbp
	mov	%r9, %rsi
	mov	%r9, %rdi
	mov	%r9, %r8
	mov	%r9, %r9
	mov	%r9, %r10
	mov	%r9, %r11
	mov	%r9, %r12
	mov	%r9, %r13
	mov	%r9, %r14
	mov	%r9, %r15
	nop
	mov	%r10,%rax
	mov	%r10,%rcx
	mov	%r10,%rdx
	mov	%r10,%rbx
	mov	%r10,%rsp
	mov	%r10,%rbp
	mov	%r10,%rsi
	mov	%r10,%rdi
	mov	%r10,%r8
	mov	%r10,%r9
	mov	%r10,%r10
	mov	%r10,%r11
	mov	%r10,%r12
	mov	%r10,%r13
	mov	%r10,%r14
	mov	%r10,%r15
	nop
	mov	%r11,%rax
	mov	%r11,%rcx
	mov	%r11,%rdx
	mov	%r11,%rbx
	mov	%r11,%rsp
	mov	%r11,%rbp
	mov	%r11,%rsi
	mov	%r11,%rdi
	mov	%r11,%r8
	mov	%r11,%r9
	mov	%r11,%r10
	mov	%r11,%r11
	mov	%r11,%r12
	mov	%r11,%r13
	mov	%r11,%r14
	mov	%r11,%r15
	nop
	mov	%r12,%rax
	mov	%r12,%rcx
	mov	%r12,%rdx
	mov	%r12,%rbx
	mov	%r12,%rsp
	mov	%r12,%rbp
	mov	%r12,%rsi
	mov	%r12,%rdi
	mov	%r12,%r8
	mov	%r12,%r9
	mov	%r12,%r10
	mov	%r12,%r11
	mov	%r12,%r12
	mov	%r12,%r13
	mov	%r12,%r14
	mov	%r12,%r15
	nop
	mov	%r13,%rax
	mov	%r13,%rcx
	mov	%r13,%rdx
	mov	%r13,%rbx
	mov	%r13,%rsp
	mov	%r13,%rbp
	mov	%r13,%rsi
	mov	%r13,%rdi
	mov	%r13,%r8
	mov	%r13,%r9
	mov	%r13,%r10
	mov	%r13,%r11
	mov	%r13,%r12
	mov	%r13,%r13
	mov	%r13,%r14
	mov	%r13,%r15
	nop
	mov	%r14,%rax
	mov	%r14,%rcx
	mov	%r14,%rdx
	mov	%r14,%rbx
	mov	%r14,%rsp
	mov	%r14,%rbp
	mov	%r14,%rsi
	mov	%r14,%rdi
	mov	%r14,%r8
	mov	%r14,%r9
	mov	%r14,%r10
	mov	%r14,%r11
	mov	%r14,%r12
	mov	%r14,%r13
	mov	%r14,%r14
	mov	%r14,%r15
	nop
	mov	%r15,%rax
	mov	%r15,%rcx
	mov	%r15,%rdx
	mov	%r15,%rbx
	mov	%r15,%rsp
	mov	%r15,%rbp
	mov	%r15,%rsi
	mov	%r15,%rdi
	mov	%r15,%r8
	mov	%r15,%r9
	mov	%r15,%r10
	mov	%r15,%r11
	mov	%r15,%r12
	mov	%r15,%r13
	mov	%r15,%r14
	mov	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	MovMemReg
	.type	MovMemReg, @function
MovMemReg:
	.cfi_startproc
	mov	%rax,(%rax)
	mov	%rax,(%rcx)
	mov	%rax,(%rdx)
	mov	%rax,(%rbx)
	mov	%rax,(%rsp)
	mov	%rax,(%rbp)
	mov	%rax,(%rsi)
	mov	%rax,(%rdi)
	mov	%rax,(%r8)
	mov	%rax,(%r9)
	mov	%rax,(%r10)
	mov	%rax,(%r11)
	mov	%rax,(%r12)
	mov	%rax,(%r13)
	mov	%rax,(%r14)
	mov	%rax,(%r15)
	nop
	mov	%rcx,(%rax)
	mov	%rcx,(%rcx)
	mov	%rcx,(%rdx)
	mov	%rcx,(%rbx)
	mov	%rcx,(%rsp)
	mov	%rcx,(%rbp)
	mov	%rcx,(%rsi)
	mov	%rcx,(%rdi)
	mov	%rcx,(%r8)
	mov	%rcx,(%r9)
	mov	%rcx,(%r10)
	mov	%rcx,(%r11)
	mov	%rcx,(%r12)
	mov	%rcx,(%r13)
	mov	%rcx,(%r14)
	mov	%rcx,(%r15)
	nop
	mov	%rdx,(%rax)
	mov	%rdx,(%rcx)
	mov	%rdx,(%rdx)
	mov	%rdx,(%rbx)
	mov	%rdx,(%rsp)
	mov	%rdx,(%rbp)
	mov	%rdx,(%rsi)
	mov	%rdx,(%rdi)
	mov	%rdx,(%r8)
	mov	%rdx,(%r9)
	mov	%rdx,(%r10)
	mov	%rdx,(%r11)
	mov	%rdx,(%r12)
	mov	%rdx,(%r13)
	mov	%rdx,(%r14)
	mov	%rdx,(%r15)
	nop
	mov	%rbx,(%rax)
	mov	%rbx,(%rcx)
	mov	%rbx,(%rdx)
	mov	%rbx,(%rbx)
	mov	%rbx,(%rsp)
	mov	%rbx,(%rbp)
	mov	%rbx,(%rsi)
	mov	%rbx,(%rdi)
	mov	%rbx,(%r8)
	mov	%rbx,(%r9)
	mov	%rbx,(%r10)
	mov	%rbx,(%r11)
	mov	%rbx,(%r12)
	mov	%rbx,(%r13)
	mov	%rbx,(%r14)
	mov	%rbx,(%r15)
	nop
	mov	%rsp,(%rax)
	mov	%rsp,(%rcx)
	mov	%rsp,(%rdx)
	mov	%rsp,(%rbx)
	mov	%rsp,(%rsp)
	mov	%rsp,(%rbp)
	mov	%rsp,(%rsi)
	mov	%rsp,(%rdi)
	mov	%rsp,(%r8)
	mov	%rsp,(%r9)
	mov	%rsp,(%r10)
	mov	%rsp,(%r11)
	mov	%rsp,(%r12)
	mov	%rsp,(%r13)
	mov	%rsp,(%r14)
	mov	%rsp,(%r15)
	nop
	mov	%rbp,(%rax)
	mov	%rbp,(%rcx)
	mov	%rbp,(%rdx)
	mov	%rbp,(%rbx)
	mov	%rbp,(%rsp)
	mov	%rbp,(%rbp)
	mov	%rbp,(%rsi)
	mov	%rbp,(%rdi)
	mov	%rbp,(%r8)
	mov	%rbp,(%r9)
	mov	%rbp,(%r10)
	mov	%rbp,(%r11)
	mov	%rbp,(%r12)
	mov	%rbp,(%r13)
	mov	%rbp,(%r14)
	mov	%rbp,(%r15)
	nop
	mov	%rsi,(%rax)
	mov	%rsi,(%rcx)
	mov	%rsi,(%rdx)
	mov	%rsi,(%rbx)
	mov	%rsi,(%rsp)
	mov	%rsi,(%rbp)
	mov	%rsi,(%rsi)
	mov	%rsi,(%rdi)
	mov	%rsi,(%r8)
	mov	%rsi,(%r9)
	mov	%rsi,(%r10)
	mov	%rsi,(%r11)
	mov	%rsi,(%r12)
	mov	%rsi,(%r13)
	mov	%rsi,(%r14)
	mov	%rsi,(%r15)
	nop
	mov	%rdi,(%rax)
	mov	%rdi,(%rcx)
	mov	%rdi,(%rdx)
	mov	%rdi,(%rbx)
	mov	%rdi,(%rsp)
	mov	%rdi,(%rbp)
	mov	%rdi,(%rsi)
	mov	%rdi,(%rdi)
	mov	%rdi,(%r8)
	mov	%rdi,(%r9)
	mov	%rdi,(%r10)
	mov	%rdi,(%r11)
	mov	%rdi,(%r12)
	mov	%rdi,(%r13)
	mov	%rdi,(%r14)
	mov	%rdi,(%r15)
	nop
	mov	%r8, (%rax)
	mov	%r8, (%rcx)
	mov	%r8, (%rdx)
	mov	%r8, (%rbx)
	mov	%r8, (%rsp)
	mov	%r8, (%rbp)
	mov	%r8, (%rsi)
	mov	%r8, (%rdi)
	mov	%r8, (%r8)
	mov	%r8, (%r9)
	mov	%r8, (%r10)
	mov	%r8, (%r11)
	mov	%r8, (%r12)
	mov	%r8, (%r13)
	mov	%r8, (%r14)
	mov	%r8, (%r15)
	nop
	mov	%r9, (%rax)
	mov	%r9, (%rcx)
	mov	%r9, (%rdx)
	mov	%r9, (%rbx)
	mov	%r9, (%rsp)
	mov	%r9, (%rbp)
	mov	%r9, (%rsi)
	mov	%r9, (%rdi)
	mov	%r9, (%r8)
	mov	%r9, (%r9)
	mov	%r9, (%r10)
	mov	%r9, (%r11)
	mov	%r9, (%r12)
	mov	%r9, (%r13)
	mov	%r9, (%r14)
	mov	%r9, (%r15)
	nop
	mov	%r10,(%rax)
	mov	%r10,(%rcx)
	mov	%r10,(%rdx)
	mov	%r10,(%rbx)
	mov	%r10,(%rsp)
	mov	%r10,(%rbp)
	mov	%r10,(%rsi)
	mov	%r10,(%rdi)
	mov	%r10,(%r8)
	mov	%r10,(%r9)
	mov	%r10,(%r10)
	mov	%r10,(%r11)
	mov	%r10,(%r12)
	mov	%r10,(%r13)
	mov	%r10,(%r14)
	mov	%r10,(%r15)
	nop
	mov	%r11,(%rax)
	mov	%r11,(%rcx)
	mov	%r11,(%rdx)
	mov	%r11,(%rbx)
	mov	%r11,(%rsp)
	mov	%r11,(%rbp)
	mov	%r11,(%rsi)
	mov	%r11,(%rdi)
	mov	%r11,(%r8)
	mov	%r11,(%r9)
	mov	%r11,(%r10)
	mov	%r11,(%r11)
	mov	%r11,(%r12)
	mov	%r11,(%r13)
	mov	%r11,(%r14)
	mov	%r11,(%r15)
	nop
	mov	%r12,(%rax)
	mov	%r12,(%rcx)
	mov	%r12,(%rdx)
	mov	%r12,(%rbx)
	mov	%r12,(%rsp)
	mov	%r12,(%rbp)
	mov	%r12,(%rsi)
	mov	%r12,(%rdi)
	mov	%r12,(%r8)
	mov	%r12,(%r9)
	mov	%r12,(%r10)
	mov	%r12,(%r11)
	mov	%r12,(%r12)
	mov	%r12,(%r13)
	mov	%r12,(%r14)
	mov	%r12,(%r15)
	nop
	mov	%r13,(%rax)
	mov	%r13,(%rcx)
	mov	%r13,(%rdx)
	mov	%r13,(%rbx)
	mov	%r13,(%rsp)
	mov	%r13,(%rbp)
	mov	%r13,(%rsi)
	mov	%r13,(%rdi)
	mov	%r13,(%r8)
	mov	%r13,(%r9)
	mov	%r13,(%r10)
	mov	%r13,(%r11)
	mov	%r13,(%r12)
	mov	%r13,(%r13)
	mov	%r13,(%r14)
	mov	%r13,(%r15)
	nop
	mov	%r14,(%rax)
	mov	%r14,(%rcx)
	mov	%r14,(%rdx)
	mov	%r14,(%rbx)
	mov	%r14,(%rsp)
	mov	%r14,(%rbp)
	mov	%r14,(%rsi)
	mov	%r14,(%rdi)
	mov	%r14,(%r8)
	mov	%r14,(%r9)
	mov	%r14,(%r10)
	mov	%r14,(%r11)
	mov	%r14,(%r12)
	mov	%r14,(%r13)
	mov	%r14,(%r14)
	mov	%r14,(%r15)
	nop
	mov	%r15,(%rax)
	mov	%r15,(%rcx)
	mov	%r15,(%rdx)
	mov	%r15,(%rbx)
	mov	%r15,(%rsp)
	mov	%r15,(%rbp)
	mov	%r15,(%rsi)
	mov	%r15,(%rdi)
	mov	%r15,(%r8)
	mov	%r15,(%r9)
	mov	%r15,(%r10)
	mov	%r15,(%r11)
	mov	%r15,(%r12)
	mov	%r15,(%r13)
	mov	%r15,(%r14)
	mov	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	MovMem8Reg
	.type	MovMem8Reg, @function
MovMem8Reg:
	.cfi_startproc
	mov	%rax,0x7f(%rax)
	mov	%rax,0x7f(%rcx)
	mov	%rax,0x7f(%rdx)
	mov	%rax,0x7f(%rbx)
	mov	%rax,0x7f(%rsp)
	mov	%rax,0x7f(%rbp)
	mov	%rax,0x7f(%rsi)
	mov	%rax,0x7f(%rdi)
	mov	%rax,0x7f(%r8)
	mov	%rax,0x7f(%r9)
	mov	%rax,0x7f(%r10)
	mov	%rax,0x7f(%r11)
	mov	%rax,0x7f(%r12)
	mov	%rax,0x7f(%r13)
	mov	%rax,0x7f(%r14)
	mov	%rax,0x7f(%r15)
	nop
	mov	%rcx,0x7f(%rax)
	mov	%rcx,0x7f(%rcx)
	mov	%rcx,0x7f(%rdx)
	mov	%rcx,0x7f(%rbx)
	mov	%rcx,0x7f(%rsp)
	mov	%rcx,0x7f(%rbp)
	mov	%rcx,0x7f(%rsi)
	mov	%rcx,0x7f(%rdi)
	mov	%rcx,0x7f(%r8)
	mov	%rcx,0x7f(%r9)
	mov	%rcx,0x7f(%r10)
	mov	%rcx,0x7f(%r11)
	mov	%rcx,0x7f(%r12)
	mov	%rcx,0x7f(%r13)
	mov	%rcx,0x7f(%r14)
	mov	%rcx,0x7f(%r15)
	nop
	mov	%rdx,0x7f(%rax)
	mov	%rdx,0x7f(%rcx)
	mov	%rdx,0x7f(%rdx)
	mov	%rdx,0x7f(%rbx)
	mov	%rdx,0x7f(%rsp)
	mov	%rdx,0x7f(%rbp)
	mov	%rdx,0x7f(%rsi)
	mov	%rdx,0x7f(%rdi)
	mov	%rdx,0x7f(%r8)
	mov	%rdx,0x7f(%r9)
	mov	%rdx,0x7f(%r10)
	mov	%rdx,0x7f(%r11)
	mov	%rdx,0x7f(%r12)
	mov	%rdx,0x7f(%r13)
	mov	%rdx,0x7f(%r14)
	mov	%rdx,0x7f(%r15)
	nop
	mov	%rbx,0x7f(%rax)
	mov	%rbx,0x7f(%rcx)
	mov	%rbx,0x7f(%rdx)
	mov	%rbx,0x7f(%rbx)
	mov	%rbx,0x7f(%rsp)
	mov	%rbx,0x7f(%rbp)
	mov	%rbx,0x7f(%rsi)
	mov	%rbx,0x7f(%rdi)
	mov	%rbx,0x7f(%r8)
	mov	%rbx,0x7f(%r9)
	mov	%rbx,0x7f(%r10)
	mov	%rbx,0x7f(%r11)
	mov	%rbx,0x7f(%r12)
	mov	%rbx,0x7f(%r13)
	mov	%rbx,0x7f(%r14)
	mov	%rbx,0x7f(%r15)
	nop
	mov	%rsp,0x7f(%rax)
	mov	%rsp,0x7f(%rcx)
	mov	%rsp,0x7f(%rdx)
	mov	%rsp,0x7f(%rbx)
	mov	%rsp,0x7f(%rsp)
	mov	%rsp,0x7f(%rbp)
	mov	%rsp,0x7f(%rsi)
	mov	%rsp,0x7f(%rdi)
	mov	%rsp,0x7f(%r8)
	mov	%rsp,0x7f(%r9)
	mov	%rsp,0x7f(%r10)
	mov	%rsp,0x7f(%r11)
	mov	%rsp,0x7f(%r12)
	mov	%rsp,0x7f(%r13)
	mov	%rsp,0x7f(%r14)
	mov	%rsp,0x7f(%r15)
	nop
	mov	%rbp,0x7f(%rax)
	mov	%rbp,0x7f(%rcx)
	mov	%rbp,0x7f(%rdx)
	mov	%rbp,0x7f(%rbx)
	mov	%rbp,0x7f(%rsp)
	mov	%rbp,0x7f(%rbp)
	mov	%rbp,0x7f(%rsi)
	mov	%rbp,0x7f(%rdi)
	mov	%rbp,0x7f(%r8)
	mov	%rbp,0x7f(%r9)
	mov	%rbp,0x7f(%r10)
	mov	%rbp,0x7f(%r11)
	mov	%rbp,0x7f(%r12)
	mov	%rbp,0x7f(%r13)
	mov	%rbp,0x7f(%r14)
	mov	%rbp,0x7f(%r15)
	nop
	mov	%rsi,0x7f(%rax)
	mov	%rsi,0x7f(%rcx)
	mov	%rsi,0x7f(%rdx)
	mov	%rsi,0x7f(%rbx)
	mov	%rsi,0x7f(%rsp)
	mov	%rsi,0x7f(%rbp)
	mov	%rsi,0x7f(%rsi)
	mov	%rsi,0x7f(%rdi)
	mov	%rsi,0x7f(%r8)
	mov	%rsi,0x7f(%r9)
	mov	%rsi,0x7f(%r10)
	mov	%rsi,0x7f(%r11)
	mov	%rsi,0x7f(%r12)
	mov	%rsi,0x7f(%r13)
	mov	%rsi,0x7f(%r14)
	mov	%rsi,0x7f(%r15)
	nop
	mov	%rdi,0x7f(%rax)
	mov	%rdi,0x7f(%rcx)
	mov	%rdi,0x7f(%rdx)
	mov	%rdi,0x7f(%rbx)
	mov	%rdi,0x7f(%rsp)
	mov	%rdi,0x7f(%rbp)
	mov	%rdi,0x7f(%rsi)
	mov	%rdi,0x7f(%rdi)
	mov	%rdi,0x7f(%r8)
	mov	%rdi,0x7f(%r9)
	mov	%rdi,0x7f(%r10)
	mov	%rdi,0x7f(%r11)
	mov	%rdi,0x7f(%r12)
	mov	%rdi,0x7f(%r13)
	mov	%rdi,0x7f(%r14)
	mov	%rdi,0x7f(%r15)
	nop
	mov	%r8, 0x7f(%rax)
	mov	%r8, 0x7f(%rcx)
	mov	%r8, 0x7f(%rdx)
	mov	%r8, 0x7f(%rbx)
	mov	%r8, 0x7f(%rsp)
	mov	%r8, 0x7f(%rbp)
	mov	%r8, 0x7f(%rsi)
	mov	%r8, 0x7f(%rdi)
	mov	%r8, 0x7f(%r8)
	mov	%r8, 0x7f(%r9)
	mov	%r8, 0x7f(%r10)
	mov	%r8, 0x7f(%r11)
	mov	%r8, 0x7f(%r12)
	mov	%r8, 0x7f(%r13)
	mov	%r8, 0x7f(%r14)
	mov	%r8, 0x7f(%r15)
	nop
	mov	%r9, 0x7f(%rax)
	mov	%r9, 0x7f(%rcx)
	mov	%r9, 0x7f(%rdx)
	mov	%r9, 0x7f(%rbx)
	mov	%r9, 0x7f(%rsp)
	mov	%r9, 0x7f(%rbp)
	mov	%r9, 0x7f(%rsi)
	mov	%r9, 0x7f(%rdi)
	mov	%r9, 0x7f(%r8)
	mov	%r9, 0x7f(%r9)
	mov	%r9, 0x7f(%r10)
	mov	%r9, 0x7f(%r11)
	mov	%r9, 0x7f(%r12)
	mov	%r9, 0x7f(%r13)
	mov	%r9, 0x7f(%r14)
	mov	%r9, 0x7f(%r15)
	nop
	mov	%r10,0x7f(%rax)
	mov	%r10,0x7f(%rcx)
	mov	%r10,0x7f(%rdx)
	mov	%r10,0x7f(%rbx)
	mov	%r10,0x7f(%rsp)
	mov	%r10,0x7f(%rbp)
	mov	%r10,0x7f(%rsi)
	mov	%r10,0x7f(%rdi)
	mov	%r10,0x7f(%r8)
	mov	%r10,0x7f(%r9)
	mov	%r10,0x7f(%r10)
	mov	%r10,0x7f(%r11)
	mov	%r10,0x7f(%r12)
	mov	%r10,0x7f(%r13)
	mov	%r10,0x7f(%r14)
	mov	%r10,0x7f(%r15)
	nop
	mov	%r11,0x7f(%rax)
	mov	%r11,0x7f(%rcx)
	mov	%r11,0x7f(%rdx)
	mov	%r11,0x7f(%rbx)
	mov	%r11,0x7f(%rsp)
	mov	%r11,0x7f(%rbp)
	mov	%r11,0x7f(%rsi)
	mov	%r11,0x7f(%rdi)
	mov	%r11,0x7f(%r8)
	mov	%r11,0x7f(%r9)
	mov	%r11,0x7f(%r10)
	mov	%r11,0x7f(%r11)
	mov	%r11,0x7f(%r12)
	mov	%r11,0x7f(%r13)
	mov	%r11,0x7f(%r14)
	mov	%r11,0x7f(%r15)
	nop
	mov	%r12,0x7f(%rax)
	mov	%r12,0x7f(%rcx)
	mov	%r12,0x7f(%rdx)
	mov	%r12,0x7f(%rbx)
	mov	%r12,0x7f(%rsp)
	mov	%r12,0x7f(%rbp)
	mov	%r12,0x7f(%rsi)
	mov	%r12,0x7f(%rdi)
	mov	%r12,0x7f(%r8)
	mov	%r12,0x7f(%r9)
	mov	%r12,0x7f(%r10)
	mov	%r12,0x7f(%r11)
	mov	%r12,0x7f(%r12)
	mov	%r12,0x7f(%r13)
	mov	%r12,0x7f(%r14)
	mov	%r12,0x7f(%r15)
	nop
	mov	%r13,0x7f(%rax)
	mov	%r13,0x7f(%rcx)
	mov	%r13,0x7f(%rdx)
	mov	%r13,0x7f(%rbx)
	mov	%r13,0x7f(%rsp)
	mov	%r13,0x7f(%rbp)
	mov	%r13,0x7f(%rsi)
	mov	%r13,0x7f(%rdi)
	mov	%r13,0x7f(%r8)
	mov	%r13,0x7f(%r9)
	mov	%r13,0x7f(%r10)
	mov	%r13,0x7f(%r11)
	mov	%r13,0x7f(%r12)
	mov	%r13,0x7f(%r13)
	mov	%r13,0x7f(%r14)
	mov	%r13,0x7f(%r15)
	nop
	mov	%r14,0x7f(%rax)
	mov	%r14,0x7f(%rcx)
	mov	%r14,0x7f(%rdx)
	mov	%r14,0x7f(%rbx)
	mov	%r14,0x7f(%rsp)
	mov	%r14,0x7f(%rbp)
	mov	%r14,0x7f(%rsi)
	mov	%r14,0x7f(%rdi)
	mov	%r14,0x7f(%r8)
	mov	%r14,0x7f(%r9)
	mov	%r14,0x7f(%r10)
	mov	%r14,0x7f(%r11)
	mov	%r14,0x7f(%r12)
	mov	%r14,0x7f(%r13)
	mov	%r14,0x7f(%r14)
	mov	%r14,0x7f(%r15)
	nop
	mov	%r15,0x7f(%rax)
	mov	%r15,0x7f(%rcx)
	mov	%r15,0x7f(%rdx)
	mov	%r15,0x7f(%rbx)
	mov	%r15,0x7f(%rsp)
	mov	%r15,0x7f(%rbp)
	mov	%r15,0x7f(%rsi)
	mov	%r15,0x7f(%rdi)
	mov	%r15,0x7f(%r8)
	mov	%r15,0x7f(%r9)
	mov	%r15,0x7f(%r10)
	mov	%r15,0x7f(%r11)
	mov	%r15,0x7f(%r12)
	mov	%r15,0x7f(%r13)
	mov	%r15,0x7f(%r14)
	mov	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	MovMem32Reg
	.type	MovMem32Reg, @function
MovMem32Reg:
	.cfi_startproc
	mov	%rax,0x7f563412(%rax)
	mov	%rax,0x7f563412(%rcx)
	mov	%rax,0x7f563412(%rdx)
	mov	%rax,0x7f563412(%rbx)
	mov	%rax,0x7f563412(%rsp)
	mov	%rax,0x7f563412(%rbp)
	mov	%rax,0x7f563412(%rsi)
	mov	%rax,0x7f563412(%rdi)
	mov	%rax,0x7f563412(%r8)
	mov	%rax,0x7f563412(%r9)
	mov	%rax,0x7f563412(%r10)
	mov	%rax,0x7f563412(%r11)
	mov	%rax,0x7f563412(%r12)
	mov	%rax,0x7f563412(%r13)
	mov	%rax,0x7f563412(%r14)
	mov	%rax,0x7f563412(%r15)
	nop
	mov	%rcx,0x7f563412(%rax)
	mov	%rcx,0x7f563412(%rcx)
	mov	%rcx,0x7f563412(%rdx)
	mov	%rcx,0x7f563412(%rbx)
	mov	%rcx,0x7f563412(%rsp)
	mov	%rcx,0x7f563412(%rbp)
	mov	%rcx,0x7f563412(%rsi)
	mov	%rcx,0x7f563412(%rdi)
	mov	%rcx,0x7f563412(%r8)
	mov	%rcx,0x7f563412(%r9)
	mov	%rcx,0x7f563412(%r10)
	mov	%rcx,0x7f563412(%r11)
	mov	%rcx,0x7f563412(%r12)
	mov	%rcx,0x7f563412(%r13)
	mov	%rcx,0x7f563412(%r14)
	mov	%rcx,0x7f563412(%r15)
	nop
	mov	%rdx,0x7f563412(%rax)
	mov	%rdx,0x7f563412(%rcx)
	mov	%rdx,0x7f563412(%rdx)
	mov	%rdx,0x7f563412(%rbx)
	mov	%rdx,0x7f563412(%rsp)
	mov	%rdx,0x7f563412(%rbp)
	mov	%rdx,0x7f563412(%rsi)
	mov	%rdx,0x7f563412(%rdi)
	mov	%rdx,0x7f563412(%r8)
	mov	%rdx,0x7f563412(%r9)
	mov	%rdx,0x7f563412(%r10)
	mov	%rdx,0x7f563412(%r11)
	mov	%rdx,0x7f563412(%r12)
	mov	%rdx,0x7f563412(%r13)
	mov	%rdx,0x7f563412(%r14)
	mov	%rdx,0x7f563412(%r15)
	nop
	mov	%rbx,0x7f563412(%rax)
	mov	%rbx,0x7f563412(%rcx)
	mov	%rbx,0x7f563412(%rdx)
	mov	%rbx,0x7f563412(%rbx)
	mov	%rbx,0x7f563412(%rsp)
	mov	%rbx,0x7f563412(%rbp)
	mov	%rbx,0x7f563412(%rsi)
	mov	%rbx,0x7f563412(%rdi)
	mov	%rbx,0x7f563412(%r8)
	mov	%rbx,0x7f563412(%r9)
	mov	%rbx,0x7f563412(%r10)
	mov	%rbx,0x7f563412(%r11)
	mov	%rbx,0x7f563412(%r12)
	mov	%rbx,0x7f563412(%r13)
	mov	%rbx,0x7f563412(%r14)
	mov	%rbx,0x7f563412(%r15)
	nop
	mov	%rsp,0x7f563412(%rax)
	mov	%rsp,0x7f563412(%rcx)
	mov	%rsp,0x7f563412(%rdx)
	mov	%rsp,0x7f563412(%rbx)
	mov	%rsp,0x7f563412(%rsp)
	mov	%rsp,0x7f563412(%rbp)
	mov	%rsp,0x7f563412(%rsi)
	mov	%rsp,0x7f563412(%rdi)
	mov	%rsp,0x7f563412(%r8)
	mov	%rsp,0x7f563412(%r9)
	mov	%rsp,0x7f563412(%r10)
	mov	%rsp,0x7f563412(%r11)
	mov	%rsp,0x7f563412(%r12)
	mov	%rsp,0x7f563412(%r13)
	mov	%rsp,0x7f563412(%r14)
	mov	%rsp,0x7f563412(%r15)
	nop
	mov	%rbp,0x7f563412(%rax)
	mov	%rbp,0x7f563412(%rcx)
	mov	%rbp,0x7f563412(%rdx)
	mov	%rbp,0x7f563412(%rbx)
	mov	%rbp,0x7f563412(%rsp)
	mov	%rbp,0x7f563412(%rbp)
	mov	%rbp,0x7f563412(%rsi)
	mov	%rbp,0x7f563412(%rdi)
	mov	%rbp,0x7f563412(%r8)
	mov	%rbp,0x7f563412(%r9)
	mov	%rbp,0x7f563412(%r10)
	mov	%rbp,0x7f563412(%r11)
	mov	%rbp,0x7f563412(%r12)
	mov	%rbp,0x7f563412(%r13)
	mov	%rbp,0x7f563412(%r14)
	mov	%rbp,0x7f563412(%r15)
	nop
	mov	%rsi,0x7f563412(%rax)
	mov	%rsi,0x7f563412(%rcx)
	mov	%rsi,0x7f563412(%rdx)
	mov	%rsi,0x7f563412(%rbx)
	mov	%rsi,0x7f563412(%rsp)
	mov	%rsi,0x7f563412(%rbp)
	mov	%rsi,0x7f563412(%rsi)
	mov	%rsi,0x7f563412(%rdi)
	mov	%rsi,0x7f563412(%r8)
	mov	%rsi,0x7f563412(%r9)
	mov	%rsi,0x7f563412(%r10)
	mov	%rsi,0x7f563412(%r11)
	mov	%rsi,0x7f563412(%r12)
	mov	%rsi,0x7f563412(%r13)
	mov	%rsi,0x7f563412(%r14)
	mov	%rsi,0x7f563412(%r15)
	nop
	mov	%rdi,0x7f563412(%rax)
	mov	%rdi,0x7f563412(%rcx)
	mov	%rdi,0x7f563412(%rdx)
	mov	%rdi,0x7f563412(%rbx)
	mov	%rdi,0x7f563412(%rsp)
	mov	%rdi,0x7f563412(%rbp)
	mov	%rdi,0x7f563412(%rsi)
	mov	%rdi,0x7f563412(%rdi)
	mov	%rdi,0x7f563412(%r8)
	mov	%rdi,0x7f563412(%r9)
	mov	%rdi,0x7f563412(%r10)
	mov	%rdi,0x7f563412(%r11)
	mov	%rdi,0x7f563412(%r12)
	mov	%rdi,0x7f563412(%r13)
	mov	%rdi,0x7f563412(%r14)
	mov	%rdi,0x7f563412(%r15)
	nop
	mov	%r8, 0x7f563412(%rax)
	mov	%r8, 0x7f563412(%rcx)
	mov	%r8, 0x7f563412(%rdx)
	mov	%r8, 0x7f563412(%rbx)
	mov	%r8, 0x7f563412(%rsp)
	mov	%r8, 0x7f563412(%rbp)
	mov	%r8, 0x7f563412(%rsi)
	mov	%r8, 0x7f563412(%rdi)
	mov	%r8, 0x7f563412(%r8)
	mov	%r8, 0x7f563412(%r9)
	mov	%r8, 0x7f563412(%r10)
	mov	%r8, 0x7f563412(%r11)
	mov	%r8, 0x7f563412(%r12)
	mov	%r8, 0x7f563412(%r13)
	mov	%r8, 0x7f563412(%r14)
	mov	%r8, 0x7f563412(%r15)
	nop
	mov	%r9, 0x7f563412(%rax)
	mov	%r9, 0x7f563412(%rcx)
	mov	%r9, 0x7f563412(%rdx)
	mov	%r9, 0x7f563412(%rbx)
	mov	%r9, 0x7f563412(%rsp)
	mov	%r9, 0x7f563412(%rbp)
	mov	%r9, 0x7f563412(%rsi)
	mov	%r9, 0x7f563412(%rdi)
	mov	%r9, 0x7f563412(%r8)
	mov	%r9, 0x7f563412(%r9)
	mov	%r9, 0x7f563412(%r10)
	mov	%r9, 0x7f563412(%r11)
	mov	%r9, 0x7f563412(%r12)
	mov	%r9, 0x7f563412(%r13)
	mov	%r9, 0x7f563412(%r14)
	mov	%r9, 0x7f563412(%r15)
	nop
	mov	%r10,0x7f563412(%rax)
	mov	%r10,0x7f563412(%rcx)
	mov	%r10,0x7f563412(%rdx)
	mov	%r10,0x7f563412(%rbx)
	mov	%r10,0x7f563412(%rsp)
	mov	%r10,0x7f563412(%rbp)
	mov	%r10,0x7f563412(%rsi)
	mov	%r10,0x7f563412(%rdi)
	mov	%r10,0x7f563412(%r8)
	mov	%r10,0x7f563412(%r9)
	mov	%r10,0x7f563412(%r10)
	mov	%r10,0x7f563412(%r11)
	mov	%r10,0x7f563412(%r12)
	mov	%r10,0x7f563412(%r13)
	mov	%r10,0x7f563412(%r14)
	mov	%r10,0x7f563412(%r15)
	nop
	mov	%r11,0x7f563412(%rax)
	mov	%r11,0x7f563412(%rcx)
	mov	%r11,0x7f563412(%rdx)
	mov	%r11,0x7f563412(%rbx)
	mov	%r11,0x7f563412(%rsp)
	mov	%r11,0x7f563412(%rbp)
	mov	%r11,0x7f563412(%rsi)
	mov	%r11,0x7f563412(%rdi)
	mov	%r11,0x7f563412(%r8)
	mov	%r11,0x7f563412(%r9)
	mov	%r11,0x7f563412(%r10)
	mov	%r11,0x7f563412(%r11)
	mov	%r11,0x7f563412(%r12)
	mov	%r11,0x7f563412(%r13)
	mov	%r11,0x7f563412(%r14)
	mov	%r11,0x7f563412(%r15)
	nop
	mov	%r12,0x7f563412(%rax)
	mov	%r12,0x7f563412(%rcx)
	mov	%r12,0x7f563412(%rdx)
	mov	%r12,0x7f563412(%rbx)
	mov	%r12,0x7f563412(%rsp)
	mov	%r12,0x7f563412(%rbp)
	mov	%r12,0x7f563412(%rsi)
	mov	%r12,0x7f563412(%rdi)
	mov	%r12,0x7f563412(%r8)
	mov	%r12,0x7f563412(%r9)
	mov	%r12,0x7f563412(%r10)
	mov	%r12,0x7f563412(%r11)
	mov	%r12,0x7f563412(%r12)
	mov	%r12,0x7f563412(%r13)
	mov	%r12,0x7f563412(%r14)
	mov	%r12,0x7f563412(%r15)
	nop
	mov	%r13,0x7f563412(%rax)
	mov	%r13,0x7f563412(%rcx)
	mov	%r13,0x7f563412(%rdx)
	mov	%r13,0x7f563412(%rbx)
	mov	%r13,0x7f563412(%rsp)
	mov	%r13,0x7f563412(%rbp)
	mov	%r13,0x7f563412(%rsi)
	mov	%r13,0x7f563412(%rdi)
	mov	%r13,0x7f563412(%r8)
	mov	%r13,0x7f563412(%r9)
	mov	%r13,0x7f563412(%r10)
	mov	%r13,0x7f563412(%r11)
	mov	%r13,0x7f563412(%r12)
	mov	%r13,0x7f563412(%r13)
	mov	%r13,0x7f563412(%r14)
	mov	%r13,0x7f563412(%r15)
	nop
	mov	%r14,0x7f563412(%rax)
	mov	%r14,0x7f563412(%rcx)
	mov	%r14,0x7f563412(%rdx)
	mov	%r14,0x7f563412(%rbx)
	mov	%r14,0x7f563412(%rsp)
	mov	%r14,0x7f563412(%rbp)
	mov	%r14,0x7f563412(%rsi)
	mov	%r14,0x7f563412(%rdi)
	mov	%r14,0x7f563412(%r8)
	mov	%r14,0x7f563412(%r9)
	mov	%r14,0x7f563412(%r10)
	mov	%r14,0x7f563412(%r11)
	mov	%r14,0x7f563412(%r12)
	mov	%r14,0x7f563412(%r13)
	mov	%r14,0x7f563412(%r14)
	mov	%r14,0x7f563412(%r15)
	nop
	mov	%r15,0x7f563412(%rax)
	mov	%r15,0x7f563412(%rcx)
	mov	%r15,0x7f563412(%rdx)
	mov	%r15,0x7f563412(%rbx)
	mov	%r15,0x7f563412(%rsp)
	mov	%r15,0x7f563412(%rbp)
	mov	%r15,0x7f563412(%rsi)
	mov	%r15,0x7f563412(%rdi)
	mov	%r15,0x7f563412(%r8)
	mov	%r15,0x7f563412(%r9)
	mov	%r15,0x7f563412(%r10)
	mov	%r15,0x7f563412(%r11)
	mov	%r15,0x7f563412(%r12)
	mov	%r15,0x7f563412(%r13)
	mov	%r15,0x7f563412(%r14)
	mov	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	MovRegMem
	.type	MovRegMem, @function
MovRegMem:
	.cfi_startproc
	mov	(%rax),%rax
	mov	(%rax),%rcx
	mov	(%rax),%rdx
	mov	(%rax),%rbx
	mov	(%rax),%rsp
	mov	(%rax),%rbp
	mov	(%rax),%rsi
	mov	(%rax),%rdi
	mov	(%rax),%r8
	mov	(%rax),%r9
	mov	(%rax),%r10
	mov	(%rax),%r11
	mov	(%rax),%r12
	mov	(%rax),%r13
	mov	(%rax),%r14
	mov	(%rax),%r15
	nop
	mov	(%rcx),%rax
	mov	(%rcx),%rcx
	mov	(%rcx),%rdx
	mov	(%rcx),%rbx
	mov	(%rcx),%rsp
	mov	(%rcx),%rbp
	mov	(%rcx),%rsi
	mov	(%rcx),%rdi
	mov	(%rcx),%r8
	mov	(%rcx),%r9
	mov	(%rcx),%r10
	mov	(%rcx),%r11
	mov	(%rcx),%r12
	mov	(%rcx),%r13
	mov	(%rcx),%r14
	mov	(%rcx),%r15
	nop
	mov	(%rdx),%rax
	mov	(%rdx),%rcx
	mov	(%rdx),%rdx
	mov	(%rdx),%rbx
	mov	(%rdx),%rsp
	mov	(%rdx),%rbp
	mov	(%rdx),%rsi
	mov	(%rdx),%rdi
	mov	(%rdx),%r8
	mov	(%rdx),%r9
	mov	(%rdx),%r10
	mov	(%rdx),%r11
	mov	(%rdx),%r12
	mov	(%rdx),%r13
	mov	(%rdx),%r14
	mov	(%rdx),%r15
	nop
	mov	(%rbx),%rax
	mov	(%rbx),%rcx
	mov	(%rbx),%rdx
	mov	(%rbx),%rbx
	mov	(%rbx),%rsp
	mov	(%rbx),%rbp
	mov	(%rbx),%rsi
	mov	(%rbx),%rdi
	mov	(%rbx),%r8
	mov	(%rbx),%r9
	mov	(%rbx),%r10
	mov	(%rbx),%r11
	mov	(%rbx),%r12
	mov	(%rbx),%r13
	mov	(%rbx),%r14
	mov	(%rbx),%r15
	nop
	mov	(%rsp),%rax
	mov	(%rsp),%rcx
	mov	(%rsp),%rdx
	mov	(%rsp),%rbx
	mov	(%rsp),%rsp
	mov	(%rsp),%rbp
	mov	(%rsp),%rsi
	mov	(%rsp),%rdi
	mov	(%rsp),%r8
	mov	(%rsp),%r9
	mov	(%rsp),%r10
	mov	(%rsp),%r11
	mov	(%rsp),%r12
	mov	(%rsp),%r13
	mov	(%rsp),%r14
	mov	(%rsp),%r15
	nop
	mov	(%rbp),%rax
	mov	(%rbp),%rcx
	mov	(%rbp),%rdx
	mov	(%rbp),%rbx
	mov	(%rbp),%rsp
	mov	(%rbp),%rbp
	mov	(%rbp),%rsi
	mov	(%rbp),%rdi
	mov	(%rbp),%r8
	mov	(%rbp),%r9
	mov	(%rbp),%r10
	mov	(%rbp),%r11
	mov	(%rbp),%r12
	mov	(%rbp),%r13
	mov	(%rbp),%r14
	mov	(%rbp),%r15
	nop
	mov	(%rsi),%rax
	mov	(%rsi),%rcx
	mov	(%rsi),%rdx
	mov	(%rsi),%rbx
	mov	(%rsi),%rsp
	mov	(%rsi),%rbp
	mov	(%rsi),%rsi
	mov	(%rsi),%rdi
	mov	(%rsi),%r8
	mov	(%rsi),%r9
	mov	(%rsi),%r10
	mov	(%rsi),%r11
	mov	(%rsi),%r12
	mov	(%rsi),%r13
	mov	(%rsi),%r14
	mov	(%rsi),%r15
	nop
	mov	(%rdi),%rax
	mov	(%rdi),%rcx
	mov	(%rdi),%rdx
	mov	(%rdi),%rbx
	mov	(%rdi),%rsp
	mov	(%rdi),%rbp
	mov	(%rdi),%rsi
	mov	(%rdi),%rdi
	mov	(%rdi),%r8
	mov	(%rdi),%r9
	mov	(%rdi),%r10
	mov	(%rdi),%r11
	mov	(%rdi),%r12
	mov	(%rdi),%r13
	mov	(%rdi),%r14
	mov	(%rdi),%r15
	nop
	mov	(%r8), %rax
	mov	(%r8), %rcx
	mov	(%r8), %rdx
	mov	(%r8), %rbx
	mov	(%r8), %rsp
	mov	(%r8), %rbp
	mov	(%r8), %rsi
	mov	(%r8), %rdi
	mov	(%r8), %r8
	mov	(%r8), %r9
	mov	(%r8), %r10
	mov	(%r8), %r11
	mov	(%r8), %r12
	mov	(%r8), %r13
	mov	(%r8), %r14
	mov	(%r8), %r15
	nop
	mov	(%r9), %rax
	mov	(%r9), %rcx
	mov	(%r9), %rdx
	mov	(%r9), %rbx
	mov	(%r9), %rsp
	mov	(%r9), %rbp
	mov	(%r9), %rsi
	mov	(%r9), %rdi
	mov	(%r9), %r8
	mov	(%r9), %r9
	mov	(%r9), %r10
	mov	(%r9), %r11
	mov	(%r9), %r12
	mov	(%r9), %r13
	mov	(%r9), %r14
	mov	(%r9), %r15
	nop
	mov	(%r10),%rax
	mov	(%r10),%rcx
	mov	(%r10),%rdx
	mov	(%r10),%rbx
	mov	(%r10),%rsp
	mov	(%r10),%rbp
	mov	(%r10),%rsi
	mov	(%r10),%rdi
	mov	(%r10),%r8
	mov	(%r10),%r9
	mov	(%r10),%r10
	mov	(%r10),%r11
	mov	(%r10),%r12
	mov	(%r10),%r13
	mov	(%r10),%r14
	mov	(%r10),%r15
	nop
	mov	(%r11),%rax
	mov	(%r11),%rcx
	mov	(%r11),%rdx
	mov	(%r11),%rbx
	mov	(%r11),%rsp
	mov	(%r11),%rbp
	mov	(%r11),%rsi
	mov	(%r11),%rdi
	mov	(%r11),%r8
	mov	(%r11),%r9
	mov	(%r11),%r10
	mov	(%r11),%r11
	mov	(%r11),%r12
	mov	(%r11),%r13
	mov	(%r11),%r14
	mov	(%r11),%r15
	nop
	mov	(%r12),%rax
	mov	(%r12),%rcx
	mov	(%r12),%rdx
	mov	(%r12),%rbx
	mov	(%r12),%rsp
	mov	(%r12),%rbp
	mov	(%r12),%rsi
	mov	(%r12),%rdi
	mov	(%r12),%r8
	mov	(%r12),%r9
	mov	(%r12),%r10
	mov	(%r12),%r11
	mov	(%r12),%r12
	mov	(%r12),%r13
	mov	(%r12),%r14
	mov	(%r12),%r15
	nop
	mov	(%r13),%rax
	mov	(%r13),%rcx
	mov	(%r13),%rdx
	mov	(%r13),%rbx
	mov	(%r13),%rsp
	mov	(%r13),%rbp
	mov	(%r13),%rsi
	mov	(%r13),%rdi
	mov	(%r13),%r8
	mov	(%r13),%r9
	mov	(%r13),%r10
	mov	(%r13),%r11
	mov	(%r13),%r12
	mov	(%r13),%r13
	mov	(%r13),%r14
	mov	(%r13),%r15
	nop
	mov	(%r14),%rax
	mov	(%r14),%rcx
	mov	(%r14),%rdx
	mov	(%r14),%rbx
	mov	(%r14),%rsp
	mov	(%r14),%rbp
	mov	(%r14),%rsi
	mov	(%r14),%rdi
	mov	(%r14),%r8
	mov	(%r14),%r9
	mov	(%r14),%r10
	mov	(%r14),%r11
	mov	(%r14),%r12
	mov	(%r14),%r13
	mov	(%r14),%r14
	mov	(%r14),%r15
	nop
	mov	(%r15),%rax
	mov	(%r15),%rcx
	mov	(%r15),%rdx
	mov	(%r15),%rbx
	mov	(%r15),%rsp
	mov	(%r15),%rbp
	mov	(%r15),%rsi
	mov	(%r15),%rdi
	mov	(%r15),%r8
	mov	(%r15),%r9
	mov	(%r15),%r10
	mov	(%r15),%r11
	mov	(%r15),%r12
	mov	(%r15),%r13
	mov	(%r15),%r14
	mov	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	MovRegMem8
	.type	MovRegMem8, @function
MovRegMem8:
	.cfi_startproc
	mov	0x7f(%rax),%rax
	mov	0x7f(%rax),%rcx
	mov	0x7f(%rax),%rdx
	mov	0x7f(%rax),%rbx
	mov	0x7f(%rax),%rsp
	mov	0x7f(%rax),%rbp
	mov	0x7f(%rax),%rsi
	mov	0x7f(%rax),%rdi
	mov	0x7f(%rax),%r8
	mov	0x7f(%rax),%r9
	mov	0x7f(%rax),%r10
	mov	0x7f(%rax),%r11
	mov	0x7f(%rax),%r12
	mov	0x7f(%rax),%r13
	mov	0x7f(%rax),%r14
	mov	0x7f(%rax),%r15
	nop
	mov	0x7f(%rcx),%rax
	mov	0x7f(%rcx),%rcx
	mov	0x7f(%rcx),%rdx
	mov	0x7f(%rcx),%rbx
	mov	0x7f(%rcx),%rsp
	mov	0x7f(%rcx),%rbp
	mov	0x7f(%rcx),%rsi
	mov	0x7f(%rcx),%rdi
	mov	0x7f(%rcx),%r8
	mov	0x7f(%rcx),%r9
	mov	0x7f(%rcx),%r10
	mov	0x7f(%rcx),%r11
	mov	0x7f(%rcx),%r12
	mov	0x7f(%rcx),%r13
	mov	0x7f(%rcx),%r14
	mov	0x7f(%rcx),%r15
	nop
	mov	0x7f(%rdx),%rax
	mov	0x7f(%rdx),%rcx
	mov	0x7f(%rdx),%rdx
	mov	0x7f(%rdx),%rbx
	mov	0x7f(%rdx),%rsp
	mov	0x7f(%rdx),%rbp
	mov	0x7f(%rdx),%rsi
	mov	0x7f(%rdx),%rdi
	mov	0x7f(%rdx),%r8
	mov	0x7f(%rdx),%r9
	mov	0x7f(%rdx),%r10
	mov	0x7f(%rdx),%r11
	mov	0x7f(%rdx),%r12
	mov	0x7f(%rdx),%r13
	mov	0x7f(%rdx),%r14
	mov	0x7f(%rdx),%r15
	nop
	mov	0x7f(%rbx),%rax
	mov	0x7f(%rbx),%rcx
	mov	0x7f(%rbx),%rdx
	mov	0x7f(%rbx),%rbx
	mov	0x7f(%rbx),%rsp
	mov	0x7f(%rbx),%rbp
	mov	0x7f(%rbx),%rsi
	mov	0x7f(%rbx),%rdi
	mov	0x7f(%rbx),%r8
	mov	0x7f(%rbx),%r9
	mov	0x7f(%rbx),%r10
	mov	0x7f(%rbx),%r11
	mov	0x7f(%rbx),%r12
	mov	0x7f(%rbx),%r13
	mov	0x7f(%rbx),%r14
	mov	0x7f(%rbx),%r15
	nop
	mov	0x7f(%rsp),%rax
	mov	0x7f(%rsp),%rcx
	mov	0x7f(%rsp),%rdx
	mov	0x7f(%rsp),%rbx
	mov	0x7f(%rsp),%rsp
	mov	0x7f(%rsp),%rbp
	mov	0x7f(%rsp),%rsi
	mov	0x7f(%rsp),%rdi
	mov	0x7f(%rsp),%r8
	mov	0x7f(%rsp),%r9
	mov	0x7f(%rsp),%r10
	mov	0x7f(%rsp),%r11
	mov	0x7f(%rsp),%r12
	mov	0x7f(%rsp),%r13
	mov	0x7f(%rsp),%r14
	mov	0x7f(%rsp),%r15
	nop
	mov	0x7f(%rbp),%rax
	mov	0x7f(%rbp),%rcx
	mov	0x7f(%rbp),%rdx
	mov	0x7f(%rbp),%rbx
	mov	0x7f(%rbp),%rsp
	mov	0x7f(%rbp),%rbp
	mov	0x7f(%rbp),%rsi
	mov	0x7f(%rbp),%rdi
	mov	0x7f(%rbp),%r8
	mov	0x7f(%rbp),%r9
	mov	0x7f(%rbp),%r10
	mov	0x7f(%rbp),%r11
	mov	0x7f(%rbp),%r12
	mov	0x7f(%rbp),%r13
	mov	0x7f(%rbp),%r14
	mov	0x7f(%rbp),%r15
	nop
	mov	0x7f(%rsi),%rax
	mov	0x7f(%rsi),%rcx
	mov	0x7f(%rsi),%rdx
	mov	0x7f(%rsi),%rbx
	mov	0x7f(%rsi),%rsp
	mov	0x7f(%rsi),%rbp
	mov	0x7f(%rsi),%rsi
	mov	0x7f(%rsi),%rdi
	mov	0x7f(%rsi),%r8
	mov	0x7f(%rsi),%r9
	mov	0x7f(%rsi),%r10
	mov	0x7f(%rsi),%r11
	mov	0x7f(%rsi),%r12
	mov	0x7f(%rsi),%r13
	mov	0x7f(%rsi),%r14
	mov	0x7f(%rsi),%r15
	nop
	mov	0x7f(%rdi),%rax
	mov	0x7f(%rdi),%rcx
	mov	0x7f(%rdi),%rdx
	mov	0x7f(%rdi),%rbx
	mov	0x7f(%rdi),%rsp
	mov	0x7f(%rdi),%rbp
	mov	0x7f(%rdi),%rsi
	mov	0x7f(%rdi),%rdi
	mov	0x7f(%rdi),%r8
	mov	0x7f(%rdi),%r9
	mov	0x7f(%rdi),%r10
	mov	0x7f(%rdi),%r11
	mov	0x7f(%rdi),%r12
	mov	0x7f(%rdi),%r13
	mov	0x7f(%rdi),%r14
	mov	0x7f(%rdi),%r15
	nop
	mov	0x7f(%r8), %rax
	mov	0x7f(%r8), %rcx
	mov	0x7f(%r8), %rdx
	mov	0x7f(%r8), %rbx
	mov	0x7f(%r8), %rsp
	mov	0x7f(%r8), %rbp
	mov	0x7f(%r8), %rsi
	mov	0x7f(%r8), %rdi
	mov	0x7f(%r8), %r8
	mov	0x7f(%r8), %r9
	mov	0x7f(%r8), %r10
	mov	0x7f(%r8), %r11
	mov	0x7f(%r8), %r12
	mov	0x7f(%r8), %r13
	mov	0x7f(%r8), %r14
	mov	0x7f(%r8), %r15
	nop
	mov	0x7f(%r9), %rax
	mov	0x7f(%r9), %rcx
	mov	0x7f(%r9), %rdx
	mov	0x7f(%r9), %rbx
	mov	0x7f(%r9), %rsp
	mov	0x7f(%r9), %rbp
	mov	0x7f(%r9), %rsi
	mov	0x7f(%r9), %rdi
	mov	0x7f(%r9), %r8
	mov	0x7f(%r9), %r9
	mov	0x7f(%r9), %r10
	mov	0x7f(%r9), %r11
	mov	0x7f(%r9), %r12
	mov	0x7f(%r9), %r13
	mov	0x7f(%r9), %r14
	mov	0x7f(%r9), %r15
	nop
	mov	0x7f(%r10),%rax
	mov	0x7f(%r10),%rcx
	mov	0x7f(%r10),%rdx
	mov	0x7f(%r10),%rbx
	mov	0x7f(%r10),%rsp
	mov	0x7f(%r10),%rbp
	mov	0x7f(%r10),%rsi
	mov	0x7f(%r10),%rdi
	mov	0x7f(%r10),%r8
	mov	0x7f(%r10),%r9
	mov	0x7f(%r10),%r10
	mov	0x7f(%r10),%r11
	mov	0x7f(%r10),%r12
	mov	0x7f(%r10),%r13
	mov	0x7f(%r10),%r14
	mov	0x7f(%r10),%r15
	nop
	mov	0x7f(%r11),%rax
	mov	0x7f(%r11),%rcx
	mov	0x7f(%r11),%rdx
	mov	0x7f(%r11),%rbx
	mov	0x7f(%r11),%rsp
	mov	0x7f(%r11),%rbp
	mov	0x7f(%r11),%rsi
	mov	0x7f(%r11),%rdi
	mov	0x7f(%r11),%r8
	mov	0x7f(%r11),%r9
	mov	0x7f(%r11),%r10
	mov	0x7f(%r11),%r11
	mov	0x7f(%r11),%r12
	mov	0x7f(%r11),%r13
	mov	0x7f(%r11),%r14
	mov	0x7f(%r11),%r15
	nop
	mov	0x7f(%r12),%rax
	mov	0x7f(%r12),%rcx
	mov	0x7f(%r12),%rdx
	mov	0x7f(%r12),%rbx
	mov	0x7f(%r12),%rsp
	mov	0x7f(%r12),%rbp
	mov	0x7f(%r12),%rsi
	mov	0x7f(%r12),%rdi
	mov	0x7f(%r12),%r8
	mov	0x7f(%r12),%r9
	mov	0x7f(%r12),%r10
	mov	0x7f(%r12),%r11
	mov	0x7f(%r12),%r12
	mov	0x7f(%r12),%r13
	mov	0x7f(%r12),%r14
	mov	0x7f(%r12),%r15
	nop
	mov	0x7f(%r13),%rax
	mov	0x7f(%r13),%rcx
	mov	0x7f(%r13),%rdx
	mov	0x7f(%r13),%rbx
	mov	0x7f(%r13),%rsp
	mov	0x7f(%r13),%rbp
	mov	0x7f(%r13),%rsi
	mov	0x7f(%r13),%rdi
	mov	0x7f(%r13),%r8
	mov	0x7f(%r13),%r9
	mov	0x7f(%r13),%r10
	mov	0x7f(%r13),%r11
	mov	0x7f(%r13),%r12
	mov	0x7f(%r13),%r13
	mov	0x7f(%r13),%r14
	mov	0x7f(%r13),%r15
	nop
	mov	0x7f(%r14),%rax
	mov	0x7f(%r14),%rcx
	mov	0x7f(%r14),%rdx
	mov	0x7f(%r14),%rbx
	mov	0x7f(%r14),%rsp
	mov	0x7f(%r14),%rbp
	mov	0x7f(%r14),%rsi
	mov	0x7f(%r14),%rdi
	mov	0x7f(%r14),%r8
	mov	0x7f(%r14),%r9
	mov	0x7f(%r14),%r10
	mov	0x7f(%r14),%r11
	mov	0x7f(%r14),%r12
	mov	0x7f(%r14),%r13
	mov	0x7f(%r14),%r14
	mov	0x7f(%r14),%r15
	nop
	mov	0x7f(%r15),%rax
	mov	0x7f(%r15),%rcx
	mov	0x7f(%r15),%rdx
	mov	0x7f(%r15),%rbx
	mov	0x7f(%r15),%rsp
	mov	0x7f(%r15),%rbp
	mov	0x7f(%r15),%rsi
	mov	0x7f(%r15),%rdi
	mov	0x7f(%r15),%r8
	mov	0x7f(%r15),%r9
	mov	0x7f(%r15),%r10
	mov	0x7f(%r15),%r11
	mov	0x7f(%r15),%r12
	mov	0x7f(%r15),%r13
	mov	0x7f(%r15),%r14
	mov	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	MovRegMem32
	.type	MovRegMem32, @function
MovRegMem32:
	.cfi_startproc
	mov	0x7f563412(%rax),%rax
	mov	0x7f563412(%rax),%rcx
	mov	0x7f563412(%rax),%rdx
	mov	0x7f563412(%rax),%rbx
	mov	0x7f563412(%rax),%rsp
	mov	0x7f563412(%rax),%rbp
	mov	0x7f563412(%rax),%rsi
	mov	0x7f563412(%rax),%rdi
	mov	0x7f563412(%rax),%r8
	mov	0x7f563412(%rax),%r9
	mov	0x7f563412(%rax),%r10
	mov	0x7f563412(%rax),%r11
	mov	0x7f563412(%rax),%r12
	mov	0x7f563412(%rax),%r13
	mov	0x7f563412(%rax),%r14
	mov	0x7f563412(%rax),%r15
	nop
	mov	0x7f563412(%rcx),%rax
	mov	0x7f563412(%rcx),%rcx
	mov	0x7f563412(%rcx),%rdx
	mov	0x7f563412(%rcx),%rbx
	mov	0x7f563412(%rcx),%rsp
	mov	0x7f563412(%rcx),%rbp
	mov	0x7f563412(%rcx),%rsi
	mov	0x7f563412(%rcx),%rdi
	mov	0x7f563412(%rcx),%r8
	mov	0x7f563412(%rcx),%r9
	mov	0x7f563412(%rcx),%r10
	mov	0x7f563412(%rcx),%r11
	mov	0x7f563412(%rcx),%r12
	mov	0x7f563412(%rcx),%r13
	mov	0x7f563412(%rcx),%r14
	mov	0x7f563412(%rcx),%r15
	nop
	mov	0x7f563412(%rdx),%rax
	mov	0x7f563412(%rdx),%rcx
	mov	0x7f563412(%rdx),%rdx
	mov	0x7f563412(%rdx),%rbx
	mov	0x7f563412(%rdx),%rsp
	mov	0x7f563412(%rdx),%rbp
	mov	0x7f563412(%rdx),%rsi
	mov	0x7f563412(%rdx),%rdi
	mov	0x7f563412(%rdx),%r8
	mov	0x7f563412(%rdx),%r9
	mov	0x7f563412(%rdx),%r10
	mov	0x7f563412(%rdx),%r11
	mov	0x7f563412(%rdx),%r12
	mov	0x7f563412(%rdx),%r13
	mov	0x7f563412(%rdx),%r14
	mov	0x7f563412(%rdx),%r15
	nop
	mov	0x7f563412(%rbx),%rax
	mov	0x7f563412(%rbx),%rcx
	mov	0x7f563412(%rbx),%rdx
	mov	0x7f563412(%rbx),%rbx
	mov	0x7f563412(%rbx),%rsp
	mov	0x7f563412(%rbx),%rbp
	mov	0x7f563412(%rbx),%rsi
	mov	0x7f563412(%rbx),%rdi
	mov	0x7f563412(%rbx),%r8
	mov	0x7f563412(%rbx),%r9
	mov	0x7f563412(%rbx),%r10
	mov	0x7f563412(%rbx),%r11
	mov	0x7f563412(%rbx),%r12
	mov	0x7f563412(%rbx),%r13
	mov	0x7f563412(%rbx),%r14
	mov	0x7f563412(%rbx),%r15
	nop
	mov	0x7f563412(%rsp),%rax
	mov	0x7f563412(%rsp),%rcx
	mov	0x7f563412(%rsp),%rdx
	mov	0x7f563412(%rsp),%rbx
	mov	0x7f563412(%rsp),%rsp
	mov	0x7f563412(%rsp),%rbp
	mov	0x7f563412(%rsp),%rsi
	mov	0x7f563412(%rsp),%rdi
	mov	0x7f563412(%rsp),%r8
	mov	0x7f563412(%rsp),%r9
	mov	0x7f563412(%rsp),%r10
	mov	0x7f563412(%rsp),%r11
	mov	0x7f563412(%rsp),%r12
	mov	0x7f563412(%rsp),%r13
	mov	0x7f563412(%rsp),%r14
	mov	0x7f563412(%rsp),%r15
	nop
	mov	0x7f563412(%rbp),%rax
	mov	0x7f563412(%rbp),%rcx
	mov	0x7f563412(%rbp),%rdx
	mov	0x7f563412(%rbp),%rbx
	mov	0x7f563412(%rbp),%rsp
	mov	0x7f563412(%rbp),%rbp
	mov	0x7f563412(%rbp),%rsi
	mov	0x7f563412(%rbp),%rdi
	mov	0x7f563412(%rbp),%r8
	mov	0x7f563412(%rbp),%r9
	mov	0x7f563412(%rbp),%r10
	mov	0x7f563412(%rbp),%r11
	mov	0x7f563412(%rbp),%r12
	mov	0x7f563412(%rbp),%r13
	mov	0x7f563412(%rbp),%r14
	mov	0x7f563412(%rbp),%r15
	nop
	mov	0x7f563412(%rsi),%rax
	mov	0x7f563412(%rsi),%rcx
	mov	0x7f563412(%rsi),%rdx
	mov	0x7f563412(%rsi),%rbx
	mov	0x7f563412(%rsi),%rsp
	mov	0x7f563412(%rsi),%rbp
	mov	0x7f563412(%rsi),%rsi
	mov	0x7f563412(%rsi),%rdi
	mov	0x7f563412(%rsi),%r8
	mov	0x7f563412(%rsi),%r9
	mov	0x7f563412(%rsi),%r10
	mov	0x7f563412(%rsi),%r11
	mov	0x7f563412(%rsi),%r12
	mov	0x7f563412(%rsi),%r13
	mov	0x7f563412(%rsi),%r14
	mov	0x7f563412(%rsi),%r15
	nop
	mov	0x7f563412(%rdi),%rax
	mov	0x7f563412(%rdi),%rcx
	mov	0x7f563412(%rdi),%rdx
	mov	0x7f563412(%rdi),%rbx
	mov	0x7f563412(%rdi),%rsp
	mov	0x7f563412(%rdi),%rbp
	mov	0x7f563412(%rdi),%rsi
	mov	0x7f563412(%rdi),%rdi
	mov	0x7f563412(%rdi),%r8
	mov	0x7f563412(%rdi),%r9
	mov	0x7f563412(%rdi),%r10
	mov	0x7f563412(%rdi),%r11
	mov	0x7f563412(%rdi),%r12
	mov	0x7f563412(%rdi),%r13
	mov	0x7f563412(%rdi),%r14
	mov	0x7f563412(%rdi),%r15
	nop
	mov	0x7f563412(%r8), %rax
	mov	0x7f563412(%r8), %rcx
	mov	0x7f563412(%r8), %rdx
	mov	0x7f563412(%r8), %rbx
	mov	0x7f563412(%r8), %rsp
	mov	0x7f563412(%r8), %rbp
	mov	0x7f563412(%r8), %rsi
	mov	0x7f563412(%r8), %rdi
	mov	0x7f563412(%r8), %r8
	mov	0x7f563412(%r8), %r9
	mov	0x7f563412(%r8), %r10
	mov	0x7f563412(%r8), %r11
	mov	0x7f563412(%r8), %r12
	mov	0x7f563412(%r8), %r13
	mov	0x7f563412(%r8), %r14
	mov	0x7f563412(%r8), %r15
	nop
	mov	0x7f563412(%r9), %rax
	mov	0x7f563412(%r9), %rcx
	mov	0x7f563412(%r9), %rdx
	mov	0x7f563412(%r9), %rbx
	mov	0x7f563412(%r9), %rsp
	mov	0x7f563412(%r9), %rbp
	mov	0x7f563412(%r9), %rsi
	mov	0x7f563412(%r9), %rdi
	mov	0x7f563412(%r9), %r8
	mov	0x7f563412(%r9), %r9
	mov	0x7f563412(%r9), %r10
	mov	0x7f563412(%r9), %r11
	mov	0x7f563412(%r9), %r12
	mov	0x7f563412(%r9), %r13
	mov	0x7f563412(%r9), %r14
	mov	0x7f563412(%r9), %r15
	nop
	mov	0x7f563412(%r10),%rax
	mov	0x7f563412(%r10),%rcx
	mov	0x7f563412(%r10),%rdx
	mov	0x7f563412(%r10),%rbx
	mov	0x7f563412(%r10),%rsp
	mov	0x7f563412(%r10),%rbp
	mov	0x7f563412(%r10),%rsi
	mov	0x7f563412(%r10),%rdi
	mov	0x7f563412(%r10),%r8
	mov	0x7f563412(%r10),%r9
	mov	0x7f563412(%r10),%r10
	mov	0x7f563412(%r10),%r11
	mov	0x7f563412(%r10),%r12
	mov	0x7f563412(%r10),%r13
	mov	0x7f563412(%r10),%r14
	mov	0x7f563412(%r10),%r15
	nop
	mov	0x7f563412(%r11),%rax
	mov	0x7f563412(%r11),%rcx
	mov	0x7f563412(%r11),%rdx
	mov	0x7f563412(%r11),%rbx
	mov	0x7f563412(%r11),%rsp
	mov	0x7f563412(%r11),%rbp
	mov	0x7f563412(%r11),%rsi
	mov	0x7f563412(%r11),%rdi
	mov	0x7f563412(%r11),%r8
	mov	0x7f563412(%r11),%r9
	mov	0x7f563412(%r11),%r10
	mov	0x7f563412(%r11),%r11
	mov	0x7f563412(%r11),%r12
	mov	0x7f563412(%r11),%r13
	mov	0x7f563412(%r11),%r14
	mov	0x7f563412(%r11),%r15
	nop
	mov	0x7f563412(%r12),%rax
	mov	0x7f563412(%r12),%rcx
	mov	0x7f563412(%r12),%rdx
	mov	0x7f563412(%r12),%rbx
	mov	0x7f563412(%r12),%rsp
	mov	0x7f563412(%r12),%rbp
	mov	0x7f563412(%r12),%rsi
	mov	0x7f563412(%r12),%rdi
	mov	0x7f563412(%r12),%r8
	mov	0x7f563412(%r12),%r9
	mov	0x7f563412(%r12),%r10
	mov	0x7f563412(%r12),%r11
	mov	0x7f563412(%r12),%r12
	mov	0x7f563412(%r12),%r13
	mov	0x7f563412(%r12),%r14
	mov	0x7f563412(%r12),%r15
	nop
	mov	0x7f563412(%r13),%rax
	mov	0x7f563412(%r13),%rcx
	mov	0x7f563412(%r13),%rdx
	mov	0x7f563412(%r13),%rbx
	mov	0x7f563412(%r13),%rsp
	mov	0x7f563412(%r13),%rbp
	mov	0x7f563412(%r13),%rsi
	mov	0x7f563412(%r13),%rdi
	mov	0x7f563412(%r13),%r8
	mov	0x7f563412(%r13),%r9
	mov	0x7f563412(%r13),%r10
	mov	0x7f563412(%r13),%r11
	mov	0x7f563412(%r13),%r12
	mov	0x7f563412(%r13),%r13
	mov	0x7f563412(%r13),%r14
	mov	0x7f563412(%r13),%r15
	nop
	mov	0x7f563412(%r14),%rax
	mov	0x7f563412(%r14),%rcx
	mov	0x7f563412(%r14),%rdx
	mov	0x7f563412(%r14),%rbx
	mov	0x7f563412(%r14),%rsp
	mov	0x7f563412(%r14),%rbp
	mov	0x7f563412(%r14),%rsi
	mov	0x7f563412(%r14),%rdi
	mov	0x7f563412(%r14),%r8
	mov	0x7f563412(%r14),%r9
	mov	0x7f563412(%r14),%r10
	mov	0x7f563412(%r14),%r11
	mov	0x7f563412(%r14),%r12
	mov	0x7f563412(%r14),%r13
	mov	0x7f563412(%r14),%r14
	mov	0x7f563412(%r14),%r15
	nop
	mov	0x7f563412(%r15),%rax
	mov	0x7f563412(%r15),%rcx
	mov	0x7f563412(%r15),%rdx
	mov	0x7f563412(%r15),%rbx
	mov	0x7f563412(%r15),%rsp
	mov	0x7f563412(%r15),%rbp
	mov	0x7f563412(%r15),%rsi
	mov	0x7f563412(%r15),%rdi
	mov	0x7f563412(%r15),%r8
	mov	0x7f563412(%r15),%r9
	mov	0x7f563412(%r15),%r10
	mov	0x7f563412(%r15),%r11
	mov	0x7f563412(%r15),%r12
	mov	0x7f563412(%r15),%r13
	mov	0x7f563412(%r15),%r14
	mov	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


