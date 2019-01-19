---
title: "MCPC (MineCraft Personal Computer) - Architecture and hardware documentation"
author: [Stefan Reiter]
date: "2019-01-19"
keywords: [MCPC, FPGA, CPU, Verilog]
titlepage: true
listings-disable-line-numbers: true
listings-no-page-break: true
...

# Introduction

This document provides a rough documentation of the MCPC (MineCraft Personal Computer) architecture and hardware. It does not go into detail on the inner workings of the different subsystems, but rather strives to provide a high-level description of components to allow developers to treat them as black-box systems in higher-abstraction systems.

# Formatting

This document uses Verilog-style number formatting. See examples below:

* *h1C* refers to "1C" in hexadecimal notation
* *b1001* refers to "1001" in binary notation
* *42* refers to "42" in decimal notation (implicit base is 10)
* *h8b2* refers to the third bit of 8~16~
* *16'h9* refers to a 16-bit number containing 0009~16~

For memory locations, the following notation is used:

* [h1234] refers to the memory cell at 25'h0001234 (memory addressing is 25 bit total)
* [p2.h1234] refers to the memory cell at 25'h0021234 ("p2" denoting page 2~10~)
