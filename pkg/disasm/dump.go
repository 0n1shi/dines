package disasm

import (
	//"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type DumpMethod uint8

const (
	DumpMethodNormal DumpMethod = DumpMethod(iota)
	DumpMethodJson
	DumpMethodYaml
)

func Dump(result *Result, method DumpMethod) {

	switch method {
	case DumpMethodJson:
		{

		}
	}
	// address := 0x8000
	// for _, section := range result.Sections {
	// 	section.
	// }
	//j, _ := json.Marshal(result)
	y, _ := yaml.Marshal(result)
	fmt.Printf("%s", y)
}
