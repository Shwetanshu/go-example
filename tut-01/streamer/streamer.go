package streamer

import (
	"database/sql"
	"time"
)

type (
	DBStreamConfig struct {
		db       *sql.DB       // source from which to read data
		interval time.Duration // how often to read from source
		query    string        //query to run against source
		dest     string        // destination to store data. Eg, file,url,REST Endpoint etc
	}
)

// Start reading config data
func (conf *DBStreamConfig) Start() error {
	return nil
}

// Stop reading config data
func (conf *DBStreamConfig) Stop() error {
	return nil
}
