package repository

import (
	"errors"
	"fmt"
	"github.com/andersonribeir0/market/db"
	"github.com/andersonribeir0/market/model"
)

type Repository interface {
	Save() error
}

type MarketRepository struct {
	market    model.Record
	conn      *db.DB
	tableName string
}

func (mr *MarketRepository) New(record model.Record) error{
	mr.market = record
	conn, err := db.NewDB()
	if err != nil {
		return errors.New(fmt.Sprintf(
			"It was not possible to get db con. error: %s",
			err.Error()))
	}
	mr.conn = conn
	return nil
}

func (mr *MarketRepository) Save() error{
	return mr.conn.PutRecord(mr.market, mr.tableName)
}

