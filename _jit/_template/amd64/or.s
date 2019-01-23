	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Or_s32
	.type	Or_s32, @function
Or_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// or	$-0x55667788,%rax
	or	$-0x55667788,%rcx
	or	$-0x55667788,%rdx
	or	$-0x55667788,%rbx
	or	$-0x55667788,%rsp
	or	$-0x55667788,%rbp
	or	$-0x55667788,%rsi
	or	$-0x55667788,%rdi
	or	$-0x55667788,%r8
	or	$-0x55667788,%r9
	or	$-0x55667788,%r10
	or	$-0x55667788,%r11
	or	$-0x55667788,%r12
	or	$-0x55667788,%r13
	or	$-0x55667788,%r14
	or	$-0x55667788,%r15
	.cfi_endproc


	.p2align 4,,15
	.globl	Or
	.type	Or, @function
Or:
	.cfi_startproc
	or	%rax,%rax
	or	%rax,%rcx
	or	%rax,%rdx
	or	%rax,%rbx
	or	%rax,%rsp
	or	%rax,%rbp
	or	%rax,%rsi
	or	%rax,%rdi
	or	%rax,%r8
	or	%rax,%r9
	or	%rax,%r10
	or	%rax,%r11
	or	%rax,%r12
	or	%rax,%r13
	or	%rax,%r14
	or	%rax,%r15
	nop
	or	%rcx,%rax
	or	%rcx,%rcx
	or	%rcx,%rdx
	or	%rcx,%rbx
	or	%rcx,%rsp
	or	%rcx,%rbp
	or	%rcx,%rsi
	or	%rcx,%rdi
	or	%rcx,%r8
	or	%rcx,%r9
	or	%rcx,%r10
	or	%rcx,%r11
	or	%rcx,%r12
	or	%rcx,%r13
	or	%rcx,%r14
	or	%rcx,%r15
	nop
	or	%rdx,%rax
	or	%rdx,%rcx
	or	%rdx,%rdx
	or	%rdx,%rbx
	or	%rdx,%rsp
	or	%rdx,%rbp
	or	%rdx,%rsi
	or	%rdx,%rdi
	or	%rdx,%r8
	or	%rdx,%r9
	or	%rdx,%r10
	or	%rdx,%r11
	or	%rdx,%r12
	or	%rdx,%r13
	or	%rdx,%r14
	or	%rdx,%r15
	nop
	or	%rbx,%rax
	or	%rbx,%rcx
	or	%rbx,%rdx
	or	%rbx,%rbx
	or	%rbx,%rsp
	or	%rbx,%rbp
	or	%rbx,%rsi
	or	%rbx,%rdi
	or	%rbx,%r8
	or	%rbx,%r9
	or	%rbx,%r10
	or	%rbx,%r11
	or	%rbx,%r12
	or	%rbx,%r13
	or	%rbx,%r14
	or	%rbx,%r15
	nop
	or	%rsp,%rax
	or	%rsp,%rcx
	or	%rsp,%rdx
	or	%rsp,%rbx
	or	%rsp,%rsp
	or	%rsp,%rbp
	or	%rsp,%rsi
	or	%rsp,%rdi
	or	%rsp,%r8
	or	%rsp,%r9
	or	%rsp,%r10
	or	%rsp,%r11
	or	%rsp,%r12
	or	%rsp,%r13
	or	%rsp,%r14
	or	%rsp,%r15
	nop
	or	%rbp,%rax
	or	%rbp,%rcx
	or	%rbp,%rdx
	or	%rbp,%rbx
	or	%rbp,%rsp
	or	%rbp,%rbp
	or	%rbp,%rsi
	or	%rbp,%rdi
	or	%rbp,%r8
	or	%rbp,%r9
	or	%rbp,%r10
	or	%rbp,%r11
	or	%rbp,%r12
	or	%rbp,%r13
	or	%rbp,%r14
	or	%rbp,%r15
	nop
	or	%rsi,%rax
	or	%rsi,%rcx
	or	%rsi,%rdx
	or	%rsi,%rbx
	or	%rsi,%rsp
	or	%rsi,%rbp
	or	%rsi,%rsi
	or	%rsi,%rdi
	or	%rsi,%r8
	or	%rsi,%r9
	or	%rsi,%r10
	or	%rsi,%r11
	or	%rsi,%r12
	or	%rsi,%r13
	or	%rsi,%r14
	or	%rsi,%r15
	nop
	or	%rdi,%rax
	or	%rdi,%rcx
	or	%rdi,%rdx
	or	%rdi,%rbx
	or	%rdi,%rsp
	or	%rdi,%rbp
	or	%rdi,%rsi
	or	%rdi,%rdi
	or	%rdi,%r8
	or	%rdi,%r9
	or	%rdi,%r10
	or	%rdi,%r11
	or	%rdi,%r12
	or	%rdi,%r13
	or	%rdi,%r14
	or	%rdi,%r15
	nop
	or	%r8, %rax
	or	%r8, %rcx
	or	%r8, %rdx
	or	%r8, %rbx
	or	%r8, %rsp
	or	%r8, %rbp
	or	%r8, %rsi
	or	%r8, %rdi
	or	%r8, %r8
	or	%r8, %r9
	or	%r8, %r10
	or	%r8, %r11
	or	%r8, %r12
	or	%r8, %r13
	or	%r8, %r14
	or	%r8, %r15
	nop
	or	%r9, %rax
	or	%r9, %rcx
	or	%r9, %rdx
	or	%r9, %rbx
	or	%r9, %rsp
	or	%r9, %rbp
	or	%r9, %rsi
	or	%r9, %rdi
	or	%r9, %r8
	or	%r9, %r9
	or	%r9, %r10
	or	%r9, %r11
	or	%r9, %r12
	or	%r9, %r13
	or	%r9, %r14
	or	%r9, %r15
	nop
	or	%r10,%rax
	or	%r10,%rcx
	or	%r10,%rdx
	or	%r10,%rbx
	or	%r10,%rsp
	or	%r10,%rbp
	or	%r10,%rsi
	or	%r10,%rdi
	or	%r10,%r8
	or	%r10,%r9
	or	%r10,%r10
	or	%r10,%r11
	or	%r10,%r12
	or	%r10,%r13
	or	%r10,%r14
	or	%r10,%r15
	nop
	or	%r11,%rax
	or	%r11,%rcx
	or	%r11,%rdx
	or	%r11,%rbx
	or	%r11,%rsp
	or	%r11,%rbp
	or	%r11,%rsi
	or	%r11,%rdi
	or	%r11,%r8
	or	%r11,%r9
	or	%r11,%r10
	or	%r11,%r11
	or	%r11,%r12
	or	%r11,%r13
	or	%r11,%r14
	or	%r11,%r15
	nop
	or	%r12,%rax
	or	%r12,%rcx
	or	%r12,%rdx
	or	%r12,%rbx
	or	%r12,%rsp
	or	%r12,%rbp
	or	%r12,%rsi
	or	%r12,%rdi
	or	%r12,%r8
	or	%r12,%r9
	or	%r12,%r10
	or	%r12,%r11
	or	%r12,%r12
	or	%r12,%r13
	or	%r12,%r14
	or	%r12,%r15
	nop
	or	%r13,%rax
	or	%r13,%rcx
	or	%r13,%rdx
	or	%r13,%rbx
	or	%r13,%rsp
	or	%r13,%rbp
	or	%r13,%rsi
	or	%r13,%rdi
	or	%r13,%r8
	or	%r13,%r9
	or	%r13,%r10
	or	%r13,%r11
	or	%r13,%r12
	or	%r13,%r13
	or	%r13,%r14
	or	%r13,%r15
	nop
	or	%r14,%rax
	or	%r14,%rcx
	or	%r14,%rdx
	or	%r14,%rbx
	or	%r14,%rsp
	or	%r14,%rbp
	or	%r14,%rsi
	or	%r14,%rdi
	or	%r14,%r8
	or	%r14,%r9
	or	%r14,%r10
	or	%r14,%r11
	or	%r14,%r12
	or	%r14,%r13
	or	%r14,%r14
	or	%r14,%r15
	nop
	or	%r15,%rax
	or	%r15,%rcx
	or	%r15,%rdx
	or	%r15,%rbx
	or	%r15,%rsp
	or	%r15,%rbp
	or	%r15,%rsi
	or	%r15,%rdi
	or	%r15,%r8
	or	%r15,%r9
	or	%r15,%r10
	or	%r15,%r11
	or	%r15,%r12
	or	%r15,%r13
	or	%r15,%r14
	or	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	OrMem
	.type	OrMem, @function
