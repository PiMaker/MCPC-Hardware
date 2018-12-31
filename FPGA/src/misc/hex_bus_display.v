module hex_bus_display(
	bus,
	hex_port,
	comma
);

input [3:0] bus;
output [7:0] hex_port;
input comma;

wire [6:0] HEXd;
assign hex_port = {comma,~HEXd};

hexto7segment hexto7segment_instance (
	.x(bus),
	.z(HEXd)
);

endmodule
