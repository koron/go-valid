package valid_test

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
	if err := valid.Parse(fs, []string{"-user", "123"}); err != nil {
		fmt.Println(err)
	}
	// Output:
}
