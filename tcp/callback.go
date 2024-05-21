package vmixtcp

// Register goroutine callback event.
func (v *vmix) OnVersion(cb func(*VersionResponse)) {
	v.callbacks.version = cb
}

func (v *vmix) OnTally(cb func(*TallyResponse)) {
	v.callbacks.tally = cb
}

func (v *vmix) OnFunction(cb func(*FunctionResponse)) {
	v.callbacks.function = cb
}

func (v *vmix) OnActs(cb func(*ActsResponse)) {
	v.callbacks.acts = cb
}

func (v *vmix) OnXML(cb func(*XMLResponse)) {
	v.callbacks.xml = cb
}

func (v *vmix) OnXMLText(cb func(*XMLTextResponse)) {
	v.callbacks.xmltext = cb
}

func (v *vmix) OnSubscribe(cb func(*SubscribeResponse)) {
	v.callbacks.subscribe = cb
}

func (v *vmix) OnUnsubscribe(cb func(*UnsubscribeResponse)) {
	v.callbacks.unsubscribe = cb
}
