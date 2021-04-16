package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/ent"
)

// Token ...
func Token(token string) {
	url := "https://discordapp.com/api/v6/users/@me"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	cli.Dispban("Discord Token Info")
	fmt.Println("")
	if res.StatusCode == 200 {
		var s = new(ent.TokenInfo)
		json.Unmarshal([]byte(string(body)), &s)
		fmt.Println("   ID: " + s.ID)
		fmt.Println("   Username: " + s.Username)
		fmt.Println("   Discriminator: " + s.Discriminator)
		fmt.Println("   MFA: " + strconv.FormatBool(s.MfaEnabled))
		fmt.Println("   Email: " + s.Email)
		fmt.Println("   Verified: " + strconv.FormatBool(s.Verified))
		fmt.Println("   Phone: " + s.Phone)
		fmt.Println("   Locale: " + s.Locale)
	} else {
		fmt.Println("   An error occurred. Please try again.")
	}
}
