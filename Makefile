.PHONY: all airbloc generate-proto clean install uninstall run test test-all
DEST = $(shell pwd)/build/bin

PROTO_DIR := proto
PROTO_SRCS := $(wildcard $(PROTO_DIR)/*.proto)

all: airbloc

airbloc:
	./env.sh go install ./cmd/airbloc
	@echo "$(DEST)/airbloc"

clean:
	@rm -rf build/

install: airbloc
	@cp -f $(DEST)/airbloc $GOPATH/bin/

generate-proto:
	@for PROTO in $(PROTO_SRCS); do protoc -I. $$PROTO --go_out=plugins=grpc:$$GOPATH/src; done

uninstall:
	@rm -f $GOPATH/bin/airbloc

run: airbloc
	@$(DEST)/airbloc

test: test-all

test-all:
	@go test -v ./...