package repository

import (
	"errors"
	"fmt"
	"github.com/andersonribeir0/market/constants"
	"github.com/andersonribeir0/market/db"
	"github.com/andersonribeir0/market/model"
)

type IMarketRepository interface {
	Save(market model.Record) error
	GetItem(id string) (model.Record, error)
	GetItemsByDistrictId(id string) ([]model.Record, error)
}

type MarketRepository struct {
	conn      *db.DB
	tableName string
}

func (mr *MarketRepository) New() error{
	conn, err := db.NewDB()
	if err != nil {
		return errors.New(fmt.Sprintf(
			"It was not possible to get db con. error: %s",
			err.Error()))
	}
	mr.conn = conn
	if mr.tableName == "" {
		mr.tableName = constants.TableName
	}
	return nil
}

func (mr *MarketRepository) Save(market model.Record) error{
	return mr.conn.PutRecord(market, mr.tableName)
}

func (mr *MarketRepository) GetItem(id string) (model.Record, error) {
	var record model.Record
	rec, err := mr.conn.GetRecordById(id, mr.tableName)
	if err != nil {
		return record, err
	}
	return record.FromRecordMap(rec)
}

func (mr *MarketRepository) GetItemsByDistrictId(id string) ([]model.Record, error) {
	var record model.Record
	rec, err := mr.conn.GetRecordByDistrictId(id, mr.tableName)
	if err != nil {
		return nil, err
	}
	return record.FromRecordMapList(rec)
}
