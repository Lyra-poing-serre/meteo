package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// https://www.infoclimat.fr/opendata/?version=2&method=get&format=json&start=2025-09-16&end=2025-09-17&token=null
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("an error occured when loading .env file %s", err.Error())
	}

	apiToken := os.Getenv("API_TOKEN")
	fmt.Printf("Hello this your token: %s\n", apiToken)

	data, err := os.ReadFile("./example.json")
	if err != nil {
		log.Fatalf("an error occured when reading the file: %s", err.Error())
	}
	var raw map[string]json.RawMessage
	err = json.Unmarshal(data, &raw)
	if err != nil {
		fmt.Printf(
			"an error ocurred when parsing request message: %s",
			err.Error(),
		)
		return
	}

	var Response InfoClimatResponse
	Response.Meteo = make(map[string]Meteo)
	err = json.Unmarshal(raw["request_state"], &Response.RequestState)
	if err != nil {
		fmt.Printf("error while parsing request_state key: %s\n", err.Error())
		return
	}
	fmt.Println(Response.Meteo)
	delete(raw, "request_state")

	err = json.Unmarshal(raw["request_key"], &Response.RequestKey)
	if err != nil {
		fmt.Printf("error while parsing request_key key: %s\n", err.Error())
		return
	}
	delete(raw, "request_key")

	err = json.Unmarshal(raw["message"], &Response.Message)
	if err != nil {
		fmt.Printf("error while parsing message key: %s\n", err.Error())
		return
	}
	delete(raw, "message")

	err = json.Unmarshal(raw["source"], &Response.Source)
	if err != nil {
		fmt.Printf("error while parsing source key: %s\n", err.Error())
		return
	}
	delete(raw, "source")

	err = json.Unmarshal(raw["model_run"], &Response.ModelRun)
	if err != nil {
		fmt.Printf("error while parsing model_run key: %s\n", err.Error())
		return
	}
	delete(raw, "model_run")

	for k, v := range raw {
		var m Meteo
		if err := json.Unmarshal(v, &m); err != nil {
			fmt.Printf("error while parsing %s\n%s\n", k, err.Error())
			continue
		}
		Response.Meteo[k] = m
	}

	fmt.Printf("Response State: %d\n", Response.RequestState)
	fmt.Printf("Response Key: %s\n", Response.RequestKey)
	fmt.Printf("Response Message: %s\n", Response.Message)
	fmt.Printf("Response Source: %s\n", Response.Source)
	fmt.Printf("Response State: %s\n", Response.ModelRun)
	fmt.Println("Response Meteo:")
	// for k, v := range Response.Meteo {
	// 	fmt.Printf(". %s\n%v\n", k, v)
	// }
}
