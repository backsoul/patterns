package database_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/backsoul/pattern/pkg"
)

// MockDatabaseConnector es un mock para DatabaseConnector.
type MockDatabaseConnector struct {
	DB *sql.DB
}

func (m *MockDatabaseConnector) Connect() (*sql.DB, error) {
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	m.DB = db

	return db, nil
}

func (m *MockDatabaseConnector) GetNow() (string, error) {
	rows := m.DB.QueryRow("SELECT NOW()")
	var now string
	if err := rows.Scan(&now); err != nil {
		return "", err
	}
	return now, nil
}

func TestDatabaseIntegration(t *testing.T) {
	var dbConnector pkg.Database = &MockDatabaseConnector{}
	db, err := dbConnector.Connect()
	if err != nil {
		t.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()
}
