package main

import (
	"flag"
	"fmt"
	"os"

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

	if write {
		cli.Banner()
		fmt.Println(cli.Dispopg("GREEN", "ACCOUNT EXISTS"))
		fmt.Println(cli.Dispop("RED", "ACCOUNT DOES NOT EXIST"))
		fmt.Println("")
		run.SendSeeker(userres, true)
		fmt.Println("")
		cli.Dispban("Go-Seek Process Complete :: Results written to " + userres + ".txt")
	} else {
		cli.Banner()
		fmt.Println(cli.Dispopg("GREEN", "ACCOUNT EXISTS"))
		fmt.Println(cli.Dispop("RED", "ACCOUNT DOES NOT EXIST"))
		fmt.Println("")
		run.SendSeeker(userres, false)
		fmt.Println("")
		cli.Dispban("Go-Seek Process Complete :: Results have not been written")
	}
}
