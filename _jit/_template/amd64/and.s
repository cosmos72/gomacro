	.file	"and.s"
	.text

	.p2align 4,,15
	.globl	And_s32
	.type	And_s32, @function
And_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// and	$-0x55667788,%rax
	and	$-0x55667788,%rcx
	and	$-0x55667788,%rdx
	and	$-0x55667788,%rbx
	and	$-0x55667788,%rsp
	and	$-0x55667788,%rbp
	and	$-0x55667788,%rsi
	and	$-0x55667788,%rdi
	and	$-0x55667788,%r8
	and	$-0x55667788,%r9
	and	$-0x55667788,%r10
	and	$-0x55667788,%r11
	and	$-0x55667788,%r12
	and	$-0x55667788,%r13
	and	$-0x55667788,%r14
	and	$-0x55667788,%r15
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
	.globl	And
	.type	And, @function
And:
	.cfi_startproc

        // reg8 += reg8
	and	%al,%al
	and	%al,%cl
	and	%al,%dl
	and	%al,%bl
	and	%al,%spl
	and	%al,%bpl
	and	%al,%sil
	and	%al,%dil
	and	%al,%r8b
	and	%al,%r9b
	and	%al,%r10b
	and	%al,%r11b
	and	%al,%r12b
	and	%al,%r13b
	and	%al,%r14b
	and	%al,%r15b
	nop
	and	%cl,%al
	and	%cl,%cl
	and	%cl,%dl
	and	%cl,%bl
	and	%cl,%spl
	and	%cl,%bpl
	and	%cl,%sil
	and	%cl,%dil
	and	%cl,%r8b
	and	%cl,%r9b
	and	%cl,%r10b
	and	%cl,%r11b
	and	%cl,%r12b
	and	%cl,%r13b
	and	%cl,%r14b
	and	%cl,%r15b
	nop
	and	%dl,%al
	and	%dl,%cl
	and	%dl,%dl
	and	%dl,%bl
	and	%dl,%spl
	and	%dl,%bpl
	and	%dl,%sil
	and	%dl,%dil
	and	%dl,%r8b
	and	%dl,%r9b
	and	%dl,%r10b
	and	%dl,%r11b
	and	%dl,%r12b
	and	%dl,%r13b
	and	%dl,%r14b
	and	%dl,%r15b
	nop
	and	%bl,%al
	and	%bl,%cl
	and	%bl,%dl
	and	%bl,%bl
	and	%bl,%spl
	and	%bl,%bpl
	and	%bl,%sil
	and	%bl,%dil
	and	%bl,%r8b
	and	%bl,%r9b
	and	%bl,%r10b
	and	%bl,%r11b
	and	%bl,%r12b
	and	%bl,%r13b
	and	%bl,%r14b
	and	%bl,%r15b
	nop
	and	%spl,%al
	and	%spl,%cl
	and	%spl,%dl
	and	%spl,%bl
	and	%spl,%spl
	and	%spl,%bpl
	and	%spl,%sil
	and	%spl,%dil
	and	%spl,%r8b
	and	%spl,%r9b
	and	%spl,%r10b
	and	%spl,%r11b
	and	%spl,%r12b
	and	%spl,%r13b
	and	%spl,%r14b
	and	%spl,%r15b
	nop
	and	%bpl,%al
	and	%bpl,%cl
	and	%bpl,%dl
	and	%bpl,%bl
	and	%bpl,%spl
	and	%bpl,%bpl
	and	%bpl,%sil
	and	%bpl,%dil
	and	%bpl,%r8b
	and	%bpl,%r9b
	and	%bpl,%r10b
	and	%bpl,%r11b
	and	%bpl,%r12b
	and	%bpl,%r13b
	and	%bpl,%r14b
	and	%bpl,%r15b
	nop
	and	%sil,%al
	and	%sil,%cl
	and	%sil,%dl
	and	%sil,%bl
	and	%sil,%spl
	and	%sil,%bpl
	and	%sil,%sil
	and	%sil,%dil
	and	%sil,%r8b
	and	%sil,%r9b
	and	%sil,%r10b
	and	%sil,%r11b
	and	%sil,%r12b
	and	%sil,%r13b
	and	%sil,%r14b
	and	%sil,%r15b
	nop
	and	%dil,%al
	and	%dil,%cl
	and	%dil,%dl
	and	%dil,%bl
	and	%dil,%spl
	and	%dil,%bpl
	and	%dil,%sil
	and	%dil,%dil
	and	%dil,%r8b
	and	%dil,%r9b
	and	%dil,%r10b
	and	%dil,%r11b
	and	%dil,%r12b
	and	%dil,%r13b
	and	%dil,%r14b
	and	%dil,%r15b
	nop
	and	%r8b, %al
	and	%r8b, %cl
	and	%r8b, %dl
	and	%r8b, %bl
	and	%r8b, %spl
	and	%r8b, %bpl
	and	%r8b, %sil
	and	%r8b, %dil
	and	%r8b, %r8b
	and	%r8b, %r9b
	and	%r8b, %r10b
	and	%r8b, %r11b
	and	%r8b, %r12b
	and	%r8b, %r13b
	and	%r8b, %r14b
	and	%r8b, %r15b
	nop
	and	%r9b, %al
	and	%r9b, %cl
	and	%r9b, %dl
	and	%r9b, %bl
	and	%r9b, %spl
	and	%r9b, %bpl
	and	%r9b, %sil
	and	%r9b, %dil
	and	%r9b, %r8b
	and	%r9b, %r9b
	and	%r9b, %r10b
	and	%r9b, %r11b
	and	%r9b, %r12b
	and	%r9b, %r13b
	and	%r9b, %r14b
	and	%r9b, %r15b
	nop
	and	%r10b,%al
	and	%r10b,%cl
	and	%r10b,%dl
	and	%r10b,%bl
	and	%r10b,%spl
	and	%r10b,%bpl
	and	%r10b,%sil
	and	%r10b,%dil
	and	%r10b,%r8b
	and	%r10b,%r9b
	and	%r10b,%r10b
	and	%r10b,%r11b
	and	%r10b,%r12b
	and	%r10b,%r13b
	and	%r10b,%r14b
	and	%r10b,%r15b
	nop
	and	%r11b,%al
	and	%r11b,%cl
	and	%r11b,%dl
	and	%r11b,%bl
	and	%r11b,%spl
	and	%r11b,%bpl
	and	%r11b,%sil
	and	%r11b,%dil
	and	%r11b,%r8b
	and	%r11b,%r9b
	and	%r11b,%r10b
	and	%r11b,%r11b
	and	%r11b,%r12b
	and	%r11b,%r13b
	and	%r11b,%r14b
	and	%r11b,%r15b
	nop
	and	%r12b,%al
	and	%r12b,%cl
	and	%r12b,%dl
	and	%r12b,%bl
	and	%r12b,%spl
	and	%r12b,%bpl
	and	%r12b,%sil
	and	%r12b,%dil
	and	%r12b,%r8b
	and	%r12b,%r9b
	and	%r12b,%r10b
	and	%r12b,%r11b
	and	%r12b,%r12b
	and	%r12b,%r13b
	and	%r12b,%r14b
	and	%r12b,%r15b
	nop
	and	%r13b,%al
	and	%r13b,%cl
	and	%r13b,%dl
	and	%r13b,%bl
	and	%r13b,%spl
	and	%r13b,%bpl
	and	%r13b,%sil
	and	%r13b,%dil
	and	%r13b,%r8b
	and	%r13b,%r9b
	and	%r13b,%r10b
	and	%r13b,%r11b
	and	%r13b,%r12b
	and	%r13b,%r13b
	and	%r13b,%r14b
	and	%r13b,%r15b
	nop
	and	%r14b,%al
	and	%r14b,%cl
	and	%r14b,%dl
	and	%r14b,%bl
	and	%r14b,%spl
	and	%r14b,%bpl
	and	%r14b,%sil
	and	%r14b,%dil
	and	%r14b,%r8b
	and	%r14b,%r9b
	and	%r14b,%r10b
	and	%r14b,%r11b
	and	%r14b,%r12b
	and	%r14b,%r13b
	and	%r14b,%r14b
	and	%r14b,%r15b
	nop
	and	%r15b,%al
	and	%r15b,%cl
	and	%r15b,%dl
	and	%r15b,%bl
	and	%r15b,%spl
	and	%r15b,%bpl
	and	%r15b,%sil
	and	%r15b,%dil
	and	%r15b,%r8b
	and	%r15b,%r9b
	and	%r15b,%r10b
	and	%r15b,%r11b
	and	%r15b,%r12b
	and	%r15b,%r13b
	and	%r15b,%r14b
	and	%r15b,%r15b
        nop
        nop
        // reg16 += reg16
	and	%ax,%ax
	and	%ax,%cx
	and	%ax,%dx
	and	%ax,%bx
	and	%ax,%sp
	and	%ax,%bp
	and	%ax,%si
	and	%ax,%di
	and	%ax,%r8w
	and	%ax,%r9w
	and	%ax,%r10w
	and	%ax,%r11w
	and	%ax,%r12w
	and	%ax,%r13w
	and	%ax,%r14w
	and	%ax,%r15w
	nop
	and	%cx,%ax
	and	%cx,%cx
	and	%cx,%dx
	and	%cx,%bx
	and	%cx,%sp
	and	%cx,%bp
	and	%cx,%si
	and	%cx,%di
	and	%cx,%r8w
	and	%cx,%r9w
	and	%cx,%r10w
	and	%cx,%r11w
	and	%cx,%r12w
	and	%cx,%r13w
	and	%cx,%r14w
	and	%cx,%r15w
	nop
	and	%dx,%ax
	and	%dx,%cx
	and	%dx,%dx
	and	%dx,%bx
	and	%dx,%sp
	and	%dx,%bp
	and	%dx,%si
	and	%dx,%di
	and	%dx,%r8w
	and	%dx,%r9w
	and	%dx,%r10w
	and	%dx,%r11w
	and	%dx,%r12w
	and	%dx,%r13w
	and	%dx,%r14w
	and	%dx,%r15w
	nop
	and	%bx,%ax
	and	%bx,%cx
	and	%bx,%dx
	and	%bx,%bx
	and	%bx,%sp
	and	%bx,%bp
	and	%bx,%si
	and	%bx,%di
	and	%bx,%r8w
	and	%bx,%r9w
	and	%bx,%r10w
	and	%bx,%r11w
	and	%bx,%r12w
	and	%bx,%r13w
	and	%bx,%r14w
	and	%bx,%r15w
	nop
	and	%sp,%ax
	and	%sp,%cx
	and	%sp,%dx
	and	%sp,%bx
	and	%sp,%sp
	and	%sp,%bp
	and	%sp,%si
	and	%sp,%di
	and	%sp,%r8w
	and	%sp,%r9w
	and	%sp,%r10w
	and	%sp,%r11w
	and	%sp,%r12w
	and	%sp,%r13w
	and	%sp,%r14w
	and	%sp,%r15w
	nop
	and	%bp,%ax
	and	%bp,%cx
	and	%bp,%dx
	and	%bp,%bx
	and	%bp,%sp
	and	%bp,%bp
	and	%bp,%si
	and	%bp,%di
	and	%bp,%r8w
	and	%bp,%r9w
	and	%bp,%r10w
	and	%bp,%r11w
	and	%bp,%r12w
	and	%bp,%r13w
	and	%bp,%r14w
	and	%bp,%r15w
	nop
	and	%si,%ax
	and	%si,%cx
	and	%si,%dx
	and	%si,%bx
	and	%si,%sp
	and	%si,%bp
	and	%si,%si
	and	%si,%di
	and	%si,%r8w
	and	%si,%r9w
	and	%si,%r10w
	and	%si,%r11w
	and	%si,%r12w
	and	%si,%r13w
	and	%si,%r14w
	and	%si,%r15w
	nop
	and	%di,%ax
	and	%di,%cx
	and	%di,%dx
	and	%di,%bx
	and	%di,%sp
	and	%di,%bp
	and	%di,%si
	and	%di,%di
	and	%di,%r8w
	and	%di,%r9w
	and	%di,%r10w
	and	%di,%r11w
	and	%di,%r12w
	and	%di,%r13w
	and	%di,%r14w
	and	%di,%r15w
	nop
	and	%r8w, %ax
	and	%r8w, %cx
	and	%r8w, %dx
	and	%r8w, %bx
	and	%r8w, %sp
	and	%r8w, %bp
	and	%r8w, %si
	and	%r8w, %di
	and	%r8w, %r8w
	and	%r8w, %r9w
	and	%r8w, %r10w
	and	%r8w, %r11w
	and	%r8w, %r12w
	and	%r8w, %r13w
	and	%r8w, %r14w
	and	%r8w, %r15w
	nop
	and	%r9w, %ax
	and	%r9w, %cx
	and	%r9w, %dx
	and	%r9w, %bx
	and	%r9w, %sp
	and	%r9w, %bp
	and	%r9w, %si
	and	%r9w, %di
	and	%r9w, %r8w
	and	%r9w, %r9w
	and	%r9w, %r10w
	and	%r9w, %r11w
	and	%r9w, %r12w
	and	%r9w, %r13w
	and	%r9w, %r14w
	and	%r9w, %r15w
	nop
	and	%r10w,%ax
	and	%r10w,%cx
	and	%r10w,%dx
	and	%r10w,%bx
	and	%r10w,%sp
	and	%r10w,%bp
	and	%r10w,%si
	and	%r10w,%di
	and	%r10w,%r8w
	and	%r10w,%r9w
	and	%r10w,%r10w
	and	%r10w,%r11w
	and	%r10w,%r12w
	and	%r10w,%r13w
	and	%r10w,%r14w
	and	%r10w,%r15w
	nop
	and	%r11w,%ax
	and	%r11w,%cx
	and	%r11w,%dx
	and	%r11w,%bx
	and	%r11w,%sp
	and	%r11w,%bp
	and	%r11w,%si
	and	%r11w,%di
	and	%r11w,%r8w
	and	%r11w,%r9w
	and	%r11w,%r10w
	and	%r11w,%r11w
	and	%r11w,%r12w
	and	%r11w,%r13w
	and	%r11w,%r14w
	and	%r11w,%r15w
	nop
	and	%r12w,%ax
	and	%r12w,%cx
	and	%r12w,%dx
	and	%r12w,%bx
	and	%r12w,%sp
	and	%r12w,%bp
	and	%r12w,%si
	and	%r12w,%di
	and	%r12w,%r8w
	and	%r12w,%r9w
	and	%r12w,%r10w
	and	%r12w,%r11w
	and	%r12w,%r12w
	and	%r12w,%r13w
	and	%r12w,%r14w
	and	%r12w,%r15w
	nop
	and	%r13w,%ax
	and	%r13w,%cx
	and	%r13w,%dx
	and	%r13w,%bx
	and	%r13w,%sp
	and	%r13w,%bp
	and	%r13w,%si
	and	%r13w,%di
	and	%r13w,%r8w
	and	%r13w,%r9w
	and	%r13w,%r10w
	and	%r13w,%r11w
	and	%r13w,%r12w
	and	%r13w,%r13w
	and	%r13w,%r14w
	and	%r13w,%r15w
	nop
	and	%r14w,%ax
	and	%r14w,%cx
	and	%r14w,%dx
	and	%r14w,%bx
	and	%r14w,%sp
	and	%r14w,%bp
	and	%r14w,%si
	and	%r14w,%di
	and	%r14w,%r8w
	and	%r14w,%r9w
	and	%r14w,%r10w
	and	%r14w,%r11w
	and	%r14w,%r12w
	and	%r14w,%r13w
	and	%r14w,%r14w
	and	%r14w,%r15w
	nop
	and	%r15w,%ax
	and	%r15w,%cx
	and	%r15w,%dx
	and	%r15w,%bx
	and	%r15w,%sp
	and	%r15w,%bp
	and	%r15w,%si
	and	%r15w,%di
	and	%r15w,%r8w
	and	%r15w,%r9w
	and	%r15w,%r10w
	and	%r15w,%r11w
	and	%r15w,%r12w
	and	%r15w,%r13w
	and	%r15w,%r14w
	and	%r15w,%r15w
        nop
        nop
        // reg32 += reg32
	and	%eax,%eax
	and	%eax,%ecx
	and	%eax,%edx
	and	%eax,%ebx
	and	%eax,%esp
	and	%eax,%ebp
	and	%eax,%esi
	and	%eax,%edi
	and	%eax,%r8d
	and	%eax,%r9d
	and	%eax,%r10d
	and	%eax,%r11d
	and	%eax,%r12d
	and	%eax,%r13d
	and	%eax,%r14d
	and	%eax,%r15d
	nop
	and	%ecx,%eax
	and	%ecx,%ecx
	and	%ecx,%edx
	and	%ecx,%ebx
	and	%ecx,%esp
	and	%ecx,%ebp
	and	%ecx,%esi
	and	%ecx,%edi
	and	%ecx,%r8d
	and	%ecx,%r9d
	and	%ecx,%r10d
	and	%ecx,%r11d
	and	%ecx,%r12d
	and	%ecx,%r13d
	and	%ecx,%r14d
	and	%ecx,%r15d
	nop
	and	%edx,%eax
	and	%edx,%ecx
	and	%edx,%edx
	and	%edx,%ebx
	and	%edx,%esp
	and	%edx,%ebp
	and	%edx,%esi
	and	%edx,%edi
	and	%edx,%r8d
	and	%edx,%r9d
	and	%edx,%r10d
	and	%edx,%r11d
	and	%edx,%r12d
	and	%edx,%r13d
	and	%edx,%r14d
	and	%edx,%r15d
	nop
	and	%ebx,%eax
	and	%ebx,%ecx
	and	%ebx,%edx
	and	%ebx,%ebx
	and	%ebx,%esp
	and	%ebx,%ebp
	and	%ebx,%esi
	and	%ebx,%edi
	and	%ebx,%r8d
	and	%ebx,%r9d
	and	%ebx,%r10d
	and	%ebx,%r11d
	and	%ebx,%r12d
	and	%ebx,%r13d
	and	%ebx,%r14d
	and	%ebx,%r15d
	nop
	and	%esp,%eax
	and	%esp,%ecx
	and	%esp,%edx
	and	%esp,%ebx
	and	%esp,%esp
	and	%esp,%ebp
	and	%esp,%esi
	and	%esp,%edi
	and	%esp,%r8d
	and	%esp,%r9d
	and	%esp,%r10d
	and	%esp,%r11d
	and	%esp,%r12d
	and	%esp,%r13d
	and	%esp,%r14d
	and	%esp,%r15d
	nop
	and	%ebp,%eax
	and	%ebp,%ecx
	and	%ebp,%edx
	and	%ebp,%ebx
	and	%ebp,%esp
	and	%ebp,%ebp
	and	%ebp,%esi
	and	%ebp,%edi
	and	%ebp,%r8d
	and	%ebp,%r9d
	and	%ebp,%r10d
	and	%ebp,%r11d
	and	%ebp,%r12d
	and	%ebp,%r13d
	and	%ebp,%r14d
	and	%ebp,%r15d
	nop
	and	%esi,%eax
	and	%esi,%ecx
	and	%esi,%edx
	and	%esi,%ebx
	and	%esi,%esp
	and	%esi,%ebp
	and	%esi,%esi
	and	%esi,%edi
	and	%esi,%r8d
	and	%esi,%r9d
	and	%esi,%r10d
	and	%esi,%r11d
	and	%esi,%r12d
	and	%esi,%r13d
	and	%esi,%r14d
	and	%esi,%r15d
	nop
	and	%edi,%eax
	and	%edi,%ecx
	and	%edi,%edx
	and	%edi,%ebx
	and	%edi,%esp
	and	%edi,%ebp
	and	%edi,%esi
	and	%edi,%edi
	and	%edi,%r8d
	and	%edi,%r9d
	and	%edi,%r10d
	and	%edi,%r11d
	and	%edi,%r12d
	and	%edi,%r13d
	and	%edi,%r14d
	and	%edi,%r15d
	nop
	and	%r8d, %eax
	and	%r8d, %ecx
	and	%r8d, %edx
	and	%r8d, %ebx
	and	%r8d, %esp
	and	%r8d, %ebp
	and	%r8d, %esi
	and	%r8d, %edi
	and	%r8d, %r8d
	and	%r8d, %r9d
	and	%r8d, %r10d
	and	%r8d, %r11d
	and	%r8d, %r12d
	and	%r8d, %r13d
	and	%r8d, %r14d
	and	%r8d, %r15d
	nop
	and	%r9d, %eax
	and	%r9d, %ecx
	and	%r9d, %edx
	and	%r9d, %ebx
	and	%r9d, %esp
	and	%r9d, %ebp
	and	%r9d, %esi
	and	%r9d, %edi
	and	%r9d, %r8d
	and	%r9d, %r9d
	and	%r9d, %r10d
	and	%r9d, %r11d
	and	%r9d, %r12d
	and	%r9d, %r13d
	and	%r9d, %r14d
	and	%r9d, %r15d
	nop
	and	%r10d,%eax
	and	%r10d,%ecx
	and	%r10d,%edx
	and	%r10d,%ebx
	and	%r10d,%esp
	and	%r10d,%ebp
	and	%r10d,%esi
	and	%r10d,%edi
	and	%r10d,%r8d
	and	%r10d,%r9d
	and	%r10d,%r10d
	and	%r10d,%r11d
	and	%r10d,%r12d
	and	%r10d,%r13d
	and	%r10d,%r14d
	and	%r10d,%r15d
	nop
	and	%r11d,%eax
	and	%r11d,%ecx
	and	%r11d,%edx
	and	%r11d,%ebx
	and	%r11d,%esp
	and	%r11d,%ebp
	and	%r11d,%esi
	and	%r11d,%edi
	and	%r11d,%r8d
	and	%r11d,%r9d
	and	%r11d,%r10d
	and	%r11d,%r11d
	and	%r11d,%r12d
	and	%r11d,%r13d
	and	%r11d,%r14d
	and	%r11d,%r15d
	nop
	and	%r12d,%eax
	and	%r12d,%ecx
	and	%r12d,%edx
	and	%r12d,%ebx
	and	%r12d,%esp
	and	%r12d,%ebp
	and	%r12d,%esi
	and	%r12d,%edi
	and	%r12d,%r8d
	and	%r12d,%r9d
	and	%r12d,%r10d
	and	%r12d,%r11d
	and	%r12d,%r12d
	and	%r12d,%r13d
	and	%r12d,%r14d
	and	%r12d,%r15d
	nop
	and	%r13d,%eax
	and	%r13d,%ecx
	and	%r13d,%edx
	and	%r13d,%ebx
	and	%r13d,%esp
	and	%r13d,%ebp
	and	%r13d,%esi
	and	%r13d,%edi
	and	%r13d,%r8d
	and	%r13d,%r9d
	and	%r13d,%r10d
	and	%r13d,%r11d
	and	%r13d,%r12d
	and	%r13d,%r13d
	and	%r13d,%r14d
	and	%r13d,%r15d
	nop
	and	%r14d,%eax
	and	%r14d,%ecx
	and	%r14d,%edx
	and	%r14d,%ebx
	and	%r14d,%esp
	and	%r14d,%ebp
	and	%r14d,%esi
	and	%r14d,%edi
	and	%r14d,%r8d
	and	%r14d,%r9d
	and	%r14d,%r10d
	and	%r14d,%r11d
	and	%r14d,%r12d
	and	%r14d,%r13d
	and	%r14d,%r14d
	and	%r14d,%r15d
	nop
	and	%r15d,%eax
	and	%r15d,%ecx
	and	%r15d,%edx
	and	%r15d,%ebx
	and	%r15d,%esp
	and	%r15d,%ebp
	and	%r15d,%esi
	and	%r15d,%edi
	and	%r15d,%r8d
	and	%r15d,%r9d
	and	%r15d,%r10d
	and	%r15d,%r11d
	and	%r15d,%r12d
	and	%r15d,%r13d
	and	%r15d,%r14d
	and	%r15d,%r15d
        nop
        nop
        // reg64 += reg64
	and	%rax,%rax
	and	%rax,%rcx
	and	%rax,%rdx
	and	%rax,%rbx
	and	%rax,%rsp
	and	%rax,%rbp
	and	%rax,%rsi
	and	%rax,%rdi
	and	%rax,%r8
	and	%rax,%r9
	and	%rax,%r10
	and	%rax,%r11
	and	%rax,%r12
	and	%rax,%r13
	and	%rax,%r14
	and	%rax,%r15
	nop
	and	%rcx,%rax
	and	%rcx,%rcx
	and	%rcx,%rdx
	and	%rcx,%rbx
	and	%rcx,%rsp
	and	%rcx,%rbp
	and	%rcx,%rsi
	and	%rcx,%rdi
	and	%rcx,%r8
	and	%rcx,%r9
	and	%rcx,%r10
	and	%rcx,%r11
	and	%rcx,%r12
	and	%rcx,%r13
	and	%rcx,%r14
	and	%rcx,%r15
	nop
	and	%rdx,%rax
	and	%rdx,%rcx
	and	%rdx,%rdx
	and	%rdx,%rbx
	and	%rdx,%rsp
	and	%rdx,%rbp
	and	%rdx,%rsi
	and	%rdx,%rdi
	and	%rdx,%r8
	and	%rdx,%r9
	and	%rdx,%r10
	and	%rdx,%r11
	and	%rdx,%r12
	and	%rdx,%r13
	and	%rdx,%r14
	and	%rdx,%r15
	nop
	and	%rbx,%rax
	and	%rbx,%rcx
	and	%rbx,%rdx
	and	%rbx,%rbx
	and	%rbx,%rsp
	and	%rbx,%rbp
	and	%rbx,%rsi
	and	%rbx,%rdi
	and	%rbx,%r8
	and	%rbx,%r9
	and	%rbx,%r10
	and	%rbx,%r11
	and	%rbx,%r12
	and	%rbx,%r13
	and	%rbx,%r14
	and	%rbx,%r15
	nop
	and	%rsp,%rax
	and	%rsp,%rcx
	and	%rsp,%rdx
	and	%rsp,%rbx
	and	%rsp,%rsp
	and	%rsp,%rbp
	and	%rsp,%rsi
	and	%rsp,%rdi
	and	%rsp,%r8
	and	%rsp,%r9
	and	%rsp,%r10
	and	%rsp,%r11
	and	%rsp,%r12
	and	%rsp,%r13
	and	%rsp,%r14
	and	%rsp,%r15
	nop
	and	%rbp,%rax
	and	%rbp,%rcx
	and	%rbp,%rdx
	and	%rbp,%rbx
	and	%rbp,%rsp
	and	%rbp,%rbp
	and	%rbp,%rsi
	and	%rbp,%rdi
	and	%rbp,%r8
	and	%rbp,%r9
	and	%rbp,%r10
	and	%rbp,%r11
	and	%rbp,%r12
	and	%rbp,%r13
	and	%rbp,%r14
	and	%rbp,%r15
	nop
	and	%rsi,%rax
	and	%rsi,%rcx
	and	%rsi,%rdx
	and	%rsi,%rbx
	and	%rsi,%rsp
	and	%rsi,%rbp
	and	%rsi,%rsi
	and	%rsi,%rdi
	and	%rsi,%r8
	and	%rsi,%r9
	and	%rsi,%r10
	and	%rsi,%r11
	and	%rsi,%r12
	and	%rsi,%r13
	and	%rsi,%r14
	and	%rsi,%r15
	nop
	and	%rdi,%rax
	and	%rdi,%rcx
	and	%rdi,%rdx
	and	%rdi,%rbx
	and	%rdi,%rsp
	and	%rdi,%rbp
	and	%rdi,%rsi
	and	%rdi,%rdi
	and	%rdi,%r8
	and	%rdi,%r9
	and	%rdi,%r10
	and	%rdi,%r11
	and	%rdi,%r12
	and	%rdi,%r13
	and	%rdi,%r14
	and	%rdi,%r15
	nop
	and	%r8, %rax
	and	%r8, %rcx
	and	%r8, %rdx
	and	%r8, %rbx
	and	%r8, %rsp
	and	%r8, %rbp
	and	%r8, %rsi
	and	%r8, %rdi
	and	%r8, %r8
	and	%r8, %r9
	and	%r8, %r10
	and	%r8, %r11
	and	%r8, %r12
	and	%r8, %r13
	and	%r8, %r14
	and	%r8, %r15
	nop
	and	%r9, %rax
	and	%r9, %rcx
	and	%r9, %rdx
	and	%r9, %rbx
	and	%r9, %rsp
	and	%r9, %rbp
	and	%r9, %rsi
	and	%r9, %rdi
	and	%r9, %r8
	and	%r9, %r9
	and	%r9, %r10
	and	%r9, %r11
	and	%r9, %r12
	and	%r9, %r13
	and	%r9, %r14
	and	%r9, %r15
	nop
	and	%r10,%rax
	and	%r10,%rcx
	and	%r10,%rdx
	and	%r10,%rbx
	and	%r10,%rsp
	and	%r10,%rbp
	and	%r10,%rsi
	and	%r10,%rdi
	and	%r10,%r8
	and	%r10,%r9
	and	%r10,%r10
	and	%r10,%r11
	and	%r10,%r12
	and	%r10,%r13
	and	%r10,%r14
	and	%r10,%r15
	nop
	and	%r11,%rax
	and	%r11,%rcx
	and	%r11,%rdx
	and	%r11,%rbx
	and	%r11,%rsp
	and	%r11,%rbp
	and	%r11,%rsi
	and	%r11,%rdi
	and	%r11,%r8
	and	%r11,%r9
	and	%r11,%r10
	and	%r11,%r11
	and	%r11,%r12
	and	%r11,%r13
	and	%r11,%r14
	and	%r11,%r15
	nop
	and	%r12,%rax
	and	%r12,%rcx
	and	%r12,%rdx
	and	%r12,%rbx
	and	%r12,%rsp
	and	%r12,%rbp
	and	%r12,%rsi
	and	%r12,%rdi
	and	%r12,%r8
	and	%r12,%r9
	and	%r12,%r10
	and	%r12,%r11
	and	%r12,%r12
	and	%r12,%r13
	and	%r12,%r14
	and	%r12,%r15
	nop
	and	%r13,%rax
	and	%r13,%rcx
	and	%r13,%rdx
	and	%r13,%rbx
	and	%r13,%rsp
	and	%r13,%rbp
	and	%r13,%rsi
	and	%r13,%rdi
	and	%r13,%r8
	and	%r13,%r9
	and	%r13,%r10
	and	%r13,%r11
	and	%r13,%r12
	and	%r13,%r13
	and	%r13,%r14
	and	%r13,%r15
	nop
	and	%r14,%rax
	and	%r14,%rcx
	and	%r14,%rdx
	and	%r14,%rbx
	and	%r14,%rsp
	and	%r14,%rbp
	and	%r14,%rsi
	and	%r14,%rdi
	and	%r14,%r8
	and	%r14,%r9
	and	%r14,%r10
	and	%r14,%r11
	and	%r14,%r12
	and	%r14,%r13
	and	%r14,%r14
	and	%r14,%r15
	nop
	and	%r15,%rax
	and	%r15,%rcx
	and	%r15,%rdx
	and	%r15,%rbx
	and	%r15,%rsp
	and	%r15,%rbp
	and	%r15,%rsi
	and	%r15,%rdi
	and	%r15,%r8
	and	%r15,%r9
	and	%r15,%r10
	and	%r15,%r11
	and	%r15,%r12
	and	%r15,%r13
	and	%r15,%r14
	and	%r15,%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndMemReg
	.type	AndMemReg, @function
