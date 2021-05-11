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

	"github.com/maraudery/goseek/internal/cli"
	"github.com/maraudery/goseek/internal/http"
	"github.com/maraudery/goseek/pkg/vehicle"
	"github.com/spf13/cobra"
)

// licenseplateCmd represents the licenseplate command
var licenseplateCmd = &cobra.Command{
	Use:   "licenseplate [plate (no dashes or spaces)] [state abbreviated]",
	Short: "Return vehicle information",
	Long:  `Return vehicle information using license plate and state`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cli.Banner()
			cmd.Usage()
		} else {
			if http.Connected() {
				cli.Banner()
				cli.Dispban("License Plate Lookup")
				fmt.Println()
				vehicle.VinLookup(vehicle.PlateToVin(args[0], args[1]))
				fmt.Println()
			} else {
				cli.Banner()
				cmd.Usage()
				fmt.Println()
				fmt.Println("An internet connection is required to run this command...")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(licenseplateCmd)
}
