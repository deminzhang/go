#!/bin/sh

# export GOPATH=/data/app/slgdev/go
# export PATH=$PATH:/data/app/slgdev/go/bin

cd protos3
protoc --version
protoc --go_out=../src/protos *.proto

echo "Done!"

