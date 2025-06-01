package symtable

import (
	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

type SymNode struct {
	Scope  Scope               // current value, Scope.Enclosed point to parent Scope
	Childs map[string]*SymNode // children scopes
}

func NewSymNode(s Scope) *SymNode {
	return &SymNode{
		Scope:  s,
		Childs: make(map[string]*SymNode),
	}
}

func (n *SymNode) Parent() Scope {
	return n.Scope.Enclosed()
}

func (n *SymNode) Child(name string) *SymNode {
	return n.Childs[name]
}

func (s *SymNode) AddChild(child *SymNode) {
	id := child.Scope.Name()
	if _, ok := s.Childs[id]; ok {
		log.Panicf("scope %s already exists in %s", id, s.Scope.Name())
		return
	}
	s.Childs[id] = child
}

type SymTable struct {
	tree   antlr.ParseTree
	root   Scope
	scopes map[string]*SymNode
}

func NewSymTable(root Scope) *SymTable {
	ret := &SymTable{
		root:   root,
		scopes: make(map[string]*SymNode),
	}
	ret.scopes[root.String()] = NewSymNode(root)
	return ret
}

func (s *SymTable) SetTree(tree antlr.ParseTree) {
	s.tree = tree
}

func (s *SymTable) Tree() antlr.ParseTree {
	return s.tree
}

func (s *SymTable) ToDotGraph() []byte {
	g := NewSymTableGraph()

	for _, node := range s.scopes {
		g.AddNode(node.Scope)
	}

	// 这是一棵树，因此每次只需加入node即可，不会存在重复node
	var dfs func(*SymNode)
	dfs = func(root *SymNode) {
		if root == nil {
			return
		}

		for _, child := range root.Childs {
			g.AddEdge(child.Scope, root.Scope)
			dfs(child)
		}
	}

	dfs(s.Root())

	return g.ToDot()
}

func (s *SymTable) Scope(name string) Scope {
	return s.Node(name).Scope
}

func (s *SymTable) Node(name string) *SymNode {
	return s.scopes[name]
}

func (s *SymTable) GlobalScope() Scope {
	return s.Root().Scope
}

func (s *SymTable) Root() *SymNode {
	return s.Node(s.root.Name())
}

func (s *SymTable) addNode(n Scope) {
	id := n.Name()
	if _, ok := s.scopes[id]; ok {
		// 已经存在
		return
	}
	s.scopes[id] = NewSymNode(n)
}

func (s *SymTable) AddEdge(child, parent Scope) {
	s.addNode(child)  // prepare child
	s.addNode(parent) // prepare parent

	pid := parent.Name()
	cid := child.Name()
	p := s.scopes[pid]
	c := s.scopes[cid]

	p.AddChild(c)
}