OrMem:
	.cfi_startproc
	or	%rax,(%rax)
	or	%rax,(%rcx)
	or	%rax,(%rdx)
	or	%rax,(%rbx)
	or	%rax,(%rsp)
	or	%rax,(%rbp)
	or	%rax,(%rsi)
	or	%rax,(%rdi)
	or	%rax,(%r8)
	or	%rax,(%r9)
	or	%rax,(%r10)
	or	%rax,(%r11)
	or	%rax,(%r12)
	or	%rax,(%r13)
	or	%rax,(%r14)
	or	%rax,(%r15)
	nop
	or	%rcx,(%rax)
	or	%rcx,(%rcx)
	or	%rcx,(%rdx)
	or	%rcx,(%rbx)
	or	%rcx,(%rsp)
	or	%rcx,(%rbp)
	or	%rcx,(%rsi)
	or	%rcx,(%rdi)
	or	%rcx,(%r8)
	or	%rcx,(%r9)
	or	%rcx,(%r10)
	or	%rcx,(%r11)
	or	%rcx,(%r12)
	or	%rcx,(%r13)
	or	%rcx,(%r14)
	or	%rcx,(%r15)
	nop
	or	%rdx,(%rax)
	or	%rdx,(%rcx)
	or	%rdx,(%rdx)
	or	%rdx,(%rbx)
	or	%rdx,(%rsp)
	or	%rdx,(%rbp)
	or	%rdx,(%rsi)
	or	%rdx,(%rdi)
	or	%rdx,(%r8)
	or	%rdx,(%r9)
	or	%rdx,(%r10)
	or	%rdx,(%r11)
	or	%rdx,(%r12)
	or	%rdx,(%r13)
	or	%rdx,(%r14)
	or	%rdx,(%r15)
	nop
	or	%rbx,(%rax)
	or	%rbx,(%rcx)
	or	%rbx,(%rdx)
	or	%rbx,(%rbx)
	or	%rbx,(%rsp)
	or	%rbx,(%rbp)
	or	%rbx,(%rsi)
	or	%rbx,(%rdi)
	or	%rbx,(%r8)
	or	%rbx,(%r9)
	or	%rbx,(%r10)
	or	%rbx,(%r11)
	or	%rbx,(%r12)
	or	%rbx,(%r13)
	or	%rbx,(%r14)
	or	%rbx,(%r15)
	nop
	or	%rsp,(%rax)
	or	%rsp,(%rcx)
	or	%rsp,(%rdx)
	or	%rsp,(%rbx)
	or	%rsp,(%rsp)
	or	%rsp,(%rbp)
	or	%rsp,(%rsi)
	or	%rsp,(%rdi)
	or	%rsp,(%r8)
	or	%rsp,(%r9)
	or	%rsp,(%r10)
	or	%rsp,(%r11)
	or	%rsp,(%r12)
	or	%rsp,(%r13)
	or	%rsp,(%r14)
	or	%rsp,(%r15)
	nop
	or	%rbp,(%rax)
	or	%rbp,(%rcx)
	or	%rbp,(%rdx)
	or	%rbp,(%rbx)
	or	%rbp,(%rsp)
	or	%rbp,(%rbp)
	or	%rbp,(%rsi)
	or	%rbp,(%rdi)
	or	%rbp,(%r8)
	or	%rbp,(%r9)
	or	%rbp,(%r10)
	or	%rbp,(%r11)
	or	%rbp,(%r12)
	or	%rbp,(%r13)
	or	%rbp,(%r14)
	or	%rbp,(%r15)
	nop
	or	%rsi,(%rax)
	or	%rsi,(%rcx)
	or	%rsi,(%rdx)
	or	%rsi,(%rbx)
	or	%rsi,(%rsp)
	or	%rsi,(%rbp)
	or	%rsi,(%rsi)
	or	%rsi,(%rdi)
	or	%rsi,(%r8)
	or	%rsi,(%r9)
	or	%rsi,(%r10)
	or	%rsi,(%r11)
	or	%rsi,(%r12)
	or	%rsi,(%r13)
	or	%rsi,(%r14)
	or	%rsi,(%r15)
	nop
	or	%rdi,(%rax)
	or	%rdi,(%rcx)
	or	%rdi,(%rdx)
	or	%rdi,(%rbx)
	or	%rdi,(%rsp)
	or	%rdi,(%rbp)
	or	%rdi,(%rsi)
	or	%rdi,(%rdi)
	or	%rdi,(%r8)
	or	%rdi,(%r9)
	or	%rdi,(%r10)
	or	%rdi,(%r11)
	or	%rdi,(%r12)
	or	%rdi,(%r13)
	or	%rdi,(%r14)
	or	%rdi,(%r15)
	nop
	or	%r8, (%rax)
	or	%r8, (%rcx)
	or	%r8, (%rdx)
	or	%r8, (%rbx)
	or	%r8, (%rsp)
	or	%r8, (%rbp)
	or	%r8, (%rsi)
	or	%r8, (%rdi)
	or	%r8, (%r8)
	or	%r8, (%r9)
	or	%r8, (%r10)
	or	%r8, (%r11)
	or	%r8, (%r12)
	or	%r8, (%r13)
	or	%r8, (%r14)
	or	%r8, (%r15)
	nop
	or	%r9, (%rax)
	or	%r9, (%rcx)
	or	%r9, (%rdx)
	or	%r9, (%rbx)
	or	%r9, (%rsp)
	or	%r9, (%rbp)
	or	%r9, (%rsi)
	or	%r9, (%rdi)
	or	%r9, (%r8)
	or	%r9, (%r9)
	or	%r9, (%r10)
	or	%r9, (%r11)
	or	%r9, (%r12)
	or	%r9, (%r13)
	or	%r9, (%r14)
	or	%r9, (%r15)
	nop
	or	%r10,(%rax)
	or	%r10,(%rcx)
	or	%r10,(%rdx)
	or	%r10,(%rbx)
	or	%r10,(%rsp)
	or	%r10,(%rbp)
	or	%r10,(%rsi)
	or	%r10,(%rdi)
	or	%r10,(%r8)
	or	%r10,(%r9)
	or	%r10,(%r10)
	or	%r10,(%r11)
	or	%r10,(%r12)
	or	%r10,(%r13)
	or	%r10,(%r14)
	or	%r10,(%r15)
	nop
	or	%r11,(%rax)
	or	%r11,(%rcx)
	or	%r11,(%rdx)
	or	%r11,(%rbx)
	or	%r11,(%rsp)
	or	%r11,(%rbp)
	or	%r11,(%rsi)
	or	%r11,(%rdi)
	or	%r11,(%r8)
	or	%r11,(%r9)
	or	%r11,(%r10)
	or	%r11,(%r11)
	or	%r11,(%r12)
	or	%r11,(%r13)
	or	%r11,(%r14)
	or	%r11,(%r15)
	nop
	or	%r12,(%rax)
	or	%r12,(%rcx)
	or	%r12,(%rdx)
	or	%r12,(%rbx)
	or	%r12,(%rsp)
	or	%r12,(%rbp)
	or	%r12,(%rsi)
	or	%r12,(%rdi)
	or	%r12,(%r8)
	or	%r12,(%r9)
	or	%r12,(%r10)
	or	%r12,(%r11)
	or	%r12,(%r12)
	or	%r12,(%r13)
	or	%r12,(%r14)
	or	%r12,(%r15)
	nop
	or	%r13,(%rax)
	or	%r13,(%rcx)
	or	%r13,(%rdx)
	or	%r13,(%rbx)
	or	%r13,(%rsp)
	or	%r13,(%rbp)
	or	%r13,(%rsi)
	or	%r13,(%rdi)
	or	%r13,(%r8)
	or	%r13,(%r9)
	or	%r13,(%r10)
	or	%r13,(%r11)
	or	%r13,(%r12)
	or	%r13,(%r13)
	or	%r13,(%r14)
	or	%r13,(%r15)
	nop
	or	%r14,(%rax)
	or	%r14,(%rcx)
	or	%r14,(%rdx)
	or	%r14,(%rbx)
	or	%r14,(%rsp)
	or	%r14,(%rbp)
	or	%r14,(%rsi)
	or	%r14,(%rdi)
	or	%r14,(%r8)
	or	%r14,(%r9)
	or	%r14,(%r10)
	or	%r14,(%r11)
	or	%r14,(%r12)
	or	%r14,(%r13)
	or	%r14,(%r14)
	or	%r14,(%r15)
	nop
	or	%r15,(%rax)
	or	%r15,(%rcx)
	or	%r15,(%rdx)
	or	%r15,(%rbx)
	or	%r15,(%rsp)
	or	%r15,(%rbp)
	or	%r15,(%rsi)
	or	%r15,(%rdi)
	or	%r15,(%r8)
	or	%r15,(%r9)
	or	%r15,(%r10)
	or	%r15,(%r11)
	or	%r15,(%r12)
	or	%r15,(%r13)
	or	%r15,(%r14)
	or	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	OrMem8
	.type	OrMem8, @function
