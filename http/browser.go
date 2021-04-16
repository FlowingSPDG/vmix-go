package vmixhttp

// BrowserBack ?
func (v *Client) BrowserBack(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserBack", params)
}

// BrowserForward ?
func (v *Client) BrowserForward(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserForward", params)
}

// BrowserKeyboardDisabled ?
func (v *Client) BrowserKeyboardDisabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserKeyboardDisabled", params)
}

// BrowserKeyboardEnabled ?
func (v *Client) BrowserKeyboardEnabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserKeyboardEnabled", params)
}

// BrowserMouseDisabled ?
func (v *Client) BrowserMouseDisabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserMouseDisabled", params)
}

// BrowserMouseEnabled ?
func (v *Client) BrowserMouseEnabled(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserMouseEnabled", params)
}

// BrowserNavigate ?
func (v *Client) BrowserNavigate(input interface{}, url string) error {
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
func (v *Client) BrowserReload(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("BrowserReload", params)
}
