package run

import (
	"fmt"
	"sync"

	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/audioo/goseek/helpers/load"
)

// SendSeeker ...
func SendSeeker(userres string, wg *sync.WaitGroup, write bool) {

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		fmt.Println(http.GetSCnoredir(v.Title, v.Domain, wg, userres, write))
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		fmt.Println(http.GetSCredir(v.Title, v.Domain, wg, userres, write))
	}
}

// SendSeekerRL ... RETURNS LIST FOR ONLINE DEMO
func SendSeekerRL(userres string, wg *sync.WaitGroup, write bool) []string {
	var arr []string

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
