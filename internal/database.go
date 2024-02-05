package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/backsoul/pattern/pkg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// MySQLFactory implements the Database interface for MySQL
type MySQLFactory struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

// Connect creates and returns a connection to the MySQL database
func (mysql *MySQLFactory) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.DBName)

	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the MySQL database!")
	return db, nil
}

func (mysql *MySQLFactory) GetNow() (string, error) {
	return "SELECT NOW()", nil
}

// PostgreSQLFactory implements the Database interface for PostgreSQL
type PostgreSQLFactory struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Connect creates and returns a connection to the PostgreSQL database
func (pg PostgreSQLFactory) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")
	return db, nil
}

func (pg *PostgreSQLFactory) GetNow() (string, error) {
	return "SELECT NOW()", nil
}

// Init initializes and returns a Database instance based on the configuration provided in the .env file
func Init() (pkg.Database, error) {
	var dbFactory pkg.Database
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file")
	}

	// Get values from environment variables
	dbType := os.Getenv("DB_TYPE")

	// Create the database factory based on the type specified in the .env file
	switch dbType {
	case "mysql":
		dbFactory = &MySQLFactory{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		}
	case "postgres":
		dbFactory = &PostgreSQLFactory{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		}
	default:
		log.Fatal("Unsupported database type specified in the .env file")
	}
	return dbFactory, nil
}
