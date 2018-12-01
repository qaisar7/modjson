package modjson

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Json is a structured represenation of the json.
type Json struct {
	obj map[string]interface{}
}

// Obj returns the structured represenation as a map.
func (j *Json) Obj() map[string]interface{} {
	if j == nil {
		return nil
	}
	return j.obj
}

// Print prints the json object.
func (j *Json) Print(prefix string, indent string) string {
	ind := Indent{prefix, indent, ""}
	st := fmt.Sprint("{\n")
	st = st + PrintMap(j, ind)
	st = st + fmt.Sprint("}\n")
	return st
}

//SearchByKey searches for they key in the Json and returns the resulting Json
func (j *Json) SearchByKey(k string) {

}

// Parse parses the given string and returns a Json object.
func Parse(st string) (*Json, error) {
	j := &Json{}
	var mr map[string]interface{}
	if err := json.Unmarshal([]byte(st), &mr); err != nil {
		return nil, fmt.Errorf("error unmarshaling, err:%v", err)
	}
	j.obj = mr
	return j, nil
}

// PrintMap prints the map of json according to the provided identation number.
func PrintMap(j *Json, ind Indent) string {
	var st string
	ind.indentIt()
	var keys []string
	if j == nil || j.Obj() == nil {
		return "j is nil"
	}
	for i := range j.obj {
		keys = append(keys, i)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if keys[i] < keys[j] {
			return true
		}
		return false
	})
	for n, i := range keys {
		switch c := j.obj[i].(type) {
		case map[string]interface{}:
			st = st + fmt.Sprintf("%s%q: {\n", ind.get(), i)
			st = st + PrintMap(&Json{obj: c}, ind)
			st = st + fmt.Sprintf("%s}%s\n", ind.get(), comma(n < len(keys)-1))
		case []interface{}:
			st = st + fmt.Sprintf("%s%q: [\n", ind.get(), i)
			st = st + PrintSlice(c, ind)
			st = st + fmt.Sprintf("%s]%s\n", ind.get(), comma(n < len(keys)-1))
		default:
			st = st + fmt.Sprintf("%s%q: %s%s\n", ind.get(), i, toString(j.obj[i]), comma(n < len(keys)-1))
		}
	}
	return st
}

func toString(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprintf("%d", i)
	case string:
		return fmt.Sprintf("%q", i)
	default:
		fmt.Println("returning default")
		return fmt.Sprintf("%v", i)
	}
}

// PrintSlice prints the slice object.
func PrintSlice(sl []interface{}, ind Indent) string {
	var st string
	ind.indentIt()
	for i, v := range sl {
		switch c := v.(type) {
		case map[string]interface{}:
			st = st + PrintMap(&Json{obj: c}, ind)
		case []interface{}:
			st = st + PrintSlice(c, ind)
		default:
			st = st + fmt.Sprintf("%s%s%s\n", ind.get(), toString(c), comma(i < len(sl)-1))
		}
	}
	return st
}
func comma(b bool) string {
	if b {
		return ","
	}
	return ""
}

// Indent specifies the indentation.
type Indent struct {
	prefix  string
	indent  string
	current string
}

func (ind *Indent) indentIt() {
	if ind.current == "" {
		ind.current = ind.prefix + ind.indent
		return
	}
	ind.current = ind.current + ind.indent
}

func (ind *Indent) get() string {
	return ind.current
}