AndMemReg:
	.cfi_startproc
        // mem8 += reg8
	and	%al,(%rax)
	and	%al,(%rcx)
	and	%al,(%rdx)
	and	%al,(%rbx)
	and	%al,(%rsp)
	and	%al,(%rbp)
	and	%al,(%rsi)
	and	%al,(%rdi)
	and	%al,(%r8)
	and	%al,(%r9)
	and	%al,(%r10)
	and	%al,(%r11)
	and	%al,(%r12)
	and	%al,(%r13)
	and	%al,(%r14)
	and	%al,(%r15)
	nop
	and	%cl,(%rax)
	and	%cl,(%rcx)
	and	%cl,(%rdx)
	and	%cl,(%rbx)
	and	%cl,(%rsp)
	and	%cl,(%rbp)
	and	%cl,(%rsi)
	and	%cl,(%rdi)
	and	%cl,(%r8)
	and	%cl,(%r9)
	and	%cl,(%r10)
	and	%cl,(%r11)
	and	%cl,(%r12)
	and	%cl,(%r13)
	and	%cl,(%r14)
	and	%cl,(%r15)
	nop
	and	%dl,(%rax)
	and	%dl,(%rcx)
	and	%dl,(%rdx)
	and	%dl,(%rbx)
	and	%dl,(%rsp)
	and	%dl,(%rbp)
	and	%dl,(%rsi)
	and	%dl,(%rdi)
	and	%dl,(%r8)
	and	%dl,(%r9)
	and	%dl,(%r10)
	and	%dl,(%r11)
	and	%dl,(%r12)
	and	%dl,(%r13)
	and	%dl,(%r14)
	and	%dl,(%r15)
	nop
	and	%bl,(%rax)
	and	%bl,(%rcx)
	and	%bl,(%rdx)
	and	%bl,(%rbx)
	and	%bl,(%rsp)
	and	%bl,(%rbp)
	and	%bl,(%rsi)
	and	%bl,(%rdi)
	and	%bl,(%r8)
	and	%bl,(%r9)
	and	%bl,(%r10)
	and	%bl,(%r11)
	and	%bl,(%r12)
	and	%bl,(%r13)
	and	%bl,(%r14)
	and	%bl,(%r15)
	nop
	and	%spl,(%rax)
	and	%spl,(%rcx)
	and	%spl,(%rdx)
	and	%spl,(%rbx)
	and	%spl,(%rsp)
	and	%spl,(%rbp)
	and	%spl,(%rsi)
	and	%spl,(%rdi)
	and	%spl,(%r8)
	and	%spl,(%r9)
	and	%spl,(%r10)
	and	%spl,(%r11)
	and	%spl,(%r12)
	and	%spl,(%r13)
	and	%spl,(%r14)
	and	%spl,(%r15)
	nop
	and	%bpl,(%rax)
	and	%bpl,(%rcx)
	and	%bpl,(%rdx)
	and	%bpl,(%rbx)
	and	%bpl,(%rsp)
	and	%bpl,(%rbp)
	and	%bpl,(%rsi)
	and	%bpl,(%rdi)
	and	%bpl,(%r8)
	and	%bpl,(%r9)
	and	%bpl,(%r10)
	and	%bpl,(%r11)
	and	%bpl,(%r12)
	and	%bpl,(%r13)
	and	%bpl,(%r14)
	and	%bpl,(%r15)
	nop
	and	%sil,(%rax)
	and	%sil,(%rcx)
	and	%sil,(%rdx)
	and	%sil,(%rbx)
	and	%sil,(%rsp)
	and	%sil,(%rbp)
	and	%sil,(%rsi)
	and	%sil,(%rdi)
	and	%sil,(%r8)
	and	%sil,(%r9)
	and	%sil,(%r10)
	and	%sil,(%r11)
	and	%sil,(%r12)
	and	%sil,(%r13)
	and	%sil,(%r14)
	and	%sil,(%r15)
	nop
	and	%dil,(%rax)
	and	%dil,(%rcx)
	and	%dil,(%rdx)
	and	%dil,(%rbx)
	and	%dil,(%rsp)
	and	%dil,(%rbp)
	and	%dil,(%rsi)
	and	%dil,(%rdi)
	and	%dil,(%r8)
	and	%dil,(%r9)
	and	%dil,(%r10)
	and	%dil,(%r11)
	and	%dil,(%r12)
	and	%dil,(%r13)
	and	%dil,(%r14)
	and	%dil,(%r15)
	nop
	and	%r8b, (%rax)
	and	%r8b, (%rcx)
	and	%r8b, (%rdx)
	and	%r8b, (%rbx)
	and	%r8b, (%rsp)
	and	%r8b, (%rbp)
	and	%r8b, (%rsi)
	and	%r8b, (%rdi)
	and	%r8b, (%r8)
	and	%r8b, (%r9)
	and	%r8b, (%r10)
	and	%r8b, (%r11)
	and	%r8b, (%r12)
	and	%r8b, (%r13)
	and	%r8b, (%r14)
	and	%r8b, (%r15)
	nop
	and	%r9b, (%rax)
	and	%r9b, (%rcx)
	and	%r9b, (%rdx)
	and	%r9b, (%rbx)
	and	%r9b, (%rsp)
	and	%r9b, (%rbp)
	and	%r9b, (%rsi)
	and	%r9b, (%rdi)
	and	%r9b, (%r8)
	and	%r9b, (%r9)
	and	%r9b, (%r10)
	and	%r9b, (%r11)
	and	%r9b, (%r12)
	and	%r9b, (%r13)
	and	%r9b, (%r14)
	and	%r9b, (%r15)
	nop
	and	%r10b,(%rax)
	and	%r10b,(%rcx)
	and	%r10b,(%rdx)
	and	%r10b,(%rbx)
	and	%r10b,(%rsp)
	and	%r10b,(%rbp)
	and	%r10b,(%rsi)
	and	%r10b,(%rdi)
	and	%r10b,(%r8)
	and	%r10b,(%r9)
	and	%r10b,(%r10)
	and	%r10b,(%r11)
	and	%r10b,(%r12)
	and	%r10b,(%r13)
	and	%r10b,(%r14)
	and	%r10b,(%r15)
	nop
	and	%r11b,(%rax)
	and	%r11b,(%rcx)
	and	%r11b,(%rdx)
	and	%r11b,(%rbx)
	and	%r11b,(%rsp)
	and	%r11b,(%rbp)
	and	%r11b,(%rsi)
	and	%r11b,(%rdi)
	and	%r11b,(%r8)
	and	%r11b,(%r9)
	and	%r11b,(%r10)
	and	%r11b,(%r11)
	and	%r11b,(%r12)
	and	%r11b,(%r13)
	and	%r11b,(%r14)
	and	%r11b,(%r15)
	nop
	and	%r12b,(%rax)
	and	%r12b,(%rcx)
	and	%r12b,(%rdx)
	and	%r12b,(%rbx)
	and	%r12b,(%rsp)
	and	%r12b,(%rbp)
	and	%r12b,(%rsi)
	and	%r12b,(%rdi)
	and	%r12b,(%r8)
	and	%r12b,(%r9)
	and	%r12b,(%r10)
	and	%r12b,(%r11)
	and	%r12b,(%r12)
	and	%r12b,(%r13)
	and	%r12b,(%r14)
	and	%r12b,(%r15)
	nop
	and	%r13b,(%rax)
	and	%r13b,(%rcx)
	and	%r13b,(%rdx)
	and	%r13b,(%rbx)
	and	%r13b,(%rsp)
	and	%r13b,(%rbp)
	and	%r13b,(%rsi)
	and	%r13b,(%rdi)
	and	%r13b,(%r8)
	and	%r13b,(%r9)
	and	%r13b,(%r10)
	and	%r13b,(%r11)
	and	%r13b,(%r12)
	and	%r13b,(%r13)
	and	%r13b,(%r14)
	and	%r13b,(%r15)
	nop
	and	%r14b,(%rax)
	and	%r14b,(%rcx)
	and	%r14b,(%rdx)
	and	%r14b,(%rbx)
	and	%r14b,(%rsp)
	and	%r14b,(%rbp)
	and	%r14b,(%rsi)
	and	%r14b,(%rdi)
	and	%r14b,(%r8)
	and	%r14b,(%r9)
	and	%r14b,(%r10)
	and	%r14b,(%r11)
	and	%r14b,(%r12)
	and	%r14b,(%r13)
	and	%r14b,(%r14)
	and	%r14b,(%r15)
	nop
	and	%r15b,(%rax)
	and	%r15b,(%rcx)
	and	%r15b,(%rdx)
	and	%r15b,(%rbx)
	and	%r15b,(%rsp)
	and	%r15b,(%rbp)
	and	%r15b,(%rsi)
	and	%r15b,(%rdi)
	and	%r15b,(%r8)
	and	%r15b,(%r9)
	and	%r15b,(%r10)
	and	%r15b,(%r11)
	and	%r15b,(%r12)
	and	%r15b,(%r13)
	and	%r15b,(%r14)
	and	%r15b,(%r15)
        nop
        nop
        // mem16 += reg16
	and	%ax,(%rax)
	and	%ax,(%rcx)
	and	%ax,(%rdx)
	and	%ax,(%rbx)
	and	%ax,(%rsp)
	and	%ax,(%rbp)
	and	%ax,(%rsi)
	and	%ax,(%rdi)
	and	%ax,(%r8)
	and	%ax,(%r9)
	and	%ax,(%r10)
	and	%ax,(%r11)
	and	%ax,(%r12)
	and	%ax,(%r13)
	and	%ax,(%r14)
	and	%ax,(%r15)
	nop
	and	%cx,(%rax)
	and	%cx,(%rcx)
	and	%cx,(%rdx)
	and	%cx,(%rbx)
	and	%cx,(%rsp)
	and	%cx,(%rbp)
	and	%cx,(%rsi)
	and	%cx,(%rdi)
	and	%cx,(%r8)
	and	%cx,(%r9)
	and	%cx,(%r10)
	and	%cx,(%r11)
	and	%cx,(%r12)
	and	%cx,(%r13)
	and	%cx,(%r14)
	and	%cx,(%r15)
	nop
	and	%dx,(%rax)
	and	%dx,(%rcx)
	and	%dx,(%rdx)
	and	%dx,(%rbx)
	and	%dx,(%rsp)
	and	%dx,(%rbp)
	and	%dx,(%rsi)
	and	%dx,(%rdi)
	and	%dx,(%r8)
	and	%dx,(%r9)
	and	%dx,(%r10)
	and	%dx,(%r11)
	and	%dx,(%r12)
	and	%dx,(%r13)
	and	%dx,(%r14)
	and	%dx,(%r15)
	nop
	and	%bx,(%rax)
	and	%bx,(%rcx)
	and	%bx,(%rdx)
	and	%bx,(%rbx)
	and	%bx,(%rsp)
	and	%bx,(%rbp)
	and	%bx,(%rsi)
	and	%bx,(%rdi)
	and	%bx,(%r8)
	and	%bx,(%r9)
	and	%bx,(%r10)
	and	%bx,(%r11)
	and	%bx,(%r12)
	and	%bx,(%r13)
	and	%bx,(%r14)
	and	%bx,(%r15)
	nop
	and	%sp,(%rax)
	and	%sp,(%rcx)
	and	%sp,(%rdx)
	and	%sp,(%rbx)
	and	%sp,(%rsp)
	and	%sp,(%rbp)
	and	%sp,(%rsi)
	and	%sp,(%rdi)
	and	%sp,(%r8)
	and	%sp,(%r9)
	and	%sp,(%r10)
	and	%sp,(%r11)
	and	%sp,(%r12)
	and	%sp,(%r13)
	and	%sp,(%r14)
	and	%sp,(%r15)
	nop
	and	%bp,(%rax)
	and	%bp,(%rcx)
	and	%bp,(%rdx)
	and	%bp,(%rbx)
	and	%bp,(%rsp)
	and	%bp,(%rbp)
	and	%bp,(%rsi)
	and	%bp,(%rdi)
	and	%bp,(%r8)
	and	%bp,(%r9)
	and	%bp,(%r10)
	and	%bp,(%r11)
	and	%bp,(%r12)
	and	%bp,(%r13)
	and	%bp,(%r14)
	and	%bp,(%r15)
	nop
	and	%si,(%rax)
	and	%si,(%rcx)
	and	%si,(%rdx)
	and	%si,(%rbx)
	and	%si,(%rsp)
	and	%si,(%rbp)
	and	%si,(%rsi)
	and	%si,(%rdi)
	and	%si,(%r8)
	and	%si,(%r9)
	and	%si,(%r10)
	and	%si,(%r11)
	and	%si,(%r12)
	and	%si,(%r13)
	and	%si,(%r14)
	and	%si,(%r15)
	nop
	and	%di,(%rax)
	and	%di,(%rcx)
	and	%di,(%rdx)
	and	%di,(%rbx)
	and	%di,(%rsp)
	and	%di,(%rbp)
	and	%di,(%rsi)
	and	%di,(%rdi)
	and	%di,(%r8)
	and	%di,(%r9)
	and	%di,(%r10)
	and	%di,(%r11)
	and	%di,(%r12)
	and	%di,(%r13)
	and	%di,(%r14)
	and	%di,(%r15)
	nop
	and	%r8w, (%rax)
	and	%r8w, (%rcx)
	and	%r8w, (%rdx)
	and	%r8w, (%rbx)
	and	%r8w, (%rsp)
	and	%r8w, (%rbp)
	and	%r8w, (%rsi)
	and	%r8w, (%rdi)
	and	%r8w, (%r8)
	and	%r8w, (%r9)
	and	%r8w, (%r10)
	and	%r8w, (%r11)
	and	%r8w, (%r12)
	and	%r8w, (%r13)
	and	%r8w, (%r14)
	and	%r8w, (%r15)
	nop
	and	%r9w, (%rax)
	and	%r9w, (%rcx)
	and	%r9w, (%rdx)
	and	%r9w, (%rbx)
	and	%r9w, (%rsp)
	and	%r9w, (%rbp)
	and	%r9w, (%rsi)
	and	%r9w, (%rdi)
	and	%r9w, (%r8)
	and	%r9w, (%r9)
	and	%r9w, (%r10)
	and	%r9w, (%r11)
	and	%r9w, (%r12)
	and	%r9w, (%r13)
	and	%r9w, (%r14)
	and	%r9w, (%r15)
	nop
	and	%r10w,(%rax)
	and	%r10w,(%rcx)
	and	%r10w,(%rdx)
	and	%r10w,(%rbx)
	and	%r10w,(%rsp)
	and	%r10w,(%rbp)
	and	%r10w,(%rsi)
	and	%r10w,(%rdi)
	and	%r10w,(%r8)
	and	%r10w,(%r9)
	and	%r10w,(%r10)
	and	%r10w,(%r11)
	and	%r10w,(%r12)
	and	%r10w,(%r13)
	and	%r10w,(%r14)
	and	%r10w,(%r15)
	nop
	and	%r11w,(%rax)
	and	%r11w,(%rcx)
	and	%r11w,(%rdx)
	and	%r11w,(%rbx)
	and	%r11w,(%rsp)
	and	%r11w,(%rbp)
	and	%r11w,(%rsi)
	and	%r11w,(%rdi)
	and	%r11w,(%r8)
	and	%r11w,(%r9)
	and	%r11w,(%r10)
	and	%r11w,(%r11)
	and	%r11w,(%r12)
	and	%r11w,(%r13)
	and	%r11w,(%r14)
	and	%r11w,(%r15)
	nop
	and	%r12w,(%rax)
	and	%r12w,(%rcx)
	and	%r12w,(%rdx)
	and	%r12w,(%rbx)
	and	%r12w,(%rsp)
	and	%r12w,(%rbp)
	and	%r12w,(%rsi)
	and	%r12w,(%rdi)
	and	%r12w,(%r8)
	and	%r12w,(%r9)
	and	%r12w,(%r10)
	and	%r12w,(%r11)
	and	%r12w,(%r12)
	and	%r12w,(%r13)
	and	%r12w,(%r14)
	and	%r12w,(%r15)
	nop
	and	%r13w,(%rax)
	and	%r13w,(%rcx)
	and	%r13w,(%rdx)
	and	%r13w,(%rbx)
	and	%r13w,(%rsp)
	and	%r13w,(%rbp)
	and	%r13w,(%rsi)
	and	%r13w,(%rdi)
	and	%r13w,(%r8)
	and	%r13w,(%r9)
	and	%r13w,(%r10)
	and	%r13w,(%r11)
	and	%r13w,(%r12)
	and	%r13w,(%r13)
	and	%r13w,(%r14)
	and	%r13w,(%r15)
	nop
	and	%r14w,(%rax)
	and	%r14w,(%rcx)
	and	%r14w,(%rdx)
	and	%r14w,(%rbx)
	and	%r14w,(%rsp)
	and	%r14w,(%rbp)
	and	%r14w,(%rsi)
	and	%r14w,(%rdi)
	and	%r14w,(%r8)
	and	%r14w,(%r9)
	and	%r14w,(%r10)
	and	%r14w,(%r11)
	and	%r14w,(%r12)
	and	%r14w,(%r13)
	and	%r14w,(%r14)
	and	%r14w,(%r15)
	nop
	and	%r15w,(%rax)
	and	%r15w,(%rcx)
	and	%r15w,(%rdx)
	and	%r15w,(%rbx)
	and	%r15w,(%rsp)
	and	%r15w,(%rbp)
	and	%r15w,(%rsi)
	and	%r15w,(%rdi)
	and	%r15w,(%r8)
	and	%r15w,(%r9)
	and	%r15w,(%r10)
	and	%r15w,(%r11)
	and	%r15w,(%r12)
	and	%r15w,(%r13)
	and	%r15w,(%r14)
	and	%r15w,(%r15)
        nop
        nop
        // mem32 += reg32
	and	%eax,(%rax)
	and	%eax,(%rcx)
	and	%eax,(%rdx)
	and	%eax,(%rbx)
	and	%eax,(%rsp)
	and	%eax,(%rbp)
	and	%eax,(%rsi)
	and	%eax,(%rdi)
	and	%eax,(%r8)
	and	%eax,(%r9)
	and	%eax,(%r10)
	and	%eax,(%r11)
	and	%eax,(%r12)
	and	%eax,(%r13)
	and	%eax,(%r14)
	and	%eax,(%r15)
	nop
	and	%ecx,(%rax)
	and	%ecx,(%rcx)
	and	%ecx,(%rdx)
	and	%ecx,(%rbx)
	and	%ecx,(%rsp)
	and	%ecx,(%rbp)
	and	%ecx,(%rsi)
	and	%ecx,(%rdi)
	and	%ecx,(%r8)
	and	%ecx,(%r9)
	and	%ecx,(%r10)
	and	%ecx,(%r11)
	and	%ecx,(%r12)
	and	%ecx,(%r13)
	and	%ecx,(%r14)
	and	%ecx,(%r15)
	nop
	and	%edx,(%rax)
	and	%edx,(%rcx)
	and	%edx,(%rdx)
	and	%edx,(%rbx)
	and	%edx,(%rsp)
	and	%edx,(%rbp)
	and	%edx,(%rsi)
	and	%edx,(%rdi)
	and	%edx,(%r8)
	and	%edx,(%r9)
	and	%edx,(%r10)
	and	%edx,(%r11)
	and	%edx,(%r12)
	and	%edx,(%r13)
	and	%edx,(%r14)
	and	%edx,(%r15)
	nop
	and	%ebx,(%rax)
	and	%ebx,(%rcx)
	and	%ebx,(%rdx)
	and	%ebx,(%rbx)
	and	%ebx,(%rsp)
	and	%ebx,(%rbp)
	and	%ebx,(%rsi)
	and	%ebx,(%rdi)
	and	%ebx,(%r8)
	and	%ebx,(%r9)
	and	%ebx,(%r10)
	and	%ebx,(%r11)
	and	%ebx,(%r12)
	and	%ebx,(%r13)
	and	%ebx,(%r14)
	and	%ebx,(%r15)
	nop
	and	%esp,(%rax)
	and	%esp,(%rcx)
	and	%esp,(%rdx)
	and	%esp,(%rbx)
	and	%esp,(%rsp)
	and	%esp,(%rbp)
	and	%esp,(%rsi)
	and	%esp,(%rdi)
	and	%esp,(%r8)
	and	%esp,(%r9)
	and	%esp,(%r10)
	and	%esp,(%r11)
	and	%esp,(%r12)
	and	%esp,(%r13)
	and	%esp,(%r14)
	and	%esp,(%r15)
	nop
	and	%ebp,(%rax)
	and	%ebp,(%rcx)
	and	%ebp,(%rdx)
	and	%ebp,(%rbx)
	and	%ebp,(%rsp)
	and	%ebp,(%rbp)
	and	%ebp,(%rsi)
	and	%ebp,(%rdi)
	and	%ebp,(%r8)
	and	%ebp,(%r9)
	and	%ebp,(%r10)
	and	%ebp,(%r11)
	and	%ebp,(%r12)
	and	%ebp,(%r13)
	and	%ebp,(%r14)
	and	%ebp,(%r15)
	nop
	and	%esi,(%rax)
	and	%esi,(%rcx)
	and	%esi,(%rdx)
	and	%esi,(%rbx)
	and	%esi,(%rsp)
	and	%esi,(%rbp)
	and	%esi,(%rsi)
	and	%esi,(%rdi)
	and	%esi,(%r8)
	and	%esi,(%r9)
	and	%esi,(%r10)
	and	%esi,(%r11)
	and	%esi,(%r12)
	and	%esi,(%r13)
	and	%esi,(%r14)
	and	%esi,(%r15)
	nop
	and	%edi,(%rax)
	and	%edi,(%rcx)
	and	%edi,(%rdx)
	and	%edi,(%rbx)
	and	%edi,(%rsp)
	and	%edi,(%rbp)
	and	%edi,(%rsi)
	and	%edi,(%rdi)
	and	%edi,(%r8)
	and	%edi,(%r9)
	and	%edi,(%r10)
	and	%edi,(%r11)
	and	%edi,(%r12)
	and	%edi,(%r13)
	and	%edi,(%r14)
	and	%edi,(%r15)
	nop
	and	%r8d, (%rax)
	and	%r8d, (%rcx)
	and	%r8d, (%rdx)
	and	%r8d, (%rbx)
	and	%r8d, (%rsp)
	and	%r8d, (%rbp)
	and	%r8d, (%rsi)
	and	%r8d, (%rdi)
	and	%r8d, (%r8)
	and	%r8d, (%r9)
	and	%r8d, (%r10)
	and	%r8d, (%r11)
	and	%r8d, (%r12)
	and	%r8d, (%r13)
	and	%r8d, (%r14)
	and	%r8d, (%r15)
	nop
	and	%r9d, (%rax)
	and	%r9d, (%rcx)
	and	%r9d, (%rdx)
	and	%r9d, (%rbx)
	and	%r9d, (%rsp)
	and	%r9d, (%rbp)
	and	%r9d, (%rsi)
	and	%r9d, (%rdi)
	and	%r9d, (%r8)
	and	%r9d, (%r9)
	and	%r9d, (%r10)
	and	%r9d, (%r11)
	and	%r9d, (%r12)
	and	%r9d, (%r13)
	and	%r9d, (%r14)
	and	%r9d, (%r15)
	nop
	and	%r10d,(%rax)
	and	%r10d,(%rcx)
	and	%r10d,(%rdx)
	and	%r10d,(%rbx)
	and	%r10d,(%rsp)
	and	%r10d,(%rbp)
	and	%r10d,(%rsi)
	and	%r10d,(%rdi)
	and	%r10d,(%r8)
	and	%r10d,(%r9)
	and	%r10d,(%r10)
	and	%r10d,(%r11)
	and	%r10d,(%r12)
	and	%r10d,(%r13)
	and	%r10d,(%r14)
	and	%r10d,(%r15)
	nop
	and	%r11d,(%rax)
	and	%r11d,(%rcx)
	and	%r11d,(%rdx)
	and	%r11d,(%rbx)
	and	%r11d,(%rsp)
	and	%r11d,(%rbp)
	and	%r11d,(%rsi)
	and	%r11d,(%rdi)
	and	%r11d,(%r8)
	and	%r11d,(%r9)
	and	%r11d,(%r10)
	and	%r11d,(%r11)
	and	%r11d,(%r12)
	and	%r11d,(%r13)
	and	%r11d,(%r14)
	and	%r11d,(%r15)
	nop
	and	%r12d,(%rax)
	and	%r12d,(%rcx)
	and	%r12d,(%rdx)
	and	%r12d,(%rbx)
	and	%r12d,(%rsp)
	and	%r12d,(%rbp)
	and	%r12d,(%rsi)
	and	%r12d,(%rdi)
	and	%r12d,(%r8)
	and	%r12d,(%r9)
	and	%r12d,(%r10)
	and	%r12d,(%r11)
	and	%r12d,(%r12)
	and	%r12d,(%r13)
	and	%r12d,(%r14)
	and	%r12d,(%r15)
	nop
	and	%r13d,(%rax)
	and	%r13d,(%rcx)
	and	%r13d,(%rdx)
	and	%r13d,(%rbx)
	and	%r13d,(%rsp)
	and	%r13d,(%rbp)
	and	%r13d,(%rsi)
	and	%r13d,(%rdi)
	and	%r13d,(%r8)
	and	%r13d,(%r9)
	and	%r13d,(%r10)
	and	%r13d,(%r11)
	and	%r13d,(%r12)
	and	%r13d,(%r13)
	and	%r13d,(%r14)
	and	%r13d,(%r15)
	nop
	and	%r14d,(%rax)
	and	%r14d,(%rcx)
	and	%r14d,(%rdx)
	and	%r14d,(%rbx)
	and	%r14d,(%rsp)
	and	%r14d,(%rbp)
	and	%r14d,(%rsi)
	and	%r14d,(%rdi)
	and	%r14d,(%r8)
	and	%r14d,(%r9)
	and	%r14d,(%r10)
	and	%r14d,(%r11)
	and	%r14d,(%r12)
	and	%r14d,(%r13)
	and	%r14d,(%r14)
	and	%r14d,(%r15)
	nop
	and	%r15d,(%rax)
	and	%r15d,(%rcx)
	and	%r15d,(%rdx)
	and	%r15d,(%rbx)
	and	%r15d,(%rsp)
	and	%r15d,(%rbp)
	and	%r15d,(%rsi)
	and	%r15d,(%rdi)
	and	%r15d,(%r8)
	and	%r15d,(%r9)
	and	%r15d,(%r10)
	and	%r15d,(%r11)
	and	%r15d,(%r12)
	and	%r15d,(%r13)
	and	%r15d,(%r14)
	and	%r15d,(%r15)
        nop
        nop
        // mem64 += reg64
	and	%rax,(%rax)
	and	%rax,(%rcx)
	and	%rax,(%rdx)
	and	%rax,(%rbx)
	and	%rax,(%rsp)
	and	%rax,(%rbp)
	and	%rax,(%rsi)
	and	%rax,(%rdi)
	and	%rax,(%r8)
	and	%rax,(%r9)
	and	%rax,(%r10)
	and	%rax,(%r11)
	and	%rax,(%r12)
	and	%rax,(%r13)
	and	%rax,(%r14)
	and	%rax,(%r15)
	nop
	and	%rcx,(%rax)
	and	%rcx,(%rcx)
	and	%rcx,(%rdx)
	and	%rcx,(%rbx)
	and	%rcx,(%rsp)
	and	%rcx,(%rbp)
	and	%rcx,(%rsi)
	and	%rcx,(%rdi)
	and	%rcx,(%r8)
	and	%rcx,(%r9)
	and	%rcx,(%r10)
	and	%rcx,(%r11)
	and	%rcx,(%r12)
	and	%rcx,(%r13)
	and	%rcx,(%r14)
	and	%rcx,(%r15)
	nop
	and	%rdx,(%rax)
	and	%rdx,(%rcx)
	and	%rdx,(%rdx)
	and	%rdx,(%rbx)
	and	%rdx,(%rsp)
	and	%rdx,(%rbp)
	and	%rdx,(%rsi)
	and	%rdx,(%rdi)
	and	%rdx,(%r8)
	and	%rdx,(%r9)
	and	%rdx,(%r10)
	and	%rdx,(%r11)
	and	%rdx,(%r12)
	and	%rdx,(%r13)
	and	%rdx,(%r14)
	and	%rdx,(%r15)
	nop
	and	%rbx,(%rax)
	and	%rbx,(%rcx)
	and	%rbx,(%rdx)
	and	%rbx,(%rbx)
	and	%rbx,(%rsp)
	and	%rbx,(%rbp)
	and	%rbx,(%rsi)
	and	%rbx,(%rdi)
	and	%rbx,(%r8)
	and	%rbx,(%r9)
	and	%rbx,(%r10)
	and	%rbx,(%r11)
	and	%rbx,(%r12)
	and	%rbx,(%r13)
	and	%rbx,(%r14)
	and	%rbx,(%r15)
	nop
	and	%rsp,(%rax)
	and	%rsp,(%rcx)
	and	%rsp,(%rdx)
	and	%rsp,(%rbx)
	and	%rsp,(%rsp)
	and	%rsp,(%rbp)
	and	%rsp,(%rsi)
	and	%rsp,(%rdi)
	and	%rsp,(%r8)
	and	%rsp,(%r9)
	and	%rsp,(%r10)
	and	%rsp,(%r11)
	and	%rsp,(%r12)
	and	%rsp,(%r13)
	and	%rsp,(%r14)
	and	%rsp,(%r15)
	nop
	and	%rbp,(%rax)
	and	%rbp,(%rcx)
	and	%rbp,(%rdx)
	and	%rbp,(%rbx)
	and	%rbp,(%rsp)
	and	%rbp,(%rbp)
	and	%rbp,(%rsi)
	and	%rbp,(%rdi)
	and	%rbp,(%r8)
	and	%rbp,(%r9)
	and	%rbp,(%r10)
	and	%rbp,(%r11)
	and	%rbp,(%r12)
	and	%rbp,(%r13)
	and	%rbp,(%r14)
	and	%rbp,(%r15)
	nop
	and	%rsi,(%rax)
	and	%rsi,(%rcx)
	and	%rsi,(%rdx)
	and	%rsi,(%rbx)
	and	%rsi,(%rsp)
	and	%rsi,(%rbp)
	and	%rsi,(%rsi)
	and	%rsi,(%rdi)
	and	%rsi,(%r8)
	and	%rsi,(%r9)
	and	%rsi,(%r10)
	and	%rsi,(%r11)
	and	%rsi,(%r12)
	and	%rsi,(%r13)
	and	%rsi,(%r14)
	and	%rsi,(%r15)
	nop
	and	%rdi,(%rax)
	and	%rdi,(%rcx)
	and	%rdi,(%rdx)
	and	%rdi,(%rbx)
	and	%rdi,(%rsp)
	and	%rdi,(%rbp)
	and	%rdi,(%rsi)
	and	%rdi,(%rdi)
	and	%rdi,(%r8)
	and	%rdi,(%r9)
	and	%rdi,(%r10)
	and	%rdi,(%r11)
	and	%rdi,(%r12)
	and	%rdi,(%r13)
	and	%rdi,(%r14)
	and	%rdi,(%r15)
	nop
	and	%r8, (%rax)
	and	%r8, (%rcx)
	and	%r8, (%rdx)
	and	%r8, (%rbx)
	and	%r8, (%rsp)
	and	%r8, (%rbp)
	and	%r8, (%rsi)
	and	%r8, (%rdi)
	and	%r8, (%r8)
	and	%r8, (%r9)
	and	%r8, (%r10)
	and	%r8, (%r11)
	and	%r8, (%r12)
	and	%r8, (%r13)
	and	%r8, (%r14)
	and	%r8, (%r15)
	nop
	and	%r9, (%rax)
	and	%r9, (%rcx)
	and	%r9, (%rdx)
	and	%r9, (%rbx)
	and	%r9, (%rsp)
	and	%r9, (%rbp)
	and	%r9, (%rsi)
	and	%r9, (%rdi)
	and	%r9, (%r8)
	and	%r9, (%r9)
	and	%r9, (%r10)
	and	%r9, (%r11)
	and	%r9, (%r12)
	and	%r9, (%r13)
	and	%r9, (%r14)
	and	%r9, (%r15)
	nop
	and	%r10,(%rax)
	and	%r10,(%rcx)
	and	%r10,(%rdx)
	and	%r10,(%rbx)
	and	%r10,(%rsp)
	and	%r10,(%rbp)
	and	%r10,(%rsi)
	and	%r10,(%rdi)
	and	%r10,(%r8)
	and	%r10,(%r9)
	and	%r10,(%r10)
	and	%r10,(%r11)
	and	%r10,(%r12)
	and	%r10,(%r13)
	and	%r10,(%r14)
	and	%r10,(%r15)
	nop
	and	%r11,(%rax)
	and	%r11,(%rcx)
	and	%r11,(%rdx)
	and	%r11,(%rbx)
	and	%r11,(%rsp)
	and	%r11,(%rbp)
	and	%r11,(%rsi)
	and	%r11,(%rdi)
	and	%r11,(%r8)
	and	%r11,(%r9)
	and	%r11,(%r10)
	and	%r11,(%r11)
	and	%r11,(%r12)
	and	%r11,(%r13)
	and	%r11,(%r14)
	and	%r11,(%r15)
	nop
	and	%r12,(%rax)
	and	%r12,(%rcx)
	and	%r12,(%rdx)
	and	%r12,(%rbx)
	and	%r12,(%rsp)
	and	%r12,(%rbp)
	and	%r12,(%rsi)
	and	%r12,(%rdi)
	and	%r12,(%r8)
	and	%r12,(%r9)
	and	%r12,(%r10)
	and	%r12,(%r11)
	and	%r12,(%r12)
	and	%r12,(%r13)
	and	%r12,(%r14)
	and	%r12,(%r15)
	nop
	and	%r13,(%rax)
	and	%r13,(%rcx)
	and	%r13,(%rdx)
	and	%r13,(%rbx)
	and	%r13,(%rsp)
	and	%r13,(%rbp)
	and	%r13,(%rsi)
	and	%r13,(%rdi)
	and	%r13,(%r8)
	and	%r13,(%r9)
	and	%r13,(%r10)
	and	%r13,(%r11)
	and	%r13,(%r12)
	and	%r13,(%r13)
	and	%r13,(%r14)
	and	%r13,(%r15)
	nop
	and	%r14,(%rax)
	and	%r14,(%rcx)
	and	%r14,(%rdx)
	and	%r14,(%rbx)
	and	%r14,(%rsp)
	and	%r14,(%rbp)
	and	%r14,(%rsi)
	and	%r14,(%rdi)
	and	%r14,(%r8)
	and	%r14,(%r9)
	and	%r14,(%r10)
	and	%r14,(%r11)
	and	%r14,(%r12)
	and	%r14,(%r13)
	and	%r14,(%r14)
	and	%r14,(%r15)
	nop
	and	%r15,(%rax)
	and	%r15,(%rcx)
	and	%r15,(%rdx)
	and	%r15,(%rbx)
	and	%r15,(%rsp)
	and	%r15,(%rbp)
	and	%r15,(%rsi)
	and	%r15,(%rdi)
	and	%r15,(%r8)
	and	%r15,(%r9)
	and	%r15,(%r10)
	and	%r15,(%r11)
	and	%r15,(%r12)
	and	%r15,(%r13)
	and	%r15,(%r14)
	and	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndMem8Reg
	.type	AndMem8Reg, @function
