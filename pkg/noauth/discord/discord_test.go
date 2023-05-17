package discord

import (
	"encoding/json"
	"testing"

	"github.com/ax-i-om/bitcrook/internal/http"
)

func TestTokenLookup(t *testing.T) {
	x, err := http.AuthGet("https://discordapp.com/api/v6/users/@me", "Authorization", "awdawd")
	if err != nil {
		t.Error(err)
	}
	clres := new(DiscordTokenInfo)
	err = json.Unmarshal([]byte(x), &clres)
	if err != nil {
		t.Error(err)
	}
}
