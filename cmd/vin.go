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
	"github.com/ax-i-om/bitcrook/pkg/vin"
	"github.com/spf13/cobra"
)

// vinCmd represents the vin command
var vinCmd = &cobra.Command{
	Use:   "vin <VIN>",
	Short: "Information about a VIN",
	Long: `
The 'vin' command takes a Vehicle Identification Number (VIN) as an
argument and scrapes data regarding the corresponding vehicle from
the information provided by www.vinfreecheck.com.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			fmt.Println(color.Colorize(color.Blue, "[i]"), "Performing request to", color.Colorize(color.Green, "www.vinfreecheck.com\n"))

			r, err := vin.VFCLookup(args[0])
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[x]"), err)
				return
			}
			if r.Make == "Error" && r.Model == "Error" && r.Trim == "Error" {
				fmt.Println("STATUS:\t\t", color.Colorize(color.Red, "FAILURE\n"))
				return
			}
			fmt.Println("STATUS:\t", color.Colorize(color.Green, "SUCCESS\n"))
			fmt.Println("Make:\t\t", r.Make)
			fmt.Println("Model:\t\t", r.Model)
			fmt.Println("Year:\t\t", r.Year)
			fmt.Println("Trim:\t\t", r.Trim)
			fmt.Println("Style/Body:\t", r.Body)
			fmt.Println("Engine:\t\t", r.Engine)
			fmt.Println("Manufactured:\t", r.ManufacturedIn)
			fmt.Println("Trim Level:\t", r.TrimLevel)
			fmt.Println("Steering Type:\t", r.SteeringType)
			fmt.Println("ABS:\t\t", r.Abs)
			fmt.Println("Tank Size:\t", r.TankSize)
			fmt.Println("Total Height:\t", r.OverallHeight)
			fmt.Println("Total Length:\t", r.OverallLength)
			fmt.Println("Total Width:\t", r.OverallWidth)
			fmt.Println("Std. Seating:\t", r.StandardSeating)
			fmt.Println("Opt. Seating:\t", r.OptionalSeating)
			fmt.Println("Hwy. Mileage:\t", r.HighwayMileage)
			fmt.Println("City Mileage:\t", r.CityMileage)
			fmt.Println("Fuel Type:\t", r.FuelType)
			fmt.Println("Drive Type:\t", r.DriveType)
			fmt.Println("Transmission:\t", r.Transmission)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(vinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
