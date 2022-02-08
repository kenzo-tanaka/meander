package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenzo-takana/meander"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	key := os.Getenv("GOOGLE_API_KEY")
	meander.APIKey = key

	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8081", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
