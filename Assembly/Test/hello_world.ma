#declare A data
#declare H addr
#declare C baseAddr
#declare E maxAddr
#declare B counter
#declare G maxCount

; Setup stack
SET SP
0x01FF ; Highest SRAM location

; Set max count
SET maxCount
0x0008

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

MOV baseAddr addr ; Reset memory position
SET maxAddr ; Set max memory position
0x0017

; Output characters from memory
#declare F tmp
.loop CALL .print ; Call print function
INC addr
JMPNQ .loop_end addr maxAddr
MOV baseAddr addr ; Reset memory position at end of data
INC counter
JMPEQ .prog_end counter maxCount
.loop_end GOTO .loop ; Loop

.print LOAD data addr ; FUNCTION print
SHFT data tmp -0x4
SHFT tmp tmp -0x4
BUS tmp 0x2 ; 1st char
AND data tmp left_only
BUS tmp 0x2 ; 2nd char
RET

.prog_end HALT