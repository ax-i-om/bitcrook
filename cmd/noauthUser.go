/*
Copyright © 2021 Bitcrook <bitcrook@protonmail.com>
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
	"strconv"

	"github.com/bitcrook/cybull/pkg/noauth/userlookup"
	"github.com/spf13/cobra"
)

// noauthUserCmd represents the noauthUser command
var noauthUserCmd = &cobra.Command{
	Use:   "noauthUser <username>",
	Short: "Returns websites that have accounts registered with the specified username. Auth: NONE",
	Long:  `Returns websites that have accounts registered with the specified username. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x := userlookup.UserLookup(args[0])
			for _, v := range x {
				fmt.Println("[" + strconv.FormatBool(v.Valid) + "] - " + v.Domain)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthUserCmd)
}
