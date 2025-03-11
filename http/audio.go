package vmixhttp

// AudioAutoMute Automatically mute the specified input when it is not in preview or active
func (v *Client) AudioAutoMute(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioAutoMute", params)
}

// AudioAutoUnMute Turn off automatic muting for the specified input
func (v *Client) AudioAutoUnMute(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioAutoUnMute", params)
}

// AudioBus Assign input to specified bus A,B,C,D,E,F,G or Master
func (v *Client) AudioBus(input interface{}, bus string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = bus
	return v.SendFunction("AudioBus", params)
}

// AudioBusOff Remove input from specified bus A,B,C,D,E,F,G or Master
func (v *Client) AudioBusOff(input interface{}, bus string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = bus
	return v.SendFunction("AudioBusOff", params)
}

// AudioBusOn Add input to specified bus A,B,C,D,E,F,G or Master
func (v *Client) AudioBusOn(input interface{}, bus string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = bus
	return v.SendFunction("AudioBusOn", params)
}

// AudioMixerShowHide Toggle Audio Mixer visibility
func (v *Client) AudioMixerShowHide() error {
	return v.SendFunction("AudioMixerShowHide", nil)
}

// AudioMixerShow Show Audio Mixer
func (v *Client) AudioMixerShow() error {
	return v.SendFunction("AudioMixerShow", nil)
}

// AudioMixerHide Hide Audio Mixer
func (v *Client) AudioMixerHide() error {
	return v.SendFunction("AudioMixerHide", nil)
}

// Audio Toggle Audio Mute On/Off
func (v *Client) Audio(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Audio", params)
}

// AudioAuto Toggle Audio Auto On/Off
func (v *Client) AudioAuto(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioAuto", params)
}

// AudioAutoOff Turn off Audio Auto
func (v *Client) AudioAutoOff(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioAutoOff", params)
}

// AudioAutoOn Turn on Audio Auto
func (v *Client) AudioAutoOn(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioAutoOn", params)
}

// AudioOff Turn off Audio
func (v *Client) AudioOff(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioOff", params)
}

// AudioOn Turn on Audio
func (v *Client) AudioOn(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("AudioOn", params)
}

// AudioPluginOff Turn off Audio Plugin, starting from 1
func (v *Client) AudioPluginOff(input interface{}, pluginNumber uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("AudioPluginOff", params)
}

// AudioPluginOn Turn on Audio Plugin, starting from 1
func (v *Client) AudioPluginOn(input interface{}, pluginNumber uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("AudioPluginOn", params)
}

// AudioPluginOnOff Toggle on/off Audio Plugin, starting from 1
func (v *Client) AudioPluginOnOff(input interface{}, pluginNumber uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("AudioPluginOnOff", params)
}

// AudioPluginShow Show Audio Plugin Editor, starting from 1
func (v *Client) AudioPluginShow(input interface{}, pluginNumber uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("AudioPluginShow", params)
}

// BusAAudio Toggle Bus A Audio On/Off
func (v *Client) BusAAudio() error {
	return v.SendFunction("BusAAudio", nil)
}

// BusAAudioOff Turn off Bus A Audio
func (v *Client) BusAAudioOff() error {
	return v.SendFunction("BusAAudioOff", nil)
}

// BusAAudioOn Turn on Bus A Audio
func (v *Client) BusAAudioOn() error {
	return v.SendFunction("BusAAudioOn", nil)
}

// BusAAudioPluginOff Turn off Audio Plugin on Bus A, starting from 1
func (v *Client) BusAAudioPluginOff(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusAAudioPluginOff", params)
}

// BusAAudioPluginOn Turn on Audio Plugin on Bus A, starting from 1
func (v *Client) BusAAudioPluginOn(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusAAudioPluginOn", params)
}

// BusAAudioPluginOnOff Toggle on/off Audio Plugin on Bus A, starting from 1
func (v *Client) BusAAudioPluginOnOff(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusAAudioPluginOnOff", params)
}

// BusAAudioPluginShow Show Audio Plugin Editor on Bus A, starting from 1
func (v *Client) BusAAudioPluginShow(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusAAudioPluginShow", params)
}

// BusBAudio Toggle Bus B Audio On/Off
func (v *Client) BusBAudio() error {
	return v.SendFunction("BusBAudio", nil)
}

// BusBAudioOff Turn off Bus B Audio
func (v *Client) BusBAudioOff() error {
	return v.SendFunction("BusBAudioOff", nil)
}

// BusBAudioOn Turn on Bus B Audio
func (v *Client) BusBAudioOn() error {
	return v.SendFunction("BusBAudioOn", nil)
}

// BusBAudioPluginOff Turn off Audio Plugin on Bus B, starting from 1
func (v *Client) BusBAudioPluginOff(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusBAudioPluginOff", params)
}

// BusBAudioPluginOn Turn on Audio Plugin on Bus B, starting from 1
func (v *Client) BusBAudioPluginOn(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusBAudioPluginOn", params)
}

