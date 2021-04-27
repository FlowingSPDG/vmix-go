package main

import (
	"flag"
	"log"

	vmixgo "github.com/FlowingSPDG/vmix-go/http"
)

var (
	host *string
	port *int
)

func init() {
	host = flag.String("host", "localhost", "vMix HTTP API Host")
	port = flag.Int("port", 8088, "vMix HTTP API Port")
	flag.Parse()
}

func main() {
	vmix, err := vmixgo.NewClient(*host, *port)
	if err != nil {
		panic(err)
	}
	log.Printf("vmix version: %v\n", vmix.Version)
	log.Printf("vmix edition: %v\n", vmix.Edition)
	for i := 0; i < len(vmix.Inputs.Input); i++ {
		log.Printf("Input %d : %v", vmix.Inputs.Input[i].Number, vmix.Inputs.Input[i])
	}
	for i := 0; i < len(vmix.Overlays.Overlay); i++ {
		log.Printf("Overlay %d : %v", vmix.Overlays.Overlay[i].Number, vmix.Overlays.Overlay[i])
	}
	for i := 0; i < len(vmix.Transitions.Transition); i++ {
		log.Printf("Transition %d : %v", vmix.Transitions.Transition[i].Number, vmix.Transitions.Transition[i])
	}
}
