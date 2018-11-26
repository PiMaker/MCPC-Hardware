#assign col 			= (x_pos)&8'h07;
#assign row			= (y_pos)&8'h0F;
#assign framebuffer_read_addr = (x_pos>>3) - 1 + (((y_pos>>4) + 15)<<6);

for y_pos in range(16, 550-32):
    for x_pos in range(8,524-8):
        col = x_pos>>3
        row = y_pos>>4
        fra = (x_pos>>3) - 1 + (((y_pos>>4) - 1)<<6)
        if (x_pos&0x07) == 1 and (y_pos&0x0F) == 1:
            print("col: %d \trow: %d \t fra: %d" % (col, row, fra))

print("Done!")
