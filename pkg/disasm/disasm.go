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

	sections, err := disassembleCode(header.ProgramBank.Count, data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sections = mergeDBLines(sections)

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

func disassembleCode(programBankCount int, data []byte) ([]*Section, error) {
	sections := []*Section{}
	section := &Section{}

	for index := HeaderSize; index < HeaderSize+(ProgramBankSize*programBankCount); {
		line := &Line{}
		line.Address = ProgramROMStartAt + index - HeaderSize

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

func mergeDBLines(sections []*Section) []*Section {
	for _, sec := range sections {
		newLines := []*Line{}
		dbLine := &Line{}
		for _, line := range sec.Lines {
			if line.Instruction != nil { // valid opcode
				if len(dbLine.Data) > 0 {
					newLines = append(newLines, dbLine)
					dbLine = &Line{}
				}
				newLines = append(newLines, line)
				continue
			}

			for _, d := range line.Data {
				dbLine.Data = append(dbLine.Data, d)
				if len(dbLine.Data) > 3 {
					newLines = append(newLines, dbLine)
					dbLine = &Line{}
				}
			}
		}
		sec.Lines = newLines
	}
	return sections
}
