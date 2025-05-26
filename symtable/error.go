package symtable

type SemanticErrType string

type SemanticErr struct {
	Err SemanticErrType
	Msg string
}

func NewTypeError(msg string) SemanticErr {
	return SemanticErr{
		Err: TypeErr,
		Msg: msg,
	}
}

const (
	TypeErr SemanticErrType = "TypeErr"
)
