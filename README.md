# Fizzbuzz REST server

### Project manipulation and layout

#### Manipulation
All commands are to be run at the root of the project

##### Running the tests
With code coverage : `go test --coverprofile=coverage.out ./... && go tool cover -func=coverage.out`

Without code coverage : `go test ./...`

##### Building the project
`go get && go build`

##### Running the app
Provided the project is built : launch `./rest-fizzbuzz`

A server should start on "localhost:8080" (TODO : make the port parametrizable)

#### Layout
- _domain_ : business logic in the form of pure functions (no side effects performed)
