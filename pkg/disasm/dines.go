package disasm

import "errors"

type Dines struct {
	romFilePath string
}

func NewDines(romFilePath string) (*Dines, error) {
	if romFilePath == "" {
		return nil, errors.New("ROM file path is empty")
	}
	return &Dines{romFilePath: romFilePath}, nil
}
