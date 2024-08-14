package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseUrl = "https://newsapi.org"
const apiRoute = "/v2/everything"
const apiKey = ""
const queryParam = "Apple"

func main() {
	var params = url.Values{}
	params.Add("apiKey", apiKey)
	params.Add("q", queryParam)

	fullURL := fmt.Sprintf("%s?%s", baseUrl+apiRoute, params.Encode())

	fmt.Println("Full url: ", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Make the request using the default HTTP client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("\nGot response with status code: ", resp.StatusCode, "; Now trying to read the response.....")

	bytesData, parsingError := io.ReadAll(resp.Body)

	if parsingError != nil {
		fmt.Println("Error during parsing: ", parsingError)
	}

	fmt.Println("\nConverting bytes data to String...")

	strData := string(bytesData)

	fmt.Println("\nString data is : ", strData)

	fmt.Println("\nConverting it to JSON data...")

	jsonData := make(map[string]interface{})

	fmt.Printf("Type of jsonData: %T\n", jsonData)

	jsonParsingError := json.Unmarshal(bytesData, &jsonData)

	if jsonParsingError != nil {
		fmt.Println("Json parsing error : ", jsonParsingError)
		return
	}

	fmt.Println("Parsed JSON: ", jsonData)

}
