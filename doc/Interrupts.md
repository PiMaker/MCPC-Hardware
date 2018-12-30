Interrupts
----------

Interrupts are added to a FIFO queue as two words:
[0] ... Interrupt number/code
[1] ... Data (arbitrary, useful in combination with [0])

To access these values from the interrupt handler, perform a read to the following CFG registers:
0x9010 for [0]
0x9011 for [1]

In CPU_STATE_PC_INC, the queue is checked, and if there is a valid entry it is removed and the interrupt handler (according to CFG[0x9000]) is triggered.
If CFG[0x9001] (interrupt enable) is equal to 16'h0, the entry is removed, but the handler is not executed. The interrupt is thus discarded.
If the FIFO queue is full (256 entries), further interrupts will be discarded.
To exit the interrupt handler, write 16'h0 to CFG[0x9002] (interrupt context). CFG[0x9002] will read 16'hFFFF iff the accessing MEMR command is executed in an interrupt handler context, 16'h0 otherwise.
To avoid breaking state for the interrupt program, all registers are temporarily stored in a backup environment.
The interrupt contexts registers will be initialized to 0 (except 0, 1, -1 and PC).
TODO: Interrupts are always executed in kernel context.