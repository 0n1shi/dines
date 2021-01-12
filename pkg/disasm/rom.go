package disasm

var MagicNumber = []byte{0x4E, 0x45, 0x53, 0x1A}

const HeaderSize = 16
const ProgramBankSize = 0x4000
const CharacterBankSize = 0x2000

const ProgramROMStartAt = 0x8000

var MapperTypeMap = map[int]string{
	0: "NROM",
	1: "SxROM, MMC1	",
	2: "UxROM",
	3: "CNROM",
	4: "TxROM, MMC3, MMC6",
	5: "ExROM, MMC5	Contains expansion sound",
	7: "AxROM",
	9: "PxROM, MMC2	",
	10: "FxROM, MMC4",
	11: "Color Dreams",
	13: "CPROM",
	15: "100-in-1 Contra Function 16 Multicart",
	16: "Bandai EPROM (24C02)",
	18: "Jaleco SS8806",
	19: "Namco 163 Contains expansion sound",
	21: "VRC4a, VRC4c",
	22: "VRC2a",
	23: "VRC2b, VRC4e",
	24: "VRC6a Contains expansion sound",
	25: "VRC4b, VRC4d",
	26: "VRC6b Contains expansion sound",
	34: "BNROM, NINA-001",
	64: "RAMBO-1 MMC3 clone with extra features",
	66: "GxROM, MxROM",
	68: "After Burner ROM-based nametables",
	69: "FME-7, Sunsoft 5B The 5B is the FME-7 with expansion sound",
	71: "Camerica/Codemasters Similar to UNROM",
	73: "VRC3",
	74: "Pirate MMC3 derivative	Has both CHR ROM and CHR RAM (2k)",
	75: "VRC1",
	76: "Namco 109 variant",
	79: "NINA-03/NINA-06 It's either 003 or 006, we don't know right now",
	85: "VRC7 Contains expansion sound",
	86: "JALECO-JF-13 ",
	94: "Senjou no Ookami ",
	105: "NES-EVENT	Similar to MMC1",
	113: "NINA-03/NINA-06??	For multicarts including mapper 79 games.",
	118: "TxSROM, MMC3	MMC3 with independent mirroring control",
	119: "TQROM, MMC3	Has both CHR ROM and CHR RAM",
	159: "Bandai EPROM (24C01)",
	166: "SUBOR",
	167: "SUBOR",
	180: "Crazy Climber	Variation of UNROM, fixed first bank at $8000",
	185: "CNROM with protection diodes",
	192: "Pirate MMC3 derivative Has both CHR ROM and CHR RAM (4k)",
	206: "DxROM, Namco 118 / MIMIC-1 Simplified MMC3 predecessor lacking some features",
	210: "Namco 175 and 340	Namco 163 with different mirroring",
	228: "Action 52	",
	232: "Camerica/Codemasters Quattro	Multicarts",
}

type Header struct {
	ProgramBank   *Bank `json:"program_bank" yaml:"program_bank"`
	CharacterBank *Bank `json:"character_bank" yaml:"character_bank"`
	Mapper        int   `json:"mapper_type" yaml:"mapper_type"`
}

type Bank struct {
	Count int `json:"count" yamll:"count"`
	Size  int `json:"size" yamll:"size"`
}
