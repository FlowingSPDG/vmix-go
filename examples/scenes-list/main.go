package main

import (
	"flag"
	"github.com/FlowingSPDG/vmix-go"
	"log"
	"time"
)

var (
	addr *string
)

func init() {
	addr = flag.String("addr", "http://localhost:8088", "vMix API Address")
	flag.Parse()
}

func main() {
	vmix, err := vmixgo.NewVmix(*addr)
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

	err = vmix.Fade(vmix.Inputs.Input[0], 500)
	if err != nil {
		log.Printf("err : %v\n", err)
	}
	time.Sleep(time.Second)

	err = vmix.Fade(0, 500)
	if err != nil {
		log.Printf("err : %v\n", err)
	}
	time.Sleep(time.Second)

	err = vmix.Fade("474e6346-1ea3-468c-8f93-70add19c354f", 500)
	if err != nil {
		log.Printf("err : %v\n", err)
	}
	time.Sleep(time.Second)

	err = vmix.Fade(nil, 500)
	if err != nil {
		log.Printf("err : %v\n", err)
	}
	time.Sleep(time.Second)
}
