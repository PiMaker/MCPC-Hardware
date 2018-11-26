module hex_bus_display(
	bus,
	hex_port
);

input [3:0] bus;
output [7:0] hex_port;

wire [6:0] HEXd;
assign hex_port = {1'b1,~HEXd};

hexto7segment hexto7segment_instance (
	.x(bus),
	.z(HEXd)
);

endmodule
