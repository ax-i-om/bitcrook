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

// Package phone contains the types and functions used for querying for
// information regarding a phone number
package phone

import (
	"encoding/json"

	"github.com/ax-i-om/bitcrook/internal/http"
)

/* ///////////////////////////////////////////////////////

   MELISSA

*/ ///////////////////////////////////////////////////////

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

// MelissaPhone is a type that represents the results of Melissa's GlobalPhone API.
type MelissaPhone struct {
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

// PhoneLookup takes a Melissa API key and a phone number as its parameters which are passed through the Melissa Global Phone API whose response is then represented by a *MelissaPhone type.
func MelissaLookup(key, phone string) (*MelissaPhone, error) {
	resp, err := http.GetReq("https://globalphone.melissadata.net/v4/WEB/GlobalPhone/doGlobalPhone?id=" + key + "&phone=" + phone)
	if err != nil {
		return nil, err
	}
	response := new(phoneResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(MelissaPhone)
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
