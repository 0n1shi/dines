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

func Dump(result *Result, method DumpMethod, colored bool) {
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
			dumpNormal(result, colored)
		}
	}

}

func dumpNormal(result *Result, colored bool) {
	dumpHeader(result.Header, colored)
	for _, section := range result.Sections {
		for _, line := range section.Lines {
			dumpAddress(line.Address)
			fmt.Printf("\t")
			dumpRawData(line, colored)
			fmt.Printf("\t")
			dumpInstruction(line, colored)
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func dumpHeader(header *Header, colored bool) {
	val := color.New()
	if colored {
		val = color.New(color.FgHiGreen)
	}
	fmt.Print("magic number: ")
	val.Println("NES")
	fmt.Print("program Bank: ")
	val.Print(header.ProgramBank.Count)
	fmt.Print(" (")
	val.Printf("%d byte", header.ProgramBank.Size)
	fmt.Println(")")
	fmt.Printf("character Bank: ")
	val.Print(header.CharacterBank.Count)
	fmt.Print(" (")
	val.Printf("%d byte", header.CharacterBank.Size)
	fmt.Println(")")
	fmt.Print("mapper: ")
	val.Print(header.Mapper)
	fmt.Print(" (")
	val.Print(MapperTypeMap[header.Mapper])
	fmt.Println(")")
}

func dumpAddress(addr int) {
	fmt.Printf("0x%04X:", addr)
}

func dumpRawData(line *Line, colored bool) {
	data := color.New()
	if colored {
		data = color.New(color.FgYellow)
	}

	for i := 0; i < 4; i++ {
		if i < len(line.Data) {
			data.Printf("%02X ", line.Data[i])
			continue
		}
		fmt.Print("   ")
	}
}

func dumpInstruction(line *Line, colored bool) {
	comment := color.New()
	reg := color.New()
	db := color.New()
	opcode := color.New()
	dollar := color.New()
	hash := color.New()
	bracket := color.New()
	args := color.New()
	if colored {
		comment = color.New(color.FgHiGreen, color.Bold)
		reg = color.New(color.FgHiMagenta)
		db = color.New(color.Bold)
		opcode = color.New(color.FgHiBlue, color.Bold)
		dollar = color.New(color.FgYellow)
		hash = color.New(color.FgHiRed)
		bracket = color.New()
		args = color.New()
	}

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
		args.Printf("%02X", arg)
	case AddressingTypeAbsolute:
		dollar.Print("$")
		args.Printf("%04X", arg)
	case AddressingTypeZeroPage:
		dollar.Print("$")
		args.Printf("%02X", arg)
	case AddressingTypeIndirect:
		bracket.Print("(")
		dollar.Print("$")
		args.Printf("%04X", arg)
		bracket.Print(")")
	case AddressingTypeAbsoluteX:
		dollar.Print("$")
		args.Printf("%04X, ", arg)
		reg.Print("X")
	case AddressingTypeAbsoluteY:
		dollar.Print("$")
		args.Printf("%04X, ", arg)
		reg.Print("Y")
	case AddressingTypeZeroPageX:
		dollar.Print("$")
		args.Printf("%02X, ", arg)
		reg.Print("X")
	case AddressingTypeZeroPageY:
		dollar.Print("$")
		args.Printf("%02X, ", arg)
		reg.Print("Y")
	case AddressingTypeIndirectX:
		bracket.Print("(")
		dollar.Print("$")
		args.Printf("%02X, ", arg)
		reg.Print("X")
		bracket.Print(")")
	case AddressingTypeIndirectY:
		bracket.Print("(")
		dollar.Print("$")
		args.Printf("%02X", arg)
		bracket.Print("), ")
		reg.Print("Y")
	case AddressingTypeRelative:
		dollar.Print("$")
		args.Printf("%04X      ", arg)
		comment.Printf("# to $%04X", (line.Address+2)+int(int8(arg)))
	}
}
