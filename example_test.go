package valid_test

import (
	"flag"
	"fmt"

	valid "github.com/koron/go-valid"
)

func ExampleInt() {
	var mode, user int
	fs := flag.NewFlagSet("example", flag.ContinueOnError)
	// "-user" must be set.
	fs.Var(valid.Int(&user, 0).MustSet(), "user", "target user ID (required)")
	// "-mode" should be chosen from 1, 2, 3, 4 or 5.
	fs.Var(valid.Int(&mode, 1).Min(1).Max(5), "mode", "choose between 1 and 5 (default 1)")
	if err := valid.Parse(fs, []string{"-user", "123"}); err != nil {
		fmt.Println(err)
	}
	// Output:
}
