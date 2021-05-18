package vin

import (
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

// VIN is the response type of the VinLookup() function, containing a plethora of information about the vehicle in question.
type VIN struct {
	Vin               string
	Year              string
	Make              string
	Model             string
	Trim              string
	ShortTrim         string
	TrimVariations    string
	MadeIn            string
	MadeInCity        string
	VehicleStyle      string
	VehicleType       string
	VehicleSize       string
	VehicleCategory   string
	Doors             string
	FuelType          string
	FuelCapacity      string
	CityMileage       string
	HighwayMileage    string
	Engine            string
	EngineSize        string
	EngineCylinders   string
	TransmissionType  string
	TransmissionGears string
	DrivenWheels      string
	AntiBrakeSystem   string
	SteeringType      string
	CurbWeight        string
	GrossWeight       string
	OverallHeight     string
	OverallLength     string
	OverallWidth      string
	StandardSeating   string
	OptionalSeating   string
	InvoicePrice      string
	DeliveryCharges   string
	MSRP              string
}

// VinLookup uses htmlquery to return a type *VIN containing list of information about a vehicle using its Vehicle Identification Number (VIN), as well as an error.
func VinLookup(vin string) (*VIN, error) {
	doc, err := htmlquery.LoadURL("https://freevinsearch.net/vin_number_results.php?vin=" + vin)
	if err != nil {
		return nil, err
	}
	// Find all VIN details.
	count := 1
	r := new(VIN)
	for count < 37 {
		c := strconv.Itoa(count)
		b, err := htmlquery.Query(doc, "/html/body/div/div[2]/div/div[1]/div["+c+"]/text()")
		if err != nil {
			return nil, err
		}
		switch count {
		case 1:
			r.Vin = strings.ReplaceAll(b.Data, " : ", "")
		case 2:
			r.Year = strings.ReplaceAll(b.Data, " : ", "")
		case 3:
			r.Make = strings.ReplaceAll(b.Data, " : ", "")
		case 4:
			r.Model = strings.ReplaceAll(b.Data, " : ", "")
		case 5:
			r.Trim = strings.ReplaceAll(b.Data, " : ", "")
		case 6:
			r.ShortTrim = strings.ReplaceAll(b.Data, " : ", "")
		case 7:
			r.TrimVariations = strings.ReplaceAll(b.Data, " : ", "")
		case 8:
			r.MadeIn = strings.ReplaceAll(b.Data, " : ", "")
		case 9:
			r.MadeInCity = strings.ReplaceAll(b.Data, " : ", "")
		case 10:
			r.VehicleStyle = strings.ReplaceAll(b.Data, " : ", "")
		case 11:
			r.VehicleType = strings.ReplaceAll(b.Data, " : ", "")
		case 12:
			r.VehicleSize = strings.ReplaceAll(b.Data, " : ", "")
		case 13:
			r.VehicleCategory = strings.ReplaceAll(b.Data, " : ", "")
		case 14:
			r.Doors = strings.ReplaceAll(b.Data, " : ", "")
		case 15:
			r.FuelType = strings.ReplaceAll(b.Data, " : ", "")
		case 16:
			r.FuelCapacity = strings.ReplaceAll(b.Data, " : ", "")
		case 17:
			r.CityMileage = strings.ReplaceAll(b.Data, " : ", "")
		case 18:
			r.HighwayMileage = strings.ReplaceAll(b.Data, " : ", "")
		case 19:
			r.Engine = strings.ReplaceAll(b.Data, " : ", "")
		case 20:
			r.EngineSize = strings.ReplaceAll(b.Data, " : ", "")
		case 21:
			r.EngineCylinders = strings.ReplaceAll(b.Data, " : ", "")
		case 22:
			r.TransmissionType = strings.ReplaceAll(b.Data, " : ", "")
		case 23:
			r.TransmissionGears = strings.ReplaceAll(b.Data, " : ", "")
		case 24:
			r.DrivenWheels = strings.ReplaceAll(b.Data, " : ", "")
		case 25:
			r.AntiBrakeSystem = strings.ReplaceAll(b.Data, " : ", "")
		case 26:
			r.SteeringType = strings.ReplaceAll(b.Data, " : ", "")
		case 27:
			r.CurbWeight = strings.ReplaceAll(b.Data, " : ", "")
		case 28:
			r.GrossWeight = strings.ReplaceAll(b.Data, " : ", "")
		case 29:
			r.OverallHeight = strings.ReplaceAll(b.Data, " : ", "")
		case 30:
			r.OverallLength = strings.ReplaceAll(b.Data, " : ", "")
		case 31:
			r.OverallWidth = strings.ReplaceAll(b.Data, " : ", "")
		case 32:
			r.StandardSeating = strings.ReplaceAll(b.Data, " : ", "")
		case 33:
			r.OptionalSeating = strings.ReplaceAll(b.Data, " : ", "")
		case 34:
			r.InvoicePrice = strings.ReplaceAll(b.Data, " : ", "")
		case 35:
			r.DeliveryCharges = strings.ReplaceAll(b.Data, " : ", "")
		case 36:
			r.MSRP = strings.ReplaceAll(b.Data, " : ", "")
		}

		count = count + 1
	}
	return r, nil
}
