BIN := rfm-server

all: deps pbgo build

build:
	go build

deps:
	go get -t ./...

pbgo: rfm/rfm.pb.go

rfm/%.pb.go: rfm-pb/%.proto
	cd rfm-pb; protoc --go_out=plugins=grpc:../rfm $*.proto

clean:
	-rm $(BIN)

.PHONY: all deps pbgo clean
