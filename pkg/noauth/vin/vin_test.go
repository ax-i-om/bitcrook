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

package vin

import (
	"testing"

	"github.com/antchfx/htmlquery"
)

func TestVinLookup(t *testing.T) {
	// Load the webpage containing the information on the passed Vehicle Identification Number (VIN)
	doc, err := htmlquery.LoadURL("https://www.vinfreecheck.com/vin/" + "1C6RD7NT7CS293032" + "/vehicle-specification")
	// Check for error, return if error is not equal to nil
	if err != nil {
		t.Error(err)
	}
	// Create a new VIN struct
	results := new(VIN)

	// Pass VIN argument to VIN.Vin value
	results.Vin = "1C6RD7NT7CS293032"

	// Begin scraping values via XPath and assigning them to their corresponding place within the VIN struct
	b, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[2]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Make = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[3]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Model = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[4]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Year = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[5]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Trim = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[6]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Body = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[7]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Engine = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[8]/div[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.ManufacturedIn = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[1]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.TrimLevel = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[2]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.SteeringType = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[3]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Abs = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[4]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.TankSize = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[5]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.OverallHeight = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[6]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.OverallLength = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[7]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.OverallWidth = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[8]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.StandardSeating = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[9]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.OptionalSeating = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[10]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.HighwayMileage = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[11]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.CityMileage = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[12]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.FuelType = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[13]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.DriveType = b.Data
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[14]/td[2]/text()")
	if err != nil {
		t.Error(err)
	}
	results.Transmission = b.Data
	if err != nil {
		t.Error(err)
	}
}
