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
	.globl	MovMemReg
	.type	MovMemReg, @function
MovMemReg:
	.cfi_startproc
        // mem8 += reg8
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
        // mem16 += reg16
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
        // mem32 += reg32
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
        // mem64 += reg64
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


