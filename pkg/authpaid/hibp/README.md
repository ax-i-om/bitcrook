# [github.com/bitcrook/cybull/pkg/authpaid/hibp](https://github.com/bitcrook/cybull/tree/main/pkg/authpaid/hibp) - paid authentication required


## Types

Type BreachList represents the names of site breaches returned by the GetBreaches() function.
``` go
type BreachList []struct {
	Name string `json:"Name"`
}
```

Type BreachedSites represents a plethora of detailed information regarding the database breaches indexed in HIBP.
``` go
type BreachedSites []struct {
	Name         string    `json:"Name"`
	Title        string    `json:"Title"`
	Domain       string    `json:"Domain"`
	Breachdate   string    `json:"BreachDate"`
	Addeddate    time.Time `json:"AddedDate"`
	Modifieddate time.Time `json:"ModifiedDate"`
	Pwncount     int       `json:"PwnCount"`
	Description  string    `json:"Description"`
	Logopath     string    `json:"LogoPath"`
	Dataclasses  []string  `json:"DataClasses"`
	Isverified   bool      `json:"IsVerified"`
	Isfabricated bool      `json:"IsFabricated"`
	Issensitive  bool      `json:"IsSensitive"`
	Isretired    bool      `json:"IsRetired"`
	Isspamlist   bool      `json:"IsSpamList"`
}
```

Type BreachedSite represents a plethora of detailed information regarding a specific breach.
``` go
type BreachedSite struct {
	Name         string    `json:"Name"`
	Title        string    `json:"Title"`
	Domain       string    `json:"Domain"`
	Breachdate   string    `json:"BreachDate"`
	Addeddate    time.Time `json:"AddedDate"`
	Modifieddate time.Time `json:"ModifiedDate"`
	Pwncount     int       `json:"PwnCount"`
	Description  string    `json:"Description"`
	Logopath     string    `json:"LogoPath"`
	Dataclasses  []string  `json:"DataClasses"`
	Isverified   bool      `json:"IsVerified"`
	Isfabricated bool      `json:"IsFabricated"`
	Issensitive  bool      `json:"IsSensitive"`
	Isretired    bool      `json:"IsRetired"`
	Isspamlist   bool      `json:"IsSpamList"`
}
```

Type AllPastes represents all of the pastes associated with the specified email.
``` go
type AllPastes []struct {
	ID         string      `json:"Id"`
	Source     string      `json:"Source"`
	Title      string      `json:"Title"`
	Date       interface{} `json:"Date"`
	Emailcount int         `json:"EmailCount"`
}
```

## Functions


GetBreaches takes a Have I Been Pwned API key and an email as its parameters and returns a list of related breaches (if any).
``` go
func GetBreaches(key string, email string) (*BreachList, error) {
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/breachedaccount/"+email, "hibp-api-key", key)
	if err != nil {
		return nil, err
	}
	response := new(BreachList)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
```

GetBreachedSites takes a Have I Been Pwned API key and an email as its parameters and returns a list of related breaches (if any).
``` go
func GetBreachedSites() (*BreachedSites, error) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breaches")
	if err != nil {
		return nil, err
	}
	response := new(BreachedSites)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
```

GetBreachedSite takes a siteName as a parameter and returns a type *BreachedSite as well as an error.
The return value contains information about the specified site and its corresponding data breach.
``` go
func GetBreachedSite(siteName string) (*BreachedSite, error) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breach/" + siteName)
	if err != nil {
		return nil, err
	}
	response := new(BreachedSite)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
```

GetAllPastes takes an email and an API key as its parameters and returns a type *AllPastes as well as an error.
The return value contains all of the pastes that are associated with the email in question.
``` go
func GetAllPastes(key string, email string) (*AllPastes, error) {
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/pasteaccount/"+email, "hibp-api-key", key)
	if err != nil {
		return nil, err
	}
	response := new(AllPastes)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
```

## Usage

The data contained in these types can all be read in a similar manner. For example, the IPLookup function:
``` go
package main

import (
	"fmt"
	"log"

	"github.com/bitcrook/cybull/pkg/authpaid/hibp"
)

func main() {
	x, err := hibp.GetBreaches("yourKeyHere", "email@example.com")
	if err != nil {
		log.Panic(err)
	}
	for _, v := range *x {
		fmt.Println(v.Name)
	}
}
```
Part of the Output from the above example:
```
2844Breaches
500px
8fit
8tracks
Adobe
```