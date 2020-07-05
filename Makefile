mod-init:
	go mod init github.com/ecshreve/cardz

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/cardz github.com/ecshreve/cardz/cmd/cardz

install:
	go install -i github.com/ecshreve/cardz/cmd/cardz

run-only:
	bin/cardz

run: build run-only

test:
	go test github.com/ecshreve/cardz/...

testv:
	go test -v github.com/ecshreve/cardz/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/cardz/...