package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	const apiBaseURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies="

	fmt.Print("Enter currency (ex: usd, eur, brl): ")
	var currency string
	fmt.Scanln(&currency)
	apiURL := fmt.Sprintf("%s%s", apiBaseURL, strings.ToLower(currency))

	type PriceResponse struct {
		Bitcoin map[string]float64 `json:"bitcoin"`
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Fail make request:", err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Fail make request:", resp.StatusCode)
		panic(resp.StatusCode)
	}

	var processPrice PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&processPrice); err != nil {
		fmt.Println("Fail decode response:", err)
		panic(err)
	}
	price, ok := processPrice.Bitcoin[strings.ToLower(currency)]
	if !ok {
		fmt.Println("Fail get price")
		panic("Fail get price")
	}

	fmt.Printf("Bitcoin price in %s is: %.2f\n", strings.ToUpper(currency), price)
}
