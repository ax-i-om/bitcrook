package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Conf is the type used for testing, which primarily contains API keys.
type Conf struct {
	MelissaKey string `json:"melissaKeyCred"`
	HibpKey    string `json:"hibpKey"`
	DataGovKey string `json:"dataGovKey"`
}

// LoadConfig returns a type *Conf by reading and unmarshalling the JSON file found at the passed path.
func LoadConfig(path string) (*Conf, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(Conf)
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
