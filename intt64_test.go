package valid

import (
	"flag"
	"testing"
)

func TestInt64MustSet(t *testing.T) {
	var opt int64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Int64(&opt, 0).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, int64(123), "-opt", "123")
	testParse(t, fs, &opt, true, int64(999), "-opt", "999")
}

func TestInt64Min(t *testing.T) {
	var opt int64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Int64(&opt, 0).Min(10), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, false, nil, "-opt", "9")
	testParse(t, fs, &opt, true, int64(10), "-opt", "10")
	testParse(t, fs, &opt, true, int64(123), "-opt", "123")
}

func TestInt64Max(t *testing.T) {
	var opt int64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Int64(&opt, 0).Max(10), "opt", "")

	testParse(t, fs, &opt, false, nil, "-opt", "11")
	testParse(t, fs, &opt, true, int64(10), "-opt", "10")
	testParse(t, fs, &opt, true, int64(0))
	testParse(t, fs, &opt, true, int64(-123), "-opt", "-123")
}

func TestInt64OneOf(t *testing.T) {
	var opt int64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Int64(&opt, 0).OneOf(100, 200, 300), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, int64(100), "-opt", "100")
	testParse(t, fs, &opt, true, int64(200), "-opt", "200")
	testParse(t, fs, &opt, true, int64(300), "-opt", "300")
	testParse(t, fs, &opt, false, nil, "-opt", "1")
	testParse(t, fs, &opt, false, nil, "-opt", "99")
	testParse(t, fs, &opt, false, nil, "-opt", "101")
	testParse(t, fs, &opt, false, nil, "-opt", "900")
}
