package main

import (
	"fmt"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/audioo/goseek/helpers/load"
)

func main() {

	var wg sync.WaitGroup
	var userres string

	cli.Banner()
	cli.Dispopg("GREEN", "ACCOUNT EXISTS")
	cli.Dispop("RED", "ACCOUNT DOES NOT EXIST")
	fmt.Println("")
	fmt.Println("   Enter Username")
	fmt.Print("   >> ")
	fmt.Scan(&userres)
	fmt.Println("")

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		http.GetSCnoredir(v.Title, v.Domain, &wg)
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		http.GetSCredir(v.Title, v.Domain, &wg)
	}

	wg.Wait()
	fmt.Println("")
	cli.Dispban("Go-Seek Process Complete")
	fmt.Println("")
	fmt.Println("")
}
