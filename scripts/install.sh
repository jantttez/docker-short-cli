#!/bin/bash

clear

PARENT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)
go build -o "${PARENT_DIR}/dcshort" "${PARENT_DIR}/main.go"
sudo mv "${PARENT_DIR}/dcshort" /usr/local/bin

echo "install success"
