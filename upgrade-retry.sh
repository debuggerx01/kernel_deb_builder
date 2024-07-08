#!/bin/bash

RETRIES=50

for i in $(seq 1 $RETRIES); do
    apt full-upgrade -y && break || {
        echo "Failed, retrying... ($i)"
        sleep 5
    }
done
