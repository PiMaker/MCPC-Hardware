; Generated using MSCR compiler version 0.3

JMP .mscr_init_main

.mscr_data __LABEL_SET
0x0
0x0
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
.mscr_function_vga_printChar_params_1 __LABEL_SET ; [Function (in func: vga_printChar)]

SETREG G 0x1 ; [Function (in func: vga_printChar)]
ADD G H H ; [Function (in func: vga_printChar)]
POP E ; [Function (in func: vga_printChar)]
POP D ; [Function (in func: vga_printChar)] (reg_alloc: checked out as dirty)
PUSH E ; [Function (in func: vga_printChar)]
MOV D F ; CALC: var char (reg_alloc: var found checked out in 3)
PUSH F ; call to $$ [FunctionCall (in func: vga_printChar)]
SETREG F 0xe000 ; CALC: literal 0xE000
PUSH F ; CALC: push operand
SETREG G 0x4
LOAD C G
MOV C F ; CALC: var vga_buf_pos_y (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x3c ; CALC: literal 60
MOV F E ; CALC: push operand
POP F
MUL F F E ; CALC: operator MUL
MOV F E
POP F
ADD F F E ; CALC: operator ADD
PUSH F
SETREG G 0x3
LOAD C G
MOV C F ; CALC: var vga_buf_pos_x (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
POP G ; call to $$ [FunctionCall (in func: vga_printChar)]
STOR G F ; call to $$ [FunctionCall (in func: vga_printChar)]
SETREG G 0x3
LOAD C G
MOV C F ; CALC: var vga_buf_pos_x (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x1 ; CALC: literal 1
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
MOV F C ; [Assignment (in func: vga_printChar)] (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
SETREG G 0x3 ; (reg_alloc: directly assigned, evicting back immediately)
STOR C G

SETREG G 0x0 ; __FLUSHSCOPE (flushing: char)
SUB H G G ; __FLUSHSCOPE (flushing: char)
STOR D G ; __FLUSHSCOPE (flushing: char)

SETREG G 0x3
LOAD D G
MOV D F ; CALC: var vga_buf_pos_x (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x50 ; CALC: literal 80
MOV F E ; CALC: push operand
POP F
EQ F F E ; CALC: operator EQ
JMPEZ .mscr_cond_else__27_5_226 F ; [IfCondition (in func: vga_printChar)]
SETREG F 0x0 ; CALC: literal 0
MOV F D ; [Assignment (in func: vga_printChar)] (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
SETREG G 0x3 ; (reg_alloc: directly assigned, evicting back immediately)
STOR D G
SETREG G 0x4
LOAD D G
MOV D F ; CALC: var vga_buf_pos_y (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x1 ; CALC: literal 1
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
MOV F D ; [Assignment (in func: vga_printChar)] (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
SETREG G 0x4 ; (reg_alloc: directly assigned, evicting back immediately)
STOR D G


SETREG G 0x4
LOAD D G
MOV D F ; CALC: var vga_buf_pos_y (reg_alloc: skipping scope search, directly-assigned) (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x3c ; CALC: literal 60
MOV F E ; CALC: push operand
POP F
EQ F F E ; CALC: operator EQ
JMPEZ .mscr_cond_else__31_9_321 F ; [IfCondition (in func: vga_printChar)]

JMP .mscr_cond_end__31_9_321 ; [BodyElse (in func: vga_printChar)]

.mscr_cond_else__31_9_321 __LABEL_SET ; [BodyElse (in func: vga_printChar)]


.mscr_cond_end__31_9_321 __LABEL_SET ; [IfCondition (in func: vga_printChar)]

JMP .mscr_cond_end__27_5_226 ; [BodyElse (in func: vga_printChar)]

.mscr_cond_else__27_5_226 __LABEL_SET ; [BodyElse (in func: vga_printChar)]


.mscr_cond_end__27_5_226 __LABEL_SET ; [IfCondition (in func: vga_printChar)]
SETREG F 0x1 ; CALC: literal 1
MOV F A ; [Body (in func: vga_printChar)]
SETREG G 0x1 ; [Body (in func: vga_printChar)]
SUB H H G ; [Body (in func: vga_printChar)]

RET ; [Body (in func: vga_printChar)]
FAULT 0x0 ; Ending function: vga_printChar [Function (in func: )]
.mscr_function_vga_printString_params_1 __LABEL_SET ; [Function (in func: vga_printString)]

SETREG G 0x3 ; [Function (in func: vga_printString)]
ADD G H H ; [Function (in func: vga_printString)]
POP E ; [Function (in func: vga_printString)]
POP D ; [Function (in func: vga_printString)] (reg_alloc: checked out as dirty)
PUSH E ; [Function (in func: vga_printString)]
SETREG F 0x0 ; CALC: literal 0
MOV F C ; [Variable (in func: vga_printString)] (reg_alloc: checked out as dirty)
MOV D F ; CALC: var str (reg_alloc: var found checked out in 3)
LOAD F F
MOV F B ; [Variable (in func: vga_printString)] (reg_alloc: checked out as dirty)

SETREG G 0x0 ; __FLUSHSCOPE (flushing: str)
SUB H G G ; __FLUSHSCOPE (flushing: str)
STOR D G ; __FLUSHSCOPE (flushing: str)
SETREG G 0x1 ; __FLUSHSCOPE (flushing: i)
SUB H G G ; __FLUSHSCOPE (flushing: i)
STOR C G ; __FLUSHSCOPE (flushing: i)
SETREG G 0x2 ; __FLUSHSCOPE (flushing: charAt)
SUB H G G ; __FLUSHSCOPE (flushing: charAt)
STOR B G ; __FLUSHSCOPE (flushing: charAt)

.mscr_while_start__44_5_492 __LABEL_SET ; [WhileLoop (in func: vga_printString)]
SETREG G 0x2
SUB H G G
LOAD D G
MOV D F ; CALC: var charAt (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x0 ; CALC: literal 0
MOV F E ; CALC: push operand
POP F
NEQ F F E ; CALC: operator NEQ
JMPEZ .mscr_while_end__44_5_492 F ; [WhileLoop (in func: vga_printString)]
MOV D F ; CALC: var charAt (reg_alloc: var found checked out in 3)
PUSH F ; [FunctionCall (in func: vga_printString)]


CALL .mscr_function_vga_printChar_params_1 ; [FunctionCall (in func: vga_printString)]

SETREG G 0x1
SUB H G G
LOAD D G
MOV D F ; CALC: var i (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
SETREG F 0x1 ; CALC: literal 1
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
MOV F D ; [Assignment (in func: vga_printString)] (reg_alloc: var found checked out in 3)
SETREG G 0x0
SUB H G G
LOAD C G
MOV C F ; CALC: var str (reg_alloc: checked out as clean)
PUSH F ; CALC: push operand
MOV D F ; CALC: var i (reg_alloc: var found checked out in 3)
MOV F E ; CALC: push operand
POP F
ADD F F E ; CALC: operator ADD
LOAD F F
MOV F C ; [Assignment (in func: vga_printString)] (reg_alloc: checked out as dirty)

SETREG G 0x1 ; __FLUSHSCOPE (flushing: i)
SUB H G G ; __FLUSHSCOPE (flushing: i)
STOR D G ; __FLUSHSCOPE (flushing: i)
SETREG G 0x2 ; __FLUSHSCOPE (flushing: charAt)
SUB H G G ; __FLUSHSCOPE (flushing: charAt)
STOR C G ; __FLUSHSCOPE (flushing: charAt)
JMP .mscr_while_start__44_5_492 ; [WhileLoop (in func: vga_printString)]
.mscr_while_end__44_5_492 __LABEL_SET ; [WhileLoop (in func: vga_printString)]

SETREG G 0x3 ; [Function (in func: )]
SUB H H G ; [Function (in func: )]

RET ; [Function (in func: )]
FAULT 0x0 ; Ending function: vga_printString [Function (in func: )]
.mscr_function_main_params_2 __LABEL_SET ; [Function (in func: main)]

SETREG G 0x2 ; [Function (in func: main)]
ADD G H H ; [Function (in func: main)]
POP E ; [Function (in func: main)]
POP D ; [Function (in func: main)] (reg_alloc: checked out as dirty)
POP C ; [Function (in func: main)] (reg_alloc: checked out as dirty)
PUSH E ; [Function (in func: main)]
SETREG F 0x5 ; CALC: literal 5
PUSH F ; [FunctionCall (in func: main)]

SETREG G 0x0 ; __FLUSHSCOPE (flushing: argp)
SUB H G G ; __FLUSHSCOPE (flushing: argp)
STOR D G ; __FLUSHSCOPE (flushing: argp)
SETREG G 0x1 ; __FLUSHSCOPE (flushing: argc)
SUB H G G ; __FLUSHSCOPE (flushing: argc)
STOR C G ; __FLUSHSCOPE (flushing: argc)

CALL .mscr_function_vga_printString_params_1 ; [FunctionCall (in func: main)]

SETREG F 0x7353 ; CALC: literal 0x7353
MOV F A ; [Body (in func: main)]
SETREG G 0x2 ; [Body (in func: main)]
SUB H H G ; [Body (in func: main)]

RET ; [Body (in func: main)]
FAULT 0x0 ; Ending function: main [Function (in func: )]
.mscr_code_end HALT