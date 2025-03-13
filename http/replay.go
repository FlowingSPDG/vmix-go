package vmixhttp

// ReplayACamera1 Select Camera 1 for Replay Channel A
func (v *Client) ReplayACamera1() error {
	return v.SendFunction("ReplayACamera1", nil)
}

// ReplayACamera2 Select Camera 2 for Replay Channel A
func (v *Client) ReplayACamera2() error {
	return v.SendFunction("ReplayACamera2", nil)
}

// ReplayACamera3 Select Camera 3 for Replay Channel A
func (v *Client) ReplayACamera3() error {
	return v.SendFunction("ReplayACamera3", nil)
}

// ReplayACamera4 Select Camera 4 for Replay Channel A
func (v *Client) ReplayACamera4() error {
	return v.SendFunction("ReplayACamera4", nil)
}

// ReplayACamera5 Select Camera 5 for Replay Channel A
func (v *Client) ReplayACamera5() error {
	return v.SendFunction("ReplayACamera5", nil)
}

// ReplayACamera6 Select Camera 6 for Replay Channel A
func (v *Client) ReplayACamera6() error {
	return v.SendFunction("ReplayACamera6", nil)
}

// ReplayACamera7 Select Camera 7 for Replay Channel A
func (v *Client) ReplayACamera7() error {
	return v.SendFunction("ReplayACamera7", nil)
}

// ReplayACamera8 Select Camera 8 for Replay Channel A
func (v *Client) ReplayACamera8() error {
	return v.SendFunction("ReplayACamera8", nil)
}

// ReplayBCamera1 Select Camera 1 for Replay Channel B
func (v *Client) ReplayBCamera1() error {
	return v.SendFunction("ReplayBCamera1", nil)
}

// ReplayBCamera2 Select Camera 2 for Replay Channel B
func (v *Client) ReplayBCamera2() error {
	return v.SendFunction("ReplayBCamera2", nil)
}

// ReplayBCamera3 Select Camera 3 for Replay Channel B
func (v *Client) ReplayBCamera3() error {
	return v.SendFunction("ReplayBCamera3", nil)
}

// ReplayBCamera4 Select Camera 4 for Replay Channel B
func (v *Client) ReplayBCamera4() error {
	return v.SendFunction("ReplayBCamera4", nil)
}

// ReplayBCamera5 Select Camera 5 for Replay Channel B
func (v *Client) ReplayBCamera5() error {
	return v.SendFunction("ReplayBCamera5", nil)
}

// ReplayBCamera6 Select Camera 6 for Replay Channel B
func (v *Client) ReplayBCamera6() error {
	return v.SendFunction("ReplayBCamera6", nil)
}

// ReplayBCamera7 Select Camera 7 for Replay Channel B
func (v *Client) ReplayBCamera7() error {
	return v.SendFunction("ReplayBCamera7", nil)
}

// ReplayBCamera8 Select Camera 8 for Replay Channel B
func (v *Client) ReplayBCamera8() error {
	return v.SendFunction("ReplayBCamera8", nil)
}

// ReplayCamera1 Select Camera 1 for Replay
func (v *Client) ReplayCamera1() error {
	return v.SendFunction("ReplayCamera1", nil)
}

// ReplayCamera2 Select Camera 2 for Replay
func (v *Client) ReplayCamera2() error {
	return v.SendFunction("ReplayCamera2", nil)
}

// ReplayCamera3 Select Camera 3 for Replay
func (v *Client) ReplayCamera3() error {
	return v.SendFunction("ReplayCamera3", nil)
}

// ReplayCamera4 Select Camera 4 for Replay
func (v *Client) ReplayCamera4() error {
	return v.SendFunction("ReplayCamera4", nil)
}

// ReplayCamera5 Select Camera 5 for Replay
func (v *Client) ReplayCamera5() error {
	return v.SendFunction("ReplayCamera5", nil)
}

// ReplayCamera6 Select Camera 6 for Replay
func (v *Client) ReplayCamera6() error {
	return v.SendFunction("ReplayCamera6", nil)
}

// ReplayCamera7 Select Camera 7 for Replay
func (v *Client) ReplayCamera7() error {
	return v.SendFunction("ReplayCamera7", nil)
}

// ReplayCamera8 Select Camera 8 for Replay
func (v *Client) ReplayCamera8() error {
	return v.SendFunction("ReplayCamera8", nil)
}

// ReplaySelectChannelA Select channel A
func (v *Client) ReplaySelectChannelA() error {
	return v.SendFunction("ReplaySelectChannelA", nil)
}

// ReplaySelectChannelAB Select channel AB
func (v *Client) ReplaySelectChannelAB() error {
	return v.SendFunction("ReplaySelectChannelAB", nil)
}

// ReplaySelectChannelB Select channel B
func (v *Client) ReplaySelectChannelB() error {
	return v.SendFunction("ReplaySelectChannelB", nil)
}

// ReplaySetAudioSource Set audio source
func (v *Client) ReplaySetAudioSource(source string) error {
	params := make(map[string]string)
	params["Value"] = source
	return v.SendFunction("ReplaySetAudioSource", params)
}

// ReplaySetChannelAToBTimecode Set A Timecode to B Timecode
func (v *Client) ReplaySetChannelAToBTimecode() error {
	return v.SendFunction("ReplaySetChannelAToBTimecode", nil)
}

// ReplaySetChannelAToBTimecodeAndCamera Set A Timecode and Camera to B
func (v *Client) ReplaySetChannelAToBTimecodeAndCamera() error {
	return v.SendFunction("ReplaySetChannelAToBTimecodeAndCamera", nil)
}

// ReplaySetChannelBToATimecode Set B Timecode to A Timecode
func (v *Client) ReplaySetChannelBToATimecode() error {
	return v.SendFunction("ReplaySetChannelBToATimecode", nil)
}

// ReplaySetChannelBToATimecodeAndCamera Set B Timecode and Camera to A
func (v *Client) ReplaySetChannelBToATimecodeAndCamera() error {
	return v.SendFunction("ReplaySetChannelBToATimecodeAndCamera", nil)
}
