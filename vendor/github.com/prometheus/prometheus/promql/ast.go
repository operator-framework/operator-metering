// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this ***REMOVED***le except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the speci***REMOVED***c language governing permissions and
// limitations under the License.

package promql

import (
	"fmt"
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/storage"
)

// Node is a generic interface for all nodes in an AST.
//
// Whenever numerous nodes are listed such as in a switch-case statement
// or a chain of function de***REMOVED***nitions (e.g. String(), expr(), etc.) convention is
// to list them as follows:
//
// 	- Statements
// 	- statement types (alphabetical)
// 	- ...
// 	- Expressions
// 	- expression types (alphabetical)
// 	- ...
//
type Node interface {
	// String representation of the node that returns the given node when parsed
	// as part of a valid query.
	String() string
}

// Statement is a generic interface for all statements.
type Statement interface {
	Node

	// stmt ensures that no other type accidentally implements the interface
	stmt()
}

// Statements is a list of statement nodes that implements Node.
type Statements []Statement

// AlertStmt represents an added alert rule.
type AlertStmt struct {
	Name        string
	Expr        Expr
	Duration    time.Duration
	Labels      labels.Labels
	Annotations labels.Labels
}

// EvalStmt holds an expression and information on the range it should
// be evaluated on.
type EvalStmt struct {
	Expr Expr // Expression to be evaluated.

	// The time boundaries for the evaluation. If Start equals End an instant
	// is evaluated.
	Start, End time.Time
	// Time between two evaluated instants for the range [Start:End].
	Interval time.Duration
}

// RecordStmt represents an added recording rule.
type RecordStmt struct {
	Name   string
	Expr   Expr
	Labels labels.Labels
}

func (*AlertStmt) stmt()  {}
func (*EvalStmt) stmt()   {}
func (*RecordStmt) stmt() {}

// Expr is a generic interface for all expression types.
type Expr interface {
	Node

	// Type returns the type the expression evaluates to. It does not perform
	// in-depth checks as this is done at parsing-time.
	Type() ValueType
	// expr ensures that no other types accidentally implement the interface.
	expr()
}

// Expressions is a list of expression nodes that implements Node.
type Expressions []Expr

// AggregateExpr represents an aggregation operation on a Vector.
type AggregateExpr struct {
	Op       itemType // The used aggregation operation.
	Expr     Expr     // The Vector expression over which is aggregated.
	Param    Expr     // Parameter used by some aggregators.
	Grouping []string // The labels by which to group the Vector.
	Without  bool     // Whether to drop the given labels rather than keep them.
}

// BinaryExpr represents a binary expression between two child expressions.
type BinaryExpr struct {
	Op       itemType // The operation of the expression.
	LHS, RHS Expr     // The operands on the respective sides of the operator.

	// The matching behavior for the operation if both operands are Vectors.
	// If they are not this ***REMOVED***eld is nil.
	VectorMatching *VectorMatching

	// If a comparison operator, return 0/1 rather than ***REMOVED***ltering.
	ReturnBool bool
}

// Call represents a function call.
type Call struct {
	Func *Function   // The function that was called.
	Args Expressions // Arguments used in the call.
}

// MatrixSelector represents a Matrix selection.
type MatrixSelector struct {
	Name          string
	Range         time.Duration
	Offset        time.Duration
	LabelMatchers []*labels.Matcher

	// The series iterators are populated at query preparation time.
	series    []storage.Series
	iterators []*storage.BufferedSeriesIterator
}

// NumberLiteral represents a number.
type NumberLiteral struct {
	Val float64
}

// ParenExpr wraps an expression so it cannot be disassembled as a consequence
// of operator precedence.
type ParenExpr struct {
	Expr Expr
}

// StringLiteral represents a string.
type StringLiteral struct {
	Val string
}

// UnaryExpr represents a unary operation on another expression.
// Currently unary operations are only supported for Scalars.
type UnaryExpr struct {
	Op   itemType
	Expr Expr
}

// VectorSelector represents a Vector selection.
type VectorSelector struct {
	Name          string
	Offset        time.Duration
	LabelMatchers []*labels.Matcher

	// The series iterators are populated at query preparation time.
	series    []storage.Series
	iterators []*storage.BufferedSeriesIterator
}

