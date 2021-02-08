.PHONY: phony

phony: test build docs

test:
	go test ./...

test-coverage:
	rm -f coverage.out && go test --coverprofile=coverage.out ./... && go tool cover -func=coverage.out

build:
	go get && go build

docs:
	rm -rf ./endpoint/docs && swag init -g ./endpoint/endpoint.go -o ./endpoint/docs

run:
	./rest-fizzbuzz