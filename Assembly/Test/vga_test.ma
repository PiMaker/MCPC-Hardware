
; Test VGA subsystem by writing alphabet on 5 lines

#declare A line
#declare B column
#declare C asciiChar
#declare D alphabetMax
#declare E lineMax
#declare G vgaBase

.main SET SP
0x7fff
SET vgaBase
0xE000
SET line
0x0000
SET alphabetMax
0x005B
SET lineMax
0x0005

.mainLoop CALL .alphabet
INC line
JMPNQ .mainLoop line lineMax
HALT

.putchar SET F
0x62 ; 98 character wide, in hex
MUL F F line
ADD F F column
ADD F F vgaBase
STOR asciiChar F
INC column
RET

.alphabet SET column
0x0000
SET asciiChar
0x0041
.alphabetLoop CALL .putchar
INC asciiChar
JMPEQ .alphabetEnd asciiChar alphabetMax
JMP .alphabetLoop

.alphabetEnd RET

