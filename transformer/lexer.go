package transformer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Token struct {
	Bounds Bounds
	Type   TokenType
	Value  any
}

func (t *Token) String() string {
	return fmt.Sprintf("(%v: %v)", t.Type, t.Value)
}

type TokenType byte

const (
	TT_EOF TokenType = iota
	TT_Err
	TT_Int
	TT_Float
	TT_String
	TT_True
	TT_False
	TT_Minus
	TT_Plus
	TT_Mul
	TT_Div
	TT_Mod
	TT_LParen
	TT_RParen
	TT_DDot
	TT_Comma
	TT_Symbol
	TT_Lt
	TT_Gt
	TT_Lte
	TT_Gte
	TT_Eq
	TT_Neq
	TT_And
	TT_Or
	TT_Not

	TT_If
	TT_Unless
	TT_For
	TT_While
	TT_Until
	TT_Loop
	TT_Fn
	TT_Typeof
	TT_Copyof
	TT_In // ITEM in value
	TT_Of // KEY in value
	TT_As
	TT_Nil
	TT_Dot
	TT_TDot
	TT_LBrack
	TT_RBrack
	TT_LCurly
	TT_RCurly
	TT_Hash
	TT_Pipe
	TT_With
	TT_By
	TT_To
	TT_Break
	TT_Continue
	TT_Return
	TT_Pub
	// TT_Priv (?)
	TT_QMark
	TT_EqSign
	TT_Arrow
	TT_Colon
	TT_Let
	TT_Else
)

func (t TokenType) String() string {
	switch t {
	case TT_EOF:
		return "eof"
	case TT_Err:
		return "err"
	case TT_Int:
		return "int"
	case TT_Float:
		return "float"
	case TT_Minus:
		return "minus"
	case TT_Plus:
		return "plus"
	case TT_Mul:
		return "mul"
	case TT_Div:
		return "div"
	case TT_Mod:
		return "mod"
	case TT_LParen:
		return "lparen"
	case TT_RParen:
		return "rparen"
	case TT_DDot:
		return "ddot"
	case TT_Symbol:
		return "symbol"
	case TT_Comma:
		return "comma"
	case TT_True:
		return "true"
	case TT_False:
		return "false"
	case TT_String:
		return "string"
	case TT_Gt:
		return "gt"
	case TT_Lt:
		return "lt"
	case TT_Lte:
		return "lte"
	case TT_Gte:
		return "gte"
	case TT_Eq:
		return "eq"
	case TT_Neq:
		return "neq"
	case TT_And:
		return "and"
	case TT_Or:
		return "or"
	case TT_Not:
		return "not"
	case TT_If:
		return "if"
	case TT_Unless:
		return "unless"
	case TT_For:
		return "for"
	case TT_While:
		return "while"
	case TT_Until:
		return "until"
	case TT_Loop:
		return "loop"
	case TT_Fn:
		return "fn"
	case TT_Typeof:
		return "typeof"
	case TT_Copyof:
		return "copyof"
	case TT_In:
		return "in"
	case TT_Of:
		return "of"
	case TT_As:
		return "as"
	case TT_Nil:
		return "nil"
	case TT_Dot:
		return "dot"
	case TT_TDot:
		return "tdot"
	case TT_LBrack:
		return "lbrack"
	case TT_RBrack:
		return "rbrack"
	case TT_LCurly:
		return "lcurly"
	case TT_RCurly:
		return "rcurly"
	case TT_Hash:
		return "hash"
	case TT_Pipe:
		return "pipe"
	case TT_With:
		return "with"
	case TT_By:
		return "by"
	case TT_To:
		return "to"
	case TT_Break:
		return "break"
	case TT_Continue:
		return "continue"
	case TT_Return:
		return "return"
	case TT_Pub:
		return "pub"
		// case TT_Priv:
		// return "priv"
	case TT_QMark:
		return "qmark"
	case TT_EqSign:
		return "eqsign"
	case TT_Arrow:
		return "arrow"
	case TT_Colon:
		return "colon"
	case TT_Let:
		return "let"
	case TT_Else:
		return "else"
	default:
		return string(t) + "?"
	}
}