AndMem8Reg:
	.cfi_startproc
	and	%rax,0x7f(%rax)
	and	%rax,0x7f(%rcx)
	and	%rax,0x7f(%rdx)
	and	%rax,0x7f(%rbx)
	and	%rax,0x7f(%rsp)
	and	%rax,0x7f(%rbp)
	and	%rax,0x7f(%rsi)
	and	%rax,0x7f(%rdi)
	and	%rax,0x7f(%r8)
	and	%rax,0x7f(%r9)
	and	%rax,0x7f(%r10)
	and	%rax,0x7f(%r11)
	and	%rax,0x7f(%r12)
	and	%rax,0x7f(%r13)
	and	%rax,0x7f(%r14)
	and	%rax,0x7f(%r15)
	nop
	and	%rcx,0x7f(%rax)
	and	%rcx,0x7f(%rcx)
	and	%rcx,0x7f(%rdx)
	and	%rcx,0x7f(%rbx)
	and	%rcx,0x7f(%rsp)
	and	%rcx,0x7f(%rbp)
	and	%rcx,0x7f(%rsi)
	and	%rcx,0x7f(%rdi)
	and	%rcx,0x7f(%r8)
	and	%rcx,0x7f(%r9)
	and	%rcx,0x7f(%r10)
	and	%rcx,0x7f(%r11)
	and	%rcx,0x7f(%r12)
	and	%rcx,0x7f(%r13)
	and	%rcx,0x7f(%r14)
	and	%rcx,0x7f(%r15)
	nop
	and	%rdx,0x7f(%rax)
	and	%rdx,0x7f(%rcx)
	and	%rdx,0x7f(%rdx)
	and	%rdx,0x7f(%rbx)
	and	%rdx,0x7f(%rsp)
	and	%rdx,0x7f(%rbp)
	and	%rdx,0x7f(%rsi)
	and	%rdx,0x7f(%rdi)
	and	%rdx,0x7f(%r8)
	and	%rdx,0x7f(%r9)
	and	%rdx,0x7f(%r10)
	and	%rdx,0x7f(%r11)
	and	%rdx,0x7f(%r12)
	and	%rdx,0x7f(%r13)
	and	%rdx,0x7f(%r14)
	and	%rdx,0x7f(%r15)
	nop
	and	%rbx,0x7f(%rax)
	and	%rbx,0x7f(%rcx)
	and	%rbx,0x7f(%rdx)
	and	%rbx,0x7f(%rbx)
	and	%rbx,0x7f(%rsp)
	and	%rbx,0x7f(%rbp)
	and	%rbx,0x7f(%rsi)
	and	%rbx,0x7f(%rdi)
	and	%rbx,0x7f(%r8)
	and	%rbx,0x7f(%r9)
	and	%rbx,0x7f(%r10)
	and	%rbx,0x7f(%r11)
	and	%rbx,0x7f(%r12)
	and	%rbx,0x7f(%r13)
	and	%rbx,0x7f(%r14)
	and	%rbx,0x7f(%r15)
	nop
	and	%rsp,0x7f(%rax)
	and	%rsp,0x7f(%rcx)
	and	%rsp,0x7f(%rdx)
	and	%rsp,0x7f(%rbx)
	and	%rsp,0x7f(%rsp)
	and	%rsp,0x7f(%rbp)
	and	%rsp,0x7f(%rsi)
	and	%rsp,0x7f(%rdi)
	and	%rsp,0x7f(%r8)
	and	%rsp,0x7f(%r9)
	and	%rsp,0x7f(%r10)
	and	%rsp,0x7f(%r11)
	and	%rsp,0x7f(%r12)
	and	%rsp,0x7f(%r13)
	and	%rsp,0x7f(%r14)
	and	%rsp,0x7f(%r15)
	nop
	and	%rbp,0x7f(%rax)
	and	%rbp,0x7f(%rcx)
	and	%rbp,0x7f(%rdx)
	and	%rbp,0x7f(%rbx)
	and	%rbp,0x7f(%rsp)
	and	%rbp,0x7f(%rbp)
	and	%rbp,0x7f(%rsi)
	and	%rbp,0x7f(%rdi)
	and	%rbp,0x7f(%r8)
	and	%rbp,0x7f(%r9)
	and	%rbp,0x7f(%r10)
	and	%rbp,0x7f(%r11)
	and	%rbp,0x7f(%r12)
	and	%rbp,0x7f(%r13)
	and	%rbp,0x7f(%r14)
	and	%rbp,0x7f(%r15)
	nop
	and	%rsi,0x7f(%rax)
	and	%rsi,0x7f(%rcx)
	and	%rsi,0x7f(%rdx)
	and	%rsi,0x7f(%rbx)
	and	%rsi,0x7f(%rsp)
	and	%rsi,0x7f(%rbp)
	and	%rsi,0x7f(%rsi)
	and	%rsi,0x7f(%rdi)
	and	%rsi,0x7f(%r8)
	and	%rsi,0x7f(%r9)
	and	%rsi,0x7f(%r10)
	and	%rsi,0x7f(%r11)
	and	%rsi,0x7f(%r12)
	and	%rsi,0x7f(%r13)
	and	%rsi,0x7f(%r14)
	and	%rsi,0x7f(%r15)
	nop
	and	%rdi,0x7f(%rax)
	and	%rdi,0x7f(%rcx)
	and	%rdi,0x7f(%rdx)
	and	%rdi,0x7f(%rbx)
	and	%rdi,0x7f(%rsp)
	and	%rdi,0x7f(%rbp)
	and	%rdi,0x7f(%rsi)
	and	%rdi,0x7f(%rdi)
	and	%rdi,0x7f(%r8)
	and	%rdi,0x7f(%r9)
	and	%rdi,0x7f(%r10)
	and	%rdi,0x7f(%r11)
	and	%rdi,0x7f(%r12)
	and	%rdi,0x7f(%r13)
	and	%rdi,0x7f(%r14)
	and	%rdi,0x7f(%r15)
	nop
	and	%r8, 0x7f(%rax)
	and	%r8, 0x7f(%rcx)
	and	%r8, 0x7f(%rdx)
	and	%r8, 0x7f(%rbx)
	and	%r8, 0x7f(%rsp)
	and	%r8, 0x7f(%rbp)
	and	%r8, 0x7f(%rsi)
	and	%r8, 0x7f(%rdi)
	and	%r8, 0x7f(%r8)
	and	%r8, 0x7f(%r9)
	and	%r8, 0x7f(%r10)
	and	%r8, 0x7f(%r11)
	and	%r8, 0x7f(%r12)
	and	%r8, 0x7f(%r13)
	and	%r8, 0x7f(%r14)
	and	%r8, 0x7f(%r15)
	nop
	and	%r9, 0x7f(%rax)
	and	%r9, 0x7f(%rcx)
	and	%r9, 0x7f(%rdx)
	and	%r9, 0x7f(%rbx)
	and	%r9, 0x7f(%rsp)
	and	%r9, 0x7f(%rbp)
	and	%r9, 0x7f(%rsi)
	and	%r9, 0x7f(%rdi)
	and	%r9, 0x7f(%r8)
	and	%r9, 0x7f(%r9)
	and	%r9, 0x7f(%r10)
	and	%r9, 0x7f(%r11)
	and	%r9, 0x7f(%r12)
	and	%r9, 0x7f(%r13)
	and	%r9, 0x7f(%r14)
	and	%r9, 0x7f(%r15)
	nop
	and	%r10,0x7f(%rax)
	and	%r10,0x7f(%rcx)
	and	%r10,0x7f(%rdx)
	and	%r10,0x7f(%rbx)
	and	%r10,0x7f(%rsp)
	and	%r10,0x7f(%rbp)
	and	%r10,0x7f(%rsi)
	and	%r10,0x7f(%rdi)
	and	%r10,0x7f(%r8)
	and	%r10,0x7f(%r9)
	and	%r10,0x7f(%r10)
	and	%r10,0x7f(%r11)
	and	%r10,0x7f(%r12)
	and	%r10,0x7f(%r13)
	and	%r10,0x7f(%r14)
	and	%r10,0x7f(%r15)
	nop
	and	%r11,0x7f(%rax)
	and	%r11,0x7f(%rcx)
	and	%r11,0x7f(%rdx)
	and	%r11,0x7f(%rbx)
	and	%r11,0x7f(%rsp)
	and	%r11,0x7f(%rbp)
	and	%r11,0x7f(%rsi)
	and	%r11,0x7f(%rdi)
	and	%r11,0x7f(%r8)
	and	%r11,0x7f(%r9)
	and	%r11,0x7f(%r10)
	and	%r11,0x7f(%r11)
	and	%r11,0x7f(%r12)
	and	%r11,0x7f(%r13)
	and	%r11,0x7f(%r14)
	and	%r11,0x7f(%r15)
	nop
	and	%r12,0x7f(%rax)
	and	%r12,0x7f(%rcx)
	and	%r12,0x7f(%rdx)
	and	%r12,0x7f(%rbx)
	and	%r12,0x7f(%rsp)
	and	%r12,0x7f(%rbp)
	and	%r12,0x7f(%rsi)
	and	%r12,0x7f(%rdi)
	and	%r12,0x7f(%r8)
	and	%r12,0x7f(%r9)
	and	%r12,0x7f(%r10)
	and	%r12,0x7f(%r11)
	and	%r12,0x7f(%r12)
	and	%r12,0x7f(%r13)
	and	%r12,0x7f(%r14)
	and	%r12,0x7f(%r15)
	nop
	and	%r13,0x7f(%rax)
	and	%r13,0x7f(%rcx)
	and	%r13,0x7f(%rdx)
	and	%r13,0x7f(%rbx)
	and	%r13,0x7f(%rsp)
	and	%r13,0x7f(%rbp)
	and	%r13,0x7f(%rsi)
	and	%r13,0x7f(%rdi)
	and	%r13,0x7f(%r8)
	and	%r13,0x7f(%r9)
	and	%r13,0x7f(%r10)
	and	%r13,0x7f(%r11)
	and	%r13,0x7f(%r12)
	and	%r13,0x7f(%r13)
	and	%r13,0x7f(%r14)
	and	%r13,0x7f(%r15)
	nop
	and	%r14,0x7f(%rax)
	and	%r14,0x7f(%rcx)
	and	%r14,0x7f(%rdx)
	and	%r14,0x7f(%rbx)
	and	%r14,0x7f(%rsp)
	and	%r14,0x7f(%rbp)
	and	%r14,0x7f(%rsi)
	and	%r14,0x7f(%rdi)
	and	%r14,0x7f(%r8)
	and	%r14,0x7f(%r9)
	and	%r14,0x7f(%r10)
	and	%r14,0x7f(%r11)
	and	%r14,0x7f(%r12)
	and	%r14,0x7f(%r13)
	and	%r14,0x7f(%r14)
	and	%r14,0x7f(%r15)
	nop
	and	%r15,0x7f(%rax)
	and	%r15,0x7f(%rcx)
	and	%r15,0x7f(%rdx)
	and	%r15,0x7f(%rbx)
	and	%r15,0x7f(%rsp)
	and	%r15,0x7f(%rbp)
	and	%r15,0x7f(%rsi)
	and	%r15,0x7f(%rdi)
	and	%r15,0x7f(%r8)
	and	%r15,0x7f(%r9)
	and	%r15,0x7f(%r10)
	and	%r15,0x7f(%r11)
	and	%r15,0x7f(%r12)
	and	%r15,0x7f(%r13)
	and	%r15,0x7f(%r14)
	and	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndMem32Reg
	.type	AndMem32Reg, @function
