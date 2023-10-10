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
	"strconv"

	"github.com/TwiN/go-color"
	"github.com/ax-i-om/bitcrook/pkg/authfree/ip2location"
	"github.com/ax-i-om/bitcrook/pkg/authfree/melissa"
	"github.com/ax-i-om/bitcrook/pkg/noauth/ip"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip <ip>",
	Short: "Information about an IP address",
	Long: `
The 'ip' command takes an IPv4/IPv6 address as an argument 
and returns any associated information that is identified.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "ip-api.com\n"))
			x, err := ip.IPLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), err)
				return
			}
			if x.Status == "fail" {
				fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
			} else {
				fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))
				fmt.Println("Continent:\t", x.Continent)
				fmt.Println("Country:\t", x.Country)
				fmt.Println("Region Name:\t", x.RegionName)
				fmt.Println("City:\t\t", x.City)
				fmt.Println("District:\t\t", x.District)
				fmt.Println("Zip:\t\t", x.Zip)
				fmt.Println("Latitude:\t", strconv.FormatFloat(x.Lat, 'f', 5, 64))
				fmt.Println("Longitude:\t", strconv.FormatFloat(x.Lon, 'f', 5, 64))
				fmt.Println("Timezone:\t", x.Timezone)
				fmt.Println("Currency:\t", x.Currency)
				fmt.Println("ISP:\t\t", x.Isp)
				fmt.Println("Org:\t\t", x.Org)
				fmt.Println("As:\t\t", x.As)
				fmt.Println("Asname:\t\t", x.Asname)
				fmt.Println("Reverse:\t", x.Reverse)
				fmt.Println("Mobile:\t\t", strconv.FormatBool(x.Mobile))
				fmt.Println("Proxy:\t\t", strconv.FormatBool(x.Proxy))
				fmt.Println("Hosting:\t", strconv.FormatBool(x.Hosting))
				fmt.Println("Query:\t\t", x.Query)
				fmt.Println("")
			}

			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "api.ip2location.io\n"))

			key1 := os.Getenv("BITCROOK_IPTL")
			if key1 == "UNSPECIFIED" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify ip2location API key in .env file")
			} else {
				x, err := ip2location.IPLookup(key1, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), err)
					return
				}

				if x.ZipCode == "" && x.TimeZone == "" && x.Asn == "" && x.As == "" {
					fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
				} else {
					fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))
					fmt.Println("Country Code:\t", x.CountryCode)
					fmt.Println("Country Name:\t", x.CountryName)
					fmt.Println("Region Name:\t", x.RegionName)
					fmt.Println("City Name:\t", x.CityName)
					fmt.Println("Latitude:\t", x.Latitude)
					fmt.Println("Longitude:\t", x.Longitude)
					fmt.Println("Zip Code:\t", x.ZipCode)
					fmt.Println("Timezone:\t", x.TimeZone)
					fmt.Println("Asn:\t\t", x.Asn)
					fmt.Println("As:\t\t", x.As)
					fmt.Println("Is Proxy:\t", x.IsProxy)
					fmt.Println("")
				}
			}

			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "globalip.melissadata.net\n"))

			key2 := os.Getenv("BITCROOK_MLSA")
			if key2 == "UNSPECIFIED" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify Melissa API key in .env file")
			} else {
				x, err := melissa.IPLookup(key2, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), err)
					return
				}
				fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))
				fmt.Println("City:\t\t", x.City)
				fmt.Println("Conn. Speed:\t", x.Connectionspeed)
				fmt.Println("Conn. Type:\t", x.Connectiontype)
				fmt.Println("Continent:\t", x.Continent)
				fmt.Println("Abrv. Country:\t", x.Countryabbreviation)
				fmt.Println("Country Name:\t", x.Countryname)
				fmt.Println("Domain Name:\t", x.Domainname)
				fmt.Println("DST:\t\t", x.Dst)
				fmt.Println("IP Address:\t", x.Ipaddress)
				fmt.Println("ISP:\t\t", x.Ispname)
				fmt.Println("Latitude:\t", x.Latitude)
				fmt.Println("Longitude:\t", x.Longitude)
				fmt.Println("Postal Code:\t", x.Postalcode)
				fmt.Println("Proxy Desc:\t", x.Proxydescription)
				fmt.Println("Proxy Type:\t", x.Proxytype)
				fmt.Println("Record ID:\t", x.Recordid)
				fmt.Println("Region:\t\t", x.Region)
				fmt.Println("Result:\t\t", x.Result)
				fmt.Println("Time Zone Code:\t", x.Timezonecode)
				fmt.Println("Time Zone Name:\t", x.Timezonename)
				fmt.Println("UTC:\t\t", x.Utc)
				fmt.Println()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
