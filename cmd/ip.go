/*
Copyright © 2021 Audio <hyperaudio@protonmail.com>

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

// Package cmd ...
package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "IP Lookup",
	Long:  `Retrieve information on an IP Address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			cli.Clear()
			cli.Banner()
			cli.Dispban("IP Lookup")
			fmt.Println("")
			var resp = http.GetReq("http://ip-api.com/json/" + args[0] + "?fields=31162361")
			var s = new(ent.IPAddress)
			json.Unmarshal([]byte(resp), &s)
			fmt.Println("   Status: " + s.Status)
			fmt.Println("   Continent: " + s.Continent)
			fmt.Println("   Country: " + s.Country)
			fmt.Println("   Region Name: " + s.RegionName)
			fmt.Println("   City: " + s.City)
			fmt.Println("   District: " + s.District)
			fmt.Println("   Zip: " + s.Zip)
			fmt.Println("   Lat: " + strconv.FormatFloat(s.Lat, 'f', -1, 64))
			fmt.Println("   Lon: " + strconv.FormatFloat(s.Lon, 'f', -1, 64))
			fmt.Println("   Time Zone: " + s.Timezone)
			fmt.Println("   Currency: " + s.Currency)
			fmt.Println("   ISP: " + s.Isp)
			fmt.Println("   Org: " + s.Org)
			fmt.Println("   As: " + s.As)
			fmt.Println("   Asname: " + s.Asname)
			fmt.Println("   Reverse: " + s.Reverse)
			fmt.Println("   Mobile: " + strconv.FormatBool(s.Mobile))
			fmt.Println("   Proxy: " + strconv.FormatBool(s.Proxy))
			fmt.Println("   Hosting: " + strconv.FormatBool(s.Hosting))
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
