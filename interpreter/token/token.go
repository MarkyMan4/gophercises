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
