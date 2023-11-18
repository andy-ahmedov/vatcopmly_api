package vatcomply

import "fmt"

type Response struct {
	Date  string             `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type Geolocate struct {
	CountryCode string `json:"country_code"`
	IP          string `json:"ip"`
}

func (g Geolocate) Info() {
	fmt.Printf("Country Code: [%s]\n", g.CountryCode)
	fmt.Printf("IP:           [%s]\n", g.IP)
}

func (r Response) Info() {
	fmt.Println("DATE:\t", r.Date)
	fmt.Println("BASE:\t", r.Base)
	for key, value := range r.Rates {
		fmt.Printf("\t[%s]: %.2f\n", key, value)
	}
}
