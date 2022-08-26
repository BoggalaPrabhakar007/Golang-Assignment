package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

// RepoLib service the repository related operations
type RepoLib interface {
	InsertMultipleRecords(ctx context.Context, databaseName string, collectionName string, resources []interface{}) (interface{}, error)
	UpdateRecord(ctx context.Context, databaseName string, collectionName string, filter map[string]interface{}, update map[string]interface{}) (int, int, error)
	GetRecord(ctx context.Context, databaseName string, collectionName string, result interface{}, filter map[string]interface{}, projection map[string]interface{}) error
	GetRecords(ctx context.Context, databaseName string, collectionName string, results interface{}, filter map[string]interface{}, projection map[string]interface{}) error
	DeleteRecordByID(ctx context.Context, databaseName string, collectionName string, id string) error
}

//RepoLibServ to do repo related operations
type RepoLibServ struct {
}

//NewRepoLibServ initialize the repo lib service
func NewRepoLibServ() RepoLibServ {
	return RepoLibServ{}
}

//ErrDocumentNotFound returned when document is not found
var ErrDocumentNotFound = errors.New("document not found")

//InsertMultipleRecords - Inserts multi records in to the collection. returns the list of inserted records and errors in case of failures.
func (r RepoLibServ) InsertMultipleRecords(ctx context.Context, databaseName string, collectionName string, resources []interface{}) (interface{}, error) {
	mClient, err := GetConnectionClient()
	if err != nil {
		return nil, err
	}
	mongoDataBase := mClient.Database(databaseName)
	mongoCollection := mongoDataBase.Collection(collectionName)
	//when SetOrdered is false - insertion will continue even if one of the document fails.
	options := options.InsertMany().SetOrdered(false)
	insertResult, err := mongoCollection.InsertMany(context.Background(), resources, options)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedIDs, err
}

//UpdateRecord - update the record in the collection after applying the filter.
func (r RepoLibServ) UpdateRecord(ctx context.Context, databaseName string, collectionName string, filter map[string]interface{}, update map[string]interface{}) (int, int, error) {
	var matchedCount int
	var modifiedCount int
	mClient, err := GetConnectionClient()
	if err != nil {
		return matchedCount, modifiedCount, err
	}
	mongoDataBase := mClient.Database(databaseName)
	mongoCollection := mongoDataBase.Collection(collectionName)
	updateResult, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return matchedCount, modifiedCount, err
	}

	matchedCount = int(updateResult.MatchedCount)
	modifiedCount = int(updateResult.ModifiedCount)
	return matchedCount, modifiedCount, nil
}

//GetRecord - returns the record from the collection after applying the filter.
func (r RepoLibServ) GetRecord(ctx context.Context, databaseName string, collectionName string, result interface{}, filter map[string]interface{}, projection map[string]interface{}) error {
	mClient, err := GetConnectionClient()
	if err != nil {
		return err
	}
	mongoDataBase := mClient.Database(databaseName)
	mongoCollection := mongoDataBase.Collection(collectionName)
	opts := &options.FindOneOptions{}
	if projection != nil {
		opts.Projection = projection
	}
	document := mongoCollection.FindOne(context.Background(), filter, opts)
	if document == nil {
		return ErrDocumentNotFound
	}
	if document.Err() != nil {
		if strings.EqualFold("mongo: no documents in result", document.Err().Error()) {
			return ErrDocumentNotFound
		}
		return document.Err()
	}
	err = document.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

//GetRecords - returns the records from the collection after applying the filter.
func (r RepoLibServ) GetRecords(ctx context.Context, databaseName string, collectionName string, results interface{}, filter map[string]interface{}, projection map[string]interface{}) error {
	mClient, err := GetConnectionClient()
	if err != nil {
		return err
	}
	mongoDataBase := mClient.Database(databaseName)
	mongoCollection := mongoDataBase.Collection(collectionName)
	opts := &options.FindOptions{}
	if projection != nil {
		opts.Projection = projection
	}
	documentsCursor, err := mongoCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return err
	}

	if err = documentsCursor.All(context.Background(), results); err != nil {
		return err
	}
	return nil
}

//DeleteRecordByID - Deletes the record with the specific ID provided.
func (r RepoLibServ) DeleteRecordByID(ctx context.Context, databaseName string, collectionName string, id string) error {
	mClient, err := GetConnectionClient()
	if err != nil {
		return err
	}
	mongoDataBase := mClient.Database(databaseName)
	mongoCollection := mongoDataBase.Collection(collectionName)
	resp, err := mongoCollection.DeleteOne(context.TODO(), bson.D{{"_id", id}}, nil)
	if err != nil {
		return err
	}
	if resp.DeletedCount == 0 {
		errMsg := fmt.Sprintf("no document found with the ID [%s]", id)
		return fmt.Errorf(errMsg)
	}
	return err
}
