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

	"github.com/TwiN/go-color"
	"github.com/ax-i-om/bitcrook/pkg/discord"
	"github.com/spf13/cobra"
)

// discordCmd represents the discord command
var discordCmd = &cobra.Command{
	Use:   "discord <token>",
	Short: "Information about a Discord authentication token",
	Long: `
The 'discord' command takes a discord authentication token as an 
argument and returns any associated information that is identified
via the discordapp API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "discordapp.com\n"))
			x, err := discord.TokenLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), err)
			}
			if x.ID == "" && x.Username == "" {
				fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
				return
			}
			fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))
			fmt.Println("ID:\t\t", x.ID)
			fmt.Println("Username:\t", x.Username)
			fmt.Println("Avatar:\t\t", x.Avatar)
			fmt.Println("Discriminator:\t", x.Discriminator)
			fmt.Println("Public Flags:\t", x.PublicFlags)
			fmt.Println("Flags:\t\t", x.Flags)
			fmt.Println("Banner:\t\t", x.Banner)
			fmt.Println("Accent Color:\t", x.AccentColor)
			fmt.Println("Global Name:\t", x.GlobalName)
			fmt.Println("Av. Decoration:\t", x.AvatarDecorationData)
			fmt.Println("Banner Color:\t", x.BannerColor)
			fmt.Println("MFA Enabled:\t", x.MfaEnabled)
			fmt.Println("Locale:\t\t", x.Locale)
			fmt.Println("Premium Type:\t", x.PremiumType)
			fmt.Println("Email:\t\t", x.Email)
			fmt.Println("Verified:\t", x.Verified)
			fmt.Println("Phone:\t\t", x.Phone)
			fmt.Println("NSFW Allowed:\t", x.NsfwAllowed)
			fmt.Println("Linked Users:\t", x.LinkedUsers)
			fmt.Println("Bought Flags:\t", x.PurchasedFlags)
			fmt.Println("Bio:\t\t", x.Bio)
			fmt.Println("Auth Types:\t", x.AuthenticatorTypes)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(discordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
