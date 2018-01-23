package valid

import (
	"flag"
	"io/ioutil"
	"testing"
	"time"
)

func TestTimeMustSet(t *testing.T) {
	var opt time.Time
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Var(Time(&opt, time.Time{}).MustSet(), "opt", "")

	testParse(t, fs, &opt, false, nil)
	testParse(t, fs, &opt, true,
		time.Date(2018, 1, 23, 13, 23, 13, 0, time.UTC),
		"-opt", "2018-01-23T13:23:13Z")
}

// TODO: more tests for Time(): Min, Max, OneOf
