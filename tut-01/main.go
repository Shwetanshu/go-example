package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/striversity/gotr/misc-ep005/ex01/streamer"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		logrus.Fatal(err)
	}
	query := `SELECT * FROM products;`
	interval := 30 * time.Minute
	dbStreamer := streamer.DBStreamConfig{ // Calling struct directly with different fields is a bad idea
		db: db, query: query, interval: interval,
	}

	dbStreamer.Start()
	//..
	dbStreamer.Stop()
}
