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
		if err := v.SUBSCRIBE(vmixtcp.EVENT_TALLY, ""); err != nil {
			return err
		}
		v.Register(vmixtcp.EVENT_TALLY, func(r *vmixtcp.Response) {
			log.Println("TALLY:", r)
		})
		// timeout
		time.Sleep(time.Second)

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
