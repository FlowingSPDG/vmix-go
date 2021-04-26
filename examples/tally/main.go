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
	defer v.Close()

	// Subscribe tally event
	_, err = v.SUBSCRIBE(vmixtcp.EVENT_TALLY)
	if err != nil {
		panic(err)
	}

	// register callback
	v.Register(vmixtcp.EVENT_TALLY, func(res *vmixtcp.Response) {
		log.Println("TALLY STATUS :", res)
	})

	lock := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for err := v.Run(ctx); err != nil; {
			// reconnect
			time.Sleep(time.Second)
			v.Run(ctx)
		}
	}()
	<-lock
}
