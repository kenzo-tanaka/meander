package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kenzo-takana/meander"
)

func main() {
	// プログラムから使用できる最大にCPU数を指定できる
	// TODO: go 1.5からはこのコードは不要
	runtime.GOMAXPROCS(runtime.NumCPU())
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
	return json.NewEncoder(w).Encode(data)
}
