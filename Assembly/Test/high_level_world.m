
//+entrypoint
int main(int argc, char* argv)
{
    int a = 3;
    int b = 7;
    int c = 84 / (doubleNumber(a) * b);

    int ascii = c + 0x30;

    // Set up display output pointers
    setupDisplay();

    outputAscii(ascii);
    outputAscii(ascii + 1);

    outputAscii('\n');

    ascii = 65;
    for(int i = 0; i < 26; i++)
    {
        outputAscii(ascii + i);
    }

    outputAscii('\n');

    if (ascii < 100)
    {
        print("Hello world!\n");
    }

    loop1: if (ascii < 65)
    {
        outputAscii(ascii);
        goto loop1;
    }

    print("END!");

    return 0;
}

int doubleNumber(int input)
{
    return input * 2;
}

int displayPointer;
#define DISPLAY_WIDTH 80
#define DISPLAY_HEIGHT 60
int displayX = 0;
int displayY = 0;

void setupDisplay()
{
    int memoryLayoutPointer = 0x8000 /* Base pointer */ + 0x65; // 0x8065 is the start address of the ascii framebuffer
    displayPointer = memoryLayoutPointer;
}

void print(char* value)
{
    while (*value != 0)
    {
        outputAscii(*value);
        value++;
    }
}

void outputAscii(char value)
{
    if (value == '\n')
    {
        displayX = 0;
        displayY++;
    }
    else
    {
        int* writeAddress = displayPointer + displayX + (displayY * DISPLAY_WIDTH);
        *writeAddress = value;

        displayX++;
    }
}
