package symtable

import (
	"fmt"
	"sync/atomic"
)

type ScopeSymTable = map[SymName]Symbol

type ScopeName = string

type Scope interface {
	Name() ScopeName
	Define(Symbol)          // 定义符号表
	Resolve(SymName) Symbol // 解析符号表
	Symbols() ScopeSymTable // 获得符号
	Enclosed() Scope        // 所依附的Scope
	Shallow(Symbol)
	Exist(Symbol) bool // 某个符号是否存在
	fmt.Stringer
}

var _ Scope = (*baseScope)(nil)

type baseScope struct {
	name     ScopeName
	enclosed Scope
	table    ScopeSymTable
}

func newBaseScope(name ScopeName, enclosed Scope) *baseScope {
	return &baseScope{
		name:     name,
		enclosed: enclosed,
		table:    make(ScopeSymTable),
	}
}

func Iter(start Scope) func(func(V Scope) bool) {
	fn := func(yield func(Scope) bool) {
		var cur Scope = start
		for {
			if !yield(cur) {
				return
			}
			cur = cur.Enclosed()
			if cur == nil {
				break
			}
		}
	}

	return fn
}

func (s *baseScope) Name() ScopeName {
	return s.name
}

func (s *baseScope) Symbols() ScopeSymTable {
	return s.table
}

func (s *baseScope) Enclosed() Scope {
	return s.enclosed
}

func (s baseScope) Exist(sym Symbol) bool {
	_, ok := s.table[sym.Name()]
	return ok
}

func (s *baseScope) Define(sym Symbol) {
	if _, ok := s.table[sym.Name()]; ok {
		msg := fmt.Sprintf("must not redefined symbol: %s", sym.Name())
		panic(msg)
	}
	s.table[sym.Name()] = sym
}

func (s *baseScope) Shallow(sym Symbol) {
	// 直接覆盖当前作用域的符号
	s.table[sym.Name()] = sym
}

func (s *baseScope) Resolve(name SymName) Symbol {
	if ans, ok := s.table[name]; ok {
		return ans
	}

	// 沿着父scope向上查找解析
	cur := s.Enclosed()
	for cur != nil {
		if r := cur.Resolve(name); r != nil {
			return r
		}
		cur = cur.Enclosed()
	}

	return nil
}

func (s baseScope) String() string {
	return s.Name()
}

type GlobalScope = *globalScopeImpl

var _ Scope = (GlobalScope)(nil)

type globalScopeImpl struct {
	*baseScope
}

func NewGlobalScope(enclosed Scope) GlobalScope {
	// 定义global 符号
	ret := globalScopeImpl{
		baseScope: newBaseScope("GlobalScope", enclosed),
	}
	ret.Define(NewBasicTypeSymbol(Int32.String(), Int32))
	return &ret
}

var _ Scope = (LocalScope)(nil)

type LocalScope = (*localScopeImpl)

type localScopeImpl struct {
	*baseScope
}

var localScopeCounter atomic.Int32

func NewLocalScope(enclosed Scope) LocalScope {
	name := fmt.Sprintf("LocalScope: %d", localScopeCounter.Load())
	localScopeCounter.Add(1)
	ret := localScopeImpl{
		baseScope: newBaseScope(name, enclosed),
	}
	return &ret
}

type FuncScope = *funcScopeImpl

var _ Scope = (FuncScope)(nil)

type funcScopeImpl struct {
	*baseScope
	f         FuncSymbol
	hasReturn bool // scope中是否完成return
}

func NewFuncScope(enclosed Scope, name string) FuncScope {
	return &funcScopeImpl{baseScope: newBaseScope(name, enclosed)}
}

func (s *funcScopeImpl) SetReturn() {
	s.hasReturn = true
}

func (s *funcScopeImpl) HasReturn() bool {
	return s.hasReturn
}

func (s funcScopeImpl) String() string {
	n := fmt.Sprintf("FuncScope: %s", s.Name())
	return n
}

func (s *funcScopeImpl) SetSymbol(f FuncSymbol) {
	s.f = f
}

func (s *funcScopeImpl) GetSymbol() FuncSymbol {
	return s.f
}

func GetFuncScope(start Scope) FuncScope {
	var funcScope FuncScope
	for s := range Iter(start) {
		if _, ok := s.(FuncScope); ok {
			funcScope = s.(FuncScope)
			break
		}
	}
	return funcScope
}
