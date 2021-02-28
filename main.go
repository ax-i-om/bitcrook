package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/run"
)

var (
	userres string
	write   bool
)

func init() {

	flag.BoolVar(&write, "w", false, "Write to File")
	flag.Parse()
	if !(len(flag.Args()) > 0) {
		fmt.Println("\nProper Command Format: ")
		fmt.Println("go run main.go [flags] username")
		fmt.Println("")
		flag.Usage()
		os.Exit(1)
	}
	userres = flag.Args()[0]

}

func main() {

	var wg sync.WaitGroup

	if write {
		cli.Banner()
		cli.Dispopg("GREEN", "ACCOUNT EXISTS")
		cli.Dispop("RED", "ACCOUNT DOES NOT EXIST")
		fmt.Println("")
		run.SendSeeker(userres, &wg, true)
		wg.Wait()
		fmt.Println("")
		cli.Dispban("Go-Seek Process Complete :: Results written to " + userres + ".txt")
	} else {
		cli.Banner()
		cli.Dispopg("GREEN", "ACCOUNT EXISTS")
		cli.Dispop("RED", "ACCOUNT DOES NOT EXIST")
		fmt.Println("")
		run.SendSeeker(userres, &wg, false)
		wg.Wait()
		fmt.Println("")
		cli.Dispban("Go-Seek Process Complete :: Results have not been written")
	}
}
