package token

type Token struct {
	Type    string
	Literal string
}

const (
	VAR    = "VAR"
	FOR    = "FOR"
	IF     = "IF"
	ELSE   = "ELSE"
	DEF    = "DEF"
	PLUS   = "+"
	MINUS  = "-"
	MULT   = "*"
	DIVIDE = "/"
	ASSIGN = "="
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"
	SEMI   = ";"
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
	EOF    = "EOF"
)

var keywords = map[string]string{
	"var":  VAR,
	"for":  FOR,
	"if":   IF,
	"else": ELSE,
	"def":  DEF,
}

// lookup a value from the input and determine if it is a keyword or an identifier
func GetIdentOrKeyword(literal string) string {
	if tok, ok := keywords[literal]; ok {
		return tok
	}

	return IDENT
}
