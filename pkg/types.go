package pkg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Database represents the database interface
type Database interface {
	Connect() (*sql.DB, error)
	GetNow() (string, error)
}
