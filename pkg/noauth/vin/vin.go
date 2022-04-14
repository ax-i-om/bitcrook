package vin

import (
	"github.com/antchfx/htmlquery"
)

// VIN is the response type of the VinLookup() function, containing a plethora of information about the vehicle in question.
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

// VinLookup uses htmlquery to return a type *VIN containing list of information about a vehicle using its Vehicle Identification Number (VIN), as well as an error.
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
	if err != nil || b == nil {
		results.Make = "Error"
	} else {
		results.Make = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[3]/div[2]/text()")
	if err != nil || b == nil {
		results.Model = "Error"
	} else {
		results.Model = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[4]/div[2]/text()")
	if err != nil || b == nil {
		results.Year = "Error"
	} else {
		results.Year = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[5]/div[2]/text()")
	if err != nil || b == nil {
		results.Trim = "Error"
	} else {
		results.Trim = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[6]/div[2]/text()")
	if err != nil || b == nil {
		results.Body = "Error"
	} else {
		results.Body = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[7]/div[2]/text()")
	if err != nil || b == nil {
		results.Engine = "Error"
	} else {
		results.Engine = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[3]/div/div/div[1]/div[8]/div[2]/text()")
	if err != nil || b == nil {
		results.ManufacturedIn = "Error"
	} else {
		results.ManufacturedIn = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[1]/td[2]/text()")
	if err != nil || b == nil {
		results.TrimLevel = "Error"
	} else {
		results.TrimLevel = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[2]/td[2]/text()")
	if err != nil || b == nil {
		results.SteeringType = "Error"
	} else {
		results.SteeringType = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[3]/td[2]/text()")
	if err != nil || b == nil {
		results.Abs = "Error"
	} else {
		results.Abs = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[4]/td[2]/text()")
	if err != nil || b == nil {
		results.TankSize = "Error"
	} else {
		results.TankSize = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[5]/td[2]/text()")
	if err != nil || b == nil {
		results.OverallHeight = "Error"
	} else {
		results.OverallHeight = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[6]/td[2]/text()")
	if err != nil || b == nil {
		results.OverallLength = "Error"
	} else {
		results.OverallLength = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[7]/td[2]/text()")
	if err != nil || b == nil {
		results.OverallWidth = "Error"
	} else {
		results.OverallWidth = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[8]/td[2]/text()")
	if err != nil || b == nil {
		results.StandardSeating = "Error"
	} else {
		results.StandardSeating = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[9]/td[2]/text()")
	if err != nil || b == nil {
		results.OptionalSeating = "Error"
	} else {
		results.OptionalSeating = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[10]/td[2]/text()")
	if err != nil || b == nil {
		results.HighwayMileage = "Error"
	} else {
		results.HighwayMileage = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[11]/td[2]/text()")
	if err != nil || b == nil {
		results.CityMileage = "Error"
	} else {
		results.CityMileage = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[12]/td[2]/text()")
	if err != nil || b == nil {
		results.FuelType = "Error"
	} else {
		results.FuelType = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[13]/td[2]/text()")
	if err != nil || b == nil {
		results.DriveType = "Error"
	} else {
		results.DriveType = b.Data
	}
	b, err = htmlquery.Query(doc, "/html/body/div[1]/div/div/div[1]/div[6]/div/table[2]/tbody/tr[14]/td[2]/text()")
	if err != nil || b == nil {
		results.Transmission = "Error"
	} else {
		results.Transmission = b.Data
	}

	// Return the struct information
	return results, nil
}
