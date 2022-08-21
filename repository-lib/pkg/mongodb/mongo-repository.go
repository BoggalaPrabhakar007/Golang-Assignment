package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

//ErrDocumentNotFound returned when document is not found
var ErrDocumentNotFound = errors.New("document not found")

//InsertRecord - inserts the "resource" param in to collection. Returns the id of the inserted records and error in case of failures
func InsertRecord(_ context.Context, databaseName string, collectionName string, resource interface{}) (interface{}, error) {
	mClient, err := GetConnectionClient()
	if err != nil {
		return nil, err
	}
	mongoDataBase := mClient.Database(databaseName)
	mongoCollection := mongoDataBase.Collection(collectionName)
	createResponse, err := mongoCollection.InsertOne(context.Background(), resource)
	if err != nil {
		return nil, err
	}
	return createResponse.InsertedID, err
}

//InsertMultipleRecords - Inserts multi records in to the collection. returns the list of inserted records and errors in case of failures.
func InsertMultipleRecords(_ context.Context, databaseName string, collectionName string, resources []interface{}) (interface{}, error) {
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

func UpdateRecord(_ context.Context, databaseName string, collectionName string, filter map[string]interface{}, update map[string]interface{}) (int, int, error) {
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

func GetRecord(_ context.Context, databaseName string, collectionName string, result interface{}, filter map[string]interface{}, projection map[string]interface{}) error {
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
func GetRecords(_ context.Context, databaseName string, collectionName string, results interface{}, filter map[string]interface{}, projection map[string]interface{}) error {
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
func DeleteRecordByID(ctx context.Context, databaseName string, collectionName string, id string) error {
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
