package vmixhttp

import "strconv"

// ReplayChangeDirection Change replay direction
func (v *Client) ReplayChangeDirection(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayChangeDirection", params)
}

// ReplayChangeSpeed Change replay speed
func (v *Client) ReplayChangeSpeed(speed float64, channel string) error {
	params := make(map[string]string)
	params["Value"] = strconv.FormatFloat(speed, 'f', 2, 64)
	params["Channel"] = channel
	return v.SendFunction("ReplayChangeSpeed", params)
}

// ReplayFastBackward Fast backward replay (1-30x)
func (v *Client) ReplayFastBackward(speed uint, channel string) error {
	if speed > 30 {
		speed = 30
	}
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(speed))
	params["Channel"] = channel
	return v.SendFunction("ReplayFastBackward", params)
}

// ReplayFastForward Fast forward replay (1-30x)
func (v *Client) ReplayFastForward(speed uint, channel string) error {
	if speed > 30 {
		speed = 30
	}
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(speed))
	params["Channel"] = channel
	return v.SendFunction("ReplayFastForward", params)
}

// ReplayJumpFrames Jump specified number of frames
func (v *Client) ReplayJumpFrames(frames int, channel string) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(frames)
	params["Channel"] = channel
	return v.SendFunction("ReplayJumpFrames", params)
}

// ReplayJumpFramesFastOff ReplayJumpFrames jumps 1 frame for each value instead of 1 second
func (v *Client) ReplayJumpFramesFastOff(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayJumpFramesFastOff", params)
}

// ReplayJumpFramesFastOn ReplayJumpFrames jumps 1 second for each value instead of 1 frame
func (v *Client) ReplayJumpFramesFastOn(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayJumpFramesFastOn", params)
}

// ReplayJumpToNow Jump to current time
func (v *Client) ReplayJumpToNow(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayJumpToNow", params)
}

// ReplayLive Switch to live mode
func (v *Client) ReplayLive() error {
	return v.SendFunction("ReplayLive", nil)
}

// ReplayLiveToggle Toggle live mode
func (v *Client) ReplayLiveToggle() error {
	return v.SendFunction("ReplayLiveToggle", nil)
}

// ReplayPause Pause replay
func (v *Client) ReplayPause(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPause", params)
}

// ReplayPlay Play replay
func (v *Client) ReplayPlay(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlay", params)
}

// ReplayPlayBackward Play backward
func (v *Client) ReplayPlayBackward(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayBackward", params)
}

// ReplayPlayForward Play forward
func (v *Client) ReplayPlayForward(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayForward", params)
}

// ReplayPlayPause Toggle play/pause
func (v *Client) ReplayPlayPause(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayPause", params)
}

// ReplayRecorded Switch to recorded mode
func (v *Client) ReplayRecorded() error {
	return v.SendFunction("ReplayRecorded", nil)
}

// ReplaySetSpeed Set replay speed (0-1)
func (v *Client) ReplaySetSpeed(speed float64, channel string) error {
	if speed < 0 {
		speed = 0
	} else if speed > 1 {
		speed = 1
	}
	params := make(map[string]string)
	params["Value"] = strconv.FormatFloat(speed, 'f', 2, 64)
	params["Channel"] = channel
	return v.SendFunction("ReplaySetSpeed", params)
}
