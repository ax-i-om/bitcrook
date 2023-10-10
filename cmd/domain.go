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

// Package cmd ...
package cmd

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/ax-i-om/bitcrook/pkg/authfree/ip2location"
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain <domain>",
	Short: "Information about a domain",
	Long: `
The 'domain' command takes a domain name as an argument and
returns any associated information that is identified via the 
IP2LOCATION Whois API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "api.ip2whois.com\n"))

			key1 := os.Getenv("BITCROOK_IPTL")
			if key1 == "UNSPECIFIED" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify ip2location API key in .env file")
			} else {
				x, err := ip2location.DomainLookup(key1, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), err)
					return
				}
				if x.DomainID == "" || len(x.DomainID) < 1 {
					fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
					return
				} else {
					// const layout = "2006-Jan-02"
					/*
						crDate, crErr := time.Parse(layout, x.CreateDate.String())
						pDate, pErr := time.Parse(layout, x.UpdateDate.String())
						xDate, xErr := time.Parse(layout, x.ExpireDate.String())
					*/
					fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))
					fmt.Println(color.Colorize(color.Blue, "--- [ General Information ] ---"))
					fmt.Println("Domain:\t\t", x.Domain)
					fmt.Println("Domain ID:\t", x.DomainID)
					fmt.Println("Status:\t\t", x.Status)
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
							fmt.Println("Expiration Date:\t",xDate.String())
						}
					*/
					fmt.Println("Domain Age:\t", fmt.Sprint(x.DomainAge), "days")
					fmt.Println("Whois Server:\t", x.WhoisServer)
					fmt.Println(color.Colorize(color.Blue, "\n--- [ Registrar ] ---"))
					fmt.Println("IanaID:\t\t", x.Registrar.IanaID)
					fmt.Println("Name:\t\t", x.Registrar.Name)
					fmt.Println("URL:\t\t", x.Registrar.URL)
					if x.Registrant.Name != "" {
						fmt.Println(color.Colorize(color.Blue, "\n--- [ Registrant ] ---"))
						fmt.Println("Name:\t\t", x.Registrant.Name)
						fmt.Println("Organization:\t", x.Registrant.Organization)
						fmt.Println("Street Address:\t", x.Registrant.StreetAddress)
						fmt.Println("City:\t\t", x.Registrant.City)
						fmt.Println("Region:\t\t", x.Registrant.Region)
						fmt.Println("Zip Code:\t", x.Registrant.ZipCode)
						fmt.Println("Country:\t", x.Registrant.Country)
						fmt.Println("Phone:\t\t", x.Registrant.Phone)
						fmt.Println("Fax:\t\t", x.Registrant.Fax)
						fmt.Println("Email:\t\t", x.Registrant.Email)
					}
					if x.Admin.Name != "" {
						fmt.Println(color.Colorize(color.Blue, "\n--- [ Admin ] ---"))
						fmt.Println("Name:\t\t", x.Admin.Name)
						fmt.Println("Organization:\t", x.Admin.Organization)
						fmt.Println("Street Address:\t", x.Admin.StreetAddress)
						fmt.Println("City:\t\t", x.Admin.City)
						fmt.Println("Region:\t\t", x.Admin.Region)
						fmt.Println("Zip Code:\t", x.Admin.ZipCode)
						fmt.Println("Country:\t", x.Admin.Country)
						fmt.Println("Phone:\t\t", x.Admin.Phone)
						fmt.Println("Fax:\t\t", x.Admin.Fax)
						fmt.Println("Email:\t\t", x.Admin.Email)
					}
					if x.Tech.Name != "" {
						fmt.Println(color.Colorize(color.Blue, "\n--- [ Tech ] ---"))
						fmt.Println("Name:\t\t", x.Tech.Name)
						fmt.Println("Organization:\t", x.Tech.Organization)
						fmt.Println("Street Address:\t", x.Tech.StreetAddress)
						fmt.Println("City:\t\t", x.Tech.City)
						fmt.Println("Region:\t\t", x.Tech.Region)
						fmt.Println("Zip Code:\t", x.Tech.ZipCode)
						fmt.Println("Country:\t", x.Tech.Country)
						fmt.Println("Phone:\t\t", x.Tech.Phone)
						fmt.Println("Fax:\t\t", x.Tech.Fax)
						fmt.Println("Email:\t\t", x.Tech.Email)
					}

					if x.Billing.Name != "" {
						fmt.Println(color.Colorize(color.Blue, "\n--- [ Billing ] ---"))
						fmt.Println("Name:\t\t", x.Billing.Name)
						fmt.Println("Organization:\t", x.Billing.Organization)
						fmt.Println("Street Address:\t", x.Billing.StreetAddress)
						fmt.Println("City:\t\t", x.Billing.City)
						fmt.Println("Region:\t\t", x.Billing.Region)
						fmt.Println("Zip Code:\t", x.Billing.ZipCode)
						fmt.Println("Country:\t", x.Billing.Country)
						fmt.Println("Phone:\t\t", x.Billing.Phone)
						fmt.Println("Fax:\t\t", x.Billing.Fax)
						fmt.Println("Email:\t\t", x.Billing.Email)
					}

					fmt.Println(color.Colorize(color.Blue, "\n--- [ Nameservers ] ---"))
					for _, v := range x.Nameservers {
						fmt.Println(v)
					}
					fmt.Println()
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
