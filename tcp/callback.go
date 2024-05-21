package vmixtcp

// Register goroutine callback event.
func (v *Vmix) OnVersion(cb func(*VersionResponse)) {
	v.callbacks.version = cb
}

func (v *Vmix) OnTally(cb func(*TallyResponse)) {
	v.callbacks.tally = cb
}

func (v *Vmix) OnFunction(cb func(*FunctionResponse)) {
	v.callbacks.function = cb
}

func (v *Vmix) OnActs(cb func(*ActsResponse)) {
	v.callbacks.acts = cb
}

func (v *Vmix) OnXML(cb func(*XMLResponse)) {
	v.callbacks.xml = cb
}

func (v *Vmix) OnXMLText(cb func(*XMLTextResponse)) {
	v.callbacks.xmltext = cb
}

func (v *Vmix) OnSubscribe(cb func(*SubscribeResponse)) {
	v.callbacks.subscribe = cb
}

func (v *Vmix) OnUnsubscribe(cb func(*UnsubscribeResponse)) {
	v.callbacks.unsubscribe = cb
}
