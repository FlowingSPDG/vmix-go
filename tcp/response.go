package vmixtcp

import "github.com/FlowingSPDG/vmix-go/common/models"

// TallyStatus alias to uint
//go:generate stringer -type=TallyStatus
type TallyStatus uint

// TallyResponse TALLY Event response
type TallyResponse struct {
	Tally []TallyStatus
}

const (
	Off TallyStatus = iota
	Program
	Preview
)

func encodeTally(b byte) TallyStatus {
	switch b {
	case '0':
		return Off
	case '1':
		return Program
	case '2':
		return Preview
	}
	return Off
}

func encodeTallies(b []byte) []TallyStatus {
	var tallies []TallyStatus
	for _, v := range b {
		tallies = append(tallies, encodeTally(v))
	}
	return tallies
}

// TODO: STATUS=ER の場合にエラーイベントを送る

type VersionResponse struct {
	Version string
}

type FunctionResponse struct {
	Response string
}

type ActsResponse struct {
	// e.g. Following examples are for "InputPreview" command.
	//
	// "InputPreview 1 1"
	// "InputPreview 1 0"
	// "InputPreview 3 1"
	//
	// this means Input 1 is used for Preview(L1), and vMix changed to input 3.
	// L2 means Input1 is no longer used for Preview.
	// L3 means Input3 is now used for Preview.
	Response string // True=1 False=0 or 32bit float.
}

type XMLResponse struct {
	XML *models.APIXML
}

type XMLTextResponse struct {
	XMLText string
}

type SubscribeResponse struct {
	Command string
}

type UnsubscribeResponse struct {
	Command string
}
