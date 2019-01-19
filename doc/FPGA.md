FPGA
====

Physical control layout
-----------------------

Status output:

* 7-segment hex display:
  * SW9: Register H (SW5 up: from IRQ regs)
  * SW8: PC
  * SW7: Instruction buffer
  * SW6: Memory FIFO data read
  * no switch up: Debug information
* Status LEDs:
  * LEDR0: Halted
  * LEDR1: In IRQ handler/context
  * LEDR2: continue_execution bit
  * LEDR3: Debugger enabled
  * LEDR4: CLK
  * LEDR5: IF_EN
  * LEDR6: STATE: INS_LOAD
  * LEDR7: STATE: WAITING
  * LEDR8: STATE: COMMIT
  * LEDR9: STATE: PC_INC
* SW0:
  * Up: Auto clock
  * Down: Manual clock
* SW1:
  * On change: Trigger IRQ for keycode 'A'


On-Chip hardware debugger
-------------------------

UART (115200 baud, n/8/1 mode)  
Pins: ARDUINO[0] = RX, ARDUINO[1] = TX  

**Command syntax**

* 8 bit
* bits 7-4: data
* bits 3-0: OP-code

**OP-codes**

| Name | Address | Description |
|---|-|----|
| DEBUGGER_OPCODE_GET | h1 | Set read address (3 bit) and print contents of specified register |
| DEBUGGER_OPCODE_SET | h2 | Set write address (3 bit) |
| DEBUGGER_OPCODE_HI | h4 | Write data to high bits at write address |
| DEBUGGER_OPCODE_LO | h8 | Write data to low bits at write address |
| DEBUGGER_OPCODE_STEP | hC | Execute a single processor instruction and print hFF when done |
| DEBUGGER_OPCODE_DUMP_ROM | hE | Dumps the entire bootloader ROM to serial |
| DEBUGGER_OPCODE_DUMP_REGS | hA | Dumps all 16 registers to serial |

**Debug-Registers**

Writeable:

* 0b0: Debug enable (dbgClk enable)
* 0b1: CPU Register address overwrite enable
* 0b2: Reset request
* 0b3: Instruction overwride enable
* 1b3-0: CPU Register address overwrite value
* 4b7-0: Low-bits of instruction overwrite buffer (first word)
* 5b7-0: High-bits of instruction overwrite buffer (first word)
* 6b7-0: Low-bits of instruction overwrite buffer (second word)
* 7b7-0: High-bits of instruction overwrite buffer (second word)

Readable:

* 8b7-0: Low-bit content of currently selected CPU register
* 9b7-0: High-bit content of currently selected CPU register
* Fb0: Halted


**Known Issues**

A program *starting* with a "SET" instruction can sometimes (unconfirmed reproducability) lead to issues in debugging mode.
