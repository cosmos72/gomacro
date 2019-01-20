	.file	"add.s"
	.text

	.p2align 4,,15
	.globl	Add_s32
	.type	Add_s32, @function
Add_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// add	$-0x55667788,%rax
	add	$-0x55667788,%rcx
	add	$-0x55667788,%rdx
	add	$-0x55667788,%rbx
	add	$-0x55667788,%rsp
	add	$-0x55667788,%rbp
	add	$-0x55667788,%rsi
	add	$-0x55667788,%rdi
	add	$-0x55667788,%r8
	add	$-0x55667788,%r9
	add	$-0x55667788,%r10
	add	$-0x55667788,%r11
	add	$-0x55667788,%r12
	add	$-0x55667788,%r13
	add	$-0x55667788,%r14
	add	$-0x55667788,%r15
        ret
	.cfi_endproc



        .p2align 4,,15
	.globl	Misc
	.type	Misc, @function
Misc:
	.cfi_startproc
        .byte 0x48, 0x01, 0xc0
        .byte 0x48, 0x09, 0xc0
        .byte 0x48, 0x11, 0xc0
        .byte 0x48, 0x19, 0xc0
        .byte 0x48, 0x21, 0xc0
        .byte 0x48, 0x29, 0xc0
        .byte 0x48, 0x31, 0xc0
        .byte 0x48, 0x39, 0xc0
	rol	%cl,%rax
	ror	%cl,%rax
	shl     %cl,%rax
	shr	%cl,%rax
	sar	%cl,%rax
	xchg	%rax,%rcx
        lea     (%rax),%rax
        inc     %rax
        dec     %rax
        push    %rax
        pop     %rax
        movzbw  %al,%cx
        movzbl  %al,%rcx
        movzwl  %ax,%rcx
        ret
	.cfi_endproc



        .p2align 4,,15
	.globl	Add
	.type	Add, @function
Add:
	.cfi_startproc
        
	add	%rax,%rax
	add	%rax,%rcx
	add	%rax,%rdx
	add	%rax,%rbx
	add	%rax,%rsp
	add	%rax,%rbp
	add	%rax,%rsi
	add	%rax,%rdi
	add	%rax,%r8
	add	%rax,%r9
	add	%rax,%r10
	add	%rax,%r11
	add	%rax,%r12
	add	%rax,%r13
	add	%rax,%r14
	add	%rax,%r15
	nop
	add	%rcx,%rax
	add	%rcx,%rcx
	add	%rcx,%rdx
	add	%rcx,%rbx
	add	%rcx,%rsp
	add	%rcx,%rbp
	add	%rcx,%rsi
	add	%rcx,%rdi
	add	%rcx,%r8
	add	%rcx,%r9
	add	%rcx,%r10
	add	%rcx,%r11
	add	%rcx,%r12
	add	%rcx,%r13
	add	%rcx,%r14
	add	%rcx,%r15
	nop
	add	%rdx,%rax
	add	%rdx,%rcx
	add	%rdx,%rdx
	add	%rdx,%rbx
	add	%rdx,%rsp
	add	%rdx,%rbp
	add	%rdx,%rsi
	add	%rdx,%rdi
	add	%rdx,%r8
	add	%rdx,%r9
	add	%rdx,%r10
	add	%rdx,%r11
	add	%rdx,%r12
	add	%rdx,%r13
	add	%rdx,%r14
	add	%rdx,%r15
	nop
	add	%rbx,%rax
	add	%rbx,%rcx
	add	%rbx,%rdx
	add	%rbx,%rbx
	add	%rbx,%rsp
	add	%rbx,%rbp
	add	%rbx,%rsi
	add	%rbx,%rdi
	add	%rbx,%r8
	add	%rbx,%r9
	add	%rbx,%r10
	add	%rbx,%r11
	add	%rbx,%r12
	add	%rbx,%r13
	add	%rbx,%r14
	add	%rbx,%r15
	nop
	add	%rsp,%rax
	add	%rsp,%rcx
	add	%rsp,%rdx
	add	%rsp,%rbx
	add	%rsp,%rsp
	add	%rsp,%rbp
	add	%rsp,%rsi
	add	%rsp,%rdi
	add	%rsp,%r8
	add	%rsp,%r9
	add	%rsp,%r10
	add	%rsp,%r11
	add	%rsp,%r12
	add	%rsp,%r13
	add	%rsp,%r14
	add	%rsp,%r15
	nop
	add	%rbp,%rax
	add	%rbp,%rcx
	add	%rbp,%rdx
	add	%rbp,%rbx
	add	%rbp,%rsp
	add	%rbp,%rbp
	add	%rbp,%rsi
	add	%rbp,%rdi
	add	%rbp,%r8
	add	%rbp,%r9
	add	%rbp,%r10
	add	%rbp,%r11
	add	%rbp,%r12
	add	%rbp,%r13
	add	%rbp,%r14
	add	%rbp,%r15
	nop
	add	%rsi,%rax
	add	%rsi,%rcx
	add	%rsi,%rdx
	add	%rsi,%rbx
	add	%rsi,%rsp
	add	%rsi,%rbp
	add	%rsi,%rsi
	add	%rsi,%rdi
	add	%rsi,%r8
	add	%rsi,%r9
	add	%rsi,%r10
	add	%rsi,%r11
	add	%rsi,%r12
	add	%rsi,%r13
	add	%rsi,%r14
	add	%rsi,%r15
	nop
	add	%rdi,%rax
	add	%rdi,%rcx
	add	%rdi,%rdx
	add	%rdi,%rbx
	add	%rdi,%rsp
	add	%rdi,%rbp
	add	%rdi,%rsi
	add	%rdi,%rdi
	add	%rdi,%r8
	add	%rdi,%r9
	add	%rdi,%r10
	add	%rdi,%r11
	add	%rdi,%r12
	add	%rdi,%r13
	add	%rdi,%r14
	add	%rdi,%r15
	nop
	add	%r8, %rax
	add	%r8, %rcx
	add	%r8, %rdx
	add	%r8, %rbx
	add	%r8, %rsp
	add	%r8, %rbp
	add	%r8, %rsi
	add	%r8, %rdi
	add	%r8, %r8
	add	%r8, %r9
	add	%r8, %r10
	add	%r8, %r11
	add	%r8, %r12
	add	%r8, %r13
	add	%r8, %r14
	add	%r8, %r15
	nop
	add	%r9, %rax
	add	%r9, %rcx
	add	%r9, %rdx
	add	%r9, %rbx
	add	%r9, %rsp
	add	%r9, %rbp
	add	%r9, %rsi
	add	%r9, %rdi
	add	%r9, %r8
	add	%r9, %r9
	add	%r9, %r10
	add	%r9, %r11
	add	%r9, %r12
	add	%r9, %r13
	add	%r9, %r14
	add	%r9, %r15
	nop
	add	%r10,%rax
	add	%r10,%rcx
	add	%r10,%rdx
	add	%r10,%rbx
	add	%r10,%rsp
	add	%r10,%rbp
	add	%r10,%rsi
	add	%r10,%rdi
	add	%r10,%r8
	add	%r10,%r9
	add	%r10,%r10
	add	%r10,%r11
	add	%r10,%r12
	add	%r10,%r13
	add	%r10,%r14
	add	%r10,%r15
	nop
	add	%r11,%rax
	add	%r11,%rcx
	add	%r11,%rdx
	add	%r11,%rbx
	add	%r11,%rsp
	add	%r11,%rbp
	add	%r11,%rsi
	add	%r11,%rdi
	add	%r11,%r8
	add	%r11,%r9
	add	%r11,%r10
	add	%r11,%r11
	add	%r11,%r12
	add	%r11,%r13
	add	%r11,%r14
	add	%r11,%r15
	nop
	add	%r12,%rax
	add	%r12,%rcx
	add	%r12,%rdx
	add	%r12,%rbx
	add	%r12,%rsp
	add	%r12,%rbp
	add	%r12,%rsi
	add	%r12,%rdi
	add	%r12,%r8
	add	%r12,%r9
	add	%r12,%r10
	add	%r12,%r11
	add	%r12,%r12
	add	%r12,%r13
	add	%r12,%r14
	add	%r12,%r15
	nop
	add	%r13,%rax
	add	%r13,%rcx
	add	%r13,%rdx
	add	%r13,%rbx
	add	%r13,%rsp
	add	%r13,%rbp
	add	%r13,%rsi
	add	%r13,%rdi
	add	%r13,%r8
	add	%r13,%r9
	add	%r13,%r10
	add	%r13,%r11
	add	%r13,%r12
	add	%r13,%r13
	add	%r13,%r14
	add	%r13,%r15
	nop
	add	%r14,%rax
	add	%r14,%rcx
	add	%r14,%rdx
	add	%r14,%rbx
	add	%r14,%rsp
	add	%r14,%rbp
	add	%r14,%rsi
	add	%r14,%rdi
	add	%r14,%r8
	add	%r14,%r9
	add	%r14,%r10
	add	%r14,%r11
	add	%r14,%r12
	add	%r14,%r13
	add	%r14,%r14
	add	%r14,%r15
	nop
	add	%r15,%rax
	add	%r15,%rcx
	add	%r15,%rdx
	add	%r15,%rbx
	add	%r15,%rsp
	add	%r15,%rbp
	add	%r15,%rsi
	add	%r15,%rdi
	add	%r15,%r8
	add	%r15,%r9
	add	%r15,%r10
	add	%r15,%r11
	add	%r15,%r12
	add	%r15,%r13
	add	%r15,%r14
	add	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AddMemReg
	.type	AddMemReg, @function
