package vmixhttp

import "strconv"

// PTZUpdatePreset Update PTZ Preset
func (v *Client) PTZUpdatePreset(input interface{}, preset uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(preset))
	return v.SendFunction("PTZUpdatePreset", params)
}

// PTZRecall Recall PTZ Preset
func (v *Client) PTZRecall(input interface{}, preset uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(preset))
	return v.SendFunction("PTZRecall", params)
}

// PTZFocusAutoToggle Toggle PTZ Auto Focus
func (v *Client) PTZFocusAutoToggle(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PTZFocusAutoToggle", params)
}

// PTZFocusManual Set PTZ Focus to Manual mode
func (v *Client) PTZFocusManual(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PTZFocusManual", params)
}

// PTZFocusAuto Set PTZ Focus to Auto mode
func (v *Client) PTZFocusAuto(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PTZFocusAuto", params)
}

// PTZHome Move PTZ to Home position
func (v *Client) PTZHome(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PTZHome", params)
}

// PTZZoomIn Start PTZ Zoom In
func (v *Client) PTZZoomIn(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PTZZoomIn", params)
}

// PTZZoomOut Start PTZ Zoom Out
func (v *Client) PTZZoomOut(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("PTZZoomOut", params)
}

// PTZPanTilt Control PTZ Pan/Tilt. Value format: "pan,tilt" where pan and tilt are between -1 and 1
func (v *Client) PTZPanTilt(input interface{}, pan, tilt float64) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	// Ensure pan and tilt values are between -1 and 1
	if pan < -1 {
		pan = -1
	} else if pan > 1 {
		pan = 1
	}
	if tilt < -1 {
		tilt = -1
	} else if tilt > 1 {
		tilt = 1
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.FormatFloat(pan, 'f', 2, 64) + "," + strconv.FormatFloat(tilt, 'f', 2, 64)
	return v.SendFunction("PTZPanTilt", params)
}
