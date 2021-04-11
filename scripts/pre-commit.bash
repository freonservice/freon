#!/usr/bin/env bash

echo "Running golangci-lint"
golangci-lint run

# $? stores exit value of the last command
if [ $? -ne 0 ]; then
 echo "Golangci-lint must pass before commit!"
 exit 1
fi

echo "Running tests"
go test -v ./...

# $? stores exit value of the last command
if [ $? -ne 0 ]; then
 echo "Tests must pass before commit!"
 exit 1
fi