AddMemReg:
	.cfi_startproc
	add	%rax,(%rax)
	add	%rax,(%rcx)
	add	%rax,(%rdx)
	add	%rax,(%rbx)
	add	%rax,(%rsp)
	add	%rax,(%rbp)
	add	%rax,(%rsi)
	add	%rax,(%rdi)
	add	%rax,(%r8)
	add	%rax,(%r9)
	add	%rax,(%r10)
	add	%rax,(%r11)
	add	%rax,(%r12)
	add	%rax,(%r13)
	add	%rax,(%r14)
	add	%rax,(%r15)
	nop
	add	%rcx,(%rax)
	add	%rcx,(%rcx)
	add	%rcx,(%rdx)
	add	%rcx,(%rbx)
	add	%rcx,(%rsp)
	add	%rcx,(%rbp)
	add	%rcx,(%rsi)
	add	%rcx,(%rdi)
	add	%rcx,(%r8)
	add	%rcx,(%r9)
	add	%rcx,(%r10)
	add	%rcx,(%r11)
	add	%rcx,(%r12)
	add	%rcx,(%r13)
	add	%rcx,(%r14)
	add	%rcx,(%r15)
	nop
	add	%rdx,(%rax)
	add	%rdx,(%rcx)
	add	%rdx,(%rdx)
	add	%rdx,(%rbx)
	add	%rdx,(%rsp)
	add	%rdx,(%rbp)
	add	%rdx,(%rsi)
	add	%rdx,(%rdi)
	add	%rdx,(%r8)
	add	%rdx,(%r9)
	add	%rdx,(%r10)
	add	%rdx,(%r11)
	add	%rdx,(%r12)
	add	%rdx,(%r13)
	add	%rdx,(%r14)
	add	%rdx,(%r15)
	nop
	add	%rbx,(%rax)
	add	%rbx,(%rcx)
	add	%rbx,(%rdx)
	add	%rbx,(%rbx)
	add	%rbx,(%rsp)
	add	%rbx,(%rbp)
	add	%rbx,(%rsi)
	add	%rbx,(%rdi)
	add	%rbx,(%r8)
	add	%rbx,(%r9)
	add	%rbx,(%r10)
	add	%rbx,(%r11)
	add	%rbx,(%r12)
	add	%rbx,(%r13)
	add	%rbx,(%r14)
	add	%rbx,(%r15)
	nop
	add	%rsp,(%rax)
	add	%rsp,(%rcx)
	add	%rsp,(%rdx)
	add	%rsp,(%rbx)
	add	%rsp,(%rsp)
	add	%rsp,(%rbp)
	add	%rsp,(%rsi)
	add	%rsp,(%rdi)
	add	%rsp,(%r8)
	add	%rsp,(%r9)
	add	%rsp,(%r10)
	add	%rsp,(%r11)
	add	%rsp,(%r12)
	add	%rsp,(%r13)
	add	%rsp,(%r14)
	add	%rsp,(%r15)
	nop
	add	%rbp,(%rax)
	add	%rbp,(%rcx)
	add	%rbp,(%rdx)
	add	%rbp,(%rbx)
	add	%rbp,(%rsp)
	add	%rbp,(%rbp)
	add	%rbp,(%rsi)
	add	%rbp,(%rdi)
	add	%rbp,(%r8)
	add	%rbp,(%r9)
	add	%rbp,(%r10)
	add	%rbp,(%r11)
	add	%rbp,(%r12)
	add	%rbp,(%r13)
	add	%rbp,(%r14)
	add	%rbp,(%r15)
	nop
	add	%rsi,(%rax)
	add	%rsi,(%rcx)
	add	%rsi,(%rdx)
	add	%rsi,(%rbx)
	add	%rsi,(%rsp)
	add	%rsi,(%rbp)
	add	%rsi,(%rsi)
	add	%rsi,(%rdi)
	add	%rsi,(%r8)
	add	%rsi,(%r9)
	add	%rsi,(%r10)
	add	%rsi,(%r11)
	add	%rsi,(%r12)
	add	%rsi,(%r13)
	add	%rsi,(%r14)
	add	%rsi,(%r15)
	nop
	add	%rdi,(%rax)
	add	%rdi,(%rcx)
	add	%rdi,(%rdx)
	add	%rdi,(%rbx)
	add	%rdi,(%rsp)
	add	%rdi,(%rbp)
	add	%rdi,(%rsi)
	add	%rdi,(%rdi)
	add	%rdi,(%r8)
	add	%rdi,(%r9)
	add	%rdi,(%r10)
	add	%rdi,(%r11)
	add	%rdi,(%r12)
	add	%rdi,(%r13)
	add	%rdi,(%r14)
	add	%rdi,(%r15)
	nop
	add	%r8, (%rax)
	add	%r8, (%rcx)
	add	%r8, (%rdx)
	add	%r8, (%rbx)
	add	%r8, (%rsp)
	add	%r8, (%rbp)
	add	%r8, (%rsi)
	add	%r8, (%rdi)
	add	%r8, (%r8)
	add	%r8, (%r9)
	add	%r8, (%r10)
	add	%r8, (%r11)
	add	%r8, (%r12)
	add	%r8, (%r13)
	add	%r8, (%r14)
	add	%r8, (%r15)
	nop
	add	%r9, (%rax)
	add	%r9, (%rcx)
	add	%r9, (%rdx)
	add	%r9, (%rbx)
	add	%r9, (%rsp)
	add	%r9, (%rbp)
	add	%r9, (%rsi)
	add	%r9, (%rdi)
	add	%r9, (%r8)
	add	%r9, (%r9)
	add	%r9, (%r10)
	add	%r9, (%r11)
	add	%r9, (%r12)
	add	%r9, (%r13)
	add	%r9, (%r14)
	add	%r9, (%r15)
	nop
	add	%r10,(%rax)
	add	%r10,(%rcx)
	add	%r10,(%rdx)
	add	%r10,(%rbx)
	add	%r10,(%rsp)
	add	%r10,(%rbp)
	add	%r10,(%rsi)
	add	%r10,(%rdi)
	add	%r10,(%r8)
	add	%r10,(%r9)
	add	%r10,(%r10)
	add	%r10,(%r11)
	add	%r10,(%r12)
	add	%r10,(%r13)
	add	%r10,(%r14)
	add	%r10,(%r15)
	nop
	add	%r11,(%rax)
	add	%r11,(%rcx)
	add	%r11,(%rdx)
	add	%r11,(%rbx)
	add	%r11,(%rsp)
	add	%r11,(%rbp)
	add	%r11,(%rsi)
	add	%r11,(%rdi)
	add	%r11,(%r8)
	add	%r11,(%r9)
	add	%r11,(%r10)
	add	%r11,(%r11)
	add	%r11,(%r12)
	add	%r11,(%r13)
	add	%r11,(%r14)
	add	%r11,(%r15)
	nop
	add	%r12,(%rax)
	add	%r12,(%rcx)
	add	%r12,(%rdx)
	add	%r12,(%rbx)
	add	%r12,(%rsp)
	add	%r12,(%rbp)
	add	%r12,(%rsi)
	add	%r12,(%rdi)
	add	%r12,(%r8)
	add	%r12,(%r9)
	add	%r12,(%r10)
	add	%r12,(%r11)
	add	%r12,(%r12)
	add	%r12,(%r13)
	add	%r12,(%r14)
	add	%r12,(%r15)
	nop
	add	%r13,(%rax)
	add	%r13,(%rcx)
	add	%r13,(%rdx)
	add	%r13,(%rbx)
	add	%r13,(%rsp)
	add	%r13,(%rbp)
	add	%r13,(%rsi)
	add	%r13,(%rdi)
	add	%r13,(%r8)
	add	%r13,(%r9)
	add	%r13,(%r10)
	add	%r13,(%r11)
	add	%r13,(%r12)
	add	%r13,(%r13)
	add	%r13,(%r14)
	add	%r13,(%r15)
	nop
	add	%r14,(%rax)
	add	%r14,(%rcx)
	add	%r14,(%rdx)
	add	%r14,(%rbx)
	add	%r14,(%rsp)
	add	%r14,(%rbp)
	add	%r14,(%rsi)
	add	%r14,(%rdi)
	add	%r14,(%r8)
	add	%r14,(%r9)
	add	%r14,(%r10)
	add	%r14,(%r11)
	add	%r14,(%r12)
	add	%r14,(%r13)
	add	%r14,(%r14)
	add	%r14,(%r15)
	nop
	add	%r15,(%rax)
	add	%r15,(%rcx)
	add	%r15,(%rdx)
	add	%r15,(%rbx)
	add	%r15,(%rsp)
	add	%r15,(%rbp)
	add	%r15,(%rsi)
	add	%r15,(%rdi)
	add	%r15,(%r8)
	add	%r15,(%r9)
	add	%r15,(%r10)
	add	%r15,(%r11)
	add	%r15,(%r12)
	add	%r15,(%r13)
	add	%r15,(%r14)
	add	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AddMem8Reg
	.type	AddMem8Reg, @function
