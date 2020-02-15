install:
	CGO_ENABLED=0 go install -v ./cmd/...

test:
	go test -v -race -cover -coverprofile=cover.out ./...

lint:
	golangci-lint run --enable-all ./...
