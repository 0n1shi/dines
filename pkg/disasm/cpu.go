package disasm

type AddressingType int

const (
	AddressingTypeImplied = AddressingType(iota)
	AddressingTypeAccumulator
	AddressingTypeImmediate
	AddressingTypeZeroPage
	AddressingTypeZeroPageX
	AddressingTypeZeroPageY
	AddressingTypeRelative
	AddressingTypeAbsolute
	AddressingTypeAbsoluteX
	AddressingTypeAbsoluteY
	AddressingTypeIndirect
	AddressingTypeIndirectX
	AddressingTypeIndirectY
)

type OpcodeType int

const (
	OpcodeLDA = OpcodeType(iota)
	OpcodeLDX
	OpcodeLDY
	OpcodeSTA
	OpcodeSTX
	OpcodeSTY
	OpcodeTXA
	OpcodeTYA
	OpcodeTXS
	OpcodeTAY
	OpcodeTAX
	OpcodeTSX
	OpcodePHA
	OpcodePHP
	OpcodePLA
	OpcodePLP
	OpcodeADC
	OpcodeSBC
	OpcodeCPX
	OpcodeCPY
	OpcodeCMP
	OpcodeAND
	OpcodeEOR
	OpcodeORA
	OpcodeBIT
	OpcodeASL
	OpcodeLSR
	OpcodeROL
	OpcodeROR
	OpcodeINX
	OpcodeINY
	OpcodeINC
	OpcodeDEX
	OpcodeDEY
	OpcodeDEC
	OpcodeCLC
	OpcodeCLI
	OpcodeCLV
	OpcodeCLD
	OpcodeSEC
	OpcodeSEI
	OpcodeSED
	OpcodeJSR
	OpcodeJMP
	OpcodeRTI
	OpcodeRTS
	OpcodeBCC
	OpcodeBCS
	OpcodeBEQ
	OpcodeBMI
	OpcodeBNE
	OpcodeBPL
	OpcodeBVC
	OpcodeBVS
	OpcodeNOP
	OpcodeBRK
)

var OpecodeMap = map[OpcodeType]string{
	OpcodeLDA: "lda",
	OpcodeLDX: "ldx",
	OpcodeLDY: "ldy",
	OpcodeSTA: "sta",
	OpcodeSTX: "stx",
	OpcodeSTY: "sty",
	OpcodeTXA: "txa",
	OpcodeTYA: "tya",
	OpcodeTXS: "txs",
	OpcodeTAY: "tay",
	OpcodeTAX: "tax",
	OpcodeTSX: "tsx",
	OpcodePHA: "pha",
	OpcodePHP: "php",
	OpcodePLA: "pla",
	OpcodePLP: "plp",
	OpcodeADC: "adc",
	OpcodeSBC: "sbc",
	OpcodeCPX: "cpx",
	OpcodeCPY: "cpy",
	OpcodeCMP: "cmp",
	OpcodeAND: "and",
	OpcodeEOR: "eor",
	OpcodeORA: "ora",
	OpcodeBIT: "bit",
	OpcodeASL: "asl",
	OpcodeLSR: "lsr",
	OpcodeROL: "rol",
	OpcodeROR: "ror",
	OpcodeINX: "inx",
	OpcodeINY: "iny",
	OpcodeINC: "inc",
	OpcodeDEX: "dex",
	OpcodeDEY: "dey",
	OpcodeDEC: "dec",
	OpcodeCLC: "clc",
	OpcodeCLI: "cli",
	OpcodeCLV: "clv",
	OpcodeCLD: "cld",
	OpcodeSEC: "sec",
	OpcodeSEI: "sei",
	OpcodeSED: "sed",
	OpcodeJSR: "jsr",
	OpcodeJMP: "jmp",
	OpcodeRTI: "rti",
	OpcodeRTS: "rts",
	OpcodeBCC: "bcc",
	OpcodeBCS: "bcs",
	OpcodeBEQ: "beq",
	OpcodeBMI: "bmi",
	OpcodeBNE: "bne",
	OpcodeBPL: "bpl",
	OpcodeBVC: "bvc",
	OpcodeBVS: "bvs",
	OpcodeNOP: "nop",
	OpcodeBRK: "brk",
}

var SectionDividers = []OpcodeType{OpcodeJMP, OpcodeRTS, OpcodeRTI}

func IsEndOfSubRoutinue(opcodeType OpcodeType) bool {
	for _, d := range SectionDividers {
		if opcodeType == d {
			return true
		}
	}
	return false
}

type Instruction struct {
	OpcodeType     OpcodeType
	AddressingType AddressingType
	Bytes          int
	Cycle          int
}

