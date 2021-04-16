package vmixgo

// BrowserBack ?
func (v *VmixHTTPClient) BrowserBack(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserBack", params)
}

// BrowserForward ?
func (v *VmixHTTPClient) BrowserForward(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserForward", params)
}

// BrowserKeyboardDisabled ?
func (v *VmixHTTPClient) BrowserKeyboardDisabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserKeyboardDisabled", params)
}

// BrowserKeyboardEnabled ?
func (v *VmixHTTPClient) BrowserKeyboardEnabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserKeyboardEnabled", params)
}

// BrowserMouseDisabled ?
func (v *VmixHTTPClient) BrowserMouseDisabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserMouseDisabled", params)
}

// BrowserMouseEnabled ?
func (v *VmixHTTPClient) BrowserMouseEnabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserMouseEnabled", params)
}

// BrowserNavigate ?
func (v *VmixHTTPClient) BrowserNavigate(input interface{}, url string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = url
	return v.SendFunction("BrowserNavigate", params)
}

// BrowserReload ?
func (v *VmixHTTPClient) BrowserReload(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserReload", params)
}
