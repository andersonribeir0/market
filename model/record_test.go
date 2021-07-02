package model

import "testing"

func TestRecord_FromCsvRecordMap(t *testing.T) {
	rec := Record{}
	district := "Goytacazes"
	codDist := int64(12)
	id := int64(1)
	kv := map[string]interface{} {
		"ID": id,
		"CODDIST": codDist,
		"DISTRITO": district,
	}
	rec, err := rec.FromCsvRecordMap(kv)
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

	recCodDist, _ := rec.CodDist.Int64()
	if recCodDist != codDist {
		t.Fatalf("Expected codDist to be equals to %d. But got %d", recCodDist, codDist)
	}
}
