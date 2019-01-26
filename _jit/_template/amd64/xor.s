	.file	"xor.s"
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
	.globl	Xor
	.type	Xor, @function
Xor:
	.cfi_startproc

        // reg8 += reg8
	xor	%al,%al
	xor	%al,%cl
	xor	%al,%dl
	xor	%al,%bl
	xor	%al,%spl
	xor	%al,%bpl
	xor	%al,%sil
	xor	%al,%dil
	xor	%al,%r8b
	xor	%al,%r9b
	xor	%al,%r10b
	xor	%al,%r11b
	xor	%al,%r12b
	xor	%al,%r13b
	xor	%al,%r14b
	xor	%al,%r15b
	nop
	xor	%cl,%al
	xor	%cl,%cl
	xor	%cl,%dl
	xor	%cl,%bl
	xor	%cl,%spl
	xor	%cl,%bpl
	xor	%cl,%sil
	xor	%cl,%dil
	xor	%cl,%r8b
	xor	%cl,%r9b
	xor	%cl,%r10b
	xor	%cl,%r11b
	xor	%cl,%r12b
	xor	%cl,%r13b
	xor	%cl,%r14b
	xor	%cl,%r15b
	nop
	xor	%dl,%al
	xor	%dl,%cl
	xor	%dl,%dl
	xor	%dl,%bl
	xor	%dl,%spl
	xor	%dl,%bpl
	xor	%dl,%sil
	xor	%dl,%dil
	xor	%dl,%r8b
	xor	%dl,%r9b
	xor	%dl,%r10b
	xor	%dl,%r11b
	xor	%dl,%r12b
	xor	%dl,%r13b
	xor	%dl,%r14b
	xor	%dl,%r15b
	nop
	xor	%bl,%al
	xor	%bl,%cl
	xor	%bl,%dl
	xor	%bl,%bl
	xor	%bl,%spl
	xor	%bl,%bpl
	xor	%bl,%sil
	xor	%bl,%dil
	xor	%bl,%r8b
	xor	%bl,%r9b
	xor	%bl,%r10b
	xor	%bl,%r11b
	xor	%bl,%r12b
	xor	%bl,%r13b
	xor	%bl,%r14b
	xor	%bl,%r15b
	nop
	xor	%spl,%al
	xor	%spl,%cl
	xor	%spl,%dl
	xor	%spl,%bl
	xor	%spl,%spl
	xor	%spl,%bpl
	xor	%spl,%sil
	xor	%spl,%dil
	xor	%spl,%r8b
	xor	%spl,%r9b
	xor	%spl,%r10b
	xor	%spl,%r11b
	xor	%spl,%r12b
	xor	%spl,%r13b
	xor	%spl,%r14b
	xor	%spl,%r15b
	nop
	xor	%bpl,%al
	xor	%bpl,%cl
	xor	%bpl,%dl
	xor	%bpl,%bl
	xor	%bpl,%spl
	xor	%bpl,%bpl
	xor	%bpl,%sil
	xor	%bpl,%dil
	xor	%bpl,%r8b
	xor	%bpl,%r9b
	xor	%bpl,%r10b
	xor	%bpl,%r11b
	xor	%bpl,%r12b
	xor	%bpl,%r13b
	xor	%bpl,%r14b
	xor	%bpl,%r15b
	nop
	xor	%sil,%al
	xor	%sil,%cl
	xor	%sil,%dl
	xor	%sil,%bl
	xor	%sil,%spl
	xor	%sil,%bpl
	xor	%sil,%sil
	xor	%sil,%dil
	xor	%sil,%r8b
	xor	%sil,%r9b
	xor	%sil,%r10b
	xor	%sil,%r11b
	xor	%sil,%r12b
	xor	%sil,%r13b
	xor	%sil,%r14b
	xor	%sil,%r15b
	nop
	xor	%dil,%al
	xor	%dil,%cl
	xor	%dil,%dl
	xor	%dil,%bl
	xor	%dil,%spl
	xor	%dil,%bpl
	xor	%dil,%sil
	xor	%dil,%dil
	xor	%dil,%r8b
	xor	%dil,%r9b
	xor	%dil,%r10b
	xor	%dil,%r11b
	xor	%dil,%r12b
	xor	%dil,%r13b
	xor	%dil,%r14b
	xor	%dil,%r15b
	nop
	xor	%r8b, %al
	xor	%r8b, %cl
	xor	%r8b, %dl
	xor	%r8b, %bl
	xor	%r8b, %spl
	xor	%r8b, %bpl
	xor	%r8b, %sil
	xor	%r8b, %dil
	xor	%r8b, %r8b
	xor	%r8b, %r9b
	xor	%r8b, %r10b
	xor	%r8b, %r11b
	xor	%r8b, %r12b
	xor	%r8b, %r13b
	xor	%r8b, %r14b
	xor	%r8b, %r15b
	nop
	xor	%r9b, %al
	xor	%r9b, %cl
	xor	%r9b, %dl
	xor	%r9b, %bl
	xor	%r9b, %spl
	xor	%r9b, %bpl
	xor	%r9b, %sil
	xor	%r9b, %dil
	xor	%r9b, %r8b
	xor	%r9b, %r9b
	xor	%r9b, %r10b
	xor	%r9b, %r11b
	xor	%r9b, %r12b
	xor	%r9b, %r13b
	xor	%r9b, %r14b
	xor	%r9b, %r15b
	nop
	xor	%r10b,%al
	xor	%r10b,%cl
	xor	%r10b,%dl
	xor	%r10b,%bl
	xor	%r10b,%spl
	xor	%r10b,%bpl
	xor	%r10b,%sil
	xor	%r10b,%dil
	xor	%r10b,%r8b
	xor	%r10b,%r9b
	xor	%r10b,%r10b
	xor	%r10b,%r11b
	xor	%r10b,%r12b
	xor	%r10b,%r13b
	xor	%r10b,%r14b
	xor	%r10b,%r15b
	nop
	xor	%r11b,%al
	xor	%r11b,%cl
	xor	%r11b,%dl
	xor	%r11b,%bl
	xor	%r11b,%spl
	xor	%r11b,%bpl
	xor	%r11b,%sil
	xor	%r11b,%dil
	xor	%r11b,%r8b
	xor	%r11b,%r9b
	xor	%r11b,%r10b
	xor	%r11b,%r11b
	xor	%r11b,%r12b
	xor	%r11b,%r13b
	xor	%r11b,%r14b
	xor	%r11b,%r15b
	nop
	xor	%r12b,%al
	xor	%r12b,%cl
	xor	%r12b,%dl
	xor	%r12b,%bl
	xor	%r12b,%spl
	xor	%r12b,%bpl
	xor	%r12b,%sil
	xor	%r12b,%dil
	xor	%r12b,%r8b
	xor	%r12b,%r9b
	xor	%r12b,%r10b
	xor	%r12b,%r11b
	xor	%r12b,%r12b
	xor	%r12b,%r13b
	xor	%r12b,%r14b
	xor	%r12b,%r15b
	nop
	xor	%r13b,%al
	xor	%r13b,%cl
	xor	%r13b,%dl
	xor	%r13b,%bl
	xor	%r13b,%spl
	xor	%r13b,%bpl
	xor	%r13b,%sil
	xor	%r13b,%dil
	xor	%r13b,%r8b
	xor	%r13b,%r9b
	xor	%r13b,%r10b
	xor	%r13b,%r11b
	xor	%r13b,%r12b
	xor	%r13b,%r13b
	xor	%r13b,%r14b
	xor	%r13b,%r15b
	nop
	xor	%r14b,%al
	xor	%r14b,%cl
	xor	%r14b,%dl
	xor	%r14b,%bl
	xor	%r14b,%spl
	xor	%r14b,%bpl
	xor	%r14b,%sil
	xor	%r14b,%dil
	xor	%r14b,%r8b
	xor	%r14b,%r9b
	xor	%r14b,%r10b
	xor	%r14b,%r11b
	xor	%r14b,%r12b
	xor	%r14b,%r13b
	xor	%r14b,%r14b
	xor	%r14b,%r15b
	nop
	xor	%r15b,%al
	xor	%r15b,%cl
	xor	%r15b,%dl
	xor	%r15b,%bl
	xor	%r15b,%spl
	xor	%r15b,%bpl
	xor	%r15b,%sil
	xor	%r15b,%dil
	xor	%r15b,%r8b
	xor	%r15b,%r9b
	xor	%r15b,%r10b
	xor	%r15b,%r11b
	xor	%r15b,%r12b
	xor	%r15b,%r13b
	xor	%r15b,%r14b
	xor	%r15b,%r15b
        nop
        nop
        // reg16 += reg16
	xor	%ax,%ax
	xor	%ax,%cx
	xor	%ax,%dx
	xor	%ax,%bx
	xor	%ax,%sp
	xor	%ax,%bp
	xor	%ax,%si
	xor	%ax,%di
	xor	%ax,%r8w
	xor	%ax,%r9w
	xor	%ax,%r10w
	xor	%ax,%r11w
	xor	%ax,%r12w
	xor	%ax,%r13w
	xor	%ax,%r14w
	xor	%ax,%r15w
	nop
	xor	%cx,%ax
	xor	%cx,%cx
	xor	%cx,%dx
	xor	%cx,%bx
	xor	%cx,%sp
	xor	%cx,%bp
	xor	%cx,%si
	xor	%cx,%di
	xor	%cx,%r8w
	xor	%cx,%r9w
	xor	%cx,%r10w
	xor	%cx,%r11w
	xor	%cx,%r12w
	xor	%cx,%r13w
	xor	%cx,%r14w
	xor	%cx,%r15w
	nop
	xor	%dx,%ax
	xor	%dx,%cx
	xor	%dx,%dx
	xor	%dx,%bx
	xor	%dx,%sp
	xor	%dx,%bp
	xor	%dx,%si
	xor	%dx,%di
	xor	%dx,%r8w
	xor	%dx,%r9w
	xor	%dx,%r10w
	xor	%dx,%r11w
	xor	%dx,%r12w
	xor	%dx,%r13w
	xor	%dx,%r14w
	xor	%dx,%r15w
	nop
	xor	%bx,%ax
	xor	%bx,%cx
	xor	%bx,%dx
	xor	%bx,%bx
	xor	%bx,%sp
	xor	%bx,%bp
	xor	%bx,%si
	xor	%bx,%di
	xor	%bx,%r8w
	xor	%bx,%r9w
	xor	%bx,%r10w
	xor	%bx,%r11w
	xor	%bx,%r12w
	xor	%bx,%r13w
	xor	%bx,%r14w
	xor	%bx,%r15w
	nop
	xor	%sp,%ax
	xor	%sp,%cx
	xor	%sp,%dx
	xor	%sp,%bx
	xor	%sp,%sp
	xor	%sp,%bp
	xor	%sp,%si
	xor	%sp,%di
	xor	%sp,%r8w
	xor	%sp,%r9w
	xor	%sp,%r10w
	xor	%sp,%r11w
	xor	%sp,%r12w
	xor	%sp,%r13w
	xor	%sp,%r14w
	xor	%sp,%r15w
	nop
	xor	%bp,%ax
	xor	%bp,%cx
	xor	%bp,%dx
	xor	%bp,%bx
	xor	%bp,%sp
	xor	%bp,%bp
	xor	%bp,%si
	xor	%bp,%di
	xor	%bp,%r8w
	xor	%bp,%r9w
	xor	%bp,%r10w
	xor	%bp,%r11w
	xor	%bp,%r12w
	xor	%bp,%r13w
	xor	%bp,%r14w
	xor	%bp,%r15w
	nop
	xor	%si,%ax
	xor	%si,%cx
	xor	%si,%dx
	xor	%si,%bx
	xor	%si,%sp
	xor	%si,%bp
	xor	%si,%si
	xor	%si,%di
	xor	%si,%r8w
	xor	%si,%r9w
	xor	%si,%r10w
	xor	%si,%r11w
	xor	%si,%r12w
	xor	%si,%r13w
	xor	%si,%r14w
	xor	%si,%r15w
	nop
	xor	%di,%ax
	xor	%di,%cx
	xor	%di,%dx
	xor	%di,%bx
	xor	%di,%sp
	xor	%di,%bp
	xor	%di,%si
	xor	%di,%di
	xor	%di,%r8w
	xor	%di,%r9w
	xor	%di,%r10w
	xor	%di,%r11w
	xor	%di,%r12w
	xor	%di,%r13w
	xor	%di,%r14w
	xor	%di,%r15w
	nop
	xor	%r8w, %ax
	xor	%r8w, %cx
	xor	%r8w, %dx
	xor	%r8w, %bx
	xor	%r8w, %sp
	xor	%r8w, %bp
	xor	%r8w, %si
	xor	%r8w, %di
	xor	%r8w, %r8w
	xor	%r8w, %r9w
	xor	%r8w, %r10w
	xor	%r8w, %r11w
	xor	%r8w, %r12w
	xor	%r8w, %r13w
	xor	%r8w, %r14w
	xor	%r8w, %r15w
	nop
	xor	%r9w, %ax
	xor	%r9w, %cx
	xor	%r9w, %dx
	xor	%r9w, %bx
	xor	%r9w, %sp
	xor	%r9w, %bp
	xor	%r9w, %si
	xor	%r9w, %di
	xor	%r9w, %r8w
	xor	%r9w, %r9w
	xor	%r9w, %r10w
	xor	%r9w, %r11w
	xor	%r9w, %r12w
	xor	%r9w, %r13w
	xor	%r9w, %r14w
	xor	%r9w, %r15w
	nop
	xor	%r10w,%ax
	xor	%r10w,%cx
	xor	%r10w,%dx
	xor	%r10w,%bx
	xor	%r10w,%sp
	xor	%r10w,%bp
	xor	%r10w,%si
	xor	%r10w,%di
	xor	%r10w,%r8w
	xor	%r10w,%r9w
	xor	%r10w,%r10w
	xor	%r10w,%r11w
	xor	%r10w,%r12w
	xor	%r10w,%r13w
	xor	%r10w,%r14w
	xor	%r10w,%r15w
	nop
	xor	%r11w,%ax
	xor	%r11w,%cx
	xor	%r11w,%dx
	xor	%r11w,%bx
	xor	%r11w,%sp
	xor	%r11w,%bp
	xor	%r11w,%si
	xor	%r11w,%di
	xor	%r11w,%r8w
	xor	%r11w,%r9w
	xor	%r11w,%r10w
	xor	%r11w,%r11w
	xor	%r11w,%r12w
	xor	%r11w,%r13w
	xor	%r11w,%r14w
	xor	%r11w,%r15w
	nop
	xor	%r12w,%ax
	xor	%r12w,%cx
	xor	%r12w,%dx
	xor	%r12w,%bx
	xor	%r12w,%sp
	xor	%r12w,%bp
	xor	%r12w,%si
	xor	%r12w,%di
	xor	%r12w,%r8w
	xor	%r12w,%r9w
	xor	%r12w,%r10w
	xor	%r12w,%r11w
	xor	%r12w,%r12w
	xor	%r12w,%r13w
	xor	%r12w,%r14w
	xor	%r12w,%r15w
	nop
	xor	%r13w,%ax
	xor	%r13w,%cx
	xor	%r13w,%dx
	xor	%r13w,%bx
	xor	%r13w,%sp
	xor	%r13w,%bp
	xor	%r13w,%si
	xor	%r13w,%di
	xor	%r13w,%r8w
	xor	%r13w,%r9w
	xor	%r13w,%r10w
	xor	%r13w,%r11w
	xor	%r13w,%r12w
	xor	%r13w,%r13w
	xor	%r13w,%r14w
	xor	%r13w,%r15w
	nop
	xor	%r14w,%ax
	xor	%r14w,%cx
	xor	%r14w,%dx
	xor	%r14w,%bx
	xor	%r14w,%sp
	xor	%r14w,%bp
	xor	%r14w,%si
	xor	%r14w,%di
	xor	%r14w,%r8w
	xor	%r14w,%r9w
	xor	%r14w,%r10w
	xor	%r14w,%r11w
	xor	%r14w,%r12w
	xor	%r14w,%r13w
	xor	%r14w,%r14w
	xor	%r14w,%r15w
	nop
	xor	%r15w,%ax
	xor	%r15w,%cx
	xor	%r15w,%dx
	xor	%r15w,%bx
	xor	%r15w,%sp
	xor	%r15w,%bp
	xor	%r15w,%si
	xor	%r15w,%di
	xor	%r15w,%r8w
	xor	%r15w,%r9w
	xor	%r15w,%r10w
	xor	%r15w,%r11w
	xor	%r15w,%r12w
	xor	%r15w,%r13w
	xor	%r15w,%r14w
	xor	%r15w,%r15w
        nop
        nop
        // reg32 += reg32
	xor	%eax,%eax
	xor	%eax,%ecx
	xor	%eax,%edx
	xor	%eax,%ebx
	xor	%eax,%esp
	xor	%eax,%ebp
	xor	%eax,%esi
	xor	%eax,%edi
	xor	%eax,%r8d
	xor	%eax,%r9d
	xor	%eax,%r10d
	xor	%eax,%r11d
	xor	%eax,%r12d
	xor	%eax,%r13d
	xor	%eax,%r14d
	xor	%eax,%r15d
	nop
	xor	%ecx,%eax
	xor	%ecx,%ecx
	xor	%ecx,%edx
	xor	%ecx,%ebx
	xor	%ecx,%esp
	xor	%ecx,%ebp
	xor	%ecx,%esi
	xor	%ecx,%edi
	xor	%ecx,%r8d
	xor	%ecx,%r9d
	xor	%ecx,%r10d
	xor	%ecx,%r11d
	xor	%ecx,%r12d
	xor	%ecx,%r13d
	xor	%ecx,%r14d
	xor	%ecx,%r15d
	nop
	xor	%edx,%eax
	xor	%edx,%ecx
	xor	%edx,%edx
	xor	%edx,%ebx
	xor	%edx,%esp
	xor	%edx,%ebp
	xor	%edx,%esi
	xor	%edx,%edi
	xor	%edx,%r8d
	xor	%edx,%r9d
	xor	%edx,%r10d
	xor	%edx,%r11d
	xor	%edx,%r12d
	xor	%edx,%r13d
	xor	%edx,%r14d
	xor	%edx,%r15d
	nop
	xor	%ebx,%eax
	xor	%ebx,%ecx
	xor	%ebx,%edx
	xor	%ebx,%ebx
	xor	%ebx,%esp
	xor	%ebx,%ebp
	xor	%ebx,%esi
	xor	%ebx,%edi
	xor	%ebx,%r8d
	xor	%ebx,%r9d
	xor	%ebx,%r10d
	xor	%ebx,%r11d
	xor	%ebx,%r12d
	xor	%ebx,%r13d
	xor	%ebx,%r14d
	xor	%ebx,%r15d
	nop
	xor	%esp,%eax
	xor	%esp,%ecx
	xor	%esp,%edx
	xor	%esp,%ebx
	xor	%esp,%esp
	xor	%esp,%ebp
	xor	%esp,%esi
	xor	%esp,%edi
	xor	%esp,%r8d
	xor	%esp,%r9d
	xor	%esp,%r10d
	xor	%esp,%r11d
	xor	%esp,%r12d
	xor	%esp,%r13d
	xor	%esp,%r14d
	xor	%esp,%r15d
	nop
	xor	%ebp,%eax
	xor	%ebp,%ecx
	xor	%ebp,%edx
	xor	%ebp,%ebx
	xor	%ebp,%esp
	xor	%ebp,%ebp
	xor	%ebp,%esi
	xor	%ebp,%edi
	xor	%ebp,%r8d
	xor	%ebp,%r9d
	xor	%ebp,%r10d
	xor	%ebp,%r11d
	xor	%ebp,%r12d
	xor	%ebp,%r13d
	xor	%ebp,%r14d
	xor	%ebp,%r15d
	nop
	xor	%esi,%eax
	xor	%esi,%ecx
	xor	%esi,%edx
	xor	%esi,%ebx
	xor	%esi,%esp
	xor	%esi,%ebp
	xor	%esi,%esi
	xor	%esi,%edi
	xor	%esi,%r8d
	xor	%esi,%r9d
	xor	%esi,%r10d
	xor	%esi,%r11d
	xor	%esi,%r12d
	xor	%esi,%r13d
	xor	%esi,%r14d
	xor	%esi,%r15d
	nop
	xor	%edi,%eax
	xor	%edi,%ecx
	xor	%edi,%edx
	xor	%edi,%ebx
	xor	%edi,%esp
	xor	%edi,%ebp
	xor	%edi,%esi
	xor	%edi,%edi
	xor	%edi,%r8d
	xor	%edi,%r9d
	xor	%edi,%r10d
	xor	%edi,%r11d
	xor	%edi,%r12d
	xor	%edi,%r13d
	xor	%edi,%r14d
	xor	%edi,%r15d
	nop
	xor	%r8d, %eax
	xor	%r8d, %ecx
	xor	%r8d, %edx
	xor	%r8d, %ebx
	xor	%r8d, %esp
	xor	%r8d, %ebp
	xor	%r8d, %esi
	xor	%r8d, %edi
	xor	%r8d, %r8d
	xor	%r8d, %r9d
	xor	%r8d, %r10d
	xor	%r8d, %r11d
	xor	%r8d, %r12d
	xor	%r8d, %r13d
	xor	%r8d, %r14d
	xor	%r8d, %r15d
	nop
	xor	%r9d, %eax
	xor	%r9d, %ecx
	xor	%r9d, %edx
	xor	%r9d, %ebx
	xor	%r9d, %esp
	xor	%r9d, %ebp
	xor	%r9d, %esi
	xor	%r9d, %edi
	xor	%r9d, %r8d
	xor	%r9d, %r9d
	xor	%r9d, %r10d
	xor	%r9d, %r11d
	xor	%r9d, %r12d
	xor	%r9d, %r13d
	xor	%r9d, %r14d
	xor	%r9d, %r15d
	nop
	xor	%r10d,%eax
	xor	%r10d,%ecx
	xor	%r10d,%edx
	xor	%r10d,%ebx
	xor	%r10d,%esp
	xor	%r10d,%ebp
	xor	%r10d,%esi
	xor	%r10d,%edi
	xor	%r10d,%r8d
	xor	%r10d,%r9d
	xor	%r10d,%r10d
	xor	%r10d,%r11d
	xor	%r10d,%r12d
	xor	%r10d,%r13d
	xor	%r10d,%r14d
	xor	%r10d,%r15d
	nop
	xor	%r11d,%eax
	xor	%r11d,%ecx
	xor	%r11d,%edx
	xor	%r11d,%ebx
	xor	%r11d,%esp
	xor	%r11d,%ebp
	xor	%r11d,%esi
	xor	%r11d,%edi
	xor	%r11d,%r8d
	xor	%r11d,%r9d
	xor	%r11d,%r10d
	xor	%r11d,%r11d
	xor	%r11d,%r12d
	xor	%r11d,%r13d
	xor	%r11d,%r14d
	xor	%r11d,%r15d
	nop
	xor	%r12d,%eax
	xor	%r12d,%ecx
	xor	%r12d,%edx
	xor	%r12d,%ebx
	xor	%r12d,%esp
	xor	%r12d,%ebp
	xor	%r12d,%esi
	xor	%r12d,%edi
	xor	%r12d,%r8d
	xor	%r12d,%r9d
	xor	%r12d,%r10d
	xor	%r12d,%r11d
	xor	%r12d,%r12d
	xor	%r12d,%r13d
	xor	%r12d,%r14d
	xor	%r12d,%r15d
	nop
	xor	%r13d,%eax
	xor	%r13d,%ecx
	xor	%r13d,%edx
	xor	%r13d,%ebx
	xor	%r13d,%esp
	xor	%r13d,%ebp
	xor	%r13d,%esi
	xor	%r13d,%edi
	xor	%r13d,%r8d
	xor	%r13d,%r9d
	xor	%r13d,%r10d
	xor	%r13d,%r11d
	xor	%r13d,%r12d
	xor	%r13d,%r13d
	xor	%r13d,%r14d
	xor	%r13d,%r15d
	nop
	xor	%r14d,%eax
	xor	%r14d,%ecx
	xor	%r14d,%edx
	xor	%r14d,%ebx
	xor	%r14d,%esp
	xor	%r14d,%ebp
	xor	%r14d,%esi
	xor	%r14d,%edi
	xor	%r14d,%r8d
	xor	%r14d,%r9d
	xor	%r14d,%r10d
	xor	%r14d,%r11d
	xor	%r14d,%r12d
	xor	%r14d,%r13d
	xor	%r14d,%r14d
	xor	%r14d,%r15d
	nop
	xor	%r15d,%eax
	xor	%r15d,%ecx
	xor	%r15d,%edx
	xor	%r15d,%ebx
	xor	%r15d,%esp
	xor	%r15d,%ebp
	xor	%r15d,%esi
	xor	%r15d,%edi
	xor	%r15d,%r8d
	xor	%r15d,%r9d
	xor	%r15d,%r10d
	xor	%r15d,%r11d
	xor	%r15d,%r12d
	xor	%r15d,%r13d
	xor	%r15d,%r14d
	xor	%r15d,%r15d
        nop
        nop
        // reg64 += reg64
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
	.globl	XorMemReg
	.type	XorMemReg, @function
