package ngaro

type Cell int

const (
	OpNop Cell = iota
	OpLit
	OpDup
	OpDrop
	OpSwap
	OpPush
	OpPop
	OpLoop
	OpJump
	OpReturn
	OpGtJump
	OpLtJump
	OpNeJump
	OpEqJump
	OpFetch
	OpStore
	OpAdd
	OpSubstract
	OpMultiply
	OpDivMod
	OpAnd
	OpOr
	OpXor
	OpShl
	OpShr
	OpZeroExit
	OpInc
	OpDec
	OpIn
	OpOut
	OpWait

	dstackSize = 128
	rstackSize = 1024
	msize      = 1048576
)

// ngaro virtual machine
type VM struct {
	ports  []Cell
	memory []Cell
	msize  int // memory size
	ip     int // instruction pointer
	dstack *stack
	rstack *stack
}

func (vm *VM) wait() {

}

func NewVM() *VM {
	return &VM{
		ports:  make([]Cell, 16),
		memory: make([]Cell, msize),
		msize:  msize,
		ip:     0,
		dstack: newStack(dstackSize),
		rstack: newStack(rstackSize),
	}
}

func (vm *VM) Run() {

}
