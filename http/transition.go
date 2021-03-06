package vmixhttp

import (
	"strconv"
)

func (v *Client) sendTransition(transition string, input interface{}, duration uint) error {
	params := make(map[string]string)
	if input != nil {
		in, err := resolveInput(input)
		if err != nil {
			return err
		}
		params["Input"] = in
	}
	params["Duration"] = strconv.Itoa(int(duration))
	if err := v.SendFunction(transition, params); err != nil {
		return err
	}
	return nil
}

// Fade transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Fade(scene interface{}, duration uint) error {
	return v.sendTransition("Fade", scene, duration)
}

// Zoom transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Zoom(scene interface{}, duration uint) error {
	return v.sendTransition("Zoom", scene, duration)
}

// Wipe transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Wipe(scene interface{}, duration uint) error {
	return v.sendTransition("Wipe", scene, duration)
}

// Slide transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Slide(scene interface{}, duration uint) error {
	return v.sendTransition("Slide", scene, duration)
}

// Fly transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Fly(scene interface{}, duration uint) error {
	return v.sendTransition("Fly", scene, duration)
}

// CrossZoom transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) CrossZoom(scene interface{}, duration uint) error {
	return v.sendTransition("CrossZoom", scene, duration)
}

// FlyRotate transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) FlyRotate(scene interface{}, duration uint) error {
	return v.sendTransition("FlyRotate", scene, duration)
}

// Cube transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Cube(scene interface{}, duration uint) error {
	return v.sendTransition("Cube", scene, duration)
}

// CubeZoom transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) CubeZoom(scene interface{}, duration uint) error {
	return v.sendTransition("CubeZoom", scene, duration)
}

// VerticalWipe transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) VerticalWipe(scene interface{}, duration uint) error {
	return v.sendTransition("VerticalWipe", scene, duration)
}

// VerticalSlide transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) VerticalSlide(scene interface{}, duration uint) error {
	return v.sendTransition("VerticalSlide", scene, duration)
}

// Merge transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Merge(scene interface{}, duration uint) error {
	return v.sendTransition("Merge", scene, duration)
}

// WipeReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) WipeReverse(scene interface{}, duration uint) error {
	return v.sendTransition("WipeReverse", scene, duration)
}

// SlideReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) SlideReverse(scene interface{}, duration uint) error {
	return v.sendTransition("SlideReverse", scene, duration)
}

// VerticalWipeReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) VerticalWipeReverse(scene interface{}, duration uint) error {
	return v.sendTransition("VerticalWipeReverse", scene, duration)
}

// VerticalSlideReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) VerticalSlideReverse(scene interface{}, duration uint) error {
	return v.sendTransition("VerticalSlideReverse", scene, duration)
}

// Cut You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Cut(scene interface{}) error {
	return v.sendTransition("Cut", scene, 0)
}

// CutDirect Cuts the input directly to Output without changing Preview. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) CutDirect(scene interface{}, duration uint) error {
	return v.sendTransition("CutDirect", scene, duration)
}

// FadeToBlack Toggle FTB On/Off
func (v *Client) FadeToBlack() error {
	return v.SendFunction("FadeToBlack", nil)
}

// QuickPlay You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) QuickPlay(scene interface{}) error {
	return v.sendTransition("QuickPlay", scene, 0)
}

// SetFader Set Master Fader T-Bar,255 will cut to Preview
func (v *Client) SetFader(fader uint8) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(fader))
	return v.SendFunction("SetFader", params)
}

// SetTransitionDuration1 Change Transition Duration for Button 1
func (v *Client) SetTransitionDuration1(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration1", params)
}

// SetTransitionDuration2 Change Transition Duration for Button 2
func (v *Client) SetTransitionDuration2(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration2", params)
}

// SetTransitionDuration3 Change Transition Duration for Button 3
func (v *Client) SetTransitionDuration3(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration3", params)
}

// SetTransitionDuration4 Change Transition Duration for Button 4
func (v *Client) SetTransitionDuration4(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration4", params)
}

// SetTransitionEffect1 Change Transition for Button 1
func (v *Client) SetTransitionEffect1(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect1", params)
}

// SetTransitionEffect2 Change Transition for Button 1
func (v *Client) SetTransitionEffect2(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect2", params)
}

// SetTransitionEffect3 Change Transition for Button 1
func (v *Client) SetTransitionEffect3(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect3", params)
}

// SetTransitionEffect4 Change Transition for Button 1
func (v *Client) SetTransitionEffect4(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect4", params)
}

// Stinger1 You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Stinger1(scene interface{}) error {
	return v.sendTransition("Stinger1", scene, 0)
}

// Stinger2 You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Client) Stinger2(scene interface{}) error {
	return v.sendTransition("Stinger2", scene, 0)
}

// Transition1 Clicks one of the four Transition buttons in the main vmix window
func (v *Client) Transition1() error {
	return v.SendFunction("Transition1", nil)
}

// Transition2 Clicks one of the four Transition buttons in the main vmix window
func (v *Client) Transition2() error {
	return v.SendFunction("Transition2", nil)
}

// Transition3 Clicks one of the four Transition buttons in the main vmix window
func (v *Client) Transition3() error {
	return v.SendFunction("Transition3", nil)
}

// Transition4 Clicks one of the four Transition buttons in the main vmix window
func (v *Client) Transition4() error {
	return v.SendFunction("Transition4", nil)
}
