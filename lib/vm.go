package ngaro

const (
	Nop = iota
	Lit
	Dup
	Drop
	Swap
	Push
	Pop
	Loop
	Jump
	Return
	GtJump
	LtJump
	NeJump
	EqJump
	Fetch
	Store
	Add
	Substract
	Multiply
	DivMod
	And
	Or
	Xor
	Shl
	Shr
	ZeroExit
	Inc
	Dec
	In
	Out
	Wait

	dstackSize = 128
	astackSize = 1024
)

// ngaro virtual machine
type VM struct {
	ports  []int
	memory []int
	dstack [dstackSize]int
	astack [astackSize]int
}

func NewVM() *VM {
	return &VM{}
}

func (vm *VM) Run() {

}
