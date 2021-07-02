package model

import "testing"

func TestRecord_FromRecordMap(t *testing.T) {
	rec := Record{}
	district := "Goytacazes"
	codDist := "12"
	id := int64(1)
	kv := map[string]interface{} {
		"ID": id,
		"CODDIST": codDist,
		"DISTRITO": district,
	}
	rec, err := rec.FromRecordMap(kv)
	if err != nil {
		t.Fatalf("Expected to get a record. %s", err.Error())
	}

	if rec.District == nil || *rec.District != district {
		t.Fatalf("Expected district to be equals to %s. But got %s", district, *rec.District)
	}

	recId, _ := rec.Id.Int64()
	if recId != id {
		t.Fatalf("Expected id to be equals to %d. But got %d", recId, id)
	}

	if codDist != *rec.CodDist {
		t.Fatalf("Expected codDist to be equals to %s. But got %s", codDist,  *rec.CodDist)
	}
}


func TestRecord_FromRecordMapList(t *testing.T) {
	var rec Record
	district := "Goytacazes"
	codDist := "12"
	id := int64(1)
	kv := map[string]interface{} {
		"ID": id,
		"CODDIST": codDist,
		"DISTRITO": district,
	}
	lkv := []map[string]interface{}{kv}

	recs, err := rec.FromRecordMapList(lkv)
	if err != nil {
		t.Fatalf("Expected to get a record. %s", err.Error())
	}

	if len(recs) != 1 {
		t.Fatalf("Expected length to be equals to 1. Got %d", len(recs))
	}
}
