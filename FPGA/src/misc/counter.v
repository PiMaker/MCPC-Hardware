module counter (
    out,
    enable,
    clk,
    reset
);

output [31:0] out;
input enable, clk, reset;

reg [31:0] out;

always @(posedge clk)
    if (reset) begin
        out <= 31'b0 ;
    end else if (enable) begin
        out <= out + 1;
    end
endmodule 
