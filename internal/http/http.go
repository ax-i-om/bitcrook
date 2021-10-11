package http

import (
	"io/ioutil"
	"net/http"
)

// GetReq performs a simple get request with a URL as its only parameter, and returns the response body as a string and an error.
func GetReq(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return (string(body)), nil
}

// AuthGet performs a simple get request with a URL, Auth Key, and an Auth Value as its parameters and returns the response body and an error.
func AuthGet(url, authkey, authval string) ([]byte, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add(authkey, authval)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
