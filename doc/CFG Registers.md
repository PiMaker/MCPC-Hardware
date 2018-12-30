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
* [0xD000-0xD800] r ... Bootloader ROM read-only access

= VGA subsystem =
* [0xE000-0xF2BF] w ... ASCII buffer (98x60 display, one ascii characted per data word, high bits unused)

= Interrupt/Timer subsystem =
* [0x9000] rw(k) ... Interrupt handler address
* [0x9001] rw(k) ... Interrupt enable
* [0x9002] rw(k) ... Interrupt context state
* [0x9010-0x9011] r(k) ... Current interrupt data (reads 0 if [0x9002] reads 0)

= Virtualization subsystem =
TBD