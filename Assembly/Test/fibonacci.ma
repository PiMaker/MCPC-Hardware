#declare A numA
#declare B numB
#declare F maxFibs
#declare C acc
#declare H output
#declare D count

; Setup variables
MOV 1 numA
MOV 1 numB
MOV 0 count
SET maxFibs
0x0020 ; Set number of fibonaccis to calculate
; Perform fibonacci algorithm
.fibstart ADD numA acc numB
MOV numB numA
MOV acc numB
; Output current number on hex display
MOV numA output
; Loop
INC count
JMPGT .fibstart maxFibs count
; End
HALT
