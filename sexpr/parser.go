package sexpr

import (
	"errors"
	// "fmt"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <sexpr>       ::= <atom> | <pars> | QUOTE <sexpr>
// <atom>        ::= NUMBER | SYMBOL
// <pars>        ::= LPAR <dotted_list> RPAR | LPAR <proper_list> RPAR
// <dotted_list> ::= <proper_list> <sexpr> DOT <sexpr>
// <proper_list> ::= <sexpr> <proper_list> | \epsilon
//
type Parser interface {
	Parse(string) (*SExpr, error)
}

// Grammar struct to implements the Parser interface
type Grammar struct {
	grammar map[*SExpr][]*SExpr    // TODO: need to change the attributes inside struct
}

// TODO: need to change the Gramar constructor
func NewParser() Parser {
	var parseGrammar Parser = Grammar{make(map[*SExpr][]*SExpr)}
	return parseGrammar
}

var tokenInd = 0
var tokenList []*token
var meetLpar = false
var quoteFlag = false

// <Start_NT> ::= <Sexpr>
// <Sexpr> ::= number | symbol | ( <List> )
// <List> ::= <Sexpr> <New> | epsilon
// <New> ::= . <Sexpr> | <Sexpr> <New> | epsilon

// Parsing table
// Row: $ number symbol ( ) Quote dot
// Column: S Sexpr List New

type nonTerminal int
const (
	Start_NT nonTerminal = iota
	Sexpr_NT 
	List_NT
	New_NT
)
var mixedArray = [][][]interface{} {{nil, {Sexpr_NT}, {Sexpr_NT}, {Sexpr_NT}, nil, nil, {Sexpr_NT}}, {nil, {tokenSymbol}, {tokenNumber}, {tokenLpar, List_NT, tokenRpar}, nil, nil, {tokenQuote, Sexpr_NT}}, {nil, {Sexpr_NT, New_NT}, {Sexpr_NT, New_NT}, {Sexpr_NT, New_NT}, {}, nil, {Sexpr_NT, New_NT}}, {nil, {Sexpr_NT, New_NT}, {Sexpr_NT, New_NT}, {Sexpr_NT, New_NT}, {}, {tokenDot, Sexpr_NT}, {Sexpr_NT, New_NT}}

// implements the Parser interface
func (g Grammar) Parse(str string) (*SExpr, error) {
	tokenInd := 0
	meetLpar := false
	quoteFlag := false
	dotFlag := false
	tokenList = nil

	// Tokennize the input string
	lex := newLexer(str)
	var	tokenList []*token
	for {
		token, err := lex.next()
		// validating the given string, return error if can't parse to token
		if err == ErrLexer {
			return nil, ErrParser
		} else {
			tokenList = append(tokenList, token)
			if token.typ == tokenEOF {
				break
			}
		}
	}
	var se = &SExpr{}
	// var se, error = parse_sexpr()
	var stack []interface{}
	stack = append(stack, Sxpr_NT)

	if len(tokenList) != 1 && tokenList[0].typ != tokenEOF {
		for len(stack) != 0 {
			ind := len(stack) - 1		// index of top element in the stack
			topOfStack := stack[ind]

			switch typ := topOfStack.(type) {
			case tokenTyp:
				if tokenList[tokenInd].typ == topOfStack {
					// number, symbol -> number, symbol
					// dot ->  mkCons (car cdr)
					if topOfStack == tokenNumber || topOfStack == tokenSymbol{ // 

					} else if topOfStack  == tokenDot{
						if tokenList[tokenInd+1].typ == tokenLpar {
							car := 
							cdr := 

						}

					} else if {

					}

				} 



			case nonTerminal:
				if parseTable[typ][tokenList[tokenInd].typ] != nil {
					// value inside the cell, find the transition to other state
					var transList = parseTable[typ][tokenList[tokenInd].typ]
					var listIndex = len(transList) -1
					stack = stack[:ind]		// pop out the top non terminal before push

					// push T -> X1 X2 X3 to the stack in reverse order
					for listIndex >= 0 {
						stack = append(stack, transList[listIndex])
					   listIndex -= 1
					}
				} else {
					return nil, ErrParser
				}
			/*
				1. <Sexpr> 
				- meetLpar: ( List_NT )
				- quoteFlag: tokenQuote <Sexpr>
				
				2. <List>
				- nil: epsilon
				- else: <Sexpr> <New>

				3. <New>
				- dotFlag: tokenDot <Sexpr>
				- nil: epsilon
				- else: <Sexpr> <New>
			*/
			}
		}
	}
	return mkConscell(car, cdr), error
}

// create the SExpr for the matched terminal
// func matchTerminal(token *Token) (*SExpr, error) {
// 	// Three types of creation


// 	// 1. create the nil s-expression
// 	se := mkNil()
// 	return se, nil


// 	// 2. create the atom s-expression
// 	if token.typ == tokenNumber {
// 		se := mkNumber(token.num)
// 		return se, nil
// 	} else if token.typ == tokenSymbol {
// 		se := mkAtom(mkSymbol(token.literal))
// 		return se, nil
// 	}


// 	// 3. create the list s-expression



// }


/*
func (g Grammar) Parse(str string) (*SExpr, error) {
	// The parseTable can be the type of
	// []interface{} which is the list in a single cell
	var finalTerm = &SExpr{}
	parseTable := mixedArray
	var stk1 = []*Sexpr{} 					
	var stk2 = [][]*Sexpr{}   			
	var stk_ptr = 0
	var termMap = map[string]*SExpr{}  	// term.toString() -> *term


	// Tokennize the input string
	lex := newLexer(str)
	var	tokenList []*token
	for {
		token, err := lex.next()

		if err == ErrLexer {
			// validating the given string, return error if can't parse to token
			return nil, ErrParser
			} else {
				if token.typ == tokenEOF {
					tokenList = append(tokenList, token)
					break
				}
				tokenList = append(tokenList, token)
			}
	}

	// pointer point to the current token in the list
	var tokenInd = 0
	// initialize the stack
	// stack needs to accept two data types, nonTerminal & tokenType
	var stack []interface{}
	stack = append(stack, Start_NT)

	if len(tokenList) != 1 && tokenList[0].typ != tokenEOF {
	 	for len(stack) != 0 {

	 		ind := len(stack) - 1		// index of top element in the stack

	 		topOfStack := stack[ind]

	 		switch typ := topOfStack.(type) { // tokenType or nonTerminal
	 		case tokenType:
	 			if tokenList[tokenInd].typ == topOfStack {
					if topOfStack == tokenAtom && tokenList[tokenInd + 1].typ == tokenLpar {
						// indicator for create the functor term and push items to two stacks
						temp := &Term{Typ: relationMap[tokenList[tokenInd].typ], Literal: tokenList[tokenInd].literal} // -------
						str := temp.String()

						if val, ok := termMap[str]; ok {
							stk1 = append(stk1, val) 		// - CHECK SYNTAX
						} else {
							termMap[str] = temp
							stk1 = append(stk1, temp)
						}

						var tempList = []*Term{}			// arguments list
						stk2 = append(stk2, tempList)		// push empty term[] to the stk2
						stk_ptr++
					} else if topOfStack == tokenRpar {
						// indicator for creating the compound Term
						if stk_ptr > 0 {
							// create the compound term
							temp := &Term{Typ: TermCompound, Functor: stk1[stk_ptr - 1] , Args: stk2[stk_ptr - 1]}

							// pop out the top of two stacks
							stk_ptr--
							stk1 = stk1[:stk_ptr]
							stk2 = stk2[:stk_ptr]

							// check if exits in the termMap avoid duplicate compound
							str := temp.String()
							if val, ok := termMap[str]; ok {
								temp = val 		// use the old *term if exits in the map
							} else {
								// put the new compound in the termMap
								termMap[str] = temp
							}

							// append the new created compound into the next level
							if stk_ptr > 0 {
								stk2[stk_ptr - 1] = append(stk2[stk_ptr - 1], temp)
							} else {
								// we create the last final compound
								finalTerm = temp
							}
						}

					} else if (topOfStack == tokenAtom || topOfStack == tokenNumber || topOfStack == tokenVariable)  {
						// general case
						temp := &Term{Typ: relationMap[tokenList[tokenInd].typ], Literal: tokenList[tokenInd].literal} // 1. Create Term struct
						str := temp.String()

						if val, ok := termMap[str]; ok {
							temp = val
						} else { 			// 3. if not, put new Term into termMap - if no duiplicate, use new Term to append to stk2
							termMap[str] = temp

						}
						finalTerm = temp

						if stk_ptr > 0 {
							stk2[stk_ptr - 1] = append(stk2[stk_ptr - 1], temp)
						}
					} else if (topOfStack == tokenEOF) {
							// only a single term left, return it
						 	if len(termMap) > 0 {
							 	return finalTerm, nil
							}
					}

	 				stack = stack[:ind]		// pop out the top element
	 				tokenInd += 1;
	 			} else {
	 				// terminal is not match
	 				return nil, ErrParser
	 			}

	 		case nonTermial:
	 			// when the top is non terminal
	 			// check the value in the parsing table with given token
	 			if parseTable[typ][tokenList[tokenInd].typ] != nil {
	 				// value inside the cell, find the transition to other state
	 				var transList = parseTable[typ][tokenList[tokenInd].typ]
	 				var listIndex = len(transList) -1
	 				stack = stack[:ind]		// pop out the top non terminal before push

	 				// push T -> X1 X2 X3 to the stack in reverse order
	 				for listIndex >= 0 {
		 				stack = append(stack, transList[listIndex])
						listIndex -= 1
		 			}
	 			} else {
	 				return nil, ErrParser
	 			}
	 		}
	 	}
	}

	// Return here because we see an empty string
	return nil, nil
}
*/

//---------------------------------------------------------------------------------------------------------------------------------------------
// grammar after left factoring
// <sexpr>       ::= NUM | SYM | ( <New1> | QUOTE <sexpr>
// <proper_list> ::= <sexpr> <proper_list> | \epsilon
// <New1>        ::= <proper_list> <New2>
// <New2>        ::= <sexpr> DOT <sexpr> ) | )

// procedure for <sexpr>
func parse_sexpr() (*SExpr, error) {
	token := tokenList[tokenInd]
	tokenTyp := token.typ

	switch tokenTyp {
	case tokenNumber:
		se := mkNumber(token.num)
		tokenInd += 1
		if !meetLpar && !quoteFlag && len(tokenList) > 2{
			return nil, ErrParser
		}
		return se, nil

	case tokenSymbol:
		se := mkAtom(mkTokenSymbol(token.literal))
		tokenInd += 1

		// check for invalid case  X ) 
		// add the quote_flag to prevent falling into the stop case
		if !meetLpar && !quoteFlag && len(tokenList) > 2{
			// fmt.Println(" *** token length ", len(tokenList))
			return nil, ErrParser
		}

		// TODO: How to deal with 'a


		return se, nil

	case tokenLpar:
		// <sexpr> -> ( <New1>
		// match the (
		meetLpar = true
		tokenInd += 1
		se, error := parse_New1()
		if error == nil {
			return se, nil
		} else {
			return nil, ErrParser
		}
		// fmt.Println(se.SExprString())

		// return parse_New1()

	case tokenQuote:
		// <sexpr> -> QUOTE <sexpr>
		// match the QUOTE
		tokenInd += 1
		quoteFlag = true
		quote_se := mkSymbol("QUOTE")
		// se, error := parse_sexpr()
		// return se, nil
		cdr, err := parse_sexpr()

		if err != nil {
			return nil, ErrParser
		}
	
		cdr2 := mkConsCell(cdr, mkNil())
		return mkConsCell(quote_se, cdr2), nil
		
		// if cdr.isAtom() {
		// 	fmt.Println(" *** Runtime info *** : Line 166")
		// 	cdr2 := mkConsCell(cdr, mkNil())
		// 	return mkConsCell(quote_se, cdr2), nil
		// } else {			
		// 	fmt.Println(" *** Runtime info *** : Line 169")
		// 	cdr2 := mkConsCell(cdr, mkNil())
		// 	return mkConsCell(quote_se, cdr2), nil
		// }
		
	default:
		return nil, ErrParser
	}
}

// procedure for <proper_list>
func parse_proper_list() (*SExpr, error) {
	// tokenTyp := tokenList[tokenInd].typ
	for i := 0; i < 3; i++ {
		// if tokenList[tokenInd+1] not exceeding the range
		token := tokenList[tokenInd+i]
		tokenTyp := token.typ
		 
		if tokenTyp == tokenDot  {
			return mkNil(), nil
		}
	}

	se, error := parse_sexpr()
	if error != nil {
		// can't go deeper
		// do back tracking and let proper_list -> e - no need to back track Rule1
		return mkNil(), nil		
	} else {
		car := se
		cdr, error := parse_proper_list()
		return mkConsCell(car, cdr), error
	}
}

// procedure for <New1>
func parse_New1() (*SExpr, error) {
	// Remember to cover New2
	se, error := parse_proper_list()
	// parse_New2()
	if error == nil {
		_, error2 := parse_New2()
		if error2 == nil {
			return se, nil
		} else {
			return nil, ErrParser
		}
	} else {
		return nil, ErrParser
	}
}

// procedure for <New2>
func parse_New2() (*SExpr, error) {
	tokenTyp := tokenList[tokenInd].typ
	if tokenTyp == tokenRpar {
		// <New2> -> )
		// match the )
		tokenInd += 1
		return mkNil(), nil
	} else {
		// <New2> -> <sexpr> DOT <sexpr> )
		// backtracking needed here
		se1, err1 := parse_sexpr()

		if err1 != nil {
			return nil, ErrParser
		}

		// match DOT 
		if tokenList[tokenInd].typ == tokenDot {
			tokenInd++;
		} else {
			return nil, ErrParser
		}

		se2, err2 := parse_sexpr()

		if err2 != nil {
			return nil, ErrParser
		} 

		// check match )
		if tokenList[tokenInd].typ == tokenRpar {
			tokenInd++;
		} else {
			return nil, ErrParser
		}

		// Double Check!!
		return mkConsCell(se1, se2), nil
	}
}

// saved mkConsCell from previous - steps
func backtrack() (*SExpr, error) {
	tokenInd = // reset the index and restart over 


	return mkConsCell(), nil // what was inside the particular step
}