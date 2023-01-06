#!/usr/bin/env bash

# reduce ACPI_MAX_LOOP_TIMEOUT value
sed -i "/ACPI_MAX_LOOP_TIMEOUT/s/30/3/g" include/acpi/acconfig.h

# make kernel version to 5.xx.0
sed -i '/SUBLEVEL =/s/[0-9].*/0/g' Makefile
