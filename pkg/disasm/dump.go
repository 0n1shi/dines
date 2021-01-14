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

// colors for output components
var (
	hdr     *color.Color = color.New(color.FgHiGreen)
	comment *color.Color = color.New(color.FgHiGreen, color.Bold)
	reg     *color.Color = color.New(color.FgHiMagenta)
	db      *color.Color = color.New(color.Bold)
	opcode  *color.Color = color.New(color.FgHiBlue, color.Bold)
	dollar  *color.Color = color.New(color.FgYellow)
	hash    *color.Color = color.New(color.FgHiRed)
	bracket *color.Color = color.New()
	args    *color.Color = color.New()
)

func Dump(result *Result, method DumpMethod, colored bool, max int) {
	switch method {
	case DumpMethodJson:
		j, _ := json.Marshal(result)
		fmt.Printf("%s", j)
	case DumpMethodYaml:
		y, _ := yaml.Marshal(result)
		fmt.Printf("%s", y)
	default:
		dumpNormal(result, colored, max)
	}

}

func dumpNormal(result *Result, colored bool, max int) {
	dumpHeader(result.Header, colored)
	lineCounter := 0
	for _, section := range result.Sections {
		for _, line := range section.Lines {
			dumpAddress(line.Address)
			fmt.Printf("\t")
			dumpRawData(line, colored)
			fmt.Printf("\t")
			dumpInstruction(line, colored)
			fmt.Printf("\n")

			lineCounter++
			if lineCounter == max {
				goto end
			}
		}
		fmt.Printf("\n")
	}

end:
}

func dumpHeader(header *Header, colored bool) {
	if !colored {
		hdr = color.New()
	}
	fmt.Print("magic number: ")
	hdr.Println("NES")
	fmt.Print("program Bank: ")
	hdr.Print(header.ProgramBank.Count)
	bracket.Print(" (")
	hdr.Printf("%d byte", header.ProgramBank.Size)
	bracket.Println(")")
	fmt.Printf("character Bank: ")
	hdr.Print(header.CharacterBank.Count)
	bracket.Print(" (")
	hdr.Printf("%d byte", header.CharacterBank.Size)
	bracket.Println(")")
	fmt.Print("mapper: ")
	hdr.Print(header.Mapper)
	bracket.Print(" (")
	hdr.Print(MapperTypeMap[header.Mapper])
	bracket.Println(")\n")
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
	if !colored {
		comment = color.New()
		reg = color.New()
		db = color.New()
		opcode = color.New()
		dollar = color.New()
		hash = color.New()
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
