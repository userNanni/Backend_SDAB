package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DimLoja struct {
	SkLoja          int
	NumeroLoja      int
	OmLoja          string
	Cidade          string
	EstadoFederacao string
	RegiaoMilitar   string
}

type QueryInfo struct {
	SQL      string
	ScanFunc func(*sql.Rows) ([]interface{}, error)
}

var queries = map[string]QueryInfo{
	"dim_loja": {
		SQL: "SELECT * FROM dim_loja;",
		ScanFunc: func(rows *sql.Rows) ([]interface{}, error) {
			var results []interface{}
			for rows.Next() {
				var d DimLoja
				err := rows.Scan(
					&d.SkLoja,
					&d.NumeroLoja,
					&d.OmLoja,
					&d.Cidade,
					&d.EstadoFederacao,
					&d.RegiaoMilitar,
				)
				if err != nil {
					return nil, err
				}
				results = append(results, d)
			}
			return results, nil
		},
	},
}

func (db *DataBase) Data(key string) ([]interface{}, error) {
	if key == "" {
		return nil, nil
	}

	queryInfo, ok := queries[key]
	if !ok {
		return nil, nil
	}

	rows, err := db.Query(queryInfo.SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := queryInfo.ScanFunc(rows)
	if err != nil {
		return nil, err
	}

	return results, nil
}
