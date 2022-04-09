lint:
	golangci-lint run --timeout 5m  ./...;

lint-fix:
	golangci-lint run --fix --timeout 5m  ./...;

lint-install:
	brew install golangci/tap/golangci-lint

lint-upgrade:
	brew upgrade golangci/tap/golangci-lint

test:
	go test ./...
