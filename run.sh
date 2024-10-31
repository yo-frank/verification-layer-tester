#!/bin/bash

if [ -f sample ]; then
    echo "Running the sample program..."
    if [ $# -eq 1 ]; then
        ./sample "$1"
    else
        ./sample
    fi
else
    echo "Executable not found. Please run ./build.sh first."
    exit 1
fi