// tokens which are one symbol and have no associated data
var TrivialTokens = map[string]TokenType{
	"+":        TT_Plus,
	"-":        TT_Minus,
	"*":        TT_Mul,
	"/":        TT_Div,
	"%":        TT_Mod,
	"(":        TT_LParen,
	")":        TT_RParen,
	"..":       TT_DDot,
	",":        TT_Comma,
	"<":        TT_Lt,
	">":        TT_Gt,
	"<=":       TT_Lte,
	">=":       TT_Gte,
	"==":       TT_Eq,
	"!=":       TT_Neq,
	"&&":       TT_And,
	"||":       TT_Or,
	"!":        TT_Not,
	"true":     TT_True,
	"false":    TT_False,
	"if":       TT_If,
	"unless":   TT_Unless,
	"for":      TT_For,
	"while":    TT_While,
	"until":    TT_Until,
	"loop":     TT_Loop,
	"fn":       TT_Fn,
	"typeof":   TT_Typeof,
	"copyof":   TT_Copyof,
	"in":       TT_In,
	"of":       TT_Of,
	"as":       TT_As,
	"nil":      TT_Nil,
	".":        TT_Dot,
	"...":      TT_TDot,
	"[":        TT_LBrack,
	"]":        TT_RBrack,
	"{":        TT_LCurly,
	"}":        TT_RCurly,
	"#":        TT_Hash,
	"|>":       TT_Pipe,
	"by":       TT_By,
	"to":       TT_To,
	"break":    TT_Break,
	"continue": TT_Continue,
	"return":   TT_Return,
	"pub":      TT_Pub,
	"?":        TT_QMark,
	"=":        TT_EqSign,
	"=>":       TT_Arrow,
	":":        TT_Colon,
	"let":      TT_Let,
	"else":     TT_Else,
	// "priv":     TT_Priv,
}

const WHITESPACE = " \t\r\n\f"

type Lexer struct {
	Tokens          []*Token
	Position        Position
	Text            string
	Errors          []*SylvaError
	OrderedByLength []string // for caching
}

func (l *Lexer) Lex() {
	for l.Char() != 0 {
		l.Logic()
	}

	eof := l.MakeToken(TT_EOF, 1)
	l.Tokens = append(l.Tokens, eof)
}

func MakeLexer(text string) *Lexer {
	// biggest first, smallest last
	orderedByLength := []string{}
	for k := range TrivialTokens {
		orderedByLength = append(orderedByLength, k)
		if len(orderedByLength) == 1 {
			continue
		}

		for i := len(orderedByLength) - 1; i > 0; i-- {
			if len(orderedByLength[i-1]) >= len(orderedByLength[i]) {
				break
			}
			orderedByLength[i-1], orderedByLength[i] = orderedByLength[i], orderedByLength[i-1]
		}
	}

	return &Lexer{
		Tokens:          []*Token{},
		Position:        Position{},
		Text:            text,
		Errors:          []*SylvaError{},
		OrderedByLength: orderedByLength,
	}
}

func (l *Lexer) Char() byte {
	if l.Position.Idx >= int32(len(l.Text)) {
		return 0
	}
	return l.Text[l.Position.Idx]
}

func (l *Lexer) Step() {
	l.Position = l.Position.Step(l.Char())
}

func (l *Lexer) MakeToken(type_ TokenType, length int) *Token {
	bounds := MakeBoundsForToken(l.Position, length)
	l.Position = bounds.End
	return &Token{Type: type_, Bounds: bounds, Value: nil}
}

func (l *Lexer) MatchesToken(token string) bool {
	oldPos := l.Position
	defer func() {
		l.Position = oldPos
	}()
	for i := range token {
		if l.Char() != token[i] {
			return false
		}
		l.Step()
	}
	return true
}

func IsWhitespace(c byte) bool {
	return strings.Contains(WHITESPACE, string(c))
}

func IsAlphanumeric(s string) bool {
	for _, char := range s {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_') {
			return false
		}
	}
	return true
}

func (l *Lexer) Token(token *Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) Logic() {
	if IsWhitespace(l.Char()) || l.Char() == 0 {
		l.Step()
		return
	}

	if l.MatchesToken("--") {
		for l.Char() != '\n' && l.Char() != 0 {
			l.Step()
		}
		return
	}

	for _, token := range l.OrderedByLength {
		if !l.MatchesToken(token) {
			continue
		}
		// symbols like "ifTrue" would break and be interpreted as (token "if"), (symbol "True")
		if IsAlphanumeric(token) {
			break
		}
		l.Token(l.MakeToken(TrivialTokens[token], len(token)))
		return
	}

	if l.Char() == '"' || l.Char() == '\'' {
		l.ParseString()
		return
	}

	if unicode.IsDigit(rune(l.Char())) {
		l.ParseNum()
		return
	}

	if unicode.IsLetter(rune(l.Char())) || l.Char() == '_' {
		l.ParseSymbol()
		return
	}

	// unknown symbol error
	l.UnknownSymbol()
}

