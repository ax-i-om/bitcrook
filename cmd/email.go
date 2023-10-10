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
	"github.com/ax-i-om/bitcrook/pkg/authfree/melissa"
	"github.com/spf13/cobra"
)

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email <email>",
	Short: "Information about an email address",
	Long: `
The 'email' command takes an email address as an argument and
returns any associated information that is identified via Melissa's
globalemail API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "globalemail.melissadata.net\n"))

			key1 := os.Getenv("BITCROOK_MLSA")
			if key1 == "UNSPECIFIED" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify Melissa API key in .env file")
			} else {
				x, err := melissa.EmailLookup(key1, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), err)
					return
				}

				if x.Emailaddress == "" {
					fmt.Println("STATUS:\t\t\t\t", color.Colorize(color.Red, "FAILURE\n"))
				} else {
					fmt.Println("STATUS:\t\t\t\t", color.Colorize(color.Green, "SUCCESS\n"))
					fmt.Println("Record ID:\t\t\t", x.Recordid)
					fmt.Println("Delivery Confidence Score:\t", x.Deliverabilityconfidencescore)
					fmt.Println("Results:\t\t\t", x.Results)
					fmt.Println("Email Address:\t\t\t", x.Emailaddress)
					fmt.Println("Mailbox Name:\t\t\t", x.Mailboxname)
					fmt.Println("Domain Name:\t\t\t", x.Domainname)
					fmt.Println("Top Level Domain:\t\t", x.Topleveldomain)
					fmt.Println("Top Level Domain Name:\t\t", x.Topleveldomainname)
					fmt.Println("Date Checked:\t\t\t", x.Datechecked)
					fmt.Println("Email Age Estimated:\t\t", x.Emailageestimated)
					fmt.Println("Domain Age Estimated:\t\t", x.Domainageestimated)
					fmt.Println("Domain Expiration Date:\t\t", x.Domainexpirationdate)
					fmt.Println("Domain Creation Date:\t\t", x.Domaincreateddate)
					fmt.Println("Domain Updated Date:\t\t", x.Domainupdateddate)
					fmt.Println("Domain Email:\t", x.Domainemail)
					fmt.Println("Domain Organization:\t\t", x.Domainorganization)
					fmt.Println("Domain Address:\t", x.Domainaddress1)
					fmt.Println("Domain Locality:\t", x.Domainlocality)
					fmt.Println("Domain Administrative Area:\t", x.Domainadministrativearea)
					fmt.Println("Domain Postal Code:\t", x.Domainpostalcode)
					fmt.Println("Domain Country:\t\t\t", x.Domaincountry)
					fmt.Println("Domain Availability:\t\t", x.Domainavailability)
					fmt.Println("Domain Country Code:\t\t", x.Domaincountrycode)
					fmt.Println("Domain Private Proxy:\t", x.Domainprivateproxy)
					fmt.Println("Privacy Flag:\t\t\t", x.Privacyflag)
					fmt.Println("MX Server:\t", x.Mxserver)
					fmt.Println()
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
