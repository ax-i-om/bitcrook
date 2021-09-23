package hibp

import (
	"encoding/json"
	"time"

	"github.com/maraudery/goseek/internal/http"
)

// BreachList is a representation of the names of site breaches returned by the GetBreaches() function.
type BreachList []struct {
	Name string `json:"Name"`
}

// BreachedSites is a representation of a plethora of detailed information regarding the database breaches indexed in HIBP.
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

// BreachedSite is a representation of a plethora of detailed information regarding a specific breach.
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

// AllPastes is a representation of all of the pastes associated with the specified email.
type AllPastes []struct {
	ID         string      `json:"Id"`
	Source     string      `json:"Source"`
	Title      string      `json:"Title"`
	Date       interface{} `json:"Date"`
	Emailcount int         `json:"EmailCount"`
}

// GetBreaches takes a Have I Been Pwned API key and an email as its parameters and returns a list of related breaches (if any).
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

// GetBreachedSites takes a Have I Been Pwned API key and an email as its parameters and returns a list of related breaches (if any).
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

// GetBreachedSite takes a siteName as a parameter and returns a type *BreachedSite as well as an error.
// The return value contains information about the specified site and its corresponding data breach.
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

// GetAllPastes takes an email and an API key as its parameters and returns a type *AllPastes as well as an error.
// The return value contains all of the pastes that are associated with the email in question.
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
