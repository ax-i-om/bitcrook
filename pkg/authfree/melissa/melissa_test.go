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

package melissa

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ax-i-om/bitcrook/internal/http"
)

func TestEmailLookup(t *testing.T) {
	key1 := os.Getenv("BITCROOK_MLSA")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	resp, err := http.GetReq("https://globalemail.melissadata.net/v4/WEB/GlobalEmail/doGlobalEmail?id=" + key1 + "&email=email@example.com&format=json")
	if err != nil {
		t.Error(err)
	}
	response := new(emailResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
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
}

func TestPhoneLookup(t *testing.T) {
	key1 := os.Getenv("BITCROOK_MLSA")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	resp, err := http.GetReq("https://globalphone.melissadata.net/v4/WEB/GlobalPhone/doGlobalPhone?id=" + key1 + "&phone=9142329901")
	if err != nil {
		t.Error(err)
	}
	response := new(phoneResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
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
}

func TestIPLookup(t *testing.T) {
	key1 := os.Getenv("BITCROOK_MLSA")
	if key1 == "UNSPECIFIED" {
		t.Error()
	}
	resp, err := http.GetReq("https://globalip.melissadata.net/v4/web/iplocation/doiplocatio?id=" + key1 + "&ip=1.1.1.1")
	if err != nil {
		t.Error(err)
	}
	response := new(ipResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
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
}
