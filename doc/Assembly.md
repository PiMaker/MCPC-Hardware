Assembly language instruction set
=================================

**Memory locations**

* EEPROM (read-only)
* SRAM
* General purpose registers (A, B, C, D, E, F, G, H) = h0-h7 (NOTE: Content of register H is displayed on a simple hex 7-segment)
* Special registers - SCR1 (Scratch) = h8, SCR2 (Scratch) = h9, SP (Stack Pointer) = hA, PC (Program Counter) = hB, 0 = hC, 1 = hD, -1 = hE, BUS (Last Received Bus Data, Deprecated) = hF

**Misc. Details**

* 16 bit words, no byte addressing
* Architecture is Little Endian: LSB is rightmost
* Scratch registers: General purpose, but not guaranteed to be consistent after library calls (but they are consistent if only base commands are used)


Instructions
------------

**Data instructions**

* (h0) HALT
* (h1) MOV from.reg to.reg - moves data between registers
* (h2) MOVNZ from.reg to.reg if.reg - performs a MOV if the value in if.reg is not 0 (contains at least one set bit)
* (h3) MOVEZ from.reg to.reg if.reg - performs a MOV if the value in if.reg is 0 (contains no set bits)
* (h4) BUS data.reg addr.lit3 - sends data from register data.reg to the bus address specified in addr.lit3 (deprecated)
* (h5) MEMR [see extended information below]
* (h6) SET <ignored*> addr.reg - writes the following instruction directly to the register in addr (and skips it, thus not executing the following instruction) - Setting PC is NOT possible!
* (h7) MEMW [see extended information below]

*ignored means 4 bits that have no effect but padding; these are not present in assembler form

**ALU instructions**

* Format: INS op1.reg out.reg op2.reg
* INS: AND, OR, XOR, ADD, SHFT*, MUL, GT, EQ (h8-hF)

*A note on SHFT: The input value is shifted by the operator input masked via 16'h00FF, the other bits (Mask: 16'hFF00) are used to determine shifting direction. A left shift is performed iff any bit in the direction-masked input is 1, a right shift otherwise. All shifts are logical, not numerical.


Memory access and layout
------------------------

**MEMR instruction reference**
```
    .... xxxx xxxx 0101 (16 bit instruction)
    |    |    |    |
    |    |    |    + MEMR instruction code h5
    |    |    |
    |    |    + Register containing memory address
    |    |
    |    + Register to write read result to
    |
    + unused
```

**MEMW instruction reference**
```
    xxxx .... xxxx 0111 (16 bit instruction)
    |    |    |    |
    |    |    |    + MEMW instruction code h7
    |    |    |
    |    |    + Register containing memory address
    |    |
    |    + unused
    |
    + Register containing data to write
```

**Memory address definition**

* [hF] ... CFG bit (if set, action will be performed on CFG table entry)
* [h0-hE] ... Access address


Fibonacci calculator
--------------------
(example assembler code)

```asm
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
```
