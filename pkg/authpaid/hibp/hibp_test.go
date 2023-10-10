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

package hibp

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ax-i-om/bitcrook/internal/http"
)

func TestGetBreaches(t *testing.T) {
	key1 := os.Getenv("BITCROOK_HIBP")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/breachedaccount/email@example.com", "hibp-api-key", key1)
	if err != nil {
		t.Error(err)
	}
	response := new(BreachList)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBreachedSites(t *testing.T) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breaches")
	if err != nil {
		t.Error(err)
	}
	response := new(BreachedSites)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBreachedSite(t *testing.T) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breach/wattpad")
	if err != nil {
		t.Error(err)
	}
	response := new(BreachedSite)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllPastes(t *testing.T) {
	key1 := os.Getenv("BITCROOK_HIBP")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/pasteaccount/email@example.com", "hibp-api-key", key1)
	if err != nil {
		t.Error()
	}
	response := new(AllPastes)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error()
	}
}
