package symtable

import (
	"bytes"
	"fmt"
	"html"
)

type SymTableGraph struct {
	nodes []string
	edges map[string]string
}

func NewSymTableGraph() *SymTableGraph {
	return &SymTableGraph{
		edges: make(map[string]string),
	}
}

func Quote(name string) string {
	return `"` + name + `"`
}

// ToScopeDot 生成单个作用域的DOT节点（HTML表格格式）
func (g SymTableGraph) ToScopeDot(scope Scope) []byte {
	if len(scope.Symbols()) == 0 {
		return []byte(Quote(scope.Name()))
	}

	const attrEachLine = 2
	var buffer bytes.Buffer
	i := 0
	count := len(scope.Symbols())
	for _, v := range scope.Symbols() {
		if i%attrEachLine == 0 { // 每行开始
			buffer.WriteString("<TR>")
		}
		elem := fmt.Appendf(nil, "<TD>%s</TD>", html.EscapeString(v.String()))
		buffer.Write(elem)
		if i%attrEachLine == attrEachLine-1 || i == count-1 {
			buffer.WriteString("</TR>") // 每行结束或遍历完成时闭合
		}
		i++
	}

	ret := fmt.Appendf(nil,
		`"%s" [label = <<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0">
			<TR><TD COLSPAN="%d">%s</TD></TR>
			%s
		</TABLE>>]`,
		scope.Name(),
		len(scope.Symbols()),
		scope,
		buffer.Bytes(),
	)

	return ret
}

func (g *SymTableGraph) AddNode(node string) {
	g.nodes = append(g.nodes, node)
}

func (g *SymTableGraph) AddEdge(child, parent string) {
	g.edges[Quote(child)] = Quote(parent)
	// 注意符号表图的结构，child -> parent而不是parent -> child
}

func (g *SymTableGraph) ToDot() []byte {
	var buf bytes.Buffer

	const declares = `
digraph G {
	rankdir = BT
	ranksep = 0.5
	edge [arrowsize = 0.5]
	node [shape = none]
`
	buf.WriteString(declares)
	buf.WriteString("\t" + `// node define` + "\n")
	for _, node := range g.nodes {
		buf.Write(fmt.Appendf(nil, "\t%s;\n", (node)))
	}

	buf.WriteString("\t" + `// edge define` + "\n")
	for child, parent := range g.edges {
		buf.Write(fmt.Appendf(nil, "\t%s -> %s;\n", child, parent))
	}

	buf.WriteString("}\n")
	return buf.Bytes()
}
