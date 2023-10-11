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

package vin

import (
	"strings"
	"testing"
)

func TestVFCLookup(t *testing.T) {
	x, err := VFCLookup("1C6RD7NT7CS293032")
	if err != nil {
		t.Error(err)
	}
	if !(strings.Contains(x.Model, "1500 Laramie")) {
		t.Error()
	}
}
