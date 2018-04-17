package ast

import (
	"github.com/benchlab/bvmUtils"
	"github.com/benchlab/asteroid/typing"
	"github.com/benchlab/asteroid/util"
)

type Node interface {
	Type() NodeType
	Start() util.Location
	End() util.Location
}

type ExpressionNode interface {
	Node
	ResolvedType() typing.Type
}

type DeclarationNode interface {
	Node
}

type StatementNode interface {
	Node
}

type ScopeNode struct {
	Parent       *ScopeNode
	ValidTypes   []NodeType
	Declarations *bvmUtils.detMap
	Sequence     []Node
	index        int
}

func (n *ScopeNode) Next() Node {
	node := n.Sequence[n.index]
	n.index++
	return node
}

func (n *ScopeNode) AddSequential(node Node) {
	if n.Sequence == nil {
		n.Sequence = make([]Node, 0)
	}
	n.Sequence = append(n.Sequence, node)
}

func (n *ScopeNode) NextDeclaration() Node {
	if n.Declarations == nil {
		return nil
	}
	return n.Declarations.Next().(Node)
}

func (n *ScopeNode) GetDeclaration(key string) Node {
	if n.Declarations == nil {
		return nil
	}
	res := n.Declarations.Get(key)
	if res == nil {
		return nil
	}
	return res.(Node)
}

func (n *ScopeNode) AddDeclaration(key string, node Node) {

	if n.Declarations == nil {
		n.Declarations = new(bvmUtils.detMap)
	}
	n.Declarations.Add(key, node)
}

func (n *ScopeNode) Type() NodeType { return Scope }

func (n *ScopeNode) IsValid(nt NodeType) bool {
	for _, t := range n.ValidTypes {
		if t == nt {
			return true
		}
	}
	return false
}

type FileNode struct {
	name string
}

func (n *FileNode) Type() NodeType { return File }

type PackageNode struct {
	name string
}

func (n *PackageNode) Type() NodeType { return Package }

type ProgramNode struct {
}

func (n *ProgramNode) Type() NodeType { return File }
