package jojivm

//opcodes

var opcodes map[string]uint64 = map[string]uint64{
	"NEWINT":   0,
	"NEWSTR":   1,
	"NEWFLOAT": 2,
	"LOAD":     3,
	"MOV":      4,
	"DEL":      5,
	"SETINT":   6,
	"SETSTR":   7,
	"SETFL":    8,
	"IADD":     9,
	"FLADD":    10,
	"STRADD":   11,
	"ISUB":     12,
	"FLSUB":    13,
	"IMUL":     14,
	"FLMUL":    15,
	"IDIV":     16,
	"FLDIV":    17,
	"SETRW":    18,
	"SETRO":    19,
}