XorMemReg:
	.cfi_startproc
        // mem8 += reg8
	xor	%al,(%rax)
	xor	%al,(%rcx)
	xor	%al,(%rdx)
	xor	%al,(%rbx)
	xor	%al,(%rsp)
	xor	%al,(%rbp)
	xor	%al,(%rsi)
	xor	%al,(%rdi)
	xor	%al,(%r8)
	xor	%al,(%r9)
	xor	%al,(%r10)
	xor	%al,(%r11)
	xor	%al,(%r12)
	xor	%al,(%r13)
	xor	%al,(%r14)
	xor	%al,(%r15)
	nop
	xor	%cl,(%rax)
	xor	%cl,(%rcx)
	xor	%cl,(%rdx)
	xor	%cl,(%rbx)
	xor	%cl,(%rsp)
	xor	%cl,(%rbp)
	xor	%cl,(%rsi)
	xor	%cl,(%rdi)
	xor	%cl,(%r8)
	xor	%cl,(%r9)
	xor	%cl,(%r10)
	xor	%cl,(%r11)
	xor	%cl,(%r12)
	xor	%cl,(%r13)
	xor	%cl,(%r14)
	xor	%cl,(%r15)
	nop
	xor	%dl,(%rax)
	xor	%dl,(%rcx)
	xor	%dl,(%rdx)
	xor	%dl,(%rbx)
	xor	%dl,(%rsp)
	xor	%dl,(%rbp)
	xor	%dl,(%rsi)
	xor	%dl,(%rdi)
	xor	%dl,(%r8)
	xor	%dl,(%r9)
	xor	%dl,(%r10)
	xor	%dl,(%r11)
	xor	%dl,(%r12)
	xor	%dl,(%r13)
	xor	%dl,(%r14)
	xor	%dl,(%r15)
	nop
	xor	%bl,(%rax)
	xor	%bl,(%rcx)
	xor	%bl,(%rdx)
	xor	%bl,(%rbx)
	xor	%bl,(%rsp)
	xor	%bl,(%rbp)
	xor	%bl,(%rsi)
	xor	%bl,(%rdi)
	xor	%bl,(%r8)
	xor	%bl,(%r9)
	xor	%bl,(%r10)
	xor	%bl,(%r11)
	xor	%bl,(%r12)
	xor	%bl,(%r13)
	xor	%bl,(%r14)
	xor	%bl,(%r15)
	nop
	xor	%spl,(%rax)
	xor	%spl,(%rcx)
	xor	%spl,(%rdx)
	xor	%spl,(%rbx)
	xor	%spl,(%rsp)
	xor	%spl,(%rbp)
	xor	%spl,(%rsi)
	xor	%spl,(%rdi)
	xor	%spl,(%r8)
	xor	%spl,(%r9)
	xor	%spl,(%r10)
	xor	%spl,(%r11)
	xor	%spl,(%r12)
	xor	%spl,(%r13)
	xor	%spl,(%r14)
	xor	%spl,(%r15)
	nop
	xor	%bpl,(%rax)
	xor	%bpl,(%rcx)
	xor	%bpl,(%rdx)
	xor	%bpl,(%rbx)
	xor	%bpl,(%rsp)
	xor	%bpl,(%rbp)
	xor	%bpl,(%rsi)
	xor	%bpl,(%rdi)
	xor	%bpl,(%r8)
	xor	%bpl,(%r9)
	xor	%bpl,(%r10)
	xor	%bpl,(%r11)
	xor	%bpl,(%r12)
	xor	%bpl,(%r13)
	xor	%bpl,(%r14)
	xor	%bpl,(%r15)
	nop
	xor	%sil,(%rax)
	xor	%sil,(%rcx)
	xor	%sil,(%rdx)
	xor	%sil,(%rbx)
	xor	%sil,(%rsp)
	xor	%sil,(%rbp)
	xor	%sil,(%rsi)
	xor	%sil,(%rdi)
	xor	%sil,(%r8)
	xor	%sil,(%r9)
	xor	%sil,(%r10)
	xor	%sil,(%r11)
	xor	%sil,(%r12)
	xor	%sil,(%r13)
	xor	%sil,(%r14)
	xor	%sil,(%r15)
	nop
	xor	%dil,(%rax)
	xor	%dil,(%rcx)
	xor	%dil,(%rdx)
	xor	%dil,(%rbx)
	xor	%dil,(%rsp)
	xor	%dil,(%rbp)
	xor	%dil,(%rsi)
	xor	%dil,(%rdi)
	xor	%dil,(%r8)
	xor	%dil,(%r9)
	xor	%dil,(%r10)
	xor	%dil,(%r11)
	xor	%dil,(%r12)
	xor	%dil,(%r13)
	xor	%dil,(%r14)
	xor	%dil,(%r15)
	nop
	xor	%r8b, (%rax)
	xor	%r8b, (%rcx)
	xor	%r8b, (%rdx)
	xor	%r8b, (%rbx)
	xor	%r8b, (%rsp)
	xor	%r8b, (%rbp)
	xor	%r8b, (%rsi)
	xor	%r8b, (%rdi)
	xor	%r8b, (%r8)
	xor	%r8b, (%r9)
	xor	%r8b, (%r10)
	xor	%r8b, (%r11)
	xor	%r8b, (%r12)
	xor	%r8b, (%r13)
	xor	%r8b, (%r14)
	xor	%r8b, (%r15)
	nop
	xor	%r9b, (%rax)
	xor	%r9b, (%rcx)
	xor	%r9b, (%rdx)
	xor	%r9b, (%rbx)
	xor	%r9b, (%rsp)
	xor	%r9b, (%rbp)
	xor	%r9b, (%rsi)
	xor	%r9b, (%rdi)
	xor	%r9b, (%r8)
	xor	%r9b, (%r9)
	xor	%r9b, (%r10)
	xor	%r9b, (%r11)
	xor	%r9b, (%r12)
	xor	%r9b, (%r13)
	xor	%r9b, (%r14)
	xor	%r9b, (%r15)
	nop
	xor	%r10b,(%rax)
	xor	%r10b,(%rcx)
	xor	%r10b,(%rdx)
	xor	%r10b,(%rbx)
	xor	%r10b,(%rsp)
	xor	%r10b,(%rbp)
	xor	%r10b,(%rsi)
	xor	%r10b,(%rdi)
	xor	%r10b,(%r8)
	xor	%r10b,(%r9)
	xor	%r10b,(%r10)
	xor	%r10b,(%r11)
	xor	%r10b,(%r12)
	xor	%r10b,(%r13)
	xor	%r10b,(%r14)
	xor	%r10b,(%r15)
	nop
	xor	%r11b,(%rax)
	xor	%r11b,(%rcx)
	xor	%r11b,(%rdx)
	xor	%r11b,(%rbx)
	xor	%r11b,(%rsp)
	xor	%r11b,(%rbp)
	xor	%r11b,(%rsi)
	xor	%r11b,(%rdi)
	xor	%r11b,(%r8)
	xor	%r11b,(%r9)
	xor	%r11b,(%r10)
	xor	%r11b,(%r11)
	xor	%r11b,(%r12)
	xor	%r11b,(%r13)
	xor	%r11b,(%r14)
	xor	%r11b,(%r15)
	nop
	xor	%r12b,(%rax)
	xor	%r12b,(%rcx)
	xor	%r12b,(%rdx)
	xor	%r12b,(%rbx)
	xor	%r12b,(%rsp)
	xor	%r12b,(%rbp)
	xor	%r12b,(%rsi)
	xor	%r12b,(%rdi)
	xor	%r12b,(%r8)
	xor	%r12b,(%r9)
	xor	%r12b,(%r10)
	xor	%r12b,(%r11)
	xor	%r12b,(%r12)
	xor	%r12b,(%r13)
	xor	%r12b,(%r14)
	xor	%r12b,(%r15)
	nop
	xor	%r13b,(%rax)
	xor	%r13b,(%rcx)
	xor	%r13b,(%rdx)
	xor	%r13b,(%rbx)
	xor	%r13b,(%rsp)
	xor	%r13b,(%rbp)
	xor	%r13b,(%rsi)
	xor	%r13b,(%rdi)
	xor	%r13b,(%r8)
	xor	%r13b,(%r9)
	xor	%r13b,(%r10)
	xor	%r13b,(%r11)
	xor	%r13b,(%r12)
	xor	%r13b,(%r13)
	xor	%r13b,(%r14)
	xor	%r13b,(%r15)
	nop
	xor	%r14b,(%rax)
	xor	%r14b,(%rcx)
	xor	%r14b,(%rdx)
	xor	%r14b,(%rbx)
	xor	%r14b,(%rsp)
	xor	%r14b,(%rbp)
	xor	%r14b,(%rsi)
	xor	%r14b,(%rdi)
	xor	%r14b,(%r8)
	xor	%r14b,(%r9)
	xor	%r14b,(%r10)
	xor	%r14b,(%r11)
	xor	%r14b,(%r12)
	xor	%r14b,(%r13)
	xor	%r14b,(%r14)
	xor	%r14b,(%r15)
	nop
	xor	%r15b,(%rax)
	xor	%r15b,(%rcx)
	xor	%r15b,(%rdx)
	xor	%r15b,(%rbx)
	xor	%r15b,(%rsp)
	xor	%r15b,(%rbp)
	xor	%r15b,(%rsi)
	xor	%r15b,(%rdi)
	xor	%r15b,(%r8)
	xor	%r15b,(%r9)
	xor	%r15b,(%r10)
	xor	%r15b,(%r11)
	xor	%r15b,(%r12)
	xor	%r15b,(%r13)
	xor	%r15b,(%r14)
	xor	%r15b,(%r15)
        nop
        nop
        // mem16 += reg16
	xor	%ax,(%rax)
	xor	%ax,(%rcx)
	xor	%ax,(%rdx)
	xor	%ax,(%rbx)
	xor	%ax,(%rsp)
	xor	%ax,(%rbp)
	xor	%ax,(%rsi)
	xor	%ax,(%rdi)
	xor	%ax,(%r8)
	xor	%ax,(%r9)
	xor	%ax,(%r10)
	xor	%ax,(%r11)
	xor	%ax,(%r12)
	xor	%ax,(%r13)
	xor	%ax,(%r14)
	xor	%ax,(%r15)
	nop
	xor	%cx,(%rax)
	xor	%cx,(%rcx)
	xor	%cx,(%rdx)
	xor	%cx,(%rbx)
	xor	%cx,(%rsp)
	xor	%cx,(%rbp)
	xor	%cx,(%rsi)
	xor	%cx,(%rdi)
	xor	%cx,(%r8)
	xor	%cx,(%r9)
	xor	%cx,(%r10)
	xor	%cx,(%r11)
	xor	%cx,(%r12)
	xor	%cx,(%r13)
	xor	%cx,(%r14)
	xor	%cx,(%r15)
	nop
	xor	%dx,(%rax)
	xor	%dx,(%rcx)
	xor	%dx,(%rdx)
	xor	%dx,(%rbx)
	xor	%dx,(%rsp)
	xor	%dx,(%rbp)
	xor	%dx,(%rsi)
	xor	%dx,(%rdi)
	xor	%dx,(%r8)
	xor	%dx,(%r9)
	xor	%dx,(%r10)
	xor	%dx,(%r11)
	xor	%dx,(%r12)
	xor	%dx,(%r13)
	xor	%dx,(%r14)
	xor	%dx,(%r15)
	nop
	xor	%bx,(%rax)
	xor	%bx,(%rcx)
	xor	%bx,(%rdx)
	xor	%bx,(%rbx)
	xor	%bx,(%rsp)
	xor	%bx,(%rbp)
	xor	%bx,(%rsi)
	xor	%bx,(%rdi)
	xor	%bx,(%r8)
	xor	%bx,(%r9)
	xor	%bx,(%r10)
	xor	%bx,(%r11)
	xor	%bx,(%r12)
	xor	%bx,(%r13)
	xor	%bx,(%r14)
	xor	%bx,(%r15)
	nop
	xor	%sp,(%rax)
	xor	%sp,(%rcx)
	xor	%sp,(%rdx)
	xor	%sp,(%rbx)
	xor	%sp,(%rsp)
	xor	%sp,(%rbp)
	xor	%sp,(%rsi)
	xor	%sp,(%rdi)
	xor	%sp,(%r8)
	xor	%sp,(%r9)
	xor	%sp,(%r10)
	xor	%sp,(%r11)
	xor	%sp,(%r12)
	xor	%sp,(%r13)
	xor	%sp,(%r14)
	xor	%sp,(%r15)
	nop
	xor	%bp,(%rax)
	xor	%bp,(%rcx)
	xor	%bp,(%rdx)
	xor	%bp,(%rbx)
	xor	%bp,(%rsp)
	xor	%bp,(%rbp)
	xor	%bp,(%rsi)
	xor	%bp,(%rdi)
	xor	%bp,(%r8)
	xor	%bp,(%r9)
	xor	%bp,(%r10)
	xor	%bp,(%r11)
	xor	%bp,(%r12)
	xor	%bp,(%r13)
	xor	%bp,(%r14)
	xor	%bp,(%r15)
	nop
	xor	%si,(%rax)
	xor	%si,(%rcx)
	xor	%si,(%rdx)
	xor	%si,(%rbx)
	xor	%si,(%rsp)
	xor	%si,(%rbp)
	xor	%si,(%rsi)
	xor	%si,(%rdi)
	xor	%si,(%r8)
	xor	%si,(%r9)
	xor	%si,(%r10)
	xor	%si,(%r11)
	xor	%si,(%r12)
	xor	%si,(%r13)
	xor	%si,(%r14)
	xor	%si,(%r15)
	nop
	xor	%di,(%rax)
	xor	%di,(%rcx)
	xor	%di,(%rdx)
	xor	%di,(%rbx)
	xor	%di,(%rsp)
	xor	%di,(%rbp)
	xor	%di,(%rsi)
	xor	%di,(%rdi)
	xor	%di,(%r8)
	xor	%di,(%r9)
	xor	%di,(%r10)
	xor	%di,(%r11)
	xor	%di,(%r12)
	xor	%di,(%r13)
	xor	%di,(%r14)
	xor	%di,(%r15)
	nop
	xor	%r8w, (%rax)
	xor	%r8w, (%rcx)
	xor	%r8w, (%rdx)
	xor	%r8w, (%rbx)
	xor	%r8w, (%rsp)
	xor	%r8w, (%rbp)
	xor	%r8w, (%rsi)
	xor	%r8w, (%rdi)
	xor	%r8w, (%r8)
	xor	%r8w, (%r9)
	xor	%r8w, (%r10)
	xor	%r8w, (%r11)
	xor	%r8w, (%r12)
	xor	%r8w, (%r13)
	xor	%r8w, (%r14)
	xor	%r8w, (%r15)
	nop
	xor	%r9w, (%rax)
	xor	%r9w, (%rcx)
	xor	%r9w, (%rdx)
	xor	%r9w, (%rbx)
	xor	%r9w, (%rsp)
	xor	%r9w, (%rbp)
	xor	%r9w, (%rsi)
	xor	%r9w, (%rdi)
	xor	%r9w, (%r8)
	xor	%r9w, (%r9)
	xor	%r9w, (%r10)
	xor	%r9w, (%r11)
	xor	%r9w, (%r12)
	xor	%r9w, (%r13)
	xor	%r9w, (%r14)
	xor	%r9w, (%r15)
	nop
	xor	%r10w,(%rax)
	xor	%r10w,(%rcx)
	xor	%r10w,(%rdx)
	xor	%r10w,(%rbx)
	xor	%r10w,(%rsp)
	xor	%r10w,(%rbp)
	xor	%r10w,(%rsi)
	xor	%r10w,(%rdi)
	xor	%r10w,(%r8)
	xor	%r10w,(%r9)
	xor	%r10w,(%r10)
	xor	%r10w,(%r11)
	xor	%r10w,(%r12)
	xor	%r10w,(%r13)
	xor	%r10w,(%r14)
	xor	%r10w,(%r15)
	nop
	xor	%r11w,(%rax)
	xor	%r11w,(%rcx)
	xor	%r11w,(%rdx)
	xor	%r11w,(%rbx)
	xor	%r11w,(%rsp)
	xor	%r11w,(%rbp)
	xor	%r11w,(%rsi)
	xor	%r11w,(%rdi)
	xor	%r11w,(%r8)
	xor	%r11w,(%r9)
	xor	%r11w,(%r10)
	xor	%r11w,(%r11)
	xor	%r11w,(%r12)
	xor	%r11w,(%r13)
	xor	%r11w,(%r14)
	xor	%r11w,(%r15)
	nop
	xor	%r12w,(%rax)
	xor	%r12w,(%rcx)
	xor	%r12w,(%rdx)
	xor	%r12w,(%rbx)
	xor	%r12w,(%rsp)
	xor	%r12w,(%rbp)
	xor	%r12w,(%rsi)
	xor	%r12w,(%rdi)
	xor	%r12w,(%r8)
	xor	%r12w,(%r9)
	xor	%r12w,(%r10)
	xor	%r12w,(%r11)
	xor	%r12w,(%r12)
	xor	%r12w,(%r13)
	xor	%r12w,(%r14)
	xor	%r12w,(%r15)
	nop
	xor	%r13w,(%rax)
	xor	%r13w,(%rcx)
	xor	%r13w,(%rdx)
	xor	%r13w,(%rbx)
	xor	%r13w,(%rsp)
	xor	%r13w,(%rbp)
	xor	%r13w,(%rsi)
	xor	%r13w,(%rdi)
	xor	%r13w,(%r8)
	xor	%r13w,(%r9)
	xor	%r13w,(%r10)
	xor	%r13w,(%r11)
	xor	%r13w,(%r12)
	xor	%r13w,(%r13)
	xor	%r13w,(%r14)
	xor	%r13w,(%r15)
	nop
	xor	%r14w,(%rax)
	xor	%r14w,(%rcx)
	xor	%r14w,(%rdx)
	xor	%r14w,(%rbx)
	xor	%r14w,(%rsp)
	xor	%r14w,(%rbp)
	xor	%r14w,(%rsi)
	xor	%r14w,(%rdi)
	xor	%r14w,(%r8)
	xor	%r14w,(%r9)
	xor	%r14w,(%r10)
	xor	%r14w,(%r11)
	xor	%r14w,(%r12)
	xor	%r14w,(%r13)
	xor	%r14w,(%r14)
	xor	%r14w,(%r15)
	nop
	xor	%r15w,(%rax)
	xor	%r15w,(%rcx)
	xor	%r15w,(%rdx)
	xor	%r15w,(%rbx)
	xor	%r15w,(%rsp)
	xor	%r15w,(%rbp)
	xor	%r15w,(%rsi)
	xor	%r15w,(%rdi)
	xor	%r15w,(%r8)
	xor	%r15w,(%r9)
	xor	%r15w,(%r10)
	xor	%r15w,(%r11)
	xor	%r15w,(%r12)
	xor	%r15w,(%r13)
	xor	%r15w,(%r14)
	xor	%r15w,(%r15)
        nop
        nop
        // mem32 += reg32
	xor	%eax,(%rax)
	xor	%eax,(%rcx)
	xor	%eax,(%rdx)
	xor	%eax,(%rbx)
	xor	%eax,(%rsp)
	xor	%eax,(%rbp)
	xor	%eax,(%rsi)
	xor	%eax,(%rdi)
	xor	%eax,(%r8)
	xor	%eax,(%r9)
	xor	%eax,(%r10)
	xor	%eax,(%r11)
	xor	%eax,(%r12)
	xor	%eax,(%r13)
	xor	%eax,(%r14)
	xor	%eax,(%r15)
	nop
	xor	%ecx,(%rax)
	xor	%ecx,(%rcx)
	xor	%ecx,(%rdx)
	xor	%ecx,(%rbx)
	xor	%ecx,(%rsp)
	xor	%ecx,(%rbp)
	xor	%ecx,(%rsi)
	xor	%ecx,(%rdi)
	xor	%ecx,(%r8)
	xor	%ecx,(%r9)
	xor	%ecx,(%r10)
	xor	%ecx,(%r11)
	xor	%ecx,(%r12)
	xor	%ecx,(%r13)
	xor	%ecx,(%r14)
	xor	%ecx,(%r15)
	nop
	xor	%edx,(%rax)
	xor	%edx,(%rcx)
	xor	%edx,(%rdx)
	xor	%edx,(%rbx)
	xor	%edx,(%rsp)
	xor	%edx,(%rbp)
	xor	%edx,(%rsi)
	xor	%edx,(%rdi)
	xor	%edx,(%r8)
	xor	%edx,(%r9)
	xor	%edx,(%r10)
	xor	%edx,(%r11)
	xor	%edx,(%r12)
	xor	%edx,(%r13)
	xor	%edx,(%r14)
	xor	%edx,(%r15)
	nop
	xor	%ebx,(%rax)
	xor	%ebx,(%rcx)
	xor	%ebx,(%rdx)
	xor	%ebx,(%rbx)
	xor	%ebx,(%rsp)
	xor	%ebx,(%rbp)
	xor	%ebx,(%rsi)
	xor	%ebx,(%rdi)
	xor	%ebx,(%r8)
	xor	%ebx,(%r9)
	xor	%ebx,(%r10)
	xor	%ebx,(%r11)
	xor	%ebx,(%r12)
	xor	%ebx,(%r13)
	xor	%ebx,(%r14)
	xor	%ebx,(%r15)
	nop
	xor	%esp,(%rax)
	xor	%esp,(%rcx)
	xor	%esp,(%rdx)
	xor	%esp,(%rbx)
	xor	%esp,(%rsp)
	xor	%esp,(%rbp)
	xor	%esp,(%rsi)
	xor	%esp,(%rdi)
	xor	%esp,(%r8)
	xor	%esp,(%r9)
	xor	%esp,(%r10)
	xor	%esp,(%r11)
	xor	%esp,(%r12)
	xor	%esp,(%r13)
	xor	%esp,(%r14)
	xor	%esp,(%r15)
	nop
	xor	%ebp,(%rax)
	xor	%ebp,(%rcx)
	xor	%ebp,(%rdx)
	xor	%ebp,(%rbx)
	xor	%ebp,(%rsp)
	xor	%ebp,(%rbp)
	xor	%ebp,(%rsi)
	xor	%ebp,(%rdi)
	xor	%ebp,(%r8)
	xor	%ebp,(%r9)
	xor	%ebp,(%r10)
	xor	%ebp,(%r11)
	xor	%ebp,(%r12)
	xor	%ebp,(%r13)
	xor	%ebp,(%r14)
	xor	%ebp,(%r15)
	nop
	xor	%esi,(%rax)
	xor	%esi,(%rcx)
	xor	%esi,(%rdx)
	xor	%esi,(%rbx)
	xor	%esi,(%rsp)
	xor	%esi,(%rbp)
	xor	%esi,(%rsi)
	xor	%esi,(%rdi)
	xor	%esi,(%r8)
	xor	%esi,(%r9)
	xor	%esi,(%r10)
	xor	%esi,(%r11)
	xor	%esi,(%r12)
	xor	%esi,(%r13)
	xor	%esi,(%r14)
	xor	%esi,(%r15)
	nop
	xor	%edi,(%rax)
	xor	%edi,(%rcx)
	xor	%edi,(%rdx)
	xor	%edi,(%rbx)
	xor	%edi,(%rsp)
	xor	%edi,(%rbp)
	xor	%edi,(%rsi)
	xor	%edi,(%rdi)
	xor	%edi,(%r8)
	xor	%edi,(%r9)
	xor	%edi,(%r10)
	xor	%edi,(%r11)
	xor	%edi,(%r12)
	xor	%edi,(%r13)
	xor	%edi,(%r14)
	xor	%edi,(%r15)
	nop
	xor	%r8d, (%rax)
	xor	%r8d, (%rcx)
	xor	%r8d, (%rdx)
	xor	%r8d, (%rbx)
	xor	%r8d, (%rsp)
	xor	%r8d, (%rbp)
	xor	%r8d, (%rsi)
	xor	%r8d, (%rdi)
	xor	%r8d, (%r8)
	xor	%r8d, (%r9)
	xor	%r8d, (%r10)
	xor	%r8d, (%r11)
	xor	%r8d, (%r12)
	xor	%r8d, (%r13)
	xor	%r8d, (%r14)
	xor	%r8d, (%r15)
	nop
	xor	%r9d, (%rax)
	xor	%r9d, (%rcx)
	xor	%r9d, (%rdx)
	xor	%r9d, (%rbx)
	xor	%r9d, (%rsp)
	xor	%r9d, (%rbp)
	xor	%r9d, (%rsi)
	xor	%r9d, (%rdi)
	xor	%r9d, (%r8)
	xor	%r9d, (%r9)
	xor	%r9d, (%r10)
	xor	%r9d, (%r11)
	xor	%r9d, (%r12)
	xor	%r9d, (%r13)
	xor	%r9d, (%r14)
	xor	%r9d, (%r15)
	nop
	xor	%r10d,(%rax)
	xor	%r10d,(%rcx)
	xor	%r10d,(%rdx)
	xor	%r10d,(%rbx)
	xor	%r10d,(%rsp)
	xor	%r10d,(%rbp)
	xor	%r10d,(%rsi)
	xor	%r10d,(%rdi)
	xor	%r10d,(%r8)
	xor	%r10d,(%r9)
	xor	%r10d,(%r10)
	xor	%r10d,(%r11)
	xor	%r10d,(%r12)
	xor	%r10d,(%r13)
	xor	%r10d,(%r14)
	xor	%r10d,(%r15)
	nop
	xor	%r11d,(%rax)
	xor	%r11d,(%rcx)
	xor	%r11d,(%rdx)
	xor	%r11d,(%rbx)
	xor	%r11d,(%rsp)
	xor	%r11d,(%rbp)
	xor	%r11d,(%rsi)
	xor	%r11d,(%rdi)
	xor	%r11d,(%r8)
	xor	%r11d,(%r9)
	xor	%r11d,(%r10)
	xor	%r11d,(%r11)
	xor	%r11d,(%r12)
	xor	%r11d,(%r13)
	xor	%r11d,(%r14)
	xor	%r11d,(%r15)
	nop
	xor	%r12d,(%rax)
	xor	%r12d,(%rcx)
	xor	%r12d,(%rdx)
	xor	%r12d,(%rbx)
	xor	%r12d,(%rsp)
	xor	%r12d,(%rbp)
	xor	%r12d,(%rsi)
	xor	%r12d,(%rdi)
	xor	%r12d,(%r8)
	xor	%r12d,(%r9)
	xor	%r12d,(%r10)
	xor	%r12d,(%r11)
	xor	%r12d,(%r12)
	xor	%r12d,(%r13)
	xor	%r12d,(%r14)
	xor	%r12d,(%r15)
	nop
	xor	%r13d,(%rax)
	xor	%r13d,(%rcx)
	xor	%r13d,(%rdx)
	xor	%r13d,(%rbx)
	xor	%r13d,(%rsp)
	xor	%r13d,(%rbp)
	xor	%r13d,(%rsi)
	xor	%r13d,(%rdi)
	xor	%r13d,(%r8)
	xor	%r13d,(%r9)
	xor	%r13d,(%r10)
	xor	%r13d,(%r11)
	xor	%r13d,(%r12)
	xor	%r13d,(%r13)
	xor	%r13d,(%r14)
	xor	%r13d,(%r15)
	nop
	xor	%r14d,(%rax)
	xor	%r14d,(%rcx)
	xor	%r14d,(%rdx)
	xor	%r14d,(%rbx)
	xor	%r14d,(%rsp)
	xor	%r14d,(%rbp)
	xor	%r14d,(%rsi)
	xor	%r14d,(%rdi)
	xor	%r14d,(%r8)
	xor	%r14d,(%r9)
	xor	%r14d,(%r10)
	xor	%r14d,(%r11)
	xor	%r14d,(%r12)
	xor	%r14d,(%r13)
	xor	%r14d,(%r14)
	xor	%r14d,(%r15)
	nop
	xor	%r15d,(%rax)
	xor	%r15d,(%rcx)
	xor	%r15d,(%rdx)
	xor	%r15d,(%rbx)
	xor	%r15d,(%rsp)
	xor	%r15d,(%rbp)
	xor	%r15d,(%rsi)
	xor	%r15d,(%rdi)
	xor	%r15d,(%r8)
	xor	%r15d,(%r9)
	xor	%r15d,(%r10)
	xor	%r15d,(%r11)
	xor	%r15d,(%r12)
	xor	%r15d,(%r13)
	xor	%r15d,(%r14)
	xor	%r15d,(%r15)
        nop
        nop
        // mem64 += reg64
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
	.globl	XorMem8Reg
	.type	XorMem8Reg, @function
