package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/contracts/domain"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/repo"
)

// PortService service for port related activities
type PortService interface {
	InsertPortData(ctx context.Context, ports map[string]models.Port) error
	GetPortsData(ctx context.Context) ([]models.PortDetails, error)
	GetPortDataByID(ctx context.Context, id string) (models.PortDetails, error)
	DeletePortByID(ctx context.Context, id string) error
	UpdatePortByID(ctx context.Context, port models.PortDetails, id string) error
}

//PortServ to do operation on port data
type PortServ struct {
	pRepo  repo.PortRepoService
	config domain.Config
}

//NewPortService initialize the port service
func NewPortService(pRepo repo.PortRepoService, config domain.Config) PortService {
	return PortServ{
		pRepo:  pRepo,
		config: config,
	}
}

// InsertPortData will read the data from the json file and insert the data in repo
func (p PortServ) InsertPortData(ctx context.Context, portsInfo map[string]models.Port) error {
	// If you want to pass the data from endpoint layer comment the code from line no 45 to 52
	/* Reading data from big JSON file we can do by using the json.NewDecoder(),empty interfafe{}, channels
	   and finally read the data from channels and insert the data in databse to do CRUD operations doing
	   like below.
	*/
	//read the port data from the file
	pData, err := ioutil.ReadFile(p.config.File.Path)
	if err != nil {
		log.Fatal(err)
	}
	err2 := json.Unmarshal(pData, &portsInfo)
	if err2 != nil {
		log.Fatal(err2)
	}
	fPortInfo := getFormattedPortData(portsInfo)
	//repo layer call for inserting data into database
	err3 := p.pRepo.InsertPorts(ctx, fPortInfo)
	return err3
}

// GetPortsData gets the port data from repo
func (p PortServ) GetPortsData(ctx context.Context) ([]models.PortDetails, error) {
	//repo layer call for getting data from database
	portDetails, err := p.pRepo.GetPorts(ctx)
	return portDetails, err

}

// GetPortDataByID gets the port data by id from repo
func (p PortServ) GetPortDataByID(ctx context.Context, id string) (models.PortDetails, error) {
	//repo layer call for getting data from database
	portDetails, err := p.pRepo.GetPortByID(ctx, id)
	return portDetails, err

}

// DeletePortByID delete the port data by id from repo
func (p PortServ) DeletePortByID(ctx context.Context, id string) error {
	//repo layer call for getting data from database
	err := p.pRepo.DeletePortByID(ctx, id)
	return err
}

// UpdatePortByID update the port data by id from repo
func (p PortServ) UpdatePortByID(ctx context.Context, port models.PortDetails, id string) error {
	//repo layer call for getting data from database
	err := p.pRepo.UpdatePortByID(ctx, id, &port)
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
