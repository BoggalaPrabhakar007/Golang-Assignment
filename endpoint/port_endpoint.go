package endpoint

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/service"

	"github.com/gorilla/mux"
)

// InsertPortDataEndPoint endpoint for insert data
func InsertPortDataEndPoint(w http.ResponseWriter, r *http.Request) {
	var portsInfo = make(map[string]models.Port)
	// if u want to read the data from http request uncomment the below lines
	/*err := json.NewDecoder(r.Body).Decode(&portsInfo)
	if err != nil {
		log.Fatal(err)
	}*/
	pServ := service.PortServ{}
	portService := service.PortService(pServ)
	err := portService.InsertPortData(context.Background(), portsInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetPortsDataEndPoint endpoint for gets  port data
func GetPortsDataEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pServ := service.PortServ{}
	portService := service.PortService(pServ)
	portDetails, err := portService.GetPortsData(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portDetails)

}

// GetPortDataByIDEndPoint endpoint for gets the port data by id
func GetPortDataByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars[constants.ID]
	pServ := service.PortServ{}
	portService := service.PortService(pServ)
	portDetails, err := portService.GetPortDataByID(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portDetails)
}

// DeletePortByIDEndPoint endpoint for delete the port data by id
func DeletePortByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars[constants.ID]
	pServ := service.PortServ{}
	portService := service.PortService(pServ)
	err := portService.DeletePortByID(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdatePortByIDEndPoint endpoint for update the port data by id
func UpdatePortByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request models.PortDetails
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
		return
	}
	vars := mux.Vars(r)
	id := vars[constants.ID]
	pServ := service.PortServ{}
	portService := service.PortService(pServ)
	err2 := portService.UpdatePortByID(context.Background(), request, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err2)
		return
	}
	w.WriteHeader(http.StatusOK)
}
