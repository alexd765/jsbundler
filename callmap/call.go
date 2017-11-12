package callmap

// Call describes a function call.
type Call struct {
	Name string `json:"name,omitempty"`
	From string `json:"from,omitempty"`
}
