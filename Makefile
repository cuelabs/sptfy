GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=sptfy
BINARY_UNIX=$(BINARY_NAME)_unix

#all: test build
all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/sptfy
#test:
#	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -fv ./cmd/sptfy/$(BINARY_NAME)
	rm -f ./cmd/sptfy$(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
