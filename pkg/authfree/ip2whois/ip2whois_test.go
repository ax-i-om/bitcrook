package ip2whois

import (
	"encoding/json"
	"testing"

	"github.com/audioo/bitcrook/internal/config"
	"github.com/audioo/bitcrook/internal/http"
)

func TestDomainLookup(t *testing.T) {
	elConf, err := config.LoadConfig("../../../keyconfig.json")
	if err != nil {
		t.Error(err)
	}
	resp, err := http.GetReq("https://api.ip2whois.com/v2?key=" + elConf.Ip2WhoisKey + "&domain=" + "bitcrook.tech" + "&format=json")
	if err != nil {
		t.Error(err)
	}
	response := new(DomainData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
	}
	if response.Domain != "bitcrook.tech" {
		t.Error()
	}
}