AddMem8Reg:
	.cfi_startproc
	add	%rax,0x7f(%rax)
	add	%rax,0x7f(%rcx)
	add	%rax,0x7f(%rdx)
	add	%rax,0x7f(%rbx)
	add	%rax,0x7f(%rsp)
	add	%rax,0x7f(%rbp)
	add	%rax,0x7f(%rsi)
	add	%rax,0x7f(%rdi)
	add	%rax,0x7f(%r8)
	add	%rax,0x7f(%r9)
	add	%rax,0x7f(%r10)
	add	%rax,0x7f(%r11)
	add	%rax,0x7f(%r12)
	add	%rax,0x7f(%r13)
	add	%rax,0x7f(%r14)
	add	%rax,0x7f(%r15)
	nop
	add	%rcx,0x7f(%rax)
	add	%rcx,0x7f(%rcx)
	add	%rcx,0x7f(%rdx)
	add	%rcx,0x7f(%rbx)
	add	%rcx,0x7f(%rsp)
	add	%rcx,0x7f(%rbp)
	add	%rcx,0x7f(%rsi)
	add	%rcx,0x7f(%rdi)
	add	%rcx,0x7f(%r8)
	add	%rcx,0x7f(%r9)
	add	%rcx,0x7f(%r10)
	add	%rcx,0x7f(%r11)
	add	%rcx,0x7f(%r12)
	add	%rcx,0x7f(%r13)
	add	%rcx,0x7f(%r14)
	add	%rcx,0x7f(%r15)
	nop
	add	%rdx,0x7f(%rax)
	add	%rdx,0x7f(%rcx)
	add	%rdx,0x7f(%rdx)
	add	%rdx,0x7f(%rbx)
	add	%rdx,0x7f(%rsp)
	add	%rdx,0x7f(%rbp)
	add	%rdx,0x7f(%rsi)
	add	%rdx,0x7f(%rdi)
	add	%rdx,0x7f(%r8)
	add	%rdx,0x7f(%r9)
	add	%rdx,0x7f(%r10)
	add	%rdx,0x7f(%r11)
	add	%rdx,0x7f(%r12)
	add	%rdx,0x7f(%r13)
	add	%rdx,0x7f(%r14)
	add	%rdx,0x7f(%r15)
	nop
	add	%rbx,0x7f(%rax)
	add	%rbx,0x7f(%rcx)
	add	%rbx,0x7f(%rdx)
	add	%rbx,0x7f(%rbx)
	add	%rbx,0x7f(%rsp)
	add	%rbx,0x7f(%rbp)
	add	%rbx,0x7f(%rsi)
	add	%rbx,0x7f(%rdi)
	add	%rbx,0x7f(%r8)
	add	%rbx,0x7f(%r9)
	add	%rbx,0x7f(%r10)
	add	%rbx,0x7f(%r11)
	add	%rbx,0x7f(%r12)
	add	%rbx,0x7f(%r13)
	add	%rbx,0x7f(%r14)
	add	%rbx,0x7f(%r15)
	nop
	add	%rsp,0x7f(%rax)
	add	%rsp,0x7f(%rcx)
	add	%rsp,0x7f(%rdx)
	add	%rsp,0x7f(%rbx)
	add	%rsp,0x7f(%rsp)
	add	%rsp,0x7f(%rbp)
	add	%rsp,0x7f(%rsi)
	add	%rsp,0x7f(%rdi)
	add	%rsp,0x7f(%r8)
	add	%rsp,0x7f(%r9)
	add	%rsp,0x7f(%r10)
	add	%rsp,0x7f(%r11)
	add	%rsp,0x7f(%r12)
	add	%rsp,0x7f(%r13)
	add	%rsp,0x7f(%r14)
	add	%rsp,0x7f(%r15)
	nop
	add	%rbp,0x7f(%rax)
	add	%rbp,0x7f(%rcx)
	add	%rbp,0x7f(%rdx)
	add	%rbp,0x7f(%rbx)
	add	%rbp,0x7f(%rsp)
	add	%rbp,0x7f(%rbp)
	add	%rbp,0x7f(%rsi)
	add	%rbp,0x7f(%rdi)
	add	%rbp,0x7f(%r8)
	add	%rbp,0x7f(%r9)
	add	%rbp,0x7f(%r10)
	add	%rbp,0x7f(%r11)
	add	%rbp,0x7f(%r12)
	add	%rbp,0x7f(%r13)
	add	%rbp,0x7f(%r14)
	add	%rbp,0x7f(%r15)
	nop
	add	%rsi,0x7f(%rax)
	add	%rsi,0x7f(%rcx)
	add	%rsi,0x7f(%rdx)
	add	%rsi,0x7f(%rbx)
	add	%rsi,0x7f(%rsp)
	add	%rsi,0x7f(%rbp)
	add	%rsi,0x7f(%rsi)
	add	%rsi,0x7f(%rdi)
	add	%rsi,0x7f(%r8)
	add	%rsi,0x7f(%r9)
	add	%rsi,0x7f(%r10)
	add	%rsi,0x7f(%r11)
	add	%rsi,0x7f(%r12)
	add	%rsi,0x7f(%r13)
	add	%rsi,0x7f(%r14)
	add	%rsi,0x7f(%r15)
	nop
	add	%rdi,0x7f(%rax)
	add	%rdi,0x7f(%rcx)
	add	%rdi,0x7f(%rdx)
	add	%rdi,0x7f(%rbx)
	add	%rdi,0x7f(%rsp)
	add	%rdi,0x7f(%rbp)
	add	%rdi,0x7f(%rsi)
	add	%rdi,0x7f(%rdi)
	add	%rdi,0x7f(%r8)
	add	%rdi,0x7f(%r9)
	add	%rdi,0x7f(%r10)
	add	%rdi,0x7f(%r11)
	add	%rdi,0x7f(%r12)
	add	%rdi,0x7f(%r13)
	add	%rdi,0x7f(%r14)
	add	%rdi,0x7f(%r15)
	nop
	add	%r8, 0x7f(%rax)
	add	%r8, 0x7f(%rcx)
	add	%r8, 0x7f(%rdx)
	add	%r8, 0x7f(%rbx)
	add	%r8, 0x7f(%rsp)
	add	%r8, 0x7f(%rbp)
	add	%r8, 0x7f(%rsi)
	add	%r8, 0x7f(%rdi)
	add	%r8, 0x7f(%r8)
	add	%r8, 0x7f(%r9)
	add	%r8, 0x7f(%r10)
	add	%r8, 0x7f(%r11)
	add	%r8, 0x7f(%r12)
	add	%r8, 0x7f(%r13)
	add	%r8, 0x7f(%r14)
	add	%r8, 0x7f(%r15)
	nop
	add	%r9, 0x7f(%rax)
	add	%r9, 0x7f(%rcx)
	add	%r9, 0x7f(%rdx)
	add	%r9, 0x7f(%rbx)
	add	%r9, 0x7f(%rsp)
	add	%r9, 0x7f(%rbp)
	add	%r9, 0x7f(%rsi)
	add	%r9, 0x7f(%rdi)
	add	%r9, 0x7f(%r8)
	add	%r9, 0x7f(%r9)
	add	%r9, 0x7f(%r10)
	add	%r9, 0x7f(%r11)
	add	%r9, 0x7f(%r12)
	add	%r9, 0x7f(%r13)
	add	%r9, 0x7f(%r14)
	add	%r9, 0x7f(%r15)
	nop
	add	%r10,0x7f(%rax)
	add	%r10,0x7f(%rcx)
	add	%r10,0x7f(%rdx)
	add	%r10,0x7f(%rbx)
	add	%r10,0x7f(%rsp)
	add	%r10,0x7f(%rbp)
	add	%r10,0x7f(%rsi)
	add	%r10,0x7f(%rdi)
	add	%r10,0x7f(%r8)
	add	%r10,0x7f(%r9)
	add	%r10,0x7f(%r10)
	add	%r10,0x7f(%r11)
	add	%r10,0x7f(%r12)
	add	%r10,0x7f(%r13)
	add	%r10,0x7f(%r14)
	add	%r10,0x7f(%r15)
	nop
	add	%r11,0x7f(%rax)
	add	%r11,0x7f(%rcx)
	add	%r11,0x7f(%rdx)
	add	%r11,0x7f(%rbx)
	add	%r11,0x7f(%rsp)
	add	%r11,0x7f(%rbp)
	add	%r11,0x7f(%rsi)
	add	%r11,0x7f(%rdi)
	add	%r11,0x7f(%r8)
	add	%r11,0x7f(%r9)
	add	%r11,0x7f(%r10)
	add	%r11,0x7f(%r11)
	add	%r11,0x7f(%r12)
	add	%r11,0x7f(%r13)
	add	%r11,0x7f(%r14)
	add	%r11,0x7f(%r15)
	nop
	add	%r12,0x7f(%rax)
	add	%r12,0x7f(%rcx)
	add	%r12,0x7f(%rdx)
	add	%r12,0x7f(%rbx)
	add	%r12,0x7f(%rsp)
	add	%r12,0x7f(%rbp)
	add	%r12,0x7f(%rsi)
	add	%r12,0x7f(%rdi)
	add	%r12,0x7f(%r8)
	add	%r12,0x7f(%r9)
	add	%r12,0x7f(%r10)
	add	%r12,0x7f(%r11)
	add	%r12,0x7f(%r12)
	add	%r12,0x7f(%r13)
	add	%r12,0x7f(%r14)
	add	%r12,0x7f(%r15)
	nop
	add	%r13,0x7f(%rax)
	add	%r13,0x7f(%rcx)
	add	%r13,0x7f(%rdx)
	add	%r13,0x7f(%rbx)
	add	%r13,0x7f(%rsp)
	add	%r13,0x7f(%rbp)
	add	%r13,0x7f(%rsi)
	add	%r13,0x7f(%rdi)
	add	%r13,0x7f(%r8)
	add	%r13,0x7f(%r9)
	add	%r13,0x7f(%r10)
	add	%r13,0x7f(%r11)
	add	%r13,0x7f(%r12)
	add	%r13,0x7f(%r13)
	add	%r13,0x7f(%r14)
	add	%r13,0x7f(%r15)
	nop
	add	%r14,0x7f(%rax)
	add	%r14,0x7f(%rcx)
	add	%r14,0x7f(%rdx)
	add	%r14,0x7f(%rbx)
	add	%r14,0x7f(%rsp)
	add	%r14,0x7f(%rbp)
	add	%r14,0x7f(%rsi)
	add	%r14,0x7f(%rdi)
	add	%r14,0x7f(%r8)
	add	%r14,0x7f(%r9)
	add	%r14,0x7f(%r10)
	add	%r14,0x7f(%r11)
	add	%r14,0x7f(%r12)
	add	%r14,0x7f(%r13)
	add	%r14,0x7f(%r14)
	add	%r14,0x7f(%r15)
	nop
	add	%r15,0x7f(%rax)
	add	%r15,0x7f(%rcx)
	add	%r15,0x7f(%rdx)
	add	%r15,0x7f(%rbx)
	add	%r15,0x7f(%rsp)
	add	%r15,0x7f(%rbp)
	add	%r15,0x7f(%rsi)
	add	%r15,0x7f(%rdi)
	add	%r15,0x7f(%r8)
	add	%r15,0x7f(%r9)
	add	%r15,0x7f(%r10)
	add	%r15,0x7f(%r11)
	add	%r15,0x7f(%r12)
	add	%r15,0x7f(%r13)
	add	%r15,0x7f(%r14)
	add	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AddMem32Reg
	.type	AddMem32Reg, @function
