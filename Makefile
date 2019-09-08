.PHONY: test
test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -timeout=10m ./...
	go tool cover -html coverage.txt -o coverage.html

.PHONY: benchmark
benchmark:
	go test -bench=. -benchtime 5s
