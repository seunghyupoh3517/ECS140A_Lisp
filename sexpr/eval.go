package sexpr

import (
	"errors"
	"math/big" // You will need to use this package in your implementation.
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error

var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
	
	if expr.isNil() {
		// nil expression, no need to evaluate
		return nil, nil
	} else if expr.isAtom() {
		// single atom expression, no need to evaluate		 
		return nil, ErrEval		// error return

	} else if expr.isConsCell() {
		// list expression, need to evaluate the element inside
		return nil, nil

	} 
	

	// if (expr.atom == nil && expr.car == nil && expr.cdr == nil){ // 1
	// 	return nil, ErrEval
	// }
	
	// if (expr.atom == nil && expr.car != nil){ // list
	// 	return mkConsCell(expr.car, expr.cdr), ErrEval
	// }
	
	// if (expr.atom != nil && expr.car == nil && expr.cdr == nil){ //single atom 2
	// 	return mkAtom(expr.atom), ErrEval
		
	// }



	// if expr.atom.literal == "+"{ //should detect the symbol first
		
	// 	plus := mkNumber(Check_add(expr.car.atom.num, expr.cdr.atom.num))
	// 	return plus, nil
	// }
	// if expr.atom.literal == "*"{ //should detect the symbol first
	// 	mul := mkNumber(Check_mult(expr.car.atom.num, expr.cdr.atom.num))
	// 	return mul, nil
	// }
	//implement num

	return nil, ErrEval
}



// func (expr *SExpr) Check_atom() *Expr{

// 	return True
// }
// func (expr *SExpr) Cons(car, cdr *SExpr) *SExpr{
// 	return &SExpr{car: expr.car, cdr: expr.cdr}
// }
func Car(expr *SExpr) *SExpr{
	if expr.atom != nil || expr == nil{
		return nil
	}else{
		return expr.car
	}
}
func Cdr(expr *SExpr) *SExpr{
	if expr.atom != nil || expr == nil{
		return nil
	}else{
		return expr.cdr
	}
}

func (expr *SExpr) Make_list() *SExpr{ //may not correct
	return mkConsCell(expr.car, expr.cdr)
}



func (expr *SExpr) Check_Num() *SExpr{

	return nil
}

//https://golang.org/pkg/math/big/?m=all
func Check_add(add1, add2 *big.Int) *big.Int{
	return new(big.Int).Add(add1, add2)
}

func Check_mult(mult1, mult2 *big.Int) *big.Int{
	return new(big.Int).Mul(mult1, mult2)
}
// func ()