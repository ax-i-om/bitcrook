package vehicle

import (
	"fmt"
	"strconv"

	"github.com/antchfx/htmlquery"
)

func plateToVin(plate string, state string) string {
	doc, err := htmlquery.LoadURL("https://www.faxvin.com/license-plate-lookup/result?plate=" + plate + "&state=" + state)
	if err != nil {
		panic(err)
	}
	// Fetch VIN from page
	vin, err := htmlquery.Query(doc, "/html/body/div/div/div[2]/div/div[2]/div[2]/text()[2]")
	if err != nil {
		panic(err)
	}
	return vin.Data
}

// VinLookup ...
func VinLookup(vin string, option int) {
	if option == 1 {
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

}
