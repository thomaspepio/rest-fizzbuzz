# Fizzbuzz REST server

### Endpoints
All endpoints are documented under http://host:port/swagger/index.html

### Project manipulation and layout

#### Project manipulation
A Makefile assists for all tasks.

Supported commands :
    - _build_ : downloads dependencies and builds the server binary
    - _start_ : starts the binary, you specify the port on which you want the app to start with `PORT=<port of your choice>`
    - _test_ : runs the app tests
    - _test-coverage_ : run the app tests and provides a `coverage.out`
    - _docs_ : make swagger documentation

#### Layout
- _domain_ : business logic in the form of pure functions (no side effects performed)

### Next steps
    - Cache : it is not necessary to compute each fizzbuzz request from scratch each time. We should implement an incremental cache.