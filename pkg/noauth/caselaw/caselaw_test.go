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

package caselaw

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestGetCases(t *testing.T) {
	url := "https://api.case.law/v1/cases/?page_size=10&search=texas&ordering=relevance"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, http.NoBody)

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	clres := new(CLResponse)
	err = json.Unmarshal([]byte(body), &clres)
	if err != nil {
		t.Error(err)
	}
}
