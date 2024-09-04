package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	const apiURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

	type PriceResponse struct {
		Bitcoin struct {
			USD float64 `json:"usd"`
		} `json:"bitcoin"`
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

	fmt.Printf("Bitcoin price: $%.2f\n", processPrice.Bitcoin.USD)
}
