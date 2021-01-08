package disasm

import "errors"

type Dines struct {
}

func NewDines() (*Dines, error) {
	return &Dines{}, nil
}

func (dines *Dines) Disassemble(data []byte) (*Result, error) {
	valid := dines.InValid(data)
	if !valid {
		return nil, errors.New("invalid rom")
	}
	return &Result{}, nil
}

func (dines *Dines) InValid(data []byte) bool {
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

func (dines *Dines) Dump(result *Result) {
}
