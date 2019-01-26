	.file	"or.s"
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
	.globl	Or
	.type	Or, @function
Or:
	.cfi_startproc

        // reg8 += reg8
	or	%al,%al
	or	%al,%cl
	or	%al,%dl
	or	%al,%bl
	or	%al,%spl
	or	%al,%bpl
	or	%al,%sil
	or	%al,%dil
	or	%al,%r8b
	or	%al,%r9b
	or	%al,%r10b
	or	%al,%r11b
	or	%al,%r12b
	or	%al,%r13b
	or	%al,%r14b
	or	%al,%r15b
	nop
	or	%cl,%al
	or	%cl,%cl
	or	%cl,%dl
	or	%cl,%bl
	or	%cl,%spl
	or	%cl,%bpl
	or	%cl,%sil
	or	%cl,%dil
	or	%cl,%r8b
	or	%cl,%r9b
	or	%cl,%r10b
	or	%cl,%r11b
	or	%cl,%r12b
	or	%cl,%r13b
	or	%cl,%r14b
	or	%cl,%r15b
	nop
	or	%dl,%al
	or	%dl,%cl
	or	%dl,%dl
	or	%dl,%bl
	or	%dl,%spl
	or	%dl,%bpl
	or	%dl,%sil
	or	%dl,%dil
	or	%dl,%r8b
	or	%dl,%r9b
	or	%dl,%r10b
	or	%dl,%r11b
	or	%dl,%r12b
	or	%dl,%r13b
	or	%dl,%r14b
	or	%dl,%r15b
	nop
	or	%bl,%al
	or	%bl,%cl
	or	%bl,%dl
	or	%bl,%bl
	or	%bl,%spl
	or	%bl,%bpl
	or	%bl,%sil
	or	%bl,%dil
	or	%bl,%r8b
	or	%bl,%r9b
	or	%bl,%r10b
	or	%bl,%r11b
	or	%bl,%r12b
	or	%bl,%r13b
	or	%bl,%r14b
	or	%bl,%r15b
	nop
	or	%spl,%al
	or	%spl,%cl
	or	%spl,%dl
	or	%spl,%bl
	or	%spl,%spl
	or	%spl,%bpl
	or	%spl,%sil
	or	%spl,%dil
	or	%spl,%r8b
	or	%spl,%r9b
	or	%spl,%r10b
	or	%spl,%r11b
	or	%spl,%r12b
	or	%spl,%r13b
	or	%spl,%r14b
	or	%spl,%r15b
	nop
	or	%bpl,%al
	or	%bpl,%cl
	or	%bpl,%dl
	or	%bpl,%bl
	or	%bpl,%spl
	or	%bpl,%bpl
	or	%bpl,%sil
	or	%bpl,%dil
	or	%bpl,%r8b
	or	%bpl,%r9b
	or	%bpl,%r10b
	or	%bpl,%r11b
	or	%bpl,%r12b
	or	%bpl,%r13b
	or	%bpl,%r14b
	or	%bpl,%r15b
	nop
	or	%sil,%al
	or	%sil,%cl
	or	%sil,%dl
	or	%sil,%bl
	or	%sil,%spl
	or	%sil,%bpl
	or	%sil,%sil
	or	%sil,%dil
	or	%sil,%r8b
	or	%sil,%r9b
	or	%sil,%r10b
	or	%sil,%r11b
	or	%sil,%r12b
	or	%sil,%r13b
	or	%sil,%r14b
	or	%sil,%r15b
	nop
	or	%dil,%al
	or	%dil,%cl
	or	%dil,%dl
	or	%dil,%bl
	or	%dil,%spl
	or	%dil,%bpl
	or	%dil,%sil
	or	%dil,%dil
	or	%dil,%r8b
	or	%dil,%r9b
	or	%dil,%r10b
	or	%dil,%r11b
	or	%dil,%r12b
	or	%dil,%r13b
	or	%dil,%r14b
	or	%dil,%r15b
	nop
	or	%r8b, %al
	or	%r8b, %cl
	or	%r8b, %dl
	or	%r8b, %bl
	or	%r8b, %spl
	or	%r8b, %bpl
	or	%r8b, %sil
	or	%r8b, %dil
	or	%r8b, %r8b
	or	%r8b, %r9b
	or	%r8b, %r10b
	or	%r8b, %r11b
	or	%r8b, %r12b
	or	%r8b, %r13b
	or	%r8b, %r14b
	or	%r8b, %r15b
	nop
	or	%r9b, %al
	or	%r9b, %cl
	or	%r9b, %dl
	or	%r9b, %bl
	or	%r9b, %spl
	or	%r9b, %bpl
	or	%r9b, %sil
	or	%r9b, %dil
	or	%r9b, %r8b
	or	%r9b, %r9b
	or	%r9b, %r10b
	or	%r9b, %r11b
	or	%r9b, %r12b
	or	%r9b, %r13b
	or	%r9b, %r14b
	or	%r9b, %r15b
	nop
	or	%r10b,%al
	or	%r10b,%cl
	or	%r10b,%dl
	or	%r10b,%bl
	or	%r10b,%spl
	or	%r10b,%bpl
	or	%r10b,%sil
	or	%r10b,%dil
	or	%r10b,%r8b
	or	%r10b,%r9b
	or	%r10b,%r10b
	or	%r10b,%r11b
	or	%r10b,%r12b
	or	%r10b,%r13b
	or	%r10b,%r14b
	or	%r10b,%r15b
	nop
	or	%r11b,%al
	or	%r11b,%cl
	or	%r11b,%dl
	or	%r11b,%bl
	or	%r11b,%spl
	or	%r11b,%bpl
	or	%r11b,%sil
	or	%r11b,%dil
	or	%r11b,%r8b
	or	%r11b,%r9b
	or	%r11b,%r10b
	or	%r11b,%r11b
	or	%r11b,%r12b
	or	%r11b,%r13b
	or	%r11b,%r14b
	or	%r11b,%r15b
	nop
	or	%r12b,%al
	or	%r12b,%cl
	or	%r12b,%dl
	or	%r12b,%bl
	or	%r12b,%spl
	or	%r12b,%bpl
	or	%r12b,%sil
	or	%r12b,%dil
	or	%r12b,%r8b
	or	%r12b,%r9b
	or	%r12b,%r10b
	or	%r12b,%r11b
	or	%r12b,%r12b
	or	%r12b,%r13b
	or	%r12b,%r14b
	or	%r12b,%r15b
	nop
	or	%r13b,%al
	or	%r13b,%cl
	or	%r13b,%dl
	or	%r13b,%bl
	or	%r13b,%spl
	or	%r13b,%bpl
	or	%r13b,%sil
	or	%r13b,%dil
	or	%r13b,%r8b
	or	%r13b,%r9b
	or	%r13b,%r10b
	or	%r13b,%r11b
	or	%r13b,%r12b
	or	%r13b,%r13b
	or	%r13b,%r14b
	or	%r13b,%r15b
	nop
	or	%r14b,%al
	or	%r14b,%cl
	or	%r14b,%dl
	or	%r14b,%bl
	or	%r14b,%spl
	or	%r14b,%bpl
	or	%r14b,%sil
	or	%r14b,%dil
	or	%r14b,%r8b
	or	%r14b,%r9b
	or	%r14b,%r10b
	or	%r14b,%r11b
	or	%r14b,%r12b
	or	%r14b,%r13b
	or	%r14b,%r14b
	or	%r14b,%r15b
	nop
	or	%r15b,%al
	or	%r15b,%cl
	or	%r15b,%dl
	or	%r15b,%bl
	or	%r15b,%spl
	or	%r15b,%bpl
	or	%r15b,%sil
	or	%r15b,%dil
	or	%r15b,%r8b
	or	%r15b,%r9b
	or	%r15b,%r10b
	or	%r15b,%r11b
	or	%r15b,%r12b
	or	%r15b,%r13b
	or	%r15b,%r14b
	or	%r15b,%r15b
        nop
        nop
        // reg16 += reg16
	or	%ax,%ax
	or	%ax,%cx
	or	%ax,%dx
	or	%ax,%bx
	or	%ax,%sp
	or	%ax,%bp
	or	%ax,%si
	or	%ax,%di
	or	%ax,%r8w
	or	%ax,%r9w
	or	%ax,%r10w
	or	%ax,%r11w
	or	%ax,%r12w
	or	%ax,%r13w
	or	%ax,%r14w
	or	%ax,%r15w
	nop
	or	%cx,%ax
	or	%cx,%cx
	or	%cx,%dx
	or	%cx,%bx
	or	%cx,%sp
	or	%cx,%bp
	or	%cx,%si
	or	%cx,%di
	or	%cx,%r8w
	or	%cx,%r9w
	or	%cx,%r10w
	or	%cx,%r11w
	or	%cx,%r12w
	or	%cx,%r13w
	or	%cx,%r14w
	or	%cx,%r15w
	nop
	or	%dx,%ax
	or	%dx,%cx
	or	%dx,%dx
	or	%dx,%bx
	or	%dx,%sp
	or	%dx,%bp
	or	%dx,%si
	or	%dx,%di
	or	%dx,%r8w
	or	%dx,%r9w
	or	%dx,%r10w
	or	%dx,%r11w
	or	%dx,%r12w
	or	%dx,%r13w
	or	%dx,%r14w
	or	%dx,%r15w
	nop
	or	%bx,%ax
	or	%bx,%cx
	or	%bx,%dx
	or	%bx,%bx
	or	%bx,%sp
	or	%bx,%bp
	or	%bx,%si
	or	%bx,%di
	or	%bx,%r8w
	or	%bx,%r9w
	or	%bx,%r10w
	or	%bx,%r11w
	or	%bx,%r12w
	or	%bx,%r13w
	or	%bx,%r14w
	or	%bx,%r15w
	nop
	or	%sp,%ax
	or	%sp,%cx
	or	%sp,%dx
	or	%sp,%bx
	or	%sp,%sp
	or	%sp,%bp
	or	%sp,%si
	or	%sp,%di
	or	%sp,%r8w
	or	%sp,%r9w
	or	%sp,%r10w
	or	%sp,%r11w
	or	%sp,%r12w
	or	%sp,%r13w
	or	%sp,%r14w
	or	%sp,%r15w
	nop
	or	%bp,%ax
	or	%bp,%cx
	or	%bp,%dx
	or	%bp,%bx
	or	%bp,%sp
	or	%bp,%bp
	or	%bp,%si
	or	%bp,%di
	or	%bp,%r8w
	or	%bp,%r9w
	or	%bp,%r10w
	or	%bp,%r11w
	or	%bp,%r12w
	or	%bp,%r13w
	or	%bp,%r14w
	or	%bp,%r15w
	nop
	or	%si,%ax
	or	%si,%cx
	or	%si,%dx
	or	%si,%bx
	or	%si,%sp
	or	%si,%bp
	or	%si,%si
	or	%si,%di
	or	%si,%r8w
	or	%si,%r9w
	or	%si,%r10w
	or	%si,%r11w
	or	%si,%r12w
	or	%si,%r13w
	or	%si,%r14w
	or	%si,%r15w
	nop
	or	%di,%ax
	or	%di,%cx
	or	%di,%dx
	or	%di,%bx
	or	%di,%sp
	or	%di,%bp
	or	%di,%si
	or	%di,%di
	or	%di,%r8w
	or	%di,%r9w
	or	%di,%r10w
	or	%di,%r11w
	or	%di,%r12w
	or	%di,%r13w
	or	%di,%r14w
	or	%di,%r15w
	nop
	or	%r8w, %ax
	or	%r8w, %cx
	or	%r8w, %dx
	or	%r8w, %bx
	or	%r8w, %sp
	or	%r8w, %bp
	or	%r8w, %si
	or	%r8w, %di
	or	%r8w, %r8w
	or	%r8w, %r9w
	or	%r8w, %r10w
	or	%r8w, %r11w
	or	%r8w, %r12w
	or	%r8w, %r13w
	or	%r8w, %r14w
	or	%r8w, %r15w
	nop
	or	%r9w, %ax
	or	%r9w, %cx
	or	%r9w, %dx
	or	%r9w, %bx
	or	%r9w, %sp
	or	%r9w, %bp
	or	%r9w, %si
	or	%r9w, %di
	or	%r9w, %r8w
	or	%r9w, %r9w
	or	%r9w, %r10w
	or	%r9w, %r11w
	or	%r9w, %r12w
	or	%r9w, %r13w
	or	%r9w, %r14w
	or	%r9w, %r15w
	nop
	or	%r10w,%ax
	or	%r10w,%cx
	or	%r10w,%dx
	or	%r10w,%bx
	or	%r10w,%sp
	or	%r10w,%bp
	or	%r10w,%si
	or	%r10w,%di
	or	%r10w,%r8w
	or	%r10w,%r9w
	or	%r10w,%r10w
	or	%r10w,%r11w
	or	%r10w,%r12w
	or	%r10w,%r13w
	or	%r10w,%r14w
	or	%r10w,%r15w
	nop
	or	%r11w,%ax
	or	%r11w,%cx
	or	%r11w,%dx
	or	%r11w,%bx
	or	%r11w,%sp
	or	%r11w,%bp
	or	%r11w,%si
	or	%r11w,%di
	or	%r11w,%r8w
	or	%r11w,%r9w
	or	%r11w,%r10w
	or	%r11w,%r11w
	or	%r11w,%r12w
	or	%r11w,%r13w
	or	%r11w,%r14w
	or	%r11w,%r15w
	nop
	or	%r12w,%ax
	or	%r12w,%cx
	or	%r12w,%dx
	or	%r12w,%bx
	or	%r12w,%sp
	or	%r12w,%bp
	or	%r12w,%si
	or	%r12w,%di
	or	%r12w,%r8w
	or	%r12w,%r9w
	or	%r12w,%r10w
	or	%r12w,%r11w
	or	%r12w,%r12w
	or	%r12w,%r13w
	or	%r12w,%r14w
	or	%r12w,%r15w
	nop
	or	%r13w,%ax
	or	%r13w,%cx
	or	%r13w,%dx
	or	%r13w,%bx
	or	%r13w,%sp
	or	%r13w,%bp
	or	%r13w,%si
	or	%r13w,%di
	or	%r13w,%r8w
	or	%r13w,%r9w
	or	%r13w,%r10w
	or	%r13w,%r11w
	or	%r13w,%r12w
	or	%r13w,%r13w
	or	%r13w,%r14w
	or	%r13w,%r15w
	nop
	or	%r14w,%ax
	or	%r14w,%cx
	or	%r14w,%dx
	or	%r14w,%bx
	or	%r14w,%sp
	or	%r14w,%bp
	or	%r14w,%si
	or	%r14w,%di
	or	%r14w,%r8w
	or	%r14w,%r9w
	or	%r14w,%r10w
	or	%r14w,%r11w
	or	%r14w,%r12w
	or	%r14w,%r13w
	or	%r14w,%r14w
	or	%r14w,%r15w
	nop
	or	%r15w,%ax
	or	%r15w,%cx
	or	%r15w,%dx
	or	%r15w,%bx
	or	%r15w,%sp
	or	%r15w,%bp
	or	%r15w,%si
	or	%r15w,%di
	or	%r15w,%r8w
	or	%r15w,%r9w
	or	%r15w,%r10w
	or	%r15w,%r11w
	or	%r15w,%r12w
	or	%r15w,%r13w
	or	%r15w,%r14w
	or	%r15w,%r15w
        nop
        nop
        // reg32 += reg32
	or	%eax,%eax
	or	%eax,%ecx
	or	%eax,%edx
	or	%eax,%ebx
	or	%eax,%esp
	or	%eax,%ebp
	or	%eax,%esi
	or	%eax,%edi
	or	%eax,%r8d
	or	%eax,%r9d
	or	%eax,%r10d
	or	%eax,%r11d
	or	%eax,%r12d
	or	%eax,%r13d
	or	%eax,%r14d
	or	%eax,%r15d
	nop
	or	%ecx,%eax
	or	%ecx,%ecx
	or	%ecx,%edx
	or	%ecx,%ebx
	or	%ecx,%esp
	or	%ecx,%ebp
	or	%ecx,%esi
	or	%ecx,%edi
	or	%ecx,%r8d
	or	%ecx,%r9d
	or	%ecx,%r10d
	or	%ecx,%r11d
	or	%ecx,%r12d
	or	%ecx,%r13d
	or	%ecx,%r14d
	or	%ecx,%r15d
	nop
	or	%edx,%eax
	or	%edx,%ecx
	or	%edx,%edx
	or	%edx,%ebx
	or	%edx,%esp
	or	%edx,%ebp
	or	%edx,%esi
	or	%edx,%edi
	or	%edx,%r8d
	or	%edx,%r9d
	or	%edx,%r10d
	or	%edx,%r11d
	or	%edx,%r12d
	or	%edx,%r13d
	or	%edx,%r14d
	or	%edx,%r15d
	nop
	or	%ebx,%eax
	or	%ebx,%ecx
	or	%ebx,%edx
	or	%ebx,%ebx
	or	%ebx,%esp
	or	%ebx,%ebp
	or	%ebx,%esi
	or	%ebx,%edi
	or	%ebx,%r8d
	or	%ebx,%r9d
	or	%ebx,%r10d
	or	%ebx,%r11d
	or	%ebx,%r12d
	or	%ebx,%r13d
	or	%ebx,%r14d
	or	%ebx,%r15d
	nop
	or	%esp,%eax
	or	%esp,%ecx
	or	%esp,%edx
	or	%esp,%ebx
	or	%esp,%esp
	or	%esp,%ebp
	or	%esp,%esi
	or	%esp,%edi
	or	%esp,%r8d
	or	%esp,%r9d
	or	%esp,%r10d
	or	%esp,%r11d
	or	%esp,%r12d
	or	%esp,%r13d
	or	%esp,%r14d
	or	%esp,%r15d
	nop
	or	%ebp,%eax
	or	%ebp,%ecx
	or	%ebp,%edx
	or	%ebp,%ebx
	or	%ebp,%esp
	or	%ebp,%ebp
	or	%ebp,%esi
	or	%ebp,%edi
	or	%ebp,%r8d
	or	%ebp,%r9d
	or	%ebp,%r10d
	or	%ebp,%r11d
	or	%ebp,%r12d
	or	%ebp,%r13d
	or	%ebp,%r14d
	or	%ebp,%r15d
	nop
	or	%esi,%eax
	or	%esi,%ecx
	or	%esi,%edx
	or	%esi,%ebx
	or	%esi,%esp
	or	%esi,%ebp
	or	%esi,%esi
	or	%esi,%edi
	or	%esi,%r8d
	or	%esi,%r9d
	or	%esi,%r10d
	or	%esi,%r11d
	or	%esi,%r12d
	or	%esi,%r13d
	or	%esi,%r14d
	or	%esi,%r15d
	nop
	or	%edi,%eax
	or	%edi,%ecx
	or	%edi,%edx
	or	%edi,%ebx
	or	%edi,%esp
	or	%edi,%ebp
	or	%edi,%esi
	or	%edi,%edi
	or	%edi,%r8d
	or	%edi,%r9d
	or	%edi,%r10d
	or	%edi,%r11d
	or	%edi,%r12d
	or	%edi,%r13d
	or	%edi,%r14d
	or	%edi,%r15d
	nop
	or	%r8d, %eax
	or	%r8d, %ecx
	or	%r8d, %edx
	or	%r8d, %ebx
	or	%r8d, %esp
	or	%r8d, %ebp
	or	%r8d, %esi
	or	%r8d, %edi
	or	%r8d, %r8d
	or	%r8d, %r9d
	or	%r8d, %r10d
	or	%r8d, %r11d
	or	%r8d, %r12d
	or	%r8d, %r13d
	or	%r8d, %r14d
	or	%r8d, %r15d
	nop
	or	%r9d, %eax
	or	%r9d, %ecx
	or	%r9d, %edx
	or	%r9d, %ebx
	or	%r9d, %esp
	or	%r9d, %ebp
	or	%r9d, %esi
	or	%r9d, %edi
	or	%r9d, %r8d
	or	%r9d, %r9d
	or	%r9d, %r10d
	or	%r9d, %r11d
	or	%r9d, %r12d
	or	%r9d, %r13d
	or	%r9d, %r14d
	or	%r9d, %r15d
	nop
	or	%r10d,%eax
	or	%r10d,%ecx
	or	%r10d,%edx
	or	%r10d,%ebx
	or	%r10d,%esp
	or	%r10d,%ebp
	or	%r10d,%esi
	or	%r10d,%edi
	or	%r10d,%r8d
	or	%r10d,%r9d
	or	%r10d,%r10d
	or	%r10d,%r11d
	or	%r10d,%r12d
	or	%r10d,%r13d
	or	%r10d,%r14d
	or	%r10d,%r15d
	nop
	or	%r11d,%eax
	or	%r11d,%ecx
	or	%r11d,%edx
	or	%r11d,%ebx
	or	%r11d,%esp
	or	%r11d,%ebp
	or	%r11d,%esi
	or	%r11d,%edi
	or	%r11d,%r8d
	or	%r11d,%r9d
	or	%r11d,%r10d
	or	%r11d,%r11d
	or	%r11d,%r12d
	or	%r11d,%r13d
	or	%r11d,%r14d
	or	%r11d,%r15d
	nop
	or	%r12d,%eax
	or	%r12d,%ecx
	or	%r12d,%edx
	or	%r12d,%ebx
	or	%r12d,%esp
	or	%r12d,%ebp
	or	%r12d,%esi
	or	%r12d,%edi
	or	%r12d,%r8d
	or	%r12d,%r9d
	or	%r12d,%r10d
	or	%r12d,%r11d
	or	%r12d,%r12d
	or	%r12d,%r13d
	or	%r12d,%r14d
	or	%r12d,%r15d
	nop
	or	%r13d,%eax
	or	%r13d,%ecx
	or	%r13d,%edx
	or	%r13d,%ebx
	or	%r13d,%esp
	or	%r13d,%ebp
	or	%r13d,%esi
	or	%r13d,%edi
	or	%r13d,%r8d
	or	%r13d,%r9d
	or	%r13d,%r10d
	or	%r13d,%r11d
	or	%r13d,%r12d
	or	%r13d,%r13d
	or	%r13d,%r14d
	or	%r13d,%r15d
	nop
	or	%r14d,%eax
	or	%r14d,%ecx
	or	%r14d,%edx
	or	%r14d,%ebx
	or	%r14d,%esp
	or	%r14d,%ebp
	or	%r14d,%esi
	or	%r14d,%edi
	or	%r14d,%r8d
	or	%r14d,%r9d
	or	%r14d,%r10d
	or	%r14d,%r11d
	or	%r14d,%r12d
	or	%r14d,%r13d
	or	%r14d,%r14d
	or	%r14d,%r15d
	nop
	or	%r15d,%eax
	or	%r15d,%ecx
	or	%r15d,%edx
	or	%r15d,%ebx
	or	%r15d,%esp
	or	%r15d,%ebp
	or	%r15d,%esi
	or	%r15d,%edi
	or	%r15d,%r8d
	or	%r15d,%r9d
	or	%r15d,%r10d
	or	%r15d,%r11d
	or	%r15d,%r12d
	or	%r15d,%r13d
	or	%r15d,%r14d
	or	%r15d,%r15d
        nop
        nop
        // reg64 += reg64
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
	.globl	OrMemReg
	.type	OrMemReg, @function
