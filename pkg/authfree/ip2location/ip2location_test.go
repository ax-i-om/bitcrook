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

package ip2location

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ax-i-om/bitcrook/internal/http"
)

func TestDomainLookup(t *testing.T) {
	key1 := os.Getenv("BITCROOK_IPTL")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	resp, err := http.GetReq("https://api.ip2whois.com/v2?key=" + key1 + "&domain=bitcrook.tech&format=json")
	if err != nil {
		t.Error(err)
	}
	response := new(DomainData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
	}
	if response.Domain != "bitcrook.tech" {
		t.Error()
	}
}

func TestIpLookup(t *testing.T) {
	key1 := os.Getenv("BITCROOK_IPTL")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	resp, err := http.GetReq("https://api.ip2location.io/?key=" + key1 + "&ip=1.1.1.1&format=json")
	if err != nil {
		t.Error(err)
	}
	response := new(IpData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
	}
}
