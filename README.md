![build](https://github.com/pgollangi/netselect/workflows/build/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/netselect)](https://goreportcard.com/report/github.com/pgollangi/netselect)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pgollangi/netselect)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/pgollangi/netselect)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pgollangi/netselect)](https://pkg.go.dev/github.com/pgollangi/netselect)
[![Maintainability](https://api.codeclimate.com/v1/badges/6236d95b7ca7a9554560/maintainability)](https://codeclimate.com/github/pgollangi/netselect/maintainability)
[![Code Climate technical debt](https://img.shields.io/codeclimate/tech-debt/pgollangi/netselect)](https://codeclimate.com/github/pgollangi/netselect/trends/technical_debt)
![License: MIT](https://img.shields.io/github/license/pgollangi/netselect)
[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/prassu158@gmail.com)
# netselect

A CLI tool as well as library to select the fastest host based on the lowest ICMP latency written in Go (golang), inspired by [apenwarr/netselect](https://github.com/apenwarr/netselect) debian package.

It’s a handy tool to choose a mirror for apt, or just to compare sites to each other. Under the hood it’s an ICMP ping.

## Features
- Finds the fastest host(s) in terms of network connectivity.
- Run desired concurent findings to get faster results. Use flag `--concurrent`.  
- Customize no. of ping attempt to perform for each host to get accurate mean response time. Use flag `--attempts`.
- Display only top `n` results on output. Use flag `--output`.
- Optionally, direct `netselect` to send "unprivileged" pings via UDP for non-sudo users. Use `--privileged=false`.

## Usage
`netselect` available as Commnad-Line tool and Go library.
### Commnad-Line

```sh
netselect [options] <host(s)>
```
#### Examples
```sh
$ netselect google.com google.in google.us
google.com       55 ms  100% ok         ( 3/ 3)
google.in        56 ms  100% ok         ( 3/ 3)
google.us        59 ms  100% ok         ( 3/ 3)
```

Read the  [documentation](https://pgollangi.github.io/netselect)  for more information on the CLI usage.

### Go Library

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

results, err := netSelector.Select()

fastestHosts := results // Fastest hosts in ASC order
```
Read the  [API documentation](https://pkg.go.dev/github.com/pgollangi/netselect) for more information on the library usage.

## Installation 

### Scoop
```sh
scoop bucket add pgollangi-bucket https://github.com/pgollangi/scoop-bucket.git
scoop install netselect
```
### Homebrew
```sh
brew install pgollangi/tap/netselect
```
Updating:
```
brew upgrade netselect
```
### Go
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

## Notice for linux users
`netelect` implements ICMP ping using both raw socket and UDP. It needs to be run as a root user.

Alternatley, you can use `setcap` to allow `netselect` to bind to raw sockets
```sh
setcap cap_net_raw=+ep /bin/netselect
```
If you do not wish to do all this, you can use flag `--privileged=false` to send an "unprivileged" ping via UDP. This must be enabled by setting

```sh
sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"
```

See [this blog](https://sturmflut.github.io/linux/ubuntu/2015/01/17/unprivileged-icmp-sockets-on-linux/) and the Go [icmp library](https://pkg.go.dev/golang.org/x/net/icmp?tab=doc) for more details.

## Contributing
Thanks for considering contributing to this project!

Please read the [Contributions](.github/CONTRIBUTING.md) and [Code of conduct](.github/CODE_OF_CONDUCT.md). 

Feel free to open an issue or submit a pull request!

## License

[MIT](LICENSE)

Copyright © [Prasanna Kumar](https://pgollangi.github.io/tabs/about)
