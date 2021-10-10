// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (this *BaseExprVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.VisitWrapper(this)
}

func (ctx *ProgContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitProg(ctx)
}

func (ctx *PrintExprContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitPrintExpr(ctx)
}

func (ctx *AssignContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitAssign(ctx)
}

func (ctx *BlankContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitBlank(ctx)
}

func (ctx *ParensContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitParens(ctx)
}

func (ctx *MulDivContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitMulDiv(ctx)
}

func (ctx *AddSubContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitAddSub(ctx)
}

func (ctx *IdContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitId(ctx)
}

func (ctx *IntContext) VisitWrapper(visitor antlr.Visitor) interface{} {
	return visitor.(*BaseExprVisitor).VisitInt(ctx)
}
