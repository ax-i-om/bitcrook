package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/run"
)

func main() {

	var wg sync.WaitGroup
	var userres string
	var writeIt bool
	flag.BoolVar(&writeIt, "w", false, "Write results to a text file.")
	flag.StringVar(&userres, "u", "abc123anonymous!!bbzzn234whatttt", "Username input.")

	flag.Parse()
	if writeIt {

		if userres == "abc123anonymous!!bbzzn234whatttt" {
			fmt.Println("")
			fmt.Println("Command format:")
			fmt.Println("go run main.go -u=user -w=write")
			fmt.Println("")
			fmt.Println("Example:")
			fmt.Println("go run main.go -u=audioo -w=false")
			fmt.Println("")
		} else {
			cli.Banner()
			cli.Dispopg("GREEN", "ACCOUNT EXISTS")
			cli.Dispop("RED", "ACCOUNT DOES NOT EXIST")
			fmt.Println("")
			run.SendSeeker(userres, &wg, true)
			wg.Wait()
			fmt.Println("")
			cli.Dispban("Go-Seek Process Complete")
		}
	} else {
		if userres == "abc123anonymous!!bbzzn234whatttt" {
			fmt.Println("")
			fmt.Println("Command format:")
			fmt.Println("go run main.go -u=user -w=write")
			fmt.Println("")
			fmt.Println("Example:")
			fmt.Println("go run main.go -u=audioo -w=false")
			fmt.Println("")
		} else {
			cli.Banner()
			cli.Dispopg("GREEN", "ACCOUNT EXISTS")
			cli.Dispop("RED", "ACCOUNT DOES NOT EXIST")
			fmt.Println("")
			run.SendSeeker(userres, &wg, false)
			wg.Wait()
			fmt.Println("")
			cli.Dispban("Go-Seek Process Complete")
		}
	}
}
