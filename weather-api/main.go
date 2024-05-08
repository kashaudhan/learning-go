package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type InnerMap map[string]string
type OuterMap map[string][]InnerMap

func main() {
	
	err := godotenv.Load()
  if err != nil {
		panic("failed to load env")
  }

	var apiKey = os.Getenv("WEATHER_API_KEY")
	var url = "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=bulk"

	placeName := os.Args[1:]

	if len(placeName) == 0 {
		fmt.Println("Please input a place name")
		return
	}
	
	place := strings.Join(placeName, " ")

	Get1(url, place)
	
}

func Get1(url string, place string) {
	body := make(OuterMap)
	body["locations"] = make([]InnerMap, 1)
	body["locations"][0] = make(InnerMap)
	body["locations"][0]["q"] = place
	body["locations"][0]["custom_id"] = "new-delhi"
	
	bodyBytes, err := json.Marshal(body)

	if err != nil {
		fmt.Println("Parsing error")
	}

	reqBody := bytes.NewBuffer(bodyBytes)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	
	defer resp.Body.Close()
	// Check the status code of the response
	if resp.StatusCode == http.StatusOK {
		
		// Unmarshal the JSON response into a Go-struct
		var result map[string]interface{}
			jsonData, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(jsonData, &result)
			if err != nil {
					fmt.Println("error: ",err)
					return
			}

			current, ok := result["bulk"].([]interface{})[0].(map[string]interface{})["query"].(map[string]interface{})["current"].(map[string]interface{})

			if !ok {
				fmt.Println("Please enter correct place name")
				return
			}

			fmt.Printf("Current temperature: %.1f°C & it's %s\n", current["temp_c"], current["condition"].(map[string]interface{})["text"])
	} else {
		// Handle the error response
		fmt.Println("Error:", resp.Status)
	}
}
