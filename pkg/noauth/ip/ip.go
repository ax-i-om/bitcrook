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

// Package ip contains the types and functions used for querying for
// information regarding an IPv4/IPv6 address
package ip

import (
	"encoding/json"

	"github.com/ax-i-om/bitcrook/internal/http"
)

// IPAddress is the response type of the IPLookup() function, containing a plethora of information about the IPV4 address in question.
type IPAddress struct {
	Status     string  `json:"status"`
	Continent  string  `json:"continent"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionname"`
	City       string  `json:"city"`
	District   string  `json:"district"`
	Zip        string  `json:"zip"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Timezone   string  `json:"timezone"`
	Currency   string  `json:"currency"`
	Isp        string  `json:"isp"`
	Org        string  `json:"org"`
	As         string  `json:"as"`
	Asname     string  `json:"asname"`
	Reverse    string  `json:"reverse"`
	Mobile     bool    `json:"mobile"`
	Proxy      bool    `json:"proxy"`
	Hosting    bool    `json:"hosting"`
	Query      string  `json:"query"`
}

// IPLookup takes an IPV4 Address (string) as a parameter and returns a plethora of information represented by type *IPAddress.
func IPLookup(ip string) (*IPAddress, error) {
	resp, err := http.GetReq("http://ip-api.com/json/" + ip + "?fields=31162361")
	if err != nil {
		return nil, err
	}
	s := new(IPAddress)
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
