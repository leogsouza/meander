package meander

// Facade is a interface which has a Public method
type Facade interface {
	Public() interface{}
}

// Public method checks if params is a Facade type
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
