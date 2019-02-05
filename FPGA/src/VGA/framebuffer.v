module framebuffer(clk, rst, blank, red_out, green_out, blue_out, line_clk, vblank, framebuffer_data, framebuffer_addr, framebuffer_write_enable, framebuffer_addr_rd, framebuffer_data_rd, framebuffer_rd_en);
	input clk, rst, blank, vblank, line_clk;
	output wire [3:0] green_out, blue_out;
    output reg [3:0] red_out;

	input [7:0] framebuffer_data;
	output wire [7:0] framebuffer_data_rd;
	input [11:0] framebuffer_addr, framebuffer_addr_rd;
	input wire framebuffer_write_enable, framebuffer_rd_en;
	
	// Framebuffer
	wire [11:0] framebuffer_read_addr;
    wire [7:0] fb_data_read;

	// Dual-port RAM implementation
	// Bus A: Reading for VGA
	// Bus B: Writing, reading for rd output
    fb_ram_dual	fb_ram_inst (
		.address_a(framebuffer_read_addr),
		.address_b(framebuffer_rd_en ? framebuffer_addr_rd : framebuffer_addr),
		.clock(clk),
		.data_a(8'h0),
		.data_b(framebuffer_data),
		.wren_a(1'h0),
		.wren_b(framebuffer_write_enable),
		.q_a(fb_data_read),
		.q_b(framebuffer_data_rd)
	);
	
	// ASCII logic
	wire [2:0] col;
	wire [3:0] row;
	wire pixel;
	wire pixel_en;
	
	// Read character font data
	reg chars[0:16383];
	initial begin
	    $readmemb("font_data.raw", chars);
	end
	

	reg [9:0] x_pos;
	reg [9:0] y_pos;
    reg [32:0] char_addr;

	always @(posedge clk)
	begin
		if (blank)
			x_pos <= 0;
		else begin
			x_pos = x_pos + 1;
            char_addr = {fb_data_read,col,row};
            red_out <= pixel_en ? {4{chars[char_addr]}} : 4'h0;
		end
	end

	always @(posedge line_clk)
	begin
		if (vblank) begin
			y_pos <= 0;
		end else begin
			y_pos <= y_pos + 1;
		end
	end


	// ASCII to pixel
	assign col 			= (x_pos)&8'h07;
	assign row			= (y_pos)&8'h0F;
	assign framebuffer_read_addr = ((x_pos + 1)>>3) - 1 + ((((y_pos-4)>>4) - 1) * 98);

	assign pixel_en		= (x_pos >= 8 && x_pos < (800-8) && y_pos >= 16 && y_pos < (592-16));
	assign green_out	= red_out;
	assign blue_out		= red_out;

endmodule // framebuffer
