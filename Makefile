#
# Standard makefile for go projects
#

# standard definitions
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
GOPATH=$(shell pwd)
SRC=$(GOPATH)/src
VENDOR=$(SRC)/vendor
GVT=$(GOPATH)/bin/gvt
LINT=$(GOPATH)/bin/golint
BIN=$(GOPATH)/bin

# project specific definitions
BASE_NAME=user-ws
SRC_TREE=userws
RUNNER=scripts/entry.sh

build: build-darwin build-linux

build-darwin:
	GOPATH=$(GOPATH) GOOS=darwin GOARCH=amd64 $(GOBUILD) -a -o $(BIN)/$(BASE_NAME).darwin $(SRC_TREE)

build-linux:
	GOPATH=$(GOPATH) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -installsuffix cgo -o $(BIN)/$(BASE_NAME).linux $(SRC_TREE)

test:
	GOPATH=$(GOPATH) $(GOTEST) -v $(SRC_TREE)/tests $(if $(TEST),-run $(TEST),)

fmt:
	GOPATH=$(GOPATH) $(GOFMT) $(SRC_TREE)/...

vet:
	GOPATH=$(GOPATH) $(GOVET) $(SRC_TREE)/...

lint:
	GOPATH=$(GOPATH) $(LINT) $(SRC_TREE)/...

clean:
	GOPATH=$(GOPATH) $(GOCLEAN)
	rm -f $(BIN)/$(BASE_NAME).*

run:
	rm -f $(BIN)/$(BASE_NAME)
	ln -s $(BIN)/$(BASE_NAME).darwin $(BIN)/$(BASE_NAME)
	$(RUNNER)

deps:
	rm -fr $(VENDOR)
	cd $(SRC); $(GOGET) -u github.com/golang/lint/golint
	cd $(SRC); $(GOGET) -u github.com/FiloSottile/gvt
	cd $(SRC); $(GVT) fetch github.com/gorilla/mux
	cd $(SRC); $(GVT) fetch github.com/nmcclain/ldap
	# for tests
	cd $(SRC); $(GVT) fetch gopkg.in/yaml.v2
	cd $(SRC); $(GVT) fetch github.com/parnurzeal/gorequest

#
# end of file
#