XorMem8Reg:
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
	.globl	XorMem32Reg
	.type	XorMem32Reg, @function
XorMem32Reg:
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
        
        
	.p2align 4,,15
	.globl	XorRegMem
	.type	XorRegMem, @function
XorRegMem:
	.cfi_startproc
	xor	(%rax),%rax
	xor	(%rax),%rcx
	xor	(%rax),%rdx
	xor	(%rax),%rbx
	xor	(%rax),%rsp
	xor	(%rax),%rbp
	xor	(%rax),%rsi
	xor	(%rax),%rdi
	xor	(%rax),%r8
	xor	(%rax),%r9
	xor	(%rax),%r10
	xor	(%rax),%r11
	xor	(%rax),%r12
	xor	(%rax),%r13
	xor	(%rax),%r14
	xor	(%rax),%r15
	nop
	xor	(%rcx),%rax
	xor	(%rcx),%rcx
	xor	(%rcx),%rdx
	xor	(%rcx),%rbx
	xor	(%rcx),%rsp
	xor	(%rcx),%rbp
	xor	(%rcx),%rsi
	xor	(%rcx),%rdi
	xor	(%rcx),%r8
	xor	(%rcx),%r9
	xor	(%rcx),%r10
	xor	(%rcx),%r11
	xor	(%rcx),%r12
	xor	(%rcx),%r13
	xor	(%rcx),%r14
	xor	(%rcx),%r15
	nop
	xor	(%rdx),%rax
	xor	(%rdx),%rcx
	xor	(%rdx),%rdx
	xor	(%rdx),%rbx
	xor	(%rdx),%rsp
	xor	(%rdx),%rbp
	xor	(%rdx),%rsi
	xor	(%rdx),%rdi
	xor	(%rdx),%r8
	xor	(%rdx),%r9
	xor	(%rdx),%r10
	xor	(%rdx),%r11
	xor	(%rdx),%r12
	xor	(%rdx),%r13
	xor	(%rdx),%r14
	xor	(%rdx),%r15
	nop
	xor	(%rbx),%rax
	xor	(%rbx),%rcx
	xor	(%rbx),%rdx
	xor	(%rbx),%rbx
	xor	(%rbx),%rsp
	xor	(%rbx),%rbp
	xor	(%rbx),%rsi
	xor	(%rbx),%rdi
	xor	(%rbx),%r8
	xor	(%rbx),%r9
	xor	(%rbx),%r10
	xor	(%rbx),%r11
	xor	(%rbx),%r12
	xor	(%rbx),%r13
	xor	(%rbx),%r14
	xor	(%rbx),%r15
	nop
	xor	(%rsp),%rax
	xor	(%rsp),%rcx
	xor	(%rsp),%rdx
	xor	(%rsp),%rbx
	xor	(%rsp),%rsp
	xor	(%rsp),%rbp
	xor	(%rsp),%rsi
	xor	(%rsp),%rdi
	xor	(%rsp),%r8
	xor	(%rsp),%r9
	xor	(%rsp),%r10
	xor	(%rsp),%r11
	xor	(%rsp),%r12
	xor	(%rsp),%r13
	xor	(%rsp),%r14
	xor	(%rsp),%r15
	nop
	xor	(%rbp),%rax
	xor	(%rbp),%rcx
	xor	(%rbp),%rdx
	xor	(%rbp),%rbx
	xor	(%rbp),%rsp
	xor	(%rbp),%rbp
	xor	(%rbp),%rsi
	xor	(%rbp),%rdi
	xor	(%rbp),%r8
	xor	(%rbp),%r9
	xor	(%rbp),%r10
	xor	(%rbp),%r11
	xor	(%rbp),%r12
	xor	(%rbp),%r13
	xor	(%rbp),%r14
	xor	(%rbp),%r15
	nop
	xor	(%rsi),%rax
	xor	(%rsi),%rcx
	xor	(%rsi),%rdx
	xor	(%rsi),%rbx
	xor	(%rsi),%rsp
	xor	(%rsi),%rbp
	xor	(%rsi),%rsi
	xor	(%rsi),%rdi
	xor	(%rsi),%r8
	xor	(%rsi),%r9
	xor	(%rsi),%r10
	xor	(%rsi),%r11
	xor	(%rsi),%r12
	xor	(%rsi),%r13
	xor	(%rsi),%r14
	xor	(%rsi),%r15
	nop
	xor	(%rdi),%rax
	xor	(%rdi),%rcx
	xor	(%rdi),%rdx
	xor	(%rdi),%rbx
	xor	(%rdi),%rsp
	xor	(%rdi),%rbp
	xor	(%rdi),%rsi
	xor	(%rdi),%rdi
	xor	(%rdi),%r8
	xor	(%rdi),%r9
	xor	(%rdi),%r10
	xor	(%rdi),%r11
	xor	(%rdi),%r12
	xor	(%rdi),%r13
	xor	(%rdi),%r14
	xor	(%rdi),%r15
	nop
	xor	(%r8), %rax
	xor	(%r8), %rcx
	xor	(%r8), %rdx
	xor	(%r8), %rbx
	xor	(%r8), %rsp
	xor	(%r8), %rbp
	xor	(%r8), %rsi
	xor	(%r8), %rdi
	xor	(%r8), %r8
	xor	(%r8), %r9
	xor	(%r8), %r10
	xor	(%r8), %r11
	xor	(%r8), %r12
	xor	(%r8), %r13
	xor	(%r8), %r14
	xor	(%r8), %r15
	nop
	xor	(%r9), %rax
	xor	(%r9), %rcx
	xor	(%r9), %rdx
	xor	(%r9), %rbx
	xor	(%r9), %rsp
	xor	(%r9), %rbp
	xor	(%r9), %rsi
	xor	(%r9), %rdi
	xor	(%r9), %r8
	xor	(%r9), %r9
	xor	(%r9), %r10
	xor	(%r9), %r11
	xor	(%r9), %r12
	xor	(%r9), %r13
	xor	(%r9), %r14
	xor	(%r9), %r15
	nop
	xor	(%r10),%rax
	xor	(%r10),%rcx
	xor	(%r10),%rdx
	xor	(%r10),%rbx
	xor	(%r10),%rsp
	xor	(%r10),%rbp
	xor	(%r10),%rsi
	xor	(%r10),%rdi
	xor	(%r10),%r8
	xor	(%r10),%r9
	xor	(%r10),%r10
	xor	(%r10),%r11
	xor	(%r10),%r12
	xor	(%r10),%r13
	xor	(%r10),%r14
	xor	(%r10),%r15
	nop
	xor	(%r11),%rax
	xor	(%r11),%rcx
	xor	(%r11),%rdx
	xor	(%r11),%rbx
	xor	(%r11),%rsp
	xor	(%r11),%rbp
	xor	(%r11),%rsi
	xor	(%r11),%rdi
	xor	(%r11),%r8
	xor	(%r11),%r9
	xor	(%r11),%r10
	xor	(%r11),%r11
	xor	(%r11),%r12
	xor	(%r11),%r13
	xor	(%r11),%r14
	xor	(%r11),%r15
	nop
	xor	(%r12),%rax
	xor	(%r12),%rcx
	xor	(%r12),%rdx
	xor	(%r12),%rbx
	xor	(%r12),%rsp
	xor	(%r12),%rbp
	xor	(%r12),%rsi
	xor	(%r12),%rdi
	xor	(%r12),%r8
	xor	(%r12),%r9
	xor	(%r12),%r10
	xor	(%r12),%r11
	xor	(%r12),%r12
	xor	(%r12),%r13
	xor	(%r12),%r14
	xor	(%r12),%r15
	nop
	xor	(%r13),%rax
	xor	(%r13),%rcx
	xor	(%r13),%rdx
	xor	(%r13),%rbx
	xor	(%r13),%rsp
	xor	(%r13),%rbp
	xor	(%r13),%rsi
	xor	(%r13),%rdi
	xor	(%r13),%r8
	xor	(%r13),%r9
	xor	(%r13),%r10
	xor	(%r13),%r11
	xor	(%r13),%r12
	xor	(%r13),%r13
	xor	(%r13),%r14
	xor	(%r13),%r15
	nop
	xor	(%r14),%rax
	xor	(%r14),%rcx
	xor	(%r14),%rdx
	xor	(%r14),%rbx
	xor	(%r14),%rsp
	xor	(%r14),%rbp
	xor	(%r14),%rsi
	xor	(%r14),%rdi
	xor	(%r14),%r8
	xor	(%r14),%r9
	xor	(%r14),%r10
	xor	(%r14),%r11
	xor	(%r14),%r12
	xor	(%r14),%r13
	xor	(%r14),%r14
	xor	(%r14),%r15
	nop
	xor	(%r15),%rax
	xor	(%r15),%rcx
	xor	(%r15),%rdx
	xor	(%r15),%rbx
	xor	(%r15),%rsp
	xor	(%r15),%rbp
	xor	(%r15),%rsi
	xor	(%r15),%rdi
	xor	(%r15),%r8
	xor	(%r15),%r9
	xor	(%r15),%r10
	xor	(%r15),%r11
	xor	(%r15),%r12
	xor	(%r15),%r13
	xor	(%r15),%r14
	xor	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XorRegMem8
	.type	XorRegMem8, @function