AndMem32Reg:
	.cfi_startproc
	and	%rax,0x7f563412(%rax)
	and	%rax,0x7f563412(%rcx)
	and	%rax,0x7f563412(%rdx)
	and	%rax,0x7f563412(%rbx)
	and	%rax,0x7f563412(%rsp)
	and	%rax,0x7f563412(%rbp)
	and	%rax,0x7f563412(%rsi)
	and	%rax,0x7f563412(%rdi)
	and	%rax,0x7f563412(%r8)
	and	%rax,0x7f563412(%r9)
	and	%rax,0x7f563412(%r10)
	and	%rax,0x7f563412(%r11)
	and	%rax,0x7f563412(%r12)
	and	%rax,0x7f563412(%r13)
	and	%rax,0x7f563412(%r14)
	and	%rax,0x7f563412(%r15)
	nop
	and	%rcx,0x7f563412(%rax)
	and	%rcx,0x7f563412(%rcx)
	and	%rcx,0x7f563412(%rdx)
	and	%rcx,0x7f563412(%rbx)
	and	%rcx,0x7f563412(%rsp)
	and	%rcx,0x7f563412(%rbp)
	and	%rcx,0x7f563412(%rsi)
	and	%rcx,0x7f563412(%rdi)
	and	%rcx,0x7f563412(%r8)
	and	%rcx,0x7f563412(%r9)
	and	%rcx,0x7f563412(%r10)
	and	%rcx,0x7f563412(%r11)
	and	%rcx,0x7f563412(%r12)
	and	%rcx,0x7f563412(%r13)
	and	%rcx,0x7f563412(%r14)
	and	%rcx,0x7f563412(%r15)
	nop
	and	%rdx,0x7f563412(%rax)
	and	%rdx,0x7f563412(%rcx)
	and	%rdx,0x7f563412(%rdx)
	and	%rdx,0x7f563412(%rbx)
	and	%rdx,0x7f563412(%rsp)
	and	%rdx,0x7f563412(%rbp)
	and	%rdx,0x7f563412(%rsi)
	and	%rdx,0x7f563412(%rdi)
	and	%rdx,0x7f563412(%r8)
	and	%rdx,0x7f563412(%r9)
	and	%rdx,0x7f563412(%r10)
	and	%rdx,0x7f563412(%r11)
	and	%rdx,0x7f563412(%r12)
	and	%rdx,0x7f563412(%r13)
	and	%rdx,0x7f563412(%r14)
	and	%rdx,0x7f563412(%r15)
	nop
	and	%rbx,0x7f563412(%rax)
	and	%rbx,0x7f563412(%rcx)
	and	%rbx,0x7f563412(%rdx)
	and	%rbx,0x7f563412(%rbx)
	and	%rbx,0x7f563412(%rsp)
	and	%rbx,0x7f563412(%rbp)
	and	%rbx,0x7f563412(%rsi)
	and	%rbx,0x7f563412(%rdi)
	and	%rbx,0x7f563412(%r8)
	and	%rbx,0x7f563412(%r9)
	and	%rbx,0x7f563412(%r10)
	and	%rbx,0x7f563412(%r11)
	and	%rbx,0x7f563412(%r12)
	and	%rbx,0x7f563412(%r13)
	and	%rbx,0x7f563412(%r14)
	and	%rbx,0x7f563412(%r15)
	nop
	and	%rsp,0x7f563412(%rax)
	and	%rsp,0x7f563412(%rcx)
	and	%rsp,0x7f563412(%rdx)
	and	%rsp,0x7f563412(%rbx)
	and	%rsp,0x7f563412(%rsp)
	and	%rsp,0x7f563412(%rbp)
	and	%rsp,0x7f563412(%rsi)
	and	%rsp,0x7f563412(%rdi)
	and	%rsp,0x7f563412(%r8)
	and	%rsp,0x7f563412(%r9)
	and	%rsp,0x7f563412(%r10)
	and	%rsp,0x7f563412(%r11)
	and	%rsp,0x7f563412(%r12)
	and	%rsp,0x7f563412(%r13)
	and	%rsp,0x7f563412(%r14)
	and	%rsp,0x7f563412(%r15)
	nop
	and	%rbp,0x7f563412(%rax)
	and	%rbp,0x7f563412(%rcx)
	and	%rbp,0x7f563412(%rdx)
	and	%rbp,0x7f563412(%rbx)
	and	%rbp,0x7f563412(%rsp)
	and	%rbp,0x7f563412(%rbp)
	and	%rbp,0x7f563412(%rsi)
	and	%rbp,0x7f563412(%rdi)
	and	%rbp,0x7f563412(%r8)
	and	%rbp,0x7f563412(%r9)
	and	%rbp,0x7f563412(%r10)
	and	%rbp,0x7f563412(%r11)
	and	%rbp,0x7f563412(%r12)
	and	%rbp,0x7f563412(%r13)
	and	%rbp,0x7f563412(%r14)
	and	%rbp,0x7f563412(%r15)
	nop
	and	%rsi,0x7f563412(%rax)
	and	%rsi,0x7f563412(%rcx)
	and	%rsi,0x7f563412(%rdx)
	and	%rsi,0x7f563412(%rbx)
	and	%rsi,0x7f563412(%rsp)
	and	%rsi,0x7f563412(%rbp)
	and	%rsi,0x7f563412(%rsi)
	and	%rsi,0x7f563412(%rdi)
	and	%rsi,0x7f563412(%r8)
	and	%rsi,0x7f563412(%r9)
	and	%rsi,0x7f563412(%r10)
	and	%rsi,0x7f563412(%r11)
	and	%rsi,0x7f563412(%r12)
	and	%rsi,0x7f563412(%r13)
	and	%rsi,0x7f563412(%r14)
	and	%rsi,0x7f563412(%r15)
	nop
	and	%rdi,0x7f563412(%rax)
	and	%rdi,0x7f563412(%rcx)
	and	%rdi,0x7f563412(%rdx)
	and	%rdi,0x7f563412(%rbx)
	and	%rdi,0x7f563412(%rsp)
	and	%rdi,0x7f563412(%rbp)
	and	%rdi,0x7f563412(%rsi)
	and	%rdi,0x7f563412(%rdi)
	and	%rdi,0x7f563412(%r8)
	and	%rdi,0x7f563412(%r9)
	and	%rdi,0x7f563412(%r10)
	and	%rdi,0x7f563412(%r11)
	and	%rdi,0x7f563412(%r12)
	and	%rdi,0x7f563412(%r13)
	and	%rdi,0x7f563412(%r14)
	and	%rdi,0x7f563412(%r15)
	nop
	and	%r8, 0x7f563412(%rax)
	and	%r8, 0x7f563412(%rcx)
	and	%r8, 0x7f563412(%rdx)
	and	%r8, 0x7f563412(%rbx)
	and	%r8, 0x7f563412(%rsp)
	and	%r8, 0x7f563412(%rbp)
	and	%r8, 0x7f563412(%rsi)
	and	%r8, 0x7f563412(%rdi)
	and	%r8, 0x7f563412(%r8)
	and	%r8, 0x7f563412(%r9)
	and	%r8, 0x7f563412(%r10)
	and	%r8, 0x7f563412(%r11)
	and	%r8, 0x7f563412(%r12)
	and	%r8, 0x7f563412(%r13)
	and	%r8, 0x7f563412(%r14)
	and	%r8, 0x7f563412(%r15)
	nop
	and	%r9, 0x7f563412(%rax)
	and	%r9, 0x7f563412(%rcx)
	and	%r9, 0x7f563412(%rdx)
	and	%r9, 0x7f563412(%rbx)
	and	%r9, 0x7f563412(%rsp)
	and	%r9, 0x7f563412(%rbp)
	and	%r9, 0x7f563412(%rsi)
	and	%r9, 0x7f563412(%rdi)
	and	%r9, 0x7f563412(%r8)
	and	%r9, 0x7f563412(%r9)
	and	%r9, 0x7f563412(%r10)
	and	%r9, 0x7f563412(%r11)
	and	%r9, 0x7f563412(%r12)
	and	%r9, 0x7f563412(%r13)
	and	%r9, 0x7f563412(%r14)
	and	%r9, 0x7f563412(%r15)
	nop
	and	%r10,0x7f563412(%rax)
	and	%r10,0x7f563412(%rcx)
	and	%r10,0x7f563412(%rdx)
	and	%r10,0x7f563412(%rbx)
	and	%r10,0x7f563412(%rsp)
	and	%r10,0x7f563412(%rbp)
	and	%r10,0x7f563412(%rsi)
	and	%r10,0x7f563412(%rdi)
	and	%r10,0x7f563412(%r8)
	and	%r10,0x7f563412(%r9)
	and	%r10,0x7f563412(%r10)
	and	%r10,0x7f563412(%r11)
	and	%r10,0x7f563412(%r12)
	and	%r10,0x7f563412(%r13)
	and	%r10,0x7f563412(%r14)
	and	%r10,0x7f563412(%r15)
	nop
	and	%r11,0x7f563412(%rax)
	and	%r11,0x7f563412(%rcx)
	and	%r11,0x7f563412(%rdx)
	and	%r11,0x7f563412(%rbx)
	and	%r11,0x7f563412(%rsp)
	and	%r11,0x7f563412(%rbp)
	and	%r11,0x7f563412(%rsi)
	and	%r11,0x7f563412(%rdi)
	and	%r11,0x7f563412(%r8)
	and	%r11,0x7f563412(%r9)
	and	%r11,0x7f563412(%r10)
	and	%r11,0x7f563412(%r11)
	and	%r11,0x7f563412(%r12)
	and	%r11,0x7f563412(%r13)
	and	%r11,0x7f563412(%r14)
	and	%r11,0x7f563412(%r15)
	nop
	and	%r12,0x7f563412(%rax)
	and	%r12,0x7f563412(%rcx)
	and	%r12,0x7f563412(%rdx)
	and	%r12,0x7f563412(%rbx)
	and	%r12,0x7f563412(%rsp)
	and	%r12,0x7f563412(%rbp)
	and	%r12,0x7f563412(%rsi)
	and	%r12,0x7f563412(%rdi)
	and	%r12,0x7f563412(%r8)
	and	%r12,0x7f563412(%r9)
	and	%r12,0x7f563412(%r10)
	and	%r12,0x7f563412(%r11)
	and	%r12,0x7f563412(%r12)
	and	%r12,0x7f563412(%r13)
	and	%r12,0x7f563412(%r14)
	and	%r12,0x7f563412(%r15)
	nop
	and	%r13,0x7f563412(%rax)
	and	%r13,0x7f563412(%rcx)
	and	%r13,0x7f563412(%rdx)
	and	%r13,0x7f563412(%rbx)
	and	%r13,0x7f563412(%rsp)
	and	%r13,0x7f563412(%rbp)
	and	%r13,0x7f563412(%rsi)
	and	%r13,0x7f563412(%rdi)
	and	%r13,0x7f563412(%r8)
	and	%r13,0x7f563412(%r9)
	and	%r13,0x7f563412(%r10)
	and	%r13,0x7f563412(%r11)
	and	%r13,0x7f563412(%r12)
	and	%r13,0x7f563412(%r13)
	and	%r13,0x7f563412(%r14)
	and	%r13,0x7f563412(%r15)
	nop
	and	%r14,0x7f563412(%rax)
	and	%r14,0x7f563412(%rcx)
	and	%r14,0x7f563412(%rdx)
	and	%r14,0x7f563412(%rbx)
	and	%r14,0x7f563412(%rsp)
	and	%r14,0x7f563412(%rbp)
	and	%r14,0x7f563412(%rsi)
	and	%r14,0x7f563412(%rdi)
	and	%r14,0x7f563412(%r8)
	and	%r14,0x7f563412(%r9)
	and	%r14,0x7f563412(%r10)
	and	%r14,0x7f563412(%r11)
	and	%r14,0x7f563412(%r12)
	and	%r14,0x7f563412(%r13)
	and	%r14,0x7f563412(%r14)
	and	%r14,0x7f563412(%r15)
	nop
	and	%r15,0x7f563412(%rax)
	and	%r15,0x7f563412(%rcx)
	and	%r15,0x7f563412(%rdx)
	and	%r15,0x7f563412(%rbx)
	and	%r15,0x7f563412(%rsp)
	and	%r15,0x7f563412(%rbp)
	and	%r15,0x7f563412(%rsi)
	and	%r15,0x7f563412(%rdi)
	and	%r15,0x7f563412(%r8)
	and	%r15,0x7f563412(%r9)
	and	%r15,0x7f563412(%r10)
	and	%r15,0x7f563412(%r11)
	and	%r15,0x7f563412(%r12)
	and	%r15,0x7f563412(%r13)
	and	%r15,0x7f563412(%r14)
	and	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	AndRegMem
	.type	AndRegMem, @function
