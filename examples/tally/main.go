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

	v := vmixtcp.New("localhost")
	// register callback
	v.OnTally(func(r *vmixtcp.TallyResponse) {
		log.Println("TALLY:", r.Tally)
	})

	retry := func() error {
		// reconnect
		if err := v.Connect(); err != nil {
			return err
		}

		// re-subscribe
		if err := v.Subscribe(vmixtcp.EventTally, ""); err != nil {
			return err
		}

		// run
		return v.Run(ctx)
	}
	go func() {
		for {
			if err := retry(); err != nil {
				log.Println("RETRY")
				time.Sleep(time.Second)
			}
		}
	}()
	<-ctx.Done()
	cancel()
	log.Println("Shutting down")
}
