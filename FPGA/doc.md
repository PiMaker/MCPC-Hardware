FPGA
====

Status output:
- 7-segment hex display:
    - SW9: Register H
    - SW8: PC
    - SW7: Instruction buffer
    - SW6: Memory FIFO data read
    - no switch up: Debug information ~0xdead~
- Status LEDs:
    - LEDR0: Halted
    - LEDR1: reg_we bit
    - LEDR2: continue_execution bit
    - LEDR3: Debugger enabled
    - LEDR4: CLK
    - LEDR5: IF_EN
    - LEDR6: STATE: INS_LOAD
    - LEDR7: STATE: WAITING
    - LEDR8: STATE: COMMIT
    - LEDR9: STATE: PC_INC
- SW0:
    - Up: Auto clock
    - Down: Manual clock


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

=> CFG table layout

= General/Memory control =
* [0x8000] r ... MCPC version number
* [0x8001] rw(k) ... RAM instruction loading (RIL; loads instruction from RAM instead of EEPROM)
* [0x8002] rw(k) ... RIL PC swap (program counter value to load on enabling/disabling RIL)
* [0x8003] rw(k) ... Protected memory page (Instructions executed as RIL when execute-select page matches this are executed with kernel priviledges, defaults to 0)
* [0x8004] rw(k) ... Kernel memory bits (see [0x8005] below)
* [0x8800] rw ... Memory high bits (RAM uses 25 bit addressing, lower 15 bit are set by MEMR/MEMW directly, the next 5 bits are loaded from the lower bits of this CFG register, the missing 5 bits are the lower kernel memory bits)

= Bootloader ROM =
* [0xD000-0xD800] r ... Bootloader ROM read-only access

= VGA subsystem =
* [0xE000-0xF2BF] w ... ASCII buffer (80x60 display, one ascii characted per data word, high bits unused)

= Interrupt/Timer subsystem =
TBD

= Virtualization subsystem =
TBD


On-Chip hardware debugger
-------------------------

UART (115200 baud, n/8/1 mode)
Pins: ARDUINO[0] = RX, ARDUINO[1] = TX

Command syntax:
 - 8 bit
 - bits 7-4: data
 - bits 3-0: OP-code

OP-codes:
 DEBUGGER_OPCODE_GET 4'h1  = Set read address (3 bit) and print contents of specified register
 DEBUGGER_OPCODE_SET 4'h2  = Set write address (3 bit)
 DEBUGGER_OPCODE_HI 4'h4   = Write data to high bits at write address
 DEBUGGER_OPCODE_LO 4'h8   = Write data to low bits at write address
 DEBUGGER_OPCODE_STEP 4'hC = Execute a single processor instruction and print FF when done

Debug-Registers:
= Writeable:
0b0: Debug enable (dbgClk enable)
0b1: CPU Register address overwrite enable
0b2: Reset request
0b3: Instruction overwride enable
1b3-0: CPU Register address overwrite value
4b7-0: Low-bits of instruction overwrite buffer (first word)
5b7-0: High-bits of instruction overwrite buffer (first word)
6b7-0: Low-bits of instruction overwrite buffer (second word)
7b7-0: High-bits of instruction overwrite buffer (second word)
= Readable:
8b7-0: Low-bit content of currently selected CPU register
9b7-0: High-bit content of currently selected CPU register
Fb0: Halted




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
