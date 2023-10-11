package domain

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ax-i-om/bitcrook/internal/http"
)

/* ///////////////////////////////////////////////////////

   HaveIBeenPwned

*/ ///////////////////////////////////////////////////////

type hibpSites []struct {
	Name               string    `json:"Name"`
	Title              string    `json:"Title"`
	Domain             string    `json:"Domain"`
	BreachDate         string    `json:"BreachDate"`
	AddedDate          time.Time `json:"AddedDate"`
	ModifiedDate       time.Time `json:"ModifiedDate"`
	PwnCount           int       `json:"PwnCount"`
	Description        string    `json:"Description"`
	LogoPath           string    `json:"LogoPath"`
	DataClasses        []string  `json:"DataClasses"`
	IsVerified         bool      `json:"IsVerified"`
	IsFabricated       bool      `json:"IsFabricated"`
	IsSensitive        bool      `json:"IsSensitive"`
	IsRetired          bool      `json:"IsRetired"`
	IsSpamList         bool      `json:"IsSpamList"`
	IsMalware          bool      `json:"IsMalware"`
	IsSubscriptionFree bool      `json:"IsSubscriptionFree"`
}

// HIBPSite is a representation of a dump of all breached sites indexed by HIBP
type HIBPSite struct {
	Name               string    `json:"Name"`
	Title              string    `json:"Title"`
	Domain             string    `json:"Domain"`
	BreachDate         string    `json:"BreachDate"`
	AddedDate          time.Time `json:"AddedDate"`
	ModifiedDate       time.Time `json:"ModifiedDate"`
	PwnCount           int       `json:"PwnCount"`
	Description        string    `json:"Description"`
	LogoPath           string    `json:"LogoPath"`
	DataClasses        []string  `json:"DataClasses"`
	IsVerified         bool      `json:"IsVerified"`
	IsFabricated       bool      `json:"IsFabricated"`
	IsSensitive        bool      `json:"IsSensitive"`
	IsRetired          bool      `json:"IsRetired"`
	IsSpamList         bool      `json:"IsSpamList"`
	IsMalware          bool      `json:"IsMalware"`
	IsSubscriptionFree bool      `json:"IsSubscriptionFree"`
}

// HIBPLookup takes a domain as a parameter and returns a type *HIBP as well as an error.
// The return value contains information about the specified site and its corresponding data breach.
func HIBPLookup(domain string) ([]HIBPSite, error) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breaches")
	if err != nil {
		return nil, err
	}
	response := new(hibpSites)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	var results []HIBPSite
	for _, v := range *response {
		if strings.Contains(strings.ToLower(domain), strings.ToLower(v.Name)) {
			x := new(HIBPSite)
			x.Name = v.Name
			x.Title = v.Title
			x.Domain = v.Domain
			x.BreachDate = v.BreachDate
			x.AddedDate = v.AddedDate
			x.ModifiedDate = v.ModifiedDate
			x.PwnCount = v.PwnCount
			x.Description = v.Description
			x.LogoPath = v.LogoPath
			x.DataClasses = v.DataClasses
			x.IsVerified = v.IsVerified
			x.IsFabricated = v.IsFabricated
			x.IsSensitive = v.IsSensitive
			x.IsRetired = v.IsRetired
			x.IsSpamList = v.IsSpamList
			x.IsMalware = v.IsMalware
			x.IsSubscriptionFree = v.IsSubscriptionFree
			results = append(results, *x)
		}
	}
	return results, nil
}

/* ///////////////////////////////////////////////////////

   IP2LOCATION

*/ ///////////////////////////////////////////////////////

// IPTLDomain is a type that represents the results of the IP2LOCATION Whois API.
type IPTLDomain struct {
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

// IPTLLookup takes an ip2location/ip2whois key and a domain as its parameters which are passed through the IP2WHOIS API whose response is then represented by a *IPTLDomain type.
func IPTLLookup(key, domain string) (*IPTLDomain, error) {
	resp, err := http.GetReq("https://api.ip2whois.com/v2?key=" + key + "&domain=" + domain + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(IPTLDomain)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