AndRegMem:
	.cfi_startproc
	and	(%rax),%rax
	and	(%rax),%rcx
	and	(%rax),%rdx
	and	(%rax),%rbx
	and	(%rax),%rsp
	and	(%rax),%rbp
	and	(%rax),%rsi
	and	(%rax),%rdi
	and	(%rax),%r8
	and	(%rax),%r9
	and	(%rax),%r10
	and	(%rax),%r11
	and	(%rax),%r12
	and	(%rax),%r13
	and	(%rax),%r14
	and	(%rax),%r15
	nop
	and	(%rcx),%rax
	and	(%rcx),%rcx
	and	(%rcx),%rdx
	and	(%rcx),%rbx
	and	(%rcx),%rsp
	and	(%rcx),%rbp
	and	(%rcx),%rsi
	and	(%rcx),%rdi
	and	(%rcx),%r8
	and	(%rcx),%r9
	and	(%rcx),%r10
	and	(%rcx),%r11
	and	(%rcx),%r12
	and	(%rcx),%r13
	and	(%rcx),%r14
	and	(%rcx),%r15
	nop
	and	(%rdx),%rax
	and	(%rdx),%rcx
	and	(%rdx),%rdx
	and	(%rdx),%rbx
	and	(%rdx),%rsp
	and	(%rdx),%rbp
	and	(%rdx),%rsi
	and	(%rdx),%rdi
	and	(%rdx),%r8
	and	(%rdx),%r9
	and	(%rdx),%r10
	and	(%rdx),%r11
	and	(%rdx),%r12
	and	(%rdx),%r13
	and	(%rdx),%r14
	and	(%rdx),%r15
	nop
	and	(%rbx),%rax
	and	(%rbx),%rcx
	and	(%rbx),%rdx
	and	(%rbx),%rbx
	and	(%rbx),%rsp
	and	(%rbx),%rbp
	and	(%rbx),%rsi
	and	(%rbx),%rdi
	and	(%rbx),%r8
	and	(%rbx),%r9
	and	(%rbx),%r10
	and	(%rbx),%r11
	and	(%rbx),%r12
	and	(%rbx),%r13
	and	(%rbx),%r14
	and	(%rbx),%r15
	nop
	and	(%rsp),%rax
	and	(%rsp),%rcx
	and	(%rsp),%rdx
	and	(%rsp),%rbx
	and	(%rsp),%rsp
	and	(%rsp),%rbp
	and	(%rsp),%rsi
	and	(%rsp),%rdi
	and	(%rsp),%r8
	and	(%rsp),%r9
	and	(%rsp),%r10
	and	(%rsp),%r11
	and	(%rsp),%r12
	and	(%rsp),%r13
	and	(%rsp),%r14
	and	(%rsp),%r15
	nop
	and	(%rbp),%rax
	and	(%rbp),%rcx
	and	(%rbp),%rdx
	and	(%rbp),%rbx
	and	(%rbp),%rsp
	and	(%rbp),%rbp
	and	(%rbp),%rsi
	and	(%rbp),%rdi
	and	(%rbp),%r8
	and	(%rbp),%r9
	and	(%rbp),%r10
	and	(%rbp),%r11
	and	(%rbp),%r12
	and	(%rbp),%r13
	and	(%rbp),%r14
	and	(%rbp),%r15
	nop
	and	(%rsi),%rax
	and	(%rsi),%rcx
	and	(%rsi),%rdx
	and	(%rsi),%rbx
	and	(%rsi),%rsp
	and	(%rsi),%rbp
	and	(%rsi),%rsi
	and	(%rsi),%rdi
	and	(%rsi),%r8
	and	(%rsi),%r9
	and	(%rsi),%r10
	and	(%rsi),%r11
	and	(%rsi),%r12
	and	(%rsi),%r13
	and	(%rsi),%r14
	and	(%rsi),%r15
	nop
	and	(%rdi),%rax
	and	(%rdi),%rcx
	and	(%rdi),%rdx
	and	(%rdi),%rbx
	and	(%rdi),%rsp
	and	(%rdi),%rbp
	and	(%rdi),%rsi
	and	(%rdi),%rdi
	and	(%rdi),%r8
	and	(%rdi),%r9
	and	(%rdi),%r10
	and	(%rdi),%r11
	and	(%rdi),%r12
	and	(%rdi),%r13
	and	(%rdi),%r14
	and	(%rdi),%r15
	nop
	and	(%r8), %rax
	and	(%r8), %rcx
	and	(%r8), %rdx
	and	(%r8), %rbx
	and	(%r8), %rsp
	and	(%r8), %rbp
	and	(%r8), %rsi
	and	(%r8), %rdi
	and	(%r8), %r8
	and	(%r8), %r9
	and	(%r8), %r10
	and	(%r8), %r11
	and	(%r8), %r12
	and	(%r8), %r13
	and	(%r8), %r14
	and	(%r8), %r15
	nop
	and	(%r9), %rax
	and	(%r9), %rcx
	and	(%r9), %rdx
	and	(%r9), %rbx
	and	(%r9), %rsp
	and	(%r9), %rbp
	and	(%r9), %rsi
	and	(%r9), %rdi
	and	(%r9), %r8
	and	(%r9), %r9
	and	(%r9), %r10
	and	(%r9), %r11
	and	(%r9), %r12
	and	(%r9), %r13
	and	(%r9), %r14
	and	(%r9), %r15
	nop
	and	(%r10),%rax
	and	(%r10),%rcx
	and	(%r10),%rdx
	and	(%r10),%rbx
	and	(%r10),%rsp
	and	(%r10),%rbp
	and	(%r10),%rsi
	and	(%r10),%rdi
	and	(%r10),%r8
	and	(%r10),%r9
	and	(%r10),%r10
	and	(%r10),%r11
	and	(%r10),%r12
	and	(%r10),%r13
	and	(%r10),%r14
	and	(%r10),%r15
	nop
	and	(%r11),%rax
	and	(%r11),%rcx
	and	(%r11),%rdx
	and	(%r11),%rbx
	and	(%r11),%rsp
	and	(%r11),%rbp
	and	(%r11),%rsi
	and	(%r11),%rdi
	and	(%r11),%r8
	and	(%r11),%r9
	and	(%r11),%r10
	and	(%r11),%r11
	and	(%r11),%r12
	and	(%r11),%r13
	and	(%r11),%r14
	and	(%r11),%r15
	nop
	and	(%r12),%rax
	and	(%r12),%rcx
	and	(%r12),%rdx
	and	(%r12),%rbx
	and	(%r12),%rsp
	and	(%r12),%rbp
	and	(%r12),%rsi
	and	(%r12),%rdi
	and	(%r12),%r8
	and	(%r12),%r9
	and	(%r12),%r10
	and	(%r12),%r11
	and	(%r12),%r12
	and	(%r12),%r13
	and	(%r12),%r14
	and	(%r12),%r15
	nop
	and	(%r13),%rax
	and	(%r13),%rcx
	and	(%r13),%rdx
	and	(%r13),%rbx
	and	(%r13),%rsp
	and	(%r13),%rbp
	and	(%r13),%rsi
	and	(%r13),%rdi
	and	(%r13),%r8
	and	(%r13),%r9
	and	(%r13),%r10
	and	(%r13),%r11
	and	(%r13),%r12
	and	(%r13),%r13
	and	(%r13),%r14
	and	(%r13),%r15
	nop
	and	(%r14),%rax
	and	(%r14),%rcx
	and	(%r14),%rdx
	and	(%r14),%rbx
	and	(%r14),%rsp
	and	(%r14),%rbp
	and	(%r14),%rsi
	and	(%r14),%rdi
	and	(%r14),%r8
	and	(%r14),%r9
	and	(%r14),%r10
	and	(%r14),%r11
	and	(%r14),%r12
	and	(%r14),%r13
	and	(%r14),%r14
	and	(%r14),%r15
	nop
	and	(%r15),%rax
	and	(%r15),%rcx
	and	(%r15),%rdx
	and	(%r15),%rbx
	and	(%r15),%rsp
	and	(%r15),%rbp
	and	(%r15),%rsi
	and	(%r15),%rdi
	and	(%r15),%r8
	and	(%r15),%r9
	and	(%r15),%r10
	and	(%r15),%r11
	and	(%r15),%r12
	and	(%r15),%r13
	and	(%r15),%r14
	and	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndRegMem8
	.type	AndRegMem8, @function
