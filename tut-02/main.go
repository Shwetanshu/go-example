package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/striversity/gotr/misc-ep005/ex02/streamer"

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

	dbStreamer, _ := streamer.New(db, query, interval, "filname.csv") // limitation is that if someone want to pass few parameters , they still have to declare all the parameters. Look into tut-03 for improvement

	dbStreamer.Start()
	//..
	dbStreamer.Stop()
}
