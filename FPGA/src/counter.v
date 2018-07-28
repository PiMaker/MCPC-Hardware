module counter (

	input CLOCK,
	input RST,
    output LED
    
);

    /* reg */
    reg [32:0] counter;
	
	/* assign */
	assign LED = counter[22];
    
    /* always */
    always @ (posedge CLOCK or posedge RST) begin
    	if (RST) begin
    		counter <= 0;
    	end else begin
    		counter <= counter + 1;
    	end
    end

endmodule