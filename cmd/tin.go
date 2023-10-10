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
	"github.com/ax-i-om/bitcrook/pkg/noauth/tin"
	"github.com/spf13/cobra"
)

// tinCmd represents the tin command
var tinCmd = &cobra.Command{
	Use:   "tin <tin|inn|ogrn|ogrnip>",
	Short: "Information about a Russian TIN",
	Long: `
The 'tin' command takes a russian Tax Identification Number (TIN) as an
argument and returns any associated information that is identified via 
the egrul.nalog.ru API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "egrul.nalog.ru\n"))

			x, err := tin.TINLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), err)
				return
			}
			if len(x.Rows) < 1 {
				fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
				return
			}
			fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))

			fmt.Println(color.Colorize(color.Blue, "[i]"), "Discovered", color.Colorize(color.Green, len(x.Rows)), "results")
			count := 0
			for i, v := range x.Rows {
				if v.INN != "" && v.OGRNIP != "" {
					fmt.Println()
					fmt.Println(color.Colorize(color.Blue, "--- [ Entry #"+fmt.Sprint(i+1)+" ] ---"))
					fmt.Println("Name:\t\t", v.Name)
					fmt.Println("Assignment:\t", v.AssignmentDate)
					fmt.Println("Termination:\t", v.TerminationDate)
					fmt.Println("OGRNIP:\t\t", v.OGRNIP)
					fmt.Println("INN:\t\t", v.INN)
					fmt.Println("Page:\t\t", v.Page)
					fmt.Println("Total:\t\t", v.Total)
					fmt.Println("Count:\t\t", v.Count)
					fmt.Println("K:\t\t", v.K)
					fmt.Println("T:\t\t", v.T)
					continue
				}
				count++
			}
			fmt.Println()
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Omitted", color.Colorize(color.Green, count), "results")
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(tinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
