package vmixtcp

// Response TCP API Response
type Response struct {
	Command        string
	StatusOrLength string
	Response       string
	Data           string
}
