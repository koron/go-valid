package valid

import (
	"flag"
	"io"
	"testing"
)

func TestStringMustSet(t *testing.T) {
	var opt string
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(String(&opt, "").MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, "abc", "-opt", "abc")
	testParse(t, fs, &opt, true, "zzz", "-opt", "zzz")
}

func TestStringMin(t *testing.T) {
	var opt string
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(String(&opt, "").Min(5), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, false, nil, "-opt", "abc")
	testParse(t, fs, &opt, false, nil, "-opt", "abcd")
	testParse(t, fs, &opt, true, "abcde", "-opt", "abcde")
	testParse(t, fs, &opt, true, "ABCDEF", "-opt", "ABCDEF")
}

func TestStringMax(t *testing.T) {
	var opt string
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(String(&opt, "").Max(5), "opt", "")

	testParse(t, fs, &opt, false, nil, "-opt", "ABCDEF")
	testParse(t, fs, &opt, true, "abcde", "-opt", "abcde")
	testParse(t, fs, &opt, true, "abcd", "-opt", "abcd")
	testParse(t, fs, &opt, true, "abc", "-opt", "abc")
	testParse(t, fs, &opt, true, "")
}

func TestStringOneOf(t *testing.T) {
	var opt string
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(String(&opt, "").OneOf("foo", "bar", "baz"), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true, "foo", "-opt", "foo")
	testParse(t, fs, &opt, true, "bar", "-opt", "bar")
	testParse(t, fs, &opt, true, "baz", "-opt", "baz")
	testParse(t, fs, &opt, false, nil, "-opt", "FOO")
	testParse(t, fs, &opt, false, nil, "-opt", "qux")
}
