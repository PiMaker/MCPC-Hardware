; Generated using MSCR compiler version 0.1.3

JMP .mscr_init_main

0x4000 ; HSP

.mscr_data __LABEL_SET
0x20

; MSCR initialization routine
.mscr_init_main __LABEL_SET
SET SP ; Stack
0x7FFF
SET H ; VarHeap
.mscr_code_end

CALL .mscr_init_userland ; Call program specific initialization

MOV 0 A
PUSH 0
CALL .mscr_function_main_params_2 ; Call userland main

; After main, put exit code to H for now
MOV A H

HALT ; After execution, halt

.mscr_init_userland __LABEL_SET
RET ;Userland init end
.mscr_function_main_params_2 __LABEL_SET ; [Function (in func: main)]
SETREG F 0x3 ; CALC: literal 3
MOV F G ; [Function (in func: main)]
ADD G H H ; [Function (in func: main)]
POP E ; [Function (in func: main)]
POP B ; [Function (in func: main)]
PUSH E ; [Function (in func: main)]
SETREG G 0x0
SUB H G G
STOR A G
SETREG G 0x4
LOAD A G
MOV A F ; CALC: var test
PUSH F ; CALC: push operand
SETREG G 0x0
SUB H G G
LOAD A G
MOV A F ; CALC: var argc
PUSH F ; CALC: push operand
POP E
POP F
SUB F F E
PUSH F
SETREG F 0x2 ; CALC: literal 2
PUSH F ; CALC: push operand
SETREG F 0x2 ; CALC: literal 2
PUSH F ; CALC: push operand
SETREG F 0x7 ; CALC: literal 7
PUSH F ; CALC: push operand
POP E
POP F
MUL F F E
PUSH F
POP E
POP F
MUL F F E
PUSH F
POP E
POP F
ADD F F E
MOV F C ; [Variable (in func: main)]
MOV C A ; [Body (in func: main)]
SETREG F 0x3 ; CALC: literal 3
MOV F G ; [Body (in func: main)]
SUB H H G ; [Body (in func: main)]
RET ; [Body (in func: main)]
FAULT 0x0 ; Ending function: main [Function (in func: )]
.mscr_code_end HALT