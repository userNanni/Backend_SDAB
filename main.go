package main

import (
	"encoding/json"
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
	server := NewAPIServer(":8080")
	db := dbConn()

	data,error := db.Query("SELECT cidade FROM fato_indicadores_nivel_servico_lojas INNER JOIN dim_tempo USING (sk_tempo) INNER JOIN dim_loja USING (sk_loja);")

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
		//fmt.Println("Cidade:", d.Cidade)
	}

	//fmt.Println("Data List:", dataList)

	server.Run()

	defer db.CloseConn()
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/helloweb", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Web!"))
	})

	router.HandleFunc("/api/v1/lojas", func(w http.ResponseWriter, r *http.Request) {
		dataResponse(w, r, dataList)
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("Starting server on", s.addr)

	return server.ListenAndServe()
}

func dataResponse(w http.ResponseWriter, r *http.Request, data ...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}