package modjson

import (
	"fmt"
	"io/ioutil"
	"testing"
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
			output: "test1_out.data",
		},
	}
	for _, tc := range tests {
		st, err := readFile(tc.input)
	}
}
