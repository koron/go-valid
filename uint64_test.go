package valid

import (
	"flag"
	"io/ioutil"
	"testing"
)

func TestUint64MustSet(t *testing.T) {
	var opt uint64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint64(&opt, 0).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, uint64(123), "-opt", "123")
	testParse(t, fs, &opt, true, uint64(999), "-opt", "999")
}

func TestUint64Min(t *testing.T) {
	var opt uint64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint64(&opt, 0).Min(10), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, false, nil, "-opt", "9")
	testParse(t, fs, &opt, true, uint64(10), "-opt", "10")
	testParse(t, fs, &opt, true, uint64(123), "-opt", "123")
}

func TestUint64Max(t *testing.T) {
	var opt uint64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint64(&opt, 0).Max(10), "opt", "")

	testParse(t, fs, &opt, false, nil, "-opt", "11")
	testParse(t, fs, &opt, true, uint64(10), "-opt", "10")
	testParse(t, fs, &opt, true, uint64(0))
	testParse(t, fs, &opt, true, uint64(5), "-opt", "5")
}

func TestUint64OneOf(t *testing.T) {
	var opt uint64
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Uint64(&opt, 0).OneOf(100, 200, 300), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, uint64(100), "-opt", "100")
	testParse(t, fs, &opt, true, uint64(200), "-opt", "200")
	testParse(t, fs, &opt, true, uint64(300), "-opt", "300")
	testParse(t, fs, &opt, false, nil, "-opt", "1")
	testParse(t, fs, &opt, false, nil, "-opt", "99")
	testParse(t, fs, &opt, false, nil, "-opt", "101")
	testParse(t, fs, &opt, false, nil, "-opt", "900")
}
