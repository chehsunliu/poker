.PHONY: test
test:
	go test -coverprofile .coverprofile -timeout 10m ./...
	go tool cover -html .coverprofile -o .coverprofile.html

.PHONY: benchmark
benchmark:
	go test -bench=. -benchtime 5s
