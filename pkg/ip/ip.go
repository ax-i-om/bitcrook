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

/* ///////////////////////////////////////////////////////

   IP-API

*/ ///////////////////////////////////////////////////////

// IPAPIAddress is the response type of the IPAPILookup() function, containing a plethora of information about the IPv4/IPv6 address queried.
type IPAPIAddress struct {
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

// IPAPILookup takes an IP Address (string) as a parameter and returns a plethora of information represented by type *IPAddress.
func IPAPILookup(ip string) (*IPAPIAddress, error) {
	resp, err := http.GetReq("http://ip-api.com/json/" + ip + "?fields=31162361")
	if err != nil {
		return nil, err
	}
	s := new(IPAPIAddress)
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type ipResponse struct {
	Version               string `json:"Version"`
	Transmissionreference string `json:"TransmissionReference"`
	Transmissionresults   string `json:"TransmissionResults"`
	Records               []struct {
		City                string `json:"City"`
		Connectionspeed     string `json:"ConnectionSpeed"`
		Connectiontype      string `json:"ConnectionType"`
		Continent           string `json:"Continent"`
		Countryabbreviation string `json:"CountryAbbreviation"`
		Countryname         string `json:"CountryName"`
		Domainname          string `json:"DomainName"`
		Dst                 string `json:"DST"`
		Ipaddress           string `json:"IPAddress"`
		Ispname             string `json:"ISPName"`
		Latitude            string `json:"Latitude"`
		Longitude           string `json:"Longitude"`
		Postalcode          string `json:"PostalCode"`
		Proxydescription    string `json:"ProxyDescription"`
		Proxytype           string `json:"ProxyType"`
		Recordid            string `json:"RecordID"`
		Region              string `json:"Region"`
		Result              string `json:"Result"`
		Timezonecode        string `json:"TimeZoneCode"`
		Timezonename        string `json:"TimeZoneName"`
		Utc                 string `json:"UTC"`
	} `json:"Records"`
}

/* ///////////////////////////////////////////////////////

   MELISSA

*/ ///////////////////////////////////////////////////////

// MelissaAddress is a type that represents the results of Melissa's GlobalIP API.
type MelissaAddress struct {
	City                string
	Connectionspeed     string
	Connectiontype      string
	Continent           string
	Countryabbreviation string
	Countryname         string
	Domainname          string
	Dst                 string
	Ipaddress           string
	Ispname             string
	Latitude            string
	Longitude           string
	Postalcode          string
	Proxydescription    string
	Proxytype           string
	Recordid            string
	Region              string
	Result              string
	Timezonecode        string
	Timezonename        string
	Utc                 string
}

// MelissaLookup takes a Melissa API key (string) and an IP address (string) as its parameters which are passed through the Melissa Global IP API whose response is then represented by a *MelissaAddress type.
func MelissaLookup(key, ip string) (*MelissaAddress, error) {
	resp, err := http.GetReq("https://globalip.melissadata.net/v4/web/iplocation/doiplocation?id=" + key + "&ip=" + ip)
	if err != nil {
		return nil, err
	}
	response := new(ipResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(MelissaAddress)
	for _, v := range response.Records {
		newRes.City = v.City
		newRes.Connectionspeed = v.Connectionspeed
		newRes.Connectiontype = v.Connectiontype
		newRes.Continent = v.Continent
		newRes.Countryabbreviation = v.Countryabbreviation
		newRes.Countryname = v.Countryname
		newRes.Domainname = v.Domainname
		newRes.Dst = v.Dst
		newRes.Ipaddress = v.Ipaddress
		newRes.Ispname = v.Ispname
		newRes.Latitude = v.Latitude
		newRes.Longitude = v.Longitude
		newRes.Postalcode = v.Postalcode
		newRes.Proxydescription = v.Proxydescription
		newRes.Proxytype = v.Proxytype
		newRes.Recordid = v.Recordid
		newRes.Region = v.Region
		newRes.Result = v.Result
		newRes.Timezonecode = v.Timezonecode
		newRes.Timezonename = v.Timezonename
		newRes.Utc = v.Utc

		break
	}
	return newRes, nil
}

/* ///////////////////////////////////////////////////////

   IP2LOCATION

*/ ///////////////////////////////////////////////////////

// IPTLAddress is a type that represents the results of the IP2LOCATION IP API.
type IPTLAddress struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionName  string  `json:"region_name"`
	CityName    string  `json:"city_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Asn         string  `json:"asn"`
	As          string  `json:"as"`
	IsProxy     bool    `json:"is_proxy"`
}

// IPTLLookup takes an ip2location/ip2whois key and an IPV4 address as its parameters which are passed through the IP2LOCATION API whose response is then represented by a *IpData type.
func IPTLLookup(key, ip string) (*IPTLAddress, error) {
	resp, err := http.GetReq("https://api.ip2location.io/?key=" + key + "&ip=" + ip + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(IPTLAddress)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
