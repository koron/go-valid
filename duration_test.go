package valid

import (
	"flag"
	"io/ioutil"
	"testing"
	"time"
)

func TestDurationMustSet(t *testing.T) {
	var opt time.Duration
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Duration(&opt, 0).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, 123*time.Second, "-opt", "123s")
	testParse(t, fs, &opt, true, 999*time.Second, "-opt", "999s")
}

func TestDurationMin(t *testing.T) {
	var opt time.Duration
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Duration(&opt, 0).Min(10*time.Second), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, false, nil, "-opt", "9s")
	testParse(t, fs, &opt, true, 10*time.Second, "-opt", "10s")
	testParse(t, fs, &opt, true, 123*time.Second, "-opt", "123s")
}

func TestDurationMax(t *testing.T) {
	var opt time.Duration
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Duration(&opt, 0).Max(10*time.Second), "opt", "")

	testParse(t, fs, &opt, false, nil, "-opt", "11s")
	testParse(t, fs, &opt, true, 10*time.Second, "-opt", "10s")
	testParse(t, fs, &opt, true, 0*time.Second)
	testParse(t, fs, &opt, true, -123*time.Second, "-opt", "-123s")
}

func TestDurationOneOf(t *testing.T) {
	var opt time.Duration
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Duration(&opt, 0).OneOf(100*time.Second, 200*time.Second, 300*time.Second), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, 100*time.Second, "-opt", "100s")
	testParse(t, fs, &opt, true, 200*time.Second, "-opt", "200s")
	testParse(t, fs, &opt, true, 300*time.Second, "-opt", "300s")
	testParse(t, fs, &opt, false, nil, "-opt", "1s")
	testParse(t, fs, &opt, false, nil, "-opt", "99s")
	testParse(t, fs, &opt, false, nil, "-opt", "101s")
	testParse(t, fs, &opt, false, nil, "-opt", "900s")
}
