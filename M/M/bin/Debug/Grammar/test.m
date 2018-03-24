#include stdlib.m

inline void printc(char c) {
	__reg("A", c);
	asm
	{
		PUSH A
		BUS A 0x2
		POP A
	}
}

void main() {
    int x = 5;
	int y = 20;
	int z = 3*y+x;
	printc((char)z);
}