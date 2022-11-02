# [github.com/audioo/bitcrook/pkg/authfree/ip2location](https://github.com/audioo/bitcrook/tree/main/pkg/authfree/ip2location) - free authentication required


## Types

Type DomainData represents the results of the IP2WHOIS API.

``` go 
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
```

## Functions

EmailLookup takes an API Key (string) and an email (string) as its parameters and returns a type *Email and an Error.
``` go
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
```

## Usage

To do...