all: test_unit benchmark build

build:
	go build permute.go

test_unit:
	go test -cover -v ./...

benchmark:
	go test -bench=. ./...