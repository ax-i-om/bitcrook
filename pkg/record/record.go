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

// Package record contains the types and functions used for querying for
// information regarding public government and court records
package record

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/ax-i-om/bitcrook/internal/http"
)

/* ///////////////////////////////////////////////////////

   CASELAW

*/ ///////////////////////////////////////////////////////

// CaselawResponse is a type that represents the results of CaselawLookup().
type CaselawResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		ID               int    `json:"id"`
		URL              string `json:"url"`
		Name             string `json:"name"`
		NameAbbreviation string `json:"name_abbreviation"`
		DecisionDate     string `json:"decision_date"`
		DocketNumber     string `json:"docket_number"`
		FirstPage        string `json:"first_page"`
		LastPage         string `json:"last_page"`
		Citations        []struct {
			Type string `json:"type"`
			Cite string `json:"cite"`
		} `json:"citations"`
		Volume struct {
			URL          string `json:"url"`
			VolumeNumber string `json:"volume_number"`
			Barcode      string `json:"barcode"`
		} `json:"volume"`
		Reporter struct {
			URL      string `json:"url"`
			FullName string `json:"full_name"`
			ID       int    `json:"id"`
		} `json:"reporter"`
		Court struct {
			URL              string `json:"url"`
			NameAbbreviation string `json:"name_abbreviation"`
			Slug             string `json:"slug"`
			ID               int    `json:"id"`
			Name             string `json:"name"`
		} `json:"court"`
		Jurisdiction struct {
			ID          int    `json:"id"`
			NameLong    string `json:"name_long"`
			URL         string `json:"url"`
			Slug        string `json:"slug"`
			Whitelisted bool   `json:"whitelisted"`
			Name        string `json:"name"`
		} `json:"jurisdiction"`
		CitesTo []struct {
			Cite      string `json:"cite"`
			Category  string `json:"category"`
			Reporter  string `json:"reporter"`
			CaseIds   []int  `json:"case_ids"`
			OpinionID int    `json:"opinion_id"`
		} `json:"cites_to"`
		FrontendURL    string   `json:"frontend_url"`
		FrontendPdfURL string   `json:"frontend_pdf_url"`
		Preview        []string `json:"preview"`
		Analysis       struct {
			WordCount     int     `json:"word_count"`
			RandomBucket  int     `json:"random_bucket"`
			Sha256        string  `json:"sha256"`
			OcrConfidence float64 `json:"ocr_confidence"`
			CharCount     int     `json:"char_count"`
			RandomID      int64   `json:"random_id"`
			Pagerank      struct {
				Percentile float64 `json:"percentile"`
				Raw        float64 `json:"raw"`
			} `json:"pagerank"`
			Cardinality int    `json:"cardinality"`
			Simhash     string `json:"simhash"`
		} `json:"analysis"`
		LastUpdated time.Time `json:"last_updated"`
		Provenance  struct {
			DateAdded string `json:"date_added"`
			Batch     string `json:"batch"`
			Source    string `json:"source"`
		} `json:"provenance"`
	} `json:"results"`
}

// CaselawLookup takes a free form search and a response count as its parameters which are passed through the Caselaw API whose response is then represented by a *CaselawResponse type.
func CaselawLookup(search string, resCount int) (*CaselawResponse, error) {
	resp, err := http.GetReq("https://api.case.law/v1/cases/?page_size=" + fmt.Sprint(resCount) + "&search=" + url.PathEscape(search) + "&ordering=relevance")
	if err != nil {
		return nil, err
	}
	response := new(CaselawResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
