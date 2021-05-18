package caselaw

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// CLResponse is a type that represents the plethora of information returned by the GetCases() function.
type CLResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
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
			Cite    string `json:"cite"`
			CaseIds []int  `json:"case_ids"`
		} `json:"cites_to"`
		FrontendURL    string   `json:"frontend_url"`
		FrontendPdfURL string   `json:"frontend_pdf_url"`
		Preview        []string `json:"preview"`
		Analysis       struct {
			WordCount     int     `json:"word_count"`
			Sha256        string  `json:"sha256"`
			OcrConfidence float64 `json:"ocr_confidence"`
			CharCount     int     `json:"char_count"`
			Pagerank      struct {
				Percentile float64 `json:"percentile"`
				Raw        float64 `json:"raw"`
			} `json:"pagerank"`
			Cardinality int    `json:"cardinality"`
			Simhash     string `json:"simhash"`
		} `json:"analysis"`
		LastUpdated time.Time `json:"last_updated"`
	} `json:"results"`
}

// GetCases takes search (string) as its only parameter and returns a type of *CLResponse and an error.
func GetCases(search string) (*CLResponse, error) {
	url := "https://api.case.law/v1/cases/?page_size=10&search=" + search + "&ordering=relevance"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Allow", "GET, HEAD, OPTIONS")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Vary", "Accept")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	clres := new(CLResponse)
	err = json.Unmarshal([]byte(body), &clres)
	if err != nil {
		return nil, err
	}
	return clres, nil
}
