package caselaw

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCases(t *testing.T) {
	url := "https://api.case.law/v1/cases/?page_size=10&search=texas&ordering=relevance"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Allow", "GET, HEAD, OPTIONS")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Vary", "Accept")

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	clres := new(CLResponse)
	err = json.Unmarshal([]byte(body), &clres)
	if err != nil {
		t.Error(err)
	}
}
