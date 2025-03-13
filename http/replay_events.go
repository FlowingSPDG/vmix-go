package vmixhttp

import "strconv"

// ReplayCopyLastEvent Copy last event to specified event list
func (v *Client) ReplayCopyLastEvent(eventList uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(eventList))
	return v.SendFunction("ReplayCopyLastEvent", params)
}

// ReplayCopySelectedEvent Copy selected event to specified event list
func (v *Client) ReplayCopySelectedEvent(eventList uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(eventList))
	return v.SendFunction("ReplayCopySelectedEvent", params)
}

// ReplayDeleteLastEvent Delete last event
func (v *Client) ReplayDeleteLastEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayDeleteLastEvent", params)
}

// ReplayDeleteSelectedEvent Delete selected event
func (v *Client) ReplayDeleteSelectedEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayDeleteSelectedEvent", params)
}

// ReplayDuplicateLastEvent Duplicate last event
func (v *Client) ReplayDuplicateLastEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayDuplicateLastEvent", params)
}

// ReplayDuplicateSelectedEvent Duplicate selected event
func (v *Client) ReplayDuplicateSelectedEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayDuplicateSelectedEvent", params)
}

// ReplayExportLastEvent Export last event to specified folder
func (v *Client) ReplayExportLastEvent(folder string, channel string) error {
	params := make(map[string]string)
	params["Value"] = folder
	params["Channel"] = channel
	return v.SendFunction("ReplayExportLastEvent", params)
}

// ReplayMarkCancel Cancel marking
func (v *Client) ReplayMarkCancel() error {
	return v.SendFunction("ReplayMarkCancel", nil)
}

// ReplayMarkIn Mark in point
func (v *Client) ReplayMarkIn() error {
	return v.SendFunction("ReplayMarkIn", nil)
}

// ReplayMarkInLive Mark in point live
func (v *Client) ReplayMarkInLive() error {
	return v.SendFunction("ReplayMarkInLive", nil)
}

// ReplayMarkInOut Create new event with specified duration
func (v *Client) ReplayMarkInOut(seconds uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(seconds))
	return v.SendFunction("ReplayMarkInOut", params)
}

// ReplayMarkInOutLive Create new event with specified duration from live
func (v *Client) ReplayMarkInOutLive(seconds uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(seconds))
	return v.SendFunction("ReplayMarkInOutLive", params)
}

// ReplayMarkInOutLiveFuture Create new event with specified duration into the future
func (v *Client) ReplayMarkInOutLiveFuture(seconds uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(seconds))
	return v.SendFunction("ReplayMarkInOutLiveFuture", params)
}

// ReplayMarkInOutRecorded Create new event with specified duration from recorded
func (v *Client) ReplayMarkInOutRecorded(seconds uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(seconds))
	return v.SendFunction("ReplayMarkInOutRecorded", params)
}

// ReplayMarkInRecorded Mark in point recorded
func (v *Client) ReplayMarkInRecorded() error {
	return v.SendFunction("ReplayMarkInRecorded", nil)
}

// ReplayMarkInRecordedNow Mark in point recorded now
func (v *Client) ReplayMarkInRecordedNow() error {
	return v.SendFunction("ReplayMarkInRecordedNow", nil)
}

// ReplayMarkOut Mark out point
func (v *Client) ReplayMarkOut() error {
	return v.SendFunction("ReplayMarkOut", nil)
}

// ReplayMoveLastEvent Move last event to specified event list
func (v *Client) ReplayMoveLastEvent(eventList uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(eventList))
	return v.SendFunction("ReplayMoveLastEvent", params)
}

// ReplayMoveSelectedEvent Move selected event to specified event list
func (v *Client) ReplayMoveSelectedEvent(eventList uint) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(eventList))
	return v.SendFunction("ReplayMoveSelectedEvent", params)
}

// ReplayMoveSelectedEventDown Move selected event down
func (v *Client) ReplayMoveSelectedEventDown() error {
	return v.SendFunction("ReplayMoveSelectedEventDown", nil)
}

// ReplayMoveSelectedEventUp Move selected event up
func (v *Client) ReplayMoveSelectedEventUp() error {
	return v.SendFunction("ReplayMoveSelectedEventUp", nil)
}

// ReplayPlayAllEvents Play all Events in active list
func (v *Client) ReplayPlayAllEvents(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayAllEvents", params)
}

// ReplayPlayAllEventsToOutput Play all Events in active list to output
func (v *Client) ReplayPlayAllEventsToOutput(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayAllEventsToOutput", params)
}

// ReplayPlayEvent Play event by number
func (v *Client) ReplayPlayEvent(eventNumber uint, channel string) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(eventNumber))
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayEvent", params)
}

