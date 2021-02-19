package sexpr

import "errors"

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


// implements the Parser interface
func (g Grammar) Parse(str string) (*SExpr, error) {

	// Tokennize the input string
	lex := newLexer(str)
	var	tokenList []*Token
	for {
		token, err := lex.next()

		// validating the given string, return error if can't parse to token
		if err == ErrLexer {
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



}

// create the SExpr for the matched terminal
func matchTerminal(token *Token) (*SExpr, error) {
	// Three type of creation


	// 1. create the nil s-expression


	// 2. create the atom s-expression


	// 3. create the list s-expression



}


// grammar after left factoring
// <sexpr>       ::= NUM | SYM | ( <New1> | QUOTE <sexpr>
// <proper_list> ::= <sexpr> <proper_list> | \epsilon
// <New1>        ::= <proper_list> <New2>
// <New2>        ::= <sexpr> DOT <sexpr> ) | )

// procedure for <sexpr>
func parse_sexpr() {
	tokenTyp := tokenList[tokenInd].typ
	if tokenTyp == tokenNumber || tokenTyp == tokenSymbol {
		// match the terminal
		// increment the tokenInd
	} else if tokenTyp == tokenLpar {
		// <sexpr> -> ( <New1>

		// match the (
		parse_New1()
	} else if tokenTyp == tokenQuote {
		// <sexpr> -> QUOTE <sexpr>

		// match the QUOTE
		parse_sexpr()
	}
}

// procedure for <proper_list>
func parse_proper_list() {
	tokenTyp := tokenList[tokenInd].typ
	if tokenTyp == tokenEOF {
		// match \epsilon
	} else {
		// <proper_list> -> <sexpr> <proper_list>
		parse_sexpr()
		parse_proper_list()
	}
}

// procedure for <New1>
func parse_New1() {
	parse_proper_list()
	parse_New2()
}

// procedure for <New2>
func parse_New2() {
	tokenTyp := tokenList[tokenInd].typ
	if tokenTyp == tokenRpar {
		// <New2> -> )

		// match the )
	} else {
		// <New2> -> <sexpr> DOT <sexpr> )
		parse_sexpr()
		// TODO: check match DOT 
		parse_sexpr()
		// TODO: check mathch )
	}
}

