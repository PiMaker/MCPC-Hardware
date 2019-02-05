.PHONY: default
default: program;

help:
	@echo "build:          Build .sof file with Quartus II"
	@echo "program:        Program FPGA device via JTAG"
	@echo "clean:          Clean unnecessary file and empty build/ directory"

# Custom Altera-specific target
OUTDIR = build/
PROJECT = MCPC
TOP_LEVEL_ENTITY = main
ASSIGNMENT_FILES = $(OUTDIR)$(PROJECT).qpf $(OUTDIR)$(PROJECT).qsf

FAMILY = "MAX10"
PART = 10M50DAF484C7G
BOARDFILE = DE10Pins

SRCS = $(shell find src/ -type f -regex '.*\.s?v$$')

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

MAP_ARGS = --read_settings_files=on $(addprefix --source=,$(SRCS))

FIT_ARGS = --part=$(PART) --read_settings_files=on
ASM_ARGS =
STA_ARGS =

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
	rm -rf build/db || true # Delete cached bootloader files
	quartus_sh --determine_smart_action $(OUTDIR)$(PROJECT) -c $(PROJECT) > $(OUTDIR)smart.log

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

program: $(OUTDIR)output_files/$(PROJECT).sof
	quartus_pgm --no_banner --mode=jtag -o "P;$(OUTDIR)output_files/$(PROJECT).sof"
