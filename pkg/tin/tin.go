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

// Package tin contains the types and functions used for querying for
// information regarding a Vehicle Identification Number (VIN)
package tin

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

/* ///////////////////////////////////////////////////////

   Nalog

*/ ///////////////////////////////////////////////////////

type postresp struct {
	T               string `json:"t"`
	CaptchaRequired bool   `json:"captchaRequired"`
}

type TinData struct {
	Rows []struct {
		AssignmentDate  string `json:"r,omitempty"`
		T               string `json:"t"`
		TerminationDate string `json:"e,omitempty"`
		Page            string `json:"pg,omitempty"`
		Total           string `json:"tot,omitempty"`
		Count           string `json:"cnt,omitempty"`
		INN             string `json:"i,omitempty"`
		K               string `json:"k"`
		Name            string `json:"n,omitempty"`
		OGRNIP          string `json:"o,omitempty"`
	} `json:"rows"`
}

// TINLookup ...
func TINLookup(tin string) (*TinData, error) {
	urlOne := "https://egrul.nalog.ru/"
	methodOne := "POST"

	payloadOne := strings.NewReader("query=" + tin)

	clientOne := &http.Client{}
	reqOne, err := http.NewRequest(methodOne, urlOne, payloadOne)

	if err != nil {
		return nil, err
	}
	reqOne.Header.Add("Content-Type", " application/x-www-form-urlencoded; charset=UTF-8")

	res, err := clientOne.Do(reqOne)
	if err != nil {
		return nil, err
	}

	bodyOne, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.ErrAbortHandler
	}
	s := new(postresp)
	err = json.Unmarshal([]byte(bodyOne), &s)
	if err != nil {
		return nil, err
	}

	urlTwo := "https://egrul.nalog.ru/search-result/" + s.T
	methodTwo := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(methodTwo, urlTwo, nil)

	if err != nil {
		return nil, err
	}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyTwo, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	sTwo := new(TinData)
	err = json.Unmarshal([]byte(bodyTwo), &sTwo)
	if err != nil {
		return nil, err
	}

	return sTwo, res.Body.Close()
}
