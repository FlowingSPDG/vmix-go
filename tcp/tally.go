package vmixtcp

// TallyStatus alias to uint
//go:generate stringer -type=TallyStatus
type TallyStatus uint

// TallyResponse TALLY Event response
type TallyResponse struct {
	Status string
	Tally  []TallyStatus
}

const (
	Off TallyStatus = iota
	Program
	Preview
)
