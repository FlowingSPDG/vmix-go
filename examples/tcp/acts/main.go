package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	vmixtcp "github.com/FlowingSPDG/vmix-go/tcp"
	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "XML", Description: "Send XML Command."},
		{Text: "VERSION", Description: "Get vMix version."},
		{Text: "TALLY", Description: "Get tally status."},
		{Text: "QUIT", Description: "Quit."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Initialize vMix
	v := vmixtcp.New("localhost")
	// register callback
	v.OnVersion(func(r *vmixtcp.VersionResponse) {
		log.Println("Version:", r.Version)

		// subscribe
		if err := v.Subscribe(vmixtcp.EventActs, ""); err != nil {
			panic(err)
		}

		// Send commands
		if err := v.XML(); err != nil {
			panic(err)
		}

		if err := v.Acts("InputPreview", nil); err != nil {
			panic(err)
		}
	})

	v.OnTally(func(r *vmixtcp.TallyResponse) {
		log.Println("TALLY:", r.Tally)
	})

	v.OnActs(func(r *vmixtcp.ActsResponse) {
		log.Println("ACTS:", r.Response)
	})

	v.OnXML(func(r *vmixtcp.XMLResponse) {
		log.Printf("XML: %#v\n", r.XML)
	})

	retry := func() error {
		// Connect TCP API
		if err := v.Connect(ctx, time.Second); err != nil {
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

	for {
		select {
		case <-ctx.Done():
			break
		default:
			t := prompt.Input("> ", completer)
			switch t {
			case "QUIT":
				return
			case "XML":
				if err := v.XML(); err != nil {
					panic(err)
				}
			case "VERSION":
				if err := v.Version(); err != nil {
					panic(err)
				}
			case "TALLY":
				if err := v.Tally(); err != nil {
					panic(err)
				}
			}
		}
	}

	<-ctx.Done()
	cancel()
	log.Println("Shutting down...")
}
