	.file	"arith.s"
	.text

        .p2align 4,,15
	.globl	Xchg
	.type	Xchg, @function
Xchg:
	.cfi_startproc
        
	xchg	%rax,%rax
	xchg	%rax,%rcx
	xchg	%rax,%rdx
	xchg	%rax,%rbx
	xchg	%rax,%rsp
	xchg	%rax,%rbp
	xchg	%rax,%rsi
	xchg	%rax,%rdi
	xchg	%rax,%r8
	xchg	%rax,%r9
	xchg	%rax,%r10
	xchg	%rax,%r11
	xchg	%rax,%r12
	xchg	%rax,%r13
	xchg	%rax,%r14
	xchg	%rax,%r15
	nop
	xchg	%rcx,%rax
	xchg	%rcx,%rcx
	xchg	%rcx,%rdx
	xchg	%rcx,%rbx
	xchg	%rcx,%rsp
	xchg	%rcx,%rbp
	xchg	%rcx,%rsi
	xchg	%rcx,%rdi
	xchg	%rcx,%r8
	xchg	%rcx,%r9
	xchg	%rcx,%r10
	xchg	%rcx,%r11
	xchg	%rcx,%r12
	xchg	%rcx,%r13
	xchg	%rcx,%r14
	xchg	%rcx,%r15
	nop
	xchg	%rdx,%rax
	xchg	%rdx,%rcx
	xchg	%rdx,%rdx
	xchg	%rdx,%rbx
	xchg	%rdx,%rsp
	xchg	%rdx,%rbp
	xchg	%rdx,%rsi
	xchg	%rdx,%rdi
	xchg	%rdx,%r8
	xchg	%rdx,%r9
	xchg	%rdx,%r10
	xchg	%rdx,%r11
	xchg	%rdx,%r12
	xchg	%rdx,%r13
	xchg	%rdx,%r14
	xchg	%rdx,%r15
	nop
	xchg	%rbx,%rax
	xchg	%rbx,%rcx
	xchg	%rbx,%rdx
	xchg	%rbx,%rbx
	xchg	%rbx,%rsp
	xchg	%rbx,%rbp
	xchg	%rbx,%rsi
	xchg	%rbx,%rdi
	xchg	%rbx,%r8
	xchg	%rbx,%r9
	xchg	%rbx,%r10
	xchg	%rbx,%r11
	xchg	%rbx,%r12
	xchg	%rbx,%r13
	xchg	%rbx,%r14
	xchg	%rbx,%r15
	nop
	xchg	%rsp,%rax
	xchg	%rsp,%rcx
	xchg	%rsp,%rdx
	xchg	%rsp,%rbx
	xchg	%rsp,%rsp
	xchg	%rsp,%rbp
	xchg	%rsp,%rsi
	xchg	%rsp,%rdi
	xchg	%rsp,%r8
	xchg	%rsp,%r9
	xchg	%rsp,%r10
	xchg	%rsp,%r11
	xchg	%rsp,%r12
	xchg	%rsp,%r13
	xchg	%rsp,%r14
	xchg	%rsp,%r15
	nop
	xchg	%rbp,%rax
	xchg	%rbp,%rcx
	xchg	%rbp,%rdx
	xchg	%rbp,%rbx
	xchg	%rbp,%rsp
	xchg	%rbp,%rbp
	xchg	%rbp,%rsi
	xchg	%rbp,%rdi
	xchg	%rbp,%r8
	xchg	%rbp,%r9
	xchg	%rbp,%r10
	xchg	%rbp,%r11
	xchg	%rbp,%r12
	xchg	%rbp,%r13
	xchg	%rbp,%r14
	xchg	%rbp,%r15
	nop
	xchg	%rsi,%rax
	xchg	%rsi,%rcx
	xchg	%rsi,%rdx
	xchg	%rsi,%rbx
	xchg	%rsi,%rsp
	xchg	%rsi,%rbp
	xchg	%rsi,%rsi
	xchg	%rsi,%rdi
	xchg	%rsi,%r8
	xchg	%rsi,%r9
	xchg	%rsi,%r10
	xchg	%rsi,%r11
	xchg	%rsi,%r12
	xchg	%rsi,%r13
	xchg	%rsi,%r14
	xchg	%rsi,%r15
	nop
	xchg	%rdi,%rax
	xchg	%rdi,%rcx
	xchg	%rdi,%rdx
	xchg	%rdi,%rbx
	xchg	%rdi,%rsp
	xchg	%rdi,%rbp
	xchg	%rdi,%rsi
	xchg	%rdi,%rdi
	xchg	%rdi,%r8
	xchg	%rdi,%r9
	xchg	%rdi,%r10
	xchg	%rdi,%r11
	xchg	%rdi,%r12
	xchg	%rdi,%r13
	xchg	%rdi,%r14
	xchg	%rdi,%r15
	nop
	xchg	%r8, %rax
	xchg	%r8, %rcx
	xchg	%r8, %rdx
	xchg	%r8, %rbx
	xchg	%r8, %rsp
	xchg	%r8, %rbp
	xchg	%r8, %rsi
	xchg	%r8, %rdi
	xchg	%r8, %r8
	xchg	%r8, %r9
	xchg	%r8, %r10
	xchg	%r8, %r11
	xchg	%r8, %r12
	xchg	%r8, %r13
	xchg	%r8, %r14
	xchg	%r8, %r15
	nop
	xchg	%r9, %rax
	xchg	%r9, %rcx
	xchg	%r9, %rdx
	xchg	%r9, %rbx
	xchg	%r9, %rsp
	xchg	%r9, %rbp
	xchg	%r9, %rsi
	xchg	%r9, %rdi
	xchg	%r9, %r8
	xchg	%r9, %r9
	xchg	%r9, %r10
	xchg	%r9, %r11
	xchg	%r9, %r12
	xchg	%r9, %r13
	xchg	%r9, %r14
	xchg	%r9, %r15
	nop
	xchg	%r10,%rax
	xchg	%r10,%rcx
	xchg	%r10,%rdx
	xchg	%r10,%rbx
	xchg	%r10,%rsp
	xchg	%r10,%rbp
	xchg	%r10,%rsi
	xchg	%r10,%rdi
	xchg	%r10,%r8
	xchg	%r10,%r9
	xchg	%r10,%r10
	xchg	%r10,%r11
	xchg	%r10,%r12
	xchg	%r10,%r13
	xchg	%r10,%r14
	xchg	%r10,%r15
	nop
	xchg	%r11,%rax
	xchg	%r11,%rcx
	xchg	%r11,%rdx
	xchg	%r11,%rbx
	xchg	%r11,%rsp
	xchg	%r11,%rbp
	xchg	%r11,%rsi
	xchg	%r11,%rdi
	xchg	%r11,%r8
	xchg	%r11,%r9
	xchg	%r11,%r10
	xchg	%r11,%r11
	xchg	%r11,%r12
	xchg	%r11,%r13
	xchg	%r11,%r14
	xchg	%r11,%r15
	nop
	xchg	%r12,%rax
	xchg	%r12,%rcx
	xchg	%r12,%rdx
	xchg	%r12,%rbx
	xchg	%r12,%rsp
	xchg	%r12,%rbp
	xchg	%r12,%rsi
	xchg	%r12,%rdi
	xchg	%r12,%r8
	xchg	%r12,%r9
	xchg	%r12,%r10
	xchg	%r12,%r11
	xchg	%r12,%r12
	xchg	%r12,%r13
	xchg	%r12,%r14
	xchg	%r12,%r15
	nop
	xchg	%r13,%rax
	xchg	%r13,%rcx
	xchg	%r13,%rdx
	xchg	%r13,%rbx
	xchg	%r13,%rsp
	xchg	%r13,%rbp
	xchg	%r13,%rsi
	xchg	%r13,%rdi
	xchg	%r13,%r8
	xchg	%r13,%r9
	xchg	%r13,%r10
	xchg	%r13,%r11
	xchg	%r13,%r12
	xchg	%r13,%r13
	xchg	%r13,%r14
	xchg	%r13,%r15
	nop
	xchg	%r14,%rax
	xchg	%r14,%rcx
	xchg	%r14,%rdx
	xchg	%r14,%rbx
	xchg	%r14,%rsp
	xchg	%r14,%rbp
	xchg	%r14,%rsi
	xchg	%r14,%rdi
	xchg	%r14,%r8
	xchg	%r14,%r9
	xchg	%r14,%r10
	xchg	%r14,%r11
	xchg	%r14,%r12
	xchg	%r14,%r13
	xchg	%r14,%r14
	xchg	%r14,%r15
	nop
	xchg	%r15,%rax
	xchg	%r15,%rcx
	xchg	%r15,%rdx
	xchg	%r15,%rbx
	xchg	%r15,%rsp
	xchg	%r15,%rbp
	xchg	%r15,%rsi
	xchg	%r15,%rdi
	xchg	%r15,%r8
	xchg	%r15,%r9
	xchg	%r15,%r10
	xchg	%r15,%r11
	xchg	%r15,%r12
	xchg	%r15,%r13
	xchg	%r15,%r14
	xchg	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XchgMemReg
	.type	XchgMemReg, @function
