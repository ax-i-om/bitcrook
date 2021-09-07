/*
Copyright © 2021 Maraudery <maraudery@protonmail.com>

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

// Package cmd ...
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/maraudery/qualear/internal/config"
	"github.com/maraudery/qualear/pkg/authfree/melissa"
	"github.com/spf13/cobra"
)

// freeauthMelissaCmd represents the freeauthMelissa command
var freeauthMelissaCmd = &cobra.Command{
	Use:   "freeauthMelissa <option (email/phone/ip)> <args>",
	Short: "Return a plethora of information about an email/phone/ip.",
	Long:  `Using the Melissa lookups API, return a plethora of information regarding an email address, phone number, or an IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
		} else {
			elConf, err := config.LoadConfig("keyconfig.json")
			if err != nil {
				log.Panic(err)
			}
			if strings.ToLower(args[0]) == "email" {
				x, err := melissa.EmailLookup(elConf.MelissaKey, args[1])
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("Record ID: " + x.Recordid)
				fmt.Println("Deliverability Confidence Score: " + x.Deliverabilityconfidencescore)
				fmt.Println("Results: " + x.Results)
				fmt.Println("Email Address: " + x.Emailaddress)
				fmt.Println("Mailbox Name: " + x.Mailboxname)
				fmt.Println("Domain Name: " + x.Domainname)
				fmt.Println("Top Level Domain: " + x.Topleveldomain)
				fmt.Println("Top Level Domain Name: " + x.Topleveldomainname)
				fmt.Println("Date Checked: " + x.Datechecked)
				fmt.Println("Email Age Estimated: " + x.Emailageestimated)
				fmt.Println("Domain Age Estimated: " + x.Domainageestimated)
				fmt.Println("Domain Expiration Date: " + x.Domainexpirationdate)
				fmt.Println("Domain Creation Date: " + x.Domaincreateddate)
				fmt.Println("Domain Updated Date: " + x.Domainupdateddate)
				fmt.Println("Domain Email: " + x.Domainemail)
				fmt.Println("Domain Organization: " + x.Domainorganization)
				fmt.Println("Domain Address: " + x.Domainaddress1)
				fmt.Println("Domain Locality: " + x.Domainlocality)
				fmt.Println("Domain Administrative Area: " + x.Domainadministrativearea)
				fmt.Println("Domain Postal Code: " + x.Domainpostalcode)
				fmt.Println("Domain Country: " + x.Domaincountry)
				fmt.Println("Domain Availability: " + x.Domainavailability)
				fmt.Println("Domain Country Code: " + x.Domaincountrycode)
				fmt.Println("Domain Private Proxy: " + x.Domainprivateproxy)
				fmt.Println("Privacy Flag: " + x.Privacyflag)
				fmt.Println("MX Server: " + x.Mxserver)
			} else if strings.ToLower(args[0]) == "phone" {
				x, err := melissa.PhoneLookup(elConf.MelissaKey, args[1])
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("Record ID: " + x.Recordid)
				fmt.Println("Results: " + x.Results)
				fmt.Println("Phone Number: " + x.Phonenumber)
				fmt.Println("Administrative Area: " + x.Administrativearea)
				fmt.Println("Country Abbreviation: " + x.Countryabbreviation)
				fmt.Println("Country Name: " + x.Countryname)
				fmt.Println("Carrier: " + x.Carrier)
				fmt.Println("Caller ID: " + x.Callerid)
				fmt.Println("DST: " + x.Dst)
				fmt.Println("International Phone Number: " + x.Internationalphonenumber)
				fmt.Println("Language: " + x.Language)
				fmt.Println("Latitude: " + x.Latitude)
				fmt.Println("Locality: " + x.Locality)
				fmt.Println("Longitude: " + x.Longitude)
				fmt.Println("Phone International Prefix: " + x.Phoneinternationalprefix)
				fmt.Println("Phone Country Dialing Code: " + x.Phonecountrydialingcode)
				fmt.Println("Phone National Prefix: " + x.Phonenationprefix)
				fmt.Println("Phone National Destination Code: " + x.Phonenationaldestinationcode)
				fmt.Println("Phone Subscriber Number: " + x.Phonesubscribernumber)
				fmt.Println("UTC: " + x.Utc)
				fmt.Println("Time Zone Code: " + x.Timezonecode)
				fmt.Println("Time Zone Name: " + x.Timezonename)
				fmt.Println("Postal Code: " + x.Postalcode)
			} else if strings.ToLower(args[0]) == "ip" {
				x, err := melissa.IPLookup(elConf.MelissaKey, args[1])
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("City: " + x.City)
				fmt.Println("Connection Speed: " + x.Connectionspeed)
				fmt.Println("Connection Type: " + x.Connectiontype)
				fmt.Println("Continent: " + x.Continent)
				fmt.Println("Country Abbreviation: " + x.Countryabbreviation)
				fmt.Println("Domain Name: " + x.Domainname)
				fmt.Println("DST: " + x.Dst)
				fmt.Println("IP Address: " + x.Ipaddress)
				fmt.Println("ISP: " + x.Ispname)
				fmt.Println("Latitude: " + x.Latitude)
				fmt.Println("Longitude: " + x.Longitude)
				fmt.Println("Postal Code: " + x.Postalcode)
				fmt.Println("Proxy Description: " + x.Proxydescription)
				fmt.Println("Proxy Type: " + x.Proxytype)
				fmt.Println("Record ID: " + x.Recordid)
				fmt.Println("Region: " + x.Region)
				fmt.Println("Result: " + x.Result)
				fmt.Println("Time Zone Code: " + x.Timezonecode)
				fmt.Println("Time Zone Name: " + x.Timezonename)
				fmt.Println("UTC: " + x.Utc)
			} else {
				cmd.Usage()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(freeauthMelissaCmd)
}
