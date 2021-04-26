package main

import (
	"log"
	"time"

	vmixtcp "github.com/FlowingSPDG/vmix-go/tcp"
)

func main() {
	v, err := vmixtcp.New("localhost")
	if err != nil {
		panic(err)
	}
	x, err := v.XML()
	if err != nil {
		panic(err)
	}
	log.Println("XML:", x)

	s, err := v.SUBSCRIBE(vmixtcp.EVENT_ACTS)
	log.Println("SUBSCRIBE:", s)
	v.Register(vmixtcp.EVENT_ACTS, func(r *vmixtcp.Response) {
		log.Println("ACT:", r)
	})
	time.Sleep(time.Second * 5)
	if err := v.Close(); err != nil {
		panic(err)
	}
}
