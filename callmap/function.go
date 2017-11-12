package callmap

// Function is a function declaration.
type Function struct {
	Calls     []Call
	Functions []Function
	Name      string
}
