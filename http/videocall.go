package vmixhttp

// VideoCallAudioSource Set audio source for video call
func (v *Client) VideoCallAudioSource(input interface{}, source string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = source
	return v.SendFunction("VideoCallAudioSource", params)
}

// VideoCallConnect Connect to video call
func (v *Client) VideoCallConnect(input interface{}, name, password string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = name + "," + password
	return v.SendFunction("VideoCallConnect", params)
}

// VideoCallReconnect Reconnect video call
func (v *Client) VideoCallReconnect(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("VideoCallReconnect", params)
}

// VideoCallVideoSource Set video source for video call
func (v *Client) VideoCallVideoSource(input interface{}, source string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = source
	return v.SendFunction("VideoCallVideoSource", params)
}

// ZoomJoinMeeting Join Zoom meeting
func (v *Client) ZoomJoinMeeting(input interface{}, meetingID, password string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = meetingID + "," + password
	return v.SendFunction("ZoomJoinMeeting", params)
}

// ZoomMuteSelf Mute self in Zoom meeting
func (v *Client) ZoomMuteSelf(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ZoomMuteSelf", params)
}

// ZoomSelectParticipantByName Select Zoom participant by name
func (v *Client) ZoomSelectParticipantByName(input interface{}, name string) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Value"] = name
	return v.SendFunction("ZoomSelectParticipantByName", params)
}

// ZoomUnMuteSelf Unmute self in Zoom meeting
func (v *Client) ZoomUnMuteSelf(input interface{}) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	return v.SendFunction("ZoomUnMuteSelf", params)
}
