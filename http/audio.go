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
