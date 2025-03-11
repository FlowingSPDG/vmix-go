package vmixhttp

import "strconv"

// SaveVideoDelay Save video clip from Video Delay
func (v *Client) SaveVideoDelay(input interface{}, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Duration"] = strconv.Itoa(int(durationMs))
	return v.SendFunction("SaveVideoDelay", params)
}

// VideoDelayStartRecording Start Video Delay Recording
func (v *Client) VideoDelayStartRecording(input interface{}, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Duration"] = strconv.Itoa(int(durationMs))
	return v.SendFunction("VideoDelayStartRecording", params)
}

// VideoDelayStartStopRecording Toggle Video Delay Recording
func (v *Client) VideoDelayStartStopRecording(input interface{}, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Duration"] = strconv.Itoa(int(durationMs))
	return v.SendFunction("VideoDelayStartStopRecording", params)
}

// VideoDelayStopRecording Stop Video Delay Recording
func (v *Client) VideoDelayStopRecording(input interface{}, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Duration"] = strconv.Itoa(int(durationMs))
	return v.SendFunction("VideoDelayStopRecording", params)
}

// WaitForCompletion Wait for a Video Input to reach the end of playback
func (v *Client) WaitForCompletion(input interface{}, durationMs uint) error {
	in, err := resolveInput(input)
	if err != nil {
		return err
	}
	params := make(map[string]string)
	params["Input"] = in
	params["Duration"] = strconv.Itoa(int(durationMs))
	return v.SendFunction("WaitForCompletion", params)
}
