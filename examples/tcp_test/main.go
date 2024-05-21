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
		v.OnVersion(func(r *vmixtcp.VersionResponse) {
			log.Println("Version:", r.Version)
		})
		v.OnActs(func(r *vmixtcp.ActsResponse) {
			log.Println("Response:", r.Response)
		})
		v.OnXML(func(r *vmixtcp.XMLResponse) {
			log.Printf("XML: %#v\n", r.XML)
		})

		// re-subscribe
		if err := v.Subscribe(vmixtcp.EventActs, ""); err != nil {
			panic(err)
		}

		if err := v.XML(); err != nil {
			panic(err)
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
