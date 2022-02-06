package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/kenzo-takana/meander"
)

func main() {
	// プログラムから使用できる最大にCPU数を指定できる
	// TODO: go 1.5からはこのコードは不要
	runtime.GOMAXPROCS(runtime.NumCPU())
	// meander.APIKey = "TODO"
	http.HandleFunc("/jouneys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
