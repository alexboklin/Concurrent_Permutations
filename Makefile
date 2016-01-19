all: test_unit build

build:
	go build permutate.go

test_unit:
	go test -cover -v ./...