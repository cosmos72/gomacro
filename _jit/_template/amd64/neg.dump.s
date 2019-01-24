
neg.o:     file format elf64-x86-64


Disassembly of section .text:

0000000000000000 <Neg>:
   0:	f6 d8                	neg    %al
   2:	f6 d9                	neg    %cl
   4:	f6 da                	neg    %dl
   6:	f6 db                	neg    %bl
   8:	40 f6 dc             	neg    %spl
   b:	40 f6 dd             	neg    %bpl
   e:	40 f6 de             	neg    %sil
  11:	40 f6 df             	neg    %dil
  14:	41 f6 d8             	neg    %r8b
  17:	41 f6 d9             	neg    %r9b
  1a:	41 f6 da             	neg    %r10b
  1d:	41 f6 db             	neg    %r11b
  20:	41 f6 dc             	neg    %r12b
  23:	41 f6 dd             	neg    %r13b
  26:	41 f6 de             	neg    %r14b
  29:	41 f6 df             	neg    %r15b
  2c:	90                   	nop
  2d:	66 f7 d8             	neg    %ax
  30:	66 f7 d9             	neg    %cx
  33:	66 f7 da             	neg    %dx
  36:	66 f7 db             	neg    %bx
  39:	66 f7 dc             	neg    %sp
  3c:	66 f7 dd             	neg    %bp
  3f:	66 f7 de             	neg    %si
  42:	66 f7 df             	neg    %di
  45:	66 41 f7 d8          	neg    %r8w
  49:	66 41 f7 d9          	neg    %r9w
  4d:	66 41 f7 da          	neg    %r10w
  51:	66 41 f7 db          	neg    %r11w
  55:	66 41 f7 dc          	neg    %r12w
  59:	66 41 f7 dd          	neg    %r13w
  5d:	66 41 f7 de          	neg    %r14w
  61:	66 41 f7 df          	neg    %r15w
  65:	90                   	nop
  66:	f7 d8                	neg    %eax
  68:	f7 d9                	neg    %ecx
  6a:	f7 da                	neg    %edx
  6c:	f7 db                	neg    %ebx
  6e:	f7 dc                	neg    %esp
  70:	f7 dd                	neg    %ebp
  72:	f7 de                	neg    %esi
  74:	f7 df                	neg    %edi
  76:	41 f7 d8             	neg    %r8d
  79:	41 f7 d9             	neg    %r9d
  7c:	41 f7 da             	neg    %r10d
  7f:	41 f7 db             	neg    %r11d
  82:	41 f7 dc             	neg    %r12d
  85:	41 f7 dd             	neg    %r13d
  88:	41 f7 de             	neg    %r14d
  8b:	41 f7 df             	neg    %r15d
  8e:	90                   	nop
  8f:	48 f7 d8             	neg    %rax
  92:	48 f7 d9             	neg    %rcx
  95:	48 f7 da             	neg    %rdx
  98:	48 f7 db             	neg    %rbx
  9b:	48 f7 dc             	neg    %rsp
  9e:	48 f7 dd             	neg    %rbp
  a1:	48 f7 de             	neg    %rsi
  a4:	48 f7 df             	neg    %rdi
  a7:	49 f7 d8             	neg    %r8
  aa:	49 f7 d9             	neg    %r9
  ad:	49 f7 da             	neg    %r10
  b0:	49 f7 db             	neg    %r11
  b3:	49 f7 dc             	neg    %r12
  b6:	49 f7 dd             	neg    %r13
  b9:	49 f7 de             	neg    %r14
  bc:	49 f7 df             	neg    %r15
  bf:	90                   	nop
  c0:	f6 18                	negb   (%rax)
  c2:	f6 19                	negb   (%rcx)
  c4:	f6 1a                	negb   (%rdx)
  c6:	f6 1b                	negb   (%rbx)
  c8:	f6 1c 24             	negb   (%rsp)
  cb:	f6 5d 00             	negb   0x0(%rbp)
  ce:	f6 1e                	negb   (%rsi)
  d0:	f6 1f                	negb   (%rdi)
  d2:	41 f6 18             	negb   (%r8)
  d5:	41 f6 19             	negb   (%r9)
  d8:	41 f6 1a             	negb   (%r10)
  db:	41 f6 1b             	negb   (%r11)
  de:	41 f6 1c 24          	negb   (%r12)
  e2:	41 f6 5d 00          	negb   0x0(%r13)
  e6:	41 f6 1e             	negb   (%r14)
  e9:	41 f6 1f             	negb   (%r15)
  ec:	90                   	nop
  ed:	66 f7 18             	negw   (%rax)
  f0:	66 f7 19             	negw   (%rcx)
  f3:	66 f7 1a             	negw   (%rdx)
  f6:	66 f7 1b             	negw   (%rbx)
  f9:	66 f7 1c 24          	negw   (%rsp)
  fd:	66 f7 5d 00          	negw   0x0(%rbp)
 101:	66 f7 1e             	negw   (%rsi)
 104:	66 f7 1f             	negw   (%rdi)
 107:	66 41 f7 18          	negw   (%r8)
 10b:	66 41 f7 19          	negw   (%r9)
 10f:	66 41 f7 1a          	negw   (%r10)
 113:	66 41 f7 1b          	negw   (%r11)
 117:	66 41 f7 1c 24       	negw   (%r12)
 11c:	66 41 f7 5d 00       	negw   0x0(%r13)
 121:	66 41 f7 1e          	negw   (%r14)
 125:	66 41 f7 1f          	negw   (%r15)
 129:	90                   	nop
 12a:	f7 18                	negl   (%rax)
 12c:	f7 19                	negl   (%rcx)
 12e:	f7 1a                	negl   (%rdx)
 130:	f7 1b                	negl   (%rbx)
 132:	f7 1c 24             	negl   (%rsp)
 135:	f7 5d 00             	negl   0x0(%rbp)
 138:	f7 1e                	negl   (%rsi)
 13a:	f7 1f                	negl   (%rdi)
 13c:	41 f7 18             	negl   (%r8)
 13f:	41 f7 19             	negl   (%r9)
 142:	41 f7 1a             	negl   (%r10)
 145:	41 f7 1b             	negl   (%r11)
 148:	41 f7 1c 24          	negl   (%r12)
 14c:	41 f7 5d 00          	negl   0x0(%r13)
 150:	41 f7 1e             	negl   (%r14)
 153:	41 f7 1f             	negl   (%r15)
 156:	90                   	nop
 157:	48 f7 18             	negq   (%rax)
 15a:	48 f7 19             	negq   (%rcx)
 15d:	48 f7 1a             	negq   (%rdx)
 160:	48 f7 1b             	negq   (%rbx)
 163:	48 f7 1c 24          	negq   (%rsp)
 167:	48 f7 5d 00          	negq   0x0(%rbp)
 16b:	48 f7 1e             	negq   (%rsi)
 16e:	48 f7 1f             	negq   (%rdi)
 171:	49 f7 18             	negq   (%r8)
 174:	49 f7 19             	negq   (%r9)
 177:	49 f7 1a             	negq   (%r10)
 17a:	49 f7 1b             	negq   (%r11)
 17d:	49 f7 1c 24          	negq   (%r12)
 181:	49 f7 5d 00          	negq   0x0(%r13)
 185:	49 f7 1e             	negq   (%r14)
 188:	49 f7 1f             	negq   (%r15)
 18b:	90                   	nop
 18c:	f6 58 7f             	negb   0x7f(%rax)
 18f:	f6 59 7f             	negb   0x7f(%rcx)
 192:	f6 5a 7f             	negb   0x7f(%rdx)
 195:	f6 5b 7f             	negb   0x7f(%rbx)
 198:	f6 5c 24 7f          	negb   0x7f(%rsp)
 19c:	f6 5d 7f             	negb   0x7f(%rbp)
 19f:	f6 5e 7f             	negb   0x7f(%rsi)
 1a2:	f6 5f 7f             	negb   0x7f(%rdi)
 1a5:	41 f6 58 7f          	negb   0x7f(%r8)
 1a9:	41 f6 59 7f          	negb   0x7f(%r9)
 1ad:	41 f6 5a 7f          	negb   0x7f(%r10)
 1b1:	41 f6 5b 7f          	negb   0x7f(%r11)
 1b5:	41 f6 5c 24 7f       	negb   0x7f(%r12)
 1ba:	41 f6 5d 7f          	negb   0x7f(%r13)
 1be:	41 f6 5e 7f          	negb   0x7f(%r14)
 1c2:	41 f6 5f 7f          	negb   0x7f(%r15)
 1c6:	90                   	nop
 1c7:	66 f7 58 7f          	negw   0x7f(%rax)
 1cb:	66 f7 59 7f          	negw   0x7f(%rcx)
 1cf:	66 f7 5a 7f          	negw   0x7f(%rdx)
 1d3:	66 f7 5b 7f          	negw   0x7f(%rbx)
 1d7:	66 f7 5c 24 7f       	negw   0x7f(%rsp)
 1dc:	66 f7 5d 7f          	negw   0x7f(%rbp)
 1e0:	66 f7 5e 7f          	negw   0x7f(%rsi)
 1e4:	66 f7 5f 7f          	negw   0x7f(%rdi)
 1e8:	66 41 f7 58 7f       	negw   0x7f(%r8)
 1ed:	66 41 f7 59 7f       	negw   0x7f(%r9)
 1f2:	66 41 f7 5a 7f       	negw   0x7f(%r10)
 1f7:	66 41 f7 5b 7f       	negw   0x7f(%r11)
 1fc:	66 41 f7 5c 24 7f    	negw   0x7f(%r12)
 202:	66 41 f7 5d 7f       	negw   0x7f(%r13)
 207:	66 41 f7 5e 7f       	negw   0x7f(%r14)
 20c:	66 41 f7 5f 7f       	negw   0x7f(%r15)
 211:	90                   	nop
 212:	f7 58 7f             	negl   0x7f(%rax)
 215:	f7 59 7f             	negl   0x7f(%rcx)
 218:	f7 5a 7f             	negl   0x7f(%rdx)
 21b:	f7 5b 7f             	negl   0x7f(%rbx)
 21e:	f7 5c 24 7f          	negl   0x7f(%rsp)
 222:	f7 5d 7f             	negl   0x7f(%rbp)
 225:	f7 5e 7f             	negl   0x7f(%rsi)
 228:	f7 5f 7f             	negl   0x7f(%rdi)
 22b:	41 f7 58 7f          	negl   0x7f(%r8)
 22f:	41 f7 59 7f          	negl   0x7f(%r9)
 233:	41 f7 5a 7f          	negl   0x7f(%r10)
 237:	41 f7 5b 7f          	negl   0x7f(%r11)
 23b:	41 f7 5c 24 7f       	negl   0x7f(%r12)
 240:	41 f7 5d 7f          	negl   0x7f(%r13)
 244:	41 f7 5e 7f          	negl   0x7f(%r14)
 248:	41 f7 5f 7f          	negl   0x7f(%r15)
 24c:	90                   	nop
 24d:	48 f7 58 7f          	negq   0x7f(%rax)
 251:	48 f7 59 7f          	negq   0x7f(%rcx)
 255:	48 f7 5a 7f          	negq   0x7f(%rdx)
 259:	48 f7 5b 7f          	negq   0x7f(%rbx)
 25d:	48 f7 5c 24 7f       	negq   0x7f(%rsp)
 262:	48 f7 5d 7f          	negq   0x7f(%rbp)
 266:	48 f7 5e 7f          	negq   0x7f(%rsi)
 26a:	48 f7 5f 7f          	negq   0x7f(%rdi)
 26e:	49 f7 58 7f          	negq   0x7f(%r8)
 272:	49 f7 59 7f          	negq   0x7f(%r9)
 276:	49 f7 5a 7f          	negq   0x7f(%r10)
 27a:	49 f7 5b 7f          	negq   0x7f(%r11)
 27e:	49 f7 5c 24 7f       	negq   0x7f(%r12)
 283:	49 f7 5d 7f          	negq   0x7f(%r13)
 287:	49 f7 5e 7f          	negq   0x7f(%r14)
 28b:	49 f7 5f 7f          	negq   0x7f(%r15)
 28f:	90                   	nop
 290:	f6 98 12 34 56 7f    	negb   0x7f563412(%rax)
 296:	f6 99 12 34 56 7f    	negb   0x7f563412(%rcx)
 29c:	f6 9a 12 34 56 7f    	negb   0x7f563412(%rdx)
 2a2:	f6 9b 12 34 56 7f    	negb   0x7f563412(%rbx)
 2a8:	f6 9c 24 12 34 56 7f 	negb   0x7f563412(%rsp)
 2af:	f6 9d 12 34 56 7f    	negb   0x7f563412(%rbp)
 2b5:	f6 9e 12 34 56 7f    	negb   0x7f563412(%rsi)
 2bb:	f6 9f 12 34 56 7f    	negb   0x7f563412(%rdi)
 2c1:	41 f6 98 12 34 56 7f 	negb   0x7f563412(%r8)
 2c8:	41 f6 99 12 34 56 7f 	negb   0x7f563412(%r9)
 2cf:	41 f6 9a 12 34 56 7f 	negb   0x7f563412(%r10)
 2d6:	41 f6 9b 12 34 56 7f 	negb   0x7f563412(%r11)
 2dd:	41 f6 9c 24 12 34 56 	negb   0x7f563412(%r12)
 2e4:	7f 
 2e5:	41 f6 9d 12 34 56 7f 	negb   0x7f563412(%r13)
 2ec:	41 f6 9e 12 34 56 7f 	negb   0x7f563412(%r14)
 2f3:	41 f6 9f 12 34 56 7f 	negb   0x7f563412(%r15)
 2fa:	90                   	nop
 2fb:	66 f7 98 12 34 56 7f 	negw   0x7f563412(%rax)
 302:	66 f7 99 12 34 56 7f 	negw   0x7f563412(%rcx)
 309:	66 f7 9a 12 34 56 7f 	negw   0x7f563412(%rdx)
 310:	66 f7 9b 12 34 56 7f 	negw   0x7f563412(%rbx)
 317:	66 f7 9c 24 12 34 56 	negw   0x7f563412(%rsp)
 31e:	7f 
 31f:	66 f7 9d 12 34 56 7f 	negw   0x7f563412(%rbp)
 326:	66 f7 9e 12 34 56 7f 	negw   0x7f563412(%rsi)
 32d:	66 f7 9f 12 34 56 7f 	negw   0x7f563412(%rdi)
 334:	66 41 f7 98 12 34 56 	negw   0x7f563412(%r8)
 33b:	7f 
 33c:	66 41 f7 99 12 34 56 	negw   0x7f563412(%r9)
 343:	7f 
 344:	66 41 f7 9a 12 34 56 	negw   0x7f563412(%r10)
 34b:	7f 
 34c:	66 41 f7 9b 12 34 56 	negw   0x7f563412(%r11)
 353:	7f 
 354:	66 41 f7 9c 24 12 34 	negw   0x7f563412(%r12)
 35b:	56 7f 
 35d:	66 41 f7 9d 12 34 56 	negw   0x7f563412(%r13)
 364:	7f 
 365:	66 41 f7 9e 12 34 56 	negw   0x7f563412(%r14)
 36c:	7f 
 36d:	66 41 f7 9f 12 34 56 	negw   0x7f563412(%r15)
 374:	7f 
 375:	90                   	nop
 376:	f7 98 12 34 56 7f    	negl   0x7f563412(%rax)
 37c:	f7 99 12 34 56 7f    	negl   0x7f563412(%rcx)
 382:	f7 9a 12 34 56 7f    	negl   0x7f563412(%rdx)
 388:	f7 9b 12 34 56 7f    	negl   0x7f563412(%rbx)
 38e:	f7 9c 24 12 34 56 7f 	negl   0x7f563412(%rsp)
 395:	f7 9d 12 34 56 7f    	negl   0x7f563412(%rbp)
 39b:	f7 9e 12 34 56 7f    	negl   0x7f563412(%rsi)
 3a1:	f7 9f 12 34 56 7f    	negl   0x7f563412(%rdi)
 3a7:	41 f7 98 12 34 56 7f 	negl   0x7f563412(%r8)
 3ae:	41 f7 99 12 34 56 7f 	negl   0x7f563412(%r9)
 3b5:	41 f7 9a 12 34 56 7f 	negl   0x7f563412(%r10)
 3bc:	41 f7 9b 12 34 56 7f 	negl   0x7f563412(%r11)
 3c3:	41 f7 9c 24 12 34 56 	negl   0x7f563412(%r12)
 3ca:	7f 
 3cb:	41 f7 9d 12 34 56 7f 	negl   0x7f563412(%r13)
 3d2:	41 f7 9e 12 34 56 7f 	negl   0x7f563412(%r14)
 3d9:	41 f7 9f 12 34 56 7f 	negl   0x7f563412(%r15)
 3e0:	90                   	nop
 3e1:	48 f7 98 12 34 56 7f 	negq   0x7f563412(%rax)
 3e8:	48 f7 99 12 34 56 7f 	negq   0x7f563412(%rcx)
 3ef:	48 f7 9a 12 34 56 7f 	negq   0x7f563412(%rdx)
 3f6:	48 f7 9b 12 34 56 7f 	negq   0x7f563412(%rbx)
 3fd:	48 f7 9c 24 12 34 56 	negq   0x7f563412(%rsp)
 404:	7f 
 405:	48 f7 9d 12 34 56 7f 	negq   0x7f563412(%rbp)
 40c:	48 f7 9e 12 34 56 7f 	negq   0x7f563412(%rsi)
 413:	48 f7 9f 12 34 56 7f 	negq   0x7f563412(%rdi)
 41a:	49 f7 98 12 34 56 7f 	negq   0x7f563412(%r8)
 421:	49 f7 99 12 34 56 7f 	negq   0x7f563412(%r9)
 428:	49 f7 9a 12 34 56 7f 	negq   0x7f563412(%r10)
 42f:	49 f7 9b 12 34 56 7f 	negq   0x7f563412(%r11)
 436:	49 f7 9c 24 12 34 56 	negq   0x7f563412(%r12)
 43d:	7f 
 43e:	49 f7 9d 12 34 56 7f 	negq   0x7f563412(%r13)
 445:	49 f7 9e 12 34 56 7f 	negq   0x7f563412(%r14)
 44c:	49 f7 9f 12 34 56 7f 	negq   0x7f563412(%r15)
 453:	c3                   	retq   
