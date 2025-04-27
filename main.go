package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

type Data struct {
	Cidade string
}

var dataList []Data

func main() {
	startPort := 8080
	port := findAvailablePort(startPort)

	server := NewAPIServer(fmt.Sprintf(":%d", port))

	db := dbConn()

	dataList = db.data("lojas")

	defer db.CloseConn()

	if err := server.Run(); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/helloweb", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Web!"))
	})

	router.HandleFunc("/api/v1/lojas", func(w http.ResponseWriter, r *http.Request) {
		dataResponse(w, r, dataList)
	})

	router.HandleFunc("/api/v1/csv/lojas", func(w http.ResponseWriter, r *http.Request) {
		csvDataResponse(w, r, dataList)
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("Starting server on", s.addr)

	return server.ListenAndServe()
}
