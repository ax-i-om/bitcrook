package melissa

import (
	"encoding/json"

	"github.com/bitcrook/goseek/internal/http"
)

/*
/ BEGIN TYPES
*/

type emailResponse struct {
	Version               string `json:"Version"`
	Transmissionreference string `json:"TransmissionReference"`
	Transmissionresults   string `json:"TransmissionResults"`
	Totalrecords          string `json:"TotalRecords"`
	Records               []struct {
		Recordid                      string `json:"RecordID"`
		Deliverabilityconfidencescore string `json:"DeliverabilityConfidenceScore"`
		Results                       string `json:"Results"`
		Emailaddress                  string `json:"EmailAddress"`
		Mailboxname                   string `json:"MailboxName"`
		Domainname                    string `json:"DomainName"`
		Topleveldomain                string `json:"TopLevelDomain"`
		Topleveldomainname            string `json:"TopLevelDomainName"`
		Datechecked                   string `json:"DateChecked"`
		Emailageestimated             string `json:"EmailAgeEstimated"`
		Domainageestimated            string `json:"DomainAgeEstimated"`
		Domainexpirationdate          string `json:"DomainExpirationDate"`
		Domaincreateddate             string `json:"DomainCreatedDate"`
		Domainupdateddate             string `json:"DomainUpdatedDate"`
		Domainemail                   string `json:"DomainEmail"`
		Domainorganization            string `json:"DomainOrganization"`
		Domainaddress1                string `json:"DomainAddress1"`
		Domainlocality                string `json:"DomainLocality"`
		Domainadministrativearea      string `json:"DomainAdministrativeArea"`
		Domainpostalcode              string `json:"DomainPostalCode"`
		Domaincountry                 string `json:"DomainCountry"`
		Domainavailability            string `json:"DomainAvailability"`
		Domaincountrycode             string `json:"DomainCountryCode"`
		Domainprivateproxy            string `json:"DomainPrivateProxy"`
		Privacyflag                   string `json:"PrivacyFlag"`
		Mxserver                      string `json:"MXServer"`
	} `json:"Records"`
}

// Email is a type that represents the results of Melissa's GlobalEmail API.
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

type phoneResponse struct {
	Version               string `json:"Version"`
	Transmissionreference string `json:"TransmissionReference"`
	Transmissionresults   string `json:"TransmissionResults"`
	Records               []struct {
		Recordid                     string      `json:"RecordID"`
		Results                      string      `json:"Results"`
		Phonenumber                  string      `json:"PhoneNumber"`
		Administrativearea           string      `json:"AdministrativeArea"`
		Countryabbreviation          string      `json:"CountryAbbreviation"`
		Countryname                  string      `json:"CountryName"`
		Carrier                      string      `json:"Carrier"`
		Callerid                     string      `json:"CallerID"`
		Dst                          string      `json:"DST"`
		Internationalphonenumber     string      `json:"InternationalPhoneNumber"`
		Language                     string      `json:"Language"`
		Latitude                     string      `json:"Latitude"`
		Locality                     string      `json:"Locality"`
		Longitude                    string      `json:"Longitude"`
		Phoneinternationalprefix     string      `json:"PhoneInternationalPrefix"`
		Phonecountrydialingcode      string      `json:"PhoneCountryDialingCode"`
		Phonenationprefix            string      `json:"PhoneNationPrefix"`
		Phonenationaldestinationcode string      `json:"PhoneNationalDestinationCode"`
		Phonesubscribernumber        string      `json:"PhoneSubscriberNumber"`
		Utc                          string      `json:"UTC"`
		Timezonecode                 string      `json:"TimeZoneCode"`
		Timezonename                 string      `json:"TimeZoneName"`
		Postalcode                   string      `json:"PostalCode"`
		Suggestions                  interface{} `json:"Suggestions"`
	} `json:"Records"`
}

// Phone is a type that represents the results of Melissa's GlobalPhone API.
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

// IPAddress is a type that represents the results of Melissa's GlobalIP API.
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

/*
/ END TYPES
*/

// EmailLookup takes a Melissa API key and an email as its parameters which are passed through the Melissa Global Email API whose response is then represented by a *Email type.
func EmailLookup(key, email string) (*Email, error) {
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

// PhoneLookup takes a Melissa API key and a phone number as its parameters which are passed through the Melissa Global Phone API whose response is then represented by a *Phone type.
func PhoneLookup(key, phone string) (*Phone, error) {
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

// IPLookup takes a Melissa API key and an IPV4 address as its parameters which are passed through the Melissa Global IP API whose response is then represented by a *IPAddress type.
func IPLookup(key, ip string) (*IPAddress, error) {
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
