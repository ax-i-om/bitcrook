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
	"github.com/ax-i-om/bitcrook/pkg/phone"
	"github.com/spf13/cobra"
)

// phoneCmd represents the phone command
var phoneCmd = &cobra.Command{
	Use:   "phone <phone>",
	Short: "Information about a phone number",
	Long: `
The 'phone' command takes a phone number as an argument and
returns any associated information that is identified via Melissa's
globalphone API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "globalphone.melissadata.net\n"))

			key1 := os.Getenv("BITCROOK_MLSA")
			if key1 == "UNSPECIFIED" || key1 == "" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify Melissa API key in .env file")
			} else {
				x, err := phone.MelissaLookup(key1, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), err)
					return
				}

				if x.Phonenumber == "" {
					fmt.Println("STATUS:\t\t\t", color.Colorize(color.Red, "FAILURE\n"))
				} else {
					fmt.Println("STATUS:\t\t\t", color.Colorize(color.Green, "SUCCESS\n"))
					fmt.Println("Record ID:\t\t", x.Recordid)
					fmt.Println("Results:\t\t", x.Results)
					fmt.Println("Phone Number:\t\t", x.Phonenumber)
					fmt.Println("Administrative Area:\t", x.Administrativearea)
					fmt.Println("Country Abbreviation:\t", x.Countryabbreviation)
					fmt.Println("Country Name:\t\t", x.Countryname)
					fmt.Println("Carrier:\t\t", x.Carrier)
					fmt.Println("Caller ID:\t\t\t", x.Callerid)
					fmt.Println("DST:\t\t\t", x.Dst)
					fmt.Println("Intnl. Phone Number:\t", x.Internationalphonenumber)
					fmt.Println("Language:\t\t", x.Language)
					fmt.Println("Latitude:\t\t", x.Latitude)
					fmt.Println("Locality:\t\t", x.Locality)
					fmt.Println("Longitude:\t\t", x.Longitude)
					fmt.Println("International Prefix:\t\t", x.Phoneinternationalprefix)
					fmt.Println("Country Dialing Code:\t", x.Phonecountrydialingcode)
					fmt.Println("Ntnl. Prefix:\t\t", x.Phonenationprefix)
					fmt.Println("Ntnl. Destination Code:\t", x.Phonenationaldestinationcode)
					fmt.Println("Subscriber Number:\t", x.Phonesubscribernumber)
					fmt.Println("UTC:\t\t\t", x.Utc)
					fmt.Println("Time Zone Code:\t\t", x.Timezonecode)
					fmt.Println("Time Zone Name:\t\t", x.Timezonename)
					fmt.Println("Postal Code:\t\t", x.Postalcode)
					fmt.Println()
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(phoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// phoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// phoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
