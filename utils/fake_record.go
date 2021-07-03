package utils

import (
	"github.com/andersonribeir0/market/model"
	"math/rand"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var GetFakeRecord = func() model.Record {
	var record model.Record
	recData := map[string]interface{} {
		"ID":         rand.Int63(),
		"LONG":       rand.Int63(),
		"LAT":        rand.Int63(),
		"SETCENS":    rand.Int63(),
		"AREAP":      rand.Int63(),
		"CODDIST":    strconv.Itoa(rand.Int()),
		"DISTRITO":   randSeqString(5),
		"CODSUBPREF": strconv.Itoa(rand.Int()),
		"SUBPREFE":   randSeqString(5),
		"REGIAO5":    randSeqString(5),
		"REGIAO8":    randSeqString(5),
		"NOME_FEIRA": randSeqString(5),
		"REGISTRO":   randSeqString(5),
	    "LOGRADOURO": randSeqString(10),
		"NUMERO":     strconv.Itoa(rand.Int()),
		"BAIRRO":     randSeqString(5),
		"REFERENCIA": randSeqString(5),
	}
	rec, _ := record.FromRecordMap(recData)
	return rec
}

func randSeqString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}