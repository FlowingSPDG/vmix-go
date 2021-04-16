package main

import (
	"log"

	vmixtcp "github.com/FlowingSPDG/vmix-go-TCP"
)

var (
	lock chan struct{}
)

func main() {
	log.Println("STARTING...")
	lock := make(chan struct{})
	v, err := vmixtcp.New("localhost")
	if err != nil {
		panic(err)
	}
	log.Println("vMix connection success")
	defer v.Close()

	// Subscribe tally event
	_, err = v.SUBSCRIBE("TALLY")
	if err != nil {
		panic(err)
	}

	// register callback
	v.RegisterTallyCallback(func(res *vmixtcp.TallyResponse) {
		log.Println("TALLY STATUS :", res.Status)
		for i := 0; i < len(res.Tally); i++ {
			log.Printf("TALLY [%d] STATUS : %s\n", i+1, res.Tally[i].String())
		}
	})
	<-lock
}
