package main

import (
	"log"

	_ "github.com/lib/pq"
)

func (db *DataBase) data(Information string) []Data {
	if Information == "" {
		log.Println("No information provided")
		return nil
	}

	var sql string

	var dataList []Data

	if Information == "lojas" {
		log.Println("Fetching data for lojas")
		sql = "SELECT cidade FROM fato_indicadores_nivel_servico_lojas INNER JOIN dim_tempo USING (sk_tempo) INNER JOIN dim_loja USING (sk_loja);"
	}

	data, error := db.Query(sql)

	if error != nil {
		log.Fatal("Error querying database:", error)
	}

	for data.Next() {
		var d Data
		if err := data.Scan(&d.Cidade); err != nil {
			log.Println("Error scanning data:", err)
			continue
		}
		dataList = append(dataList, d)
	}

	return dataList
}
