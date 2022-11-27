package tin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type postresp struct {
	T               string `json:"t"`
	CaptchaRequired bool   `json:"captchaRequired"`
}

type TinData struct {
	Rows []struct {
		AssignmentDate  string `json:"r,omitempty"`
		T               string `json:"t"`
		TerminationDate string `json:"e,omitempty"`
		Page            string `json:"pg,omitempty"`
		Total           string `json:"tot,omitempty"`
		Count           string `json:"cnt,omitempty"`
		INN             string `json:"i,omitempty"`
		K               string `json:"k"`
		Name            string `json:"n,omitempty"`
		OGRNIP          string `json:"o,omitempty"`
	} `json:"rows"`
}

// TINLookup ...
func TINLookup(tin string) (*TinData, error) {
	urlOne := "https://egrul.nalog.ru/"
	methodOne := "POST"

	payloadOne := strings.NewReader("query=" + tin)

	clientOne := &http.Client{}
	reqOne, err := http.NewRequest(methodOne, urlOne, payloadOne)

	if err != nil {
		return nil, err
	}
	reqOne.Header.Add("Content-Type", " application/x-www-form-urlencoded; charset=UTF-8")

	resOne, err := clientOne.Do(reqOne)
	if err != nil {
		return nil, err
	}
	defer resOne.Body.Close()

	bodyOne, err := ioutil.ReadAll(resOne.Body)
	if err != nil {
		return nil, http.ErrAbortHandler
	}
	s := new(postresp)
	err = json.Unmarshal([]byte(bodyOne), &s)
	if err != nil {
		return nil, err
	}

	urlTwo := "https://egrul.nalog.ru/search-result/" + s.T
	methodTwo := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(methodTwo, urlTwo, nil)

	if err != nil {
		return nil, err
	}
	resTwo, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resTwo.Body.Close()

	bodyTwo, err := ioutil.ReadAll(resTwo.Body)
	if err != nil {
		return nil, err
	}
	sTwo := new(TinData)
	err = json.Unmarshal([]byte(bodyTwo), &sTwo)
	if err != nil {
		return nil, err
	}

	return sTwo, nil
}
