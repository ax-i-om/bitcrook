/*
Copyright © 2021 audioo <audioo@protonmail.com>
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

	"github.com/audioo/bitcrook/internal/config"
	"github.com/audioo/bitcrook/pkg/authfree/ip2whois"
	"github.com/spf13/cobra"
)

// freeauthDomainCmd represents the freeauthDomain command
var freeauthDomainCmd = &cobra.Command{
	Use:   "freeauthDomain <domain name>",
	Short: "Returns public information regarding a domain. Auth: FREE",
	Long:  `Returns public information regarding a domain. Auth: FREE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			elConf, err := config.LoadConfig("keyconfig.json")
			if err != nil {
				log.Panic(err)
			}
			x, err := ip2whois.DomainLookup(elConf.Ip2WhoisKey, args[0])
			if err != nil {
				log.Panic(err)
			}
			if x.DomainID == "" || len(x.DomainID) < 1 {
				fmt.Println("Invalid Domain")
			} else {
				// const layout = "2006-Jan-02"
				/*
					crDate, crErr := time.Parse(layout, x.CreateDate.String())
					pDate, pErr := time.Parse(layout, x.UpdateDate.String())
					xDate, xErr := time.Parse(layout, x.ExpireDate.String())
				*/
				fmt.Println("General Information:\n--------------------")
				fmt.Println("Domain: " + x.Domain)
				fmt.Println("Domain ID: " + x.DomainID)
				fmt.Println("Status: " + x.Status)
				/*
					if crErr != nil {
						fmt.Println("Creation Date: ")
					} else {
						fmt.Println("Creation Date: " + crDate.String())
					}
					if pErr != nil {
						fmt.Println("Update Date: ")
					} else {
						fmt.Println("Update Date: " + pDate.String())
					}
					if xErr != nil {
						fmt.Println("Expiration Date: ")
					} else {
						fmt.Println("Expiration Date: " + xDate.String())
					}
				*/
				fmt.Println("Domain Age (Days): " + fmt.Sprint(x.DomainAge))
				fmt.Println("Whois Server: " + x.WhoisServer)
				fmt.Println("\nRegistrar:\n--------------------")
				fmt.Println("IanaID: " + x.Registrar.IanaID)
				fmt.Println("Name: " + x.Registrar.Name)
				fmt.Println("URL: " + x.Registrar.URL)
				fmt.Println("\nRegistrant:\n--------------------")
				fmt.Println("Name: " + x.Registrant.Name)
				fmt.Println("Organization: " + x.Registrant.Organization)
				fmt.Println("Street Address: " + x.Registrant.StreetAddress)
				fmt.Println("City: " + x.Registrant.City)
				fmt.Println("Region: " + x.Registrant.Region)
				fmt.Println("Zip Code: " + x.Registrant.ZipCode)
				fmt.Println("Country: " + x.Registrant.Country)
				fmt.Println("Phone: " + x.Registrant.Phone)
				fmt.Println("Fax: " + x.Registrant.Fax)
				fmt.Println("Email: " + x.Registrant.Email)
				fmt.Println("\nAdmin:\n--------------------")
				fmt.Println("Name: " + x.Admin.Name)
				fmt.Println("Organization: " + x.Admin.Organization)
				fmt.Println("Street Address: " + x.Admin.StreetAddress)
				fmt.Println("City: " + x.Admin.City)
				fmt.Println("Region: " + x.Admin.Region)
				fmt.Println("Zip Code: " + x.Admin.ZipCode)
				fmt.Println("Country: " + x.Admin.Country)
				fmt.Println("Phone: " + x.Admin.Phone)
				fmt.Println("Fax: " + x.Admin.Fax)
				fmt.Println("Email: " + x.Admin.Email)
				fmt.Println("\nTech:\n--------------------")
				fmt.Println("Name: " + x.Tech.Name)
				fmt.Println("Organization: " + x.Tech.Organization)
				fmt.Println("Street Address: " + x.Tech.StreetAddress)
				fmt.Println("City: " + x.Tech.City)
				fmt.Println("Region: " + x.Tech.Region)
				fmt.Println("Zip Code: " + x.Tech.ZipCode)
				fmt.Println("Country: " + x.Tech.Country)
				fmt.Println("Phone: " + x.Tech.Phone)
				fmt.Println("Fax: " + x.Tech.Fax)
				fmt.Println("Email: " + x.Tech.Email)
				fmt.Println("\nBilling:\n--------------------")
				fmt.Println("Name: " + x.Billing.Name)
				fmt.Println("Organization: " + x.Billing.Organization)
				fmt.Println("Street Address: " + x.Billing.StreetAddress)
				fmt.Println("City: " + x.Billing.City)
				fmt.Println("Region: " + x.Billing.Region)
				fmt.Println("Zip Code: " + x.Billing.ZipCode)
				fmt.Println("Country: " + x.Billing.Country)
				fmt.Println("Phone: " + x.Billing.Phone)
				fmt.Println("Fax: " + x.Billing.Fax)
				fmt.Println("Email: " + x.Billing.Email)
				fmt.Println("\nName Servers:\n--------------------")
				for _, v := range x.Nameservers {
					fmt.Println(v)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(freeauthDomainCmd)
}
