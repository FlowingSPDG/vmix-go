package main

import (
	"flag"
	"fmt"

	"github.com/FlowingSPDG/vmix-go/scraper"
)

func main() {
	helpVer := 0
	flag.IntVar(&helpVer, "helpver", 28, "vMix Help Version")

	fmt.Println("Start scraping vMix Help Version", helpVer)

	sc, err := scraper.GetShortcuts(helpVer)
	if err != nil {
		panic(err)
	}

	for _, s := range sc {
		fmt.Printf("%#v: ", s)
		for _, p := range s.Parameters {
			fmt.Println(p)
		}
	}

	fmt.Println("Scraping completed")
}
