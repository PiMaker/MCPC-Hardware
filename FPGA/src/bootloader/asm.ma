; Hand-crafted ASM (for interrupts and performance optimized routines)

; Interrupt handler
.IRQ_HANDLER __LABEL_SET
;STORLA 1 0xFFFF
;LOADLA A 0x9010
;STORLA_P A 0x0 0x1
;LOADLA A 0x9011
;STORLA_P A 0x1 0x1

LOADLA H 0x9011
HALT

.IRQ_RETURN STORLA 0 0x9002

.mscr_code_end HALT
; Manually insert end label, since we sed' that out in the Makefile
; (to avoid MSCR-compiled code to override our code in RAM)