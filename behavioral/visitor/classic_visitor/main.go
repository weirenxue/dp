package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (de *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(de)
}

type AdditionExpression struct {
	left, right Expression
}

func (de *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(de)
}

type ExpressionPrinter struct {
	sb *strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{sb: &strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	ep.sb.WriteString("(")
	ae.left.Accept(ep)
	ep.sb.WriteString("+")
	ae.right.Accept(ep)
	ep.sb.WriteString(")")
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(de *DoubleExpression) {
	ee.result = de.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(ae *AdditionExpression) {
	ae.left.Accept(ee)
	result := ee.result
	ae.right.Accept(ee)
	ee.result += result
}

func main() {
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{4},
			right: &DoubleExpression{3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s=%g\n", ep, ee.result)
}
