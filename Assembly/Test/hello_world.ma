#declare A data
#declare B addr
#declare C baseAddr
#declare E maxAddr

; Setup stack
SET SP
0x01FF ; Highest SRAM location

; Set bit masks
#declare D left_only
SET left_only
0xFF00

; Set memory address
SET baseAddr
0x0010
MOV baseAddr addr

; Write "Hello, world!\n" to memory
CHAR addr 'He'
CHAR addr 'll'
CHAR addr 'o,'
CHAR addr ' w'
CHAR addr 'or'
CHAR addr 'ld'
CHAR addr '!\n'

BRK

MOV baseAddr addr ; Reset memory position
SET maxAddr ; Set max memory position
0x0017

; Output characters from memory
#declare F tmp
.loop BRK
CALL .print ; Call print function
BRK
INC addr
JMPNQ .loop_end addr maxAddr
MOV baseAddr addr ; Reset memory position at end of data
.loop_end GOTO .loop ; Loop

.print BRK
LOAD data addr ; FUNCTION print
SHFT data tmp -0x4
SHFT tmp tmp -0x4
BUS tmp 0x2 ; 1st char
AND data tmp left_only
BUS tmp 0x2 ; 2nd char
BRK
RET