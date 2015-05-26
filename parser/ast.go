package parser

import (
	"fmt"
)

type Node interface {
	String() string
}

type Stat interface {
	Node
	statNode()
}

type Expr interface {
	Node
	exprNode()
}

type Decl interface {
	Node
	declNode()
}

type Variable struct {
	Type    Type
	Name    string
	Mutable bool
}

func (v *Variable) String() string {
	return "[Variable " + v.Name + "]" // TODO mut, type
}

//
// Nodes
//

/**
 * Declarations
 */

type VariableDecl struct {
	Variable   *Variable
	Assignment Expr
}

func (v *VariableDecl) declNode() {}

func (v *VariableDecl) String() string {
	if v.Assignment == nil {
		return "[VariableDecl " + v.Variable.String() + "]"
	} else {
		return "[VariableDecl " + v.Variable.String() +
			" = " + v.Assignment.String() + "]"
	}
}

/**
 * Expressions
 */

type RuneLiteral struct {
	Value rune
}

func (v *RuneLiteral) exprNode() {}

func (v *RuneLiteral) String() string {
	return fmt.Sprintf("[RuneLiteral: %c]", v.Value)
}

type IntegerLiteral struct {
	Value uint64
}

func (v *IntegerLiteral) exprNode() {}

func (v *IntegerLiteral) String() string {
	return fmt.Sprintf("[IntegerLiteral: %d]", v.Value)
}

type FloatingLiteral struct {
	Value float64
}

func (v *FloatingLiteral) exprNode() {}

func (v *FloatingLiteral) String() string {
	return fmt.Sprintf("[FloatingLiteral: %f]", v.Value)
}

type StringLiteral struct {
	Value string
}

func (v *StringLiteral) exprNode() {}

func (v *StringLiteral) String() string {
	return "[StringLiteral: " + v.Value + "]"
}

type BinaryExpr struct {
	Lhand, Rhand Expr
	Op           BinOpType
}

func (v *BinaryExpr) exprNode() {}

func (v *BinaryExpr) String() string {
	return "[BinaryLiteral: " + v.Lhand.String() + " " +
		v.Op.String() + " "	+
		v.Rhand.String() + "]"
}

