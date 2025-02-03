package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // current position in our input (pointer to current char)
	line         int  // current line in our input
	readPosition int  // current reading position in input (after position char)
	ch           byte // current char
}

func New(input string) *Lexer {
    l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		l.line++
	} else {
		l.ch = l.input[l.readPosition]
	}

	if l.ch == '\n' {
		l.line++
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)
			tok.Line = l.line
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.line)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.line)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.line)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = string(ch) + string(l.ch)
			tok.Line = l.line
		} else {
			tok = newToken(token.BANG, l.ch, l.line)
		}
	case '/':
		tok = newToken(token.FSLASH, l.ch, l.line)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.line)
	case '<':
		tok = newToken(token.LT, l.ch, l.line)
	case '>':
		tok = newToken(token.GT, l.ch, l.line)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) || isPeriod(l.ch) {
			tok.Type, tok.Literal = l.readNumber()
			tok.Line = l.line
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// rules:
// if we encounter a . assume float
// if we already encountered a . reject
// else int
func (l *Lexer) readNumber() (token.TokenType, string) {
	position := l.position
	var tokenType token.TokenType = token.INT
	for isDigit(l.ch) || isPeriod(l.ch) {
		if tokenType == token.FLOAT && isPeriod(l.ch) {
			return token.ILLEGAL, l.input[position:l.position]
		}

		if isPeriod(l.ch) {
			tokenType = token.FLOAT
		}

		l.readChar()
	}

	return tokenType, l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isPeriod(ch byte) bool {
	return '.' == ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte, line int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line}
}
