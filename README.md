![build](https://github.com/pgollangi/netselect-go/workflows/build/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/netselect-go)](https://goreportcard.com/report/github.com/pgollangi/netselect-go)
![License: MIT](https://img.shields.io/github/license/pgollangi/netselect-go)

# netselect

A CLI tool to select the fastest mirror based on the lowest ICMP latency written in Go (golang), inspired by [apenwarr/netselect](https://github.com/apenwarr/netselect) debian package.

## Usage

```sh
netselect [flags] <mirror(s)>
```
### Examples
```sh
netselect google.com twitter.com
netselect -v
netselect -h
```
## Learn More

Read the  [documentation](https://pgollangi.com/netselect)  for more information on the CLI tool.

## Installation

Download a binary suitable for your OS at the [releases page](https://github.com/pgollangi/netselect/releases/latest).

### NPM
```sh
npm install netselect
```

## Building from source

netselect CLI is written in the [Go programming language](https://golang.org/), so to build the CLI yourself, you first need to have Go installed and configured on your machine.

 ### Install Go

To download and install  `Go`, please refer to the  [Go documentation](https://golang.org/doc/install). Please download  `Go 1.14.x`  or above.

### Clone this repository
```sh
$ git clone https://gitlab.com/pgollangi/netselect.git
$ cd netselect
```
### Build project
On Unix based systems run:
```sh
build/build.sh
```
On Windows run:
```console
build/build.bat
```
Once completed, you will find the netselect CLI executable at `bin` directory. 
Run `bin/netselect -v` to check if it worked.

## License

`netselect` is under MIT License. See the [LICENSE](https://github.com/pgollangi/netselect-go/blob/main/LICENSE) file for details.