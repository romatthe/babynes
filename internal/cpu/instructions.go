package cpu

type AddressingMode int
type Instruction int

// Addressing modes
const (
	ABS AddressingMode = iota // Absolute
	ABX                       // Absolute (Indexed)
	ABY                       // Absolute (Indexed)
	IND                       // Indirect
	IDX                       // Indirect (Indexed)
	IDY                       // Indirect (Indexed)
	ZPG                       // Zero Page
	ZPX                       // Zero Page (Indexed)
	ZPY                       // Zero Page (Indexed)
	ACC                       // Accumulator
	IMM                       // Immediate
	IMP                       // Implied
	REL                       // Relative
)

// Instruction mnemonics
const (
	INVALID Instruction = iota
	ADC
	AND
	ASL
	BCC
	BCS
	BEQ
	BIT
	BMI
	BNE
	BPL
	BRK
	BVC
	BVS
	CLC
	CLD
	CLI
	CLV
	CMP
	CPX
	CPY
	DEC
	DEX
	DEY
	EOR
	INC
	INX
	INY
	JMP
	JSR
	LDA
	LDX
	LDY
	LSR
	NOP
	ORA
	PHA
	PHP
	PLA
	PLP
	ROL
	ROR
	RTI
	RTS
	SBC
	SEC
	SED
	SEI
	STA
	STX
	STY
	TAX
	TAY
	TSX
	TXA
	TXS
	TYA
)

// Opcode and accompanying data
type Opcode struct {
	bytes       uint8          // raw opcode byte
	instruction Instruction    // Instruction
	mode        AddressingMode // Addressing Mode
	cycles      uint8          // Base cycles
	pageCross   bool
}

func NewOpcode(bytes uint8, instr Instruction, mode AddressingMode, cycles uint8, pageCross bool) Opcode {
	return Opcode{
		bytes:       bytes,
		instruction: instr,
		mode:        mode,
		cycles:      cycles,
		pageCross:   pageCross,
	}
}

func BadOpcode(bytes uint8) Opcode {
	return Opcode{
		bytes:       bytes,
		instruction: INVALID,
		mode:        IMM,
		cycles:      0,
		pageCross:   false,
	}
}

