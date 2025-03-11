package vmixhttp

// SetDynamicValue1 Set Dynamic Value to use when specifying Dynamic in a shortcut value
func (v *Client) SetDynamicValue1(value string) error {
	params := make(map[string]string)
	params["Value"] = value
	return v.SendFunction("SetDynamicValue1", params)
}

// SetDynamicValue2 Set Dynamic Value to use when specifying Dynamic2 in a shortcut value
func (v *Client) SetDynamicValue2(value string) error {
	params := make(map[string]string)
	params["Value"] = value
	return v.SendFunction("SetDynamicValue2", params)
}

// SetDynamicValue3 Set Dynamic Value to use when specifying Dynamic3 in a shortcut value
func (v *Client) SetDynamicValue3(value string) error {
	params := make(map[string]string)
	params["Value"] = value
	return v.SendFunction("SetDynamicValue3", params)
}

// SetDynamicValue4 Set Dynamic Value to use when specifying Dynamic4 in a shortcut value
func (v *Client) SetDynamicValue4(value string) error {
	params := make(map[string]string)
	params["Value"] = value
	return v.SendFunction("SetDynamicValue4", params)
}
