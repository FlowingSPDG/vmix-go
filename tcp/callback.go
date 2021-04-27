package vmixtcp

import (
	"fmt"
)

// Register goroutine callback event.
func (v *Vmix) Register(command string, cb func(*Response)) error {
	if _, exist := v.cbhandler[command]; exist {
		return fmt.Errorf("Handler exist")
	}
	v.cbhandler[command] = append(v.cbhandler[command], cb)
	return nil
}
