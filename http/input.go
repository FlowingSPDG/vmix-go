package vmixhttp

import "strconv"

// SelectInput Select Input by Number or Name
func (v *Client) SelectInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("SelectInput", params)
}

// PreviewInput Preview Input by Number or Name
func (v *Client) PreviewInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PreviewInput", params)
}

// ActiveInput Set Input to Program by Number or Name
func (v *Client) ActiveInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ActiveInput", params)
}

// RestartInput Restart Input
func (v *Client) RestartInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("RestartInput", params)
}

// StopInput Stop Input
func (v *Client) StopInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("StopInput", params)
}

// PlayInput Play Input
func (v *Client) PlayInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PlayInput", params)
}

// PauseInput Pause Input
func (v *Client) PauseInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PauseInput", params)
}

// SetInputPosition Set Input Position in Pixels (X,Y)
func (v *Client) SetInputPosition(input interface{}, x, y int) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return v.SendFunction("SetPosition", params)
}

// SetInputVolume Set Input Volume (0-100)
func (v *Client) SetInputVolume(input interface{}, volume uint8) error {
	if volume > 100 {
		volume = 100
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(volume))
	return v.SendFunction("SetVolume", params)
}

// NextInput Switch to Next Input
func (v *Client) NextInput() error {
	return v.SendFunction("NextInput", nil)
}

// PreviousInput Switch to Previous Input
func (v *Client) PreviousInput() error {
	return v.SendFunction("PreviousInput", nil)
}

// FirstInput Switch to First Input
func (v *Client) FirstInput() error {
	return v.SendFunction("FirstInput", nil)
}

// LastInput Switch to Last Input
func (v *Client) LastInput() error {
	return v.SendFunction("LastInput", nil)
}

// SetInputName Set the name of an input
func (v *Client) SetInputName(input interface{}, name string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = name
	return v.SendFunction("SetInputName", params)
}

// ResetInput Reset an input to its default state
func (v *Client) ResetInput(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ResetInput", params)
}
