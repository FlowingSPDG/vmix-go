package vmixtcp

// Register goroutine callback event.
func (v *vmix) OnVersion(cb func(*VersionResponse, error)) {
	v.callbacks.version = cb
}

func (v *vmix) OnTally(cb func(*TallyResponse, error)) {
	v.callbacks.tally = cb
}

func (v *vmix) OnFunction(cb func(*FunctionResponse, error)) {
	v.callbacks.function = cb
}

func (v *vmix) OnActs(cb func(*ActsResponse, error)) {
	v.callbacks.acts = cb
}

func (v *vmix) OnXML(cb func(*XMLResponse, error)) {
	v.callbacks.xml = cb
}

func (v *vmix) OnXMLText(cb func(*XMLTextResponse, error)) {
	v.callbacks.xmltext = cb
}

func (v *vmix) OnSubscribe(cb func(*SubscribeResponse, error)) {
	v.callbacks.subscribe = cb
}

func (v *vmix) OnUnsubscribe(cb func(*UnsubscribeResponse, error)) {
	v.callbacks.unsubscribe = cb
}
