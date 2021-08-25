package ip

import (
	"encoding/json"
	"testing"

	"github.com/maraudery/ra/internal/http"
)

func TestIPLookup(t *testing.T) {
	resp, err := http.GetReq("http://ip-api.com/json/1.1.1.1?fields=31162361")
	if err != nil {
		t.Error(err)
	}
	s := new(IPAddress)
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		t.Error(err)
	}
	if len(s.Reverse) < 1 {
		t.Fail()
	}

}
