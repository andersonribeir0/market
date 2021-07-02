package repository

import (
	"encoding/json"
	"fmt"
	"github.com/andersonribeir0/market/db"
	"github.com/andersonribeir0/market/model"
	"os"
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
	district := "Sana"
	id := json.Number('1')
	codDist := json.Number('5')
	err := repo.New(model.Record{
		Id:           &id,
		CodDist:      &codDist,
		District:     &district,
	})
	if err != nil {
		t.Fatalf("Impossible to create MarketRepository %s", err.Error())
	}
}

func TestMarketRepository_Save(t *testing.T) {
	conn = &db.DB{}
	repo := &MarketRepository{
		market: model.Record{},
		conn:   nil,
		tableName: tableName,
	}
	district := "Sanaaaaa"
	id := json.Number("1455588585588555554")
	codDist := json.Number("5122")
	areaP := json.Number("564558814158")
	repo.New(model.Record{
	    Id:     &id,
		CodDist:      &codDist,
		District:     &district,
		AreaP:        &areaP,
	})

	err := repo.Save()
	if err != nil {
		t.Fatalf("Error when saving a new record %s", err.Error())
	}
}