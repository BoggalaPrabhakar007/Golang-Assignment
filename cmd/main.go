package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/repo"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/service"
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

	//Load the config file
	config := config.LoadConfig(constants.ConfigPath)
	//initialize the repo library
	repoLib := repository.NewRepoLibServ()
	//initialize the repo layer
	repoServ := repo.NewPortRepoServ(repoLib)
	//initialize the port service library
	pServ := service.NewPortService(repoServ, config)
	//initialize the transport endpoints
	transport.InitTransport(r, pServ)

	//connecting to database
	_, err := repository.GetConnectionClient()
	if err != nil {
		log.Fatal(err)
	}

	//starting the http server
	fmt.Println(fmt.Sprintf(Msg, config.Server.Port))
	err = http.ListenAndServe(config.Server.Port, r)
	if err != nil {
		log.Fatal(err)
	}

}
