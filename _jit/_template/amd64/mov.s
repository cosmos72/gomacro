	.file	"mov.s"
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
        movzx  %al,%cx
        movzx  %al,%rcx
        movzx  %ax,%rcx
        movsx  %al,%cx
        movsx  %al,%rcx
        movsx  %ax,%rcx
        movsx  %eax,%rcx
        ret
	.cfi_endproc



        .p2align 4,,15
	.globl	Mov
	.type	Mov, @function
Mov:
	.cfi_startproc
        // reg8 += reg8
	mov	%al,%al
	mov	%al,%cl
	mov	%al,%dl
	mov	%al,%bl
	mov	%al,%spl
	mov	%al,%bpl
	mov	%al,%sil
	mov	%al,%dil
	mov	%al,%r8b
	mov	%al,%r9b
	mov	%al,%r10b
	mov	%al,%r11b
	mov	%al,%r12b
	mov	%al,%r13b
	mov	%al,%r14b
	mov	%al,%r15b
	nop
	mov	%cl,%al
	mov	%cl,%cl
	mov	%cl,%dl
	mov	%cl,%bl
	mov	%cl,%spl
	mov	%cl,%bpl
	mov	%cl,%sil
	mov	%cl,%dil
	mov	%cl,%r8b
	mov	%cl,%r9b
	mov	%cl,%r10b
	mov	%cl,%r11b
	mov	%cl,%r12b
	mov	%cl,%r13b
	mov	%cl,%r14b
	mov	%cl,%r15b
	nop
	mov	%dl,%al
	mov	%dl,%cl
	mov	%dl,%dl
	mov	%dl,%bl
	mov	%dl,%spl
	mov	%dl,%bpl
	mov	%dl,%sil
	mov	%dl,%dil
	mov	%dl,%r8b
	mov	%dl,%r9b
	mov	%dl,%r10b
	mov	%dl,%r11b
	mov	%dl,%r12b
	mov	%dl,%r13b
	mov	%dl,%r14b
	mov	%dl,%r15b
	nop
	mov	%bl,%al
	mov	%bl,%cl
	mov	%bl,%dl
	mov	%bl,%bl
	mov	%bl,%spl
	mov	%bl,%bpl
	mov	%bl,%sil
	mov	%bl,%dil
	mov	%bl,%r8b
	mov	%bl,%r9b
	mov	%bl,%r10b
	mov	%bl,%r11b
	mov	%bl,%r12b
	mov	%bl,%r13b
	mov	%bl,%r14b
	mov	%bl,%r15b
	nop
	mov	%spl,%al
	mov	%spl,%cl
	mov	%spl,%dl
	mov	%spl,%bl
	mov	%spl,%spl
	mov	%spl,%bpl
	mov	%spl,%sil
	mov	%spl,%dil
	mov	%spl,%r8b
	mov	%spl,%r9b
	mov	%spl,%r10b
	mov	%spl,%r11b
	mov	%spl,%r12b
	mov	%spl,%r13b
	mov	%spl,%r14b
	mov	%spl,%r15b
	nop
	mov	%bpl,%al
	mov	%bpl,%cl
	mov	%bpl,%dl
	mov	%bpl,%bl
	mov	%bpl,%spl
	mov	%bpl,%bpl
	mov	%bpl,%sil
	mov	%bpl,%dil
	mov	%bpl,%r8b
	mov	%bpl,%r9b
	mov	%bpl,%r10b
	mov	%bpl,%r11b
	mov	%bpl,%r12b
	mov	%bpl,%r13b
	mov	%bpl,%r14b
	mov	%bpl,%r15b
	nop
	mov	%sil,%al
	mov	%sil,%cl
	mov	%sil,%dl
	mov	%sil,%bl
	mov	%sil,%spl
	mov	%sil,%bpl
	mov	%sil,%sil
	mov	%sil,%dil
	mov	%sil,%r8b
	mov	%sil,%r9b
	mov	%sil,%r10b
	mov	%sil,%r11b
	mov	%sil,%r12b
	mov	%sil,%r13b
	mov	%sil,%r14b
	mov	%sil,%r15b
	nop
	mov	%dil,%al
	mov	%dil,%cl
	mov	%dil,%dl
	mov	%dil,%bl
	mov	%dil,%spl
	mov	%dil,%bpl
	mov	%dil,%sil
	mov	%dil,%dil
	mov	%dil,%r8b
	mov	%dil,%r9b
	mov	%dil,%r10b
	mov	%dil,%r11b
	mov	%dil,%r12b
	mov	%dil,%r13b
	mov	%dil,%r14b
	mov	%dil,%r15b
	nop
	mov	%r8b, %al
	mov	%r8b, %cl
	mov	%r8b, %dl
	mov	%r8b, %bl
	mov	%r8b, %spl
	mov	%r8b, %bpl
	mov	%r8b, %sil
	mov	%r8b, %dil
	mov	%r8b, %r8b
	mov	%r8b, %r9b
	mov	%r8b, %r10b
	mov	%r8b, %r11b
	mov	%r8b, %r12b
	mov	%r8b, %r13b
	mov	%r8b, %r14b
	mov	%r8b, %r15b
	nop
	mov	%r9b, %al
	mov	%r9b, %cl
	mov	%r9b, %dl
	mov	%r9b, %bl
	mov	%r9b, %spl
	mov	%r9b, %bpl
	mov	%r9b, %sil
	mov	%r9b, %dil
	mov	%r9b, %r8b
	mov	%r9b, %r9b
	mov	%r9b, %r10b
	mov	%r9b, %r11b
	mov	%r9b, %r12b
	mov	%r9b, %r13b
	mov	%r9b, %r14b
	mov	%r9b, %r15b
	nop
	mov	%r10b,%al
	mov	%r10b,%cl
	mov	%r10b,%dl
	mov	%r10b,%bl
	mov	%r10b,%spl
	mov	%r10b,%bpl
	mov	%r10b,%sil
	mov	%r10b,%dil
	mov	%r10b,%r8b
	mov	%r10b,%r9b
	mov	%r10b,%r10b
	mov	%r10b,%r11b
	mov	%r10b,%r12b
	mov	%r10b,%r13b
	mov	%r10b,%r14b
	mov	%r10b,%r15b
	nop
	mov	%r11b,%al
	mov	%r11b,%cl
	mov	%r11b,%dl
	mov	%r11b,%bl
	mov	%r11b,%spl
	mov	%r11b,%bpl
	mov	%r11b,%sil
	mov	%r11b,%dil
	mov	%r11b,%r8b
	mov	%r11b,%r9b
	mov	%r11b,%r10b
	mov	%r11b,%r11b
	mov	%r11b,%r12b
	mov	%r11b,%r13b
	mov	%r11b,%r14b
	mov	%r11b,%r15b
	nop
	mov	%r12b,%al
	mov	%r12b,%cl
	mov	%r12b,%dl
	mov	%r12b,%bl
	mov	%r12b,%spl
	mov	%r12b,%bpl
	mov	%r12b,%sil
	mov	%r12b,%dil
	mov	%r12b,%r8b
	mov	%r12b,%r9b
	mov	%r12b,%r10b
	mov	%r12b,%r11b
	mov	%r12b,%r12b
	mov	%r12b,%r13b
	mov	%r12b,%r14b
	mov	%r12b,%r15b
	nop
	mov	%r13b,%al
	mov	%r13b,%cl
	mov	%r13b,%dl
	mov	%r13b,%bl
	mov	%r13b,%spl
	mov	%r13b,%bpl
	mov	%r13b,%sil
	mov	%r13b,%dil
	mov	%r13b,%r8b
	mov	%r13b,%r9b
	mov	%r13b,%r10b
	mov	%r13b,%r11b
	mov	%r13b,%r12b
	mov	%r13b,%r13b
	mov	%r13b,%r14b
	mov	%r13b,%r15b
	nop
	mov	%r14b,%al
	mov	%r14b,%cl
	mov	%r14b,%dl
	mov	%r14b,%bl
	mov	%r14b,%spl
	mov	%r14b,%bpl
	mov	%r14b,%sil
	mov	%r14b,%dil
	mov	%r14b,%r8b
	mov	%r14b,%r9b
	mov	%r14b,%r10b
	mov	%r14b,%r11b
	mov	%r14b,%r12b
	mov	%r14b,%r13b
	mov	%r14b,%r14b
	mov	%r14b,%r15b
	nop
	mov	%r15b,%al
	mov	%r15b,%cl
	mov	%r15b,%dl
	mov	%r15b,%bl
	mov	%r15b,%spl
	mov	%r15b,%bpl
	mov	%r15b,%sil
	mov	%r15b,%dil
	mov	%r15b,%r8b
	mov	%r15b,%r9b
	mov	%r15b,%r10b
	mov	%r15b,%r11b
	mov	%r15b,%r12b
	mov	%r15b,%r13b
	mov	%r15b,%r14b
	mov	%r15b,%r15b
        nop
        nop
        // reg16 += reg16
	mov	%ax,%ax
	mov	%ax,%cx
	mov	%ax,%dx
	mov	%ax,%bx
	mov	%ax,%sp
	mov	%ax,%bp
	mov	%ax,%si
	mov	%ax,%di
	mov	%ax,%r8w
	mov	%ax,%r9w
	mov	%ax,%r10w
	mov	%ax,%r11w
	mov	%ax,%r12w
	mov	%ax,%r13w
	mov	%ax,%r14w
	mov	%ax,%r15w
	nop
	mov	%cx,%ax
	mov	%cx,%cx
	mov	%cx,%dx
	mov	%cx,%bx
	mov	%cx,%sp
	mov	%cx,%bp
	mov	%cx,%si
	mov	%cx,%di
	mov	%cx,%r8w
	mov	%cx,%r9w
	mov	%cx,%r10w
	mov	%cx,%r11w
	mov	%cx,%r12w
	mov	%cx,%r13w
	mov	%cx,%r14w
	mov	%cx,%r15w
	nop
	mov	%dx,%ax
	mov	%dx,%cx
	mov	%dx,%dx
	mov	%dx,%bx
	mov	%dx,%sp
	mov	%dx,%bp
	mov	%dx,%si
	mov	%dx,%di
	mov	%dx,%r8w
	mov	%dx,%r9w
	mov	%dx,%r10w
	mov	%dx,%r11w
	mov	%dx,%r12w
	mov	%dx,%r13w
	mov	%dx,%r14w
	mov	%dx,%r15w
	nop
	mov	%bx,%ax
	mov	%bx,%cx
	mov	%bx,%dx
	mov	%bx,%bx
	mov	%bx,%sp
	mov	%bx,%bp
	mov	%bx,%si
	mov	%bx,%di
	mov	%bx,%r8w
	mov	%bx,%r9w
	mov	%bx,%r10w
	mov	%bx,%r11w
	mov	%bx,%r12w
	mov	%bx,%r13w
	mov	%bx,%r14w
	mov	%bx,%r15w
	nop
	mov	%sp,%ax
	mov	%sp,%cx
	mov	%sp,%dx
	mov	%sp,%bx
	mov	%sp,%sp
	mov	%sp,%bp
	mov	%sp,%si
	mov	%sp,%di
	mov	%sp,%r8w
	mov	%sp,%r9w
	mov	%sp,%r10w
	mov	%sp,%r11w
	mov	%sp,%r12w
	mov	%sp,%r13w
	mov	%sp,%r14w
	mov	%sp,%r15w
	nop
	mov	%bp,%ax
	mov	%bp,%cx
	mov	%bp,%dx
	mov	%bp,%bx
	mov	%bp,%sp
	mov	%bp,%bp
	mov	%bp,%si
	mov	%bp,%di
	mov	%bp,%r8w
	mov	%bp,%r9w
	mov	%bp,%r10w
	mov	%bp,%r11w
	mov	%bp,%r12w
	mov	%bp,%r13w
	mov	%bp,%r14w
	mov	%bp,%r15w
	nop
	mov	%si,%ax
	mov	%si,%cx
	mov	%si,%dx
	mov	%si,%bx
	mov	%si,%sp
	mov	%si,%bp
	mov	%si,%si
	mov	%si,%di
	mov	%si,%r8w
	mov	%si,%r9w
	mov	%si,%r10w
	mov	%si,%r11w
	mov	%si,%r12w
	mov	%si,%r13w
	mov	%si,%r14w
	mov	%si,%r15w
	nop
	mov	%di,%ax
	mov	%di,%cx
	mov	%di,%dx
	mov	%di,%bx
	mov	%di,%sp
	mov	%di,%bp
	mov	%di,%si
	mov	%di,%di
	mov	%di,%r8w
	mov	%di,%r9w
	mov	%di,%r10w
	mov	%di,%r11w
	mov	%di,%r12w
	mov	%di,%r13w
	mov	%di,%r14w
	mov	%di,%r15w
	nop
	mov	%r8w, %ax
	mov	%r8w, %cx
	mov	%r8w, %dx
	mov	%r8w, %bx
	mov	%r8w, %sp
	mov	%r8w, %bp
	mov	%r8w, %si
	mov	%r8w, %di
	mov	%r8w, %r8w
	mov	%r8w, %r9w
	mov	%r8w, %r10w
	mov	%r8w, %r11w
	mov	%r8w, %r12w
	mov	%r8w, %r13w
	mov	%r8w, %r14w
	mov	%r8w, %r15w
	nop
	mov	%r9w, %ax
	mov	%r9w, %cx
	mov	%r9w, %dx
	mov	%r9w, %bx
	mov	%r9w, %sp
	mov	%r9w, %bp
	mov	%r9w, %si
	mov	%r9w, %di
	mov	%r9w, %r8w
	mov	%r9w, %r9w
	mov	%r9w, %r10w
	mov	%r9w, %r11w
	mov	%r9w, %r12w
	mov	%r9w, %r13w
	mov	%r9w, %r14w
	mov	%r9w, %r15w
	nop
	mov	%r10w,%ax
	mov	%r10w,%cx
	mov	%r10w,%dx
	mov	%r10w,%bx
	mov	%r10w,%sp
	mov	%r10w,%bp
	mov	%r10w,%si
	mov	%r10w,%di
	mov	%r10w,%r8w
	mov	%r10w,%r9w
	mov	%r10w,%r10w
	mov	%r10w,%r11w
	mov	%r10w,%r12w
	mov	%r10w,%r13w
	mov	%r10w,%r14w
	mov	%r10w,%r15w
	nop
	mov	%r11w,%ax
	mov	%r11w,%cx
	mov	%r11w,%dx
	mov	%r11w,%bx
	mov	%r11w,%sp
	mov	%r11w,%bp
	mov	%r11w,%si
	mov	%r11w,%di
	mov	%r11w,%r8w
	mov	%r11w,%r9w
	mov	%r11w,%r10w
	mov	%r11w,%r11w
	mov	%r11w,%r12w
	mov	%r11w,%r13w
	mov	%r11w,%r14w
	mov	%r11w,%r15w
	nop
	mov	%r12w,%ax
	mov	%r12w,%cx
	mov	%r12w,%dx
	mov	%r12w,%bx
	mov	%r12w,%sp
	mov	%r12w,%bp
	mov	%r12w,%si
	mov	%r12w,%di
	mov	%r12w,%r8w
	mov	%r12w,%r9w
	mov	%r12w,%r10w
	mov	%r12w,%r11w
	mov	%r12w,%r12w
	mov	%r12w,%r13w
	mov	%r12w,%r14w
	mov	%r12w,%r15w
	nop
	mov	%r13w,%ax
	mov	%r13w,%cx
	mov	%r13w,%dx
	mov	%r13w,%bx
	mov	%r13w,%sp
	mov	%r13w,%bp
	mov	%r13w,%si
	mov	%r13w,%di
	mov	%r13w,%r8w
	mov	%r13w,%r9w
	mov	%r13w,%r10w
	mov	%r13w,%r11w
	mov	%r13w,%r12w
	mov	%r13w,%r13w
	mov	%r13w,%r14w
	mov	%r13w,%r15w
	nop
	mov	%r14w,%ax
	mov	%r14w,%cx
	mov	%r14w,%dx
	mov	%r14w,%bx
	mov	%r14w,%sp
	mov	%r14w,%bp
	mov	%r14w,%si
	mov	%r14w,%di
	mov	%r14w,%r8w
	mov	%r14w,%r9w
	mov	%r14w,%r10w
	mov	%r14w,%r11w
	mov	%r14w,%r12w
	mov	%r14w,%r13w
	mov	%r14w,%r14w
	mov	%r14w,%r15w
	nop
	mov	%r15w,%ax
	mov	%r15w,%cx
	mov	%r15w,%dx
	mov	%r15w,%bx
	mov	%r15w,%sp
	mov	%r15w,%bp
	mov	%r15w,%si
	mov	%r15w,%di
	mov	%r15w,%r8w
	mov	%r15w,%r9w
	mov	%r15w,%r10w
	mov	%r15w,%r11w
	mov	%r15w,%r12w
	mov	%r15w,%r13w
	mov	%r15w,%r14w
	mov	%r15w,%r15w
        nop
        nop
        // reg32 += reg32
	mov	%eax,%eax
	mov	%eax,%ecx
	mov	%eax,%edx
	mov	%eax,%ebx
	mov	%eax,%esp
	mov	%eax,%ebp
	mov	%eax,%esi
	mov	%eax,%edi
	mov	%eax,%r8d
	mov	%eax,%r9d
	mov	%eax,%r10d
	mov	%eax,%r11d
	mov	%eax,%r12d
	mov	%eax,%r13d
	mov	%eax,%r14d
	mov	%eax,%r15d
	nop
	mov	%ecx,%eax
	mov	%ecx,%ecx
	mov	%ecx,%edx
	mov	%ecx,%ebx
	mov	%ecx,%esp
	mov	%ecx,%ebp
	mov	%ecx,%esi
	mov	%ecx,%edi
	mov	%ecx,%r8d
	mov	%ecx,%r9d
	mov	%ecx,%r10d
	mov	%ecx,%r11d
	mov	%ecx,%r12d
	mov	%ecx,%r13d
	mov	%ecx,%r14d
	mov	%ecx,%r15d
	nop
	mov	%edx,%eax
	mov	%edx,%ecx
	mov	%edx,%edx
	mov	%edx,%ebx
	mov	%edx,%esp
	mov	%edx,%ebp
	mov	%edx,%esi
	mov	%edx,%edi
	mov	%edx,%r8d
	mov	%edx,%r9d
	mov	%edx,%r10d
	mov	%edx,%r11d
	mov	%edx,%r12d
	mov	%edx,%r13d
	mov	%edx,%r14d
	mov	%edx,%r15d
	nop
	mov	%ebx,%eax
	mov	%ebx,%ecx
	mov	%ebx,%edx
	mov	%ebx,%ebx
	mov	%ebx,%esp
	mov	%ebx,%ebp
	mov	%ebx,%esi
	mov	%ebx,%edi
	mov	%ebx,%r8d
	mov	%ebx,%r9d
	mov	%ebx,%r10d
	mov	%ebx,%r11d
	mov	%ebx,%r12d
	mov	%ebx,%r13d
	mov	%ebx,%r14d
	mov	%ebx,%r15d
	nop
	mov	%esp,%eax
	mov	%esp,%ecx
	mov	%esp,%edx
	mov	%esp,%ebx
	mov	%esp,%esp
	mov	%esp,%ebp
	mov	%esp,%esi
	mov	%esp,%edi
	mov	%esp,%r8d
	mov	%esp,%r9d
	mov	%esp,%r10d
	mov	%esp,%r11d
	mov	%esp,%r12d
	mov	%esp,%r13d
	mov	%esp,%r14d
	mov	%esp,%r15d
	nop
	mov	%ebp,%eax
	mov	%ebp,%ecx
	mov	%ebp,%edx
	mov	%ebp,%ebx
	mov	%ebp,%esp
	mov	%ebp,%ebp
	mov	%ebp,%esi
	mov	%ebp,%edi
	mov	%ebp,%r8d
	mov	%ebp,%r9d
	mov	%ebp,%r10d
	mov	%ebp,%r11d
	mov	%ebp,%r12d
	mov	%ebp,%r13d
	mov	%ebp,%r14d
	mov	%ebp,%r15d
	nop
	mov	%esi,%eax
	mov	%esi,%ecx
	mov	%esi,%edx
	mov	%esi,%ebx
	mov	%esi,%esp
	mov	%esi,%ebp
	mov	%esi,%esi
	mov	%esi,%edi
	mov	%esi,%r8d
	mov	%esi,%r9d
	mov	%esi,%r10d
	mov	%esi,%r11d
	mov	%esi,%r12d
	mov	%esi,%r13d
	mov	%esi,%r14d
	mov	%esi,%r15d
	nop
	mov	%edi,%eax
	mov	%edi,%ecx
	mov	%edi,%edx
	mov	%edi,%ebx
	mov	%edi,%esp
	mov	%edi,%ebp
	mov	%edi,%esi
	mov	%edi,%edi
	mov	%edi,%r8d
	mov	%edi,%r9d
	mov	%edi,%r10d
	mov	%edi,%r11d
	mov	%edi,%r12d
	mov	%edi,%r13d
	mov	%edi,%r14d
	mov	%edi,%r15d
	nop
	mov	%r8d, %eax
	mov	%r8d, %ecx
	mov	%r8d, %edx
	mov	%r8d, %ebx
	mov	%r8d, %esp
	mov	%r8d, %ebp
	mov	%r8d, %esi
	mov	%r8d, %edi
	mov	%r8d, %r8d
	mov	%r8d, %r9d
	mov	%r8d, %r10d
	mov	%r8d, %r11d
	mov	%r8d, %r12d
	mov	%r8d, %r13d
	mov	%r8d, %r14d
	mov	%r8d, %r15d
	nop
	mov	%r9d, %eax
	mov	%r9d, %ecx
	mov	%r9d, %edx
	mov	%r9d, %ebx
	mov	%r9d, %esp
	mov	%r9d, %ebp
	mov	%r9d, %esi
	mov	%r9d, %edi
	mov	%r9d, %r8d
	mov	%r9d, %r9d
	mov	%r9d, %r10d
	mov	%r9d, %r11d
	mov	%r9d, %r12d
	mov	%r9d, %r13d
	mov	%r9d, %r14d
	mov	%r9d, %r15d
	nop
	mov	%r10d,%eax
	mov	%r10d,%ecx
	mov	%r10d,%edx
	mov	%r10d,%ebx
	mov	%r10d,%esp
	mov	%r10d,%ebp
	mov	%r10d,%esi
	mov	%r10d,%edi
	mov	%r10d,%r8d
	mov	%r10d,%r9d
	mov	%r10d,%r10d
	mov	%r10d,%r11d
	mov	%r10d,%r12d
	mov	%r10d,%r13d
	mov	%r10d,%r14d
	mov	%r10d,%r15d
	nop
	mov	%r11d,%eax
	mov	%r11d,%ecx
	mov	%r11d,%edx
	mov	%r11d,%ebx
	mov	%r11d,%esp
	mov	%r11d,%ebp
	mov	%r11d,%esi
	mov	%r11d,%edi
	mov	%r11d,%r8d
	mov	%r11d,%r9d
	mov	%r11d,%r10d
	mov	%r11d,%r11d
	mov	%r11d,%r12d
	mov	%r11d,%r13d
	mov	%r11d,%r14d
	mov	%r11d,%r15d
	nop
	mov	%r12d,%eax
	mov	%r12d,%ecx
	mov	%r12d,%edx
	mov	%r12d,%ebx
	mov	%r12d,%esp
	mov	%r12d,%ebp
	mov	%r12d,%esi
	mov	%r12d,%edi
	mov	%r12d,%r8d
	mov	%r12d,%r9d
	mov	%r12d,%r10d
	mov	%r12d,%r11d
	mov	%r12d,%r12d
	mov	%r12d,%r13d
	mov	%r12d,%r14d
	mov	%r12d,%r15d
	nop
	mov	%r13d,%eax
	mov	%r13d,%ecx
	mov	%r13d,%edx
	mov	%r13d,%ebx
	mov	%r13d,%esp
	mov	%r13d,%ebp
	mov	%r13d,%esi
	mov	%r13d,%edi
	mov	%r13d,%r8d
	mov	%r13d,%r9d
	mov	%r13d,%r10d
	mov	%r13d,%r11d
	mov	%r13d,%r12d
	mov	%r13d,%r13d
	mov	%r13d,%r14d
	mov	%r13d,%r15d
	nop
	mov	%r14d,%eax
	mov	%r14d,%ecx
	mov	%r14d,%edx
	mov	%r14d,%ebx
	mov	%r14d,%esp
	mov	%r14d,%ebp
	mov	%r14d,%esi
	mov	%r14d,%edi
	mov	%r14d,%r8d
	mov	%r14d,%r9d
	mov	%r14d,%r10d
	mov	%r14d,%r11d
	mov	%r14d,%r12d
	mov	%r14d,%r13d
	mov	%r14d,%r14d
	mov	%r14d,%r15d
	nop
	mov	%r15d,%eax
	mov	%r15d,%ecx
	mov	%r15d,%edx
	mov	%r15d,%ebx
	mov	%r15d,%esp
	mov	%r15d,%ebp
	mov	%r15d,%esi
	mov	%r15d,%edi
	mov	%r15d,%r8d
	mov	%r15d,%r9d
	mov	%r15d,%r10d
	mov	%r15d,%r11d
	mov	%r15d,%r12d
	mov	%r15d,%r13d
	mov	%r15d,%r14d
	mov	%r15d,%r15d
        nop
        nop
        // reg64 += reg64
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
	.globl	MovMemReg8
	.type	MovMemReg8, @function
