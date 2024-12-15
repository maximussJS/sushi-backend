package di

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Token       string
}
