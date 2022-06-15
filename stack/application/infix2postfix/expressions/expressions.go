package expressions

const (
	INFIX                 = "Infix"
	POSTFIX               = "Postfix"
	OPEN_FLOWER_BRACKET   = "{"
	CLOSED_FLOWER_BRACKET = "}"
	OPEN_ROUND_BRACKET    = "("
	CLOSED_ROUND_BRACKET  = ")"
	OPEN_SQUARE_BRACKET   = "["
	CLOSED_SQUARE_BRACKET = "]"
	PLUS                  = "+"
	MINUS                 = "-"
	STAR                  = "*"
	DEVIDE                = "/"
	POWER                 = "^"
)

func Mapping(i string) string {
	switch i {
	case OPEN_FLOWER_BRACKET:
		return CLOSED_FLOWER_BRACKET
	case OPEN_SQUARE_BRACKET:
		return CLOSED_SQUARE_BRACKET
	case OPEN_ROUND_BRACKET:
		return CLOSED_ROUND_BRACKET
	default:
		return ""
	}
}

var Exps map[int]map[string]string = map[int]map[string]string{
	1: {
		INFIX:   `a+b*c-d/e`,
		POSTFIX: ``,
	},
	2: {
		INFIX:   `((a+b)*c)-d^e^f`,
		POSTFIX: ``,
	},
	3: {
		INFIX:   `a+b*(c^d-e)^(f+g*h)-i`,
		POSTFIX: `abcd^e-fgh*+^*+i-`,
	},
	4: {
		INFIX:   `A*B+C`,
		POSTFIX: `AB*C+`,
	},
	5: {
		INFIX:   `(A+B)*(C/D)`,
		POSTFIX: `AB+CD/*`,
	},
	6: {
		INFIX:   `A*(B*C+D*E)+F`,
		POSTFIX: `ABC*DE*+*F+`,
	},
	7: {
		INFIX:   `(A+B)*C+(D-E)/F+G`,
		POSTFIX: `AB+C*DE-F/+G+`,
	},
	8: {
		INFIX:   `(p+q)*(m-n)`,
		POSTFIX: `pq+mn-*`,
	},
	9: {
		INFIX:   `a+b*(c/d*e)-f-g`,
		POSTFIX: `abcd/e**+f-g-`,
	},
	10: {
		INFIX:   `a+b*c^d^e-f/g*h`,
		POSTFIX: `abcde^^*+fg/h*-`,
	},
	11: {
		INFIX:   `x^y/(a*z)+b`,
		POSTFIX: `xy^az*/b+`,
	},
}

func Precedence(ch string) int {
	switch ch {
	case PLUS, MINUS:
		return 1
	case STAR, DEVIDE:
		return 2
	case POWER:
		return 3
	default:
		return -1
	}
}

func IsOperand(ch string) bool {
	return (ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z")
}
