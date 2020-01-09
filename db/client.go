package db

import (
	"database/sql"
	"fmt"
	"github/social-network/config"
	"github/social-network/models"

	"github.com/google/logger"
	_ "github.com/lib/pq"
)

// Datastore is an interface used to hold db methods
type Datastore interface {
	// SaveAccount is used to save the account details into DB
	SaveAccount(account *models.Account) error
	// GetAccount is used to retrieve account details from DB
	GetAccount(id, email string) ([]models.Account, error)
}

// hold db connection values
type dbClient struct {
	db *sql.DB
}

// NewClient is used to establish the connection between postgresql and our service
func NewClient() Datastore {
	var err error
	// enable database connections
	db, err := sql.Open(config.DBName, fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUserName, config.DBPassword, config.DBName))
	if err != nil {
		panic(err)
	}
	logger.Info("Successfully established DB connection")
	return &dbClient{db: db}
}
