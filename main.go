package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
    "os"
)

const apiURL = "https://www.alphavantage.co/query?"

type GDPData struct {
    Date string `json:"date"`
    Value string `json:"value"`
}

type ApiResponse struct {
    Name     string    `json:"name"`
    Interval string    `json:"interval"`
    Unit     string    `json:"unit"`
    Data     []GDPData `json:"data"`
}

func main() {
	interval := "annual"
	fetchRealGDPData(interval)
}

func fetchRealGDPData(interval string) {
    apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	// build req url
	url := fmt.Sprintf("%sfunction=REAL_GDP&interval=%s&apikey=%s", apiURL, interval, apiKey)

	//get request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	//read the res body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	//parse json res
	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Error parsing json", err)
		return
	}

    // Example of printing out the fetched data
    fmt.Println("Name:", apiResponse.Name)
    fmt.Println("Interval:", apiResponse.Interval)
    fmt.Println("Unit:", apiResponse.Unit)
    for _, data := range apiResponse.Data {
        fmt.Printf("Date: %s, Value: %s\n", data.Date, data.Value)
    }

}
