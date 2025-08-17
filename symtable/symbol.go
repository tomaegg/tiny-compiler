package symtable

import (
	"fmt"
	"strconv"
	"strings"
	"tj-compiler/utils/kv"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

type SymName = string

type SymType interface {
	String() string
	SameWith(SymType) bool
	isSymType()
}

// --- 基础类型 ---
type (
	SymToInfer struct{} // 待推断类型
	SymInt32   struct{} // i32
	SymVoid    struct{} // void
	SymFunc    struct{} // 函数
	SymError   struct{} // 错误类型
)

// --- 复合类型：数组 ---
type SymArray struct {
	ElemType SymType // 数组元素类型（如 SymInt32）
	Length   int32   // 数组长度（如 10）
}

func IsNumberTypes(t ...SymType) bool {
	for _, s := range t {
		if !IsNumberType(s) {
			return false
		}
	}
	return true
}

func IsNumberType(t SymType) bool {
	switch t.(type) {
	case SymInt32:
		return true
	default:
		return false
	}
}

// 实现 SymType 接口
func (t SymToInfer) String() string { return "toInfer" }
func (t SymInt32) String() string   { return "i32" }
func (t SymVoid) String() string    { return "void" }
func (t SymFunc) String() string    { return "func" }
func (t SymError) String() string   { return "errType" }
func (t SymArray) String() string {
	lenStr := "invalid"
	if t.Length >= 0 {
		lenStr = strconv.Itoa(int(t.Length))
	}
	return fmt.Sprintf("[%s;%s]", t.ElemType, lenStr)
}

func (t SymToInfer) SameWith(o SymType) bool { return t == o }
func (t SymVoid) SameWith(o SymType) bool    { return t == o }
func (t SymInt32) SameWith(o SymType) bool   { return t == o }
func (t SymError) SameWith(o SymType) bool   { return t == o }
func (t SymFunc) SameWith(o SymType) bool    { return t == o }

func (t SymArray) SameWith(o SymType) bool {
	other, ok := o.(SymArray)
	if !ok {
		return false
	}
	if t.Length != other.Length {
		return false
	}

	switch lhs := t.ElemType.(type) {
	case SymArray:
		rhs, ok := other.ElemType.(SymArray)
		return ok && lhs.SameWith(rhs)

	default:
		return t.ElemType == other.ElemType
	}
}

func NewSymArray(etype SymType, length int32) SymArray {
	return SymArray{ElemType: etype, Length: length}
}

// 标记方法（防止外部类型实现 SymType）
func (t SymToInfer) isSymType() {}
func (t SymInt32) isSymType()   {}
func (t SymVoid) isSymType()    {}
func (t SymFunc) isSymType()    {}
func (t SymError) isSymType()   {}
func (t SymArray) isSymType()   {}

var (
	ToInfer SymType = SymToInfer{}
	Int32   SymType = SymInt32{}
	Void    SymType = SymVoid{}
	Func    SymType = SymFunc{}
	Error   SymType = SymError{}
)

type Symbol interface {
	kv.KeyMapInterface
	Name() SymName
	Token() antlr.Token
	Type() SymType
	Infer(SymType)
	fmt.Stringer
}

type BaseSymbol = *baseSymbolImpl

type baseSymbolImpl struct {
	*kv.KeyMap
	name    SymName
	stype   SymType
	mutable bool
	token   antlr.Token
}

var _ Symbol = (BaseSymbol)(nil)

func NewBaseSymbol(name SymName, stype SymType, mutable bool, token antlr.Token) BaseSymbol {
	ret := &baseSymbolImpl{name: name, stype: stype, mutable: mutable, token: token}
	ret.KeyMap = kv.NewKeyMap()
	return ret
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
	if t == ToInfer {
		log.Panicf("cannot infer type for symbol <%s>: current type is <%s>", bs.name, bs.stype)
	}
	bs.stype = t
}

type FuncSymbol = *funcSymbolImpl

var _ Symbol = (FuncSymbol)(nil)

type funcSymbolImpl struct {
	*kv.KeyMap
	name     SymName
	enclosed Scope
	params   []BaseSymbol
	retType  SymType
	token    antlr.Token
}

func NewFuncSymbol(name SymName, enclosed Scope, params []BaseSymbol, retType SymType, atToken antlr.Token) FuncSymbol {
	ret := &funcSymbolImpl{
		name:     name,
		enclosed: enclosed,
		params:   params,
		retType:  retType,
		token:    atToken,
	}
	ret.KeyMap = kv.NewKeyMap()
	return ret
}

func (f funcSymbolImpl) Type() SymType {
	return Func
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
	*kv.KeyMap
	name  SymName
	stype SymType
}

func NewBasicTypeSymbol(name SymName, stype SymType) BasicTypeSymbol {
	ret := BasicTypeSymbol{name: name, stype: stype}
	ret.KeyMap = kv.NewKeyMap()
	return ret
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

type ArraySymbol struct {
	baseSymbolImpl
}
