#!/usr/bin/env bash

# reduce ACPI_MAX_LOOP_TIMEOUT value
sed -i "/ACPI_MAX_LOOP_TIMEOUT/s/30/3/g" include/acpi/acconfig.h