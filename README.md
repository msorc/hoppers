UPDATE: the solution now is considered incomplete with the case test/broken.input. We can return to the visited position with different speed to jump over an obstacle with a higher velocity.

# Hopper

We hope to hop with hopping hope.

## Run

``` sh
cat test/test.input | go run cmd/hoppers/main.go
cat test/test2.input | go run cmd/hoppers/main.go
cat test/wrong.input | go run cmd/hoppers/main.go
```

## Tests

``` sh
go test ./pkg/** -v
```

