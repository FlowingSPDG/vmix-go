package vmixtcp

import "fmt"

const (
	// Terminate letter
	Terminate = "\r\n"
)

func newXMLCommand() []byte {
	return []byte(EVENT_XML + Terminate)
}

func newXMLTEXTCommand(xpath string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", EVENT_XMLTEXT, xpath, Terminate))
}

func newTALLYCommand() []byte {
	return []byte(EVENT_TALLY + Terminate)
}

func newFUNCTIONCommand(name string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", EVENT_FUNCTION, name, Terminate))
}

func newSUBSCRIBECommand(event, option string) []byte {
	if option != "" {
		return []byte(fmt.Sprintf("%s %s %s%s", EVENT_SUBSCRIBE, event, option, Terminate))
	}
	return []byte(fmt.Sprintf("%s %s%s", EVENT_SUBSCRIBE, event, Terminate))
}

func newUNSUBSCRIBECommand(event string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", EVENT_UNSUBSCRIBE, event, Terminate))
}

func newQUITCommand() []byte {
	return []byte("QUIT" + Terminate)
}
