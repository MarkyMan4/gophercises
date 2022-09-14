package object

const (
	INTEGER_OBJ = "INTEGER"
	FLOAT_OBJ   = "FLOAT"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	ERROR_OBJ   = "ERROR"
)

type Object interface {
	Type() string
	ToString() string
}