OrMem8:
	.cfi_startproc
	or	%rax,0x7f(%rax)
	or	%rax,0x7f(%rcx)
	or	%rax,0x7f(%rdx)
	or	%rax,0x7f(%rbx)
	or	%rax,0x7f(%rsp)
	or	%rax,0x7f(%rbp)
	or	%rax,0x7f(%rsi)
	or	%rax,0x7f(%rdi)
	or	%rax,0x7f(%r8)
	or	%rax,0x7f(%r9)
	or	%rax,0x7f(%r10)
	or	%rax,0x7f(%r11)
	or	%rax,0x7f(%r12)
	or	%rax,0x7f(%r13)
	or	%rax,0x7f(%r14)
	or	%rax,0x7f(%r15)
	nop
	or	%rcx,0x7f(%rax)
	or	%rcx,0x7f(%rcx)
	or	%rcx,0x7f(%rdx)
	or	%rcx,0x7f(%rbx)
	or	%rcx,0x7f(%rsp)
	or	%rcx,0x7f(%rbp)
	or	%rcx,0x7f(%rsi)
	or	%rcx,0x7f(%rdi)
	or	%rcx,0x7f(%r8)
	or	%rcx,0x7f(%r9)
	or	%rcx,0x7f(%r10)
	or	%rcx,0x7f(%r11)
	or	%rcx,0x7f(%r12)
	or	%rcx,0x7f(%r13)
	or	%rcx,0x7f(%r14)
	or	%rcx,0x7f(%r15)
	nop
	or	%rdx,0x7f(%rax)
	or	%rdx,0x7f(%rcx)
	or	%rdx,0x7f(%rdx)
	or	%rdx,0x7f(%rbx)
	or	%rdx,0x7f(%rsp)
	or	%rdx,0x7f(%rbp)
	or	%rdx,0x7f(%rsi)
	or	%rdx,0x7f(%rdi)
	or	%rdx,0x7f(%r8)
	or	%rdx,0x7f(%r9)
	or	%rdx,0x7f(%r10)
	or	%rdx,0x7f(%r11)
	or	%rdx,0x7f(%r12)
	or	%rdx,0x7f(%r13)
	or	%rdx,0x7f(%r14)
	or	%rdx,0x7f(%r15)
	nop
	or	%rbx,0x7f(%rax)
	or	%rbx,0x7f(%rcx)
	or	%rbx,0x7f(%rdx)
	or	%rbx,0x7f(%rbx)
	or	%rbx,0x7f(%rsp)
	or	%rbx,0x7f(%rbp)
	or	%rbx,0x7f(%rsi)
	or	%rbx,0x7f(%rdi)
	or	%rbx,0x7f(%r8)
	or	%rbx,0x7f(%r9)
	or	%rbx,0x7f(%r10)
	or	%rbx,0x7f(%r11)
	or	%rbx,0x7f(%r12)
	or	%rbx,0x7f(%r13)
	or	%rbx,0x7f(%r14)
	or	%rbx,0x7f(%r15)
	nop
	or	%rsp,0x7f(%rax)
	or	%rsp,0x7f(%rcx)
	or	%rsp,0x7f(%rdx)
	or	%rsp,0x7f(%rbx)
	or	%rsp,0x7f(%rsp)
	or	%rsp,0x7f(%rbp)
	or	%rsp,0x7f(%rsi)
	or	%rsp,0x7f(%rdi)
	or	%rsp,0x7f(%r8)
	or	%rsp,0x7f(%r9)
	or	%rsp,0x7f(%r10)
	or	%rsp,0x7f(%r11)
	or	%rsp,0x7f(%r12)
	or	%rsp,0x7f(%r13)
	or	%rsp,0x7f(%r14)
	or	%rsp,0x7f(%r15)
	nop
	or	%rbp,0x7f(%rax)
	or	%rbp,0x7f(%rcx)
	or	%rbp,0x7f(%rdx)
	or	%rbp,0x7f(%rbx)
	or	%rbp,0x7f(%rsp)
	or	%rbp,0x7f(%rbp)
	or	%rbp,0x7f(%rsi)
	or	%rbp,0x7f(%rdi)
	or	%rbp,0x7f(%r8)
	or	%rbp,0x7f(%r9)
	or	%rbp,0x7f(%r10)
	or	%rbp,0x7f(%r11)
	or	%rbp,0x7f(%r12)
	or	%rbp,0x7f(%r13)
	or	%rbp,0x7f(%r14)
	or	%rbp,0x7f(%r15)
	nop
	or	%rsi,0x7f(%rax)
	or	%rsi,0x7f(%rcx)
	or	%rsi,0x7f(%rdx)
	or	%rsi,0x7f(%rbx)
	or	%rsi,0x7f(%rsp)
	or	%rsi,0x7f(%rbp)
	or	%rsi,0x7f(%rsi)
	or	%rsi,0x7f(%rdi)
	or	%rsi,0x7f(%r8)
	or	%rsi,0x7f(%r9)
	or	%rsi,0x7f(%r10)
	or	%rsi,0x7f(%r11)
	or	%rsi,0x7f(%r12)
	or	%rsi,0x7f(%r13)
	or	%rsi,0x7f(%r14)
	or	%rsi,0x7f(%r15)
	nop
	or	%rdi,0x7f(%rax)
	or	%rdi,0x7f(%rcx)
	or	%rdi,0x7f(%rdx)
	or	%rdi,0x7f(%rbx)
	or	%rdi,0x7f(%rsp)
	or	%rdi,0x7f(%rbp)
	or	%rdi,0x7f(%rsi)
	or	%rdi,0x7f(%rdi)
	or	%rdi,0x7f(%r8)
	or	%rdi,0x7f(%r9)
	or	%rdi,0x7f(%r10)
	or	%rdi,0x7f(%r11)
	or	%rdi,0x7f(%r12)
	or	%rdi,0x7f(%r13)
	or	%rdi,0x7f(%r14)
	or	%rdi,0x7f(%r15)
	nop
	or	%r8, 0x7f(%rax)
	or	%r8, 0x7f(%rcx)
	or	%r8, 0x7f(%rdx)
	or	%r8, 0x7f(%rbx)
	or	%r8, 0x7f(%rsp)
	or	%r8, 0x7f(%rbp)
	or	%r8, 0x7f(%rsi)
	or	%r8, 0x7f(%rdi)
	or	%r8, 0x7f(%r8)
	or	%r8, 0x7f(%r9)
	or	%r8, 0x7f(%r10)
	or	%r8, 0x7f(%r11)
	or	%r8, 0x7f(%r12)
	or	%r8, 0x7f(%r13)
	or	%r8, 0x7f(%r14)
	or	%r8, 0x7f(%r15)
	nop
	or	%r9, 0x7f(%rax)
	or	%r9, 0x7f(%rcx)
	or	%r9, 0x7f(%rdx)
	or	%r9, 0x7f(%rbx)
	or	%r9, 0x7f(%rsp)
	or	%r9, 0x7f(%rbp)
	or	%r9, 0x7f(%rsi)
	or	%r9, 0x7f(%rdi)
	or	%r9, 0x7f(%r8)
	or	%r9, 0x7f(%r9)
	or	%r9, 0x7f(%r10)
	or	%r9, 0x7f(%r11)
	or	%r9, 0x7f(%r12)
	or	%r9, 0x7f(%r13)
	or	%r9, 0x7f(%r14)
	or	%r9, 0x7f(%r15)
	nop
	or	%r10,0x7f(%rax)
	or	%r10,0x7f(%rcx)
	or	%r10,0x7f(%rdx)
	or	%r10,0x7f(%rbx)
	or	%r10,0x7f(%rsp)
	or	%r10,0x7f(%rbp)
	or	%r10,0x7f(%rsi)
	or	%r10,0x7f(%rdi)
	or	%r10,0x7f(%r8)
	or	%r10,0x7f(%r9)
	or	%r10,0x7f(%r10)
	or	%r10,0x7f(%r11)
	or	%r10,0x7f(%r12)
	or	%r10,0x7f(%r13)
	or	%r10,0x7f(%r14)
	or	%r10,0x7f(%r15)
	nop
	or	%r11,0x7f(%rax)
	or	%r11,0x7f(%rcx)
	or	%r11,0x7f(%rdx)
	or	%r11,0x7f(%rbx)
	or	%r11,0x7f(%rsp)
	or	%r11,0x7f(%rbp)
	or	%r11,0x7f(%rsi)
	or	%r11,0x7f(%rdi)
	or	%r11,0x7f(%r8)
	or	%r11,0x7f(%r9)
	or	%r11,0x7f(%r10)
	or	%r11,0x7f(%r11)
	or	%r11,0x7f(%r12)
	or	%r11,0x7f(%r13)
	or	%r11,0x7f(%r14)
	or	%r11,0x7f(%r15)
	nop
	or	%r12,0x7f(%rax)
	or	%r12,0x7f(%rcx)
	or	%r12,0x7f(%rdx)
	or	%r12,0x7f(%rbx)
	or	%r12,0x7f(%rsp)
	or	%r12,0x7f(%rbp)
	or	%r12,0x7f(%rsi)
	or	%r12,0x7f(%rdi)
	or	%r12,0x7f(%r8)
	or	%r12,0x7f(%r9)
	or	%r12,0x7f(%r10)
	or	%r12,0x7f(%r11)
	or	%r12,0x7f(%r12)
	or	%r12,0x7f(%r13)
	or	%r12,0x7f(%r14)
	or	%r12,0x7f(%r15)
	nop
	or	%r13,0x7f(%rax)
	or	%r13,0x7f(%rcx)
	or	%r13,0x7f(%rdx)
	or	%r13,0x7f(%rbx)
	or	%r13,0x7f(%rsp)
	or	%r13,0x7f(%rbp)
	or	%r13,0x7f(%rsi)
	or	%r13,0x7f(%rdi)
	or	%r13,0x7f(%r8)
	or	%r13,0x7f(%r9)
	or	%r13,0x7f(%r10)
	or	%r13,0x7f(%r11)
	or	%r13,0x7f(%r12)
	or	%r13,0x7f(%r13)
	or	%r13,0x7f(%r14)
	or	%r13,0x7f(%r15)
	nop
	or	%r14,0x7f(%rax)
	or	%r14,0x7f(%rcx)
	or	%r14,0x7f(%rdx)
	or	%r14,0x7f(%rbx)
	or	%r14,0x7f(%rsp)
	or	%r14,0x7f(%rbp)
	or	%r14,0x7f(%rsi)
	or	%r14,0x7f(%rdi)
	or	%r14,0x7f(%r8)
	or	%r14,0x7f(%r9)
	or	%r14,0x7f(%r10)
	or	%r14,0x7f(%r11)
	or	%r14,0x7f(%r12)
	or	%r14,0x7f(%r13)
	or	%r14,0x7f(%r14)
	or	%r14,0x7f(%r15)
	nop
	or	%r15,0x7f(%rax)
	or	%r15,0x7f(%rcx)
	or	%r15,0x7f(%rdx)
	or	%r15,0x7f(%rbx)
	or	%r15,0x7f(%rsp)
	or	%r15,0x7f(%rbp)
	or	%r15,0x7f(%rsi)
	or	%r15,0x7f(%rdi)
	or	%r15,0x7f(%r8)
	or	%r15,0x7f(%r9)
	or	%r15,0x7f(%r10)
	or	%r15,0x7f(%r11)
	or	%r15,0x7f(%r12)
	or	%r15,0x7f(%r13)
	or	%r15,0x7f(%r14)
	or	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	OrMem32
	.type	OrMem32, @function
