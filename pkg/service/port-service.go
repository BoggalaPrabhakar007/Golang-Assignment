package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/repo"

	"github.com/gorilla/mux"
)

var repoPortInfo = make(map[string]models.Port)

// InsertPortData will read the data from the json file and insert the data in repo
func InsertPortData(w http.ResponseWriter, r *http.Request) {
	//read the port data from the file
	pData, err := ioutil.ReadFile(config.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	var portsInfo = make(map[string]models.Port)
	err2 := json.Unmarshal(pData, &portsInfo)
	if err2 != nil {
		log.Fatal(err2)
	}
	fPortInfo := getFormattedPortData(portsInfo)
	//repo layer call for inserting data into database
	err3 := repo.InsertPorts(context.Background(), fPortInfo)
	if err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err3)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetPortData gets the port data from repo
func GetPortData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//repo layer call for getting data from database
	portDetails, err := repo.GetPorts(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portDetails)

}

// GetPortDataByID gets the port data by id from repo
func GetPortDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars[constants.ID]
	//repo layer call for getting data from database
	portDetails, err := repo.GetPortByID(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portDetails)
}

// DeletePortByID delete the port data by id from repo
func DeletePortByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars[constants.ID]
	//repo layer call for getting data from database
	err := repo.DeletePortByID(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdatePortByID delete the port data by id from repo
func UpdatePortByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request models.PortDetails
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
		return
	}
	vars := mux.Vars(r)
	id := vars[constants.ID]
	//repo layer call for getting data from database
	err2 := repo.UpdatePortByID(context.Background(), id, &request)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err2)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//getFormattedPortData gets the port details in clean format
func getFormattedPortData(portsInfo map[string]models.Port) []models.PortDetails {
	pDetails := []models.PortDetails{}
	for k, v := range portsInfo {
		posrtDetails := models.PortDetails{
			ID: k,
			Port: models.Port{
				Name:        v.Name,
				City:        v.City,
				Country:     v.Country,
				Alias:       v.Alias,
				Regions:     v.Regions,
				Coordinates: v.Coordinates,
				Province:    v.Province,
				Timezone:    v.Timezone,
				Unlocs:      v.Unlocs,
				Code:        v.Code,
			},
		}
		pDetails = append(pDetails, posrtDetails)
	}
	return pDetails
}
