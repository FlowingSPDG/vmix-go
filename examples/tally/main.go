package main

import (
	"log"

	vmixtcp "github.com/FlowingSPDG/vmix-go/tcp"
)

var (
	lock chan struct{}
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
	<-lock
}
