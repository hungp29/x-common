GO          := go

format:
	$(GO) fmt ./...

deps:
	$(GO) mod tidy
	$(GO) mod download
