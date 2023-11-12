.PHONY: test
test: *.go
	go test ./...

.PHONY: show-example
show-example:
	go run ./example
