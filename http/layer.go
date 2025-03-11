package vmixhttp

import (
	"fmt"
	"strconv"
)

// 内部関数：レイヤー番号のバリデーション
func validateLayer(layer uint) uint {
	if layer < 1 {
		return 1
	}
	if layer > 10 {
		return 10
	}
	return layer
}

// 内部関数：レイヤー関連の共通パラメータ設定
func (v *Client) setLayerParam(funcName string, input interface{}, value *string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	if value != nil {
		params["Value"] = *value
	}
	return v.SendFunction(funcName, params)
}

// 内部関数：Pan値のバリデーション
func validatePan(pan float64) float64 {
	if pan < -2 {
		return -2
	}
	if pan > 2 {
		return 2
	}
	return pan
}

// 内部関数：Zoom値のバリデーション
func validateZoom(zoom uint16) uint16 {
	if zoom > 5 {
		return 5
	}
	return zoom
}

// 内部関数：crop値のバリデーション
func validateCrop(crop float64) float64 {
	if crop > 1 {
		return 1
	}
	if crop < 0 {
		return 0
	}
	return crop
}

// LayerOff Turn Layer Off
func (v *Client) LayerOff(input interface{}, layer uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(validateLayer(layer)))
	return v.SendFunction("LayerOff", params)
}

// LayerOn Turn Layer On
func (v *Client) LayerOn(input interface{}, layer uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = strconv.Itoa(int(validateLayer(layer)))
	return v.SendFunction("LayerOn", params)
}

// Layer1-10のCrop実装
func (v *Client) SetLayer1Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer1Crop", input, &value)
}

func (v *Client) SetLayer2Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer2Crop", input, &value)
}

func (v *Client) SetLayer3Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer3Crop", input, &value)
}

func (v *Client) SetLayer4Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer4Crop", input, &value)
}

func (v *Client) SetLayer5Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer5Crop", input, &value)
}

func (v *Client) SetLayer6Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer6Crop", input, &value)
}

func (v *Client) SetLayer7Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer7Crop", input, &value)
}

func (v *Client) SetLayer8Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer8Crop", input, &value)
}

func (v *Client) SetLayer9Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer9Crop", input, &value)
}

func (v *Client) SetLayer10Crop(input interface{}, left, top, right, bottom float64) error {
	value := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", validateCrop(left), validateCrop(top), validateCrop(right), validateCrop(bottom))
	return v.setLayerParam("SetLayer10Crop", input, &value)
}

func (v *Client) SetLayer1CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer1CropX1", input, &value)
}

func (v *Client) SetLayer2CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer2CropX1", input, &value)
}

func (v *Client) SetLayer3CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer3CropX1", input, &value)
}

func (v *Client) SetLayer4CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer4CropX1", input, &value)
}

func (v *Client) SetLayer5CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer5CropX1", input, &value)
}

func (v *Client) SetLayer6CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer6CropX1", input, &value)
}

func (v *Client) SetLayer7CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer7CropX1", input, &value)
}

func (v *Client) SetLayer8CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer8CropX1", input, &value)
}

func (v *Client) SetLayer9CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer9CropX1", input, &value)
}

func (v *Client) SetLayer10CropX1(input interface{}, x1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x1))
	return v.setLayerParam("SetLayer10CropX1", input, &value)
}

func (v *Client) SetLayer1CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer1CropX2", input, &value)
}

func (v *Client) SetLayer2CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer2CropX2", input, &value)
}

func (v *Client) SetLayer3CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer3CropX2", input, &value)
}

func (v *Client) SetLayer4CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer4CropX2", input, &value)
}

func (v *Client) SetLayer5CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer5CropX2", input, &value)
}

func (v *Client) SetLayer6CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer6CropX2", input, &value)
}

func (v *Client) SetLayer7CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer7CropX2", input, &value)
}

func (v *Client) SetLayer8CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer8CropX2", input, &value)
}

func (v *Client) SetLayer9CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer9CropX2", input, &value)
}

func (v *Client) SetLayer10CropX2(input interface{}, x2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(x2))
	return v.setLayerParam("SetLayer10CropX2", input, &value)
}

func (v *Client) SetLayer1CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer1CropY1", input, &value)
}

func (v *Client) SetLayer2CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer2CropY1", input, &value)
}