AndRegMem8:
	.cfi_startproc
	and	0x7f(%rax),%rax
	and	0x7f(%rax),%rcx
	and	0x7f(%rax),%rdx
	and	0x7f(%rax),%rbx
	and	0x7f(%rax),%rsp
	and	0x7f(%rax),%rbp
	and	0x7f(%rax),%rsi
	and	0x7f(%rax),%rdi
	and	0x7f(%rax),%r8
	and	0x7f(%rax),%r9
	and	0x7f(%rax),%r10
	and	0x7f(%rax),%r11
	and	0x7f(%rax),%r12
	and	0x7f(%rax),%r13
	and	0x7f(%rax),%r14
	and	0x7f(%rax),%r15
	nop
	and	0x7f(%rcx),%rax
	and	0x7f(%rcx),%rcx
	and	0x7f(%rcx),%rdx
	and	0x7f(%rcx),%rbx
	and	0x7f(%rcx),%rsp
	and	0x7f(%rcx),%rbp
	and	0x7f(%rcx),%rsi
	and	0x7f(%rcx),%rdi
	and	0x7f(%rcx),%r8
	and	0x7f(%rcx),%r9
	and	0x7f(%rcx),%r10
	and	0x7f(%rcx),%r11
	and	0x7f(%rcx),%r12
	and	0x7f(%rcx),%r13
	and	0x7f(%rcx),%r14
	and	0x7f(%rcx),%r15
	nop
	and	0x7f(%rdx),%rax
	and	0x7f(%rdx),%rcx
	and	0x7f(%rdx),%rdx
	and	0x7f(%rdx),%rbx
	and	0x7f(%rdx),%rsp
	and	0x7f(%rdx),%rbp
	and	0x7f(%rdx),%rsi
	and	0x7f(%rdx),%rdi
	and	0x7f(%rdx),%r8
	and	0x7f(%rdx),%r9
	and	0x7f(%rdx),%r10
	and	0x7f(%rdx),%r11
	and	0x7f(%rdx),%r12
	and	0x7f(%rdx),%r13
	and	0x7f(%rdx),%r14
	and	0x7f(%rdx),%r15
	nop
	and	0x7f(%rbx),%rax
	and	0x7f(%rbx),%rcx
	and	0x7f(%rbx),%rdx
	and	0x7f(%rbx),%rbx
	and	0x7f(%rbx),%rsp
	and	0x7f(%rbx),%rbp
	and	0x7f(%rbx),%rsi
	and	0x7f(%rbx),%rdi
	and	0x7f(%rbx),%r8
	and	0x7f(%rbx),%r9
	and	0x7f(%rbx),%r10
	and	0x7f(%rbx),%r11
	and	0x7f(%rbx),%r12
	and	0x7f(%rbx),%r13
	and	0x7f(%rbx),%r14
	and	0x7f(%rbx),%r15
	nop
	and	0x7f(%rsp),%rax
	and	0x7f(%rsp),%rcx
	and	0x7f(%rsp),%rdx
	and	0x7f(%rsp),%rbx
	and	0x7f(%rsp),%rsp
	and	0x7f(%rsp),%rbp
	and	0x7f(%rsp),%rsi
	and	0x7f(%rsp),%rdi
	and	0x7f(%rsp),%r8
	and	0x7f(%rsp),%r9
	and	0x7f(%rsp),%r10
	and	0x7f(%rsp),%r11
	and	0x7f(%rsp),%r12
	and	0x7f(%rsp),%r13
	and	0x7f(%rsp),%r14
	and	0x7f(%rsp),%r15
	nop
	and	0x7f(%rbp),%rax
	and	0x7f(%rbp),%rcx
	and	0x7f(%rbp),%rdx
	and	0x7f(%rbp),%rbx
	and	0x7f(%rbp),%rsp
	and	0x7f(%rbp),%rbp
	and	0x7f(%rbp),%rsi
	and	0x7f(%rbp),%rdi
	and	0x7f(%rbp),%r8
	and	0x7f(%rbp),%r9
	and	0x7f(%rbp),%r10
	and	0x7f(%rbp),%r11
	and	0x7f(%rbp),%r12
	and	0x7f(%rbp),%r13
	and	0x7f(%rbp),%r14
	and	0x7f(%rbp),%r15
	nop
	and	0x7f(%rsi),%rax
	and	0x7f(%rsi),%rcx
	and	0x7f(%rsi),%rdx
	and	0x7f(%rsi),%rbx
	and	0x7f(%rsi),%rsp
	and	0x7f(%rsi),%rbp
	and	0x7f(%rsi),%rsi
	and	0x7f(%rsi),%rdi
	and	0x7f(%rsi),%r8
	and	0x7f(%rsi),%r9
	and	0x7f(%rsi),%r10
	and	0x7f(%rsi),%r11
	and	0x7f(%rsi),%r12
	and	0x7f(%rsi),%r13
	and	0x7f(%rsi),%r14
	and	0x7f(%rsi),%r15
	nop
	and	0x7f(%rdi),%rax
	and	0x7f(%rdi),%rcx
	and	0x7f(%rdi),%rdx
	and	0x7f(%rdi),%rbx
	and	0x7f(%rdi),%rsp
	and	0x7f(%rdi),%rbp
	and	0x7f(%rdi),%rsi
	and	0x7f(%rdi),%rdi
	and	0x7f(%rdi),%r8
	and	0x7f(%rdi),%r9
	and	0x7f(%rdi),%r10
	and	0x7f(%rdi),%r11
	and	0x7f(%rdi),%r12
	and	0x7f(%rdi),%r13
	and	0x7f(%rdi),%r14
	and	0x7f(%rdi),%r15
	nop
	and	0x7f(%r8), %rax
	and	0x7f(%r8), %rcx
	and	0x7f(%r8), %rdx
	and	0x7f(%r8), %rbx
	and	0x7f(%r8), %rsp
	and	0x7f(%r8), %rbp
	and	0x7f(%r8), %rsi
	and	0x7f(%r8), %rdi
	and	0x7f(%r8), %r8
	and	0x7f(%r8), %r9
	and	0x7f(%r8), %r10
	and	0x7f(%r8), %r11
	and	0x7f(%r8), %r12
	and	0x7f(%r8), %r13
	and	0x7f(%r8), %r14
	and	0x7f(%r8), %r15
	nop
	and	0x7f(%r9), %rax
	and	0x7f(%r9), %rcx
	and	0x7f(%r9), %rdx
	and	0x7f(%r9), %rbx
	and	0x7f(%r9), %rsp
	and	0x7f(%r9), %rbp
	and	0x7f(%r9), %rsi
	and	0x7f(%r9), %rdi
	and	0x7f(%r9), %r8
	and	0x7f(%r9), %r9
	and	0x7f(%r9), %r10
	and	0x7f(%r9), %r11
	and	0x7f(%r9), %r12
	and	0x7f(%r9), %r13
	and	0x7f(%r9), %r14
	and	0x7f(%r9), %r15
	nop
	and	0x7f(%r10),%rax
	and	0x7f(%r10),%rcx
	and	0x7f(%r10),%rdx
	and	0x7f(%r10),%rbx
	and	0x7f(%r10),%rsp
	and	0x7f(%r10),%rbp
	and	0x7f(%r10),%rsi
	and	0x7f(%r10),%rdi
	and	0x7f(%r10),%r8
	and	0x7f(%r10),%r9
	and	0x7f(%r10),%r10
	and	0x7f(%r10),%r11
	and	0x7f(%r10),%r12
	and	0x7f(%r10),%r13
	and	0x7f(%r10),%r14
	and	0x7f(%r10),%r15
	nop
	and	0x7f(%r11),%rax
	and	0x7f(%r11),%rcx
	and	0x7f(%r11),%rdx
	and	0x7f(%r11),%rbx
	and	0x7f(%r11),%rsp
	and	0x7f(%r11),%rbp
	and	0x7f(%r11),%rsi
	and	0x7f(%r11),%rdi
	and	0x7f(%r11),%r8
	and	0x7f(%r11),%r9
	and	0x7f(%r11),%r10
	and	0x7f(%r11),%r11
	and	0x7f(%r11),%r12
	and	0x7f(%r11),%r13
	and	0x7f(%r11),%r14
	and	0x7f(%r11),%r15
	nop
	and	0x7f(%r12),%rax
	and	0x7f(%r12),%rcx
	and	0x7f(%r12),%rdx
	and	0x7f(%r12),%rbx
	and	0x7f(%r12),%rsp
	and	0x7f(%r12),%rbp
	and	0x7f(%r12),%rsi
	and	0x7f(%r12),%rdi
	and	0x7f(%r12),%r8
	and	0x7f(%r12),%r9
	and	0x7f(%r12),%r10
	and	0x7f(%r12),%r11
	and	0x7f(%r12),%r12
	and	0x7f(%r12),%r13
	and	0x7f(%r12),%r14
	and	0x7f(%r12),%r15
	nop
	and	0x7f(%r13),%rax
	and	0x7f(%r13),%rcx
	and	0x7f(%r13),%rdx
	and	0x7f(%r13),%rbx
	and	0x7f(%r13),%rsp
	and	0x7f(%r13),%rbp
	and	0x7f(%r13),%rsi
	and	0x7f(%r13),%rdi
	and	0x7f(%r13),%r8
	and	0x7f(%r13),%r9
	and	0x7f(%r13),%r10
	and	0x7f(%r13),%r11
	and	0x7f(%r13),%r12
	and	0x7f(%r13),%r13
	and	0x7f(%r13),%r14
	and	0x7f(%r13),%r15
	nop
	and	0x7f(%r14),%rax
	and	0x7f(%r14),%rcx
	and	0x7f(%r14),%rdx
	and	0x7f(%r14),%rbx
	and	0x7f(%r14),%rsp
	and	0x7f(%r14),%rbp
	and	0x7f(%r14),%rsi
	and	0x7f(%r14),%rdi
	and	0x7f(%r14),%r8
	and	0x7f(%r14),%r9
	and	0x7f(%r14),%r10
	and	0x7f(%r14),%r11
	and	0x7f(%r14),%r12
	and	0x7f(%r14),%r13
	and	0x7f(%r14),%r14
	and	0x7f(%r14),%r15
	nop
	and	0x7f(%r15),%rax
	and	0x7f(%r15),%rcx
	and	0x7f(%r15),%rdx
	and	0x7f(%r15),%rbx
	and	0x7f(%r15),%rsp
	and	0x7f(%r15),%rbp
	and	0x7f(%r15),%rsi
	and	0x7f(%r15),%rdi
	and	0x7f(%r15),%r8
	and	0x7f(%r15),%r9
	and	0x7f(%r15),%r10
	and	0x7f(%r15),%r11
	and	0x7f(%r15),%r12
	and	0x7f(%r15),%r13
	and	0x7f(%r15),%r14
	and	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	AndRegMem32
	.type	AndRegMem32, @function
