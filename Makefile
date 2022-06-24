check:
	make test
	make vet
	make fmt

test:
	go test -v ./...

start:
	go run cmd/app/*.go

race:
	go run -race cmd/app/*.go

vet:
	go vet ./...

fmt:
	gofmt -w .
