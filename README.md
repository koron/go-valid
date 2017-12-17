# validate flag constraints

[![GoDoc](https://godoc.org/github.com/koron/go-valid?status.svg)](https://godoc.org/github.com/koron/go-valid)
[![CircleCI](https://circleci.com/gh/koron/go-valid.svg?style=svg)](https://circleci.com/gh/koron/go-valid)

Package "valid" provides constraints validatable values for "flag" package.

```go
package main

import (
    "flag"
    "fmt"

    valid "github.com/koron/go-valid"
)

func ExampleInt() {
    fs := flag.NewFlagSet("example", flag.ContinueOnError)
    // "-user" must be set.
    var user int
    fs.Var(valid.Int(&user, 0).MustSet(), "user",
        "target user ID (required)")
    // "-mode" should be chosen from 1, 2, 3, 4 or 5.
    var mode int
    fs.Var(valid.Int(&mode, 1).Min(1).Max(5), "mode",
        "choose between 1 and 5 (default 1)")
    // parse option with validation
    if err := valid.Parse(fs, os.Args); err != nil {
        panic(err)
    }
    // run your procedure with "user" and "mode".
}
```

## Description

Package "valid" provides `flag.Value` compatible values with validation for
some constraints. `flag.Value` can be used with `flag.Var()` or
`flag.FlagSet#Var()`.

Functions `{Type}(p *type, defaultValue type)` returns a new `*{Type}Value`.
This value have methods which declare some constraints like `MustSet()`,
`Min()` or so. For example "int" value which is between 1 and 5 with default 1:

```go
var user int
uv := valid.IntValue(&user, 1).Min(1).Max(5)
```

Another example, "int" value which must be set:

```go
var mode int
mv := valid.IntValue(&mode, 0).MustSet()
```

Then you should combine those values with the "flag" package.

```go
flag.Var(uv, "user", "user ID (mandatory)")
flag.Var(mv, "mode", "between 1 and 5")
flag.Parse()
```
