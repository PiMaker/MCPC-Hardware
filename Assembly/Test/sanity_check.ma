SET B
0xDFFF
SET D
0x0041

; Print something
.start INC B
STOR D B

; Set value to H
SET H
0xa1b8

; Store H in RAM
STORLA -1 0x0b0
STORLA H 0x0b1
STORLA -1 0x0b2

; Load zero into H again
MOV 0 H

; Restore H from RAM
LOADLA SCR1 0x0b0
LOADLA H 0x0b1
LOADLA SCR2 0x0b2

; Test ALU
OR 1 H H
SET C
0xa1b9

; Jump to beginning
JMPEQ .start C H
