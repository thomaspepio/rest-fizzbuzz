.PHONY: phony

phony: build test docs

test:
	go test ./...

test-coverage:
	rm -f coverage.out && go test --coverprofile=coverage.out ./... && go tool cover -func=coverage.out

build:
	go get && go build && go get github.com/swaggo/swag/cmd/swag

docs:
	swag init -g ./endpoint/endpoint.go -o ./endpoint/docs

run:
	./rest-fizzbuzz
