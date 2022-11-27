/*
Copyright © 2021 audioo <contactaudio@pm.me>
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

	"github.com/audioo/bitcrook/pkg/noauth/tin"
	"github.com/spf13/cobra"
)

// noauthTinCmd represents the noauthTin command
var noauthTinCmd = &cobra.Command{
	Use:   "noauthTin <TIN | INN | OGRN | OGRNIP>",
	Short: "Returns public information regarding a Russian INN. Auth: NONE",
	Long:  `Returns public information regarding a Russian INN. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x, err := tin.TINLookup(args[0])
			if err != nil {
				log.Panic(err)
			} else {
				for i, v := range x.Rows {
					fmt.Println("--- Result #" + fmt.Sprint(i) + " ---")
					fmt.Println("Name: " + v.Name)
					fmt.Println("Assignment Date: " + v.AssignmentDate)
					fmt.Println("TerminationDate: " + v.TerminationDate)
					fmt.Println("OGRNIP: " + v.OGRNIP)
					fmt.Println("INN: " + v.INN)
					fmt.Println("Page: " + v.Page)
					fmt.Println("Total: " + v.Total)
					fmt.Println("Count: " + v.Count)
					fmt.Println("K: " + v.K)
					fmt.Println("T: " + v.T)
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(noauthTinCmd)
}
