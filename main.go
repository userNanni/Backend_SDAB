package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"reflect"

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

	data, error := db.Query("SELECT cidade FROM fato_indicadores_nivel_servico_lojas INNER JOIN dim_tempo USING (sk_tempo) INNER JOIN dim_loja USING (sk_loja);")

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

func dataResponse(w http.ResponseWriter, _ *http.Request, data ...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func csvDataResponse(w http.ResponseWriter, _ *http.Request, dataList any) {

	v := reflect.ValueOf(dataList)

	if v.Kind() != reflect.Slice {
		http.Error(w, "Invalid data type: expected slice", http.StatusBadRequest)
		return
	}

	if v.Len() == 0 {
		http.Error(w, "Empty data list", http.StatusBadRequest)
		return
	}

	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	firstElem := v.Index(0)
	if firstElem.Kind() == reflect.Ptr {
		firstElem = firstElem.Elem()
	}

	t := firstElem.Type()

	headers := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	if err := writer.Write(headers); err != nil {
		http.Error(w, "Failed to write CSV header", http.StatusInternalServerError)
		return
	}

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		row := make([]string, t.NumField())
		for j := 0; j < t.NumField(); j++ {
			field := elem.Field(j)
			row[j] = fieldToString(field)
		}

		if err := writer.Write(row); err != nil {
			http.Error(w, "Failed to write CSV row", http.StatusInternalServerError)
			return
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		http.Error(w, "Failed to flush CSV data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

func fieldToString(v reflect.Value) string {
	if !v.IsValid() {
		return ""
	}
	return InterfaceAsString(v)
}

func InterfaceAsString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	default:
		return ""
	}
}
