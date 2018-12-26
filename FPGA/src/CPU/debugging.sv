// Debugger Operation Codes
`define DEBUGGER_OPCODE_GET 4'h1
`define DEBUGGER_OPCODE_SET 4'h2
`define DEBUGGER_OPCODE_HI 4'h4
`define DEBUGGER_OPCODE_LO 4'h8
`define DEBUGGER_OPCODE_STEP 4'hC
`define DEBUGGER_OPCODE_DUMP_ROM 4'hE
`define DEBUGGER_OPCODE_DUMP_REGS 4'hA

`define BOOTLOADER_ROM_SIZE 2047

module cpu_debugger(
	//////////// GLOBALS ////////////
	input wire clk,
	input wire clk50,
	input wire rst,

    //////////// DEBUGGER-SPECIFIC ////////////
    output wire debugEn,

    output wire cpuClk,
    input wire [7:0] cpuState,
    input wire halted,

    input wire [15:0] regRead,
    output reg [3:0] regAddrOvr,
    output reg regAddrOvrEn,

    output wire [31:0] instructionOverride,
    output wire instructionOverrideEn,

    output wire [15:0] rom_addr,
    input wire [15:0] rom_data,

    output wire rstReq,

	//////////// GPIO ////////////
	inout 		    [15:0]		ARDUINO_IO,
    output reg      [15:0] dbgDbg
);

    // UART
    reg [7:0] uart_din;
    wire [7:0] uart_dout, uart_dout_rev, uart_din_rev;
    reg uart_wr_en = 1'b0;
    reg uart_rdy_clr = 1'b0;
    wire uart_rdy, uart_tx_busy;
    uart uart_instance(
        .rst(1'h0),
        .din(uart_din_rev),
	    .wr_en(uart_wr_en),
	    .clk_50m(clk50),
	    .tx(ARDUINO_IO[1]),
	    .tx_busy(uart_tx_busy),
	    .rx(ARDUINO_IO[0]),
	    .rdy(uart_rdy),
	    .rdy_clr(uart_rdy_clr),
	    .dout(uart_dout_rev)
    );

    // Legacy
    genvar i;
    generate
        for (i=0; i<8; i++) begin:uart_reverse_loop
            assign uart_din_rev[i] = uart_din[i];
            assign uart_dout[i] = uart_dout_rev[i];
        end
    endgenerate

    // Internal values
    reg [7:0] dbgRegs [0:15];
    reg [2:0] dbgRegWrite;

    reg stepDone = 1'h0;
    reg dbgClk = 1'h0;
    reg stepReq = 1'h0;

    reg dumpingRom = 1'h0;
    reg dumpingHiBits = 1'h0;
    reg [15:0] currentDumpAddr = 16'h0;

    reg dumpingRegs = 1'h0;
    reg [3:0] currentDumpReg = 4'h0;

    assign rom_addr = currentDumpAddr;

    assign debugEn = dbgRegs[0][0];
    assign cpuClk = debugEn && ~rst ? dbgClk : clk;

    assign dbgDbg[12] = dumpingRom;
    assign dbgDbg[13] = dumpingRegs;

    // Debug Register Assignments
    assign regAddrOvrEn =          debugEn && (dbgRegs[0][1] || dumpingRegs);
    assign regAddrOvr =            dumpingRegs ? currentDumpReg : dbgRegs[1][3:0];
    assign rstReq =                dbgRegs[0][2];
    assign instructionOverrideEn = dbgRegs[0][3];
    assign instructionOverride   = {dbgRegs[5], dbgRegs[4], dbgRegs[7], dbgRegs[6]};

    always_comb begin
        dbgRegs[8] <= regRead[7:0];
        dbgRegs[9] <= regRead[15:8];
        dbgRegs[4'hF][0] <= halted;
    end

    // Internal states
    reg stepStarted;
    reg prevUartRdy = 1'b0;

    reg dumpStarted = 1'b0;
    reg [7:0] checksum = 8'h0;
    reg checksumWrite = 1'h0;

    reg [3:0] resetHoldCounter = 4'h0;

    // Clocked debugger logic
    always @(posedge clk50) begin

        if (dumpingRom) begin

            // ROM dump logic
            if (~uart_tx_busy && ~dumpStarted) begin
                dumpStarted <= 1'b1;

                if (checksumWrite) begin
                    checksumWrite <= 1'b0;
                    dumpingRom <= 1'b0; // Safe to exit directly because we clear uart_wr_en in non-rom-dump path as well and dumpStarted is set to 0 on initialization of next dump
                    uart_din <= checksum;
                    uart_wr_en <= 1'b1;
                end else begin
                    if (dumpingHiBits) begin
                        uart_din = rom_data[15:8];
                    end else begin
                        uart_din = rom_data[7:0];
                    end

                    // Simple XOR based checksum
                    checksum <= checksum ^ uart_din;

                    if (dumpingHiBits) begin
                        currentDumpAddr = currentDumpAddr + 16'h1;
                    end

                    if (currentDumpAddr == (`BOOTLOADER_ROM_SIZE + 1)) begin
                        checksumWrite <= 1'h1;
                    end

                    dumpingHiBits <= ~dumpingHiBits;
                    uart_wr_en <= 1'b1;
                end

            end else begin

                if (uart_tx_busy) begin
                    dumpStarted <= 1'b0;
                end

                uart_wr_en <= 1'b0;
            end

        end else if (dumpingRegs) begin

            // Register dump logic
            if (~uart_tx_busy && ~dumpStarted) begin
                dumpStarted <= 1'b1;

                if (dumpingHiBits) begin
                    uart_din = regRead[15:8];
                end else begin
                    uart_din = regRead[7:0];
                end

                if (dumpingHiBits) begin
                    currentDumpReg = currentDumpReg + 4'h1;

                    if (currentDumpReg == 0) begin
                        dumpingRegs <= 1'h0;
                    end
                end

                dumpingHiBits <= ~dumpingHiBits;
                uart_wr_en <= 1'b1;

            end else begin

                if (uart_tx_busy) begin
                    dumpStarted <= 1'b0;
                end

                uart_wr_en <= 1'b0;
            end

        end else begin

            uart_din <= 8'hAB;

            // Command decoding
            uart_wr_en <= 1'b0;
            uart_rdy_clr <= 1'b0;
            if (uart_rdy && ~prevUartRdy) begin
                uart_rdy_clr <= 1'b1;
                resetHoldCounter <= 4'h0;

                dbgDbg[7:0] <= uart_dout;

                case (uart_dout[3:0])
                    `DEBUGGER_OPCODE_GET: begin
                        uart_din <= dbgRegs[{1'b1,uart_dout[6:4]}];
                        uart_wr_en <= 1'b1;
                    end

                    `DEBUGGER_OPCODE_SET: begin
                        dbgRegWrite <= uart_dout[6:4];
                        uart_wr_en <= 1'b1;
                    end

                    `DEBUGGER_OPCODE_HI: begin
                        dbgRegs[{1'b0,dbgRegWrite}][7:4] <= uart_dout[7:4];
                        uart_wr_en <= 1'b1;
                    end

                    `DEBUGGER_OPCODE_LO: begin
                        dbgRegs[{1'b0,dbgRegWrite}][3:0] <= uart_dout[7:4];
                        uart_wr_en <= 1'b1;
                    end

                    `DEBUGGER_OPCODE_STEP: begin
                        stepReq <= 1'h1;
                        stepStarted <= 1'h1;
                        stepDone <= 1'h0;
                    end

                    `DEBUGGER_OPCODE_DUMP_ROM: begin
                        dumpingRom <= 1'h1;
                        currentDumpAddr <= 16'h0;
                        dumpingHiBits <= 1'h0;
                        dumpStarted <= 1'h0;
                        checksumWrite <= 1'h0;
                        checksum <= 8'h0;
                    end

                    `DEBUGGER_OPCODE_DUMP_REGS: begin
                        dumpingRegs <= 1'h1;
                        currentDumpReg <= 4'h0;
                        dumpingHiBits <= 1'h0;
                        dumpStarted <= 1'h0;
                    end
                endcase
            end else begin
                // Reset logic
                if (rst) begin

                    stepDone <= 1'h0;
                    dbgClk <= 1'h0;
                    stepStarted <= 1'h0;
                    stepReq <= 1'h0;

                    dumpingRom <= 1'h0;
                    currentDumpAddr <= 16'h0;
                    dumpingHiBits <= 1'h0;
                    dumpStarted <= 1'h0;
                    checksumWrite <= 1'h0;
                    checksum <= 8'h0;

                    // Reset debugger state if reset was not triggered by us
                    /*if (~dbgRegs[0][2]) begin
                        for (int i=0; i<8; i++) begin
                            dbgRegs[i] <= 8'h0;
                        end
                    end*/

                // Step logic
                end else if (stepReq && ~stepDone) begin

                    if (halted || (cpuState == 0 && ~stepStarted)) begin
                        stepDone <= 1'h1;
                        stepReq <= 1'h0;
                    end

                    dbgClk <= ~dbgClk;
                    stepStarted <= 1'h0;

                end else if (stepDone) begin
                    uart_wr_en <= 1'b1;
                    stepDone <= 1'b0;
                end

            end

            prevUartRdy = uart_rdy;

        end

    end

endmodule