XchgMemReg:
	.cfi_startproc
	xchg	%rax,(%rax)
	xchg	%rax,(%rcx)
	xchg	%rax,(%rdx)
	xchg	%rax,(%rbx)
	xchg	%rax,(%rsp)
	xchg	%rax,(%rbp)
	xchg	%rax,(%rsi)
	xchg	%rax,(%rdi)
	xchg	%rax,(%r8)
	xchg	%rax,(%r9)
	xchg	%rax,(%r10)
	xchg	%rax,(%r11)
	xchg	%rax,(%r12)
	xchg	%rax,(%r13)
	xchg	%rax,(%r14)
	xchg	%rax,(%r15)
	nop
	xchg	%rcx,(%rax)
	xchg	%rcx,(%rcx)
	xchg	%rcx,(%rdx)
	xchg	%rcx,(%rbx)
	xchg	%rcx,(%rsp)
	xchg	%rcx,(%rbp)
	xchg	%rcx,(%rsi)
	xchg	%rcx,(%rdi)
	xchg	%rcx,(%r8)
	xchg	%rcx,(%r9)
	xchg	%rcx,(%r10)
	xchg	%rcx,(%r11)
	xchg	%rcx,(%r12)
	xchg	%rcx,(%r13)
	xchg	%rcx,(%r14)
	xchg	%rcx,(%r15)
	nop
	xchg	%rdx,(%rax)
	xchg	%rdx,(%rcx)
	xchg	%rdx,(%rdx)
	xchg	%rdx,(%rbx)
	xchg	%rdx,(%rsp)
	xchg	%rdx,(%rbp)
	xchg	%rdx,(%rsi)
	xchg	%rdx,(%rdi)
	xchg	%rdx,(%r8)
	xchg	%rdx,(%r9)
	xchg	%rdx,(%r10)
	xchg	%rdx,(%r11)
	xchg	%rdx,(%r12)
	xchg	%rdx,(%r13)
	xchg	%rdx,(%r14)
	xchg	%rdx,(%r15)
	nop
	xchg	%rbx,(%rax)
	xchg	%rbx,(%rcx)
	xchg	%rbx,(%rdx)
	xchg	%rbx,(%rbx)
	xchg	%rbx,(%rsp)
	xchg	%rbx,(%rbp)
	xchg	%rbx,(%rsi)
	xchg	%rbx,(%rdi)
	xchg	%rbx,(%r8)
	xchg	%rbx,(%r9)
	xchg	%rbx,(%r10)
	xchg	%rbx,(%r11)
	xchg	%rbx,(%r12)
	xchg	%rbx,(%r13)
	xchg	%rbx,(%r14)
	xchg	%rbx,(%r15)
	nop
	xchg	%rsp,(%rax)
	xchg	%rsp,(%rcx)
	xchg	%rsp,(%rdx)
	xchg	%rsp,(%rbx)
	xchg	%rsp,(%rsp)
	xchg	%rsp,(%rbp)
	xchg	%rsp,(%rsi)
	xchg	%rsp,(%rdi)
	xchg	%rsp,(%r8)
	xchg	%rsp,(%r9)
	xchg	%rsp,(%r10)
	xchg	%rsp,(%r11)
	xchg	%rsp,(%r12)
	xchg	%rsp,(%r13)
	xchg	%rsp,(%r14)
	xchg	%rsp,(%r15)
	nop
	xchg	%rbp,(%rax)
	xchg	%rbp,(%rcx)
	xchg	%rbp,(%rdx)
	xchg	%rbp,(%rbx)
	xchg	%rbp,(%rsp)
	xchg	%rbp,(%rbp)
	xchg	%rbp,(%rsi)
	xchg	%rbp,(%rdi)
	xchg	%rbp,(%r8)
	xchg	%rbp,(%r9)
	xchg	%rbp,(%r10)
	xchg	%rbp,(%r11)
	xchg	%rbp,(%r12)
	xchg	%rbp,(%r13)
	xchg	%rbp,(%r14)
	xchg	%rbp,(%r15)
	nop
	xchg	%rsi,(%rax)
	xchg	%rsi,(%rcx)
	xchg	%rsi,(%rdx)
	xchg	%rsi,(%rbx)
	xchg	%rsi,(%rsp)
	xchg	%rsi,(%rbp)
	xchg	%rsi,(%rsi)
	xchg	%rsi,(%rdi)
	xchg	%rsi,(%r8)
	xchg	%rsi,(%r9)
	xchg	%rsi,(%r10)
	xchg	%rsi,(%r11)
	xchg	%rsi,(%r12)
	xchg	%rsi,(%r13)
	xchg	%rsi,(%r14)
	xchg	%rsi,(%r15)
	nop
	xchg	%rdi,(%rax)
	xchg	%rdi,(%rcx)
	xchg	%rdi,(%rdx)
	xchg	%rdi,(%rbx)
	xchg	%rdi,(%rsp)
	xchg	%rdi,(%rbp)
	xchg	%rdi,(%rsi)
	xchg	%rdi,(%rdi)
	xchg	%rdi,(%r8)
	xchg	%rdi,(%r9)
	xchg	%rdi,(%r10)
	xchg	%rdi,(%r11)
	xchg	%rdi,(%r12)
	xchg	%rdi,(%r13)
	xchg	%rdi,(%r14)
	xchg	%rdi,(%r15)
	nop
	xchg	%r8, (%rax)
	xchg	%r8, (%rcx)
	xchg	%r8, (%rdx)
	xchg	%r8, (%rbx)
	xchg	%r8, (%rsp)
	xchg	%r8, (%rbp)
	xchg	%r8, (%rsi)
	xchg	%r8, (%rdi)
	xchg	%r8, (%r8)
	xchg	%r8, (%r9)
	xchg	%r8, (%r10)
	xchg	%r8, (%r11)
	xchg	%r8, (%r12)
	xchg	%r8, (%r13)
	xchg	%r8, (%r14)
	xchg	%r8, (%r15)
	nop
	xchg	%r9, (%rax)
	xchg	%r9, (%rcx)
	xchg	%r9, (%rdx)
	xchg	%r9, (%rbx)
	xchg	%r9, (%rsp)
	xchg	%r9, (%rbp)
	xchg	%r9, (%rsi)
	xchg	%r9, (%rdi)
	xchg	%r9, (%r8)
	xchg	%r9, (%r9)
	xchg	%r9, (%r10)
	xchg	%r9, (%r11)
	xchg	%r9, (%r12)
	xchg	%r9, (%r13)
	xchg	%r9, (%r14)
	xchg	%r9, (%r15)
	nop
	xchg	%r10,(%rax)
	xchg	%r10,(%rcx)
	xchg	%r10,(%rdx)
	xchg	%r10,(%rbx)
	xchg	%r10,(%rsp)
	xchg	%r10,(%rbp)
	xchg	%r10,(%rsi)
	xchg	%r10,(%rdi)
	xchg	%r10,(%r8)
	xchg	%r10,(%r9)
	xchg	%r10,(%r10)
	xchg	%r10,(%r11)
	xchg	%r10,(%r12)
	xchg	%r10,(%r13)
	xchg	%r10,(%r14)
	xchg	%r10,(%r15)
	nop
	xchg	%r11,(%rax)
	xchg	%r11,(%rcx)
	xchg	%r11,(%rdx)
	xchg	%r11,(%rbx)
	xchg	%r11,(%rsp)
	xchg	%r11,(%rbp)
	xchg	%r11,(%rsi)
	xchg	%r11,(%rdi)
	xchg	%r11,(%r8)
	xchg	%r11,(%r9)
	xchg	%r11,(%r10)
	xchg	%r11,(%r11)
	xchg	%r11,(%r12)
	xchg	%r11,(%r13)
	xchg	%r11,(%r14)
	xchg	%r11,(%r15)
	nop
	xchg	%r12,(%rax)
	xchg	%r12,(%rcx)
	xchg	%r12,(%rdx)
	xchg	%r12,(%rbx)
	xchg	%r12,(%rsp)
	xchg	%r12,(%rbp)
	xchg	%r12,(%rsi)
	xchg	%r12,(%rdi)
	xchg	%r12,(%r8)
	xchg	%r12,(%r9)
	xchg	%r12,(%r10)
	xchg	%r12,(%r11)
	xchg	%r12,(%r12)
	xchg	%r12,(%r13)
	xchg	%r12,(%r14)
	xchg	%r12,(%r15)
	nop
	xchg	%r13,(%rax)
	xchg	%r13,(%rcx)
	xchg	%r13,(%rdx)
	xchg	%r13,(%rbx)
	xchg	%r13,(%rsp)
	xchg	%r13,(%rbp)
	xchg	%r13,(%rsi)
	xchg	%r13,(%rdi)
	xchg	%r13,(%r8)
	xchg	%r13,(%r9)
	xchg	%r13,(%r10)
	xchg	%r13,(%r11)
	xchg	%r13,(%r12)
	xchg	%r13,(%r13)
	xchg	%r13,(%r14)
	xchg	%r13,(%r15)
	nop
	xchg	%r14,(%rax)
	xchg	%r14,(%rcx)
	xchg	%r14,(%rdx)
	xchg	%r14,(%rbx)
	xchg	%r14,(%rsp)
	xchg	%r14,(%rbp)
	xchg	%r14,(%rsi)
	xchg	%r14,(%rdi)
	xchg	%r14,(%r8)
	xchg	%r14,(%r9)
	xchg	%r14,(%r10)
	xchg	%r14,(%r11)
	xchg	%r14,(%r12)
	xchg	%r14,(%r13)
	xchg	%r14,(%r14)
	xchg	%r14,(%r15)
	nop
	xchg	%r15,(%rax)
	xchg	%r15,(%rcx)
	xchg	%r15,(%rdx)
	xchg	%r15,(%rbx)
	xchg	%r15,(%rsp)
	xchg	%r15,(%rbp)
	xchg	%r15,(%rsi)
	xchg	%r15,(%rdi)
	xchg	%r15,(%r8)
	xchg	%r15,(%r9)
	xchg	%r15,(%r10)
	xchg	%r15,(%r11)
	xchg	%r15,(%r12)
	xchg	%r15,(%r13)
	xchg	%r15,(%r14)
	xchg	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XchgMem8Reg
	.type	XchgMem8Reg, @function
