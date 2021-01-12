package disasm

type Result struct {
	Header   *Header    `json:"header" yaml:"header"`
	Sections []*Section `json:"sections" yaml:"sections"`
}

type Section struct {
	Lines            []*Line `json:"lines" yaml:"lines"`
	HasInvalidOpcode bool    `json:"has_invalid_opcode" yaml:"has_invalid_opcode"`
}

type Line struct {
	Address     int          `json:"address" yaml:"address"`
	Data        []int        `json:"data" yaml:"data"`
	Instruction *Instruction `json:"instruction" yaml:"instruction"`
}
