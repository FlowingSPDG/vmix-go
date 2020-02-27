package vmixgo

import (
	"fmt"
	"reflect"
	"strconv"
)

func (v *Vmix) sendTransition(transition string, input interface{}) error {
	s := reflect.ValueOf(input)
	if !s.IsValid() {
		if err := v.SendFunction(transition, nil); err != nil {
			return err
		}
	}
	// fmt.Printf("type : %v", s.Type().String())
	switch s.Type().String() {
	case "int":
		params := make(map[string]string)
		params["Input"] = strconv.Itoa(input.(int))
		if err := v.SendFunction(transition, params); err != nil {
			return err
		}
	case "string":
		params := make(map[string]string)
		params["Input"] = input.(string)
		if err := v.SendFunction(transition, params); err != nil {
			return err
		}
	case "vmixgo.Input":
		params := make(map[string]string)
		in := input.(Input)
		params["Input"] = in.Key
		if err := v.SendFunction(transition, params); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Interface type not correct")
	}
	return nil
}

// Fade transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Fade(scene interface{}) error {
	return v.sendTransition("Fade", scene)
}

// Zoom transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Zoom(scene interface{}) error {
	return v.sendTransition("Zoom", scene)
}

// Wipe transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Wipe(scene interface{}) error {
	return v.sendTransition("Wipe", scene)
}

// Slide transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Slide(scene interface{}) error {
	return v.sendTransition("Slide", scene)
}

// Fly transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Fly(scene interface{}) error {
	return v.sendTransition("Fly", scene)
}

// CrossZoom transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) CrossZoom(scene interface{}) error {
	return v.sendTransition("CrossZoom", scene)
}

// FlyRotate transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) FlyRotate(scene interface{}) error {
	return v.sendTransition("FlyRotate", scene)
}

// Cube transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Cube(scene interface{}) error {
	return v.sendTransition("Cube", scene)
}

// CubeZoom transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) CubeZoom(scene interface{}) error {
	return v.sendTransition("CubeZoom", scene)
}

// VerticalWipe transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) VerticalWipe(scene interface{}) error {
	return v.sendTransition("VerticalWipe", scene)
}

// VerticalSlide transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) VerticalSlide(scene interface{}) error {
	return v.sendTransition("VerticalSlide", scene)
}

// Merge transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Merge(scene interface{}) error {
	return v.sendTransition("Merge", scene)
}

// WipeReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) WipeReverse(scene interface{}) error {
	return v.sendTransition("WipeReverse", scene)
}

// SlideReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) SlideReverse(scene interface{}) error {
	return v.sendTransition("SlideReverse", scene)
}

// VerticalWipeReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) VerticalWipeReverse(scene interface{}) error {
	return v.sendTransition("VerticalWipeReverse", scene)
}

// VerticalSlideReverse transition. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) VerticalSlideReverse(scene interface{}) error {
	return v.sendTransition("VerticalSlideReverse", scene)
}

// Cut You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Cut(scene interface{}) error {
	return v.sendTransition("Cut", scene)
}

// CutDirect Cuts the input directly to Output without changing Preview. You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) CutDirect(scene interface{}) error {
	return v.sendTransition("CutDirect", scene)
}

// FadeToBlack Toggle FTB On/Off
func (v *Vmix) FadeToBlack() error {
	return v.SendFunction("FadeToBlack", nil)
}

// QuickPlay You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) QuickPlay(scene interface{}) error {
	return v.sendTransition("QuickPlay", scene)
}

// SetFader Set Master Fader T-Bar,255 will cut to Preview
func (v *Vmix) SetFader(fader uint8) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(fader))
	return v.SendFunction("SetFader", params)
}

// SetTransitionDuration1 Change Transition Duration for Button 1
func (v *Vmix) SetTransitionDuration1(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration1", params)
}

// SetTransitionDuration2 Change Transition Duration for Button 2
func (v *Vmix) SetTransitionDuration2(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration2", params)
}

// SetTransitionDuration3 Change Transition Duration for Button 3
func (v *Vmix) SetTransitionDuration3(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration3", params)
}

// SetTransitionDuration4 Change Transition Duration for Button 4
func (v *Vmix) SetTransitionDuration4(duration uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(duration))
	return v.SendFunction("SetTransitionDuration4", params)
}

// SetTransitionEffect1 Change Transition for Button 1
func (v *Vmix) SetTransitionEffect1(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect1", params)
}

// SetTransitionEffect2 Change Transition for Button 1
func (v *Vmix) SetTransitionEffect2(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect2", params)
}

// SetTransitionEffect3 Change Transition for Button 1
func (v *Vmix) SetTransitionEffect3(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect3", params)
}

// SetTransitionEffect4 Change Transition for Button 1
func (v *Vmix) SetTransitionEffect4(transition string) error {
	params := make(map[string]string)
	params["Value"] = transition
	return v.SendFunction("SetTransitionEffect4", params)
}

// Stinger1 You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Stinger1(scene interface{}) error {
	return v.sendTransition("Stinger1", scene)
}

// Stinger2 You can use string scene-key, int scene-number or vmixgo.Scene struct.
func (v *Vmix) Stinger2(scene interface{}) error {
	return v.sendTransition("Stinger2", scene)
}

// Transition1 Clicks one of the four Transition buttons in the main vmix window
func (v *Vmix) Transition1() error {
	return v.SendFunction("Transition1", nil)
}

// Transition2 Clicks one of the four Transition buttons in the main vmix window
func (v *Vmix) Transition2() error {
	return v.SendFunction("Transition2", nil)
}

// Transition3 Clicks one of the four Transition buttons in the main vmix window
func (v *Vmix) Transition3() error {
	return v.SendFunction("Transition3", nil)
}

// Transition4 Clicks one of the four Transition buttons in the main vmix window
func (v *Vmix) Transition4() error {
	return v.SendFunction("Transition4", nil)
}
