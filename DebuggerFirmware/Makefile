### PROJECT_DIR
PROJECT_DIR       = $(PWD)

### ARDMK_DIR
### Path to the Arduino-Makefile directory.
ARDMK_DIR         = /usr/share/arduino
ARDUINO_DIR       = /usr/share/arduino

### Path to where the your project's libraries are stored.
USER_LIB_PATH    :=  $(PROJECT_DIR)/libraries

### BOARD_TAG
BOARD_TAG         = mega
#BOARD_SUB         = 16MHzatmega328
MCU               = atmega2560

### MONITOR_BAUDRATE
MONITOR_BAUDRATE  = 115200

### AVR_TOOLS_DIR
AVR_TOOLS_DIR     = /usr

### AVRDUDE
AVRDUDE          = /usr/bin/avrdude
AVRDUDE_CONF     = /etc/avrdude.conf
AVRDUDE_MCU      = m2560
AVRDUDE_ARD_PROGRAMMER = stk500v2
AVRDUDE_ARD_BAUDRATE   = 115200

### CFLAGS_STD
### Set the C standard to be used during compilation. Documentation (https://github.com/WeAreLeka/Arduino-Makefile/blob/std-flags/arduino-mk-vars.md#cflags_std)
CFLAGS_STD        = -std=gnu11

### CXXFLAGS_STD
### Set the C++ standard to be used during compilation. Documentation (https://github.com/WeAreLeka/Arduino-Makefile/blob/std-flags/arduino-mk-vars.md#cxxflags_std)
CXXFLAGS_STD      = -std=gnu++11

### CXXFLAGS
### Flags you might want to set for debugging purpose. Comment to stop.
CXXFLAGS         += -pedantic -Wall -Wextra

### MONITOR_PORT
### The port your board is connected to. Using an '*' tries all the ports and finds the right one.
MONITOR_PORT      = /dev/ttyACM0

### CURRENT_DIR
### Do not touch - used for binaries path
CURRENT_DIR       = $(shell basename $(CURDIR))

### OBJDIR
OBJDIR            = $(PROJECT_DIR)/bin/$(BOARD_TAG)/$(CURRENT_DIR)

### Do not touch - the path to Arduino.mk, inside the ARDMK_DIR
include $(ARDMK_DIR)/Arduino.mk
