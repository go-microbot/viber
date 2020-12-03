all: lint

## golangci-lint run
lint:
	golangci-lint cache clean
	golangci-lint run --config .golangci.yml --timeout=5m

## generate mocks
gen-mocks:
	rm -rf ./api/mocks/*.go
	mockery --dir ./api/ --all --case underscore --output ./api/mocks --disable-version-string

## run tests
test:
	go test -p 1 -timeout 15m -covermode=count -coverprofile=coverage.out -coverpkg=./api/.,./bot/... ./api/. ./bot/...
