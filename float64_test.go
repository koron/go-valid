package valid

import (
	"flag"
	"testing"
)

func TestFloat64MustSet(t *testing.T) {
	var opt float64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Float64(&opt, 0).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, 123.0, "-opt", "123")
	testParse(t, fs, &opt, true, 999.0, "-opt", "999")
}

func TestFloat64Min(t *testing.T) {
	var opt float64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Float64(&opt, 0).Min(10), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, false, nil, "-opt", "9")
	testParse(t, fs, &opt, true, 10.0, "-opt", "10")
	testParse(t, fs, &opt, true, 123.0, "-opt", "123")
}

func TestFloat64Max(t *testing.T) {
	var opt float64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Float64(&opt, 0).Max(10), "opt", "")

	testParse(t, fs, &opt, false, nil, "-opt", "11")
	testParse(t, fs, &opt, true, 10.0, "-opt", "10")
	testParse(t, fs, &opt, true, 0.0)
	testParse(t, fs, &opt, true, -123.0, "-opt", "-123")
}

func TestFloat64OneOf(t *testing.T) {
	var opt float64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Var(Float64(&opt, 0).OneOf(100, 200, 300), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, 100.0, "-opt", "100")
	testParse(t, fs, &opt, true, 200.0, "-opt", "200")
	testParse(t, fs, &opt, true, 300.0, "-opt", "300")
	testParse(t, fs, &opt, false, nil, "-opt", "1")
	testParse(t, fs, &opt, false, nil, "-opt", "99")
	testParse(t, fs, &opt, false, nil, "-opt", "101")
	testParse(t, fs, &opt, false, nil, "-opt", "900")
}
