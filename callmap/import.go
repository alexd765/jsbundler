package callmap

// Import describes a javascript import
type Import struct {
	Name string `json:"name,omitempty"`
	From string `json:"from,omitempty"`
}
