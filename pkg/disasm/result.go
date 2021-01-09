package disasm

type Result struct {
	Header   *Header    `json:"header" yaml:"header"`
	Sections []*Section `json:"sections" yaml:"sections"`
}

type Section struct {
	Lines []*Line `json:"lines" yaml:"lines"`
}

type Line struct {
}