XchgMem8Reg:
	.cfi_startproc
	xchg	%rax,0x7f(%rax)
	xchg	%rax,0x7f(%rcx)
	xchg	%rax,0x7f(%rdx)
	xchg	%rax,0x7f(%rbx)
	xchg	%rax,0x7f(%rsp)
	xchg	%rax,0x7f(%rbp)
	xchg	%rax,0x7f(%rsi)
	xchg	%rax,0x7f(%rdi)
	xchg	%rax,0x7f(%r8)
	xchg	%rax,0x7f(%r9)
	xchg	%rax,0x7f(%r10)
	xchg	%rax,0x7f(%r11)
	xchg	%rax,0x7f(%r12)
	xchg	%rax,0x7f(%r13)
	xchg	%rax,0x7f(%r14)
	xchg	%rax,0x7f(%r15)
	nop
	xchg	%rcx,0x7f(%rax)
	xchg	%rcx,0x7f(%rcx)
	xchg	%rcx,0x7f(%rdx)
	xchg	%rcx,0x7f(%rbx)
	xchg	%rcx,0x7f(%rsp)
	xchg	%rcx,0x7f(%rbp)
	xchg	%rcx,0x7f(%rsi)
	xchg	%rcx,0x7f(%rdi)
	xchg	%rcx,0x7f(%r8)
	xchg	%rcx,0x7f(%r9)
	xchg	%rcx,0x7f(%r10)
	xchg	%rcx,0x7f(%r11)
	xchg	%rcx,0x7f(%r12)
	xchg	%rcx,0x7f(%r13)
	xchg	%rcx,0x7f(%r14)
	xchg	%rcx,0x7f(%r15)
	nop
	xchg	%rdx,0x7f(%rax)
	xchg	%rdx,0x7f(%rcx)
	xchg	%rdx,0x7f(%rdx)
	xchg	%rdx,0x7f(%rbx)
	xchg	%rdx,0x7f(%rsp)
	xchg	%rdx,0x7f(%rbp)
	xchg	%rdx,0x7f(%rsi)
	xchg	%rdx,0x7f(%rdi)
	xchg	%rdx,0x7f(%r8)
	xchg	%rdx,0x7f(%r9)
	xchg	%rdx,0x7f(%r10)
	xchg	%rdx,0x7f(%r11)
	xchg	%rdx,0x7f(%r12)
	xchg	%rdx,0x7f(%r13)
	xchg	%rdx,0x7f(%r14)
	xchg	%rdx,0x7f(%r15)
	nop
	xchg	%rbx,0x7f(%rax)
	xchg	%rbx,0x7f(%rcx)
	xchg	%rbx,0x7f(%rdx)
	xchg	%rbx,0x7f(%rbx)
	xchg	%rbx,0x7f(%rsp)
	xchg	%rbx,0x7f(%rbp)
	xchg	%rbx,0x7f(%rsi)
	xchg	%rbx,0x7f(%rdi)
	xchg	%rbx,0x7f(%r8)
	xchg	%rbx,0x7f(%r9)
	xchg	%rbx,0x7f(%r10)
	xchg	%rbx,0x7f(%r11)
	xchg	%rbx,0x7f(%r12)
	xchg	%rbx,0x7f(%r13)
	xchg	%rbx,0x7f(%r14)
	xchg	%rbx,0x7f(%r15)
	nop
	xchg	%rsp,0x7f(%rax)
	xchg	%rsp,0x7f(%rcx)
	xchg	%rsp,0x7f(%rdx)
	xchg	%rsp,0x7f(%rbx)
	xchg	%rsp,0x7f(%rsp)
	xchg	%rsp,0x7f(%rbp)
	xchg	%rsp,0x7f(%rsi)
	xchg	%rsp,0x7f(%rdi)
	xchg	%rsp,0x7f(%r8)
	xchg	%rsp,0x7f(%r9)
	xchg	%rsp,0x7f(%r10)
	xchg	%rsp,0x7f(%r11)
	xchg	%rsp,0x7f(%r12)
	xchg	%rsp,0x7f(%r13)
	xchg	%rsp,0x7f(%r14)
	xchg	%rsp,0x7f(%r15)
	nop
	xchg	%rbp,0x7f(%rax)
	xchg	%rbp,0x7f(%rcx)
	xchg	%rbp,0x7f(%rdx)
	xchg	%rbp,0x7f(%rbx)
	xchg	%rbp,0x7f(%rsp)
	xchg	%rbp,0x7f(%rbp)
	xchg	%rbp,0x7f(%rsi)
	xchg	%rbp,0x7f(%rdi)
	xchg	%rbp,0x7f(%r8)
	xchg	%rbp,0x7f(%r9)
	xchg	%rbp,0x7f(%r10)
	xchg	%rbp,0x7f(%r11)
	xchg	%rbp,0x7f(%r12)
	xchg	%rbp,0x7f(%r13)
	xchg	%rbp,0x7f(%r14)
	xchg	%rbp,0x7f(%r15)
	nop
	xchg	%rsi,0x7f(%rax)
	xchg	%rsi,0x7f(%rcx)
	xchg	%rsi,0x7f(%rdx)
	xchg	%rsi,0x7f(%rbx)
	xchg	%rsi,0x7f(%rsp)
	xchg	%rsi,0x7f(%rbp)
	xchg	%rsi,0x7f(%rsi)
	xchg	%rsi,0x7f(%rdi)
	xchg	%rsi,0x7f(%r8)
	xchg	%rsi,0x7f(%r9)
	xchg	%rsi,0x7f(%r10)
	xchg	%rsi,0x7f(%r11)
	xchg	%rsi,0x7f(%r12)
	xchg	%rsi,0x7f(%r13)
	xchg	%rsi,0x7f(%r14)
	xchg	%rsi,0x7f(%r15)
	nop
	xchg	%rdi,0x7f(%rax)
	xchg	%rdi,0x7f(%rcx)
	xchg	%rdi,0x7f(%rdx)
	xchg	%rdi,0x7f(%rbx)
	xchg	%rdi,0x7f(%rsp)
	xchg	%rdi,0x7f(%rbp)
	xchg	%rdi,0x7f(%rsi)
	xchg	%rdi,0x7f(%rdi)
	xchg	%rdi,0x7f(%r8)
	xchg	%rdi,0x7f(%r9)
	xchg	%rdi,0x7f(%r10)
	xchg	%rdi,0x7f(%r11)
	xchg	%rdi,0x7f(%r12)
	xchg	%rdi,0x7f(%r13)
	xchg	%rdi,0x7f(%r14)
	xchg	%rdi,0x7f(%r15)
	nop
	xchg	%r8, 0x7f(%rax)
	xchg	%r8, 0x7f(%rcx)
	xchg	%r8, 0x7f(%rdx)
	xchg	%r8, 0x7f(%rbx)
	xchg	%r8, 0x7f(%rsp)
	xchg	%r8, 0x7f(%rbp)
	xchg	%r8, 0x7f(%rsi)
	xchg	%r8, 0x7f(%rdi)
	xchg	%r8, 0x7f(%r8)
	xchg	%r8, 0x7f(%r9)
	xchg	%r8, 0x7f(%r10)
	xchg	%r8, 0x7f(%r11)
	xchg	%r8, 0x7f(%r12)
	xchg	%r8, 0x7f(%r13)
	xchg	%r8, 0x7f(%r14)
	xchg	%r8, 0x7f(%r15)
	nop
	xchg	%r9, 0x7f(%rax)
	xchg	%r9, 0x7f(%rcx)
	xchg	%r9, 0x7f(%rdx)
	xchg	%r9, 0x7f(%rbx)
	xchg	%r9, 0x7f(%rsp)
	xchg	%r9, 0x7f(%rbp)
	xchg	%r9, 0x7f(%rsi)
	xchg	%r9, 0x7f(%rdi)
	xchg	%r9, 0x7f(%r8)
	xchg	%r9, 0x7f(%r9)
	xchg	%r9, 0x7f(%r10)
	xchg	%r9, 0x7f(%r11)
	xchg	%r9, 0x7f(%r12)
	xchg	%r9, 0x7f(%r13)
	xchg	%r9, 0x7f(%r14)
	xchg	%r9, 0x7f(%r15)
	nop
	xchg	%r10,0x7f(%rax)
	xchg	%r10,0x7f(%rcx)
	xchg	%r10,0x7f(%rdx)
	xchg	%r10,0x7f(%rbx)
	xchg	%r10,0x7f(%rsp)
	xchg	%r10,0x7f(%rbp)
	xchg	%r10,0x7f(%rsi)
	xchg	%r10,0x7f(%rdi)
	xchg	%r10,0x7f(%r8)
	xchg	%r10,0x7f(%r9)
	xchg	%r10,0x7f(%r10)
	xchg	%r10,0x7f(%r11)
	xchg	%r10,0x7f(%r12)
	xchg	%r10,0x7f(%r13)
	xchg	%r10,0x7f(%r14)
	xchg	%r10,0x7f(%r15)
	nop
	xchg	%r11,0x7f(%rax)
	xchg	%r11,0x7f(%rcx)
	xchg	%r11,0x7f(%rdx)
	xchg	%r11,0x7f(%rbx)
	xchg	%r11,0x7f(%rsp)
	xchg	%r11,0x7f(%rbp)
	xchg	%r11,0x7f(%rsi)
	xchg	%r11,0x7f(%rdi)
	xchg	%r11,0x7f(%r8)
	xchg	%r11,0x7f(%r9)
	xchg	%r11,0x7f(%r10)
	xchg	%r11,0x7f(%r11)
	xchg	%r11,0x7f(%r12)
	xchg	%r11,0x7f(%r13)
	xchg	%r11,0x7f(%r14)
	xchg	%r11,0x7f(%r15)
	nop
	xchg	%r12,0x7f(%rax)
	xchg	%r12,0x7f(%rcx)
	xchg	%r12,0x7f(%rdx)
	xchg	%r12,0x7f(%rbx)
	xchg	%r12,0x7f(%rsp)
	xchg	%r12,0x7f(%rbp)
	xchg	%r12,0x7f(%rsi)
	xchg	%r12,0x7f(%rdi)
	xchg	%r12,0x7f(%r8)
	xchg	%r12,0x7f(%r9)
	xchg	%r12,0x7f(%r10)
	xchg	%r12,0x7f(%r11)
	xchg	%r12,0x7f(%r12)
	xchg	%r12,0x7f(%r13)
	xchg	%r12,0x7f(%r14)
	xchg	%r12,0x7f(%r15)
	nop
	xchg	%r13,0x7f(%rax)
	xchg	%r13,0x7f(%rcx)
	xchg	%r13,0x7f(%rdx)
	xchg	%r13,0x7f(%rbx)
	xchg	%r13,0x7f(%rsp)
	xchg	%r13,0x7f(%rbp)
	xchg	%r13,0x7f(%rsi)
	xchg	%r13,0x7f(%rdi)
	xchg	%r13,0x7f(%r8)
	xchg	%r13,0x7f(%r9)
	xchg	%r13,0x7f(%r10)
	xchg	%r13,0x7f(%r11)
	xchg	%r13,0x7f(%r12)
	xchg	%r13,0x7f(%r13)
	xchg	%r13,0x7f(%r14)
	xchg	%r13,0x7f(%r15)
	nop
	xchg	%r14,0x7f(%rax)
	xchg	%r14,0x7f(%rcx)
	xchg	%r14,0x7f(%rdx)
	xchg	%r14,0x7f(%rbx)
	xchg	%r14,0x7f(%rsp)
	xchg	%r14,0x7f(%rbp)
	xchg	%r14,0x7f(%rsi)
	xchg	%r14,0x7f(%rdi)
	xchg	%r14,0x7f(%r8)
	xchg	%r14,0x7f(%r9)
	xchg	%r14,0x7f(%r10)
	xchg	%r14,0x7f(%r11)
	xchg	%r14,0x7f(%r12)
	xchg	%r14,0x7f(%r13)
	xchg	%r14,0x7f(%r14)
	xchg	%r14,0x7f(%r15)
	nop
	xchg	%r15,0x7f(%rax)
	xchg	%r15,0x7f(%rcx)
	xchg	%r15,0x7f(%rdx)
	xchg	%r15,0x7f(%rbx)
	xchg	%r15,0x7f(%rsp)
	xchg	%r15,0x7f(%rbp)
	xchg	%r15,0x7f(%rsi)
	xchg	%r15,0x7f(%rdi)
	xchg	%r15,0x7f(%r8)
	xchg	%r15,0x7f(%r9)
	xchg	%r15,0x7f(%r10)
	xchg	%r15,0x7f(%r11)
	xchg	%r15,0x7f(%r12)
	xchg	%r15,0x7f(%r13)
	xchg	%r15,0x7f(%r14)
	xchg	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XchgMem32Reg
	.type	XchgMem32Reg, @function
