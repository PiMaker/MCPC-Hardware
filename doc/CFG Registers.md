CFG register layout
-------------------

= General/Memory control =
* [0x8000] r ... MCPC version number
* [0x8001] rw(k) ... RAM instruction loading (RIL; loads instruction from RAM instead of EEPROM)
* [0x8002] rw(k) ... RIL PC swap (program counter value to load on enabling/disabling RIL)
* [0x8003] rw(k) ... Protected memory page (Instructions executed as RIL when execute-select page matches this are executed with kernel priviledges, defaults to 0)
* [0x8004] rw(k) ... Kernel memory bits (see [0x8800] below)
* [0x8800] rw ... Memory high bits (RAM uses 25 bit addressing, lower 15 bit are set by MEMR/MEMW directly, the next 5 bits are loaded from the lower bits of this CFG register, the missing 5 bits are the lower kernel memory bits)

= Bootloader ROM =
* [0xD000-0xD800] r ... Bootloader ROM read-only access (up to 0x800 relative)

= VGA subsystem =
* [0xDFFD] r ... Width of display (maximum x, exclusive, 98 on hardware, 120 max)
* [0xDFFE] r ... Height of display (maximum y, exclusive, 35 on hardware, 65 max)
* [0xDFFF] r ... Last valid address of ASCII buffer
* [0xE000-{0xDFFF}] rw ... ASCII buffer (one ascii characted per data word, high bits unused)

= Interrupt/Timer subsystem =
* [0x9000] rw(k) ... Interrupt handler address
* [0x9001] rw(k) ... Interrupt enable
* [0x9002] rw(k) ... Interrupt context state (can only write 0 to exit IRQ, other writes are ignored)
* [0x9010-0x9011] r(k) ... Current interrupt data (reads 0 if [0x9002] reads 0)

= Virtualization subsystem =
TBD

= Debugging =
* [0xFFFF] w ... Break Execution (enable manual clock)