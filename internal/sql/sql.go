package db

import (
	"context"
	"errors"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
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
	Conn *sqlx.DB
}

const (
	Get                 int = 0
	Select                  = 1
	Exec                    = 2
	ExecWithReturningId     = 3
)

// New initial new database connection with give configurations.
func New(cfg *Config) (db *Database, err error) {

	db = new(Database)

	if db.Conn, err = sqlx.Connect(cfg.Driver,
		cfg.URL); err != nil {
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

// Exec query -
// queryName = Get/Select/Exec/ExecReturnId
func (db *Database) ExecQuery(ctx context.Context,
	query int, sqlQuery string,
	dest interface{}, args ...interface{},
) (insertID uint64, err error) {

	tx, err := db.Conn.BeginTxx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.PreparexContext(ctx, sqlQuery)
	if err != nil {
		return
	}
	defer stmt.Close()
	switch query {
	case Get:
		err = stmt.GetContext(ctx, dest, args...)
		if err != nil {
			return
		}

	case Select:
		err = stmt.SelectContext(ctx, dest, args...)
		if err != nil {
			return
		}
	case Exec:
		_, err = stmt.ExecContext(ctx, args...)
		if err != nil {
			return
		}
	case ExecWithReturningId:
		err = stmt.QueryRowContext(ctx, args...).Scan(&insertID)
		if err != nil {
			return
		}
	default:
		return insertID,
			errors.New("unknown query name. Use: Get, Select, Exec, ExecReturnId")
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return

}
