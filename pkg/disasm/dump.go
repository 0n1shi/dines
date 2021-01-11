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
			dumpNormal(result)
		}
	}

}

func dumpNormal(result *Result) {
	dumpHeader(result.Header)
	address := 0x8000
	for _, section := range result.Sections {
		for _, line := range section.Lines {
			dumpAddress(address)
			fmt.Printf("\t")
			dumpRawData(line)
			fmt.Printf("\t")
			dumpInstruction(line, address)

			if line.Instruction == nil {
				address++
			} else {
				address += line.Instruction.Bytes
			}

			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func dumpHeader(header *Header) {
	fmt.Println("magic number: NES")
	fmt.Printf("program Bank: %d (%d byte)\n", header.ProgramBank.Count, header.ProgramBank.Size)
	fmt.Printf("character Bank: %d (%d byte)\n", header.CharacterBank.Count, header.CharacterBank.Size)
	fmt.Printf("mapper: %d (%s)\n\n", header.Mapper, MapperTypeMap[header.Mapper])
}

func dumpAddress(addr int) {
	fmt.Printf("0x%04X:", addr)
}

func dumpRawData(line *Line) {
	for i := 0; i < 4; i++ {
		if i < len(line.Data) {
			fmt.Printf("%02X ", line.Data[i])
			continue
		}
		fmt.Print("   ")
	}
}

func dumpInstruction(line *Line, currentAddr int) {
	if line.Instruction == nil { // invalid opcode, must be .db
		fmt.Print("db ")
		for _, d := range line.Data {
			fmt.Printf("%02X ", d)
		}
		return
	}

	arg := 0
	for i := len(line.Data) - 1; i > 0; i-- {
		arg = (arg << 8) | line.Data[i]
	}

	ins := line.Instruction

	fmt.Printf("%s ", OpecodeMap[ins.OpcodeType])
	switch line.Instruction.AddressingType {
	case AddressingTypeImmediate:
		fmt.Printf("#$%02X", arg)
	case AddressingTypeAbsolute:
		fmt.Printf("$%04X", arg)
	case AddressingTypeZeroPage:
		fmt.Printf("$%02X", arg)
	case AddressingTypeIndirect:
		fmt.Printf("($%04X)", arg)
	case AddressingTypeAbsoluteX:
		fmt.Printf("$%04X, X", arg)
	case AddressingTypeAbsoluteY:
		fmt.Printf("$%04X, Y", arg)
	case AddressingTypeZeroPageX:
		fmt.Printf("$%02X, X", arg)
	case AddressingTypeZeroPageY:
		fmt.Printf("$%02X, Y", arg)
	case AddressingTypeIndirectX:
		fmt.Printf("($%02X, X)", arg)
	case AddressingTypeIndirectY:
		fmt.Printf("($%02X), Y", arg)
	case AddressingTypeRelative:
		fmt.Printf("$%04X      # to $%04X", arg, (currentAddr+2)+int(int8(arg)))
	}
}
