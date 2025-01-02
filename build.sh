#!/bin/sh
SCRIPT_DIRNAME=$(dirname "$0")
PROJ_DIR=$(cd "${SCRIPT_DIRNAME}" || exit 1; pwd)
if [ ! -e "${PROJ_DIR}/bin" ]; then
    mkdir "${PROJ_DIR}/bin"
fi
cd "${PROJ_DIR}/src" || exit 1
GOOS=linux GOARCH=arm64 go build -o "${PROJ_DIR}/bin/rpi-version" main.go
