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

package http

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetReq(t *testing.T) {
	resp, err := http.Get("http://azenv.net/")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error("Non 200 Status Code")
	}
}

func TestAuthGet(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Error(err)
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, "https://haveibeenpwned.com/api/v3/breachedaccount/email@example.com", http.NoBody)

	if err != nil {
		t.Error(err)
	}
	req.Header.Add("hibp-api-key", os.Getenv("BITCROOK_HIBP"))

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}
