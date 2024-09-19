package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = newTokenByte(token.ASSIGN, l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = newTokenByte(token.BANG, l.ch)
		}
	case ';':
		tok = newTokenByte(token.SEMICOLON, l.ch)
	case '(':
		tok = newTokenByte(token.LPAREN, l.ch)
	case ')':
		tok = newTokenByte(token.RPAREN, l.ch)
	case ',':
		tok = newTokenByte(token.COMMA, l.ch)
	case '+':
		tok = newTokenByte(token.PLUS, l.ch)
	case '-':
		tok = newTokenByte(token.MINUS, l.ch)
	case '*':
		tok = newTokenByte(token.ASTERISK, l.ch)
	case '/':
		tok = newTokenByte(token.SLASH, l.ch)
	case '<':
		tok = newTokenByte(token.LT, l.ch)
	case '>':
		tok = newTokenByte(token.GT, l.ch)
	case '{':
		tok = newTokenByte(token.LBRACE, l.ch)
	case '}':
		tok = newTokenByte(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readWhile(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		if isDigit(l.ch) {
			tok.Literal = l.readWhile(isDigit)
			tok.Type = token.INT
			return tok
		}

		tok = newTokenByte(token.ILLEGAL, l.ch)
	}

	l.readChar()

	return tok
}

func (l *Lexer) readWhile(fn func(byte) bool) string {
	position := l.position

	for fn(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func newTokenByte(tokenType token.TokenType, ch byte) token.Token {
	return newToken(tokenType, string(ch))
}
