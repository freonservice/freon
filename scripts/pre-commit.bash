#!/usr/bin/env bash

echo "Running goimports test"
if test -z "$(goimports -d $(find . -type f -name '*.go'|grep -v -f .goimportsignore) 2>&1 | tee /dev/fd/2)"; then
  echo "Goimports completed successful"
else
  echo "Goimports must pass before commit!"
  exit 1
fi

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