MovMemReg8:
	.cfi_startproc
        // mem8[0] += reg8
	mov	%al,(%rax)
	mov	%al,(%rcx)
	mov	%al,(%rdx)
	mov	%al,(%rbx)
	mov	%al,(%rsp)
	mov	%al,(%rbp)
	mov	%al,(%rsi)
	mov	%al,(%rdi)
	mov	%al,(%r8)
	mov	%al,(%r9)
	mov	%al,(%r10)
	mov	%al,(%r11)
	mov	%al,(%r12)
	mov	%al,(%r13)
	mov	%al,(%r14)
	mov	%al,(%r15)
	nop
	mov	%cl,(%rax)
	mov	%cl,(%rcx)
	mov	%cl,(%rdx)
	mov	%cl,(%rbx)
	mov	%cl,(%rsp)
	mov	%cl,(%rbp)
	mov	%cl,(%rsi)
	mov	%cl,(%rdi)
	mov	%cl,(%r8)
	mov	%cl,(%r9)
	mov	%cl,(%r10)
	mov	%cl,(%r11)
	mov	%cl,(%r12)
	mov	%cl,(%r13)
	mov	%cl,(%r14)
	mov	%cl,(%r15)
	nop
	mov	%dl,(%rax)
	mov	%dl,(%rcx)
	mov	%dl,(%rdx)
	mov	%dl,(%rbx)
	mov	%dl,(%rsp)
	mov	%dl,(%rbp)
	mov	%dl,(%rsi)
	mov	%dl,(%rdi)
	mov	%dl,(%r8)
	mov	%dl,(%r9)
	mov	%dl,(%r10)
	mov	%dl,(%r11)
	mov	%dl,(%r12)
	mov	%dl,(%r13)
	mov	%dl,(%r14)
	mov	%dl,(%r15)
	nop
	mov	%bl,(%rax)
	mov	%bl,(%rcx)
	mov	%bl,(%rdx)
	mov	%bl,(%rbx)
	mov	%bl,(%rsp)
	mov	%bl,(%rbp)
	mov	%bl,(%rsi)
	mov	%bl,(%rdi)
	mov	%bl,(%r8)
	mov	%bl,(%r9)
	mov	%bl,(%r10)
	mov	%bl,(%r11)
	mov	%bl,(%r12)
	mov	%bl,(%r13)
	mov	%bl,(%r14)
	mov	%bl,(%r15)
	nop
	mov	%spl,(%rax)
	mov	%spl,(%rcx)
	mov	%spl,(%rdx)
	mov	%spl,(%rbx)
	mov	%spl,(%rsp)
	mov	%spl,(%rbp)
	mov	%spl,(%rsi)
	mov	%spl,(%rdi)
	mov	%spl,(%r8)
	mov	%spl,(%r9)
	mov	%spl,(%r10)
	mov	%spl,(%r11)
	mov	%spl,(%r12)
	mov	%spl,(%r13)
	mov	%spl,(%r14)
	mov	%spl,(%r15)
	nop
	mov	%bpl,(%rax)
	mov	%bpl,(%rcx)
	mov	%bpl,(%rdx)
	mov	%bpl,(%rbx)
	mov	%bpl,(%rsp)
	mov	%bpl,(%rbp)
	mov	%bpl,(%rsi)
	mov	%bpl,(%rdi)
	mov	%bpl,(%r8)
	mov	%bpl,(%r9)
	mov	%bpl,(%r10)
	mov	%bpl,(%r11)
	mov	%bpl,(%r12)
	mov	%bpl,(%r13)
	mov	%bpl,(%r14)
	mov	%bpl,(%r15)
	nop
	mov	%sil,(%rax)
	mov	%sil,(%rcx)
	mov	%sil,(%rdx)
	mov	%sil,(%rbx)
	mov	%sil,(%rsp)
	mov	%sil,(%rbp)
	mov	%sil,(%rsi)
	mov	%sil,(%rdi)
	mov	%sil,(%r8)
	mov	%sil,(%r9)
	mov	%sil,(%r10)
	mov	%sil,(%r11)
	mov	%sil,(%r12)
	mov	%sil,(%r13)
	mov	%sil,(%r14)
	mov	%sil,(%r15)
	nop
	mov	%dil,(%rax)
	mov	%dil,(%rcx)
	mov	%dil,(%rdx)
	mov	%dil,(%rbx)
	mov	%dil,(%rsp)
	mov	%dil,(%rbp)
	mov	%dil,(%rsi)
	mov	%dil,(%rdi)
	mov	%dil,(%r8)
	mov	%dil,(%r9)
	mov	%dil,(%r10)
	mov	%dil,(%r11)
	mov	%dil,(%r12)
	mov	%dil,(%r13)
	mov	%dil,(%r14)
	mov	%dil,(%r15)
	nop
	mov	%r8b, (%rax)
	mov	%r8b, (%rcx)
	mov	%r8b, (%rdx)
	mov	%r8b, (%rbx)
	mov	%r8b, (%rsp)
	mov	%r8b, (%rbp)
	mov	%r8b, (%rsi)
	mov	%r8b, (%rdi)
	mov	%r8b, (%r8)
	mov	%r8b, (%r9)
	mov	%r8b, (%r10)
	mov	%r8b, (%r11)
	mov	%r8b, (%r12)
	mov	%r8b, (%r13)
	mov	%r8b, (%r14)
	mov	%r8b, (%r15)
	nop
	mov	%r9b, (%rax)
	mov	%r9b, (%rcx)
	mov	%r9b, (%rdx)
	mov	%r9b, (%rbx)
	mov	%r9b, (%rsp)
	mov	%r9b, (%rbp)
	mov	%r9b, (%rsi)
	mov	%r9b, (%rdi)
	mov	%r9b, (%r8)
	mov	%r9b, (%r9)
	mov	%r9b, (%r10)
	mov	%r9b, (%r11)
	mov	%r9b, (%r12)
	mov	%r9b, (%r13)
	mov	%r9b, (%r14)
	mov	%r9b, (%r15)
	nop
	mov	%r10b,(%rax)
	mov	%r10b,(%rcx)
	mov	%r10b,(%rdx)
	mov	%r10b,(%rbx)
	mov	%r10b,(%rsp)
	mov	%r10b,(%rbp)
	mov	%r10b,(%rsi)
	mov	%r10b,(%rdi)
	mov	%r10b,(%r8)
	mov	%r10b,(%r9)
	mov	%r10b,(%r10)
	mov	%r10b,(%r11)
	mov	%r10b,(%r12)
	mov	%r10b,(%r13)
	mov	%r10b,(%r14)
	mov	%r10b,(%r15)
	nop
	mov	%r11b,(%rax)
	mov	%r11b,(%rcx)
	mov	%r11b,(%rdx)
	mov	%r11b,(%rbx)
	mov	%r11b,(%rsp)
	mov	%r11b,(%rbp)
	mov	%r11b,(%rsi)
	mov	%r11b,(%rdi)
	mov	%r11b,(%r8)
	mov	%r11b,(%r9)
	mov	%r11b,(%r10)
	mov	%r11b,(%r11)
	mov	%r11b,(%r12)
	mov	%r11b,(%r13)
	mov	%r11b,(%r14)
	mov	%r11b,(%r15)
	nop
	mov	%r12b,(%rax)
	mov	%r12b,(%rcx)
	mov	%r12b,(%rdx)
	mov	%r12b,(%rbx)
	mov	%r12b,(%rsp)
	mov	%r12b,(%rbp)
	mov	%r12b,(%rsi)
	mov	%r12b,(%rdi)
	mov	%r12b,(%r8)
	mov	%r12b,(%r9)
	mov	%r12b,(%r10)
	mov	%r12b,(%r11)
	mov	%r12b,(%r12)
	mov	%r12b,(%r13)
	mov	%r12b,(%r14)
	mov	%r12b,(%r15)
	nop
	mov	%r13b,(%rax)
	mov	%r13b,(%rcx)
	mov	%r13b,(%rdx)
	mov	%r13b,(%rbx)
	mov	%r13b,(%rsp)
	mov	%r13b,(%rbp)
	mov	%r13b,(%rsi)
	mov	%r13b,(%rdi)
	mov	%r13b,(%r8)
	mov	%r13b,(%r9)
	mov	%r13b,(%r10)
	mov	%r13b,(%r11)
	mov	%r13b,(%r12)
	mov	%r13b,(%r13)
	mov	%r13b,(%r14)
	mov	%r13b,(%r15)
	nop
	mov	%r14b,(%rax)
	mov	%r14b,(%rcx)
	mov	%r14b,(%rdx)
	mov	%r14b,(%rbx)
	mov	%r14b,(%rsp)
	mov	%r14b,(%rbp)
	mov	%r14b,(%rsi)
	mov	%r14b,(%rdi)
	mov	%r14b,(%r8)
	mov	%r14b,(%r9)
	mov	%r14b,(%r10)
	mov	%r14b,(%r11)
	mov	%r14b,(%r12)
	mov	%r14b,(%r13)
	mov	%r14b,(%r14)
	mov	%r14b,(%r15)
	nop
	mov	%r15b,(%rax)
	mov	%r15b,(%rcx)
	mov	%r15b,(%rdx)
	mov	%r15b,(%rbx)
	mov	%r15b,(%rsp)
	mov	%r15b,(%rbp)
	mov	%r15b,(%rsi)
	mov	%r15b,(%rdi)
	mov	%r15b,(%r8)
	mov	%r15b,(%r9)
	mov	%r15b,(%r10)
	mov	%r15b,(%r11)
	mov	%r15b,(%r12)
	mov	%r15b,(%r13)
	mov	%r15b,(%r14)
	mov	%r15b,(%r15)
        nop
        nop
        // mem8[off8] += reg8
	mov	%al,0x7F(%rax)
	mov	%al,0x7F(%rcx)
	mov	%al,0x7F(%rdx)
	mov	%al,0x7F(%rbx)
	mov	%al,0x7F(%rsp)
	mov	%al,0x7F(%rbp)
	mov	%al,0x7F(%rsi)
	mov	%al,0x7F(%rdi)
	mov	%al,0x7F(%r8)
	mov	%al,0x7F(%r9)
	mov	%al,0x7F(%r10)
	mov	%al,0x7F(%r11)
	mov	%al,0x7F(%r12)
	mov	%al,0x7F(%r13)
	mov	%al,0x7F(%r14)
	mov	%al,0x7F(%r15)
	nop
	mov	%cl,0x7F(%rax)
	mov	%cl,0x7F(%rcx)
	mov	%cl,0x7F(%rdx)
	mov	%cl,0x7F(%rbx)
	mov	%cl,0x7F(%rsp)
	mov	%cl,0x7F(%rbp)
	mov	%cl,0x7F(%rsi)
	mov	%cl,0x7F(%rdi)
	mov	%cl,0x7F(%r8)
	mov	%cl,0x7F(%r9)
	mov	%cl,0x7F(%r10)
	mov	%cl,0x7F(%r11)
	mov	%cl,0x7F(%r12)
	mov	%cl,0x7F(%r13)
	mov	%cl,0x7F(%r14)
	mov	%cl,0x7F(%r15)
	nop
	mov	%dl,0x7F(%rax)
	mov	%dl,0x7F(%rcx)
	mov	%dl,0x7F(%rdx)
	mov	%dl,0x7F(%rbx)
	mov	%dl,0x7F(%rsp)
	mov	%dl,0x7F(%rbp)
	mov	%dl,0x7F(%rsi)
	mov	%dl,0x7F(%rdi)
	mov	%dl,0x7F(%r8)
	mov	%dl,0x7F(%r9)
	mov	%dl,0x7F(%r10)
	mov	%dl,0x7F(%r11)
	mov	%dl,0x7F(%r12)
	mov	%dl,0x7F(%r13)
	mov	%dl,0x7F(%r14)
	mov	%dl,0x7F(%r15)
	nop
	mov	%bl,0x7F(%rax)
	mov	%bl,0x7F(%rcx)
	mov	%bl,0x7F(%rdx)
	mov	%bl,0x7F(%rbx)
	mov	%bl,0x7F(%rsp)
	mov	%bl,0x7F(%rbp)
	mov	%bl,0x7F(%rsi)
	mov	%bl,0x7F(%rdi)
	mov	%bl,0x7F(%r8)
	mov	%bl,0x7F(%r9)
	mov	%bl,0x7F(%r10)
	mov	%bl,0x7F(%r11)
	mov	%bl,0x7F(%r12)
	mov	%bl,0x7F(%r13)
	mov	%bl,0x7F(%r14)
	mov	%bl,0x7F(%r15)
	nop
	mov	%spl,0x7F(%rax)
	mov	%spl,0x7F(%rcx)
	mov	%spl,0x7F(%rdx)
	mov	%spl,0x7F(%rbx)
	mov	%spl,0x7F(%rsp)
	mov	%spl,0x7F(%rbp)
	mov	%spl,0x7F(%rsi)
	mov	%spl,0x7F(%rdi)
	mov	%spl,0x7F(%r8)
	mov	%spl,0x7F(%r9)
	mov	%spl,0x7F(%r10)
	mov	%spl,0x7F(%r11)
	mov	%spl,0x7F(%r12)
	mov	%spl,0x7F(%r13)
	mov	%spl,0x7F(%r14)
	mov	%spl,0x7F(%r15)
	nop
	mov	%bpl,0x7F(%rax)
	mov	%bpl,0x7F(%rcx)
	mov	%bpl,0x7F(%rdx)
	mov	%bpl,0x7F(%rbx)
	mov	%bpl,0x7F(%rsp)
	mov	%bpl,0x7F(%rbp)
	mov	%bpl,0x7F(%rsi)
	mov	%bpl,0x7F(%rdi)
	mov	%bpl,0x7F(%r8)
	mov	%bpl,0x7F(%r9)
	mov	%bpl,0x7F(%r10)
	mov	%bpl,0x7F(%r11)
	mov	%bpl,0x7F(%r12)
	mov	%bpl,0x7F(%r13)
	mov	%bpl,0x7F(%r14)
	mov	%bpl,0x7F(%r15)
	nop
	mov	%sil,0x7F(%rax)
	mov	%sil,0x7F(%rcx)
	mov	%sil,0x7F(%rdx)
	mov	%sil,0x7F(%rbx)
	mov	%sil,0x7F(%rsp)
	mov	%sil,0x7F(%rbp)
	mov	%sil,0x7F(%rsi)
	mov	%sil,0x7F(%rdi)
	mov	%sil,0x7F(%r8)
	mov	%sil,0x7F(%r9)
	mov	%sil,0x7F(%r10)
	mov	%sil,0x7F(%r11)
	mov	%sil,0x7F(%r12)
	mov	%sil,0x7F(%r13)
	mov	%sil,0x7F(%r14)
	mov	%sil,0x7F(%r15)
	nop
	mov	%dil,0x7F(%rax)
	mov	%dil,0x7F(%rcx)
	mov	%dil,0x7F(%rdx)
	mov	%dil,0x7F(%rbx)
	mov	%dil,0x7F(%rsp)
	mov	%dil,0x7F(%rbp)
	mov	%dil,0x7F(%rsi)
	mov	%dil,0x7F(%rdi)
	mov	%dil,0x7F(%r8)
	mov	%dil,0x7F(%r9)
	mov	%dil,0x7F(%r10)
	mov	%dil,0x7F(%r11)
	mov	%dil,0x7F(%r12)
	mov	%dil,0x7F(%r13)
	mov	%dil,0x7F(%r14)
	mov	%dil,0x7F(%r15)
	nop
	mov	%r8b, 0x7F(%rax)
	mov	%r8b, 0x7F(%rcx)
	mov	%r8b, 0x7F(%rdx)
	mov	%r8b, 0x7F(%rbx)
	mov	%r8b, 0x7F(%rsp)
	mov	%r8b, 0x7F(%rbp)
	mov	%r8b, 0x7F(%rsi)
	mov	%r8b, 0x7F(%rdi)
	mov	%r8b, 0x7F(%r8)
	mov	%r8b, 0x7F(%r9)
	mov	%r8b, 0x7F(%r10)
	mov	%r8b, 0x7F(%r11)
	mov	%r8b, 0x7F(%r12)
	mov	%r8b, 0x7F(%r13)
	mov	%r8b, 0x7F(%r14)
	mov	%r8b, 0x7F(%r15)
	nop
	mov	%r9b, 0x7F(%rax)
	mov	%r9b, 0x7F(%rcx)
	mov	%r9b, 0x7F(%rdx)
	mov	%r9b, 0x7F(%rbx)
	mov	%r9b, 0x7F(%rsp)
	mov	%r9b, 0x7F(%rbp)
	mov	%r9b, 0x7F(%rsi)
	mov	%r9b, 0x7F(%rdi)
	mov	%r9b, 0x7F(%r8)
	mov	%r9b, 0x7F(%r9)
	mov	%r9b, 0x7F(%r10)
	mov	%r9b, 0x7F(%r11)
	mov	%r9b, 0x7F(%r12)
	mov	%r9b, 0x7F(%r13)
	mov	%r9b, 0x7F(%r14)
	mov	%r9b, 0x7F(%r15)
	nop
	mov	%r10b,0x7F(%rax)
	mov	%r10b,0x7F(%rcx)
	mov	%r10b,0x7F(%rdx)
	mov	%r10b,0x7F(%rbx)
	mov	%r10b,0x7F(%rsp)
	mov	%r10b,0x7F(%rbp)
	mov	%r10b,0x7F(%rsi)
	mov	%r10b,0x7F(%rdi)
	mov	%r10b,0x7F(%r8)
	mov	%r10b,0x7F(%r9)
	mov	%r10b,0x7F(%r10)
	mov	%r10b,0x7F(%r11)
	mov	%r10b,0x7F(%r12)
	mov	%r10b,0x7F(%r13)
	mov	%r10b,0x7F(%r14)
	mov	%r10b,0x7F(%r15)
	nop
	mov	%r11b,0x7F(%rax)
	mov	%r11b,0x7F(%rcx)
	mov	%r11b,0x7F(%rdx)
	mov	%r11b,0x7F(%rbx)
	mov	%r11b,0x7F(%rsp)
	mov	%r11b,0x7F(%rbp)
	mov	%r11b,0x7F(%rsi)
	mov	%r11b,0x7F(%rdi)
	mov	%r11b,0x7F(%r8)
	mov	%r11b,0x7F(%r9)
	mov	%r11b,0x7F(%r10)
	mov	%r11b,0x7F(%r11)
	mov	%r11b,0x7F(%r12)
	mov	%r11b,0x7F(%r13)
	mov	%r11b,0x7F(%r14)
	mov	%r11b,0x7F(%r15)
	nop
	mov	%r12b,0x7F(%rax)
	mov	%r12b,0x7F(%rcx)
	mov	%r12b,0x7F(%rdx)
	mov	%r12b,0x7F(%rbx)
	mov	%r12b,0x7F(%rsp)
	mov	%r12b,0x7F(%rbp)
	mov	%r12b,0x7F(%rsi)
	mov	%r12b,0x7F(%rdi)
	mov	%r12b,0x7F(%r8)
	mov	%r12b,0x7F(%r9)
	mov	%r12b,0x7F(%r10)
	mov	%r12b,0x7F(%r11)
	mov	%r12b,0x7F(%r12)
	mov	%r12b,0x7F(%r13)
	mov	%r12b,0x7F(%r14)
	mov	%r12b,0x7F(%r15)
	nop
	mov	%r13b,0x7F(%rax)
	mov	%r13b,0x7F(%rcx)
	mov	%r13b,0x7F(%rdx)
	mov	%r13b,0x7F(%rbx)
	mov	%r13b,0x7F(%rsp)
	mov	%r13b,0x7F(%rbp)
	mov	%r13b,0x7F(%rsi)
	mov	%r13b,0x7F(%rdi)
	mov	%r13b,0x7F(%r8)
	mov	%r13b,0x7F(%r9)
	mov	%r13b,0x7F(%r10)
	mov	%r13b,0x7F(%r11)
	mov	%r13b,0x7F(%r12)
	mov	%r13b,0x7F(%r13)
	mov	%r13b,0x7F(%r14)
	mov	%r13b,0x7F(%r15)
	nop
	mov	%r14b,0x7F(%rax)
	mov	%r14b,0x7F(%rcx)
	mov	%r14b,0x7F(%rdx)
	mov	%r14b,0x7F(%rbx)
	mov	%r14b,0x7F(%rsp)
	mov	%r14b,0x7F(%rbp)
	mov	%r14b,0x7F(%rsi)
	mov	%r14b,0x7F(%rdi)
	mov	%r14b,0x7F(%r8)
	mov	%r14b,0x7F(%r9)
	mov	%r14b,0x7F(%r10)
	mov	%r14b,0x7F(%r11)
	mov	%r14b,0x7F(%r12)
	mov	%r14b,0x7F(%r13)
	mov	%r14b,0x7F(%r14)
	mov	%r14b,0x7F(%r15)
	nop
	mov	%r15b,0x7F(%rax)
	mov	%r15b,0x7F(%rcx)
	mov	%r15b,0x7F(%rdx)
	mov	%r15b,0x7F(%rbx)
	mov	%r15b,0x7F(%rsp)
	mov	%r15b,0x7F(%rbp)
	mov	%r15b,0x7F(%rsi)
	mov	%r15b,0x7F(%rdi)
	mov	%r15b,0x7F(%r8)
	mov	%r15b,0x7F(%r9)
	mov	%r15b,0x7F(%r10)
	mov	%r15b,0x7F(%r11)
	mov	%r15b,0x7F(%r12)
	mov	%r15b,0x7F(%r13)
	mov	%r15b,0x7F(%r14)
	mov	%r15b,0x7F(%r15)
        nop
        nop
        // mem8[off32] += reg8
	mov	%al,0x12345678(%rax)
	mov	%al,0x12345678(%rcx)
	mov	%al,0x12345678(%rdx)
	mov	%al,0x12345678(%rbx)
	mov	%al,0x12345678(%rsp)
	mov	%al,0x12345678(%rbp)
	mov	%al,0x12345678(%rsi)
	mov	%al,0x12345678(%rdi)
	mov	%al,0x12345678(%r8)
	mov	%al,0x12345678(%r9)
	mov	%al,0x12345678(%r10)
	mov	%al,0x12345678(%r11)
	mov	%al,0x12345678(%r12)
	mov	%al,0x12345678(%r13)
	mov	%al,0x12345678(%r14)
	mov	%al,0x12345678(%r15)
	nop
	mov	%cl,0x12345678(%rax)
	mov	%cl,0x12345678(%rcx)
	mov	%cl,0x12345678(%rdx)
	mov	%cl,0x12345678(%rbx)
	mov	%cl,0x12345678(%rsp)
	mov	%cl,0x12345678(%rbp)
	mov	%cl,0x12345678(%rsi)
	mov	%cl,0x12345678(%rdi)
	mov	%cl,0x12345678(%r8)
	mov	%cl,0x12345678(%r9)
	mov	%cl,0x12345678(%r10)
	mov	%cl,0x12345678(%r11)
	mov	%cl,0x12345678(%r12)
	mov	%cl,0x12345678(%r13)
	mov	%cl,0x12345678(%r14)
	mov	%cl,0x12345678(%r15)
	nop
	mov	%dl,0x12345678(%rax)
	mov	%dl,0x12345678(%rcx)
	mov	%dl,0x12345678(%rdx)
	mov	%dl,0x12345678(%rbx)
	mov	%dl,0x12345678(%rsp)
	mov	%dl,0x12345678(%rbp)
	mov	%dl,0x12345678(%rsi)
	mov	%dl,0x12345678(%rdi)
	mov	%dl,0x12345678(%r8)
	mov	%dl,0x12345678(%r9)
	mov	%dl,0x12345678(%r10)
	mov	%dl,0x12345678(%r11)
	mov	%dl,0x12345678(%r12)
	mov	%dl,0x12345678(%r13)
	mov	%dl,0x12345678(%r14)
	mov	%dl,0x12345678(%r15)
	nop
	mov	%bl,0x12345678(%rax)
	mov	%bl,0x12345678(%rcx)
	mov	%bl,0x12345678(%rdx)
	mov	%bl,0x12345678(%rbx)
	mov	%bl,0x12345678(%rsp)
	mov	%bl,0x12345678(%rbp)
	mov	%bl,0x12345678(%rsi)
	mov	%bl,0x12345678(%rdi)
	mov	%bl,0x12345678(%r8)
	mov	%bl,0x12345678(%r9)
	mov	%bl,0x12345678(%r10)
	mov	%bl,0x12345678(%r11)
	mov	%bl,0x12345678(%r12)
	mov	%bl,0x12345678(%r13)
	mov	%bl,0x12345678(%r14)
	mov	%bl,0x12345678(%r15)
	nop
	mov	%spl,0x12345678(%rax)
	mov	%spl,0x12345678(%rcx)
	mov	%spl,0x12345678(%rdx)
	mov	%spl,0x12345678(%rbx)
	mov	%spl,0x12345678(%rsp)
	mov	%spl,0x12345678(%rbp)
	mov	%spl,0x12345678(%rsi)
	mov	%spl,0x12345678(%rdi)
	mov	%spl,0x12345678(%r8)
	mov	%spl,0x12345678(%r9)
	mov	%spl,0x12345678(%r10)
	mov	%spl,0x12345678(%r11)
	mov	%spl,0x12345678(%r12)
	mov	%spl,0x12345678(%r13)
	mov	%spl,0x12345678(%r14)
	mov	%spl,0x12345678(%r15)
	nop
	mov	%bpl,0x12345678(%rax)
	mov	%bpl,0x12345678(%rcx)
	mov	%bpl,0x12345678(%rdx)
	mov	%bpl,0x12345678(%rbx)
	mov	%bpl,0x12345678(%rsp)
	mov	%bpl,0x12345678(%rbp)
	mov	%bpl,0x12345678(%rsi)
	mov	%bpl,0x12345678(%rdi)
	mov	%bpl,0x12345678(%r8)
	mov	%bpl,0x12345678(%r9)
	mov	%bpl,0x12345678(%r10)
	mov	%bpl,0x12345678(%r11)
	mov	%bpl,0x12345678(%r12)
	mov	%bpl,0x12345678(%r13)
	mov	%bpl,0x12345678(%r14)
	mov	%bpl,0x12345678(%r15)
	nop
	mov	%sil,0x12345678(%rax)
	mov	%sil,0x12345678(%rcx)
	mov	%sil,0x12345678(%rdx)
	mov	%sil,0x12345678(%rbx)
	mov	%sil,0x12345678(%rsp)
	mov	%sil,0x12345678(%rbp)
	mov	%sil,0x12345678(%rsi)
	mov	%sil,0x12345678(%rdi)
	mov	%sil,0x12345678(%r8)
	mov	%sil,0x12345678(%r9)
	mov	%sil,0x12345678(%r10)
	mov	%sil,0x12345678(%r11)
	mov	%sil,0x12345678(%r12)
	mov	%sil,0x12345678(%r13)
	mov	%sil,0x12345678(%r14)
	mov	%sil,0x12345678(%r15)
	nop
	mov	%dil,0x12345678(%rax)
	mov	%dil,0x12345678(%rcx)
	mov	%dil,0x12345678(%rdx)
	mov	%dil,0x12345678(%rbx)
	mov	%dil,0x12345678(%rsp)
	mov	%dil,0x12345678(%rbp)
	mov	%dil,0x12345678(%rsi)
	mov	%dil,0x12345678(%rdi)
	mov	%dil,0x12345678(%r8)
	mov	%dil,0x12345678(%r9)
	mov	%dil,0x12345678(%r10)
	mov	%dil,0x12345678(%r11)
	mov	%dil,0x12345678(%r12)
	mov	%dil,0x12345678(%r13)
	mov	%dil,0x12345678(%r14)
	mov	%dil,0x12345678(%r15)
	nop
	mov	%r8b, 0x12345678(%rax)
	mov	%r8b, 0x12345678(%rcx)
	mov	%r8b, 0x12345678(%rdx)
	mov	%r8b, 0x12345678(%rbx)
	mov	%r8b, 0x12345678(%rsp)
	mov	%r8b, 0x12345678(%rbp)
	mov	%r8b, 0x12345678(%rsi)
	mov	%r8b, 0x12345678(%rdi)
	mov	%r8b, 0x12345678(%r8)
	mov	%r8b, 0x12345678(%r9)
	mov	%r8b, 0x12345678(%r10)
	mov	%r8b, 0x12345678(%r11)
	mov	%r8b, 0x12345678(%r12)
	mov	%r8b, 0x12345678(%r13)
	mov	%r8b, 0x12345678(%r14)
	mov	%r8b, 0x12345678(%r15)
	nop
	mov	%r9b, 0x12345678(%rax)
	mov	%r9b, 0x12345678(%rcx)
	mov	%r9b, 0x12345678(%rdx)
	mov	%r9b, 0x12345678(%rbx)
	mov	%r9b, 0x12345678(%rsp)
	mov	%r9b, 0x12345678(%rbp)
	mov	%r9b, 0x12345678(%rsi)
	mov	%r9b, 0x12345678(%rdi)
	mov	%r9b, 0x12345678(%r8)
	mov	%r9b, 0x12345678(%r9)
	mov	%r9b, 0x12345678(%r10)
	mov	%r9b, 0x12345678(%r11)
	mov	%r9b, 0x12345678(%r12)
	mov	%r9b, 0x12345678(%r13)
	mov	%r9b, 0x12345678(%r14)
	mov	%r9b, 0x12345678(%r15)
	nop
	mov	%r10b,0x12345678(%rax)
	mov	%r10b,0x12345678(%rcx)
	mov	%r10b,0x12345678(%rdx)
	mov	%r10b,0x12345678(%rbx)
	mov	%r10b,0x12345678(%rsp)
	mov	%r10b,0x12345678(%rbp)
	mov	%r10b,0x12345678(%rsi)
	mov	%r10b,0x12345678(%rdi)
	mov	%r10b,0x12345678(%r8)
	mov	%r10b,0x12345678(%r9)
	mov	%r10b,0x12345678(%r10)
	mov	%r10b,0x12345678(%r11)
	mov	%r10b,0x12345678(%r12)
	mov	%r10b,0x12345678(%r13)
	mov	%r10b,0x12345678(%r14)
	mov	%r10b,0x12345678(%r15)
	nop
	mov	%r11b,0x12345678(%rax)
	mov	%r11b,0x12345678(%rcx)
	mov	%r11b,0x12345678(%rdx)
	mov	%r11b,0x12345678(%rbx)
	mov	%r11b,0x12345678(%rsp)
	mov	%r11b,0x12345678(%rbp)
	mov	%r11b,0x12345678(%rsi)
	mov	%r11b,0x12345678(%rdi)
	mov	%r11b,0x12345678(%r8)
	mov	%r11b,0x12345678(%r9)
	mov	%r11b,0x12345678(%r10)
	mov	%r11b,0x12345678(%r11)
	mov	%r11b,0x12345678(%r12)
	mov	%r11b,0x12345678(%r13)
	mov	%r11b,0x12345678(%r14)
	mov	%r11b,0x12345678(%r15)
	nop
	mov	%r12b,0x12345678(%rax)
	mov	%r12b,0x12345678(%rcx)
	mov	%r12b,0x12345678(%rdx)
	mov	%r12b,0x12345678(%rbx)
	mov	%r12b,0x12345678(%rsp)
	mov	%r12b,0x12345678(%rbp)
	mov	%r12b,0x12345678(%rsi)
	mov	%r12b,0x12345678(%rdi)
	mov	%r12b,0x12345678(%r8)
	mov	%r12b,0x12345678(%r9)
	mov	%r12b,0x12345678(%r10)
	mov	%r12b,0x12345678(%r11)
	mov	%r12b,0x12345678(%r12)
	mov	%r12b,0x12345678(%r13)
	mov	%r12b,0x12345678(%r14)
	mov	%r12b,0x12345678(%r15)
	nop
	mov	%r13b,0x12345678(%rax)
	mov	%r13b,0x12345678(%rcx)
	mov	%r13b,0x12345678(%rdx)
	mov	%r13b,0x12345678(%rbx)
	mov	%r13b,0x12345678(%rsp)
	mov	%r13b,0x12345678(%rbp)
	mov	%r13b,0x12345678(%rsi)
	mov	%r13b,0x12345678(%rdi)
	mov	%r13b,0x12345678(%r8)
	mov	%r13b,0x12345678(%r9)
	mov	%r13b,0x12345678(%r10)
	mov	%r13b,0x12345678(%r11)
	mov	%r13b,0x12345678(%r12)
	mov	%r13b,0x12345678(%r13)
	mov	%r13b,0x12345678(%r14)
	mov	%r13b,0x12345678(%r15)
	nop
	mov	%r14b,0x12345678(%rax)
	mov	%r14b,0x12345678(%rcx)
	mov	%r14b,0x12345678(%rdx)
	mov	%r14b,0x12345678(%rbx)
	mov	%r14b,0x12345678(%rsp)
	mov	%r14b,0x12345678(%rbp)
	mov	%r14b,0x12345678(%rsi)
	mov	%r14b,0x12345678(%rdi)
	mov	%r14b,0x12345678(%r8)
	mov	%r14b,0x12345678(%r9)
	mov	%r14b,0x12345678(%r10)
	mov	%r14b,0x12345678(%r11)
	mov	%r14b,0x12345678(%r12)
	mov	%r14b,0x12345678(%r13)
	mov	%r14b,0x12345678(%r14)
	mov	%r14b,0x12345678(%r15)
	nop
	mov	%r15b,0x12345678(%rax)
	mov	%r15b,0x12345678(%rcx)
	mov	%r15b,0x12345678(%rdx)
	mov	%r15b,0x12345678(%rbx)
	mov	%r15b,0x12345678(%rsp)
	mov	%r15b,0x12345678(%rbp)
	mov	%r15b,0x12345678(%rsi)
	mov	%r15b,0x12345678(%rdi)
	mov	%r15b,0x12345678(%r8)
	mov	%r15b,0x12345678(%r9)
	mov	%r15b,0x12345678(%r10)
	mov	%r15b,0x12345678(%r11)
	mov	%r15b,0x12345678(%r12)
	mov	%r15b,0x12345678(%r13)
	mov	%r15b,0x12345678(%r14)
	mov	%r15b,0x12345678(%r15)
        ret
	.cfi_endproc

	.p2align 4,,15
	.globl	MovMemReg16
	.type	MovMemReg16, @function
