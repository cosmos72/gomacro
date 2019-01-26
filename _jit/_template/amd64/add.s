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
	.globl	Add
	.type	Add, @function
Add:
	.cfi_startproc

        // reg8 += reg8
	add	%al,%al
	add	%al,%cl
	add	%al,%dl
	add	%al,%bl
	add	%al,%spl
	add	%al,%bpl
	add	%al,%sil
	add	%al,%dil
	add	%al,%r8b
	add	%al,%r9b
	add	%al,%r10b
	add	%al,%r11b
	add	%al,%r12b
	add	%al,%r13b
	add	%al,%r14b
	add	%al,%r15b
	nop
	add	%cl,%al
	add	%cl,%cl
	add	%cl,%dl
	add	%cl,%bl
	add	%cl,%spl
	add	%cl,%bpl
	add	%cl,%sil
	add	%cl,%dil
	add	%cl,%r8b
	add	%cl,%r9b
	add	%cl,%r10b
	add	%cl,%r11b
	add	%cl,%r12b
	add	%cl,%r13b
	add	%cl,%r14b
	add	%cl,%r15b
	nop
	add	%dl,%al
	add	%dl,%cl
	add	%dl,%dl
	add	%dl,%bl
	add	%dl,%spl
	add	%dl,%bpl
	add	%dl,%sil
	add	%dl,%dil
	add	%dl,%r8b
	add	%dl,%r9b
	add	%dl,%r10b
	add	%dl,%r11b
	add	%dl,%r12b
	add	%dl,%r13b
	add	%dl,%r14b
	add	%dl,%r15b
	nop
	add	%bl,%al
	add	%bl,%cl
	add	%bl,%dl
	add	%bl,%bl
	add	%bl,%spl
	add	%bl,%bpl
	add	%bl,%sil
	add	%bl,%dil
	add	%bl,%r8b
	add	%bl,%r9b
	add	%bl,%r10b
	add	%bl,%r11b
	add	%bl,%r12b
	add	%bl,%r13b
	add	%bl,%r14b
	add	%bl,%r15b
	nop
	add	%spl,%al
	add	%spl,%cl
	add	%spl,%dl
	add	%spl,%bl
	add	%spl,%spl
	add	%spl,%bpl
	add	%spl,%sil
	add	%spl,%dil
	add	%spl,%r8b
	add	%spl,%r9b
	add	%spl,%r10b
	add	%spl,%r11b
	add	%spl,%r12b
	add	%spl,%r13b
	add	%spl,%r14b
	add	%spl,%r15b
	nop
	add	%bpl,%al
	add	%bpl,%cl
	add	%bpl,%dl
	add	%bpl,%bl
	add	%bpl,%spl
	add	%bpl,%bpl
	add	%bpl,%sil
	add	%bpl,%dil
	add	%bpl,%r8b
	add	%bpl,%r9b
	add	%bpl,%r10b
	add	%bpl,%r11b
	add	%bpl,%r12b
	add	%bpl,%r13b
	add	%bpl,%r14b
	add	%bpl,%r15b
	nop
	add	%sil,%al
	add	%sil,%cl
	add	%sil,%dl
	add	%sil,%bl
	add	%sil,%spl
	add	%sil,%bpl
	add	%sil,%sil
	add	%sil,%dil
	add	%sil,%r8b
	add	%sil,%r9b
	add	%sil,%r10b
	add	%sil,%r11b
	add	%sil,%r12b
	add	%sil,%r13b
	add	%sil,%r14b
	add	%sil,%r15b
	nop
	add	%dil,%al
	add	%dil,%cl
	add	%dil,%dl
	add	%dil,%bl
	add	%dil,%spl
	add	%dil,%bpl
	add	%dil,%sil
	add	%dil,%dil
	add	%dil,%r8b
	add	%dil,%r9b
	add	%dil,%r10b
	add	%dil,%r11b
	add	%dil,%r12b
	add	%dil,%r13b
	add	%dil,%r14b
	add	%dil,%r15b
	nop
	add	%r8b, %al
	add	%r8b, %cl
	add	%r8b, %dl
	add	%r8b, %bl
	add	%r8b, %spl
	add	%r8b, %bpl
	add	%r8b, %sil
	add	%r8b, %dil
	add	%r8b, %r8b
	add	%r8b, %r9b
	add	%r8b, %r10b
	add	%r8b, %r11b
	add	%r8b, %r12b
	add	%r8b, %r13b
	add	%r8b, %r14b
	add	%r8b, %r15b
	nop
	add	%r9b, %al
	add	%r9b, %cl
	add	%r9b, %dl
	add	%r9b, %bl
	add	%r9b, %spl
	add	%r9b, %bpl
	add	%r9b, %sil
	add	%r9b, %dil
	add	%r9b, %r8b
	add	%r9b, %r9b
	add	%r9b, %r10b
	add	%r9b, %r11b
	add	%r9b, %r12b
	add	%r9b, %r13b
	add	%r9b, %r14b
	add	%r9b, %r15b
	nop
	add	%r10b,%al
	add	%r10b,%cl
	add	%r10b,%dl
	add	%r10b,%bl
	add	%r10b,%spl
	add	%r10b,%bpl
	add	%r10b,%sil
	add	%r10b,%dil
	add	%r10b,%r8b
	add	%r10b,%r9b
	add	%r10b,%r10b
	add	%r10b,%r11b
	add	%r10b,%r12b
	add	%r10b,%r13b
	add	%r10b,%r14b
	add	%r10b,%r15b
	nop
	add	%r11b,%al
	add	%r11b,%cl
	add	%r11b,%dl
	add	%r11b,%bl
	add	%r11b,%spl
	add	%r11b,%bpl
	add	%r11b,%sil
	add	%r11b,%dil
	add	%r11b,%r8b
	add	%r11b,%r9b
	add	%r11b,%r10b
	add	%r11b,%r11b
	add	%r11b,%r12b
	add	%r11b,%r13b
	add	%r11b,%r14b
	add	%r11b,%r15b
	nop
	add	%r12b,%al
	add	%r12b,%cl
	add	%r12b,%dl
	add	%r12b,%bl
	add	%r12b,%spl
	add	%r12b,%bpl
	add	%r12b,%sil
	add	%r12b,%dil
	add	%r12b,%r8b
	add	%r12b,%r9b
	add	%r12b,%r10b
	add	%r12b,%r11b
	add	%r12b,%r12b
	add	%r12b,%r13b
	add	%r12b,%r14b
	add	%r12b,%r15b
	nop
	add	%r13b,%al
	add	%r13b,%cl
	add	%r13b,%dl
	add	%r13b,%bl
	add	%r13b,%spl
	add	%r13b,%bpl
	add	%r13b,%sil
	add	%r13b,%dil
	add	%r13b,%r8b
	add	%r13b,%r9b
	add	%r13b,%r10b
	add	%r13b,%r11b
	add	%r13b,%r12b
	add	%r13b,%r13b
	add	%r13b,%r14b
	add	%r13b,%r15b
	nop
	add	%r14b,%al
	add	%r14b,%cl
	add	%r14b,%dl
	add	%r14b,%bl
	add	%r14b,%spl
	add	%r14b,%bpl
	add	%r14b,%sil
	add	%r14b,%dil
	add	%r14b,%r8b
	add	%r14b,%r9b
	add	%r14b,%r10b
	add	%r14b,%r11b
	add	%r14b,%r12b
	add	%r14b,%r13b
	add	%r14b,%r14b
	add	%r14b,%r15b
	nop
	add	%r15b,%al
	add	%r15b,%cl
	add	%r15b,%dl
	add	%r15b,%bl
	add	%r15b,%spl
	add	%r15b,%bpl
	add	%r15b,%sil
	add	%r15b,%dil
	add	%r15b,%r8b
	add	%r15b,%r9b
	add	%r15b,%r10b
	add	%r15b,%r11b
	add	%r15b,%r12b
	add	%r15b,%r13b
	add	%r15b,%r14b
	add	%r15b,%r15b
        nop
        nop
        // reg16 += reg16
	add	%ax,%ax
	add	%ax,%cx
	add	%ax,%dx
	add	%ax,%bx
	add	%ax,%sp
	add	%ax,%bp
	add	%ax,%si
	add	%ax,%di
	add	%ax,%r8w
	add	%ax,%r9w
	add	%ax,%r10w
	add	%ax,%r11w
	add	%ax,%r12w
	add	%ax,%r13w
	add	%ax,%r14w
	add	%ax,%r15w
	nop
	add	%cx,%ax
	add	%cx,%cx
	add	%cx,%dx
	add	%cx,%bx
	add	%cx,%sp
	add	%cx,%bp
	add	%cx,%si
	add	%cx,%di
	add	%cx,%r8w
	add	%cx,%r9w
	add	%cx,%r10w
	add	%cx,%r11w
	add	%cx,%r12w
	add	%cx,%r13w
	add	%cx,%r14w
	add	%cx,%r15w
	nop
	add	%dx,%ax
	add	%dx,%cx
	add	%dx,%dx
	add	%dx,%bx
	add	%dx,%sp
	add	%dx,%bp
	add	%dx,%si
	add	%dx,%di
	add	%dx,%r8w
	add	%dx,%r9w
	add	%dx,%r10w
	add	%dx,%r11w
	add	%dx,%r12w
	add	%dx,%r13w
	add	%dx,%r14w
	add	%dx,%r15w
	nop
	add	%bx,%ax
	add	%bx,%cx
	add	%bx,%dx
	add	%bx,%bx
	add	%bx,%sp
	add	%bx,%bp
	add	%bx,%si
	add	%bx,%di
	add	%bx,%r8w
	add	%bx,%r9w
	add	%bx,%r10w
	add	%bx,%r11w
	add	%bx,%r12w
	add	%bx,%r13w
	add	%bx,%r14w
	add	%bx,%r15w
	nop
	add	%sp,%ax
	add	%sp,%cx
	add	%sp,%dx
	add	%sp,%bx
	add	%sp,%sp
	add	%sp,%bp
	add	%sp,%si
	add	%sp,%di
	add	%sp,%r8w
	add	%sp,%r9w
	add	%sp,%r10w
	add	%sp,%r11w
	add	%sp,%r12w
	add	%sp,%r13w
	add	%sp,%r14w
	add	%sp,%r15w
	nop
	add	%bp,%ax
	add	%bp,%cx
	add	%bp,%dx
	add	%bp,%bx
	add	%bp,%sp
	add	%bp,%bp
	add	%bp,%si
	add	%bp,%di
	add	%bp,%r8w
	add	%bp,%r9w
	add	%bp,%r10w
	add	%bp,%r11w
	add	%bp,%r12w
	add	%bp,%r13w
	add	%bp,%r14w
	add	%bp,%r15w
	nop
	add	%si,%ax
	add	%si,%cx
	add	%si,%dx
	add	%si,%bx
	add	%si,%sp
	add	%si,%bp
	add	%si,%si
	add	%si,%di
	add	%si,%r8w
	add	%si,%r9w
	add	%si,%r10w
	add	%si,%r11w
	add	%si,%r12w
	add	%si,%r13w
	add	%si,%r14w
	add	%si,%r15w
	nop
	add	%di,%ax
	add	%di,%cx
	add	%di,%dx
	add	%di,%bx
	add	%di,%sp
	add	%di,%bp
	add	%di,%si
	add	%di,%di
	add	%di,%r8w
	add	%di,%r9w
	add	%di,%r10w
	add	%di,%r11w
	add	%di,%r12w
	add	%di,%r13w
	add	%di,%r14w
	add	%di,%r15w
	nop
	add	%r8w, %ax
	add	%r8w, %cx
	add	%r8w, %dx
	add	%r8w, %bx
	add	%r8w, %sp
	add	%r8w, %bp
	add	%r8w, %si
	add	%r8w, %di
	add	%r8w, %r8w
	add	%r8w, %r9w
	add	%r8w, %r10w
	add	%r8w, %r11w
	add	%r8w, %r12w
	add	%r8w, %r13w
	add	%r8w, %r14w
	add	%r8w, %r15w
	nop
	add	%r9w, %ax
	add	%r9w, %cx
	add	%r9w, %dx
	add	%r9w, %bx
	add	%r9w, %sp
	add	%r9w, %bp
	add	%r9w, %si
	add	%r9w, %di
	add	%r9w, %r8w
	add	%r9w, %r9w
	add	%r9w, %r10w
	add	%r9w, %r11w
	add	%r9w, %r12w
	add	%r9w, %r13w
	add	%r9w, %r14w
	add	%r9w, %r15w
	nop
	add	%r10w,%ax
	add	%r10w,%cx
	add	%r10w,%dx
	add	%r10w,%bx
	add	%r10w,%sp
	add	%r10w,%bp
	add	%r10w,%si
	add	%r10w,%di
	add	%r10w,%r8w
	add	%r10w,%r9w
	add	%r10w,%r10w
	add	%r10w,%r11w
	add	%r10w,%r12w
	add	%r10w,%r13w
	add	%r10w,%r14w
	add	%r10w,%r15w
	nop
	add	%r11w,%ax
	add	%r11w,%cx
	add	%r11w,%dx
	add	%r11w,%bx
	add	%r11w,%sp
	add	%r11w,%bp
	add	%r11w,%si
	add	%r11w,%di
	add	%r11w,%r8w
	add	%r11w,%r9w
	add	%r11w,%r10w
	add	%r11w,%r11w
	add	%r11w,%r12w
	add	%r11w,%r13w
	add	%r11w,%r14w
	add	%r11w,%r15w
	nop
	add	%r12w,%ax
	add	%r12w,%cx
	add	%r12w,%dx
	add	%r12w,%bx
	add	%r12w,%sp
	add	%r12w,%bp
	add	%r12w,%si
	add	%r12w,%di
	add	%r12w,%r8w
	add	%r12w,%r9w
	add	%r12w,%r10w
	add	%r12w,%r11w
	add	%r12w,%r12w
	add	%r12w,%r13w
	add	%r12w,%r14w
	add	%r12w,%r15w
	nop
	add	%r13w,%ax
	add	%r13w,%cx
	add	%r13w,%dx
	add	%r13w,%bx
	add	%r13w,%sp
	add	%r13w,%bp
	add	%r13w,%si
	add	%r13w,%di
	add	%r13w,%r8w
	add	%r13w,%r9w
	add	%r13w,%r10w
	add	%r13w,%r11w
	add	%r13w,%r12w
	add	%r13w,%r13w
	add	%r13w,%r14w
	add	%r13w,%r15w
	nop
	add	%r14w,%ax
	add	%r14w,%cx
	add	%r14w,%dx
	add	%r14w,%bx
	add	%r14w,%sp
	add	%r14w,%bp
	add	%r14w,%si
	add	%r14w,%di
	add	%r14w,%r8w
	add	%r14w,%r9w
	add	%r14w,%r10w
	add	%r14w,%r11w
	add	%r14w,%r12w
	add	%r14w,%r13w
	add	%r14w,%r14w
	add	%r14w,%r15w
	nop
	add	%r15w,%ax
	add	%r15w,%cx
	add	%r15w,%dx
	add	%r15w,%bx
	add	%r15w,%sp
	add	%r15w,%bp
	add	%r15w,%si
	add	%r15w,%di
	add	%r15w,%r8w
	add	%r15w,%r9w
	add	%r15w,%r10w
	add	%r15w,%r11w
	add	%r15w,%r12w
	add	%r15w,%r13w
	add	%r15w,%r14w
	add	%r15w,%r15w
        nop
        nop
        // reg32 += reg32
	add	%eax,%eax
	add	%eax,%ecx
	add	%eax,%edx
	add	%eax,%ebx
	add	%eax,%esp
	add	%eax,%ebp
	add	%eax,%esi
	add	%eax,%edi
	add	%eax,%r8d
	add	%eax,%r9d
	add	%eax,%r10d
	add	%eax,%r11d
	add	%eax,%r12d
	add	%eax,%r13d
	add	%eax,%r14d
	add	%eax,%r15d
	nop
	add	%ecx,%eax
	add	%ecx,%ecx
	add	%ecx,%edx
	add	%ecx,%ebx
	add	%ecx,%esp
	add	%ecx,%ebp
	add	%ecx,%esi
	add	%ecx,%edi
	add	%ecx,%r8d
	add	%ecx,%r9d
	add	%ecx,%r10d
	add	%ecx,%r11d
	add	%ecx,%r12d
	add	%ecx,%r13d
	add	%ecx,%r14d
	add	%ecx,%r15d
	nop
	add	%edx,%eax
	add	%edx,%ecx
	add	%edx,%edx
	add	%edx,%ebx
	add	%edx,%esp
	add	%edx,%ebp
	add	%edx,%esi
	add	%edx,%edi
	add	%edx,%r8d
	add	%edx,%r9d
	add	%edx,%r10d
	add	%edx,%r11d
	add	%edx,%r12d
	add	%edx,%r13d
	add	%edx,%r14d
	add	%edx,%r15d
	nop
	add	%ebx,%eax
	add	%ebx,%ecx
	add	%ebx,%edx
	add	%ebx,%ebx
	add	%ebx,%esp
	add	%ebx,%ebp
	add	%ebx,%esi
	add	%ebx,%edi
	add	%ebx,%r8d
	add	%ebx,%r9d
	add	%ebx,%r10d
	add	%ebx,%r11d
	add	%ebx,%r12d
	add	%ebx,%r13d
	add	%ebx,%r14d
	add	%ebx,%r15d
	nop
	add	%esp,%eax
	add	%esp,%ecx
	add	%esp,%edx
	add	%esp,%ebx
	add	%esp,%esp
	add	%esp,%ebp
	add	%esp,%esi
	add	%esp,%edi
	add	%esp,%r8d
	add	%esp,%r9d
	add	%esp,%r10d
	add	%esp,%r11d
	add	%esp,%r12d
	add	%esp,%r13d
	add	%esp,%r14d
	add	%esp,%r15d
	nop
	add	%ebp,%eax
	add	%ebp,%ecx
	add	%ebp,%edx
	add	%ebp,%ebx
	add	%ebp,%esp
	add	%ebp,%ebp
	add	%ebp,%esi
	add	%ebp,%edi
	add	%ebp,%r8d
	add	%ebp,%r9d
	add	%ebp,%r10d
	add	%ebp,%r11d
	add	%ebp,%r12d
	add	%ebp,%r13d
	add	%ebp,%r14d
	add	%ebp,%r15d
	nop
	add	%esi,%eax
	add	%esi,%ecx
	add	%esi,%edx
	add	%esi,%ebx
	add	%esi,%esp
	add	%esi,%ebp
	add	%esi,%esi
	add	%esi,%edi
	add	%esi,%r8d
	add	%esi,%r9d
	add	%esi,%r10d
	add	%esi,%r11d
	add	%esi,%r12d
	add	%esi,%r13d
	add	%esi,%r14d
	add	%esi,%r15d
	nop
	add	%edi,%eax
	add	%edi,%ecx
	add	%edi,%edx
	add	%edi,%ebx
	add	%edi,%esp
	add	%edi,%ebp
	add	%edi,%esi
	add	%edi,%edi
	add	%edi,%r8d
	add	%edi,%r9d
	add	%edi,%r10d
	add	%edi,%r11d
	add	%edi,%r12d
	add	%edi,%r13d
	add	%edi,%r14d
	add	%edi,%r15d
	nop
	add	%r8d, %eax
	add	%r8d, %ecx
	add	%r8d, %edx
	add	%r8d, %ebx
	add	%r8d, %esp
	add	%r8d, %ebp
	add	%r8d, %esi
	add	%r8d, %edi
	add	%r8d, %r8d
	add	%r8d, %r9d
	add	%r8d, %r10d
	add	%r8d, %r11d
	add	%r8d, %r12d
	add	%r8d, %r13d
	add	%r8d, %r14d
	add	%r8d, %r15d
	nop
	add	%r9d, %eax
	add	%r9d, %ecx
	add	%r9d, %edx
	add	%r9d, %ebx
	add	%r9d, %esp
	add	%r9d, %ebp
	add	%r9d, %esi
	add	%r9d, %edi
	add	%r9d, %r8d
	add	%r9d, %r9d
	add	%r9d, %r10d
	add	%r9d, %r11d
	add	%r9d, %r12d
	add	%r9d, %r13d
	add	%r9d, %r14d
	add	%r9d, %r15d
	nop
	add	%r10d,%eax
	add	%r10d,%ecx
	add	%r10d,%edx
	add	%r10d,%ebx
	add	%r10d,%esp
	add	%r10d,%ebp
	add	%r10d,%esi
	add	%r10d,%edi
	add	%r10d,%r8d
	add	%r10d,%r9d
	add	%r10d,%r10d
	add	%r10d,%r11d
	add	%r10d,%r12d
	add	%r10d,%r13d
	add	%r10d,%r14d
	add	%r10d,%r15d
	nop
	add	%r11d,%eax
	add	%r11d,%ecx
	add	%r11d,%edx
	add	%r11d,%ebx
	add	%r11d,%esp
	add	%r11d,%ebp
	add	%r11d,%esi
	add	%r11d,%edi
	add	%r11d,%r8d
	add	%r11d,%r9d
	add	%r11d,%r10d
	add	%r11d,%r11d
	add	%r11d,%r12d
	add	%r11d,%r13d
	add	%r11d,%r14d
	add	%r11d,%r15d
	nop
	add	%r12d,%eax
	add	%r12d,%ecx
	add	%r12d,%edx
	add	%r12d,%ebx
	add	%r12d,%esp
	add	%r12d,%ebp
	add	%r12d,%esi
	add	%r12d,%edi
	add	%r12d,%r8d
	add	%r12d,%r9d
	add	%r12d,%r10d
	add	%r12d,%r11d
	add	%r12d,%r12d
	add	%r12d,%r13d
	add	%r12d,%r14d
	add	%r12d,%r15d
	nop
	add	%r13d,%eax
	add	%r13d,%ecx
	add	%r13d,%edx
	add	%r13d,%ebx
	add	%r13d,%esp
	add	%r13d,%ebp
	add	%r13d,%esi
	add	%r13d,%edi
	add	%r13d,%r8d
	add	%r13d,%r9d
	add	%r13d,%r10d
	add	%r13d,%r11d
	add	%r13d,%r12d
	add	%r13d,%r13d
	add	%r13d,%r14d
	add	%r13d,%r15d
	nop
	add	%r14d,%eax
	add	%r14d,%ecx
	add	%r14d,%edx
	add	%r14d,%ebx
	add	%r14d,%esp
	add	%r14d,%ebp
	add	%r14d,%esi
	add	%r14d,%edi
	add	%r14d,%r8d
	add	%r14d,%r9d
	add	%r14d,%r10d
	add	%r14d,%r11d
	add	%r14d,%r12d
	add	%r14d,%r13d
	add	%r14d,%r14d
	add	%r14d,%r15d
	nop
	add	%r15d,%eax
	add	%r15d,%ecx
	add	%r15d,%edx
	add	%r15d,%ebx
	add	%r15d,%esp
	add	%r15d,%ebp
	add	%r15d,%esi
	add	%r15d,%edi
	add	%r15d,%r8d
	add	%r15d,%r9d
	add	%r15d,%r10d
	add	%r15d,%r11d
	add	%r15d,%r12d
	add	%r15d,%r13d
	add	%r15d,%r14d
	add	%r15d,%r15d
        nop
        nop
        // reg64 += reg64
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
        // mem8 += reg8
	add	%al,(%rax)
	add	%al,(%rcx)
	add	%al,(%rdx)
	add	%al,(%rbx)
	add	%al,(%rsp)
	add	%al,(%rbp)
	add	%al,(%rsi)
	add	%al,(%rdi)
	add	%al,(%r8)
	add	%al,(%r9)
	add	%al,(%r10)
	add	%al,(%r11)
	add	%al,(%r12)
	add	%al,(%r13)
	add	%al,(%r14)
	add	%al,(%r15)
	nop
	add	%cl,(%rax)
	add	%cl,(%rcx)
	add	%cl,(%rdx)
	add	%cl,(%rbx)
	add	%cl,(%rsp)
	add	%cl,(%rbp)
	add	%cl,(%rsi)
	add	%cl,(%rdi)
	add	%cl,(%r8)
	add	%cl,(%r9)
	add	%cl,(%r10)
	add	%cl,(%r11)
	add	%cl,(%r12)
	add	%cl,(%r13)
	add	%cl,(%r14)
	add	%cl,(%r15)
	nop
	add	%dl,(%rax)
	add	%dl,(%rcx)
	add	%dl,(%rdx)
	add	%dl,(%rbx)
	add	%dl,(%rsp)
	add	%dl,(%rbp)
	add	%dl,(%rsi)
	add	%dl,(%rdi)
	add	%dl,(%r8)
	add	%dl,(%r9)
	add	%dl,(%r10)
	add	%dl,(%r11)
	add	%dl,(%r12)
	add	%dl,(%r13)
	add	%dl,(%r14)
	add	%dl,(%r15)
	nop
	add	%bl,(%rax)
	add	%bl,(%rcx)
	add	%bl,(%rdx)
	add	%bl,(%rbx)
	add	%bl,(%rsp)
	add	%bl,(%rbp)
	add	%bl,(%rsi)
	add	%bl,(%rdi)
	add	%bl,(%r8)
	add	%bl,(%r9)
	add	%bl,(%r10)
	add	%bl,(%r11)
	add	%bl,(%r12)
	add	%bl,(%r13)
	add	%bl,(%r14)
	add	%bl,(%r15)
	nop
	add	%spl,(%rax)
	add	%spl,(%rcx)
	add	%spl,(%rdx)
	add	%spl,(%rbx)
	add	%spl,(%rsp)
	add	%spl,(%rbp)
	add	%spl,(%rsi)
	add	%spl,(%rdi)
	add	%spl,(%r8)
	add	%spl,(%r9)
	add	%spl,(%r10)
	add	%spl,(%r11)
	add	%spl,(%r12)
	add	%spl,(%r13)
	add	%spl,(%r14)
	add	%spl,(%r15)
	nop
	add	%bpl,(%rax)
	add	%bpl,(%rcx)
	add	%bpl,(%rdx)
	add	%bpl,(%rbx)
	add	%bpl,(%rsp)
	add	%bpl,(%rbp)
	add	%bpl,(%rsi)
	add	%bpl,(%rdi)
	add	%bpl,(%r8)
	add	%bpl,(%r9)
	add	%bpl,(%r10)
	add	%bpl,(%r11)
	add	%bpl,(%r12)
	add	%bpl,(%r13)
	add	%bpl,(%r14)
	add	%bpl,(%r15)
	nop
	add	%sil,(%rax)
	add	%sil,(%rcx)
	add	%sil,(%rdx)
	add	%sil,(%rbx)
	add	%sil,(%rsp)
	add	%sil,(%rbp)
	add	%sil,(%rsi)
	add	%sil,(%rdi)
	add	%sil,(%r8)
	add	%sil,(%r9)
	add	%sil,(%r10)
	add	%sil,(%r11)
	add	%sil,(%r12)
	add	%sil,(%r13)
	add	%sil,(%r14)
	add	%sil,(%r15)
	nop
	add	%dil,(%rax)
	add	%dil,(%rcx)
	add	%dil,(%rdx)
	add	%dil,(%rbx)
	add	%dil,(%rsp)
	add	%dil,(%rbp)
	add	%dil,(%rsi)
	add	%dil,(%rdi)
	add	%dil,(%r8)
	add	%dil,(%r9)
	add	%dil,(%r10)
	add	%dil,(%r11)
	add	%dil,(%r12)
	add	%dil,(%r13)
	add	%dil,(%r14)
	add	%dil,(%r15)
	nop
	add	%r8b, (%rax)
	add	%r8b, (%rcx)
	add	%r8b, (%rdx)
	add	%r8b, (%rbx)
	add	%r8b, (%rsp)
	add	%r8b, (%rbp)
	add	%r8b, (%rsi)
	add	%r8b, (%rdi)
	add	%r8b, (%r8)
	add	%r8b, (%r9)
	add	%r8b, (%r10)
	add	%r8b, (%r11)
	add	%r8b, (%r12)
	add	%r8b, (%r13)
	add	%r8b, (%r14)
	add	%r8b, (%r15)
	nop
	add	%r9b, (%rax)
	add	%r9b, (%rcx)
	add	%r9b, (%rdx)
	add	%r9b, (%rbx)
	add	%r9b, (%rsp)
	add	%r9b, (%rbp)
	add	%r9b, (%rsi)
	add	%r9b, (%rdi)
	add	%r9b, (%r8)
	add	%r9b, (%r9)
	add	%r9b, (%r10)
	add	%r9b, (%r11)
	add	%r9b, (%r12)
	add	%r9b, (%r13)
	add	%r9b, (%r14)
	add	%r9b, (%r15)
	nop
	add	%r10b,(%rax)
	add	%r10b,(%rcx)
	add	%r10b,(%rdx)
	add	%r10b,(%rbx)
	add	%r10b,(%rsp)
	add	%r10b,(%rbp)
	add	%r10b,(%rsi)
	add	%r10b,(%rdi)
	add	%r10b,(%r8)
	add	%r10b,(%r9)
	add	%r10b,(%r10)
	add	%r10b,(%r11)
	add	%r10b,(%r12)
	add	%r10b,(%r13)
	add	%r10b,(%r14)
	add	%r10b,(%r15)
	nop
	add	%r11b,(%rax)
	add	%r11b,(%rcx)
	add	%r11b,(%rdx)
	add	%r11b,(%rbx)
	add	%r11b,(%rsp)
	add	%r11b,(%rbp)
	add	%r11b,(%rsi)
	add	%r11b,(%rdi)
	add	%r11b,(%r8)
	add	%r11b,(%r9)
	add	%r11b,(%r10)
	add	%r11b,(%r11)
	add	%r11b,(%r12)
	add	%r11b,(%r13)
	add	%r11b,(%r14)
	add	%r11b,(%r15)
	nop
	add	%r12b,(%rax)
	add	%r12b,(%rcx)
	add	%r12b,(%rdx)
	add	%r12b,(%rbx)
	add	%r12b,(%rsp)
	add	%r12b,(%rbp)
	add	%r12b,(%rsi)
	add	%r12b,(%rdi)
	add	%r12b,(%r8)
	add	%r12b,(%r9)
	add	%r12b,(%r10)
	add	%r12b,(%r11)
	add	%r12b,(%r12)
	add	%r12b,(%r13)
	add	%r12b,(%r14)
	add	%r12b,(%r15)
	nop
	add	%r13b,(%rax)
	add	%r13b,(%rcx)
	add	%r13b,(%rdx)
	add	%r13b,(%rbx)
	add	%r13b,(%rsp)
	add	%r13b,(%rbp)
	add	%r13b,(%rsi)
	add	%r13b,(%rdi)
	add	%r13b,(%r8)
	add	%r13b,(%r9)
	add	%r13b,(%r10)
	add	%r13b,(%r11)
	add	%r13b,(%r12)
	add	%r13b,(%r13)
	add	%r13b,(%r14)
	add	%r13b,(%r15)
	nop
	add	%r14b,(%rax)
	add	%r14b,(%rcx)
	add	%r14b,(%rdx)
	add	%r14b,(%rbx)
	add	%r14b,(%rsp)
	add	%r14b,(%rbp)
	add	%r14b,(%rsi)
	add	%r14b,(%rdi)
	add	%r14b,(%r8)
	add	%r14b,(%r9)
	add	%r14b,(%r10)
	add	%r14b,(%r11)
	add	%r14b,(%r12)
	add	%r14b,(%r13)
	add	%r14b,(%r14)
	add	%r14b,(%r15)
	nop
	add	%r15b,(%rax)
	add	%r15b,(%rcx)
	add	%r15b,(%rdx)
	add	%r15b,(%rbx)
	add	%r15b,(%rsp)
	add	%r15b,(%rbp)
	add	%r15b,(%rsi)
	add	%r15b,(%rdi)
	add	%r15b,(%r8)
	add	%r15b,(%r9)
	add	%r15b,(%r10)
	add	%r15b,(%r11)
	add	%r15b,(%r12)
	add	%r15b,(%r13)
	add	%r15b,(%r14)
	add	%r15b,(%r15)
        nop
        nop
        // mem16 += reg16
	add	%ax,(%rax)
	add	%ax,(%rcx)
	add	%ax,(%rdx)
	add	%ax,(%rbx)
	add	%ax,(%rsp)
	add	%ax,(%rbp)
	add	%ax,(%rsi)
	add	%ax,(%rdi)
	add	%ax,(%r8)
	add	%ax,(%r9)
	add	%ax,(%r10)
	add	%ax,(%r11)
	add	%ax,(%r12)
	add	%ax,(%r13)
	add	%ax,(%r14)
	add	%ax,(%r15)
	nop
	add	%cx,(%rax)
	add	%cx,(%rcx)
	add	%cx,(%rdx)
	add	%cx,(%rbx)
	add	%cx,(%rsp)
	add	%cx,(%rbp)
	add	%cx,(%rsi)
	add	%cx,(%rdi)
	add	%cx,(%r8)
	add	%cx,(%r9)
	add	%cx,(%r10)
	add	%cx,(%r11)
	add	%cx,(%r12)
	add	%cx,(%r13)
	add	%cx,(%r14)
	add	%cx,(%r15)
	nop
	add	%dx,(%rax)
	add	%dx,(%rcx)
	add	%dx,(%rdx)
	add	%dx,(%rbx)
	add	%dx,(%rsp)
	add	%dx,(%rbp)
	add	%dx,(%rsi)
	add	%dx,(%rdi)
	add	%dx,(%r8)
	add	%dx,(%r9)
	add	%dx,(%r10)
	add	%dx,(%r11)
	add	%dx,(%r12)
	add	%dx,(%r13)
	add	%dx,(%r14)
	add	%dx,(%r15)
	nop
	add	%bx,(%rax)
	add	%bx,(%rcx)
	add	%bx,(%rdx)
	add	%bx,(%rbx)
	add	%bx,(%rsp)
	add	%bx,(%rbp)
	add	%bx,(%rsi)
	add	%bx,(%rdi)
	add	%bx,(%r8)
	add	%bx,(%r9)
	add	%bx,(%r10)
	add	%bx,(%r11)
	add	%bx,(%r12)
	add	%bx,(%r13)
	add	%bx,(%r14)
	add	%bx,(%r15)
	nop
	add	%sp,(%rax)
	add	%sp,(%rcx)
	add	%sp,(%rdx)
	add	%sp,(%rbx)
	add	%sp,(%rsp)
	add	%sp,(%rbp)
	add	%sp,(%rsi)
	add	%sp,(%rdi)
	add	%sp,(%r8)
	add	%sp,(%r9)
	add	%sp,(%r10)
	add	%sp,(%r11)
	add	%sp,(%r12)
	add	%sp,(%r13)
	add	%sp,(%r14)
	add	%sp,(%r15)
	nop
	add	%bp,(%rax)
	add	%bp,(%rcx)
	add	%bp,(%rdx)
	add	%bp,(%rbx)
	add	%bp,(%rsp)
	add	%bp,(%rbp)
	add	%bp,(%rsi)
	add	%bp,(%rdi)
	add	%bp,(%r8)
	add	%bp,(%r9)
	add	%bp,(%r10)
	add	%bp,(%r11)
	add	%bp,(%r12)
	add	%bp,(%r13)
	add	%bp,(%r14)
	add	%bp,(%r15)
	nop
	add	%si,(%rax)
	add	%si,(%rcx)
	add	%si,(%rdx)
	add	%si,(%rbx)
	add	%si,(%rsp)
	add	%si,(%rbp)
	add	%si,(%rsi)
	add	%si,(%rdi)
	add	%si,(%r8)
	add	%si,(%r9)
	add	%si,(%r10)
	add	%si,(%r11)
	add	%si,(%r12)
	add	%si,(%r13)
	add	%si,(%r14)
	add	%si,(%r15)
	nop
	add	%di,(%rax)
	add	%di,(%rcx)
	add	%di,(%rdx)
	add	%di,(%rbx)
	add	%di,(%rsp)
	add	%di,(%rbp)
	add	%di,(%rsi)
	add	%di,(%rdi)
	add	%di,(%r8)
	add	%di,(%r9)
	add	%di,(%r10)
	add	%di,(%r11)
	add	%di,(%r12)
	add	%di,(%r13)
	add	%di,(%r14)
	add	%di,(%r15)
	nop
	add	%r8w, (%rax)
	add	%r8w, (%rcx)
	add	%r8w, (%rdx)
	add	%r8w, (%rbx)
	add	%r8w, (%rsp)
	add	%r8w, (%rbp)
	add	%r8w, (%rsi)
	add	%r8w, (%rdi)
	add	%r8w, (%r8)
	add	%r8w, (%r9)
	add	%r8w, (%r10)
	add	%r8w, (%r11)
	add	%r8w, (%r12)
	add	%r8w, (%r13)
	add	%r8w, (%r14)
	add	%r8w, (%r15)
	nop
	add	%r9w, (%rax)
	add	%r9w, (%rcx)
	add	%r9w, (%rdx)
	add	%r9w, (%rbx)
	add	%r9w, (%rsp)
	add	%r9w, (%rbp)
	add	%r9w, (%rsi)
	add	%r9w, (%rdi)
	add	%r9w, (%r8)
	add	%r9w, (%r9)
	add	%r9w, (%r10)
	add	%r9w, (%r11)
	add	%r9w, (%r12)
	add	%r9w, (%r13)
	add	%r9w, (%r14)
	add	%r9w, (%r15)
	nop
	add	%r10w,(%rax)
	add	%r10w,(%rcx)
	add	%r10w,(%rdx)
	add	%r10w,(%rbx)
	add	%r10w,(%rsp)
	add	%r10w,(%rbp)
	add	%r10w,(%rsi)
	add	%r10w,(%rdi)
	add	%r10w,(%r8)
	add	%r10w,(%r9)
	add	%r10w,(%r10)
	add	%r10w,(%r11)
	add	%r10w,(%r12)
	add	%r10w,(%r13)
	add	%r10w,(%r14)
	add	%r10w,(%r15)
	nop
	add	%r11w,(%rax)
	add	%r11w,(%rcx)
	add	%r11w,(%rdx)
	add	%r11w,(%rbx)
	add	%r11w,(%rsp)
	add	%r11w,(%rbp)
	add	%r11w,(%rsi)
	add	%r11w,(%rdi)
	add	%r11w,(%r8)
	add	%r11w,(%r9)
	add	%r11w,(%r10)
	add	%r11w,(%r11)
	add	%r11w,(%r12)
	add	%r11w,(%r13)
	add	%r11w,(%r14)
	add	%r11w,(%r15)
	nop
	add	%r12w,(%rax)
	add	%r12w,(%rcx)
	add	%r12w,(%rdx)
	add	%r12w,(%rbx)
	add	%r12w,(%rsp)
	add	%r12w,(%rbp)
	add	%r12w,(%rsi)
	add	%r12w,(%rdi)
	add	%r12w,(%r8)
	add	%r12w,(%r9)
	add	%r12w,(%r10)
	add	%r12w,(%r11)
	add	%r12w,(%r12)
	add	%r12w,(%r13)
	add	%r12w,(%r14)
	add	%r12w,(%r15)
	nop
	add	%r13w,(%rax)
	add	%r13w,(%rcx)
	add	%r13w,(%rdx)
	add	%r13w,(%rbx)
	add	%r13w,(%rsp)
	add	%r13w,(%rbp)
	add	%r13w,(%rsi)
	add	%r13w,(%rdi)
	add	%r13w,(%r8)
	add	%r13w,(%r9)
	add	%r13w,(%r10)
	add	%r13w,(%r11)
	add	%r13w,(%r12)
	add	%r13w,(%r13)
	add	%r13w,(%r14)
	add	%r13w,(%r15)
	nop
	add	%r14w,(%rax)
	add	%r14w,(%rcx)
	add	%r14w,(%rdx)
	add	%r14w,(%rbx)
	add	%r14w,(%rsp)
	add	%r14w,(%rbp)
	add	%r14w,(%rsi)
	add	%r14w,(%rdi)
	add	%r14w,(%r8)
	add	%r14w,(%r9)
	add	%r14w,(%r10)
	add	%r14w,(%r11)
	add	%r14w,(%r12)
	add	%r14w,(%r13)
	add	%r14w,(%r14)
	add	%r14w,(%r15)
	nop
	add	%r15w,(%rax)
	add	%r15w,(%rcx)
	add	%r15w,(%rdx)
	add	%r15w,(%rbx)
	add	%r15w,(%rsp)
	add	%r15w,(%rbp)
	add	%r15w,(%rsi)
	add	%r15w,(%rdi)
	add	%r15w,(%r8)
	add	%r15w,(%r9)
	add	%r15w,(%r10)
	add	%r15w,(%r11)
	add	%r15w,(%r12)
	add	%r15w,(%r13)
	add	%r15w,(%r14)
	add	%r15w,(%r15)
        nop
        nop
        // mem32 += reg32
	add	%eax,(%rax)
	add	%eax,(%rcx)
	add	%eax,(%rdx)
	add	%eax,(%rbx)
	add	%eax,(%rsp)
	add	%eax,(%rbp)
	add	%eax,(%rsi)
	add	%eax,(%rdi)
	add	%eax,(%r8)
	add	%eax,(%r9)
	add	%eax,(%r10)
	add	%eax,(%r11)
	add	%eax,(%r12)
	add	%eax,(%r13)
	add	%eax,(%r14)
	add	%eax,(%r15)
	nop
	add	%ecx,(%rax)
	add	%ecx,(%rcx)
	add	%ecx,(%rdx)
	add	%ecx,(%rbx)
	add	%ecx,(%rsp)
	add	%ecx,(%rbp)
	add	%ecx,(%rsi)
	add	%ecx,(%rdi)
	add	%ecx,(%r8)
	add	%ecx,(%r9)
	add	%ecx,(%r10)
	add	%ecx,(%r11)
	add	%ecx,(%r12)
	add	%ecx,(%r13)
	add	%ecx,(%r14)
	add	%ecx,(%r15)
	nop
	add	%edx,(%rax)
	add	%edx,(%rcx)
	add	%edx,(%rdx)
	add	%edx,(%rbx)
	add	%edx,(%rsp)
	add	%edx,(%rbp)
	add	%edx,(%rsi)
	add	%edx,(%rdi)
	add	%edx,(%r8)
	add	%edx,(%r9)
	add	%edx,(%r10)
	add	%edx,(%r11)
	add	%edx,(%r12)
	add	%edx,(%r13)
	add	%edx,(%r14)
	add	%edx,(%r15)
	nop
	add	%ebx,(%rax)
	add	%ebx,(%rcx)
	add	%ebx,(%rdx)
	add	%ebx,(%rbx)
	add	%ebx,(%rsp)
	add	%ebx,(%rbp)
	add	%ebx,(%rsi)
	add	%ebx,(%rdi)
	add	%ebx,(%r8)
	add	%ebx,(%r9)
	add	%ebx,(%r10)
	add	%ebx,(%r11)
	add	%ebx,(%r12)
	add	%ebx,(%r13)
	add	%ebx,(%r14)
	add	%ebx,(%r15)
	nop
	add	%esp,(%rax)
	add	%esp,(%rcx)
	add	%esp,(%rdx)
	add	%esp,(%rbx)
	add	%esp,(%rsp)
	add	%esp,(%rbp)
	add	%esp,(%rsi)
	add	%esp,(%rdi)
	add	%esp,(%r8)
	add	%esp,(%r9)
	add	%esp,(%r10)
	add	%esp,(%r11)
	add	%esp,(%r12)
	add	%esp,(%r13)
	add	%esp,(%r14)
	add	%esp,(%r15)
	nop
	add	%ebp,(%rax)
	add	%ebp,(%rcx)
	add	%ebp,(%rdx)
	add	%ebp,(%rbx)
	add	%ebp,(%rsp)
	add	%ebp,(%rbp)
	add	%ebp,(%rsi)
	add	%ebp,(%rdi)
	add	%ebp,(%r8)
	add	%ebp,(%r9)
	add	%ebp,(%r10)
	add	%ebp,(%r11)
	add	%ebp,(%r12)
	add	%ebp,(%r13)
	add	%ebp,(%r14)
	add	%ebp,(%r15)
	nop
	add	%esi,(%rax)
	add	%esi,(%rcx)
	add	%esi,(%rdx)
	add	%esi,(%rbx)
	add	%esi,(%rsp)
	add	%esi,(%rbp)
	add	%esi,(%rsi)
	add	%esi,(%rdi)
	add	%esi,(%r8)
	add	%esi,(%r9)
	add	%esi,(%r10)
	add	%esi,(%r11)
	add	%esi,(%r12)
	add	%esi,(%r13)
	add	%esi,(%r14)
	add	%esi,(%r15)
	nop
	add	%edi,(%rax)
	add	%edi,(%rcx)
	add	%edi,(%rdx)
	add	%edi,(%rbx)
	add	%edi,(%rsp)
	add	%edi,(%rbp)
	add	%edi,(%rsi)
	add	%edi,(%rdi)
	add	%edi,(%r8)
	add	%edi,(%r9)
	add	%edi,(%r10)
	add	%edi,(%r11)
	add	%edi,(%r12)
	add	%edi,(%r13)
	add	%edi,(%r14)
	add	%edi,(%r15)
	nop
	add	%r8d, (%rax)
	add	%r8d, (%rcx)
	add	%r8d, (%rdx)
	add	%r8d, (%rbx)
	add	%r8d, (%rsp)
	add	%r8d, (%rbp)
	add	%r8d, (%rsi)
	add	%r8d, (%rdi)
	add	%r8d, (%r8)
	add	%r8d, (%r9)
	add	%r8d, (%r10)
	add	%r8d, (%r11)
	add	%r8d, (%r12)
	add	%r8d, (%r13)
	add	%r8d, (%r14)
	add	%r8d, (%r15)
	nop
	add	%r9d, (%rax)
	add	%r9d, (%rcx)
	add	%r9d, (%rdx)
	add	%r9d, (%rbx)
	add	%r9d, (%rsp)
	add	%r9d, (%rbp)
	add	%r9d, (%rsi)
	add	%r9d, (%rdi)
	add	%r9d, (%r8)
	add	%r9d, (%r9)
	add	%r9d, (%r10)
	add	%r9d, (%r11)
	add	%r9d, (%r12)
	add	%r9d, (%r13)
	add	%r9d, (%r14)
	add	%r9d, (%r15)
	nop
	add	%r10d,(%rax)
	add	%r10d,(%rcx)
	add	%r10d,(%rdx)
	add	%r10d,(%rbx)
	add	%r10d,(%rsp)
	add	%r10d,(%rbp)
	add	%r10d,(%rsi)
	add	%r10d,(%rdi)
	add	%r10d,(%r8)
	add	%r10d,(%r9)
	add	%r10d,(%r10)
	add	%r10d,(%r11)
	add	%r10d,(%r12)
	add	%r10d,(%r13)
	add	%r10d,(%r14)
	add	%r10d,(%r15)
	nop
	add	%r11d,(%rax)
	add	%r11d,(%rcx)
	add	%r11d,(%rdx)
	add	%r11d,(%rbx)
	add	%r11d,(%rsp)
	add	%r11d,(%rbp)
	add	%r11d,(%rsi)
	add	%r11d,(%rdi)
	add	%r11d,(%r8)
	add	%r11d,(%r9)
	add	%r11d,(%r10)
	add	%r11d,(%r11)
	add	%r11d,(%r12)
	add	%r11d,(%r13)
	add	%r11d,(%r14)
	add	%r11d,(%r15)
	nop
	add	%r12d,(%rax)
	add	%r12d,(%rcx)
	add	%r12d,(%rdx)
	add	%r12d,(%rbx)
	add	%r12d,(%rsp)
	add	%r12d,(%rbp)
	add	%r12d,(%rsi)
	add	%r12d,(%rdi)
	add	%r12d,(%r8)
	add	%r12d,(%r9)
	add	%r12d,(%r10)
	add	%r12d,(%r11)
	add	%r12d,(%r12)
	add	%r12d,(%r13)
	add	%r12d,(%r14)
	add	%r12d,(%r15)
	nop
	add	%r13d,(%rax)
	add	%r13d,(%rcx)
	add	%r13d,(%rdx)
	add	%r13d,(%rbx)
	add	%r13d,(%rsp)
	add	%r13d,(%rbp)
	add	%r13d,(%rsi)
	add	%r13d,(%rdi)
	add	%r13d,(%r8)
	add	%r13d,(%r9)
	add	%r13d,(%r10)
	add	%r13d,(%r11)
	add	%r13d,(%r12)
	add	%r13d,(%r13)
	add	%r13d,(%r14)
	add	%r13d,(%r15)
	nop
	add	%r14d,(%rax)
	add	%r14d,(%rcx)
	add	%r14d,(%rdx)
	add	%r14d,(%rbx)
	add	%r14d,(%rsp)
	add	%r14d,(%rbp)
	add	%r14d,(%rsi)
	add	%r14d,(%rdi)
	add	%r14d,(%r8)
	add	%r14d,(%r9)
	add	%r14d,(%r10)
	add	%r14d,(%r11)
	add	%r14d,(%r12)
	add	%r14d,(%r13)
	add	%r14d,(%r14)
	add	%r14d,(%r15)
	nop
	add	%r15d,(%rax)
	add	%r15d,(%rcx)
	add	%r15d,(%rdx)
	add	%r15d,(%rbx)
	add	%r15d,(%rsp)
	add	%r15d,(%rbp)
	add	%r15d,(%rsi)
	add	%r15d,(%rdi)
	add	%r15d,(%r8)
	add	%r15d,(%r9)
	add	%r15d,(%r10)
	add	%r15d,(%r11)
	add	%r15d,(%r12)
	add	%r15d,(%r13)
	add	%r15d,(%r14)
	add	%r15d,(%r15)
        nop
        nop
        // mem64 += reg64
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


