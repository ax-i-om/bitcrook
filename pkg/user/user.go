package user

import (
	"sync"

	"github.com/maraudery/goseek/internal/http"
	load "github.com/maraudery/goseek/internal/siteloader"
	"github.com/maraudery/goseek/pkg/ent"
)

// SendSeeker checks all of the websites that have been
// loaded to see if the passed username correlates with any,
// and writes them to a .txt file if write == true.
// This will be updated in the future to return a []ent.Website slice.
func SendSeeker(userres string, write bool) {
	var wg sync.WaitGroup

	var arrNo []ent.Website = load.NoRedirSites(userres)
	for _, v := range arrNo {
		go http.CheckUser(v, userres, write, false, &wg)
	}
	var arrYes []ent.Website = load.RedirSites(userres)
	for _, v := range arrYes {
		go http.CheckUser(v, userres, write, true, &wg)
	}
	wg.Wait()
}
