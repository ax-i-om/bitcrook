/*
Copyright © 2021 Maraudery <maraudery@protonmail.com>

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
	"fmt"
	"log"
	"strconv"

	"github.com/maraudery/qualear/pkg/noauth/caselaw"
	"github.com/spf13/cobra"
)

// noauthCaselawCmd represents the noauthCaselaw command
var noauthCaselawCmd = &cobra.Command{
	Use:   "noauthCaselaw <\"Search Terms\">",
	Short: "Returns court case information regarding the value specified. Auth: NONE",
	Long:  `Returns court case information regarding the value specified. Auth: NONE`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			x, err := caselaw.GetCases(args[0])
			if err != nil {
				log.Panic(err)
			} else {
				fmt.Println("Count: " + strconv.Itoa(x.Count))
				for i, v := range x.Results {
					fmt.Println()
					fmt.Println("--------------------")
					fmt.Println("CASE NUMBER " + strconv.Itoa(i))
					fmt.Println()
					fmt.Println("ID: " + strconv.Itoa(v.ID))
					fmt.Println("URL: " + v.URL)
					fmt.Println("Name: " + v.Name)
					fmt.Println("Name Abbreviation: " + v.NameAbbreviation)
					fmt.Println("Decision Date: " + v.DecisionDate)
					fmt.Println("Docket Number: " + v.DocketNumber)
					fmt.Println("First Page: " + v.FirstPage)
					fmt.Println("Last Page: " + v.LastPage)
					fmt.Println("CITATIONS: ")
					for ia, w := range v.Citations {
						fmt.Println("CITATION NUMBER: " + strconv.Itoa(ia))
						fmt.Println("--- Type: " + w.Type)
						fmt.Println("--- Cite: " + w.Cite)
					}
					fmt.Println("VOLUME:")
					fmt.Println("---URL: " + v.Volume.URL)
					fmt.Println("---Volume Number: " + v.Volume.VolumeNumber)
					fmt.Println("---Barcode: " + v.Volume.Barcode)
					fmt.Println("REPORTER:")
					fmt.Println("---URL: " + v.Reporter.URL)
					fmt.Println("---Full Name: " + v.Reporter.FullName)
					fmt.Println("---ID: " + strconv.Itoa(v.Reporter.ID))
					fmt.Println("COURT:")
					fmt.Println("---URL: " + v.Court.URL)
					fmt.Println("---Name Abbreviation: " + v.Court.NameAbbreviation)
					fmt.Println("---Slug: " + v.Court.Slug)
					fmt.Println("---ID: " + strconv.Itoa(v.Court.ID))
					fmt.Println("---Name: " + v.Court.Name)
					fmt.Println("JURISDICTION:")
					fmt.Println("---ID: " + strconv.Itoa(v.Jurisdiction.ID))
					fmt.Println("---Name Long: " + v.Jurisdiction.NameLong)
					fmt.Println("---URL: " + v.Jurisdiction.URL)
					fmt.Println("---Slug: " + v.Jurisdiction.Slug)
					fmt.Println("---Whitelisted: " + strconv.FormatBool(v.Jurisdiction.Whitelisted))
					fmt.Println("---Name: " + v.Jurisdiction.Name)
					for ib, y := range v.CitesTo {
						fmt.Println("CITES TO: " + strconv.Itoa(ib))
						fmt.Println("--- Cite: " + y.Cite)
					}
					fmt.Println("Front-End URL: " + v.FrontendURL)
					fmt.Println("Front-End PDF URL: " + v.FrontendPdfURL)
					fmt.Println("PREVIEW: ")
					for _, s := range v.Preview {
						fmt.Println("--- " + s)
					}
					fmt.Println("ANALYSIS:")
					fmt.Println("---Word Count: " + strconv.Itoa(v.Analysis.WordCount))
					fmt.Println("---Sha256: " + v.Analysis.Sha256)
					fmt.Println("---OCR Confidence: " + strconv.FormatFloat(v.Analysis.OcrConfidence, 'f', 5, 64))
					fmt.Println("---Char Count: " + strconv.Itoa(v.Analysis.CharCount))
					fmt.Println("---PAGE RANK:")
					fmt.Println("------Percentile: " + strconv.FormatFloat(v.Analysis.Pagerank.Percentile, 'f', 5, 64))
					fmt.Println("------Raw: " + strconv.FormatFloat(v.Analysis.Pagerank.Raw, 'f', 5, 64))
					fmt.Println("---Cardinality: " + strconv.Itoa(v.Analysis.Cardinality))
					fmt.Println("---Simhash: " + v.Analysis.Simhash)
					fmt.Println("LAST UPDATED: " + v.LastUpdated.String())
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(noauthCaselawCmd)
}
