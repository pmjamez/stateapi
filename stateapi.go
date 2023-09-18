package main

import (
	"encoding/json"
	"net/http"
)

type StateTax struct {
	State string    `json:"state"`
	Rates []float64 `json:"rates"`
}

var stateTaxRates = []StateTax{
	{"Alabama", []float64{2.0, 4.0, 5.0}},
	{"Alaska", []float64{0.0}},
	{"Arizona", []float64{2.59, 3.34, 4.17, 4.5, 4.9, 5.9, 8.0}},
	{"Arkansas", []float64{2.0, 4.0, 5.0, 5.9, 6.6}},
	{"California", []float64{1.0, 2.0, 4.0, 6.0, 8.0, 9.3, 10.3, 11.3, 12.3}},
	{"Colorado", []float64{4.5}},
	{"Connecticut", []float64{3.0, 5.0, 5.5, 6.0, 6.5, 6.9}},
	{"Delaware", []float64{2.2, 3.9, 4.8, 5.2, 5.55, 5.95, 6.6, 7.05, 7.8}},
	{"Florida", []float64{4.0, 5.5, 6.0, 6.5, 7.0, 7.5}},
	{"Georgia", []float64{1.0, 2.0, 3.0, 4.0, 5.0, 5.75, 6.0}},
	{"Hawaii", []float64{1.4, 3.2, 5.5, 6.4, 6.8, 7.2, 7.6, 7.9, 8.25}},
	{"Idaho", []float64{1.125, 3.125, 3.25, 3.5, 4.0, 5.0, 6.0, 6.925}},
	{"Illinois", []float64{4.95}},
	{"Indiana", []float64{3.23, 3.33, 3.4, 3.23}},
	{"Iowa", []float64{0.33, 0.67, 2.25, 4.14, 5.63, 6.67}},
	{"Kansas", []float64{3.1, 5.25}},
	{"Kentucky", []float64{2.0, 3.0, 4.0, 5.0, 5.8, 6.0}},
	{"Louisiana", []float64{2.0, 4.0, 6.0}},
	{"Maine", []float64{5.8, 6.75}},
	{"Maryland", []float64{2.0, 3.0, 4.0, 4.75, 5.0, 5.25, 5.5, 5.75, 6.0}},
	{"Massachusetts", []float64{5.0}},
	{"Michigan", []float64{4.25}},
	{"Minnesota", []float64{5.35, 7.05, 7.85, 9.85}},
	{"Mississippi", []float64{3.0, 4.0, 5.0}},
	{"Missouri", []float64{1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 5.4}},
	{"Montana", []float64{1.0, 2.0, 3.0, 4.0, 5.4, 6.9}},
	{"Nebraska", []float64{2.46, 3.22, 3.57, 5.01, 6.84}},
	{"Nevada", []float64{0.0}},
	{"New Hampshire", []float64{0.0}},
	{"New Jersey", []float64{1.4, 1.75, 2.45, 3.5, 5.53, 6.37, 8.97}},
	{"New Mexico", []float64{1.7, 3.2, 4.7, 4.9, 5.9}},
	{"New York", []float64{4.0, 4.5, 5.25, 5.9, 6.33, 6.57, 6.85, 8.82, 9.65}},
	{"North Carolina", []float64{5.25}},
	{"North Dakota", []float64{1.1, 2.04, 2.27, 2.64, 2.9}},
	{"Ohio", []float64{0.5, 1.0, 2.0, 2.5, 3.0, 3.5, 4.5}},
	{"Oklahoma", []float64{0.5, 1.0, 2.0, 3.0, 4.0, 5.0, 5.25, 5.5, 5.65}},
	{"Oregon", []float64{5.0, 7.0, 9.0, 9.9, 10.8, 11.0, 11.25}},
	{"Pennsylvania", []float64{3.07}},
	{"Rhode Island", []float64{3.75, 4.75, 5.99, 5.9}},
	{"South Carolina", []float64{0.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
	{"South Dakota", []float64{0.0}},
	{"Tennessee", []float64{1.0, 2.0, 3.0, 4.0, 4.5, 5.0, 6.0}},
	{"Texas", []float64{0.0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0}},
	{"Utah", []float64{4.95}},
	{"Vermont", []float64{3.35, 6.6, 7.6, 8.75, 8.95}},
	{"Virginia", []float64{2.0, 3.0, 5.75}},
	{"Washington", []float64{0.0}},
	{"West Virginia", []float64{3.0, 4.0, 4.5, 6.0}},
	{"Wisconsin", []float64{3.54, 4.65, 6.27}},
	{"Wyoming", []float64{0.0}},
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
