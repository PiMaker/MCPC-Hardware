;autotest reg=0 val=7;
; Generated using MSCR compiler version 0.2.1

JMP .mscr_init_main

.mscr_data __LABEL_SET
0x48
0x65
0x6c
0x6c
0x6f
0x2c
0x20
0x77
0x6f
0x72
0x6c
0x64
0x21
0x0

; MSCR initialization routine
.mscr_init_main __LABEL_SET
SET SP ; Stack
.mscr_data_end __LABEL_SET ; _data_end label has to be one word after code start, because reading in bootloaderInitAsm is technically off by one for performance reasons
0x7FFF
SET H ; VarHeap
.mscr_code_end

CALL .mscr_init_userland ; Call program specific initialization

PUSH 0 ; argp
PUSH 0 ; argc
CALL .mscr_function_main_params_2 ; Call userland main

; After main, copy exit code to H to show on hex-display (but keep in A for autotest!)
MOV A H

HALT ; After execution, halt


; MSCR bootloader static value loader
.mscr_init_bootloader SET A
.mscr_data_end ; Data block end address = Code block start address
SETREG B 0x0003 ; Data start
SETREG C 0xD003 ; Start of readonly CFG region for bootloader ROM + offset for data start

.mscr_init_bootloader_loop_start __LABEL_SET
LOAD D C ; Read from ROM to regD
STOR D B ; Write to RAM
INC C ; Increment read address
INC B ; Increment write address
EQ A D B ; Check if we reached end of data and jump accordingly
JMPNZ .mscr_init_bootloader_return D
JMP .mscr_init_bootloader_loop_start

.mscr_init_bootloader_return RET ; Return out



.mscr_init_userland __LABEL_SET
CALL .mscr_init_bootloader
RET ;Userland init end
.mscr_function_main_params_2 __LABEL_SET ; [Function (in func: main)]

SETREG F 0x2 ; CALC: literal 2
MOV F G ; [Function (in func: main)]
ADD G H H ; [Function (in func: main)]
POP E ; [Function (in func: main)]
POP D ; [Function (in func: main)] (reg_alloc: checked out as dirty)
POP C ; [Function (in func: main)] (reg_alloc: checked out as dirty)
PUSH E ; [Function (in func: main)]
SETREG F 0x3 ; CALC: literal 3
PUSH F ; CALC: push operand
SETREG F 0x77 ; CALC: literal 119
PUSH F ; CALC: push operand

SETREG G 0x0 ; __FLUSHSCOPE (flushing: argp)
SUB H G G ; __FLUSHSCOPE (flushing: argp)
STOR D G ; __FLUSHSCOPE (flushing: argp)
SETREG G 0x1 ; __FLUSHSCOPE (flushing: argc)
SUB H G G ; __FLUSHSCOPE (flushing: argc)
STOR C G ; __FLUSHSCOPE (flushing: argc)

CALL .mscr_function_indexOf_params_2
PUSH A

POP F
MOV F A ; [Body (in func: main)]
SETREG F 0x2 ; CALC: literal 2
MOV F G ; [Body (in func: main)]
SUB H H G ; [Body (in func: main)]

RET ; [Body (in func: main)]
FAULT 0x0 ; Ending function: main [Function (in func: )]
.mscr_function_indexOf_params_2 __LABEL_SET ; [Function (in func: indexOf)]

SETREG F 0x5 ; CALC: literal 5
MOV F G ; [Function (in func: indexOf)]
ADD G H H ; [Function (in func: indexOf)]
POP E ; [Function (in func: indexOf)]
POP D ; [Function (in func: indexOf)] (reg_alloc: checked out as dirty)
POP C ; [Function (in func: indexOf)] (reg_alloc: checked out as dirty)
PUSH E ; [Function (in func: indexOf)]
SETREG F 0x0 ; CALC: literal 0
MOV F B ; [Variable (in func: indexOf)] (reg_alloc: checked out as dirty)

