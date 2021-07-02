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
	dynamo    *dynamodb.DynamoDB
}

func NewDB() (*DB, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint:                          aws.String("http://localhost:8000"),
			Region:                            aws.String("us-east-1"),
			LogLevel:                          aws.LogLevel(aws.LogDebug),
		},
	}))
	return &DB{ dynamo: dynamodb.New(sess) }, nil
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
			{
				AttributeName: aws.String("CODDIST"),
				KeyType:       aws.String("RANGE"),
			},
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{
				IndexName: aws.String("districtIdIdx"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("ID"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("CODDIST"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType:   aws.String("INCLUDE"),
					NonKeyAttributes: []*string{aws.String("CODDIST")},
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