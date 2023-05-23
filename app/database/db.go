package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

// Database struct
type Database struct {
	// Db, config can be added here
	Dbpool *pgxpool.Pool
}

// Db is the global database connection pool that is used by all handlers
var Db Database

// InitDb initializes the database connection pool and stores it in the global Db variable
func InitDb() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error(os.Stderr, "Unable to create connection pool:", err)
		os.Exit(1)
	}
	Db = Database{Dbpool: dbpool}

	var greeting string
	greeting, err = Db.GetHelloWorld()
	if err != nil {
		log.Error(os.Stderr, "InitDb failed: ", err)
		os.Exit(1)
	}
	log.Infof("InitDb succesfull: %v", greeting)
}

// CloseDbPool closes the database connection pool
func (db *Database) CloseDbPool() {
	db.Dbpool.Close()
}