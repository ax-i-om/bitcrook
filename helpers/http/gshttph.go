package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/audioo/goseek/helpers/cli"
)

// GetSCredir ... Automatically follows HTTP redirects.
func GetSCredir(title string, url string, wg *sync.WaitGroup, userres string, write bool) string {
	file, err := os.OpenFile(userres+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		// fmt.Println(err)
	}
	if !write {
		file.Close()
		os.Remove(userres + ".txt")
	} else {
		defer file.Close()
	}
	wg.Add(1)

	defer wg.Done()
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return cli.Dispop(title, url)
	}
	res, err := client.Do(req)
	if err != nil {
		return cli.Dispop(title, url)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return cli.Dispop(title, url)
	}
	strconv.AppendBool(body, true)
	if res.StatusCode == 200 {
		if write {
			fmt.Fprintf(file, "%s\n", title+" => "+url)
		}
		return cli.Dispopg(title, url)
	}
	return cli.Dispop(title, url)

}

// GetSCnoredir ... Does not automatically follow HTTP redirects.
func GetSCnoredir(title string, url string, wg *sync.WaitGroup, userres string, write bool) string {
	file, err := os.OpenFile(userres+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		// fmt.Println(err)
		// return
	}
	defer file.Close()
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
		return cli.Dispop(title, url)
	}
	res, err := client.Do(req)
	if err != nil {
		return cli.Dispop(title, url)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return cli.Dispop(title, url)
	}
	strconv.AppendBool(body, true)
	if res.StatusCode == 200 {
		if write {
			fmt.Fprintf(file, "%s\n", title+" => "+url)
		}
		return cli.Dispopg(title, url)
	}
	return cli.Dispop(title, url)
}
