module main (
	input MAX10_CLK1_50,
	input [1:0] KEY,
	output [9:0] LEDR
);

	wire rkey;
	assign rkey = ! (KEY[0] & KEY[1]);

	counter counter_instance (
		.CLOCK(MAX10_CLK1_50),
		.RST(rkey),
		.LED(LEDR[0])
	);

endmodule