var InstructionMap = map[int]Instruction{
	0x00: {
		OpcodeType:     OpcodeBRK,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          7,
	},
	0x01: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0x05: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x06: {
		OpcodeType:     OpcodeASL,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          5,
	},
	0x08: {
		OpcodeType:     OpcodePHP,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          3,
	},
	0x09: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0x0A: {
		OpcodeType:     OpcodeASL,
		AddressingType: AddressingTypeAccumulator,
		Bytes:          1,
		Cycle:          2,
	},
	0x0D: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x0E: {
		OpcodeType:     OpcodeASL,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0x10: {
		OpcodeType:     OpcodeBPL,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0x11: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          5,
	},
	0x15: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0x16: {
		OpcodeType:     OpcodeASL,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          6,
	},
	0x18: {
		OpcodeType:     OpcodeCLC,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x19: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0x1D: {
		OpcodeType:     OpcodeORA,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0x1E: {
		OpcodeType:     OpcodeASL,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          7,
	},
	0x20: {
		OpcodeType:     OpcodeJSR,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0x21: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0x24: {
		OpcodeType:     OpcodeBIT,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x25: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x26: {
		OpcodeType:     OpcodeROL,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          5,
	},
	0x28: {
		OpcodeType:     OpcodePLP,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          4,
	},
	0x29: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0x2A: {
		OpcodeType:     OpcodeROL,
		AddressingType: AddressingTypeAccumulator,
		Bytes:          1,
		Cycle:          2,
	},
	0x2C: {
		OpcodeType:     OpcodeBIT,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x2D: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x2E: {
		OpcodeType:     OpcodeROL,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0x30: {
		OpcodeType:     OpcodeBMI,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0x31: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          5,
	},
	0x35: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0x36: {
		OpcodeType:     OpcodeROL,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          6,
	},
	0x38: {
		OpcodeType:     OpcodeSEC,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x39: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0x3D: {
		OpcodeType:     OpcodeAND,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0x3E: {
		OpcodeType:     OpcodeROL,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          7,
	},
	0x40: {
		OpcodeType:     OpcodeRTI,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          6,
	},
	0x41: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0x45: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x46: {
		OpcodeType:     OpcodeLSR,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          5,
	},
	0x48: {
		OpcodeType:     OpcodePHA,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          3,
	},
	0x49: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          3,
	},
	0x4A: {
		OpcodeType:     OpcodeLSR,
		AddressingType: AddressingTypeAccumulator,
		Bytes:          1,
		Cycle:          2,
	},
	0x4C: {
		OpcodeType:     OpcodeJMP,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          3,
	},
	0x4D: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x4E: {
		OpcodeType:     OpcodeLSR,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0x50: {
		OpcodeType:     OpcodeBVC,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0x51: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          5,
	},
	0x55: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0x56: {
		OpcodeType:     OpcodeLSR,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          6,
	},
	0x58: {
		OpcodeType:     OpcodeCLI,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x59: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0x5D: {
		OpcodeType:     OpcodeEOR,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0x5E: {
		OpcodeType:     OpcodeLSR,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          7,
	},
	0x60: {
		OpcodeType:     OpcodeRTS,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          6,
	},
	0x61: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0x65: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x66: {
		OpcodeType:     OpcodeROR,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          5,
	},
	0x68: {
		OpcodeType:     OpcodePLA,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          4,
	},
	0x69: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0x6A: {
		OpcodeType:     OpcodeROR,
		AddressingType: AddressingTypeAccumulator,
		Bytes:          1,
		Cycle:          2,
	},
	0x6C: {
		OpcodeType:     OpcodeJMP,
		AddressingType: AddressingTypeIndirect,
		Bytes:          3,
		Cycle:          5,
	},
	0x6D: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x6E: {
		OpcodeType:     OpcodeROR,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0x70: {
		OpcodeType:     OpcodeBVS,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0x71: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          2,
		Cycle:          5,
	},
	0x75: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0x76: {
		OpcodeType:     OpcodeROR,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          6,
	},
	0x78: {
		OpcodeType:     OpcodeSEI,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x79: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0x7D: {
		OpcodeType:     OpcodeADC,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0x7E: {
		OpcodeType:     OpcodeROR,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          7,
	},
	0x81: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0x84: {
		OpcodeType:     OpcodeSTY,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x85: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x86: {
		OpcodeType:     OpcodeSTX,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0x88: {
		OpcodeType:     OpcodeDEY,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x8A: {
		OpcodeType:     OpcodeTXA,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x8C: {
		OpcodeType:     OpcodeSTY,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x8D: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x8E: {
		OpcodeType:     OpcodeSTX,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0x90: {
		OpcodeType:     OpcodeBCC,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0x91: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          6,
	},
	0x94: {
		OpcodeType:     OpcodeSTY,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0x95: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0x96: {
		OpcodeType:     OpcodeSTX,
		AddressingType: AddressingTypeZeroPageY,
		Bytes:          2,
		Cycle:          4,
	},
	0x98: {
		OpcodeType:     OpcodeTYA,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x99: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          5,
	},
	0x9A: {
		OpcodeType:     OpcodeTXS,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0x9D: {
		OpcodeType:     OpcodeSTA,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          5,
	},
	0xA0: {
		OpcodeType:     OpcodeLDY,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xA1: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0xA2: {
		OpcodeType:     OpcodeLDX,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xA4: {
		OpcodeType:     OpcodeLDY,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xA5: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xA6: {
		OpcodeType:     OpcodeLDX,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xA8: {
		OpcodeType:     OpcodeTAY,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xA9: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xAA: {
		OpcodeType:     OpcodeTAX,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xAC: {
		OpcodeType:     OpcodeLDY,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xAD: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xAE: {
		OpcodeType:     OpcodeLDX,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xB0: {
		OpcodeType:     OpcodeBCS,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0xB1: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          5,
	},
	0xB4: {
		OpcodeType:     OpcodeLDY,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0xB5: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0xB6: {
		OpcodeType:     OpcodeLDX,
		AddressingType: AddressingTypeZeroPageY,
		Bytes:          2,
		Cycle:          4,
	},
	0xB8: {
		OpcodeType:     OpcodeCLV,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xB9: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0xBA: {
		OpcodeType:     OpcodeTSX,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xBC: {
		OpcodeType:     OpcodeLDY,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xBD: {
		OpcodeType:     OpcodeLDA,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0xBE: {
		OpcodeType:     OpcodeLDX,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0xC0: {
		OpcodeType:     OpcodeCPY,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xC1: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0xC4: {
		OpcodeType:     OpcodeCPY,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xC5: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xC6: {
		OpcodeType:     OpcodeDEC,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          5,
	},
	0xC8: {
		OpcodeType:     OpcodeINY,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xC9: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xCA: {
		OpcodeType:     OpcodeDEX,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xCC: {
		OpcodeType:     OpcodeCPY,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xCD: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xCE: {
		OpcodeType:     OpcodeDEC,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0xD0: {
		OpcodeType:     OpcodeBNE,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0xD1: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          5,
	},
	0xD5: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0xD6: {
		OpcodeType:     OpcodeDEC,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          6,
	},
	0xD8: {
		OpcodeType:     OpcodeCLD,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xD9: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0xDD: {
		OpcodeType:     OpcodeCMP,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0xDE: {
		OpcodeType:     OpcodeDEC,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          7,
	},
	0xE0: {
		OpcodeType:     OpcodeCPX,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xE1: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeIndirectX,
		Bytes:          2,
		Cycle:          6,
	},
	0xE4: {
		OpcodeType:     OpcodeCPX,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xE5: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          3,
	},
	0xE6: {
		OpcodeType:     OpcodeINC,
		AddressingType: AddressingTypeZeroPage,
		Bytes:          2,
		Cycle:          5,
	},
	0xE8: {
		OpcodeType:     OpcodeINX,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xE9: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeImmediate,
		Bytes:          2,
		Cycle:          2,
	},
	0xEA: {
		OpcodeType:     OpcodeNOP,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xEC: {
		OpcodeType:     OpcodeCPX,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xED: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          4,
	},
	0xEE: {
		OpcodeType:     OpcodeINC,
		AddressingType: AddressingTypeAbsolute,
		Bytes:          3,
		Cycle:          6,
	},
	0xF0: {
		OpcodeType:     OpcodeBEQ,
		AddressingType: AddressingTypeRelative,
		Bytes:          2,
		Cycle:          2,
	},
	0xF1: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeIndirectY,
		Bytes:          2,
		Cycle:          5,
	},
	0xF5: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          4,
	},
	0xF6: {
		OpcodeType:     OpcodeINC,
		AddressingType: AddressingTypeZeroPageX,
		Bytes:          2,
		Cycle:          6,
	},
	0xF8: {
		OpcodeType:     OpcodeSED,
		AddressingType: AddressingTypeImplied,
		Bytes:          1,
		Cycle:          2,
	},
	0xF9: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeAbsoluteY,
		Bytes:          3,
		Cycle:          4,
	},
	0xFD: {
		OpcodeType:     OpcodeSBC,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          4,
	},
	0xFE: {
		OpcodeType:     OpcodeINC,
		AddressingType: AddressingTypeAbsoluteX,
		Bytes:          3,
		Cycle:          7,
	},
}