// BusBAudioPluginOnOff Toggle on/off Audio Plugin on Bus B, starting from 1
func (v *Client) BusBAudioPluginOnOff(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusBAudioPluginOnOff", params)
}

// BusBAudioPluginShow Show Audio Plugin Editor on Bus B, starting from 1
func (v *Client) BusBAudioPluginShow(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("BusBAudioPluginShow", params)
}

// BusXAudio Toggle Bus Audio On/Off for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXAudio(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXAudio", params)
}

// BusXAudioOff Turn off Bus Audio for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXAudioOff(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXAudioOff", params)
}

// BusXAudioOn Turn on Bus Audio for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXAudioOn(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXAudioOn", params)
}

// BusXAudioPluginOff Turn off Audio Plugin for specified bus and plugin number
func (v *Client) BusXAudioPluginOff(bus string, pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = bus + "," + itoa(pluginNumber)
	return v.SendFunction("BusXAudioPluginOff", params)
}

// BusXAudioPluginOn Turn on Audio Plugin for specified bus and plugin number
func (v *Client) BusXAudioPluginOn(bus string, pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = bus + "," + itoa(pluginNumber)
	return v.SendFunction("BusXAudioPluginOn", params)
}

// BusXAudioPluginOnOff Toggle on/off Audio Plugin for specified bus and plugin number
func (v *Client) BusXAudioPluginOnOff(bus string, pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = bus + "," + itoa(pluginNumber)
	return v.SendFunction("BusXAudioPluginOnOff", params)
}

// BusXAudioPluginShow Show Audio Plugin Editor for specified bus and plugin number
func (v *Client) BusXAudioPluginShow(bus string, pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = bus + "," + itoa(pluginNumber)
	return v.SendFunction("BusXAudioPluginShow", params)
}

// BusXSendToMaster Toggle sending Bus to Master for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXSendToMaster(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXSendToMaster", params)
}

// BusXSendToMasterOff Turn off sending Bus to Master for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXSendToMasterOff(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXSendToMasterOff", params)
}

// BusXSendToMasterOn Turn on sending Bus to Master for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXSendToMasterOn(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXSendToMasterOn", params)
}

// BusXSolo Toggle Solo for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXSolo(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXSolo", params)
}

// BusXSoloOff Turn off Solo for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXSoloOff(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXSoloOff", params)
}

// BusXSoloOn Turn on Solo for specified bus M,A,B,C,D,E,F,G
func (v *Client) BusXSoloOn(bus string) error {
	params := make(map[string]string)
	params["Value"] = bus
	return v.SendFunction("BusXSoloOn", params)
}

// MasterAudio Toggle Master Audio On/Off
func (v *Client) MasterAudio() error {
	return v.SendFunction("MasterAudio", nil)
}

// MasterAudioOff Turn off Master Audio
func (v *Client) MasterAudioOff() error {
	return v.SendFunction("MasterAudioOff", nil)
}

// MasterAudioOn Turn on Master Audio
func (v *Client) MasterAudioOn() error {
	return v.SendFunction("MasterAudioOn", nil)
}

// MasterAudioPluginOff Turn off Audio Plugin on Master, starting from 1
func (v *Client) MasterAudioPluginOff(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("MasterAudioPluginOff", params)
}

// MasterAudioPluginOn Turn on Audio Plugin on Master, starting from 1
func (v *Client) MasterAudioPluginOn(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("MasterAudioPluginOn", params)
}

// MasterAudioPluginOnOff Toggle on/off Audio Plugin on Master, starting from 1
func (v *Client) MasterAudioPluginOnOff(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("MasterAudioPluginOnOff", params)
}

// MasterAudioPluginShow Show Audio Plugin Editor on Master, starting from 1
func (v *Client) MasterAudioPluginShow(pluginNumber uint) error {
	params := make(map[string]string)
	params["Value"] = itoa(pluginNumber)
	return v.SendFunction("MasterAudioPluginShow", params)
}

// Solo Toggle Solo On/Off
func (v *Client) Solo(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("Solo", params)
}

// SoloAllOff Turn off Solo for all Inputs and Busses
func (v *Client) SoloAllOff() error {
	return v.SendFunction("SoloAllOff", nil)
}

// SoloOff Turn off Solo
func (v *Client) SoloOff(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("SoloOff", params)
}

// SoloOn Turn on Solo
func (v *Client) SoloOn(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("SoloOn", params)
}

// SoloPFL Toggle between AFL or PFL mode for Solo
func (v *Client) SoloPFL(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("SoloPFL", params)
}

// SoloPFLOff Turn off PFL mode for Solo
func (v *Client) SoloPFLOff(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("SoloPFLOff", params)
}

// SoloPFLOn Turn on PFL mode for Solo
func (v *Client) SoloPFLOn(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("SoloPFLOn", params)
}
