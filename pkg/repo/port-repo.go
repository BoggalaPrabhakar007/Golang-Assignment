package repo

import (
	"context"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"

	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	repository "github.com/BoggalaPrabhakar007/golang-assignment/repository-lib/pkg/mongodb"
)

const CollectionName = "PortsCollection"

// InsertPorts insert the data into database
func InsertPorts(_ context.Context, portsDetails []models.PortDetails) error {
	portsDetailsDocs := make([]interface{}, len(portsDetails))
	for i, val := range portsDetails {
		portsDetailsDocs[i] = val
	}
	_, err := repository.InsertMultipleRecords(context.Background(), config.DatabaseName, CollectionName, portsDetailsDocs)
	if err != nil {
		return err
	}
	return nil
}

// GetPorts gets the data from database
func GetPorts(_ context.Context) ([]models.PortDetails, error) {
	var portsDetails []models.PortDetails
	var filter = make(map[string]interface{})
	err := repository.GetRecords(context.Background(), config.DatabaseName, CollectionName, &portsDetails, filter, nil)
	return portsDetails, err
}

// GetPortByID gets the data from database using id
func GetPortByID(_ context.Context, id string) (models.PortDetails, error) {
	var portDetails models.PortDetails
	var filter = make(map[string]interface{})
	filter[constants.DBID] = id
	err := repository.GetRecord(context.Background(), config.DatabaseName, CollectionName, &portDetails, filter, nil)
	return portDetails, err
}

// DeletePortByID delete the data from database using id
func DeletePortByID(_ context.Context, id string) error {
	err := repository.DeleteRecordByID(context.Background(), config.DatabaseName, CollectionName, id)
	return err
}

//UpdatePortByID will edit the port Details
func UpdatePortByID(ctx context.Context, id string, port models.Port) error {
	updateMap := make(map[string]interface{})
	filter := make(map[string]interface{})
	updateMap[constants.Name] = port.Name
	updateMap[constants.City] = port.City
	updateMap[constants.Code] = port.Code
	updateMap[constants.Unlocs] = port.Unlocs
	updateMap[constants.Province] = port.Province
	updateMap[constants.Alias] = port.Alias
	updateMap[constants.Regions] = port.Regions
	updateMap[constants.TimeZone] = port.Timezone
	updateMap[constants.Coordinates] = port.Coordinates
	updateMap[constants.Country] = port.Country
	filter[constants.DBID] = id

	_, _, err := repository.UpdateRecord(ctx, config.DatabaseName, CollectionName, filter, updateMap)
	if err != nil {
		return err
	}
	return nil
}
