	.file	"sub.s"
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
	.globl	Sub
	.type	Sub, @function
Sub:
	.cfi_startproc

        // reg8 += reg8
	sub	%al,%al
	sub	%al,%cl
	sub	%al,%dl
	sub	%al,%bl
	sub	%al,%spl
	sub	%al,%bpl
	sub	%al,%sil
	sub	%al,%dil
	sub	%al,%r8b
	sub	%al,%r9b
	sub	%al,%r10b
	sub	%al,%r11b
	sub	%al,%r12b
	sub	%al,%r13b
	sub	%al,%r14b
	sub	%al,%r15b
	nop
	sub	%cl,%al
	sub	%cl,%cl
	sub	%cl,%dl
	sub	%cl,%bl
	sub	%cl,%spl
	sub	%cl,%bpl
	sub	%cl,%sil
	sub	%cl,%dil
	sub	%cl,%r8b
	sub	%cl,%r9b
	sub	%cl,%r10b
	sub	%cl,%r11b
	sub	%cl,%r12b
	sub	%cl,%r13b
	sub	%cl,%r14b
	sub	%cl,%r15b
	nop
	sub	%dl,%al
	sub	%dl,%cl
	sub	%dl,%dl
	sub	%dl,%bl
	sub	%dl,%spl
	sub	%dl,%bpl
	sub	%dl,%sil
	sub	%dl,%dil
	sub	%dl,%r8b
	sub	%dl,%r9b
	sub	%dl,%r10b
	sub	%dl,%r11b
	sub	%dl,%r12b
	sub	%dl,%r13b
	sub	%dl,%r14b
	sub	%dl,%r15b
	nop
	sub	%bl,%al
	sub	%bl,%cl
	sub	%bl,%dl
	sub	%bl,%bl
	sub	%bl,%spl
	sub	%bl,%bpl
	sub	%bl,%sil
	sub	%bl,%dil
	sub	%bl,%r8b
	sub	%bl,%r9b
	sub	%bl,%r10b
	sub	%bl,%r11b
	sub	%bl,%r12b
	sub	%bl,%r13b
	sub	%bl,%r14b
	sub	%bl,%r15b
	nop
	sub	%spl,%al
	sub	%spl,%cl
	sub	%spl,%dl
	sub	%spl,%bl
	sub	%spl,%spl
	sub	%spl,%bpl
	sub	%spl,%sil
	sub	%spl,%dil
	sub	%spl,%r8b
	sub	%spl,%r9b
	sub	%spl,%r10b
	sub	%spl,%r11b
	sub	%spl,%r12b
	sub	%spl,%r13b
	sub	%spl,%r14b
	sub	%spl,%r15b
	nop
	sub	%bpl,%al
	sub	%bpl,%cl
	sub	%bpl,%dl
	sub	%bpl,%bl
	sub	%bpl,%spl
	sub	%bpl,%bpl
	sub	%bpl,%sil
	sub	%bpl,%dil
	sub	%bpl,%r8b
	sub	%bpl,%r9b
	sub	%bpl,%r10b
	sub	%bpl,%r11b
	sub	%bpl,%r12b
	sub	%bpl,%r13b
	sub	%bpl,%r14b
	sub	%bpl,%r15b
	nop
	sub	%sil,%al
	sub	%sil,%cl
	sub	%sil,%dl
	sub	%sil,%bl
	sub	%sil,%spl
	sub	%sil,%bpl
	sub	%sil,%sil
	sub	%sil,%dil
	sub	%sil,%r8b
	sub	%sil,%r9b
	sub	%sil,%r10b
	sub	%sil,%r11b
	sub	%sil,%r12b
	sub	%sil,%r13b
	sub	%sil,%r14b
	sub	%sil,%r15b
	nop
	sub	%dil,%al
	sub	%dil,%cl
	sub	%dil,%dl
	sub	%dil,%bl
	sub	%dil,%spl
	sub	%dil,%bpl
	sub	%dil,%sil
	sub	%dil,%dil
	sub	%dil,%r8b
	sub	%dil,%r9b
	sub	%dil,%r10b
	sub	%dil,%r11b
	sub	%dil,%r12b
	sub	%dil,%r13b
	sub	%dil,%r14b
	sub	%dil,%r15b
	nop
	sub	%r8b, %al
	sub	%r8b, %cl
	sub	%r8b, %dl
	sub	%r8b, %bl
	sub	%r8b, %spl
	sub	%r8b, %bpl
	sub	%r8b, %sil
	sub	%r8b, %dil
	sub	%r8b, %r8b
	sub	%r8b, %r9b
	sub	%r8b, %r10b
	sub	%r8b, %r11b
	sub	%r8b, %r12b
	sub	%r8b, %r13b
	sub	%r8b, %r14b
	sub	%r8b, %r15b
	nop
	sub	%r9b, %al
	sub	%r9b, %cl
	sub	%r9b, %dl
	sub	%r9b, %bl
	sub	%r9b, %spl
	sub	%r9b, %bpl
	sub	%r9b, %sil
	sub	%r9b, %dil
	sub	%r9b, %r8b
	sub	%r9b, %r9b
	sub	%r9b, %r10b
	sub	%r9b, %r11b
	sub	%r9b, %r12b
	sub	%r9b, %r13b
	sub	%r9b, %r14b
	sub	%r9b, %r15b
	nop
	sub	%r10b,%al
	sub	%r10b,%cl
	sub	%r10b,%dl
	sub	%r10b,%bl
	sub	%r10b,%spl
	sub	%r10b,%bpl
	sub	%r10b,%sil
	sub	%r10b,%dil
	sub	%r10b,%r8b
	sub	%r10b,%r9b
	sub	%r10b,%r10b
	sub	%r10b,%r11b
	sub	%r10b,%r12b
	sub	%r10b,%r13b
	sub	%r10b,%r14b
	sub	%r10b,%r15b
	nop
	sub	%r11b,%al
	sub	%r11b,%cl
	sub	%r11b,%dl
	sub	%r11b,%bl
	sub	%r11b,%spl
	sub	%r11b,%bpl
	sub	%r11b,%sil
	sub	%r11b,%dil
	sub	%r11b,%r8b
	sub	%r11b,%r9b
	sub	%r11b,%r10b
	sub	%r11b,%r11b
	sub	%r11b,%r12b
	sub	%r11b,%r13b
	sub	%r11b,%r14b
	sub	%r11b,%r15b
	nop
	sub	%r12b,%al
	sub	%r12b,%cl
	sub	%r12b,%dl
	sub	%r12b,%bl
	sub	%r12b,%spl
	sub	%r12b,%bpl
	sub	%r12b,%sil
	sub	%r12b,%dil
	sub	%r12b,%r8b
	sub	%r12b,%r9b
	sub	%r12b,%r10b
	sub	%r12b,%r11b
	sub	%r12b,%r12b
	sub	%r12b,%r13b
	sub	%r12b,%r14b
	sub	%r12b,%r15b
	nop
	sub	%r13b,%al
	sub	%r13b,%cl
	sub	%r13b,%dl
	sub	%r13b,%bl
	sub	%r13b,%spl
	sub	%r13b,%bpl
	sub	%r13b,%sil
	sub	%r13b,%dil
	sub	%r13b,%r8b
	sub	%r13b,%r9b
	sub	%r13b,%r10b
	sub	%r13b,%r11b
	sub	%r13b,%r12b
	sub	%r13b,%r13b
	sub	%r13b,%r14b
	sub	%r13b,%r15b
	nop
	sub	%r14b,%al
	sub	%r14b,%cl
	sub	%r14b,%dl
	sub	%r14b,%bl
	sub	%r14b,%spl
	sub	%r14b,%bpl
	sub	%r14b,%sil
	sub	%r14b,%dil
	sub	%r14b,%r8b
	sub	%r14b,%r9b
	sub	%r14b,%r10b
	sub	%r14b,%r11b
	sub	%r14b,%r12b
	sub	%r14b,%r13b
	sub	%r14b,%r14b
	sub	%r14b,%r15b
	nop
	sub	%r15b,%al
	sub	%r15b,%cl
	sub	%r15b,%dl
	sub	%r15b,%bl
	sub	%r15b,%spl
	sub	%r15b,%bpl
	sub	%r15b,%sil
	sub	%r15b,%dil
	sub	%r15b,%r8b
	sub	%r15b,%r9b
	sub	%r15b,%r10b
	sub	%r15b,%r11b
	sub	%r15b,%r12b
	sub	%r15b,%r13b
	sub	%r15b,%r14b
	sub	%r15b,%r15b
        nop
        nop
        // reg16 += reg16
	sub	%ax,%ax
	sub	%ax,%cx
	sub	%ax,%dx
	sub	%ax,%bx
	sub	%ax,%sp
	sub	%ax,%bp
	sub	%ax,%si
	sub	%ax,%di
	sub	%ax,%r8w
	sub	%ax,%r9w
	sub	%ax,%r10w
	sub	%ax,%r11w
	sub	%ax,%r12w
	sub	%ax,%r13w
	sub	%ax,%r14w
	sub	%ax,%r15w
	nop
	sub	%cx,%ax
	sub	%cx,%cx
	sub	%cx,%dx
	sub	%cx,%bx
	sub	%cx,%sp
	sub	%cx,%bp
	sub	%cx,%si
	sub	%cx,%di
	sub	%cx,%r8w
	sub	%cx,%r9w
	sub	%cx,%r10w
	sub	%cx,%r11w
	sub	%cx,%r12w
	sub	%cx,%r13w
	sub	%cx,%r14w
	sub	%cx,%r15w
	nop
	sub	%dx,%ax
	sub	%dx,%cx
	sub	%dx,%dx
	sub	%dx,%bx
	sub	%dx,%sp
	sub	%dx,%bp
	sub	%dx,%si
	sub	%dx,%di
	sub	%dx,%r8w
	sub	%dx,%r9w
	sub	%dx,%r10w
	sub	%dx,%r11w
	sub	%dx,%r12w
	sub	%dx,%r13w
	sub	%dx,%r14w
	sub	%dx,%r15w
	nop
	sub	%bx,%ax
	sub	%bx,%cx
	sub	%bx,%dx
	sub	%bx,%bx
	sub	%bx,%sp
	sub	%bx,%bp
	sub	%bx,%si
	sub	%bx,%di
	sub	%bx,%r8w
	sub	%bx,%r9w
	sub	%bx,%r10w
	sub	%bx,%r11w
	sub	%bx,%r12w
	sub	%bx,%r13w
	sub	%bx,%r14w
	sub	%bx,%r15w
	nop
	sub	%sp,%ax
	sub	%sp,%cx
	sub	%sp,%dx
	sub	%sp,%bx
	sub	%sp,%sp
	sub	%sp,%bp
	sub	%sp,%si
	sub	%sp,%di
	sub	%sp,%r8w
	sub	%sp,%r9w
	sub	%sp,%r10w
	sub	%sp,%r11w
	sub	%sp,%r12w
	sub	%sp,%r13w
	sub	%sp,%r14w
	sub	%sp,%r15w
	nop
	sub	%bp,%ax
	sub	%bp,%cx
	sub	%bp,%dx
	sub	%bp,%bx
	sub	%bp,%sp
	sub	%bp,%bp
	sub	%bp,%si
	sub	%bp,%di
	sub	%bp,%r8w
	sub	%bp,%r9w
	sub	%bp,%r10w
	sub	%bp,%r11w
	sub	%bp,%r12w
	sub	%bp,%r13w
	sub	%bp,%r14w
	sub	%bp,%r15w
	nop
	sub	%si,%ax
	sub	%si,%cx
	sub	%si,%dx
	sub	%si,%bx
	sub	%si,%sp
	sub	%si,%bp
	sub	%si,%si
	sub	%si,%di
	sub	%si,%r8w
	sub	%si,%r9w
	sub	%si,%r10w
	sub	%si,%r11w
	sub	%si,%r12w
	sub	%si,%r13w
	sub	%si,%r14w
	sub	%si,%r15w
	nop
	sub	%di,%ax
	sub	%di,%cx
	sub	%di,%dx
	sub	%di,%bx
	sub	%di,%sp
	sub	%di,%bp
	sub	%di,%si
	sub	%di,%di
	sub	%di,%r8w
	sub	%di,%r9w
	sub	%di,%r10w
	sub	%di,%r11w
	sub	%di,%r12w
	sub	%di,%r13w
	sub	%di,%r14w
	sub	%di,%r15w
	nop
	sub	%r8w, %ax
	sub	%r8w, %cx
	sub	%r8w, %dx
	sub	%r8w, %bx
	sub	%r8w, %sp
	sub	%r8w, %bp
	sub	%r8w, %si
	sub	%r8w, %di
	sub	%r8w, %r8w
	sub	%r8w, %r9w
	sub	%r8w, %r10w
	sub	%r8w, %r11w
	sub	%r8w, %r12w
	sub	%r8w, %r13w
	sub	%r8w, %r14w
	sub	%r8w, %r15w
	nop
	sub	%r9w, %ax
	sub	%r9w, %cx
	sub	%r9w, %dx
	sub	%r9w, %bx
	sub	%r9w, %sp
	sub	%r9w, %bp
	sub	%r9w, %si
	sub	%r9w, %di
	sub	%r9w, %r8w
	sub	%r9w, %r9w
	sub	%r9w, %r10w
	sub	%r9w, %r11w
	sub	%r9w, %r12w
	sub	%r9w, %r13w
	sub	%r9w, %r14w
	sub	%r9w, %r15w
	nop
	sub	%r10w,%ax
	sub	%r10w,%cx
	sub	%r10w,%dx
	sub	%r10w,%bx
	sub	%r10w,%sp
	sub	%r10w,%bp
	sub	%r10w,%si
	sub	%r10w,%di
	sub	%r10w,%r8w
	sub	%r10w,%r9w
	sub	%r10w,%r10w
	sub	%r10w,%r11w
	sub	%r10w,%r12w
	sub	%r10w,%r13w
	sub	%r10w,%r14w
	sub	%r10w,%r15w
	nop
	sub	%r11w,%ax
	sub	%r11w,%cx
	sub	%r11w,%dx
	sub	%r11w,%bx
	sub	%r11w,%sp
	sub	%r11w,%bp
	sub	%r11w,%si
	sub	%r11w,%di
	sub	%r11w,%r8w
	sub	%r11w,%r9w
	sub	%r11w,%r10w
	sub	%r11w,%r11w
	sub	%r11w,%r12w
	sub	%r11w,%r13w
	sub	%r11w,%r14w
	sub	%r11w,%r15w
	nop
	sub	%r12w,%ax
	sub	%r12w,%cx
	sub	%r12w,%dx
	sub	%r12w,%bx
	sub	%r12w,%sp
	sub	%r12w,%bp
	sub	%r12w,%si
	sub	%r12w,%di
	sub	%r12w,%r8w
	sub	%r12w,%r9w
	sub	%r12w,%r10w
	sub	%r12w,%r11w
	sub	%r12w,%r12w
	sub	%r12w,%r13w
	sub	%r12w,%r14w
	sub	%r12w,%r15w
	nop
	sub	%r13w,%ax
	sub	%r13w,%cx
	sub	%r13w,%dx
	sub	%r13w,%bx
	sub	%r13w,%sp
	sub	%r13w,%bp
	sub	%r13w,%si
	sub	%r13w,%di
	sub	%r13w,%r8w
	sub	%r13w,%r9w
	sub	%r13w,%r10w
	sub	%r13w,%r11w
	sub	%r13w,%r12w
	sub	%r13w,%r13w
	sub	%r13w,%r14w
	sub	%r13w,%r15w
	nop
	sub	%r14w,%ax
	sub	%r14w,%cx
	sub	%r14w,%dx
	sub	%r14w,%bx
	sub	%r14w,%sp
	sub	%r14w,%bp
	sub	%r14w,%si
	sub	%r14w,%di
	sub	%r14w,%r8w
	sub	%r14w,%r9w
	sub	%r14w,%r10w
	sub	%r14w,%r11w
	sub	%r14w,%r12w
	sub	%r14w,%r13w
	sub	%r14w,%r14w
	sub	%r14w,%r15w
	nop
	sub	%r15w,%ax
	sub	%r15w,%cx
	sub	%r15w,%dx
	sub	%r15w,%bx
	sub	%r15w,%sp
	sub	%r15w,%bp
	sub	%r15w,%si
	sub	%r15w,%di
	sub	%r15w,%r8w
	sub	%r15w,%r9w
	sub	%r15w,%r10w
	sub	%r15w,%r11w
	sub	%r15w,%r12w
	sub	%r15w,%r13w
	sub	%r15w,%r14w
	sub	%r15w,%r15w
        nop
        nop
        // reg32 += reg32
	sub	%eax,%eax
	sub	%eax,%ecx
	sub	%eax,%edx
	sub	%eax,%ebx
	sub	%eax,%esp
	sub	%eax,%ebp
	sub	%eax,%esi
	sub	%eax,%edi
	sub	%eax,%r8d
	sub	%eax,%r9d
	sub	%eax,%r10d
	sub	%eax,%r11d
	sub	%eax,%r12d
	sub	%eax,%r13d
	sub	%eax,%r14d
	sub	%eax,%r15d
	nop
	sub	%ecx,%eax
	sub	%ecx,%ecx
	sub	%ecx,%edx
	sub	%ecx,%ebx
	sub	%ecx,%esp
	sub	%ecx,%ebp
	sub	%ecx,%esi
	sub	%ecx,%edi
	sub	%ecx,%r8d
	sub	%ecx,%r9d
	sub	%ecx,%r10d
	sub	%ecx,%r11d
	sub	%ecx,%r12d
	sub	%ecx,%r13d
	sub	%ecx,%r14d
	sub	%ecx,%r15d
	nop
	sub	%edx,%eax
	sub	%edx,%ecx
	sub	%edx,%edx
	sub	%edx,%ebx
	sub	%edx,%esp
	sub	%edx,%ebp
	sub	%edx,%esi
	sub	%edx,%edi
	sub	%edx,%r8d
	sub	%edx,%r9d
	sub	%edx,%r10d
	sub	%edx,%r11d
	sub	%edx,%r12d
	sub	%edx,%r13d
	sub	%edx,%r14d
	sub	%edx,%r15d
	nop
	sub	%ebx,%eax
	sub	%ebx,%ecx
	sub	%ebx,%edx
	sub	%ebx,%ebx
	sub	%ebx,%esp
	sub	%ebx,%ebp
	sub	%ebx,%esi
	sub	%ebx,%edi
	sub	%ebx,%r8d
	sub	%ebx,%r9d
	sub	%ebx,%r10d
	sub	%ebx,%r11d
	sub	%ebx,%r12d
	sub	%ebx,%r13d
	sub	%ebx,%r14d
	sub	%ebx,%r15d
	nop
	sub	%esp,%eax
	sub	%esp,%ecx
	sub	%esp,%edx
	sub	%esp,%ebx
	sub	%esp,%esp
	sub	%esp,%ebp
	sub	%esp,%esi
	sub	%esp,%edi
	sub	%esp,%r8d
	sub	%esp,%r9d
	sub	%esp,%r10d
	sub	%esp,%r11d
	sub	%esp,%r12d
	sub	%esp,%r13d
	sub	%esp,%r14d
	sub	%esp,%r15d
	nop
	sub	%ebp,%eax
	sub	%ebp,%ecx
	sub	%ebp,%edx
	sub	%ebp,%ebx
	sub	%ebp,%esp
	sub	%ebp,%ebp
	sub	%ebp,%esi
	sub	%ebp,%edi
	sub	%ebp,%r8d
	sub	%ebp,%r9d
	sub	%ebp,%r10d
	sub	%ebp,%r11d
	sub	%ebp,%r12d
	sub	%ebp,%r13d
	sub	%ebp,%r14d
	sub	%ebp,%r15d
	nop
	sub	%esi,%eax
	sub	%esi,%ecx
	sub	%esi,%edx
	sub	%esi,%ebx
	sub	%esi,%esp
	sub	%esi,%ebp
	sub	%esi,%esi
	sub	%esi,%edi
	sub	%esi,%r8d
	sub	%esi,%r9d
	sub	%esi,%r10d
	sub	%esi,%r11d
	sub	%esi,%r12d
	sub	%esi,%r13d
	sub	%esi,%r14d
	sub	%esi,%r15d
	nop
	sub	%edi,%eax
	sub	%edi,%ecx
	sub	%edi,%edx
	sub	%edi,%ebx
	sub	%edi,%esp
	sub	%edi,%ebp
	sub	%edi,%esi
	sub	%edi,%edi
	sub	%edi,%r8d
	sub	%edi,%r9d
	sub	%edi,%r10d
	sub	%edi,%r11d
	sub	%edi,%r12d
	sub	%edi,%r13d
	sub	%edi,%r14d
	sub	%edi,%r15d
	nop
	sub	%r8d, %eax
	sub	%r8d, %ecx
	sub	%r8d, %edx
	sub	%r8d, %ebx
	sub	%r8d, %esp
	sub	%r8d, %ebp
	sub	%r8d, %esi
	sub	%r8d, %edi
	sub	%r8d, %r8d
	sub	%r8d, %r9d
	sub	%r8d, %r10d
	sub	%r8d, %r11d
	sub	%r8d, %r12d
	sub	%r8d, %r13d
	sub	%r8d, %r14d
	sub	%r8d, %r15d
	nop
	sub	%r9d, %eax
	sub	%r9d, %ecx
	sub	%r9d, %edx
	sub	%r9d, %ebx
	sub	%r9d, %esp
	sub	%r9d, %ebp
	sub	%r9d, %esi
	sub	%r9d, %edi
	sub	%r9d, %r8d
	sub	%r9d, %r9d
	sub	%r9d, %r10d
	sub	%r9d, %r11d
	sub	%r9d, %r12d
	sub	%r9d, %r13d
	sub	%r9d, %r14d
	sub	%r9d, %r15d
	nop
	sub	%r10d,%eax
	sub	%r10d,%ecx
	sub	%r10d,%edx
	sub	%r10d,%ebx
	sub	%r10d,%esp
	sub	%r10d,%ebp
	sub	%r10d,%esi
	sub	%r10d,%edi
	sub	%r10d,%r8d
	sub	%r10d,%r9d
	sub	%r10d,%r10d
	sub	%r10d,%r11d
	sub	%r10d,%r12d
	sub	%r10d,%r13d
	sub	%r10d,%r14d
	sub	%r10d,%r15d
	nop
	sub	%r11d,%eax
	sub	%r11d,%ecx
	sub	%r11d,%edx
	sub	%r11d,%ebx
	sub	%r11d,%esp
	sub	%r11d,%ebp
	sub	%r11d,%esi
	sub	%r11d,%edi
	sub	%r11d,%r8d
	sub	%r11d,%r9d
	sub	%r11d,%r10d
	sub	%r11d,%r11d
	sub	%r11d,%r12d
	sub	%r11d,%r13d
	sub	%r11d,%r14d
	sub	%r11d,%r15d
	nop
	sub	%r12d,%eax
	sub	%r12d,%ecx
	sub	%r12d,%edx
	sub	%r12d,%ebx
	sub	%r12d,%esp
	sub	%r12d,%ebp
	sub	%r12d,%esi
	sub	%r12d,%edi
	sub	%r12d,%r8d
	sub	%r12d,%r9d
	sub	%r12d,%r10d
	sub	%r12d,%r11d
	sub	%r12d,%r12d
	sub	%r12d,%r13d
	sub	%r12d,%r14d
	sub	%r12d,%r15d
	nop
	sub	%r13d,%eax
	sub	%r13d,%ecx
	sub	%r13d,%edx
	sub	%r13d,%ebx
	sub	%r13d,%esp
	sub	%r13d,%ebp
	sub	%r13d,%esi
	sub	%r13d,%edi
	sub	%r13d,%r8d
	sub	%r13d,%r9d
	sub	%r13d,%r10d
	sub	%r13d,%r11d
	sub	%r13d,%r12d
	sub	%r13d,%r13d
	sub	%r13d,%r14d
	sub	%r13d,%r15d
	nop
	sub	%r14d,%eax
	sub	%r14d,%ecx
	sub	%r14d,%edx
	sub	%r14d,%ebx
	sub	%r14d,%esp
	sub	%r14d,%ebp
	sub	%r14d,%esi
	sub	%r14d,%edi
	sub	%r14d,%r8d
	sub	%r14d,%r9d
	sub	%r14d,%r10d
	sub	%r14d,%r11d
	sub	%r14d,%r12d
	sub	%r14d,%r13d
	sub	%r14d,%r14d
	sub	%r14d,%r15d
	nop
	sub	%r15d,%eax
	sub	%r15d,%ecx
	sub	%r15d,%edx
	sub	%r15d,%ebx
	sub	%r15d,%esp
	sub	%r15d,%ebp
	sub	%r15d,%esi
	sub	%r15d,%edi
	sub	%r15d,%r8d
	sub	%r15d,%r9d
	sub	%r15d,%r10d
	sub	%r15d,%r11d
	sub	%r15d,%r12d
	sub	%r15d,%r13d
	sub	%r15d,%r14d
	sub	%r15d,%r15d
        nop
        nop
        // reg64 += reg64
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
        // mem8 += reg8
	sub	%al,(%rax)
	sub	%al,(%rcx)
	sub	%al,(%rdx)
	sub	%al,(%rbx)
	sub	%al,(%rsp)
	sub	%al,(%rbp)
	sub	%al,(%rsi)
	sub	%al,(%rdi)
	sub	%al,(%r8)
	sub	%al,(%r9)
	sub	%al,(%r10)
	sub	%al,(%r11)
	sub	%al,(%r12)
	sub	%al,(%r13)
	sub	%al,(%r14)
	sub	%al,(%r15)
	nop
	sub	%cl,(%rax)
	sub	%cl,(%rcx)
	sub	%cl,(%rdx)
	sub	%cl,(%rbx)
	sub	%cl,(%rsp)
	sub	%cl,(%rbp)
	sub	%cl,(%rsi)
	sub	%cl,(%rdi)
	sub	%cl,(%r8)
	sub	%cl,(%r9)
	sub	%cl,(%r10)
	sub	%cl,(%r11)
	sub	%cl,(%r12)
	sub	%cl,(%r13)
	sub	%cl,(%r14)
	sub	%cl,(%r15)
	nop
	sub	%dl,(%rax)
	sub	%dl,(%rcx)
	sub	%dl,(%rdx)
	sub	%dl,(%rbx)
	sub	%dl,(%rsp)
	sub	%dl,(%rbp)
	sub	%dl,(%rsi)
	sub	%dl,(%rdi)
	sub	%dl,(%r8)
	sub	%dl,(%r9)
	sub	%dl,(%r10)
	sub	%dl,(%r11)
	sub	%dl,(%r12)
	sub	%dl,(%r13)
	sub	%dl,(%r14)
	sub	%dl,(%r15)
	nop
	sub	%bl,(%rax)
	sub	%bl,(%rcx)
	sub	%bl,(%rdx)
	sub	%bl,(%rbx)
	sub	%bl,(%rsp)
	sub	%bl,(%rbp)
	sub	%bl,(%rsi)
	sub	%bl,(%rdi)
	sub	%bl,(%r8)
	sub	%bl,(%r9)
	sub	%bl,(%r10)
	sub	%bl,(%r11)
	sub	%bl,(%r12)
	sub	%bl,(%r13)
	sub	%bl,(%r14)
	sub	%bl,(%r15)
	nop
	sub	%spl,(%rax)
	sub	%spl,(%rcx)
	sub	%spl,(%rdx)
	sub	%spl,(%rbx)
	sub	%spl,(%rsp)
	sub	%spl,(%rbp)
	sub	%spl,(%rsi)
	sub	%spl,(%rdi)
	sub	%spl,(%r8)
	sub	%spl,(%r9)
	sub	%spl,(%r10)
	sub	%spl,(%r11)
	sub	%spl,(%r12)
	sub	%spl,(%r13)
	sub	%spl,(%r14)
	sub	%spl,(%r15)
	nop
	sub	%bpl,(%rax)
	sub	%bpl,(%rcx)
	sub	%bpl,(%rdx)
	sub	%bpl,(%rbx)
	sub	%bpl,(%rsp)
	sub	%bpl,(%rbp)
	sub	%bpl,(%rsi)
	sub	%bpl,(%rdi)
	sub	%bpl,(%r8)
	sub	%bpl,(%r9)
	sub	%bpl,(%r10)
	sub	%bpl,(%r11)
	sub	%bpl,(%r12)
	sub	%bpl,(%r13)
	sub	%bpl,(%r14)
	sub	%bpl,(%r15)
	nop
	sub	%sil,(%rax)
	sub	%sil,(%rcx)
	sub	%sil,(%rdx)
	sub	%sil,(%rbx)
	sub	%sil,(%rsp)
	sub	%sil,(%rbp)
	sub	%sil,(%rsi)
	sub	%sil,(%rdi)
	sub	%sil,(%r8)
	sub	%sil,(%r9)
	sub	%sil,(%r10)
	sub	%sil,(%r11)
	sub	%sil,(%r12)
	sub	%sil,(%r13)
	sub	%sil,(%r14)
	sub	%sil,(%r15)
	nop
	sub	%dil,(%rax)
	sub	%dil,(%rcx)
	sub	%dil,(%rdx)
	sub	%dil,(%rbx)
	sub	%dil,(%rsp)
	sub	%dil,(%rbp)
	sub	%dil,(%rsi)
	sub	%dil,(%rdi)
	sub	%dil,(%r8)
	sub	%dil,(%r9)
	sub	%dil,(%r10)
	sub	%dil,(%r11)
	sub	%dil,(%r12)
	sub	%dil,(%r13)
	sub	%dil,(%r14)
	sub	%dil,(%r15)
	nop
	sub	%r8b, (%rax)
	sub	%r8b, (%rcx)
	sub	%r8b, (%rdx)
	sub	%r8b, (%rbx)
	sub	%r8b, (%rsp)
	sub	%r8b, (%rbp)
	sub	%r8b, (%rsi)
	sub	%r8b, (%rdi)
	sub	%r8b, (%r8)
	sub	%r8b, (%r9)
	sub	%r8b, (%r10)
	sub	%r8b, (%r11)
	sub	%r8b, (%r12)
	sub	%r8b, (%r13)
	sub	%r8b, (%r14)
	sub	%r8b, (%r15)
	nop
	sub	%r9b, (%rax)
	sub	%r9b, (%rcx)
	sub	%r9b, (%rdx)
	sub	%r9b, (%rbx)
	sub	%r9b, (%rsp)
	sub	%r9b, (%rbp)
	sub	%r9b, (%rsi)
	sub	%r9b, (%rdi)
	sub	%r9b, (%r8)
	sub	%r9b, (%r9)
	sub	%r9b, (%r10)
	sub	%r9b, (%r11)
	sub	%r9b, (%r12)
	sub	%r9b, (%r13)
	sub	%r9b, (%r14)
	sub	%r9b, (%r15)
	nop
	sub	%r10b,(%rax)
	sub	%r10b,(%rcx)
	sub	%r10b,(%rdx)
	sub	%r10b,(%rbx)
	sub	%r10b,(%rsp)
	sub	%r10b,(%rbp)
	sub	%r10b,(%rsi)
	sub	%r10b,(%rdi)
	sub	%r10b,(%r8)
	sub	%r10b,(%r9)
	sub	%r10b,(%r10)
	sub	%r10b,(%r11)
	sub	%r10b,(%r12)
	sub	%r10b,(%r13)
	sub	%r10b,(%r14)
	sub	%r10b,(%r15)
	nop
	sub	%r11b,(%rax)
	sub	%r11b,(%rcx)
	sub	%r11b,(%rdx)
	sub	%r11b,(%rbx)
	sub	%r11b,(%rsp)
	sub	%r11b,(%rbp)
	sub	%r11b,(%rsi)
	sub	%r11b,(%rdi)
	sub	%r11b,(%r8)
	sub	%r11b,(%r9)
	sub	%r11b,(%r10)
	sub	%r11b,(%r11)
	sub	%r11b,(%r12)
	sub	%r11b,(%r13)
	sub	%r11b,(%r14)
	sub	%r11b,(%r15)
	nop
	sub	%r12b,(%rax)
	sub	%r12b,(%rcx)
	sub	%r12b,(%rdx)
	sub	%r12b,(%rbx)
	sub	%r12b,(%rsp)
	sub	%r12b,(%rbp)
	sub	%r12b,(%rsi)
	sub	%r12b,(%rdi)
	sub	%r12b,(%r8)
	sub	%r12b,(%r9)
	sub	%r12b,(%r10)
	sub	%r12b,(%r11)
	sub	%r12b,(%r12)
	sub	%r12b,(%r13)
	sub	%r12b,(%r14)
	sub	%r12b,(%r15)
	nop
	sub	%r13b,(%rax)
	sub	%r13b,(%rcx)
	sub	%r13b,(%rdx)
	sub	%r13b,(%rbx)
	sub	%r13b,(%rsp)
	sub	%r13b,(%rbp)
	sub	%r13b,(%rsi)
	sub	%r13b,(%rdi)
	sub	%r13b,(%r8)
	sub	%r13b,(%r9)
	sub	%r13b,(%r10)
	sub	%r13b,(%r11)
	sub	%r13b,(%r12)
	sub	%r13b,(%r13)
	sub	%r13b,(%r14)
	sub	%r13b,(%r15)
	nop
	sub	%r14b,(%rax)
	sub	%r14b,(%rcx)
	sub	%r14b,(%rdx)
	sub	%r14b,(%rbx)
	sub	%r14b,(%rsp)
	sub	%r14b,(%rbp)
	sub	%r14b,(%rsi)
	sub	%r14b,(%rdi)
	sub	%r14b,(%r8)
	sub	%r14b,(%r9)
	sub	%r14b,(%r10)
	sub	%r14b,(%r11)
	sub	%r14b,(%r12)
	sub	%r14b,(%r13)
	sub	%r14b,(%r14)
	sub	%r14b,(%r15)
	nop
	sub	%r15b,(%rax)
	sub	%r15b,(%rcx)
	sub	%r15b,(%rdx)
	sub	%r15b,(%rbx)
	sub	%r15b,(%rsp)
	sub	%r15b,(%rbp)
	sub	%r15b,(%rsi)
	sub	%r15b,(%rdi)
	sub	%r15b,(%r8)
	sub	%r15b,(%r9)
	sub	%r15b,(%r10)
	sub	%r15b,(%r11)
	sub	%r15b,(%r12)
	sub	%r15b,(%r13)
	sub	%r15b,(%r14)
	sub	%r15b,(%r15)
        nop
        nop
        // mem16 += reg16
	sub	%ax,(%rax)
	sub	%ax,(%rcx)
	sub	%ax,(%rdx)
	sub	%ax,(%rbx)
	sub	%ax,(%rsp)
	sub	%ax,(%rbp)
	sub	%ax,(%rsi)
	sub	%ax,(%rdi)
	sub	%ax,(%r8)
	sub	%ax,(%r9)
	sub	%ax,(%r10)
	sub	%ax,(%r11)
	sub	%ax,(%r12)
	sub	%ax,(%r13)
	sub	%ax,(%r14)
	sub	%ax,(%r15)
	nop
	sub	%cx,(%rax)
	sub	%cx,(%rcx)
	sub	%cx,(%rdx)
	sub	%cx,(%rbx)
	sub	%cx,(%rsp)
	sub	%cx,(%rbp)
	sub	%cx,(%rsi)
	sub	%cx,(%rdi)
	sub	%cx,(%r8)
	sub	%cx,(%r9)
	sub	%cx,(%r10)
	sub	%cx,(%r11)
	sub	%cx,(%r12)
	sub	%cx,(%r13)
	sub	%cx,(%r14)
	sub	%cx,(%r15)
	nop
	sub	%dx,(%rax)
	sub	%dx,(%rcx)
	sub	%dx,(%rdx)
	sub	%dx,(%rbx)
	sub	%dx,(%rsp)
	sub	%dx,(%rbp)
	sub	%dx,(%rsi)
	sub	%dx,(%rdi)
	sub	%dx,(%r8)
	sub	%dx,(%r9)
	sub	%dx,(%r10)
	sub	%dx,(%r11)
	sub	%dx,(%r12)
	sub	%dx,(%r13)
	sub	%dx,(%r14)
	sub	%dx,(%r15)
	nop
	sub	%bx,(%rax)
	sub	%bx,(%rcx)
	sub	%bx,(%rdx)
	sub	%bx,(%rbx)
	sub	%bx,(%rsp)
	sub	%bx,(%rbp)
	sub	%bx,(%rsi)
	sub	%bx,(%rdi)
	sub	%bx,(%r8)
	sub	%bx,(%r9)
	sub	%bx,(%r10)
	sub	%bx,(%r11)
	sub	%bx,(%r12)
	sub	%bx,(%r13)
	sub	%bx,(%r14)
	sub	%bx,(%r15)
	nop
	sub	%sp,(%rax)
	sub	%sp,(%rcx)
	sub	%sp,(%rdx)
	sub	%sp,(%rbx)
	sub	%sp,(%rsp)
	sub	%sp,(%rbp)
	sub	%sp,(%rsi)
	sub	%sp,(%rdi)
	sub	%sp,(%r8)
	sub	%sp,(%r9)
	sub	%sp,(%r10)
	sub	%sp,(%r11)
	sub	%sp,(%r12)
	sub	%sp,(%r13)
	sub	%sp,(%r14)
	sub	%sp,(%r15)
	nop
	sub	%bp,(%rax)
	sub	%bp,(%rcx)
	sub	%bp,(%rdx)
	sub	%bp,(%rbx)
	sub	%bp,(%rsp)
	sub	%bp,(%rbp)
	sub	%bp,(%rsi)
	sub	%bp,(%rdi)
	sub	%bp,(%r8)
	sub	%bp,(%r9)
	sub	%bp,(%r10)
	sub	%bp,(%r11)
	sub	%bp,(%r12)
	sub	%bp,(%r13)
	sub	%bp,(%r14)
	sub	%bp,(%r15)
	nop
	sub	%si,(%rax)
	sub	%si,(%rcx)
	sub	%si,(%rdx)
	sub	%si,(%rbx)
	sub	%si,(%rsp)
	sub	%si,(%rbp)
	sub	%si,(%rsi)
	sub	%si,(%rdi)
	sub	%si,(%r8)
	sub	%si,(%r9)
	sub	%si,(%r10)
	sub	%si,(%r11)
	sub	%si,(%r12)
	sub	%si,(%r13)
	sub	%si,(%r14)
	sub	%si,(%r15)
	nop
	sub	%di,(%rax)
	sub	%di,(%rcx)
	sub	%di,(%rdx)
	sub	%di,(%rbx)
	sub	%di,(%rsp)
	sub	%di,(%rbp)
	sub	%di,(%rsi)
	sub	%di,(%rdi)
	sub	%di,(%r8)
	sub	%di,(%r9)
	sub	%di,(%r10)
	sub	%di,(%r11)
	sub	%di,(%r12)
	sub	%di,(%r13)
	sub	%di,(%r14)
	sub	%di,(%r15)
	nop
	sub	%r8w, (%rax)
	sub	%r8w, (%rcx)
	sub	%r8w, (%rdx)
	sub	%r8w, (%rbx)
	sub	%r8w, (%rsp)
	sub	%r8w, (%rbp)
	sub	%r8w, (%rsi)
	sub	%r8w, (%rdi)
	sub	%r8w, (%r8)
	sub	%r8w, (%r9)
	sub	%r8w, (%r10)
	sub	%r8w, (%r11)
	sub	%r8w, (%r12)
	sub	%r8w, (%r13)
	sub	%r8w, (%r14)
	sub	%r8w, (%r15)
	nop
	sub	%r9w, (%rax)
	sub	%r9w, (%rcx)
	sub	%r9w, (%rdx)
	sub	%r9w, (%rbx)
	sub	%r9w, (%rsp)
	sub	%r9w, (%rbp)
	sub	%r9w, (%rsi)
	sub	%r9w, (%rdi)
	sub	%r9w, (%r8)
	sub	%r9w, (%r9)
	sub	%r9w, (%r10)
	sub	%r9w, (%r11)
	sub	%r9w, (%r12)
	sub	%r9w, (%r13)
	sub	%r9w, (%r14)
	sub	%r9w, (%r15)
	nop
	sub	%r10w,(%rax)
	sub	%r10w,(%rcx)
	sub	%r10w,(%rdx)
	sub	%r10w,(%rbx)
	sub	%r10w,(%rsp)
	sub	%r10w,(%rbp)
	sub	%r10w,(%rsi)
	sub	%r10w,(%rdi)
	sub	%r10w,(%r8)
	sub	%r10w,(%r9)
	sub	%r10w,(%r10)
	sub	%r10w,(%r11)
	sub	%r10w,(%r12)
	sub	%r10w,(%r13)
	sub	%r10w,(%r14)
	sub	%r10w,(%r15)
	nop
	sub	%r11w,(%rax)
	sub	%r11w,(%rcx)
	sub	%r11w,(%rdx)
	sub	%r11w,(%rbx)
	sub	%r11w,(%rsp)
	sub	%r11w,(%rbp)
	sub	%r11w,(%rsi)
	sub	%r11w,(%rdi)
	sub	%r11w,(%r8)
	sub	%r11w,(%r9)
	sub	%r11w,(%r10)
	sub	%r11w,(%r11)
	sub	%r11w,(%r12)
	sub	%r11w,(%r13)
	sub	%r11w,(%r14)
	sub	%r11w,(%r15)
	nop
	sub	%r12w,(%rax)
	sub	%r12w,(%rcx)
	sub	%r12w,(%rdx)
	sub	%r12w,(%rbx)
	sub	%r12w,(%rsp)
	sub	%r12w,(%rbp)
	sub	%r12w,(%rsi)
	sub	%r12w,(%rdi)
	sub	%r12w,(%r8)
	sub	%r12w,(%r9)
	sub	%r12w,(%r10)
	sub	%r12w,(%r11)
	sub	%r12w,(%r12)
	sub	%r12w,(%r13)
	sub	%r12w,(%r14)
	sub	%r12w,(%r15)
	nop
	sub	%r13w,(%rax)
	sub	%r13w,(%rcx)
	sub	%r13w,(%rdx)
	sub	%r13w,(%rbx)
	sub	%r13w,(%rsp)
	sub	%r13w,(%rbp)
	sub	%r13w,(%rsi)
	sub	%r13w,(%rdi)
	sub	%r13w,(%r8)
	sub	%r13w,(%r9)
	sub	%r13w,(%r10)
	sub	%r13w,(%r11)
	sub	%r13w,(%r12)
	sub	%r13w,(%r13)
	sub	%r13w,(%r14)
	sub	%r13w,(%r15)
	nop
	sub	%r14w,(%rax)
	sub	%r14w,(%rcx)
	sub	%r14w,(%rdx)
	sub	%r14w,(%rbx)
	sub	%r14w,(%rsp)
	sub	%r14w,(%rbp)
	sub	%r14w,(%rsi)
	sub	%r14w,(%rdi)
	sub	%r14w,(%r8)
	sub	%r14w,(%r9)
	sub	%r14w,(%r10)
	sub	%r14w,(%r11)
	sub	%r14w,(%r12)
	sub	%r14w,(%r13)
	sub	%r14w,(%r14)
	sub	%r14w,(%r15)
	nop
	sub	%r15w,(%rax)
	sub	%r15w,(%rcx)
	sub	%r15w,(%rdx)
	sub	%r15w,(%rbx)
	sub	%r15w,(%rsp)
	sub	%r15w,(%rbp)
	sub	%r15w,(%rsi)
	sub	%r15w,(%rdi)
	sub	%r15w,(%r8)
	sub	%r15w,(%r9)
	sub	%r15w,(%r10)
	sub	%r15w,(%r11)
	sub	%r15w,(%r12)
	sub	%r15w,(%r13)
	sub	%r15w,(%r14)
	sub	%r15w,(%r15)
        nop
        nop
        // mem32 += reg32
	sub	%eax,(%rax)
	sub	%eax,(%rcx)
	sub	%eax,(%rdx)
	sub	%eax,(%rbx)
	sub	%eax,(%rsp)
	sub	%eax,(%rbp)
	sub	%eax,(%rsi)
	sub	%eax,(%rdi)
	sub	%eax,(%r8)
	sub	%eax,(%r9)
	sub	%eax,(%r10)
	sub	%eax,(%r11)
	sub	%eax,(%r12)
	sub	%eax,(%r13)
	sub	%eax,(%r14)
	sub	%eax,(%r15)
	nop
	sub	%ecx,(%rax)
	sub	%ecx,(%rcx)
	sub	%ecx,(%rdx)
	sub	%ecx,(%rbx)
	sub	%ecx,(%rsp)
	sub	%ecx,(%rbp)
	sub	%ecx,(%rsi)
	sub	%ecx,(%rdi)
	sub	%ecx,(%r8)
	sub	%ecx,(%r9)
	sub	%ecx,(%r10)
	sub	%ecx,(%r11)
	sub	%ecx,(%r12)
	sub	%ecx,(%r13)
	sub	%ecx,(%r14)
	sub	%ecx,(%r15)
	nop
	sub	%edx,(%rax)
	sub	%edx,(%rcx)
	sub	%edx,(%rdx)
	sub	%edx,(%rbx)
	sub	%edx,(%rsp)
	sub	%edx,(%rbp)
	sub	%edx,(%rsi)
	sub	%edx,(%rdi)
	sub	%edx,(%r8)
	sub	%edx,(%r9)
	sub	%edx,(%r10)
	sub	%edx,(%r11)
	sub	%edx,(%r12)
	sub	%edx,(%r13)
	sub	%edx,(%r14)
	sub	%edx,(%r15)
	nop
	sub	%ebx,(%rax)
	sub	%ebx,(%rcx)
	sub	%ebx,(%rdx)
	sub	%ebx,(%rbx)
	sub	%ebx,(%rsp)
	sub	%ebx,(%rbp)
	sub	%ebx,(%rsi)
	sub	%ebx,(%rdi)
	sub	%ebx,(%r8)
	sub	%ebx,(%r9)
	sub	%ebx,(%r10)
	sub	%ebx,(%r11)
	sub	%ebx,(%r12)
	sub	%ebx,(%r13)
	sub	%ebx,(%r14)
	sub	%ebx,(%r15)
	nop
	sub	%esp,(%rax)
	sub	%esp,(%rcx)
	sub	%esp,(%rdx)
	sub	%esp,(%rbx)
	sub	%esp,(%rsp)
	sub	%esp,(%rbp)
	sub	%esp,(%rsi)
	sub	%esp,(%rdi)
	sub	%esp,(%r8)
	sub	%esp,(%r9)
	sub	%esp,(%r10)
	sub	%esp,(%r11)
	sub	%esp,(%r12)
	sub	%esp,(%r13)
	sub	%esp,(%r14)
	sub	%esp,(%r15)
	nop
	sub	%ebp,(%rax)
	sub	%ebp,(%rcx)
	sub	%ebp,(%rdx)
	sub	%ebp,(%rbx)
	sub	%ebp,(%rsp)
	sub	%ebp,(%rbp)
	sub	%ebp,(%rsi)
	sub	%ebp,(%rdi)
	sub	%ebp,(%r8)
	sub	%ebp,(%r9)
	sub	%ebp,(%r10)
	sub	%ebp,(%r11)
	sub	%ebp,(%r12)
	sub	%ebp,(%r13)
	sub	%ebp,(%r14)
	sub	%ebp,(%r15)
	nop
	sub	%esi,(%rax)
	sub	%esi,(%rcx)
	sub	%esi,(%rdx)
	sub	%esi,(%rbx)
	sub	%esi,(%rsp)
	sub	%esi,(%rbp)
	sub	%esi,(%rsi)
	sub	%esi,(%rdi)
	sub	%esi,(%r8)
	sub	%esi,(%r9)
	sub	%esi,(%r10)
	sub	%esi,(%r11)
	sub	%esi,(%r12)
	sub	%esi,(%r13)
	sub	%esi,(%r14)
	sub	%esi,(%r15)
	nop
	sub	%edi,(%rax)
	sub	%edi,(%rcx)
	sub	%edi,(%rdx)
	sub	%edi,(%rbx)
	sub	%edi,(%rsp)
	sub	%edi,(%rbp)
	sub	%edi,(%rsi)
	sub	%edi,(%rdi)
	sub	%edi,(%r8)
	sub	%edi,(%r9)
	sub	%edi,(%r10)
	sub	%edi,(%r11)
	sub	%edi,(%r12)
	sub	%edi,(%r13)
	sub	%edi,(%r14)
	sub	%edi,(%r15)
	nop
	sub	%r8d, (%rax)
	sub	%r8d, (%rcx)
	sub	%r8d, (%rdx)
	sub	%r8d, (%rbx)
	sub	%r8d, (%rsp)
	sub	%r8d, (%rbp)
	sub	%r8d, (%rsi)
	sub	%r8d, (%rdi)
	sub	%r8d, (%r8)
	sub	%r8d, (%r9)
	sub	%r8d, (%r10)
	sub	%r8d, (%r11)
	sub	%r8d, (%r12)
	sub	%r8d, (%r13)
	sub	%r8d, (%r14)
	sub	%r8d, (%r15)
	nop
	sub	%r9d, (%rax)
	sub	%r9d, (%rcx)
	sub	%r9d, (%rdx)
	sub	%r9d, (%rbx)
	sub	%r9d, (%rsp)
	sub	%r9d, (%rbp)
	sub	%r9d, (%rsi)
	sub	%r9d, (%rdi)
	sub	%r9d, (%r8)
	sub	%r9d, (%r9)
	sub	%r9d, (%r10)
	sub	%r9d, (%r11)
	sub	%r9d, (%r12)
	sub	%r9d, (%r13)
	sub	%r9d, (%r14)
	sub	%r9d, (%r15)
	nop
	sub	%r10d,(%rax)
	sub	%r10d,(%rcx)
	sub	%r10d,(%rdx)
	sub	%r10d,(%rbx)
	sub	%r10d,(%rsp)
	sub	%r10d,(%rbp)
	sub	%r10d,(%rsi)
	sub	%r10d,(%rdi)
	sub	%r10d,(%r8)
	sub	%r10d,(%r9)
	sub	%r10d,(%r10)
	sub	%r10d,(%r11)
	sub	%r10d,(%r12)
	sub	%r10d,(%r13)
	sub	%r10d,(%r14)
	sub	%r10d,(%r15)
	nop
	sub	%r11d,(%rax)
	sub	%r11d,(%rcx)
	sub	%r11d,(%rdx)
	sub	%r11d,(%rbx)
	sub	%r11d,(%rsp)
	sub	%r11d,(%rbp)
	sub	%r11d,(%rsi)
	sub	%r11d,(%rdi)
	sub	%r11d,(%r8)
	sub	%r11d,(%r9)
	sub	%r11d,(%r10)
	sub	%r11d,(%r11)
	sub	%r11d,(%r12)
	sub	%r11d,(%r13)
	sub	%r11d,(%r14)
	sub	%r11d,(%r15)
	nop
	sub	%r12d,(%rax)
	sub	%r12d,(%rcx)
	sub	%r12d,(%rdx)
	sub	%r12d,(%rbx)
	sub	%r12d,(%rsp)
	sub	%r12d,(%rbp)
	sub	%r12d,(%rsi)
	sub	%r12d,(%rdi)
	sub	%r12d,(%r8)
	sub	%r12d,(%r9)
	sub	%r12d,(%r10)
	sub	%r12d,(%r11)
	sub	%r12d,(%r12)
	sub	%r12d,(%r13)
	sub	%r12d,(%r14)
	sub	%r12d,(%r15)
	nop
	sub	%r13d,(%rax)
	sub	%r13d,(%rcx)
	sub	%r13d,(%rdx)
	sub	%r13d,(%rbx)
	sub	%r13d,(%rsp)
	sub	%r13d,(%rbp)
	sub	%r13d,(%rsi)
	sub	%r13d,(%rdi)
	sub	%r13d,(%r8)
	sub	%r13d,(%r9)
	sub	%r13d,(%r10)
	sub	%r13d,(%r11)
	sub	%r13d,(%r12)
	sub	%r13d,(%r13)
	sub	%r13d,(%r14)
	sub	%r13d,(%r15)
	nop
	sub	%r14d,(%rax)
	sub	%r14d,(%rcx)
	sub	%r14d,(%rdx)
	sub	%r14d,(%rbx)
	sub	%r14d,(%rsp)
	sub	%r14d,(%rbp)
	sub	%r14d,(%rsi)
	sub	%r14d,(%rdi)
	sub	%r14d,(%r8)
	sub	%r14d,(%r9)
	sub	%r14d,(%r10)
	sub	%r14d,(%r11)
	sub	%r14d,(%r12)
	sub	%r14d,(%r13)
	sub	%r14d,(%r14)
	sub	%r14d,(%r15)
	nop
	sub	%r15d,(%rax)
	sub	%r15d,(%rcx)
	sub	%r15d,(%rdx)
	sub	%r15d,(%rbx)
	sub	%r15d,(%rsp)
	sub	%r15d,(%rbp)
	sub	%r15d,(%rsi)
	sub	%r15d,(%rdi)
	sub	%r15d,(%r8)
	sub	%r15d,(%r9)
	sub	%r15d,(%r10)
	sub	%r15d,(%r11)
	sub	%r15d,(%r12)
	sub	%r15d,(%r13)
	sub	%r15d,(%r14)
	sub	%r15d,(%r15)
        nop
        nop
        // mem64 += reg64
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


