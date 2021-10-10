// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

var memory map[string]int

func init() {
	memory = make(map[string]int)
}

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by ExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by ExprParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by ExprParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by ExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by ExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ExprParser#int.
	VisitInt(ctx *IntContext) interface{}
}

// VisitProg is called when the visitor visits Prog
func (v *BaseExprVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.Visit(ctx.Stat(0))
}

// VisitPrintExpr is called when the visitor visits PrintExpr
func (v *BaseExprVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr())
	fmt.Println(value)
	return value
}

// VisitAssign is called when the visitor visits Assign
func (v *BaseExprVisitor) VisitAssign(ctx *AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(int)
	memory[id] = value
	return value
}

// VisitBlank is called when the visitor visits Blank
func (v *BaseExprVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

// VisitParens is called when the visitor visits Parens
func (v *BaseExprVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

// VisitMulDiv is called when the visitor visits MulDiv
func (v *BaseExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)
	if ctx.op.GetTokenType() == ExprLexerMUL {
		return left * right
	}
	return left  / right
}

// VisitAddSub is called when the visitor visits AddSub
func (v *BaseExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)
	if ctx.op.GetTokenType() == ExprLexerADD {
		return left + right
	}
	return left - right
}

// VisitId is called when the visitor visits Id
func (v *BaseExprVisitor) VisitId(ctx *IdContext) interface{} {
	id := ctx.ID().GetText()
	if value, ok := memory[id]; ok {
		return value
	}
	return 0
}

// VisitInt is called when the visitor visits Int
func (v *BaseExprVisitor) VisitInt(ctx *IntContext) interface{} {
	value, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		panic(err)
	}
	return value
}
