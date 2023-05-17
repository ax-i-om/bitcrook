/*
Copyright © 2021 ax-i-om <contactaudio@pm.me>
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

	"github.com/ax-i-om/bitcrook/pkg/noauth/vin"
	"github.com/spf13/cobra"
)

// noauthVinCmd represents the noauthVin command
var noauthVinCmd = &cobra.Command{
	Use:   "noauthVin <VIN>",
	Short: "Returns public information regarding a VIN. Auth: NONE",
	Long:  `Returns public information regarding a VIN. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			r, err := vin.VinLookup(args[0])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Make: " + r.Make)
			fmt.Println("Model: " + r.Model)
			fmt.Println("Year: " + r.Year)
			fmt.Println("Trim: " + r.Trim)
			fmt.Println("Style/Body: " + r.Body)
			fmt.Println("Engine: " + r.Engine)
			fmt.Println("Manufactured In: " + r.ManufacturedIn)
			fmt.Println("Trim Level: " + r.TrimLevel)
			fmt.Println("Steering Type: " + r.SteeringType)
			fmt.Println("Anti-Brake System: " + r.Abs)
			fmt.Println("Tank Size: " + r.TankSize)
			fmt.Println("Overall Height: " + r.OverallHeight)
			fmt.Println("Overall Length: " + r.OverallLength)
			fmt.Println("Overall Width: " + r.OverallWidth)
			fmt.Println("Standard Seating: " + r.StandardSeating)
			fmt.Println("Optional Seating: " + r.OptionalSeating)
			fmt.Println("Highway Mileage: " + r.HighwayMileage)
			fmt.Println("City Mileage: " + r.CityMileage)
			fmt.Println("Fuel Type: " + r.FuelType)
			fmt.Println("Drive Type: " + r.DriveType)
			fmt.Println("Transmission: " + r.Transmission)
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthVinCmd)
}
