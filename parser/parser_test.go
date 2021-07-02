package parser

import (
	"os"
	"path/filepath"
	"testing"
)


func getCsvFileInstance() CsvFileParser {
	path := filepath.Base("./mock_input")
	return CsvFileParser{
		path:     path,
		fileName: string(os.PathSeparator) + "csv_test.csv",
	}
}

func getInvalidCsvFileInstance() CsvFileParser {
	path := filepath.Base("./mock_input")
	return CsvFileParser{
		path:     path,
		fileName: string(os.PathSeparator) + "invalid_csv_test.csv",
	}
}

func TestShouldParseCsvRecordSuccessfully(t *testing.T) {
	csvFileParser := getCsvFileInstance()

	records, err := csvFileParser.Parse()
	if err != nil {
		t.Fatalf("Error when trying to parse file %s with path %s: %s",
			csvFileParser.fileName,
			csvFileParser.path,
			err.Error())
	}

	if records == nil {
		t.Fatalf("Empty records")
	}

}

func TestShouldReturnErrorWhenCsvFileIsInvalid(t *testing.T) {
	csvFileParser := getInvalidCsvFileInstance()
	_, err := csvFileParser.Parse()
	if err == nil {
		t.Fatalf("Should return an error when csv file is invalid")
	}
}
