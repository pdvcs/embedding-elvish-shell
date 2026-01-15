# embedding-elvish-shell

A sample Go application that embeds the [Elvish](https://elv.sh/) interpreter to provide shell scripting capabilities.

## Features

- **Embedded Elvish**: Run Elvish scripts directly from a Go binary.
- **Argument Passing**: Pass command-line arguments to Elvish scripts via the native `$args` variable.
- **Standard I/O**: Full support for standard input, output, and error streams.

## Prerequisites

- [Go](https://go.dev/) 1.22 or later.

## Quick Start

```shell
cd ~/go/src/embedding-elvish-shell
go build -o embedded-elvish-example
./embedded-elvish-example sample.elv hello elvish world
```

## Usage

```shell
./embedded-elvish-example <script.elv> [arguments...]
```

### Accessing Arguments in Elvish

Inside your Elvish script, you can access the arguments passed to
`embedded-elvish-example` using the `$args` variable:

```elvish
echo "The script is:" (src)[name]
for arg $args[0..] {
    echo "Additional argument: "$arg
}
```

## How it Works

The application uses the `src.elv.sh/pkg/eval` package to:
1. Initialize a new `eval.Evaler`.
2. Inject the `$args` variable into the global namespace.
3. Configure standard I/O ports.
4. Parse and evaluate the provided script.

## Project Structure

- `main.go`: The Go entry point that embeds Elvish.
- `sample.elv`: A sample Elvish script.
- `go.mod` / `go.sum`: Go module configuration.