MovMemReg16:
	.cfi_startproc
        // mem16[0] += reg16
	mov	%ax,(%rax)
	mov	%ax,(%rcx)
	mov	%ax,(%rdx)
	mov	%ax,(%rbx)
	mov	%ax,(%rsp)
	mov	%ax,(%rbp)
	mov	%ax,(%rsi)
	mov	%ax,(%rdi)
	mov	%ax,(%r8)
	mov	%ax,(%r9)
	mov	%ax,(%r10)
	mov	%ax,(%r11)
	mov	%ax,(%r12)
	mov	%ax,(%r13)
	mov	%ax,(%r14)
	mov	%ax,(%r15)
	nop
	mov	%cx,(%rax)
	mov	%cx,(%rcx)
	mov	%cx,(%rdx)
	mov	%cx,(%rbx)
	mov	%cx,(%rsp)
	mov	%cx,(%rbp)
	mov	%cx,(%rsi)
	mov	%cx,(%rdi)
	mov	%cx,(%r8)
	mov	%cx,(%r9)
	mov	%cx,(%r10)
	mov	%cx,(%r11)
	mov	%cx,(%r12)
	mov	%cx,(%r13)
	mov	%cx,(%r14)
	mov	%cx,(%r15)
	nop
	mov	%dx,(%rax)
	mov	%dx,(%rcx)
	mov	%dx,(%rdx)
	mov	%dx,(%rbx)
	mov	%dx,(%rsp)
	mov	%dx,(%rbp)
	mov	%dx,(%rsi)
	mov	%dx,(%rdi)
	mov	%dx,(%r8)
	mov	%dx,(%r9)
	mov	%dx,(%r10)
	mov	%dx,(%r11)
	mov	%dx,(%r12)
	mov	%dx,(%r13)
	mov	%dx,(%r14)
	mov	%dx,(%r15)
	nop
	mov	%bx,(%rax)
	mov	%bx,(%rcx)
	mov	%bx,(%rdx)
	mov	%bx,(%rbx)
	mov	%bx,(%rsp)
	mov	%bx,(%rbp)
	mov	%bx,(%rsi)
	mov	%bx,(%rdi)
	mov	%bx,(%r8)
	mov	%bx,(%r9)
	mov	%bx,(%r10)
	mov	%bx,(%r11)
	mov	%bx,(%r12)
	mov	%bx,(%r13)
	mov	%bx,(%r14)
	mov	%bx,(%r15)
	nop
	mov	%sp,(%rax)
	mov	%sp,(%rcx)
	mov	%sp,(%rdx)
	mov	%sp,(%rbx)
	mov	%sp,(%rsp)
	mov	%sp,(%rbp)
	mov	%sp,(%rsi)
	mov	%sp,(%rdi)
	mov	%sp,(%r8)
	mov	%sp,(%r9)
	mov	%sp,(%r10)
	mov	%sp,(%r11)
	mov	%sp,(%r12)
	mov	%sp,(%r13)
	mov	%sp,(%r14)
	mov	%sp,(%r15)
	nop
	mov	%bp,(%rax)
	mov	%bp,(%rcx)
	mov	%bp,(%rdx)
	mov	%bp,(%rbx)
	mov	%bp,(%rsp)
	mov	%bp,(%rbp)
	mov	%bp,(%rsi)
	mov	%bp,(%rdi)
	mov	%bp,(%r8)
	mov	%bp,(%r9)
	mov	%bp,(%r10)
	mov	%bp,(%r11)
	mov	%bp,(%r12)
	mov	%bp,(%r13)
	mov	%bp,(%r14)
	mov	%bp,(%r15)
	nop
	mov	%si,(%rax)
	mov	%si,(%rcx)
	mov	%si,(%rdx)
	mov	%si,(%rbx)
	mov	%si,(%rsp)
	mov	%si,(%rbp)
	mov	%si,(%rsi)
	mov	%si,(%rdi)
	mov	%si,(%r8)
	mov	%si,(%r9)
	mov	%si,(%r10)
	mov	%si,(%r11)
	mov	%si,(%r12)
	mov	%si,(%r13)
	mov	%si,(%r14)
	mov	%si,(%r15)
	nop
	mov	%di,(%rax)
	mov	%di,(%rcx)
	mov	%di,(%rdx)
	mov	%di,(%rbx)
	mov	%di,(%rsp)
	mov	%di,(%rbp)
	mov	%di,(%rsi)
	mov	%di,(%rdi)
	mov	%di,(%r8)
	mov	%di,(%r9)
	mov	%di,(%r10)
	mov	%di,(%r11)
	mov	%di,(%r12)
	mov	%di,(%r13)
	mov	%di,(%r14)
	mov	%di,(%r15)
	nop
	mov	%r8w, (%rax)
	mov	%r8w, (%rcx)
	mov	%r8w, (%rdx)
	mov	%r8w, (%rbx)
	mov	%r8w, (%rsp)
	mov	%r8w, (%rbp)
	mov	%r8w, (%rsi)
	mov	%r8w, (%rdi)
	mov	%r8w, (%r8)
	mov	%r8w, (%r9)
	mov	%r8w, (%r10)
	mov	%r8w, (%r11)
	mov	%r8w, (%r12)
	mov	%r8w, (%r13)
	mov	%r8w, (%r14)
	mov	%r8w, (%r15)
	nop
	mov	%r9w, (%rax)
	mov	%r9w, (%rcx)
	mov	%r9w, (%rdx)
	mov	%r9w, (%rbx)
	mov	%r9w, (%rsp)
	mov	%r9w, (%rbp)
	mov	%r9w, (%rsi)
	mov	%r9w, (%rdi)
	mov	%r9w, (%r8)
	mov	%r9w, (%r9)
	mov	%r9w, (%r10)
	mov	%r9w, (%r11)
	mov	%r9w, (%r12)
	mov	%r9w, (%r13)
	mov	%r9w, (%r14)
	mov	%r9w, (%r15)
	nop
	mov	%r10w,(%rax)
	mov	%r10w,(%rcx)
	mov	%r10w,(%rdx)
	mov	%r10w,(%rbx)
	mov	%r10w,(%rsp)
	mov	%r10w,(%rbp)
	mov	%r10w,(%rsi)
	mov	%r10w,(%rdi)
	mov	%r10w,(%r8)
	mov	%r10w,(%r9)
	mov	%r10w,(%r10)
	mov	%r10w,(%r11)
	mov	%r10w,(%r12)
	mov	%r10w,(%r13)
	mov	%r10w,(%r14)
	mov	%r10w,(%r15)
	nop
	mov	%r11w,(%rax)
	mov	%r11w,(%rcx)
	mov	%r11w,(%rdx)
	mov	%r11w,(%rbx)
	mov	%r11w,(%rsp)
	mov	%r11w,(%rbp)
	mov	%r11w,(%rsi)
	mov	%r11w,(%rdi)
	mov	%r11w,(%r8)
	mov	%r11w,(%r9)
	mov	%r11w,(%r10)
	mov	%r11w,(%r11)
	mov	%r11w,(%r12)
	mov	%r11w,(%r13)
	mov	%r11w,(%r14)
	mov	%r11w,(%r15)
	nop
	mov	%r12w,(%rax)
	mov	%r12w,(%rcx)
	mov	%r12w,(%rdx)
	mov	%r12w,(%rbx)
	mov	%r12w,(%rsp)
	mov	%r12w,(%rbp)
	mov	%r12w,(%rsi)
	mov	%r12w,(%rdi)
	mov	%r12w,(%r8)
	mov	%r12w,(%r9)
	mov	%r12w,(%r10)
	mov	%r12w,(%r11)
	mov	%r12w,(%r12)
	mov	%r12w,(%r13)
	mov	%r12w,(%r14)
	mov	%r12w,(%r15)
	nop
	mov	%r13w,(%rax)
	mov	%r13w,(%rcx)
	mov	%r13w,(%rdx)
	mov	%r13w,(%rbx)
	mov	%r13w,(%rsp)
	mov	%r13w,(%rbp)
	mov	%r13w,(%rsi)
	mov	%r13w,(%rdi)
	mov	%r13w,(%r8)
	mov	%r13w,(%r9)
	mov	%r13w,(%r10)
	mov	%r13w,(%r11)
	mov	%r13w,(%r12)
	mov	%r13w,(%r13)
	mov	%r13w,(%r14)
	mov	%r13w,(%r15)
	nop
	mov	%r14w,(%rax)
	mov	%r14w,(%rcx)
	mov	%r14w,(%rdx)
	mov	%r14w,(%rbx)
	mov	%r14w,(%rsp)
	mov	%r14w,(%rbp)
	mov	%r14w,(%rsi)
	mov	%r14w,(%rdi)
	mov	%r14w,(%r8)
	mov	%r14w,(%r9)
	mov	%r14w,(%r10)
	mov	%r14w,(%r11)
	mov	%r14w,(%r12)
	mov	%r14w,(%r13)
	mov	%r14w,(%r14)
	mov	%r14w,(%r15)
	nop
	mov	%r15w,(%rax)
	mov	%r15w,(%rcx)
	mov	%r15w,(%rdx)
	mov	%r15w,(%rbx)
	mov	%r15w,(%rsp)
	mov	%r15w,(%rbp)
	mov	%r15w,(%rsi)
	mov	%r15w,(%rdi)
	mov	%r15w,(%r8)
	mov	%r15w,(%r9)
	mov	%r15w,(%r10)
	mov	%r15w,(%r11)
	mov	%r15w,(%r12)
	mov	%r15w,(%r13)
	mov	%r15w,(%r14)
	mov	%r15w,(%r15)
        nop
        nop
        // mem16[off8] += reg16
	mov	%ax,0x7F(%rax)
	mov	%ax,0x7F(%rcx)
	mov	%ax,0x7F(%rdx)
	mov	%ax,0x7F(%rbx)
	mov	%ax,0x7F(%rsp)
	mov	%ax,0x7F(%rbp)
	mov	%ax,0x7F(%rsi)
	mov	%ax,0x7F(%rdi)
	mov	%ax,0x7F(%r8)
	mov	%ax,0x7F(%r9)
	mov	%ax,0x7F(%r10)
	mov	%ax,0x7F(%r11)
	mov	%ax,0x7F(%r12)
	mov	%ax,0x7F(%r13)
	mov	%ax,0x7F(%r14)
	mov	%ax,0x7F(%r15)
	nop
	mov	%cx,0x7F(%rax)
	mov	%cx,0x7F(%rcx)
	mov	%cx,0x7F(%rdx)
	mov	%cx,0x7F(%rbx)
	mov	%cx,0x7F(%rsp)
	mov	%cx,0x7F(%rbp)
	mov	%cx,0x7F(%rsi)
	mov	%cx,0x7F(%rdi)
	mov	%cx,0x7F(%r8)
	mov	%cx,0x7F(%r9)
	mov	%cx,0x7F(%r10)
	mov	%cx,0x7F(%r11)
	mov	%cx,0x7F(%r12)
	mov	%cx,0x7F(%r13)
	mov	%cx,0x7F(%r14)
	mov	%cx,0x7F(%r15)
	nop
	mov	%dx,0x7F(%rax)
	mov	%dx,0x7F(%rcx)
	mov	%dx,0x7F(%rdx)
	mov	%dx,0x7F(%rbx)
	mov	%dx,0x7F(%rsp)
	mov	%dx,0x7F(%rbp)
	mov	%dx,0x7F(%rsi)
	mov	%dx,0x7F(%rdi)
	mov	%dx,0x7F(%r8)
	mov	%dx,0x7F(%r9)
	mov	%dx,0x7F(%r10)
	mov	%dx,0x7F(%r11)
	mov	%dx,0x7F(%r12)
	mov	%dx,0x7F(%r13)
	mov	%dx,0x7F(%r14)
	mov	%dx,0x7F(%r15)
	nop
	mov	%bx,0x7F(%rax)
	mov	%bx,0x7F(%rcx)
	mov	%bx,0x7F(%rdx)
	mov	%bx,0x7F(%rbx)
	mov	%bx,0x7F(%rsp)
	mov	%bx,0x7F(%rbp)
	mov	%bx,0x7F(%rsi)
	mov	%bx,0x7F(%rdi)
	mov	%bx,0x7F(%r8)
	mov	%bx,0x7F(%r9)
	mov	%bx,0x7F(%r10)
	mov	%bx,0x7F(%r11)
	mov	%bx,0x7F(%r12)
	mov	%bx,0x7F(%r13)
	mov	%bx,0x7F(%r14)
	mov	%bx,0x7F(%r15)
	nop
	mov	%sp,0x7F(%rax)
	mov	%sp,0x7F(%rcx)
	mov	%sp,0x7F(%rdx)
	mov	%sp,0x7F(%rbx)
	mov	%sp,0x7F(%rsp)
	mov	%sp,0x7F(%rbp)
	mov	%sp,0x7F(%rsi)
	mov	%sp,0x7F(%rdi)
	mov	%sp,0x7F(%r8)
	mov	%sp,0x7F(%r9)
	mov	%sp,0x7F(%r10)
	mov	%sp,0x7F(%r11)
	mov	%sp,0x7F(%r12)
	mov	%sp,0x7F(%r13)
	mov	%sp,0x7F(%r14)
	mov	%sp,0x7F(%r15)
	nop
	mov	%bp,0x7F(%rax)
	mov	%bp,0x7F(%rcx)
	mov	%bp,0x7F(%rdx)
	mov	%bp,0x7F(%rbx)
	mov	%bp,0x7F(%rsp)
	mov	%bp,0x7F(%rbp)
	mov	%bp,0x7F(%rsi)
	mov	%bp,0x7F(%rdi)
	mov	%bp,0x7F(%r8)
	mov	%bp,0x7F(%r9)
	mov	%bp,0x7F(%r10)
	mov	%bp,0x7F(%r11)
	mov	%bp,0x7F(%r12)
	mov	%bp,0x7F(%r13)
	mov	%bp,0x7F(%r14)
	mov	%bp,0x7F(%r15)
	nop
	mov	%si,0x7F(%rax)
	mov	%si,0x7F(%rcx)
	mov	%si,0x7F(%rdx)
	mov	%si,0x7F(%rbx)
	mov	%si,0x7F(%rsp)
	mov	%si,0x7F(%rbp)
	mov	%si,0x7F(%rsi)
	mov	%si,0x7F(%rdi)
	mov	%si,0x7F(%r8)
	mov	%si,0x7F(%r9)
	mov	%si,0x7F(%r10)
	mov	%si,0x7F(%r11)
	mov	%si,0x7F(%r12)
	mov	%si,0x7F(%r13)
	mov	%si,0x7F(%r14)
	mov	%si,0x7F(%r15)
	nop
	mov	%di,0x7F(%rax)
	mov	%di,0x7F(%rcx)
	mov	%di,0x7F(%rdx)
	mov	%di,0x7F(%rbx)
	mov	%di,0x7F(%rsp)
	mov	%di,0x7F(%rbp)
	mov	%di,0x7F(%rsi)
	mov	%di,0x7F(%rdi)
	mov	%di,0x7F(%r8)
	mov	%di,0x7F(%r9)
	mov	%di,0x7F(%r10)
	mov	%di,0x7F(%r11)
	mov	%di,0x7F(%r12)
	mov	%di,0x7F(%r13)
	mov	%di,0x7F(%r14)
	mov	%di,0x7F(%r15)
	nop
	mov	%r8w, 0x7F(%rax)
	mov	%r8w, 0x7F(%rcx)
	mov	%r8w, 0x7F(%rdx)
	mov	%r8w, 0x7F(%rbx)
	mov	%r8w, 0x7F(%rsp)
	mov	%r8w, 0x7F(%rbp)
	mov	%r8w, 0x7F(%rsi)
	mov	%r8w, 0x7F(%rdi)
	mov	%r8w, 0x7F(%r8)
	mov	%r8w, 0x7F(%r9)
	mov	%r8w, 0x7F(%r10)
	mov	%r8w, 0x7F(%r11)
	mov	%r8w, 0x7F(%r12)
	mov	%r8w, 0x7F(%r13)
	mov	%r8w, 0x7F(%r14)
	mov	%r8w, 0x7F(%r15)
	nop
	mov	%r9w, 0x7F(%rax)
	mov	%r9w, 0x7F(%rcx)
	mov	%r9w, 0x7F(%rdx)
	mov	%r9w, 0x7F(%rbx)
	mov	%r9w, 0x7F(%rsp)
	mov	%r9w, 0x7F(%rbp)
	mov	%r9w, 0x7F(%rsi)
	mov	%r9w, 0x7F(%rdi)
	mov	%r9w, 0x7F(%r8)
	mov	%r9w, 0x7F(%r9)
	mov	%r9w, 0x7F(%r10)
	mov	%r9w, 0x7F(%r11)
	mov	%r9w, 0x7F(%r12)
	mov	%r9w, 0x7F(%r13)
	mov	%r9w, 0x7F(%r14)
	mov	%r9w, 0x7F(%r15)
	nop
	mov	%r10w,0x7F(%rax)
	mov	%r10w,0x7F(%rcx)
	mov	%r10w,0x7F(%rdx)
	mov	%r10w,0x7F(%rbx)
	mov	%r10w,0x7F(%rsp)
	mov	%r10w,0x7F(%rbp)
	mov	%r10w,0x7F(%rsi)
	mov	%r10w,0x7F(%rdi)
	mov	%r10w,0x7F(%r8)
	mov	%r10w,0x7F(%r9)
	mov	%r10w,0x7F(%r10)
	mov	%r10w,0x7F(%r11)
	mov	%r10w,0x7F(%r12)
	mov	%r10w,0x7F(%r13)
	mov	%r10w,0x7F(%r14)
	mov	%r10w,0x7F(%r15)
	nop
	mov	%r11w,0x7F(%rax)
	mov	%r11w,0x7F(%rcx)
	mov	%r11w,0x7F(%rdx)
	mov	%r11w,0x7F(%rbx)
	mov	%r11w,0x7F(%rsp)
	mov	%r11w,0x7F(%rbp)
	mov	%r11w,0x7F(%rsi)
	mov	%r11w,0x7F(%rdi)
	mov	%r11w,0x7F(%r8)
	mov	%r11w,0x7F(%r9)
	mov	%r11w,0x7F(%r10)
	mov	%r11w,0x7F(%r11)
	mov	%r11w,0x7F(%r12)
	mov	%r11w,0x7F(%r13)
	mov	%r11w,0x7F(%r14)
	mov	%r11w,0x7F(%r15)
	nop
	mov	%r12w,0x7F(%rax)
	mov	%r12w,0x7F(%rcx)
	mov	%r12w,0x7F(%rdx)
	mov	%r12w,0x7F(%rbx)
	mov	%r12w,0x7F(%rsp)
	mov	%r12w,0x7F(%rbp)
	mov	%r12w,0x7F(%rsi)
	mov	%r12w,0x7F(%rdi)
	mov	%r12w,0x7F(%r8)
	mov	%r12w,0x7F(%r9)
	mov	%r12w,0x7F(%r10)
	mov	%r12w,0x7F(%r11)
	mov	%r12w,0x7F(%r12)
	mov	%r12w,0x7F(%r13)
	mov	%r12w,0x7F(%r14)
	mov	%r12w,0x7F(%r15)
	nop
	mov	%r13w,0x7F(%rax)
	mov	%r13w,0x7F(%rcx)
	mov	%r13w,0x7F(%rdx)
	mov	%r13w,0x7F(%rbx)
	mov	%r13w,0x7F(%rsp)
	mov	%r13w,0x7F(%rbp)
	mov	%r13w,0x7F(%rsi)
	mov	%r13w,0x7F(%rdi)
	mov	%r13w,0x7F(%r8)
	mov	%r13w,0x7F(%r9)
	mov	%r13w,0x7F(%r10)
	mov	%r13w,0x7F(%r11)
	mov	%r13w,0x7F(%r12)
	mov	%r13w,0x7F(%r13)
	mov	%r13w,0x7F(%r14)
	mov	%r13w,0x7F(%r15)
	nop
	mov	%r14w,0x7F(%rax)
	mov	%r14w,0x7F(%rcx)
	mov	%r14w,0x7F(%rdx)
	mov	%r14w,0x7F(%rbx)
	mov	%r14w,0x7F(%rsp)
	mov	%r14w,0x7F(%rbp)
	mov	%r14w,0x7F(%rsi)
	mov	%r14w,0x7F(%rdi)
	mov	%r14w,0x7F(%r8)
	mov	%r14w,0x7F(%r9)
	mov	%r14w,0x7F(%r10)
	mov	%r14w,0x7F(%r11)
	mov	%r14w,0x7F(%r12)
	mov	%r14w,0x7F(%r13)
	mov	%r14w,0x7F(%r14)
	mov	%r14w,0x7F(%r15)
	nop
	mov	%r15w,0x7F(%rax)
	mov	%r15w,0x7F(%rcx)
	mov	%r15w,0x7F(%rdx)
	mov	%r15w,0x7F(%rbx)
	mov	%r15w,0x7F(%rsp)
	mov	%r15w,0x7F(%rbp)
	mov	%r15w,0x7F(%rsi)
	mov	%r15w,0x7F(%rdi)
	mov	%r15w,0x7F(%r8)
	mov	%r15w,0x7F(%r9)
	mov	%r15w,0x7F(%r10)
	mov	%r15w,0x7F(%r11)
	mov	%r15w,0x7F(%r12)
	mov	%r15w,0x7F(%r13)
	mov	%r15w,0x7F(%r14)
	mov	%r15w,0x7F(%r15)
        nop
        nop
        // mem16[off32] += reg16
	mov	%ax,0x12345678(%rax)
	mov	%ax,0x12345678(%rcx)
	mov	%ax,0x12345678(%rdx)
	mov	%ax,0x12345678(%rbx)
	mov	%ax,0x12345678(%rsp)
	mov	%ax,0x12345678(%rbp)
	mov	%ax,0x12345678(%rsi)
	mov	%ax,0x12345678(%rdi)
	mov	%ax,0x12345678(%r8)
	mov	%ax,0x12345678(%r9)
	mov	%ax,0x12345678(%r10)
	mov	%ax,0x12345678(%r11)
	mov	%ax,0x12345678(%r12)
	mov	%ax,0x12345678(%r13)
	mov	%ax,0x12345678(%r14)
	mov	%ax,0x12345678(%r15)
	nop
	mov	%cx,0x12345678(%rax)
	mov	%cx,0x12345678(%rcx)
	mov	%cx,0x12345678(%rdx)
	mov	%cx,0x12345678(%rbx)
	mov	%cx,0x12345678(%rsp)
	mov	%cx,0x12345678(%rbp)
	mov	%cx,0x12345678(%rsi)
	mov	%cx,0x12345678(%rdi)
	mov	%cx,0x12345678(%r8)
	mov	%cx,0x12345678(%r9)
	mov	%cx,0x12345678(%r10)
	mov	%cx,0x12345678(%r11)
	mov	%cx,0x12345678(%r12)
	mov	%cx,0x12345678(%r13)
	mov	%cx,0x12345678(%r14)
	mov	%cx,0x12345678(%r15)
	nop
	mov	%dx,0x12345678(%rax)
	mov	%dx,0x12345678(%rcx)
	mov	%dx,0x12345678(%rdx)
	mov	%dx,0x12345678(%rbx)
	mov	%dx,0x12345678(%rsp)
	mov	%dx,0x12345678(%rbp)
	mov	%dx,0x12345678(%rsi)
	mov	%dx,0x12345678(%rdi)
	mov	%dx,0x12345678(%r8)
	mov	%dx,0x12345678(%r9)
	mov	%dx,0x12345678(%r10)
	mov	%dx,0x12345678(%r11)
	mov	%dx,0x12345678(%r12)
	mov	%dx,0x12345678(%r13)
	mov	%dx,0x12345678(%r14)
	mov	%dx,0x12345678(%r15)
	nop
	mov	%bx,0x12345678(%rax)
	mov	%bx,0x12345678(%rcx)
	mov	%bx,0x12345678(%rdx)
	mov	%bx,0x12345678(%rbx)
	mov	%bx,0x12345678(%rsp)
	mov	%bx,0x12345678(%rbp)
	mov	%bx,0x12345678(%rsi)
	mov	%bx,0x12345678(%rdi)
	mov	%bx,0x12345678(%r8)
	mov	%bx,0x12345678(%r9)
	mov	%bx,0x12345678(%r10)
	mov	%bx,0x12345678(%r11)
	mov	%bx,0x12345678(%r12)
	mov	%bx,0x12345678(%r13)
	mov	%bx,0x12345678(%r14)
	mov	%bx,0x12345678(%r15)
	nop
	mov	%sp,0x12345678(%rax)
	mov	%sp,0x12345678(%rcx)
	mov	%sp,0x12345678(%rdx)
	mov	%sp,0x12345678(%rbx)
	mov	%sp,0x12345678(%rsp)
	mov	%sp,0x12345678(%rbp)
	mov	%sp,0x12345678(%rsi)
	mov	%sp,0x12345678(%rdi)
	mov	%sp,0x12345678(%r8)
	mov	%sp,0x12345678(%r9)
	mov	%sp,0x12345678(%r10)
	mov	%sp,0x12345678(%r11)
	mov	%sp,0x12345678(%r12)
	mov	%sp,0x12345678(%r13)
	mov	%sp,0x12345678(%r14)
	mov	%sp,0x12345678(%r15)
	nop
	mov	%bp,0x12345678(%rax)
	mov	%bp,0x12345678(%rcx)
	mov	%bp,0x12345678(%rdx)
	mov	%bp,0x12345678(%rbx)
	mov	%bp,0x12345678(%rsp)
	mov	%bp,0x12345678(%rbp)
	mov	%bp,0x12345678(%rsi)
	mov	%bp,0x12345678(%rdi)
	mov	%bp,0x12345678(%r8)
	mov	%bp,0x12345678(%r9)
	mov	%bp,0x12345678(%r10)
	mov	%bp,0x12345678(%r11)
	mov	%bp,0x12345678(%r12)
	mov	%bp,0x12345678(%r13)
	mov	%bp,0x12345678(%r14)
	mov	%bp,0x12345678(%r15)
	nop
	mov	%si,0x12345678(%rax)
	mov	%si,0x12345678(%rcx)
	mov	%si,0x12345678(%rdx)
	mov	%si,0x12345678(%rbx)
	mov	%si,0x12345678(%rsp)
	mov	%si,0x12345678(%rbp)
	mov	%si,0x12345678(%rsi)
	mov	%si,0x12345678(%rdi)
	mov	%si,0x12345678(%r8)
	mov	%si,0x12345678(%r9)
	mov	%si,0x12345678(%r10)
	mov	%si,0x12345678(%r11)
	mov	%si,0x12345678(%r12)
	mov	%si,0x12345678(%r13)
	mov	%si,0x12345678(%r14)
	mov	%si,0x12345678(%r15)
	nop
	mov	%di,0x12345678(%rax)
	mov	%di,0x12345678(%rcx)
	mov	%di,0x12345678(%rdx)
	mov	%di,0x12345678(%rbx)
	mov	%di,0x12345678(%rsp)
	mov	%di,0x12345678(%rbp)
	mov	%di,0x12345678(%rsi)
	mov	%di,0x12345678(%rdi)
	mov	%di,0x12345678(%r8)
	mov	%di,0x12345678(%r9)
	mov	%di,0x12345678(%r10)
	mov	%di,0x12345678(%r11)
	mov	%di,0x12345678(%r12)
	mov	%di,0x12345678(%r13)
	mov	%di,0x12345678(%r14)
	mov	%di,0x12345678(%r15)
	nop
	mov	%r8w, 0x12345678(%rax)
	mov	%r8w, 0x12345678(%rcx)
	mov	%r8w, 0x12345678(%rdx)
	mov	%r8w, 0x12345678(%rbx)
	mov	%r8w, 0x12345678(%rsp)
	mov	%r8w, 0x12345678(%rbp)
	mov	%r8w, 0x12345678(%rsi)
	mov	%r8w, 0x12345678(%rdi)
	mov	%r8w, 0x12345678(%r8)
	mov	%r8w, 0x12345678(%r9)
	mov	%r8w, 0x12345678(%r10)
	mov	%r8w, 0x12345678(%r11)
	mov	%r8w, 0x12345678(%r12)
	mov	%r8w, 0x12345678(%r13)
	mov	%r8w, 0x12345678(%r14)
	mov	%r8w, 0x12345678(%r15)
	nop
	mov	%r9w, 0x12345678(%rax)
	mov	%r9w, 0x12345678(%rcx)
	mov	%r9w, 0x12345678(%rdx)
	mov	%r9w, 0x12345678(%rbx)
	mov	%r9w, 0x12345678(%rsp)
	mov	%r9w, 0x12345678(%rbp)
	mov	%r9w, 0x12345678(%rsi)
	mov	%r9w, 0x12345678(%rdi)
	mov	%r9w, 0x12345678(%r8)
	mov	%r9w, 0x12345678(%r9)
	mov	%r9w, 0x12345678(%r10)
	mov	%r9w, 0x12345678(%r11)
	mov	%r9w, 0x12345678(%r12)
	mov	%r9w, 0x12345678(%r13)
	mov	%r9w, 0x12345678(%r14)
	mov	%r9w, 0x12345678(%r15)
	nop
	mov	%r10w,0x12345678(%rax)
	mov	%r10w,0x12345678(%rcx)
	mov	%r10w,0x12345678(%rdx)
	mov	%r10w,0x12345678(%rbx)
	mov	%r10w,0x12345678(%rsp)
	mov	%r10w,0x12345678(%rbp)
	mov	%r10w,0x12345678(%rsi)
	mov	%r10w,0x12345678(%rdi)
	mov	%r10w,0x12345678(%r8)
	mov	%r10w,0x12345678(%r9)
	mov	%r10w,0x12345678(%r10)
	mov	%r10w,0x12345678(%r11)
	mov	%r10w,0x12345678(%r12)
	mov	%r10w,0x12345678(%r13)
	mov	%r10w,0x12345678(%r14)
	mov	%r10w,0x12345678(%r15)
	nop
	mov	%r11w,0x12345678(%rax)
	mov	%r11w,0x12345678(%rcx)
	mov	%r11w,0x12345678(%rdx)
	mov	%r11w,0x12345678(%rbx)
	mov	%r11w,0x12345678(%rsp)
	mov	%r11w,0x12345678(%rbp)
	mov	%r11w,0x12345678(%rsi)
	mov	%r11w,0x12345678(%rdi)
	mov	%r11w,0x12345678(%r8)
	mov	%r11w,0x12345678(%r9)
	mov	%r11w,0x12345678(%r10)
	mov	%r11w,0x12345678(%r11)
	mov	%r11w,0x12345678(%r12)
	mov	%r11w,0x12345678(%r13)
	mov	%r11w,0x12345678(%r14)
	mov	%r11w,0x12345678(%r15)
	nop
	mov	%r12w,0x12345678(%rax)
	mov	%r12w,0x12345678(%rcx)
	mov	%r12w,0x12345678(%rdx)
	mov	%r12w,0x12345678(%rbx)
	mov	%r12w,0x12345678(%rsp)
	mov	%r12w,0x12345678(%rbp)
	mov	%r12w,0x12345678(%rsi)
	mov	%r12w,0x12345678(%rdi)
	mov	%r12w,0x12345678(%r8)
	mov	%r12w,0x12345678(%r9)
	mov	%r12w,0x12345678(%r10)
	mov	%r12w,0x12345678(%r11)
	mov	%r12w,0x12345678(%r12)
	mov	%r12w,0x12345678(%r13)
	mov	%r12w,0x12345678(%r14)
	mov	%r12w,0x12345678(%r15)
	nop
	mov	%r13w,0x12345678(%rax)
	mov	%r13w,0x12345678(%rcx)
	mov	%r13w,0x12345678(%rdx)
	mov	%r13w,0x12345678(%rbx)
	mov	%r13w,0x12345678(%rsp)
	mov	%r13w,0x12345678(%rbp)
	mov	%r13w,0x12345678(%rsi)
	mov	%r13w,0x12345678(%rdi)
	mov	%r13w,0x12345678(%r8)
	mov	%r13w,0x12345678(%r9)
	mov	%r13w,0x12345678(%r10)
	mov	%r13w,0x12345678(%r11)
	mov	%r13w,0x12345678(%r12)
	mov	%r13w,0x12345678(%r13)
	mov	%r13w,0x12345678(%r14)
	mov	%r13w,0x12345678(%r15)
	nop
	mov	%r14w,0x12345678(%rax)
	mov	%r14w,0x12345678(%rcx)
	mov	%r14w,0x12345678(%rdx)
	mov	%r14w,0x12345678(%rbx)
	mov	%r14w,0x12345678(%rsp)
	mov	%r14w,0x12345678(%rbp)
	mov	%r14w,0x12345678(%rsi)
	mov	%r14w,0x12345678(%rdi)
	mov	%r14w,0x12345678(%r8)
	mov	%r14w,0x12345678(%r9)
	mov	%r14w,0x12345678(%r10)
	mov	%r14w,0x12345678(%r11)
	mov	%r14w,0x12345678(%r12)
	mov	%r14w,0x12345678(%r13)
	mov	%r14w,0x12345678(%r14)
	mov	%r14w,0x12345678(%r15)
	nop
	mov	%r15w,0x12345678(%rax)
	mov	%r15w,0x12345678(%rcx)
	mov	%r15w,0x12345678(%rdx)
	mov	%r15w,0x12345678(%rbx)
	mov	%r15w,0x12345678(%rsp)
	mov	%r15w,0x12345678(%rbp)
	mov	%r15w,0x12345678(%rsi)
	mov	%r15w,0x12345678(%rdi)
	mov	%r15w,0x12345678(%r8)
	mov	%r15w,0x12345678(%r9)
	mov	%r15w,0x12345678(%r10)
	mov	%r15w,0x12345678(%r11)
	mov	%r15w,0x12345678(%r12)
	mov	%r15w,0x12345678(%r13)
	mov	%r15w,0x12345678(%r14)
	mov	%r15w,0x12345678(%r15)
        ret
	.cfi_endproc

	.p2align 4,,15
	.globl	MovMemReg32
	.type	MovMemReg32, @function
