; Setup variables
MOV 1 A
MOV 1 B
MOV 1 D
; Set number of fibonacci numbers to calculate
SET F
0x0020
; Perform fibonacci algorithm
.fibstart
ADD A B C
MOV B A
MOV C B
; Store output in consecutive RAM addresses
STOR A D
INC D
; Loop
JMPGT .fibstart F D
; End