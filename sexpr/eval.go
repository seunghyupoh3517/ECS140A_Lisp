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


//go test eval_test.go lexer.go sexpr.go  eval.go parser.go
//Use this to run test



var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
	//fmt.Println("1111111111111111111111111111111111111111111111111111111")
	if expr.isNil() {
		// nil expression, no need to evaluate
		return nil, nil
	} else if expr.isAtom() {
		// single atom expression, no need to evaluate		 
	 	return nil, ErrEval		// error return

	//common line 33 - 66 could pass most of the invalid test
	} else if expr.isConsCell() {
		//fmt.Println("222222222222")
		//list expression, need to evaluate the element inside
		if expr.car.atom.literal == "QUOTE"  { 
			//may need iterate through to see if there is any error term
			se, _ := expr.cdr.Eval()
			if se == nil {
				return nil, ErrEval
			}
			return se, nil

		}else if expr.car.atom.literal == "+" { 
			//two situation
			//first is sum use Addup()
			//second is positive number
			return nil, nil

		}else if expr.car.atom.literal == "*" {
			//multiple use Multp()
			return nil, nil

		}else if expr.car.atom.literal == "CAR" {
			var se = expr.Cdr()
			// var temp = se

			// detecting the QUOTE in correct position
			if se != nil && se.isConsCell() {
				if se.Car() != nil && se.Car().isConsCell() {
					se = se.Car()
					if se.Car().isAtom && se.Car().literal == "QUOTE" {
						se = se.Cdr()
						if se.Car().isConsCell {
							se = se.Car().Car()
							return se == nil ? mkNil() : se , nil
						}
					}
				}
			}

			return nil, ErrEval




			// fmt.Println("caaaaaaaaar " + expr.SExprString())

			// fmt.Println("line 56 ----- " + expr.cdr.SExprString())

			// if expr.cdr.isConsCell() {
			// 	fmt.Println("????????????????????????")
			// 	fmt.Println("what are you?", expr.cdr.car.SExprString())

			// 	if expr.cdr.car.isNil() {
			// 		fmt.Println("zhe bu shi tao wa ma?")
			// 	} else if expr.cdr.car.isAtom() {
			// 		fmt.Println("ATOMMMMMMMMMMMMMMMMMM")
			// 	}
			// }

			// if !expr.cdr.isNil() && expr.cdr.car.atom.literal == "QUOTE" {
			// 	// se := 

			// 	fmt.Println("------------ line 59")
			// 	// return expr.cdr, nil
			// 	return nil, nil
			// } else if expr.cdr.isNil() {
			// 	fmt.Println("------------ line 63")
			// 	return nil, nil
			// 	// return expr.cdr, nil
			// }

			// fmt.Println("------------ line 68")
			// // return nil, ErrEval
			// return nil, nil

			// if !expr.cdr.isConsCell() || expr.cdr.isNil() {
			// 	return nil, ErrEval
			// } else {
			// 	se := expr.car
			// 	return se, nil
			// }

		}else if expr.car.atom.literal == "CDR" {
			fmt.Println("cdddddddddr " + expr.SExprString())
			if !expr.cdr.isConsCell() || expr.cdr.isNil() {
				return nil, ErrEval
			} else {
				se := expr. cdr
				return se, nil
			}

		}else if expr.car.atom.literal == "LENGTH" {
			//check if there is anything wrong in cdr
			return nil, nil

		}else if expr.car.atom.literal == "ATOM" {
			//check if there is more than one thing after atom
			return nil, nil

		}else if expr.car.atom.literal == "ZEROP" {
			//check if there any empty 
			return nil, nil

		}else if expr.car.atom.literal == "LISTP" {
			//not sure
			return nil, nil
		}

		return nil, ErrEval
	}
	
	/*
	 I think our parser will reject the single number like line 116 - 122 from eval_test
	 fmt.Println(expr.car.atom.num) can only get 1 from line 10
	*/
	fmt.Println(expr.car.atom.literal) //this will give the leader symbol and it seems correct
	if expr.isNumber(){ //Can not get the number
		fmt.Println("1111111111111111111111111111111111111111111111111111111")
		number := mkNumber(expr.car.atom.num)
		return number, nil
	}

	
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
func (expr *SExpr) Car() *SExpr{ //should be used for recur function
	if expr == nil || expr.atom != nil{
		return nil
	}else{
		return expr.car
	}
}
func (expr *SExpr) Cdr() *SExpr{ //should be used for recur function
	if expr == nil || expr.atom != nil{
		return nil
	}else{
		return expr.cdr
	}
}

func (expr *SExpr) Make_list() *SExpr{ //may not correct
	return mkConsCell(expr.car, expr.cdr)
}



func (expr *SExpr) Check_Num() *SExpr{

	//if expr.car.atom.literal == "+"
	// if the first is + or no symbol, return atom.num
	//if expr.car.atom.literal == "-"
	// if the first is -, then reverse the atom.num
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