package vmixtcp

import "fmt"

const (
	// Terminate letter
	Terminate = "\r\n"
)

const (
	commandVersion     string = "VERSION"
	commandTally       string = "TALLY"
	commandFunction    string = "FUNCTION"
	commandXML         string = "XML"
	commandXMLText     string = "XMLTEXT"
	commandSubscribe   string = "SUBSCRIBE"
	commandUnsubscribe string = "UNSUBSCRIBE"
	commandQuit        string = "QUIT"
	commandActs        string = "ACTS"
)

const (
	// For SUBSCRIBE Event
	EventTally = commandTally
	EventActs  = commandActs
)

const (
	statusOK string = "OK"
	statusER string = "ER"
)

func newXMLCommand() []byte {
	return []byte(commandXML + Terminate)
}

func newXMLTEXTCommand(xpath string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", commandXMLText, xpath, Terminate))
}

func newTALLYCommand() []byte {
	return []byte(commandTally + Terminate)
}

func newFUNCTIONCommand(name string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", commandFunction, name, Terminate))
}

func newSUBSCRIBECommand(event, option string) []byte {
	if option != "" {
		return []byte(fmt.Sprintf("%s %s %s%s", commandSubscribe, event, option, Terminate))
	}
	return []byte(fmt.Sprintf("%s %s%s", commandSubscribe, event, Terminate))
}

func newUNSUBSCRIBECommand(event string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", commandUnsubscribe, event, Terminate))
}

func newQUITCommand() []byte {
	return []byte(commandQuit + Terminate)
}
