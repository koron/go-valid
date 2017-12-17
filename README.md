# validate flag constraints

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
