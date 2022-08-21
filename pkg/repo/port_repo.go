package repo

import (
	"context"
	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	repository "github.com/BoggalaPrabhakar007/golang-assignment/repository-lib/pkg/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

const CollectionName = "PortsCollection"

// PortRepoService service for port repo related activities
type PortRepoService interface {
	InsertPorts(ctx context.Context, portsDetails []models.PortDetails) error
	GetPorts(ctx context.Context) ([]models.PortDetails, error)
	GetPortByID(ctx context.Context, id string) (models.PortDetails, error)
	DeletePortByID(ctx context.Context, id string) error
	UpdatePortByID(ctx context.Context, id string, port *models.PortDetails) error
}

//PortRepoServ to do operation on port repo data
type PortRepoServ struct {
}

// InsertPorts insert the data into database
func (p PortRepoServ) InsertPorts(ctx context.Context, portsDetails []models.PortDetails) error {
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
func (p PortRepoServ) GetPorts(ctx context.Context) ([]models.PortDetails, error) {
	var portsDetails []models.PortDetails
	var filter = make(map[string]interface{})
	err := repository.GetRecords(context.Background(), config.DatabaseName, CollectionName, &portsDetails, filter, nil)
	return portsDetails, err
}

// GetPortByID gets the data from database using id
func (p PortRepoServ) GetPortByID(ctx context.Context, id string) (models.PortDetails, error) {
	var portDetails models.PortDetails
	var filter = make(map[string]interface{})
	filter[constants.DBID] = id
	err := repository.GetRecord(context.Background(), config.DatabaseName, CollectionName, &portDetails, filter, nil)
	return portDetails, err
}

// DeletePortByID delete the data from database using id
func (p PortRepoServ) DeletePortByID(ctx context.Context, id string) error {
	err := repository.DeleteRecordByID(context.Background(), config.DatabaseName, CollectionName, id)
	return err
}

//UpdatePortByID will edit the port Details
func (p PortRepoServ) UpdatePortByID(ctx context.Context, id string, port *models.PortDetails) error {
	filter := bson.M{constants.DBID: id}
	update, err := getUpdateObject(port)
	if err != nil {
		return err
	}
	_, _, err = repository.UpdateRecord(ctx, config.DatabaseName, CollectionName, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

// Gets update object in BSON format
func getUpdateObject(v interface{}) (interface{}, error) {
	pByte, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return nil, err
	}
	return update, nil
}