XchgMem32Reg:
	.cfi_startproc
	xchg	%rax,0x7f563412(%rax)
	xchg	%rax,0x7f563412(%rcx)
	xchg	%rax,0x7f563412(%rdx)
	xchg	%rax,0x7f563412(%rbx)
	xchg	%rax,0x7f563412(%rsp)
	xchg	%rax,0x7f563412(%rbp)
	xchg	%rax,0x7f563412(%rsi)
	xchg	%rax,0x7f563412(%rdi)
	xchg	%rax,0x7f563412(%r8)
	xchg	%rax,0x7f563412(%r9)
	xchg	%rax,0x7f563412(%r10)
	xchg	%rax,0x7f563412(%r11)
	xchg	%rax,0x7f563412(%r12)
	xchg	%rax,0x7f563412(%r13)
	xchg	%rax,0x7f563412(%r14)
	xchg	%rax,0x7f563412(%r15)
	nop
	xchg	%rcx,0x7f563412(%rax)
	xchg	%rcx,0x7f563412(%rcx)
	xchg	%rcx,0x7f563412(%rdx)
	xchg	%rcx,0x7f563412(%rbx)
	xchg	%rcx,0x7f563412(%rsp)
	xchg	%rcx,0x7f563412(%rbp)
	xchg	%rcx,0x7f563412(%rsi)
	xchg	%rcx,0x7f563412(%rdi)
	xchg	%rcx,0x7f563412(%r8)
	xchg	%rcx,0x7f563412(%r9)
	xchg	%rcx,0x7f563412(%r10)
	xchg	%rcx,0x7f563412(%r11)
	xchg	%rcx,0x7f563412(%r12)
	xchg	%rcx,0x7f563412(%r13)
	xchg	%rcx,0x7f563412(%r14)
	xchg	%rcx,0x7f563412(%r15)
	nop
	xchg	%rdx,0x7f563412(%rax)
	xchg	%rdx,0x7f563412(%rcx)
	xchg	%rdx,0x7f563412(%rdx)
	xchg	%rdx,0x7f563412(%rbx)
	xchg	%rdx,0x7f563412(%rsp)
	xchg	%rdx,0x7f563412(%rbp)
	xchg	%rdx,0x7f563412(%rsi)
	xchg	%rdx,0x7f563412(%rdi)
	xchg	%rdx,0x7f563412(%r8)
	xchg	%rdx,0x7f563412(%r9)
	xchg	%rdx,0x7f563412(%r10)
	xchg	%rdx,0x7f563412(%r11)
	xchg	%rdx,0x7f563412(%r12)
	xchg	%rdx,0x7f563412(%r13)
	xchg	%rdx,0x7f563412(%r14)
	xchg	%rdx,0x7f563412(%r15)
	nop
	xchg	%rbx,0x7f563412(%rax)
	xchg	%rbx,0x7f563412(%rcx)
	xchg	%rbx,0x7f563412(%rdx)
	xchg	%rbx,0x7f563412(%rbx)
	xchg	%rbx,0x7f563412(%rsp)
	xchg	%rbx,0x7f563412(%rbp)
	xchg	%rbx,0x7f563412(%rsi)
	xchg	%rbx,0x7f563412(%rdi)
	xchg	%rbx,0x7f563412(%r8)
	xchg	%rbx,0x7f563412(%r9)
	xchg	%rbx,0x7f563412(%r10)
	xchg	%rbx,0x7f563412(%r11)
	xchg	%rbx,0x7f563412(%r12)
	xchg	%rbx,0x7f563412(%r13)
	xchg	%rbx,0x7f563412(%r14)
	xchg	%rbx,0x7f563412(%r15)
	nop
	xchg	%rsp,0x7f563412(%rax)
	xchg	%rsp,0x7f563412(%rcx)
	xchg	%rsp,0x7f563412(%rdx)
	xchg	%rsp,0x7f563412(%rbx)
	xchg	%rsp,0x7f563412(%rsp)
	xchg	%rsp,0x7f563412(%rbp)
	xchg	%rsp,0x7f563412(%rsi)
	xchg	%rsp,0x7f563412(%rdi)
	xchg	%rsp,0x7f563412(%r8)
	xchg	%rsp,0x7f563412(%r9)
	xchg	%rsp,0x7f563412(%r10)
	xchg	%rsp,0x7f563412(%r11)
	xchg	%rsp,0x7f563412(%r12)
	xchg	%rsp,0x7f563412(%r13)
	xchg	%rsp,0x7f563412(%r14)
	xchg	%rsp,0x7f563412(%r15)
	nop
	xchg	%rbp,0x7f563412(%rax)
	xchg	%rbp,0x7f563412(%rcx)
	xchg	%rbp,0x7f563412(%rdx)
	xchg	%rbp,0x7f563412(%rbx)
	xchg	%rbp,0x7f563412(%rsp)
	xchg	%rbp,0x7f563412(%rbp)
	xchg	%rbp,0x7f563412(%rsi)
	xchg	%rbp,0x7f563412(%rdi)
	xchg	%rbp,0x7f563412(%r8)
	xchg	%rbp,0x7f563412(%r9)
	xchg	%rbp,0x7f563412(%r10)
	xchg	%rbp,0x7f563412(%r11)
	xchg	%rbp,0x7f563412(%r12)
	xchg	%rbp,0x7f563412(%r13)
	xchg	%rbp,0x7f563412(%r14)
	xchg	%rbp,0x7f563412(%r15)
	nop
	xchg	%rsi,0x7f563412(%rax)
	xchg	%rsi,0x7f563412(%rcx)
	xchg	%rsi,0x7f563412(%rdx)
	xchg	%rsi,0x7f563412(%rbx)
	xchg	%rsi,0x7f563412(%rsp)
	xchg	%rsi,0x7f563412(%rbp)
	xchg	%rsi,0x7f563412(%rsi)
	xchg	%rsi,0x7f563412(%rdi)
	xchg	%rsi,0x7f563412(%r8)
	xchg	%rsi,0x7f563412(%r9)
	xchg	%rsi,0x7f563412(%r10)
	xchg	%rsi,0x7f563412(%r11)
	xchg	%rsi,0x7f563412(%r12)
	xchg	%rsi,0x7f563412(%r13)
	xchg	%rsi,0x7f563412(%r14)
	xchg	%rsi,0x7f563412(%r15)
	nop
	xchg	%rdi,0x7f563412(%rax)
	xchg	%rdi,0x7f563412(%rcx)
	xchg	%rdi,0x7f563412(%rdx)
	xchg	%rdi,0x7f563412(%rbx)
	xchg	%rdi,0x7f563412(%rsp)
	xchg	%rdi,0x7f563412(%rbp)
	xchg	%rdi,0x7f563412(%rsi)
	xchg	%rdi,0x7f563412(%rdi)
	xchg	%rdi,0x7f563412(%r8)
	xchg	%rdi,0x7f563412(%r9)
	xchg	%rdi,0x7f563412(%r10)
	xchg	%rdi,0x7f563412(%r11)
	xchg	%rdi,0x7f563412(%r12)
	xchg	%rdi,0x7f563412(%r13)
	xchg	%rdi,0x7f563412(%r14)
	xchg	%rdi,0x7f563412(%r15)
	nop
	xchg	%r8, 0x7f563412(%rax)
	xchg	%r8, 0x7f563412(%rcx)
	xchg	%r8, 0x7f563412(%rdx)
	xchg	%r8, 0x7f563412(%rbx)
	xchg	%r8, 0x7f563412(%rsp)
	xchg	%r8, 0x7f563412(%rbp)
	xchg	%r8, 0x7f563412(%rsi)
	xchg	%r8, 0x7f563412(%rdi)
	xchg	%r8, 0x7f563412(%r8)
	xchg	%r8, 0x7f563412(%r9)
	xchg	%r8, 0x7f563412(%r10)
	xchg	%r8, 0x7f563412(%r11)
	xchg	%r8, 0x7f563412(%r12)
	xchg	%r8, 0x7f563412(%r13)
	xchg	%r8, 0x7f563412(%r14)
	xchg	%r8, 0x7f563412(%r15)
	nop
	xchg	%r9, 0x7f563412(%rax)
	xchg	%r9, 0x7f563412(%rcx)
	xchg	%r9, 0x7f563412(%rdx)
	xchg	%r9, 0x7f563412(%rbx)
	xchg	%r9, 0x7f563412(%rsp)
	xchg	%r9, 0x7f563412(%rbp)
	xchg	%r9, 0x7f563412(%rsi)
	xchg	%r9, 0x7f563412(%rdi)
	xchg	%r9, 0x7f563412(%r8)
	xchg	%r9, 0x7f563412(%r9)
	xchg	%r9, 0x7f563412(%r10)
	xchg	%r9, 0x7f563412(%r11)
	xchg	%r9, 0x7f563412(%r12)
	xchg	%r9, 0x7f563412(%r13)
	xchg	%r9, 0x7f563412(%r14)
	xchg	%r9, 0x7f563412(%r15)
	nop
	xchg	%r10,0x7f563412(%rax)
	xchg	%r10,0x7f563412(%rcx)
	xchg	%r10,0x7f563412(%rdx)
	xchg	%r10,0x7f563412(%rbx)
	xchg	%r10,0x7f563412(%rsp)
	xchg	%r10,0x7f563412(%rbp)
	xchg	%r10,0x7f563412(%rsi)
	xchg	%r10,0x7f563412(%rdi)
	xchg	%r10,0x7f563412(%r8)
	xchg	%r10,0x7f563412(%r9)
	xchg	%r10,0x7f563412(%r10)
	xchg	%r10,0x7f563412(%r11)
	xchg	%r10,0x7f563412(%r12)
	xchg	%r10,0x7f563412(%r13)
	xchg	%r10,0x7f563412(%r14)
	xchg	%r10,0x7f563412(%r15)
	nop
	xchg	%r11,0x7f563412(%rax)
	xchg	%r11,0x7f563412(%rcx)
	xchg	%r11,0x7f563412(%rdx)
	xchg	%r11,0x7f563412(%rbx)
	xchg	%r11,0x7f563412(%rsp)
	xchg	%r11,0x7f563412(%rbp)
	xchg	%r11,0x7f563412(%rsi)
	xchg	%r11,0x7f563412(%rdi)
	xchg	%r11,0x7f563412(%r8)
	xchg	%r11,0x7f563412(%r9)
	xchg	%r11,0x7f563412(%r10)
	xchg	%r11,0x7f563412(%r11)
	xchg	%r11,0x7f563412(%r12)
	xchg	%r11,0x7f563412(%r13)
	xchg	%r11,0x7f563412(%r14)
	xchg	%r11,0x7f563412(%r15)
	nop
	xchg	%r12,0x7f563412(%rax)
	xchg	%r12,0x7f563412(%rcx)
	xchg	%r12,0x7f563412(%rdx)
	xchg	%r12,0x7f563412(%rbx)
	xchg	%r12,0x7f563412(%rsp)
	xchg	%r12,0x7f563412(%rbp)
	xchg	%r12,0x7f563412(%rsi)
	xchg	%r12,0x7f563412(%rdi)
	xchg	%r12,0x7f563412(%r8)
	xchg	%r12,0x7f563412(%r9)
	xchg	%r12,0x7f563412(%r10)
	xchg	%r12,0x7f563412(%r11)
	xchg	%r12,0x7f563412(%r12)
	xchg	%r12,0x7f563412(%r13)
	xchg	%r12,0x7f563412(%r14)
	xchg	%r12,0x7f563412(%r15)
	nop
	xchg	%r13,0x7f563412(%rax)
	xchg	%r13,0x7f563412(%rcx)
	xchg	%r13,0x7f563412(%rdx)
	xchg	%r13,0x7f563412(%rbx)
	xchg	%r13,0x7f563412(%rsp)
	xchg	%r13,0x7f563412(%rbp)
	xchg	%r13,0x7f563412(%rsi)
	xchg	%r13,0x7f563412(%rdi)
	xchg	%r13,0x7f563412(%r8)
	xchg	%r13,0x7f563412(%r9)
	xchg	%r13,0x7f563412(%r10)
	xchg	%r13,0x7f563412(%r11)
	xchg	%r13,0x7f563412(%r12)
	xchg	%r13,0x7f563412(%r13)
	xchg	%r13,0x7f563412(%r14)
	xchg	%r13,0x7f563412(%r15)
	nop
	xchg	%r14,0x7f563412(%rax)
	xchg	%r14,0x7f563412(%rcx)
	xchg	%r14,0x7f563412(%rdx)
	xchg	%r14,0x7f563412(%rbx)
	xchg	%r14,0x7f563412(%rsp)
	xchg	%r14,0x7f563412(%rbp)
	xchg	%r14,0x7f563412(%rsi)
	xchg	%r14,0x7f563412(%rdi)
	xchg	%r14,0x7f563412(%r8)
	xchg	%r14,0x7f563412(%r9)
	xchg	%r14,0x7f563412(%r10)
	xchg	%r14,0x7f563412(%r11)
	xchg	%r14,0x7f563412(%r12)
	xchg	%r14,0x7f563412(%r13)
	xchg	%r14,0x7f563412(%r14)
	xchg	%r14,0x7f563412(%r15)
	nop
	xchg	%r15,0x7f563412(%rax)
	xchg	%r15,0x7f563412(%rcx)
	xchg	%r15,0x7f563412(%rdx)
	xchg	%r15,0x7f563412(%rbx)
	xchg	%r15,0x7f563412(%rsp)
	xchg	%r15,0x7f563412(%rbp)
	xchg	%r15,0x7f563412(%rsi)
	xchg	%r15,0x7f563412(%rdi)
	xchg	%r15,0x7f563412(%r8)
	xchg	%r15,0x7f563412(%r9)
	xchg	%r15,0x7f563412(%r10)
	xchg	%r15,0x7f563412(%r11)
	xchg	%r15,0x7f563412(%r12)
	xchg	%r15,0x7f563412(%r13)
	xchg	%r15,0x7f563412(%r14)
	xchg	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	XchgRegMem
	.type	XchgRegMem, @function
