package vehicle

import (
	"fmt"
	"strconv"

	"github.com/antchfx/htmlquery"
)

// PlateToVin uses htmlquery to convert a United States license plate to a Vehicle Identification Number (VIN).
func PlateToVin(plate string, state string) string {
	doc, err := htmlquery.LoadURL("https://licenseplatelookup.io/lookup/" + plate + ":" + state)
	if err != nil {
		panic(err)
	}
	// Fetch VIN from page
	b, _ := htmlquery.Query(doc, "/html/body/div/div/main/div[1]/text()[2]")
	return b.Data
}

// VinLookup uses htmlquery to return a list of information about a vehicle using its Vehicle Identification Number (VIN).
// This will be updated soon.
func VinLookup(vin string) {

	doc, err := htmlquery.LoadURL("https://freevinsearch.net/vin_number_results.php?vin=" + vin)
	if err != nil {
		panic(err)
	}
	// Find all VIN details.
	count := 1
	for count < 37 {
		c := strconv.Itoa(count)
		b, _ := htmlquery.Query(doc, "/html/body/div/div[2]/div/div[1]/div["+c+"]/text()")
		n, _ := htmlquery.Query(doc, "//*[@id=\"Content\"]/div/div[1]/div["+c+"]/b")
		fmt.Println("   " + n.FirstChild.Data + b.Data)
		count = count + 1
	}

}
