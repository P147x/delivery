# Makefile based on template
# https://sohlich.github.io/post/go_makefile/

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOPATH=$(shell pwd)
export GOPATH

BINARY_NAME=delivery
    
all: deps test build

build:
	$(GOBUILD) -o $(BINARY_NAME) delivery 
test: 
	$(GOTEST) delivery
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v delivery
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/0xAX/notificator
	$(GOGET) github.com/getlantern/systray
	$(GOGET) github.com/gocolly/colly