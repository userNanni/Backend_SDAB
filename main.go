package main

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *DataBase
}

func NewAPIServer(addr string, db *DataBase) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func main() {
	startPort := 8080
	port := findAvailablePort(startPort)

	db := dbConn()
	defer db.CloseConn()

	server := NewAPIServer(fmt.Sprintf(":%d", port), db)

	if err := server.Run(); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/helloweb", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Web!"))
	})

	router.HandleFunc("/api/v1/dim_loja", func(w http.ResponseWriter, r *http.Request) {
		data, err := s.db.Data("dim_loja")
		if err != nil {
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return





























































			
		}
		dataResponse(w, r, data)
	})

	router.HandleFunc("/api/v1/csv/dim_loja", func(w http.ResponseWriter, r *http.Request) {
		data, err := s.db.Data("dim_loja")
		if err != nil {
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return
		}
		csvDataResponse(w, r, data)
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("Starting server on", s.addr)

	return server.ListenAndServe()
}
