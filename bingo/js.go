package main

type JsTokenType string

const(
	ILLEGAL JsTokenType = "ILLEGAL"
	EOF = "EOF"
	INDENT = "INDENT"
	INT = "INT"
	ASSIGN = "="
	PLUS ="+"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type JsLexer struct{
	input string
	pos int
	readPos int
	ch byte
}

func (l *JsLexer) readChar(){
	if l.readPos >= len(l.input){
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *JsLexer) NextToken() *JsToken{
	var t *JsToken
	switch l.ch{
	case '=':
		t = newJsToken(ASSIGN, '=')
	case ';':
		t = newJsToken(SEMICOLON, ';')
	case '(':
		t = newJsToken(LPAREN, '(')
	case ')':
		t = newJsToken(RPAREN, ')')
	case ',':
		t = newJsToken(COMMA, ',')
	case '+':
		t = newJsToken(PLUS, '+')
	case '{':
		t = newJsToken(LBRACE, '{')
	case '}':
		t = newJsToken(RBRACE, '}')
	case 0:
		// EOFを表現する
		t = &JsToken{JsTokenType:EOF, Literal:""}
	}
	return t
}

type JsToken struct{
	JsTokenType JsTokenType
	Literal string
}

func newJsToken(t JsTokenType, ch byte) *JsToken{
	return &JsToken{
		JsTokenType:t,
		Literal:string(ch),
	}
}
