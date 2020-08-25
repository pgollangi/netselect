# WARNING!!

This tool is under development. Not yet ready for production use.

![build](https://github.com/pgollangi/netselect/workflows/build/badge.svg?branch=master)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pgollangi/netselect)](https://pkg.go.dev/github.com/pgollangi/netselect)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/netselect)](https://goreportcard.com/report/github.com/pgollangi/netselect)
![License: MIT](https://img.shields.io/github/license/pgollangi/netselect)

# netselect

A CLI tool as well as library to select the fastest host based on the lowest ICMP latency written in Go (golang), inspired by [apenwarr/netselect](https://github.com/apenwarr/netselect) debian package.

## Usage

```sh
netselect [options] <host(s)>
```
### Examples
```sh
netselect m1.example.com m2.example.com m3.example.com
```
Output:
```
m2.example.com       56 ms         100% ok         ( 3/ 3)
m3.example.com      136 ms         100% ok         ( 3/ 3)
m1.example.com      294 ms         100% ok         ( 3/ 3)
```
## Learn More

Read the  [documentation](https://pgollangi.com/netselect)  for more information on the CLI tool.

<!---
## Installation 

Download a binary suitable for your OS at the [releases page](https://github.com/pgollangi/netselect/releases/latest).

### NPM
```sh
npm install netselect
```
--->

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

## Contributions
Thanks for considering contributing to this project!

Please read the [Contributions](.github/CONTRIBUTING.md) and [Code of conduct](.github/CODE_OF_CONDUCT.md). 

Feel free to open an issue or submit a pull request!

## License

[MIT](LICENSE)

Copyright © [Prasanna Kumar](https://pgollangi.com)