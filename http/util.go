package vmixhttp

import (
	"fmt"
	"strconv"

	vmixgo "github.com/FlowingSPDG/vmix-go"
)

// resolveInput resolves vmix keys, number, scene name to string.
func resolveInput(input any) (string, error) {
	switch input := input.(type) {
	case int:
		return strconv.Itoa(input), nil
	case string:
		return input, nil
	case vmixgo.Input:
		return input.Key, nil
	default:
		return "", fmt.Errorf("Interface type not correct(%v)", input)
	}
}
