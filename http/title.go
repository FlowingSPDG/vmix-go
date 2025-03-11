package vmixhttp

import "strconv"

// SetTitle Set Title text
func (v *Client) SetTitle(input interface{}, text string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = text
	return v.SendFunction("SetTitle", params)
}

// SetText Set any text field in a title
func (v *Client) SetText(input interface{}, selectedName, text string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["SelectedName"] = selectedName
	params["Value"] = text
	return v.SendFunction("SetText", params)
}

// TitleBeginAnimation Start Title Animation
func (v *Client) TitleBeginAnimation(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("TitleBeginAnimation", params)
}

// TitlePreset Load Title Preset by number
func (v *Client) TitlePreset(input interface{}, preset uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(preset))
	return v.SendFunction("TitlePreset", params)
}

// TitlePresetNext Load Next Title Preset
func (v *Client) TitlePresetNext(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("TitlePresetNext", params)
}

// TitlePresetPrevious Load Previous Title Preset
func (v *Client) TitlePresetPrevious(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("TitlePresetPrevious", params)
}