func (l *Lexer) UnknownSymbol() {
	start := l.Position
	unknownSymbol := ""
	for {
		if IsWhitespace(l.Char()) || l.Char() == 0 {
			break
		}
		matchesSomeToken := false
		for _, token := range l.OrderedByLength {
			if !l.MatchesToken(token) || IsAlphanumeric(token) {
				continue
			}
			matchesSomeToken = true
			break
		}
		if matchesSomeToken {
			break
		}

		unknownSymbol += string(l.Char())
		l.Step()
	}

	bounds := Bounds{Start: start, End: l.Position}
	l.Token(&Token{
		Type:   TT_Err,
		Bounds: bounds,
		Value:  unknownSymbol,
	})
	l.Errors = append(l.Errors, &SylvaError{
		Bounds:  bounds,
		Type:    "lex error",
		Message: fmt.Sprintf("unknown symbol '%v'", unknownSymbol),
	})
}

func (l *Lexer) ParseNum() {
	start := l.Position
	numString := ""
	hasDot := false
	for unicode.IsDigit(rune(l.Char())) || l.Char() == '.' {
		if l.Char() == '.' {
			if hasDot {
				break
			}
			hasDot = true
		}
		numString += string(l.Char())
		l.Step()
	}

	if !hasDot {
		num, _ := strconv.ParseInt(numString, 10, 64)
		l.Token(&Token{
			Bounds: Bounds{Start: start, End: l.Position},
			Type:   TT_Int,
			Value:  num,
		})
	} else {
		num, _ := strconv.ParseFloat(numString, 64)
		l.Token(&Token{
			Bounds: Bounds{Start: start, End: l.Position},
			Type:   TT_Float,
			Value:  num,
		})
	}
}

func (l *Lexer) ParseSymbol() {
	start := l.Position
	symbol := ""

	for l.Char() == '_' {
		symbol += string(l.Char())
		l.Step()
	}

	for unicode.IsLetter(rune(l.Char())) || (symbol != "" && unicode.IsDigit(rune(l.Char()))) {
		symbol += string(l.Char())
		l.Step()
	}

	if token, ok := TrivialTokens[symbol]; ok {
		l.Token(&Token{
			Bounds: MakeBounds(start, l.Position),
			Type:   token,
			Value:  nil,
		})
		return
	}

	l.Token(&Token{
		Bounds: MakeBounds(start, l.Position),
		Type:   TT_Symbol,
		Value:  symbol,
	})
}

func escapeSequence(b byte) (byte, error) {
	switch b {
	case '\\':
		return '\\', nil
	case '0':
		return 0, nil
	case 'n':
		return '\n', nil
	case 't':
		return '\t', nil
	case 'f':
		return '\f', nil
	case 'b':
		return '\b', nil
	case '\'':
		return '\'', nil
	case '"':
		return '"', nil
	case 'a':
		return '\a', nil
	case 'v':
		return '\v', nil
	case 'r':
		return '\r', nil
	default:
		return 0, fmt.Errorf("'\\%v' is not a valid escape sequence", b)
	}
}

func (l *Lexer) ParseString() {
	start := l.Position
	quote := l.Char()
	if !(quote == '\'' || quote == '"') {
		fmt.Println("[WARNING] something really bad just happened when parsing a string (from Lexer.ParseString)")
		l.UnknownSymbol()
		return
	}

	l.Step()

	chars := []byte{}

	wereErrors := false
	for l.Char() != quote && l.Char() != 0 && l.Char() != '\n' {
		if l.Char() != '\\' {
			chars = append(chars, l.Char())
			l.Step()
		} else {
			escapeStart := l.Position
			l.Step()
			if l.Char() == 0 || l.Char() == '\n' {
				wereErrors = true
				l.Errors = append(l.Errors, &SylvaError{
					Bounds:  MakeBounds(l.Position, l.Position.Next()),
					Type:    "lex error",
					Message: "expected another character",
				})
				break
			}
			currChar := l.Char()
			l.Step()
			res, err := escapeSequence(currChar)
			if err != nil {
				l.Errors = append(l.Errors, &SylvaError{
					Bounds:  MakeBounds(escapeStart, l.Position),
					Type:    "lex error",
					Message: err.Error(),
				})
				res = 0
			}
			chars = append(chars, res)
		}
	}

	token := &Token{
		Type:   TT_String,
		Bounds: MakeBounds(start, l.Position),
		Value:  string(chars),
	}

	if l.Char() != quote && !wereErrors {
		l.Errors = append(l.Errors, &SylvaError{
			Bounds:  MakeBounds(l.Position, l.Position.Next()),
			Type:    "lex error",
			Message: "expected closing quote",
		})
	}

	l.Step()
	l.Token(token)
}