func (e *AggregateExpr) Type() ValueType  { return ValueTypeVector }
func (e *Call) Type() ValueType           { return e.Func.ReturnType }
func (e *MatrixSelector) Type() ValueType { return ValueTypeMatrix }
func (e *NumberLiteral) Type() ValueType  { return ValueTypeScalar }
func (e *ParenExpr) Type() ValueType      { return e.Expr.Type() }
func (e *StringLiteral) Type() ValueType  { return ValueTypeString }
func (e *UnaryExpr) Type() ValueType      { return e.Expr.Type() }
func (e *VectorSelector) Type() ValueType { return ValueTypeVector }
func (e *BinaryExpr) Type() ValueType {
	if e.LHS.Type() == ValueTypeScalar && e.RHS.Type() == ValueTypeScalar {
		return ValueTypeScalar
	}
	return ValueTypeVector
}

func (*AggregateExpr) expr()  {}
func (*BinaryExpr) expr()     {}
func (*Call) expr()           {}
func (*MatrixSelector) expr() {}
func (*NumberLiteral) expr()  {}
func (*ParenExpr) expr()      {}
func (*StringLiteral) expr()  {}
func (*UnaryExpr) expr()      {}
func (*VectorSelector) expr() {}

// VectorMatchCardinality describes the cardinality relationship
// of two Vectors in a binary operation.
type VectorMatchCardinality int

const (
	CardOneToOne VectorMatchCardinality = iota
	CardManyToOne
	CardOneToMany
	CardManyToMany
)

func (vmc VectorMatchCardinality) String() string {
	switch vmc {
	case CardOneToOne:
		return "one-to-one"
	case CardManyToOne:
		return "many-to-one"
	case CardOneToMany:
		return "one-to-many"
	case CardManyToMany:
		return "many-to-many"
	}
	panic("promql.VectorMatchCardinality.String: unknown match cardinality")
}

// VectorMatching describes how elements from two Vectors in a binary
// operation are supposed to be matched.
type VectorMatching struct {
	// The cardinality of the two Vectors.
	Card VectorMatchCardinality
	// MatchingLabels contains the labels which de***REMOVED***ne equality of a pair of
	// elements from the Vectors.
	MatchingLabels []string
	// On includes the given label names from matching,
	// rather than excluding them.
	On bool
	// Include contains additional labels that should be included in
	// the result from the side with the lower cardinality.
	Include []string
}

// Visitor allows visiting a Node and its child nodes. The Visit method is
// invoked for each node encountered by Walk. If the result visitor w is not
// nil, Walk visits each of the children of node with the visitor w, followed
// by a call of w.Visit(nil).
type Visitor interface {
	Visit(node Node) (w Visitor)
}

// Walk traverses an AST in depth-***REMOVED***rst order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	switch n := node.(type) {
	case Statements:
		for _, s := range n {
			Walk(v, s)
		}
	case *AlertStmt:
		Walk(v, n.Expr)

	case *EvalStmt:
		Walk(v, n.Expr)

	case *RecordStmt:
		Walk(v, n.Expr)

	case Expressions:
		for _, e := range n {
			Walk(v, e)
		}
	case *AggregateExpr:
		Walk(v, n.Expr)

	case *BinaryExpr:
		Walk(v, n.LHS)
		Walk(v, n.RHS)

	case *Call:
		Walk(v, n.Args)

	case *ParenExpr:
		Walk(v, n.Expr)

	case *UnaryExpr:
		Walk(v, n.Expr)

	case *MatrixSelector, *NumberLiteral, *StringLiteral, *VectorSelector:
		// nothing to do

	default:
		panic(fmt.Errorf("promql.Walk: unhandled node type %T", node))
	}

	v.Visit(nil)
}

type inspector func(Node) bool

func (f inspector) Visit(node Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}

// Inspect traverses an AST in depth-***REMOVED***rst order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// for all the non-nil children of node, recursively.
func Inspect(node Node, f func(Node) bool) {
	Walk(inspector(f), node)
}
