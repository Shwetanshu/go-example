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

	config func(c dbStreamConfig) dbStreamConfig
)

var (
	defaultInterval = 5 * time.Second
)

func WithDb(c *sql.DB) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.db = db
		return c
	}
}

func WithQuery(q string) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.query = q
		return c
	}
}

func WithOutputFile(fn string) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.dest = fn
		return c
	}
}

func WithInterval(interval time.Duration) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.interval = interval
		return c
	}
}

func New(options ...config) (*dbStreamConfig, error) {
	conf := dbStreamConfig{interval: defaultInterval}

	for _, opt := range options {
		opt(conf)
	}

	return &conf, nil
}

// Start reading config data
func (conf *DBStreamConfig) Start() error {
	go func(c *dbStreamConfig) {
		for {
			logrus.Info("Fetching Data at %v", time.Now())
			time.Sleep(c.interval)
		}
	}(conf)
	return nil
}

// Stop reading config data
func (conf *DBStreamConfig) Stop() error {
	return nil
}
