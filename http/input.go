package vmixhttp

import (
	"fmt"
	"strconv"
)

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
	params["Value"] = fmt.Sprintf("%s|%s", inputType, value)
	return v.SendFunction("AddInput", params)
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

// SetInputLayer Set input layer (0-10)
func (v *Client) SetInputLayer(input any, index uint8, layer any) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	l, err := resolveInput(layer)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%d%s", index, l)
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
