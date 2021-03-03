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
func SendSeeker(userres string, wg *sync.WaitGroup, write bool) {

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		if http.GetSCnoredir(v.Title, v.Domain, wg, userres, write).Valid {
			fmt.Println(cli.Dispopg(v.Title, v.Domain))
		} else {
			fmt.Println(cli.Dispop(v.Title, v.Domain))
		}
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		if http.GetSCredir(v.Title, v.Domain, wg, userres, write).Valid {
			fmt.Println(cli.Dispopg(v.Title, v.Domain))
		} else {
			fmt.Println(cli.Dispop(v.Title, v.Domain))
		}
	}
}

// SendSeekerRL ... RETURNS LIST FOR ONLINE DEMO
func SendSeekerRL(userres string, wg *sync.WaitGroup, write bool) []ent.WebsiteRes {
	var arr []ent.WebsiteRes

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		arr = append(arr, http.GetSCnoredir(v.Title, v.Domain, wg, userres, write))
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		arr = append(arr, http.GetSCredir(v.Title, v.Domain, wg, userres, write))
	}

	return arr
}
