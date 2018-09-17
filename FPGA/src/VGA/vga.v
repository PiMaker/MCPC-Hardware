module vsync(line_clk, vsync_out, blank_out);
	input line_clk;
	output vsync_out;
	output blank_out;

	reg [10:0] count = 10'b0000000000;
	reg vsync  = 0;
	reg blank  = 0;

	always @(posedge line_clk)
	if (count < 666)
		count <= count + 11'h1;
	else
		count <= 0;

	always @(posedge line_clk)
	if (count < 600)
		blank 		<= 0;
	else
		blank 		<= 1;

	always @(posedge line_clk)
	begin
		if (count < 637)
			vsync 	<= 1;
		else if (count >= 637 && count < 643)
			vsync 	<= 0;
		else if (count >= 643)
			vsync 	<= 1;
	end

	assign vsync_out  = vsync;
	assign blank_out  = blank;

endmodule // vsync   

module hsync(clk50, hsync_out, blank_out, newline_out);
	input clk50;
	output hsync_out, blank_out, newline_out;

	reg [10:0] count = 10'b0000000000;
	reg hsync 	= 0;
	reg blank 	= 0;
	reg newline 	= 0;

	always @(posedge clk50)
	begin
		if (count < 1040)
			count  <= count + 11'h1;
		else
			count  <= 0;
	end

	always @(posedge clk50)
	begin
		if (count == 0)
			newline <= 1;
		else
			newline <= 0;
	end

	always @(posedge clk50)
	begin
		if (count >= 800)
			blank  <= 1;
		else
			blank  <= 0;
	end

	always @(posedge clk50)
	begin
		if (count < 856) // pixel data plus front porch
			hsync <= 1;
		else if (count >= 856 && count < 976)
			hsync <= 0;
		else if (count >= 976)
			hsync <= 1;
	end // always @ (posedge clk50)
				 
   assign hsync_out    = hsync;
	assign blank_out    = blank;
	assign newline_out  = newline;

endmodule // hsync


module vga_controller(clk50, rst, hsync_out, vsync_out, red_out, blue_out, green_out, fb_data, fb_addr, fb_we);
	input clk50;
	input rst;
	input [7:0] fb_data;
	input [11:0] fb_addr;
	input wire fb_we;
	output hsync_out, vsync_out;
	output [3:0] red_out, blue_out, green_out;
	wire line_clk, blank, hblank, vblank;

	hsync   hs(clk50, hsync_out, hblank, line_clk);
	vsync   vs(line_clk, vsync_out, vblank);
	framebuffer   fb(clk50, rst, blank, red_out, green_out, blue_out, line_clk, vblank, fb_data, fb_addr, fb_we);

	assign blank 	 = hblank || vblank;

endmodule // vga_controller













