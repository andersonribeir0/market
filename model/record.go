package model

import (
	"encoding/json"
	"errors"
	"fmt"
)


type Record struct {
	Id               *json.Number     `json:"ID" validate:"required"`
	Long             *json.Number     `json:"LONG,omitempty"`
	Lat              *json.Number     `json:"LAT,omitempty"`
	SetCens          *json.Number  	  `json:"SETCENS,omitempty"`
	AreaP            *json.Number  	  `json:"AREAP,omitempty"`
	CodDist          *string     	  `json:"CODDIST" validate:"required"`
	District         *string 		  `json:"DISTRITO,omitempty"`
	CodSubPref       *json.Number     `json:"CODSUBPREF,omitempty"`
	SubPref          *string 		  `json:"SUBPREFE,omitempty"`
	Region5          *string 		  `json:"REGIAO5,omitempty"`
	Region8          *string 		  `json:"REGIAO8,omitempty"`
	MarketName       *string 		  `json:"NOME_FEIRA,omitempty"`
	Record           *string 		  `json:"REGISTRO,omitempty"`
	Street           *string 		  `json:"LOGRADOURO,omitempty"`
	Number           *string 		  `json:"NUMERO,omitempty"`
	Neighborhood     *string 		  `json:"BAIRRO,omitempty"`
	AddressRef       *string 		  `json:"REFERENCIA,omitempty"`
}

func (r Record) FromRecordMap(kv map[string]interface{}) (Record, error) {
	record := Record{}
	jsonString, err := json.Marshal(kv)
	if err != nil {
		return record, errors.New(fmt.Sprintf("Error when marshalling map to json string %#v %s", kv, err.Error()))
	}
	err = json.Unmarshal(jsonString, &record)
	if err != nil {
		return record, errors.New(fmt.Sprintf("Error when parsing map to struct %#v %s", kv, err.Error()))
	}
	return record, nil
}

func (r Record) FromRecordMapList(lkv []map[string]interface{}) ([]Record, error) {
	var record []Record
	jsonString, err := json.Marshal(lkv)
	if err != nil {
		return record, errors.New(fmt.Sprintf("Error when marshalling map list to json string %#v", lkv))
	}
	err = json.Unmarshal(jsonString, &record)
	if err != nil {
		return record, errors.New(fmt.Sprintf("Error when parsing map list to struct %#v", lkv))
	}
	return record, nil
}