// ReplayPlayEventsByID Play events by ID list
func (v *Client) ReplayPlayEventsByID(ids string, channel string) error {
	params := make(map[string]string)
	params["Value"] = ids
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayEventsByID", params)
}

// ReplayPlayEventsByIDToOutput Play events by ID list to output
func (v *Client) ReplayPlayEventsByIDToOutput(ids string, channel string) error {
	params := make(map[string]string)
	params["Value"] = ids
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayEventsByIDToOutput", params)
}

// ReplayPlayEventToOutput Play event by number to output
func (v *Client) ReplayPlayEventToOutput(eventNumber uint, channel string) error {
	params := make(map[string]string)
	params["Value"] = strconv.Itoa(int(eventNumber))
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayEventToOutput", params)
}

// ReplayPlayLastEvent Play last event
func (v *Client) ReplayPlayLastEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayLastEvent", params)
}

// ReplayPlayLastEventToOutput Play last event to output
func (v *Client) ReplayPlayLastEventToOutput(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayLastEventToOutput", params)
}

// ReplayPlayNext Play next event
func (v *Client) ReplayPlayNext(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayNext", params)
}

// ReplayPlayPrevious Play previous event
func (v *Client) ReplayPlayPrevious(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlayPrevious", params)
}

// ReplayPlaySelectedEvent Play selected event
func (v *Client) ReplayPlaySelectedEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlaySelectedEvent", params)
}

// ReplayPlaySelectedEventToOutput Play selected event to output
func (v *Client) ReplayPlaySelectedEventToOutput(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplayPlaySelectedEventToOutput", params)
}

// ReplaySelectAllEvents Select all events in active channel
func (v *Client) ReplaySelectAllEvents() error {
	return v.SendFunction("ReplaySelectAllEvents", nil)
}

// ReplaySelectEvents1 Select events list 1
func (v *Client) ReplaySelectEvents1(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents1", params)
}

// ReplaySelectEvents2 Select events list 2
func (v *Client) ReplaySelectEvents2(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents2", params)
}

// ReplaySelectEvents3 Select events list 3
func (v *Client) ReplaySelectEvents3(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents3", params)
}

// ReplaySelectEvents4 Select events list 4
func (v *Client) ReplaySelectEvents4(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents4", params)
}

// ReplaySelectEvents5 Select events list 5
func (v *Client) ReplaySelectEvents5(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents5", params)
}

// ReplaySelectEvents6 Select events list 6
func (v *Client) ReplaySelectEvents6(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents6", params)
}

// ReplaySelectEvents7 Select events list 7
func (v *Client) ReplaySelectEvents7(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents7", params)
}

// ReplaySelectEvents8 Select events list 8
func (v *Client) ReplaySelectEvents8(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents8", params)
}

// ReplaySelectEvents9 Select events list 9
func (v *Client) ReplaySelectEvents9(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents9", params)
}

// ReplaySelectEvents10 Select events list 10
func (v *Client) ReplaySelectEvents10(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents10", params)
}

// ReplaySelectEvents11 Select events list 11
func (v *Client) ReplaySelectEvents11(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents11", params)
}

// ReplaySelectEvents12 Select events list 12
func (v *Client) ReplaySelectEvents12(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents12", params)
}

// ReplaySelectEvents13 Select events list 13
func (v *Client) ReplaySelectEvents13(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents13", params)
}

// ReplaySelectEvents14 Select events list 14
func (v *Client) ReplaySelectEvents14(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents14", params)
}

// ReplaySelectEvents15 Select events list 15
func (v *Client) ReplaySelectEvents15(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents15", params)
}

// ReplaySelectEvents16 Select events list 16
func (v *Client) ReplaySelectEvents16(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents16", params)
}

// ReplaySelectEvents17 Select events list 17
func (v *Client) ReplaySelectEvents17(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents17", params)
}

// ReplaySelectEvents18 Select events list 18
func (v *Client) ReplaySelectEvents18(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents18", params)
}

// ReplaySelectEvents19 Select events list 19
func (v *Client) ReplaySelectEvents19(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents19", params)
}

// ReplaySelectEvents20 Select events list 20
func (v *Client) ReplaySelectEvents20(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectEvents20", params)
}

// ReplaySelectFirstEvent Select first event
func (v *Client) ReplaySelectFirstEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectFirstEvent", params)
}

// ReplaySelectLastEvent Select last event
func (v *Client) ReplaySelectLastEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectLastEvent", params)
}

// ReplaySelectNextEvent Select next event
func (v *Client) ReplaySelectNextEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectNextEvent", params)
}

// ReplaySelectPreviousEvent Select previous event
func (v *Client) ReplaySelectPreviousEvent(channel string) error {
	params := make(map[string]string)
	params["Channel"] = channel
	return v.SendFunction("ReplaySelectPreviousEvent", params)
}
