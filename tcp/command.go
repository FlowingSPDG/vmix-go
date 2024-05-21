package vmixtcp

import "fmt"

const (
	terminate = "\r\n"
)

// TODO: 専用の型定義をする

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

func newTallyCommand() []byte {
	return []byte(commandTally + terminate)
}

func newFunctionCommand(name string, query string) []byte {
	return []byte(fmt.Sprintf("%s %s %s%s", commandFunction, name, query, terminate))
}

func newActsCommand(name string, input ...int) []byte {
	return []byte(fmt.Sprintf("%s %s %d%s", commandActs, name, input, terminate))
}

func newXMLCommand() []byte {
	return []byte(commandXML + terminate)
}

func newXMLTEXTCommand(xpath string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", commandXMLText, xpath, terminate))
}

func newSubscribeCommand(event, option string) []byte {
	if option != "" {
		return []byte(fmt.Sprintf("%s %s %s%s", commandSubscribe, event, option, terminate))
	}
	return []byte(fmt.Sprintf("%s %s%s", commandSubscribe, event, terminate))
}

func newUnsubscribeCommand(event string) []byte {
	return []byte(fmt.Sprintf("%s %s%s", commandUnsubscribe, event, terminate))
}

func newQuitCommand() []byte {
	return []byte(commandQuit + terminate)
}
