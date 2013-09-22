export GOPATH:=$(GOPATH):$(shell pwd)

all:
		rm src/pane/*.pb.go
		protoc --proto_path=protos/ --go_out=src/pane/ protos/*.proto
		go build
