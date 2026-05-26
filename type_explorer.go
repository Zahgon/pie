package main

import (
	"go/ast"
)

type TypeExplorer struct {
	TypeName    string
	Methods     []string
	IsInterface bool
}

func NewTypeExplorer(pkg *ast.Package, typeName string) *TypeExplorer {
	_ = "STUB: not implemented"
	return nil
}

func (explorer *TypeExplorer) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func (explorer *TypeExplorer) HasEquals() bool { _ = "STUB: not implemented"; return false }

func (explorer *TypeExplorer) HasString() bool { _ = "STUB: not implemented"; return false }

func (explorer *TypeExplorer) HasMethod(lookingFor string) bool {
	_ = "STUB: not implemented"
	return false
}
