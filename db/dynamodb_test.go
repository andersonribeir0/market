package db

import "testing"

func TestTableOperations(t *testing.T) {
	tableName := "a_table_name"
	conn := &DB{}
	err := conn.CreateTable(tableName)
	if err != nil {
		t.Fatalf("An error occurred when trying to create a table in dynamodb %s", err.Error())
	}
	err = conn.DeleteTable(tableName)
	if err != nil {
		t.Fatalf("An error occurred when trying to delete table in dynamodb %s", err.Error())
	}
}