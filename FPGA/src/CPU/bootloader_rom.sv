module bootloader_rom (
    input [(addr_width-1):0] addr1, addr2,
    input clk, 
    output reg [(data_width-1):0] q1, q2
);
    parameter data_width = 16;
    parameter addr_width = 15;

    reg [data_width-1:0] rom[2**addr_width-1:0];

    initial begin
        $readmemh("bootloader.hex", rom);
    end

    always @ (posedge clk) begin
        q1 <= rom[addr1];
        q2 <= rom[addr2];
    end
endmodule