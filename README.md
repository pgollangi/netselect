![build](https://github.com/pgollangi/netselect/workflows/build/badge.svg?branch=master)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pgollangi/netselect)](https://pkg.go.dev/github.com/pgollangi/netselect)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/netselect)](https://goreportcard.com/report/github.com/pgollangi/netselect)
![License: MIT](https://img.shields.io/github/license/pgollangi/netselect)

# netselect

A CLI tool as well as library to select the fastest host based on the lowest ICMP latency written in Go (golang), inspired by [apenwarr/netselect](https://github.com/apenwarr/netselect) debian package.

## Usage (CLI)

```sh
netselect [options] <host(s)>
```
### Examples
```sh
$ netselect google.com google.in google.us
google.com       55 ms  100% ok         ( 3/ 3)
google.in        56 ms  100% ok         ( 3/ 3)
google.us        59 ms  100% ok         ( 3/ 3)
```
### Learn More

Read the  [documentation](https://pgollangi.com/netselect)  for more information on the CLI tool.

## Usage (Go Library)

Here is a simple example that finds fastest hosts:

```go
hosts := []*netselect.Host{
    &netselect.Host{Address: "google.in"},
    &netselect.Host{Address: "google.com"},
}

netSelector, err :=netselect.NewNetSelector(hosts)
if err != nil {
    panic(err)
}

fastestHosts := netSelector.Hosts // Fastest hosts in ASC order
```


## Installation 

```sh
$ go get github.com/pgollangi/netselect/cmd/netselect
$ netselect
```

### Manual

1. Download and install binary from the [latest release](https://github.com/pgollangi/netselect/releases/latest).
2. Recommended: add `netselect` executable to your $PATH.

## Building from source

`netselect` CLI is written in the [Go programming language](https://golang.org/), so to build the CLI yourself, you first need to have Go installed and configured on your machine.

 ### Install Go

To download and install  `Go`, please refer to the  [Go documentation](https://golang.org/doc/install). Please download  `Go 1.14.x`  or above.

### Clone this repository
```sh
$ git clone https://gitlab.com/pgollangi/netselect.git
$ cd netselect
```
### Build

```sh
$ go build cmd/netselect/netselect.go
$ netselect

```

## Contributions
Thanks for considering contributing to this project!

Please read the [Contributions](.github/CONTRIBUTING.md) and [Code of conduct](.github/CODE_OF_CONDUCT.md). 

Feel free to open an issue or submit a pull request!

## License

[MIT](LICENSE)

Copyright © [Prasanna Kumar](https://pgollangi.com)
