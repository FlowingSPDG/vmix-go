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

// AddInput Add a new input
func (v *Client) AddInput(inputType, value string) error {
	params := make(map[string]string)
	params["Value"] = value
	return v.SendFunction("AddInput"+inputType, params)
}

// SetCrop Set input crop (Left,Top,Right,Bottom) in pixels
func (v *Client) SetCrop(input interface{}, left, top, right, bottom uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(left)) + "," + strconv.Itoa(int(top)) + "," + strconv.Itoa(int(right)) + "," + strconv.Itoa(int(bottom))
	return v.SendFunction("SetCrop", params)
}

// SetBalance Set input balance (-1 to 1)
func (v *Client) SetBalance(input interface{}, balance float64) error {
	if balance < -1 {
		balance = -1
	} else if balance > 1 {
		balance = 1
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.FormatFloat(balance, 'f', 2, 64)
	return v.SendFunction("SetBalance", params)
}

// SetColorLevel Set input color level (0-400)
func (v *Client) SetColorLevel(input interface{}, level uint16) error {
	if level > 400 {
		level = 400
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(level))
	return v.SendFunction("SetColorLevel", params)
}

// SetColorCorrectionPreset Set input color correction preset (1-12)
func (v *Client) SetColorCorrectionPreset(input interface{}, preset uint8) error {
	if preset < 1 {
		preset = 1
	} else if preset > 12 {
		preset = 12
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(preset))
	return v.SendFunction("SetColorCorrectionPreset", params)
}

// SetLayer Set input layer (0-10)
func (v *Client) SetLayer(input interface{}, layer uint8) error {
	if layer > 10 {
		layer = 10
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(layer))
	return v.SendFunction("SetLayer", params)
}

// SetZoom Set input zoom (0-500)
func (v *Client) SetZoom(input interface{}, zoom uint16) error {
	if zoom > 500 {
		zoom = 500
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(zoom))
	return v.SendFunction("SetZoom", params)
}

// SetPanX Set input pan X (-2 to 2)
func (v *Client) SetPanX(input interface{}, pan float64) error {
	if pan < -2 {
		pan = -2
	} else if pan > 2 {
		pan = 2
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.FormatFloat(pan, 'f', 2, 64)
	return v.SendFunction("SetPanX", params)
}

// SetPanY Set input pan Y (-2 to 2)
func (v *Client) SetPanY(input interface{}, pan float64) error {
	if pan < -2 {
		pan = -2
	} else if pan > 2 {
		pan = 2
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.FormatFloat(pan, 'f', 2, 64)
	return v.SendFunction("SetPanY", params)
}

// SetRotation Set input rotation (0-360)
func (v *Client) SetRotation(input interface{}, rotation uint16) error {
	if rotation > 360 {
		rotation = 360
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(rotation))
	return v.SendFunction("SetRotation", params)
}

// Loop Toggle input loop on/off
func (v *Client) Loop(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Loop", params)
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

// Undo Undo closing Input
func (v *Client) Undo() error {
	return v.SendFunction("Undo", nil)
}

// ColourCorrectionAuto Basic Auto Colour Correction
func (v *Client) ColourCorrectionAuto(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ColourCorrectionAuto", params)
}

// ColourCorrectionReset Reset Colour Correction to Default Values
func (v *Client) ColourCorrectionReset(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ColourCorrectionReset", params)
}

// Effect1 Toggle Effect 1 On/Off
func (v *Client) Effect1(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect1", params)
}

// Effect1Off Turn Effect 1 Off
func (v *Client) Effect1Off(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect1Off", params)
}

// Effect1On Turn Effect 1 On
func (v *Client) Effect1On(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect1On", params)
}

// Effect2 Toggle Effect 2 On/Off
func (v *Client) Effect2(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect2", params)
}

// Effect2Off Turn Effect 2 Off
func (v *Client) Effect2Off(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect2Off", params)
}

// Effect2On Turn Effect 2 On
func (v *Client) Effect2On(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect2On", params)
}

// Effect3 Toggle Effect 3 On/Off
func (v *Client) Effect3(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect3", params)
}

// Effect3Off Turn Effect 3 Off
func (v *Client) Effect3Off(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect3Off", params)
}

// Effect3On Turn Effect 3 On
func (v *Client) Effect3On(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect3On", params)
}

// Effect4 Toggle Effect 4 On/Off
func (v *Client) Effect4(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect4", params)
}

// Effect4Off Turn Effect 4 Off
func (v *Client) Effect4Off(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect4Off", params)
}

// Effect4On Turn Effect 4 On
func (v *Client) Effect4On(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Effect4On", params)
}

// InputPreviewHide Hides large preview of input
func (v *Client) InputPreviewHide(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("InputPreviewHide", params)
}

// InputPreviewShow Shows large preview of input
func (v *Client) InputPreviewShow(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("InputPreviewShow", params)
}

// InputPreviewShowHide Toggles large preview of input
func (v *Client) InputPreviewShowHide(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("InputPreviewShowHide", params)
}

// ListAdd Add Filename to List
func (v *Client) ListAdd(input interface{}, filename string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = filename
	return v.SendFunction("ListAdd", params)
}

// ListExport Export List as M3U to Filename
func (v *Client) ListExport(input interface{}, filename string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = filename
	return v.SendFunction("ListExport", params)
}

// ListPlayOut Play out the current list
func (v *Client) ListPlayOut(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ListPlayOut", params)
}

// ListRemove Remove from List by Index starting from 1
func (v *Client) ListRemove(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(index))
	return v.SendFunction("ListRemove", params)
}

// ListRemoveAll Remove all items from List
func (v *Client) ListRemoveAll(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ListRemoveAll", params)
}

// ListShowHide Toggle List visibility
func (v *Client) ListShowHide(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ListShowHide", params)
}

// ListShuffle Shuffle (randomize) List
func (v *Client) ListShuffle(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ListShuffle", params)
}

// MirrorOff Turn off Mirror
func (v *Client) MirrorOff(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("MirrorOff", params)
}

// MirrorOn Turn on Mirror
func (v *Client) MirrorOn(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("MirrorOn", params)
}
