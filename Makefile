SHELL := /bin/bash
SCRIPT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
PROJ_DIR := $(shell cd $(SCRIPT_DIR) && pwd)
BIN_DIR := $(PROJ_DIR)/bin
SRC_DIR := $(PROJ_DIR)/src

all: $(BIN_DIR)/rpi-version

$(BIN_DIR)/rpi-version:
	@mkdir -p $(BIN_DIR)
	cd $(SRC_DIR) && \
	GOOS=linux GOARCH=arm64 go build -o $(BIN_DIR)/rpi-version main.go
