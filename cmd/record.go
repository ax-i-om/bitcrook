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
	"github.com/ax-i-om/bitcrook/pkg/record"
	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record <query>",
	Short: "Search U.S. court data via Case.Law",
	Long: `
The 'record' command takes a free form query that supports conditional
operators as an argument and returns any associated information that is 
identified via Harvard's Case.Law API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "api.case.law\n"))

			x, err := record.CaselawLookup(args[0], 11)
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), err)
				return
			}
			if len(x.Results) < 1 {
				fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
				return
			}
			fmt.Println("STATUS:\t\t", color.Colorize(color.Green, "SUCCESS\n"))

			fmt.Println(color.Colorize(color.Blue, "[i]"), "Displaying", color.Colorize(color.Green, len(x.Results)), "results")
			for i, v := range x.Results {
				fmt.Println()
				fmt.Println(color.Colorize(color.Blue, "--- [ Entry #"+fmt.Sprint(i+1)+" ] ---"))
				fmt.Println("Name:\t\t", v.Name)
				fmt.Println("Decision Date:\t", v.DecisionDate)
				fmt.Println("Jurisdiction:\t", v.Jurisdiction.NameLong)
				fmt.Println("Court:\t\t", v.Court.Name)
				fmt.Println("Source:\t\t", v.Provenance.Source)
				fmt.Println("URL:\t\t", v.URL)
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
