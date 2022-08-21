package http

import (
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/endpoint"

	"github.com/gorilla/mux"
)

const (
	//BaseURL url for port service
	BaseURL = "/api/v1/port_domain_service"
	//RegisterPortsURI is used to register the ports
	RegisterPortsURI = BaseURL + "/ports"
	//GetPortsURI is used to get the ports
	GetPortsURI = BaseURL + "/ports"
	//GetPortByIDURI gets the port by id
	GetPortByIDURI = BaseURL + "/port/{id}"
	//UpdatePortURI updates the port for given port id
	UpdatePortURI = BaseURL + "/port/{id}"
	//DeletePortIRI deletes the port for given port id
	DeletePortIRI = BaseURL + "/port/{id}"
)

//InitTransport http endpoints
func InitTransport(r *mux.Router) {
	r.HandleFunc(RegisterPortsURI, endpoint.InsertPortDataEndPoint).Methods(http.MethodPost)
	r.HandleFunc(GetPortsURI, endpoint.GetPortsDataEndPoint).Methods(http.MethodGet)
	r.HandleFunc(GetPortByIDURI, endpoint.GetPortDataByIDEndPoint).Methods(http.MethodGet)
	r.HandleFunc(UpdatePortURI, endpoint.UpdatePortByIDEndPoint).Methods(http.MethodPut)
	r.HandleFunc(DeletePortIRI, endpoint.DeletePortByIDEndPoint).Methods(http.MethodDelete)
}