AddMem32Reg:
	.cfi_startproc
	add	%rax,0x7f563412(%rax)
	add	%rax,0x7f563412(%rcx)
	add	%rax,0x7f563412(%rdx)
	add	%rax,0x7f563412(%rbx)
	add	%rax,0x7f563412(%rsp)
	add	%rax,0x7f563412(%rbp)
	add	%rax,0x7f563412(%rsi)
	add	%rax,0x7f563412(%rdi)
	add	%rax,0x7f563412(%r8)
	add	%rax,0x7f563412(%r9)
	add	%rax,0x7f563412(%r10)
	add	%rax,0x7f563412(%r11)
	add	%rax,0x7f563412(%r12)
	add	%rax,0x7f563412(%r13)
	add	%rax,0x7f563412(%r14)
	add	%rax,0x7f563412(%r15)
	nop
	add	%rcx,0x7f563412(%rax)
	add	%rcx,0x7f563412(%rcx)
	add	%rcx,0x7f563412(%rdx)
	add	%rcx,0x7f563412(%rbx)
	add	%rcx,0x7f563412(%rsp)
	add	%rcx,0x7f563412(%rbp)
	add	%rcx,0x7f563412(%rsi)
	add	%rcx,0x7f563412(%rdi)
	add	%rcx,0x7f563412(%r8)
	add	%rcx,0x7f563412(%r9)
	add	%rcx,0x7f563412(%r10)
	add	%rcx,0x7f563412(%r11)
	add	%rcx,0x7f563412(%r12)
	add	%rcx,0x7f563412(%r13)
	add	%rcx,0x7f563412(%r14)
	add	%rcx,0x7f563412(%r15)
	nop
	add	%rdx,0x7f563412(%rax)
	add	%rdx,0x7f563412(%rcx)
	add	%rdx,0x7f563412(%rdx)
	add	%rdx,0x7f563412(%rbx)
	add	%rdx,0x7f563412(%rsp)
	add	%rdx,0x7f563412(%rbp)
	add	%rdx,0x7f563412(%rsi)
	add	%rdx,0x7f563412(%rdi)
	add	%rdx,0x7f563412(%r8)
	add	%rdx,0x7f563412(%r9)
	add	%rdx,0x7f563412(%r10)
	add	%rdx,0x7f563412(%r11)
	add	%rdx,0x7f563412(%r12)
	add	%rdx,0x7f563412(%r13)
	add	%rdx,0x7f563412(%r14)
	add	%rdx,0x7f563412(%r15)
	nop
	add	%rbx,0x7f563412(%rax)
	add	%rbx,0x7f563412(%rcx)
	add	%rbx,0x7f563412(%rdx)
	add	%rbx,0x7f563412(%rbx)
	add	%rbx,0x7f563412(%rsp)
	add	%rbx,0x7f563412(%rbp)
	add	%rbx,0x7f563412(%rsi)
	add	%rbx,0x7f563412(%rdi)
	add	%rbx,0x7f563412(%r8)
	add	%rbx,0x7f563412(%r9)
	add	%rbx,0x7f563412(%r10)
	add	%rbx,0x7f563412(%r11)
	add	%rbx,0x7f563412(%r12)
	add	%rbx,0x7f563412(%r13)
	add	%rbx,0x7f563412(%r14)
	add	%rbx,0x7f563412(%r15)
	nop
	add	%rsp,0x7f563412(%rax)
	add	%rsp,0x7f563412(%rcx)
	add	%rsp,0x7f563412(%rdx)
	add	%rsp,0x7f563412(%rbx)
	add	%rsp,0x7f563412(%rsp)
	add	%rsp,0x7f563412(%rbp)
	add	%rsp,0x7f563412(%rsi)
	add	%rsp,0x7f563412(%rdi)
	add	%rsp,0x7f563412(%r8)
	add	%rsp,0x7f563412(%r9)
	add	%rsp,0x7f563412(%r10)
	add	%rsp,0x7f563412(%r11)
	add	%rsp,0x7f563412(%r12)
	add	%rsp,0x7f563412(%r13)
	add	%rsp,0x7f563412(%r14)
	add	%rsp,0x7f563412(%r15)
	nop
	add	%rbp,0x7f563412(%rax)
	add	%rbp,0x7f563412(%rcx)
	add	%rbp,0x7f563412(%rdx)
	add	%rbp,0x7f563412(%rbx)
	add	%rbp,0x7f563412(%rsp)
	add	%rbp,0x7f563412(%rbp)
	add	%rbp,0x7f563412(%rsi)
	add	%rbp,0x7f563412(%rdi)
	add	%rbp,0x7f563412(%r8)
	add	%rbp,0x7f563412(%r9)
	add	%rbp,0x7f563412(%r10)
	add	%rbp,0x7f563412(%r11)
	add	%rbp,0x7f563412(%r12)
	add	%rbp,0x7f563412(%r13)
	add	%rbp,0x7f563412(%r14)
	add	%rbp,0x7f563412(%r15)
	nop
	add	%rsi,0x7f563412(%rax)
	add	%rsi,0x7f563412(%rcx)
	add	%rsi,0x7f563412(%rdx)
	add	%rsi,0x7f563412(%rbx)
	add	%rsi,0x7f563412(%rsp)
	add	%rsi,0x7f563412(%rbp)
	add	%rsi,0x7f563412(%rsi)
	add	%rsi,0x7f563412(%rdi)
	add	%rsi,0x7f563412(%r8)
	add	%rsi,0x7f563412(%r9)
	add	%rsi,0x7f563412(%r10)
	add	%rsi,0x7f563412(%r11)
	add	%rsi,0x7f563412(%r12)
	add	%rsi,0x7f563412(%r13)
	add	%rsi,0x7f563412(%r14)
	add	%rsi,0x7f563412(%r15)
	nop
	add	%rdi,0x7f563412(%rax)
	add	%rdi,0x7f563412(%rcx)
	add	%rdi,0x7f563412(%rdx)
	add	%rdi,0x7f563412(%rbx)
	add	%rdi,0x7f563412(%rsp)
	add	%rdi,0x7f563412(%rbp)
	add	%rdi,0x7f563412(%rsi)
	add	%rdi,0x7f563412(%rdi)
	add	%rdi,0x7f563412(%r8)
	add	%rdi,0x7f563412(%r9)
	add	%rdi,0x7f563412(%r10)
	add	%rdi,0x7f563412(%r11)
	add	%rdi,0x7f563412(%r12)
	add	%rdi,0x7f563412(%r13)
	add	%rdi,0x7f563412(%r14)
	add	%rdi,0x7f563412(%r15)
	nop
	add	%r8, 0x7f563412(%rax)
	add	%r8, 0x7f563412(%rcx)
	add	%r8, 0x7f563412(%rdx)
	add	%r8, 0x7f563412(%rbx)
	add	%r8, 0x7f563412(%rsp)
	add	%r8, 0x7f563412(%rbp)
	add	%r8, 0x7f563412(%rsi)
	add	%r8, 0x7f563412(%rdi)
	add	%r8, 0x7f563412(%r8)
	add	%r8, 0x7f563412(%r9)
	add	%r8, 0x7f563412(%r10)
	add	%r8, 0x7f563412(%r11)
	add	%r8, 0x7f563412(%r12)
	add	%r8, 0x7f563412(%r13)
	add	%r8, 0x7f563412(%r14)
	add	%r8, 0x7f563412(%r15)
	nop
	add	%r9, 0x7f563412(%rax)
	add	%r9, 0x7f563412(%rcx)
	add	%r9, 0x7f563412(%rdx)
	add	%r9, 0x7f563412(%rbx)
	add	%r9, 0x7f563412(%rsp)
	add	%r9, 0x7f563412(%rbp)
	add	%r9, 0x7f563412(%rsi)
	add	%r9, 0x7f563412(%rdi)
	add	%r9, 0x7f563412(%r8)
	add	%r9, 0x7f563412(%r9)
	add	%r9, 0x7f563412(%r10)
	add	%r9, 0x7f563412(%r11)
	add	%r9, 0x7f563412(%r12)
	add	%r9, 0x7f563412(%r13)
	add	%r9, 0x7f563412(%r14)
	add	%r9, 0x7f563412(%r15)
	nop
	add	%r10,0x7f563412(%rax)
	add	%r10,0x7f563412(%rcx)
	add	%r10,0x7f563412(%rdx)
	add	%r10,0x7f563412(%rbx)
	add	%r10,0x7f563412(%rsp)
	add	%r10,0x7f563412(%rbp)
	add	%r10,0x7f563412(%rsi)
	add	%r10,0x7f563412(%rdi)
	add	%r10,0x7f563412(%r8)
	add	%r10,0x7f563412(%r9)
	add	%r10,0x7f563412(%r10)
	add	%r10,0x7f563412(%r11)
	add	%r10,0x7f563412(%r12)
	add	%r10,0x7f563412(%r13)
	add	%r10,0x7f563412(%r14)
	add	%r10,0x7f563412(%r15)
	nop
	add	%r11,0x7f563412(%rax)
	add	%r11,0x7f563412(%rcx)
	add	%r11,0x7f563412(%rdx)
	add	%r11,0x7f563412(%rbx)
	add	%r11,0x7f563412(%rsp)
	add	%r11,0x7f563412(%rbp)
	add	%r11,0x7f563412(%rsi)
	add	%r11,0x7f563412(%rdi)
	add	%r11,0x7f563412(%r8)
	add	%r11,0x7f563412(%r9)
	add	%r11,0x7f563412(%r10)
	add	%r11,0x7f563412(%r11)
	add	%r11,0x7f563412(%r12)
	add	%r11,0x7f563412(%r13)
	add	%r11,0x7f563412(%r14)
	add	%r11,0x7f563412(%r15)
	nop
	add	%r12,0x7f563412(%rax)
	add	%r12,0x7f563412(%rcx)
	add	%r12,0x7f563412(%rdx)
	add	%r12,0x7f563412(%rbx)
	add	%r12,0x7f563412(%rsp)
	add	%r12,0x7f563412(%rbp)
	add	%r12,0x7f563412(%rsi)
	add	%r12,0x7f563412(%rdi)
	add	%r12,0x7f563412(%r8)
	add	%r12,0x7f563412(%r9)
	add	%r12,0x7f563412(%r10)
	add	%r12,0x7f563412(%r11)
	add	%r12,0x7f563412(%r12)
	add	%r12,0x7f563412(%r13)
	add	%r12,0x7f563412(%r14)
	add	%r12,0x7f563412(%r15)
	nop
	add	%r13,0x7f563412(%rax)
	add	%r13,0x7f563412(%rcx)
	add	%r13,0x7f563412(%rdx)
	add	%r13,0x7f563412(%rbx)
	add	%r13,0x7f563412(%rsp)
	add	%r13,0x7f563412(%rbp)
	add	%r13,0x7f563412(%rsi)
	add	%r13,0x7f563412(%rdi)
	add	%r13,0x7f563412(%r8)
	add	%r13,0x7f563412(%r9)
	add	%r13,0x7f563412(%r10)
	add	%r13,0x7f563412(%r11)
	add	%r13,0x7f563412(%r12)
	add	%r13,0x7f563412(%r13)
	add	%r13,0x7f563412(%r14)
	add	%r13,0x7f563412(%r15)
	nop
	add	%r14,0x7f563412(%rax)
	add	%r14,0x7f563412(%rcx)
	add	%r14,0x7f563412(%rdx)
	add	%r14,0x7f563412(%rbx)
	add	%r14,0x7f563412(%rsp)
	add	%r14,0x7f563412(%rbp)
	add	%r14,0x7f563412(%rsi)
	add	%r14,0x7f563412(%rdi)
	add	%r14,0x7f563412(%r8)
	add	%r14,0x7f563412(%r9)
	add	%r14,0x7f563412(%r10)
	add	%r14,0x7f563412(%r11)
	add	%r14,0x7f563412(%r12)
	add	%r14,0x7f563412(%r13)
	add	%r14,0x7f563412(%r14)
	add	%r14,0x7f563412(%r15)
	nop
	add	%r15,0x7f563412(%rax)
	add	%r15,0x7f563412(%rcx)
	add	%r15,0x7f563412(%rdx)
	add	%r15,0x7f563412(%rbx)
	add	%r15,0x7f563412(%rsp)
	add	%r15,0x7f563412(%rbp)
	add	%r15,0x7f563412(%rsi)
	add	%r15,0x7f563412(%rdi)
	add	%r15,0x7f563412(%r8)
	add	%r15,0x7f563412(%r9)
	add	%r15,0x7f563412(%r10)
	add	%r15,0x7f563412(%r11)
	add	%r15,0x7f563412(%r12)
	add	%r15,0x7f563412(%r13)
	add	%r15,0x7f563412(%r14)
	add	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	AddRegMem
	.type	AddRegMem, @function
