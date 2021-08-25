package http

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/maraudery/ra/internal/config"
)

func TestGetReq(t *testing.T) {
	resp, err := http.Get("http://azenv.net/")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error("Non 200 Status Code")
	}
}

func TestAuthGet(t *testing.T) {
	elConf, err := config.LoadConfig("../../keyconfig.json")
	if err != nil {
		t.Error(err)
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, "https://haveibeenpwned.com/api/v3/breachedaccount/email@example.com", nil)

	if err != nil {
		t.Error(err)
	}
	req.Header.Add("hibp-api-key", elConf.HibpKey)

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}
