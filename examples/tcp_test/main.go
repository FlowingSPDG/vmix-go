package main

import (
	"context"
	"log"
	"time"

	vmixtcp "github.com/FlowingSPDG/vmix-go/tcp"
)

func main() {
	var v *vmixtcp.Vmix
	retry := func() error {
		// reconnect
		var err error
		v, err = vmixtcp.New("localhost")
		if err != nil {
			return err
		}

		// re-subscribe
		s, err := v.SUBSCRIBE(vmixtcp.EVENT_ACTS)
		log.Println("SUBSCRIBE:", s)
		v.Register(vmixtcp.EVENT_ACTS, func(r *vmixtcp.Response) {
			log.Println("ACT:", r)
		})
		// timeout
		time.Sleep(time.Second)

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

		// run
		return v.Run(context.TODO())
	}
	go func() {
		for err := retry(); err != nil; {
			log.Println("RETRY")
			time.Sleep(time.Second)
			err = retry()
		}
	}()
	lock := make(chan struct{})
	<-lock
}
