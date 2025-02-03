package lexer_test

import (
	"testing"

	"monkey/lexer"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let tenPointFive = 10.5;

let add = fn(x, y) {
    x + y;
}

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}

10 == 10;
10 != 9;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
	}{
		{token.LET, "let", 1},
		{token.IDENT, "five", 1},
		{token.ASSIGN, "=", 1},
		{token.INT, "5", 1},
		{token.SEMICOLON, ";", 1},
		{token.LET, "let", 2},
		{token.IDENT, "ten", 2},
		{token.ASSIGN, "=", 2},
		{token.INT, "10", 2},
		{token.SEMICOLON, ";", 2},
		{token.LET, "let", 3},
		{token.IDENT, "tenPointFive", 3},
		{token.ASSIGN, "=", 3},
		{token.FLOAT, "10.5", 3},
		{token.SEMICOLON, ";", 3},
		{token.LET, "let", 5},
		{token.IDENT, "add", 5},
		{token.ASSIGN, "=", 5},
		{token.FUNCTION, "fn", 5},
		{token.LPAREN, "(", 5},
		{token.IDENT, "x", 5},
		{token.COMMA, ",", 5},
		{token.IDENT, "y", 5},
		{token.RPAREN, ")", 5},
		{token.LBRACE, "{", 5},
		{token.IDENT, "x", 6},
		{token.PLUS, "+", 6},
		{token.IDENT, "y", 6},
		{token.SEMICOLON, ";", 6},
		{token.RBRACE, "}", 7},
		{token.LET, "let", 9},
		{token.IDENT, "result", 9},
		{token.ASSIGN, "=", 9},
		{token.IDENT, "add", 9},
		{token.LPAREN, "(", 9},
		{token.IDENT, "five", 9},
		{token.COMMA, ",", 9},
		{token.IDENT, "ten", 9},
		{token.RPAREN, ")", 9},
		{token.SEMICOLON, ";", 9},
		{token.BANG, "!", 10},
		{token.MINUS, "-", 10},
		{token.FSLASH, "/", 10},
		{token.ASTERISK, "*", 10},
		{token.INT, "5", 10},
		{token.SEMICOLON, ";", 10},
		{token.INT, "5", 11},
		{token.LT, "<", 11},
		{token.INT, "10", 11},
		{token.GT, ">", 11},
		{token.INT, "5", 11},
		{token.SEMICOLON, ";", 11},
		{token.IF, "if", 13},
		{token.LPAREN, "(", 13},
		{token.INT, "5", 13},
		{token.LT, "<", 13},
		{token.INT, "10", 13},
		{token.RPAREN, ")", 13},
		{token.LBRACE, "{", 13},
		{token.RETURN, "return", 14},
		{token.TRUE, "true", 14},
		{token.SEMICOLON, ";", 14},
		{token.RBRACE, "}", 15},
		{token.ELSE, "else", 15},
		{token.LBRACE, "{", 15},
		{token.RETURN, "return", 16},
		{token.FALSE, "false", 16},
		{token.SEMICOLON, ";", 16},
		{token.RBRACE, "}", 17},
		{token.INT, "10", 19},
		{token.EQ, "==", 19},
		{token.INT, "10", 19},
		{token.SEMICOLON, ";", 19},
		{token.INT, "10", 20},
		{token.NOT_EQ, "!=", 20},
		{token.INT, "9", 20},
		{token.SEMICOLON, ";", 20},
		{token.EOF, "", 21},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong on line=%d. expected=%q, got=%q", i, tok.Line, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong on line=%d. expected=%q, got=%q", i, tok.Line, tt.expectedLiteral, tok.Literal)
		}

        if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. expected=%d, got=%d", i, tt.expectedLine, tok.Line)
        }
	}
}
