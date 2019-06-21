install:
	@dep ensure -v
	
run:
	@go run main.go

test:
	@go test -v ./... -race -cover

build:
	go build -v -a -o challenge
