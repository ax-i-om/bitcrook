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