OrMemReg:
	.cfi_startproc
        // mem8 += reg8
	or	%al,(%rax)
	or	%al,(%rcx)
	or	%al,(%rdx)
	or	%al,(%rbx)
	or	%al,(%rsp)
	or	%al,(%rbp)
	or	%al,(%rsi)
	or	%al,(%rdi)
	or	%al,(%r8)
	or	%al,(%r9)
	or	%al,(%r10)
	or	%al,(%r11)
	or	%al,(%r12)
	or	%al,(%r13)
	or	%al,(%r14)
	or	%al,(%r15)
	nop
	or	%cl,(%rax)
	or	%cl,(%rcx)
	or	%cl,(%rdx)
	or	%cl,(%rbx)
	or	%cl,(%rsp)
	or	%cl,(%rbp)
	or	%cl,(%rsi)
	or	%cl,(%rdi)
	or	%cl,(%r8)
	or	%cl,(%r9)
	or	%cl,(%r10)
	or	%cl,(%r11)
	or	%cl,(%r12)
	or	%cl,(%r13)
	or	%cl,(%r14)
	or	%cl,(%r15)
	nop
	or	%dl,(%rax)
	or	%dl,(%rcx)
	or	%dl,(%rdx)
	or	%dl,(%rbx)
	or	%dl,(%rsp)
	or	%dl,(%rbp)
	or	%dl,(%rsi)
	or	%dl,(%rdi)
	or	%dl,(%r8)
	or	%dl,(%r9)
	or	%dl,(%r10)
	or	%dl,(%r11)
	or	%dl,(%r12)
	or	%dl,(%r13)
	or	%dl,(%r14)
	or	%dl,(%r15)
	nop
	or	%bl,(%rax)
	or	%bl,(%rcx)
	or	%bl,(%rdx)
	or	%bl,(%rbx)
	or	%bl,(%rsp)
	or	%bl,(%rbp)
	or	%bl,(%rsi)
	or	%bl,(%rdi)
	or	%bl,(%r8)
	or	%bl,(%r9)
	or	%bl,(%r10)
	or	%bl,(%r11)
	or	%bl,(%r12)
	or	%bl,(%r13)
	or	%bl,(%r14)
	or	%bl,(%r15)
	nop
	or	%spl,(%rax)
	or	%spl,(%rcx)
	or	%spl,(%rdx)
	or	%spl,(%rbx)
	or	%spl,(%rsp)
	or	%spl,(%rbp)
	or	%spl,(%rsi)
	or	%spl,(%rdi)
	or	%spl,(%r8)
	or	%spl,(%r9)
	or	%spl,(%r10)
	or	%spl,(%r11)
	or	%spl,(%r12)
	or	%spl,(%r13)
	or	%spl,(%r14)
	or	%spl,(%r15)
	nop
	or	%bpl,(%rax)
	or	%bpl,(%rcx)
	or	%bpl,(%rdx)
	or	%bpl,(%rbx)
	or	%bpl,(%rsp)
	or	%bpl,(%rbp)
	or	%bpl,(%rsi)
	or	%bpl,(%rdi)
	or	%bpl,(%r8)
	or	%bpl,(%r9)
	or	%bpl,(%r10)
	or	%bpl,(%r11)
	or	%bpl,(%r12)
	or	%bpl,(%r13)
	or	%bpl,(%r14)
	or	%bpl,(%r15)
	nop
	or	%sil,(%rax)
	or	%sil,(%rcx)
	or	%sil,(%rdx)
	or	%sil,(%rbx)
	or	%sil,(%rsp)
	or	%sil,(%rbp)
	or	%sil,(%rsi)
	or	%sil,(%rdi)
	or	%sil,(%r8)
	or	%sil,(%r9)
	or	%sil,(%r10)
	or	%sil,(%r11)
	or	%sil,(%r12)
	or	%sil,(%r13)
	or	%sil,(%r14)
	or	%sil,(%r15)
	nop
	or	%dil,(%rax)
	or	%dil,(%rcx)
	or	%dil,(%rdx)
	or	%dil,(%rbx)
	or	%dil,(%rsp)
	or	%dil,(%rbp)
	or	%dil,(%rsi)
	or	%dil,(%rdi)
	or	%dil,(%r8)
	or	%dil,(%r9)
	or	%dil,(%r10)
	or	%dil,(%r11)
	or	%dil,(%r12)
	or	%dil,(%r13)
	or	%dil,(%r14)
	or	%dil,(%r15)
	nop
	or	%r8b, (%rax)
	or	%r8b, (%rcx)
	or	%r8b, (%rdx)
	or	%r8b, (%rbx)
	or	%r8b, (%rsp)
	or	%r8b, (%rbp)
	or	%r8b, (%rsi)
	or	%r8b, (%rdi)
	or	%r8b, (%r8)
	or	%r8b, (%r9)
	or	%r8b, (%r10)
	or	%r8b, (%r11)
	or	%r8b, (%r12)
	or	%r8b, (%r13)
	or	%r8b, (%r14)
	or	%r8b, (%r15)
	nop
	or	%r9b, (%rax)
	or	%r9b, (%rcx)
	or	%r9b, (%rdx)
	or	%r9b, (%rbx)
	or	%r9b, (%rsp)
	or	%r9b, (%rbp)
	or	%r9b, (%rsi)
	or	%r9b, (%rdi)
	or	%r9b, (%r8)
	or	%r9b, (%r9)
	or	%r9b, (%r10)
	or	%r9b, (%r11)
	or	%r9b, (%r12)
	or	%r9b, (%r13)
	or	%r9b, (%r14)
	or	%r9b, (%r15)
	nop
	or	%r10b,(%rax)
	or	%r10b,(%rcx)
	or	%r10b,(%rdx)
	or	%r10b,(%rbx)
	or	%r10b,(%rsp)
	or	%r10b,(%rbp)
	or	%r10b,(%rsi)
	or	%r10b,(%rdi)
	or	%r10b,(%r8)
	or	%r10b,(%r9)
	or	%r10b,(%r10)
	or	%r10b,(%r11)
	or	%r10b,(%r12)
	or	%r10b,(%r13)
	or	%r10b,(%r14)
	or	%r10b,(%r15)
	nop
	or	%r11b,(%rax)
	or	%r11b,(%rcx)
	or	%r11b,(%rdx)
	or	%r11b,(%rbx)
	or	%r11b,(%rsp)
	or	%r11b,(%rbp)
	or	%r11b,(%rsi)
	or	%r11b,(%rdi)
	or	%r11b,(%r8)
	or	%r11b,(%r9)
	or	%r11b,(%r10)
	or	%r11b,(%r11)
	or	%r11b,(%r12)
	or	%r11b,(%r13)
	or	%r11b,(%r14)
	or	%r11b,(%r15)
	nop
	or	%r12b,(%rax)
	or	%r12b,(%rcx)
	or	%r12b,(%rdx)
	or	%r12b,(%rbx)
	or	%r12b,(%rsp)
	or	%r12b,(%rbp)
	or	%r12b,(%rsi)
	or	%r12b,(%rdi)
	or	%r12b,(%r8)
	or	%r12b,(%r9)
	or	%r12b,(%r10)
	or	%r12b,(%r11)
	or	%r12b,(%r12)
	or	%r12b,(%r13)
	or	%r12b,(%r14)
	or	%r12b,(%r15)
	nop
	or	%r13b,(%rax)
	or	%r13b,(%rcx)
	or	%r13b,(%rdx)
	or	%r13b,(%rbx)
	or	%r13b,(%rsp)
	or	%r13b,(%rbp)
	or	%r13b,(%rsi)
	or	%r13b,(%rdi)
	or	%r13b,(%r8)
	or	%r13b,(%r9)
	or	%r13b,(%r10)
	or	%r13b,(%r11)
	or	%r13b,(%r12)
	or	%r13b,(%r13)
	or	%r13b,(%r14)
	or	%r13b,(%r15)
	nop
	or	%r14b,(%rax)
	or	%r14b,(%rcx)
	or	%r14b,(%rdx)
	or	%r14b,(%rbx)
	or	%r14b,(%rsp)
	or	%r14b,(%rbp)
	or	%r14b,(%rsi)
	or	%r14b,(%rdi)
	or	%r14b,(%r8)
	or	%r14b,(%r9)
	or	%r14b,(%r10)
	or	%r14b,(%r11)
	or	%r14b,(%r12)
	or	%r14b,(%r13)
	or	%r14b,(%r14)
	or	%r14b,(%r15)
	nop
	or	%r15b,(%rax)
	or	%r15b,(%rcx)
	or	%r15b,(%rdx)
	or	%r15b,(%rbx)
	or	%r15b,(%rsp)
	or	%r15b,(%rbp)
	or	%r15b,(%rsi)
	or	%r15b,(%rdi)
	or	%r15b,(%r8)
	or	%r15b,(%r9)
	or	%r15b,(%r10)
	or	%r15b,(%r11)
	or	%r15b,(%r12)
	or	%r15b,(%r13)
	or	%r15b,(%r14)
	or	%r15b,(%r15)
        nop
        nop
        // mem16 += reg16
	or	%ax,(%rax)
	or	%ax,(%rcx)
	or	%ax,(%rdx)
	or	%ax,(%rbx)
	or	%ax,(%rsp)
	or	%ax,(%rbp)
	or	%ax,(%rsi)
	or	%ax,(%rdi)
	or	%ax,(%r8)
	or	%ax,(%r9)
	or	%ax,(%r10)
	or	%ax,(%r11)
	or	%ax,(%r12)
	or	%ax,(%r13)
	or	%ax,(%r14)
	or	%ax,(%r15)
	nop
	or	%cx,(%rax)
	or	%cx,(%rcx)
	or	%cx,(%rdx)
	or	%cx,(%rbx)
	or	%cx,(%rsp)
	or	%cx,(%rbp)
	or	%cx,(%rsi)
	or	%cx,(%rdi)
	or	%cx,(%r8)
	or	%cx,(%r9)
	or	%cx,(%r10)
	or	%cx,(%r11)
	or	%cx,(%r12)
	or	%cx,(%r13)
	or	%cx,(%r14)
	or	%cx,(%r15)
	nop
	or	%dx,(%rax)
	or	%dx,(%rcx)
	or	%dx,(%rdx)
	or	%dx,(%rbx)
	or	%dx,(%rsp)
	or	%dx,(%rbp)
	or	%dx,(%rsi)
	or	%dx,(%rdi)
	or	%dx,(%r8)
	or	%dx,(%r9)
	or	%dx,(%r10)
	or	%dx,(%r11)
	or	%dx,(%r12)
	or	%dx,(%r13)
	or	%dx,(%r14)
	or	%dx,(%r15)
	nop
	or	%bx,(%rax)
	or	%bx,(%rcx)
	or	%bx,(%rdx)
	or	%bx,(%rbx)
	or	%bx,(%rsp)
	or	%bx,(%rbp)
	or	%bx,(%rsi)
	or	%bx,(%rdi)
	or	%bx,(%r8)
	or	%bx,(%r9)
	or	%bx,(%r10)
	or	%bx,(%r11)
	or	%bx,(%r12)
	or	%bx,(%r13)
	or	%bx,(%r14)
	or	%bx,(%r15)
	nop
	or	%sp,(%rax)
	or	%sp,(%rcx)
	or	%sp,(%rdx)
	or	%sp,(%rbx)
	or	%sp,(%rsp)
	or	%sp,(%rbp)
	or	%sp,(%rsi)
	or	%sp,(%rdi)
	or	%sp,(%r8)
	or	%sp,(%r9)
	or	%sp,(%r10)
	or	%sp,(%r11)
	or	%sp,(%r12)
	or	%sp,(%r13)
	or	%sp,(%r14)
	or	%sp,(%r15)
	nop
	or	%bp,(%rax)
	or	%bp,(%rcx)
	or	%bp,(%rdx)
	or	%bp,(%rbx)
	or	%bp,(%rsp)
	or	%bp,(%rbp)
	or	%bp,(%rsi)
	or	%bp,(%rdi)
	or	%bp,(%r8)
	or	%bp,(%r9)
	or	%bp,(%r10)
	or	%bp,(%r11)
	or	%bp,(%r12)
	or	%bp,(%r13)
	or	%bp,(%r14)
	or	%bp,(%r15)
	nop
	or	%si,(%rax)
	or	%si,(%rcx)
	or	%si,(%rdx)
	or	%si,(%rbx)
	or	%si,(%rsp)
	or	%si,(%rbp)
	or	%si,(%rsi)
	or	%si,(%rdi)
	or	%si,(%r8)
	or	%si,(%r9)
	or	%si,(%r10)
	or	%si,(%r11)
	or	%si,(%r12)
	or	%si,(%r13)
	or	%si,(%r14)
	or	%si,(%r15)
	nop
	or	%di,(%rax)
	or	%di,(%rcx)
	or	%di,(%rdx)
	or	%di,(%rbx)
	or	%di,(%rsp)
	or	%di,(%rbp)
	or	%di,(%rsi)
	or	%di,(%rdi)
	or	%di,(%r8)
	or	%di,(%r9)
	or	%di,(%r10)
	or	%di,(%r11)
	or	%di,(%r12)
	or	%di,(%r13)
	or	%di,(%r14)
	or	%di,(%r15)
	nop
	or	%r8w, (%rax)
	or	%r8w, (%rcx)
	or	%r8w, (%rdx)
	or	%r8w, (%rbx)
	or	%r8w, (%rsp)
	or	%r8w, (%rbp)
	or	%r8w, (%rsi)
	or	%r8w, (%rdi)
	or	%r8w, (%r8)
	or	%r8w, (%r9)
	or	%r8w, (%r10)
	or	%r8w, (%r11)
	or	%r8w, (%r12)
	or	%r8w, (%r13)
	or	%r8w, (%r14)
	or	%r8w, (%r15)
	nop
	or	%r9w, (%rax)
	or	%r9w, (%rcx)
	or	%r9w, (%rdx)
	or	%r9w, (%rbx)
	or	%r9w, (%rsp)
	or	%r9w, (%rbp)
	or	%r9w, (%rsi)
	or	%r9w, (%rdi)
	or	%r9w, (%r8)
	or	%r9w, (%r9)
	or	%r9w, (%r10)
	or	%r9w, (%r11)
	or	%r9w, (%r12)
	or	%r9w, (%r13)
	or	%r9w, (%r14)
	or	%r9w, (%r15)
	nop
	or	%r10w,(%rax)
	or	%r10w,(%rcx)
	or	%r10w,(%rdx)
	or	%r10w,(%rbx)
	or	%r10w,(%rsp)
	or	%r10w,(%rbp)
	or	%r10w,(%rsi)
	or	%r10w,(%rdi)
	or	%r10w,(%r8)
	or	%r10w,(%r9)
	or	%r10w,(%r10)
	or	%r10w,(%r11)
	or	%r10w,(%r12)
	or	%r10w,(%r13)
	or	%r10w,(%r14)
	or	%r10w,(%r15)
	nop
	or	%r11w,(%rax)
	or	%r11w,(%rcx)
	or	%r11w,(%rdx)
	or	%r11w,(%rbx)
	or	%r11w,(%rsp)
	or	%r11w,(%rbp)
	or	%r11w,(%rsi)
	or	%r11w,(%rdi)
	or	%r11w,(%r8)
	or	%r11w,(%r9)
	or	%r11w,(%r10)
	or	%r11w,(%r11)
	or	%r11w,(%r12)
	or	%r11w,(%r13)
	or	%r11w,(%r14)
	or	%r11w,(%r15)
	nop
	or	%r12w,(%rax)
	or	%r12w,(%rcx)
	or	%r12w,(%rdx)
	or	%r12w,(%rbx)
	or	%r12w,(%rsp)
	or	%r12w,(%rbp)
	or	%r12w,(%rsi)
	or	%r12w,(%rdi)
	or	%r12w,(%r8)
	or	%r12w,(%r9)
	or	%r12w,(%r10)
	or	%r12w,(%r11)
	or	%r12w,(%r12)
	or	%r12w,(%r13)
	or	%r12w,(%r14)
	or	%r12w,(%r15)
	nop
	or	%r13w,(%rax)
	or	%r13w,(%rcx)
	or	%r13w,(%rdx)
	or	%r13w,(%rbx)
	or	%r13w,(%rsp)
	or	%r13w,(%rbp)
	or	%r13w,(%rsi)
	or	%r13w,(%rdi)
	or	%r13w,(%r8)
	or	%r13w,(%r9)
	or	%r13w,(%r10)
	or	%r13w,(%r11)
	or	%r13w,(%r12)
	or	%r13w,(%r13)
	or	%r13w,(%r14)
	or	%r13w,(%r15)
	nop
	or	%r14w,(%rax)
	or	%r14w,(%rcx)
	or	%r14w,(%rdx)
	or	%r14w,(%rbx)
	or	%r14w,(%rsp)
	or	%r14w,(%rbp)
	or	%r14w,(%rsi)
	or	%r14w,(%rdi)
	or	%r14w,(%r8)
	or	%r14w,(%r9)
	or	%r14w,(%r10)
	or	%r14w,(%r11)
	or	%r14w,(%r12)
	or	%r14w,(%r13)
	or	%r14w,(%r14)
	or	%r14w,(%r15)
	nop
	or	%r15w,(%rax)
	or	%r15w,(%rcx)
	or	%r15w,(%rdx)
	or	%r15w,(%rbx)
	or	%r15w,(%rsp)
	or	%r15w,(%rbp)
	or	%r15w,(%rsi)
	or	%r15w,(%rdi)
	or	%r15w,(%r8)
	or	%r15w,(%r9)
	or	%r15w,(%r10)
	or	%r15w,(%r11)
	or	%r15w,(%r12)
	or	%r15w,(%r13)
	or	%r15w,(%r14)
	or	%r15w,(%r15)
        nop
        nop
        // mem32 += reg32
	or	%eax,(%rax)
	or	%eax,(%rcx)
	or	%eax,(%rdx)
	or	%eax,(%rbx)
	or	%eax,(%rsp)
	or	%eax,(%rbp)
	or	%eax,(%rsi)
	or	%eax,(%rdi)
	or	%eax,(%r8)
	or	%eax,(%r9)
	or	%eax,(%r10)
	or	%eax,(%r11)
	or	%eax,(%r12)
	or	%eax,(%r13)
	or	%eax,(%r14)
	or	%eax,(%r15)
	nop
	or	%ecx,(%rax)
	or	%ecx,(%rcx)
	or	%ecx,(%rdx)
	or	%ecx,(%rbx)
	or	%ecx,(%rsp)
	or	%ecx,(%rbp)
	or	%ecx,(%rsi)
	or	%ecx,(%rdi)
	or	%ecx,(%r8)
	or	%ecx,(%r9)
	or	%ecx,(%r10)
	or	%ecx,(%r11)
	or	%ecx,(%r12)
	or	%ecx,(%r13)
	or	%ecx,(%r14)
	or	%ecx,(%r15)
	nop
	or	%edx,(%rax)
	or	%edx,(%rcx)
	or	%edx,(%rdx)
	or	%edx,(%rbx)
	or	%edx,(%rsp)
	or	%edx,(%rbp)
	or	%edx,(%rsi)
	or	%edx,(%rdi)
	or	%edx,(%r8)
	or	%edx,(%r9)
	or	%edx,(%r10)
	or	%edx,(%r11)
	or	%edx,(%r12)
	or	%edx,(%r13)
	or	%edx,(%r14)
	or	%edx,(%r15)
	nop
	or	%ebx,(%rax)
	or	%ebx,(%rcx)
	or	%ebx,(%rdx)
	or	%ebx,(%rbx)
	or	%ebx,(%rsp)
	or	%ebx,(%rbp)
	or	%ebx,(%rsi)
	or	%ebx,(%rdi)
	or	%ebx,(%r8)
	or	%ebx,(%r9)
	or	%ebx,(%r10)
	or	%ebx,(%r11)
	or	%ebx,(%r12)
	or	%ebx,(%r13)
	or	%ebx,(%r14)
	or	%ebx,(%r15)
	nop
	or	%esp,(%rax)
	or	%esp,(%rcx)
	or	%esp,(%rdx)
	or	%esp,(%rbx)
	or	%esp,(%rsp)
	or	%esp,(%rbp)
	or	%esp,(%rsi)
	or	%esp,(%rdi)
	or	%esp,(%r8)
	or	%esp,(%r9)
	or	%esp,(%r10)
	or	%esp,(%r11)
	or	%esp,(%r12)
	or	%esp,(%r13)
	or	%esp,(%r14)
	or	%esp,(%r15)
	nop
	or	%ebp,(%rax)
	or	%ebp,(%rcx)
	or	%ebp,(%rdx)
	or	%ebp,(%rbx)
	or	%ebp,(%rsp)
	or	%ebp,(%rbp)
	or	%ebp,(%rsi)
	or	%ebp,(%rdi)
	or	%ebp,(%r8)
	or	%ebp,(%r9)
	or	%ebp,(%r10)
	or	%ebp,(%r11)
	or	%ebp,(%r12)
	or	%ebp,(%r13)
	or	%ebp,(%r14)
	or	%ebp,(%r15)
	nop
	or	%esi,(%rax)
	or	%esi,(%rcx)
	or	%esi,(%rdx)
	or	%esi,(%rbx)
	or	%esi,(%rsp)
	or	%esi,(%rbp)
	or	%esi,(%rsi)
	or	%esi,(%rdi)
	or	%esi,(%r8)
	or	%esi,(%r9)
	or	%esi,(%r10)
	or	%esi,(%r11)
	or	%esi,(%r12)
	or	%esi,(%r13)
	or	%esi,(%r14)
	or	%esi,(%r15)
	nop
	or	%edi,(%rax)
	or	%edi,(%rcx)
	or	%edi,(%rdx)
	or	%edi,(%rbx)
	or	%edi,(%rsp)
	or	%edi,(%rbp)
	or	%edi,(%rsi)
	or	%edi,(%rdi)
	or	%edi,(%r8)
	or	%edi,(%r9)
	or	%edi,(%r10)
	or	%edi,(%r11)
	or	%edi,(%r12)
	or	%edi,(%r13)
	or	%edi,(%r14)
	or	%edi,(%r15)
	nop
	or	%r8d, (%rax)
	or	%r8d, (%rcx)
	or	%r8d, (%rdx)
	or	%r8d, (%rbx)
	or	%r8d, (%rsp)
	or	%r8d, (%rbp)
	or	%r8d, (%rsi)
	or	%r8d, (%rdi)
	or	%r8d, (%r8)
	or	%r8d, (%r9)
	or	%r8d, (%r10)
	or	%r8d, (%r11)
	or	%r8d, (%r12)
	or	%r8d, (%r13)
	or	%r8d, (%r14)
	or	%r8d, (%r15)
	nop
	or	%r9d, (%rax)
	or	%r9d, (%rcx)
	or	%r9d, (%rdx)
	or	%r9d, (%rbx)
	or	%r9d, (%rsp)
	or	%r9d, (%rbp)
	or	%r9d, (%rsi)
	or	%r9d, (%rdi)
	or	%r9d, (%r8)
	or	%r9d, (%r9)
	or	%r9d, (%r10)
	or	%r9d, (%r11)
	or	%r9d, (%r12)
	or	%r9d, (%r13)
	or	%r9d, (%r14)
	or	%r9d, (%r15)
	nop
	or	%r10d,(%rax)
	or	%r10d,(%rcx)
	or	%r10d,(%rdx)
	or	%r10d,(%rbx)
	or	%r10d,(%rsp)
	or	%r10d,(%rbp)
	or	%r10d,(%rsi)
	or	%r10d,(%rdi)
	or	%r10d,(%r8)
	or	%r10d,(%r9)
	or	%r10d,(%r10)
	or	%r10d,(%r11)
	or	%r10d,(%r12)
	or	%r10d,(%r13)
	or	%r10d,(%r14)
	or	%r10d,(%r15)
	nop
	or	%r11d,(%rax)
	or	%r11d,(%rcx)
	or	%r11d,(%rdx)
	or	%r11d,(%rbx)
	or	%r11d,(%rsp)
	or	%r11d,(%rbp)
	or	%r11d,(%rsi)
	or	%r11d,(%rdi)
	or	%r11d,(%r8)
	or	%r11d,(%r9)
	or	%r11d,(%r10)
	or	%r11d,(%r11)
	or	%r11d,(%r12)
	or	%r11d,(%r13)
	or	%r11d,(%r14)
	or	%r11d,(%r15)
	nop
	or	%r12d,(%rax)
	or	%r12d,(%rcx)
	or	%r12d,(%rdx)
	or	%r12d,(%rbx)
	or	%r12d,(%rsp)
	or	%r12d,(%rbp)
	or	%r12d,(%rsi)
	or	%r12d,(%rdi)
	or	%r12d,(%r8)
	or	%r12d,(%r9)
	or	%r12d,(%r10)
	or	%r12d,(%r11)
	or	%r12d,(%r12)
	or	%r12d,(%r13)
	or	%r12d,(%r14)
	or	%r12d,(%r15)
	nop
	or	%r13d,(%rax)
	or	%r13d,(%rcx)
	or	%r13d,(%rdx)
	or	%r13d,(%rbx)
	or	%r13d,(%rsp)
	or	%r13d,(%rbp)
	or	%r13d,(%rsi)
	or	%r13d,(%rdi)
	or	%r13d,(%r8)
	or	%r13d,(%r9)
	or	%r13d,(%r10)
	or	%r13d,(%r11)
	or	%r13d,(%r12)
	or	%r13d,(%r13)
	or	%r13d,(%r14)
	or	%r13d,(%r15)
	nop
	or	%r14d,(%rax)
	or	%r14d,(%rcx)
	or	%r14d,(%rdx)
	or	%r14d,(%rbx)
	or	%r14d,(%rsp)
	or	%r14d,(%rbp)
	or	%r14d,(%rsi)
	or	%r14d,(%rdi)
	or	%r14d,(%r8)
	or	%r14d,(%r9)
	or	%r14d,(%r10)
	or	%r14d,(%r11)
	or	%r14d,(%r12)
	or	%r14d,(%r13)
	or	%r14d,(%r14)
	or	%r14d,(%r15)
	nop
	or	%r15d,(%rax)
	or	%r15d,(%rcx)
	or	%r15d,(%rdx)
	or	%r15d,(%rbx)
	or	%r15d,(%rsp)
	or	%r15d,(%rbp)
	or	%r15d,(%rsi)
	or	%r15d,(%rdi)
	or	%r15d,(%r8)
	or	%r15d,(%r9)
	or	%r15d,(%r10)
	or	%r15d,(%r11)
	or	%r15d,(%r12)
	or	%r15d,(%r13)
	or	%r15d,(%r14)
	or	%r15d,(%r15)
        nop
        nop
        // mem64 += reg64
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
	.globl	OrMem8Reg
	.type	OrMem8Reg, @function
