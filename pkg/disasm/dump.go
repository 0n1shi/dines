package disasm

import (
	//"encoding/json"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
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
	for _, section := range result.Sections {
		for _, line := range section.Lines {
			dumpAddress(line.Address)
			fmt.Printf("\t")
			dumpRawData(line)
			fmt.Printf("\t")
			dumpInstruction(line)
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func dumpHeader(header *Header) {
	grn := color.New(color.FgHiGreen)
	fmt.Print("magic number: ")
	grn.Println("NES")
	fmt.Print("program Bank: ")
	grn.Print(header.ProgramBank.Count)
	fmt.Print(" (")
	grn.Printf("%d byte", header.ProgramBank.Size)
	fmt.Println(")")
	fmt.Printf("character Bank: ")
	grn.Print(header.CharacterBank.Count)
	fmt.Print(" (")
	grn.Printf("%d byte", header.CharacterBank.Size)
	fmt.Println(")")
	fmt.Print("mapper: ")
	grn.Print(header.Mapper)
	fmt.Print(" (")
	grn.Print(MapperTypeMap[header.Mapper])
	fmt.Println(")")
}

func dumpAddress(addr int) {
	fmt.Printf("0x%04X:", addr)
}

func dumpRawData(line *Line) {
	yellow := color.New(color.FgYellow)
	for i := 0; i < 4; i++ {
		if i < len(line.Data) {
			yellow.Printf("%02X ", line.Data[i])
			continue
		}
		fmt.Print("   ")
	}
}

func dumpInstruction(line *Line) {
	comment := color.New(color.FgHiGreen, color.Bold)
	reg := color.New(color.FgHiMagenta)
	db := color.New(color.Bold)
	opcode := color.New(color.FgHiBlue, color.Bold)
	dollar := color.New(color.FgYellow)
	hash := color.New(color.FgHiRed)

	if line.Instruction == nil { // invalid opcode, must be .db
		db.Print("db ")
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

	opcode.Printf("%s ", OpecodeMap[ins.OpcodeType])

	switch line.Instruction.AddressingType {
	case AddressingTypeImmediate:
		hash.Print("#")
		dollar.Print("$")
		fmt.Printf("%02X", arg)
	case AddressingTypeAbsolute:
		dollar.Print("$")
		fmt.Printf("%04X", arg)
	case AddressingTypeZeroPage:
		dollar.Print("$")
		fmt.Printf("%02X", arg)
	case AddressingTypeIndirect:
		fmt.Print("(")
		dollar.Print("$")
		fmt.Printf("%04X)", arg)
	case AddressingTypeAbsoluteX:
		dollar.Print("$")
		fmt.Printf("%04X, ", arg)
		reg.Print("X")
	case AddressingTypeAbsoluteY:
		dollar.Print("$")
		fmt.Printf("%04X, ", arg)
		reg.Print("Y")
	case AddressingTypeZeroPageX:
		dollar.Print("$")
		fmt.Printf("%02X, ", arg)
		reg.Print("X")
	case AddressingTypeZeroPageY:
		dollar.Print("$")
		fmt.Printf("%02X, ", arg)
		reg.Print("Y")
	case AddressingTypeIndirectX:
		fmt.Print("(")
		dollar.Print("$")
		fmt.Printf("%02X, ", arg)
		reg.Print("X")
		fmt.Print(")")
	case AddressingTypeIndirectY:
		fmt.Print("(")
		dollar.Print("$")
		fmt.Printf("%02X), ", arg)
		reg.Print("Y")
	case AddressingTypeRelative:
		dollar.Print("$")
		fmt.Printf("%04X      ", arg)
		comment.Printf("# to $%04X", (line.Address+2)+int(int8(arg)))
	}
}
