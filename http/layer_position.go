package vmixhttp

import (
	"fmt"
	"strconv"
)

func validateHeight(h int) int {
	if h > 4096 {
		return 4096
	}
	if h < -4096 {
		return -4096
	}
	return h
}

func validateWidth(w int) int {
	if w > 4096 {
		return 4096
	}
	if w < -4096 {
		return -4096
	}
	return w
}
func (v *Client) SetLayer1Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer1Height", input, &value)
}

func (v *Client) SetLayer2Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer2Height", input, &value)
}

func (v *Client) SetLayer3Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer3Height", input, &value)
}

func (v *Client) SetLayer4Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer4Height", input, &value)
}

func (v *Client) SetLayer5Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer5Height", input, &value)
}

func (v *Client) SetLayer6Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer6Height", input, &value)
}

func (v *Client) SetLayer7Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer7Height", input, &value)
}

func (v *Client) SetLayer8Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer8Height", input, &value)
}

func (v *Client) SetLayer9Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer9Height", input, &value)
}

func (v *Client) SetLayer10Height(input interface{}, height int) error {
	value := strconv.Itoa(validateHeight(height))
	return v.setLayerParam("SetLayer10Height", input, &value)
}

// Layer1-10のWidth実装
func (v *Client) SetLayer1Width(input interface{}, width int) error {
	value := strconv.Itoa(validateWidth(width))
	return v.setLayerParam("SetLayer1Width", input, &value)
}

func (v *Client) SetLayer2Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer2Width", input, &value)
}

func (v *Client) SetLayer3Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer3Width", input, &value)
}

func (v *Client) SetLayer4Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer4Width", input, &value)
}

func (v *Client) SetLayer5Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer5Width", input, &value)
}

func (v *Client) SetLayer6Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer6Width", input, &value)
}

func (v *Client) SetLayer7Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer7Width", input, &value)
}

func (v *Client) SetLayer8Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer8Width", input, &value)
}

func (v *Client) SetLayer9Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer9Width", input, &value)
}

func (v *Client) SetLayer10Width(input interface{}, width uint) error {
	value := strconv.Itoa(int(width))
	return v.setLayerParam("SetLayer10Width", input, &value)
}

// Layer1-10のX座標実装
func (v *Client) SetLayer1X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer1X", input, &value)
}

func (v *Client) SetLayer2X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer2X", input, &value)
}

func (v *Client) SetLayer3X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer3X", input, &value)
}

func (v *Client) SetLayer4X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer4X", input, &value)
}

func (v *Client) SetLayer5X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer5X", input, &value)
}

func (v *Client) SetLayer6X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer6X", input, &value)
}

func (v *Client) SetLayer7X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer7X", input, &value)
}

func (v *Client) SetLayer8X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer8X", input, &value)
}

func (v *Client) SetLayer9X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer9X", input, &value)
}

func (v *Client) SetLayer10X(input interface{}, x int) error {
	value := strconv.Itoa(x)
	return v.setLayerParam("SetLayer10X", input, &value)
}

// Layer1-10のY座標実装
func (v *Client) SetLayer1Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer1Y", input, &value)
}

func (v *Client) SetLayer2Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer2Y", input, &value)
}

func (v *Client) SetLayer3Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer3Y", input, &value)
}

func (v *Client) SetLayer4Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer4Y", input, &value)
}

func (v *Client) SetLayer5Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer5Y", input, &value)
}

func (v *Client) SetLayer6Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer6Y", input, &value)
}

func (v *Client) SetLayer7Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer7Y", input, &value)
}

func (v *Client) SetLayer8Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer8Y", input, &value)
}

func (v *Client) SetLayer9Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer9Y", input, &value)
}

func (v *Client) SetLayer10Y(input interface{}, y int) error {
	value := strconv.Itoa(y)
	return v.setLayerParam("SetLayer10Y", input, &value)
}

// Layer1-10のRectangle実装
func (v *Client) SetLayer1Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer1Rectangle", input, &value)
}

func (v *Client) SetLayer2Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer2Rectangle", input, &value)
}

func (v *Client) SetLayer3Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer3Rectangle", input, &value)
}

func (v *Client) SetLayer4Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer4Rectangle", input, &value)
}

func (v *Client) SetLayer5Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer5Rectangle", input, &value)
}

func (v *Client) SetLayer6Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer6Rectangle", input, &value)
}

func (v *Client) SetLayer7Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer7Rectangle", input, &value)
}

func (v *Client) SetLayer8Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer8Rectangle", input, &value)
}

func (v *Client) SetLayer9Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer9Rectangle", input, &value)
}

func (v *Client) SetLayer10Rectangle(input interface{}, x, y, width, height uint) error {
	value := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	return v.setLayerParam("SetLayer10Rectangle", input, &value)
}
