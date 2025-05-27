package symtable

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type (
	SymName = string
	SymType int
)

const (
	SymToInfer SymType = iota
	SymInt32
	SymVoid
)

func (t SymType) String() string {
	switch t {
	case SymInt32:
		return "i32"
	case SymToInfer:
		return "toInfer"
	case SymVoid:
		return "void"
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
	token   antlr.Token
}

var _ Symbol = BaseSymbol{}

func NewBaseSymbol(name SymName, stype SymType, mutable bool, token antlr.Token) BaseSymbol {
	return BaseSymbol{name: name, stype: stype, mutable: mutable, token: token}
}

func (bs BaseSymbol) Token() antlr.Token {
	return bs.token
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
		"<%s%s,%s>",
		bs.name, mut, bs.stype,
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
	retType  SymType
}

func NewFuncSymbol(name SymName, enclosed Scope, params []BaseSymbol, retType SymType) FuncSymbol {
	return FuncSymbol{name: name, enclosed: enclosed, params: params, retType: retType}
}

func (f FuncSymbol) Name() SymName {
	return f.name
}

func (f FuncSymbol) RetType() SymType {
	return f.retType
}

func (f FuncSymbol) String() string {
	p := make([]string, len(f.params))
	for i, s := range f.params {
		p[i] = s.String()
	}
	s := fmt.Sprintf("%s:func(%s) %s",
		f.Name(), strings.Join(p, ", "), f.retType,
	)
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
