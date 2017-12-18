package valid

import (
	"flag"
	"io/ioutil"
	"testing"
)

func TestUintMustSet(t *testing.T) {
	var opt uint
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint(&opt, 0).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, uint(123), "-opt", "123")
	testParse(t, fs, &opt, true, uint(999), "-opt", "999")
}

func TestUintMin(t *testing.T) {
	var opt uint
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint(&opt, 0).Min(10), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, false, nil, "-opt", "9")
	testParse(t, fs, &opt, true, uint(10), "-opt", "10")
	testParse(t, fs, &opt, true, uint(123), "-opt", "123")
}

func TestUintMax(t *testing.T) {
	var opt uint
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint(&opt, 0).Max(10), "opt", "")

	testParse(t, fs, &opt, false, nil, "-opt", "11")
	testParse(t, fs, &opt, true, uint(10), "-opt", "10")
	testParse(t, fs, &opt, true, uint(0))
	testParse(t, fs, &opt, true, uint(5), "-opt", "5")
}

func TestUintOneOf(t *testing.T) {
	var opt uint
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint(&opt, 0).OneOf(100, 200, 300), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, uint(100), "-opt", "100")
	testParse(t, fs, &opt, true, uint(200), "-opt", "200")
	testParse(t, fs, &opt, true, uint(300), "-opt", "300")
	testParse(t, fs, &opt, false, nil, "-opt", "1")
	testParse(t, fs, &opt, false, nil, "-opt", "99")
	testParse(t, fs, &opt, false, nil, "-opt", "101")
	testParse(t, fs, &opt, false, nil, "-opt", "900")
}
