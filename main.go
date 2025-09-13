package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("an error occured when loading .env file %s", err.Error())
	}

	apiToken := os.Getenv("API_TOKEN")
	fmt.Printf("Hello this your token: %s\n", apiToken)

	defaultCity := strings.Split(os.Getenv("DEFAULT_CITY"), ",")
	if defaultCity[0] != "" {
		fmt.Printf("Il y a %d ville(s) par défault\n", len(defaultCity))
		for i, city := range defaultCity {
			fmt.Printf("%d. %s\n", i, city)
		}
	} else {
		fmt.Println("Warning: empty default value.")
	}
}
