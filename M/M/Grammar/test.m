#include stdlib.m

/*
  printc prints one single ascii character
  to the terminal output
*/
void printc(char c) {
    c << 8;
	asm
	{
		BUS __reg(c) 0x2
	}
}

// Program entry point
void main() {
    int x = 5;
	int y = 20;
	int z = 3*y+x;
	printc(char(z));
}