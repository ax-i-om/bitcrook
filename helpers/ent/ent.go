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

// TokenInfo is ...
type TokenInfo struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
	Flags         int    `json:"flags"`
	Locale        string `json:"locale"`
	NsfwAllowed   bool   `json:"nsfw_allowed"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	Email         string `json:"email"`
	Verified      bool   `json:"verified"`
	Phone         string `json:"phone"`
}
