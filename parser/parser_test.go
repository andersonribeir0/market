package parser

import (
	"fmt"
	"testing"
)

const FilePath = "./"
const FileName = "DEINFO_AB_FEIRASLIVRES_2014.csv"

func getCsvFileInstance() CsvFileParser {
	return CsvFileParser{
		path:     FilePath,
		fileName: FileName,
	}
}

func TestShouldParseCsvRecordSuccessfully(t *testing.T) {
	csvFileParser := getCsvFileInstance()

	if err := csvFileParser.Open(); err != nil {
		t.Fatalf(fmt.Sprintf("Error when trying to open file %s with path %s: %s",
			csvFileParser.fileName,
			csvFileParser.path,
			err.Error()))
	}

	_, err := csvFileParser.ParseLine()
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error when trying reading parse line: %s", err.Error()))
	}
}