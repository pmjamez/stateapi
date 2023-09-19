package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type StateTax struct {
	State string    `json:"state"`
	Rates []float64 `json:"rates"`
	// Amounts []float64 `json:"amounts"`
	Incomes []float64 `json:"incomes"`
}

var stateTaxRates = []StateTax{
	{"Alabama", []float64{0.0200, 0.0400, 0.0500}, []float64{1000.0000, 5000.0000, 6000.0000}},
	{"Alaska", []float64{0.0000}, nil},
	{"Arizona", []float64{0.0259, 0.0334, 0.0417, 0.0450, 0.0490, 0.0590, 0.0800}, []float64{10318.0000, 25872.0000, 51744.0000, 155232.0000, 517440.0000, 1552320.0000}},
	{"Arkansas", []float64{0.0200, 0.0400, 0.0500, 0.0590, 0.0660}, []float64{4299.0000, 8499.0000, 12699.0000, 21165.0000}},
	{"California", []float64{0.0100, 0.0200, 0.0400, 0.0600, 0.0800, 0.0930, 0.1030, 0.1130, 0.1230}, []float64{8542.0000, 20255.0000, 31969.0000, 44377.0000, 56085.0000, 286492.0000, 343791.0000}},
	{"Colorado", []float64{0.0450}, nil},
	{"Connecticut", []float64{0.0300, 0.0500, 0.0550, 0.0600, 0.0650, 0.0690}, []float64{10000.0000, 50000.0000, 100000.0000, 200000.0000}},
	{"Delaware", []float64{0.0220, 0.0390, 0.0480, 0.0520, 0.0555, 0.0595, 0.0660, 0.0705, 0.0780}, []float64{2000.0000, 5000.0000, 10000.0000, 20000.0000, 25000.0000, 60000.0000}},
	{"Florida", []float64{0.0400, 0.0550, 0.0600, 0.0650, 0.0700, 0.0750}, []float64{5000.0000, 25000.0000, 50000.0000, 75000.0000, 100000.0000}},
	{"Georgia", []float64{0.0100, 0.0200, 0.0300, 0.0400, 0.0500, 0.0575, 0.0600}, []float64{750.0000, 2250.0000, 3750.0000, 5250.0000, 7000.0000}},
	{"Hawaii", []float64{0.0140, 0.0320, 0.0550, 0.0640, 0.0680, 0.0720, 0.0760, 0.0790, 0.0825}, []float64{2400.0000, 4800.0000, 9600.0000, 14400.0000, 19200.0000}},
	{"Idaho", []float64{0.01125, 0.03125, 0.03250, 0.03500, 0.04000, 0.05000, 0.06000, 0.06925}, []float64{1550.0000, 3100.0000, 4650.0000, 6200.0000, 7750.0000}},
	{"Illinois", []float64{0.0495}, nil},
	{"Indiana", []float64{0.0323, 0.0333, 0.0340, 0.0323}, []float64{10000.0000, 25000.0000, 50000.0000}},
	{"Iowa", []float64{0.0033, 0.0067, 0.0225, 0.0414, 0.0563, 0.0667}, []float64{1000.0000, 5000.0000, 10000.0000, 20000.0000}},
	{"Kansas", []float64{0.0310, 0.0525}, []float64{15000.0000, 30000.0000}},
	{"Kentucky", []float64{0.0200, 0.0300, 0.0400, 0.0500, 0.0580, 0.0600}, []float64{3000.0000, 4000.0000, 5000.0000, 5800.0000, 6000.0000}},
	{"Louisiana", []float64{0.0200, 0.0400, 0.0600}, []float64{12500.0000, 50000.0000}},
	{"Maine", []float64{0.0580, 0.0675}, []float64{21050.0000}},
	{"Maryland", []float64{0.0200, 0.0300, 0.0400, 0.0475, 0.0500, 0.0525, 0.0550, 0.0575, 0.0600}, []float64{1000.0000, 2000.0000, 3000.0000, 4000.0000, 5000.0000}},
	{"Massachusetts", []float64{0.0500}, nil},
	{"Michigan", []float64{0.0425}, nil},
	{"Minnesota", []float64{0.0535, 0.0705, 0.0785, 0.0985}, []float64{27090.0000, 180251.0000, 280251.0000}},
	{"Mississippi", []float64{0.0300, 0.0400, 0.0500}, []float64{5000.0000, 10000.0000}},
	{"Missouri", []float64{0.0150, 0.0200, 0.0250, 0.0300, 0.0350, 0.0400, 0.0540}, []float64{1008.0000, 2108.0000, 4216.0000, 8429.0000, 10000.0000}},
	{"Montana", []float64{0.0100, 0.0200, 0.0300, 0.0400, 0.0540, 0.0690}, []float64{2900.0000, 5800.0000, 8700.0000, 11599.0000}},
	{"Nebraska", []float64{0.0246, 0.0322, 0.0357, 0.0501, 0.0684}, []float64{3200.0000, 16400.0000, 29750.0000, 31900.0000}},
	{"Nevada", []float64{0.0000}, nil},
	{"New Hampshire", []float64{0.0000}, nil},
	{"New Jersey", []float64{0.0140, 0.0175, 0.0245, 0.0350, 0.0553, 0.0637, 0.0897}, []float64{20000.0000, 35000.0000, 40000.0000, 75000.0000, 500000.0000}},
	{"New Mexico", []float64{0.0170, 0.0320, 0.0470, 0.0490, 0.0590}, []float64{5500.0000, 11000.0000, 16000.0000}},
	{"New York", []float64{0.0400, 0.0450, 0.0525, 0.0590, 0.0633, 0.0657, 0.0685, 0.0882, 0.0965}, []float64{8500.0000, 11700.0000, 13900.0000, 21400.0000, 80650.0000, 215400.0000, 1077550.0000}},
	{"North Carolina", []float64{0.0525}, nil},
	{"North Dakota", []float64{0.0110, 0.0204, 0.0227, 0.0264, 0.0290}, []float64{3950.0000, 9775.0000, 39475.0000}},
	{"Ohio", []float64{0.0050, 0.0100, 0.0200, 0.0250, 0.0300, 0.0350, 0.0450}, []float64{10750.0000, 21500.0000, 43000.0000}},
	{"Oklahoma", []float64{0.0050, 0.0100, 0.0200, 0.0300, 0.0400, 0.0500, 0.0525, 0.0550, 0.0565}, []float64{1000.0000, 5000.0000, 10000.0000, 20000.0000, 25000.0000, 50000.0000}},
	{"Oregon", []float64{0.0500, 0.0700, 0.0900, 0.0990, 0.1080, 0.1100, 0.1125}, []float64{3400.0000, 15000.0000, 20000.0000, 125000.0000}},
	{"Pennsylvania", []float64{0.0307}, nil},
	{"Rhode Island", []float64{0.0375, 0.0475, 0.0599, 0.0590}, []float64{6150.0000, 15750.0000, 67450.0000}},
	{"South Carolina", []float64{0.0300, 0.0400, 0.0500, 0.0600, 0.0700}, []float64{3000.0000, 6000.0000, 12000.0000, 17000.0000}},
	{"South Dakota", []float64{0.0000}, nil},
	{"Tennessee", []float64{0.0100, 0.0200, 0.0300, 0.0400, 0.0450, 0.0500, 0.0600}, []float64{10000.0000, 20000.0000, 30000.0000, 45000.0000}},
	{"Texas", []float64{0.0000, 0.0050, 0.0100, 0.0150, 0.0200, 0.0250, 0.0300}, []float64{5000.0000, 10000.0000, 15000.0000, 20000.0000}},
	{"Utah", []float64{0.0495}, nil},
	{"Vermont", []float64{0.0335, 0.0660, 0.0760, 0.0875, 0.0895}, []float64{39150.0000, 195750.0000, 260250.0000}},
	{"Virginia", []float64{0.0200, 0.0300, 0.0575}, []float64{3000.0000, 5000.0000, 17000.0000}},
	{"Washington", []float64{0.0000}, nil},
	{"West Virginia", []float64{0.0300, 0.0400, 0.0450, 0.0600}, []float64{10000.0000, 25000.0000, 40000.0000}},
	{"Wisconsin", []float64{0.0354, 0.0465, 0.0627}, []float64{22410.0000, 117950.0000}},
	{"Wyoming", []float64{0.0000}, nil},
}

