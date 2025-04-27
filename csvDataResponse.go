package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"reflect"

	_ "github.com/lib/pq"
)

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