XorRegMem8:
	.cfi_startproc
	xor	0x7f(%rax),%rax
	xor	0x7f(%rax),%rcx
	xor	0x7f(%rax),%rdx
	xor	0x7f(%rax),%rbx
	xor	0x7f(%rax),%rsp
	xor	0x7f(%rax),%rbp
	xor	0x7f(%rax),%rsi
	xor	0x7f(%rax),%rdi
	xor	0x7f(%rax),%r8
	xor	0x7f(%rax),%r9
	xor	0x7f(%rax),%r10
	xor	0x7f(%rax),%r11
	xor	0x7f(%rax),%r12
	xor	0x7f(%rax),%r13
	xor	0x7f(%rax),%r14
	xor	0x7f(%rax),%r15
	nop
	xor	0x7f(%rcx),%rax
	xor	0x7f(%rcx),%rcx
	xor	0x7f(%rcx),%rdx
	xor	0x7f(%rcx),%rbx
	xor	0x7f(%rcx),%rsp
	xor	0x7f(%rcx),%rbp
	xor	0x7f(%rcx),%rsi
	xor	0x7f(%rcx),%rdi
	xor	0x7f(%rcx),%r8
	xor	0x7f(%rcx),%r9
	xor	0x7f(%rcx),%r10
	xor	0x7f(%rcx),%r11
	xor	0x7f(%rcx),%r12
	xor	0x7f(%rcx),%r13
	xor	0x7f(%rcx),%r14
	xor	0x7f(%rcx),%r15
	nop
	xor	0x7f(%rdx),%rax
	xor	0x7f(%rdx),%rcx
	xor	0x7f(%rdx),%rdx
	xor	0x7f(%rdx),%rbx
	xor	0x7f(%rdx),%rsp
	xor	0x7f(%rdx),%rbp
	xor	0x7f(%rdx),%rsi
	xor	0x7f(%rdx),%rdi
	xor	0x7f(%rdx),%r8
	xor	0x7f(%rdx),%r9
	xor	0x7f(%rdx),%r10
	xor	0x7f(%rdx),%r11
	xor	0x7f(%rdx),%r12
	xor	0x7f(%rdx),%r13
	xor	0x7f(%rdx),%r14
	xor	0x7f(%rdx),%r15
	nop
	xor	0x7f(%rbx),%rax
	xor	0x7f(%rbx),%rcx
	xor	0x7f(%rbx),%rdx
	xor	0x7f(%rbx),%rbx
	xor	0x7f(%rbx),%rsp
	xor	0x7f(%rbx),%rbp
	xor	0x7f(%rbx),%rsi
	xor	0x7f(%rbx),%rdi
	xor	0x7f(%rbx),%r8
	xor	0x7f(%rbx),%r9
	xor	0x7f(%rbx),%r10
	xor	0x7f(%rbx),%r11
	xor	0x7f(%rbx),%r12
	xor	0x7f(%rbx),%r13
	xor	0x7f(%rbx),%r14
	xor	0x7f(%rbx),%r15
	nop
	xor	0x7f(%rsp),%rax
	xor	0x7f(%rsp),%rcx
	xor	0x7f(%rsp),%rdx
	xor	0x7f(%rsp),%rbx
	xor	0x7f(%rsp),%rsp
	xor	0x7f(%rsp),%rbp
	xor	0x7f(%rsp),%rsi
	xor	0x7f(%rsp),%rdi
	xor	0x7f(%rsp),%r8
	xor	0x7f(%rsp),%r9
	xor	0x7f(%rsp),%r10
	xor	0x7f(%rsp),%r11
	xor	0x7f(%rsp),%r12
	xor	0x7f(%rsp),%r13
	xor	0x7f(%rsp),%r14
	xor	0x7f(%rsp),%r15
	nop
	xor	0x7f(%rbp),%rax
	xor	0x7f(%rbp),%rcx
	xor	0x7f(%rbp),%rdx
	xor	0x7f(%rbp),%rbx
	xor	0x7f(%rbp),%rsp
	xor	0x7f(%rbp),%rbp
	xor	0x7f(%rbp),%rsi
	xor	0x7f(%rbp),%rdi
	xor	0x7f(%rbp),%r8
	xor	0x7f(%rbp),%r9
	xor	0x7f(%rbp),%r10
	xor	0x7f(%rbp),%r11
	xor	0x7f(%rbp),%r12
	xor	0x7f(%rbp),%r13
	xor	0x7f(%rbp),%r14
	xor	0x7f(%rbp),%r15
	nop
	xor	0x7f(%rsi),%rax
	xor	0x7f(%rsi),%rcx
	xor	0x7f(%rsi),%rdx
	xor	0x7f(%rsi),%rbx
	xor	0x7f(%rsi),%rsp
	xor	0x7f(%rsi),%rbp
	xor	0x7f(%rsi),%rsi
	xor	0x7f(%rsi),%rdi
	xor	0x7f(%rsi),%r8
	xor	0x7f(%rsi),%r9
	xor	0x7f(%rsi),%r10
	xor	0x7f(%rsi),%r11
	xor	0x7f(%rsi),%r12
	xor	0x7f(%rsi),%r13
	xor	0x7f(%rsi),%r14
	xor	0x7f(%rsi),%r15
	nop
	xor	0x7f(%rdi),%rax
	xor	0x7f(%rdi),%rcx
	xor	0x7f(%rdi),%rdx
	xor	0x7f(%rdi),%rbx
	xor	0x7f(%rdi),%rsp
	xor	0x7f(%rdi),%rbp
	xor	0x7f(%rdi),%rsi
	xor	0x7f(%rdi),%rdi
	xor	0x7f(%rdi),%r8
	xor	0x7f(%rdi),%r9
	xor	0x7f(%rdi),%r10
	xor	0x7f(%rdi),%r11
	xor	0x7f(%rdi),%r12
	xor	0x7f(%rdi),%r13
	xor	0x7f(%rdi),%r14
	xor	0x7f(%rdi),%r15
	nop
	xor	0x7f(%r8), %rax
	xor	0x7f(%r8), %rcx
	xor	0x7f(%r8), %rdx
	xor	0x7f(%r8), %rbx
	xor	0x7f(%r8), %rsp
	xor	0x7f(%r8), %rbp
	xor	0x7f(%r8), %rsi
	xor	0x7f(%r8), %rdi
	xor	0x7f(%r8), %r8
	xor	0x7f(%r8), %r9
	xor	0x7f(%r8), %r10
	xor	0x7f(%r8), %r11
	xor	0x7f(%r8), %r12
	xor	0x7f(%r8), %r13
	xor	0x7f(%r8), %r14
	xor	0x7f(%r8), %r15
	nop
	xor	0x7f(%r9), %rax
	xor	0x7f(%r9), %rcx
	xor	0x7f(%r9), %rdx
	xor	0x7f(%r9), %rbx
	xor	0x7f(%r9), %rsp
	xor	0x7f(%r9), %rbp
	xor	0x7f(%r9), %rsi
	xor	0x7f(%r9), %rdi
	xor	0x7f(%r9), %r8
	xor	0x7f(%r9), %r9
	xor	0x7f(%r9), %r10
	xor	0x7f(%r9), %r11
	xor	0x7f(%r9), %r12
	xor	0x7f(%r9), %r13
	xor	0x7f(%r9), %r14
	xor	0x7f(%r9), %r15
	nop
	xor	0x7f(%r10),%rax
	xor	0x7f(%r10),%rcx
	xor	0x7f(%r10),%rdx
	xor	0x7f(%r10),%rbx
	xor	0x7f(%r10),%rsp
	xor	0x7f(%r10),%rbp
	xor	0x7f(%r10),%rsi
	xor	0x7f(%r10),%rdi
	xor	0x7f(%r10),%r8
	xor	0x7f(%r10),%r9
	xor	0x7f(%r10),%r10
	xor	0x7f(%r10),%r11
	xor	0x7f(%r10),%r12
	xor	0x7f(%r10),%r13
	xor	0x7f(%r10),%r14
	xor	0x7f(%r10),%r15
	nop
	xor	0x7f(%r11),%rax
	xor	0x7f(%r11),%rcx
	xor	0x7f(%r11),%rdx
	xor	0x7f(%r11),%rbx
	xor	0x7f(%r11),%rsp
	xor	0x7f(%r11),%rbp
	xor	0x7f(%r11),%rsi
	xor	0x7f(%r11),%rdi
	xor	0x7f(%r11),%r8
	xor	0x7f(%r11),%r9
	xor	0x7f(%r11),%r10
	xor	0x7f(%r11),%r11
	xor	0x7f(%r11),%r12
	xor	0x7f(%r11),%r13
	xor	0x7f(%r11),%r14
	xor	0x7f(%r11),%r15
	nop
	xor	0x7f(%r12),%rax
	xor	0x7f(%r12),%rcx
	xor	0x7f(%r12),%rdx
	xor	0x7f(%r12),%rbx
	xor	0x7f(%r12),%rsp
	xor	0x7f(%r12),%rbp
	xor	0x7f(%r12),%rsi
	xor	0x7f(%r12),%rdi
	xor	0x7f(%r12),%r8
	xor	0x7f(%r12),%r9
	xor	0x7f(%r12),%r10
	xor	0x7f(%r12),%r11
	xor	0x7f(%r12),%r12
	xor	0x7f(%r12),%r13
	xor	0x7f(%r12),%r14
	xor	0x7f(%r12),%r15
	nop
	xor	0x7f(%r13),%rax
	xor	0x7f(%r13),%rcx
	xor	0x7f(%r13),%rdx
	xor	0x7f(%r13),%rbx
	xor	0x7f(%r13),%rsp
	xor	0x7f(%r13),%rbp
	xor	0x7f(%r13),%rsi
	xor	0x7f(%r13),%rdi
	xor	0x7f(%r13),%r8
	xor	0x7f(%r13),%r9
	xor	0x7f(%r13),%r10
	xor	0x7f(%r13),%r11
	xor	0x7f(%r13),%r12
	xor	0x7f(%r13),%r13
	xor	0x7f(%r13),%r14
	xor	0x7f(%r13),%r15
	nop
	xor	0x7f(%r14),%rax
	xor	0x7f(%r14),%rcx
	xor	0x7f(%r14),%rdx
	xor	0x7f(%r14),%rbx
	xor	0x7f(%r14),%rsp
	xor	0x7f(%r14),%rbp
	xor	0x7f(%r14),%rsi
	xor	0x7f(%r14),%rdi
	xor	0x7f(%r14),%r8
	xor	0x7f(%r14),%r9
	xor	0x7f(%r14),%r10
	xor	0x7f(%r14),%r11
	xor	0x7f(%r14),%r12
	xor	0x7f(%r14),%r13
	xor	0x7f(%r14),%r14
	xor	0x7f(%r14),%r15
	nop
	xor	0x7f(%r15),%rax
	xor	0x7f(%r15),%rcx
	xor	0x7f(%r15),%rdx
	xor	0x7f(%r15),%rbx
	xor	0x7f(%r15),%rsp
	xor	0x7f(%r15),%rbp
	xor	0x7f(%r15),%rsi
	xor	0x7f(%r15),%rdi
	xor	0x7f(%r15),%r8
	xor	0x7f(%r15),%r9
	xor	0x7f(%r15),%r10
	xor	0x7f(%r15),%r11
	xor	0x7f(%r15),%r12
	xor	0x7f(%r15),%r13
	xor	0x7f(%r15),%r14
	xor	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	XorRegMem32
	.type	XorRegMem32, @function