XchgRegMem:
	.cfi_startproc
	xchg	(%rax),%rax
	xchg	(%rax),%rcx
	xchg	(%rax),%rdx
	xchg	(%rax),%rbx
	xchg	(%rax),%rsp
	xchg	(%rax),%rbp
	xchg	(%rax),%rsi
	xchg	(%rax),%rdi
	xchg	(%rax),%r8
	xchg	(%rax),%r9
	xchg	(%rax),%r10
	xchg	(%rax),%r11
	xchg	(%rax),%r12
	xchg	(%rax),%r13
	xchg	(%rax),%r14
	xchg	(%rax),%r15
	nop
	xchg	(%rcx),%rax
	xchg	(%rcx),%rcx
	xchg	(%rcx),%rdx
	xchg	(%rcx),%rbx
	xchg	(%rcx),%rsp
	xchg	(%rcx),%rbp
	xchg	(%rcx),%rsi
	xchg	(%rcx),%rdi
	xchg	(%rcx),%r8
	xchg	(%rcx),%r9
	xchg	(%rcx),%r10
	xchg	(%rcx),%r11
	xchg	(%rcx),%r12
	xchg	(%rcx),%r13
	xchg	(%rcx),%r14
	xchg	(%rcx),%r15
	nop
	xchg	(%rdx),%rax
	xchg	(%rdx),%rcx
	xchg	(%rdx),%rdx
	xchg	(%rdx),%rbx
	xchg	(%rdx),%rsp
	xchg	(%rdx),%rbp
	xchg	(%rdx),%rsi
	xchg	(%rdx),%rdi
	xchg	(%rdx),%r8
	xchg	(%rdx),%r9
	xchg	(%rdx),%r10
	xchg	(%rdx),%r11
	xchg	(%rdx),%r12
	xchg	(%rdx),%r13
	xchg	(%rdx),%r14
	xchg	(%rdx),%r15
	nop
	xchg	(%rbx),%rax
	xchg	(%rbx),%rcx
	xchg	(%rbx),%rdx
	xchg	(%rbx),%rbx
	xchg	(%rbx),%rsp
	xchg	(%rbx),%rbp
	xchg	(%rbx),%rsi
	xchg	(%rbx),%rdi
	xchg	(%rbx),%r8
	xchg	(%rbx),%r9
	xchg	(%rbx),%r10
	xchg	(%rbx),%r11
	xchg	(%rbx),%r12
	xchg	(%rbx),%r13
	xchg	(%rbx),%r14
	xchg	(%rbx),%r15
	nop
	xchg	(%rsp),%rax
	xchg	(%rsp),%rcx
	xchg	(%rsp),%rdx
	xchg	(%rsp),%rbx
	xchg	(%rsp),%rsp
	xchg	(%rsp),%rbp
	xchg	(%rsp),%rsi
	xchg	(%rsp),%rdi
	xchg	(%rsp),%r8
	xchg	(%rsp),%r9
	xchg	(%rsp),%r10
	xchg	(%rsp),%r11
	xchg	(%rsp),%r12
	xchg	(%rsp),%r13
	xchg	(%rsp),%r14
	xchg	(%rsp),%r15
	nop
	xchg	(%rbp),%rax
	xchg	(%rbp),%rcx
	xchg	(%rbp),%rdx
	xchg	(%rbp),%rbx
	xchg	(%rbp),%rsp
	xchg	(%rbp),%rbp
	xchg	(%rbp),%rsi
	xchg	(%rbp),%rdi
	xchg	(%rbp),%r8
	xchg	(%rbp),%r9
	xchg	(%rbp),%r10
	xchg	(%rbp),%r11
	xchg	(%rbp),%r12
	xchg	(%rbp),%r13
	xchg	(%rbp),%r14
	xchg	(%rbp),%r15
	nop
	xchg	(%rsi),%rax
	xchg	(%rsi),%rcx
	xchg	(%rsi),%rdx
	xchg	(%rsi),%rbx
	xchg	(%rsi),%rsp
	xchg	(%rsi),%rbp
	xchg	(%rsi),%rsi
	xchg	(%rsi),%rdi
	xchg	(%rsi),%r8
	xchg	(%rsi),%r9
	xchg	(%rsi),%r10
	xchg	(%rsi),%r11
	xchg	(%rsi),%r12
	xchg	(%rsi),%r13
	xchg	(%rsi),%r14
	xchg	(%rsi),%r15
	nop
	xchg	(%rdi),%rax
	xchg	(%rdi),%rcx
	xchg	(%rdi),%rdx
	xchg	(%rdi),%rbx
	xchg	(%rdi),%rsp
	xchg	(%rdi),%rbp
	xchg	(%rdi),%rsi
	xchg	(%rdi),%rdi
	xchg	(%rdi),%r8
	xchg	(%rdi),%r9
	xchg	(%rdi),%r10
	xchg	(%rdi),%r11
	xchg	(%rdi),%r12
	xchg	(%rdi),%r13
	xchg	(%rdi),%r14
	xchg	(%rdi),%r15
	nop
	xchg	(%r8), %rax
	xchg	(%r8), %rcx
	xchg	(%r8), %rdx
	xchg	(%r8), %rbx
	xchg	(%r8), %rsp
	xchg	(%r8), %rbp
	xchg	(%r8), %rsi
	xchg	(%r8), %rdi
	xchg	(%r8), %r8
	xchg	(%r8), %r9
	xchg	(%r8), %r10
	xchg	(%r8), %r11
	xchg	(%r8), %r12
	xchg	(%r8), %r13
	xchg	(%r8), %r14
	xchg	(%r8), %r15
	nop
	xchg	(%r9), %rax
	xchg	(%r9), %rcx
	xchg	(%r9), %rdx
	xchg	(%r9), %rbx
	xchg	(%r9), %rsp
	xchg	(%r9), %rbp
	xchg	(%r9), %rsi
	xchg	(%r9), %rdi
	xchg	(%r9), %r8
	xchg	(%r9), %r9
	xchg	(%r9), %r10
	xchg	(%r9), %r11
	xchg	(%r9), %r12
	xchg	(%r9), %r13
	xchg	(%r9), %r14
	xchg	(%r9), %r15
	nop
	xchg	(%r10),%rax
	xchg	(%r10),%rcx
	xchg	(%r10),%rdx
	xchg	(%r10),%rbx
	xchg	(%r10),%rsp
	xchg	(%r10),%rbp
	xchg	(%r10),%rsi
	xchg	(%r10),%rdi
	xchg	(%r10),%r8
	xchg	(%r10),%r9
	xchg	(%r10),%r10
	xchg	(%r10),%r11
	xchg	(%r10),%r12
	xchg	(%r10),%r13
	xchg	(%r10),%r14
	xchg	(%r10),%r15
	nop
	xchg	(%r11),%rax
	xchg	(%r11),%rcx
	xchg	(%r11),%rdx
	xchg	(%r11),%rbx
	xchg	(%r11),%rsp
	xchg	(%r11),%rbp
	xchg	(%r11),%rsi
	xchg	(%r11),%rdi
	xchg	(%r11),%r8
	xchg	(%r11),%r9
	xchg	(%r11),%r10
	xchg	(%r11),%r11
	xchg	(%r11),%r12
	xchg	(%r11),%r13
	xchg	(%r11),%r14
	xchg	(%r11),%r15
	nop
	xchg	(%r12),%rax
	xchg	(%r12),%rcx
	xchg	(%r12),%rdx
	xchg	(%r12),%rbx
	xchg	(%r12),%rsp
	xchg	(%r12),%rbp
	xchg	(%r12),%rsi
	xchg	(%r12),%rdi
	xchg	(%r12),%r8
	xchg	(%r12),%r9
	xchg	(%r12),%r10
	xchg	(%r12),%r11
	xchg	(%r12),%r12
	xchg	(%r12),%r13
	xchg	(%r12),%r14
	xchg	(%r12),%r15
	nop
	xchg	(%r13),%rax
	xchg	(%r13),%rcx
	xchg	(%r13),%rdx
	xchg	(%r13),%rbx
	xchg	(%r13),%rsp
	xchg	(%r13),%rbp
	xchg	(%r13),%rsi
	xchg	(%r13),%rdi
	xchg	(%r13),%r8
	xchg	(%r13),%r9
	xchg	(%r13),%r10
	xchg	(%r13),%r11
	xchg	(%r13),%r12
	xchg	(%r13),%r13
	xchg	(%r13),%r14
	xchg	(%r13),%r15
	nop
	xchg	(%r14),%rax
	xchg	(%r14),%rcx
	xchg	(%r14),%rdx
	xchg	(%r14),%rbx
	xchg	(%r14),%rsp
	xchg	(%r14),%rbp
	xchg	(%r14),%rsi
	xchg	(%r14),%rdi
	xchg	(%r14),%r8
	xchg	(%r14),%r9
	xchg	(%r14),%r10
	xchg	(%r14),%r11
	xchg	(%r14),%r12
	xchg	(%r14),%r13
	xchg	(%r14),%r14
	xchg	(%r14),%r15
	nop
	xchg	(%r15),%rax
	xchg	(%r15),%rcx
	xchg	(%r15),%rdx
	xchg	(%r15),%rbx
	xchg	(%r15),%rsp
	xchg	(%r15),%rbp
	xchg	(%r15),%rsi
	xchg	(%r15),%rdi
	xchg	(%r15),%r8
	xchg	(%r15),%r9
	xchg	(%r15),%r10
	xchg	(%r15),%r11
	xchg	(%r15),%r12
	xchg	(%r15),%r13
	xchg	(%r15),%r14
	xchg	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XchgRegMem8
	.type	XchgRegMem8, @function