SETREG G 0x0 ; __FLUSHSCOPE (flushing: needle)
SUB H G G ; __FLUSHSCOPE (flushing: needle)
STOR D G ; __FLUSHSCOPE (flushing: needle)
SETREG G 0x1 ; __FLUSHSCOPE (flushing: haystack)
SUB H G G ; __FLUSHSCOPE (flushing: haystack)
STOR C G ; __FLUSHSCOPE (flushing: haystack)
SETREG G 0x2 ; __FLUSHSCOPE (flushing: i)
SUB H G G ; __FLUSHSCOPE (flushing: i)
STOR B G ; __FLUSHSCOPE (flushing: i)

.mscr_while_start__12_5_157 __LABEL_SET ; [WhileLoop (in func: indexOf)]
SETREG F 0x1 ; CALC: literal 1
JMPEZ .mscr_while_end__12_5_157 F ; [WhileLoop (in func: indexOf)]
SETREG G 0x1
SUB H G G
LOAD D G
MOV D F ; CALC: var haystack (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG G 0x2
SUB H G G
LOAD D G
MOV D F ; CALC: var i (reg_alloc: checked out as clean)
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
MOV F D ; [Variable (in func: indexOf)] (reg_alloc: checked out as dirty)
LOAD F F
MOV F C ; [Variable (in func: indexOf)] (reg_alloc: checked out as dirty)

SETREG G 0x3 ; __FLUSHSCOPE (flushing: hsi)
SUB H G G ; __FLUSHSCOPE (flushing: hsi)
STOR D G ; __FLUSHSCOPE (flushing: hsi)
SETREG G 0x4 ; __FLUSHSCOPE (flushing: charAt)
SUB H G G ; __FLUSHSCOPE (flushing: charAt)
STOR C G ; __FLUSHSCOPE (flushing: charAt)

SETREG G 0x4
SUB H G G
LOAD D G
MOV D F ; CALC: var charAt (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG G 0x0
SUB H G G
LOAD D G
MOV D F ; CALC: var needle (reg_alloc: checked out as clean)
MOV F E ; CALC: push operand
POP F
EQ F F E ; CALC: operator EQ
PUSH F
SETREG G 0x4
SUB H G G
LOAD D G
MOV D F ; CALC: var charAt (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x0 ; CALC: literal 0
MOV F E ; CALC: push operand
POP F
EQ F F E ; CALC: operator EQ
MOV F E
POP F
OR F F E ; CALC: operator OR
JMPEZ .mscr_cond_else__15_9_242 F ; [IfCondition (in func: indexOf)]
SETREG G 0x2
SUB H G G
LOAD D G
MOV D A ; [BodyIf (in func: indexOf)] (reg_alloc: checked out as clean)
SETREG F 0x5 ; CALC: literal 5
MOV F G ; [BodyIf (in func: indexOf)]
SUB H H G ; [BodyIf (in func: indexOf)]

RET ; [BodyIf (in func: indexOf)]

JMP .mscr_cond_end__15_9_242 ; [BodyElse (in func: indexOf)]

.mscr_cond_else__15_9_242 __LABEL_SET ; [BodyElse (in func: indexOf)]


.mscr_cond_end__15_9_242 __LABEL_SET ; [IfCondition (in func: indexOf)]
SETREG G 0x2
SUB H G G
LOAD D G
MOV D F ; CALC: var i (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x1 ; CALC: literal 1
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
MOV F D ; [Assignment (in func: indexOf)] (reg_alloc: var found checked out in 3)

SETREG G 0x2 ; __FLUSHSCOPE (flushing: i)
SUB H G G ; __FLUSHSCOPE (flushing: i)
STOR D G ; __FLUSHSCOPE (flushing: i)
JMP .mscr_while_start__12_5_157 ; [WhileLoop (in func: indexOf)]
.mscr_while_end__12_5_157 __LABEL_SET ; [WhileLoop (in func: indexOf)]

SETREG F 0x1 ; CALC: literal 1
NEG F F
MOV F A ; [Body (in func: indexOf)]
SETREG F 0x5 ; CALC: literal 5
MOV F G ; [Body (in func: indexOf)]
SUB H H G ; [Body (in func: indexOf)]

RET ; [Body (in func: indexOf)]
FAULT 0x0 ; Ending function: indexOf [Function (in func: )]
.mscr_code_end HALT