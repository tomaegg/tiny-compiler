package symtable

import (
	"fmt"
	"strconv"
)

type ScopeSymTable = map[SymName]Symbol

type ScopeName = string

type Scope interface {
	Name() ScopeName
	Define(Symbol)          // 定义符号表
	Resolve(SymName) Symbol // 解析符号表
	Symbols() ScopeSymTable // 获得符号
	Enclosed() Scope        // 所依附的Scope
	Exist(Symbol) bool      // 某个符号是否存在
}

var _ Scope = (*BaseScope)(nil)

type BaseScope struct {
	name     ScopeName
	enclosed Scope
	table    ScopeSymTable
}

func NewBaseScope(name ScopeName, enclosed Scope) *BaseScope {
	return &BaseScope{
		name:     name,
		enclosed: enclosed,
		table:    make(ScopeSymTable),
	}
}

func (s *BaseScope) Name() ScopeName {
	return s.name
}

func (s *BaseScope) Symbols() ScopeSymTable {
	return s.table
}

func (s *BaseScope) Enclosed() Scope {
	return s.enclosed
}

func (s BaseScope) Exist(sym Symbol) bool {
	_, ok := s.table[sym.Name()]
	return ok
}

func (s *BaseScope) Define(sym Symbol) {
	if _, ok := s.table[sym.Name()]; ok {
		msg := fmt.Sprintf("must not redefined symbol: %s", sym.Name())
		panic(msg)
	}
	s.table[sym.Name()] = sym
}

func (s *BaseScope) Resolve(name SymName) Symbol {
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

var _ Scope = (*GlobalScope)(nil)

type GlobalScope struct {
	*BaseScope
}

func NewGlobalScope(enclosed Scope) *GlobalScope {
	// 定义global 符号
	ret := GlobalScope{
		BaseScope: NewBaseScope("GlobalScope", enclosed),
	}
	ret.Define(NewBasicTypeSymbol(BasicInt32))
	return &ret
}

type LocalScope struct {
	*BaseScope
}

var localScopeCounter = 0

func NewLocalScope(enclosed Scope) *LocalScope {
	name := fmt.Sprintf("LocalScope_%d", localScopeCounter)
	localScopeCounter++
	ret := LocalScope{
		BaseScope: NewBaseScope(name, enclosed),
	}
	return &ret
}

type FuncScope struct {
	*BaseScope
}

func NewFuncScope(enclosed Scope, name string) *FuncScope {
	return &FuncScope{NewBaseScope(name, enclosed)}
}

func (s FuncScope) Name() string {
	n := fmt.Sprintf("FuncScope: %s", s.BaseScope.Name())
	return strconv.Quote(n)
}
