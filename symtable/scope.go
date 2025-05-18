package symtable

import "fmt"

type SymTable = map[SymName]Symbol

type ScopeName = string

type Scope interface {
	Name() ScopeName
	Define(Symbol)          // 定义符号表
	Resolve(SymName) Symbol // 解析符号表
	Symbols() SymTable      // 获得符号
	Enclosed() Scope        // 所依附的Scope
}

var _ Scope = (*BaseScope)(nil)

type BaseScope struct {
	name     ScopeName
	enclosed Scope
	table    SymTable
}

func NewBaseScope(name ScopeName, enclosed Scope) *BaseScope {
	return &BaseScope{
		name:     name,
		enclosed: enclosed,
		table:    make(SymTable),
	}
}

func (s *BaseScope) Name() ScopeName {
	return s.name
}

func (s *BaseScope) Symbols() SymTable {
	return s.table
}

func (s *BaseScope) Enclosed() Scope {
	return s.enclosed
}

func (s *BaseScope) Define(sym Symbol) {
	if _, ok := s.table[sym.Name()]; ok {
		panic("must not redefined symbol")
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
	enclosed.Define(NewBasicTypeSymbol(SymInt32))
	ret := GlobalScope{
		BaseScope: NewBaseScope("GlobalScope", enclosed),
	}
	return &ret
}

type LocalScope struct {
	*BaseScope
}

var localScopeCounter = 0

func NewLocalScope(enclosed Scope) *LocalScope {
	name := fmt.Sprintf("LocalScope:%d", localScopeCounter)
	localScopeCounter++
	ret := LocalScope{
		BaseScope: NewBaseScope(name, enclosed),
	}
	return &ret
}
