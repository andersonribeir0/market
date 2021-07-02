package db

import (
	"errors"
	"fmt"
	"github.com/andersonribeir0/market/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DB struct {
	dynamo *dynamodb.DynamoDB
}

var conn *DB

func GetConn() *DB {
	if conn == nil || conn.dynamo == nil {
		conn, _ = NewDB()
	}
	return conn
}

func NewDB() (*DB, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint: aws.String("http://dynamodb:8000"),
			Region:   aws.String("us-east-1"),
			LogLevel: aws.LogLevel(aws.LogOff),
		},
	}))
	return &DB{dynamo: dynamodb.New(sess)}, nil
}

func (db *DB) PutRecord(item model.Record, tableName string) error {
	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		return errors.New(fmt.Sprintf(
			"It was not possible to parse map %#v error: %s",
			item,
			err.Error()))
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = db.dynamo.PutItem(input)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"It was not possible to put item %#v error: %s",
			item,
			err.Error()))
	}

	return nil
}

func (db *DB) GetRecordById(id string, tableName string) (map[string]interface{}, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
		TableName:      aws.String(tableName),
		ConsistentRead: aws.Bool(true),
	}

	item, err := db.dynamo.GetItem(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when getting by Id %s", err.Error()))
	}

	var record map[string]interface{}
	if item.Item != nil {
		err = dynamodbattribute.UnmarshalMap(item.Item, &record)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error when unmarhalling map %s", err.Error()))
		}
	}

	return record, nil
}

func (db *DB) GetRecordByDistrictId(id string, tableName string) ([]map[string]interface{}, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":districtId": {
				S: aws.String(id),
			},
		},
		IndexName:              aws.String("districtIdIdx"),
		KeyConditionExpression: aws.String("CODDIST = :districtId"),
		TableName:      		aws.String(tableName),
	}

	item, err := db.dynamo.Query(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when getting by Id %s", err.Error()))
	}

	var record []map[string]interface{}
	if item.Items != nil {
		err = dynamodbattribute.UnmarshalListOfMaps(item.Items, &record)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error when unmarhalling map %s", err.Error()))
		}
	}
	return record, nil
}

func (db *DB) DeleteTable(tableName string) error {
	newDB, err := NewDB()
	if err != nil {
		return err
	}
	_, err = newDB.dynamo.DeleteTable(&dynamodb.DeleteTableInput{TableName: &tableName})
	return err
}

func (db *DB) CreateTable(tableName string) error {
	newDB, err := NewDB()
	if err != nil {
		return err
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("CODDIST"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("districtIdIdx"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("CODDIST"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(1),
					WriteCapacityUnits: aws.Int64(2),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(2),
		},
		TableName: aws.String(tableName),
	}

	if _, err := newDB.dynamo.CreateTable(input); err != nil {
		return errors.New(fmt.Sprintf("Got error calling CreateTable: %s", err.Error()))
	}
	return nil
}
