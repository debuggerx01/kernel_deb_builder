#!/usr/bin/env bash

VERSION=$(grep 'Kernel Configuration' < /boot/config-"$(uname -r)" | awk '{print $3}')

# add deb-src to sources.list
echo "deb-src https://home-packages.chinauos.com/home plum main contrib non-free" >> /etc/apt/sources.list

# install dep
sudo apt update
sudo apt install -y wget
sudo apt build-dep -y linux

# download kernel source
wget http://www.kernel.org/pub/linux/kernel/v5.x/linux-"$VERSION".tar.xz
tar -xf linux-"$VERSION".tar.xz
cd linux-"$VERSION" || exit

# copy config file
cp /boot/config-"$(uname -r)" .config

# change version string
sed -i "s/-desktop/-debuggerx/g" .config

# reduce ACPI_MAX_LOOP_TIMEOUT value
sed -i "/ACPI_MAX_LOOP_TIMEOUT/s/30/3/g" include/acpi/acconfig.h

# disable DEBUG_INFO to speedup build
scripts/config --disable DEBUG_INFO

CPU_CORES=$(grep -c processor < /proc/cpuinfo)
make deb-pkg -j"$CPU_CORES"