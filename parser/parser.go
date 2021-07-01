package parser

import (
	"errors"
	"github.com/andersonribeir0/market/model"
)


type FileHandler interface {
	Open() error
	ParseLine() (model.Record, error)
	Close() error
}

type CsvFileParser struct {
	path     string
	fileName string
}

func (fp *CsvFileParser) Open() error {
	return errors.New("implement me")
}

func (fp *CsvFileParser) ParseLine() (model.Record, error) {
	return model.Record{}, errors.New("implement me")

}

func (fp *CsvFileParser) Close() error{
	return errors.New("implement me")
}