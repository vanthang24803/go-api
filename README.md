# Running the Application

This is a basic guide on how to run the Go API application.

## Prerequisites

To run the application, you will need to have the following installed on your machine:

- Go >= 1.22.4 : [Download and Install Go SDK](https://go.dev/doc/install)
- Git: [Download and Install Git](https://git-scm.com/downloads)


## License

This project is licensed under the [MIT License](LICENSE).


## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```