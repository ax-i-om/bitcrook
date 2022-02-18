package hibp

import (
	"encoding/json"
	"testing"

	"github.com/bitcrook/cybull/internal/config"
	"github.com/bitcrook/cybull/internal/http"
)

func TestGetBreaches(t *testing.T) {
	elConf, err := config.LoadConfig("../../../keyconfig.json")
	if err != nil {
		t.Error(err)
	}
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/breachedaccount/email@example.com", "hibp-api-key", elConf.HibpKey)
	if err != nil {
		t.Error(err)
	}
	response := new(BreachList)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBreachedSites(t *testing.T) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breaches")
	if err != nil {
		t.Error(err)
	}
	response := new(BreachedSites)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBreachedSite(t *testing.T) {
	body, err := http.GetReq("https://haveibeenpwned.com/api/v3/breach/wattpad")
	if err != nil {
		t.Error(err)
	}
	response := new(BreachedSite)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllPastes(t *testing.T) {
	elConf, err := config.LoadConfig("../../../keyconfig.json")
	if err != nil {
		t.Error(err)
	}
	body, err := http.AuthGet("https://haveibeenpwned.com/api/v3/pasteaccount/email@example.com", "hibp-api-key", elConf.HibpKey)
	if err != nil {
		t.Error()
	}
	response := new(AllPastes)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		t.Error()
	}
}
