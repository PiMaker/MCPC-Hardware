module RegisterDataArbitrator (
    input clk,
    input rst,

    input [16:0] inputs [0:15],
    input [15:0] default_input,
    output [15:0] out
);

integer i;
always @(posedge clk) begin
    out <= default_input;

    for (i = 0; i < 16; i = i + 1) begin
        if (inputs[i][0]) begin
            // Latest "<=" wins, or so I've been told
            out <= inputs[i][1 +: 16];
        end
    end
end

endmodule // RegisterDataArbitrator