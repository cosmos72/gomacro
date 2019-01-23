
neg.o:     file format elf64-x86-64


Disassembly of section .text:

0000000000000000 <Neg>:
   0:	48 f7 d8             	neg    %rax
   3:	48 f7 d9             	neg    %rcx
   6:	48 f7 da             	neg    %rdx
   9:	48 f7 db             	neg    %rbx
   c:	48 f7 dc             	neg    %rsp
   f:	48 f7 dd             	neg    %rbp
  12:	48 f7 de             	neg    %rsi
  15:	48 f7 df             	neg    %rdi
  18:	49 f7 d8             	neg    %r8
  1b:	49 f7 d9             	neg    %r9
  1e:	49 f7 da             	neg    %r10
  21:	49 f7 db             	neg    %r11
  24:	49 f7 dc             	neg    %r12
  27:	49 f7 dd             	neg    %r13
  2a:	49 f7 de             	neg    %r14
  2d:	49 f7 df             	neg    %r15
  30:	90                   	nop
  31:	f6 18                	negb   (%rax)
  33:	f6 19                	negb   (%rcx)
  35:	f6 1a                	negb   (%rdx)
  37:	f6 1b                	negb   (%rbx)
  39:	f6 1c 24             	negb   (%rsp)
  3c:	f6 5d 00             	negb   0x0(%rbp)
  3f:	f6 1e                	negb   (%rsi)
  41:	f6 1f                	negb   (%rdi)
  43:	41 f6 18             	negb   (%r8)
  46:	41 f6 19             	negb   (%r9)
  49:	41 f6 1a             	negb   (%r10)
  4c:	41 f6 1b             	negb   (%r11)
  4f:	41 f6 1c 24          	negb   (%r12)
  53:	41 f6 5d 00          	negb   0x0(%r13)
  57:	41 f6 1e             	negb   (%r14)
  5a:	41 f6 1f             	negb   (%r15)
  5d:	90                   	nop
  5e:	66 f7 18             	negw   (%rax)
  61:	66 f7 19             	negw   (%rcx)
  64:	66 f7 1a             	negw   (%rdx)
  67:	66 f7 1b             	negw   (%rbx)
  6a:	66 f7 1c 24          	negw   (%rsp)
  6e:	66 f7 5d 00          	negw   0x0(%rbp)
  72:	66 f7 1e             	negw   (%rsi)
  75:	66 f7 1f             	negw   (%rdi)
  78:	66 41 f7 18          	negw   (%r8)
  7c:	66 41 f7 19          	negw   (%r9)
  80:	66 41 f7 1a          	negw   (%r10)
  84:	66 41 f7 1b          	negw   (%r11)
  88:	66 41 f7 1c 24       	negw   (%r12)
  8d:	66 41 f7 5d 00       	negw   0x0(%r13)
  92:	66 41 f7 1e          	negw   (%r14)
  96:	66 41 f7 1f          	negw   (%r15)
  9a:	90                   	nop
  9b:	f7 18                	negl   (%rax)
  9d:	f7 19                	negl   (%rcx)
  9f:	f7 1a                	negl   (%rdx)
  a1:	f7 1b                	negl   (%rbx)
  a3:	f7 1c 24             	negl   (%rsp)
  a6:	f7 5d 00             	negl   0x0(%rbp)
  a9:	f7 1e                	negl   (%rsi)
  ab:	f7 1f                	negl   (%rdi)
  ad:	41 f7 18             	negl   (%r8)
  b0:	41 f7 19             	negl   (%r9)
  b3:	41 f7 1a             	negl   (%r10)
  b6:	41 f7 1b             	negl   (%r11)
  b9:	41 f7 1c 24          	negl   (%r12)
  bd:	41 f7 5d 00          	negl   0x0(%r13)
  c1:	41 f7 1e             	negl   (%r14)
  c4:	41 f7 1f             	negl   (%r15)
  c7:	90                   	nop
  c8:	48 f7 18             	negq   (%rax)
  cb:	48 f7 19             	negq   (%rcx)
  ce:	48 f7 1a             	negq   (%rdx)
  d1:	48 f7 1b             	negq   (%rbx)
  d4:	48 f7 1c 24          	negq   (%rsp)
  d8:	48 f7 5d 00          	negq   0x0(%rbp)
  dc:	48 f7 1e             	negq   (%rsi)
  df:	48 f7 1f             	negq   (%rdi)
  e2:	49 f7 18             	negq   (%r8)
  e5:	49 f7 19             	negq   (%r9)
  e8:	49 f7 1a             	negq   (%r10)
  eb:	49 f7 1b             	negq   (%r11)
  ee:	49 f7 1c 24          	negq   (%r12)
  f2:	49 f7 5d 00          	negq   0x0(%r13)
  f6:	49 f7 1e             	negq   (%r14)
  f9:	49 f7 1f             	negq   (%r15)
  fc:	90                   	nop
  fd:	f6 58 7f             	negb   0x7f(%rax)
 100:	f6 59 7f             	negb   0x7f(%rcx)
 103:	f6 5a 7f             	negb   0x7f(%rdx)
 106:	f6 5b 7f             	negb   0x7f(%rbx)
 109:	f6 5c 24 7f          	negb   0x7f(%rsp)
 10d:	f6 5d 7f             	negb   0x7f(%rbp)
 110:	f6 5e 7f             	negb   0x7f(%rsi)
 113:	f6 5f 7f             	negb   0x7f(%rdi)
 116:	41 f6 58 7f          	negb   0x7f(%r8)
 11a:	41 f6 59 7f          	negb   0x7f(%r9)
 11e:	41 f6 5a 7f          	negb   0x7f(%r10)
 122:	41 f6 5b 7f          	negb   0x7f(%r11)
 126:	41 f6 5c 24 7f       	negb   0x7f(%r12)
 12b:	41 f6 5d 7f          	negb   0x7f(%r13)
 12f:	41 f6 5e 7f          	negb   0x7f(%r14)
 133:	41 f6 5f 7f          	negb   0x7f(%r15)
 137:	90                   	nop
 138:	66 f7 58 7f          	negw   0x7f(%rax)
 13c:	66 f7 59 7f          	negw   0x7f(%rcx)
 140:	66 f7 5a 7f          	negw   0x7f(%rdx)
 144:	66 f7 5b 7f          	negw   0x7f(%rbx)
 148:	66 f7 5c 24 7f       	negw   0x7f(%rsp)
 14d:	66 f7 5d 7f          	negw   0x7f(%rbp)
 151:	66 f7 5e 7f          	negw   0x7f(%rsi)
 155:	66 f7 5f 7f          	negw   0x7f(%rdi)
 159:	66 41 f7 58 7f       	negw   0x7f(%r8)
 15e:	66 41 f7 59 7f       	negw   0x7f(%r9)
 163:	66 41 f7 5a 7f       	negw   0x7f(%r10)
 168:	66 41 f7 5b 7f       	negw   0x7f(%r11)
 16d:	66 41 f7 5c 24 7f    	negw   0x7f(%r12)
 173:	66 41 f7 5d 7f       	negw   0x7f(%r13)
 178:	66 41 f7 5e 7f       	negw   0x7f(%r14)
 17d:	66 41 f7 5f 7f       	negw   0x7f(%r15)
 182:	90                   	nop
 183:	f7 58 7f             	negl   0x7f(%rax)
 186:	f7 59 7f             	negl   0x7f(%rcx)
 189:	f7 5a 7f             	negl   0x7f(%rdx)
 18c:	f7 5b 7f             	negl   0x7f(%rbx)
 18f:	f7 5c 24 7f          	negl   0x7f(%rsp)
 193:	f7 5d 7f             	negl   0x7f(%rbp)
 196:	f7 5e 7f             	negl   0x7f(%rsi)
 199:	f7 5f 7f             	negl   0x7f(%rdi)
 19c:	41 f7 58 7f          	negl   0x7f(%r8)
 1a0:	41 f7 59 7f          	negl   0x7f(%r9)
 1a4:	41 f7 5a 7f          	negl   0x7f(%r10)
 1a8:	41 f7 5b 7f          	negl   0x7f(%r11)
 1ac:	41 f7 5c 24 7f       	negl   0x7f(%r12)
 1b1:	41 f7 5d 7f          	negl   0x7f(%r13)
 1b5:	41 f7 5e 7f          	negl   0x7f(%r14)
 1b9:	41 f7 5f 7f          	negl   0x7f(%r15)
 1bd:	90                   	nop
 1be:	48 f7 58 7f          	negq   0x7f(%rax)
 1c2:	48 f7 59 7f          	negq   0x7f(%rcx)
 1c6:	48 f7 5a 7f          	negq   0x7f(%rdx)
 1ca:	48 f7 5b 7f          	negq   0x7f(%rbx)
 1ce:	48 f7 5c 24 7f       	negq   0x7f(%rsp)
 1d3:	48 f7 5d 7f          	negq   0x7f(%rbp)
 1d7:	48 f7 5e 7f          	negq   0x7f(%rsi)
 1db:	48 f7 5f 7f          	negq   0x7f(%rdi)
 1df:	49 f7 58 7f          	negq   0x7f(%r8)
 1e3:	49 f7 59 7f          	negq   0x7f(%r9)
 1e7:	49 f7 5a 7f          	negq   0x7f(%r10)
 1eb:	49 f7 5b 7f          	negq   0x7f(%r11)
 1ef:	49 f7 5c 24 7f       	negq   0x7f(%r12)
 1f4:	49 f7 5d 7f          	negq   0x7f(%r13)
 1f8:	49 f7 5e 7f          	negq   0x7f(%r14)
 1fc:	49 f7 5f 7f          	negq   0x7f(%r15)
 200:	90                   	nop
 201:	f6 98 12 34 56 7f    	negb   0x7f563412(%rax)
 207:	f6 99 12 34 56 7f    	negb   0x7f563412(%rcx)
 20d:	f6 9a 12 34 56 7f    	negb   0x7f563412(%rdx)
 213:	f6 9b 12 34 56 7f    	negb   0x7f563412(%rbx)
 219:	f6 9c 24 12 34 56 7f 	negb   0x7f563412(%rsp)
 220:	f6 9d 12 34 56 7f    	negb   0x7f563412(%rbp)
 226:	f6 9e 12 34 56 7f    	negb   0x7f563412(%rsi)
 22c:	f6 9f 12 34 56 7f    	negb   0x7f563412(%rdi)
 232:	41 f6 98 12 34 56 7f 	negb   0x7f563412(%r8)
 239:	41 f6 99 12 34 56 7f 	negb   0x7f563412(%r9)
 240:	41 f6 9a 12 34 56 7f 	negb   0x7f563412(%r10)
 247:	41 f6 9b 12 34 56 7f 	negb   0x7f563412(%r11)
 24e:	41 f6 9c 24 12 34 56 	negb   0x7f563412(%r12)
 255:	7f 
 256:	41 f6 9d 12 34 56 7f 	negb   0x7f563412(%r13)
 25d:	41 f6 9e 12 34 56 7f 	negb   0x7f563412(%r14)
 264:	41 f6 9f 12 34 56 7f 	negb   0x7f563412(%r15)
 26b:	90                   	nop
 26c:	66 f7 98 12 34 56 7f 	negw   0x7f563412(%rax)
 273:	66 f7 99 12 34 56 7f 	negw   0x7f563412(%rcx)
 27a:	66 f7 9a 12 34 56 7f 	negw   0x7f563412(%rdx)
 281:	66 f7 9b 12 34 56 7f 	negw   0x7f563412(%rbx)
 288:	66 f7 9c 24 12 34 56 	negw   0x7f563412(%rsp)
 28f:	7f 
 290:	66 f7 9d 12 34 56 7f 	negw   0x7f563412(%rbp)
 297:	66 f7 9e 12 34 56 7f 	negw   0x7f563412(%rsi)
 29e:	66 f7 9f 12 34 56 7f 	negw   0x7f563412(%rdi)
 2a5:	66 41 f7 98 12 34 56 	negw   0x7f563412(%r8)
 2ac:	7f 
 2ad:	66 41 f7 99 12 34 56 	negw   0x7f563412(%r9)
 2b4:	7f 
 2b5:	66 41 f7 9a 12 34 56 	negw   0x7f563412(%r10)
 2bc:	7f 
 2bd:	66 41 f7 9b 12 34 56 	negw   0x7f563412(%r11)
 2c4:	7f 
 2c5:	66 41 f7 9c 24 12 34 	negw   0x7f563412(%r12)
 2cc:	56 7f 
 2ce:	66 41 f7 9d 12 34 56 	negw   0x7f563412(%r13)
 2d5:	7f 
 2d6:	66 41 f7 9e 12 34 56 	negw   0x7f563412(%r14)
 2dd:	7f 
 2de:	66 41 f7 9f 12 34 56 	negw   0x7f563412(%r15)
 2e5:	7f 
 2e6:	90                   	nop
 2e7:	f7 98 12 34 56 7f    	negl   0x7f563412(%rax)
 2ed:	f7 99 12 34 56 7f    	negl   0x7f563412(%rcx)
 2f3:	f7 9a 12 34 56 7f    	negl   0x7f563412(%rdx)
 2f9:	f7 9b 12 34 56 7f    	negl   0x7f563412(%rbx)
 2ff:	f7 9c 24 12 34 56 7f 	negl   0x7f563412(%rsp)
 306:	f7 9d 12 34 56 7f    	negl   0x7f563412(%rbp)
 30c:	f7 9e 12 34 56 7f    	negl   0x7f563412(%rsi)
 312:	f7 9f 12 34 56 7f    	negl   0x7f563412(%rdi)
 318:	41 f7 98 12 34 56 7f 	negl   0x7f563412(%r8)
 31f:	41 f7 99 12 34 56 7f 	negl   0x7f563412(%r9)
 326:	41 f7 9a 12 34 56 7f 	negl   0x7f563412(%r10)
 32d:	41 f7 9b 12 34 56 7f 	negl   0x7f563412(%r11)
 334:	41 f7 9c 24 12 34 56 	negl   0x7f563412(%r12)
 33b:	7f 
 33c:	41 f7 9d 12 34 56 7f 	negl   0x7f563412(%r13)
 343:	41 f7 9e 12 34 56 7f 	negl   0x7f563412(%r14)
 34a:	41 f7 9f 12 34 56 7f 	negl   0x7f563412(%r15)
 351:	90                   	nop
 352:	48 f7 98 12 34 56 7f 	negq   0x7f563412(%rax)
 359:	48 f7 99 12 34 56 7f 	negq   0x7f563412(%rcx)
 360:	48 f7 9a 12 34 56 7f 	negq   0x7f563412(%rdx)
 367:	48 f7 9b 12 34 56 7f 	negq   0x7f563412(%rbx)
 36e:	48 f7 9c 24 12 34 56 	negq   0x7f563412(%rsp)
 375:	7f 
 376:	48 f7 9d 12 34 56 7f 	negq   0x7f563412(%rbp)
 37d:	48 f7 9e 12 34 56 7f 	negq   0x7f563412(%rsi)
 384:	48 f7 9f 12 34 56 7f 	negq   0x7f563412(%rdi)
 38b:	49 f7 98 12 34 56 7f 	negq   0x7f563412(%r8)
 392:	49 f7 99 12 34 56 7f 	negq   0x7f563412(%r9)
 399:	49 f7 9a 12 34 56 7f 	negq   0x7f563412(%r10)
 3a0:	49 f7 9b 12 34 56 7f 	negq   0x7f563412(%r11)
 3a7:	49 f7 9c 24 12 34 56 	negq   0x7f563412(%r12)
 3ae:	7f 
 3af:	49 f7 9d 12 34 56 7f 	negq   0x7f563412(%r13)
 3b6:	49 f7 9e 12 34 56 7f 	negq   0x7f563412(%r14)
 3bd:	49 f7 9f 12 34 56 7f 	negq   0x7f563412(%r15)
 3c4:	c3                   	retq   
