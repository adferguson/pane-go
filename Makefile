export GOPATH:=$(GOPATH):$(shell pwd)

all:
	rm src/pane/pane.go
#	protoc --proto_path=protos/ --go_out=src/pane/ protos/*.proto
	generator pane.thrift src/
	go build

simpleclient:
	rm -rf gen-py
	thrift --gen py pane.thrift