XorRegMem32:
	.cfi_startproc
	xor	0x7f563412(%rax),%rax
	xor	0x7f563412(%rax),%rcx
	xor	0x7f563412(%rax),%rdx
	xor	0x7f563412(%rax),%rbx
	xor	0x7f563412(%rax),%rsp
	xor	0x7f563412(%rax),%rbp
	xor	0x7f563412(%rax),%rsi
	xor	0x7f563412(%rax),%rdi
	xor	0x7f563412(%rax),%r8
	xor	0x7f563412(%rax),%r9
	xor	0x7f563412(%rax),%r10
	xor	0x7f563412(%rax),%r11
	xor	0x7f563412(%rax),%r12
	xor	0x7f563412(%rax),%r13
	xor	0x7f563412(%rax),%r14
	xor	0x7f563412(%rax),%r15
	nop
	xor	0x7f563412(%rcx),%rax
	xor	0x7f563412(%rcx),%rcx
	xor	0x7f563412(%rcx),%rdx
	xor	0x7f563412(%rcx),%rbx
	xor	0x7f563412(%rcx),%rsp
	xor	0x7f563412(%rcx),%rbp
	xor	0x7f563412(%rcx),%rsi
	xor	0x7f563412(%rcx),%rdi
	xor	0x7f563412(%rcx),%r8
	xor	0x7f563412(%rcx),%r9
	xor	0x7f563412(%rcx),%r10
	xor	0x7f563412(%rcx),%r11
	xor	0x7f563412(%rcx),%r12
	xor	0x7f563412(%rcx),%r13
	xor	0x7f563412(%rcx),%r14
	xor	0x7f563412(%rcx),%r15
	nop
	xor	0x7f563412(%rdx),%rax
	xor	0x7f563412(%rdx),%rcx
	xor	0x7f563412(%rdx),%rdx
	xor	0x7f563412(%rdx),%rbx
	xor	0x7f563412(%rdx),%rsp
	xor	0x7f563412(%rdx),%rbp
	xor	0x7f563412(%rdx),%rsi
	xor	0x7f563412(%rdx),%rdi
	xor	0x7f563412(%rdx),%r8
	xor	0x7f563412(%rdx),%r9
	xor	0x7f563412(%rdx),%r10
	xor	0x7f563412(%rdx),%r11
	xor	0x7f563412(%rdx),%r12
	xor	0x7f563412(%rdx),%r13
	xor	0x7f563412(%rdx),%r14
	xor	0x7f563412(%rdx),%r15
	nop
	xor	0x7f563412(%rbx),%rax
	xor	0x7f563412(%rbx),%rcx
	xor	0x7f563412(%rbx),%rdx
	xor	0x7f563412(%rbx),%rbx
	xor	0x7f563412(%rbx),%rsp
	xor	0x7f563412(%rbx),%rbp
	xor	0x7f563412(%rbx),%rsi
	xor	0x7f563412(%rbx),%rdi
	xor	0x7f563412(%rbx),%r8
	xor	0x7f563412(%rbx),%r9
	xor	0x7f563412(%rbx),%r10
	xor	0x7f563412(%rbx),%r11
	xor	0x7f563412(%rbx),%r12
	xor	0x7f563412(%rbx),%r13
	xor	0x7f563412(%rbx),%r14
	xor	0x7f563412(%rbx),%r15
	nop
	xor	0x7f563412(%rsp),%rax
	xor	0x7f563412(%rsp),%rcx
	xor	0x7f563412(%rsp),%rdx
	xor	0x7f563412(%rsp),%rbx
	xor	0x7f563412(%rsp),%rsp
	xor	0x7f563412(%rsp),%rbp
	xor	0x7f563412(%rsp),%rsi
	xor	0x7f563412(%rsp),%rdi
	xor	0x7f563412(%rsp),%r8
	xor	0x7f563412(%rsp),%r9
	xor	0x7f563412(%rsp),%r10
	xor	0x7f563412(%rsp),%r11
	xor	0x7f563412(%rsp),%r12
	xor	0x7f563412(%rsp),%r13
	xor	0x7f563412(%rsp),%r14
	xor	0x7f563412(%rsp),%r15
	nop
	xor	0x7f563412(%rbp),%rax
	xor	0x7f563412(%rbp),%rcx
	xor	0x7f563412(%rbp),%rdx
	xor	0x7f563412(%rbp),%rbx
	xor	0x7f563412(%rbp),%rsp
	xor	0x7f563412(%rbp),%rbp
	xor	0x7f563412(%rbp),%rsi
	xor	0x7f563412(%rbp),%rdi
	xor	0x7f563412(%rbp),%r8
	xor	0x7f563412(%rbp),%r9
	xor	0x7f563412(%rbp),%r10
	xor	0x7f563412(%rbp),%r11
	xor	0x7f563412(%rbp),%r12
	xor	0x7f563412(%rbp),%r13
	xor	0x7f563412(%rbp),%r14
	xor	0x7f563412(%rbp),%r15
	nop
	xor	0x7f563412(%rsi),%rax
	xor	0x7f563412(%rsi),%rcx
	xor	0x7f563412(%rsi),%rdx
	xor	0x7f563412(%rsi),%rbx
	xor	0x7f563412(%rsi),%rsp
	xor	0x7f563412(%rsi),%rbp
	xor	0x7f563412(%rsi),%rsi
	xor	0x7f563412(%rsi),%rdi
	xor	0x7f563412(%rsi),%r8
	xor	0x7f563412(%rsi),%r9
	xor	0x7f563412(%rsi),%r10
	xor	0x7f563412(%rsi),%r11
	xor	0x7f563412(%rsi),%r12
	xor	0x7f563412(%rsi),%r13
	xor	0x7f563412(%rsi),%r14
	xor	0x7f563412(%rsi),%r15
	nop
	xor	0x7f563412(%rdi),%rax
	xor	0x7f563412(%rdi),%rcx
	xor	0x7f563412(%rdi),%rdx
	xor	0x7f563412(%rdi),%rbx
	xor	0x7f563412(%rdi),%rsp
	xor	0x7f563412(%rdi),%rbp
	xor	0x7f563412(%rdi),%rsi
	xor	0x7f563412(%rdi),%rdi
	xor	0x7f563412(%rdi),%r8
	xor	0x7f563412(%rdi),%r9
	xor	0x7f563412(%rdi),%r10
	xor	0x7f563412(%rdi),%r11
	xor	0x7f563412(%rdi),%r12
	xor	0x7f563412(%rdi),%r13
	xor	0x7f563412(%rdi),%r14
	xor	0x7f563412(%rdi),%r15
	nop
	xor	0x7f563412(%r8), %rax
	xor	0x7f563412(%r8), %rcx
	xor	0x7f563412(%r8), %rdx
	xor	0x7f563412(%r8), %rbx
	xor	0x7f563412(%r8), %rsp
	xor	0x7f563412(%r8), %rbp
	xor	0x7f563412(%r8), %rsi
	xor	0x7f563412(%r8), %rdi
	xor	0x7f563412(%r8), %r8
	xor	0x7f563412(%r8), %r9
	xor	0x7f563412(%r8), %r10
	xor	0x7f563412(%r8), %r11
	xor	0x7f563412(%r8), %r12
	xor	0x7f563412(%r8), %r13
	xor	0x7f563412(%r8), %r14
	xor	0x7f563412(%r8), %r15
	nop
	xor	0x7f563412(%r9), %rax
	xor	0x7f563412(%r9), %rcx
	xor	0x7f563412(%r9), %rdx
	xor	0x7f563412(%r9), %rbx
	xor	0x7f563412(%r9), %rsp
	xor	0x7f563412(%r9), %rbp
	xor	0x7f563412(%r9), %rsi
	xor	0x7f563412(%r9), %rdi
	xor	0x7f563412(%r9), %r8
	xor	0x7f563412(%r9), %r9
	xor	0x7f563412(%r9), %r10
	xor	0x7f563412(%r9), %r11
	xor	0x7f563412(%r9), %r12
	xor	0x7f563412(%r9), %r13
	xor	0x7f563412(%r9), %r14
	xor	0x7f563412(%r9), %r15
	nop
	xor	0x7f563412(%r10),%rax
	xor	0x7f563412(%r10),%rcx
	xor	0x7f563412(%r10),%rdx
	xor	0x7f563412(%r10),%rbx
	xor	0x7f563412(%r10),%rsp
	xor	0x7f563412(%r10),%rbp
	xor	0x7f563412(%r10),%rsi
	xor	0x7f563412(%r10),%rdi
	xor	0x7f563412(%r10),%r8
	xor	0x7f563412(%r10),%r9
	xor	0x7f563412(%r10),%r10
	xor	0x7f563412(%r10),%r11
	xor	0x7f563412(%r10),%r12
	xor	0x7f563412(%r10),%r13
	xor	0x7f563412(%r10),%r14
	xor	0x7f563412(%r10),%r15
	nop
	xor	0x7f563412(%r11),%rax
	xor	0x7f563412(%r11),%rcx
	xor	0x7f563412(%r11),%rdx
	xor	0x7f563412(%r11),%rbx
	xor	0x7f563412(%r11),%rsp
	xor	0x7f563412(%r11),%rbp
	xor	0x7f563412(%r11),%rsi
	xor	0x7f563412(%r11),%rdi
	xor	0x7f563412(%r11),%r8
	xor	0x7f563412(%r11),%r9
	xor	0x7f563412(%r11),%r10
	xor	0x7f563412(%r11),%r11
	xor	0x7f563412(%r11),%r12
	xor	0x7f563412(%r11),%r13
	xor	0x7f563412(%r11),%r14
	xor	0x7f563412(%r11),%r15
	nop
	xor	0x7f563412(%r12),%rax
	xor	0x7f563412(%r12),%rcx
	xor	0x7f563412(%r12),%rdx
	xor	0x7f563412(%r12),%rbx
	xor	0x7f563412(%r12),%rsp
	xor	0x7f563412(%r12),%rbp
	xor	0x7f563412(%r12),%rsi
	xor	0x7f563412(%r12),%rdi
	xor	0x7f563412(%r12),%r8
	xor	0x7f563412(%r12),%r9
	xor	0x7f563412(%r12),%r10
	xor	0x7f563412(%r12),%r11
	xor	0x7f563412(%r12),%r12
	xor	0x7f563412(%r12),%r13
	xor	0x7f563412(%r12),%r14
	xor	0x7f563412(%r12),%r15
	nop
	xor	0x7f563412(%r13),%rax
	xor	0x7f563412(%r13),%rcx
	xor	0x7f563412(%r13),%rdx
	xor	0x7f563412(%r13),%rbx
	xor	0x7f563412(%r13),%rsp
	xor	0x7f563412(%r13),%rbp
	xor	0x7f563412(%r13),%rsi
	xor	0x7f563412(%r13),%rdi
	xor	0x7f563412(%r13),%r8
	xor	0x7f563412(%r13),%r9
	xor	0x7f563412(%r13),%r10
	xor	0x7f563412(%r13),%r11
	xor	0x7f563412(%r13),%r12
	xor	0x7f563412(%r13),%r13
	xor	0x7f563412(%r13),%r14
	xor	0x7f563412(%r13),%r15
	nop
	xor	0x7f563412(%r14),%rax
	xor	0x7f563412(%r14),%rcx
	xor	0x7f563412(%r14),%rdx
	xor	0x7f563412(%r14),%rbx
	xor	0x7f563412(%r14),%rsp
	xor	0x7f563412(%r14),%rbp
	xor	0x7f563412(%r14),%rsi
	xor	0x7f563412(%r14),%rdi
	xor	0x7f563412(%r14),%r8
	xor	0x7f563412(%r14),%r9
	xor	0x7f563412(%r14),%r10
	xor	0x7f563412(%r14),%r11
	xor	0x7f563412(%r14),%r12
	xor	0x7f563412(%r14),%r13
	xor	0x7f563412(%r14),%r14
	xor	0x7f563412(%r14),%r15
	nop
	xor	0x7f563412(%r15),%rax
	xor	0x7f563412(%r15),%rcx
	xor	0x7f563412(%r15),%rdx
	xor	0x7f563412(%r15),%rbx
	xor	0x7f563412(%r15),%rsp
	xor	0x7f563412(%r15),%rbp
	xor	0x7f563412(%r15),%rsi
	xor	0x7f563412(%r15),%rdi
	xor	0x7f563412(%r15),%r8
	xor	0x7f563412(%r15),%r9
	xor	0x7f563412(%r15),%r10
	xor	0x7f563412(%r15),%r11
	xor	0x7f563412(%r15),%r12
	xor	0x7f563412(%r15),%r13
	xor	0x7f563412(%r15),%r14
	xor	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


