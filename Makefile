test:
	@go test ./...

coverage:
	@go test ./... -coverprofile=coverage.txt
