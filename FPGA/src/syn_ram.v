module syn_ram(
    output [DATA_WIDTH-1:0] data_out,
    input [ADDR_WIDTH-1:0] addr_in,
    input [ADDR_WIDTH-1:0] addr_out,
    input [DATA_WIDTH-1:0] data_in,
    input write_enable,
    input clk
);

	parameter DATA_WIDTH = 8;
	parameter DATA_COUNT = 256;
	parameter ADDR_WIDTH = 8;

    reg [DATA_WIDTH-1:0] memory [0:DATA_COUNT-1];

    always @(posedge clk) begin
        if (write_enable) begin
            memory[addr_in] <= data_in;
        end
    end

    assign data_out = memory[addr_out];

endmodule
