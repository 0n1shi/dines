package disasm

import (
	//"encoding/json"
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type DumpMethod string

const (
	DumpMethodNormal DumpMethod = DumpMethod("normal")
	DumpMethodJson   DumpMethod = DumpMethod("json")
	DumpMethodYaml   DumpMethod = DumpMethod("yaml")
)

func Dump(result *Result, method DumpMethod) {
	switch method {
	case DumpMethodJson:
		{
			j, _ := json.Marshal(result)
			fmt.Printf("%s", j)
		}
	case DumpMethodYaml:
		{
			y, _ := yaml.Marshal(result)
			fmt.Printf("%s", y)
		}
	default:
		{
			fmt.Println("hello world")
		}
	}

}
