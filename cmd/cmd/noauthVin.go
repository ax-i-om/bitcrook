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

	"github.com/maraudery/goseek/pkg/noauth/vin"
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
			x, err := vin.VinLookup(args[0])
			if err != nil {
				log.Panic(err)
			} else {
				fmt.Println("VIN: " + x.Vin)
				fmt.Println("Year: " + x.Year)
				fmt.Println("Make: " + x.Make)
				fmt.Println("Model: " + x.Model)
				fmt.Println("Trim: " + x.Trim)
				fmt.Println("Short Trim: " + x.ShortTrim)
				fmt.Println("Trim Variations: " + x.TrimVariations)
				fmt.Println("Made In: " + x.MadeIn)
				fmt.Println("Made In City: " + x.MadeInCity)
				fmt.Println("Vehicle Style: " + x.VehicleStyle)
				fmt.Println("Vehicle Type: " + x.VehicleType)
				fmt.Println("Vehicle Size: " + x.VehicleSize)
				fmt.Println("Vehicle Category: " + x.VehicleCategory)
				fmt.Println("Doors: " + x.Doors)
				fmt.Println("Fuel Type: " + x.FuelType)
				fmt.Println("Fuel Capacity: " + x.FuelCapacity)
				fmt.Println("City Mileage: " + x.CityMileage)
				fmt.Println("Highway Mileage: " + x.HighwayMileage)
				fmt.Println("Engine: " + x.Engine)
				fmt.Println("Engine Size: " + x.EngineSize)
				fmt.Println("Engine Cylinders: " + x.EngineCylinders)
				fmt.Println("Transmission Type: " + x.TransmissionType)
				fmt.Println("Transmission Gears: " + x.TransmissionGears)
				fmt.Println("Driven Wheels: " + x.DrivenWheels)
				fmt.Println("Anti-Brake System: " + x.AntiBrakeSystem)
				fmt.Println("Steering Type: " + x.SteeringType)
				fmt.Println("Curb Weight: " + x.CurbWeight)
				fmt.Println("Gross Weight: " + x.GrossWeight)
				fmt.Println("Overall Height: " + x.OverallHeight)
				fmt.Println("Overall Length: " + x.OverallLength)
				fmt.Println("Overall Width: " + x.OverallWidth)
				fmt.Println("Standard Seating: " + x.StandardSeating)
				fmt.Println("Optional Seating: " + x.OptionalSeating)
				fmt.Println("Invoice Price: " + x.InvoicePrice)
				fmt.Println("Delivery Charges: " + x.DeliveryCharges)
				fmt.Println("MSRP: " + x.MSRP)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthVinCmd)
}
