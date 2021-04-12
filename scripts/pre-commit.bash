#!/usr/bin/env bash

echo "Running golangci-lint"
if golangci-lint run; then
  echo "Golangci-lint completed successful"
else
  echo "Golangci-lint must pass before commit!"
  exit 1
fi

echo "Running tests"
if go test -race -v ./...; then
  echo "Tests completed successful"
else
  echo "Tests must pass before commit!"
  exit 1
fi

echo "Code is successful committed"
