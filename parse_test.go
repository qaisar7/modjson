package modjson

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kylelemons/godebug/diff"
)

const (
	testDataPath = "./testdata"
)

func readFile(f string) (string, error) {
	f = testDataPath + "/" + f
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("unable to read file :%s", f)
	}
	return string(b), nil
}

func TestParseJson(t *testing.T) {
	tests := []struct {
		desc   string
		input  string
		output string
	}{
		{
			desc:   "success, parse json with maps and slices",
			input:  "test1.data",
			output: "test1.out.data",
		},
	}
	for _, tc := range tests {
		st, err := readFile(tc.input)
		if err != nil {
			t.Fatalf("unable to read from %s", tc.input)
		}
		want, err := readFile(tc.output)
		if err != nil {
			t.Fatalf("unable to read file %s", tc.output)
		}
		j, err := Parse(st)
		if err != nil {
			t.Errorf("unable to parse to json %v", err)
		}
		if diff := diff.Diff(want, j.Print(" ", " ")); diff != "" {
			t.Errorf("unexpected diff, (-want, +got)%s", diff)
		}

	}
}
