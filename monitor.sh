#!/usr/bin/env bash

while [ -f monitor.sh ]; do
  echo ================================================
  df -h
  echo ================================================
  sleep 5
done