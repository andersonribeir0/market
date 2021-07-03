package db

import (
	"fmt"
	"github.com/andersonribeir0/market/utils"
	"os"
	"testing"
)

const tableName = "a_table_name"
var connTest, _ = NewDB()

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	connTest.DeleteTable(tableName)
	if err := connTest.CreateTable(tableName); err != nil {
		fmt.Fprintf(os.Stdout, "error creating table: %v\n", err)
		os.Exit(1)
	}
}

func tearDown() {
	connTest.DeleteTable(tableName)
}

func TestDB(t *testing.T) {
	record := utils.GetFakeRecord()
	err := connTest.PutRecord(record, tableName)
	if err != nil {
		t.Fatalf("Error on PutRecord %s", err.Error())
	}

	result, err := connTest.GetRecordById(record.Id.String(), tableName)
	if err != nil {
		t.Fatalf("Error on GetRecordById %s", err.Error())
	}
	if result == nil {
		t.Fatalf("Error on GetRecordById. Should get one record with id %s. But got nil", record.Id.String())
	}
	if result["ID"].(string) != record.Id.String() {
		t.Fatalf("Error on GetRecordById. Should get one record with id %s. But got %s",
			record.Id.String(), result["ID"])
	}

	resultList, err := connTest.GetRecordByDistrictId(*record.CodDist, tableName)
	if len(resultList) != 1 {
		t.Fatalf("Error on GetRecordByDistrictId. Result list must have size of 1. Got %d", len(resultList))
	}

	expectedCodDist := record.CodDist
	codDist := resultList[0]["CODDIST"]
	if *expectedCodDist != codDist.(string){
		t.Fatalf("Error on GetRecordByDistrictId. CodList should be %s. Got %s",
			*record.CodDist, resultList[0]["CODDIST"])
	}
}
