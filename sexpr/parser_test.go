package sexpr

import (
	"math/big"
	"testing"
)

func buildNilExample() *SExpr {
	return mkNil()
}

func buildNumberExample() *SExpr {
	return mkNumber(big.NewInt(100))
}

func buildSymbolExample() *SExpr {
	return mkSymbol("+")
}

func buildConsCellExample() *SExpr {
	return mkConsCell(mkSymbol("A"), mkSymbol("B"))
}

func buildProperListExample() *SExpr {
	return mkConsCell(mkSymbol("A"), mkConsCell(mkSymbol("B"), mkConsCell(mkSymbol("C"), mkNil())))
}

func buildDottedListExample() *SExpr {
	return mkConsCell(mkSymbol("A"), mkConsCell(mkSymbol("B"), mkSymbol("C")))
}

func buildQuoteExample() *SExpr {
	return mkConsCell(mkSymbol("QUOTE"), mkConsCell(mkSymbol("A"), mkNil()))
}

type SExprBuilder func() *SExpr

func TestParseExample(t *testing.T) {
	for idx, test := range []struct {
		input                string
		expectedSExprBuilder SExprBuilder
	}{
		{"()", buildNilExample},
		{"100", buildNumberExample},
		{"+", buildSymbolExample},
		{"(A . B)", buildConsCellExample},
		{"(A B C)", buildProperListExample},
		{"(A B . C)", buildDottedListExample},
		{"'A", buildQuoteExample},
	} {
		actual, err := NewParser().Parse(test.input)
		expected := test.expectedSExprBuilder()
		if err != nil {
			t.Errorf("\nin test %d\nunexpected error", idx)
			continue
		}
		if actual.SExprString() != expected.SExprString() {
			t.Errorf("\nin test %d (\"%s\")\nerror: got      %s\n       expected %s",
				idx, test.input, actual.SExprString(), expected.SExprString())
		}
	}
}

func TestParserInvalid(t *testing.T) {
	for idx, test := range []string{
		"",
		"(",
		"'",
		")",
		"x)",
		"( ) ( )",
		"(a . () . () . ())",
		"((x .",
		"(x",
	} {
		_, err := NewParser().Parse(test)
		if err == nil {
			t.Errorf("\nin test %d\nshould error", idx)
		}
	}
}

func TestParserProperList(t *testing.T) {
	for idx, test := range []struct {
		input, expectedSExprString string
	}{
		{"()", "NIL"},
		{"a", "A"},

		// proper lists
		{"(())", "(NIL . NIL)"},
		{"(a)", "(A . NIL)"},
		{"((a))", "((A . NIL) . NIL)"},
		{"(a b c)", "(A . (B . (C . NIL)))"},
		{"(a b c d)", "(A . (B . (C . (D . NIL))))"},
		{"(  a b c  d  e)", "(A . (B . (C . (D . (E . NIL)))))"},
		{"(  (a b )c)", "((A . (B . NIL)) . (C . NIL))"},
		{"(a b (c d))", "(A . (B . ((C . (D . NIL)) . NIL)))"},
		{"(a () () a)", "(A . (NIL . (NIL . (A . NIL))))"},

		// dotted lists
		{"(a (b . c))", "(A . ((B . C) . NIL))"}, // parenthesis: 2 pairs dot: 1 sym: 3 -> PL 1 // Rule1 o
		{"(a . b)", "(A . B)"}, // p: 2 dot: 1 sym: 2 -> PL 0 // Rule1 o
		{"(a . (b . c))", "(A . (B . C))"}, // p: 2 dot: 2 sym: 3 -> PL 0  // Rule1 o
		{"(a . (b . (c . d)))", "(A . (B . (C . D)))"}, // p: 3 dot: 3 sym: 4 -> PL 0 // Rule1 o 
		{"(a . ((b . c) . d))", "(A . ((B . C) . D))"}, // p: 3 dot: 3 sym: 4 -> PL 0 // Not sure if our grammar works for this
		{"(a b . c)", "(A . (B . C))"}, // p: 2 dot : 1 sym: 3 -> PL  1 // Rule1 o
		{"(a b c . d)", "(A . (B . (C . D)))"}, // p: 2 dot: 1 sym: 4 -> PL 2 // Rule1 o
		{"(a (b c) d . e)", "(A . ((B . (C . NIL)) . (D . E)))"}, // p: 2 dot: 1 sym:5 -> PL ... // Rule1 o
		{"(a b (c d) . e)", "(A . (B . ((C . (D . NIL)) . E)))"}, // p: 2 dot: 1 sym:5 
		{"(a b c . (d e))", "(A . (B . (C . (D . (E . NIL)))))"}, // p: 2 dot: 1 sym:5
		{"(a b c . (d . e))", "(A . (B . (C . (D . E))))"}, // p: 2 dot: 2 sym: 5
		{"(a(b.c  ))", "(A . ((B . C) . NIL))"}, // p: 2  dot: 1 sym: 3
		{"(a . ( (  ) . ( ( ) . a)))", "(A . (NIL . (NIL . A)))"}, // p: 4 dot: 3 sym: 2
		// No relevance between parenthesis and dots but <new> and parenthesis has relevance 
		// Need to find relation between proper_list and dot

		// from the current tokenIndex if there are dot between [current tokenIndex, current tokenIndex + 2]
		// then <Proper_list> will always have to go to production rule of epsilon - properFlag = True
		// properFlag = False back to off as the tokenIndex increase and there's no dot in the range - Rule1

		// N2 production Rule  - Rule 2

		// Whether S to have terminal or N2 to have terminal - Rule 3
		// if the [current tokenIndex, current tokenIndex + 1]

		// Not using the rule - backtracking but where to stop and where to come back and what data to be restored and begin from
		// It would be better to ( tokenIndex - a ) start from the one of the beginning step with regulation on the production rules
		// when there has been an error 

		
	} {
		actual, err := NewParser().Parse(test.input)
		if err != nil {
			t.Errorf("\nin test %d\nunexpected error", idx)
			continue
		}
		if actual.SExprString() != test.expectedSExprString {
			t.Errorf("\nin test %d (\"%s\")\nerror: got      %s\n       expected %s",
				idx, test.input, actual.SExprString(), test.expectedSExprString)
		}
	}
}

func TestParseQuote(t *testing.T) {
	for idx, test := range []struct {
		input, expectedSExpr string
	}{
		{"'(1 2)", "(QUOTE . ((1 . (2 . NIL)) . NIL))"},
		{"'(1 . 2)", "(QUOTE . ((1 . 2) . NIL))"},
		{"(quote . (1 . 2))", "(QUOTE . (1 . 2))"},
		{"'a", "(QUOTE . (A . NIL))"},
		{"'(a)", "(QUOTE . ((A . NIL) . NIL))"},
		{"''a", "(QUOTE . ((QUOTE . (A . NIL)) . NIL))"},
		{"''(a)", "(QUOTE . ((QUOTE . ((A . NIL) . NIL)) . NIL))"},
		{"(' a 'b '  c)", "((QUOTE . (A . NIL)) . ((QUOTE . (B . NIL)) . ((QUOTE . (C . NIL)) . NIL)))"},
	} {
		actual, _ := NewParser().Parse(test.input)
		if actual.SExprString() != test.expectedSExpr {
			t.Errorf("\nerror: in test %d (\"%s\"):\n\texpected: %s\n\tgot      %s",
				idx, test.input, test.expectedSExpr, actual.SExprString())
		}
	}
}
