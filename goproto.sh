#!/bin/sh

export GOPATH=/data/app/slgdev/go
export PATH=$PATH:/data/app/slgdev/go/bin

cd protocol
protoc --version
protoc --go_out=../src/protocol *.proto

echo "Done!"

