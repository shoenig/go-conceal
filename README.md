go-conceal
==========

Guard against logging secrets in Go.

```
conceal (verb):
  1. to hide; withdraw or remove from observation
  2. to keep secret; to prevent or avoid disclosing
```

[![Go Reference](https://pkg.go.dev/badge/shoenig/go-conceal.svg)](https://pkg.go.dev/shoenig/go-conceal)
[![GitHub](https://img.shields.io/github/license/shoenig/go-conceal.svg)](LICENSE)

# Project Overview

Module `github.com/shoenig/go-conceal` can be used to help protect against sensitive
values from being exposed in places they shouldn't be, particularly in log lines.

# Getting Started

The `conceal` module can be installed by running
```
$ go get github.com/shoenig/go-conceal
```

#### Example Usage
```golang
// protect a string value
text := secrets.New("abc123")
fmt.Sprintf("%s", text) // prints "<redacted>"
fmt.Sprintf("%#v", text) // prints "conceal.Text{}"

// protect a byte slice
b := secrets.NewBytes([]byte{1, 2, 3})
fmt.Sprintf("%s", b) // prints "<redacted>"
fmt.Sprintf("%#v", b) // prints "conceal.Bytes{}"

// get access to the underlying secret values
doThings(text.Unveil(), b.Unveil())
```

# Contributing

The `github.com/shoenig/go-conceal` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# License

The `github.com/shoenig/go-conceal` module is open source under the [BSD-3-Clause](LICENSE) license.
