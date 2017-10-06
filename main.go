package parsejson

import (
	"encoding/json"
	"fmt"
)

func main() {
	st := `
	{
		"bgp": {
			"neighbor": {
				"family": "inet",
				"address": "192.168.1.1",
				"type": "ibgp",
				"acl": {
					"10.0.0.0/8": ["accept", "reject"],
					"class" : {
						"ip": "1.1.1.1",
						"v4": "2.2.2.2"
					}
				}
			}
		}
	}
	`
	var mr map[string]interface{}
	err := json.Unmarshal([]byte(st), &mr)
	if err != nil {
		fmt.Errorf("error unmarshaling, err:%v", err)
	}
	fmt.Println(mr)
	printMap(mr, " ")
	fmt.Println(mr)
	s, _ := json.Marshal(mr)
	fmt.Println(string(s))
}

func printMap(mr map[string]interface{}, indent string) {
	indent2 := indent + "  "
	for i, v := range mr {
		switch c := v.(type) {
		case map[string]interface{}:
			if i == "acl" {
				mr["acl1"] = v
			}
			delete(mr, "acl")
			//continue
			fmt.Printf("%s%s: {\n", indent2, i)
			printMap(c, indent2)
			fmt.Printf("%s}\n", indent2)
		case []interface{}:
			fmt.Printf("%s%s: [\n", indent2, i)
			printSlice(c, indent2)
			fmt.Printf("%s]\n", indent2)
		default:
			//fmt.Println(c)
			fmt.Printf("%s%s: %s\n", indent2, i, v)
		}
	}
}

func printSlice(sl []interface{}, indent string) {
	indent2 := indent + "  "
	for _, v := range sl {
		switch c := v.(type) {
		case map[string]interface{}:
			printMap(c, indent2)
		case []interface{}:
			printSlice(c, indent2)
		default:
			//fmt.Println(c)
			fmt.Printf("%s%s\n", indent2, c)
		}
	}
}
