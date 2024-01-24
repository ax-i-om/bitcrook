/*
Copyright © 2021 ax-i-om <addressaxiom@pm.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package handlers contains internal functions for handling requests
package http

import (
	"errors"
	"io"
	"net/http"
)

// GetReq performs a simple get request with a URL as its only parameter, and returns the response body as a string and an error.
func GetReq(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), resp.Body.Close()
}

// CustomGet performs a simple get request with a URL and specified headers/queries
func CustomGet(url string, headers []string) ([]byte, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, http.NoBody)

	if err != nil {
		return nil, err
	}

	if len(headers)%2 == 1 {
		return nil, errors.New("Uneven headers")
	}
	count := 0
	for count < len(headers) {
		req.Header.Add(headers[count], headers[count+1])
		count += 2
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, res.Body.Close()
}
