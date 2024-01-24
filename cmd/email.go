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
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/ax-i-om/bitcrook/internal/http"
	"github.com/ax-i-om/bitcrook/pkg/domain"
	"github.com/ax-i-om/bitcrook/pkg/email"
	"github.com/spf13/cobra"
)

func getByName(sitename string) (*domain.HIBPSite, error) {
	x, err := http.GetReq("https://haveibeenpwned.com/api/v3/breach/" + sitename)
	if err != nil {
		return nil, err
	}
	response := new(domain.HIBPSite)
	err = json.Unmarshal([]byte(x), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

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
			key1 := os.Getenv("BITCROOK_MLSA")
			if key1 == "UNSPECIFIED" || key1 == "" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify Melissa API key in .env file")
				fmt.Println()
			} else {
				x, err := email.MelissaLookup(key1, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), "Melissa:", err)
					fmt.Println()
				} else {
					if x.Emailaddress == "" {
						fmt.Println("STATUS:\t\t\t\t", color.Colorize(color.Red, "FAILURE\n"))
					} else {
						fmt.Println("STATUS:\t\t\t\t", color.Colorize(color.Green, "SUCCESS\n"))
						fmt.Println("Delivery Confidence Score:\t", x.Deliverabilityconfidencescore)
						fmt.Println("Email Address:\t\t\t", x.Emailaddress)
						fmt.Println("Mailbox Name:\t\t\t", x.Mailboxname)
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

			key2 := os.Getenv("BITCROOK_HIBP")
			if key2 == "UNSPECIFIED" || strings.ReplaceAll(key2, " ", "") == "" {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Failed to specify HaveIBeenPwned API key in .env file")
				fmt.Println()
			} else {
				x, err := email.HIBPLookup(key2, args[0])
				if err != nil {
					fmt.Println(color.Colorize(color.Red, "[x]"), "HaveIBeenPwned:", err)
					fmt.Println()
				} else {
					fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))

					fmt.Println(color.Colorize(color.Blue, "[i]"), "Email discovered in", color.Colorize(color.Green, len(x)), "breaches")
					fmt.Println()
					for _, v := range x {
						dat, err := getByName(v.Name)
						if err != nil {
							fmt.Println(color.Colorize(color.Red, "[x]"), "HaveIBeenPwned:", err)
							return
						}
						fmt.Println(color.Colorize(color.Green, v.Name+":"), dat.BreachDate)
						for _, v := range dat.DataClasses {
							fmt.Print(v + ", ")
						}
						fmt.Println()
						fmt.Println()
					}
				}
			}

			pint, err := email.PinterestLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Pinterest:", err)
				fmt.Println()
			} else {
				if pint.ResourceResponse.Data {
					fmt.Println(color.Colorize(color.Green, "[+]"), "The email", color.Colorize(color.Blue, args[0]), "is linked to a Pinterest account!")
				} else {
					fmt.Println(color.Colorize(color.Red, "[x]"), "The email", color.Colorize(color.Blue, args[0]), "is not linked to a Pinterest account!")
				}
			}
			fmt.Println()

			twit, err := email.TwitterValidityLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Twitter:", err)
				fmt.Println()
			} else {
				if twit.Taken {
					fmt.Println(color.Colorize(color.Green, "[+]"), "The email", color.Colorize(color.Blue, args[0]), "is linked to a Twitter account!")
				} else {
					fmt.Println(color.Colorize(color.Red, "[x]"), "The email", color.Colorize(color.Blue, args[0]), "is not linked to a Twitter account!")
				}
			}
			fmt.Println()

			grav, err := email.GravatarLookup(args[0])
			if err != nil {
				if strings.Contains(err.Error(), "cannot unmarshal string into Go value of type email.GravatarResponse") {
					fmt.Println(color.Colorize(color.Red, "[x]"), "The email", color.Colorize(color.Blue, args[0]), "is not linked to a Gravatar account!")
				} else {
					fmt.Println(color.Colorize(color.Red, "[x]"), "Gravatar:", err)
					fmt.Println()
				}
			} else {
				chkd := md5.Sum([]byte(args[0]))
				if grav.Entry[0].RequestHash == hex.EncodeToString(chkd[:]) {
					fmt.Println(color.Colorize(color.Green, "[+]"), "The email", color.Colorize(color.Blue, args[0]), "is linked to a Gravatar account!")
					fmt.Println()
					fmt.Println("Profile URL:\t\t", grav.Entry[0].ProfileURL)
					fmt.Println("Name:\t\t\t", grav.Entry[0].Name.Formatted)

					fmt.Println("Username:\t\t", grav.Entry[0].PreferredUsername)
					fmt.Println("Display Name:\t\t", grav.Entry[0].DisplayName)
					fmt.Println("Current Location:\t", grav.Entry[0].CurrentLocation)
					fmt.Println("Phone Numbers:\t\t")
					for _, v := range grav.Entry[0].PhoneNumbers {
						fmt.Println("\t- ("+v.Type+"):", v.Value)
					}
					fmt.Println("Emails:\t\t")
					for _, v := range grav.Entry[0].Emails {
						fmt.Println("\t-", v.Value)
					}
					fmt.Println("Websites:\t\t")
					for _, v := range grav.Entry[0].Urls {
						fmt.Println("\t- ("+v.Title+"):", v.Value)
					}
				} else {
					fmt.Println(color.Colorize(color.Red, "[x]"), "The email", color.Colorize(color.Blue, args[0]), "is not linked to a Gravatar account!")
				}
			}
			fmt.Println()

			word, err := email.WordpressLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), "Wordpress:", err)
				fmt.Println()
			} else {
				if word {
					fmt.Println(color.Colorize(color.Green, "[+]"), "The email", color.Colorize(color.Blue, args[0]), "is linked to a Wordpress account!")
				} else {
					fmt.Println(color.Colorize(color.Red, "[x]"), "The email", color.Colorize(color.Blue, args[0]), "is not linked to a Wordpress account!")
				}
			}
			fmt.Println()

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
