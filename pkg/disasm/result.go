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
	Data        []byte       `json:"data" yaml:"data"`
	Instruction *Instruction `json:"instruction" yaml:"instruction"`
}
