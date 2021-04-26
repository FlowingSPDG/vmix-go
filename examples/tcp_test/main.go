package main

import (
	"context"
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

	xpathPreview, err := v.XMLPATH("vmix/preview")
	if err != nil {
		panic(err)
	}
	log.Println("XPATH vmix/preview:", xpathPreview)

	xpathProgram, err := v.XMLPATH("vmix/active")
	if err != nil {
		panic(err)
	}
	log.Println("XPATH vmix/preview:", xpathProgram)

	s, err := v.SUBSCRIBE(vmixtcp.EVENT_ACTS)
	log.Println("SUBSCRIBE:", s)
	v.Register(vmixtcp.EVENT_ACTS, func(r *vmixtcp.Response) {
		log.Println("ACT:", r)
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for err := v.Run(ctx); err != nil; {
			// reconnect
			time.Sleep(time.Second)
			v.Run(ctx)
		}
	}()
	time.Sleep(time.Second * 5)
}
