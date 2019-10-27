secrets
=======

Guard against logging secrets in Go.

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/pkgs/secrets)](https://goreportcard.com/report/gophers.dev/pkgs/secrets)
[![Build Status](https://travis-ci.com/shoenig/secrets.svg?branch=master)](https://travis-ci.com/shoenig/secrets)
[![GoDoc](https://godoc.org/gophers.dev/pkgs/secrets?status.svg)](https://godoc.org/gophers.dev/pkgs/secrets)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/secrets.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/secrets.svg)](LICENSE)

# Project Overview

Module `gophers.dev/pkgs/secrets` can be used to help protect against secret
values from being exposed in places they shouldn't be, particularly in log lines.

# Getting Started

The `secrets` module can be installed by running
```
$ go get gophers.dev/pkgs/secrets
```

#### Example Usage
```golang
// protect a string value
text := secrets.New("abc123")
fmt.Sprintf("%s", text) // prints "<redacted>"
fmt.Sprintf("%#v", text) // prints "secrets.Text{}"

// protect a byte slice
b := secrets.NewBytes([]byte{1, 2, 3})
fmt.Sprintf("%s", b) // prints "<redacted>"
fmt.Sprintf("%#v", b) // prints "secrets.Bytes{}"

// get access to the underlying secret values
doThings(text.Secret(), b.Secret())
```

# Contributing

The `gophers.dev/pkgs/secrets` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# License

The `gophers.dev/pkgs/secrets` module is open source under the [BSD-3-Clause](LICENSE) license.
