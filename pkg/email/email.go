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

// Package email contains the types and functions used for querying for
// information regarding an email address
package email

import (
	"encoding/json"

	"github.com/ax-i-om/bitcrook/internal/http"
)

/* ///////////////////////////////////////////////////////

   MELISSA

*/ ///////////////////////////////////////////////////////

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

// MelissaEmail is a type that represents the results of Melissa's GlobalEmail API.
type MelissaEmail struct {
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

// MelissaLookup takes a Melissa API key and an email as its parameters which are passed through the Melissa Global Email API whose response is then represented by a *MelissaEmail type.
func MelissaLookup(key, email string) (*MelissaEmail, error) {
	resp, err := http.GetReq("https://globalemail.melissadata.net/v4/WEB/GlobalEmail/doGlobalEmail?id=" + key + "&email=" + email + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(emailResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(MelissaEmail)
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

/* ///////////////////////////////////////////////////////

   HaveIBeenPwned

*/ ///////////////////////////////////////////////////////

type breachList []struct {
	Name string `json:"Name"`
}

// Acc is a representation of the names of site breaches returned by the HIBPLookup() function.
type AccBreaches struct {
	Name string `json:"Name"`
}

// HIBPLookup takes a Have I Been Pwned API key and an email as its parameters and returns a list of related breaches (if any).
func HIBPLookup(key, email string) ([]AccBreaches, error) {
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/breachedaccount/"+email, "hibp-api-key", key)
	if err != nil {
		return nil, err
	}
	response := new(breachList)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}

	var results []AccBreaches
	for _, v := range *response {
		results = append(results, v)
	}
	return results, nil
}
