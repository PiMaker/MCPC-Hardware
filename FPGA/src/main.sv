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


// Blank HEX display
assign HEX5 = 8'hFF;
assign HEX4 = 8'h7F;

hex_bus_display hex_bus_display_instance_0 (
	.bus(debug_bus[0 +: 4]),
	.hex_port(HEX0)
);

hex_bus_display hex_bus_display_instance_1 (
	.bus(debug_bus[4 +: 4]),
	.hex_port(HEX1)
);

hex_bus_display hex_bus_display_instance_2 (
	.bus(debug_bus[8 +: 4]),
	.hex_port(HEX2)
);

hex_bus_display hex_bus_display_instance_3 (
	.bus(debug_bus[12 +: 4]),
	.hex_port(HEX3)
);



// FB vars
wire [7:0] fb_data_bus;
wire [11:0] fb_addr_bus;
wire fb_we;


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
	.fb_we(fb_we)
);


// Manual clock for testing/debugging
reg manual_clock;
always @(*) begin
	manual_clock <= !KEY[1];
end

wire [31:0] debug_clock;
counter counter_instance_debug_clock (
	.clk(MAX10_CLK1_50),
	.reset(rst),
	.enable(1'b1),
	.out(debug_clock)
);

wire core_clock = SW[0] ? debug_clock[2] : manual_clock;


// Main CPU instance
wire [15:0] debug_bus;
wire debugEn;
wire rstReq;

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
	
	.DEBUG_BUS(debug_bus),
	.LEDR(LEDR),
	.SW(SW),

    .fb_data(fb_data_bus),
    .fb_addr(fb_addr_bus),
    .fb_we(fb_we),

	.debugEnOut(debugEn),
	.rstReq(rstReq),

	.ARDUINO_IO(ARDUINO_IO),
	.GPIO(GPIO)
);

	
endmodule
