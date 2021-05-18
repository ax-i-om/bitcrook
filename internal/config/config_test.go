package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	jsonFile, err := os.Open("../../keyconfig.json")
	if err != nil && len(err.Error()) > 0 {
		t.Error(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Error(err)
	}
	conf := new(Conf)
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		t.Error(err)
	}
}
