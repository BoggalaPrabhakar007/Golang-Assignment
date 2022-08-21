package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	repository "github.com/BoggalaPrabhakar007/golang-assignment/repository-lib/pkg/mongodb"
	transport "github.com/BoggalaPrabhakar007/golang-assignment/transport/http"

	"github.com/gorilla/mux"
)

const (
	//Msg is ued to display the info messages
	Msg string = "Staring application on port %s"
)

// main is the main function
func main() {
	//mux router
	r := mux.NewRouter()
	transport.InitTransport(r)

	//connecting to database
	_, err := repository.GetConnectionClient()
	if err != nil {
		log.Fatal(err)
	}

	//starting the http server
	fmt.Println(fmt.Sprintf(Msg, config.PortNumber))
	err = http.ListenAndServe(config.PortNumber, r)
	if err != nil {
		log.Fatal(err)
	}

}