AddRegMem:
	.cfi_startproc
	add	(%rax),%rax
	add	(%rax),%rcx
	add	(%rax),%rdx
	add	(%rax),%rbx
	add	(%rax),%rsp
	add	(%rax),%rbp
	add	(%rax),%rsi
	add	(%rax),%rdi
	add	(%rax),%r8
	add	(%rax),%r9
	add	(%rax),%r10
	add	(%rax),%r11
	add	(%rax),%r12
	add	(%rax),%r13
	add	(%rax),%r14
	add	(%rax),%r15
	nop
	add	(%rcx),%rax
	add	(%rcx),%rcx
	add	(%rcx),%rdx
	add	(%rcx),%rbx
	add	(%rcx),%rsp
	add	(%rcx),%rbp
	add	(%rcx),%rsi
	add	(%rcx),%rdi
	add	(%rcx),%r8
	add	(%rcx),%r9
	add	(%rcx),%r10
	add	(%rcx),%r11
	add	(%rcx),%r12
	add	(%rcx),%r13
	add	(%rcx),%r14
	add	(%rcx),%r15
	nop
	add	(%rdx),%rax
	add	(%rdx),%rcx
	add	(%rdx),%rdx
	add	(%rdx),%rbx
	add	(%rdx),%rsp
	add	(%rdx),%rbp
	add	(%rdx),%rsi
	add	(%rdx),%rdi
	add	(%rdx),%r8
	add	(%rdx),%r9
	add	(%rdx),%r10
	add	(%rdx),%r11
	add	(%rdx),%r12
	add	(%rdx),%r13
	add	(%rdx),%r14
	add	(%rdx),%r15
	nop
	add	(%rbx),%rax
	add	(%rbx),%rcx
	add	(%rbx),%rdx
	add	(%rbx),%rbx
	add	(%rbx),%rsp
	add	(%rbx),%rbp
	add	(%rbx),%rsi
	add	(%rbx),%rdi
	add	(%rbx),%r8
	add	(%rbx),%r9
	add	(%rbx),%r10
	add	(%rbx),%r11
	add	(%rbx),%r12
	add	(%rbx),%r13
	add	(%rbx),%r14
	add	(%rbx),%r15
	nop
	add	(%rsp),%rax
	add	(%rsp),%rcx
	add	(%rsp),%rdx
	add	(%rsp),%rbx
	add	(%rsp),%rsp
	add	(%rsp),%rbp
	add	(%rsp),%rsi
	add	(%rsp),%rdi
	add	(%rsp),%r8
	add	(%rsp),%r9
	add	(%rsp),%r10
	add	(%rsp),%r11
	add	(%rsp),%r12
	add	(%rsp),%r13
	add	(%rsp),%r14
	add	(%rsp),%r15
	nop
	add	(%rbp),%rax
	add	(%rbp),%rcx
	add	(%rbp),%rdx
	add	(%rbp),%rbx
	add	(%rbp),%rsp
	add	(%rbp),%rbp
	add	(%rbp),%rsi
	add	(%rbp),%rdi
	add	(%rbp),%r8
	add	(%rbp),%r9
	add	(%rbp),%r10
	add	(%rbp),%r11
	add	(%rbp),%r12
	add	(%rbp),%r13
	add	(%rbp),%r14
	add	(%rbp),%r15
	nop
	add	(%rsi),%rax
	add	(%rsi),%rcx
	add	(%rsi),%rdx
	add	(%rsi),%rbx
	add	(%rsi),%rsp
	add	(%rsi),%rbp
	add	(%rsi),%rsi
	add	(%rsi),%rdi
	add	(%rsi),%r8
	add	(%rsi),%r9
	add	(%rsi),%r10
	add	(%rsi),%r11
	add	(%rsi),%r12
	add	(%rsi),%r13
	add	(%rsi),%r14
	add	(%rsi),%r15
	nop
	add	(%rdi),%rax
	add	(%rdi),%rcx
	add	(%rdi),%rdx
	add	(%rdi),%rbx
	add	(%rdi),%rsp
	add	(%rdi),%rbp
	add	(%rdi),%rsi
	add	(%rdi),%rdi
	add	(%rdi),%r8
	add	(%rdi),%r9
	add	(%rdi),%r10
	add	(%rdi),%r11
	add	(%rdi),%r12
	add	(%rdi),%r13
	add	(%rdi),%r14
	add	(%rdi),%r15
	nop
	add	(%r8), %rax
	add	(%r8), %rcx
	add	(%r8), %rdx
	add	(%r8), %rbx
	add	(%r8), %rsp
	add	(%r8), %rbp
	add	(%r8), %rsi
	add	(%r8), %rdi
	add	(%r8), %r8
	add	(%r8), %r9
	add	(%r8), %r10
	add	(%r8), %r11
	add	(%r8), %r12
	add	(%r8), %r13
	add	(%r8), %r14
	add	(%r8), %r15
	nop
	add	(%r9), %rax
	add	(%r9), %rcx
	add	(%r9), %rdx
	add	(%r9), %rbx
	add	(%r9), %rsp
	add	(%r9), %rbp
	add	(%r9), %rsi
	add	(%r9), %rdi
	add	(%r9), %r8
	add	(%r9), %r9
	add	(%r9), %r10
	add	(%r9), %r11
	add	(%r9), %r12
	add	(%r9), %r13
	add	(%r9), %r14
	add	(%r9), %r15
	nop
	add	(%r10),%rax
	add	(%r10),%rcx
	add	(%r10),%rdx
	add	(%r10),%rbx
	add	(%r10),%rsp
	add	(%r10),%rbp
	add	(%r10),%rsi
	add	(%r10),%rdi
	add	(%r10),%r8
	add	(%r10),%r9
	add	(%r10),%r10
	add	(%r10),%r11
	add	(%r10),%r12
	add	(%r10),%r13
	add	(%r10),%r14
	add	(%r10),%r15
	nop
	add	(%r11),%rax
	add	(%r11),%rcx
	add	(%r11),%rdx
	add	(%r11),%rbx
	add	(%r11),%rsp
	add	(%r11),%rbp
	add	(%r11),%rsi
	add	(%r11),%rdi
	add	(%r11),%r8
	add	(%r11),%r9
	add	(%r11),%r10
	add	(%r11),%r11
	add	(%r11),%r12
	add	(%r11),%r13
	add	(%r11),%r14
	add	(%r11),%r15
	nop
	add	(%r12),%rax
	add	(%r12),%rcx
	add	(%r12),%rdx
	add	(%r12),%rbx
	add	(%r12),%rsp
	add	(%r12),%rbp
	add	(%r12),%rsi
	add	(%r12),%rdi
	add	(%r12),%r8
	add	(%r12),%r9
	add	(%r12),%r10
	add	(%r12),%r11
	add	(%r12),%r12
	add	(%r12),%r13
	add	(%r12),%r14
	add	(%r12),%r15
	nop
	add	(%r13),%rax
	add	(%r13),%rcx
	add	(%r13),%rdx
	add	(%r13),%rbx
	add	(%r13),%rsp
	add	(%r13),%rbp
	add	(%r13),%rsi
	add	(%r13),%rdi
	add	(%r13),%r8
	add	(%r13),%r9
	add	(%r13),%r10
	add	(%r13),%r11
	add	(%r13),%r12
	add	(%r13),%r13
	add	(%r13),%r14
	add	(%r13),%r15
	nop
	add	(%r14),%rax
	add	(%r14),%rcx
	add	(%r14),%rdx
	add	(%r14),%rbx
	add	(%r14),%rsp
	add	(%r14),%rbp
	add	(%r14),%rsi
	add	(%r14),%rdi
	add	(%r14),%r8
	add	(%r14),%r9
	add	(%r14),%r10
	add	(%r14),%r11
	add	(%r14),%r12
	add	(%r14),%r13
	add	(%r14),%r14
	add	(%r14),%r15
	nop
	add	(%r15),%rax
	add	(%r15),%rcx
	add	(%r15),%rdx
	add	(%r15),%rbx
	add	(%r15),%rsp
	add	(%r15),%rbp
	add	(%r15),%rsi
	add	(%r15),%rdi
	add	(%r15),%r8
	add	(%r15),%r9
	add	(%r15),%r10
	add	(%r15),%r11
	add	(%r15),%r12
	add	(%r15),%r13
	add	(%r15),%r14
	add	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AddRegMem8
	.type	AddRegMem8, @function
