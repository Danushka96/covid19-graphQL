package main

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Record struct {
	FIPS          string `json:"fips"`
	Admin2        string `json:"admin_2"`
	ProvinceState string `json:"province_state"`
	CountryRegion string `json:"country_region"`
	LastUpdate    string `json:"last_update"`
	Lat           string `json:"lat"`
	Long_         string `json:"long_"`
	Confirmed     string `json:"confirmed"`
	Deaths        string `json:"deaths"`
	Recovered     string `json:"recovered"`
	Active        string `json:"active"`
	CombinedKey   string `json:"combined_key"`
}
