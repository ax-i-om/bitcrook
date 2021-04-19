package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/audioo/goseek/helpers/ent"
)

// GetSCredir ...
func GetSCredir(title string, url string, userres string, write bool, redirect bool) ent.Website {
	file, _ := os.OpenFile(userres+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if !write {
		file.Close()
		os.Remove(userres + ".txt")
	} else {
		defer file.Close()
	}

	method := "GET"
	var client *http.Client
	if redirect {
		client = &http.Client{
			Timeout: 5 * time.Second,
		}
	} else {
		client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 5 * time.Second,
		}
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return ent.Website{Title: title, Domain: url, Valid: false}
	}
	res, err := client.Do(req)
	if err != nil {
		return ent.Website{Title: title, Domain: url, Valid: false}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ent.Website{Title: title, Domain: url, Valid: false}
	}
	strconv.AppendBool(body, true)
	if res.StatusCode == 200 {
		if write {
			fmt.Fprintf(file, "%s\n", title+" => "+url)
		}
		return ent.Website{Title: title, Domain: url, Valid: true}
	}
	return ent.Website{Title: title, Domain: url, Valid: false}

}

// Connected ...
func Connected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	return err == nil
}

// GetReq ...
func GetReq(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	return (string(body))
}
