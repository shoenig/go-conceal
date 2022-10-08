secrets
=======

Guard against logging secrets in Go.

[![Go Reference](https://pkg.go.dev/badge/shoenig/secrets.svg)](https://pkg.go.dev/shoenig/secrets)
[![GitHub](https://img.shields.io/github/license/shoenig/secrets.svg)](LICENSE)

# Project Overview

Module `github.com/shoenig/secrets` can be used to help protect against secret
values from being exposed in places they shouldn't be, particularly in log lines.

# Getting Started

The `secrets` module can be installed by running
```
$ go get github.com/shoenig/secrets
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

The `github.com/shoenig/secrets` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# License

The `github.com/shoenig/secrets` module is open source under the [BSD-3-Clause](LICENSE) license.
