; Hand-crafted ASM (for interrupts and performance optimized routines)


; Memory addresses:
; p1.0x100: rd ptr
; p1.0x101: wr ptr
; p1.0x0-0xFF: input fifo


; Interrupt handler
.irq_handler __LABEL_SET

; Get irq type and preload IRQ payload
loadla A 0x9010
loadla H 0x9011

; Keyboard
setreg B 0xA
jmpeq .irq_keycode_known A B

; Unknown keycode (VM only)
setreg B 0xB
jmpeq .irq_keycode_unknown A B

; Exit IRQ
.irq_return STORLA 0 0x9002


; Keyboard key pressed, move into fifo queue
.irq_keycode_known __LABEL_SET
mov H A
jmp .irq_append_to_fifo
.irq_keycode_unknown  __LABEL_SET
setreg A 0x3f ; '?'
jmp .irq_append_to_fifo


; FIFO controller
; regA contains keycode
.irq_append_to_fifo  __LABEL_SET

setpage 0x1 ; set page for rest of function
loadla B 0x100 ; rd
loadla C 0x101 ; wr

; Check for overflow and ignore IRQ if necessary
setreg G 0x00FF
jmpeq .irq_return C G

; Reset rd and wr if possible
jmpnq .irq_skip_reset B C
jmpez .irq_skip_reset B 
storla 0 0x100 ; rd is invalid from here on out
mov 0 C

.irq_skip_reset __LABEL_SET
; Write to FIFO and increment wr ptr
stor A C
inc C

; Write back wr ptr
storla C 0x101

; FIFO controller end
jmp .irq_return


.mscr_code_end HALT
; Manually insert end label, since we sed' that out in the Makefile
; (to avoid MSCR-compiled code to override our meticulously hand-crafted code in RAM)
