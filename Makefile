BIN := bin
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./node_modules/*")
NAME := j

TAGS ?=

ifeq ($(OS), Windows_NT)
	EXECUTABLE := $(NAME).exe
	UNAME := Windows
else
	EXECUTABLE := $(NAME)
	UNAME := $(shell uname -s)
endif

ifeq ($(UNAME), Darwin)
	GOBUILD ?= go build -i
else
	GOBUILD ?= go build
endif

.PHONY: build
build: $(BIN)/$(EXECUTABLE) $(BIN)/$(EXECUTABLE)-debug

$(BIN)/$(EXECUTABLE): $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ .

$(BIN)/$(EXECUTABLE)-debug: $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -gcflags '$(GCFLAGS)' -o $@ .