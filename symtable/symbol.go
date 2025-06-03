package symtable

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
	"tinygo.org/x/go-llvm"
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
	SymError
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
	case SymError:
		return "errType"
	}
	panic("invalid type")
}

type Symbol interface {
	Name() SymName
	Token() antlr.Token
	Type() SymType
	Infer(SymType)
	SetLLVMValue(llvm.Value)
	LLVMValue() llvm.Value
	fmt.Stringer
}

type BaseSymbol = *baseSymbolImpl

type baseSymbolImpl struct {
	name    SymName
	stype   SymType
	mutable bool
	token   antlr.Token

	llvmVal llvm.Value
}

var _ Symbol = (BaseSymbol)(nil)

func NewBaseSymbol(name SymName, stype SymType, mutable bool, token antlr.Token) BaseSymbol {
	return &baseSymbolImpl{name: name, stype: stype, mutable: mutable, token: token}
}

func (bs *baseSymbolImpl) SetLLVMValue(val llvm.Value) {
	bs.llvmVal = val
}

func (bs *baseSymbolImpl) LLVMValue() llvm.Value {
	if bs.llvmVal.IsNil() {
		panic(fmt.Sprintf("llvm value for symbol <%s> is nil", bs.name))
	}
	return bs.llvmVal
}

func (bs baseSymbolImpl) Token() antlr.Token {
	return bs.token
}

func (bs baseSymbolImpl) Mutable() bool {
	return bs.mutable
}

func (bs baseSymbolImpl) Name() SymName {
	return bs.name
}

func (bs baseSymbolImpl) String() string {
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

func (bs baseSymbolImpl) Type() SymType {
	return bs.stype
}

func (bs *baseSymbolImpl) Infer(t SymType) {
	if bs.stype != SymToInfer {
		log.Panicf("cannot infer type for symbol <%s>: current type is <%s>", bs.name, bs.stype)
	}
	bs.stype = t
}

type FuncSymbol = *funcSymbolImpl

var _ Symbol = (FuncSymbol)(nil)

type funcSymbolImpl struct {
	name     SymName
	enclosed Scope
	params   []BaseSymbol
	retType  SymType
	token    antlr.Token

	llvmFn llvm.Value
}

func NewFuncSymbol(name SymName, enclosed Scope, params []BaseSymbol, retType SymType, atToken antlr.Token) FuncSymbol {
	return &funcSymbolImpl{
		name:     name,
		enclosed: enclosed,
		params:   params,
		retType:  retType,
		token:    atToken,
	}
}

func (f *funcSymbolImpl) SetLLVMValue(val llvm.Value) {
	f.llvmFn = val
}

func (f *funcSymbolImpl) LLVMValue() llvm.Value {
	if f.llvmFn.IsNil() {
		panic(fmt.Sprintf("llvm value for function symbol <%s> is nil", f.name))
	}
	return f.llvmFn
}

func (f funcSymbolImpl) Type() SymType {
	return SymFunc
}

func (f funcSymbolImpl) Token() antlr.Token {
	return f.token
}

func (f funcSymbolImpl) Name() SymName {
	return f.name
}

func (f funcSymbolImpl) RetType() SymType {
	return f.retType
}

func (f funcSymbolImpl) Params() []BaseSymbol {
	return f.params
}

func (f funcSymbolImpl) String() string {
	p := make([]string, len(f.params))
	for i, s := range f.params {
		p[i] = s.String()
	}
	s := fmt.Sprintf("%s: fn(%s) %s",
		f.Name(), strings.Join(p, ", "), f.retType,
	)
	return s
}

func (f funcSymbolImpl) Infer(extType SymType) {
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

func (b BasicTypeSymbol) LLVMValue() llvm.Value {
	return llvm.Value{}
}

func (b BasicTypeSymbol) SetLLVMValue(val llvm.Value) {
	panic(fmt.Sprintf("cannot set llvm value for basic type symbol <%s>", b.name))
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
