package main

/*

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestData(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	database := &DataBase{db}

	t.Run("Empty Information", func(t *testing.T) {
		result := database.data("")
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("Valid Information - lojas", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"cidade"}).
			AddRow("São Paulo").
			AddRow("Rio de Janeiro")

		mock.ExpectQuery("SELECT cidade FROM fato_indicadores_nivel_servico_lojas").
			WillReturnRows(rows)

		result := database.data("lojas")
		if len(result) != 2 {
			t.Errorf("Expected 2 rows, got %d", len(result))
		}

		if result[0].Cidade != "São Paulo" || result[1].Cidade != "Rio de Janeiro" {
			t.Errorf("Unexpected data: %v", result)
		}
	})

	t.Run("Query Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT cidade FROM fato_indicadores_nivel_servico_lojas").
			WillReturnError(sql.ErrConnDone)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected log.Fatal to be called, but it did not")
			}
		}()

		log.SetFlags(0)
		log.SetOutput(nil) // Suppress log output
		database.data("lojas")
	})
}
*/
