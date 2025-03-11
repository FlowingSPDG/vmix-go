package vmixhttp

import (
	"fmt"
	"strconv"
)

// LayerOff Turn Off Layer For Input At Index (starting from 1)
func (v *Client) LayerOff(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(index)
	return v.SendFunction("LayerOff", params)
}

// LayerOn Turn On Layer For Input At Index (starting from 1)
func (v *Client) LayerOn(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(index)
	return v.SendFunction("LayerOn", params)
}

// LayerOnOff Toggle On/Off Layer For Input At Index (starting from 1)
func (v *Client) LayerOnOff(input interface{}, index uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(index)
	return v.SendFunction("LayerOnOff", params)
}

// SetLayerAnimated Change Layer Index to Input. Animate if input exists in another layer
func (v *Client) SetLayerAnimated(input interface{}, index uint, layerInput interface{}, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	layerIn, err := resolveInput(layerInput)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(index) + "," + layerIn + "," + itoa(durationMs)
	return v.SendFunction("SetLayerAnimated", params)
}

// SwapLayerAnimated Animate swapping the Layers in Input
func (v *Client) SwapLayerAnimated(input interface{}, fromIndex, toIndex, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(fromIndex) + "," + itoa(toIndex) + "," + itoa(durationMs)
	return v.SendFunction("SwapLayerAnimated", params)
}

// SetLayerCrop Change current Crop value of Input Layer
func (v *Client) SetLayerCrop(input interface{}, layerIndex uint, x1, y1, x2, y2 float64) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%f,%f,%f,%f", x1, y1, x2, y2)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"Crop", params)
}

// SetLayerCropX1 Change current Crop X1 value of Input Layer
func (v *Client) SetLayerCropX1(input interface{}, layerIndex uint, x1 float64) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%f", x1)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"CropX1", params)
}

// SetLayerCropX2 Change current Crop X2 value of Input Layer
func (v *Client) SetLayerCropX2(input interface{}, layerIndex uint, x2 float64) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%f", x2)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"CropX2", params)
}

// SetLayerCropY1 Change current Crop Y1 value of Input Layer
func (v *Client) SetLayerCropY1(input interface{}, layerIndex uint, y1 float64) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%f", y1)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"CropY1", params)
}

// SetLayerCropY2 Change current Crop Y2 value of Input Layer
func (v *Client) SetLayerCropY2(input interface{}, layerIndex uint, y2 float64) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%f", y2)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"CropY2", params)
}

// SetLayerHeight Change current Height value of Input Layer
func (v *Client) SetLayerHeight(input interface{}, layerIndex uint, pixels int) error {
	if pixels < -4096 {
		pixels = -4096
	} else if pixels > 4096 {
		pixels = 4096
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(pixels)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"Height", params)
}

// SetLayerPanX Change current PanX value of Input Layer
func (v *Client) SetLayerPanX(input interface{}, layerIndex uint, pan float64) error {
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
	params["Value"] = fmt.Sprintf("%f", pan)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"PanX", params)
}

// SetLayerPanY Change current PanY value of Input Layer
func (v *Client) SetLayerPanY(input interface{}, layerIndex uint, pan float64) error {
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
	params["Value"] = fmt.Sprintf("%f", pan)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"PanY", params)
}

// SetLayerRectangle Change current Rectangle values of Input Layer in pixels
func (v *Client) SetLayerRectangle(input interface{}, layerIndex uint, x, y, width, height int) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"Rectangle", params)
}

// SetLayerWidth Change current Width value of Input Layer
func (v *Client) SetLayerWidth(input interface{}, layerIndex uint, pixels int) error {
	if pixels < -4096 {
		pixels = -4096
	} else if pixels > 4096 {
		pixels = 4096
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(pixels)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"Width", params)
}

// SetLayerX Change current X value of Input Layer
func (v *Client) SetLayerX(input interface{}, layerIndex uint, pixels int) error {
	if pixels < -4096 {
		pixels = -4096
	} else if pixels > 4096 {
		pixels = 4096
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(pixels)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"X", params)
}

// SetLayerY Change current Y value of Input Layer
func (v *Client) SetLayerY(input interface{}, layerIndex uint, pixels int) error {
	if pixels < -4096 {
		pixels = -4096
	} else if pixels > 4096 {
		pixels = 4096
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(pixels)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"Y", params)
}

// SetLayerZoom Change current Zoom level of Input Layer
func (v *Client) SetLayerZoom(input interface{}, layerIndex uint, zoom float64) error {
	if zoom < 0 {
		zoom = 0
	} else if zoom > 5 {
		zoom = 5
	}
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = fmt.Sprintf("%f", zoom)
	return v.SendFunction("SetLayer"+itoa(layerIndex)+"Zoom", params)
}