XchgRegMem8:
	.cfi_startproc
	xchg	0x7f(%rax),%rax
	xchg	0x7f(%rax),%rcx
	xchg	0x7f(%rax),%rdx
	xchg	0x7f(%rax),%rbx
	xchg	0x7f(%rax),%rsp
	xchg	0x7f(%rax),%rbp
	xchg	0x7f(%rax),%rsi
	xchg	0x7f(%rax),%rdi
	xchg	0x7f(%rax),%r8
	xchg	0x7f(%rax),%r9
	xchg	0x7f(%rax),%r10
	xchg	0x7f(%rax),%r11
	xchg	0x7f(%rax),%r12
	xchg	0x7f(%rax),%r13
	xchg	0x7f(%rax),%r14
	xchg	0x7f(%rax),%r15
	nop
	xchg	0x7f(%rcx),%rax
	xchg	0x7f(%rcx),%rcx
	xchg	0x7f(%rcx),%rdx
	xchg	0x7f(%rcx),%rbx
	xchg	0x7f(%rcx),%rsp
	xchg	0x7f(%rcx),%rbp
	xchg	0x7f(%rcx),%rsi
	xchg	0x7f(%rcx),%rdi
	xchg	0x7f(%rcx),%r8
	xchg	0x7f(%rcx),%r9
	xchg	0x7f(%rcx),%r10
	xchg	0x7f(%rcx),%r11
	xchg	0x7f(%rcx),%r12
	xchg	0x7f(%rcx),%r13
	xchg	0x7f(%rcx),%r14
	xchg	0x7f(%rcx),%r15
	nop
	xchg	0x7f(%rdx),%rax
	xchg	0x7f(%rdx),%rcx
	xchg	0x7f(%rdx),%rdx
	xchg	0x7f(%rdx),%rbx
	xchg	0x7f(%rdx),%rsp
	xchg	0x7f(%rdx),%rbp
	xchg	0x7f(%rdx),%rsi
	xchg	0x7f(%rdx),%rdi
	xchg	0x7f(%rdx),%r8
	xchg	0x7f(%rdx),%r9
	xchg	0x7f(%rdx),%r10
	xchg	0x7f(%rdx),%r11
	xchg	0x7f(%rdx),%r12
	xchg	0x7f(%rdx),%r13
	xchg	0x7f(%rdx),%r14
	xchg	0x7f(%rdx),%r15
	nop
	xchg	0x7f(%rbx),%rax
	xchg	0x7f(%rbx),%rcx
	xchg	0x7f(%rbx),%rdx
	xchg	0x7f(%rbx),%rbx
	xchg	0x7f(%rbx),%rsp
	xchg	0x7f(%rbx),%rbp
	xchg	0x7f(%rbx),%rsi
	xchg	0x7f(%rbx),%rdi
	xchg	0x7f(%rbx),%r8
	xchg	0x7f(%rbx),%r9
	xchg	0x7f(%rbx),%r10
	xchg	0x7f(%rbx),%r11
	xchg	0x7f(%rbx),%r12
	xchg	0x7f(%rbx),%r13
	xchg	0x7f(%rbx),%r14
	xchg	0x7f(%rbx),%r15
	nop
	xchg	0x7f(%rsp),%rax
	xchg	0x7f(%rsp),%rcx
	xchg	0x7f(%rsp),%rdx
	xchg	0x7f(%rsp),%rbx
	xchg	0x7f(%rsp),%rsp
	xchg	0x7f(%rsp),%rbp
	xchg	0x7f(%rsp),%rsi
	xchg	0x7f(%rsp),%rdi
	xchg	0x7f(%rsp),%r8
	xchg	0x7f(%rsp),%r9
	xchg	0x7f(%rsp),%r10
	xchg	0x7f(%rsp),%r11
	xchg	0x7f(%rsp),%r12
	xchg	0x7f(%rsp),%r13
	xchg	0x7f(%rsp),%r14
	xchg	0x7f(%rsp),%r15
	nop
	xchg	0x7f(%rbp),%rax
	xchg	0x7f(%rbp),%rcx
	xchg	0x7f(%rbp),%rdx
	xchg	0x7f(%rbp),%rbx
	xchg	0x7f(%rbp),%rsp
	xchg	0x7f(%rbp),%rbp
	xchg	0x7f(%rbp),%rsi
	xchg	0x7f(%rbp),%rdi
	xchg	0x7f(%rbp),%r8
	xchg	0x7f(%rbp),%r9
	xchg	0x7f(%rbp),%r10
	xchg	0x7f(%rbp),%r11
	xchg	0x7f(%rbp),%r12
	xchg	0x7f(%rbp),%r13
	xchg	0x7f(%rbp),%r14
	xchg	0x7f(%rbp),%r15
	nop
	xchg	0x7f(%rsi),%rax
	xchg	0x7f(%rsi),%rcx
	xchg	0x7f(%rsi),%rdx
	xchg	0x7f(%rsi),%rbx
	xchg	0x7f(%rsi),%rsp
	xchg	0x7f(%rsi),%rbp
	xchg	0x7f(%rsi),%rsi
	xchg	0x7f(%rsi),%rdi
	xchg	0x7f(%rsi),%r8
	xchg	0x7f(%rsi),%r9
	xchg	0x7f(%rsi),%r10
	xchg	0x7f(%rsi),%r11
	xchg	0x7f(%rsi),%r12
	xchg	0x7f(%rsi),%r13
	xchg	0x7f(%rsi),%r14
	xchg	0x7f(%rsi),%r15
	nop
	xchg	0x7f(%rdi),%rax
	xchg	0x7f(%rdi),%rcx
	xchg	0x7f(%rdi),%rdx
	xchg	0x7f(%rdi),%rbx
	xchg	0x7f(%rdi),%rsp
	xchg	0x7f(%rdi),%rbp
	xchg	0x7f(%rdi),%rsi
	xchg	0x7f(%rdi),%rdi
	xchg	0x7f(%rdi),%r8
	xchg	0x7f(%rdi),%r9
	xchg	0x7f(%rdi),%r10
	xchg	0x7f(%rdi),%r11
	xchg	0x7f(%rdi),%r12
	xchg	0x7f(%rdi),%r13
	xchg	0x7f(%rdi),%r14
	xchg	0x7f(%rdi),%r15
	nop
	xchg	0x7f(%r8), %rax
	xchg	0x7f(%r8), %rcx
	xchg	0x7f(%r8), %rdx
	xchg	0x7f(%r8), %rbx
	xchg	0x7f(%r8), %rsp
	xchg	0x7f(%r8), %rbp
	xchg	0x7f(%r8), %rsi
	xchg	0x7f(%r8), %rdi
	xchg	0x7f(%r8), %r8
	xchg	0x7f(%r8), %r9
	xchg	0x7f(%r8), %r10
	xchg	0x7f(%r8), %r11
	xchg	0x7f(%r8), %r12
	xchg	0x7f(%r8), %r13
	xchg	0x7f(%r8), %r14
	xchg	0x7f(%r8), %r15
	nop
	xchg	0x7f(%r9), %rax
	xchg	0x7f(%r9), %rcx
	xchg	0x7f(%r9), %rdx
	xchg	0x7f(%r9), %rbx
	xchg	0x7f(%r9), %rsp
	xchg	0x7f(%r9), %rbp
	xchg	0x7f(%r9), %rsi
	xchg	0x7f(%r9), %rdi
	xchg	0x7f(%r9), %r8
	xchg	0x7f(%r9), %r9
	xchg	0x7f(%r9), %r10
	xchg	0x7f(%r9), %r11
	xchg	0x7f(%r9), %r12
	xchg	0x7f(%r9), %r13
	xchg	0x7f(%r9), %r14
	xchg	0x7f(%r9), %r15
	nop
	xchg	0x7f(%r10),%rax
	xchg	0x7f(%r10),%rcx
	xchg	0x7f(%r10),%rdx
	xchg	0x7f(%r10),%rbx
	xchg	0x7f(%r10),%rsp
	xchg	0x7f(%r10),%rbp
	xchg	0x7f(%r10),%rsi
	xchg	0x7f(%r10),%rdi
	xchg	0x7f(%r10),%r8
	xchg	0x7f(%r10),%r9
	xchg	0x7f(%r10),%r10
	xchg	0x7f(%r10),%r11
	xchg	0x7f(%r10),%r12
	xchg	0x7f(%r10),%r13
	xchg	0x7f(%r10),%r14
	xchg	0x7f(%r10),%r15
	nop
	xchg	0x7f(%r11),%rax
	xchg	0x7f(%r11),%rcx
	xchg	0x7f(%r11),%rdx
	xchg	0x7f(%r11),%rbx
	xchg	0x7f(%r11),%rsp
	xchg	0x7f(%r11),%rbp
	xchg	0x7f(%r11),%rsi
	xchg	0x7f(%r11),%rdi
	xchg	0x7f(%r11),%r8
	xchg	0x7f(%r11),%r9
	xchg	0x7f(%r11),%r10
	xchg	0x7f(%r11),%r11
	xchg	0x7f(%r11),%r12
	xchg	0x7f(%r11),%r13
	xchg	0x7f(%r11),%r14
	xchg	0x7f(%r11),%r15
	nop
	xchg	0x7f(%r12),%rax
	xchg	0x7f(%r12),%rcx
	xchg	0x7f(%r12),%rdx
	xchg	0x7f(%r12),%rbx
	xchg	0x7f(%r12),%rsp
	xchg	0x7f(%r12),%rbp
	xchg	0x7f(%r12),%rsi
	xchg	0x7f(%r12),%rdi
	xchg	0x7f(%r12),%r8
	xchg	0x7f(%r12),%r9
	xchg	0x7f(%r12),%r10
	xchg	0x7f(%r12),%r11
	xchg	0x7f(%r12),%r12
	xchg	0x7f(%r12),%r13
	xchg	0x7f(%r12),%r14
	xchg	0x7f(%r12),%r15
	nop
	xchg	0x7f(%r13),%rax
	xchg	0x7f(%r13),%rcx
	xchg	0x7f(%r13),%rdx
	xchg	0x7f(%r13),%rbx
	xchg	0x7f(%r13),%rsp
	xchg	0x7f(%r13),%rbp
	xchg	0x7f(%r13),%rsi
	xchg	0x7f(%r13),%rdi
	xchg	0x7f(%r13),%r8
	xchg	0x7f(%r13),%r9
	xchg	0x7f(%r13),%r10
	xchg	0x7f(%r13),%r11
	xchg	0x7f(%r13),%r12
	xchg	0x7f(%r13),%r13
	xchg	0x7f(%r13),%r14
	xchg	0x7f(%r13),%r15
	nop
	xchg	0x7f(%r14),%rax
	xchg	0x7f(%r14),%rcx
	xchg	0x7f(%r14),%rdx
	xchg	0x7f(%r14),%rbx
	xchg	0x7f(%r14),%rsp
	xchg	0x7f(%r14),%rbp
	xchg	0x7f(%r14),%rsi
	xchg	0x7f(%r14),%rdi
	xchg	0x7f(%r14),%r8
	xchg	0x7f(%r14),%r9
	xchg	0x7f(%r14),%r10
	xchg	0x7f(%r14),%r11
	xchg	0x7f(%r14),%r12
	xchg	0x7f(%r14),%r13
	xchg	0x7f(%r14),%r14
	xchg	0x7f(%r14),%r15
	nop
	xchg	0x7f(%r15),%rax
	xchg	0x7f(%r15),%rcx
	xchg	0x7f(%r15),%rdx
	xchg	0x7f(%r15),%rbx
	xchg	0x7f(%r15),%rsp
	xchg	0x7f(%r15),%rbp
	xchg	0x7f(%r15),%rsi
	xchg	0x7f(%r15),%rdi
	xchg	0x7f(%r15),%r8
	xchg	0x7f(%r15),%r9
	xchg	0x7f(%r15),%r10
	xchg	0x7f(%r15),%r11
	xchg	0x7f(%r15),%r12
	xchg	0x7f(%r15),%r13
	xchg	0x7f(%r15),%r14
	xchg	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XchgRegMem32
	.type	XchgRegMem32, @function
