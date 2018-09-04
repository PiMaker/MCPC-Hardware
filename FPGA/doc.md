FPGA
====

Status output:
- 7-segment hex display:
    - SW9: Register H
    - SW8: PC
    - SW7: Instruction buffer
    - no switch up: 0xdead
- Status LEDs:
    - LEDR0: Halted
    - LEDR1: reg_we bit
    - LEDR2: continue_execution bit
    - LEDR6: STATE: INS_LOAD
    - LEDR7: STATE: WAITING
    - LEDR8: STATE: COMMIT
    - LEDR9: STATE: PC_INC
- SW0:
    - Up:
        - SW1:
            - Up: Auto clock
            - Down: Debug clock
    - Down: Manual clock



Assembly language instruction set
=================================

Memory locations:
- (EEPROM)
- SRAM
- Registers (A, B, C, D, E, F, G, H) = 0x0-0x7 (NOTE: Content of register H is displayed on a simple hex 7-segment)
    - Special Registers - SCR1 (Scratch) = 0x8, SCR2 (Scratch) = 0x9, SP (Stack Pointer) = 0xA, PC (Program Counter) = 0xB, 0 = 0xC, 1 = 0xD, -1 = 0xE, BUS (Last Received Bus Data) = 0xF

Architecture is Little Endian: LSB is rightmost
Scratch registers: General purpose, but not guaranteed to be consistent after library calls (but they are consistent if only base commands are used)

Instructions
------------

Data instructions:
+ (0x0) HALT
+ (0x1) MOV from.reg to.reg - moves data between registers
+ (0x2) MOVNZ from.reg to.reg if.reg - performs a MOV if the value in if.reg is not 0 (contains at least one set bit)
+ (0x3) MOVEZ from.reg to.reg if.reg - performs a MOV if the value in if.reg is 0 (contains no set bits)
+ (0x4) BUS data.reg addr.lit3 - sends data from register data.reg to the bus address specified in addr.lit3
+ (0x5) LROM from.reg to.reg - writes data to the EEPROM
+ (0x6) SET <ignored*> addr.reg - writes the following instruction directly to the register in addr (and skips it, thus not executing the following instruction) - Setting PC is NOT possible!
+ (0x7) Extended instruction set header (extensible platform protocol)

*ignored means 4 bit that have to effect but padding; these are not present in assembler form

ALU instructions:
+ Format: INS op1.reg out.reg <op2.reg/op2.lit4*>
+ INS: AND, OR, NOT, ADD, SHFT*, MUL, GT, EQ
+ (0x8-0xF)

Compiled instructions (multiple instructions, single assembly command):
- NEG from.reg to.reg = NOT from 0 to, ADD to 1 to - negates a 2's complement number, also converts from and to 2's complement
- JMP to.reg = MOV to PC - Jump unconditionally to an address in a register
- JMP to.lit16 = SET SCR, to, MOV SCR PC - Jump unconditionally to a literal address
- JMPNZ to.reg if.reg = MOVNZ to PC if - Jump to register to.reg if if.reg is not zero
- JMPEZ to.reg if.reg = MOVEZ to PC if - Jump to register to.reg if if.reg is equal to zero
- JMPGT to.reg val.reg cmpto.reg = GT val cmpto SCR, MOVNZ to PC SCR - Jump to register to.reg if val.reg is greater than cmpto.reg
- CALL addr.reg = ADD SP -1 SP, MOV PC SP, MOV addr PC, ADD SP 1 SP - Calls a function at addr.reg
- RET = MOV SP PC - Returns from a function
- STOR data.reg addr.reg = SHFT addr -0x7 SCR, SHFT SCR 0x3 SCR, OR SCR 1 SCR, BUS SCR 0x1, BUS data 0x1 - Stores data in SRAM
- LOAD data.reg addr.reg = SHFT addr -0x7 SCR, SHFT SCR 0x3 SCR, BUS SCR 0x1, HOLD, MOV BUS data - Retrieves data from SRAM
- INC val.reg = ADD val 1 val - Increments by one
- DEC val.reg = ADD val -1 val - Decrements by one
- NOOP = MOV SCR SCR - does nothing for one cycle

*NOTE: JMP/JMPxx instructions always jump to the instruction address AFTER the one specified

Fibonacci calculator
--------------------

; Setup variables
MOV 1 A
MOV 1 B
MOV 1 D
; Set jump address
SET E
0x0007
; Set number of fibonacci numbers to calculate
SET F
0x0020
; Perform fibonacci algorithm
ADD A B C
MOV B A
MOV C B
; Store output in consecutive RAM addresses
STOR A D
INC D
; Loop
JMPGT E F D
; End
