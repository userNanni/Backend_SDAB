package main

import (
	"testing"
)

func TestDB(t *testing.T) {
	db := dbConn()
	defer db.CloseConn()
}
