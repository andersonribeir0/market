package repository

import (
	"encoding/json"
	"fmt"
	"github.com/andersonribeir0/market/db"
	"github.com/andersonribeir0/market/model"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

const tableName = "market_test_table"
var conn *db.DB

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	conn = &db.DB{}
	conn.DeleteTable(tableName)
	if err := conn.CreateTable(tableName); err != nil {
		fmt.Fprintf(os.Stdout, "error creating table: %v\n", err)
		os.Exit(1)
	}
}

func tearDown() {
	conn = &db.DB{}
	conn.DeleteTable(tableName)
}

func TestMarketRepository_New(t *testing.T) {
	repo := &MarketRepository{}
	err := repo.New()
	if err != nil {
		t.Fatalf("Impossible to create MarketRepository %s", err.Error())
	}
}

func putRecord(district string, id string, codDist string, areaP string) (*db.DB, error){
	conn = db.GetConn()
	repo := &MarketRepository{
		conn:   nil,
		tableName: tableName,
	}
	mid := json.Number(id)
	cod := json.Number(codDist)
	area := json.Number(areaP)
	repo.New()
	return conn, repo.Save(model.Record{
		Id:           &mid,
		CodDist:      &cod,
		District:     &district,
		AreaP:        &area,
	})
}

func TestMarketRepository_Save(t *testing.T) {
	conn = &db.DB{}
	repo := &MarketRepository{
		conn:   nil,
		tableName: tableName,
	}
	district := "Sanaaaaa"
	id := json.Number("1455588585588555554")
	codDist := json.Number("5122")
	areaP := json.Number("564558814158")
	repo.New()

	err := repo.Save(model.Record{
		Id:           &id,
		CodDist:      &codDist,
		District:     &district,
		AreaP:        &areaP,
	})
	if err != nil {
		t.Fatalf("Error when saving a new record %s", err.Error())
	}
}


func TestMarketRepository_GetById(t *testing.T) {
	id := strconv.Itoa(rand.Int())
	conn, err := putRecord("Goytacazes", id, "14", "41")
	if err != nil {
		t.Fatalf("Error when saving a new record %s", err.Error())
	}
	rec, err := conn.GetRecordById(id, tableName)
	if err != nil {
		t.Fatalf("Error when getting record %s", err.Error())
	}
	if rec["ID"].(string) != id {
		t.Fatalf("Expected to return record with id %s. Got %s", id, rec["ID"])
	}
}

func TestMarketRepository_GetByDistrictId(t *testing.T) {
	id := strconv.Itoa(rand.Int())
	districtId := "14"
	conn, err := putRecord("Goytacazes", id, districtId, "41")
	if err != nil {
		t.Fatalf("Error when saving a new record %s", err.Error())
	}
	rec, err := conn.GetRecordByDistrictId(districtId, tableName)
	if err != nil {
		t.Fatalf("Error when getting record %s", err.Error())
	}

	if rec[0]["CODDIST"].(string) != districtId {
		t.Fatalf("Expected to return record with id %s. Got %s", id, rec[0]["CODDIST"].(string))
	}
}