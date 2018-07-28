help:
	@echo "check:          Check code validity with iverilog"
	@echo "simulate:       Simulate with textual output"
	@echo "display:        Simulate and display waveform with gtkwave"
	@echo "build:          Build .sof file with Quartus II"
	@echo "program:        Program FPGA device via JTAG"
	@echo "clean:          Clean unnecessary file and empty build/ directory"

###############################################################################
#
# ICARUS VERILOG & GTKWAVE MAKEFILE
# MADE BY WILLIAM GIBB FOR HACDC
# williamgibb@gmail.com
#
# USE THE FOLLOWING COMMANDS WITH THIS MAKEFILE
#	"make check" - compiles your verilog design - good for checking code
#	"make simulate" - compiles your design+TB & simulates your design
#	"make display" - compiles, simulates and displays waveforms
#
###############################################################################
#
# CHANGE THESE THREE LINES FOR YOUR DESIGN
#
#TOOL INPUT
SOURCEDIR = src/
SRC = $(shell find src/ -type f \( -name '*.v' ! -name 'testbench.v' \))
TESTBENCH = src/testbench.v
TBOUTPUT = waves.lxt	#THIS NEEDS TO MATCH THE OUTPUT FILE
			#FROM YOUR TESTBENCH
###############################################################################
# BE CAREFUL WHEN CHANGING ITEMS BELOW THIS LINE
###############################################################################
#TOOLS
COMPILER = iverilog
SIMULATOR = vvp
VIEWER = gtkwave
#TOOL OPTIONS
COFLAGS = -o
SFLAGS =
SOUTPUT = -lxt		#SIMULATOR OUTPUT TYPE
#TOOL OUTPUT
COUTPUT = icarus.out			#COMPILER OUTPUT
###############################################################################
#MAKE DIRECTIVES
check : $(TESTBENCH) $(SRC)
	$(COMPILER) $(SRC)

simulate: $(COUTPUT)
	$(SIMULATOR) $(SFLAGS) $(COUTPUT) $(SOUTPUT)

display: $(TBOUTPUT)
	$(VIEWER) $(TBOUTPUT) &
#MAKE DEPENDANCIES
$(TBOUTPUT): $(COUTPUT)
	$(SIMULATOR) $(SOPTIONS) $(COUTPUT) $(SOUTPUT)

$(COUTPUT): $(TESTBENCH) $(SRC)
	$(COMPILER) $(COFLAGS) $(COUTPUT) $(TESTBENCH) $(SRC)



# Custom Altera-specific target


###################################################################
# Project Configuration:
#
# Specify the name of the design (project) and the Quartus II
# Settings File (.qsf)
###################################################################

OUTDIR = build/
PROJECT = MCPC
TOP_LEVEL_ENTITY = main
ASSIGNMENT_FILES = $(OUTDIR)$(PROJECT).qpf $(OUTDIR)$(PROJECT).qsf

###################################################################
# Part, Family, Boardfile DE1 or DE2
FAMILY = "MAX10"
PART = 10M50DAF484C7G
BOARDFILE = DE10Pins
###################################################################

###################################################################
# Setup your sources here
SRCS = $(shell find src/ -type f \( -name '*.v' ! -name 'testbench.v' \))
###################################################################
# Main Targets
#
# build: build everything
# clean: remove output files and database
# program: program your device with the compiled design
###################################################################

build: $(OUTDIR)smart.log $(OUTDIR)$(PROJECT).asm.rpt $(OUTDIR)$(PROJECT).sta.rpt

clean:
	cd $(OUTDIR); \
	rm -rf *
	rm -f $(TBOUTPUT) a.out $(COUTPUT)

map: $(OUTDIR)smart.log $(OUTDIR)$(PROJECT).map.rpt
fit: $(OUTDIR)smart.log $(OUTDIR)$(PROJECT).fit.rpt
asm: $(OUTDIR)smart.log $(OUTDIR)$(PROJECT).asm.rpt
sta: $(OUTDIR)smart.log $(OUTDIR)$(PROJECT).sta.rpt
smart: $(OUTDIR)smart.log

###################################################################
# Executable Configuration
###################################################################

MAP_ARGS = --read_settings_files=on $(addprefix --source=,$(SRCS))

FIT_ARGS = --part=$(PART) --read_settings_files=on
ASM_ARGS =
STA_ARGS =

###################################################################
# Target implementations
###################################################################

STAMP = echo done >

$(OUTDIR)$(PROJECT).map.rpt: $(OUTDIR)map.chg $(SOURCE_FILES)
	quartus_map $(MAP_ARGS) $(OUTDIR)$(PROJECT)
	$(STAMP) $(OUTDIR)fit.chg

$(OUTDIR)$(PROJECT).fit.rpt: $(OUTDIR)fit.chg $(OUTDIR)$(PROJECT).map.rpt
	quartus_fit $(FIT_ARGS) $(OUTDIR)$(PROJECT)
	$(STAMP) $(OUTDIR)asm.chg
	$(STAMP) $(OUTDIR)sta.chg

$(OUTDIR)$(PROJECT).asm.rpt: $(OUTDIR)asm.chg $(OUTDIR)$(PROJECT).fit.rpt
	quartus_asm $(ASM_ARGS) $(OUTDIR)$(PROJECT)

$(OUTDIR)$(PROJECT).sta.rpt: $(OUTDIR)sta.chg $(OUTDIR)$(PROJECT).fit.rpt
	quartus_sta $(STA_ARGS) $(OUTDIR)$(PROJECT)

$(OUTDIR)smart.log: $(ASSIGNMENT_FILES)
	cp -n DE10_Lite.SDC build/
	ln -s ../src build/src || true
	quartus_sh --determine_smart_action $(OUTDIR)$(PROJECT) -c $(PROJECT) > $(OUTDIR)smart.log

###################################################################
# Project initialization
###################################################################

$(ASSIGNMENT_FILES):
	quartus_sh --prepare -f $(FAMILY) -t $(TOP_LEVEL_ENTITY) $(OUTDIR)$(PROJECT)
	-cat $(BOARDFILE) >> $(OUTDIR)$(PROJECT).qsf
$(OUTDIR)map.chg:
	$(STAMP) $(OUTDIR)map.chg
$(OUTDIR)fit.chg:
	$(STAMP) $(OUTDIR)fit.chg
$(OUTDIR)sta.chg:
	$(STAMP) $(OUTDIR)sta.chg
$(OUTDIR)asm.chg:
	$(STAMP) $(OUTDIR)asm.chg

###################################################################
# Programming the device
###################################################################

program: $(OUTDIR)output_files/$(PROJECT).sof
	quartus_pgm --no_banner --mode=jtag -o "P;$(OUTDIR)output_files/$(PROJECT).sof"
