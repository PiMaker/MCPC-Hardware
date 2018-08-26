// CPU States
`define CPU_STATE_INS_LOAD 8'h00
`define CPU_STATE_WAITING 8'h01
`define CPU_STATE_COMMIT 8'h02
`define CPU_STATE_PC_INC 8'h03

// Instructions
`define INS_HALT 4'h0
`define INS_MOV 4'h1
`define INS_MOVNZ 4'h2
`define INS_MOVEZ 4'h3
`define INS_BUS 4'h4
`define INS_RES1 4'h5
`define INS_SET 4'h6
`define INS_RES2 4'h7
`define INS_ALU_AND 4'h8
`define INS_ALU_OR 4'h9
`define INS_ALU_NOT 4'hA
`define INS_ALU_ADD 4'hB
`define INS_ALU_SHFT 4'hC
`define INS_ALU_MUL 4'hD
`define INS_ALU_GT 4'hE
`define INS_ALU_EQ 4'hF

// Instruction decomposition
`define INSDECOMP_INS(ins) ins[0+:4]
`define INSDECOMP_FROM(ins) ins[4+:4]
`define INSDECOMP_TO(ins) ins[8+:4]
`define INSDECOMP_IF(ins) ins[12+:4]


module cpu(
	//////////// GLOBALS ////////////
	input clk,
	input clk50,
	input rst,
	
	
	//////////// SDRAM ////////////
	output		    [12:0]		DRAM_ADDR,
	output		     [1:0]		DRAM_BA,
	output		          		DRAM_CAS_N,
	output		          		DRAM_CKE,
	output		          		DRAM_CLK,
	output		          		DRAM_CS_N,
	inout 		    [15:0]		DRAM_DQ,
	output		          		DRAM_LDQM,
	output		          		DRAM_RAS_N,
	output		          		DRAM_UDQM,
	output		          		DRAM_WE_N,
	
	//////////// MISC ////////////
	output			[15:0]		DEBUG_BUS,
	output			[9:0]		LEDR,
	input			[9:0]		SW
);


	//////////// SDRAM ////////////
	wire  [15:0]  writedata;
	wire  [15:0]  readdata;
	wire          write;
	wire          read;
	wire  [24:0]  writeaddr;
	wire  [24:0]  readaddr;

	// Initialize SDRAM controller
	Sdram_Control ram_controller_instance (
		.REF_CLK(clk50),
		.RESET_N(rst),
		// FIFO Write Side 
		.WR_DATA(writedata),
		.WR(write),
		.WR_ADDR(writeaddr),
		.WR_MAX_ADDR(25'h1ffffff),
		.WR_LENGTH(9'h80),
		.WR_LOAD(),
		.WR_CLK(clk),
		// FIFO Read Side 
		.RD_DATA(readdata),
		.RD(read),
		.RD_ADDR(readaddr),
		.RD_MAX_ADDR(25'h1ffffff),
		.RD_LENGTH(9'h80),
		.RD_LOAD(),
		.RD_CLK(clk),
		// SDRAM Side
		.SA(DRAM_ADDR),
		.BA(DRAM_BA),
		.CS_N(DRAM_CS_N),
		.CKE(DRAM_CKE),
		.RAS_N(DRAM_RAS_N),
		.CAS_N(DRAM_CAS_N),
		.WE_N(DRAM_WE_N),
		.DQ(DRAM_DQ),
		.DQM({DRAM_UDQM,DRAM_LDQM}),
		.SDR_CLK(DRAM_CLK)
	);
		
		
	// CORE BUSSES
	wire [15:0]
		reg_data_read,
		reg_data_write,
		
		reg_h_out,
		pc_out,

		bus_datain;
		
	reg [3:0]
		reg_addr_read,
		reg_addr_write,
		reg_addr_if;
		
	wire
		reg_we,
		bus_fromin,
		if_en = 0;

	reg
		pc_inc;
	
	
	// CORE COMPONENTS
	wire [3:0] reg_addr_write_wire;
	assign reg_addr_write_wire = if_en ? reg_addr_if : reg_addr_write;

	core_registers  u_core_registers (
		.clk                     ( clk            ),
		.rst                     ( rst            ),
		.addr_read               ( reg_addr_read      ),
		.addr_write              ( reg_addr_write_wire     ),
		.data_write              ( reg_data_read     ),
		.write_enable            ( reg_we   ),
		.bus_datain              ( bus_datain     ),
		.bus_fromin              ( bus_fromin     ),
		.pc_inc                  ( pc_inc         ),

		.data_read               ( reg_data_read      ),
		.pc_out                  ( pc_out         ),
		.reg_h_out               ( reg_h_out      )
	);


	// CORE LOGIC
	reg [7:0] cpu_state;
	reg [15:0] instruction_buffer;
	reg [15:0] bootloader_rom [0:2048];

	reg [15:0] continue_execution_register;
	reg [15:0] write_enable_register;


	wire continue_execution;
	assign continue_execution = |continue_execution_register;

	wire reg_we_requested;
	reg reg_we_approved;
	assign reg_we_requested = |write_enable_register;
	assign reg_we = reg_we_requested && reg_we_approved;


	initial begin
		// Dummy program
		bootloader_rom[0] <= 16'h07E1;
		bootloader_rom[1] <= 16'h05D1;
		bootloader_rom[2] <= 16'h0751;
		bootloader_rom[3] <= 16'h0BC1;
	end


	always @(posedge clk) begin
	  
		if (rst) begin

			cpu_state <= `CPU_STATE_INS_LOAD;
			continue_execution_register <= 16'h0;
			write_enable_register <= 16'h0;
			reg_we_approved <= 0;
			halted <= 0;
			instruction_buffer <= 16'h0;
			pc_inc <= 0;

		end else begin

			// Main CPU logic
			case (cpu_state)

				`CPU_STATE_INS_LOAD: begin
					cpu_state <= `CPU_STATE_WAITING;
					instruction_buffer <= bootloader_rom[pc_out];
					end

				`CPU_STATE_WAITING: begin
					if (continue_execution) begin
						cpu_state <= `CPU_STATE_COMMIT;
						continue_execution_register <= 16'h0;
					end else begin

						// Decompose instruction into different register busses
						reg_addr_write <= `INSDECOMP_TO(instruction_buffer);
						reg_addr_read <= `INSDECOMP_FROM(instruction_buffer);
						reg_addr_if <= `INSDECOMP_IF(instruction_buffer);

						// Instruction decoding
						case (`INSDECOMP_INS(instruction_buffer))

							`INS_HALT: task_halt();
							`INS_MOV: task_mov();
							default: task_halt();

						endcase
					end
					end

				`CPU_STATE_COMMIT: begin
					reg_we_approved <= 1'b1;
					pc_inc <= 1'b1;
					cpu_state <= `CPU_STATE_PC_INC;
					end
				
				`CPU_STATE_PC_INC: begin
					pc_inc <= 1'b0;
					reg_we_approved <= 1'b0;
					write_enable_register <= 16'h0;
					cpu_state <= `CPU_STATE_INS_LOAD;
					end

				default: cpu_state <= `CPU_STATE_INS_LOAD;

			endcase

		end

	end


	// INSTRUCTION LOGIC
	reg halted;
	task task_halt;
	begin
		// do nothing, because halt
		halted <= 1'b1;
	end
	endtask

	task task_mov;
	begin
		// move data
		write_enable_register[`INS_MOV] <= 1'b1;
		continue_execution_register[`INS_MOV] <= 1'b1;
	end
	endtask


	// DEBUG OUTPUT LOGIC
	assign DEBUG_BUS = SW[9] ? reg_h_out : (SW[8] ? pc_out : (SW[7] ? instruction_buffer : 16'hDEAD));

	assign LEDR[0] = halted;
	assign LEDR[1] = reg_we;
	assign LEDR[2] = continue_execution;
	assign LEDR[6] = cpu_state == `CPU_STATE_INS_LOAD;
	assign LEDR[7] = cpu_state == `CPU_STATE_WAITING;
	assign LEDR[8] = cpu_state == `CPU_STATE_COMMIT;
	assign LEDR[9] = cpu_state == `CPU_STATE_PC_INC;

endmodule
