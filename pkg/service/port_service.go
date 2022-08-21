package service

import (
	"context"
	"encoding/json"
	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/repo"
	"io/ioutil"
	"log"
	"net/http"
)

// PortService service for port related activities
type PortService interface {
	InsertPortData(ctx context.Context, r *http.Request) error
	GetPortsData(ctx context.Context) ([]models.PortDetails, error)
	GetPortDataByID(ctx context.Context, id string) (models.PortDetails, error)
	DeletePortByID(ctx context.Context, id string) error
	UpdatePortByID(ctx context.Context, port models.PortDetails, id string) error
}

//PortServ to do operation on port data
type PortServ struct {
}

var repoPortInfo = make(map[string]models.Port)

// InsertPortData will read the data from the json file and insert the data in repo
func (p PortServ) InsertPortData(ctx context.Context, r *http.Request) error {
	// if u want to read the data from http request uncomment the below lines and comment from line no 40 to 48
	/*var portsInfo = make(map[string]models.Port)
	err := json.NewDecoder(r.Body).Decode(&portsInfo)
	if err != nil {
		log.Fatal(err)
	}*/
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
	rServ := repo.PortRepoServ{}
	repoService := repo.PortRepoService(rServ)
	err3 := repoService.InsertPorts(ctx, fPortInfo)
	return err3
}

// GetPortsData gets the port data from repo
func (p PortServ) GetPortsData(ctx context.Context) ([]models.PortDetails, error) {
	//repo layer call for getting data from database
	rServ := repo.PortRepoServ{}
	repoService := repo.PortRepoService(rServ)
	portDetails, err := repoService.GetPorts(ctx)
	return portDetails, err

}

// GetPortDataByID gets the port data by id from repo
func (p PortServ) GetPortDataByID(ctx context.Context, id string) (models.PortDetails, error) {
	//repo layer call for getting data from database
	rServ := repo.PortRepoServ{}
	repoService := repo.PortRepoService(rServ)
	portDetails, err := repoService.GetPortByID(ctx, id)
	return portDetails, err

}

// DeletePortByID delete the port data by id from repo
func (p PortServ) DeletePortByID(ctx context.Context, id string) error {
	//repo layer call for getting data from database
	rServ := repo.PortRepoServ{}
	repoService := repo.PortRepoService(rServ)
	err := repoService.DeletePortByID(ctx, id)
	return err
}

// UpdatePortByID update the port data by id from repo
func (p PortServ) UpdatePortByID(ctx context.Context, port models.PortDetails, id string) error {
	//repo layer call for getting data from database
	rServ := repo.PortRepoServ{}
	repoService := repo.PortRepoService(rServ)
	err := repoService.UpdatePortByID(ctx, id, &port)
	return err
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
