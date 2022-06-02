package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	Driver string `yaml:"driver" mapstructure:"driver"`
	URL    string `yaml:"url" mapstructure:"url"`
}

// NewConfig returns new default database configuration.
func NewConfig() *Config {
	return &Config{}
}

// The Database repsents Database instance.
type Database struct {
	Conn *sql.DB
}

// New initial new database connection with give configurations.
func New(conf *Config) (db *Database, err error) {

	db = new(Database)
	if db.Conn, err = sql.Open(conf.Driver, conf.URL); err != nil {
		return nil, err
	}

	if err = db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Close database connection.
func (db *Database) Close() {
	db.Conn.Close()
}

// Ping database server.
func (db *Database) Ping(ctx context.Context) error {
	return db.Conn.PingContext(ctx)
}
