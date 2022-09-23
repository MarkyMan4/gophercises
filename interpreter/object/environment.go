package object

type Environment struct {
	definitions map[string]Object
	parent      *Environment
}

func NewEnvironment() *Environment {
	defs := make(map[string]Object)

	return &Environment{definitions: defs}
}

func CreateChildEnvironment(env *Environment) *Environment {
	defs := make(map[string]Object)
	child := &Environment{definitions: defs, parent: env}

	return child
}

func (e *Environment) Get(ident string) Object {
	return e.definitions[ident]
}

func (e *Environment) Set(ident string, obj Object) {
	e.definitions[ident] = obj
}

func (e *Environment) GetEnvMap() map[string]Object {
	return e.definitions
}

func (e *Environment) GetParentEnv() *Environment {
	return e.parent
}
