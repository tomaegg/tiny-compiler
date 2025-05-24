package symtable

type (
	SymName = string
	SymType = int
)

const (
	SymToInfer SymType = iota
	SymInt32
)

const (
	BasicInt32 SymName = "int32"
)

type Symbol interface {
	Name() SymName
}

type BaseSymbol struct {
	name    SymName
	stype   SymType
	mutable bool
}

var _ Symbol = BaseSymbol{}

func NewBaseSymbol(name SymName, stype SymType, mutable bool) BaseSymbol {
	return BaseSymbol{name: name, stype: stype, mutable: mutable}
}

func (bs BaseSymbol) Mutable() bool {
	return bs.mutable
}

func (bs BaseSymbol) Name() SymName {
	return bs.name
}

func (bs BaseSymbol) Type() SymType {
	return bs.stype
}

var _ Symbol = FuncSymbol{}

type FuncSymbol struct {
	name     SymName
	enclosed Scope
}

func NewFuncSymbol(name SymName, enclosed Scope) FuncSymbol {
	return FuncSymbol{name: name, enclosed: enclosed}
}

func (f FuncSymbol) Name() SymName {
	return f.name
}

var _ Symbol = BasicTypeSymbol{}

type BasicTypeSymbol struct {
	name SymName
}

func NewBasicTypeSymbol(name SymName) BasicTypeSymbol {
	return BasicTypeSymbol{name: name}
}

func (b BasicTypeSymbol) Name() SymName {
	return b.name
}
