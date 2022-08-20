package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/service"
	repository "github.com/BoggalaPrabhakar007/golang-assignment/repository-lib/pkg/mongodb"

	"github.com/gorilla/mux"
)

const (
	//Msg is ued to display the info messages
	Msg string = "Staring application on port %s"
	//RegisterPortsURI is used to register the ports
	RegisterPortsURI = "/ports"
	//GetPortsURI is used to get the ports
	GetPortsURI = "/ports"
	//GetPortByIDURI gets the port by id
	GetPortByIDURI = "/port/{id}"
	//UpdatePortURI updates the port for given port id
	UpdatePortURI = "/port/{id}"
	//DeletePortIRI deletes the port for given port id
	DeletePortIRI = "/port/{id}"
)

// main is the main function
func main() {
	r := mux.NewRouter()
	r.HandleFunc(RegisterPortsURI, service.InsertPortData).Methods(http.MethodPost)
	r.HandleFunc(GetPortsURI, service.GetPortData).Methods(http.MethodGet)
	r.HandleFunc(GetPortByIDURI, service.GetPortDataByID).Methods(http.MethodGet)
	r.HandleFunc(UpdatePortURI, service.UpdatePortByID).Methods(http.MethodPut)
	r.HandleFunc(DeletePortIRI, service.DeletePortByID).Methods(http.MethodDelete)
	//starting the http server
	fmt.Println(fmt.Sprintf(Msg, config.PortNumber))
	err := http.ListenAndServe(config.PortNumber, r)
	if err != nil {
		log.Fatal(err)
	}
	//connecting to database
	_, err = repository.GetConnectionClient()
	if err != nil {
		log.Fatal(err)
	}
}
