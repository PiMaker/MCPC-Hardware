/*
  MCPC Debugger Firmware

  Wiring:
    = Description = Arduino Pin = FPGA Pin =
    - Ground - GND - GND -
    - Debug enable - 2 - IO0
    - SDA - SDA (20 on mega) - IO1
    - SCL - SCL (21 on mega) - IO2

  Serial interface commands (<Newline> terminates):
    - ping ... pong
    - enable ... Enable debugger
    - disable ... Disable debugger
    - step ... perform a single step
    - read <reg> ... returns value of MCPC register number <reg>
                     <reg> is in hex without "0x"
    - run <op_code> <op_code> ... execute two op codes (does not change PC,
                                  second op code is arbitrary, but probably most useful as SET value or NOOP)
                                  <op_code> is in hex without "0x", e.g. "5A31" for a MOV (opcode 0x1)

*/

#define DEBUGGER_OPCODE_GET 1
#define DEBUGGER_OPCODE_SET 2
#define DEBUGGER_OPCODE_HI 4
#define DEBUGGER_OPCODE_LO 8
#define DEBUGGER_OPCODE_STEP 0xC

#include "Arduino.h"
#include "HardwareSerial.h"
#include "./Cmd.h"
#include "avr/iom2560.h"

extern HardwareSerial Serial;

boolean enabled = false;

// Ping-Pong
void fn_ping(int arg_cnt, char **args) {
  Serial.println("pong");
}


// Comm functions
void sendRaw(uint8_t data) {
  DDRA = ~data;
  delay(1);
  pinMode(9, OUTPUT);
  delay(1);
  pinMode(9, INPUT);
  delay(15);
}

void setReg(uint8_t reg, uint8_t value) {
  sendRaw(DEBUGGER_OPCODE_SET | (reg<<4));
  sendRaw(DEBUGGER_OPCODE_LO | ((value & 0x0F) << 4));
  sendRaw(DEBUGGER_OPCODE_HI | (value & 0xF0));
}

uint8_t getReg(uint8_t reg) {
  sendRaw(DEBUGGER_OPCODE_GET | (reg<<4));
  return (uint8_t)(PINC);
}


// Cmd functions
void fn_enable(int arg_cnt, char **args) {
  pinMode(8, OUTPUT);
  enabled = true;

  Serial.println("success");
}

void fn_disable(int arg_cnt, char **args) {
  pinMode(8, INPUT);
  enabled = false;

  Serial.println("success");
}

void fn_read(int arg_cnt, char **args) {
  if (arg_cnt < 2) {
    Serial.println("error:too few arguments, need reg number");
  } else if (!enabled) {
    Serial.println("error:debugger disabled");
  } else {

    uint8_t reg = (uint8_t)cmdStr2Num(args[1], 16);
    setReg(0, 1);
    setReg(1, reg);
    uint8_t regLow = getReg(8);
    uint8_t regHigh = getReg(9);
    setReg(0, 0);
    
    Serial.print(reg, HEX);
    Serial.print(":");
    Serial.print(regHigh, HEX);
    if (regLow < 0x10) {
      Serial.print("0");
    }
    Serial.print(regLow, HEX);

    Serial.println();
    Serial.println("success");
  }
}

void fn_readAll(int arg_cnt, char **args) {

  if (!enabled) {
    Serial.println("error:debugger disabled");
  } else {

    for(uint8_t reg = 0; reg < 16; reg++)
    {
      setReg(0, 1);
      setReg(1, reg);
      uint8_t regLow = getReg(8);
      uint8_t regHigh = getReg(9);
      setReg(0, 0);
      
      Serial.print(reg, HEX);
      Serial.print(":");
      Serial.print(regHigh, HEX);
      if (regLow < 0x10) {
        Serial.print("0");
      }
      Serial.print(regLow, HEX);
      Serial.println();
    }

    Serial.println("success");
    
  }
}

void fn_step(int arg_cnt, char **args) {

  if (!enabled) {
    Serial.println("error:debugger disabled");
  } else {

    sendRaw(DEBUGGER_OPCODE_STEP);
    Serial.println("waiting");
    delay(5);
    while (!digitalRead(10)) {}
    Serial.println("success");
  }
}

void fn_reset(int arg_cnt, char **args) {

  if (!enabled) {
    Serial.println("error:debugger disabled");
  } else {

    setReg(7, 0xFF);
    Serial.println("success");
  }
}

void setup() {
  // Initialize comm pins
  PORTA = 0x00;
  DDRA = 0x00;
  PORTC = 0x00;
  DDRC = 0x00;
  
  // Initialize other pins
  digitalWrite(8, LOW);
  digitalWrite(9, LOW);
  digitalWrite(10, LOW);

  // Initialize Serial
  cmdInit(115200);

  // Initialize different commands
  cmdAdd("ping", &fn_ping);
  cmdAdd("enable", &fn_enable);
  cmdAdd("disable", &fn_disable);
  cmdAdd("regs", &fn_readAll);
  cmdAdd("read", &fn_read);
  cmdAdd("step", &fn_step);
  cmdAdd("reset", &fn_reset);

  Serial.println("mcpc debugger 0.1");
}

void loop() {
  cmdPoll();
}
