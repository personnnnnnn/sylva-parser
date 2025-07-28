package transformer

import "fmt"

type BinaryOperator string

const (
	BO_WTF    BinaryOperator = ""
	BO_Plus   BinaryOperator = "+"
	BO_Minus  BinaryOperator = "-"
	BO_Mul    BinaryOperator = "*"
	BO_Div    BinaryOperator = "/"
	BO_Mod    BinaryOperator = "%"
	BO_Concat BinaryOperator = ".."
	BO_Lt     BinaryOperator = "<"
	BO_Gt     BinaryOperator = ">"
	BO_Lte    BinaryOperator = "<="
	BO_Gte    BinaryOperator = ">="
	BO_Eq     BinaryOperator = "=="
	BO_Neq    BinaryOperator = "!="
	BO_And    BinaryOperator = "&&"
	BO_Or     BinaryOperator = "||"
)

func TTtoBinaryOperator(tt TokenType) (BinaryOperator, error) {
	switch tt {
	case TT_Plus:
		return BO_Plus, nil
	case TT_Minus:
		return BO_Minus, nil
	case TT_Mul:
		return BO_Mul, nil
	case TT_Div:
		return BO_Div, nil
	case TT_Mod:
		return BO_Mod, nil
	case TT_DDot:
		return BO_Concat, nil
	case TT_Lt:
		return BO_Lt, nil
	case TT_Gt:
		return BO_Gt, nil
	case TT_Lte:
		return BO_Lte, nil
	case TT_Gte:
		return BO_Gte, nil
	case TT_Eq:
		return BO_Eq, nil
	case TT_Neq:
		return BO_Neq, nil
	case TT_And:
		return BO_And, nil
	case TT_Or:
		return BO_Or, nil
	}
	return BO_WTF, fmt.Errorf("token type %v has no corresponding binary operator", tt)
}
