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
SET data
0x4865
STOR data addr
INC addr
SET data
0x6C6C
STOR data addr
INC addr
SET data
0x6F2C
STOR data addr
INC addr
SET data
0x2077
STOR data addr
INC addr
SET data
0x6F72
STOR data addr
INC addr
SET data
0x6C64
STOR data addr
INC addr
SET data
0x210A
STOR data addr

MOV baseAddr addr ; Reset memory position
SET maxAddr ; Set max memory position
0x0017

; Output characters from memory
#declare F tmp
.loop CALL .print ; Call print function
INC addr
JMPNQ .loop_end addr maxAddr
MOV baseAddr addr ; Reset memory position at end of data
.loop_end GOTO .loop ; Loop

.print LOAD data addr ; FUNCTION print
AND data tmp left_only
BUS tmp 0x2 ; 1st char
SHFT data tmp -0x4
SHFT tmp tmp -0x4
BUS tmp 0x2 ; 2nd char
RET