package postgres

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	// postgres driver
	_ "github.com/lib/pq"
)

// DB represents a database for the code to interface with.
type DB struct {
	*sqlx.DB
}

// CreateDB connects to a Postgres database.
func CreateDB() (*DB, error) {
	dataSource := os.Getenv("DB_URL")
	fmt.Printf("Endpoint: %v\n", dataSource)
	conn, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		return nil, errors.Wrapf(err, "Error creating db")
	}
	db := DB{
		DB: conn,
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &db, nil
}

// Transact provides a wrapper for transactions.
func (db *DB) Transact(txFunc func(*sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				err = r
			default:
				err = fmt.Errorf("%s", r)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	err = txFunc(tx)
	return err
}

// Close closes all DB connections.
func (db *DB) Close() error {
	err := db.DB.Close()
	if err != nil {
		return errors.Wrap(err, "cannot close db")
	}
	return nil
}

func buildValues(numCols int) string {
	placeholders := make([]string, numCols)
	for c := 0; c < numCols; c++ {
		placeholders[c] = "?"
	}
	placeholderStr := fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
	return sqlx.Rebind(sqlx.DOLLAR, placeholderStr)
}
