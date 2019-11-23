package cpu

type AddressingMode int
type Instruction int

// Addressing modes
const (
	INVALID AddressingMode = iota
	ABS                    // Absolute
	ABSX                   // Absolute (Indexed)
	ABSY                   // Absolute (Indexed)
	IND                    // Indirect
	INDX                   // INdirect (Indexed)
	INDY                   // Indirect (Indexed)
	ZPG                    // Zero Page
	ZPGX                   // Zero Page (Indexed)
	ZPGY                   // Zero Page (Indexed)
	ACC                    // Accumulator
	IMM                    // Immediate
	IMPL                   // Implied
	REL                    // Relative
)

// Instruction mnemonics
const (
	ADC Instruction = iota
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
}

func NewOpcode(bytes uint8, instr Instruction, mode AddressingMode, cycles uint8) Opcode {
	return Opcode{
		bytes:       bytes,
		instruction: instr,
		mode:        mode,
		cycles:      cycles,
	}
}

var OpcodeTable = [256]Opcode{
	NewOpcode(0x00, BRK, IMPL, 0),
}
