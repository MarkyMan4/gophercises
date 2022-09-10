package token

type Token struct {
	Type    string
	Literal string
}

const (
	VAR    = "VAR"
	PLUS   = "+"
	MINUS  = "-"
	MULT   = "*"
	DIVIDE = "/"
	LPAREN = "("
	RPAREN = ")"
	SEMI   = ";"
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
)

var keywords = map[string]string{
	"var": VAR,
}

// lookup a value from the input and determine if it is a keyword or an identifier
func GetIdentOrKeyword(literal string) string {
	if tok, ok := keywords[literal]; ok {
		return tok
	}

	return IDENT
}
