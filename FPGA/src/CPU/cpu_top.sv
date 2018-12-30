// CPU States
`define CPU_STATE_INS_LOAD 8'h00
`define CPU_STATE_WAITING 8'h01
`define CPU_STATE_COMMIT 8'h02
`define CPU_STATE_PC_INC 8'h03
`define CPU_STATE_IRQ_ENTER 8'h04

// Instructions
`define INS_HALT 4'h0 // done
`define INS_MOV 4'h1 // done
`define INS_MOVNZ 4'h2 // done
`define INS_MOVEZ 4'h3 // done
`define INS_BUS 4'h4 // not implemented
`define INS_MEMR 4'h5 // done
`define INS_SET 4'h6 // done
`define INS_MEMW 4'h7 // done
`define INS_ALU_AND 4'h8 // done
`define INS_ALU_OR 4'h9 // done
`define INS_ALU_XOR 4'hA // done
`define INS_ALU_ADD 4'hB // done
`define INS_ALU_SHFT 4'hC // done
`define INS_ALU_MUL 4'hD // done
`define INS_ALU_GT 4'hE // done
`define INS_ALU_EQ 4'hF // done

// Instruction decomposition
`define INSDECOMP_INS(ins) ins[0+:4]
`define INSDECOMP_FROM(ins) ins[4+:4]
`define INSDECOMP_TO(ins) ins[8+:4]
`define INSDECOMP_IF(ins) ins[12+:4]

// General defines
`define BOOTLOADER_ROM_SIZE 2047


