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
	"strconv"

	"github.com/maraudery/ra/pkg/noauth/discord"
	"github.com/spf13/cobra"
)

// noauthDiscordTokenCmd represents the noauthDiscordToken command
var noauthDiscordTokenCmd = &cobra.Command{
	Use:   "noauthDiscordToken <token>",
	Short: "Returns information regarding the passed token. Auth: NONE",
	Long:  `Returns information regarding the passed token. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x, err := discord.TokenLookup(args[0])
			if err != nil {
				log.Panic(err)
			} else {
				fmt.Println("ID: " + x.ID)
				fmt.Println("Username: " + x.Username)
				fmt.Println("Avatar: " + x.Avatar)
				fmt.Println("Discriminator: " + x.Discriminator)
				fmt.Println("Public Flags: " + strconv.Itoa(x.PublicFlags))
				fmt.Println("Flags: " + strconv.Itoa(x.Flags))
				fmt.Println("Purchased Flags: " + strconv.Itoa(x.PurchasedFlags))
				fmt.Println("Locale: " + x.Locale)
				fmt.Println("NSFW Allowed: " + strconv.FormatBool(x.NsfwAllowed))
				fmt.Println("MFA Enabled: " + strconv.FormatBool(x.MfaEnabled))
				fmt.Println("Email: " + x.Email)
				fmt.Println("Verified: " + strconv.FormatBool(x.Verified))
				fmt.Println("Phone: " + x.Phone)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthDiscordTokenCmd)
}
