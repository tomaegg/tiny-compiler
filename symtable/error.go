package symtable

import "fmt"

type SemanticErrType string

type BasicSemanticErr struct {
	Err SemanticErrType
	Msg string
}

type SemanticErr interface {
	Message(string, ...any) SemanticErr
	fmt.Stringer
}

func (e *BasicSemanticErr) Message(format string, a ...any) SemanticErr {
	msg := fmt.Sprintf(format, a...)
	e.Msg = msg
	return e
}

func (e *BasicSemanticErr) String() string {
	return fmt.Sprintf("%s: %s", e.Err, e.Msg)
}

func NewSematicErr(errType SemanticErrType) SemanticErr {
	return &BasicSemanticErr{
		Err: errType,
	}
}

const (
	TypeErr        SemanticErrType = "TypeError"
	IntOverflowErr SemanticErrType = "IntOverflowError"
	ArgsErr        SemanticErrType = "ArgsError"
	RetErr         SemanticErrType = "RetError"
	BreakErr       SemanticErrType = "BreakErr"
	ContiErr       SemanticErrType = "ContinueErr"
)
