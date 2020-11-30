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
	go test -p 1 -timeout 15m -covermode=count -coverprofile=coverage.out -coverpkg=./api/.,./bot/...,./query/...,./form/... ./api/. ./bot/... ./query/... ./form/...

## start docker containers
start_images:
	docker run --publish 8081:8081 -it --rm -d --name telegram-bot-api huntechio/telegram-bot-api:master-7cf91e4 --api-id=${TEST_API_ID} --api-hash=${TEST_API_HASH}

## stop docker containers
stop_images:
	docker rm -f telegram-bot-api || true
