package sexpr

import (
	"errors"
	"math/big" // You will need to use this package in your implementation.
	"fmt"
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error

var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
	fmt.Println("1111111111111111111111111111111111111111111111111111111")
	if expr.isNil() {
		// nil expression, no need to evaluate
		return nil, nil
	} else if expr.isAtom() {
		// single atom expression, no need to evaluate		 
	 	return nil, ErrEval		// error return
	} else if expr.isConsCell() {
		fmt.Println("222222222222")
		//list expression, need to evaluate the element inside
		if expr.car.atom.token.literal == "'" || expr.atom.token.literal == "'" { //also tried QUOTE does not work
			return nil, nil
		}else if expr.car.atom.token.literal == "+" || expr.atom.token.literal == "+"{ //check if the symbol is only the symbol or not
			
			return nil, nil
		}else if expr.car.atom.token.literal == "*" || expr.atom.token.literal == "*"{
			return nil, nil
		}
		return nil, ErrEval
	}
// I don't know why that literal can not get the symbo
	

	
	// if (expr.atom == nil && expr.car != nil){ // list
	// 	return mkConsCell(expr.car, expr.cdr), ErrEval
	// }
	
	// if (expr.atom != nil && expr.car == nil && expr.cdr == nil){ //single atom 2
	// 	return mkAtom(expr.atom), ErrEval
		
	// }
	//fmt.Println("1111111111111111111111111111111111111111111111111111111")
// 	if expr.isSymbol() { // check add or mult
// 		//fmt.Println(expr.car.atom.num) //not sure if you guys implement + and *
// 	    fmt.Println("444444444444444444444444444")
// 	   if expr.atom.typ.tokenSymbol == "+"{ //should detect the symbol first
// 		fmt.Println("333333333333333333333")
// 		   add := mkNumber(Addup(expr.car.atom.num, expr.cdr.atom.num))
// 		   //fmt.Println(expr.car.atom.num)
// 		   return add,nil
// 	   }else if expr.atom.toekn.literal == "*"{ //should detect the symbol first
// 		fmt.Println("5555555555555555555")
// 		   mul := mkNumber(Multp(expr.car.atom.num, expr.cdr.atom.num))
// 		   return mul, nil
// 	   }else if expr.atom.token.literal == "Quote"{
// 		fmt.Println("6666666666666666666666")
// 		   return mkNumber(expr.cdr.atom.num), nil
// 	   }
//    }



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
func Addup(add1, add2 *big.Int) *big.Int{
	return new(big.Int).Add(add1, add2)
}

func Multp(mult1, mult2 *big.Int) *big.Int{
	return new(big.Int).Mul(mult1, mult2)
}
// func ()