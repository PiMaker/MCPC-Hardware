module framebuffer(clk, rst, blank, red_out, green_out, blue_out, line_clk, vblank, framebuffer_data, framebuffer_addr, framebuffer_write_enable);
	input clk, rst, blank, vblank, line_clk;
	output [3:0] red_out, green_out, blue_out;

	input [7:0] framebuffer_data;
	input [11:0] framebuffer_addr;
	input wire framebuffer_write_enable;
	
	// Framebuffer
	reg [7:0] fb_mem [0:4096];
	wire [11:0] framebuffer_read_addr;
	
	reg [12:0] i;
	
	always @(posedge clk)
	begin
		if (rst) begin
			fb_mem[0] = 8'h5F;
			for (i = 1; i < 4096; i = i + 1)
			begin
				fb_mem[i] <= 8'h0;
			end
		end else if (framebuffer_write_enable)
		begin
			fb_mem[framebuffer_addr] <= framebuffer_data;
		end
	end

	// ASCII logic
	wire [2:0] col;
	wire [3:0] row;
	wire pixel;
	wire pixel_en;
	
	// Read character font data
	reg chars[0:16384];
	initial begin
	    $readmemb("font_data.raw", chars, 0, 16384);
	end
	

	reg [9:0] x_pos;
	reg [9:0] y_pos;

	always @(posedge clk)
	begin
		if (blank)
			x_pos <= 0;
		else begin
			x_pos <= x_pos + 1;
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
	assign col 			= (x_pos)%8;
	assign row			= (y_pos)%16;
	assign framebuffer_read_addr = (x_pos/8) - 1 + ((y_pos/16) - 1) * 98;
	
	assign pixel_en		= (x_pos >= 8 && x_pos < 800-8 && y_pos >= 16 && y_pos < 600-24);
	assign red_out 		= pixel_en ? {4{chars[{fb_mem[framebuffer_read_addr],col,row}]}} : 4'h0;
	assign green_out	= red_out;
	assign blue_out		= red_out;

endmodule // framebuffer