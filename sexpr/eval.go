package sexpr

import (
	"errors"
	//"math/big" // You will need to use this package in your implementation.
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error

var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
	//panic("TODO: implement Eval")
	
	
	expr.mkAtom("1")

	expr.isNil()
	if atom 
	
	return nil
}

func Car(e *SExpr) *SExpr {
	if e.atom == nil || isAtom == nil{
		return nil
	}
	else {
		return e.car
	}
}



func (expr *SExpr) Check_atom() *Expr{

	return True
}

func (expr *SExpr) Check_atom() *Expr{

	return True
}

func (expr *SExpr) Check_Num() *Expr{

	return True
}

func (expr *SExpr) Check_add() *Expr{

}
// func ()