var OpcodeTable = [256]Opcode{
	NewOpcode(0x00, BRK, IMP, 0, false),
	NewOpcode(0x01, ORA, IDX, 6, false),
	BadOpcode(0x02),
	BadOpcode(0x03),
	BadOpcode(0x04),
	NewOpcode(0x05, ORA, ZPG, 3, false),
	NewOpcode(0x06, ASL, ZPG, 5, false),
	BadOpcode(0x07),
	NewOpcode(0x08, PHP, IMP, 3, false),
	NewOpcode(0x09, ORA, IMM, 2, false),
	NewOpcode(0x0A, ASL, ACC, 2, false),
	BadOpcode(0x0B),
	BadOpcode(0x0C),
	NewOpcode(0x0D, ORA, ABS, 4, false),
	NewOpcode(0x0E, ASL, ABS, 6, false),
	BadOpcode(0x0F),
	NewOpcode(0x10, BPL, REL, 2, false),
	NewOpcode(0x11, ORA, IDY, 5, true),
	BadOpcode(0x12),
	BadOpcode(0x13),
	BadOpcode(0x14),
	NewOpcode(0x15, ORA, ZPX, 4, false),
	NewOpcode(0x16, ASL, ZPX, 6, false),
	BadOpcode(0x17),
	NewOpcode(0x18, CLC, IMP, 2, false),
	NewOpcode(0x19, ORA, ABY, 4, true),
	BadOpcode(0x1A),
	BadOpcode(0x1B),
	BadOpcode(0x1C),
	NewOpcode(0x1D, ORA, ABX, 4, true),
	NewOpcode(0x1E, ASL, ABX, 7, false),
	BadOpcode(0x1F),
	NewOpcode(0x20, JSR, ABS, 6, false),
	NewOpcode(0x21, AND, IDX, 6, false),
	BadOpcode(0x22),
	BadOpcode(0x23),
	NewOpcode(0x24, BIT, ZPG, 3, false),
	NewOpcode(0x25, AND, ZPG, 3, false),
	NewOpcode(0x26, ROL, ZPG, 5, false),
	BadOpcode(0x27),
	NewOpcode(0x28, PLP, IMP, 4, false),
	NewOpcode(0x29, AND, IMM, 2, false),
	NewOpcode(0x2A, ROL, ACC, 2, false),
	BadOpcode(0x2B),
	NewOpcode(0x2C, BIT, ABS, 4, false),
	NewOpcode(0x2D, AND, ABS, 4, false),
	NewOpcode(0x2E, ROL, ABS, 6, false),
	BadOpcode(0x2F),
	NewOpcode(0x30, BMI, REL, 2, false),
	NewOpcode(0x31, AND, IDY, 5, true),
	BadOpcode(0x32),
	BadOpcode(0x33),
	BadOpcode(0x34),
	NewOpcode(0x35, AND, ZPX, 4, false),
	NewOpcode(0x36, ROL, ZPX, 6, false),
	BadOpcode(0x37),
	NewOpcode(0x38, SEC, IMP, 2, false),
	NewOpcode(0x39, AND, ABY, 4, true),
	BadOpcode(0x3A),
	BadOpcode(0x3B),
	BadOpcode(0x3C),
	NewOpcode(0x3D, AND, ABX, 4, true),
	NewOpcode(0x3E, ROL, ABX, 7, false),
	BadOpcode(0x3F),
	NewOpcode(0x40, RTI, IMP, 6, false),
	NewOpcode(0x41, EOR, IDX, 6, false),
	BadOpcode(0x42),
	BadOpcode(0x43),
	BadOpcode(0x44),
	NewOpcode(0x45, EOR, ZPG, 3, false),
	NewOpcode(0x46, LSR, ZPG, 5, false),
	BadOpcode(0x47),
	NewOpcode(0x48, PHA, IMP, 3, false),
	NewOpcode(0x49, EOR, IMM, 2, false),
	NewOpcode(0x4A, LSR, ACC, 2, false),
	BadOpcode(0x4B),
	NewOpcode(0x4C, JMP, ABS, 3, false),
	NewOpcode(0x4D, EOR, ABS, 4, false),
	NewOpcode(0x4E, LSR, ABS, 6, false),
	BadOpcode(0x4F),
	NewOpcode(0x50, BVC, REL, 2, false),
	NewOpcode(0x51, EOR, IDY, 5, true),
	BadOpcode(0x52),
	BadOpcode(0x53),
	BadOpcode(0x54),
	NewOpcode(0x55, EOR, ZPX, 4, false),
	NewOpcode(0x56, LSR, ZPX, 6, false),
	BadOpcode(0x57),
	NewOpcode(0x58, CLI, IMP, 2, false),
	NewOpcode(0x59, EOR, ABY, 4, true),
	BadOpcode(0x5A),
	BadOpcode(0x5B),
	BadOpcode(0x5C),
	NewOpcode(0x5D, EOR, ABX, 4, true),
	NewOpcode(0x5E, LSR, ABX, 7, false),
	BadOpcode(0x5F),
	NewOpcode(0x60, RTS, IMP, 6, false),
	NewOpcode(0x61, ADC, IDX, 6, false),
	BadOpcode(0x62),
	BadOpcode(0x63),
	BadOpcode(0x64),
	NewOpcode(0x65, ADC, ZPG, 3, false),
	NewOpcode(0x66, ROR, ZPG, 5, false),
	BadOpcode(0x67),
	NewOpcode(0x68, PLA, IMP, 4, false),
	NewOpcode(0x69, ADC, IMM, 2, false),
	NewOpcode(0x6A, ROR, ACC, 2, false),
	BadOpcode(0x6B),
	NewOpcode(0x6C, JMP, IND, 5, false),
	NewOpcode(0x6D, ADC, ABS, 4, false),
	NewOpcode(0x6E, ROR, ABS, 6, false),
	BadOpcode(0x6F),
	NewOpcode(0x70, BVS, REL, 2, false),
	NewOpcode(0x71, ADC, IDY, 5, true),
	BadOpcode(0x72),
	BadOpcode(0x73),
	BadOpcode(0x74),
	NewOpcode(0x75, ADC, ZPX, 4, false),
	NewOpcode(0x76, ROR, ZPX, 6, false),
	BadOpcode(0x77),
	NewOpcode(0x78, SEI, IMP, 2, false),
	NewOpcode(0x79, ADC, ABY, 4, true),
	BadOpcode(0x7A),
	BadOpcode(0x7B),
	BadOpcode(0x7C),
	NewOpcode(0x7D, ADC, ABX, 4, true),
	NewOpcode(0x7E, ROR, ABX, 7, false),
	BadOpcode(0x7F),
	BadOpcode(0x80),
	NewOpcode(0x81, STA, IDX, 6, false),
	BadOpcode(0x82),
	BadOpcode(0x83),
	NewOpcode(0x84, STY, ZPG, 3, false),
	NewOpcode(0x85, STA, ZPG, 3, false),
	NewOpcode(0x86, STX, ZPG, 3, false),
	BadOpcode(0x87),
	NewOpcode(0x88, DEY, IMP, 2, false),
	BadOpcode(0x89),
	NewOpcode(0x8A, TXA, IMP, 2, false),
	BadOpcode(0x8B),
	NewOpcode(0x8C, STY, ABS, 4, false),
	NewOpcode(0x8D, STA, ABS, 4, false),
	NewOpcode(0x8E, STX, ABS, 4, false),
	BadOpcode(0x8F),
	NewOpcode(0x90, BCC, REL, 2, false),
	NewOpcode(0x91, STA, IDY, 6, false),
	BadOpcode(0x92),
	BadOpcode(0x93),
	NewOpcode(0x94, STY, ZPX, 4, false),
	NewOpcode(0x95, STA, ZPX, 4, false),
	NewOpcode(0x96, STX, ZPY, 4, false),
	BadOpcode(0x97),
	NewOpcode(0x98, TYA, IMP, 2, false),
	NewOpcode(0x99, STA, ABY, 5, false),
	NewOpcode(0x9A, TXS, IMP, 2, false),
	BadOpcode(0x9B),
	BadOpcode(0x9C),
	NewOpcode(0x9D, STA, ABX, 5, false),
	BadOpcode(0x9E),
	BadOpcode(0x9F),
	NewOpcode(0xA0, LDY, IMM, 2, false),
	NewOpcode(0xA1, LDA, IDX, 6, false),
	NewOpcode(0xA2, LDX, IMM, 2, false),
	BadOpcode(0xA3),
	NewOpcode(0xA4, LDY, ZPG, 3, false),
	NewOpcode(0xA5, LDA, ZPG, 3, false),
	NewOpcode(0xA6, LDX, ZPG, 3, false),
	BadOpcode(0xA7),
	NewOpcode(0xA8, TAY, IMP, 2, false),
	NewOpcode(0xA9, LDA, IMM, 2, false),
	NewOpcode(0xAA, TAX, IMP, 2, false),
	BadOpcode(0xAB),
	NewOpcode(0xAC, LDY, ABS, 4, false),
	NewOpcode(0xAD, LDA, ABS, 4, false),
	NewOpcode(0xAE, LDX, ABS, 4, false),
	BadOpcode(0xAF),
	NewOpcode(0xB0, BCS, REL, 2, false),
	NewOpcode(0xB1, LDA, IDY, 5, true),
	BadOpcode(0xB2),
	BadOpcode(0xB3),
	NewOpcode(0xB4, LDY, ZPX, 4, false),
	NewOpcode(0xB5, LDA, ZPX, 4, false),
	NewOpcode(0xB6, LDX, ZPY, 4, false),
	BadOpcode(0xB7),
	NewOpcode(0xB8, CLV, IMP, 2, false),
	NewOpcode(0xB9, LDA, ABY, 4, true),
	NewOpcode(0xBA, TSX, IMP, 2, false),
	BadOpcode(0xBB),
	NewOpcode(0xBC, LDY, ABX, 4, true),
	NewOpcode(0xBD, LDA, ABX, 4, true),
	NewOpcode(0xBE, LDX, ABY, 4, true),
	BadOpcode(0xBF),
	NewOpcode(0xC0, CPY, IMM, 2, false),
	NewOpcode(0xC1, CMP, IDX, 6, false),
	BadOpcode(0xC2),
	BadOpcode(0xC3),
	NewOpcode(0xC4, CPY, ZPG, 3, false),
	NewOpcode(0xC5, CMP, ZPG, 3, false),
	NewOpcode(0xC6, DEC, ZPG, 5, false),
	BadOpcode(0xC7),
	NewOpcode(0xC8, INY, IMP, 2, false),
	NewOpcode(0xC9, CMP, IMM, 2, false),
	NewOpcode(0xCA, DEX, IMP, 2, false),
	BadOpcode(0xCB),
	NewOpcode(0xCC, CPY, ABS, 4, false),
	NewOpcode(0xCD, CMP, ABS, 4, false),
	NewOpcode(0xCE, DEC, ABS, 6, false),
	BadOpcode(0xCF),
	NewOpcode(0xD0, BNE, REL, 2, false),
	NewOpcode(0xD1, CMP, IDY, 5, true),
	BadOpcode(0xD2),
	BadOpcode(0xD3),
	BadOpcode(0xD4),
	NewOpcode(0xD5, CMP, ZPX, 4, false),
	NewOpcode(0xD6, DEC, ZPX, 6, false),
	BadOpcode(0xD7),
	NewOpcode(0xD8, CLD, IMP, 2, false),
	NewOpcode(0xD9, CMP, ABY, 4, true),
	BadOpcode(0xDA),
	BadOpcode(0xDB),
	BadOpcode(0xDC),
	NewOpcode(0xDD, CMP, ABX, 4, true),
	NewOpcode(0xDE, DEC, ABX, 7, false),
	BadOpcode(0xDF),
	NewOpcode(0xE0, CPX, IMM, 2, false),
	NewOpcode(0xE1, SBC, IDX, 6, false),
	BadOpcode(0xE2),
	BadOpcode(0xE3),
	NewOpcode(0xE4, CPX, ZPG, 3, false),
	NewOpcode(0xE5, SBC, ZPG, 3, false),
	NewOpcode(0xE6, INC, ZPG, 5, false),
	BadOpcode(0xE7),
	NewOpcode(0xE8, INX, IMP, 2, false),
	NewOpcode(0xE9, SBC, IMM, 2, false),
	NewOpcode(0xEA, NOP, IMP, 2, false),
	BadOpcode(0xEB),
	NewOpcode(0xEC, CPX, ABS, 4, false),
	NewOpcode(0xED, SBC, ABS, 4, false),
	NewOpcode(0xEE, INC, ABS, 6, false),
	BadOpcode(0xEF),
	NewOpcode(0xF0, BEQ, REL, 2, false),
	NewOpcode(0xF1, SBC, IDY, 5, true),
	BadOpcode(0xF2),
	BadOpcode(0xF3),
	BadOpcode(0xF4),
	NewOpcode(0xF5, SBC, ZPX, 4, false),
	NewOpcode(0xF6, INC, ZPX, 6, false),
	BadOpcode(0xF7),
	NewOpcode(0xF8, SED, IMP, 2, false),
	NewOpcode(0xF9, SBC, ABY, 4, true),
	BadOpcode(0xFA),
	BadOpcode(0xFB),
	BadOpcode(0xFC),
	NewOpcode(0xFD, SBC, ABX, 4, true),
	NewOpcode(0xFE, INC, ABX, 7, false),
	BadOpcode(0xFF),
}
