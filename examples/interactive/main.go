package main

import (
	"flag"
	"fmt"
	"github.com/FlowingSPDG/vmix-go"
	"github.com/c-bata/go-prompt"
)

var (
	Inputs []vmixgo.Input
	addr   *string
)

func TransitionCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Fade", Description: "Fade to next input in 500ms."},
		{Text: "Cut", Description: "Switches next input immidiately."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func InputCompleter(d prompt.Document) []prompt.Suggest {
	s := make([]prompt.Suggest, 0, len(Inputs))
	for i := 0; i < len(Inputs); i++ {
		s = append(s, prompt.Suggest{
			Text:        Inputs[i].Name,
			Description: Inputs[i].Title,
		})
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func init() {
	addr = flag.String("addr", "http://localhost:8088", "vMix API Address")
}

func main() {
	vmix, err := vmixgo.NewVmix(*addr)
	if err != nil {
		panic(err)
	}
	Inputs = vmix.Inputs.Input
	fmt.Println("Please select transition.")
	transition := prompt.Input(">>> ", TransitionCompleter)
	fmt.Println("Please select Input to switch.")
	in := prompt.Input(">>> ", InputCompleter)

	switch transition {
	case "Fade":
		vmix.Fade(in, 500)
	case "Cut":
		vmix.Cut(in)
	}
}
