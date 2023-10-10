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

// Package discord contains the types and functions used for querying for
// information regarding a Discord authentication token
package discord

import (
	"encoding/json"

	"github.com/ax-i-om/bitcrook/internal/http"
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
		return nil, err
	}
	clres := new(DiscordTokenInfo)
	err = json.Unmarshal([]byte(x), &clres)
	if err != nil {
		return nil, err
	}
	return clres, nil
}