AddRegMem8:
	.cfi_startproc
	add	0x7f(%rax),%rax
	add	0x7f(%rax),%rcx
	add	0x7f(%rax),%rdx
	add	0x7f(%rax),%rbx
	add	0x7f(%rax),%rsp
	add	0x7f(%rax),%rbp
	add	0x7f(%rax),%rsi
	add	0x7f(%rax),%rdi
	add	0x7f(%rax),%r8
	add	0x7f(%rax),%r9
	add	0x7f(%rax),%r10
	add	0x7f(%rax),%r11
	add	0x7f(%rax),%r12
	add	0x7f(%rax),%r13
	add	0x7f(%rax),%r14
	add	0x7f(%rax),%r15
	nop
	add	0x7f(%rcx),%rax
	add	0x7f(%rcx),%rcx
	add	0x7f(%rcx),%rdx
	add	0x7f(%rcx),%rbx
	add	0x7f(%rcx),%rsp
	add	0x7f(%rcx),%rbp
	add	0x7f(%rcx),%rsi
	add	0x7f(%rcx),%rdi
	add	0x7f(%rcx),%r8
	add	0x7f(%rcx),%r9
	add	0x7f(%rcx),%r10
	add	0x7f(%rcx),%r11
	add	0x7f(%rcx),%r12
	add	0x7f(%rcx),%r13
	add	0x7f(%rcx),%r14
	add	0x7f(%rcx),%r15
	nop
	add	0x7f(%rdx),%rax
	add	0x7f(%rdx),%rcx
	add	0x7f(%rdx),%rdx
	add	0x7f(%rdx),%rbx
	add	0x7f(%rdx),%rsp
	add	0x7f(%rdx),%rbp
	add	0x7f(%rdx),%rsi
	add	0x7f(%rdx),%rdi
	add	0x7f(%rdx),%r8
	add	0x7f(%rdx),%r9
	add	0x7f(%rdx),%r10
	add	0x7f(%rdx),%r11
	add	0x7f(%rdx),%r12
	add	0x7f(%rdx),%r13
	add	0x7f(%rdx),%r14
	add	0x7f(%rdx),%r15
	nop
	add	0x7f(%rbx),%rax
	add	0x7f(%rbx),%rcx
	add	0x7f(%rbx),%rdx
	add	0x7f(%rbx),%rbx
	add	0x7f(%rbx),%rsp
	add	0x7f(%rbx),%rbp
	add	0x7f(%rbx),%rsi
	add	0x7f(%rbx),%rdi
	add	0x7f(%rbx),%r8
	add	0x7f(%rbx),%r9
	add	0x7f(%rbx),%r10
	add	0x7f(%rbx),%r11
	add	0x7f(%rbx),%r12
	add	0x7f(%rbx),%r13
	add	0x7f(%rbx),%r14
	add	0x7f(%rbx),%r15
	nop
	add	0x7f(%rsp),%rax
	add	0x7f(%rsp),%rcx
	add	0x7f(%rsp),%rdx
	add	0x7f(%rsp),%rbx
	add	0x7f(%rsp),%rsp
	add	0x7f(%rsp),%rbp
	add	0x7f(%rsp),%rsi
	add	0x7f(%rsp),%rdi
	add	0x7f(%rsp),%r8
	add	0x7f(%rsp),%r9
	add	0x7f(%rsp),%r10
	add	0x7f(%rsp),%r11
	add	0x7f(%rsp),%r12
	add	0x7f(%rsp),%r13
	add	0x7f(%rsp),%r14
	add	0x7f(%rsp),%r15
	nop
	add	0x7f(%rbp),%rax
	add	0x7f(%rbp),%rcx
	add	0x7f(%rbp),%rdx
	add	0x7f(%rbp),%rbx
	add	0x7f(%rbp),%rsp
	add	0x7f(%rbp),%rbp
	add	0x7f(%rbp),%rsi
	add	0x7f(%rbp),%rdi
	add	0x7f(%rbp),%r8
	add	0x7f(%rbp),%r9
	add	0x7f(%rbp),%r10
	add	0x7f(%rbp),%r11
	add	0x7f(%rbp),%r12
	add	0x7f(%rbp),%r13
	add	0x7f(%rbp),%r14
	add	0x7f(%rbp),%r15
	nop
	add	0x7f(%rsi),%rax
	add	0x7f(%rsi),%rcx
	add	0x7f(%rsi),%rdx
	add	0x7f(%rsi),%rbx
	add	0x7f(%rsi),%rsp
	add	0x7f(%rsi),%rbp
	add	0x7f(%rsi),%rsi
	add	0x7f(%rsi),%rdi
	add	0x7f(%rsi),%r8
	add	0x7f(%rsi),%r9
	add	0x7f(%rsi),%r10
	add	0x7f(%rsi),%r11
	add	0x7f(%rsi),%r12
	add	0x7f(%rsi),%r13
	add	0x7f(%rsi),%r14
	add	0x7f(%rsi),%r15
	nop
	add	0x7f(%rdi),%rax
	add	0x7f(%rdi),%rcx
	add	0x7f(%rdi),%rdx
	add	0x7f(%rdi),%rbx
	add	0x7f(%rdi),%rsp
	add	0x7f(%rdi),%rbp
	add	0x7f(%rdi),%rsi
	add	0x7f(%rdi),%rdi
	add	0x7f(%rdi),%r8
	add	0x7f(%rdi),%r9
	add	0x7f(%rdi),%r10
	add	0x7f(%rdi),%r11
	add	0x7f(%rdi),%r12
	add	0x7f(%rdi),%r13
	add	0x7f(%rdi),%r14
	add	0x7f(%rdi),%r15
	nop
	add	0x7f(%r8), %rax
	add	0x7f(%r8), %rcx
	add	0x7f(%r8), %rdx
	add	0x7f(%r8), %rbx
	add	0x7f(%r8), %rsp
	add	0x7f(%r8), %rbp
	add	0x7f(%r8), %rsi
	add	0x7f(%r8), %rdi
	add	0x7f(%r8), %r8
	add	0x7f(%r8), %r9
	add	0x7f(%r8), %r10
	add	0x7f(%r8), %r11
	add	0x7f(%r8), %r12
	add	0x7f(%r8), %r13
	add	0x7f(%r8), %r14
	add	0x7f(%r8), %r15
	nop
	add	0x7f(%r9), %rax
	add	0x7f(%r9), %rcx
	add	0x7f(%r9), %rdx
	add	0x7f(%r9), %rbx
	add	0x7f(%r9), %rsp
	add	0x7f(%r9), %rbp
	add	0x7f(%r9), %rsi
	add	0x7f(%r9), %rdi
	add	0x7f(%r9), %r8
	add	0x7f(%r9), %r9
	add	0x7f(%r9), %r10
	add	0x7f(%r9), %r11
	add	0x7f(%r9), %r12
	add	0x7f(%r9), %r13
	add	0x7f(%r9), %r14
	add	0x7f(%r9), %r15
	nop
	add	0x7f(%r10),%rax
	add	0x7f(%r10),%rcx
	add	0x7f(%r10),%rdx
	add	0x7f(%r10),%rbx
	add	0x7f(%r10),%rsp
	add	0x7f(%r10),%rbp
	add	0x7f(%r10),%rsi
	add	0x7f(%r10),%rdi
	add	0x7f(%r10),%r8
	add	0x7f(%r10),%r9
	add	0x7f(%r10),%r10
	add	0x7f(%r10),%r11
	add	0x7f(%r10),%r12
	add	0x7f(%r10),%r13
	add	0x7f(%r10),%r14
	add	0x7f(%r10),%r15
	nop
	add	0x7f(%r11),%rax
	add	0x7f(%r11),%rcx
	add	0x7f(%r11),%rdx
	add	0x7f(%r11),%rbx
	add	0x7f(%r11),%rsp
	add	0x7f(%r11),%rbp
	add	0x7f(%r11),%rsi
	add	0x7f(%r11),%rdi
	add	0x7f(%r11),%r8
	add	0x7f(%r11),%r9
	add	0x7f(%r11),%r10
	add	0x7f(%r11),%r11
	add	0x7f(%r11),%r12
	add	0x7f(%r11),%r13
	add	0x7f(%r11),%r14
	add	0x7f(%r11),%r15
	nop
	add	0x7f(%r12),%rax
	add	0x7f(%r12),%rcx
	add	0x7f(%r12),%rdx
	add	0x7f(%r12),%rbx
	add	0x7f(%r12),%rsp
	add	0x7f(%r12),%rbp
	add	0x7f(%r12),%rsi
	add	0x7f(%r12),%rdi
	add	0x7f(%r12),%r8
	add	0x7f(%r12),%r9
	add	0x7f(%r12),%r10
	add	0x7f(%r12),%r11
	add	0x7f(%r12),%r12
	add	0x7f(%r12),%r13
	add	0x7f(%r12),%r14
	add	0x7f(%r12),%r15
	nop
	add	0x7f(%r13),%rax
	add	0x7f(%r13),%rcx
	add	0x7f(%r13),%rdx
	add	0x7f(%r13),%rbx
	add	0x7f(%r13),%rsp
	add	0x7f(%r13),%rbp
	add	0x7f(%r13),%rsi
	add	0x7f(%r13),%rdi
	add	0x7f(%r13),%r8
	add	0x7f(%r13),%r9
	add	0x7f(%r13),%r10
	add	0x7f(%r13),%r11
	add	0x7f(%r13),%r12
	add	0x7f(%r13),%r13
	add	0x7f(%r13),%r14
	add	0x7f(%r13),%r15
	nop
	add	0x7f(%r14),%rax
	add	0x7f(%r14),%rcx
	add	0x7f(%r14),%rdx
	add	0x7f(%r14),%rbx
	add	0x7f(%r14),%rsp
	add	0x7f(%r14),%rbp
	add	0x7f(%r14),%rsi
	add	0x7f(%r14),%rdi
	add	0x7f(%r14),%r8
	add	0x7f(%r14),%r9
	add	0x7f(%r14),%r10
	add	0x7f(%r14),%r11
	add	0x7f(%r14),%r12
	add	0x7f(%r14),%r13
	add	0x7f(%r14),%r14
	add	0x7f(%r14),%r15
	nop
	add	0x7f(%r15),%rax
	add	0x7f(%r15),%rcx
	add	0x7f(%r15),%rdx
	add	0x7f(%r15),%rbx
	add	0x7f(%r15),%rsp
	add	0x7f(%r15),%rbp
	add	0x7f(%r15),%rsi
	add	0x7f(%r15),%rdi
	add	0x7f(%r15),%r8
	add	0x7f(%r15),%r9
	add	0x7f(%r15),%r10
	add	0x7f(%r15),%r11
	add	0x7f(%r15),%r12
	add	0x7f(%r15),%r13
	add	0x7f(%r15),%r14
	add	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AddRegMem32
	.type	AddRegMem32, @function