XchgRegMem32:
	.cfi_startproc
	xchg	0x7f563412(%rax),%rax
	xchg	0x7f563412(%rax),%rcx
	xchg	0x7f563412(%rax),%rdx
	xchg	0x7f563412(%rax),%rbx
	xchg	0x7f563412(%rax),%rsp
	xchg	0x7f563412(%rax),%rbp
	xchg	0x7f563412(%rax),%rsi
	xchg	0x7f563412(%rax),%rdi
	xchg	0x7f563412(%rax),%r8
	xchg	0x7f563412(%rax),%r9
	xchg	0x7f563412(%rax),%r10
	xchg	0x7f563412(%rax),%r11
	xchg	0x7f563412(%rax),%r12
	xchg	0x7f563412(%rax),%r13
	xchg	0x7f563412(%rax),%r14
	xchg	0x7f563412(%rax),%r15
	nop
	xchg	0x7f563412(%rcx),%rax
	xchg	0x7f563412(%rcx),%rcx
	xchg	0x7f563412(%rcx),%rdx
	xchg	0x7f563412(%rcx),%rbx
	xchg	0x7f563412(%rcx),%rsp
	xchg	0x7f563412(%rcx),%rbp
	xchg	0x7f563412(%rcx),%rsi
	xchg	0x7f563412(%rcx),%rdi
	xchg	0x7f563412(%rcx),%r8
	xchg	0x7f563412(%rcx),%r9
	xchg	0x7f563412(%rcx),%r10
	xchg	0x7f563412(%rcx),%r11
	xchg	0x7f563412(%rcx),%r12
	xchg	0x7f563412(%rcx),%r13
	xchg	0x7f563412(%rcx),%r14
	xchg	0x7f563412(%rcx),%r15
	nop
	xchg	0x7f563412(%rdx),%rax
	xchg	0x7f563412(%rdx),%rcx
	xchg	0x7f563412(%rdx),%rdx
	xchg	0x7f563412(%rdx),%rbx
	xchg	0x7f563412(%rdx),%rsp
	xchg	0x7f563412(%rdx),%rbp
	xchg	0x7f563412(%rdx),%rsi
	xchg	0x7f563412(%rdx),%rdi
	xchg	0x7f563412(%rdx),%r8
	xchg	0x7f563412(%rdx),%r9
	xchg	0x7f563412(%rdx),%r10
	xchg	0x7f563412(%rdx),%r11
	xchg	0x7f563412(%rdx),%r12
	xchg	0x7f563412(%rdx),%r13
	xchg	0x7f563412(%rdx),%r14
	xchg	0x7f563412(%rdx),%r15
	nop
	xchg	0x7f563412(%rbx),%rax
	xchg	0x7f563412(%rbx),%rcx
	xchg	0x7f563412(%rbx),%rdx
	xchg	0x7f563412(%rbx),%rbx
	xchg	0x7f563412(%rbx),%rsp
	xchg	0x7f563412(%rbx),%rbp
	xchg	0x7f563412(%rbx),%rsi
	xchg	0x7f563412(%rbx),%rdi
	xchg	0x7f563412(%rbx),%r8
	xchg	0x7f563412(%rbx),%r9
	xchg	0x7f563412(%rbx),%r10
	xchg	0x7f563412(%rbx),%r11
	xchg	0x7f563412(%rbx),%r12
	xchg	0x7f563412(%rbx),%r13
	xchg	0x7f563412(%rbx),%r14
	xchg	0x7f563412(%rbx),%r15
	nop
	xchg	0x7f563412(%rsp),%rax
	xchg	0x7f563412(%rsp),%rcx
	xchg	0x7f563412(%rsp),%rdx
	xchg	0x7f563412(%rsp),%rbx
	xchg	0x7f563412(%rsp),%rsp
	xchg	0x7f563412(%rsp),%rbp
	xchg	0x7f563412(%rsp),%rsi
	xchg	0x7f563412(%rsp),%rdi
	xchg	0x7f563412(%rsp),%r8
	xchg	0x7f563412(%rsp),%r9
	xchg	0x7f563412(%rsp),%r10
	xchg	0x7f563412(%rsp),%r11
	xchg	0x7f563412(%rsp),%r12
	xchg	0x7f563412(%rsp),%r13
	xchg	0x7f563412(%rsp),%r14
	xchg	0x7f563412(%rsp),%r15
	nop
	xchg	0x7f563412(%rbp),%rax
	xchg	0x7f563412(%rbp),%rcx
	xchg	0x7f563412(%rbp),%rdx
	xchg	0x7f563412(%rbp),%rbx
	xchg	0x7f563412(%rbp),%rsp
	xchg	0x7f563412(%rbp),%rbp
	xchg	0x7f563412(%rbp),%rsi
	xchg	0x7f563412(%rbp),%rdi
	xchg	0x7f563412(%rbp),%r8
	xchg	0x7f563412(%rbp),%r9
	xchg	0x7f563412(%rbp),%r10
	xchg	0x7f563412(%rbp),%r11
	xchg	0x7f563412(%rbp),%r12
	xchg	0x7f563412(%rbp),%r13
	xchg	0x7f563412(%rbp),%r14
	xchg	0x7f563412(%rbp),%r15
	nop
	xchg	0x7f563412(%rsi),%rax
	xchg	0x7f563412(%rsi),%rcx
	xchg	0x7f563412(%rsi),%rdx
	xchg	0x7f563412(%rsi),%rbx
	xchg	0x7f563412(%rsi),%rsp
	xchg	0x7f563412(%rsi),%rbp
	xchg	0x7f563412(%rsi),%rsi
	xchg	0x7f563412(%rsi),%rdi
	xchg	0x7f563412(%rsi),%r8
	xchg	0x7f563412(%rsi),%r9
	xchg	0x7f563412(%rsi),%r10
	xchg	0x7f563412(%rsi),%r11
	xchg	0x7f563412(%rsi),%r12
	xchg	0x7f563412(%rsi),%r13
	xchg	0x7f563412(%rsi),%r14
	xchg	0x7f563412(%rsi),%r15
	nop
	xchg	0x7f563412(%rdi),%rax
	xchg	0x7f563412(%rdi),%rcx
	xchg	0x7f563412(%rdi),%rdx
	xchg	0x7f563412(%rdi),%rbx
	xchg	0x7f563412(%rdi),%rsp
	xchg	0x7f563412(%rdi),%rbp
	xchg	0x7f563412(%rdi),%rsi
	xchg	0x7f563412(%rdi),%rdi
	xchg	0x7f563412(%rdi),%r8
	xchg	0x7f563412(%rdi),%r9
	xchg	0x7f563412(%rdi),%r10
	xchg	0x7f563412(%rdi),%r11
	xchg	0x7f563412(%rdi),%r12
	xchg	0x7f563412(%rdi),%r13
	xchg	0x7f563412(%rdi),%r14
	xchg	0x7f563412(%rdi),%r15
	nop
	xchg	0x7f563412(%r8), %rax
	xchg	0x7f563412(%r8), %rcx
	xchg	0x7f563412(%r8), %rdx
	xchg	0x7f563412(%r8), %rbx
	xchg	0x7f563412(%r8), %rsp
	xchg	0x7f563412(%r8), %rbp
	xchg	0x7f563412(%r8), %rsi
	xchg	0x7f563412(%r8), %rdi
	xchg	0x7f563412(%r8), %r8
	xchg	0x7f563412(%r8), %r9
	xchg	0x7f563412(%r8), %r10
	xchg	0x7f563412(%r8), %r11
	xchg	0x7f563412(%r8), %r12
	xchg	0x7f563412(%r8), %r13
	xchg	0x7f563412(%r8), %r14
	xchg	0x7f563412(%r8), %r15
	nop
	xchg	0x7f563412(%r9), %rax
	xchg	0x7f563412(%r9), %rcx
	xchg	0x7f563412(%r9), %rdx
	xchg	0x7f563412(%r9), %rbx
	xchg	0x7f563412(%r9), %rsp
	xchg	0x7f563412(%r9), %rbp
	xchg	0x7f563412(%r9), %rsi
	xchg	0x7f563412(%r9), %rdi
	xchg	0x7f563412(%r9), %r8
	xchg	0x7f563412(%r9), %r9
	xchg	0x7f563412(%r9), %r10
	xchg	0x7f563412(%r9), %r11
	xchg	0x7f563412(%r9), %r12
	xchg	0x7f563412(%r9), %r13
	xchg	0x7f563412(%r9), %r14
	xchg	0x7f563412(%r9), %r15
	nop
	xchg	0x7f563412(%r10),%rax
	xchg	0x7f563412(%r10),%rcx
	xchg	0x7f563412(%r10),%rdx
	xchg	0x7f563412(%r10),%rbx
	xchg	0x7f563412(%r10),%rsp
	xchg	0x7f563412(%r10),%rbp
	xchg	0x7f563412(%r10),%rsi
	xchg	0x7f563412(%r10),%rdi
	xchg	0x7f563412(%r10),%r8
	xchg	0x7f563412(%r10),%r9
	xchg	0x7f563412(%r10),%r10
	xchg	0x7f563412(%r10),%r11
	xchg	0x7f563412(%r10),%r12
	xchg	0x7f563412(%r10),%r13
	xchg	0x7f563412(%r10),%r14
	xchg	0x7f563412(%r10),%r15
	nop
	xchg	0x7f563412(%r11),%rax
	xchg	0x7f563412(%r11),%rcx
	xchg	0x7f563412(%r11),%rdx
	xchg	0x7f563412(%r11),%rbx
	xchg	0x7f563412(%r11),%rsp
	xchg	0x7f563412(%r11),%rbp
	xchg	0x7f563412(%r11),%rsi
	xchg	0x7f563412(%r11),%rdi
	xchg	0x7f563412(%r11),%r8
	xchg	0x7f563412(%r11),%r9
	xchg	0x7f563412(%r11),%r10
	xchg	0x7f563412(%r11),%r11
	xchg	0x7f563412(%r11),%r12
	xchg	0x7f563412(%r11),%r13
	xchg	0x7f563412(%r11),%r14
	xchg	0x7f563412(%r11),%r15
	nop
	xchg	0x7f563412(%r12),%rax
	xchg	0x7f563412(%r12),%rcx
	xchg	0x7f563412(%r12),%rdx
	xchg	0x7f563412(%r12),%rbx
	xchg	0x7f563412(%r12),%rsp
	xchg	0x7f563412(%r12),%rbp
	xchg	0x7f563412(%r12),%rsi
	xchg	0x7f563412(%r12),%rdi
	xchg	0x7f563412(%r12),%r8
	xchg	0x7f563412(%r12),%r9
	xchg	0x7f563412(%r12),%r10
	xchg	0x7f563412(%r12),%r11
	xchg	0x7f563412(%r12),%r12
	xchg	0x7f563412(%r12),%r13
	xchg	0x7f563412(%r12),%r14
	xchg	0x7f563412(%r12),%r15
	nop
	xchg	0x7f563412(%r13),%rax
	xchg	0x7f563412(%r13),%rcx
	xchg	0x7f563412(%r13),%rdx
	xchg	0x7f563412(%r13),%rbx
	xchg	0x7f563412(%r13),%rsp
	xchg	0x7f563412(%r13),%rbp
	xchg	0x7f563412(%r13),%rsi
	xchg	0x7f563412(%r13),%rdi
	xchg	0x7f563412(%r13),%r8
	xchg	0x7f563412(%r13),%r9
	xchg	0x7f563412(%r13),%r10
	xchg	0x7f563412(%r13),%r11
	xchg	0x7f563412(%r13),%r12
	xchg	0x7f563412(%r13),%r13
	xchg	0x7f563412(%r13),%r14
	xchg	0x7f563412(%r13),%r15
	nop
	xchg	0x7f563412(%r14),%rax
	xchg	0x7f563412(%r14),%rcx
	xchg	0x7f563412(%r14),%rdx
	xchg	0x7f563412(%r14),%rbx
	xchg	0x7f563412(%r14),%rsp
	xchg	0x7f563412(%r14),%rbp
	xchg	0x7f563412(%r14),%rsi
	xchg	0x7f563412(%r14),%rdi
	xchg	0x7f563412(%r14),%r8
	xchg	0x7f563412(%r14),%r9
	xchg	0x7f563412(%r14),%r10
	xchg	0x7f563412(%r14),%r11
	xchg	0x7f563412(%r14),%r12
	xchg	0x7f563412(%r14),%r13
	xchg	0x7f563412(%r14),%r14
	xchg	0x7f563412(%r14),%r15
	nop
	xchg	0x7f563412(%r15),%rax
	xchg	0x7f563412(%r15),%rcx
	xchg	0x7f563412(%r15),%rdx
	xchg	0x7f563412(%r15),%rbx
	xchg	0x7f563412(%r15),%rsp
	xchg	0x7f563412(%r15),%rbp
	xchg	0x7f563412(%r15),%rsi
	xchg	0x7f563412(%r15),%rdi
	xchg	0x7f563412(%r15),%r8
	xchg	0x7f563412(%r15),%r9
	xchg	0x7f563412(%r15),%r10
	xchg	0x7f563412(%r15),%r11
	xchg	0x7f563412(%r15),%r12
	xchg	0x7f563412(%r15),%r13
	xchg	0x7f563412(%r15),%r14
	xchg	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


