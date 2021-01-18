package main

import (
"encoding/json"
"fmt"
_ "github.com/joho/godotenv/autoload"
"log"
"net/http"
"os"
)

// Response types
type Response struct {
		Status int `json:"status"`
		Response string `json:"response"`
}

func main() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := Response{}
		res.Status = http.StatusOK
		res.Response = "Hello Golang"
		json.NewEncoder(w).Encode(res)
})

log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}