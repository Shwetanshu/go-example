package streamer

import (
	"database/sql"
	"time"
)

type (
	dbStreamConfig struct {
		db       *sql.DB       // source from which to read data
		interval time.Duration // how often to read from source
		query    string        //query to run against source
		dest     string        // destination to store data. Eg, file,url,REST Endpoint etc
	}
)

func New(db *sql.Db, query string, interval time.Duration, dest string) (*dbStreamConfig, error) {
	conf := &dbStreamConfig{
		db: db, 
		query: query, 
		interval: interval, 
		dest: dest
	}

	return &conf, nil
}

// Start reading config data
func (conf *DBStreamConfig) Start() error {
	return nil
}

// Stop reading config data
func (conf *DBStreamConfig) Stop() error {
	return nil
}
