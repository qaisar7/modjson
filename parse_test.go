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

func writeFile(f string, data []byte) error {
	f = testDataPath + "/" + f
	if err := ioutil.WriteFile(f, data, 0666); err != nil {
		return fmt.Errorf("unable to write to file :%s error:%v", f, err)
	}
	return nil
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
		{
			desc:   "success, parse json with only slice",
			input:  "test2.data",
			output: "test2.out.data",
		},
	}
	for _, tc := range tests {
		st, err := readFile(tc.input)
		if err != nil {
			t.Fatalf("%q: unable to read from %s", tc.desc, tc.input)
		}
		want, err := readFile(tc.output)
		if err != nil {
			t.Fatalf("%q: unable to read file %s", tc.desc, tc.output)
		}
		j, err := Parse(st)
		if err != nil {
			t.Errorf("%q: unable to parse to json %v", tc.desc, err)
		}
		got := j.Print(" ", " ")
		if diff := diff.Diff(want, got); diff != "" {
			t.Errorf("%q: unexpected diff, (-want, +got)\n%s", tc.desc, diff)
		}
		j1, err := Parse(got)
		if err != nil {
			t.Errorf("%q: unable to parse to json again %v", tc.desc, err)
		}
		got1 := j1.Print(" ", " ")
		if diff := diff.Diff(got, got1); diff != "" {
			t.Errorf("%q: unexpected diff, (-got previous, +got new)\n%s", tc.desc, diff)
		}
	}
}

func TestSearchByKey(t *testing.T) {
	t.Log("running this test")
	tests := []struct {
		desc  string
		input string
		key   string
		want  string
	}{
		{
			desc:  "success, search for a key in map",
			input: "test1.data",
			key:   "acl",
			want: ` "10.0.0.0/8": [
  "accept",
  "reject"
 ],
 "class": {
  "ip": "1.1.1.1",
  "v4": "2.2.2.2"
 }
`,
		},
		{
			desc:  "success, search for a key in slices",
			input: "test2.data",
			key:   "teams",
			want: ` "white",
 "blue"
`,
		},
	}
	for _, tc := range tests {
		st, err := readFile(tc.input)
		if err != nil {
			t.Fatalf("unable to read from %s", tc.input)
		}
		j, err := Parse(st)
		if err != nil {
			t.Errorf("unable to parse to json %v", err)
		}
		got := j.SearchByKey(tc.key)
		if diff := diff.Diff(tc.want, *got); diff != "" {
			t.Errorf("SearchByKey failed (-want, +got)\n%s", diff)
		}
	}
}
