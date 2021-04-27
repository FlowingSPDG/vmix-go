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
		if err := v.SUBSCRIBE(vmixtcp.EVENT_ACTS, ""); err != nil {
			panic(err)
		}

		v.Register(vmixtcp.EVENT_ACTS, func(r *vmixtcp.Response) {
			log.Println("ACT:", r)
			if err := v.XML(); err != nil {
				panic(err)
			}
		})

		v.Register(vmixtcp.EVENT_XML, func(r *vmixtcp.Response) {
			log.Println("XML:", r)
		})
		// timeout
		time.Sleep(time.Second)

		if err := v.XML(); err != nil {
			panic(err)
		}

		if err := v.XMLPATH("vmix/preview"); err != nil {
			panic(err)
		}

		if err := v.XMLPATH("vmix/active"); err != nil {
			panic(err)
		}

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
