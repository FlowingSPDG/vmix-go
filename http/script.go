package vmixhttp

// ScriptStart ?
func (v *Client) ScriptStart(scriptname string) error {
	params := make(map[string]string)
	params["Value"] = scriptname
	return v.SendFunction("ScriptStart", params)
}

// ScriptStartDynamic Start a dynamic script using code specified as the Value
func (v *Client) ScriptStartDynamic(code string) error {
	params := make(map[string]string)
	params["Value"] = code
	return v.SendFunction("ScriptStartDynamic", params)
}

// ScriptStop ?
func (v *Client) ScriptStop(scriptname string) error {
	params := make(map[string]string)
	params["Value"] = scriptname
	return v.SendFunction("ScriptStop", params)
}

// ScriptStopAll ?
func (v *Client) ScriptStopAll() error {
	return v.SendFunction("ScriptStopAll", nil)
}

// ScriptStopDynamic ?
func (v *Client) ScriptStopDynamic() error {
	return v.SendFunction("ScriptStopDynamic", nil)
}
