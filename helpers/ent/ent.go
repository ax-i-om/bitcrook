package ent

// Website is ...
type Website struct {
	Title  string
	Domain string
}

// WebsiteRes is ...
type WebsiteRes struct {
	Title  string
	Domain string
	Valid  bool
}

// Printer is ...
type Printer struct{}

// IPAddress is ...
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

// ScanResult is ...
type ScanResult struct {
	Port    string
	State   string
	Service string
}
