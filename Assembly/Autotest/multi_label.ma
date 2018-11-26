;autotest reg=1 val=0x2;

SET B
0x1

JMP .test
HALT

.test1 __LABEL_SET
.test2 __LABEL_SET
.test __LABEL_SET
.testus SET B
0x2
HALT