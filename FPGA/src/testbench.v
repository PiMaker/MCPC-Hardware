module testbench;
	
	reg clk;
	wire led;
	reg rst;
	
	main main_instance (
		.CLOCK_50(clk),
		.RESET(rst),
		.LED(led)
	);
	
	initial
	begin
		rst = 1;
		#1 rst = 0;
		clk = 0;
		$display("== Starting simulation! ==");
		$monitor("clk=%d,led=%d",clk,led);
		#100 $finish();
	end
	
	always
	begin
		#2 clk = !clk;
	end
 
endmodule