# Go Math Quiz

This is a simple math quiz web site that can be deployed as a single binary. 

## Building
This project only uses the Go standard library so cloning it and running `go build` from inside the `go-math-quiz` directory should build a binary.

## Flags

- `max` = The maximum integer value used in the math problems
- `level` = Level indicates what operators are used in the quiz questions, 1 = add, 2 = sub, 3 = mul, 4 = div
- `port` = The port that the web server is running on

## Logging

As the game is played the server will log the questions and whether or no they were answered correctly as a CSV string