OrMem8Reg:
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
	.globl	OrMem32Reg
	.type	OrMem32Reg, @function
OrMem32Reg:
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
        
        
	.p2align 4,,15
	.globl	OrRegMem
	.type	OrRegMem, @function
OrRegMem:
	.cfi_startproc
	or	(%rax),%rax
	or	(%rax),%rcx
	or	(%rax),%rdx
	or	(%rax),%rbx
	or	(%rax),%rsp
	or	(%rax),%rbp
	or	(%rax),%rsi
	or	(%rax),%rdi
	or	(%rax),%r8
	or	(%rax),%r9
	or	(%rax),%r10
	or	(%rax),%r11
	or	(%rax),%r12
	or	(%rax),%r13
	or	(%rax),%r14
	or	(%rax),%r15
	nop
	or	(%rcx),%rax
	or	(%rcx),%rcx
	or	(%rcx),%rdx
	or	(%rcx),%rbx
	or	(%rcx),%rsp
	or	(%rcx),%rbp
	or	(%rcx),%rsi
	or	(%rcx),%rdi
	or	(%rcx),%r8
	or	(%rcx),%r9
	or	(%rcx),%r10
	or	(%rcx),%r11
	or	(%rcx),%r12
	or	(%rcx),%r13
	or	(%rcx),%r14
	or	(%rcx),%r15
	nop
	or	(%rdx),%rax
	or	(%rdx),%rcx
	or	(%rdx),%rdx
	or	(%rdx),%rbx
	or	(%rdx),%rsp
	or	(%rdx),%rbp
	or	(%rdx),%rsi
	or	(%rdx),%rdi
	or	(%rdx),%r8
	or	(%rdx),%r9
	or	(%rdx),%r10
	or	(%rdx),%r11
	or	(%rdx),%r12
	or	(%rdx),%r13
	or	(%rdx),%r14
	or	(%rdx),%r15
	nop
	or	(%rbx),%rax
	or	(%rbx),%rcx
	or	(%rbx),%rdx
	or	(%rbx),%rbx
	or	(%rbx),%rsp
	or	(%rbx),%rbp
	or	(%rbx),%rsi
	or	(%rbx),%rdi
	or	(%rbx),%r8
	or	(%rbx),%r9
	or	(%rbx),%r10
	or	(%rbx),%r11
	or	(%rbx),%r12
	or	(%rbx),%r13
	or	(%rbx),%r14
	or	(%rbx),%r15
	nop
	or	(%rsp),%rax
	or	(%rsp),%rcx
	or	(%rsp),%rdx
	or	(%rsp),%rbx
	or	(%rsp),%rsp
	or	(%rsp),%rbp
	or	(%rsp),%rsi
	or	(%rsp),%rdi
	or	(%rsp),%r8
	or	(%rsp),%r9
	or	(%rsp),%r10
	or	(%rsp),%r11
	or	(%rsp),%r12
	or	(%rsp),%r13
	or	(%rsp),%r14
	or	(%rsp),%r15
	nop
	or	(%rbp),%rax
	or	(%rbp),%rcx
	or	(%rbp),%rdx
	or	(%rbp),%rbx
	or	(%rbp),%rsp
	or	(%rbp),%rbp
	or	(%rbp),%rsi
	or	(%rbp),%rdi
	or	(%rbp),%r8
	or	(%rbp),%r9
	or	(%rbp),%r10
	or	(%rbp),%r11
	or	(%rbp),%r12
	or	(%rbp),%r13
	or	(%rbp),%r14
	or	(%rbp),%r15
	nop
	or	(%rsi),%rax
	or	(%rsi),%rcx
	or	(%rsi),%rdx
	or	(%rsi),%rbx
	or	(%rsi),%rsp
	or	(%rsi),%rbp
	or	(%rsi),%rsi
	or	(%rsi),%rdi
	or	(%rsi),%r8
	or	(%rsi),%r9
	or	(%rsi),%r10
	or	(%rsi),%r11
	or	(%rsi),%r12
	or	(%rsi),%r13
	or	(%rsi),%r14
	or	(%rsi),%r15
	nop
	or	(%rdi),%rax
	or	(%rdi),%rcx
	or	(%rdi),%rdx
	or	(%rdi),%rbx
	or	(%rdi),%rsp
	or	(%rdi),%rbp
	or	(%rdi),%rsi
	or	(%rdi),%rdi
	or	(%rdi),%r8
	or	(%rdi),%r9
	or	(%rdi),%r10
	or	(%rdi),%r11
	or	(%rdi),%r12
	or	(%rdi),%r13
	or	(%rdi),%r14
	or	(%rdi),%r15
	nop
	or	(%r8), %rax
	or	(%r8), %rcx
	or	(%r8), %rdx
	or	(%r8), %rbx
	or	(%r8), %rsp
	or	(%r8), %rbp
	or	(%r8), %rsi
	or	(%r8), %rdi
	or	(%r8), %r8
	or	(%r8), %r9
	or	(%r8), %r10
	or	(%r8), %r11
	or	(%r8), %r12
	or	(%r8), %r13
	or	(%r8), %r14
	or	(%r8), %r15
	nop
	or	(%r9), %rax
	or	(%r9), %rcx
	or	(%r9), %rdx
	or	(%r9), %rbx
	or	(%r9), %rsp
	or	(%r9), %rbp
	or	(%r9), %rsi
	or	(%r9), %rdi
	or	(%r9), %r8
	or	(%r9), %r9
	or	(%r9), %r10
	or	(%r9), %r11
	or	(%r9), %r12
	or	(%r9), %r13
	or	(%r9), %r14
	or	(%r9), %r15
	nop
	or	(%r10),%rax
	or	(%r10),%rcx
	or	(%r10),%rdx
	or	(%r10),%rbx
	or	(%r10),%rsp
	or	(%r10),%rbp
	or	(%r10),%rsi
	or	(%r10),%rdi
	or	(%r10),%r8
	or	(%r10),%r9
	or	(%r10),%r10
	or	(%r10),%r11
	or	(%r10),%r12
	or	(%r10),%r13
	or	(%r10),%r14
	or	(%r10),%r15
	nop
	or	(%r11),%rax
	or	(%r11),%rcx
	or	(%r11),%rdx
	or	(%r11),%rbx
	or	(%r11),%rsp
	or	(%r11),%rbp
	or	(%r11),%rsi
	or	(%r11),%rdi
	or	(%r11),%r8
	or	(%r11),%r9
	or	(%r11),%r10
	or	(%r11),%r11
	or	(%r11),%r12
	or	(%r11),%r13
	or	(%r11),%r14
	or	(%r11),%r15
	nop
	or	(%r12),%rax
	or	(%r12),%rcx
	or	(%r12),%rdx
	or	(%r12),%rbx
	or	(%r12),%rsp
	or	(%r12),%rbp
	or	(%r12),%rsi
	or	(%r12),%rdi
	or	(%r12),%r8
	or	(%r12),%r9
	or	(%r12),%r10
	or	(%r12),%r11
	or	(%r12),%r12
	or	(%r12),%r13
	or	(%r12),%r14
	or	(%r12),%r15
	nop
	or	(%r13),%rax
	or	(%r13),%rcx
	or	(%r13),%rdx
	or	(%r13),%rbx
	or	(%r13),%rsp
	or	(%r13),%rbp
	or	(%r13),%rsi
	or	(%r13),%rdi
	or	(%r13),%r8
	or	(%r13),%r9
	or	(%r13),%r10
	or	(%r13),%r11
	or	(%r13),%r12
	or	(%r13),%r13
	or	(%r13),%r14
	or	(%r13),%r15
	nop
	or	(%r14),%rax
	or	(%r14),%rcx
	or	(%r14),%rdx
	or	(%r14),%rbx
	or	(%r14),%rsp
	or	(%r14),%rbp
	or	(%r14),%rsi
	or	(%r14),%rdi
	or	(%r14),%r8
	or	(%r14),%r9
	or	(%r14),%r10
	or	(%r14),%r11
	or	(%r14),%r12
	or	(%r14),%r13
	or	(%r14),%r14
	or	(%r14),%r15
	nop
	or	(%r15),%rax
	or	(%r15),%rcx
	or	(%r15),%rdx
	or	(%r15),%rbx
	or	(%r15),%rsp
	or	(%r15),%rbp
	or	(%r15),%rsi
	or	(%r15),%rdi
	or	(%r15),%r8
	or	(%r15),%r9
	or	(%r15),%r10
	or	(%r15),%r11
	or	(%r15),%r12
	or	(%r15),%r13
	or	(%r15),%r14
	or	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	OrRegMem8
	.type	OrRegMem8, @function
