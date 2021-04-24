package email

import (
	"fmt"
	"strconv"

	"github.com/antchfx/htmlquery"
)

// Email returns information on an email. This will be updated soon.
func Email(email string) { // WORK IN PROGRESS
	doc, err := htmlquery.LoadURL("https://thatsthem.com/email/" + email)
	if err != nil {
		panic(err)
	}
	count := 1
	res := 0
	for {
		b, _ := htmlquery.Query(doc, "/html/body/div[2]/div/main/section[3]/div[2]/div["+strconv.Itoa(count)+"]")

		if b == nil {
			break
		}
		count++
		res++
	}
	if res == 0 {
		fmt.Println("ThatsThem yielded no results or you have been rate limited.")
	} else {
		fmt.Println("Got " + strconv.Itoa(res) + "+ results via thatsthem. Link: https://thatsthem.com/email/" + email)
	}
}
