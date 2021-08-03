#!/usr/bin/env bash

# add deb-src to sources.list
echo "deb-src https://home-packages.chinauos.com/home plum main contrib non-free" >> /etc/apt/sources.list

# install dep
sudo apt update
sudo apt install -y wget
sudo apt build-dep -y linux

# download kernel source
wget http://www.kernel.org/pub/linux/kernel/v5.x/linux-5.10.41.tar.xz
tar -xf linux-5.10.41.tar.xz
cd linux-5.10.41 || exit

# copy config file
cp /boot/config-5.10.41-amd64-desktop .config

# change version string
sed -i "s/-desktop/-debuggerx/g" .config

# reduce ACPI_MAX_LOOP_TIMEOUT value
sed -i "/ACPI_MAX_LOOP_TIMEOUT/s/30/3/g" include/acpi/acconfig.h

CPU_CORES=$(grep -c processor < /proc/cpuinfo)
make deb-pkg -j"$CPU_CORES"