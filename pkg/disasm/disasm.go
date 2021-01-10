package disasm

import (
	errs "errors"

	"github.com/pkg/errors"
)

func Disassemble(data []byte) (*Result, error) {
	valid := isValidROM(data)
	if !valid {
		return nil, errs.New("invalid rom")
	}

	header, err := disassembleHeader(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sections, err := disassembleCode(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	//sections = validateSections(sections) // not perfect, may kill valid opcodes

	return &Result{
		Header:   header,
		Sections: sections,
	}, nil
}

func isValidROM(data []byte) bool {
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

func disassembleHeader(data []byte) (*Header, error) {
	header := &Header{}
	header.ProgramBank = &Bank{
		Count: int(data[4]),
		Size:  int(data[4]) * ProgramBankSize,
	}
	header.CharacterBank = &Bank{
		Count: int(data[5]),
		Size:  int(data[5]) * ProgramBankSize,
	}
	header.Mapper = int(data[7]&0xF0) | int((data[6]&0xF0)>>4)
	return header, nil
}

func disassembleCode(data []byte) ([]*Section, error) {
	sections := []*Section{}
	section := &Section{}

	for index := HeaderSize; index < len(data); {
		line := &Line{}
		opcode := data[index] // opcode byte

		ins, ok := InstructionMap[int(opcode)] // get opcode info

		// invalid opcode, just store a data (byte)
		if !ok {
			line.Data = append(line.Data, int(opcode))
			index++
			section.Lines = append(section.Lines, line)
			section.HasInvalidOpcode = true
			continue
		}

		// store the instruction and the binary
		line.Instruction = &ins
		for i := index; i < index+line.Instruction.Bytes; i++ {
			line.Data = append(line.Data, int(data[i]))
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

func validateSections(sections []*Section) []*Section {
	for _, section := range sections {
		if !section.HasInvalidOpcode {
			// all opcodes are valid
			continue
		}

		// convert all into "db" opcode
		newLines := []*Line{}
		newLine := &Line{}
		for _, line := range section.Lines {
			for _, d := range line.Data {
				newLine.Data = append(newLine.Data, d)
				if len(newLine.Data) > 3 {
					newLines = append(newLines, newLine)
					newLine = &Line{}
				}
			}
		}
		section.Lines = newLines
	}

	return sections
}
