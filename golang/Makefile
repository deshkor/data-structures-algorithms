all: tidy vendor test

tidy:
	go mod tidy

vendor:
	go mod vendor

test:
	go test -count=1 ./...