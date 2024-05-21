package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	vmixtcp "github.com/FlowingSPDG/vmix-go/tcp"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	var v *vmixtcp.Vmix
	retry := func() error {
		// reconnect
		var err error
		v, err = vmixtcp.New("localhost")
		if err != nil {
			return err
		}
		// register callback
		v.OnTally(func(r *vmixtcp.TallyResponse) {
			log.Println("TALLY:", r.Tally)
		})

		// re-subscribe
		if err := v.Subscribe(vmixtcp.EventTally, ""); err != nil {
			return err
		}

		// run
		return v.Run(ctx)
	}
	go func() {
		for err := retry(); err != nil; {
			log.Println("RETRY")
			time.Sleep(time.Second)
			err = retry()
		}
	}()
	<-ctx.Done()
	cancel()
	log.Println("Shutting down")
}
