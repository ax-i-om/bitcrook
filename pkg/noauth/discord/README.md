# [github.com/audioo/bitcrook/pkg/noauth/discord](https://github.com/audioo/bitcrook/tree/main/pkg/noauth/discord) - no authentication required


## Types

Type DiscordTokenInfo represents a plethora of information returned by the TokenLookup() function.
``` go
type DiscordTokenInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Avatar         string `json:"avatar"`
	Discriminator  string `json:"discriminator"`
	PublicFlags    int    `json:"public_flags"`
	Flags          int    `json:"flags"`
	PurchasedFlags int    `json:"purchased_flags"`
	Locale         string `json:"locale"`
	NsfwAllowed    bool   `json:"nsfw_allowed"`
	MfaEnabled     bool   `json:"mfa_enabled"`
	Email          string `json:"email"`
	Verified       bool   `json:"verified"`
	Phone          string `json:"phone"`
}
```

## Functions


TokenLookup takes a Discord account token (string) as its only parameter and returns a type of *DiscordTokenInfo and an error.
``` go
func TokenLookup(token string) (*DiscordTokenInfo, error) {
	x, err := http.AuthGet("https://discordapp.com/api/v6/users/@me", "Authorization", token)
	if err != nil {
		log.Panic(err)
	}
	clres := new(DiscordTokenInfo)
	err = json.Unmarshal([]byte(x), &clres)
	if err != nil {
		return nil, err
	}
	return clres, nil
}
```

## Usage

An example of the TokenLookup() function in use:
``` go
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/audioo/bitcrook/pkg/noauth/caselaw"
)

func main() {
	x, err := discord.TokenLookup("REDACTED")
	if err != nil {
		log.Panic(err)
	}
    fmt.Println("ID: " + x.ID)
}
```
Part of the output from the above example:
```
ID: 738890436319510539
```