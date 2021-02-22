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

// implements the Parser interface
func (g Grammar) Parse(str string) (*SExpr, error) {

	// Tokennize the input string
	lex := newLexer(str)
	// var	tokenList []*token
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
	// pointer point to the current token in the list
	// var tokenInd = 0

	// var se = &SExpr{}
	var se, error = parse_sexpr()
	tokenInd = 0
	meetLpar = false
	tokenList = nil
	return se, error

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
		if !meetLpar && len(tokenList) > 1{
			//meetLpar = false
			return nil, ErrParser
		}
		return se, nil

	case tokenSymbol:
		se := mkAtom(mkTokenSymbol(token.literal))
		tokenInd += 1
		if !meetLpar && len(tokenList) > 1{
			//meetLpar = false
			return nil, ErrParser
		}
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
		// se, error := parse_sexpr()
		// return se, nil
		return parse_sexpr()
		
	default:
		return nil, ErrParser
	}
}

// procedure for <proper_list>
func parse_proper_list() (*SExpr, error) {
	// tokenTyp := tokenList[tokenInd].typ
	se, error := parse_sexpr()
	
	if error != nil {
		return mkNil(), ErrParser
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
		parse_New2()
		return se, nil
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
		parse_sexpr()
		// TODO: check match DOT 
		parse_sexpr()
		// TODO: check mathch )
		// Double Check!!
		return mkNil(), nil
	}
}