AddRegMem32:
	.cfi_startproc
	add	0x7f563412(%rax),%rax
	add	0x7f563412(%rax),%rcx
	add	0x7f563412(%rax),%rdx
	add	0x7f563412(%rax),%rbx
	add	0x7f563412(%rax),%rsp
	add	0x7f563412(%rax),%rbp
	add	0x7f563412(%rax),%rsi
	add	0x7f563412(%rax),%rdi
	add	0x7f563412(%rax),%r8
	add	0x7f563412(%rax),%r9
	add	0x7f563412(%rax),%r10
	add	0x7f563412(%rax),%r11
	add	0x7f563412(%rax),%r12
	add	0x7f563412(%rax),%r13
	add	0x7f563412(%rax),%r14
	add	0x7f563412(%rax),%r15
	nop
	add	0x7f563412(%rcx),%rax
	add	0x7f563412(%rcx),%rcx
	add	0x7f563412(%rcx),%rdx
	add	0x7f563412(%rcx),%rbx
	add	0x7f563412(%rcx),%rsp
	add	0x7f563412(%rcx),%rbp
	add	0x7f563412(%rcx),%rsi
	add	0x7f563412(%rcx),%rdi
	add	0x7f563412(%rcx),%r8
	add	0x7f563412(%rcx),%r9
	add	0x7f563412(%rcx),%r10
	add	0x7f563412(%rcx),%r11
	add	0x7f563412(%rcx),%r12
	add	0x7f563412(%rcx),%r13
	add	0x7f563412(%rcx),%r14
	add	0x7f563412(%rcx),%r15
	nop
	add	0x7f563412(%rdx),%rax
	add	0x7f563412(%rdx),%rcx
	add	0x7f563412(%rdx),%rdx
	add	0x7f563412(%rdx),%rbx
	add	0x7f563412(%rdx),%rsp
	add	0x7f563412(%rdx),%rbp
	add	0x7f563412(%rdx),%rsi
	add	0x7f563412(%rdx),%rdi
	add	0x7f563412(%rdx),%r8
	add	0x7f563412(%rdx),%r9
	add	0x7f563412(%rdx),%r10
	add	0x7f563412(%rdx),%r11
	add	0x7f563412(%rdx),%r12
	add	0x7f563412(%rdx),%r13
	add	0x7f563412(%rdx),%r14
	add	0x7f563412(%rdx),%r15
	nop
	add	0x7f563412(%rbx),%rax
	add	0x7f563412(%rbx),%rcx
	add	0x7f563412(%rbx),%rdx
	add	0x7f563412(%rbx),%rbx
	add	0x7f563412(%rbx),%rsp
	add	0x7f563412(%rbx),%rbp
	add	0x7f563412(%rbx),%rsi
	add	0x7f563412(%rbx),%rdi
	add	0x7f563412(%rbx),%r8
	add	0x7f563412(%rbx),%r9
	add	0x7f563412(%rbx),%r10
	add	0x7f563412(%rbx),%r11
	add	0x7f563412(%rbx),%r12
	add	0x7f563412(%rbx),%r13
	add	0x7f563412(%rbx),%r14
	add	0x7f563412(%rbx),%r15
	nop
	add	0x7f563412(%rsp),%rax
	add	0x7f563412(%rsp),%rcx
	add	0x7f563412(%rsp),%rdx
	add	0x7f563412(%rsp),%rbx
	add	0x7f563412(%rsp),%rsp
	add	0x7f563412(%rsp),%rbp
	add	0x7f563412(%rsp),%rsi
	add	0x7f563412(%rsp),%rdi
	add	0x7f563412(%rsp),%r8
	add	0x7f563412(%rsp),%r9
	add	0x7f563412(%rsp),%r10
	add	0x7f563412(%rsp),%r11
	add	0x7f563412(%rsp),%r12
	add	0x7f563412(%rsp),%r13
	add	0x7f563412(%rsp),%r14
	add	0x7f563412(%rsp),%r15
	nop
	add	0x7f563412(%rbp),%rax
	add	0x7f563412(%rbp),%rcx
	add	0x7f563412(%rbp),%rdx
	add	0x7f563412(%rbp),%rbx
	add	0x7f563412(%rbp),%rsp
	add	0x7f563412(%rbp),%rbp
	add	0x7f563412(%rbp),%rsi
	add	0x7f563412(%rbp),%rdi
	add	0x7f563412(%rbp),%r8
	add	0x7f563412(%rbp),%r9
	add	0x7f563412(%rbp),%r10
	add	0x7f563412(%rbp),%r11
	add	0x7f563412(%rbp),%r12
	add	0x7f563412(%rbp),%r13
	add	0x7f563412(%rbp),%r14
	add	0x7f563412(%rbp),%r15
	nop
	add	0x7f563412(%rsi),%rax
	add	0x7f563412(%rsi),%rcx
	add	0x7f563412(%rsi),%rdx
	add	0x7f563412(%rsi),%rbx
	add	0x7f563412(%rsi),%rsp
	add	0x7f563412(%rsi),%rbp
	add	0x7f563412(%rsi),%rsi
	add	0x7f563412(%rsi),%rdi
	add	0x7f563412(%rsi),%r8
	add	0x7f563412(%rsi),%r9
	add	0x7f563412(%rsi),%r10
	add	0x7f563412(%rsi),%r11
	add	0x7f563412(%rsi),%r12
	add	0x7f563412(%rsi),%r13
	add	0x7f563412(%rsi),%r14
	add	0x7f563412(%rsi),%r15
	nop
	add	0x7f563412(%rdi),%rax
	add	0x7f563412(%rdi),%rcx
	add	0x7f563412(%rdi),%rdx
	add	0x7f563412(%rdi),%rbx
	add	0x7f563412(%rdi),%rsp
	add	0x7f563412(%rdi),%rbp
	add	0x7f563412(%rdi),%rsi
	add	0x7f563412(%rdi),%rdi
	add	0x7f563412(%rdi),%r8
	add	0x7f563412(%rdi),%r9
	add	0x7f563412(%rdi),%r10
	add	0x7f563412(%rdi),%r11
	add	0x7f563412(%rdi),%r12
	add	0x7f563412(%rdi),%r13
	add	0x7f563412(%rdi),%r14
	add	0x7f563412(%rdi),%r15
	nop
	add	0x7f563412(%r8), %rax
	add	0x7f563412(%r8), %rcx
	add	0x7f563412(%r8), %rdx
	add	0x7f563412(%r8), %rbx
	add	0x7f563412(%r8), %rsp
	add	0x7f563412(%r8), %rbp
	add	0x7f563412(%r8), %rsi
	add	0x7f563412(%r8), %rdi
	add	0x7f563412(%r8), %r8
	add	0x7f563412(%r8), %r9
	add	0x7f563412(%r8), %r10
	add	0x7f563412(%r8), %r11
	add	0x7f563412(%r8), %r12
	add	0x7f563412(%r8), %r13
	add	0x7f563412(%r8), %r14
	add	0x7f563412(%r8), %r15
	nop
	add	0x7f563412(%r9), %rax
	add	0x7f563412(%r9), %rcx
	add	0x7f563412(%r9), %rdx
	add	0x7f563412(%r9), %rbx
	add	0x7f563412(%r9), %rsp
	add	0x7f563412(%r9), %rbp
	add	0x7f563412(%r9), %rsi
	add	0x7f563412(%r9), %rdi
	add	0x7f563412(%r9), %r8
	add	0x7f563412(%r9), %r9
	add	0x7f563412(%r9), %r10
	add	0x7f563412(%r9), %r11
	add	0x7f563412(%r9), %r12
	add	0x7f563412(%r9), %r13
	add	0x7f563412(%r9), %r14
	add	0x7f563412(%r9), %r15
	nop
	add	0x7f563412(%r10),%rax
	add	0x7f563412(%r10),%rcx
	add	0x7f563412(%r10),%rdx
	add	0x7f563412(%r10),%rbx
	add	0x7f563412(%r10),%rsp
	add	0x7f563412(%r10),%rbp
	add	0x7f563412(%r10),%rsi
	add	0x7f563412(%r10),%rdi
	add	0x7f563412(%r10),%r8
	add	0x7f563412(%r10),%r9
	add	0x7f563412(%r10),%r10
	add	0x7f563412(%r10),%r11
	add	0x7f563412(%r10),%r12
	add	0x7f563412(%r10),%r13
	add	0x7f563412(%r10),%r14
	add	0x7f563412(%r10),%r15
	nop
	add	0x7f563412(%r11),%rax
	add	0x7f563412(%r11),%rcx
	add	0x7f563412(%r11),%rdx
	add	0x7f563412(%r11),%rbx
	add	0x7f563412(%r11),%rsp
	add	0x7f563412(%r11),%rbp
	add	0x7f563412(%r11),%rsi
	add	0x7f563412(%r11),%rdi
	add	0x7f563412(%r11),%r8
	add	0x7f563412(%r11),%r9
	add	0x7f563412(%r11),%r10
	add	0x7f563412(%r11),%r11
	add	0x7f563412(%r11),%r12
	add	0x7f563412(%r11),%r13
	add	0x7f563412(%r11),%r14
	add	0x7f563412(%r11),%r15
	nop
	add	0x7f563412(%r12),%rax
	add	0x7f563412(%r12),%rcx
	add	0x7f563412(%r12),%rdx
	add	0x7f563412(%r12),%rbx
	add	0x7f563412(%r12),%rsp
	add	0x7f563412(%r12),%rbp
	add	0x7f563412(%r12),%rsi
	add	0x7f563412(%r12),%rdi
	add	0x7f563412(%r12),%r8
	add	0x7f563412(%r12),%r9
	add	0x7f563412(%r12),%r10
	add	0x7f563412(%r12),%r11
	add	0x7f563412(%r12),%r12
	add	0x7f563412(%r12),%r13
	add	0x7f563412(%r12),%r14
	add	0x7f563412(%r12),%r15
	nop
	add	0x7f563412(%r13),%rax
	add	0x7f563412(%r13),%rcx
	add	0x7f563412(%r13),%rdx
	add	0x7f563412(%r13),%rbx
	add	0x7f563412(%r13),%rsp
	add	0x7f563412(%r13),%rbp
	add	0x7f563412(%r13),%rsi
	add	0x7f563412(%r13),%rdi
	add	0x7f563412(%r13),%r8
	add	0x7f563412(%r13),%r9
	add	0x7f563412(%r13),%r10
	add	0x7f563412(%r13),%r11
	add	0x7f563412(%r13),%r12
	add	0x7f563412(%r13),%r13
	add	0x7f563412(%r13),%r14
	add	0x7f563412(%r13),%r15
	nop
	add	0x7f563412(%r14),%rax
	add	0x7f563412(%r14),%rcx
	add	0x7f563412(%r14),%rdx
	add	0x7f563412(%r14),%rbx
	add	0x7f563412(%r14),%rsp
	add	0x7f563412(%r14),%rbp
	add	0x7f563412(%r14),%rsi
	add	0x7f563412(%r14),%rdi
	add	0x7f563412(%r14),%r8
	add	0x7f563412(%r14),%r9
	add	0x7f563412(%r14),%r10
	add	0x7f563412(%r14),%r11
	add	0x7f563412(%r14),%r12
	add	0x7f563412(%r14),%r13
	add	0x7f563412(%r14),%r14
	add	0x7f563412(%r14),%r15
	nop
	add	0x7f563412(%r15),%rax
	add	0x7f563412(%r15),%rcx
	add	0x7f563412(%r15),%rdx
	add	0x7f563412(%r15),%rbx
	add	0x7f563412(%r15),%rsp
	add	0x7f563412(%r15),%rbp
	add	0x7f563412(%r15),%rsi
	add	0x7f563412(%r15),%rdi
	add	0x7f563412(%r15),%r8
	add	0x7f563412(%r15),%r9
	add	0x7f563412(%r15),%r10
	add	0x7f563412(%r15),%r11
	add	0x7f563412(%r15),%r12
	add	0x7f563412(%r15),%r13
	add	0x7f563412(%r15),%r14
	add	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


