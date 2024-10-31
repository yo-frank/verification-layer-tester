#!/bin/bash

echo "Building the sample program..."
go build -o sample sample.go

if [ $? -eq 0 ]; then
    echo "Build successful."
else
    echo "Build failed."
    exit 1
fi