OrRegMem8:
	.cfi_startproc
	or	0x7f(%rax),%rax
	or	0x7f(%rax),%rcx
	or	0x7f(%rax),%rdx
	or	0x7f(%rax),%rbx
	or	0x7f(%rax),%rsp
	or	0x7f(%rax),%rbp
	or	0x7f(%rax),%rsi
	or	0x7f(%rax),%rdi
	or	0x7f(%rax),%r8
	or	0x7f(%rax),%r9
	or	0x7f(%rax),%r10
	or	0x7f(%rax),%r11
	or	0x7f(%rax),%r12
	or	0x7f(%rax),%r13
	or	0x7f(%rax),%r14
	or	0x7f(%rax),%r15
	nop
	or	0x7f(%rcx),%rax
	or	0x7f(%rcx),%rcx
	or	0x7f(%rcx),%rdx
	or	0x7f(%rcx),%rbx
	or	0x7f(%rcx),%rsp
	or	0x7f(%rcx),%rbp
	or	0x7f(%rcx),%rsi
	or	0x7f(%rcx),%rdi
	or	0x7f(%rcx),%r8
	or	0x7f(%rcx),%r9
	or	0x7f(%rcx),%r10
	or	0x7f(%rcx),%r11
	or	0x7f(%rcx),%r12
	or	0x7f(%rcx),%r13
	or	0x7f(%rcx),%r14
	or	0x7f(%rcx),%r15
	nop
	or	0x7f(%rdx),%rax
	or	0x7f(%rdx),%rcx
	or	0x7f(%rdx),%rdx
	or	0x7f(%rdx),%rbx
	or	0x7f(%rdx),%rsp
	or	0x7f(%rdx),%rbp
	or	0x7f(%rdx),%rsi
	or	0x7f(%rdx),%rdi
	or	0x7f(%rdx),%r8
	or	0x7f(%rdx),%r9
	or	0x7f(%rdx),%r10
	or	0x7f(%rdx),%r11
	or	0x7f(%rdx),%r12
	or	0x7f(%rdx),%r13
	or	0x7f(%rdx),%r14
	or	0x7f(%rdx),%r15
	nop
	or	0x7f(%rbx),%rax
	or	0x7f(%rbx),%rcx
	or	0x7f(%rbx),%rdx
	or	0x7f(%rbx),%rbx
	or	0x7f(%rbx),%rsp
	or	0x7f(%rbx),%rbp
	or	0x7f(%rbx),%rsi
	or	0x7f(%rbx),%rdi
	or	0x7f(%rbx),%r8
	or	0x7f(%rbx),%r9
	or	0x7f(%rbx),%r10
	or	0x7f(%rbx),%r11
	or	0x7f(%rbx),%r12
	or	0x7f(%rbx),%r13
	or	0x7f(%rbx),%r14
	or	0x7f(%rbx),%r15
	nop
	or	0x7f(%rsp),%rax
	or	0x7f(%rsp),%rcx
	or	0x7f(%rsp),%rdx
	or	0x7f(%rsp),%rbx
	or	0x7f(%rsp),%rsp
	or	0x7f(%rsp),%rbp
	or	0x7f(%rsp),%rsi
	or	0x7f(%rsp),%rdi
	or	0x7f(%rsp),%r8
	or	0x7f(%rsp),%r9
	or	0x7f(%rsp),%r10
	or	0x7f(%rsp),%r11
	or	0x7f(%rsp),%r12
	or	0x7f(%rsp),%r13
	or	0x7f(%rsp),%r14
	or	0x7f(%rsp),%r15
	nop
	or	0x7f(%rbp),%rax
	or	0x7f(%rbp),%rcx
	or	0x7f(%rbp),%rdx
	or	0x7f(%rbp),%rbx
	or	0x7f(%rbp),%rsp
	or	0x7f(%rbp),%rbp
	or	0x7f(%rbp),%rsi
	or	0x7f(%rbp),%rdi
	or	0x7f(%rbp),%r8
	or	0x7f(%rbp),%r9
	or	0x7f(%rbp),%r10
	or	0x7f(%rbp),%r11
	or	0x7f(%rbp),%r12
	or	0x7f(%rbp),%r13
	or	0x7f(%rbp),%r14
	or	0x7f(%rbp),%r15
	nop
	or	0x7f(%rsi),%rax
	or	0x7f(%rsi),%rcx
	or	0x7f(%rsi),%rdx
	or	0x7f(%rsi),%rbx
	or	0x7f(%rsi),%rsp
	or	0x7f(%rsi),%rbp
	or	0x7f(%rsi),%rsi
	or	0x7f(%rsi),%rdi
	or	0x7f(%rsi),%r8
	or	0x7f(%rsi),%r9
	or	0x7f(%rsi),%r10
	or	0x7f(%rsi),%r11
	or	0x7f(%rsi),%r12
	or	0x7f(%rsi),%r13
	or	0x7f(%rsi),%r14
	or	0x7f(%rsi),%r15
	nop
	or	0x7f(%rdi),%rax
	or	0x7f(%rdi),%rcx
	or	0x7f(%rdi),%rdx
	or	0x7f(%rdi),%rbx
	or	0x7f(%rdi),%rsp
	or	0x7f(%rdi),%rbp
	or	0x7f(%rdi),%rsi
	or	0x7f(%rdi),%rdi
	or	0x7f(%rdi),%r8
	or	0x7f(%rdi),%r9
	or	0x7f(%rdi),%r10
	or	0x7f(%rdi),%r11
	or	0x7f(%rdi),%r12
	or	0x7f(%rdi),%r13
	or	0x7f(%rdi),%r14
	or	0x7f(%rdi),%r15
	nop
	or	0x7f(%r8), %rax
	or	0x7f(%r8), %rcx
	or	0x7f(%r8), %rdx
	or	0x7f(%r8), %rbx
	or	0x7f(%r8), %rsp
	or	0x7f(%r8), %rbp
	or	0x7f(%r8), %rsi
	or	0x7f(%r8), %rdi
	or	0x7f(%r8), %r8
	or	0x7f(%r8), %r9
	or	0x7f(%r8), %r10
	or	0x7f(%r8), %r11
	or	0x7f(%r8), %r12
	or	0x7f(%r8), %r13
	or	0x7f(%r8), %r14
	or	0x7f(%r8), %r15
	nop
	or	0x7f(%r9), %rax
	or	0x7f(%r9), %rcx
	or	0x7f(%r9), %rdx
	or	0x7f(%r9), %rbx
	or	0x7f(%r9), %rsp
	or	0x7f(%r9), %rbp
	or	0x7f(%r9), %rsi
	or	0x7f(%r9), %rdi
	or	0x7f(%r9), %r8
	or	0x7f(%r9), %r9
	or	0x7f(%r9), %r10
	or	0x7f(%r9), %r11
	or	0x7f(%r9), %r12
	or	0x7f(%r9), %r13
	or	0x7f(%r9), %r14
	or	0x7f(%r9), %r15
	nop
	or	0x7f(%r10),%rax
	or	0x7f(%r10),%rcx
	or	0x7f(%r10),%rdx
	or	0x7f(%r10),%rbx
	or	0x7f(%r10),%rsp
	or	0x7f(%r10),%rbp
	or	0x7f(%r10),%rsi
	or	0x7f(%r10),%rdi
	or	0x7f(%r10),%r8
	or	0x7f(%r10),%r9
	or	0x7f(%r10),%r10
	or	0x7f(%r10),%r11
	or	0x7f(%r10),%r12
	or	0x7f(%r10),%r13
	or	0x7f(%r10),%r14
	or	0x7f(%r10),%r15
	nop
	or	0x7f(%r11),%rax
	or	0x7f(%r11),%rcx
	or	0x7f(%r11),%rdx
	or	0x7f(%r11),%rbx
	or	0x7f(%r11),%rsp
	or	0x7f(%r11),%rbp
	or	0x7f(%r11),%rsi
	or	0x7f(%r11),%rdi
	or	0x7f(%r11),%r8
	or	0x7f(%r11),%r9
	or	0x7f(%r11),%r10
	or	0x7f(%r11),%r11
	or	0x7f(%r11),%r12
	or	0x7f(%r11),%r13
	or	0x7f(%r11),%r14
	or	0x7f(%r11),%r15
	nop
	or	0x7f(%r12),%rax
	or	0x7f(%r12),%rcx
	or	0x7f(%r12),%rdx
	or	0x7f(%r12),%rbx
	or	0x7f(%r12),%rsp
	or	0x7f(%r12),%rbp
	or	0x7f(%r12),%rsi
	or	0x7f(%r12),%rdi
	or	0x7f(%r12),%r8
	or	0x7f(%r12),%r9
	or	0x7f(%r12),%r10
	or	0x7f(%r12),%r11
	or	0x7f(%r12),%r12
	or	0x7f(%r12),%r13
	or	0x7f(%r12),%r14
	or	0x7f(%r12),%r15
	nop
	or	0x7f(%r13),%rax
	or	0x7f(%r13),%rcx
	or	0x7f(%r13),%rdx
	or	0x7f(%r13),%rbx
	or	0x7f(%r13),%rsp
	or	0x7f(%r13),%rbp
	or	0x7f(%r13),%rsi
	or	0x7f(%r13),%rdi
	or	0x7f(%r13),%r8
	or	0x7f(%r13),%r9
	or	0x7f(%r13),%r10
	or	0x7f(%r13),%r11
	or	0x7f(%r13),%r12
	or	0x7f(%r13),%r13
	or	0x7f(%r13),%r14
	or	0x7f(%r13),%r15
	nop
	or	0x7f(%r14),%rax
	or	0x7f(%r14),%rcx
	or	0x7f(%r14),%rdx
	or	0x7f(%r14),%rbx
	or	0x7f(%r14),%rsp
	or	0x7f(%r14),%rbp
	or	0x7f(%r14),%rsi
	or	0x7f(%r14),%rdi
	or	0x7f(%r14),%r8
	or	0x7f(%r14),%r9
	or	0x7f(%r14),%r10
	or	0x7f(%r14),%r11
	or	0x7f(%r14),%r12
	or	0x7f(%r14),%r13
	or	0x7f(%r14),%r14
	or	0x7f(%r14),%r15
	nop
	or	0x7f(%r15),%rax
	or	0x7f(%r15),%rcx
	or	0x7f(%r15),%rdx
	or	0x7f(%r15),%rbx
	or	0x7f(%r15),%rsp
	or	0x7f(%r15),%rbp
	or	0x7f(%r15),%rsi
	or	0x7f(%r15),%rdi
	or	0x7f(%r15),%r8
	or	0x7f(%r15),%r9
	or	0x7f(%r15),%r10
	or	0x7f(%r15),%r11
	or	0x7f(%r15),%r12
	or	0x7f(%r15),%r13
	or	0x7f(%r15),%r14
	or	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	OrRegMem32
	.type	OrRegMem32, @function
