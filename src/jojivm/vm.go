package jojivm

//Actual VM implementation, also contains implementation of registers, etc.

type Register struct {
	Value            *JojiObject
	ValueHeapAddress uint64
	IsWritable       bool
}

type Instruction struct {
	Opcode uint64
	Args   [1]string
}

type VM struct {
	Heap               [1]JojiObject
	InstructionPointer uint64
	Code               [1]Instruction
	//registers
	Reg1  Register
	Reg2  Register
	Reg3  Register
	Reg4  Register
	Reg5  Register
	Reg6  Register
	Reg7  Register
	Reg8  Register
	Reg9  Register
	Reg10 Register
}
