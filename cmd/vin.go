/*
Copyright © 2021 Audio <hyperaudio@protonmail.com>

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

	"github.com/audioo/goseek/internal/cli"
	"github.com/audioo/goseek/pkg/vehicle"
	"github.com/spf13/cobra"
)

// vinCmd represents the vin command
var vinCmd = &cobra.Command{
	Use:   "vin [vin]",
	Short: "Return vehicle information from VIN",
	Long:  `Return vehicle information using its Vehicle Identification Number`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cli.Banner()
			cmd.Usage()
		} else {
			cli.Banner()
			cli.Dispban("VIN Lookup")
			fmt.Println()
			vehicle.VinLookup(args[0])
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(vinCmd)
}
