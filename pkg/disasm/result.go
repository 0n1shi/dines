package disasm

type Result struct {
	Header   *Header
	Sections []*Section
}

type Section struct {
	Lines []*Line
}

type Line struct {
}
