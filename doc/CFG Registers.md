CFG register layout
===================

General/Memory control
----------------------

* [h8000] r ... MCPC version number
* [h8001] rw(k) ... RAM instruction loading (RIL; loads instruction from RAM instead of EEPROM)
* [h8002] rw(k) ... RIL PC swap (program counter value to load on enabling/disabling RIL)
* [h8003] rw(k) ... Protected memory page (Instructions executed as RIL when execute-select page matches this are executed with kernel priviledges, defaults to 0)
* [h8004] rw(k) ... Kernel memory bits (see [h8800] below)
* [h8800] rw ... Memory high bits (RAM uses 25 bit addressing, lower 15 bit are set by MEMR/MEMW directly, the next 5 bits are loaded from the lower bits of this CFG register, the missing 5 bits are the lower kernel memory bits)

Bootloader ROM
--------------

* [hD000-hD800] r ... Bootloader ROM read-only access (up to h800 relative, e.g. 2048 words)

VGA subsystem
-------------

* [hDFFD] r ... Width of display (maximum x, exclusive, 98 on hardware, 120 max)
* [hDFFE] r ... Height of display (maximum y, exclusive, 35 on hardware, 65 max)
* [hDFFF] r ... Last valid address of ASCII buffer
* [hE000-{hDFFF}] rw ... ASCII buffer (one ascii characted per data word, high bits unused)

Interrupt/Timer subsystem
-------------------------

* [h9000] rw(k) ... Interrupt handler address
* [h9001] rw(k) ... Interrupt enable
* [h9002] rw(k) ... Interrupt context state (can only write 0 to exit IRQ, other writes are ignored)
* [h9010-h9011] r(k) ... Current interrupt data (reads 0 if [h9002] reads 0)

Virtualization subsystem
------------------------

TBD

Debugging
---------

* [hFFFF] w ... Break Execution (enable manual clock)
