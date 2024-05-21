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

	// Initialize vMix
	v := vmixtcp.New("localhost")
	// register callback
	v.OnVersion(func(r *vmixtcp.VersionResponse) {
		log.Println("Version:", r.Version)
	})
	v.OnActs(func(r *vmixtcp.ActsResponse) {
		log.Println("Response:", r.Response)
	})
	v.OnXML(func(r *vmixtcp.XMLResponse) {
		log.Printf("XML: %#v\n", r.XML)
	})

	retry := func() error {
		// Connect TCP API
		if err := v.Connect(); err != nil {
			return err
		}

		// subscribe
		if err := v.Subscribe(vmixtcp.EventActs, ""); err != nil {
			return err
		}

		// Send commands
		if err := v.XML(); err != nil {
			return err
		}

		if err := v.Acts("InputPreview"); err != nil {
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
