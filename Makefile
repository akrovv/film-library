build:
	mkdir api && go build -o ./api ./...

test-api:
	go test --cover ./... | grep -v 'no test files'

lint:
	golangci-lint run ./...
