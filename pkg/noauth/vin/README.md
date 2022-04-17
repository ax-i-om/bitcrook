# [github.com/audioo/bitcrook/noauth/vin](https://github.com/audioo/bitcrook/tree/main/noauth/vin) - no authentication required


## Types

VIN is the response type of the VinLookup() function, containing a plethora of information about the vehicle in question.
``` go 
type VIN struct {
	Vin             string
	Make            string
	Model           string
	Year            string
	Trim            string
	Body            string
	Engine          string
	ManufacturedIn  string
	TrimLevel       string
	SteeringType    string
	Abs             string
	TankSize        string
	OverallHeight   string
	OverallLength   string
	OverallWidth    string
	StandardSeating string
	OptionalSeating string
	HighwayMileage  string
	CityMileage     string
	FuelType        string
	DriveType       string
	Transmission    string
}
```

## Functions


VinLookup uses htmlquery to return a type *VIN containing list of information about a vehicle using its Vehicle Identification Number (VIN), as well as an error.
``` go
func VinLookup(vin string) (*VIN, error) {
	// Load the webpage containing the information on the passed Vehicle Identification Number (VIN)
	doc, err := htmlquery.LoadURL("https://www.vinfreecheck.com/vin/" + vin + "/vehicle-specification")
	// Check for error, return if error is not equal to nil
	if err != nil {
		return nil, err
	}
	// Create a new VIN struct
	results := new(VIN)

	// Pass VIN argument to VIN.Vin value
	results.Vin = vin

	// Begin scraping values via XPath and assigning them to their corresponding place within the VIN struct
	b, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[2]/div[2]/text()")
	results.Make = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[3]/div[2]/text()")
	results.Model = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[4]/div[2]/text()")
	results.Year = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[5]/div[2]/text()")
	results.Trim = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[6]/div[2]/text()")
	results.Body = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[7]/div[2]/text()")
	results.Engine = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[8]/div[2]/text()")
	results.ManufacturedIn = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[1]/td[2]/text()")
	results.TrimLevel = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[2]/td[2]/text()")
	results.SteeringType = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[3]/td[2]/text()")
	results.Abs = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[4]/td[2]/text()")
	results.TankSize = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[5]/td[2]/text()")
	results.OverallHeight = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[6]/td[2]/text()")
	results.OverallLength = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[7]/td[2]/text()")
	results.OverallWidth = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[8]/td[2]/text()")
	results.StandardSeating = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[9]/td[2]/text()")
	results.OptionalSeating = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[10]/td[2]/text()")
	results.HighwayMileage = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[11]/td[2]/text()")
	results.CityMileage = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[12]/td[2]/text()")
	results.FuelType = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[13]/td[2]/text()")
	results.DriveType = check(b, err)
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[14]/td[2]/text()")
	results.Transmission = check(b, err)

	// Return the struct information
	return results, nil
}
```

## Usage

The VinLookup() function in use:
``` go
package main

import (
	"fmt"
	"log"

	"github.com/audioo/bitcrook/noauth/vin"
)

func main() {
	x, err := vin.VinLookup("11VG815E5GA000037")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x.Make)
}
```

Output from the above example: 
```
KALMAR INDUSTRIES LLC
```