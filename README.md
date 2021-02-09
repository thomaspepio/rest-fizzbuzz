# Fizzbuzz REST server

### Endpoints
- /1/fizzbuzz?limit=`limit`?fizzer=`fizzer`?buzzer=`buzzer`?fizz=`fizz`?buzz=`buzz` : returns a list of strings from 1 to `limit`, replacing all numbers that are multiples of `fizzer` by `fizz`, all multiples of `buzzer` by `buzz`, and all multiples of `fizzer` and `buzzer` by a concatenation of `fizz` and `buzz`

- /1/stats : returns statistics and the parameters of the most frequent fizzbuzz request

A more detailed documentation can be found under http://localhost:port/swagger/index.html

### Project manipulation and layout

#### Requirements
- Golang >= 1.15.6
- `$GOPATH` should be set
- `$PATH` should reference `$GOPATH/bin`

#### Project manipulation
A Makefile assists for all tasks.

Supported commands :
- _build_ : downloads dependencies and builds the server binary

- _start_ : starts the binary, you specify the port on which you want the app to start with `PORT=port of your choice`

- _test_ : runs the app tests

- _test-coverage_ : run the app tests and provides a `coverage.out`

- _docs_ : make swagger documentation

By default, typing `make` at the root of the project should run all the necessary steps for a `make run` to succeed.

#### Layout
- _domain_ : business logic (the concrete fizzbuzz algorithm) in the form of pure functions (no side effects performed)
- _service_ : implementation of the fizzbuzz service, with specific input and ouput types
- _endopint_ : HTTP API layer, matches HTTP semantics with the service layer
- _constant_ : constants used across various packages

### Next steps
- Cache : it is not necessary to compute each fizzbuzz request from scratch each time. We should implement an incremental cache.