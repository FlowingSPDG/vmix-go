package vmixhttp

import "strconv"

// MoveMultiViewOverlay Move Overlay in Input MultiView
func (v *Client) MoveMultiViewOverlay(input interface{}, fromIndex, toIndex uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(fromIndex)) + "," + strconv.Itoa(int(toIndex))
	return v.SendFunction("MoveMultiViewOverlay", params)
}

// MultiViewOverlay Toggle On/Off MultiView Overlay For Input At Index (starting from 1)
func (v *Client) MultiViewOverlay(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(index))
	return v.SendFunction("MultiViewOverlay", params)
}

// MultiViewOverlayOff Turn Off MultiView Overlay For Input At Index (starting from 1)
func (v *Client) MultiViewOverlayOff(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(index))
	return v.SendFunction("MultiViewOverlayOff", params)
}

// MultiViewOverlayOn Turn On MultiView Overlay For Input At Index (starting from 1)
func (v *Client) MultiViewOverlayOn(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(index))
	return v.SendFunction("MultiViewOverlayOn", params)
}

// SetMultiViewOverlay Change Layer in Input MultiView
func (v *Client) SetMultiViewOverlay(input interface{}, index uint, overlayInput interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	overlayIn, err := resolveInput(overlayInput)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(index)) + "," + overlayIn
	return v.SendFunction("SetMultiViewOverlay", params)
}
