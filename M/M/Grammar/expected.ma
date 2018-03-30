; Declare statements
#declare C var_printc_c
#declare C var_main_x
#declare D var_main_y
#declare E var_main_z


; Static setup code
SET SP
0x01FF
STORL .halt SP
CALL .func_main
.halt HALT


; function printc / load parameter c (to regA)
.func_printc POP H
POP var_printc_c

; calc and set c << 8 (shift is special case)
SHFT var_printc_c G -0x4
SHFT G G -0x4
MOV G var_printc_c

; asm block
BUS C 0x2

; auto-return
RETH


; function main / set x
.func_main POP H
SET G
0x5
MOV G var_main_x

; set y
SET G
0x14
MOV G var_main_y

; calc z = 3 y * x +
SET G
0x3
PUSH G
PUSH var_main_y
POP A
POP B
MUL A G B
PUSH G
PUSH var_main_x
POP A
POP B
ADD A G B
PUSH G
POP var_main_z

; calculate parameter (char)z
SET A
0x00FF
AND var_main_z G A

; push parameter to stack
PUSH G

; call printc
CALL .func_printc

; auto-return, but end of main
RETH