module cpu(
	//////////// GLOBALS ////////////
	input clkCore,
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
	

    //////////// VGA ////////////
    output reg [7:0] fb_data,
    output reg [11:0] fb_addr,
    output reg fb_we,

	//////////// PS2 ////////////
	input [7:0] ps2_data,
	input ps2_data_en,

	//////////// MISC ////////////
	output			[15:0]		DEBUG_BUS,
	output			[9:0]		LEDR,
	input			[9:0]		SW,
	output						debugEnOut,
	output						rstReq,

	//////////// GPIO ////////////
	inout 		    [15:0]		ARDUINO_IO,
	inout			[35:0]		GPIO
);


    // SDRAM new
    wire ram_clk;
    assign DRAM_CLK = ram_clk;
    sdram_pll0 ram_pll_instance (
        .areset(rst),
        .inclk0(clk50),
        .c0(ram_clk)
    );

    sdram_controller ram_controller_instance (
        .wr_addr(mem_writeaddr),
        .wr_data(mem_writedata),
        .wr_enable(mem_write),

        .rd_addr(mem_readaddr),
        .rd_data(mem_readdata),
        .rd_enable(mem_read),
        .rd_ready(mem_read_ready),

        .busy(mem_busy),

        .rst_n(!rst),
        .clk(ram_clk),

        .addr(DRAM_ADDR),
        .bank_addr(DRAM_BA),
        .data(DRAM_DQ),
        .clock_enable(DRAM_CKE),
        .cs_n(DRAM_CS_N),
        .ras_n(DRAM_RAS_N),
        .cas_n(DRAM_CAS_N),
        .we_n(DRAM_WE_N),
        .data_mask_low(DRAM_LDQM),
        .data_mask_high(DRAM_UDQM)
    );

	// SDRAM Vars
	wire  [15:0] mem_writedata;
	wire  [15:0] mem_readdata;
	reg          mem_write;
	reg          mem_read;
    reg          mem_read_ready;
	reg  [24:0]  mem_writeaddr;
	reg  [24:0]  mem_readaddr;
    reg          mem_busy;

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
		reg_addr_if,
		reg_addr_dbg;
		
	wire
		clk,

		debugEn,
		debugRegRead,
		debugInsOvr,

		reg_we,
		bus_fromin,
		reg_if_en;

	reg
		pc_inc,
		pc_dbg_virt;

	wire [31:0]
		debugInsOvrData;


	// DEBUGGER
	assign debugEnOut = debugEn;
    wire [15:0] dbgDbg, dbgRomAddr;
	cpu_debugger debugger (
		.clk(clkCore),
		.clk50(clk50),
		.rst(rst),

		.debugEn(debugEn),
		.rstReq(rstReq),

		.cpuClk(clk),
		.cpuState(cpu_state),
		.halted(halted),

		.instructionOverride(debugInsOvrData),
		.instructionOverrideEn(debugInsOvr),

		.rom_addr(dbgRomAddr),
		.rom_data(bootloader_rom[dbgRomAddr]),

		.regRead(reg_data_read),
		.regAddrOvr(reg_addr_dbg),
		.regAddrOvrEn(debugRegRead),

		.ARDUINO_IO(ARDUINO_IO),
        .dbgDbg(dbgDbg)
	);


	// INTERRUPTS
	reg irq_en = 1'h0;
	reg irq_rd_en;
	wire irq_full, irq_empty;
	wire [31:0] irq_dout;
	irq_fifo u_irq_fifo(
		.clock(clk50),
		.reset(rst),
		.wr(ps2_data_en),
		.rd(ps2_rd_en),
		.din({8'h0,ps2_data,16'h000A}),
		.empty(irq_empty),
		.full(irq_empty),
		.dout(irq_dout)
    );

	
	// CORE COMPONENTS
	wire [3:0] reg_addr_read_wire;
	assign reg_addr_read_wire = debugRegRead ? reg_addr_dbg : (reg_if_en ? reg_addr_if : reg_addr_read);

	wire pc_inc_real;
	assign pc_inc_real = debugInsOvr ? 1'b0 : pc_inc;

	reg in_irq_context = 1'h0;
	reg in_irq_context_prev = 1'h0;

	reg [15:0] irq_handler_addr;

	core_registers  u_core_registers (
		.clk                     ( clk            ),
		.rst                     ( rst            ),
		.addr_read               ( reg_addr_read_wire      ),
		.addr_write              ( reg_addr_write     ),
		.data_write              ( reg_data_write     ),
		.write_enable            ( reg_we_irq_off   ),
		.bus_datain              ( bus_datain     ),
		.bus_fromin              ( bus_fromin     ),
		.pc_inc                  ( pc_inc_real_irq_off    ),
		.pc_default              ( 16'h0 ),

		.data_read               ( reg_data_read_irq_off  ),
		.pc_out                  ( pc_out_irq_off         ),
		.reg_h_out               ( reg_h_out      )
	);

	// In IRQ context we use seperate registers to avoid dirtying state
	core_registers  u_core_registers_irq (
		.clk                     ( clk            ),
		.rst                     ( rst_irq_regs            ),
		.addr_read               ( reg_addr_read_wire      ),
		.addr_write              ( reg_addr_write     ),
		.data_write              ( reg_data_write     ),
		.write_enable            ( reg_we_irq_on   ),
		.bus_datain              (      ),
		.bus_fromin              (      ),
		.pc_inc                  ( pc_inc_real_irq_on    ),
		.pc_default              ( irq_handler_addr ),

		.data_read               ( reg_data_read_irq_on  ),
		.pc_out                  ( pc_out_irq_on         ),
		.reg_h_out               ()
	);

	// IRQ and regular register combining
	wire reg_we_irq_on, reg_we_irq_off;
	wire pc_inc_real_irq_on, pc_inc_real_irq_off;
	wire reg_data_read_irq_on, reg_data_read_irq_off;
	wire pc_out_irq_on, pc_out_irq_off;

	assign reg_we_irq_on = in_irq_context ? reg_we : 1'h0;
	assign reg_we_irq_off = !in_irq_context ? reg_we : 1'h0;
	assign pc_inc_real_irq_on = in_irq_context ? pc_inc_real : 1'h0;
	assign pc_inc_real_irq_off = !in_irq_context ? pc_inc_real : 1'h0;

	assign reg_data_read = in_irq_context ? reg_data_read_irq_on : reg_data_read_irq_off;
	assign pc_out = in_irq_context ? pc_out_irq_on : pc_out_irq_off;

	// Reset logic for IRQ registers
	reg rst_irq_regs = 1'h0;
	always @(posedge clk) begin
		if (in_irq_context && !in_irq_context_prev) begin
			// We entered irq context, reset irq registers
			rst_irq_regs = 1'h1;
		end else if (in_irq_context && in_irq_context_prev) begin
			rst_irq_regs = 1'h0;
		end

		in_irq_context_prev = in_irq_context;
	end

	// Register data write arbitration
	reg [16:0] reg_data_write_inputs [0:15];
	RegisterDataArbitrator  u_RegisterDataArbitrator (
		.clk                     ( clk             ),
		.rst                     ( rst             ),
		.inputs                  ( reg_data_write_inputs ),
		.default_input           ( reg_data_read ),

		.out                     ( reg_data_write  )
	);


	// CORE LOGIC
	reg [7:0] cpu_state;
	reg [15:0] instruction_buffer;
	reg [15:0] bootloader_rom [0:`BOOTLOADER_ROM_SIZE];

	reg [15:0] continue_execution_register;
	reg [15:0] write_enable_register;


	wire continue_execution;
	assign continue_execution = |continue_execution_register;

	wire reg_we_requested;
	reg reg_we_approved;
	assign reg_we_requested = |write_enable_register;
	assign reg_we = reg_we_requested && reg_we_approved && ~debugRegRead;


	initial begin
		$readmemh("bootloader.hex", bootloader_rom);
	end


	always @(posedge clk) begin
	  
		if (rst) begin

			cpu_state <= `CPU_STATE_INS_LOAD;
			continue_execution_register <= 16'h0;
			write_enable_register <= 16'h0;
			reg_we_approved <= 0;
			instruction_buffer <= 16'h0;
			pc_inc <= 0;
			pc_dbg_virt <= 0;
			mem_addr_ext_kernel <= 0;
			mem_addr_ext_user <= 0;

			reset_instruction_tasks();

		end else begin

			// Main CPU logic
			case (cpu_state)

				`CPU_STATE_INS_LOAD: begin
					cpu_state <= `CPU_STATE_WAITING;

					if (debugInsOvr) begin
						if (pc_dbg_virt) begin
							instruction_buffer = debugInsOvrData[15:0];
						end else begin
							instruction_buffer = debugInsOvrData[31:16];
						end
					end else begin
						instruction_buffer = bootloader_rom[pc_out];
					end

					// Decompose instruction into different register busses
					reg_addr_write <= `INSDECOMP_TO(instruction_buffer);
					reg_addr_read <= `INSDECOMP_FROM(instruction_buffer);
					reg_addr_if <= `INSDECOMP_IF(instruction_buffer);
					end

				`CPU_STATE_WAITING: begin
					if (continue_execution) begin
						cpu_state <= `CPU_STATE_COMMIT;
						continue_execution_register <= 16'h0;
						pc_inc <= 1'b0; // A bit messy but avoids triple increments during SET
					end else begin

						// Instruction decoding
						case (`INSDECOMP_INS(instruction_buffer))

							`INS_HALT: task_halt();
							`INS_MOV: task_mov();
							`INS_MOVNZ: task_movnz();
							`INS_MOVEZ: task_movez();
							`INS_SET: task_set();
							`INS_MEMR: task_memr();
							`INS_MEMW: task_memw();
							`INS_BUS: task_halt();
							default: task_alu(); // ALU here, regular instructions exhausted

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

					// Debug instruction override virtual PC increment
					if (debugInsOvr) begin
						pc_dbg_virt <= pc_dbg_virt + 1'b1;
					end
					
					reset_instruction_tasks();

					end

				default: cpu_state <= `CPU_STATE_INS_LOAD;

			endcase

		end

	end


	// INSTRUCTION LOGIC

	// HALT
	reg halted;
	task task_halt;
	begin
		// do nothing, because halt
		halted <= 1'b1;
	end
	endtask

	// MOV
	task task_mov;
	begin
		// move data
		write_enable_register[`INS_MOV] <= 1'b1;
		continue_execution_register[`INS_MOV] <= 1'b1;
	end
	endtask

	// MOVNZ
	reg task_movnz_state;
	task task_movnz;
	begin
		if (task_movnz_state == 0) begin
			reg_if_en <= 1'b1;
			task_movnz_state <= 1'b1;
		end else begin
			if (|reg_data_read) begin
				write_enable_register[`INS_MOVNZ] <= 1'b1;
			end

			reg_if_en <= 1'b0;
			task_movnz_state <= 1'b0;
			continue_execution_register[`INS_MOVNZ] <= 1'b1;
		end
	end
	endtask

	// MOVEZ
	reg task_movez_state;
	task task_movez;
	begin
		if (task_movez_state == 0) begin
			reg_if_en <= 1'b1;
			task_movez_state <= 1'b1;
		end else begin
			if (&(~reg_data_read)) begin
				write_enable_register[`INS_MOVEZ] <= 1'b1;
			end

			reg_if_en <= 1'b0;
			task_movez_state <= 1'b0;
			continue_execution_register[`INS_MOVEZ] <= 1'b1;
		end
	end
	endtask

	// SET
	reg task_set_state;
	reg [15:0] task_set_buffer;
	task task_set;
	begin
		case (task_set_state)
			1'h0: begin
				task_set_state <= 1'h1;

				if (debugInsOvr) begin
					if (~pc_dbg_virt) begin
						task_set_buffer <= debugInsOvrData[15:0];
					end else begin
						task_set_buffer <= debugInsOvrData[31:16];
					end
				end else begin
					task_set_buffer <= bootloader_rom[pc_out + 1];
				end
			end
			1'h1: begin
				task_set_state <= 1'h0;

				reg_data_write_inputs[`INS_SET] <= {task_set_buffer, 1'b1};
				write_enable_register[`INS_SET] <= 1'b1;

				pc_inc <= 1; // Skip data value

				// Debug instruction override virtual PC increment (for data skipping)
				if (debugInsOvr) begin
					pc_dbg_virt += 1'b1;
				end

				continue_execution_register[`INS_SET] <= 1'b1;
			end
		endcase
	end
	endtask

	// ALU
	reg task_alu_state;
	reg [15:0] task_alu_input_buffer;
	reg [15:0] task_alu_temp_buffer;
	task task_alu;
	begin
		case (task_alu_state)
			1'h0: begin
				task_alu_state <= 1'h1;
				task_alu_input_buffer = reg_data_read;
				reg_if_en <= 1'b1;
			end
			1'h1: begin
				task_alu_state <= 1'h0;
				continue_execution_register[`INS_ALU_AND] <= 1'b1; // Use any ALU instruction, it's all the same
				write_enable_register[`INS_ALU_AND] <= 1'b1;

				case (`INSDECOMP_INS(instruction_buffer))
					`INS_ALU_AND: task_alu_temp_buffer =  task_alu_input_buffer & reg_data_read;
					`INS_ALU_OR: task_alu_temp_buffer  =  task_alu_input_buffer | reg_data_read;
					`INS_ALU_ADD: task_alu_temp_buffer =  task_alu_input_buffer + reg_data_read;
					`INS_ALU_XOR: task_alu_temp_buffer =  task_alu_input_buffer ^ reg_data_read;
					`INS_ALU_MUL: task_alu_temp_buffer =  task_alu_input_buffer * reg_data_read;
					`INS_ALU_EQ: task_alu_temp_buffer  =  (task_alu_input_buffer ==  reg_data_read ? 16'hFFFF : 16'h0);
					`INS_ALU_GT: task_alu_temp_buffer  =  (task_alu_input_buffer > reg_data_read ? 16'hFFFF : 16'h0);
					`INS_ALU_SHFT: begin
						if (reg_addr_if & 4'h8) begin
							task_alu_temp_buffer = task_alu_input_buffer << (reg_addr_if & 4'b0111);
						end else begin
							task_alu_temp_buffer = task_alu_input_buffer >> reg_addr_if;
						end
					end
				endcase

				reg_data_write_inputs[`INS_ALU_AND] <= {task_alu_temp_buffer, 1'b1};
			end
		endcase
	end
	endtask


	// MEM(R|W)
    reg mem_read_ready_stor = 0;
    reg mem_read_ready_stor_reset = 0;
    always @(*) begin
        if (mem_read_ready) begin
            mem_read_ready_stor <= 1'b1;
        end else if (mem_read_ready_stor_reset) begin
            mem_read_ready_stor <= 0;
        end
    end

	// Address extension registers
	reg [4:0]
		mem_addr_ext_user,
		mem_addr_ext_kernel;

	// Read (MEMR)
    reg [16:0] task_memr_state;
    reg task_memr_is_cfg;
	task task_memr;
	begin
		if (reg_data_read[15] || task_memr_is_cfg) begin
            task_memr_is_cfg <= 1'b1;
        end

		case (task_memr_state)
			0: begin
				mem_readaddr <= {mem_addr_ext_kernel,mem_addr_ext_user,reg_data_read[14:0]};
				mem_read <= 1'h1;
				task_memr_state <= task_memr_is_cfg ? 2 : 1;
			end
            2: begin
				if (task_memr_is_cfg) begin
					// CFG register reads

					if (reg_data_read >= 16'hD000 && reg_data_read < (16'hD000 + `BOOTLOADER_ROM_SIZE)) begin
						reg_data_write_inputs[`INS_MEMR] <= {bootloader_rom[reg_data_read - 16'hD000], 1'b1}; // Bootloader ROM read
					end else if (reg_data_read == 16'h8004) begin
						reg_data_write_inputs[`INS_MEMR] <= {11'h0, mem_addr_ext_kernel, 1'b1};
					end else if (reg_data_read == 16'h8800) begin
						reg_data_write_inputs[`INS_MEMR] <= {11'h0, mem_addr_ext_user, 1'b1};
					end else if (reg_data_read == 16'h9000) begin
						reg_data_write_inputs[`INS_MEMR] <= irq_handler_addr;
					end else if (reg_data_read == 16'h9001) begin
						reg_data_write_inputs[`INS_MEMR] <= irq_en ? 16'hFFFF : 16'h0;
					end else if (reg_data_read == 16'h9002) begin
						reg_data_write_inputs[`INS_MEMR] <= in_irq_context ? 16'hFFFF : 16'h0;
					end else if (reg_data_read == 16'h9010) begin
						reg_data_write_inputs[`INS_MEMR] <= irq_dout[0+:16];
					end else if (reg_data_read == 16'h9011) begin
						reg_data_write_inputs[`INS_MEMR] <= irq_dout[16+:16];
					end

					mem_read <= 1'h0;
                    write_enable_register[`INS_MEMR] <= 1'h1;
                    continue_execution_register[`INS_MEMR] <= 1'h1;

				end else if (mem_read_ready_stor) begin
					// Direct RAM read
                    mem_read_ready_stor_reset <= 1'b1;
                    reg_data_write_inputs[`INS_MEMR] = {mem_readdata, 1'b1};

					mem_read <= 1'h0;

                    write_enable_register[`INS_MEMR] <= 1'h1;
                    continue_execution_register[`INS_MEMR] <= 1'h1;
                end
            end
			default: begin
				task_memr_state <= task_memr_state + 1;
			end
		endcase
	end
	endtask

	// Write (MEMW)
	`define MEMW_DELAY 3
	reg [16:0] task_memw_state;
	reg [15:0] task_memw_addr_buffer;
    reg task_memw_is_cfg;
	task task_memw;
	begin
        task_memw_is_cfg = task_memw_addr_buffer[15] || task_memw_is_cfg;

        case (task_memw_state)
            0: begin
                task_memw_addr_buffer <= reg_data_read;
                reg_if_en <= 1'b1;
                task_memw_state <= 1;
            end
            1: begin
                if (task_memw_is_cfg) begin
					// CFG register writes
                    if (task_memw_addr_buffer >= 16'hE000 && task_memw_addr_buffer <= 16'hF2BF) begin
						// Framebuffer write
                        fb_addr = task_memw_addr_buffer - 16'hE000;
                        fb_data = reg_data_read[0 +: 8];
                        fb_we = 1'b1;
                    end else if (task_memw_addr_buffer == 16'h8004) begin
						mem_addr_ext_kernel <= reg_data_read[4:0];
					end else if (task_memw_addr_buffer == 16'h8800) begin
						mem_addr_ext_user <= reg_data_read[4:0];
					end else if (task_memw_addr_buffer == 16'h9000) begin
						irq_handler_addr <= reg_data_read;
					end else if (task_memw_addr_buffer == 16'h9001) begin
						irq_en <= (|reg_addr_read) ? 1'h1 : 1'h0;
					end else if (task_memw_addr_buffer == 16'h9002) begin
						if (in_irq_context) begin
							in_irq_context <= (|reg_addr_read) ? 1'h1 : 1'h0;
						end
					end

					task_memw_state <= `MEMW_DELAY;
                end else begin
					// Direct RAM write
                    mem_writeaddr = {mem_addr_ext_kernel,mem_addr_ext_user,task_memw_addr_buffer[14:0]};
                    mem_writedata = reg_data_read;
                    mem_write <= 1'h1;

					task_memw_state <= 2;
                end
            end
            `MEMW_DELAY: begin
                mem_write <= 1'b0;
                fb_we <= 1'b0;
                continue_execution_register[`INS_MEMW] <= 1'h1;
            end
			default: begin
				task_memw_state <= task_memw_state + 1;
			end
        endcase
	end
	endtask


	// CALLED AFTER EACH INSTRUCTION AND ON RST
	integer ix;
	task reset_instruction_tasks;
	begin
		halted <= 0;

		reg_if_en <= 0;

		task_movnz_state <= 0;
		task_movez_state <= 0;

		task_set_state <= 0;

		task_memr_state <= 0;
		task_memr_is_cfg <= 0;
		task_memw_state <= 0;
		task_memw_addr_buffer <= 0;
        task_memw_is_cfg <= 0;
        mem_read_ready_stor_reset <= 0;

		mem_read <= 0;
		mem_readaddr <= 0;
		mem_write <= 0;
		mem_writeaddr <= 0;
		mem_writedata <= 0;

		for (ix = 0; ix < 16; ix = ix + 1) begin
		  reg_data_write_inputs[ix] <= 0;
		end
	end
	endtask


	// DEBUG OUTPUT LOGIC
	assign DEBUG_BUS = SW[9] ? reg_h_out : (SW[8] ? pc_out : (SW[7] ? instruction_buffer : (SW[6] ? mem_readdata :
		dbgDbg // Display debug stats if nothing else selected
	)));

	assign LEDR[0] = halted;
	assign LEDR[1] = reg_we;
	assign LEDR[2] = continue_execution;
	assign LEDR[3] = debugEn;
	assign LEDR[4] = clk;
	assign LEDR[5] = reg_if_en;
	assign LEDR[6] = cpu_state == `CPU_STATE_INS_LOAD;
	assign LEDR[7] = cpu_state == `CPU_STATE_WAITING;
	assign LEDR[8] = cpu_state == `CPU_STATE_COMMIT;
	assign LEDR[9] = cpu_state == `CPU_STATE_PC_INC;

endmodule