OrRegMem32:
	.cfi_startproc
	or	0x7f563412(%rax),%rax
	or	0x7f563412(%rax),%rcx
	or	0x7f563412(%rax),%rdx
	or	0x7f563412(%rax),%rbx
	or	0x7f563412(%rax),%rsp
	or	0x7f563412(%rax),%rbp
	or	0x7f563412(%rax),%rsi
	or	0x7f563412(%rax),%rdi
	or	0x7f563412(%rax),%r8
	or	0x7f563412(%rax),%r9
	or	0x7f563412(%rax),%r10
	or	0x7f563412(%rax),%r11
	or	0x7f563412(%rax),%r12
	or	0x7f563412(%rax),%r13
	or	0x7f563412(%rax),%r14
	or	0x7f563412(%rax),%r15
	nop
	or	0x7f563412(%rcx),%rax
	or	0x7f563412(%rcx),%rcx
	or	0x7f563412(%rcx),%rdx
	or	0x7f563412(%rcx),%rbx
	or	0x7f563412(%rcx),%rsp
	or	0x7f563412(%rcx),%rbp
	or	0x7f563412(%rcx),%rsi
	or	0x7f563412(%rcx),%rdi
	or	0x7f563412(%rcx),%r8
	or	0x7f563412(%rcx),%r9
	or	0x7f563412(%rcx),%r10
	or	0x7f563412(%rcx),%r11
	or	0x7f563412(%rcx),%r12
	or	0x7f563412(%rcx),%r13
	or	0x7f563412(%rcx),%r14
	or	0x7f563412(%rcx),%r15
	nop
	or	0x7f563412(%rdx),%rax
	or	0x7f563412(%rdx),%rcx
	or	0x7f563412(%rdx),%rdx
	or	0x7f563412(%rdx),%rbx
	or	0x7f563412(%rdx),%rsp
	or	0x7f563412(%rdx),%rbp
	or	0x7f563412(%rdx),%rsi
	or	0x7f563412(%rdx),%rdi
	or	0x7f563412(%rdx),%r8
	or	0x7f563412(%rdx),%r9
	or	0x7f563412(%rdx),%r10
	or	0x7f563412(%rdx),%r11
	or	0x7f563412(%rdx),%r12
	or	0x7f563412(%rdx),%r13
	or	0x7f563412(%rdx),%r14
	or	0x7f563412(%rdx),%r15
	nop
	or	0x7f563412(%rbx),%rax
	or	0x7f563412(%rbx),%rcx
	or	0x7f563412(%rbx),%rdx
	or	0x7f563412(%rbx),%rbx
	or	0x7f563412(%rbx),%rsp
	or	0x7f563412(%rbx),%rbp
	or	0x7f563412(%rbx),%rsi
	or	0x7f563412(%rbx),%rdi
	or	0x7f563412(%rbx),%r8
	or	0x7f563412(%rbx),%r9
	or	0x7f563412(%rbx),%r10
	or	0x7f563412(%rbx),%r11
	or	0x7f563412(%rbx),%r12
	or	0x7f563412(%rbx),%r13
	or	0x7f563412(%rbx),%r14
	or	0x7f563412(%rbx),%r15
	nop
	or	0x7f563412(%rsp),%rax
	or	0x7f563412(%rsp),%rcx
	or	0x7f563412(%rsp),%rdx
	or	0x7f563412(%rsp),%rbx
	or	0x7f563412(%rsp),%rsp
	or	0x7f563412(%rsp),%rbp
	or	0x7f563412(%rsp),%rsi
	or	0x7f563412(%rsp),%rdi
	or	0x7f563412(%rsp),%r8
	or	0x7f563412(%rsp),%r9
	or	0x7f563412(%rsp),%r10
	or	0x7f563412(%rsp),%r11
	or	0x7f563412(%rsp),%r12
	or	0x7f563412(%rsp),%r13
	or	0x7f563412(%rsp),%r14
	or	0x7f563412(%rsp),%r15
	nop
	or	0x7f563412(%rbp),%rax
	or	0x7f563412(%rbp),%rcx
	or	0x7f563412(%rbp),%rdx
	or	0x7f563412(%rbp),%rbx
	or	0x7f563412(%rbp),%rsp
	or	0x7f563412(%rbp),%rbp
	or	0x7f563412(%rbp),%rsi
	or	0x7f563412(%rbp),%rdi
	or	0x7f563412(%rbp),%r8
	or	0x7f563412(%rbp),%r9
	or	0x7f563412(%rbp),%r10
	or	0x7f563412(%rbp),%r11
	or	0x7f563412(%rbp),%r12
	or	0x7f563412(%rbp),%r13
	or	0x7f563412(%rbp),%r14
	or	0x7f563412(%rbp),%r15
	nop
	or	0x7f563412(%rsi),%rax
	or	0x7f563412(%rsi),%rcx
	or	0x7f563412(%rsi),%rdx
	or	0x7f563412(%rsi),%rbx
	or	0x7f563412(%rsi),%rsp
	or	0x7f563412(%rsi),%rbp
	or	0x7f563412(%rsi),%rsi
	or	0x7f563412(%rsi),%rdi
	or	0x7f563412(%rsi),%r8
	or	0x7f563412(%rsi),%r9
	or	0x7f563412(%rsi),%r10
	or	0x7f563412(%rsi),%r11
	or	0x7f563412(%rsi),%r12
	or	0x7f563412(%rsi),%r13
	or	0x7f563412(%rsi),%r14
	or	0x7f563412(%rsi),%r15
	nop
	or	0x7f563412(%rdi),%rax
	or	0x7f563412(%rdi),%rcx
	or	0x7f563412(%rdi),%rdx
	or	0x7f563412(%rdi),%rbx
	or	0x7f563412(%rdi),%rsp
	or	0x7f563412(%rdi),%rbp
	or	0x7f563412(%rdi),%rsi
	or	0x7f563412(%rdi),%rdi
	or	0x7f563412(%rdi),%r8
	or	0x7f563412(%rdi),%r9
	or	0x7f563412(%rdi),%r10
	or	0x7f563412(%rdi),%r11
	or	0x7f563412(%rdi),%r12
	or	0x7f563412(%rdi),%r13
	or	0x7f563412(%rdi),%r14
	or	0x7f563412(%rdi),%r15
	nop
	or	0x7f563412(%r8), %rax
	or	0x7f563412(%r8), %rcx
	or	0x7f563412(%r8), %rdx
	or	0x7f563412(%r8), %rbx
	or	0x7f563412(%r8), %rsp
	or	0x7f563412(%r8), %rbp
	or	0x7f563412(%r8), %rsi
	or	0x7f563412(%r8), %rdi
	or	0x7f563412(%r8), %r8
	or	0x7f563412(%r8), %r9
	or	0x7f563412(%r8), %r10
	or	0x7f563412(%r8), %r11
	or	0x7f563412(%r8), %r12
	or	0x7f563412(%r8), %r13
	or	0x7f563412(%r8), %r14
	or	0x7f563412(%r8), %r15
	nop
	or	0x7f563412(%r9), %rax
	or	0x7f563412(%r9), %rcx
	or	0x7f563412(%r9), %rdx
	or	0x7f563412(%r9), %rbx
	or	0x7f563412(%r9), %rsp
	or	0x7f563412(%r9), %rbp
	or	0x7f563412(%r9), %rsi
	or	0x7f563412(%r9), %rdi
	or	0x7f563412(%r9), %r8
	or	0x7f563412(%r9), %r9
	or	0x7f563412(%r9), %r10
	or	0x7f563412(%r9), %r11
	or	0x7f563412(%r9), %r12
	or	0x7f563412(%r9), %r13
	or	0x7f563412(%r9), %r14
	or	0x7f563412(%r9), %r15
	nop
	or	0x7f563412(%r10),%rax
	or	0x7f563412(%r10),%rcx
	or	0x7f563412(%r10),%rdx
	or	0x7f563412(%r10),%rbx
	or	0x7f563412(%r10),%rsp
	or	0x7f563412(%r10),%rbp
	or	0x7f563412(%r10),%rsi
	or	0x7f563412(%r10),%rdi
	or	0x7f563412(%r10),%r8
	or	0x7f563412(%r10),%r9
	or	0x7f563412(%r10),%r10
	or	0x7f563412(%r10),%r11
	or	0x7f563412(%r10),%r12
	or	0x7f563412(%r10),%r13
	or	0x7f563412(%r10),%r14
	or	0x7f563412(%r10),%r15
	nop
	or	0x7f563412(%r11),%rax
	or	0x7f563412(%r11),%rcx
	or	0x7f563412(%r11),%rdx
	or	0x7f563412(%r11),%rbx
	or	0x7f563412(%r11),%rsp
	or	0x7f563412(%r11),%rbp
	or	0x7f563412(%r11),%rsi
	or	0x7f563412(%r11),%rdi
	or	0x7f563412(%r11),%r8
	or	0x7f563412(%r11),%r9
	or	0x7f563412(%r11),%r10
	or	0x7f563412(%r11),%r11
	or	0x7f563412(%r11),%r12
	or	0x7f563412(%r11),%r13
	or	0x7f563412(%r11),%r14
	or	0x7f563412(%r11),%r15
	nop
	or	0x7f563412(%r12),%rax
	or	0x7f563412(%r12),%rcx
	or	0x7f563412(%r12),%rdx
	or	0x7f563412(%r12),%rbx
	or	0x7f563412(%r12),%rsp
	or	0x7f563412(%r12),%rbp
	or	0x7f563412(%r12),%rsi
	or	0x7f563412(%r12),%rdi
	or	0x7f563412(%r12),%r8
	or	0x7f563412(%r12),%r9
	or	0x7f563412(%r12),%r10
	or	0x7f563412(%r12),%r11
	or	0x7f563412(%r12),%r12
	or	0x7f563412(%r12),%r13
	or	0x7f563412(%r12),%r14
	or	0x7f563412(%r12),%r15
	nop
	or	0x7f563412(%r13),%rax
	or	0x7f563412(%r13),%rcx
	or	0x7f563412(%r13),%rdx
	or	0x7f563412(%r13),%rbx
	or	0x7f563412(%r13),%rsp
	or	0x7f563412(%r13),%rbp
	or	0x7f563412(%r13),%rsi
	or	0x7f563412(%r13),%rdi
	or	0x7f563412(%r13),%r8
	or	0x7f563412(%r13),%r9
	or	0x7f563412(%r13),%r10
	or	0x7f563412(%r13),%r11
	or	0x7f563412(%r13),%r12
	or	0x7f563412(%r13),%r13
	or	0x7f563412(%r13),%r14
	or	0x7f563412(%r13),%r15
	nop
	or	0x7f563412(%r14),%rax
	or	0x7f563412(%r14),%rcx
	or	0x7f563412(%r14),%rdx
	or	0x7f563412(%r14),%rbx
	or	0x7f563412(%r14),%rsp
	or	0x7f563412(%r14),%rbp
	or	0x7f563412(%r14),%rsi
	or	0x7f563412(%r14),%rdi
	or	0x7f563412(%r14),%r8
	or	0x7f563412(%r14),%r9
	or	0x7f563412(%r14),%r10
	or	0x7f563412(%r14),%r11
	or	0x7f563412(%r14),%r12
	or	0x7f563412(%r14),%r13
	or	0x7f563412(%r14),%r14
	or	0x7f563412(%r14),%r15
	nop
	or	0x7f563412(%r15),%rax
	or	0x7f563412(%r15),%rcx
	or	0x7f563412(%r15),%rdx
	or	0x7f563412(%r15),%rbx
	or	0x7f563412(%r15),%rsp
	or	0x7f563412(%r15),%rbp
	or	0x7f563412(%r15),%rsi
	or	0x7f563412(%r15),%rdi
	or	0x7f563412(%r15),%r8
	or	0x7f563412(%r15),%r9
	or	0x7f563412(%r15),%r10
	or	0x7f563412(%r15),%r11
	or	0x7f563412(%r15),%r12
	or	0x7f563412(%r15),%r13
	or	0x7f563412(%r15),%r14
	or	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


