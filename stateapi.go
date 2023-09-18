package main

import (
	"encoding/json"
	"net/http"
)

type StateTax struct {
	State string  `json:"state"`
	Rate  float64 `json:"rate"`
}

var stateTaxRates = []StateTax{
	{"Alabama", 4.0},
	{"Alaska", 0.0},
	{"Arizona", 5.6},
	{"Arkansas", 6.5},
	{"California", 7.25},
	{"Colorado", 2.9},
	{"Connecticut", 6.35},
	{"Delaware", 0.0},
	{"Florida", 6.0},
	{"Georgia", 4.0},
	{"Hawaii", 4.0},
	{"Idaho", 6.0},
	{"Illinois", 6.25},
	{"Indiana", 7.0},
	{"Iowa", 6.0},
	{"Kansas", 6.5},
	{"Kentucky", 6.0},
	{"Louisiana", 4.45},
	{"Maine", 5.5},
	{"Maryland", 6.0},
	{"Massachusetts", 6.25},
	{"Michigan", 6.0},
	{"Minnesota", 6.88},
	{"Mississippi", 7.0},
	{"Missouri", 4.225},
	{"Montana", 0.0},
	{"Nebraska", 5.5},
	{"Nevada", 6.85},
	{"New Hampshire", 0.0},
	{"New Jersey", 6.625},
	{"New Mexico", 5.125},
	{"New York", 4.0},
	{"North Carolina", 4.75},
	{"North Dakota", 5.0},
	{"Ohio", 5.75},
	{"Oklahoma", 4.5},
	{"Oregon", 0.0},
	{"Pennsylvania", 6.0},
	{"Rhode Island", 7.0},
	{"South Carolina", 6.0},
	{"South Dakota", 4.5},
	{"Tennessee", 7.0},
	{"Texas", 6.25},
	{"Utah", 4.7},
	{"Vermont", 6.0},
	{"Virginia", 4.3},
	{"Washington", 6.5},
	{"West Virginia", 6.0},
	{"Wisconsin", 5.0},
	{"Wyoming", 4.0},
}

func main() {
	http.HandleFunc("/tax-rates", getTaxRates)
	http.ListenAndServe(":8080", nil)
}

func getTaxRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Serve the static tax rates.
	taxRates := stateTaxRates

	// Encode the tax rates as indented JSON and send the response.
	prettyJSON, err := json.MarshalIndent(taxRates, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(prettyJSON)
}
