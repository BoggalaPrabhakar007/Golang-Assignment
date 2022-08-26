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

//portEndpoints expose endpoint for the port service
type portEndpoints struct {
	pServ service.PortService
}

//NewEndpoint initialize endpoint
func NewEndpoint(pServ service.PortService) portEndpoints {
	return portEndpoints{
		pServ: pServ,
	}
}

// InsertPortDataEndPoint endpoint for insert data
func (p portEndpoints) InsertPortDataEndPoint(w http.ResponseWriter, r *http.Request) {
	var portsInfo = make(map[string]models.Port)
	// if u want to read the data from http request uncomment the below lines
	/*err := json.NewDecoder(r.Body).Decode(&portsInfo)
	if err != nil {
		log.Fatal(err)
	}*/

	err := p.pServ.InsertPortData(context.Background(), portsInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetPortsDataEndPoint endpoint for gets  port data
func (p portEndpoints) GetPortsDataEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	portDetails, err := p.pServ.GetPortsData(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portDetails)

}

// GetPortDataByIDEndPoint endpoint for gets the port data by id
func (p portEndpoints) GetPortDataByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars[constants.ID]
	portDetails, err := p.pServ.GetPortDataByID(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portDetails)
}

// DeletePortByIDEndPoint endpoint for delete the port data by id
func (p portEndpoints) DeletePortByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars[constants.ID]
	err := p.pServ.DeletePortByID(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdatePortByIDEndPoint endpoint for update the port data by id
func (p portEndpoints) UpdatePortByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request models.PortDetails
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
		return
	}
	vars := mux.Vars(r)
	id := vars[constants.ID]
	err2 := p.pServ.UpdatePortByID(context.Background(), request, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err2)
		return
	}
	w.WriteHeader(http.StatusOK)
}
