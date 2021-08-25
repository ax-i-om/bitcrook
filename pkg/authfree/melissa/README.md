# [github.com/maraudery/ra/pkg/authfree/melissa](https://github.com/maraudery/ra/tree/main/pkg/authfree/melissa) - free authentication required


## Types

Type Email represents the results of Melissa's GlobalEmail API.
``` go
type Email struct {
	Recordid                      string
	Deliverabilityconfidencescore string
	Results                       string
	Emailaddress                  string
	Mailboxname                   string
	Domainname                    string
	Topleveldomain                string
	Topleveldomainname            string
	Datechecked                   string
	Emailageestimated             string
	Domainageestimated            string
	Domainexpirationdate          string
	Domaincreateddate             string
	Domainupdateddate             string
	Domainemail                   string
	Domainorganization            string
	Domainaddress1                string
	Domainlocality                string
	Domainadministrativearea      string
	Domainpostalcode              string
	Domaincountry                 string
	Domainavailability            string
	Domaincountrycode             string
	Domainprivateproxy            string
	Privacyflag                   string
	Mxserver                      string
}
```

Type Phone represents the results of Melissa's GlobalPhone API.
``` go
type Phone struct {
	Recordid                     string
	Results                      string
	Phonenumber                  string
	Administrativearea           string
	Countryabbreviation          string
	Countryname                  string
	Carrier                      string
	Callerid                     string
	Dst                          string
	Internationalphonenumber     string
	Language                     string
	Latitude                     string
	Locality                     string
	Longitude                    string
	Phoneinternationalprefix     string
	Phonecountrydialingcode      string
	Phonenationprefix            string
	Phonenationaldestinationcode string
	Phonesubscribernumber        string
	Utc                          string
	Timezonecode                 string
	Timezonename                 string
	Postalcode                   string
}
```

Type IPAddress represents the results of Melissa's GlobalIP API.
``` go
type IPAddress struct {
	City                string
	Connectionspeed     string
	Connectiontype      string
	Continent           string
	Countryabbreviation string
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
```

## Functions


EmailLookup takes an API Key (string) and an email (string) as its parameters and returns a type *Email and an Error.
``` go
func EmailLookup(key string, email string) (*Email, error) {
	resp, err := http.GetReq("https://globalemail.melissadata.net/v4/WEB/GlobalEmail/doGlobalEmail?id=" + key + "&email=" + email + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(emailResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(Email)
	for _, v := range response.Records {
		newRes.Recordid = v.Recordid
		newRes.Deliverabilityconfidencescore = v.Deliverabilityconfidencescore
		newRes.Results = v.Results
		newRes.Emailaddress = v.Emailaddress
		newRes.Mailboxname = v.Mailboxname
		newRes.Domainname = v.Domainname
		newRes.Topleveldomain = v.Topleveldomain
		newRes.Topleveldomainname = v.Topleveldomainname
		newRes.Datechecked = v.Datechecked
		newRes.Emailageestimated = v.Emailageestimated
		newRes.Domainageestimated = v.Domainageestimated
		newRes.Domainexpirationdate = v.Domainexpirationdate
		newRes.Domaincreateddate = v.Domaincreateddate
		newRes.Domainupdateddate = v.Domainupdateddate
		newRes.Domainemail = v.Domainemail
		newRes.Domainorganization = v.Domainorganization
		newRes.Domainaddress1 = v.Domainaddress1
		newRes.Domainlocality = v.Domainlocality
		newRes.Domainadministrativearea = v.Domainadministrativearea
		newRes.Domainpostalcode = v.Domainpostalcode
		newRes.Domaincountry = v.Domaincountry
		newRes.Domainavailability = v.Domainavailability
		newRes.Domaincountrycode = v.Domaincountrycode
		newRes.Domainprivateproxy = v.Domainprivateproxy
		newRes.Privacyflag = v.Privacyflag
		newRes.Mxserver = v.Mxserver

		break
	}
	return newRes, nil
}
```

PhoneLookup takes a Melissa API key and a phone number as its parameters which are passed through the Melissa Global Phone API whose response is then represented by a *Phone type and an Error.
``` go
func PhoneLookup(key string, phone string) (*Phone, error) {
	resp, err := http.GetReq("https://globalphone.melissadata.net/v4/WEB/GlobalPhone/doGlobalPhone?id=" + key + "&phone=" + phone)
	if err != nil {
		return nil, err
	}
	response := new(phoneResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(Phone)
	for _, v := range response.Records {
		newRes.Recordid = v.Recordid
		newRes.Results = v.Results
		newRes.Phonenumber = v.Phonenumber
		newRes.Administrativearea = v.Administrativearea
		newRes.Countryabbreviation = v.Countryabbreviation
		newRes.Countryname = v.Countryname
		newRes.Carrier = v.Carrier
		newRes.Callerid = v.Callerid
		newRes.Dst = v.Dst
		newRes.Internationalphonenumber = v.Internationalphonenumber
		newRes.Language = v.Language
		newRes.Latitude = v.Latitude
		newRes.Locality = v.Locality
		newRes.Longitude = v.Longitude
		newRes.Phoneinternationalprefix = v.Phoneinternationalprefix
		newRes.Phonecountrydialingcode = v.Phonecountrydialingcode
		newRes.Phonenationprefix = v.Phonenationprefix
		newRes.Phonenationaldestinationcode = v.Phonenationaldestinationcode
		newRes.Phonesubscribernumber = v.Phonesubscribernumber
		newRes.Utc = v.Utc
		newRes.Timezonecode = v.Timezonecode
		newRes.Timezonename = v.Timezonename
		newRes.Postalcode = v.Postalcode

		break
	}
	return newRes, nil
}
```

IPLookup takes a Melissa API key and an IPV4 address as its parameters which are passed through the Melissa Global IP API whose response is then represented by a *IPAddress 
type and an Error.
``` go
func IPLookup(key string, ip string) (*IPAddress, error) {
	resp, err := http.GetReq("https://globalip.melissadata.net/v4/web/iplocation/doiplocatio?id=" + key + "&ip=" + ip)
	if err != nil {
		return nil, err
	}
	response := new(ipResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(IPAddress)
	for _, v := range response.Records {
		newRes.City = v.City
		newRes.Connectionspeed = v.Connectionspeed
		newRes.Connectiontype = v.Connectiontype
		newRes.Continent = v.Continent
		newRes.Countryabbreviation = v.Countryabbreviation
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
```

## Usage

The data contained in these types can all be read in a similar manner. For example, the IPLookup function:
``` go
package main

import (
	"fmt"
	"log"

	"github.com/maraudery/ra/pkg/authfree/melissa"
)

func main() {
	x, err := melissa.IPLookup("yourKeyHere", "1.1.1.1")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(x.Ispname)
}
```
Output from the above example:
```
Apnic And Cloudflare Dns Resolver Project
```