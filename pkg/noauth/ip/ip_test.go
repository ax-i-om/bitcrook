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

package ip

import (
	"encoding/json"
	"testing"

	"github.com/ax-i-om/bitcrook/internal/http"
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
