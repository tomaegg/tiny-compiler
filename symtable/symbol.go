package symtable

import (
	"fmt"
	"strings"
)

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

func GetType(t SymType) string {
	switch t {
	case SymInt32:
		return "int32"
	case SymToInfer:
		return "toInfer"
	}
	panic("invalid type")
}

type Symbol interface {
	Name() SymName
	fmt.Stringer
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

func (bs BaseSymbol) String() string {
	var mut string
	if bs.mutable {
		mut = "(mut)"
	}
	s := fmt.Sprintf(
		"%s%s:%s",
		bs.name, mut, GetType(bs.stype),
	)
	return s
}

func (bs BaseSymbol) Type() SymType {
	return bs.stype
}

var _ Symbol = FuncSymbol{}

type FuncSymbol struct {
	name     SymName
	enclosed Scope
	params   []BaseSymbol
}

func NewFuncSymbol(name SymName, enclosed Scope, params []BaseSymbol) FuncSymbol {
	return FuncSymbol{name: name, enclosed: enclosed, params: params}
}

func (f FuncSymbol) Name() SymName {
	return f.name
}

func (f FuncSymbol) String() string {
	p := make([]string, len(f.params))
	for i, s := range f.params {
		p[i] = s.String()
	}
	s := fmt.Sprintf("%s:func(%s)", f.Name(), strings.Join(p, ", "))
	return s
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

func (b BasicTypeSymbol) String() string {
	return b.Name()
}
