package vmixgo

import (
	"fmt"
	"strconv"

	"github.com/FlowingSPDG/vmix-go/common/models"
)

// resolveInput resolves vmix keys, number, scene name to string.
func resolveInput(input interface{}) (string, error) {
	switch input := input.(type) {
	case int:
		return strconv.Itoa(input), nil
	case string:
		return input, nil
	case models.Input:
		return input.Key, nil
	default:
		return "", fmt.Errorf("Interface type not correct(%v)", input)
	}
}
