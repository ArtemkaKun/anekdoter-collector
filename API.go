package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/joke", GetJoke).Methods("GET")
}

func GetJoke(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetRandomJoke())
	if err != nil {
		EncodingJSONError(err)
	}
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while encoding JSON: %v\n", err))
}
