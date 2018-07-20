TEST_FLAGS = -short
EXTRA_FLAGS = 

.PHONY: test
test:
	go test -coverprofile .coverprofile -timeout 10m $(TEST_FLAGS) $(EXTRA_FLAGS) ./...
	go tool cover -html .coverprofile -o .coverprofile.html

.PHONY: benchmark
benchmark:
	go test -bench=.