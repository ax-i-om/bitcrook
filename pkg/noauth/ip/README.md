# [github.com/bitcrook/goseek/pkg/noauth/ip](https://github.com/bitcrook/goseek/tree/main/pkg/noauth/ip) - no authentication required


## Types

``` go
type IPAddress struct {
	Status     string  `json:"status"`
	Continent  string  `json:"continent"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionname"`
	City       string  `json:"city"`
	District   string  `json:"district"`
	Zip        string  `json:"zip"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Timezone   string  `json:"timezone"`
	Currency   string  `json:"currency"`
	Isp        string  `json:"isp"`
	Org        string  `json:"org"`
	As         string  `json:"as"`
	Asname     string  `json:"asname"`
	Reverse    string  `json:"reverse"`
	Mobile     bool    `json:"mobile"`
	Proxy      bool    `json:"proxy"`
	Hosting    bool    `json:"hosting"`
	Query      string  `json:"query"`
} 
```

## Functions


IPLookup takes an IPV4 Address (string) as a parameter and returns a plethora of information represented by type *IPAddress.
``` go
func IPLookup(ip string) (*IPAddress, error) {
	resp, err := http.GetReq("http://ip-api.com/json/" + ip + "?fields=31162361")
	if err != nil {
		return nil, err
	}
	s := new(IPAddress)
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
```

## Usage

Reading from the returned struct can be done like so:
``` go
package main

import (
	"fmt"
	"log"

	"github.com/bitcrook/goseek/pkg/noauth/ip"
)

func main() {
	res, err := ip.IPLookup("1.1.1.1")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res.Lat)
	fmt.Println(res.Lat)
}
```
Output from the above example:
```
-27.4766
-27.4766
```