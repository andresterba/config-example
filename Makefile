GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=config-example

all: build
build: 
	$(GOBUILD) -o $(BINARY_NAME)
test: 
	$(GOTEST) ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: build
	./config-example

