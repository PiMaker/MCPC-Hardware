Assembly language instruction set
=================================

Memory locations:
- (EEPROM)
- SRAM
- Registers (A, B, C, D, E, F, G, H) = 0x0-0x7 (NOTE: Content of register H is displayed on a simple hex 7-segment)
    - Special Registers - SCR1 (Scratch) = 0x8, SCR2 (Scratch) = 0x9, SP (Stack Pointer) = 0xA, PC (Program Counter) = 0xB, 0 = 0xC, 1 = 0xD, -1 = 0xE, BUS (Last Received Bus Data) = 0xF

16 bit words, no byte addressing
Architecture is Little Endian: LSB is rightmost
Scratch registers: General purpose, but not guaranteed to be consistent after library calls (but they are consistent if only base commands are used)


Instructions
------------

Data instructions:
+ (0x0) HALT
+ (0x1) MOV from.reg to.reg - moves data between registers
+ (0x2) MOVNZ from.reg to.reg if.reg - performs a MOV if the value in if.reg is not 0 (contains at least one set bit)
+ (0x3) MOVEZ from.reg to.reg if.reg - performs a MOV if the value in if.reg is 0 (contains no set bits)
+ (0x4) BUS data.reg addr.lit3 - sends data from register data.reg to the bus address specified in addr.lit3 (deprecated)
+ (0x5) MEMR [see extended information below]
+ (0x6) SET <ignored*> addr.reg - writes the following instruction directly to the register in addr (and skips it, thus not executing the following instruction) - Setting PC is NOT possible!
+ (0x7) MEMW [see extended information below]

*ignored means 4 bit that have to effect but padding; these are not present in assembler form

ALU instructions:
+ Format: INS op1.reg out.reg <op2.reg/op2.lit4*>
+ INS: AND, OR, XOR, ADD, SHFT*, MUL, GT, EQ
+ (0x8-0xF)

Compiled instructions (multiple instructions, single assembly command): See base.mlib


Memory access and layout
------------------------

=> MEMR instruction reference:

    .... xxxx xxxx 0101 (16 bit instruction)
    |    |    |    |
    |    |    |    + MEMR instruction code 0x5
    |    |    |
    |    |    + Register containing memory address
    |    |
    |    + Register to write read result to
    |
    + unused

=> MEMW instruction reference

    xxxx .... xxxx 0111 (16 bit instruction)
    |    |    |    |
    |    |    |    + MEMW instruction code 0x7
    |    |    |
    |    |    + Register containing memory address
    |    |
    |    + unused
    |
    + Register containing data to write

=> Memory address definition

* [0xF] ... CFG bit (if set, action will be performed on CFG table entry)
* [0x0-0xE] ... Access address


Fibonacci calculator
--------------------
(example assembler code)

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
