package run

import (
	"sync"

	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/audioo/goseek/helpers/load"
)

// SendSeeker ...
func SendSeeker(userres string, wg *sync.WaitGroup) {

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		http.GetSCnoredir(v.Title, v.Domain, wg)
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		http.GetSCredir(v.Title, v.Domain, wg)
	}
}

// SendSeekerWrite ...
func SendSeekerWrite(userres string, wg *sync.WaitGroup) {

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		http.GetSCnoredirWrite(v.Title, v.Domain, wg, userres)
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		http.GetSCredirWrite(v.Title, v.Domain, wg, userres)
	}
}
