package vmixtcp

// Response TCP API Response
type Response struct {
	Command        string // Command. e.g. "TALLY"
	StatusOrLength string // Status or command length. e.g.  "OK" or "27"(XML)
	Response       string // Command response. e.g. "PreviewInput"(FUNCTION) or "0121..."(TALLY)
	Data           string // Optional data. e.g. XML.
}