MovMemReg32:
	.cfi_startproc
        // mem32[0] += reg32
	mov	%eax,(%rax)
	mov	%eax,(%rcx)
	mov	%eax,(%rdx)
	mov	%eax,(%rbx)
	mov	%eax,(%rsp)
	mov	%eax,(%rbp)
	mov	%eax,(%rsi)
	mov	%eax,(%rdi)
	mov	%eax,(%r8)
	mov	%eax,(%r9)
	mov	%eax,(%r10)
	mov	%eax,(%r11)
	mov	%eax,(%r12)
	mov	%eax,(%r13)
	mov	%eax,(%r14)
	mov	%eax,(%r15)
	nop
	mov	%ecx,(%rax)
	mov	%ecx,(%rcx)
	mov	%ecx,(%rdx)
	mov	%ecx,(%rbx)
	mov	%ecx,(%rsp)
	mov	%ecx,(%rbp)
	mov	%ecx,(%rsi)
	mov	%ecx,(%rdi)
	mov	%ecx,(%r8)
	mov	%ecx,(%r9)
	mov	%ecx,(%r10)
	mov	%ecx,(%r11)
	mov	%ecx,(%r12)
	mov	%ecx,(%r13)
	mov	%ecx,(%r14)
	mov	%ecx,(%r15)
	nop
	mov	%edx,(%rax)
	mov	%edx,(%rcx)
	mov	%edx,(%rdx)
	mov	%edx,(%rbx)
	mov	%edx,(%rsp)
	mov	%edx,(%rbp)
	mov	%edx,(%rsi)
	mov	%edx,(%rdi)
	mov	%edx,(%r8)
	mov	%edx,(%r9)
	mov	%edx,(%r10)
	mov	%edx,(%r11)
	mov	%edx,(%r12)
	mov	%edx,(%r13)
	mov	%edx,(%r14)
	mov	%edx,(%r15)
	nop
	mov	%ebx,(%rax)
	mov	%ebx,(%rcx)
	mov	%ebx,(%rdx)
	mov	%ebx,(%rbx)
	mov	%ebx,(%rsp)
	mov	%ebx,(%rbp)
	mov	%ebx,(%rsi)
	mov	%ebx,(%rdi)
	mov	%ebx,(%r8)
	mov	%ebx,(%r9)
	mov	%ebx,(%r10)
	mov	%ebx,(%r11)
	mov	%ebx,(%r12)
	mov	%ebx,(%r13)
	mov	%ebx,(%r14)
	mov	%ebx,(%r15)
	nop
	mov	%esp,(%rax)
	mov	%esp,(%rcx)
	mov	%esp,(%rdx)
	mov	%esp,(%rbx)
	mov	%esp,(%rsp)
	mov	%esp,(%rbp)
	mov	%esp,(%rsi)
	mov	%esp,(%rdi)
	mov	%esp,(%r8)
	mov	%esp,(%r9)
	mov	%esp,(%r10)
	mov	%esp,(%r11)
	mov	%esp,(%r12)
	mov	%esp,(%r13)
	mov	%esp,(%r14)
	mov	%esp,(%r15)
	nop
	mov	%ebp,(%rax)
	mov	%ebp,(%rcx)
	mov	%ebp,(%rdx)
	mov	%ebp,(%rbx)
	mov	%ebp,(%rsp)
	mov	%ebp,(%rbp)
	mov	%ebp,(%rsi)
	mov	%ebp,(%rdi)
	mov	%ebp,(%r8)
	mov	%ebp,(%r9)
	mov	%ebp,(%r10)
	mov	%ebp,(%r11)
	mov	%ebp,(%r12)
	mov	%ebp,(%r13)
	mov	%ebp,(%r14)
	mov	%ebp,(%r15)
	nop
	mov	%esi,(%rax)
	mov	%esi,(%rcx)
	mov	%esi,(%rdx)
	mov	%esi,(%rbx)
	mov	%esi,(%rsp)
	mov	%esi,(%rbp)
	mov	%esi,(%rsi)
	mov	%esi,(%rdi)
	mov	%esi,(%r8)
	mov	%esi,(%r9)
	mov	%esi,(%r10)
	mov	%esi,(%r11)
	mov	%esi,(%r12)
	mov	%esi,(%r13)
	mov	%esi,(%r14)
	mov	%esi,(%r15)
	nop
	mov	%edi,(%rax)
	mov	%edi,(%rcx)
	mov	%edi,(%rdx)
	mov	%edi,(%rbx)
	mov	%edi,(%rsp)
	mov	%edi,(%rbp)
	mov	%edi,(%rsi)
	mov	%edi,(%rdi)
	mov	%edi,(%r8)
	mov	%edi,(%r9)
	mov	%edi,(%r10)
	mov	%edi,(%r11)
	mov	%edi,(%r12)
	mov	%edi,(%r13)
	mov	%edi,(%r14)
	mov	%edi,(%r15)
	nop
	mov	%r8d, (%rax)
	mov	%r8d, (%rcx)
	mov	%r8d, (%rdx)
	mov	%r8d, (%rbx)
	mov	%r8d, (%rsp)
	mov	%r8d, (%rbp)
	mov	%r8d, (%rsi)
	mov	%r8d, (%rdi)
	mov	%r8d, (%r8)
	mov	%r8d, (%r9)
	mov	%r8d, (%r10)
	mov	%r8d, (%r11)
	mov	%r8d, (%r12)
	mov	%r8d, (%r13)
	mov	%r8d, (%r14)
	mov	%r8d, (%r15)
	nop
	mov	%r9d, (%rax)
	mov	%r9d, (%rcx)
	mov	%r9d, (%rdx)
	mov	%r9d, (%rbx)
	mov	%r9d, (%rsp)
	mov	%r9d, (%rbp)
	mov	%r9d, (%rsi)
	mov	%r9d, (%rdi)
	mov	%r9d, (%r8)
	mov	%r9d, (%r9)
	mov	%r9d, (%r10)
	mov	%r9d, (%r11)
	mov	%r9d, (%r12)
	mov	%r9d, (%r13)
	mov	%r9d, (%r14)
	mov	%r9d, (%r15)
	nop
	mov	%r10d,(%rax)
	mov	%r10d,(%rcx)
	mov	%r10d,(%rdx)
	mov	%r10d,(%rbx)
	mov	%r10d,(%rsp)
	mov	%r10d,(%rbp)
	mov	%r10d,(%rsi)
	mov	%r10d,(%rdi)
	mov	%r10d,(%r8)
	mov	%r10d,(%r9)
	mov	%r10d,(%r10)
	mov	%r10d,(%r11)
	mov	%r10d,(%r12)
	mov	%r10d,(%r13)
	mov	%r10d,(%r14)
	mov	%r10d,(%r15)
	nop
	mov	%r11d,(%rax)
	mov	%r11d,(%rcx)
	mov	%r11d,(%rdx)
	mov	%r11d,(%rbx)
	mov	%r11d,(%rsp)
	mov	%r11d,(%rbp)
	mov	%r11d,(%rsi)
	mov	%r11d,(%rdi)
	mov	%r11d,(%r8)
	mov	%r11d,(%r9)
	mov	%r11d,(%r10)
	mov	%r11d,(%r11)
	mov	%r11d,(%r12)
	mov	%r11d,(%r13)
	mov	%r11d,(%r14)
	mov	%r11d,(%r15)
	nop
	mov	%r12d,(%rax)
	mov	%r12d,(%rcx)
	mov	%r12d,(%rdx)
	mov	%r12d,(%rbx)
	mov	%r12d,(%rsp)
	mov	%r12d,(%rbp)
	mov	%r12d,(%rsi)
	mov	%r12d,(%rdi)
	mov	%r12d,(%r8)
	mov	%r12d,(%r9)
	mov	%r12d,(%r10)
	mov	%r12d,(%r11)
	mov	%r12d,(%r12)
	mov	%r12d,(%r13)
	mov	%r12d,(%r14)
	mov	%r12d,(%r15)
	nop
	mov	%r13d,(%rax)
	mov	%r13d,(%rcx)
	mov	%r13d,(%rdx)
	mov	%r13d,(%rbx)
	mov	%r13d,(%rsp)
	mov	%r13d,(%rbp)
	mov	%r13d,(%rsi)
	mov	%r13d,(%rdi)
	mov	%r13d,(%r8)
	mov	%r13d,(%r9)
	mov	%r13d,(%r10)
	mov	%r13d,(%r11)
	mov	%r13d,(%r12)
	mov	%r13d,(%r13)
	mov	%r13d,(%r14)
	mov	%r13d,(%r15)
	nop
	mov	%r14d,(%rax)
	mov	%r14d,(%rcx)
	mov	%r14d,(%rdx)
	mov	%r14d,(%rbx)
	mov	%r14d,(%rsp)
	mov	%r14d,(%rbp)
	mov	%r14d,(%rsi)
	mov	%r14d,(%rdi)
	mov	%r14d,(%r8)
	mov	%r14d,(%r9)
	mov	%r14d,(%r10)
	mov	%r14d,(%r11)
	mov	%r14d,(%r12)
	mov	%r14d,(%r13)
	mov	%r14d,(%r14)
	mov	%r14d,(%r15)
	nop
	mov	%r15d,(%rax)
	mov	%r15d,(%rcx)
	mov	%r15d,(%rdx)
	mov	%r15d,(%rbx)
	mov	%r15d,(%rsp)
	mov	%r15d,(%rbp)
	mov	%r15d,(%rsi)
	mov	%r15d,(%rdi)
	mov	%r15d,(%r8)
	mov	%r15d,(%r9)
	mov	%r15d,(%r10)
	mov	%r15d,(%r11)
	mov	%r15d,(%r12)
	mov	%r15d,(%r13)
	mov	%r15d,(%r14)
	mov	%r15d,(%r15)
        nop
        nop
        // mem32[off8] += reg32
	mov	%eax,0x7F(%rax)
	mov	%eax,0x7F(%rcx)
	mov	%eax,0x7F(%rdx)
	mov	%eax,0x7F(%rbx)
	mov	%eax,0x7F(%rsp)
	mov	%eax,0x7F(%rbp)
	mov	%eax,0x7F(%rsi)
	mov	%eax,0x7F(%rdi)
	mov	%eax,0x7F(%r8)
	mov	%eax,0x7F(%r9)
	mov	%eax,0x7F(%r10)
	mov	%eax,0x7F(%r11)
	mov	%eax,0x7F(%r12)
	mov	%eax,0x7F(%r13)
	mov	%eax,0x7F(%r14)
	mov	%eax,0x7F(%r15)
	nop
	mov	%ecx,0x7F(%rax)
	mov	%ecx,0x7F(%rcx)
	mov	%ecx,0x7F(%rdx)
	mov	%ecx,0x7F(%rbx)
	mov	%ecx,0x7F(%rsp)
	mov	%ecx,0x7F(%rbp)
	mov	%ecx,0x7F(%rsi)
	mov	%ecx,0x7F(%rdi)
	mov	%ecx,0x7F(%r8)
	mov	%ecx,0x7F(%r9)
	mov	%ecx,0x7F(%r10)
	mov	%ecx,0x7F(%r11)
	mov	%ecx,0x7F(%r12)
	mov	%ecx,0x7F(%r13)
	mov	%ecx,0x7F(%r14)
	mov	%ecx,0x7F(%r15)
	nop
	mov	%edx,0x7F(%rax)
	mov	%edx,0x7F(%rcx)
	mov	%edx,0x7F(%rdx)
	mov	%edx,0x7F(%rbx)
	mov	%edx,0x7F(%rsp)
	mov	%edx,0x7F(%rbp)
	mov	%edx,0x7F(%rsi)
	mov	%edx,0x7F(%rdi)
	mov	%edx,0x7F(%r8)
	mov	%edx,0x7F(%r9)
	mov	%edx,0x7F(%r10)
	mov	%edx,0x7F(%r11)
	mov	%edx,0x7F(%r12)
	mov	%edx,0x7F(%r13)
	mov	%edx,0x7F(%r14)
	mov	%edx,0x7F(%r15)
	nop
	mov	%ebx,0x7F(%rax)
	mov	%ebx,0x7F(%rcx)
	mov	%ebx,0x7F(%rdx)
	mov	%ebx,0x7F(%rbx)
	mov	%ebx,0x7F(%rsp)
	mov	%ebx,0x7F(%rbp)
	mov	%ebx,0x7F(%rsi)
	mov	%ebx,0x7F(%rdi)
	mov	%ebx,0x7F(%r8)
	mov	%ebx,0x7F(%r9)
	mov	%ebx,0x7F(%r10)
	mov	%ebx,0x7F(%r11)
	mov	%ebx,0x7F(%r12)
	mov	%ebx,0x7F(%r13)
	mov	%ebx,0x7F(%r14)
	mov	%ebx,0x7F(%r15)
	nop
	mov	%esp,0x7F(%rax)
	mov	%esp,0x7F(%rcx)
	mov	%esp,0x7F(%rdx)
	mov	%esp,0x7F(%rbx)
	mov	%esp,0x7F(%rsp)
	mov	%esp,0x7F(%rbp)
	mov	%esp,0x7F(%rsi)
	mov	%esp,0x7F(%rdi)
	mov	%esp,0x7F(%r8)
	mov	%esp,0x7F(%r9)
	mov	%esp,0x7F(%r10)
	mov	%esp,0x7F(%r11)
	mov	%esp,0x7F(%r12)
	mov	%esp,0x7F(%r13)
	mov	%esp,0x7F(%r14)
	mov	%esp,0x7F(%r15)
	nop
	mov	%ebp,0x7F(%rax)
	mov	%ebp,0x7F(%rcx)
	mov	%ebp,0x7F(%rdx)
	mov	%ebp,0x7F(%rbx)
	mov	%ebp,0x7F(%rsp)
	mov	%ebp,0x7F(%rbp)
	mov	%ebp,0x7F(%rsi)
	mov	%ebp,0x7F(%rdi)
	mov	%ebp,0x7F(%r8)
	mov	%ebp,0x7F(%r9)
	mov	%ebp,0x7F(%r10)
	mov	%ebp,0x7F(%r11)
	mov	%ebp,0x7F(%r12)
	mov	%ebp,0x7F(%r13)
	mov	%ebp,0x7F(%r14)
	mov	%ebp,0x7F(%r15)
	nop
	mov	%esi,0x7F(%rax)
	mov	%esi,0x7F(%rcx)
	mov	%esi,0x7F(%rdx)
	mov	%esi,0x7F(%rbx)
	mov	%esi,0x7F(%rsp)
	mov	%esi,0x7F(%rbp)
	mov	%esi,0x7F(%rsi)
	mov	%esi,0x7F(%rdi)
	mov	%esi,0x7F(%r8)
	mov	%esi,0x7F(%r9)
	mov	%esi,0x7F(%r10)
	mov	%esi,0x7F(%r11)
	mov	%esi,0x7F(%r12)
	mov	%esi,0x7F(%r13)
	mov	%esi,0x7F(%r14)
	mov	%esi,0x7F(%r15)
	nop
	mov	%edi,0x7F(%rax)
	mov	%edi,0x7F(%rcx)
	mov	%edi,0x7F(%rdx)
	mov	%edi,0x7F(%rbx)
	mov	%edi,0x7F(%rsp)
	mov	%edi,0x7F(%rbp)
	mov	%edi,0x7F(%rsi)
	mov	%edi,0x7F(%rdi)
	mov	%edi,0x7F(%r8)
	mov	%edi,0x7F(%r9)
	mov	%edi,0x7F(%r10)
	mov	%edi,0x7F(%r11)
	mov	%edi,0x7F(%r12)
	mov	%edi,0x7F(%r13)
	mov	%edi,0x7F(%r14)
	mov	%edi,0x7F(%r15)
	nop
	mov	%r8d, 0x7F(%rax)
	mov	%r8d, 0x7F(%rcx)
	mov	%r8d, 0x7F(%rdx)
	mov	%r8d, 0x7F(%rbx)
	mov	%r8d, 0x7F(%rsp)
	mov	%r8d, 0x7F(%rbp)
	mov	%r8d, 0x7F(%rsi)
	mov	%r8d, 0x7F(%rdi)
	mov	%r8d, 0x7F(%r8)
	mov	%r8d, 0x7F(%r9)
	mov	%r8d, 0x7F(%r10)
	mov	%r8d, 0x7F(%r11)
	mov	%r8d, 0x7F(%r12)
	mov	%r8d, 0x7F(%r13)
	mov	%r8d, 0x7F(%r14)
	mov	%r8d, 0x7F(%r15)
	nop
	mov	%r9d, 0x7F(%rax)
	mov	%r9d, 0x7F(%rcx)
	mov	%r9d, 0x7F(%rdx)
	mov	%r9d, 0x7F(%rbx)
	mov	%r9d, 0x7F(%rsp)
	mov	%r9d, 0x7F(%rbp)
	mov	%r9d, 0x7F(%rsi)
	mov	%r9d, 0x7F(%rdi)
	mov	%r9d, 0x7F(%r8)
	mov	%r9d, 0x7F(%r9)
	mov	%r9d, 0x7F(%r10)
	mov	%r9d, 0x7F(%r11)
	mov	%r9d, 0x7F(%r12)
	mov	%r9d, 0x7F(%r13)
	mov	%r9d, 0x7F(%r14)
	mov	%r9d, 0x7F(%r15)
	nop
	mov	%r10d,0x7F(%rax)
	mov	%r10d,0x7F(%rcx)
	mov	%r10d,0x7F(%rdx)
	mov	%r10d,0x7F(%rbx)
	mov	%r10d,0x7F(%rsp)
	mov	%r10d,0x7F(%rbp)
	mov	%r10d,0x7F(%rsi)
	mov	%r10d,0x7F(%rdi)
	mov	%r10d,0x7F(%r8)
	mov	%r10d,0x7F(%r9)
	mov	%r10d,0x7F(%r10)
	mov	%r10d,0x7F(%r11)
	mov	%r10d,0x7F(%r12)
	mov	%r10d,0x7F(%r13)
	mov	%r10d,0x7F(%r14)
	mov	%r10d,0x7F(%r15)
	nop
	mov	%r11d,0x7F(%rax)
	mov	%r11d,0x7F(%rcx)
	mov	%r11d,0x7F(%rdx)
	mov	%r11d,0x7F(%rbx)
	mov	%r11d,0x7F(%rsp)
	mov	%r11d,0x7F(%rbp)
	mov	%r11d,0x7F(%rsi)
	mov	%r11d,0x7F(%rdi)
	mov	%r11d,0x7F(%r8)
	mov	%r11d,0x7F(%r9)
	mov	%r11d,0x7F(%r10)
	mov	%r11d,0x7F(%r11)
	mov	%r11d,0x7F(%r12)
	mov	%r11d,0x7F(%r13)
	mov	%r11d,0x7F(%r14)
	mov	%r11d,0x7F(%r15)
	nop
	mov	%r12d,0x7F(%rax)
	mov	%r12d,0x7F(%rcx)
	mov	%r12d,0x7F(%rdx)
	mov	%r12d,0x7F(%rbx)
	mov	%r12d,0x7F(%rsp)
	mov	%r12d,0x7F(%rbp)
	mov	%r12d,0x7F(%rsi)
	mov	%r12d,0x7F(%rdi)
	mov	%r12d,0x7F(%r8)
	mov	%r12d,0x7F(%r9)
	mov	%r12d,0x7F(%r10)
	mov	%r12d,0x7F(%r11)
	mov	%r12d,0x7F(%r12)
	mov	%r12d,0x7F(%r13)
	mov	%r12d,0x7F(%r14)
	mov	%r12d,0x7F(%r15)
	nop
	mov	%r13d,0x7F(%rax)
	mov	%r13d,0x7F(%rcx)
	mov	%r13d,0x7F(%rdx)
	mov	%r13d,0x7F(%rbx)
	mov	%r13d,0x7F(%rsp)
	mov	%r13d,0x7F(%rbp)
	mov	%r13d,0x7F(%rsi)
	mov	%r13d,0x7F(%rdi)
	mov	%r13d,0x7F(%r8)
	mov	%r13d,0x7F(%r9)
	mov	%r13d,0x7F(%r10)
	mov	%r13d,0x7F(%r11)
	mov	%r13d,0x7F(%r12)
	mov	%r13d,0x7F(%r13)
	mov	%r13d,0x7F(%r14)
	mov	%r13d,0x7F(%r15)
	nop
	mov	%r14d,0x7F(%rax)
	mov	%r14d,0x7F(%rcx)
	mov	%r14d,0x7F(%rdx)
	mov	%r14d,0x7F(%rbx)
	mov	%r14d,0x7F(%rsp)
	mov	%r14d,0x7F(%rbp)
	mov	%r14d,0x7F(%rsi)
	mov	%r14d,0x7F(%rdi)
	mov	%r14d,0x7F(%r8)
	mov	%r14d,0x7F(%r9)
	mov	%r14d,0x7F(%r10)
	mov	%r14d,0x7F(%r11)
	mov	%r14d,0x7F(%r12)
	mov	%r14d,0x7F(%r13)
	mov	%r14d,0x7F(%r14)
	mov	%r14d,0x7F(%r15)
	nop
	mov	%r15d,0x7F(%rax)
	mov	%r15d,0x7F(%rcx)
	mov	%r15d,0x7F(%rdx)
	mov	%r15d,0x7F(%rbx)
	mov	%r15d,0x7F(%rsp)
	mov	%r15d,0x7F(%rbp)
	mov	%r15d,0x7F(%rsi)
	mov	%r15d,0x7F(%rdi)
	mov	%r15d,0x7F(%r8)
	mov	%r15d,0x7F(%r9)
	mov	%r15d,0x7F(%r10)
	mov	%r15d,0x7F(%r11)
	mov	%r15d,0x7F(%r12)
	mov	%r15d,0x7F(%r13)
	mov	%r15d,0x7F(%r14)
	mov	%r15d,0x7F(%r15)
        nop
        nop
        // mem32[off32] += reg32
	mov	%eax,0x12345678(%rax)
	mov	%eax,0x12345678(%rcx)
	mov	%eax,0x12345678(%rdx)
	mov	%eax,0x12345678(%rbx)
	mov	%eax,0x12345678(%rsp)
	mov	%eax,0x12345678(%rbp)
	mov	%eax,0x12345678(%rsi)
	mov	%eax,0x12345678(%rdi)
	mov	%eax,0x12345678(%r8)
	mov	%eax,0x12345678(%r9)
	mov	%eax,0x12345678(%r10)
	mov	%eax,0x12345678(%r11)
	mov	%eax,0x12345678(%r12)
	mov	%eax,0x12345678(%r13)
	mov	%eax,0x12345678(%r14)
	mov	%eax,0x12345678(%r15)
	nop
	mov	%ecx,0x12345678(%rax)
	mov	%ecx,0x12345678(%rcx)
	mov	%ecx,0x12345678(%rdx)
	mov	%ecx,0x12345678(%rbx)
	mov	%ecx,0x12345678(%rsp)
	mov	%ecx,0x12345678(%rbp)
	mov	%ecx,0x12345678(%rsi)
	mov	%ecx,0x12345678(%rdi)
	mov	%ecx,0x12345678(%r8)
	mov	%ecx,0x12345678(%r9)
	mov	%ecx,0x12345678(%r10)
	mov	%ecx,0x12345678(%r11)
	mov	%ecx,0x12345678(%r12)
	mov	%ecx,0x12345678(%r13)
	mov	%ecx,0x12345678(%r14)
	mov	%ecx,0x12345678(%r15)
	nop
	mov	%edx,0x12345678(%rax)
	mov	%edx,0x12345678(%rcx)
	mov	%edx,0x12345678(%rdx)
	mov	%edx,0x12345678(%rbx)
	mov	%edx,0x12345678(%rsp)
	mov	%edx,0x12345678(%rbp)
	mov	%edx,0x12345678(%rsi)
	mov	%edx,0x12345678(%rdi)
	mov	%edx,0x12345678(%r8)
	mov	%edx,0x12345678(%r9)
	mov	%edx,0x12345678(%r10)
	mov	%edx,0x12345678(%r11)
	mov	%edx,0x12345678(%r12)
	mov	%edx,0x12345678(%r13)
	mov	%edx,0x12345678(%r14)
	mov	%edx,0x12345678(%r15)
	nop
	mov	%ebx,0x12345678(%rax)
	mov	%ebx,0x12345678(%rcx)
	mov	%ebx,0x12345678(%rdx)
	mov	%ebx,0x12345678(%rbx)
	mov	%ebx,0x12345678(%rsp)
	mov	%ebx,0x12345678(%rbp)
	mov	%ebx,0x12345678(%rsi)
	mov	%ebx,0x12345678(%rdi)
	mov	%ebx,0x12345678(%r8)
	mov	%ebx,0x12345678(%r9)
	mov	%ebx,0x12345678(%r10)
	mov	%ebx,0x12345678(%r11)
	mov	%ebx,0x12345678(%r12)
	mov	%ebx,0x12345678(%r13)
	mov	%ebx,0x12345678(%r14)
	mov	%ebx,0x12345678(%r15)
	nop
	mov	%esp,0x12345678(%rax)
	mov	%esp,0x12345678(%rcx)
	mov	%esp,0x12345678(%rdx)
	mov	%esp,0x12345678(%rbx)
	mov	%esp,0x12345678(%rsp)
	mov	%esp,0x12345678(%rbp)
	mov	%esp,0x12345678(%rsi)
	mov	%esp,0x12345678(%rdi)
	mov	%esp,0x12345678(%r8)
	mov	%esp,0x12345678(%r9)
	mov	%esp,0x12345678(%r10)
	mov	%esp,0x12345678(%r11)
	mov	%esp,0x12345678(%r12)
	mov	%esp,0x12345678(%r13)
	mov	%esp,0x12345678(%r14)
	mov	%esp,0x12345678(%r15)
	nop
	mov	%ebp,0x12345678(%rax)
	mov	%ebp,0x12345678(%rcx)
	mov	%ebp,0x12345678(%rdx)
	mov	%ebp,0x12345678(%rbx)
	mov	%ebp,0x12345678(%rsp)
	mov	%ebp,0x12345678(%rbp)
	mov	%ebp,0x12345678(%rsi)
	mov	%ebp,0x12345678(%rdi)
	mov	%ebp,0x12345678(%r8)
	mov	%ebp,0x12345678(%r9)
	mov	%ebp,0x12345678(%r10)
	mov	%ebp,0x12345678(%r11)
	mov	%ebp,0x12345678(%r12)
	mov	%ebp,0x12345678(%r13)
	mov	%ebp,0x12345678(%r14)
	mov	%ebp,0x12345678(%r15)
	nop
	mov	%esi,0x12345678(%rax)
	mov	%esi,0x12345678(%rcx)
	mov	%esi,0x12345678(%rdx)
	mov	%esi,0x12345678(%rbx)
	mov	%esi,0x12345678(%rsp)
	mov	%esi,0x12345678(%rbp)
	mov	%esi,0x12345678(%rsi)
	mov	%esi,0x12345678(%rdi)
	mov	%esi,0x12345678(%r8)
	mov	%esi,0x12345678(%r9)
	mov	%esi,0x12345678(%r10)
	mov	%esi,0x12345678(%r11)
	mov	%esi,0x12345678(%r12)
	mov	%esi,0x12345678(%r13)
	mov	%esi,0x12345678(%r14)
	mov	%esi,0x12345678(%r15)
	nop
	mov	%edi,0x12345678(%rax)
	mov	%edi,0x12345678(%rcx)
	mov	%edi,0x12345678(%rdx)
	mov	%edi,0x12345678(%rbx)
	mov	%edi,0x12345678(%rsp)
	mov	%edi,0x12345678(%rbp)
	mov	%edi,0x12345678(%rsi)
	mov	%edi,0x12345678(%rdi)
	mov	%edi,0x12345678(%r8)
	mov	%edi,0x12345678(%r9)
	mov	%edi,0x12345678(%r10)
	mov	%edi,0x12345678(%r11)
	mov	%edi,0x12345678(%r12)
	mov	%edi,0x12345678(%r13)
	mov	%edi,0x12345678(%r14)
	mov	%edi,0x12345678(%r15)
	nop
	mov	%r8d, 0x12345678(%rax)
	mov	%r8d, 0x12345678(%rcx)
	mov	%r8d, 0x12345678(%rdx)
	mov	%r8d, 0x12345678(%rbx)
	mov	%r8d, 0x12345678(%rsp)
	mov	%r8d, 0x12345678(%rbp)
	mov	%r8d, 0x12345678(%rsi)
	mov	%r8d, 0x12345678(%rdi)
	mov	%r8d, 0x12345678(%r8)
	mov	%r8d, 0x12345678(%r9)
	mov	%r8d, 0x12345678(%r10)
	mov	%r8d, 0x12345678(%r11)
	mov	%r8d, 0x12345678(%r12)
	mov	%r8d, 0x12345678(%r13)
	mov	%r8d, 0x12345678(%r14)
	mov	%r8d, 0x12345678(%r15)
	nop
	mov	%r9d, 0x12345678(%rax)
	mov	%r9d, 0x12345678(%rcx)
	mov	%r9d, 0x12345678(%rdx)
	mov	%r9d, 0x12345678(%rbx)
	mov	%r9d, 0x12345678(%rsp)
	mov	%r9d, 0x12345678(%rbp)
	mov	%r9d, 0x12345678(%rsi)
	mov	%r9d, 0x12345678(%rdi)
	mov	%r9d, 0x12345678(%r8)
	mov	%r9d, 0x12345678(%r9)
	mov	%r9d, 0x12345678(%r10)
	mov	%r9d, 0x12345678(%r11)
	mov	%r9d, 0x12345678(%r12)
	mov	%r9d, 0x12345678(%r13)
	mov	%r9d, 0x12345678(%r14)
	mov	%r9d, 0x12345678(%r15)
	nop
	mov	%r10d,0x12345678(%rax)
	mov	%r10d,0x12345678(%rcx)
	mov	%r10d,0x12345678(%rdx)
	mov	%r10d,0x12345678(%rbx)
	mov	%r10d,0x12345678(%rsp)
	mov	%r10d,0x12345678(%rbp)
	mov	%r10d,0x12345678(%rsi)
	mov	%r10d,0x12345678(%rdi)
	mov	%r10d,0x12345678(%r8)
	mov	%r10d,0x12345678(%r9)
	mov	%r10d,0x12345678(%r10)
	mov	%r10d,0x12345678(%r11)
	mov	%r10d,0x12345678(%r12)
	mov	%r10d,0x12345678(%r13)
	mov	%r10d,0x12345678(%r14)
	mov	%r10d,0x12345678(%r15)
	nop
	mov	%r11d,0x12345678(%rax)
	mov	%r11d,0x12345678(%rcx)
	mov	%r11d,0x12345678(%rdx)
	mov	%r11d,0x12345678(%rbx)
	mov	%r11d,0x12345678(%rsp)
	mov	%r11d,0x12345678(%rbp)
	mov	%r11d,0x12345678(%rsi)
	mov	%r11d,0x12345678(%rdi)
	mov	%r11d,0x12345678(%r8)
	mov	%r11d,0x12345678(%r9)
	mov	%r11d,0x12345678(%r10)
	mov	%r11d,0x12345678(%r11)
	mov	%r11d,0x12345678(%r12)
	mov	%r11d,0x12345678(%r13)
	mov	%r11d,0x12345678(%r14)
	mov	%r11d,0x12345678(%r15)
	nop
	mov	%r12d,0x12345678(%rax)
	mov	%r12d,0x12345678(%rcx)
	mov	%r12d,0x12345678(%rdx)
	mov	%r12d,0x12345678(%rbx)
	mov	%r12d,0x12345678(%rsp)
	mov	%r12d,0x12345678(%rbp)
	mov	%r12d,0x12345678(%rsi)
	mov	%r12d,0x12345678(%rdi)
	mov	%r12d,0x12345678(%r8)
	mov	%r12d,0x12345678(%r9)
	mov	%r12d,0x12345678(%r10)
	mov	%r12d,0x12345678(%r11)
	mov	%r12d,0x12345678(%r12)
	mov	%r12d,0x12345678(%r13)
	mov	%r12d,0x12345678(%r14)
	mov	%r12d,0x12345678(%r15)
	nop
	mov	%r13d,0x12345678(%rax)
	mov	%r13d,0x12345678(%rcx)
	mov	%r13d,0x12345678(%rdx)
	mov	%r13d,0x12345678(%rbx)
	mov	%r13d,0x12345678(%rsp)
	mov	%r13d,0x12345678(%rbp)
	mov	%r13d,0x12345678(%rsi)
	mov	%r13d,0x12345678(%rdi)
	mov	%r13d,0x12345678(%r8)
	mov	%r13d,0x12345678(%r9)
	mov	%r13d,0x12345678(%r10)
	mov	%r13d,0x12345678(%r11)
	mov	%r13d,0x12345678(%r12)
	mov	%r13d,0x12345678(%r13)
	mov	%r13d,0x12345678(%r14)
	mov	%r13d,0x12345678(%r15)
	nop
	mov	%r14d,0x12345678(%rax)
	mov	%r14d,0x12345678(%rcx)
	mov	%r14d,0x12345678(%rdx)
	mov	%r14d,0x12345678(%rbx)
	mov	%r14d,0x12345678(%rsp)
	mov	%r14d,0x12345678(%rbp)
	mov	%r14d,0x12345678(%rsi)
	mov	%r14d,0x12345678(%rdi)
	mov	%r14d,0x12345678(%r8)
	mov	%r14d,0x12345678(%r9)
	mov	%r14d,0x12345678(%r10)
	mov	%r14d,0x12345678(%r11)
	mov	%r14d,0x12345678(%r12)
	mov	%r14d,0x12345678(%r13)
	mov	%r14d,0x12345678(%r14)
	mov	%r14d,0x12345678(%r15)
	nop
	mov	%r15d,0x12345678(%rax)
	mov	%r15d,0x12345678(%rcx)
	mov	%r15d,0x12345678(%rdx)
	mov	%r15d,0x12345678(%rbx)
	mov	%r15d,0x12345678(%rsp)
	mov	%r15d,0x12345678(%rbp)
	mov	%r15d,0x12345678(%rsi)
	mov	%r15d,0x12345678(%rdi)
	mov	%r15d,0x12345678(%r8)
	mov	%r15d,0x12345678(%r9)
	mov	%r15d,0x12345678(%r10)
	mov	%r15d,0x12345678(%r11)
	mov	%r15d,0x12345678(%r12)
	mov	%r15d,0x12345678(%r13)
	mov	%r15d,0x12345678(%r14)
	mov	%r15d,0x12345678(%r15)
        ret
	.cfi_endproc

	.p2align 4,,15
	.globl	MovMemReg64
	.type	MovMemReg64, @function
