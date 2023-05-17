package ip2location

import (
	"encoding/json"

	"github.com/ax-i-om/bitcrook/internal/http"
)

/*
/ BEGIN TYPES
*/

// DomainData is a type that represents the results of the IP2WHOIS API.
type DomainData struct {
	Domain   string `json:"domain"`
	DomainID string `json:"domain_id"`
	Status   string `json:"status"`
	// CreateDate  time.Time `json:"create_date"`
	// UpdateDate  time.Time `json:"update_date"`
	// ExpireDate  time.Time `json:"expire_date"`
	DomainAge   int    `json:"domain_age"`
	WhoisServer string `json:"whois_server"`
	Registrar   struct {
		IanaID string `json:"iana_id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
	} `json:"registrar"`
	Registrant struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"registrant"`
	Admin struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"admin"`
	Tech struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"tech"`
	Billing struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"billing"`
	Nameservers []string `json:"nameservers"`
}

type IpData struct {
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

// DomainLookup takes an ip2location/ip2whois key and a domain as its parameters which are passed through the IP2WHOIS API whose response is then represented by a *domainResponse type.
func DomainLookup(key, domain string) (*DomainData, error) {
	resp, err := http.GetReq("https://api.ip2whois.com/v2?key=" + key + "&domain=" + domain + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(DomainData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// IPLookup takes an ip2location/ip2whois key and an IPV4 address as its parameters which are passed through the IP2LOCATION API whose response is then represented by a *IpData type.
func IPLookup(key, ip string) (*IpData, error) {
	resp, err := http.GetReq("https://api.ip2location.io/?key=" + key + "&ip=" + ip + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(IpData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
