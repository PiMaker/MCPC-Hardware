// Main module ("entry point")

module main (
	//////////// CLOCK //////////
	input 		          		ADC_CLK_10,
	input 		          		MAX10_CLK1_50,
	input 		          		MAX10_CLK2_50,

	//////////// SDRAM //////////
	output		    [12:0]		DRAM_ADDR,
	output		     [1:0]		DRAM_BA,
	output		          		DRAM_CAS_N,
	output		          		DRAM_CKE,
	output		          		DRAM_CLK,
	output		          		DRAM_CS_N,
	inout 		    [15:0]		DRAM_DQ,
	output		          		DRAM_LDQM,
	output		          		DRAM_RAS_N,
	output		          		DRAM_UDQM,
	output		          		DRAM_WE_N,

	//////////// SEG7 //////////
	output		     [7:0]		HEX0,
	output		     [7:0]		HEX1,
	output		     [7:0]		HEX2,
	output		     [7:0]		HEX3,
	output		     [7:0]		HEX4,
	output		     [7:0]		HEX5,

	//////////// KEY //////////
	input 		     [1:0]		KEY,

	//////////// LED //////////
	output		     [9:0]		LEDR,

	//////////// SW //////////
	input 		     [9:0]		SW,

	//////////// VGA //////////
	output		     [3:0]		VGA_B,
	output		     [3:0]		VGA_G,
	output		          		VGA_HS,
	output		     [3:0]		VGA_R,
	output		          		VGA_VS,

	//////////// Accelerometer //////////
	output		          		GSENSOR_CS_N,
	input 		     [2:1]		GSENSOR_INT,
	output		          		GSENSOR_SCLK,
	inout 		          		GSENSOR_SDI,
	inout 		          		GSENSOR_SDO,

	//////////// Arduino //////////
	inout 		    [15:0]		ARDUINO_IO,
	inout			[35:0]		GPIO,
	inout						CLK_I2C_SCL,
	inout						CLK_I2C_SDA,
	inout 		          		ARDUINO_RESET_N
);


// Reset on Power up
reg [32:0] PWR_RST = 0;
always @(posedge MAX10_CLK1_50) begin
	if (rstReq) begin
		PWR_RST <= 0;
	end else if (!PWR_RST[25]) begin
		PWR_RST <= PWR_RST + 1'b1;
	end
end


// "Global" nets
wire RST;
assign RST = (RST & core_clock) | (~PWR_RST[25]) | (~KEY[0]);


// HEX display
hex_bus_display hex_bus_display_instance_0 (
	.bus(debug_bus[0 +: 4]),
	.hex_port(HEX0),
	.comma(1'h1)
);

hex_bus_display hex_bus_display_instance_1 (
	.bus(debug_bus[4 +: 4]),
	.hex_port(HEX1),
	.comma(1'h1)
);

hex_bus_display hex_bus_display_instance_2 (
	.bus(debug_bus[8 +: 4]),
	.hex_port(HEX2),
	.comma(1'h1)
);

hex_bus_display hex_bus_display_instance_3 (
	.bus(debug_bus[12 +: 4]),
	.hex_port(HEX3),
	.comma(1'h1)
);

hex_bus_display hex_bus_display_instance_4 (
	.bus(irq_count[0 +: 4]),
	.hex_port(HEX4),
	.comma(1'h0)
);

hex_bus_display hex_bus_display_instance_5 (
	.bus(irq_count[4 +: 4]),
	.hex_port(HEX5),
	.comma(1'h1)
);



// FB vars
wire [7:0] fb_data_bus;
wire [11:0] fb_addr_bus;
wire fb_we;
wire [11:0] framebuffer_addr_rd;
wire [7:0] framebuffer_data_rd;
wire framebuffer_rd_en;


// VGA controller
vga_controller vga_controller_instance (
	.clk50(MAX10_CLK1_50),
	.rst(RST),
	.hsync_out(VGA_HS),
	.vsync_out(VGA_VS),
	.red_out(VGA_R),
	.blue_out(VGA_B),
	.green_out(VGA_G),
	.fb_data(fb_data_bus),
	.fb_addr(fb_addr_bus),
	.fb_we(fb_we),
	.framebuffer_addr_rd(framebuffer_addr_rd),
	.framebuffer_data_rd(framebuffer_data_rd),
	.framebuffer_rd_en(framebuffer_rd_en)
);


// Clocking
reg manual_clock;
always @(*) begin
	manual_clock <= !KEY[1];
end

wire [31:0] debug_clock;
counter counter_instance_debug_clock (
	.clk(MAX10_CLK1_50),
	.reset(),
	.enable(1'b1),
	.out(debug_clock)
);

wire core_clock_fast = debug_clock[1];
wire core_clock;
wire clkbreak_act = clkbreak && !RST;

always @(posedge core_clock_fast) begin
	if (SW[0] && !clkbreak_act) begin
		core_clock <= ~core_clock;
	end else if (debug_clock[1]) begin
		core_clock <= manual_clock;
	end
end


// Keyboard
wire [7:0] ps2_data;
wire ps2_data_en;

PS2_Controller ps2_instance (
	// Inputs
	.CLOCK_50(MAX10_CLK1_50),
	.reset(|PWR_RST[23:0]),

	.the_command(16'hFF),
	.send_command(RST), // Reset triggers keyboard reset instead

	// Bidirectionals
	.PS2_CLK(ARDUINO_IO[6]),					// PS2 Clock
 	.PS2_DAT(ARDUINO_IO[7]),					// PS2 Data

	// Outputs
	.command_was_sent(a),
	.error_communication_timed_out(b),

	.received_data(ps2_data),
	.received_data_en(ps2_data_en)			// If 1 - new data has been received
);

wire a, b;
//assign debug_bus = {3'h0, ps2_data_en, 2'h0, b, a, ps2_data};


// Main CPU instance
wire [15:0] debug_bus;
wire [7:0] irq_count;
wire debugEn;
wire rstReq;
wire clkbreak;

cpu cpu_instance (
	.clkCore(core_clock),
	.clk50(MAX10_CLK1_50),
	.rst(RST),
	
	.DRAM_ADDR(DRAM_ADDR),
	.DRAM_BA(DRAM_BA),
	.DRAM_CAS_N(DRAM_CAS_N),
	.DRAM_CKE(DRAM_CKE),
	.DRAM_CLK(DRAM_CLK),
	.DRAM_CS_N(DRAM_CS_N),
	.DRAM_DQ(DRAM_DQ),
	.DRAM_LDQM(DRAM_LDQM),
	.DRAM_RAS_N(DRAM_RAS_N),
	.DRAM_UDQM(DRAM_UDQM),
	.DRAM_WE_N(DRAM_WE_N),

	.ps2_data(ps2_data),
	.ps2_data_en(ps2_data_en),
	
	.DEBUG_BUS(debug_bus),
	.LEDR(LEDR),
	.SW(SW),
	.irq_count(irq_count),

    .fb_data(fb_data_bus),
    .fb_addr(fb_addr_bus),
    .fb_we(fb_we),
	.framebuffer_rd_en(framebuffer_rd_en),
	.framebuffer_data_rd(framebuffer_data_rd),
	.framebuffer_addr_rd(framebuffer_addr_rd),

	.debugEnOut(debugEn),
	.rstReq(rstReq),
	.clkbreak(clkbreak),

	.ARDUINO_IO(ARDUINO_IO),
	.GPIO(GPIO)
);

endmodule
