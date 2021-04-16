package main

import (
	"log"

	vmixtcp "github.com/FlowingSPDG/vmix-go-TCP"
)

func main() {
	v, err := vmixtcp.New("localhost")
	if err != nil {
		panic(err)
	}
	defer v.Close()

	resp1, err := v.TALLY()
	if err != nil {
		panic(err)
	}
	log.Printf("TALLY RESPONSE1 : %s\n", resp1)

	resp, err := v.XML()
	if err != nil {
		panic(err)
	}
	log.Printf("XML RESPONSE : %s\n", resp)

	// If you want to parse XML, Comment-out following code to parse it
	// https://github.com/FlowingSPDG/vmix-go/blob/master/models.go#L12-L45
	//
	/*
		import {
			"encoding/xml"
			vmixgo "github.com/FlowingSPDG/vmix-go"
		}
		v := vmixgo.Vmix{}
		if err := xml.Unmarshal([]byte(resp), &v); err != nil {
			return err
		}
	*/

	resp1, err = v.FUNCTION("PreviewInput Input=1")
	if err != nil {
		panic(err)
	}
	log.Printf("FUNCTION RESPONSE : %s\n", resp1)

	if err := v.QUIT(); err != nil {
		panic(err)
	}
}
