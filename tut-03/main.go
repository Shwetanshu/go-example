package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		logrus.Fatal(err)
	}
	query := `SELECT * FROM products;`
	interval := 30 * time.Minute

	// dbStreamer := streamer.DBStreamConfig{
	// 	db: db, query: query, interval: interval,
	// }

	dbStreamer, _ := streamer.New(
		streamer.WithDb(db),
		streamer.WithQuery(query),
		streamer.WithOutputFile("filname.csv")) // considering interval is optional. and here one can pass in any order

	dbStreamer.Start()
	time.Sleep(20 * time.Second)
	dbStreamer.Stop()
}
