package main

import (
	"fmt"
	"net/http"
	
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/handlers"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/ports", handlers.InsertPortData)
	http.HandleFunc("/getports", handlers.GetPortData)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
