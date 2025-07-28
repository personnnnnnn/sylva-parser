package transformer

import "fmt"

type UnaryOperator string

const (
	UO_WTF    UnaryOperator = ""
	UO_Plus   UnaryOperator = "+"
	UO_Minus  UnaryOperator = "-"
	UO_Len    UnaryOperator = "#"
	UO_Not    UnaryOperator = "!"
	UO_Typeof UnaryOperator = "typeof"
	UO_Copyof UnaryOperator = "copyof"
)

func TTtoUnaryOperator(tt TokenType) (UnaryOperator, error) {
	switch tt {
	case TT_Plus:
		return UO_Plus, nil
	case TT_Minus:
		return UO_Minus, nil
	case TT_Not:
		return UO_Not, nil
	case TT_Hash:
		return UO_Len, nil
	case TT_Copyof:
		return UO_Copyof, nil
	case TT_Typeof:
		return UO_Typeof, nil
	}
	return UO_WTF, fmt.Errorf("token type %v has no corresponding unary operator", tt)
}
