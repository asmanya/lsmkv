.PHONY: tools fmt lint test test-all race fuzz-smoke crash-smoke crash-long bench bench-compare cover clean

tools:
	@echo "--- tool versions ---"
	go version
	golangci-lint --version
	benchstat -h >/dev/null 2>&1 && echo "benchstat: ok"
	govulncheck -version
	gh --version
	make --version
	git --version

fmt:
	gofmt -l .
	test -z "$$(gofmt -l .)"

lint: fmt
	golangci-lint run ./...

test:
	go test -short ./...

test-all:
	go test ./...

race:
	go test -race -timeout 15m ./...

fuzz-smoke:
	@echo "not yet implemented (Phase 2)"

crash-smoke:
	@echo "not yet implemented (Phase 6)"

crash-long:
	@echo "not yet implemented (Phase 6)"

bench:
	go test -run '^$$' -bench . -benchmem ./...

bench-compare:
	@echo "not yet implemented (needs benchstat baseline, Phase 2+)"

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -f coverage.out coverage.html
	rm -rf bench-results/ crash-artifacts/
	go clean ./...
