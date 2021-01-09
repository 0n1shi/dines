package disasm

import (
	"encoding/json"
	errs "errors"
	"fmt"

	"github.com/pkg/errors"
)

type Dines struct {
}

func (dines *Dines) Disassemble(data []byte) (*Result, error) {
	valid := dines.isValid(data)
	if !valid {
		return nil, errs.New("invalid rom")
	}

	header, err := dines.disassembleHeader(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sections, err := dines.disassembleCode(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Result{
		Header:   header,
		Sections: sections,
	}, nil
}

func (dines *Dines) isValid(data []byte) bool {
	if len(data) < 4 {
		return false
	}

	for i := 0; i < len(MagicNumber); i++ {
		if data[i] != MagicNumber[i] {
			return false
		}
	}

	return true
}

func (dines *Dines) disassembleHeader(data []byte) (*Header, error) {
	header := &Header{}
	header.ProgBankCount = int(data[4])
	header.CharBankCount = int(data[5])
	header.Mapper = int(data[7]&0xF0) | int((data[6]&0xF0)>>4)
	return header, nil
}

func (dines *Dines) disassembleCode(data []byte) ([]*Section, error) {
	sections := []*Section{}
	section := &Section{}

	for index := HeaderSize; index < len(data); {
		line := &Line{}
		opcode := data[index] // opcode byte

		ins, ok := InstructionMap[int(opcode)] // get opcode info

		// invalid opcode, just store a data (byte)
		if !ok {
			line.Data = append(line.Data, opcode)
			index++
			section.Lines = append(section.Lines, line)
			section.HasInvalidOpcode = true
			continue
		}

		// store the instruction and the binary
		line.Instruction = &ins
		for i := index; i < index+line.Instruction.Bytes; i++ {
			line.Data = append(line.Data, data[i])
		}
		section.Lines = append(section.Lines, line)
		index += ins.Bytes

		// if find jmp or rts and so on, it may be end of subroutinue
		if IsEndOfSubRoutinue(ins.OpcodeType) {
			sections = append(sections, section)
			section = &Section{}
		}
	}

	sections = append(sections, section)

	return sections, nil
}

func (dines *Dines) Dump(result *Result) {
	d, _ := json.Marshal(result)
	fmt.Print(string(d))
}
