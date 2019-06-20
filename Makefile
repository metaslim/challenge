install:
	@dep ensure
	
run:
	@go run main.go

test:
	@go test -v ./... -race -cover
