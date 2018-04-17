package ast

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestAddDeclaration(t *testing.T) {
	n := &ClassDeclarationNode{}
	s := ScopeNode{}
	s.AddDeclaration("hi", n)
	bvmUtil.Assert(t, s.Declarations.Length() == 1, "wrong length")
}

func TestGetDeclaration(t *testing.T) {
	n := &ClassDeclarationNode{}
	s := ScopeNode{}
	s.AddDeclaration("hi", n)
	bvmUtil.Assert(t, s.Declarations.Length() == 1, "wrong length")
	x := s.GetDeclaration("hi")
	c := x.(*ClassDeclarationNode)
	c.Identifier = "Sending..."
	bvmUtil.Assert(t, n.Identifier == "Sending...", "message not received")
}
