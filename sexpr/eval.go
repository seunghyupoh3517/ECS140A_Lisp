package sexpr

import (
	"errors"
	"math/big" // You will need to use this package in your implementation.
	// "fmt"
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error


//go test eval_test.go lexer.go sexpr.go  eval.go parser.go
//Use this to run test



var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
	if expr.isNil() {
		// nil expression, no need to evaluate
		return nil, nil
	} else if expr.isAtom() {
		// single atom expression, no need to evaluate	
		
		if expr.isNumber() {
			return expr, nil
		}

	 	return nil, ErrEval		// error return

	//common line 33 - 66 could pass most of the invalid test
	} else if expr.isConsCell() {
		
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

			// detecting the QUOTE in correct position
			if se != nil && se.isConsCell() {

				if se.Car() != nil && se.Car().isConsCell() && se.Cdr().isNil() {
					se = se.Car()
					if se.Car().isAtom() && se.Car().atom.literal == "QUOTE" {
						se = se.Cdr()
						// unwrap the cons the extract the car in the cons
						if se.Car().isConsCell() {
							se = se.Car().Car()
							// check for the Nil 
							if se == nil {
								return mkNil(), nil
							} else {
								return se, nil
							}
						}
					}
				} else if se.Car().atom.literal == "NIL" {
					// TODO: Double check if we find Nil by using .atom()
					return mkNil(), nil
				}
			}

			return nil, ErrEval

		}else if expr.car.atom.literal == "CDR" {
			var se = expr.Cdr()
			
			// detecting the QUOTE in correct position
			if se != nil && se.isConsCell() {
				if se.Car() != nil && se.Car().isConsCell() && se.Cdr().isNil() {
					// fmt.Println(" ********* Debug info: Line 85")
					se =se.Car()
					if se.Car().isAtom() && se.Car().atom.literal == "QUOTE" {
						// fmt.Println(" ********* Debug info: Line 88")
						se = se.Cdr()
						// unwrap the cons, extract the cdr in the cons
						if se.Car().isConsCell() {
							// fmt.Println(" ********* Debug info: Line 92")
							se = se.Car().Cdr()
							if se == nil {
								return mkNil(), nil
							} else {
								return se, nil
							}
						}

					}
				} else if se.Car().atom.literal == "NIL" {
					return mkNil(), nil
				}
			}

			return nil, ErrEval

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
	// fmt.Println(expr.car.atom.literal) //this will give the leader symbol and it seems correct
	// if expr.isNumber(){ //Can not get the number
	// 	fmt.Println("1111111111111111111111111111111111111111111111111111111")
	// 	number := mkNumber(expr.car.atom.num)
	// 	return number, nil
	// }

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
