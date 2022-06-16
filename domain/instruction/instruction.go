package instruction

type Instruction int

const (
	TurnLeft Instruction = iota
	TurnRight
	Advance
)