func (v *Client) SetLayer3CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer3CropY1", input, &value)
}

func (v *Client) SetLayer4CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer4CropY1", input, &value)
}

func (v *Client) SetLayer5CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer5CropY1", input, &value)
}

func (v *Client) SetLayer6CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer6CropY1", input, &value)
}

func (v *Client) SetLayer7CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer7CropY1", input, &value)
}

func (v *Client) SetLayer8CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer8CropY1", input, &value)
}

func (v *Client) SetLayer9CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer9CropY1", input, &value)
}

func (v *Client) SetLayer10CropY1(input interface{}, y1 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y1))
	return v.setLayerParam("SetLayer10CropY1", input, &value)
}

func (v *Client) SetLayer1CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer1CropY2", input, &value)
}

func (v *Client) SetLayer2CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer2CropY2", input, &value)
}

func (v *Client) SetLayer3CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer3CropY2", input, &value)
}

func (v *Client) SetLayer4CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer4CropY2", input, &value)
}

func (v *Client) SetLayer5CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer5CropY2", input, &value)
}

func (v *Client) SetLayer6CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer6CropY2", input, &value)
}

func (v *Client) SetLayer7CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer7CropY2", input, &value)
}

func (v *Client) SetLayer8CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer8CropY2", input, &value)
}

func (v *Client) SetLayer9CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer9CropY2", input, &value)
}

func (v *Client) SetLayer10CropY2(input interface{}, y2 float64) error {
	value := fmt.Sprintf("%.3f", validateCrop(y2))
	return v.setLayerParam("SetLayer10CropY2", input, &value)
}

// Layer1-10のPanX実装
func (v *Client) SetLayer1PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer1PanX", input, &value)
}

func (v *Client) SetLayer2PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer2PanX", input, &value)
}

func (v *Client) SetLayer3PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer3PanX", input, &value)
}

func (v *Client) SetLayer4PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer4PanX", input, &value)
}

func (v *Client) SetLayer5PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer5PanX", input, &value)
}

func (v *Client) SetLayer6PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer6PanX", input, &value)
}

func (v *Client) SetLayer7PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer7PanX", input, &value)
}

func (v *Client) SetLayer8PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer8PanX", input, &value)
}

func (v *Client) SetLayer9PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer9PanX", input, &value)
}

func (v *Client) SetLayer10PanX(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer10PanX", input, &value)
}

// Layer1-10のPanY実装
func (v *Client) SetLayer1PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer1PanY", input, &value)
}

func (v *Client) SetLayer2PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer2PanY", input, &value)
}

func (v *Client) SetLayer3PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer3PanY", input, &value)
}

func (v *Client) SetLayer4PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer4PanY", input, &value)
}

func (v *Client) SetLayer5PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer5PanY", input, &value)
}

func (v *Client) SetLayer6PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer6PanY", input, &value)
}

func (v *Client) SetLayer7PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer7PanY", input, &value)
}

func (v *Client) SetLayer8PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer8PanY", input, &value)
}

func (v *Client) SetLayer9PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer9PanY", input, &value)
}

func (v *Client) SetLayer10PanY(input interface{}, pan float64) error {
	value := fmt.Sprintf("%.2f", validatePan(pan))
	return v.setLayerParam("SetLayer10PanY", input, &value)
}

// Layer1-10のZoom実装
func (v *Client) SetLayer1Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer1Zoom", input, &value)
}

func (v *Client) SetLayer2Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer2Zoom", input, &value)
}

func (v *Client) SetLayer3Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer3Zoom", input, &value)
}

func (v *Client) SetLayer4Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer4Zoom", input, &value)
}

func (v *Client) SetLayer5Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer5Zoom", input, &value)
}

func (v *Client) SetLayer6Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer6Zoom", input, &value)
}

func (v *Client) SetLayer7Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer7Zoom", input, &value)
}

func (v *Client) SetLayer8Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer8Zoom", input, &value)
}

func (v *Client) SetLayer9Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer9Zoom", input, &value)
}

func (v *Client) SetLayer10Zoom(input interface{}, zoom uint16) error {
	value := fmt.Sprintf("%d", validateZoom(zoom))
	return v.setLayerParam("SetLayer10Zoom", input, &value)
}
