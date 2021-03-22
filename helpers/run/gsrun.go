package run

import (
	"fmt"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/audioo/goseek/helpers/load"
)

// SendSeeker ...
func SendSeeker(userres string, write bool) {
	var wg sync.WaitGroup

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		go check(v.Title, v.Domain, userres, write, false, &wg)
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		go check(v.Title, v.Domain, userres, write, true, &wg)
	}
	wg.Wait()
}

func check(title string, domain string, userres string, write bool, redirect bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	x := http.GetSCredir(title, domain, userres, write, redirect).Valid
	if x {
		fmt.Println(cli.Dispopg(title, domain))
	} else {
		fmt.Println(cli.Dispop(title, domain))
	}
}
