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
package domain

import (
	"os"
	"strings"
	"testing"
)

func TestHIBPLookup(t *testing.T) {
	x, err := HIBPLookup("canva.com")
	if err != nil {
		t.Error(err)
	}
	if len(x) < 1 {
		t.Error()
	}
	if !strings.Contains(strings.ToLower(x[0].Name), "canva") {
		t.Error()
	}
}

func TestIPTLLookup(t *testing.T) {
	x, err := IPTLLookup(os.Getenv("BITCROOK_IPTL"), "canva.com")
	if err != nil {
		t.Error(err)
	}
	if x.Domain != "canva.com" {
		t.Error()
	}
}
