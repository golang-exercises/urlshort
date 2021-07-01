# urlshort

A command line tool that redirects given paths to a specific url based on given key-value pairs. `urlshort` supports `YAML`, `JSON` and `BoltDB`. Reference to [gophercise - Exercise 2](https://github.com/gophercises/urlshort).

## Installation

Create a new directory in your `$GOPATH` and make sure that in your `$GOPATH` exists a `bin` directory.

```bash
# Project structure
$GOPATH
|
--- bin
--- src
     |
     --- <YOUR-PROJECT>
```

## Build

Install the `make` tool and run :

```bash
make build
```

## Run the program

```bash
$GOPATH/bin/urlshort --yaml <path-to-yaml> --json <path-to-json> --port <int server starts on this port>
```

Examples of the accepted .yaml and .json can be found in the `/assets` folder.
