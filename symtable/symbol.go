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
	SymFunc            // 泛指function
	SymInt32
	SymVoid
	SymUndefined
)

func (t SymType) String() string {
	switch t {
	case SymInt32:
		return "i32"
	case SymToInfer:
		return "toInfer"
	case SymVoid:
		return "void"
	case SymFunc:
		return "func"
	case SymUndefined:
		return "undefined"
	}
	panic("invalid type")
}

type Symbol interface {
	Name() SymName
	Token() antlr.Token
	Type() SymType
	Infer(exttype SymType)
	fmt.Stringer
}

type BaseSymbol struct {
	name    SymName
	stype   SymType
	mutable bool
	token   antlr.Token
}

var _ Symbol = &BaseSymbol{}

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

func (bs *BaseSymbol) Infer(extType SymType) {
	if bs.stype != SymToInfer {
		panic(fmt.Sprintf("cannot infer type for symbol <%s>: current type is <%s>", bs.name, bs.stype))
	}
	bs.stype = extType
}

var _ Symbol = FuncSymbol{}

type FuncSymbol struct {
	name     SymName
	enclosed Scope
	params   []BaseSymbol
	retType  SymType
	token    antlr.Token
}

func NewFuncSymbol(name SymName, enclosed Scope, params []BaseSymbol, retType SymType, atToken antlr.Token) FuncSymbol {
	return FuncSymbol{
		name:     name,
		enclosed: enclosed,
		params:   params,
		retType:  retType,
		token:    atToken,
	}
}

func (f FuncSymbol) Type() SymType {
	return SymFunc
}

func (f FuncSymbol) Token() antlr.Token {
	return f.token
}

func (f FuncSymbol) Name() SymName {
	return f.name
}

func (f FuncSymbol) RetType() SymType {
	return f.retType
}

func (f FuncSymbol) Params() []BaseSymbol {
	return f.params
}

func (f FuncSymbol) String() string {
	p := make([]string, len(f.params))
	for i, s := range f.params {
		p[i] = s.String()
	}
	s := fmt.Sprintf("%s: func(%s) %s",
		f.Name(), strings.Join(p, ", "), f.retType,
	)
	return s
}

func (f FuncSymbol) Infer(extType SymType) {
	panic(fmt.Sprintf("cannot infer type for function symbol <%s>", f.name))
}

var _ Symbol = BasicTypeSymbol{}

type BasicTypeSymbol struct {
	name  SymName
	stype SymType
}

func NewBasicTypeSymbol(name SymName, stype SymType) BasicTypeSymbol {
	return BasicTypeSymbol{name: name, stype: stype}
}

func (b BasicTypeSymbol) Token() antlr.Token {
	return nil
}

func (b BasicTypeSymbol) Type() SymType {
	return b.stype
}

func (b BasicTypeSymbol) Name() SymName {
	return b.name
}

func (b BasicTypeSymbol) String() string {
	return b.Name()
}

func (b BasicTypeSymbol) Infer(extType SymType) {
	panic(fmt.Sprintf("cannot infer type for basic type symbol <%s>", b.name))
}
