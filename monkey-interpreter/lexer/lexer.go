package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position input (points to current char)
	readPosition int  // current reading position in input (next to current char)
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Start the lexer, initialize other fields
	return l
}

func (self *Lexer) peekChar() byte {
	if self.readPosition > len(self.input) {
		return 0
	} else {
		return self.input[self.readPosition]
	}
}

func (self *Lexer) readChar() {
	if self.readPosition >= len(self.input) {
		self.ch = 0
	} else {
		self.ch = self.input[self.readPosition]
	}

	self.position = self.readPosition
	self.readPosition += 1
}

func (self *Lexer) NextToken() token.Token {
	var tok token.Token

	self.skipWhile(isWhitespace)

	switch self.ch {
	case '=':
		if self.peekChar() == '=' {
			self.readChar()
			tok.Type = token.EQ
			tok.Literal = "=="
		} else {
			tok = newToken(token.ASSIGN, self.ch)
		}
	case '+':
		tok = newToken(token.PLUS, self.ch)
	case '-':
		tok = newToken(token.MINUS, self.ch)
	case '!':
		if self.peekChar() == '=' {
			self.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = "!="
		} else {
			tok = newToken(token.BANG, self.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, self.ch)
	case '/':
		tok = newToken(token.SLASH, self.ch)
	case '<':
		tok = newToken(token.LT, self.ch)
	case '>':
		tok = newToken(token.GT, self.ch)

	case ',':
		tok = newToken(token.COMMA, self.ch)
	case ';':
		tok = newToken(token.SEMICOLON, self.ch)
	case '(':
		tok = newToken(token.LPAREN, self.ch)
	case ')':
		tok = newToken(token.RPAREN, self.ch)
	case '{':
		tok = newToken(token.LBRACE, self.ch)
	case '}':
		tok = newToken(token.RBRACE, self.ch)
	case 0: // EOF
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(self.ch) {
			tok.Literal = self.readWhile(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(self.ch) {
			tok.Type = token.INT
			tok.Literal = self.readWhile(isDigit)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, self.ch)
		}
	}

	self.readChar()
	return tok
}

// readWhile advances the Lexer for as long as the `test` func returns true.
// Returns the input from where the test starts until it ends(`test` returns false).
func (self *Lexer) readWhile(test func(byte) bool) string {
	startPos := self.position

	for test(self.ch) {
		self.readChar()
	}

	return self.input[startPos:self.position]
}

// skipWhile advances the Lexer for as long as the `test` func returns true.
func (self *Lexer) skipWhile(test func(byte) bool) {
	for test(self.ch) {
		self.readChar()
	}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