MovMemReg64:
	.cfi_startproc
        // mem64[0] += reg64
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
        nop
        nop
        // mem64[off8] += reg64
	mov	%rax,0x7F(%rax)
	mov	%rax,0x7F(%rcx)
	mov	%rax,0x7F(%rdx)
	mov	%rax,0x7F(%rbx)
	mov	%rax,0x7F(%rsp)
	mov	%rax,0x7F(%rbp)
	mov	%rax,0x7F(%rsi)
	mov	%rax,0x7F(%rdi)
	mov	%rax,0x7F(%r8)
	mov	%rax,0x7F(%r9)
	mov	%rax,0x7F(%r10)
	mov	%rax,0x7F(%r11)
	mov	%rax,0x7F(%r12)
	mov	%rax,0x7F(%r13)
	mov	%rax,0x7F(%r14)
	mov	%rax,0x7F(%r15)
	nop
	mov	%rcx,0x7F(%rax)
	mov	%rcx,0x7F(%rcx)
	mov	%rcx,0x7F(%rdx)
	mov	%rcx,0x7F(%rbx)
	mov	%rcx,0x7F(%rsp)
	mov	%rcx,0x7F(%rbp)
	mov	%rcx,0x7F(%rsi)
	mov	%rcx,0x7F(%rdi)
	mov	%rcx,0x7F(%r8)
	mov	%rcx,0x7F(%r9)
	mov	%rcx,0x7F(%r10)
	mov	%rcx,0x7F(%r11)
	mov	%rcx,0x7F(%r12)
	mov	%rcx,0x7F(%r13)
	mov	%rcx,0x7F(%r14)
	mov	%rcx,0x7F(%r15)
	nop
	mov	%rdx,0x7F(%rax)
	mov	%rdx,0x7F(%rcx)
	mov	%rdx,0x7F(%rdx)
	mov	%rdx,0x7F(%rbx)
	mov	%rdx,0x7F(%rsp)
	mov	%rdx,0x7F(%rbp)
	mov	%rdx,0x7F(%rsi)
	mov	%rdx,0x7F(%rdi)
	mov	%rdx,0x7F(%r8)
	mov	%rdx,0x7F(%r9)
	mov	%rdx,0x7F(%r10)
	mov	%rdx,0x7F(%r11)
	mov	%rdx,0x7F(%r12)
	mov	%rdx,0x7F(%r13)
	mov	%rdx,0x7F(%r14)
	mov	%rdx,0x7F(%r15)
	nop
	mov	%rbx,0x7F(%rax)
	mov	%rbx,0x7F(%rcx)
	mov	%rbx,0x7F(%rdx)
	mov	%rbx,0x7F(%rbx)
	mov	%rbx,0x7F(%rsp)
	mov	%rbx,0x7F(%rbp)
	mov	%rbx,0x7F(%rsi)
	mov	%rbx,0x7F(%rdi)
	mov	%rbx,0x7F(%r8)
	mov	%rbx,0x7F(%r9)
	mov	%rbx,0x7F(%r10)
	mov	%rbx,0x7F(%r11)
	mov	%rbx,0x7F(%r12)
	mov	%rbx,0x7F(%r13)
	mov	%rbx,0x7F(%r14)
	mov	%rbx,0x7F(%r15)
	nop
	mov	%rsp,0x7F(%rax)
	mov	%rsp,0x7F(%rcx)
	mov	%rsp,0x7F(%rdx)
	mov	%rsp,0x7F(%rbx)
	mov	%rsp,0x7F(%rsp)
	mov	%rsp,0x7F(%rbp)
	mov	%rsp,0x7F(%rsi)
	mov	%rsp,0x7F(%rdi)
	mov	%rsp,0x7F(%r8)
	mov	%rsp,0x7F(%r9)
	mov	%rsp,0x7F(%r10)
	mov	%rsp,0x7F(%r11)
	mov	%rsp,0x7F(%r12)
	mov	%rsp,0x7F(%r13)
	mov	%rsp,0x7F(%r14)
	mov	%rsp,0x7F(%r15)
	nop
	mov	%rbp,0x7F(%rax)
	mov	%rbp,0x7F(%rcx)
	mov	%rbp,0x7F(%rdx)
	mov	%rbp,0x7F(%rbx)
	mov	%rbp,0x7F(%rsp)
	mov	%rbp,0x7F(%rbp)
	mov	%rbp,0x7F(%rsi)
	mov	%rbp,0x7F(%rdi)
	mov	%rbp,0x7F(%r8)
	mov	%rbp,0x7F(%r9)
	mov	%rbp,0x7F(%r10)
	mov	%rbp,0x7F(%r11)
	mov	%rbp,0x7F(%r12)
	mov	%rbp,0x7F(%r13)
	mov	%rbp,0x7F(%r14)
	mov	%rbp,0x7F(%r15)
	nop
	mov	%rsi,0x7F(%rax)
	mov	%rsi,0x7F(%rcx)
	mov	%rsi,0x7F(%rdx)
	mov	%rsi,0x7F(%rbx)
	mov	%rsi,0x7F(%rsp)
	mov	%rsi,0x7F(%rbp)
	mov	%rsi,0x7F(%rsi)
	mov	%rsi,0x7F(%rdi)
	mov	%rsi,0x7F(%r8)
	mov	%rsi,0x7F(%r9)
	mov	%rsi,0x7F(%r10)
	mov	%rsi,0x7F(%r11)
	mov	%rsi,0x7F(%r12)
	mov	%rsi,0x7F(%r13)
	mov	%rsi,0x7F(%r14)
	mov	%rsi,0x7F(%r15)
	nop
	mov	%rdi,0x7F(%rax)
	mov	%rdi,0x7F(%rcx)
	mov	%rdi,0x7F(%rdx)
	mov	%rdi,0x7F(%rbx)
	mov	%rdi,0x7F(%rsp)
	mov	%rdi,0x7F(%rbp)
	mov	%rdi,0x7F(%rsi)
	mov	%rdi,0x7F(%rdi)
	mov	%rdi,0x7F(%r8)
	mov	%rdi,0x7F(%r9)
	mov	%rdi,0x7F(%r10)
	mov	%rdi,0x7F(%r11)
	mov	%rdi,0x7F(%r12)
	mov	%rdi,0x7F(%r13)
	mov	%rdi,0x7F(%r14)
	mov	%rdi,0x7F(%r15)
	nop
	mov	%r8, 0x7F(%rax)
	mov	%r8, 0x7F(%rcx)
	mov	%r8, 0x7F(%rdx)
	mov	%r8, 0x7F(%rbx)
	mov	%r8, 0x7F(%rsp)
	mov	%r8, 0x7F(%rbp)
	mov	%r8, 0x7F(%rsi)
	mov	%r8, 0x7F(%rdi)
	mov	%r8, 0x7F(%r8)
	mov	%r8, 0x7F(%r9)
	mov	%r8, 0x7F(%r10)
	mov	%r8, 0x7F(%r11)
	mov	%r8, 0x7F(%r12)
	mov	%r8, 0x7F(%r13)
	mov	%r8, 0x7F(%r14)
	mov	%r8, 0x7F(%r15)
	nop
	mov	%r9, 0x7F(%rax)
	mov	%r9, 0x7F(%rcx)
	mov	%r9, 0x7F(%rdx)
	mov	%r9, 0x7F(%rbx)
	mov	%r9, 0x7F(%rsp)
	mov	%r9, 0x7F(%rbp)
	mov	%r9, 0x7F(%rsi)
	mov	%r9, 0x7F(%rdi)
	mov	%r9, 0x7F(%r8)
	mov	%r9, 0x7F(%r9)
	mov	%r9, 0x7F(%r10)
	mov	%r9, 0x7F(%r11)
	mov	%r9, 0x7F(%r12)
	mov	%r9, 0x7F(%r13)
	mov	%r9, 0x7F(%r14)
	mov	%r9, 0x7F(%r15)
	nop
	mov	%r10,0x7F(%rax)
	mov	%r10,0x7F(%rcx)
	mov	%r10,0x7F(%rdx)
	mov	%r10,0x7F(%rbx)
	mov	%r10,0x7F(%rsp)
	mov	%r10,0x7F(%rbp)
	mov	%r10,0x7F(%rsi)
	mov	%r10,0x7F(%rdi)
	mov	%r10,0x7F(%r8)
	mov	%r10,0x7F(%r9)
	mov	%r10,0x7F(%r10)
	mov	%r10,0x7F(%r11)
	mov	%r10,0x7F(%r12)
	mov	%r10,0x7F(%r13)
	mov	%r10,0x7F(%r14)
	mov	%r10,0x7F(%r15)
	nop
	mov	%r11,0x7F(%rax)
	mov	%r11,0x7F(%rcx)
	mov	%r11,0x7F(%rdx)
	mov	%r11,0x7F(%rbx)
	mov	%r11,0x7F(%rsp)
	mov	%r11,0x7F(%rbp)
	mov	%r11,0x7F(%rsi)
	mov	%r11,0x7F(%rdi)
	mov	%r11,0x7F(%r8)
	mov	%r11,0x7F(%r9)
	mov	%r11,0x7F(%r10)
	mov	%r11,0x7F(%r11)
	mov	%r11,0x7F(%r12)
	mov	%r11,0x7F(%r13)
	mov	%r11,0x7F(%r14)
	mov	%r11,0x7F(%r15)
	nop
	mov	%r12,0x7F(%rax)
	mov	%r12,0x7F(%rcx)
	mov	%r12,0x7F(%rdx)
	mov	%r12,0x7F(%rbx)
	mov	%r12,0x7F(%rsp)
	mov	%r12,0x7F(%rbp)
	mov	%r12,0x7F(%rsi)
	mov	%r12,0x7F(%rdi)
	mov	%r12,0x7F(%r8)
	mov	%r12,0x7F(%r9)
	mov	%r12,0x7F(%r10)
	mov	%r12,0x7F(%r11)
	mov	%r12,0x7F(%r12)
	mov	%r12,0x7F(%r13)
	mov	%r12,0x7F(%r14)
	mov	%r12,0x7F(%r15)
	nop
	mov	%r13,0x7F(%rax)
	mov	%r13,0x7F(%rcx)
	mov	%r13,0x7F(%rdx)
	mov	%r13,0x7F(%rbx)
	mov	%r13,0x7F(%rsp)
	mov	%r13,0x7F(%rbp)
	mov	%r13,0x7F(%rsi)
	mov	%r13,0x7F(%rdi)
	mov	%r13,0x7F(%r8)
	mov	%r13,0x7F(%r9)
	mov	%r13,0x7F(%r10)
	mov	%r13,0x7F(%r11)
	mov	%r13,0x7F(%r12)
	mov	%r13,0x7F(%r13)
	mov	%r13,0x7F(%r14)
	mov	%r13,0x7F(%r15)
	nop
	mov	%r14,0x7F(%rax)
	mov	%r14,0x7F(%rcx)
	mov	%r14,0x7F(%rdx)
	mov	%r14,0x7F(%rbx)
	mov	%r14,0x7F(%rsp)
	mov	%r14,0x7F(%rbp)
	mov	%r14,0x7F(%rsi)
	mov	%r14,0x7F(%rdi)
	mov	%r14,0x7F(%r8)
	mov	%r14,0x7F(%r9)
	mov	%r14,0x7F(%r10)
	mov	%r14,0x7F(%r11)
	mov	%r14,0x7F(%r12)
	mov	%r14,0x7F(%r13)
	mov	%r14,0x7F(%r14)
	mov	%r14,0x7F(%r15)
	nop
	mov	%r15,0x7F(%rax)
	mov	%r15,0x7F(%rcx)
	mov	%r15,0x7F(%rdx)
	mov	%r15,0x7F(%rbx)
	mov	%r15,0x7F(%rsp)
	mov	%r15,0x7F(%rbp)
	mov	%r15,0x7F(%rsi)
	mov	%r15,0x7F(%rdi)
	mov	%r15,0x7F(%r8)
	mov	%r15,0x7F(%r9)
	mov	%r15,0x7F(%r10)
	mov	%r15,0x7F(%r11)
	mov	%r15,0x7F(%r12)
	mov	%r15,0x7F(%r13)
	mov	%r15,0x7F(%r14)
	mov	%r15,0x7F(%r15)
        nop
        nop
        // mem64[off32] += reg64
	mov	%rax,0x12345678(%rax)
	mov	%rax,0x12345678(%rcx)
	mov	%rax,0x12345678(%rdx)
	mov	%rax,0x12345678(%rbx)
	mov	%rax,0x12345678(%rsp)
	mov	%rax,0x12345678(%rbp)
	mov	%rax,0x12345678(%rsi)
	mov	%rax,0x12345678(%rdi)
	mov	%rax,0x12345678(%r8)
	mov	%rax,0x12345678(%r9)
	mov	%rax,0x12345678(%r10)
	mov	%rax,0x12345678(%r11)
	mov	%rax,0x12345678(%r12)
	mov	%rax,0x12345678(%r13)
	mov	%rax,0x12345678(%r14)
	mov	%rax,0x12345678(%r15)
	nop
	mov	%rcx,0x12345678(%rax)
	mov	%rcx,0x12345678(%rcx)
	mov	%rcx,0x12345678(%rdx)
	mov	%rcx,0x12345678(%rbx)
	mov	%rcx,0x12345678(%rsp)
	mov	%rcx,0x12345678(%rbp)
	mov	%rcx,0x12345678(%rsi)
	mov	%rcx,0x12345678(%rdi)
	mov	%rcx,0x12345678(%r8)
	mov	%rcx,0x12345678(%r9)
	mov	%rcx,0x12345678(%r10)
	mov	%rcx,0x12345678(%r11)
	mov	%rcx,0x12345678(%r12)
	mov	%rcx,0x12345678(%r13)
	mov	%rcx,0x12345678(%r14)
	mov	%rcx,0x12345678(%r15)
	nop
	mov	%rdx,0x12345678(%rax)
	mov	%rdx,0x12345678(%rcx)
	mov	%rdx,0x12345678(%rdx)
	mov	%rdx,0x12345678(%rbx)
	mov	%rdx,0x12345678(%rsp)
	mov	%rdx,0x12345678(%rbp)
	mov	%rdx,0x12345678(%rsi)
	mov	%rdx,0x12345678(%rdi)
	mov	%rdx,0x12345678(%r8)
	mov	%rdx,0x12345678(%r9)
	mov	%rdx,0x12345678(%r10)
	mov	%rdx,0x12345678(%r11)
	mov	%rdx,0x12345678(%r12)
	mov	%rdx,0x12345678(%r13)
	mov	%rdx,0x12345678(%r14)
	mov	%rdx,0x12345678(%r15)
	nop
	mov	%rbx,0x12345678(%rax)
	mov	%rbx,0x12345678(%rcx)
	mov	%rbx,0x12345678(%rdx)
	mov	%rbx,0x12345678(%rbx)
	mov	%rbx,0x12345678(%rsp)
	mov	%rbx,0x12345678(%rbp)
	mov	%rbx,0x12345678(%rsi)
	mov	%rbx,0x12345678(%rdi)
	mov	%rbx,0x12345678(%r8)
	mov	%rbx,0x12345678(%r9)
	mov	%rbx,0x12345678(%r10)
	mov	%rbx,0x12345678(%r11)
	mov	%rbx,0x12345678(%r12)
	mov	%rbx,0x12345678(%r13)
	mov	%rbx,0x12345678(%r14)
	mov	%rbx,0x12345678(%r15)
	nop
	mov	%rsp,0x12345678(%rax)
	mov	%rsp,0x12345678(%rcx)
	mov	%rsp,0x12345678(%rdx)
	mov	%rsp,0x12345678(%rbx)
	mov	%rsp,0x12345678(%rsp)
	mov	%rsp,0x12345678(%rbp)
	mov	%rsp,0x12345678(%rsi)
	mov	%rsp,0x12345678(%rdi)
	mov	%rsp,0x12345678(%r8)
	mov	%rsp,0x12345678(%r9)
	mov	%rsp,0x12345678(%r10)
	mov	%rsp,0x12345678(%r11)
	mov	%rsp,0x12345678(%r12)
	mov	%rsp,0x12345678(%r13)
	mov	%rsp,0x12345678(%r14)
	mov	%rsp,0x12345678(%r15)
	nop
	mov	%rbp,0x12345678(%rax)
	mov	%rbp,0x12345678(%rcx)
	mov	%rbp,0x12345678(%rdx)
	mov	%rbp,0x12345678(%rbx)
	mov	%rbp,0x12345678(%rsp)
	mov	%rbp,0x12345678(%rbp)
	mov	%rbp,0x12345678(%rsi)
	mov	%rbp,0x12345678(%rdi)
	mov	%rbp,0x12345678(%r8)
	mov	%rbp,0x12345678(%r9)
	mov	%rbp,0x12345678(%r10)
	mov	%rbp,0x12345678(%r11)
	mov	%rbp,0x12345678(%r12)
	mov	%rbp,0x12345678(%r13)
	mov	%rbp,0x12345678(%r14)
	mov	%rbp,0x12345678(%r15)
	nop
	mov	%rsi,0x12345678(%rax)
	mov	%rsi,0x12345678(%rcx)
	mov	%rsi,0x12345678(%rdx)
	mov	%rsi,0x12345678(%rbx)
	mov	%rsi,0x12345678(%rsp)
	mov	%rsi,0x12345678(%rbp)
	mov	%rsi,0x12345678(%rsi)
	mov	%rsi,0x12345678(%rdi)
	mov	%rsi,0x12345678(%r8)
	mov	%rsi,0x12345678(%r9)
	mov	%rsi,0x12345678(%r10)
	mov	%rsi,0x12345678(%r11)
	mov	%rsi,0x12345678(%r12)
	mov	%rsi,0x12345678(%r13)
	mov	%rsi,0x12345678(%r14)
	mov	%rsi,0x12345678(%r15)
	nop
	mov	%rdi,0x12345678(%rax)
	mov	%rdi,0x12345678(%rcx)
	mov	%rdi,0x12345678(%rdx)
	mov	%rdi,0x12345678(%rbx)
	mov	%rdi,0x12345678(%rsp)
	mov	%rdi,0x12345678(%rbp)
	mov	%rdi,0x12345678(%rsi)
	mov	%rdi,0x12345678(%rdi)
	mov	%rdi,0x12345678(%r8)
	mov	%rdi,0x12345678(%r9)
	mov	%rdi,0x12345678(%r10)
	mov	%rdi,0x12345678(%r11)
	mov	%rdi,0x12345678(%r12)
	mov	%rdi,0x12345678(%r13)
	mov	%rdi,0x12345678(%r14)
	mov	%rdi,0x12345678(%r15)
	nop
	mov	%r8, 0x12345678(%rax)
	mov	%r8, 0x12345678(%rcx)
	mov	%r8, 0x12345678(%rdx)
	mov	%r8, 0x12345678(%rbx)
	mov	%r8, 0x12345678(%rsp)
	mov	%r8, 0x12345678(%rbp)
	mov	%r8, 0x12345678(%rsi)
	mov	%r8, 0x12345678(%rdi)
	mov	%r8, 0x12345678(%r8)
	mov	%r8, 0x12345678(%r9)
	mov	%r8, 0x12345678(%r10)
	mov	%r8, 0x12345678(%r11)
	mov	%r8, 0x12345678(%r12)
	mov	%r8, 0x12345678(%r13)
	mov	%r8, 0x12345678(%r14)
	mov	%r8, 0x12345678(%r15)
	nop
	mov	%r9, 0x12345678(%rax)
	mov	%r9, 0x12345678(%rcx)
	mov	%r9, 0x12345678(%rdx)
	mov	%r9, 0x12345678(%rbx)
	mov	%r9, 0x12345678(%rsp)
	mov	%r9, 0x12345678(%rbp)
	mov	%r9, 0x12345678(%rsi)
	mov	%r9, 0x12345678(%rdi)
	mov	%r9, 0x12345678(%r8)
	mov	%r9, 0x12345678(%r9)
	mov	%r9, 0x12345678(%r10)
	mov	%r9, 0x12345678(%r11)
	mov	%r9, 0x12345678(%r12)
	mov	%r9, 0x12345678(%r13)
	mov	%r9, 0x12345678(%r14)
	mov	%r9, 0x12345678(%r15)
	nop
	mov	%r10,0x12345678(%rax)
	mov	%r10,0x12345678(%rcx)
	mov	%r10,0x12345678(%rdx)
	mov	%r10,0x12345678(%rbx)
	mov	%r10,0x12345678(%rsp)
	mov	%r10,0x12345678(%rbp)
	mov	%r10,0x12345678(%rsi)
	mov	%r10,0x12345678(%rdi)
	mov	%r10,0x12345678(%r8)
	mov	%r10,0x12345678(%r9)
	mov	%r10,0x12345678(%r10)
	mov	%r10,0x12345678(%r11)
	mov	%r10,0x12345678(%r12)
	mov	%r10,0x12345678(%r13)
	mov	%r10,0x12345678(%r14)
	mov	%r10,0x12345678(%r15)
	nop
	mov	%r11,0x12345678(%rax)
	mov	%r11,0x12345678(%rcx)
	mov	%r11,0x12345678(%rdx)
	mov	%r11,0x12345678(%rbx)
	mov	%r11,0x12345678(%rsp)
	mov	%r11,0x12345678(%rbp)
	mov	%r11,0x12345678(%rsi)
	mov	%r11,0x12345678(%rdi)
	mov	%r11,0x12345678(%r8)
	mov	%r11,0x12345678(%r9)
	mov	%r11,0x12345678(%r10)
	mov	%r11,0x12345678(%r11)
	mov	%r11,0x12345678(%r12)
	mov	%r11,0x12345678(%r13)
	mov	%r11,0x12345678(%r14)
	mov	%r11,0x12345678(%r15)
	nop
	mov	%r12,0x12345678(%rax)
	mov	%r12,0x12345678(%rcx)
	mov	%r12,0x12345678(%rdx)
	mov	%r12,0x12345678(%rbx)
	mov	%r12,0x12345678(%rsp)
	mov	%r12,0x12345678(%rbp)
	mov	%r12,0x12345678(%rsi)
	mov	%r12,0x12345678(%rdi)
	mov	%r12,0x12345678(%r8)
	mov	%r12,0x12345678(%r9)
	mov	%r12,0x12345678(%r10)
	mov	%r12,0x12345678(%r11)
	mov	%r12,0x12345678(%r12)
	mov	%r12,0x12345678(%r13)
	mov	%r12,0x12345678(%r14)
	mov	%r12,0x12345678(%r15)
	nop
	mov	%r13,0x12345678(%rax)
	mov	%r13,0x12345678(%rcx)
	mov	%r13,0x12345678(%rdx)
	mov	%r13,0x12345678(%rbx)
	mov	%r13,0x12345678(%rsp)
	mov	%r13,0x12345678(%rbp)
	mov	%r13,0x12345678(%rsi)
	mov	%r13,0x12345678(%rdi)
	mov	%r13,0x12345678(%r8)
	mov	%r13,0x12345678(%r9)
	mov	%r13,0x12345678(%r10)
	mov	%r13,0x12345678(%r11)
	mov	%r13,0x12345678(%r12)
	mov	%r13,0x12345678(%r13)
	mov	%r13,0x12345678(%r14)
	mov	%r13,0x12345678(%r15)
	nop
	mov	%r14,0x12345678(%rax)
	mov	%r14,0x12345678(%rcx)
	mov	%r14,0x12345678(%rdx)
	mov	%r14,0x12345678(%rbx)
	mov	%r14,0x12345678(%rsp)
	mov	%r14,0x12345678(%rbp)
	mov	%r14,0x12345678(%rsi)
	mov	%r14,0x12345678(%rdi)
	mov	%r14,0x12345678(%r8)
	mov	%r14,0x12345678(%r9)
	mov	%r14,0x12345678(%r10)
	mov	%r14,0x12345678(%r11)
	mov	%r14,0x12345678(%r12)
	mov	%r14,0x12345678(%r13)
	mov	%r14,0x12345678(%r14)
	mov	%r14,0x12345678(%r15)
	nop
	mov	%r15,0x12345678(%rax)
	mov	%r15,0x12345678(%rcx)
	mov	%r15,0x12345678(%rdx)
	mov	%r15,0x12345678(%rbx)
	mov	%r15,0x12345678(%rsp)
	mov	%r15,0x12345678(%rbp)
	mov	%r15,0x12345678(%rsi)
	mov	%r15,0x12345678(%rdi)
	mov	%r15,0x12345678(%r8)
	mov	%r15,0x12345678(%r9)
	mov	%r15,0x12345678(%r10)
	mov	%r15,0x12345678(%r11)
	mov	%r15,0x12345678(%r12)
	mov	%r15,0x12345678(%r13)
	mov	%r15,0x12345678(%r14)
	mov	%r15,0x12345678(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	MovRegMem8
	.type	MovRegMem8, @function
MovRegMem8:
	.cfi_startproc
        // reg8 = mem8[0]
	mov	(%rax),%al
	mov	(%rax),%cl
	mov	(%rax),%dl
	mov	(%rax),%bl
	mov	(%rax),%spl
	mov	(%rax),%bpl
	mov	(%rax),%sil
	mov	(%rax),%dil
	mov	(%rax),%r8b
	mov	(%rax),%r9b
	mov	(%rax),%r10b
	mov	(%rax),%r11b
	mov	(%rax),%r12b
	mov	(%rax),%r13b
	mov	(%rax),%r14b
	mov	(%rax),%r15b
	nop
	mov	(%rcx),%al
	mov	(%rcx),%cl
	mov	(%rcx),%dl
	mov	(%rcx),%bl
	mov	(%rcx),%spl
	mov	(%rcx),%bpl
	mov	(%rcx),%sil
	mov	(%rcx),%dil
	mov	(%rcx),%r8b
	mov	(%rcx),%r9b
	mov	(%rcx),%r10b
	mov	(%rcx),%r11b
	mov	(%rcx),%r12b
	mov	(%rcx),%r13b
	mov	(%rcx),%r14b
	mov	(%rcx),%r15b
	nop
	mov	(%rdx),%al
	mov	(%rdx),%cl
	mov	(%rdx),%dl
	mov	(%rdx),%bl
	mov	(%rdx),%spl
	mov	(%rdx),%bpl
	mov	(%rdx),%sil
	mov	(%rdx),%dil
	mov	(%rdx),%r8b
	mov	(%rdx),%r9b
	mov	(%rdx),%r10b
	mov	(%rdx),%r11b
	mov	(%rdx),%r12b
	mov	(%rdx),%r13b
	mov	(%rdx),%r14b
	mov	(%rdx),%r15b
	nop
	mov	(%rbx),%al
	mov	(%rbx),%cl
	mov	(%rbx),%dl
	mov	(%rbx),%bl
	mov	(%rbx),%spl
	mov	(%rbx),%bpl
	mov	(%rbx),%sil
	mov	(%rbx),%dil
	mov	(%rbx),%r8b
	mov	(%rbx),%r9b
	mov	(%rbx),%r10b
	mov	(%rbx),%r11b
	mov	(%rbx),%r12b
	mov	(%rbx),%r13b
	mov	(%rbx),%r14b
	mov	(%rbx),%r15b
	nop
	mov	(%rsp),%al
	mov	(%rsp),%cl
	mov	(%rsp),%dl
	mov	(%rsp),%bl
	mov	(%rsp),%spl
	mov	(%rsp),%bpl
	mov	(%rsp),%sil
	mov	(%rsp),%dil
	mov	(%rsp),%r8b
	mov	(%rsp),%r9b
	mov	(%rsp),%r10b
	mov	(%rsp),%r11b
	mov	(%rsp),%r12b
	mov	(%rsp),%r13b
	mov	(%rsp),%r14b
	mov	(%rsp),%r15b
	nop
	mov	(%rbp),%al
	mov	(%rbp),%cl
	mov	(%rbp),%dl
	mov	(%rbp),%bl
	mov	(%rbp),%spl
	mov	(%rbp),%bpl
	mov	(%rbp),%sil
	mov	(%rbp),%dil
	mov	(%rbp),%r8b
	mov	(%rbp),%r9b
	mov	(%rbp),%r10b
	mov	(%rbp),%r11b
	mov	(%rbp),%r12b
	mov	(%rbp),%r13b
	mov	(%rbp),%r14b
	mov	(%rbp),%r15b
	nop
	mov	(%rsi),%al
	mov	(%rsi),%cl
	mov	(%rsi),%dl
	mov	(%rsi),%bl
	mov	(%rsi),%spl
	mov	(%rsi),%bpl
	mov	(%rsi),%sil
	mov	(%rsi),%dil
	mov	(%rsi),%r8b
	mov	(%rsi),%r9b
	mov	(%rsi),%r10b
	mov	(%rsi),%r11b
	mov	(%rsi),%r12b
	mov	(%rsi),%r13b
	mov	(%rsi),%r14b
	mov	(%rsi),%r15b
	nop
	mov	(%rdi),%al
	mov	(%rdi),%cl
	mov	(%rdi),%dl
	mov	(%rdi),%bl
	mov	(%rdi),%spl
	mov	(%rdi),%bpl
	mov	(%rdi),%sil
	mov	(%rdi),%dil
	mov	(%rdi),%r8b
	mov	(%rdi),%r9b
	mov	(%rdi),%r10b
	mov	(%rdi),%r11b
	mov	(%rdi),%r12b
	mov	(%rdi),%r13b
	mov	(%rdi),%r14b
	mov	(%rdi),%r15b
	nop
	mov	(%r8), %al
	mov	(%r8), %cl
	mov	(%r8), %dl
	mov	(%r8), %bl
	mov	(%r8), %spl
	mov	(%r8), %bpl
	mov	(%r8), %sil
	mov	(%r8), %dil
	mov	(%r8), %r8b
	mov	(%r8), %r9b
	mov	(%r8), %r10b
	mov	(%r8), %r11b
	mov	(%r8), %r12b
	mov	(%r8), %r13b
	mov	(%r8), %r14b
	mov	(%r8), %r15b
	nop
	mov	(%r9), %al
	mov	(%r9), %cl
	mov	(%r9), %dl
	mov	(%r9), %bl
	mov	(%r9), %spl
	mov	(%r9), %bpl
	mov	(%r9), %sil
	mov	(%r9), %dil
	mov	(%r9), %r8b
	mov	(%r9), %r9b
	mov	(%r9), %r10b
	mov	(%r9), %r11b
	mov	(%r9), %r12b
	mov	(%r9), %r13b
	mov	(%r9), %r14b
	mov	(%r9), %r15b
	nop
	mov	(%r10),%al
	mov	(%r10),%cl
	mov	(%r10),%dl
	mov	(%r10),%bl
	mov	(%r10),%spl
	mov	(%r10),%bpl
	mov	(%r10),%sil
	mov	(%r10),%dil
	mov	(%r10),%r8b
	mov	(%r10),%r9b
	mov	(%r10),%r10b
	mov	(%r10),%r11b
	mov	(%r10),%r12b
	mov	(%r10),%r13b
	mov	(%r10),%r14b
	mov	(%r10),%r15b
	nop
	mov	(%r11),%al
	mov	(%r11),%cl
	mov	(%r11),%dl
	mov	(%r11),%bl
	mov	(%r11),%spl
	mov	(%r11),%bpl
	mov	(%r11),%sil
	mov	(%r11),%dil
	mov	(%r11),%r8b
	mov	(%r11),%r9b
	mov	(%r11),%r10b
	mov	(%r11),%r11b
	mov	(%r11),%r12b
	mov	(%r11),%r13b
	mov	(%r11),%r14b
	mov	(%r11),%r15b
	nop
	mov	(%r12),%al
	mov	(%r12),%cl
	mov	(%r12),%dl
	mov	(%r12),%bl
	mov	(%r12),%spl
	mov	(%r12),%bpl
	mov	(%r12),%sil
	mov	(%r12),%dil
	mov	(%r12),%r8b
	mov	(%r12),%r9b
	mov	(%r12),%r10b
	mov	(%r12),%r11b
	mov	(%r12),%r12b
	mov	(%r12),%r13b
	mov	(%r12),%r14b
	mov	(%r12),%r15b
	nop
	mov	(%r13),%al
	mov	(%r13),%cl
	mov	(%r13),%dl
	mov	(%r13),%bl
	mov	(%r13),%spl
	mov	(%r13),%bpl
	mov	(%r13),%sil
	mov	(%r13),%dil
	mov	(%r13),%r8b
	mov	(%r13),%r9b
	mov	(%r13),%r10b
	mov	(%r13),%r11b
	mov	(%r13),%r12b
	mov	(%r13),%r13b
	mov	(%r13),%r14b
	mov	(%r13),%r15b
	nop
	mov	(%r14),%al
	mov	(%r14),%cl
	mov	(%r14),%dl
	mov	(%r14),%bl
	mov	(%r14),%spl
	mov	(%r14),%bpl
	mov	(%r14),%sil
	mov	(%r14),%dil
	mov	(%r14),%r8b
	mov	(%r14),%r9b
	mov	(%r14),%r10b
	mov	(%r14),%r11b
	mov	(%r14),%r12b
	mov	(%r14),%r13b
	mov	(%r14),%r14b
	mov	(%r14),%r15b
	nop
	mov	(%r15),%al
	mov	(%r15),%cl
	mov	(%r15),%dl
	mov	(%r15),%bl
	mov	(%r15),%spl
	mov	(%r15),%bpl
	mov	(%r15),%sil
	mov	(%r15),%dil
	mov	(%r15),%r8b
	mov	(%r15),%r9b
	mov	(%r15),%r10b
	mov	(%r15),%r11b
	mov	(%r15),%r12b
	mov	(%r15),%r13b
	mov	(%r15),%r14b
	mov	(%r15),%r15b
        nop
        nop
        // mem8[off8] += reg8
	mov	0x7F(%rax),%al
	mov	0x7F(%rax),%cl
	mov	0x7F(%rax),%dl
	mov	0x7F(%rax),%bl
	mov	0x7F(%rax),%spl
	mov	0x7F(%rax),%bpl
	mov	0x7F(%rax),%sil
	mov	0x7F(%rax),%dil
	mov	0x7F(%rax),%r8b
	mov	0x7F(%rax),%r9b
	mov	0x7F(%rax),%r10b
	mov	0x7F(%rax),%r11b
	mov	0x7F(%rax),%r12b
	mov	0x7F(%rax),%r13b
	mov	0x7F(%rax),%r14b
	mov	0x7F(%rax),%r15b
	nop
	mov	0x7F(%rcx),%al
	mov	0x7F(%rcx),%cl
	mov	0x7F(%rcx),%dl
	mov	0x7F(%rcx),%bl
	mov	0x7F(%rcx),%spl
	mov	0x7F(%rcx),%bpl
	mov	0x7F(%rcx),%sil
	mov	0x7F(%rcx),%dil
	mov	0x7F(%rcx),%r8b
	mov	0x7F(%rcx),%r9b
	mov	0x7F(%rcx),%r10b
	mov	0x7F(%rcx),%r11b
	mov	0x7F(%rcx),%r12b
	mov	0x7F(%rcx),%r13b
	mov	0x7F(%rcx),%r14b
	mov	0x7F(%rcx),%r15b
	nop
	mov	0x7F(%rdx),%al
	mov	0x7F(%rdx),%cl
	mov	0x7F(%rdx),%dl
	mov	0x7F(%rdx),%bl
	mov	0x7F(%rdx),%spl
	mov	0x7F(%rdx),%bpl
	mov	0x7F(%rdx),%sil
	mov	0x7F(%rdx),%dil
	mov	0x7F(%rdx),%r8b
	mov	0x7F(%rdx),%r9b
	mov	0x7F(%rdx),%r10b
	mov	0x7F(%rdx),%r11b
	mov	0x7F(%rdx),%r12b
	mov	0x7F(%rdx),%r13b
	mov	0x7F(%rdx),%r14b
	mov	0x7F(%rdx),%r15b
	nop
	mov	0x7F(%rbx),%al
	mov	0x7F(%rbx),%cl
	mov	0x7F(%rbx),%dl
	mov	0x7F(%rbx),%bl
	mov	0x7F(%rbx),%spl
	mov	0x7F(%rbx),%bpl
	mov	0x7F(%rbx),%sil
	mov	0x7F(%rbx),%dil
	mov	0x7F(%rbx),%r8b
	mov	0x7F(%rbx),%r9b
	mov	0x7F(%rbx),%r10b
	mov	0x7F(%rbx),%r11b
	mov	0x7F(%rbx),%r12b
	mov	0x7F(%rbx),%r13b
	mov	0x7F(%rbx),%r14b
	mov	0x7F(%rbx),%r15b
	nop
	mov	0x7F(%rsp),%al
	mov	0x7F(%rsp),%cl
	mov	0x7F(%rsp),%dl
	mov	0x7F(%rsp),%bl
	mov	0x7F(%rsp),%spl
	mov	0x7F(%rsp),%bpl
	mov	0x7F(%rsp),%sil
	mov	0x7F(%rsp),%dil
	mov	0x7F(%rsp),%r8b
	mov	0x7F(%rsp),%r9b
	mov	0x7F(%rsp),%r10b
	mov	0x7F(%rsp),%r11b
	mov	0x7F(%rsp),%r12b
	mov	0x7F(%rsp),%r13b
	mov	0x7F(%rsp),%r14b
	mov	0x7F(%rsp),%r15b
	nop
	mov	0x7F(%rbp),%al
	mov	0x7F(%rbp),%cl
	mov	0x7F(%rbp),%dl
	mov	0x7F(%rbp),%bl
	mov	0x7F(%rbp),%spl
	mov	0x7F(%rbp),%bpl
	mov	0x7F(%rbp),%sil
	mov	0x7F(%rbp),%dil
	mov	0x7F(%rbp),%r8b
	mov	0x7F(%rbp),%r9b
	mov	0x7F(%rbp),%r10b
	mov	0x7F(%rbp),%r11b
	mov	0x7F(%rbp),%r12b
	mov	0x7F(%rbp),%r13b
	mov	0x7F(%rbp),%r14b
	mov	0x7F(%rbp),%r15b
	nop
	mov	0x7F(%rsi),%al
	mov	0x7F(%rsi),%cl
	mov	0x7F(%rsi),%dl
	mov	0x7F(%rsi),%bl
	mov	0x7F(%rsi),%spl
	mov	0x7F(%rsi),%bpl
	mov	0x7F(%rsi),%sil
	mov	0x7F(%rsi),%dil
	mov	0x7F(%rsi),%r8b
	mov	0x7F(%rsi),%r9b
	mov	0x7F(%rsi),%r10b
	mov	0x7F(%rsi),%r11b
	mov	0x7F(%rsi),%r12b
	mov	0x7F(%rsi),%r13b
	mov	0x7F(%rsi),%r14b
	mov	0x7F(%rsi),%r15b
	nop
	mov	0x7F(%rdi),%al
	mov	0x7F(%rdi),%cl
	mov	0x7F(%rdi),%dl
	mov	0x7F(%rdi),%bl
	mov	0x7F(%rdi),%spl
	mov	0x7F(%rdi),%bpl
	mov	0x7F(%rdi),%sil
	mov	0x7F(%rdi),%dil
	mov	0x7F(%rdi),%r8b
	mov	0x7F(%rdi),%r9b
	mov	0x7F(%rdi),%r10b
	mov	0x7F(%rdi),%r11b
	mov	0x7F(%rdi),%r12b
	mov	0x7F(%rdi),%r13b
	mov	0x7F(%rdi),%r14b
	mov	0x7F(%rdi),%r15b
	nop
	mov	0x7F(%r8), %al
	mov	0x7F(%r8), %cl
	mov	0x7F(%r8), %dl
	mov	0x7F(%r8), %bl
	mov	0x7F(%r8), %spl
	mov	0x7F(%r8), %bpl
	mov	0x7F(%r8), %sil
	mov	0x7F(%r8), %dil
	mov	0x7F(%r8), %r8b
	mov	0x7F(%r8), %r9b
	mov	0x7F(%r8), %r10b
	mov	0x7F(%r8), %r11b
	mov	0x7F(%r8), %r12b
	mov	0x7F(%r8), %r13b
	mov	0x7F(%r8), %r14b
	mov	0x7F(%r8), %r15b
	nop
	mov	0x7F(%r9), %al
	mov	0x7F(%r9), %cl
	mov	0x7F(%r9), %dl
	mov	0x7F(%r9), %bl
	mov	0x7F(%r9), %spl
	mov	0x7F(%r9), %bpl
	mov	0x7F(%r9), %sil
	mov	0x7F(%r9), %dil
	mov	0x7F(%r9), %r8b
	mov	0x7F(%r9), %r9b
	mov	0x7F(%r9), %r10b
	mov	0x7F(%r9), %r11b
	mov	0x7F(%r9), %r12b
	mov	0x7F(%r9), %r13b
	mov	0x7F(%r9), %r14b
	mov	0x7F(%r9), %r15b
	nop
	mov	0x7F(%r10),%al
	mov	0x7F(%r10),%cl
	mov	0x7F(%r10),%dl
	mov	0x7F(%r10),%bl
	mov	0x7F(%r10),%spl
	mov	0x7F(%r10),%bpl
	mov	0x7F(%r10),%sil
	mov	0x7F(%r10),%dil
	mov	0x7F(%r10),%r8b
	mov	0x7F(%r10),%r9b
	mov	0x7F(%r10),%r10b
	mov	0x7F(%r10),%r11b
	mov	0x7F(%r10),%r12b
	mov	0x7F(%r10),%r13b
	mov	0x7F(%r10),%r14b
	mov	0x7F(%r10),%r15b
	nop
	mov	0x7F(%r11),%al
	mov	0x7F(%r11),%cl
	mov	0x7F(%r11),%dl
	mov	0x7F(%r11),%bl
	mov	0x7F(%r11),%spl
	mov	0x7F(%r11),%bpl
	mov	0x7F(%r11),%sil
	mov	0x7F(%r11),%dil
	mov	0x7F(%r11),%r8b
	mov	0x7F(%r11),%r9b
	mov	0x7F(%r11),%r10b
	mov	0x7F(%r11),%r11b
	mov	0x7F(%r11),%r12b
	mov	0x7F(%r11),%r13b
	mov	0x7F(%r11),%r14b
	mov	0x7F(%r11),%r15b
	nop
	mov	0x7F(%r12),%al
	mov	0x7F(%r12),%cl
	mov	0x7F(%r12),%dl
	mov	0x7F(%r12),%bl
	mov	0x7F(%r12),%spl
	mov	0x7F(%r12),%bpl
	mov	0x7F(%r12),%sil
	mov	0x7F(%r12),%dil
	mov	0x7F(%r12),%r8b
	mov	0x7F(%r12),%r9b
	mov	0x7F(%r12),%r10b
	mov	0x7F(%r12),%r11b
	mov	0x7F(%r12),%r12b
	mov	0x7F(%r12),%r13b
	mov	0x7F(%r12),%r14b
	mov	0x7F(%r12),%r15b
	nop
	mov	0x7F(%r13),%al
	mov	0x7F(%r13),%cl
	mov	0x7F(%r13),%dl
	mov	0x7F(%r13),%bl
	mov	0x7F(%r13),%spl
	mov	0x7F(%r13),%bpl
	mov	0x7F(%r13),%sil
	mov	0x7F(%r13),%dil
	mov	0x7F(%r13),%r8b
	mov	0x7F(%r13),%r9b
	mov	0x7F(%r13),%r10b
	mov	0x7F(%r13),%r11b
	mov	0x7F(%r13),%r12b
	mov	0x7F(%r13),%r13b
	mov	0x7F(%r13),%r14b
	mov	0x7F(%r13),%r15b
	nop
	mov	0x7F(%r14),%al
	mov	0x7F(%r14),%cl
	mov	0x7F(%r14),%dl
	mov	0x7F(%r14),%bl
	mov	0x7F(%r14),%spl
	mov	0x7F(%r14),%bpl
	mov	0x7F(%r14),%sil
	mov	0x7F(%r14),%dil
	mov	0x7F(%r14),%r8b
	mov	0x7F(%r14),%r9b
	mov	0x7F(%r14),%r10b
	mov	0x7F(%r14),%r11b
	mov	0x7F(%r14),%r12b
	mov	0x7F(%r14),%r13b
	mov	0x7F(%r14),%r14b
	mov	0x7F(%r14),%r15b
	nop
	mov	0x7F(%r15),%al
	mov	0x7F(%r15),%cl
	mov	0x7F(%r15),%dl
	mov	0x7F(%r15),%bl
	mov	0x7F(%r15),%spl
	mov	0x7F(%r15),%bpl
	mov	0x7F(%r15),%sil
	mov	0x7F(%r15),%dil
	mov	0x7F(%r15),%r8b
	mov	0x7F(%r15),%r9b
	mov	0x7F(%r15),%r10b
	mov	0x7F(%r15),%r11b
	mov	0x7F(%r15),%r12b
	mov	0x7F(%r15),%r13b
	mov	0x7F(%r15),%r14b
	mov	0x7F(%r15),%r15b
        nop
        nop
        // mem8[off32] += reg8
	mov	0x12345678(%rax),%al
	mov	0x12345678(%rax),%cl
	mov	0x12345678(%rax),%dl
	mov	0x12345678(%rax),%bl
	mov	0x12345678(%rax),%spl
	mov	0x12345678(%rax),%bpl
	mov	0x12345678(%rax),%sil
	mov	0x12345678(%rax),%dil
	mov	0x12345678(%rax),%r8b
	mov	0x12345678(%rax),%r9b
	mov	0x12345678(%rax),%r10b
	mov	0x12345678(%rax),%r11b
	mov	0x12345678(%rax),%r12b
	mov	0x12345678(%rax),%r13b
	mov	0x12345678(%rax),%r14b
	mov	0x12345678(%rax),%r15b
	nop
	mov	0x12345678(%rcx),%al
	mov	0x12345678(%rcx),%cl
	mov	0x12345678(%rcx),%dl
	mov	0x12345678(%rcx),%bl
	mov	0x12345678(%rcx),%spl
	mov	0x12345678(%rcx),%bpl
	mov	0x12345678(%rcx),%sil
	mov	0x12345678(%rcx),%dil
	mov	0x12345678(%rcx),%r8b
	mov	0x12345678(%rcx),%r9b
	mov	0x12345678(%rcx),%r10b
	mov	0x12345678(%rcx),%r11b
	mov	0x12345678(%rcx),%r12b
	mov	0x12345678(%rcx),%r13b
	mov	0x12345678(%rcx),%r14b
	mov	0x12345678(%rcx),%r15b
	nop
	mov	0x12345678(%rdx),%al
	mov	0x12345678(%rdx),%cl
	mov	0x12345678(%rdx),%dl
	mov	0x12345678(%rdx),%bl
	mov	0x12345678(%rdx),%spl
	mov	0x12345678(%rdx),%bpl
	mov	0x12345678(%rdx),%sil
	mov	0x12345678(%rdx),%dil
	mov	0x12345678(%rdx),%r8b
	mov	0x12345678(%rdx),%r9b
	mov	0x12345678(%rdx),%r10b
	mov	0x12345678(%rdx),%r11b
	mov	0x12345678(%rdx),%r12b
	mov	0x12345678(%rdx),%r13b
	mov	0x12345678(%rdx),%r14b
	mov	0x12345678(%rdx),%r15b
	nop
	mov	0x12345678(%rbx),%al
	mov	0x12345678(%rbx),%cl
	mov	0x12345678(%rbx),%dl
	mov	0x12345678(%rbx),%bl
	mov	0x12345678(%rbx),%spl
	mov	0x12345678(%rbx),%bpl
	mov	0x12345678(%rbx),%sil
	mov	0x12345678(%rbx),%dil
	mov	0x12345678(%rbx),%r8b
	mov	0x12345678(%rbx),%r9b
	mov	0x12345678(%rbx),%r10b
	mov	0x12345678(%rbx),%r11b
	mov	0x12345678(%rbx),%r12b
	mov	0x12345678(%rbx),%r13b
	mov	0x12345678(%rbx),%r14b
	mov	0x12345678(%rbx),%r15b
	nop
	mov	0x12345678(%rsp),%al
	mov	0x12345678(%rsp),%cl
	mov	0x12345678(%rsp),%dl
	mov	0x12345678(%rsp),%bl
	mov	0x12345678(%rsp),%spl
	mov	0x12345678(%rsp),%bpl
	mov	0x12345678(%rsp),%sil
	mov	0x12345678(%rsp),%dil
	mov	0x12345678(%rsp),%r8b
	mov	0x12345678(%rsp),%r9b
	mov	0x12345678(%rsp),%r10b
	mov	0x12345678(%rsp),%r11b
	mov	0x12345678(%rsp),%r12b
	mov	0x12345678(%rsp),%r13b
	mov	0x12345678(%rsp),%r14b
	mov	0x12345678(%rsp),%r15b
	nop
	mov	0x12345678(%rbp),%al
	mov	0x12345678(%rbp),%cl
	mov	0x12345678(%rbp),%dl
	mov	0x12345678(%rbp),%bl
	mov	0x12345678(%rbp),%spl
	mov	0x12345678(%rbp),%bpl
	mov	0x12345678(%rbp),%sil
	mov	0x12345678(%rbp),%dil
	mov	0x12345678(%rbp),%r8b
	mov	0x12345678(%rbp),%r9b
	mov	0x12345678(%rbp),%r10b
	mov	0x12345678(%rbp),%r11b
	mov	0x12345678(%rbp),%r12b
	mov	0x12345678(%rbp),%r13b
	mov	0x12345678(%rbp),%r14b
	mov	0x12345678(%rbp),%r15b
	nop
	mov	0x12345678(%rsi),%al
	mov	0x12345678(%rsi),%cl
	mov	0x12345678(%rsi),%dl
	mov	0x12345678(%rsi),%bl
	mov	0x12345678(%rsi),%spl
	mov	0x12345678(%rsi),%bpl
	mov	0x12345678(%rsi),%sil
	mov	0x12345678(%rsi),%dil
	mov	0x12345678(%rsi),%r8b
	mov	0x12345678(%rsi),%r9b
	mov	0x12345678(%rsi),%r10b
	mov	0x12345678(%rsi),%r11b
	mov	0x12345678(%rsi),%r12b
	mov	0x12345678(%rsi),%r13b
	mov	0x12345678(%rsi),%r14b
	mov	0x12345678(%rsi),%r15b
	nop
	mov	0x12345678(%rdi),%al
	mov	0x12345678(%rdi),%cl
	mov	0x12345678(%rdi),%dl
	mov	0x12345678(%rdi),%bl
	mov	0x12345678(%rdi),%spl
	mov	0x12345678(%rdi),%bpl
	mov	0x12345678(%rdi),%sil
	mov	0x12345678(%rdi),%dil
	mov	0x12345678(%rdi),%r8b
	mov	0x12345678(%rdi),%r9b
	mov	0x12345678(%rdi),%r10b
	mov	0x12345678(%rdi),%r11b
	mov	0x12345678(%rdi),%r12b
	mov	0x12345678(%rdi),%r13b
	mov	0x12345678(%rdi),%r14b
	mov	0x12345678(%rdi),%r15b
	nop
	mov	0x12345678(%r8), %al
	mov	0x12345678(%r8), %cl
	mov	0x12345678(%r8), %dl
	mov	0x12345678(%r8), %bl
	mov	0x12345678(%r8), %spl
	mov	0x12345678(%r8), %bpl
	mov	0x12345678(%r8), %sil
	mov	0x12345678(%r8), %dil
	mov	0x12345678(%r8), %r8b
	mov	0x12345678(%r8), %r9b
	mov	0x12345678(%r8), %r10b
	mov	0x12345678(%r8), %r11b
	mov	0x12345678(%r8), %r12b
	mov	0x12345678(%r8), %r13b
	mov	0x12345678(%r8), %r14b
	mov	0x12345678(%r8), %r15b
	nop
	mov	0x12345678(%r9), %al
	mov	0x12345678(%r9), %cl
	mov	0x12345678(%r9), %dl
	mov	0x12345678(%r9), %bl
	mov	0x12345678(%r9), %spl
	mov	0x12345678(%r9), %bpl
	mov	0x12345678(%r9), %sil
	mov	0x12345678(%r9), %dil
	mov	0x12345678(%r9), %r8b
	mov	0x12345678(%r9), %r9b
	mov	0x12345678(%r9), %r10b
	mov	0x12345678(%r9), %r11b
	mov	0x12345678(%r9), %r12b
	mov	0x12345678(%r9), %r13b
	mov	0x12345678(%r9), %r14b
	mov	0x12345678(%r9), %r15b
	nop
	mov	0x12345678(%r10),%al
	mov	0x12345678(%r10),%cl
	mov	0x12345678(%r10),%dl
	mov	0x12345678(%r10),%bl
	mov	0x12345678(%r10),%spl
	mov	0x12345678(%r10),%bpl
	mov	0x12345678(%r10),%sil
	mov	0x12345678(%r10),%dil
	mov	0x12345678(%r10),%r8b
	mov	0x12345678(%r10),%r9b
	mov	0x12345678(%r10),%r10b
	mov	0x12345678(%r10),%r11b
	mov	0x12345678(%r10),%r12b
	mov	0x12345678(%r10),%r13b
	mov	0x12345678(%r10),%r14b
	mov	0x12345678(%r10),%r15b
	nop
	mov	0x12345678(%r11),%al
	mov	0x12345678(%r11),%cl
	mov	0x12345678(%r11),%dl
	mov	0x12345678(%r11),%bl
	mov	0x12345678(%r11),%spl
	mov	0x12345678(%r11),%bpl
	mov	0x12345678(%r11),%sil
	mov	0x12345678(%r11),%dil
	mov	0x12345678(%r11),%r8b
	mov	0x12345678(%r11),%r9b
	mov	0x12345678(%r11),%r10b
	mov	0x12345678(%r11),%r11b
	mov	0x12345678(%r11),%r12b
	mov	0x12345678(%r11),%r13b
	mov	0x12345678(%r11),%r14b
	mov	0x12345678(%r11),%r15b
	nop
	mov	0x12345678(%r12),%al
	mov	0x12345678(%r12),%cl
	mov	0x12345678(%r12),%dl
	mov	0x12345678(%r12),%bl
	mov	0x12345678(%r12),%spl
	mov	0x12345678(%r12),%bpl
	mov	0x12345678(%r12),%sil
	mov	0x12345678(%r12),%dil
	mov	0x12345678(%r12),%r8b
	mov	0x12345678(%r12),%r9b
	mov	0x12345678(%r12),%r10b
	mov	0x12345678(%r12),%r11b
	mov	0x12345678(%r12),%r12b
	mov	0x12345678(%r12),%r13b
	mov	0x12345678(%r12),%r14b
	mov	0x12345678(%r12),%r15b
	nop
	mov	0x12345678(%r13),%al
	mov	0x12345678(%r13),%cl
	mov	0x12345678(%r13),%dl
	mov	0x12345678(%r13),%bl
	mov	0x12345678(%r13),%spl
	mov	0x12345678(%r13),%bpl
	mov	0x12345678(%r13),%sil
	mov	0x12345678(%r13),%dil
	mov	0x12345678(%r13),%r8b
	mov	0x12345678(%r13),%r9b
	mov	0x12345678(%r13),%r10b
	mov	0x12345678(%r13),%r11b
	mov	0x12345678(%r13),%r12b
	mov	0x12345678(%r13),%r13b
	mov	0x12345678(%r13),%r14b
	mov	0x12345678(%r13),%r15b
	nop
	mov	0x12345678(%r14),%al
	mov	0x12345678(%r14),%cl
	mov	0x12345678(%r14),%dl
	mov	0x12345678(%r14),%bl
	mov	0x12345678(%r14),%spl
	mov	0x12345678(%r14),%bpl
	mov	0x12345678(%r14),%sil
	mov	0x12345678(%r14),%dil
	mov	0x12345678(%r14),%r8b
	mov	0x12345678(%r14),%r9b
	mov	0x12345678(%r14),%r10b
	mov	0x12345678(%r14),%r11b
	mov	0x12345678(%r14),%r12b
	mov	0x12345678(%r14),%r13b
	mov	0x12345678(%r14),%r14b
	mov	0x12345678(%r14),%r15b
	nop
	mov	0x12345678(%r15),%al
	mov	0x12345678(%r15),%cl
	mov	0x12345678(%r15),%dl
	mov	0x12345678(%r15),%bl
	mov	0x12345678(%r15),%spl
	mov	0x12345678(%r15),%bpl
	mov	0x12345678(%r15),%sil
	mov	0x12345678(%r15),%dil
	mov	0x12345678(%r15),%r8b
	mov	0x12345678(%r15),%r9b
	mov	0x12345678(%r15),%r10b
	mov	0x12345678(%r15),%r11b
	mov	0x12345678(%r15),%r12b
	mov	0x12345678(%r15),%r13b
	mov	0x12345678(%r15),%r14b
	mov	0x12345678(%r15),%r15b
        ret
	.cfi_endproc
        
	.p2align 4,,15
	.globl	MovRegMem16
	.type	MovRegMem16, @function
MovRegMem16:
	.cfi_startproc
        // reg16 = mem16[0]
	mov	(%rax),%ax
	mov	(%rax),%cx
	mov	(%rax),%dx
	mov	(%rax),%bx
	mov	(%rax),%sp
	mov	(%rax),%bp
	mov	(%rax),%si
	mov	(%rax),%di
	mov	(%rax),%r8w
	mov	(%rax),%r9w
	mov	(%rax),%r10w
	mov	(%rax),%r11w
	mov	(%rax),%r12w
	mov	(%rax),%r13w
	mov	(%rax),%r14w
	mov	(%rax),%r15w
	nop
	mov	(%rcx),%ax
	mov	(%rcx),%cx
	mov	(%rcx),%dx
	mov	(%rcx),%bx
	mov	(%rcx),%sp
	mov	(%rcx),%bp
	mov	(%rcx),%si
	mov	(%rcx),%di
	mov	(%rcx),%r8w
	mov	(%rcx),%r9w
	mov	(%rcx),%r10w
	mov	(%rcx),%r11w
	mov	(%rcx),%r12w
	mov	(%rcx),%r13w
	mov	(%rcx),%r14w
	mov	(%rcx),%r15w
	nop
	mov	(%rdx),%ax
	mov	(%rdx),%cx
	mov	(%rdx),%dx
	mov	(%rdx),%bx
	mov	(%rdx),%sp
	mov	(%rdx),%bp
	mov	(%rdx),%si
	mov	(%rdx),%di
	mov	(%rdx),%r8w
	mov	(%rdx),%r9w
	mov	(%rdx),%r10w
	mov	(%rdx),%r11w
	mov	(%rdx),%r12w
	mov	(%rdx),%r13w
	mov	(%rdx),%r14w
	mov	(%rdx),%r15w
	nop
	mov	(%rbx),%ax
	mov	(%rbx),%cx
	mov	(%rbx),%dx
	mov	(%rbx),%bx
	mov	(%rbx),%sp
	mov	(%rbx),%bp
	mov	(%rbx),%si
	mov	(%rbx),%di
	mov	(%rbx),%r8w
	mov	(%rbx),%r9w
	mov	(%rbx),%r10w
	mov	(%rbx),%r11w
	mov	(%rbx),%r12w
	mov	(%rbx),%r13w
	mov	(%rbx),%r14w
	mov	(%rbx),%r15w
	nop
	mov	(%rsp),%ax
	mov	(%rsp),%cx
	mov	(%rsp),%dx
	mov	(%rsp),%bx
	mov	(%rsp),%sp
	mov	(%rsp),%bp
	mov	(%rsp),%si
	mov	(%rsp),%di
	mov	(%rsp),%r8w
	mov	(%rsp),%r9w
	mov	(%rsp),%r10w
	mov	(%rsp),%r11w
	mov	(%rsp),%r12w
	mov	(%rsp),%r13w
	mov	(%rsp),%r14w
	mov	(%rsp),%r15w
	nop
	mov	(%rbp),%ax
	mov	(%rbp),%cx
	mov	(%rbp),%dx
	mov	(%rbp),%bx
	mov	(%rbp),%sp
	mov	(%rbp),%bp
	mov	(%rbp),%si
	mov	(%rbp),%di
	mov	(%rbp),%r8w
	mov	(%rbp),%r9w
	mov	(%rbp),%r10w
	mov	(%rbp),%r11w
	mov	(%rbp),%r12w
	mov	(%rbp),%r13w
	mov	(%rbp),%r14w
	mov	(%rbp),%r15w
	nop
	mov	(%rsi),%ax
	mov	(%rsi),%cx
	mov	(%rsi),%dx
	mov	(%rsi),%bx
	mov	(%rsi),%sp
	mov	(%rsi),%bp
	mov	(%rsi),%si
	mov	(%rsi),%di
	mov	(%rsi),%r8w
	mov	(%rsi),%r9w
	mov	(%rsi),%r10w
	mov	(%rsi),%r11w
	mov	(%rsi),%r12w
	mov	(%rsi),%r13w
	mov	(%rsi),%r14w
	mov	(%rsi),%r15w
	nop
	mov	(%rdi),%ax
	mov	(%rdi),%cx
	mov	(%rdi),%dx
	mov	(%rdi),%bx
	mov	(%rdi),%sp
	mov	(%rdi),%bp
	mov	(%rdi),%si
	mov	(%rdi),%di
	mov	(%rdi),%r8w
	mov	(%rdi),%r9w
	mov	(%rdi),%r10w
	mov	(%rdi),%r11w
	mov	(%rdi),%r12w
	mov	(%rdi),%r13w
	mov	(%rdi),%r14w
	mov	(%rdi),%r15w
	nop
	mov	(%r8), %ax
	mov	(%r8), %cx
	mov	(%r8), %dx
	mov	(%r8), %bx
	mov	(%r8), %sp
	mov	(%r8), %bp
	mov	(%r8), %si
	mov	(%r8), %di
	mov	(%r8), %r8w
	mov	(%r8), %r9w
	mov	(%r8), %r10w
	mov	(%r8), %r11w
	mov	(%r8), %r12w
	mov	(%r8), %r13w
	mov	(%r8), %r14w
	mov	(%r8), %r15w
	nop
	mov	(%r9), %ax
	mov	(%r9), %cx
	mov	(%r9), %dx
	mov	(%r9), %bx
	mov	(%r9), %sp
	mov	(%r9), %bp
	mov	(%r9), %si
	mov	(%r9), %di
	mov	(%r9), %r8w
	mov	(%r9), %r9w
	mov	(%r9), %r10w
	mov	(%r9), %r11w
	mov	(%r9), %r12w
	mov	(%r9), %r13w
	mov	(%r9), %r14w
	mov	(%r9), %r15w
	nop
	mov	(%r10),%ax
	mov	(%r10),%cx
	mov	(%r10),%dx
	mov	(%r10),%bx
	mov	(%r10),%sp
	mov	(%r10),%bp
	mov	(%r10),%si
	mov	(%r10),%di
	mov	(%r10),%r8w
	mov	(%r10),%r9w
	mov	(%r10),%r10w
	mov	(%r10),%r11w
	mov	(%r10),%r12w
	mov	(%r10),%r13w
	mov	(%r10),%r14w
	mov	(%r10),%r15w
	nop
	mov	(%r11),%ax
	mov	(%r11),%cx
	mov	(%r11),%dx
	mov	(%r11),%bx
	mov	(%r11),%sp
	mov	(%r11),%bp
	mov	(%r11),%si
	mov	(%r11),%di
	mov	(%r11),%r8w
	mov	(%r11),%r9w
	mov	(%r11),%r10w
	mov	(%r11),%r11w
	mov	(%r11),%r12w
	mov	(%r11),%r13w
	mov	(%r11),%r14w
	mov	(%r11),%r15w
	nop
	mov	(%r12),%ax
	mov	(%r12),%cx
	mov	(%r12),%dx
	mov	(%r12),%bx
	mov	(%r12),%sp
	mov	(%r12),%bp
	mov	(%r12),%si
	mov	(%r12),%di
	mov	(%r12),%r8w
	mov	(%r12),%r9w
	mov	(%r12),%r10w
	mov	(%r12),%r11w
	mov	(%r12),%r12w
	mov	(%r12),%r13w
	mov	(%r12),%r14w
	mov	(%r12),%r15w
	nop
	mov	(%r13),%ax
	mov	(%r13),%cx
	mov	(%r13),%dx
	mov	(%r13),%bx
	mov	(%r13),%sp
	mov	(%r13),%bp
	mov	(%r13),%si
	mov	(%r13),%di
	mov	(%r13),%r8w
	mov	(%r13),%r9w
	mov	(%r13),%r10w
	mov	(%r13),%r11w
	mov	(%r13),%r12w
	mov	(%r13),%r13w
	mov	(%r13),%r14w
	mov	(%r13),%r15w
	nop
	mov	(%r14),%ax
	mov	(%r14),%cx
	mov	(%r14),%dx
	mov	(%r14),%bx
	mov	(%r14),%sp
	mov	(%r14),%bp
	mov	(%r14),%si
	mov	(%r14),%di
	mov	(%r14),%r8w
	mov	(%r14),%r9w
	mov	(%r14),%r10w
	mov	(%r14),%r11w
	mov	(%r14),%r12w
	mov	(%r14),%r13w
	mov	(%r14),%r14w
	mov	(%r14),%r15w
	nop
	mov	(%r15),%ax
	mov	(%r15),%cx
	mov	(%r15),%dx
	mov	(%r15),%bx
	mov	(%r15),%sp
	mov	(%r15),%bp
	mov	(%r15),%si
	mov	(%r15),%di
	mov	(%r15),%r8w
	mov	(%r15),%r9w
	mov	(%r15),%r10w
	mov	(%r15),%r11w
	mov	(%r15),%r12w
	mov	(%r15),%r13w
	mov	(%r15),%r14w
	mov	(%r15),%r15w
        nop
        nop
        // mem16[off8] += reg16
	mov	0x7F(%rax),%ax
	mov	0x7F(%rax),%cx
	mov	0x7F(%rax),%dx
	mov	0x7F(%rax),%bx
	mov	0x7F(%rax),%sp
	mov	0x7F(%rax),%bp
	mov	0x7F(%rax),%si
	mov	0x7F(%rax),%di
	mov	0x7F(%rax),%r8w
	mov	0x7F(%rax),%r9w
	mov	0x7F(%rax),%r10w
	mov	0x7F(%rax),%r11w
	mov	0x7F(%rax),%r12w
	mov	0x7F(%rax),%r13w
	mov	0x7F(%rax),%r14w
	mov	0x7F(%rax),%r15w
	nop
	mov	0x7F(%rcx),%ax
	mov	0x7F(%rcx),%cx
	mov	0x7F(%rcx),%dx
	mov	0x7F(%rcx),%bx
	mov	0x7F(%rcx),%sp
	mov	0x7F(%rcx),%bp
	mov	0x7F(%rcx),%si
	mov	0x7F(%rcx),%di
	mov	0x7F(%rcx),%r8w
	mov	0x7F(%rcx),%r9w
	mov	0x7F(%rcx),%r10w
	mov	0x7F(%rcx),%r11w
	mov	0x7F(%rcx),%r12w
	mov	0x7F(%rcx),%r13w
	mov	0x7F(%rcx),%r14w
	mov	0x7F(%rcx),%r15w
	nop
	mov	0x7F(%rdx),%ax
	mov	0x7F(%rdx),%cx
	mov	0x7F(%rdx),%dx
	mov	0x7F(%rdx),%bx
	mov	0x7F(%rdx),%sp
	mov	0x7F(%rdx),%bp
	mov	0x7F(%rdx),%si
	mov	0x7F(%rdx),%di
	mov	0x7F(%rdx),%r8w
	mov	0x7F(%rdx),%r9w
	mov	0x7F(%rdx),%r10w
	mov	0x7F(%rdx),%r11w
	mov	0x7F(%rdx),%r12w
	mov	0x7F(%rdx),%r13w
	mov	0x7F(%rdx),%r14w
	mov	0x7F(%rdx),%r15w
	nop
	mov	0x7F(%rbx),%ax
	mov	0x7F(%rbx),%cx
	mov	0x7F(%rbx),%dx
	mov	0x7F(%rbx),%bx
	mov	0x7F(%rbx),%sp
	mov	0x7F(%rbx),%bp
	mov	0x7F(%rbx),%si
	mov	0x7F(%rbx),%di
	mov	0x7F(%rbx),%r8w
	mov	0x7F(%rbx),%r9w
	mov	0x7F(%rbx),%r10w
	mov	0x7F(%rbx),%r11w
	mov	0x7F(%rbx),%r12w
	mov	0x7F(%rbx),%r13w
	mov	0x7F(%rbx),%r14w
	mov	0x7F(%rbx),%r15w
	nop
	mov	0x7F(%rsp),%ax
	mov	0x7F(%rsp),%cx
	mov	0x7F(%rsp),%dx
	mov	0x7F(%rsp),%bx
	mov	0x7F(%rsp),%sp
	mov	0x7F(%rsp),%bp
	mov	0x7F(%rsp),%si
	mov	0x7F(%rsp),%di
	mov	0x7F(%rsp),%r8w
	mov	0x7F(%rsp),%r9w
	mov	0x7F(%rsp),%r10w
	mov	0x7F(%rsp),%r11w
	mov	0x7F(%rsp),%r12w
	mov	0x7F(%rsp),%r13w
	mov	0x7F(%rsp),%r14w
	mov	0x7F(%rsp),%r15w
	nop
	mov	0x7F(%rbp),%ax
	mov	0x7F(%rbp),%cx
	mov	0x7F(%rbp),%dx
	mov	0x7F(%rbp),%bx
	mov	0x7F(%rbp),%sp
	mov	0x7F(%rbp),%bp
	mov	0x7F(%rbp),%si
	mov	0x7F(%rbp),%di
	mov	0x7F(%rbp),%r8w
	mov	0x7F(%rbp),%r9w
	mov	0x7F(%rbp),%r10w
	mov	0x7F(%rbp),%r11w
	mov	0x7F(%rbp),%r12w
	mov	0x7F(%rbp),%r13w
	mov	0x7F(%rbp),%r14w
	mov	0x7F(%rbp),%r15w
	nop
	mov	0x7F(%rsi),%ax
	mov	0x7F(%rsi),%cx
	mov	0x7F(%rsi),%dx
	mov	0x7F(%rsi),%bx
	mov	0x7F(%rsi),%sp
	mov	0x7F(%rsi),%bp
	mov	0x7F(%rsi),%si
	mov	0x7F(%rsi),%di
	mov	0x7F(%rsi),%r8w
	mov	0x7F(%rsi),%r9w
	mov	0x7F(%rsi),%r10w
	mov	0x7F(%rsi),%r11w
	mov	0x7F(%rsi),%r12w
	mov	0x7F(%rsi),%r13w
	mov	0x7F(%rsi),%r14w
	mov	0x7F(%rsi),%r15w
	nop
	mov	0x7F(%rdi),%ax
	mov	0x7F(%rdi),%cx
	mov	0x7F(%rdi),%dx
	mov	0x7F(%rdi),%bx
	mov	0x7F(%rdi),%sp
	mov	0x7F(%rdi),%bp
	mov	0x7F(%rdi),%si
	mov	0x7F(%rdi),%di
	mov	0x7F(%rdi),%r8w
	mov	0x7F(%rdi),%r9w
	mov	0x7F(%rdi),%r10w
	mov	0x7F(%rdi),%r11w
	mov	0x7F(%rdi),%r12w
	mov	0x7F(%rdi),%r13w
	mov	0x7F(%rdi),%r14w
	mov	0x7F(%rdi),%r15w
	nop
	mov	0x7F(%r8), %ax
	mov	0x7F(%r8), %cx
	mov	0x7F(%r8), %dx
	mov	0x7F(%r8), %bx
	mov	0x7F(%r8), %sp
	mov	0x7F(%r8), %bp
	mov	0x7F(%r8), %si
	mov	0x7F(%r8), %di
	mov	0x7F(%r8), %r8w
	mov	0x7F(%r8), %r9w
	mov	0x7F(%r8), %r10w
	mov	0x7F(%r8), %r11w
	mov	0x7F(%r8), %r12w
	mov	0x7F(%r8), %r13w
	mov	0x7F(%r8), %r14w
	mov	0x7F(%r8), %r15w
	nop
	mov	0x7F(%r9), %ax
	mov	0x7F(%r9), %cx
	mov	0x7F(%r9), %dx
	mov	0x7F(%r9), %bx
	mov	0x7F(%r9), %sp
	mov	0x7F(%r9), %bp
	mov	0x7F(%r9), %si
	mov	0x7F(%r9), %di
	mov	0x7F(%r9), %r8w
	mov	0x7F(%r9), %r9w
	mov	0x7F(%r9), %r10w
	mov	0x7F(%r9), %r11w
	mov	0x7F(%r9), %r12w
	mov	0x7F(%r9), %r13w
	mov	0x7F(%r9), %r14w
	mov	0x7F(%r9), %r15w
	nop
	mov	0x7F(%r10),%ax
	mov	0x7F(%r10),%cx
	mov	0x7F(%r10),%dx
	mov	0x7F(%r10),%bx
	mov	0x7F(%r10),%sp
	mov	0x7F(%r10),%bp
	mov	0x7F(%r10),%si
	mov	0x7F(%r10),%di
	mov	0x7F(%r10),%r8w
	mov	0x7F(%r10),%r9w
	mov	0x7F(%r10),%r10w
	mov	0x7F(%r10),%r11w
	mov	0x7F(%r10),%r12w
	mov	0x7F(%r10),%r13w
	mov	0x7F(%r10),%r14w
	mov	0x7F(%r10),%r15w
	nop
	mov	0x7F(%r11),%ax
	mov	0x7F(%r11),%cx
	mov	0x7F(%r11),%dx
	mov	0x7F(%r11),%bx
	mov	0x7F(%r11),%sp
	mov	0x7F(%r11),%bp
	mov	0x7F(%r11),%si
	mov	0x7F(%r11),%di
	mov	0x7F(%r11),%r8w
	mov	0x7F(%r11),%r9w
	mov	0x7F(%r11),%r10w
	mov	0x7F(%r11),%r11w
	mov	0x7F(%r11),%r12w
	mov	0x7F(%r11),%r13w
	mov	0x7F(%r11),%r14w
	mov	0x7F(%r11),%r15w
	nop
	mov	0x7F(%r12),%ax
	mov	0x7F(%r12),%cx
	mov	0x7F(%r12),%dx
	mov	0x7F(%r12),%bx
	mov	0x7F(%r12),%sp
	mov	0x7F(%r12),%bp
	mov	0x7F(%r12),%si
	mov	0x7F(%r12),%di
	mov	0x7F(%r12),%r8w
	mov	0x7F(%r12),%r9w
	mov	0x7F(%r12),%r10w
	mov	0x7F(%r12),%r11w
	mov	0x7F(%r12),%r12w
	mov	0x7F(%r12),%r13w
	mov	0x7F(%r12),%r14w
	mov	0x7F(%r12),%r15w
	nop
	mov	0x7F(%r13),%ax
	mov	0x7F(%r13),%cx
	mov	0x7F(%r13),%dx
	mov	0x7F(%r13),%bx
	mov	0x7F(%r13),%sp
	mov	0x7F(%r13),%bp
	mov	0x7F(%r13),%si
	mov	0x7F(%r13),%di
	mov	0x7F(%r13),%r8w
	mov	0x7F(%r13),%r9w
	mov	0x7F(%r13),%r10w
	mov	0x7F(%r13),%r11w
	mov	0x7F(%r13),%r12w
	mov	0x7F(%r13),%r13w
	mov	0x7F(%r13),%r14w
	mov	0x7F(%r13),%r15w
	nop
	mov	0x7F(%r14),%ax
	mov	0x7F(%r14),%cx
	mov	0x7F(%r14),%dx
	mov	0x7F(%r14),%bx
	mov	0x7F(%r14),%sp
	mov	0x7F(%r14),%bp
	mov	0x7F(%r14),%si
	mov	0x7F(%r14),%di
	mov	0x7F(%r14),%r8w
	mov	0x7F(%r14),%r9w
	mov	0x7F(%r14),%r10w
	mov	0x7F(%r14),%r11w
	mov	0x7F(%r14),%r12w
	mov	0x7F(%r14),%r13w
	mov	0x7F(%r14),%r14w
	mov	0x7F(%r14),%r15w
	nop
	mov	0x7F(%r15),%ax
	mov	0x7F(%r15),%cx
	mov	0x7F(%r15),%dx
	mov	0x7F(%r15),%bx
	mov	0x7F(%r15),%sp
	mov	0x7F(%r15),%bp
	mov	0x7F(%r15),%si
	mov	0x7F(%r15),%di
	mov	0x7F(%r15),%r8w
	mov	0x7F(%r15),%r9w
	mov	0x7F(%r15),%r10w
	mov	0x7F(%r15),%r11w
	mov	0x7F(%r15),%r12w
	mov	0x7F(%r15),%r13w
	mov	0x7F(%r15),%r14w
	mov	0x7F(%r15),%r15w
        nop
        nop
        // mem16[off32] += reg16
	mov	0x12345678(%rax),%ax
	mov	0x12345678(%rax),%cx
	mov	0x12345678(%rax),%dx
	mov	0x12345678(%rax),%bx
	mov	0x12345678(%rax),%sp
	mov	0x12345678(%rax),%bp
	mov	0x12345678(%rax),%si
	mov	0x12345678(%rax),%di
	mov	0x12345678(%rax),%r8w
	mov	0x12345678(%rax),%r9w
	mov	0x12345678(%rax),%r10w
	mov	0x12345678(%rax),%r11w
	mov	0x12345678(%rax),%r12w
	mov	0x12345678(%rax),%r13w
	mov	0x12345678(%rax),%r14w
	mov	0x12345678(%rax),%r15w
	nop
	mov	0x12345678(%rcx),%ax
	mov	0x12345678(%rcx),%cx
	mov	0x12345678(%rcx),%dx
	mov	0x12345678(%rcx),%bx
	mov	0x12345678(%rcx),%sp
	mov	0x12345678(%rcx),%bp
	mov	0x12345678(%rcx),%si
	mov	0x12345678(%rcx),%di
	mov	0x12345678(%rcx),%r8w
	mov	0x12345678(%rcx),%r9w
	mov	0x12345678(%rcx),%r10w
	mov	0x12345678(%rcx),%r11w
	mov	0x12345678(%rcx),%r12w
	mov	0x12345678(%rcx),%r13w
	mov	0x12345678(%rcx),%r14w
	mov	0x12345678(%rcx),%r15w
	nop
	mov	0x12345678(%rdx),%ax
	mov	0x12345678(%rdx),%cx
	mov	0x12345678(%rdx),%dx
	mov	0x12345678(%rdx),%bx
	mov	0x12345678(%rdx),%sp
	mov	0x12345678(%rdx),%bp
	mov	0x12345678(%rdx),%si
	mov	0x12345678(%rdx),%di
	mov	0x12345678(%rdx),%r8w
	mov	0x12345678(%rdx),%r9w
	mov	0x12345678(%rdx),%r10w
	mov	0x12345678(%rdx),%r11w
	mov	0x12345678(%rdx),%r12w
	mov	0x12345678(%rdx),%r13w
	mov	0x12345678(%rdx),%r14w
	mov	0x12345678(%rdx),%r15w
	nop
	mov	0x12345678(%rbx),%ax
	mov	0x12345678(%rbx),%cx
	mov	0x12345678(%rbx),%dx
	mov	0x12345678(%rbx),%bx
	mov	0x12345678(%rbx),%sp
	mov	0x12345678(%rbx),%bp
	mov	0x12345678(%rbx),%si
	mov	0x12345678(%rbx),%di
	mov	0x12345678(%rbx),%r8w
	mov	0x12345678(%rbx),%r9w
	mov	0x12345678(%rbx),%r10w
	mov	0x12345678(%rbx),%r11w
	mov	0x12345678(%rbx),%r12w
	mov	0x12345678(%rbx),%r13w
	mov	0x12345678(%rbx),%r14w
	mov	0x12345678(%rbx),%r15w
	nop
	mov	0x12345678(%rsp),%ax
	mov	0x12345678(%rsp),%cx
	mov	0x12345678(%rsp),%dx
	mov	0x12345678(%rsp),%bx
	mov	0x12345678(%rsp),%sp
	mov	0x12345678(%rsp),%bp
	mov	0x12345678(%rsp),%si
	mov	0x12345678(%rsp),%di
	mov	0x12345678(%rsp),%r8w
	mov	0x12345678(%rsp),%r9w
	mov	0x12345678(%rsp),%r10w
	mov	0x12345678(%rsp),%r11w
	mov	0x12345678(%rsp),%r12w
	mov	0x12345678(%rsp),%r13w
	mov	0x12345678(%rsp),%r14w
	mov	0x12345678(%rsp),%r15w
	nop
	mov	0x12345678(%rbp),%ax
	mov	0x12345678(%rbp),%cx
	mov	0x12345678(%rbp),%dx
	mov	0x12345678(%rbp),%bx
	mov	0x12345678(%rbp),%sp
	mov	0x12345678(%rbp),%bp
	mov	0x12345678(%rbp),%si
	mov	0x12345678(%rbp),%di
	mov	0x12345678(%rbp),%r8w
	mov	0x12345678(%rbp),%r9w
	mov	0x12345678(%rbp),%r10w
	mov	0x12345678(%rbp),%r11w
	mov	0x12345678(%rbp),%r12w
	mov	0x12345678(%rbp),%r13w
	mov	0x12345678(%rbp),%r14w
	mov	0x12345678(%rbp),%r15w
	nop
	mov	0x12345678(%rsi),%ax
	mov	0x12345678(%rsi),%cx
	mov	0x12345678(%rsi),%dx
	mov	0x12345678(%rsi),%bx
	mov	0x12345678(%rsi),%sp
	mov	0x12345678(%rsi),%bp
	mov	0x12345678(%rsi),%si
	mov	0x12345678(%rsi),%di
	mov	0x12345678(%rsi),%r8w
	mov	0x12345678(%rsi),%r9w
	mov	0x12345678(%rsi),%r10w
	mov	0x12345678(%rsi),%r11w
	mov	0x12345678(%rsi),%r12w
	mov	0x12345678(%rsi),%r13w
	mov	0x12345678(%rsi),%r14w
	mov	0x12345678(%rsi),%r15w
	nop
	mov	0x12345678(%rdi),%ax
	mov	0x12345678(%rdi),%cx
	mov	0x12345678(%rdi),%dx
	mov	0x12345678(%rdi),%bx
	mov	0x12345678(%rdi),%sp
	mov	0x12345678(%rdi),%bp
	mov	0x12345678(%rdi),%si
	mov	0x12345678(%rdi),%di
	mov	0x12345678(%rdi),%r8w
	mov	0x12345678(%rdi),%r9w
	mov	0x12345678(%rdi),%r10w
	mov	0x12345678(%rdi),%r11w
	mov	0x12345678(%rdi),%r12w
	mov	0x12345678(%rdi),%r13w
	mov	0x12345678(%rdi),%r14w
	mov	0x12345678(%rdi),%r15w
	nop
	mov	0x12345678(%r8), %ax
	mov	0x12345678(%r8), %cx
	mov	0x12345678(%r8), %dx
	mov	0x12345678(%r8), %bx
	mov	0x12345678(%r8), %sp
	mov	0x12345678(%r8), %bp
	mov	0x12345678(%r8), %si
	mov	0x12345678(%r8), %di
	mov	0x12345678(%r8), %r8w
	mov	0x12345678(%r8), %r9w
	mov	0x12345678(%r8), %r10w
	mov	0x12345678(%r8), %r11w
	mov	0x12345678(%r8), %r12w
	mov	0x12345678(%r8), %r13w
	mov	0x12345678(%r8), %r14w
	mov	0x12345678(%r8), %r15w
	nop
	mov	0x12345678(%r9), %ax
	mov	0x12345678(%r9), %cx
	mov	0x12345678(%r9), %dx
	mov	0x12345678(%r9), %bx
	mov	0x12345678(%r9), %sp
	mov	0x12345678(%r9), %bp
	mov	0x12345678(%r9), %si
	mov	0x12345678(%r9), %di
	mov	0x12345678(%r9), %r8w
	mov	0x12345678(%r9), %r9w
	mov	0x12345678(%r9), %r10w
	mov	0x12345678(%r9), %r11w
	mov	0x12345678(%r9), %r12w
	mov	0x12345678(%r9), %r13w
	mov	0x12345678(%r9), %r14w
	mov	0x12345678(%r9), %r15w
	nop
	mov	0x12345678(%r10),%ax
	mov	0x12345678(%r10),%cx
	mov	0x12345678(%r10),%dx
	mov	0x12345678(%r10),%bx
	mov	0x12345678(%r10),%sp
	mov	0x12345678(%r10),%bp
	mov	0x12345678(%r10),%si
	mov	0x12345678(%r10),%di
	mov	0x12345678(%r10),%r8w
	mov	0x12345678(%r10),%r9w
	mov	0x12345678(%r10),%r10w
	mov	0x12345678(%r10),%r11w
	mov	0x12345678(%r10),%r12w
	mov	0x12345678(%r10),%r13w
	mov	0x12345678(%r10),%r14w
	mov	0x12345678(%r10),%r15w
	nop
	mov	0x12345678(%r11),%ax
	mov	0x12345678(%r11),%cx
	mov	0x12345678(%r11),%dx
	mov	0x12345678(%r11),%bx
	mov	0x12345678(%r11),%sp
	mov	0x12345678(%r11),%bp
	mov	0x12345678(%r11),%si
	mov	0x12345678(%r11),%di
	mov	0x12345678(%r11),%r8w
	mov	0x12345678(%r11),%r9w
	mov	0x12345678(%r11),%r10w
	mov	0x12345678(%r11),%r11w
	mov	0x12345678(%r11),%r12w
	mov	0x12345678(%r11),%r13w
	mov	0x12345678(%r11),%r14w
	mov	0x12345678(%r11),%r15w
	nop
	mov	0x12345678(%r12),%ax
	mov	0x12345678(%r12),%cx
	mov	0x12345678(%r12),%dx
	mov	0x12345678(%r12),%bx
	mov	0x12345678(%r12),%sp
	mov	0x12345678(%r12),%bp
	mov	0x12345678(%r12),%si
	mov	0x12345678(%r12),%di
	mov	0x12345678(%r12),%r8w
	mov	0x12345678(%r12),%r9w
	mov	0x12345678(%r12),%r10w
	mov	0x12345678(%r12),%r11w
	mov	0x12345678(%r12),%r12w
	mov	0x12345678(%r12),%r13w
	mov	0x12345678(%r12),%r14w
	mov	0x12345678(%r12),%r15w
	nop
	mov	0x12345678(%r13),%ax
	mov	0x12345678(%r13),%cx
	mov	0x12345678(%r13),%dx
	mov	0x12345678(%r13),%bx
	mov	0x12345678(%r13),%sp
	mov	0x12345678(%r13),%bp
	mov	0x12345678(%r13),%si
	mov	0x12345678(%r13),%di
	mov	0x12345678(%r13),%r8w
	mov	0x12345678(%r13),%r9w
	mov	0x12345678(%r13),%r10w
	mov	0x12345678(%r13),%r11w
	mov	0x12345678(%r13),%r12w
	mov	0x12345678(%r13),%r13w
	mov	0x12345678(%r13),%r14w
	mov	0x12345678(%r13),%r15w
	nop
	mov	0x12345678(%r14),%ax
	mov	0x12345678(%r14),%cx
	mov	0x12345678(%r14),%dx
	mov	0x12345678(%r14),%bx
	mov	0x12345678(%r14),%sp
	mov	0x12345678(%r14),%bp
	mov	0x12345678(%r14),%si
	mov	0x12345678(%r14),%di
	mov	0x12345678(%r14),%r8w
	mov	0x12345678(%r14),%r9w
	mov	0x12345678(%r14),%r10w
	mov	0x12345678(%r14),%r11w
	mov	0x12345678(%r14),%r12w
	mov	0x12345678(%r14),%r13w
	mov	0x12345678(%r14),%r14w
	mov	0x12345678(%r14),%r15w
	nop
	mov	0x12345678(%r15),%ax
	mov	0x12345678(%r15),%cx
	mov	0x12345678(%r15),%dx
	mov	0x12345678(%r15),%bx
	mov	0x12345678(%r15),%sp
	mov	0x12345678(%r15),%bp
	mov	0x12345678(%r15),%si
	mov	0x12345678(%r15),%di
	mov	0x12345678(%r15),%r8w
	mov	0x12345678(%r15),%r9w
	mov	0x12345678(%r15),%r10w
	mov	0x12345678(%r15),%r11w
	mov	0x12345678(%r15),%r12w
	mov	0x12345678(%r15),%r13w
	mov	0x12345678(%r15),%r14w
	mov	0x12345678(%r15),%r15w
        ret
	.cfi_endproc

        
	.p2align 4,,15
	.globl	MovRegMem32
	.type	MovRegMem32, @function
MovRegMem32:
	.cfi_startproc
        // reg32 = mem32[0]
	mov	(%rax),%eax
	mov	(%rax),%ecx
	mov	(%rax),%edx
	mov	(%rax),%ebx
	mov	(%rax),%esp
	mov	(%rax),%ebp
	mov	(%rax),%esi
	mov	(%rax),%edi
	mov	(%rax),%r8d
	mov	(%rax),%r9d
	mov	(%rax),%r10d
	mov	(%rax),%r11d
	mov	(%rax),%r12d
	mov	(%rax),%r13d
	mov	(%rax),%r14d
	mov	(%rax),%r15d
	nop
	mov	(%rcx),%eax
	mov	(%rcx),%ecx
	mov	(%rcx),%edx
	mov	(%rcx),%ebx
	mov	(%rcx),%esp
	mov	(%rcx),%ebp
	mov	(%rcx),%esi
	mov	(%rcx),%edi
	mov	(%rcx),%r8d
	mov	(%rcx),%r9d
	mov	(%rcx),%r10d
	mov	(%rcx),%r11d
	mov	(%rcx),%r12d
	mov	(%rcx),%r13d
	mov	(%rcx),%r14d
	mov	(%rcx),%r15d
	nop
	mov	(%rdx),%eax
	mov	(%rdx),%ecx
	mov	(%rdx),%edx
	mov	(%rdx),%ebx
	mov	(%rdx),%esp
	mov	(%rdx),%ebp
	mov	(%rdx),%esi
	mov	(%rdx),%edi
	mov	(%rdx),%r8d
	mov	(%rdx),%r9d
	mov	(%rdx),%r10d
	mov	(%rdx),%r11d
	mov	(%rdx),%r12d
	mov	(%rdx),%r13d
	mov	(%rdx),%r14d
	mov	(%rdx),%r15d
	nop
	mov	(%rbx),%eax
	mov	(%rbx),%ecx
	mov	(%rbx),%edx
	mov	(%rbx),%ebx
	mov	(%rbx),%esp
	mov	(%rbx),%ebp
	mov	(%rbx),%esi
	mov	(%rbx),%edi
	mov	(%rbx),%r8d
	mov	(%rbx),%r9d
	mov	(%rbx),%r10d
	mov	(%rbx),%r11d
	mov	(%rbx),%r12d
	mov	(%rbx),%r13d
	mov	(%rbx),%r14d
	mov	(%rbx),%r15d
	nop
	mov	(%rsp),%eax
	mov	(%rsp),%ecx
	mov	(%rsp),%edx
	mov	(%rsp),%ebx
	mov	(%rsp),%esp
	mov	(%rsp),%ebp
	mov	(%rsp),%esi
	mov	(%rsp),%edi
	mov	(%rsp),%r8d
	mov	(%rsp),%r9d
	mov	(%rsp),%r10d
	mov	(%rsp),%r11d
	mov	(%rsp),%r12d
	mov	(%rsp),%r13d
	mov	(%rsp),%r14d
	mov	(%rsp),%r15d
	nop
	mov	(%rbp),%eax
	mov	(%rbp),%ecx
	mov	(%rbp),%edx
	mov	(%rbp),%ebx
	mov	(%rbp),%esp
	mov	(%rbp),%ebp
	mov	(%rbp),%esi
	mov	(%rbp),%edi
	mov	(%rbp),%r8d
	mov	(%rbp),%r9d
	mov	(%rbp),%r10d
	mov	(%rbp),%r11d
	mov	(%rbp),%r12d
	mov	(%rbp),%r13d
	mov	(%rbp),%r14d
	mov	(%rbp),%r15d
	nop
	mov	(%rsi),%eax
	mov	(%rsi),%ecx
	mov	(%rsi),%edx
	mov	(%rsi),%ebx
	mov	(%rsi),%esp
	mov	(%rsi),%ebp
	mov	(%rsi),%esi
	mov	(%rsi),%edi
	mov	(%rsi),%r8d
	mov	(%rsi),%r9d
	mov	(%rsi),%r10d
	mov	(%rsi),%r11d
	mov	(%rsi),%r12d
	mov	(%rsi),%r13d
	mov	(%rsi),%r14d
	mov	(%rsi),%r15d
	nop
	mov	(%rdi),%eax
	mov	(%rdi),%ecx
	mov	(%rdi),%edx
	mov	(%rdi),%ebx
	mov	(%rdi),%esp
	mov	(%rdi),%ebp
	mov	(%rdi),%esi
	mov	(%rdi),%edi
	mov	(%rdi),%r8d
	mov	(%rdi),%r9d
	mov	(%rdi),%r10d
	mov	(%rdi),%r11d
	mov	(%rdi),%r12d
	mov	(%rdi),%r13d
	mov	(%rdi),%r14d
	mov	(%rdi),%r15d
	nop
	mov	(%r8), %eax
	mov	(%r8), %ecx
	mov	(%r8), %edx
	mov	(%r8), %ebx
	mov	(%r8), %esp
	mov	(%r8), %ebp
	mov	(%r8), %esi
	mov	(%r8), %edi
	mov	(%r8), %r8d
	mov	(%r8), %r9d
	mov	(%r8), %r10d
	mov	(%r8), %r11d
	mov	(%r8), %r12d
	mov	(%r8), %r13d
	mov	(%r8), %r14d
	mov	(%r8), %r15d
	nop
	mov	(%r9), %eax
	mov	(%r9), %ecx
	mov	(%r9), %edx
	mov	(%r9), %ebx
	mov	(%r9), %esp
	mov	(%r9), %ebp
	mov	(%r9), %esi
	mov	(%r9), %edi
	mov	(%r9), %r8d
	mov	(%r9), %r9d
	mov	(%r9), %r10d
	mov	(%r9), %r11d
	mov	(%r9), %r12d
	mov	(%r9), %r13d
	mov	(%r9), %r14d
	mov	(%r9), %r15d
	nop
	mov	(%r10),%eax
	mov	(%r10),%ecx
	mov	(%r10),%edx
	mov	(%r10),%ebx
	mov	(%r10),%esp
	mov	(%r10),%ebp
	mov	(%r10),%esi
	mov	(%r10),%edi
	mov	(%r10),%r8d
	mov	(%r10),%r9d
	mov	(%r10),%r10d
	mov	(%r10),%r11d
	mov	(%r10),%r12d
	mov	(%r10),%r13d
	mov	(%r10),%r14d
	mov	(%r10),%r15d
	nop
	mov	(%r11),%eax
	mov	(%r11),%ecx
	mov	(%r11),%edx
	mov	(%r11),%ebx
	mov	(%r11),%esp
	mov	(%r11),%ebp
	mov	(%r11),%esi
	mov	(%r11),%edi
	mov	(%r11),%r8d
	mov	(%r11),%r9d
	mov	(%r11),%r10d
	mov	(%r11),%r11d
	mov	(%r11),%r12d
	mov	(%r11),%r13d
	mov	(%r11),%r14d
	mov	(%r11),%r15d
	nop
	mov	(%r12),%eax
	mov	(%r12),%ecx
	mov	(%r12),%edx
	mov	(%r12),%ebx
	mov	(%r12),%esp
	mov	(%r12),%ebp
	mov	(%r12),%esi
	mov	(%r12),%edi
	mov	(%r12),%r8d
	mov	(%r12),%r9d
	mov	(%r12),%r10d
	mov	(%r12),%r11d
	mov	(%r12),%r12d
	mov	(%r12),%r13d
	mov	(%r12),%r14d
	mov	(%r12),%r15d
	nop
	mov	(%r13),%eax
	mov	(%r13),%ecx
	mov	(%r13),%edx
	mov	(%r13),%ebx
	mov	(%r13),%esp
	mov	(%r13),%ebp
	mov	(%r13),%esi
	mov	(%r13),%edi
	mov	(%r13),%r8d
	mov	(%r13),%r9d
	mov	(%r13),%r10d
	mov	(%r13),%r11d
	mov	(%r13),%r12d
	mov	(%r13),%r13d
	mov	(%r13),%r14d
	mov	(%r13),%r15d
	nop
	mov	(%r14),%eax
	mov	(%r14),%ecx
	mov	(%r14),%edx
	mov	(%r14),%ebx
	mov	(%r14),%esp
	mov	(%r14),%ebp
	mov	(%r14),%esi
	mov	(%r14),%edi
	mov	(%r14),%r8d
	mov	(%r14),%r9d
	mov	(%r14),%r10d
	mov	(%r14),%r11d
	mov	(%r14),%r12d
	mov	(%r14),%r13d
	mov	(%r14),%r14d
	mov	(%r14),%r15d
	nop
	mov	(%r15),%eax
	mov	(%r15),%ecx
	mov	(%r15),%edx
	mov	(%r15),%ebx
	mov	(%r15),%esp
	mov	(%r15),%ebp
	mov	(%r15),%esi
	mov	(%r15),%edi
	mov	(%r15),%r8d
	mov	(%r15),%r9d
	mov	(%r15),%r10d
	mov	(%r15),%r11d
	mov	(%r15),%r12d
	mov	(%r15),%r13d
	mov	(%r15),%r14d
	mov	(%r15),%r15d
        nop
        nop
        // mem32[off8] += reg32
	mov	0x7F(%rax),%eax
	mov	0x7F(%rax),%ecx
	mov	0x7F(%rax),%edx
	mov	0x7F(%rax),%ebx
	mov	0x7F(%rax),%esp
	mov	0x7F(%rax),%ebp
	mov	0x7F(%rax),%esi
	mov	0x7F(%rax),%edi
	mov	0x7F(%rax),%r8d
	mov	0x7F(%rax),%r9d
	mov	0x7F(%rax),%r10d
	mov	0x7F(%rax),%r11d
	mov	0x7F(%rax),%r12d
	mov	0x7F(%rax),%r13d
	mov	0x7F(%rax),%r14d
	mov	0x7F(%rax),%r15d
	nop
	mov	0x7F(%rcx),%eax
	mov	0x7F(%rcx),%ecx
	mov	0x7F(%rcx),%edx
	mov	0x7F(%rcx),%ebx
	mov	0x7F(%rcx),%esp
	mov	0x7F(%rcx),%ebp
	mov	0x7F(%rcx),%esi
	mov	0x7F(%rcx),%edi
	mov	0x7F(%rcx),%r8d
	mov	0x7F(%rcx),%r9d
	mov	0x7F(%rcx),%r10d
	mov	0x7F(%rcx),%r11d
	mov	0x7F(%rcx),%r12d
	mov	0x7F(%rcx),%r13d
	mov	0x7F(%rcx),%r14d
	mov	0x7F(%rcx),%r15d
	nop
	mov	0x7F(%rdx),%eax
	mov	0x7F(%rdx),%ecx
	mov	0x7F(%rdx),%edx
	mov	0x7F(%rdx),%ebx
	mov	0x7F(%rdx),%esp
	mov	0x7F(%rdx),%ebp
	mov	0x7F(%rdx),%esi
	mov	0x7F(%rdx),%edi
	mov	0x7F(%rdx),%r8d
	mov	0x7F(%rdx),%r9d
	mov	0x7F(%rdx),%r10d
	mov	0x7F(%rdx),%r11d
	mov	0x7F(%rdx),%r12d
	mov	0x7F(%rdx),%r13d
	mov	0x7F(%rdx),%r14d
	mov	0x7F(%rdx),%r15d
	nop
	mov	0x7F(%rbx),%eax
	mov	0x7F(%rbx),%ecx
	mov	0x7F(%rbx),%edx
	mov	0x7F(%rbx),%ebx
	mov	0x7F(%rbx),%esp
	mov	0x7F(%rbx),%ebp
	mov	0x7F(%rbx),%esi
	mov	0x7F(%rbx),%edi
	mov	0x7F(%rbx),%r8d
	mov	0x7F(%rbx),%r9d
	mov	0x7F(%rbx),%r10d
	mov	0x7F(%rbx),%r11d
	mov	0x7F(%rbx),%r12d
	mov	0x7F(%rbx),%r13d
	mov	0x7F(%rbx),%r14d
	mov	0x7F(%rbx),%r15d
	nop
	mov	0x7F(%rsp),%eax
	mov	0x7F(%rsp),%ecx
	mov	0x7F(%rsp),%edx
	mov	0x7F(%rsp),%ebx
	mov	0x7F(%rsp),%esp
	mov	0x7F(%rsp),%ebp
	mov	0x7F(%rsp),%esi
	mov	0x7F(%rsp),%edi
	mov	0x7F(%rsp),%r8d
	mov	0x7F(%rsp),%r9d
	mov	0x7F(%rsp),%r10d
	mov	0x7F(%rsp),%r11d
	mov	0x7F(%rsp),%r12d
	mov	0x7F(%rsp),%r13d
	mov	0x7F(%rsp),%r14d
	mov	0x7F(%rsp),%r15d
	nop
	mov	0x7F(%rbp),%eax
	mov	0x7F(%rbp),%ecx
	mov	0x7F(%rbp),%edx
	mov	0x7F(%rbp),%ebx
	mov	0x7F(%rbp),%esp
	mov	0x7F(%rbp),%ebp
	mov	0x7F(%rbp),%esi
	mov	0x7F(%rbp),%edi
	mov	0x7F(%rbp),%r8d
	mov	0x7F(%rbp),%r9d
	mov	0x7F(%rbp),%r10d
	mov	0x7F(%rbp),%r11d
	mov	0x7F(%rbp),%r12d
	mov	0x7F(%rbp),%r13d
	mov	0x7F(%rbp),%r14d
	mov	0x7F(%rbp),%r15d
	nop
	mov	0x7F(%rsi),%eax
	mov	0x7F(%rsi),%ecx
	mov	0x7F(%rsi),%edx
	mov	0x7F(%rsi),%ebx
	mov	0x7F(%rsi),%esp
	mov	0x7F(%rsi),%ebp
	mov	0x7F(%rsi),%esi
	mov	0x7F(%rsi),%edi
	mov	0x7F(%rsi),%r8d
	mov	0x7F(%rsi),%r9d
	mov	0x7F(%rsi),%r10d
	mov	0x7F(%rsi),%r11d
	mov	0x7F(%rsi),%r12d
	mov	0x7F(%rsi),%r13d
	mov	0x7F(%rsi),%r14d
	mov	0x7F(%rsi),%r15d
	nop
	mov	0x7F(%rdi),%eax
	mov	0x7F(%rdi),%ecx
	mov	0x7F(%rdi),%edx
	mov	0x7F(%rdi),%ebx
	mov	0x7F(%rdi),%esp
	mov	0x7F(%rdi),%ebp
	mov	0x7F(%rdi),%esi
	mov	0x7F(%rdi),%edi
	mov	0x7F(%rdi),%r8d
	mov	0x7F(%rdi),%r9d
	mov	0x7F(%rdi),%r10d
	mov	0x7F(%rdi),%r11d
	mov	0x7F(%rdi),%r12d
	mov	0x7F(%rdi),%r13d
	mov	0x7F(%rdi),%r14d
	mov	0x7F(%rdi),%r15d
	nop
	mov	0x7F(%r8), %eax
	mov	0x7F(%r8), %ecx
	mov	0x7F(%r8), %edx
	mov	0x7F(%r8), %ebx
	mov	0x7F(%r8), %esp
	mov	0x7F(%r8), %ebp
	mov	0x7F(%r8), %esi
	mov	0x7F(%r8), %edi
	mov	0x7F(%r8), %r8d
	mov	0x7F(%r8), %r9d
	mov	0x7F(%r8), %r10d
	mov	0x7F(%r8), %r11d
	mov	0x7F(%r8), %r12d
	mov	0x7F(%r8), %r13d
	mov	0x7F(%r8), %r14d
	mov	0x7F(%r8), %r15d
	nop
	mov	0x7F(%r9), %eax
	mov	0x7F(%r9), %ecx
	mov	0x7F(%r9), %edx
	mov	0x7F(%r9), %ebx
	mov	0x7F(%r9), %esp
	mov	0x7F(%r9), %ebp
	mov	0x7F(%r9), %esi
	mov	0x7F(%r9), %edi
	mov	0x7F(%r9), %r8d
	mov	0x7F(%r9), %r9d
	mov	0x7F(%r9), %r10d
	mov	0x7F(%r9), %r11d
	mov	0x7F(%r9), %r12d
	mov	0x7F(%r9), %r13d
	mov	0x7F(%r9), %r14d
	mov	0x7F(%r9), %r15d
	nop
	mov	0x7F(%r10),%eax
	mov	0x7F(%r10),%ecx
	mov	0x7F(%r10),%edx
	mov	0x7F(%r10),%ebx
	mov	0x7F(%r10),%esp
	mov	0x7F(%r10),%ebp
	mov	0x7F(%r10),%esi
	mov	0x7F(%r10),%edi
	mov	0x7F(%r10),%r8d
	mov	0x7F(%r10),%r9d
	mov	0x7F(%r10),%r10d
	mov	0x7F(%r10),%r11d
	mov	0x7F(%r10),%r12d
	mov	0x7F(%r10),%r13d
	mov	0x7F(%r10),%r14d
	mov	0x7F(%r10),%r15d
	nop
	mov	0x7F(%r11),%eax
	mov	0x7F(%r11),%ecx
	mov	0x7F(%r11),%edx
	mov	0x7F(%r11),%ebx
	mov	0x7F(%r11),%esp
	mov	0x7F(%r11),%ebp
	mov	0x7F(%r11),%esi
	mov	0x7F(%r11),%edi
	mov	0x7F(%r11),%r8d
	mov	0x7F(%r11),%r9d
	mov	0x7F(%r11),%r10d
	mov	0x7F(%r11),%r11d
	mov	0x7F(%r11),%r12d
	mov	0x7F(%r11),%r13d
	mov	0x7F(%r11),%r14d
	mov	0x7F(%r11),%r15d
	nop
	mov	0x7F(%r12),%eax
	mov	0x7F(%r12),%ecx
	mov	0x7F(%r12),%edx
	mov	0x7F(%r12),%ebx
	mov	0x7F(%r12),%esp
	mov	0x7F(%r12),%ebp
	mov	0x7F(%r12),%esi
	mov	0x7F(%r12),%edi
	mov	0x7F(%r12),%r8d
	mov	0x7F(%r12),%r9d
	mov	0x7F(%r12),%r10d
	mov	0x7F(%r12),%r11d
	mov	0x7F(%r12),%r12d
	mov	0x7F(%r12),%r13d
	mov	0x7F(%r12),%r14d
	mov	0x7F(%r12),%r15d
	nop
	mov	0x7F(%r13),%eax
	mov	0x7F(%r13),%ecx
	mov	0x7F(%r13),%edx
	mov	0x7F(%r13),%ebx
	mov	0x7F(%r13),%esp
	mov	0x7F(%r13),%ebp
	mov	0x7F(%r13),%esi
	mov	0x7F(%r13),%edi
	mov	0x7F(%r13),%r8d
	mov	0x7F(%r13),%r9d
	mov	0x7F(%r13),%r10d
	mov	0x7F(%r13),%r11d
	mov	0x7F(%r13),%r12d
	mov	0x7F(%r13),%r13d
	mov	0x7F(%r13),%r14d
	mov	0x7F(%r13),%r15d
	nop
	mov	0x7F(%r14),%eax
	mov	0x7F(%r14),%ecx
	mov	0x7F(%r14),%edx
	mov	0x7F(%r14),%ebx
	mov	0x7F(%r14),%esp
	mov	0x7F(%r14),%ebp
	mov	0x7F(%r14),%esi
	mov	0x7F(%r14),%edi
	mov	0x7F(%r14),%r8d
	mov	0x7F(%r14),%r9d
	mov	0x7F(%r14),%r10d
	mov	0x7F(%r14),%r11d
	mov	0x7F(%r14),%r12d
	mov	0x7F(%r14),%r13d
	mov	0x7F(%r14),%r14d
	mov	0x7F(%r14),%r15d
	nop
	mov	0x7F(%r15),%eax
	mov	0x7F(%r15),%ecx
	mov	0x7F(%r15),%edx
	mov	0x7F(%r15),%ebx
	mov	0x7F(%r15),%esp
	mov	0x7F(%r15),%ebp
	mov	0x7F(%r15),%esi
	mov	0x7F(%r15),%edi
	mov	0x7F(%r15),%r8d
	mov	0x7F(%r15),%r9d
	mov	0x7F(%r15),%r10d
	mov	0x7F(%r15),%r11d
	mov	0x7F(%r15),%r12d
	mov	0x7F(%r15),%r13d
	mov	0x7F(%r15),%r14d
	mov	0x7F(%r15),%r15d
        nop
        nop
        // mem32[off32] += reg32
	mov	0x12345678(%rax),%eax
	mov	0x12345678(%rax),%ecx
	mov	0x12345678(%rax),%edx
	mov	0x12345678(%rax),%ebx
	mov	0x12345678(%rax),%esp
	mov	0x12345678(%rax),%ebp
	mov	0x12345678(%rax),%esi
	mov	0x12345678(%rax),%edi
	mov	0x12345678(%rax),%r8d
	mov	0x12345678(%rax),%r9d
	mov	0x12345678(%rax),%r10d
	mov	0x12345678(%rax),%r11d
	mov	0x12345678(%rax),%r12d
	mov	0x12345678(%rax),%r13d
	mov	0x12345678(%rax),%r14d
	mov	0x12345678(%rax),%r15d
	nop
	mov	0x12345678(%rcx),%eax
	mov	0x12345678(%rcx),%ecx
	mov	0x12345678(%rcx),%edx
	mov	0x12345678(%rcx),%ebx
	mov	0x12345678(%rcx),%esp
	mov	0x12345678(%rcx),%ebp
	mov	0x12345678(%rcx),%esi
	mov	0x12345678(%rcx),%edi
	mov	0x12345678(%rcx),%r8d
	mov	0x12345678(%rcx),%r9d
	mov	0x12345678(%rcx),%r10d
	mov	0x12345678(%rcx),%r11d
	mov	0x12345678(%rcx),%r12d
	mov	0x12345678(%rcx),%r13d
	mov	0x12345678(%rcx),%r14d
	mov	0x12345678(%rcx),%r15d
	nop
	mov	0x12345678(%rdx),%eax
	mov	0x12345678(%rdx),%ecx
	mov	0x12345678(%rdx),%edx
	mov	0x12345678(%rdx),%ebx
	mov	0x12345678(%rdx),%esp
	mov	0x12345678(%rdx),%ebp
	mov	0x12345678(%rdx),%esi
	mov	0x12345678(%rdx),%edi
	mov	0x12345678(%rdx),%r8d
	mov	0x12345678(%rdx),%r9d
	mov	0x12345678(%rdx),%r10d
	mov	0x12345678(%rdx),%r11d
	mov	0x12345678(%rdx),%r12d
	mov	0x12345678(%rdx),%r13d
	mov	0x12345678(%rdx),%r14d
	mov	0x12345678(%rdx),%r15d
	nop
	mov	0x12345678(%rbx),%eax
	mov	0x12345678(%rbx),%ecx
	mov	0x12345678(%rbx),%edx
	mov	0x12345678(%rbx),%ebx
	mov	0x12345678(%rbx),%esp
	mov	0x12345678(%rbx),%ebp
	mov	0x12345678(%rbx),%esi
	mov	0x12345678(%rbx),%edi
	mov	0x12345678(%rbx),%r8d
	mov	0x12345678(%rbx),%r9d
	mov	0x12345678(%rbx),%r10d
	mov	0x12345678(%rbx),%r11d
	mov	0x12345678(%rbx),%r12d
	mov	0x12345678(%rbx),%r13d
	mov	0x12345678(%rbx),%r14d
	mov	0x12345678(%rbx),%r15d
	nop
	mov	0x12345678(%rsp),%eax
	mov	0x12345678(%rsp),%ecx
	mov	0x12345678(%rsp),%edx
	mov	0x12345678(%rsp),%ebx
	mov	0x12345678(%rsp),%esp
	mov	0x12345678(%rsp),%ebp
	mov	0x12345678(%rsp),%esi
	mov	0x12345678(%rsp),%edi
	mov	0x12345678(%rsp),%r8d
	mov	0x12345678(%rsp),%r9d
	mov	0x12345678(%rsp),%r10d
	mov	0x12345678(%rsp),%r11d
	mov	0x12345678(%rsp),%r12d
	mov	0x12345678(%rsp),%r13d
	mov	0x12345678(%rsp),%r14d
	mov	0x12345678(%rsp),%r15d
	nop
	mov	0x12345678(%rbp),%eax
	mov	0x12345678(%rbp),%ecx
	mov	0x12345678(%rbp),%edx
	mov	0x12345678(%rbp),%ebx
	mov	0x12345678(%rbp),%esp
	mov	0x12345678(%rbp),%ebp
	mov	0x12345678(%rbp),%esi
	mov	0x12345678(%rbp),%edi
	mov	0x12345678(%rbp),%r8d
	mov	0x12345678(%rbp),%r9d
	mov	0x12345678(%rbp),%r10d
	mov	0x12345678(%rbp),%r11d
	mov	0x12345678(%rbp),%r12d
	mov	0x12345678(%rbp),%r13d
	mov	0x12345678(%rbp),%r14d
	mov	0x12345678(%rbp),%r15d
	nop
	mov	0x12345678(%rsi),%eax
	mov	0x12345678(%rsi),%ecx
	mov	0x12345678(%rsi),%edx
	mov	0x12345678(%rsi),%ebx
	mov	0x12345678(%rsi),%esp
	mov	0x12345678(%rsi),%ebp
	mov	0x12345678(%rsi),%esi
	mov	0x12345678(%rsi),%edi
	mov	0x12345678(%rsi),%r8d
	mov	0x12345678(%rsi),%r9d
	mov	0x12345678(%rsi),%r10d
	mov	0x12345678(%rsi),%r11d
	mov	0x12345678(%rsi),%r12d
	mov	0x12345678(%rsi),%r13d
	mov	0x12345678(%rsi),%r14d
	mov	0x12345678(%rsi),%r15d
	nop
	mov	0x12345678(%rdi),%eax
	mov	0x12345678(%rdi),%ecx
	mov	0x12345678(%rdi),%edx
	mov	0x12345678(%rdi),%ebx
	mov	0x12345678(%rdi),%esp
	mov	0x12345678(%rdi),%ebp
	mov	0x12345678(%rdi),%esi
	mov	0x12345678(%rdi),%edi
	mov	0x12345678(%rdi),%r8d
	mov	0x12345678(%rdi),%r9d
	mov	0x12345678(%rdi),%r10d
	mov	0x12345678(%rdi),%r11d
	mov	0x12345678(%rdi),%r12d
	mov	0x12345678(%rdi),%r13d
	mov	0x12345678(%rdi),%r14d
	mov	0x12345678(%rdi),%r15d
	nop
	mov	0x12345678(%r8), %eax
	mov	0x12345678(%r8), %ecx
	mov	0x12345678(%r8), %edx
	mov	0x12345678(%r8), %ebx
	mov	0x12345678(%r8), %esp
	mov	0x12345678(%r8), %ebp
	mov	0x12345678(%r8), %esi
	mov	0x12345678(%r8), %edi
	mov	0x12345678(%r8), %r8d
	mov	0x12345678(%r8), %r9d
	mov	0x12345678(%r8), %r10d
	mov	0x12345678(%r8), %r11d
	mov	0x12345678(%r8), %r12d
	mov	0x12345678(%r8), %r13d
	mov	0x12345678(%r8), %r14d
	mov	0x12345678(%r8), %r15d
	nop
	mov	0x12345678(%r9), %eax
	mov	0x12345678(%r9), %ecx
	mov	0x12345678(%r9), %edx
	mov	0x12345678(%r9), %ebx
	mov	0x12345678(%r9), %esp
	mov	0x12345678(%r9), %ebp
	mov	0x12345678(%r9), %esi
	mov	0x12345678(%r9), %edi
	mov	0x12345678(%r9), %r8d
	mov	0x12345678(%r9), %r9d
	mov	0x12345678(%r9), %r10d
	mov	0x12345678(%r9), %r11d
	mov	0x12345678(%r9), %r12d
	mov	0x12345678(%r9), %r13d
	mov	0x12345678(%r9), %r14d
	mov	0x12345678(%r9), %r15d
	nop
	mov	0x12345678(%r10),%eax
	mov	0x12345678(%r10),%ecx
	mov	0x12345678(%r10),%edx
	mov	0x12345678(%r10),%ebx
	mov	0x12345678(%r10),%esp
	mov	0x12345678(%r10),%ebp
	mov	0x12345678(%r10),%esi
	mov	0x12345678(%r10),%edi
	mov	0x12345678(%r10),%r8d
	mov	0x12345678(%r10),%r9d
	mov	0x12345678(%r10),%r10d
	mov	0x12345678(%r10),%r11d
	mov	0x12345678(%r10),%r12d
	mov	0x12345678(%r10),%r13d
	mov	0x12345678(%r10),%r14d
	mov	0x12345678(%r10),%r15d
	nop
	mov	0x12345678(%r11),%eax
	mov	0x12345678(%r11),%ecx
	mov	0x12345678(%r11),%edx
	mov	0x12345678(%r11),%ebx
	mov	0x12345678(%r11),%esp
	mov	0x12345678(%r11),%ebp
	mov	0x12345678(%r11),%esi
	mov	0x12345678(%r11),%edi
	mov	0x12345678(%r11),%r8d
	mov	0x12345678(%r11),%r9d
	mov	0x12345678(%r11),%r10d
	mov	0x12345678(%r11),%r11d
	mov	0x12345678(%r11),%r12d
	mov	0x12345678(%r11),%r13d
	mov	0x12345678(%r11),%r14d
	mov	0x12345678(%r11),%r15d
	nop
	mov	0x12345678(%r12),%eax
	mov	0x12345678(%r12),%ecx
	mov	0x12345678(%r12),%edx
	mov	0x12345678(%r12),%ebx
	mov	0x12345678(%r12),%esp
	mov	0x12345678(%r12),%ebp
	mov	0x12345678(%r12),%esi
	mov	0x12345678(%r12),%edi
	mov	0x12345678(%r12),%r8d
	mov	0x12345678(%r12),%r9d
	mov	0x12345678(%r12),%r10d
	mov	0x12345678(%r12),%r11d
	mov	0x12345678(%r12),%r12d
	mov	0x12345678(%r12),%r13d
	mov	0x12345678(%r12),%r14d
	mov	0x12345678(%r12),%r15d
	nop
	mov	0x12345678(%r13),%eax
	mov	0x12345678(%r13),%ecx
	mov	0x12345678(%r13),%edx
	mov	0x12345678(%r13),%ebx
	mov	0x12345678(%r13),%esp
	mov	0x12345678(%r13),%ebp
	mov	0x12345678(%r13),%esi
	mov	0x12345678(%r13),%edi
	mov	0x12345678(%r13),%r8d
	mov	0x12345678(%r13),%r9d
	mov	0x12345678(%r13),%r10d
	mov	0x12345678(%r13),%r11d
	mov	0x12345678(%r13),%r12d
	mov	0x12345678(%r13),%r13d
	mov	0x12345678(%r13),%r14d
	mov	0x12345678(%r13),%r15d
	nop
	mov	0x12345678(%r14),%eax
	mov	0x12345678(%r14),%ecx
	mov	0x12345678(%r14),%edx
	mov	0x12345678(%r14),%ebx
	mov	0x12345678(%r14),%esp
	mov	0x12345678(%r14),%ebp
	mov	0x12345678(%r14),%esi
	mov	0x12345678(%r14),%edi
	mov	0x12345678(%r14),%r8d
	mov	0x12345678(%r14),%r9d
	mov	0x12345678(%r14),%r10d
	mov	0x12345678(%r14),%r11d
	mov	0x12345678(%r14),%r12d
	mov	0x12345678(%r14),%r13d
	mov	0x12345678(%r14),%r14d
	mov	0x12345678(%r14),%r15d
	nop
	mov	0x12345678(%r15),%eax
	mov	0x12345678(%r15),%ecx
	mov	0x12345678(%r15),%edx
	mov	0x12345678(%r15),%ebx
	mov	0x12345678(%r15),%esp
	mov	0x12345678(%r15),%ebp
	mov	0x12345678(%r15),%esi
	mov	0x12345678(%r15),%edi
	mov	0x12345678(%r15),%r8d
	mov	0x12345678(%r15),%r9d
	mov	0x12345678(%r15),%r10d
	mov	0x12345678(%r15),%r11d
	mov	0x12345678(%r15),%r12d
	mov	0x12345678(%r15),%r13d
	mov	0x12345678(%r15),%r14d
	mov	0x12345678(%r15),%r15d
        ret
	.cfi_endproc
