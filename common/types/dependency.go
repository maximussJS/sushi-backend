package types

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Token       string
}