OrMem32:
	.cfi_startproc
	or	%rax,0x7f563412(%rax)
	or	%rax,0x7f563412(%rcx)
	or	%rax,0x7f563412(%rdx)
	or	%rax,0x7f563412(%rbx)
	or	%rax,0x7f563412(%rsp)
	or	%rax,0x7f563412(%rbp)
	or	%rax,0x7f563412(%rsi)
	or	%rax,0x7f563412(%rdi)
	or	%rax,0x7f563412(%r8)
	or	%rax,0x7f563412(%r9)
	or	%rax,0x7f563412(%r10)
	or	%rax,0x7f563412(%r11)
	or	%rax,0x7f563412(%r12)
	or	%rax,0x7f563412(%r13)
	or	%rax,0x7f563412(%r14)
	or	%rax,0x7f563412(%r15)
	nop
	or	%rcx,0x7f563412(%rax)
	or	%rcx,0x7f563412(%rcx)
	or	%rcx,0x7f563412(%rdx)
	or	%rcx,0x7f563412(%rbx)
	or	%rcx,0x7f563412(%rsp)
	or	%rcx,0x7f563412(%rbp)
	or	%rcx,0x7f563412(%rsi)
	or	%rcx,0x7f563412(%rdi)
	or	%rcx,0x7f563412(%r8)
	or	%rcx,0x7f563412(%r9)
	or	%rcx,0x7f563412(%r10)
	or	%rcx,0x7f563412(%r11)
	or	%rcx,0x7f563412(%r12)
	or	%rcx,0x7f563412(%r13)
	or	%rcx,0x7f563412(%r14)
	or	%rcx,0x7f563412(%r15)
	nop
	or	%rdx,0x7f563412(%rax)
	or	%rdx,0x7f563412(%rcx)
	or	%rdx,0x7f563412(%rdx)
	or	%rdx,0x7f563412(%rbx)
	or	%rdx,0x7f563412(%rsp)
	or	%rdx,0x7f563412(%rbp)
	or	%rdx,0x7f563412(%rsi)
	or	%rdx,0x7f563412(%rdi)
	or	%rdx,0x7f563412(%r8)
	or	%rdx,0x7f563412(%r9)
	or	%rdx,0x7f563412(%r10)
	or	%rdx,0x7f563412(%r11)
	or	%rdx,0x7f563412(%r12)
	or	%rdx,0x7f563412(%r13)
	or	%rdx,0x7f563412(%r14)
	or	%rdx,0x7f563412(%r15)
	nop
	or	%rbx,0x7f563412(%rax)
	or	%rbx,0x7f563412(%rcx)
	or	%rbx,0x7f563412(%rdx)
	or	%rbx,0x7f563412(%rbx)
	or	%rbx,0x7f563412(%rsp)
	or	%rbx,0x7f563412(%rbp)
	or	%rbx,0x7f563412(%rsi)
	or	%rbx,0x7f563412(%rdi)
	or	%rbx,0x7f563412(%r8)
	or	%rbx,0x7f563412(%r9)
	or	%rbx,0x7f563412(%r10)
	or	%rbx,0x7f563412(%r11)
	or	%rbx,0x7f563412(%r12)
	or	%rbx,0x7f563412(%r13)
	or	%rbx,0x7f563412(%r14)
	or	%rbx,0x7f563412(%r15)
	nop
	or	%rsp,0x7f563412(%rax)
	or	%rsp,0x7f563412(%rcx)
	or	%rsp,0x7f563412(%rdx)
	or	%rsp,0x7f563412(%rbx)
	or	%rsp,0x7f563412(%rsp)
	or	%rsp,0x7f563412(%rbp)
	or	%rsp,0x7f563412(%rsi)
	or	%rsp,0x7f563412(%rdi)
	or	%rsp,0x7f563412(%r8)
	or	%rsp,0x7f563412(%r9)
	or	%rsp,0x7f563412(%r10)
	or	%rsp,0x7f563412(%r11)
	or	%rsp,0x7f563412(%r12)
	or	%rsp,0x7f563412(%r13)
	or	%rsp,0x7f563412(%r14)
	or	%rsp,0x7f563412(%r15)
	nop
	or	%rbp,0x7f563412(%rax)
	or	%rbp,0x7f563412(%rcx)
	or	%rbp,0x7f563412(%rdx)
	or	%rbp,0x7f563412(%rbx)
	or	%rbp,0x7f563412(%rsp)
	or	%rbp,0x7f563412(%rbp)
	or	%rbp,0x7f563412(%rsi)
	or	%rbp,0x7f563412(%rdi)
	or	%rbp,0x7f563412(%r8)
	or	%rbp,0x7f563412(%r9)
	or	%rbp,0x7f563412(%r10)
	or	%rbp,0x7f563412(%r11)
	or	%rbp,0x7f563412(%r12)
	or	%rbp,0x7f563412(%r13)
	or	%rbp,0x7f563412(%r14)
	or	%rbp,0x7f563412(%r15)
	nop
	or	%rsi,0x7f563412(%rax)
	or	%rsi,0x7f563412(%rcx)
	or	%rsi,0x7f563412(%rdx)
	or	%rsi,0x7f563412(%rbx)
	or	%rsi,0x7f563412(%rsp)
	or	%rsi,0x7f563412(%rbp)
	or	%rsi,0x7f563412(%rsi)
	or	%rsi,0x7f563412(%rdi)
	or	%rsi,0x7f563412(%r8)
	or	%rsi,0x7f563412(%r9)
	or	%rsi,0x7f563412(%r10)
	or	%rsi,0x7f563412(%r11)
	or	%rsi,0x7f563412(%r12)
	or	%rsi,0x7f563412(%r13)
	or	%rsi,0x7f563412(%r14)
	or	%rsi,0x7f563412(%r15)
	nop
	or	%rdi,0x7f563412(%rax)
	or	%rdi,0x7f563412(%rcx)
	or	%rdi,0x7f563412(%rdx)
	or	%rdi,0x7f563412(%rbx)
	or	%rdi,0x7f563412(%rsp)
	or	%rdi,0x7f563412(%rbp)
	or	%rdi,0x7f563412(%rsi)
	or	%rdi,0x7f563412(%rdi)
	or	%rdi,0x7f563412(%r8)
	or	%rdi,0x7f563412(%r9)
	or	%rdi,0x7f563412(%r10)
	or	%rdi,0x7f563412(%r11)
	or	%rdi,0x7f563412(%r12)
	or	%rdi,0x7f563412(%r13)
	or	%rdi,0x7f563412(%r14)
	or	%rdi,0x7f563412(%r15)
	nop
	or	%r8, 0x7f563412(%rax)
	or	%r8, 0x7f563412(%rcx)
	or	%r8, 0x7f563412(%rdx)
	or	%r8, 0x7f563412(%rbx)
	or	%r8, 0x7f563412(%rsp)
	or	%r8, 0x7f563412(%rbp)
	or	%r8, 0x7f563412(%rsi)
	or	%r8, 0x7f563412(%rdi)
	or	%r8, 0x7f563412(%r8)
	or	%r8, 0x7f563412(%r9)
	or	%r8, 0x7f563412(%r10)
	or	%r8, 0x7f563412(%r11)
	or	%r8, 0x7f563412(%r12)
	or	%r8, 0x7f563412(%r13)
	or	%r8, 0x7f563412(%r14)
	or	%r8, 0x7f563412(%r15)
	nop
	or	%r9, 0x7f563412(%rax)
	or	%r9, 0x7f563412(%rcx)
	or	%r9, 0x7f563412(%rdx)
	or	%r9, 0x7f563412(%rbx)
	or	%r9, 0x7f563412(%rsp)
	or	%r9, 0x7f563412(%rbp)
	or	%r9, 0x7f563412(%rsi)
	or	%r9, 0x7f563412(%rdi)
	or	%r9, 0x7f563412(%r8)
	or	%r9, 0x7f563412(%r9)
	or	%r9, 0x7f563412(%r10)
	or	%r9, 0x7f563412(%r11)
	or	%r9, 0x7f563412(%r12)
	or	%r9, 0x7f563412(%r13)
	or	%r9, 0x7f563412(%r14)
	or	%r9, 0x7f563412(%r15)
	nop
	or	%r10,0x7f563412(%rax)
	or	%r10,0x7f563412(%rcx)
	or	%r10,0x7f563412(%rdx)
	or	%r10,0x7f563412(%rbx)
	or	%r10,0x7f563412(%rsp)
	or	%r10,0x7f563412(%rbp)
	or	%r10,0x7f563412(%rsi)
	or	%r10,0x7f563412(%rdi)
	or	%r10,0x7f563412(%r8)
	or	%r10,0x7f563412(%r9)
	or	%r10,0x7f563412(%r10)
	or	%r10,0x7f563412(%r11)
	or	%r10,0x7f563412(%r12)
	or	%r10,0x7f563412(%r13)
	or	%r10,0x7f563412(%r14)
	or	%r10,0x7f563412(%r15)
	nop
	or	%r11,0x7f563412(%rax)
	or	%r11,0x7f563412(%rcx)
	or	%r11,0x7f563412(%rdx)
	or	%r11,0x7f563412(%rbx)
	or	%r11,0x7f563412(%rsp)
	or	%r11,0x7f563412(%rbp)
	or	%r11,0x7f563412(%rsi)
	or	%r11,0x7f563412(%rdi)
	or	%r11,0x7f563412(%r8)
	or	%r11,0x7f563412(%r9)
	or	%r11,0x7f563412(%r10)
	or	%r11,0x7f563412(%r11)
	or	%r11,0x7f563412(%r12)
	or	%r11,0x7f563412(%r13)
	or	%r11,0x7f563412(%r14)
	or	%r11,0x7f563412(%r15)
	nop
	or	%r12,0x7f563412(%rax)
	or	%r12,0x7f563412(%rcx)
	or	%r12,0x7f563412(%rdx)
	or	%r12,0x7f563412(%rbx)
	or	%r12,0x7f563412(%rsp)
	or	%r12,0x7f563412(%rbp)
	or	%r12,0x7f563412(%rsi)
	or	%r12,0x7f563412(%rdi)
	or	%r12,0x7f563412(%r8)
	or	%r12,0x7f563412(%r9)
	or	%r12,0x7f563412(%r10)
	or	%r12,0x7f563412(%r11)
	or	%r12,0x7f563412(%r12)
	or	%r12,0x7f563412(%r13)
	or	%r12,0x7f563412(%r14)
	or	%r12,0x7f563412(%r15)
	nop
	or	%r13,0x7f563412(%rax)
	or	%r13,0x7f563412(%rcx)
	or	%r13,0x7f563412(%rdx)
	or	%r13,0x7f563412(%rbx)
	or	%r13,0x7f563412(%rsp)
	or	%r13,0x7f563412(%rbp)
	or	%r13,0x7f563412(%rsi)
	or	%r13,0x7f563412(%rdi)
	or	%r13,0x7f563412(%r8)
	or	%r13,0x7f563412(%r9)
	or	%r13,0x7f563412(%r10)
	or	%r13,0x7f563412(%r11)
	or	%r13,0x7f563412(%r12)
	or	%r13,0x7f563412(%r13)
	or	%r13,0x7f563412(%r14)
	or	%r13,0x7f563412(%r15)
	nop
	or	%r14,0x7f563412(%rax)
	or	%r14,0x7f563412(%rcx)
	or	%r14,0x7f563412(%rdx)
	or	%r14,0x7f563412(%rbx)
	or	%r14,0x7f563412(%rsp)
	or	%r14,0x7f563412(%rbp)
	or	%r14,0x7f563412(%rsi)
	or	%r14,0x7f563412(%rdi)
	or	%r14,0x7f563412(%r8)
	or	%r14,0x7f563412(%r9)
	or	%r14,0x7f563412(%r10)
	or	%r14,0x7f563412(%r11)
	or	%r14,0x7f563412(%r12)
	or	%r14,0x7f563412(%r13)
	or	%r14,0x7f563412(%r14)
	or	%r14,0x7f563412(%r15)
	nop
	or	%r15,0x7f563412(%rax)
	or	%r15,0x7f563412(%rcx)
	or	%r15,0x7f563412(%rdx)
	or	%r15,0x7f563412(%rbx)
	or	%r15,0x7f563412(%rsp)
	or	%r15,0x7f563412(%rbp)
	or	%r15,0x7f563412(%rsi)
	or	%r15,0x7f563412(%rdi)
	or	%r15,0x7f563412(%r8)
	or	%r15,0x7f563412(%r9)
	or	%r15,0x7f563412(%r10)
	or	%r15,0x7f563412(%r11)
	or	%r15,0x7f563412(%r12)
	or	%r15,0x7f563412(%r13)
	or	%r15,0x7f563412(%r14)
	or	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
