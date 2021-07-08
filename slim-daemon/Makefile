# Copyright 2017-2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

QUIET := @
GO_BUILD := go build -mod=vendor
IMAGE_VERSION ?= latest

TARGET := anet-agent
WIN_TARGET := anet-agent.exe

all: $(TARGET) $(WIN_TARGET)

windows: $(WIN_TARGET)

.PHONY: all windows $(TARGET) $(WIN_TARGET)

$(TARGET):
	$(QUIET)$(GO_BUILD) -o $(TARGET)

$(WIN_TARGET):
	$(QUIET)GOOS=windows $(GO_BUILD) -o $(WIN_TARGET)

clean:
	$(QUIET)rm -f $(TARGET) $(WIN_TARGET)

.PHONY: clean

windows-image:
	$(QUIET)docker build -t anet-agent-windows:$(IMAGE_VERSION) -f Dockerfile.windows ..

.PHONY: windows-image