AndRegMem32:
	.cfi_startproc
	and	0x7f563412(%rax),%rax
	and	0x7f563412(%rax),%rcx
	and	0x7f563412(%rax),%rdx
	and	0x7f563412(%rax),%rbx
	and	0x7f563412(%rax),%rsp
	and	0x7f563412(%rax),%rbp
	and	0x7f563412(%rax),%rsi
	and	0x7f563412(%rax),%rdi
	and	0x7f563412(%rax),%r8
	and	0x7f563412(%rax),%r9
	and	0x7f563412(%rax),%r10
	and	0x7f563412(%rax),%r11
	and	0x7f563412(%rax),%r12
	and	0x7f563412(%rax),%r13
	and	0x7f563412(%rax),%r14
	and	0x7f563412(%rax),%r15
	nop
	and	0x7f563412(%rcx),%rax
	and	0x7f563412(%rcx),%rcx
	and	0x7f563412(%rcx),%rdx
	and	0x7f563412(%rcx),%rbx
	and	0x7f563412(%rcx),%rsp
	and	0x7f563412(%rcx),%rbp
	and	0x7f563412(%rcx),%rsi
	and	0x7f563412(%rcx),%rdi
	and	0x7f563412(%rcx),%r8
	and	0x7f563412(%rcx),%r9
	and	0x7f563412(%rcx),%r10
	and	0x7f563412(%rcx),%r11
	and	0x7f563412(%rcx),%r12
	and	0x7f563412(%rcx),%r13
	and	0x7f563412(%rcx),%r14
	and	0x7f563412(%rcx),%r15
	nop
	and	0x7f563412(%rdx),%rax
	and	0x7f563412(%rdx),%rcx
	and	0x7f563412(%rdx),%rdx
	and	0x7f563412(%rdx),%rbx
	and	0x7f563412(%rdx),%rsp
	and	0x7f563412(%rdx),%rbp
	and	0x7f563412(%rdx),%rsi
	and	0x7f563412(%rdx),%rdi
	and	0x7f563412(%rdx),%r8
	and	0x7f563412(%rdx),%r9
	and	0x7f563412(%rdx),%r10
	and	0x7f563412(%rdx),%r11
	and	0x7f563412(%rdx),%r12
	and	0x7f563412(%rdx),%r13
	and	0x7f563412(%rdx),%r14
	and	0x7f563412(%rdx),%r15
	nop
	and	0x7f563412(%rbx),%rax
	and	0x7f563412(%rbx),%rcx
	and	0x7f563412(%rbx),%rdx
	and	0x7f563412(%rbx),%rbx
	and	0x7f563412(%rbx),%rsp
	and	0x7f563412(%rbx),%rbp
	and	0x7f563412(%rbx),%rsi
	and	0x7f563412(%rbx),%rdi
	and	0x7f563412(%rbx),%r8
	and	0x7f563412(%rbx),%r9
	and	0x7f563412(%rbx),%r10
	and	0x7f563412(%rbx),%r11
	and	0x7f563412(%rbx),%r12
	and	0x7f563412(%rbx),%r13
	and	0x7f563412(%rbx),%r14
	and	0x7f563412(%rbx),%r15
	nop
	and	0x7f563412(%rsp),%rax
	and	0x7f563412(%rsp),%rcx
	and	0x7f563412(%rsp),%rdx
	and	0x7f563412(%rsp),%rbx
	and	0x7f563412(%rsp),%rsp
	and	0x7f563412(%rsp),%rbp
	and	0x7f563412(%rsp),%rsi
	and	0x7f563412(%rsp),%rdi
	and	0x7f563412(%rsp),%r8
	and	0x7f563412(%rsp),%r9
	and	0x7f563412(%rsp),%r10
	and	0x7f563412(%rsp),%r11
	and	0x7f563412(%rsp),%r12
	and	0x7f563412(%rsp),%r13
	and	0x7f563412(%rsp),%r14
	and	0x7f563412(%rsp),%r15
	nop
	and	0x7f563412(%rbp),%rax
	and	0x7f563412(%rbp),%rcx
	and	0x7f563412(%rbp),%rdx
	and	0x7f563412(%rbp),%rbx
	and	0x7f563412(%rbp),%rsp
	and	0x7f563412(%rbp),%rbp
	and	0x7f563412(%rbp),%rsi
	and	0x7f563412(%rbp),%rdi
	and	0x7f563412(%rbp),%r8
	and	0x7f563412(%rbp),%r9
	and	0x7f563412(%rbp),%r10
	and	0x7f563412(%rbp),%r11
	and	0x7f563412(%rbp),%r12
	and	0x7f563412(%rbp),%r13
	and	0x7f563412(%rbp),%r14
	and	0x7f563412(%rbp),%r15
	nop
	and	0x7f563412(%rsi),%rax
	and	0x7f563412(%rsi),%rcx
	and	0x7f563412(%rsi),%rdx
	and	0x7f563412(%rsi),%rbx
	and	0x7f563412(%rsi),%rsp
	and	0x7f563412(%rsi),%rbp
	and	0x7f563412(%rsi),%rsi
	and	0x7f563412(%rsi),%rdi
	and	0x7f563412(%rsi),%r8
	and	0x7f563412(%rsi),%r9
	and	0x7f563412(%rsi),%r10
	and	0x7f563412(%rsi),%r11
	and	0x7f563412(%rsi),%r12
	and	0x7f563412(%rsi),%r13
	and	0x7f563412(%rsi),%r14
	and	0x7f563412(%rsi),%r15
	nop
	and	0x7f563412(%rdi),%rax
	and	0x7f563412(%rdi),%rcx
	and	0x7f563412(%rdi),%rdx
	and	0x7f563412(%rdi),%rbx
	and	0x7f563412(%rdi),%rsp
	and	0x7f563412(%rdi),%rbp
	and	0x7f563412(%rdi),%rsi
	and	0x7f563412(%rdi),%rdi
	and	0x7f563412(%rdi),%r8
	and	0x7f563412(%rdi),%r9
	and	0x7f563412(%rdi),%r10
	and	0x7f563412(%rdi),%r11
	and	0x7f563412(%rdi),%r12
	and	0x7f563412(%rdi),%r13
	and	0x7f563412(%rdi),%r14
	and	0x7f563412(%rdi),%r15
	nop
	and	0x7f563412(%r8), %rax
	and	0x7f563412(%r8), %rcx
	and	0x7f563412(%r8), %rdx
	and	0x7f563412(%r8), %rbx
	and	0x7f563412(%r8), %rsp
	and	0x7f563412(%r8), %rbp
	and	0x7f563412(%r8), %rsi
	and	0x7f563412(%r8), %rdi
	and	0x7f563412(%r8), %r8
	and	0x7f563412(%r8), %r9
	and	0x7f563412(%r8), %r10
	and	0x7f563412(%r8), %r11
	and	0x7f563412(%r8), %r12
	and	0x7f563412(%r8), %r13
	and	0x7f563412(%r8), %r14
	and	0x7f563412(%r8), %r15
	nop
	and	0x7f563412(%r9), %rax
	and	0x7f563412(%r9), %rcx
	and	0x7f563412(%r9), %rdx
	and	0x7f563412(%r9), %rbx
	and	0x7f563412(%r9), %rsp
	and	0x7f563412(%r9), %rbp
	and	0x7f563412(%r9), %rsi
	and	0x7f563412(%r9), %rdi
	and	0x7f563412(%r9), %r8
	and	0x7f563412(%r9), %r9
	and	0x7f563412(%r9), %r10
	and	0x7f563412(%r9), %r11
	and	0x7f563412(%r9), %r12
	and	0x7f563412(%r9), %r13
	and	0x7f563412(%r9), %r14
	and	0x7f563412(%r9), %r15
	nop
	and	0x7f563412(%r10),%rax
	and	0x7f563412(%r10),%rcx
	and	0x7f563412(%r10),%rdx
	and	0x7f563412(%r10),%rbx
	and	0x7f563412(%r10),%rsp
	and	0x7f563412(%r10),%rbp
	and	0x7f563412(%r10),%rsi
	and	0x7f563412(%r10),%rdi
	and	0x7f563412(%r10),%r8
	and	0x7f563412(%r10),%r9
	and	0x7f563412(%r10),%r10
	and	0x7f563412(%r10),%r11
	and	0x7f563412(%r10),%r12
	and	0x7f563412(%r10),%r13
	and	0x7f563412(%r10),%r14
	and	0x7f563412(%r10),%r15
	nop
	and	0x7f563412(%r11),%rax
	and	0x7f563412(%r11),%rcx
	and	0x7f563412(%r11),%rdx
	and	0x7f563412(%r11),%rbx
	and	0x7f563412(%r11),%rsp
	and	0x7f563412(%r11),%rbp
	and	0x7f563412(%r11),%rsi
	and	0x7f563412(%r11),%rdi
	and	0x7f563412(%r11),%r8
	and	0x7f563412(%r11),%r9
	and	0x7f563412(%r11),%r10
	and	0x7f563412(%r11),%r11
	and	0x7f563412(%r11),%r12
	and	0x7f563412(%r11),%r13
	and	0x7f563412(%r11),%r14
	and	0x7f563412(%r11),%r15
	nop
	and	0x7f563412(%r12),%rax
	and	0x7f563412(%r12),%rcx
	and	0x7f563412(%r12),%rdx
	and	0x7f563412(%r12),%rbx
	and	0x7f563412(%r12),%rsp
	and	0x7f563412(%r12),%rbp
	and	0x7f563412(%r12),%rsi
	and	0x7f563412(%r12),%rdi
	and	0x7f563412(%r12),%r8
	and	0x7f563412(%r12),%r9
	and	0x7f563412(%r12),%r10
	and	0x7f563412(%r12),%r11
	and	0x7f563412(%r12),%r12
	and	0x7f563412(%r12),%r13
	and	0x7f563412(%r12),%r14
	and	0x7f563412(%r12),%r15
	nop
	and	0x7f563412(%r13),%rax
	and	0x7f563412(%r13),%rcx
	and	0x7f563412(%r13),%rdx
	and	0x7f563412(%r13),%rbx
	and	0x7f563412(%r13),%rsp
	and	0x7f563412(%r13),%rbp
	and	0x7f563412(%r13),%rsi
	and	0x7f563412(%r13),%rdi
	and	0x7f563412(%r13),%r8
	and	0x7f563412(%r13),%r9
	and	0x7f563412(%r13),%r10
	and	0x7f563412(%r13),%r11
	and	0x7f563412(%r13),%r12
	and	0x7f563412(%r13),%r13
	and	0x7f563412(%r13),%r14
	and	0x7f563412(%r13),%r15
	nop
	and	0x7f563412(%r14),%rax
	and	0x7f563412(%r14),%rcx
	and	0x7f563412(%r14),%rdx
	and	0x7f563412(%r14),%rbx
	and	0x7f563412(%r14),%rsp
	and	0x7f563412(%r14),%rbp
	and	0x7f563412(%r14),%rsi
	and	0x7f563412(%r14),%rdi
	and	0x7f563412(%r14),%r8
	and	0x7f563412(%r14),%r9
	and	0x7f563412(%r14),%r10
	and	0x7f563412(%r14),%r11
	and	0x7f563412(%r14),%r12
	and	0x7f563412(%r14),%r13
	and	0x7f563412(%r14),%r14
	and	0x7f563412(%r14),%r15
	nop
	and	0x7f563412(%r15),%rax
	and	0x7f563412(%r15),%rcx
	and	0x7f563412(%r15),%rdx
	and	0x7f563412(%r15),%rbx
	and	0x7f563412(%r15),%rsp
	and	0x7f563412(%r15),%rbp
	and	0x7f563412(%r15),%rsi
	and	0x7f563412(%r15),%rdi
	and	0x7f563412(%r15),%r8
	and	0x7f563412(%r15),%r9
	and	0x7f563412(%r15),%r10
	and	0x7f563412(%r15),%r11
	and	0x7f563412(%r15),%r12
	and	0x7f563412(%r15),%r13
	and	0x7f563412(%r15),%r14
	and	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


