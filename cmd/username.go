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

	"github.com/ax-i-om/bitcrook/pkg/noauth/userlookup"
	"github.com/spf13/cobra"
)

// usernameCmd represents the username command
var usernameCmd = &cobra.Command{
	Use:   "username <username>",
	Short: "Returns websites that have accounts registered with the specified username. Auth: NONE",
	Long:  `Returns websites that have accounts registered with the specified username. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			userlookup.CliSearch(args[0])
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(usernameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usernameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usernameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
