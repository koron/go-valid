package valid

import (
	"flag"
	"io"
	"testing"
)

func TestBoolMustSet(t *testing.T) {
	var opt bool
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(Bool(&opt, false).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, true, "-opt")
	testParse(t, fs, &opt, true, true, "-opt=true")
	testParse(t, fs, &opt, true, false, "-opt=false")
}
