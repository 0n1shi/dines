package disasm

import (
	errs "errors"

	"github.com/pkg/errors"
)

type Dines struct {
}

func NewDines() *Dines {
	return &Dines{}
}

func (dines *Dines) Disassemble(data []byte) (*Result, error) {
	valid := dines.isValid(data)
	if !valid {
		return nil, errs.New("invalid rom")
	}

	header, err := dines.header(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Result{
		Header: header,
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

func (dines *Dines) header(data []byte) (*Header, error) {
	header := &Header{}
	header.ProgBankCount = int(data[4])
	header.CharBankCount = int(data[5])
	header.Mapper = int(data[7]&0xF0) | int((data[6]&0xF0)>>4)
	return header, nil
}

func (dines *Dines) Dump(result *Result) {
}
