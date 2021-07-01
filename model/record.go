package model

type Record struct {
	Id         int    `json:"ID"`
	Long       int    `json:"LONG"`
	Lat        int    `json:"LAT"`
	Setcens    int64  `json:"SETCENS"`
	AreaP      int64  `json:"AREAP"`
	CodDist    int    `json:"CODDIST"`
	Distrito   string `json:"DISTRITO"`
	CodSubPref int    `json:"CODSUBPREF"`
	SubPref    string `json:"SUBPREFE"`
	Regiao5    string `json:"REGIAO5"`
	Regiao8    string `json:"REGIAO8"`
	NomeFeira  string `json:"NOME_FEIRA"`
	Registro   string `json:"REGISTRO"`
	Logradouro string `json:"LOGRADOURO"`
	Numero     string `json:"NUMERO"`
	Bairro     string `json:"BAIRRO"`
	Referencia string `json:"REFERENCIA"`
}

