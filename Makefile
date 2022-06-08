# Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

BUILD_DIR = $(CURDIR)/bin
PKG_NAME = pcache

clean:
	rm -rf $(CURDIR)/bin

build:
	go build -o $(BUILD_DIR)/$(PKG_NAME) $(CURDIR)/cmd/$(PKG_NAME)/$(PKG_NAME).go
