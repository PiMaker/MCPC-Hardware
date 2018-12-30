module core_registers(
	input clk,
	input rst,
	
	input [3:0] addr_read,
	input [3:0] addr_write,
	input [15:0] data_write,
	input write_enable,
	input [15:0] bus_datain,
	input bus_fromin,
	input pc_inc,
	input [15:0] pc_default,
	
	output [15:0] data_read,
	output [15:0] pc_out,
	output [15:0] reg_h_out
);

	reg [15:0] gp_reg [0:10];
	reg [15:0] bus_reg;
	reg [15:0] pc_reg;
	
	assign reg_h_out = gp_reg[7];
	assign pc_out = pc_reg;

	integer i;

	// Read logic
	assign data_read = 	addr_read < 11 ? gp_reg[addr_read] :
						addr_read == 11 ? pc_reg :
						addr_read == 12 ? 1'b0 :
						addr_read == 13 ? 1'b1 :
						addr_read == 14 ? 16'hFFFF :
						bus_reg;
						
	 // Write logic
	 always @(posedge clk) begin
	 	
		if (rst) begin
			for (i = 0; i < 11; i = i + 1) begin
				gp_reg[i] <= 0;
			end
		end else if (write_enable) begin
		  	if (addr_write < 11) begin
	 			gp_reg[addr_write] <= data_write;
 			end
		end
	 	
 		
 		// PC logic
 		if (rst) begin
 			pc_reg <= pc_default;
 		end else begin
	 		if (write_enable && (addr_write == 11)) begin
	 			pc_reg <= data_write;
	 		end else if (pc_inc) begin
	 			pc_reg <= pc_reg + 1'b1;
	 		end
 		end
	 	
	 end
	 
	 // Bus logic
	 /*always @(posedge clk) begin
	 	
	 	if (bus_fromin) begin
	 		bus_reg <= bus_datain;
	 	end 
	 	
	 end*/
	 
							
endmodule