func getTaxRates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stateTaxRates)
}

// func getTaxRatesByState(c *gin.Context) {
// 	// Get the state name from the request URL
// 	state := c.Param("state")

// 	// Find the tax rates for the specified state
// 	var taxRates []float64
// 	for _, tax := range stateTaxRates {
// 		if tax.State == state {
// 			taxRates = tax.Rates
// 			break
// 		}
// 	}

// 	// Check if tax rates were found
// 	if len(taxRates) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
// 		return
// 	}

//		// Return the tax rates in indented JSON format
//		c.IndentedJSON(http.StatusOK, gin.H{"state": state, "rates": taxRates})
//	}
func calculateStateTax(income float64, rates []float64, incomes []float64) float64 {
	tax := 0.0
	for i := 0; i < len(incomes); i++ {
		if income <= incomes[i] {
			tax += income * rates[i]
			break
		} else {
			tax += incomes[i] * rates[i]
			income -= incomes[i]
		}
	}
	return tax
}

func getStateTax(c *gin.Context) {
	state := c.Param("state")
	incomeStr := c.DefaultQuery("income", "0")

	income, err := strconv.ParseFloat(incomeStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid income amount"})
		return
	}

	var stateTax StateTax
	for _, st := range stateTaxRates {
		if st.State == state {
			stateTax = st
			break
		}
	}

	if stateTax.State == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State not found"})
		return
	}

	tax := calculateStateTax(income, stateTax.Rates, stateTax.Incomes)

	c.JSON(http.StatusOK, gin.H{
		"state_tax": tax,
	})
}

func main() {
	router := gin.Default()
	router.GET("/taxrates", getTaxRates)
	router.GET("/taxrates/:state", getStateTax)
	router.Run("localhost:8080")
}
