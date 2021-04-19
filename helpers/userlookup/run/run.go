package run

import (
	"fmt"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/audioo/goseek/helpers/userlookup/load"
)

// SendSeeker ...
func SendSeeker(userres string, write bool) {
	var wg sync.WaitGroup

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		go check(v, userres, write, false, &wg)
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		go check(v, userres, write, true, &wg)
	}
	wg.Wait()
}

func check(site ent.Website, userres string, write bool, redirect bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	x := http.GetSCredir(site.Title, site.Domain, userres, write, redirect).Valid
	if x {
		fmt.Println(cli.Dispopg(site.Title, site.Domain+" | "+site.Delete))
	} else {
		fmt.Println(cli.Dispop(site.Title, site.Domain+" | "+site.Delete))
	}
}
