package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type FileHandler interface {
	Parse() error
}

type CsvFileParser struct {
	path               string
	fileName           string
	csvDataMap		   []map[string]interface{}
	keys			   []string
}

func (fp *CsvFileParser) New(path string, fileName string) *CsvFileParser {
	return &CsvFileParser{
		path:       path,
		fileName:   fileName,
	}
}

func (fp *CsvFileParser) Parse() ([]map[string]interface{}, error) {
	f, err := os.Open(fp.path + fp.fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when trying to open file %s from path %s: %s",
			fp.path,
			fp.fileName,
			err))
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	keys, err := csvReader.Read()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting csv record header from file %s from path %s: %s",
			fp.path,
			fp.fileName,
			err))
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf(
			"An error occurred processing csv file %#v err: %s",
			fp,
			err.Error()))
	}

	for i := range records {
		fp.addCsvDataMapItem(keys, records[i])
	}

	return fp.csvDataMap, nil
}

func (fp *CsvFileParser) GetCsvDataMap() []map[string]interface{} {
	return fp.csvDataMap
}

func (fp *CsvFileParser) addCsvDataMapItem(keys []string, record []string) {
	recordItem := make(map[string]interface{})
	for i := range record {
		recordItem[keys[i]] = record[i]
	}
	fp.csvDataMap = append(fp.csvDataMap, recordItem)
}
