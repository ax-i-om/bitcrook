package http

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
)

// GetSCredir ... Automatically follows HTTP redirects.
func GetSCredir(title string, url string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		cli.Dispop(title, url)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		cli.Dispop(title, url)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		cli.Dispop(title, url)
		return
	}
	strconv.AppendBool(body, true)
	if res.StatusCode == 200 {
		cli.Dispopg(title, url)
	} else {
		cli.Dispop(title, url)
	}
}

// GetSCnoredir ... Does not automatically follow HTTP redirects.
func GetSCnoredir(title string, url string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		cli.Dispop(title, url)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		cli.Dispop(title, url)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		cli.Dispop(title, url)
		return
	}
	strconv.AppendBool(body, true)
	if res.StatusCode == 200 {
		cli.Dispopg(title, url)
	} else {
		cli.Dispop(title, url)
	}
}
