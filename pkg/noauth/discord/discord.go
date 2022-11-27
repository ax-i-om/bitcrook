package discord

import (
	"encoding/json"
	"log"

	"github.com/audioo/bitcrook/internal/http"
)

// DiscordTokenInfo is a represntation of a plethora of information returned by the TokenLookup() function.
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

// TokenLookup takes a Discord account token (string) as its only parameter and returns a type of *DiscordTokenInfo and an error.
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
